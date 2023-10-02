package sample

import (
	"fmt"

	"github.com/pkg/errors"
	vpcapi "go.githedgehog.com/fabric/api/vpc/v1alpha2"
	wiringapi "go.githedgehog.com/fabric/api/wiring/v1alpha2"
	"go.githedgehog.com/fabric/pkg/wiring"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Preset string

const (
	SAMPLE_CC_VLAB Preset = "vlab"
	SAMPLE_CC_LAB  Preset = "lab"
)

var PresetsAll = []Preset{
	SAMPLE_CC_VLAB,
	SAMPLE_CC_LAB,
}

func CollapsedCore(preset Preset) (*wiring.Data, error) {
	if preset == "" {
		preset = SAMPLE_CC_VLAB
	}

	oniePort := "eth0" // we're using mgmt port for now

	data, err := wiring.New()
	if err != nil {
		return nil, err
	}

	_, err = createRack(data, "rack-1", wiringapi.RackSpec{})
	if err != nil {
		return nil, err
	}

	_, err = createNAT(data, "default", vpcapi.NATSpec{
		DNAT: vpcapi.DNAT{
			Pool: []string{
				"192.168.91.192/26", // 192.168.91.193 - 192.168.91.254
			},
		},
	})
	if err != nil {
		return nil, err
	}

	switch1 := wiringapi.SwitchSpec{
		Location: location("1"),
	}
	if preset == SAMPLE_CC_LAB {
		switch1.PortGroupSpeeds = map[string]string{
			"1":  "SPEED_10GB",
			"12": "SPEED_10GB",
		}
	}
	_, err = createSwitch(data, "switch-1", "rack-1", switch1)
	if err != nil {
		return nil, err
	}

	switch2 := wiringapi.SwitchSpec{
		Location: location("2"),
	}
	if preset == SAMPLE_CC_LAB {
		switch2.PortGroupSpeeds = map[string]string{
			"1":  "SPEED_10GB",
			"12": "SPEED_10GB",
		}
	}
	_, err = createSwitch(data, "switch-2", "rack-1", switch2)
	if err != nil {
		return nil, err
	}

	_, err = createServer(data, "control-1", "rack-1", wiringapi.ServerSpec{
		Type: wiringapi.ServerTypeControl,
	})
	if err != nil {
		return nil, err
	}

	_, err = createServer(data, "server-1", "rack-1", wiringapi.ServerSpec{})
	if err != nil {
		return nil, err
	}
	_, err = createServer(data, "server-2", "rack-1", wiringapi.ServerSpec{})
	if err != nil {
		return nil, err
	}

	if preset == SAMPLE_CC_VLAB {
		_, err = createServer(data, "server-3", "rack-1", wiringapi.ServerSpec{})
		if err != nil {
			return nil, err
		}
		_, err = createServer(data, "server-4", "rack-1", wiringapi.ServerSpec{})
		if err != nil {
			return nil, err
		}
	} else if preset == SAMPLE_CC_LAB {
		_, err = createServer(data, "server-3", "rack-1", wiringapi.ServerSpec{})
		if err != nil {
			return nil, err
		}
	}

	natPort := "Ethernet0"
	if preset == SAMPLE_CC_LAB {
		natPort = "Ethernet0"
	}

	// nat connection switch-1
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		NAT: &wiringapi.ConnNAT{
			Link: wiringapi.ConnNATLink{
				Switch: wiringapi.ConnNATLinkSwitch{
					BasePortName: wiringapi.NewBasePortName("switch-1/" + natPort),
					IP:           "192.168.91.0/31",
					NeighborIP:   "192.168.91.1",
					RemoteAS:     65102,
					AnchorIP:     "192.168.91.129/27",
					SNAT: wiringapi.SNAT{
						Pool: []string{
							"192.168.91.128/27",
						},
					},
				},
				NAT: wiringapi.NewBasePortName("default"),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// nat connection switch-2
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		NAT: &wiringapi.ConnNAT{
			Link: wiringapi.ConnNATLink{
				Switch: wiringapi.ConnNATLinkSwitch{
					BasePortName: wiringapi.NewBasePortName("switch-2/" + natPort),
					IP:           "192.168.91.65/31",
					NeighborIP:   "192.168.91.64",
					RemoteAS:     65102,
					AnchorIP:     "192.168.91.97/27",
					SNAT: wiringapi.SNAT{
						Pool: []string{
							"192.168.91.96/27",
						},
					},
				},
				NAT: wiringapi.NewBasePortName("default"),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	ctrlSwitchPort := func(portID int) string {
		if preset == SAMPLE_CC_VLAB {
			// return fmt.Sprintf("eth%d", portID)
			return fmt.Sprintf("enp0s%d", portID+2)
		}
		if preset == SAMPLE_CC_LAB {
			return fmt.Sprintf("eno%d", portID+1)
		}

		return "<invalid>"
	}

	// control-1 <> switch-1
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		Management: &wiringapi.ConnMgmt{
			Link: wiringapi.ConnMgmtLink{
				Server: wiringapi.ConnMgmtLinkServer{
					BasePortName: wiringapi.NewBasePortName("control-1/" + ctrlSwitchPort(1)),
					IP:           "192.168.101.1/31",
				},
				Switch: wiringapi.ConnMgmtLinkSwitch{
					BasePortName: wiringapi.NewBasePortName("switch-1/Management0"),
					IP:           "192.168.101.0/31",
					ONIEPortName: oniePort,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// control-1 <> switch-2
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		Management: &wiringapi.ConnMgmt{
			Link: wiringapi.ConnMgmtLink{
				Server: wiringapi.ConnMgmtLinkServer{
					BasePortName: wiringapi.NewBasePortName("control-1/" + ctrlSwitchPort(2)),
					IP:           "192.168.102.1/31",
				},
				Switch: wiringapi.ConnMgmtLinkSwitch{
					BasePortName: wiringapi.NewBasePortName("switch-2/Management0"),
					IP:           "192.168.102.0/31",
					ONIEPortName: oniePort,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	mclagPeerPort1 := "Ethernet1"
	mclagPeerPort2 := "Ethernet2"
	mclagSessionPort1 := "Ethernet3"
	mclagSessionPort2 := "Ethernet4"
	if preset == SAMPLE_CC_LAB {
		mclagPeerPort1 = "Ethernet48"
		mclagPeerPort2 = "Ethernet56"
		mclagSessionPort1 = "Ethernet64"
		mclagSessionPort2 = "Ethernet68"
	}

	// MCLAG Domain peer link
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		MCLAGDomain: &wiringapi.ConnMCLAGDomain{
			PeerLinks: []wiringapi.SwitchToSwitchLink{
				{
					Switch1: wiringapi.NewBasePortName("switch-1/" + mclagPeerPort1),
					Switch2: wiringapi.NewBasePortName("switch-2/" + mclagPeerPort1),
				},
				{
					Switch1: wiringapi.NewBasePortName("switch-1/" + mclagPeerPort2),
					Switch2: wiringapi.NewBasePortName("switch-2/" + mclagPeerPort2),
				},
			},
			SessionLinks: []wiringapi.SwitchToSwitchLink{
				{
					Switch1: wiringapi.NewBasePortName("switch-1/" + mclagSessionPort1),
					Switch2: wiringapi.NewBasePortName("switch-2/" + mclagSessionPort1),
				},
				{
					Switch1: wiringapi.NewBasePortName("switch-1/" + mclagSessionPort2),
					Switch2: wiringapi.NewBasePortName("switch-2/" + mclagSessionPort2),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	server1Port1 := "enp0s2"
	server1Port2 := "enp0s3"
	server1Switch1Port := "Ethernet5"
	server1Switch2Port := "Ethernet5"
	if preset == SAMPLE_CC_LAB {
		server1Port1 = "enp7s0"
		server1Port2 = "enp8s0"
		server1Switch1Port = "Ethernet47"
		server1Switch2Port = "Ethernet46"
	}

	// server-1 <MCLAG> (switch-1, switch-2)
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		MCLAG: &wiringapi.ConnMCLAG{
			Links: []wiringapi.ServerToSwitchLink{
				{
					Server: wiringapi.NewBasePortName("server-1/" + server1Port1),
					Switch: wiringapi.NewBasePortName("switch-1/" + server1Switch1Port),
				},
				{
					Server: wiringapi.NewBasePortName("server-1/" + server1Port2),
					Switch: wiringapi.NewBasePortName("switch-2/" + server1Switch2Port),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	server2Port1 := "enp0s2"
	server2Port2 := "enp0s3"
	server2Switch1Port := "Ethernet6"
	server2Switch2Port := "Ethernet6"
	if preset == SAMPLE_CC_LAB {
		server2Port1 = "enp7s0"
		server2Port2 = "enp8s0"
		server2Switch1Port = "Ethernet46"
		server2Switch2Port = "Ethernet47"
	}

	// server-2 <MCLAG> (switch-1, switch-2)
	_, err = createConnection(data, wiringapi.ConnectionSpec{
		MCLAG: &wiringapi.ConnMCLAG{
			Links: []wiringapi.ServerToSwitchLink{
				{
					Server: wiringapi.NewBasePortName("server-2/" + server2Port1),
					Switch: wiringapi.NewBasePortName("switch-1/" + server2Switch1Port),
				},
				{
					Server: wiringapi.NewBasePortName("server-2/" + server2Port2),
					Switch: wiringapi.NewBasePortName("switch-2/" + server2Switch2Port),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if preset == SAMPLE_CC_VLAB {
		// server-3 <> switch-1
		_, err = createConnection(data, wiringapi.ConnectionSpec{
			Unbundled: &wiringapi.ConnUnbundled{
				Link: wiringapi.ServerToSwitchLink{
					Server: wiringapi.NewBasePortName("server-3/nic0/port0"),
					Switch: wiringapi.NewBasePortName("switch-1/Ethernet7"),
				},
			},
		})
		if err != nil {
			return nil, err
		}

		// server-4 <> switch-2
		_, err = createConnection(data, wiringapi.ConnectionSpec{
			Unbundled: &wiringapi.ConnUnbundled{
				Link: wiringapi.ServerToSwitchLink{
					Server: wiringapi.NewBasePortName("server-4/nic0/port0"),
					Switch: wiringapi.NewBasePortName("switch-2/Ethernet7"),
				},
			},
		})
		if err != nil {
			return nil, err
		}
	} else if preset == SAMPLE_CC_LAB {
		// server-3 <MCLAG> (switch-1, switch-2)
		_, err = createConnection(data, wiringapi.ConnectionSpec{
			MCLAG: &wiringapi.ConnMCLAG{
				Links: []wiringapi.ServerToSwitchLink{
					{
						Server: wiringapi.NewBasePortName("server-3/ens3f0"),
						Switch: wiringapi.NewBasePortName("switch-1/Ethernet4"),
					},
					{
						Server: wiringapi.NewBasePortName("server-3/ens3f1"),
						Switch: wiringapi.NewBasePortName("switch-2/Ethernet4"),
					},
				},
			},
		})
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func location(slot string) wiringapi.Location {
	return wiringapi.Location{
		Location: "LOC",
		Aisle:    "1",
		Row:      "1",
		Rack:     "1",
		Slot:     slot,
	}
}

func createRack(data *wiring.Data, name string, spec wiringapi.RackSpec) (*wiringapi.Rack, error) {
	sw := &wiringapi.Rack{
		TypeMeta: meta.TypeMeta{
			Kind:       wiringapi.KindRack,
			APIVersion: wiringapi.GroupVersion.String(),
		},
		ObjectMeta: meta.ObjectMeta{
			Name:   name,
			Labels: map[string]string{},
		},
		Spec: spec,
	}

	return sw, errors.Wrapf(data.Add(sw), "error creating switch %s", name)
}

func createNAT(data *wiring.Data, name string, spec vpcapi.NATSpec) (*vpcapi.NAT, error) {
	nat := &vpcapi.NAT{
		TypeMeta: meta.TypeMeta{
			Kind:       "NAT",
			APIVersion: vpcapi.GroupVersion.String(),
		},
		ObjectMeta: meta.ObjectMeta{
			Name:   name,
			Labels: map[string]string{},
		},
		Spec: spec,
	}

	return nat, errors.Wrapf(data.Add(nat), "error creating nat %s", name)
}

func createSwitch(data *wiring.Data, name string, rack string, spec wiringapi.SwitchSpec) (*wiringapi.Switch, error) {
	sw := &wiringapi.Switch{
		TypeMeta: meta.TypeMeta{
			Kind:       wiringapi.KindSwitch,
			APIVersion: wiringapi.GroupVersion.String(),
		},
		ObjectMeta: meta.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				wiringapi.LabelRack: rack,
			},
		},
		Spec: spec,
	}

	sw.Default()

	return sw, errors.Wrapf(data.Add(sw), "error creating switch %s", name)
}

func createServer(data *wiring.Data, name string, rack string, spec wiringapi.ServerSpec) (*wiringapi.Server, error) {
	server := &wiringapi.Server{
		TypeMeta: meta.TypeMeta{
			Kind:       wiringapi.KindServer,
			APIVersion: wiringapi.GroupVersion.String(),
		},
		ObjectMeta: meta.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				wiringapi.LabelRack: rack,
			},
		},
		Spec: spec,
	}

	server.Default()

	return server, errors.Wrapf(data.Add(server), "error creating server %s", name)
}

func createConnection(data *wiring.Data, spec wiringapi.ConnectionSpec) (*wiringapi.Connection, error) {
	name := spec.GenerateName()

	conn := &wiringapi.Connection{
		TypeMeta: meta.TypeMeta{
			Kind:       wiringapi.KindConnection,
			APIVersion: wiringapi.GroupVersion.String(),
		},
		ObjectMeta: meta.ObjectMeta{
			Name:   name,
			Labels: map[string]string{},
		},
		Spec: spec,
	}
	conn.Labels[wiringapi.ListLabelRack("rack-1")] = wiringapi.ListLabelValue

	conn.Default()

	return conn, errors.Wrapf(data.Add(conn), "error creating connection %s", name)
}
