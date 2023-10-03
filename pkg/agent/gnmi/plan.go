package gnmi

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"go.githedgehog.com/fabric/pkg/agent/gnmi/bcom/oc"
	"go.githedgehog.com/fabric/pkg/util/iputil"
)

type Plan struct {
	Hostname        string
	ManagementIface string
	ManagementIP    string
	MCLAGDomain     MCLAGDomain
	PortChannels    []PortChannel
	InterfaceIPs    []InterfaceIP
	Users           []User
	VPCs            []VPC
	PortGroupSpeeds map[string]string
	NAT             NAT
	StaticNAT       map[string]string
}

type PortChannel struct {
	ID             uint16 // 1..256
	Description    string
	TrunkVLANRange *string
	Members        []string // Interfaces
}

func PortChannelName(id uint16) string {
	return fmt.Sprintf("PortChannel%d", id)
}

func (pChan *PortChannel) Name() string {
	return PortChannelName(pChan.ID)
}

type InterfaceIP struct {
	Name string
	IP   string
}

type MCLAGDomain struct {
	ID       uint32 // 1..4095
	SourceIP string
	PeerIP   string
	PeerLink string
	Members  []string // PortChannels
	// MCLAGSystemMac string // TODO evaluate if we need it
}

type User struct {
	Name     string
	Password string
	Role     string
	SSHKey   string
}

type VPC struct {
	Name       string
	Subnet     string
	VLAN       uint16
	DHCP       bool
	DHCPRelay  string
	DHCPSource string
	Peers      []string
	SNAT       bool
}

type NAT struct {
	PublicIface string
	PublicIP    string
	AnchorIP    string
	Pool        []string
	Neighbor    string
	RemoteAS    uint32
}

const (
	VRF_PREFIX     = "Vrf"
	VRF_PREFIX_VPC = "V"
	VRF_NAT        = "Ndefault"

	ASN                 uint32 = 65101
	ANCHOR_VLAN         uint16 = 500
	NAT_ZONE_PUBLIC            = 1
	NAT_ZONE_OTHER             = 0
	NAT_INSTANCE_ID            = 0
	NAT_ACL_POOL_PREFIX        = "public"
)

func (plan *Plan) Entries() ([]*Entry, []*Entry, error) {
	earlyApply := []*Entry{}

	earlyApply = append(earlyApply, EntDisableZtp())
	earlyApply = append(earlyApply, EntHostname(plan.Hostname))

	for _, user := range plan.Users {
		earlyApply = append(earlyApply, EntUser(user.Name, user.Password, user.Role, user.SSHKey))
	}

	readyApply := []*Entry{}
	{
		ip, ipNet, err := net.ParseCIDR(plan.ManagementIP)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse management ip %s", plan.ManagementIP)
		}
		prefixLen, _ := ipNet.Mask.Size()

		readyApply = append(readyApply, EntInterfaceIP(plan.ManagementIface, ip.String(), uint8(prefixLen)))
	}

	for group, speedStr := range plan.PortGroupSpeeds {
		speed := oc.OpenconfigIfEthernet_ETHERNET_SPEED_UNSET

		for id, enum := range oc.ΛEnum["E_OpenconfigIfEthernet_ETHERNET_SPEED"] {
			if enum.Name == speedStr {
				speed = oc.E_OpenconfigIfEthernet_ETHERNET_SPEED(id)
				break
			}
		}

		// TODO add speed validation to the API
		if speed == oc.OpenconfigIfEthernet_ETHERNET_SPEED_UNSET || speed == oc.OpenconfigIfEthernet_ETHERNET_SPEED_SPEED_UNKNOWN {
			return nil, nil, errors.Errorf("unset or unknown speed %s for portgroup %s", speedStr, group)
		}

		// TODO add some good validation and probably different formats like w/o SPEED_ prefix and show options in error

		readyApply = append(readyApply, EntPortGroupSpeed(group, speedStr, speed))
	}

	for _, pChan := range plan.PortChannels {
		if pChan.TrunkVLANRange != nil {
			readyApply = append(readyApply, EntPortChannel(pChan.Name(), pChan.Description, *pChan.TrunkVLANRange))
		} else {
			readyApply = append(readyApply, EntL3PortChannel(pChan.Name(), pChan.Description))
		}

		for _, member := range pChan.Members {
			readyApply = append(readyApply, EntPortChannelMember(pChan.Name(), member))
			readyApply = append(readyApply, EntInterfaceUP(member))
		}
	}

	for _, ifIP := range plan.InterfaceIPs {
		ip, ipNet, err := net.ParseCIDR(ifIP.IP)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse CIDR %s for %s", ifIP.IP, ifIP.Name)
		}
		prefixLen, _ := ipNet.Mask.Size()

		readyApply = append(readyApply, EntInterfaceIP(ifIP.Name, ip.String(), uint8(prefixLen)))
	}

	readyApply = append(readyApply, EntMCLAGDomain(plan.MCLAGDomain.ID, plan.MCLAGDomain.SourceIP, plan.MCLAGDomain.PeerIP, plan.MCLAGDomain.PeerLink))

	for _, member := range plan.MCLAGDomain.Members {
		readyApply = append(readyApply, EntMCLAGMember(plan.MCLAGDomain.ID, member))
	}

	// TODO per Vrf policy
	policyName := "vpc-no-advertise"
	readyApply = append(readyApply, EntBGPRoutingPolicy(policyName,
		[]oc.OpenconfigRoutingPolicy_RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config_Communities_Union{
			oc.OpenconfigBgpTypes_BGP_WELL_KNOWN_STD_COMMUNITY_NO_ADVERTISE,
		},
	))
	for _, vpc := range plan.VPCs {
		vrfName := VRF_PREFIX + VRF_PREFIX_VPC + vpc.Name
		// policyName := vrfName + "_route_map"

		readyApply = append(readyApply, EntVrf(vrfName))
		readyApply = append(readyApply, EntVrfBGP(vrfName, ASN, []string{}, "", 0))

		readyApply = append(readyApply, EntBGPRouteDistribution(vrfName, policyName))

		cidr, err := iputil.ParseCIDR(vpc.Subnet)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse subnet %s for vpc %s", vpc.Subnet, vpc.Name)
		}
		prefixLen, _ := cidr.Subnet.Mask.Size()

		readyApply = append(readyApply, EntVLANInterface(vpc.VLAN, vpc.Name))
		readyApply = append(readyApply, EntVLANVrfMember(vrfName, vpc.VLAN))
		readyApply = append(readyApply, EntVLANInterfaceConf(vpc.VLAN, cidr.Gateway.String(), uint8(prefixLen)))

		if vpc.DHCP {
			ip, _, err := net.ParseCIDR(vpc.DHCPRelay)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "failed to parse DHCP relay %s for vpc %s", vpc.DHCPRelay, vpc.Name)
			}
			readyApply = append(readyApply, EntDHCPRelay(vpc.VLAN, ip.String(), vpc.DHCPSource))
		}
	}

	for _, vpc := range plan.VPCs {
		vrfName := VRF_PREFIX + VRF_PREFIX_VPC + vpc.Name
		peers := []string{}
		for _, peer := range vpc.Peers {
			peers = append(peers, VRF_PREFIX+VRF_PREFIX_VPC+peer)
		}
		if len(peers) > 0 { // TODO what about case when we removing all peers?
			readyApply = append(readyApply, EntVrfImportRoutes(vrfName, peers))
		}
	}

	nat := len(plan.StaticNAT) > 0
	for _, vpc := range plan.VPCs {
		if vpc.SNAT {
			nat = true
			break
		}
	}

	if plan.NAT.PublicIface != "" && nat {
		readyApply = append(readyApply, EntVrf(VRF_PREFIX+VRF_NAT))

		publicIP, err := iputil.ParseCIDR(plan.NAT.PublicIP)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse nat public ip %s", plan.NAT.PublicIP)
		}
		publicPrefixLen, _ := publicIP.Subnet.Mask.Size()

		readyApply = append(readyApply, EntVrfMember(VRF_PREFIX+VRF_NAT, plan.NAT.PublicIface))
		readyApply = append(readyApply, EntInterfaceIP(plan.NAT.PublicIface, publicIP.IP.String(), uint8(publicPrefixLen)))
		readyApply = append(readyApply, EntInterfaceNATZone(plan.NAT.PublicIface, NAT_ZONE_PUBLIC))
		readyApply = append(readyApply, EntInterfaceUP(plan.NAT.PublicIface))

		anchorIP, err := iputil.ParseCIDR(plan.NAT.AnchorIP)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse nat anchor ip %s", plan.NAT.AnchorIP)
		}
		anchorPrefixLen, _ := anchorIP.Subnet.Mask.Size()

		anchorIface := fmt.Sprintf("Vlan%d", ANCHOR_VLAN)
		readyApply = append(readyApply, EntVLANInterface(ANCHOR_VLAN, "NAT Anchor Interface"))
		readyApply = append(readyApply, EntVrfMember(VRF_PREFIX+VRF_NAT, anchorIface))
		readyApply = append(readyApply, EntVLANInterfaceConf(ANCHOR_VLAN, anchorIP.IP.String(), uint8(anchorPrefixLen)))
		readyApply = append(readyApply, EntInterfaceNATZone(anchorIface, NAT_ZONE_PUBLIC))

		pool := []string{}
		for idx, cidr := range plan.NAT.Pool {
			first, last, err := iputil.Range(cidr)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "failed to parse nat pool cidr #%d %s", idx, cidr)
			}
			pool = append(pool, fmt.Sprintf("%s-%s", first, last))
		}

		readyApply = append(readyApply, EntNATInstance(NAT_INSTANCE_ID, NAT_ZONE_PUBLIC, NAT_ACL_POOL_PREFIX, pool))

		// Temp hack to avoid NAT on out NAT Public IP
		readyApply = append(readyApply, EntStaticNAT(NAT_INSTANCE_ID, publicIP.IP.String(), publicIP.IP.String()))

		readyApply = append(readyApply, EntVrfBGP(VRF_PREFIX+VRF_NAT, ASN, plan.NAT.Pool, plan.NAT.Neighbor, plan.NAT.RemoteAS))
		readyApply = append(readyApply, EntBGPRouteDistribution(VRF_PREFIX+VRF_NAT, ""))

		// Temp hack to avoid NAT on MCLAG Session IPs
		sourceIP, _, err := net.ParseCIDR(plan.MCLAGDomain.SourceIP)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse mclag source ip %s", plan.MCLAGDomain.SourceIP)
		}
		peerIP, _, err := net.ParseCIDR(plan.MCLAGDomain.PeerIP)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to parse mclag peer ip %s", plan.MCLAGDomain.PeerIP)
		}
		readyApply = append(readyApply, EntStaticNAT(NAT_INSTANCE_ID, sourceIP.String(), sourceIP.String()))
		readyApply = append(readyApply, EntStaticNAT(NAT_INSTANCE_ID, peerIP.String(), peerIP.String()))

		for _, vpc := range plan.VPCs {
			readyApply = append(readyApply, EntInterfaceNATZone(fmt.Sprintf("Vlan%d", vpc.VLAN), NAT_ZONE_OTHER))

			if vpc.SNAT {
				// readyApply = append(readyApply, EntVrfImportRoutes(VRF_PREFIX+VRF_NAT, []string{VRF_PREFIX + VRF_PREFIX_VPC + vpc.Name}))
				// readyApply = append(readyApply, EntVrfImportRoutes(VRF_PREFIX+VRF_PREFIX_VPC+vpc.Name, []string{VRF_PREFIX + VRF_NAT}))
				readyApply = append(readyApply, EntVrfImportRoutes(VRF_PREFIX+VRF_NAT, []string{VRF_PREFIX + VRF_PREFIX_VPC + vpc.Name}))
				readyApply = append(readyApply, EntVrfImportRoutes(VRF_PREFIX+VRF_PREFIX_VPC+vpc.Name, []string{VRF_PREFIX + VRF_NAT}))
			}
		}

		for privateIP, externalIP := range plan.StaticNAT {
			readyApply = append(readyApply, EntStaticNAT(NAT_INSTANCE_ID, privateIP, externalIP))
			readyApply = append(readyApply, EntVrfBGP(VRF_PREFIX+VRF_NAT, ASN, []string{externalIP + "/32"}, "", plan.NAT.RemoteAS))
		}
	}

	return earlyApply, readyApply, nil
}

func subnetsToRanges(subnets []string) ([]string, error) {
	ranges := []string{}
	for _, subnet := range subnets {
		first, last, err := iputil.Range(subnet)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse subnet %s", subnet)
		}

		if first == last {
			ranges = append(ranges, first)
		} else {
			ranges = append(ranges, fmt.Sprintf("%s-%s", first, last))
		}
	}

	return ranges, nil
}
