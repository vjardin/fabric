package agent

import (
	"context"
	"crypto/tls"
	_ "embed"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	agentapi "go.githedgehog.com/fabric/api/agent/v1alpha2"
	"go.githedgehog.com/fabric/pkg/agent/gnmi"
	"go.githedgehog.com/fabric/pkg/util/uefiutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/clientcmd"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	CONF_FILE       = "agent-config.yaml"
	KUBECONFIG_FILE = "agent-kubeconfig"
)

//go:embed motd.txt
var motd []byte

type Service struct {
	Basedir string
	Version string

	DryRun          bool
	SkipControlLink bool
	ApplyOnce       bool
	SkipActions     bool

	client    *gnmi.Client
	name      string
	installID string
	runID     string
}

func (svc *Service) Run(ctx context.Context, getClient func() (*gnmi.Client, error)) error {
	if svc.Basedir == "" {
		return errors.New("basedir is required")
	}
	if svc.Version == "" {
		return errors.New("version is required")
	}

	agent, err := svc.loadConfigFromFile()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	svc.client, err = getClient()
	if err != nil {
		return errors.Wrap(err, "failed to create gNMI client")
	}
	defer svc.client.Close()

	err = svc.processAgent(ctx, agent)
	if err != nil {
		return errors.Wrap(err, "failed to process agent config from file")
	}

	if svc.DryRun {
		slog.Warn("Dry run, exiting")
		return nil
	}

	err = os.WriteFile("/etc/motd", motd, 0o644)
	if err != nil {
		slog.Warn("Failed to write motd", "err", err)
	}

	if !svc.ApplyOnce {
		err := svc.setInstallAndRunIDs()
		if err != nil {
			return errors.Wrap(err, "failed to set install and run IDs")
		}

		slog.Info("Starting watch for config changes in K8s")

		kube, err := svc.kubeClient()
		if err != nil {
			return err
		}

		currentGen := agent.Generation

		err = kube.Get(ctx, client.ObjectKey{Name: agent.Name, Namespace: "default"}, agent) // TODO ns
		if err != nil {
			return errors.Wrapf(err, "failed to get initial agent config from k8s")
		}

		// reset observability state
		now := metav1.Time{Time: time.Now()}
		agent.Status.LastHeartbeat = now
		agent.Status.LastAttemptTime = now
		agent.Status.LastAttemptGen = currentGen
		agent.Status.LastAppliedTime = now
		agent.Status.LastAppliedGen = currentGen
		agent.Status.InstallID = svc.installID
		agent.Status.RunID = svc.runID
		agent.Status.Version = svc.Version
		agent.Status.StatusUpdates = agent.Spec.StatusUpdates

		nosInfo, err := svc.client.GetNOSInfo(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get initial NOS info")
		}
		agent.Status.NOSInfo = *nosInfo

		err = kube.Status().Update(ctx, agent) // TODO maybe use patch for such status updates?
		if err != nil {
			return errors.Wrapf(err, "failed to reset agent observability status") // TODO gracefully handle case if resourceVersion changed
		}

		watcher, err := kube.Watch(context.TODO(), &agentapi.AgentList{}, client.InNamespace("default"), client.MatchingFields{ // TODO ns
			"metadata.name": svc.name,
		})
		if err != nil {
			return err
		}
		defer watcher.Stop()

		for {
			// process currently loaded agent from K8s
			err = svc.processAgentFromKube(ctx, kube, agent, &currentGen)
			if err != nil {
				return errors.Wrap(err, "failed to process agent config from k8s")
			}

			select {
			case <-ctx.Done():
				slog.Info("Context done, exiting")
				return nil
			case <-time.After(30 * time.Second):
				slog.Debug("Sending heartbeat")

				nosInfo, err := svc.client.GetNOSInfo(ctx)
				if err != nil {
					return errors.Wrapf(err, "failed to get heartbeat NOS info")
				}
				agent.Status.NOSInfo = *nosInfo
				agent.Status.LastHeartbeat = metav1.Time{Time: time.Now()}

				err = kube.Status().Update(ctx, agent)
				if err != nil {
					return errors.Wrapf(err, "failed to update agent heartbeat") // TODO gracefully handle case if resourceVersion changed
				}
			case event := <-watcher.ResultChan():
				// TODO handle bookmarks and delete events
				if event.Type == watch.Deleted || event.Type == watch.Bookmark {
					slog.Info("Received watch event, ignoring", "event", event.Type)
					continue
				}
				if event.Type == watch.Error {
					slog.Error("Received watch error", "event", event.Type, "object", event.Object)
					if err, ok := event.Object.(error); ok {
						return errors.Wrapf(err, "watch error")
					}

					return errors.New("watch error")
				}

				agent = event.Object.(*agentapi.Agent)
			}
		}
	}

	return nil
}

func (svc *Service) setInstallAndRunIDs() error {
	svc.runID = uuid.New().String()

	installIDFile := filepath.Join(svc.Basedir, "install-id")
	installID, err := os.ReadFile(installIDFile)
	if os.IsNotExist(err) {
		newInstallID := uuid.New().String()
		err = os.WriteFile(installIDFile, []byte(newInstallID), 0o640)
		if err != nil {
			return errors.Wrapf(err, "failed to write install ID file %q", installIDFile)
		}
	} else if err != nil {
		return errors.Wrapf(err, "failed to read install ID file %q", installIDFile)
	} else {
		svc.installID = string(installID)
	}

	slog.Info("IDs ready", "install", svc.installID, "run", svc.runID)
	return nil
}

func (svc *Service) processAgent(ctx context.Context, agent *agentapi.Agent) error {
	slog.Info("Processing agent config", "name", agent.Name, "gen", agent.Generation, "res", agent.ResourceVersion)

	plan, err := PreparePlan(agent)
	if err != nil {
		return errors.Wrap(err, "failed to process config to prepare plan")
	}

	_, err = plan.Entries()
	if err != nil {
		return errors.Wrap(err, "failed to generate plan entries")
	}
	slog.Debug("Plan entries generated")

	if svc.DryRun {
		return nil
	}

	if !svc.SkipControlLink {
		if err := svc.ensureControlLink(agent); err != nil {
			return errors.Wrap(err, "failed to ensure control link")
		}
		slog.Info("Control link configuration applied")
	} else {
		slog.Info("Control link configuration is skipped")
	}

	err = plan.ApplyWith(ctx, svc.client)
	if err != nil {
		return errors.Wrap(err, "failed to apply config")
	}
	slog.Info("Config applied", "name", agent.Name, "gen", agent.Generation, "res", agent.ResourceVersion)

	return nil
}

func (svc *Service) processAgentFromKube(ctx context.Context, kube client.Client, agent *agentapi.Agent, currentGen *int64) error {
	if agent.Generation == *currentGen {
		return nil
	}

	slog.Info("Agent config changed", "current", currentGen, "new", agent.Generation)

	// demonstrating that we're going to try to apply config
	agent.Status.LastAttemptGen = agent.Generation
	agent.Status.LastAttemptTime = metav1.Time{Time: time.Now()}

	err := kube.Status().Update(ctx, agent)
	if err != nil {
		return err // TODO gracefully handle case if resourceVersion changed
	}

	err = svc.processAgent(ctx, agent)
	if err != nil {
		return errors.Wrap(err, "failed to process agent config loaded from k8s")
	}

	err = svc.saveConfigToFile(agent)
	if err != nil {
		return errors.Wrap(err, "failed to save agent config to file")
	}

	// report that we've been able to apply config
	agent.Status.LastAppliedGen = agent.Generation
	agent.Status.LastAppliedTime = metav1.Time{Time: time.Now()}

	err = kube.Status().Update(ctx, agent)
	if err != nil {
		return err // TODO gracefully handle case if resourceVersion changed
	}

	err = svc.processActions(ctx, agent)
	if err != nil {
		return errors.Wrap(err, "failed to process agent actions from k8s")
	}

	*currentGen = agent.Generation

	return nil
}

func (svc *Service) processActions(ctx context.Context, agent *agentapi.Agent) error {
	reboot := false
	if agent.Spec.Reinstall != "" && agent.Spec.Reinstall == svc.installID {
		slog.Info("Reinstall requested", "installID", agent.Spec.Reinstall)
		if !svc.SkipActions {
			err := uefiutil.MakeONIEDefaultBootEntryAndCleanup()
			if err != nil {
				slog.Warn("Failed to make ONIE default boot entry", "err", err)
			} else {
				slog.Info("Rebooting into ONIE")
				reboot = true
			}
		}
	}

	if agent.Spec.Reboot != "" && agent.Spec.Reboot == svc.runID {
		slog.Info("Reboot requested", "runID", agent.Spec.Reboot)
		if !svc.SkipActions {
			slog.Info("Rebooting")
			reboot = true
		}
	}

	if reboot {
		cmd := exec.CommandContext(ctx, "wall", "Hedgehog Agent initiated reboot")
		err := cmd.Run()
		if err != nil {
			slog.Warn("Failed to send wall message", "err", err)
		}

		cmd = exec.CommandContext(ctx, "reboot")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return errors.Wrap(err, "failed to reboot")
		}
	}

	desiredVersion := ""
	if agent.Spec.Version.Default != "" {
		desiredVersion = agent.Spec.Version.Default
	}
	if agent.Spec.Version.Override != "" {
		desiredVersion = agent.Spec.Version.Override
	}
	if desiredVersion != "" && svc.Version != desiredVersion {
		slog.Info("Desired version is different from current", "desired", desiredVersion, "current", svc.Version)
		if !svc.SkipActions {
			slog.Info("Attempting to upgrade Agent")

			err := svc.agentUpgrade(ctx, agent, desiredVersion)
			if err != nil {
				slog.Warn("Failed to upgrade Agent", "err", err)
			} else {
				slog.Info("Agent upgraded")
				os.Exit(0) // TODO graceful agent restart
			}
		}
	}

	return nil
}

func (svc *Service) agentUpgrade(ctx context.Context, agent *agentapi.Agent, desiredVersion string) error {
	path, err := os.MkdirTemp("/tmp", "agent-upgrade-*")
	if err != nil {
		return errors.Wrap(err, "failed to create temp dir")
	}
	defer os.RemoveAll(path)

	fs, err := file.New(path)
	if err != nil {
		return errors.Wrapf(err, "error creating oras file store in %s", path)
	}
	defer fs.Close()

	repo, err := remote.NewRepository(agent.Spec.Version.Repo)
	if err != nil {
		return errors.Wrapf(err, "error creating oras remote repo %s", agent.Spec.Version.Repo)
	}

	baseTransport := http.DefaultTransport.(*http.Transport).Clone()
	baseTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// TODO load CA
	// config.RootCAs, err = crypto.LoadCertPool(opts.CACertFilePath)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	repo.Client = &auth.Client{
		Client: &http.Client{
			Transport: baseTransport,
		},
	}

	_, err = oras.Copy(context.Background(), repo, desiredVersion, fs, desiredVersion, oras.CopyOptions{
		CopyGraphOptions: oras.CopyGraphOptions{
			Concurrency: 2,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "error downloading new agent %s from %s", desiredVersion, agent.Spec.Version.Repo)
	}

	agentPath := filepath.Join(path, "agent")

	err = os.Chmod(agentPath, 0o755)
	if err != nil {
		return errors.Wrapf(err, "failed to chmod new agent binary in %s", path)
	}

	cmd := exec.CommandContext(ctx, agentPath, "apply", "--dry-run=true")
	err = cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "failed to run new agent binary in %s", path)
	}

	err = os.Rename(agentPath, "/opt/hedgehog/bin/agent")
	if err != nil {
		return errors.Wrapf(err, "failed to move new agent binary from %s to /opt/hedgehog/bin/agent", path)
	}

	return nil
}

func (svc *Service) kubeClient() (client.WithWatch, error) {
	kubeconfigPath := filepath.Join(svc.Basedir, KUBECONFIG_FILE)
	cfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		nil,
	).ClientConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load kubeconfig from %s", kubeconfigPath)
	}

	scheme := runtime.NewScheme()
	err = agentapi.AddToScheme(scheme)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add agent scheme")
	}

	kubeClient, err := client.NewWithWatch(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kube client")
	}

	return kubeClient, nil
}

func (s *Service) configFilePath() string {
	return filepath.Join(s.Basedir, CONF_FILE)
}

func (s *Service) loadConfigFromFile() (*agentapi.Agent, error) {
	data, err := os.ReadFile(s.configFilePath())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read config file %s", s.configFilePath())
	}

	config := &agentapi.Agent{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal config file %s", s.configFilePath())
	}
	s.name = config.Name

	return config, nil
}

func (s *Service) saveConfigToFile(agent *agentapi.Agent) error {
	if agent == nil {
		return errors.New("no config to save")
	}

	data, err := yaml.Marshal(agent)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal config")
	}

	err = os.WriteFile(s.configFilePath(), data, 0o640)
	if err != nil {
		return errors.Wrapf(err, "failed to write config file %s", s.configFilePath())
	}

	return nil
}
