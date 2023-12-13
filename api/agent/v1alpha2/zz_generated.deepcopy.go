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
	vpcv1alpha2 "go.githedgehog.com/fabric/api/vpc/v1alpha2"
	wiringv1alpha2 "go.githedgehog.com/fabric/api/wiring/v1alpha2"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentList.
func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	out.Version = in.Version
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]UserCreds, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Switch.DeepCopyInto(&out.Switch)
	if in.Switches != nil {
		in, out := &in.Switches, &out.Switches
		*out = make(map[string]wiringv1alpha2.SwitchSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make(map[string]wiringv1alpha2.ConnectionSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.VPCs != nil {
		in, out := &in.VPCs, &out.VPCs
		*out = make(map[string]vpcv1alpha2.VPCSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.VPCAttachments != nil {
		in, out := &in.VPCAttachments, &out.VPCAttachments
		*out = make(map[string]vpcv1alpha2.VPCAttachmentSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VPCPeerings != nil {
		in, out := &in.VPCPeerings, &out.VPCPeerings
		*out = make(map[string]vpcv1alpha2.VPCPeeringSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.VPCLoopbackLinks != nil {
		in, out := &in.VPCLoopbackLinks, &out.VPCLoopbackLinks
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VPCLoopbackVLANs != nil {
		in, out := &in.VPCLoopbackVLANs, &out.VPCLoopbackVLANs
		*out = make(map[string]uint16, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Externals != nil {
		in, out := &in.Externals, &out.Externals
		*out = make(map[string]vpcv1alpha2.ExternalSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalAttachments != nil {
		in, out := &in.ExternalAttachments, &out.ExternalAttachments
		*out = make(map[string]vpcv1alpha2.ExternalAttachmentSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalPeerings != nil {
		in, out := &in.ExternalPeerings, &out.ExternalPeerings
		*out = make(map[string]vpcv1alpha2.ExternalPeeringSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ConfiguredVPCSubnets != nil {
		in, out := &in.ConfiguredVPCSubnets, &out.ConfiguredVPCSubnets
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MCLAGAttachedVPCs != nil {
		in, out := &in.MCLAGAttachedVPCs, &out.MCLAGAttachedVPCs
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VNIs != nil {
		in, out := &in.VNIs, &out.VNIs
		*out = make(map[string]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IRBVLANs != nil {
		in, out := &in.IRBVLANs, &out.IRBVLANs
		*out = make(map[string]uint16, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalPeeringPrefixIDs != nil {
		in, out := &in.ExternalPeeringPrefixIDs, &out.ExternalPeeringPrefixIDs
		*out = make(map[string]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalSeqs != nil {
		in, out := &in.ExternalSeqs, &out.ExternalSeqs
		*out = make(map[string]uint16, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PortChannels != nil {
		in, out := &in.PortChannels, &out.PortChannels
		*out = make(map[string]uint16, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StatusUpdates != nil {
		in, out := &in.StatusUpdates, &out.StatusUpdates
		*out = make([]ApplyStatusUpdate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpecConfig) DeepCopyInto(out *AgentSpecConfig) {
	*out = *in
	if in.CollapsedCore != nil {
		in, out := &in.CollapsedCore, &out.CollapsedCore
		*out = new(AgentSpecConfigCollapsedCore)
		**out = **in
	}
	if in.SpineLeaf != nil {
		in, out := &in.SpineLeaf, &out.SpineLeaf
		*out = new(AgentSpecConfigSpineLeaf)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpecConfig.
func (in *AgentSpecConfig) DeepCopy() *AgentSpecConfig {
	if in == nil {
		return nil
	}
	out := new(AgentSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpecConfigCollapsedCore) DeepCopyInto(out *AgentSpecConfigCollapsedCore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpecConfigCollapsedCore.
func (in *AgentSpecConfigCollapsedCore) DeepCopy() *AgentSpecConfigCollapsedCore {
	if in == nil {
		return nil
	}
	out := new(AgentSpecConfigCollapsedCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpecConfigSpineLeaf) DeepCopyInto(out *AgentSpecConfigSpineLeaf) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpecConfigSpineLeaf.
func (in *AgentSpecConfigSpineLeaf) DeepCopy() *AgentSpecConfigSpineLeaf {
	if in == nil {
		return nil
	}
	out := new(AgentSpecConfigSpineLeaf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	in.LastHeartbeat.DeepCopyInto(&out.LastHeartbeat)
	in.LastAttemptTime.DeepCopyInto(&out.LastAttemptTime)
	in.LastAppliedTime.DeepCopyInto(&out.LastAppliedTime)
	out.NOSInfo = in.NOSInfo
	if in.StatusUpdates != nil {
		in, out := &in.StatusUpdates, &out.StatusUpdates
		*out = make([]ApplyStatusUpdate, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentStatus.
func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentVersion) DeepCopyInto(out *AgentVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentVersion.
func (in *AgentVersion) DeepCopy() *AgentVersion {
	if in == nil {
		return nil
	}
	out := new(AgentVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyStatusUpdate) DeepCopyInto(out *ApplyStatusUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyStatusUpdate.
func (in *ApplyStatusUpdate) DeepCopy() *ApplyStatusUpdate {
	if in == nil {
		return nil
	}
	out := new(ApplyStatusUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlAgent) DeepCopyInto(out *ControlAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlAgent.
func (in *ControlAgent) DeepCopy() *ControlAgent {
	if in == nil {
		return nil
	}
	out := new(ControlAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlAgentList) DeepCopyInto(out *ControlAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControlAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlAgentList.
func (in *ControlAgentList) DeepCopy() *ControlAgentList {
	if in == nil {
		return nil
	}
	out := new(ControlAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlAgentSpec) DeepCopyInto(out *ControlAgentSpec) {
	*out = *in
	out.Version = in.Version
	if in.Networkd != nil {
		in, out := &in.Networkd, &out.Networkd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlAgentSpec.
func (in *ControlAgentSpec) DeepCopy() *ControlAgentSpec {
	if in == nil {
		return nil
	}
	out := new(ControlAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlAgentStatus) DeepCopyInto(out *ControlAgentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastHeartbeat.DeepCopyInto(&out.LastHeartbeat)
	in.LastAttemptTime.DeepCopyInto(&out.LastAttemptTime)
	in.LastAppliedTime.DeepCopyInto(&out.LastAppliedTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlAgentStatus.
func (in *ControlAgentStatus) DeepCopy() *ControlAgentStatus {
	if in == nil {
		return nil
	}
	out := new(ControlAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NOSInfo) DeepCopyInto(out *NOSInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NOSInfo.
func (in *NOSInfo) DeepCopy() *NOSInfo {
	if in == nil {
		return nil
	}
	out := new(NOSInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserCreds) DeepCopyInto(out *UserCreds) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserCreds.
func (in *UserCreds) DeepCopy() *UserCreds {
	if in == nil {
		return nil
	}
	out := new(UserCreds)
	in.DeepCopyInto(out)
	return out
}
