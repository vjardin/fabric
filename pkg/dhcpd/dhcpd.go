package dhcpd

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	dhcpapi "go.githedgehog.com/fabric/api/dhcp/v1alpha2"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/pkg/errors"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(dhcpapi.AddToScheme(scheme))
}

func kubeClient() (client.WithWatch, error) {
	k8scfg, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}
	client, err := client.NewWithWatch(k8scfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

type Service struct {
	Verbose bool
	Config  string

	kubeUpdates  chan Event
	updateStatus func(dhcpapi.DHCPSubnet) error
}

type Event struct {
	Type   EventType
	Subnet *dhcpapi.DHCPSubnet
}

type EventType string

const (
	EventTypeAdded    EventType = "ADDED"
	EventTypeModified EventType = "MODIFIED"
	EventTypeDeleted  EventType = "DELETED"
)

func (d *Service) Run(ctx context.Context) error {
	kube, err := kubeClient()
	if err != nil {
		return errors.Wrap(err, "cannot create kube client")
	}

	d.kubeUpdates = make(chan Event, 100)
	d.updateStatus = func(d dhcpapi.DHCPSubnet) error {
		// TODO download latest obj and try to update its status
		return kube.Status().Update(ctx, &d)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := d.runKubeWatcher(ctx, kube); err != nil {
			slog.Error("KubeWatcher", "error", err)
		}

		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := d.runCoreDHCP(ctx); err != nil {
			slog.Error("CoreDHCP", "error", err)
		}

		time.Sleep(1 * time.Second)
		os.Exit(2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		handleExpiredLeases()
	}()
	wg.Wait()

	return nil
}

func (d *Service) runKubeWatcher(ctx context.Context, kube client.WithWatch) error {
	var err error
	var watcher watch.Interface

	for {
		if watcher == nil {
			slog.Info("Starting K8s watcher")
			if watcher, err = kube.Watch(context.TODO(), &dhcpapi.DHCPSubnetList{}, client.InNamespace("default")); err != nil { // TODO ns
				return err
			}
			defer watcher.Stop()
		}

		select {
		case <-ctx.Done():
			return nil
		case event, ok := <-watcher.ResultChan():
			if !ok {
				slog.Warn("K8s watch channel closed, restarting watcher")
				watcher = nil
				continue
			}

			if event.Object == nil {
				slog.Warn("Received nil object from K8s, restarting watcher")
				watcher = nil
				continue
			}

			if event.Type == watch.Bookmark {
				slog.Info("Received watch event, ignoring", "event", event.Type)
				continue
			}

			if event.Type == watch.Error {
				slog.Error("Received watch error", "event", event.Type, "object", event.Object)
				if err, ok := event.Object.(error); ok {
					slog.Error("Watch error", "error", err)
				}

				watcher = nil
				continue
			}

			subnet := event.Object.(*dhcpapi.DHCPSubnet)
			slog.Debug("Received watch event", "event", event.Type, "subnet", subnet.Name)
			d.kubeUpdates <- Event{
				Type:   EventType(event.Type),
				Subnet: subnet,
			}
		}
	}
}
