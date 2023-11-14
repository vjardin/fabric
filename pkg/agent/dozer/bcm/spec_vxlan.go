package bcm

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/openconfig/ygot/ygot"
	"github.com/pkg/errors"
	"go.githedgehog.com/fabric/pkg/agent/dozer"
	"go.githedgehog.com/fabric/pkg/agent/dozer/bcm/gnmi"
	"go.githedgehog.com/fabric/pkg/agent/dozer/bcm/gnmi/oc"
)

var specVXLANTunnelsEnforcer = &DefaultMapEnforcer[string, *dozer.SpecVXLANTunnel]{
	Summary:      "VXLAN tunnel",
	ValueHandler: specVXLANTunnelEnforcer,
}

var specVXLANTunnelEnforcer = &DefaultValueEnforcer[string, *dozer.SpecVXLANTunnel]{
	Summary:      "VXLAN tunnel %s",
	CreatePath:   "/sonic-vxlan/VXLAN_TUNNEL/VXLAN_TUNNEL_LIST",
	Path:         "/sonic-vxlan/VXLAN_TUNNEL/VXLAN_TUNNEL_LIST[name=%s]",
	UpdateWeight: ActionWeightVXLANTunnelUpdate,
	DeleteWeight: ActionWeightVXLANTunnelDelete,
	Marshal: func(name string, value *dozer.SpecVXLANTunnel) (ygot.ValidatedGoStruct, error) {
		return &oc.SonicVxlan_SonicVxlan_VXLAN_TUNNEL{
			VXLAN_TUNNEL_LIST: map[string]*oc.SonicVxlan_SonicVxlan_VXLAN_TUNNEL_VXLAN_TUNNEL_LIST{
				name: {
					Name:    ygot.String(name),
					SrcIp:   value.SourceIP,
					SrcIntf: value.SourceInterface,
				},
			},
		}, nil
	},
}

var specVXLANEVPNNVOsEnforcer = &DefaultMapEnforcer[string, *dozer.SpecVXLANEVPNNVO]{
	Summary:      "VXLAN EVPN NVO",
	ValueHandler: specVXLANEVPNNVOEnforcer,
}

var specVXLANEVPNNVOEnforcer = &DefaultValueEnforcer[string, *dozer.SpecVXLANEVPNNVO]{
	Summary:      "VXLAN EVPN NVO %s",
	CreatePath:   "/sonic-vxlan/VXLAN_EVPN_NVO/VXLAN_EVPN_NVO_LIST",
	Path:         "/sonic-vxlan/VXLAN_EVPN_NVO/VXLAN_EVPN_NVO_LIST[name=%s]",
	UpdateWeight: ActionWeightVXLANEVPNNVOUpdate,
	DeleteWeight: ActionWeightVXLANEVPNNVODelete,
	Marshal: func(name string, value *dozer.SpecVXLANEVPNNVO) (ygot.ValidatedGoStruct, error) {
		return &oc.SonicVxlan_SonicVxlan_VXLAN_EVPN_NVO{
			VXLAN_EVPN_NVO_LIST: map[string]*oc.SonicVxlan_SonicVxlan_VXLAN_EVPN_NVO_VXLAN_EVPN_NVO_LIST{
				name: {
					Name:       ygot.String(name),
					SourceVtep: value.SourceVTEP,
				},
			},
		}, nil
	},
}

var specVXLANTunnelMapsEnforcer = &DefaultMapEnforcer[string, *dozer.SpecVXLANTunnelMap]{
	Summary:      "VXLAN tunnel map",
	ValueHandler: specVXLANTunnelMapEnforcer,
}

var specVXLANTunnelMapEnforcer = &DefaultValueEnforcer[string, *dozer.SpecVXLANTunnelMap]{
	Summary:      "VXLAN tunnel map %s",
	CreatePath:   "/sonic-vxlan/VXLAN_TUNNEL_MAP/VXLAN_TUNNEL_MAP_LIST",
	Path:         "/sonic-vxlan/VXLAN_TUNNEL_MAP/VXLAN_TUNNEL_MAP_LIST[name=vtepfabric][mapname=%s]", // TODO unhardcode vtepfabric, but it's always only single vtep configured
	UpdateWeight: ActionWeightVXLANTunnelMapUpdate,
	DeleteWeight: ActionWeightVXLANTunnelMapDelete,
	Marshal: func(name string, value *dozer.SpecVXLANTunnelMap) (ygot.ValidatedGoStruct, error) {
		if value.VTEP == nil {
			return nil, errors.Errorf("missing VTEP")
		}
		if *value.VTEP != "vtepfabric" {
			return nil, errors.Errorf("unsupported VTEP %q", *value.VTEP)
		}
		if value.VLAN == nil {
			return nil, errors.Errorf("missing VLAN")
		}

		return &oc.SonicVxlan_SonicVxlan_VXLAN_TUNNEL_MAP{
			VXLAN_TUNNEL_MAP_LIST: map[oc.SonicVxlan_SonicVxlan_VXLAN_TUNNEL_MAP_VXLAN_TUNNEL_MAP_LIST_Key]*oc.SonicVxlan_SonicVxlan_VXLAN_TUNNEL_MAP_VXLAN_TUNNEL_MAP_LIST{
				{
					Name:    "vtepfabric", // TODO ditto
					Mapname: name,
				}: {
					Name:    ygot.String("vtepfabric"), // TODO ditto
					Mapname: ygot.String(name),
					Vlan:    ygot.String(fmt.Sprintf("Vlan%d", *value.VLAN)),
					Vni:     value.VNI,
				},
			},
		}, nil
	},
}

var specVRFVNIMapEnforcer = &DefaultMapEnforcer[string, *dozer.SpecVRFVNIEntry]{
	Summary:      "VRF VNI",
	ValueHandler: specVRFVNIEnforcer,
}

var specVRFVNIEnforcer = &DefaultValueEnforcer[string, *dozer.SpecVRFVNIEntry]{
	Summary:      "VRF VNI %s",
	Path:         "/sonic-vrf/VRF/VRF_LIST[vrf_name=%s]/vni",
	UpdateWeight: ActionWeightVRFVNIUpdate,
	DeleteWeight: ActionWeightVRFVNIDelete,
	Marshal: func(name string, value *dozer.SpecVRFVNIEntry) (ygot.ValidatedGoStruct, error) {
		return &oc.SonicVrf_SonicVrf_VRF_VRF_LIST{
			VrfName: ygot.String(name),
			Vni:     value.VNI,
		}, nil
	},
}

func loadActualVXLANs(ctx context.Context, client *gnmi.Client, spec *dozer.Spec) error {
	ocVal := &oc.Device{}
	err := client.Get(ctx, "/sonic-vxlan", ocVal)
	if err != nil {
		return errors.Wrapf(err, "failed to get vxlan")
	}

	spec.VXLANTunnels, spec.VXLANEVPNNVOs, spec.VXLANTunnelMap, err = unmarshalActualVXLANs(ocVal.SonicVxlan)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal vxlan")
	}

	return nil
}

func unmarshalActualVXLANs(ocVal *oc.SonicVxlan_SonicVxlan) (map[string]*dozer.SpecVXLANTunnel, map[string]*dozer.SpecVXLANEVPNNVO, map[string]*dozer.SpecVXLANTunnelMap, error) {
	vxlanTunnels := map[string]*dozer.SpecVXLANTunnel{}
	vxlanEvpnNvos := map[string]*dozer.SpecVXLANEVPNNVO{}
	vxlanTunnelMaps := map[string]*dozer.SpecVXLANTunnelMap{}

	if ocVal == nil {
		return vxlanTunnels, vxlanEvpnNvos, vxlanTunnelMaps, nil
	}

	if ocVal.VXLAN_TUNNEL != nil {
		for name, vxlanTunnel := range ocVal.VXLAN_TUNNEL.VXLAN_TUNNEL_LIST {
			vxlanTunnels[name] = &dozer.SpecVXLANTunnel{
				SourceIP:        vxlanTunnel.SrcIp,
				SourceInterface: vxlanTunnel.SrcIntf,
			}
		}
	}

	if ocVal.VXLAN_EVPN_NVO != nil {
		for name, vxlanEvpnNvo := range ocVal.VXLAN_EVPN_NVO.VXLAN_EVPN_NVO_LIST {
			vxlanEvpnNvos[name] = &dozer.SpecVXLANEVPNNVO{
				SourceVTEP: vxlanEvpnNvo.SourceVtep,
			}
		}
	}

	if ocVal.VXLAN_TUNNEL_MAP != nil {
		for key, vxlanTunnelMap := range ocVal.VXLAN_TUNNEL_MAP.VXLAN_TUNNEL_MAP_LIST {
			if key.Name != "vtepfabric" { // TODO ditto
				continue
			}

			var vlan *uint16
			if vxlanTunnelMap.Vlan != nil {
				value, err := strconv.ParseUint(strings.TrimPrefix(*vxlanTunnelMap.Vlan, "Vlan"), 10, 16)
				if err != nil {
					return nil, nil, nil, errors.Wrapf(err, "can't parse vlan %s", *vxlanTunnelMap.Vlan)
				}
				vlan = ygot.Uint16(uint16(value))
			}

			vxlanTunnelMaps[key.Mapname] = &dozer.SpecVXLANTunnelMap{
				VTEP: vxlanTunnelMap.Name,
				VLAN: vlan,
				VNI:  vxlanTunnelMap.Vni,
			}
		}
	}

	return vxlanTunnels, vxlanEvpnNvos, vxlanTunnelMaps, nil
}

func loadActualVRFVNIMap(ctx context.Context, client *gnmi.Client, spec *dozer.Spec) error {
	ocVRFMap := &oc.SonicVrf_SonicVrf_VRF{}
	err := client.Get(ctx, "/sonic-vrf/VRF/VRF_LIST", ocVRFMap)
	if err != nil {
		return errors.Wrapf(err, "failed to get vrfs for vni map")
	}

	spec.VRFVNIMap, err = unmarshalActualVRFVNIMap(ocVRFMap)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal vrf vni map")
	}

	return nil
}

func unmarshalActualVRFVNIMap(ocVal *oc.SonicVrf_SonicVrf_VRF) (map[string]*dozer.SpecVRFVNIEntry, error) {
	vrfVnis := map[string]*dozer.SpecVRFVNIEntry{}

	if ocVal == nil {
		return vrfVnis, nil
	}

	for name, vrf := range ocVal.VRF_LIST {
		if vrf.Vni == nil {
			continue
		}
		vrfVnis[name] = &dozer.SpecVRFVNIEntry{
			VNI: vrf.Vni,
		}
	}

	return vrfVnis, nil
}
