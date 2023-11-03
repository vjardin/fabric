//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Hedgehog.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyStatus) DeepCopyInto(out *ApplyStatus) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	if in.Detailed != nil {
		in, out := &in.Detailed, &out.Detailed
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyStatus.
func (in *ApplyStatus) DeepCopy() *ApplyStatus {
	if in == nil {
		return nil
	}
	out := new(ApplyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasePortName) DeepCopyInto(out *BasePortName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasePortName.
func (in *BasePortName) DeepCopy() *BasePortName {
	if in == nil {
		return nil
	}
	out := new(BasePortName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnFabric) DeepCopyInto(out *ConnFabric) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]FabricLink, len(*in))
		copy(*out, *in)
	}
	out.Spine = in.Spine
	out.Leaf = in.Leaf
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnFabric.
func (in *ConnFabric) DeepCopy() *ConnFabric {
	if in == nil {
		return nil
	}
	out := new(ConnFabric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnFabricLeaf) DeepCopyInto(out *ConnFabricLeaf) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnFabricLeaf.
func (in *ConnFabricLeaf) DeepCopy() *ConnFabricLeaf {
	if in == nil {
		return nil
	}
	out := new(ConnFabricLeaf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnFabricLinkSwitch) DeepCopyInto(out *ConnFabricLinkSwitch) {
	*out = *in
	out.BasePortName = in.BasePortName
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnFabricLinkSwitch.
func (in *ConnFabricLinkSwitch) DeepCopy() *ConnFabricLinkSwitch {
	if in == nil {
		return nil
	}
	out := new(ConnFabricLinkSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnFabricSpine) DeepCopyInto(out *ConnFabricSpine) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnFabricSpine.
func (in *ConnFabricSpine) DeepCopy() *ConnFabricSpine {
	if in == nil {
		return nil
	}
	out := new(ConnFabricSpine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnMCLAG) DeepCopyInto(out *ConnMCLAG) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]ServerToSwitchLink, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnMCLAG.
func (in *ConnMCLAG) DeepCopy() *ConnMCLAG {
	if in == nil {
		return nil
	}
	out := new(ConnMCLAG)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnMCLAGDomain) DeepCopyInto(out *ConnMCLAGDomain) {
	*out = *in
	if in.PeerLinks != nil {
		in, out := &in.PeerLinks, &out.PeerLinks
		*out = make([]SwitchToSwitchLink, len(*in))
		copy(*out, *in)
	}
	if in.SessionLinks != nil {
		in, out := &in.SessionLinks, &out.SessionLinks
		*out = make([]SwitchToSwitchLink, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnMCLAGDomain.
func (in *ConnMCLAGDomain) DeepCopy() *ConnMCLAGDomain {
	if in == nil {
		return nil
	}
	out := new(ConnMCLAGDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnMgmt) DeepCopyInto(out *ConnMgmt) {
	*out = *in
	out.Link = in.Link
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnMgmt.
func (in *ConnMgmt) DeepCopy() *ConnMgmt {
	if in == nil {
		return nil
	}
	out := new(ConnMgmt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnMgmtLink) DeepCopyInto(out *ConnMgmtLink) {
	*out = *in
	out.Server = in.Server
	out.Switch = in.Switch
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnMgmtLink.
func (in *ConnMgmtLink) DeepCopy() *ConnMgmtLink {
	if in == nil {
		return nil
	}
	out := new(ConnMgmtLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnMgmtLinkServer) DeepCopyInto(out *ConnMgmtLinkServer) {
	*out = *in
	out.BasePortName = in.BasePortName
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnMgmtLinkServer.
func (in *ConnMgmtLinkServer) DeepCopy() *ConnMgmtLinkServer {
	if in == nil {
		return nil
	}
	out := new(ConnMgmtLinkServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnMgmtLinkSwitch) DeepCopyInto(out *ConnMgmtLinkSwitch) {
	*out = *in
	out.BasePortName = in.BasePortName
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnMgmtLinkSwitch.
func (in *ConnMgmtLinkSwitch) DeepCopy() *ConnMgmtLinkSwitch {
	if in == nil {
		return nil
	}
	out := new(ConnMgmtLinkSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnNAT) DeepCopyInto(out *ConnNAT) {
	*out = *in
	in.Link.DeepCopyInto(&out.Link)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnNAT.
func (in *ConnNAT) DeepCopy() *ConnNAT {
	if in == nil {
		return nil
	}
	out := new(ConnNAT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnNATLink) DeepCopyInto(out *ConnNATLink) {
	*out = *in
	in.Switch.DeepCopyInto(&out.Switch)
	out.NAT = in.NAT
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnNATLink.
func (in *ConnNATLink) DeepCopy() *ConnNATLink {
	if in == nil {
		return nil
	}
	out := new(ConnNATLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnNATLinkSwitch) DeepCopyInto(out *ConnNATLinkSwitch) {
	*out = *in
	out.BasePortName = in.BasePortName
	in.SNAT.DeepCopyInto(&out.SNAT)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnNATLinkSwitch.
func (in *ConnNATLinkSwitch) DeepCopy() *ConnNATLinkSwitch {
	if in == nil {
		return nil
	}
	out := new(ConnNATLinkSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnUnbundled) DeepCopyInto(out *ConnUnbundled) {
	*out = *in
	out.Link = in.Link
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnUnbundled.
func (in *ConnUnbundled) DeepCopy() *ConnUnbundled {
	if in == nil {
		return nil
	}
	out := new(ConnUnbundled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnVPCLoopback) DeepCopyInto(out *ConnVPCLoopback) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]SwitchToSwitchLink, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnVPCLoopback.
func (in *ConnVPCLoopback) DeepCopy() *ConnVPCLoopback {
	if in == nil {
		return nil
	}
	out := new(ConnVPCLoopback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Connection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionList) DeepCopyInto(out *ConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Connection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionList.
func (in *ConnectionList) DeepCopy() *ConnectionList {
	if in == nil {
		return nil
	}
	out := new(ConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpec) DeepCopyInto(out *ConnectionSpec) {
	*out = *in
	if in.Unbundled != nil {
		in, out := &in.Unbundled, &out.Unbundled
		*out = new(ConnUnbundled)
		**out = **in
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(ConnMgmt)
		**out = **in
	}
	if in.MCLAG != nil {
		in, out := &in.MCLAG, &out.MCLAG
		*out = new(ConnMCLAG)
		(*in).DeepCopyInto(*out)
	}
	if in.MCLAGDomain != nil {
		in, out := &in.MCLAGDomain, &out.MCLAGDomain
		*out = new(ConnMCLAGDomain)
		(*in).DeepCopyInto(*out)
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(ConnNAT)
		(*in).DeepCopyInto(*out)
	}
	if in.Fabric != nil {
		in, out := &in.Fabric, &out.Fabric
		*out = new(ConnFabric)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCLoopback != nil {
		in, out := &in.VPCLoopback, &out.VPCLoopback
		*out = new(ConnVPCLoopback)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpec.
func (in *ConnectionSpec) DeepCopy() *ConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStatus) DeepCopyInto(out *ConnectionStatus) {
	*out = *in
	in.Applied.DeepCopyInto(&out.Applied)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStatus.
func (in *ConnectionStatus) DeepCopy() *ConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricLink) DeepCopyInto(out *FabricLink) {
	*out = *in
	out.Spine = in.Spine
	out.Leaf = in.Leaf
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricLink.
func (in *FabricLink) DeepCopy() *FabricLink {
	if in == nil {
		return nil
	}
	out := new(FabricLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLDPConfig) DeepCopyInto(out *LLDPConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLDPConfig.
func (in *LLDPConfig) DeepCopy() *LLDPConfig {
	if in == nil {
		return nil
	}
	out := new(LLDPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Location) DeepCopyInto(out *Location) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Location.
func (in *Location) DeepCopy() *Location {
	if in == nil {
		return nil
	}
	out := new(Location)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationSig) DeepCopyInto(out *LocationSig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationSig.
func (in *LocationSig) DeepCopy() *LocationSig {
	if in == nil {
		return nil
	}
	out := new(LocationSig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rack.
func (in *Rack) DeepCopy() *Rack {
	if in == nil {
		return nil
	}
	out := new(Rack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackList) DeepCopyInto(out *RackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackList.
func (in *RackList) DeepCopy() *RackList {
	if in == nil {
		return nil
	}
	out := new(RackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackPosition) DeepCopyInto(out *RackPosition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackPosition.
func (in *RackPosition) DeepCopy() *RackPosition {
	if in == nil {
		return nil
	}
	out := new(RackPosition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackSpec) DeepCopyInto(out *RackSpec) {
	*out = *in
	out.Position = in.Position
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackSpec.
func (in *RackSpec) DeepCopy() *RackSpec {
	if in == nil {
		return nil
	}
	out := new(RackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackStatus) DeepCopyInto(out *RackStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackStatus.
func (in *RackStatus) DeepCopy() *RackStatus {
	if in == nil {
		return nil
	}
	out := new(RackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNAT) DeepCopyInto(out *SNAT) {
	*out = *in
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNAT.
func (in *SNAT) DeepCopy() *SNAT {
	if in == nil {
		return nil
	}
	out := new(SNAT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfile) DeepCopyInto(out *ServerProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfile.
func (in *ServerProfile) DeepCopy() *ServerProfile {
	if in == nil {
		return nil
	}
	out := new(ServerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfileList) DeepCopyInto(out *ServerProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfileList.
func (in *ServerProfileList) DeepCopy() *ServerProfileList {
	if in == nil {
		return nil
	}
	out := new(ServerProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfileNIC) DeepCopyInto(out *ServerProfileNIC) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ServerProfileNICPort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfileNIC.
func (in *ServerProfileNIC) DeepCopy() *ServerProfileNIC {
	if in == nil {
		return nil
	}
	out := new(ServerProfileNIC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfileNICPort) DeepCopyInto(out *ServerProfileNICPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfileNICPort.
func (in *ServerProfileNICPort) DeepCopy() *ServerProfileNICPort {
	if in == nil {
		return nil
	}
	out := new(ServerProfileNICPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfileSpec) DeepCopyInto(out *ServerProfileSpec) {
	*out = *in
	if in.NICs != nil {
		in, out := &in.NICs, &out.NICs
		*out = make([]ServerProfileNIC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfileSpec.
func (in *ServerProfileSpec) DeepCopy() *ServerProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ServerProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfileStatus) DeepCopyInto(out *ServerProfileStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfileStatus.
func (in *ServerProfileStatus) DeepCopy() *ServerProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ServerProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerToSwitchLink) DeepCopyInto(out *ServerToSwitchLink) {
	*out = *in
	out.Server = in.Server
	out.Switch = in.Switch
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerToSwitchLink.
func (in *ServerToSwitchLink) DeepCopy() *ServerToSwitchLink {
	if in == nil {
		return nil
	}
	out := new(ServerToSwitchLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Switch) DeepCopyInto(out *Switch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Switch.
func (in *Switch) DeepCopy() *Switch {
	if in == nil {
		return nil
	}
	out := new(Switch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Switch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchList) DeepCopyInto(out *SwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Switch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchList.
func (in *SwitchList) DeepCopy() *SwitchList {
	if in == nil {
		return nil
	}
	out := new(SwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchProfile) DeepCopyInto(out *SwitchProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchProfile.
func (in *SwitchProfile) DeepCopy() *SwitchProfile {
	if in == nil {
		return nil
	}
	out := new(SwitchProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchProfileLimits) DeepCopyInto(out *SwitchProfileLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchProfileLimits.
func (in *SwitchProfileLimits) DeepCopy() *SwitchProfileLimits {
	if in == nil {
		return nil
	}
	out := new(SwitchProfileLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchProfileList) DeepCopyInto(out *SwitchProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwitchProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchProfileList.
func (in *SwitchProfileList) DeepCopy() *SwitchProfileList {
	if in == nil {
		return nil
	}
	out := new(SwitchProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchProfilePort) DeepCopyInto(out *SwitchProfilePort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchProfilePort.
func (in *SwitchProfilePort) DeepCopy() *SwitchProfilePort {
	if in == nil {
		return nil
	}
	out := new(SwitchProfilePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchProfileSpec) DeepCopyInto(out *SwitchProfileSpec) {
	*out = *in
	out.Limits = in.Limits
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]SwitchProfilePort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchProfileSpec.
func (in *SwitchProfileSpec) DeepCopy() *SwitchProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchProfileStatus) DeepCopyInto(out *SwitchProfileStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchProfileStatus.
func (in *SwitchProfileStatus) DeepCopy() *SwitchProfileStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSpec) DeepCopyInto(out *SwitchSpec) {
	*out = *in
	out.Location = in.Location
	out.LocationSig = in.LocationSig
	out.LLDPConfig = in.LLDPConfig
	if in.PortGroupSpeeds != nil {
		in, out := &in.PortGroupSpeeds, &out.PortGroupSpeeds
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PortBreakouts != nil {
		in, out := &in.PortBreakouts, &out.PortBreakouts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSpec.
func (in *SwitchSpec) DeepCopy() *SwitchSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchStatus) DeepCopyInto(out *SwitchStatus) {
	*out = *in
	in.Applied.DeepCopyInto(&out.Applied)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchStatus.
func (in *SwitchStatus) DeepCopy() *SwitchStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchToSwitchLink) DeepCopyInto(out *SwitchToSwitchLink) {
	*out = *in
	out.Switch1 = in.Switch1
	out.Switch2 = in.Switch2
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchToSwitchLink.
func (in *SwitchToSwitchLink) DeepCopy() *SwitchToSwitchLink {
	if in == nil {
		return nil
	}
	out := new(SwitchToSwitchLink)
	in.DeepCopyInto(out)
	return out
}
