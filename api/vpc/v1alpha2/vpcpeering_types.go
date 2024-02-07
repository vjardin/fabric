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

package v1alpha2

import (
	"context"
	"slices"
	"sort"

	"github.com/pkg/errors"
	wiringapi "go.githedgehog.com/fabric/api/wiring/v1alpha2"
	"go.githedgehog.com/fabric/pkg/manager/validation"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VPCPeeringSpec defines the desired state of VPCPeering
type VPCPeeringSpec struct {
	Remote string `json:"remote,omitempty"`
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=10
	// Permit defines a list of the peering policies - which VPC subnets will have access to the peer VPC subnets.
	Permit []map[string]VPCPeer `json:"permit,omitempty"`
}

type VPCPeer struct {
	//+kubebuilder:validation:MinItems=1
	//+kubebuilder:validation:MaxItems=10
	// Subnets is the list of subnets to advertise from current VPC to the peer VPC
	Subnets []string `json:"subnets,omitempty"`
}

// VPCPeeringStatus defines the observed state of VPCPeering
type VPCPeeringStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=hedgehog;fabric,shortName=vpcpeer
// +kubebuilder:printcolumn:name="VPC1",type=string,JSONPath=`.metadata.labels.fabric\.githedgehog\.com/vpc1`,priority=0
// +kubebuilder:printcolumn:name="VPC2",type=string,JSONPath=`.metadata.labels.fabric\.githedgehog\.com/vpc2`,priority=0
// +kubebuilder:printcolumn:name="Remote",type=string,JSONPath=`.spec.remote`,priority=0
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,priority=0
// VPCPeering represents a peering between two VPCs with corresponding filtering rules.
// Minimal example of the VPC peering showing vpc-1 to vpc-2 peering with all subnets allowed:
//
//	spec:
//	  permit:
//	  - vpc-1: {}
//	    vpc-2: {}
type VPCPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the desired state of the VPCPeering
	Spec VPCPeeringSpec `json:"spec,omitempty"`
	// Status is the observed state of the VPCPeering
	Status VPCPeeringStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VPCPeeringList contains a list of VPCPeering
type VPCPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeering `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCPeering{}, &VPCPeeringList{})
}

func (s *VPCPeeringSpec) VPCs() (string, string, error) {
	vpcs := []string{}
	for _, permit := range s.Permit {
		for vpc := range permit {
			if slices.Contains(vpcs, vpc) {
				return "", "", errors.Errorf("vpc peering must have 2 unique VPCs")
			}
			vpcs = append(vpcs, vpc)
		}
	}

	if len(vpcs) != 2 {
		return "", "", errors.Errorf("vpc peering must have exactly 2 VPCs")
	}

	sort.Strings(vpcs)

	return vpcs[0], vpcs[1], nil
}

func (peering *VPCPeering) Default() {
	if peering.Labels == nil {
		peering.Labels = map[string]string{}
	}

	wiringapi.CleanupFabricLabels(peering.Labels)

	vpc1, vpc2, err := peering.Spec.VPCs()
	if err != nil {
		return // it'll be handled in validation stage
	}

	peering.Labels[ListLabelVPC(vpc1)] = ListLabelValue
	peering.Labels[ListLabelVPC(vpc2)] = ListLabelValue
	peering.Labels[LabelVPC1] = vpc1
	peering.Labels[LabelVPC2] = vpc2
}

func (peering *VPCPeering) Validate(ctx context.Context, client validation.Client, vpcPeeringDisabled bool) (admission.Warnings, error) {
	if vpcPeeringDisabled {
		return nil, errors.Errorf("vpc peering is not allowed")
	}

	vpc1, vpc2, err := peering.Spec.VPCs()
	if err != nil {
		return nil, err
	}

	for idx, permit := range peering.Spec.Permit {
		if len(permit) != 2 {
			return nil, errors.Errorf("permit must have exactly 2 VPCs (idx %d)", idx)
		}
	}

	// TODO validate overlaps in subnets

	if client != nil {
		other := &VPCPeeringList{}
		err := client.List(ctx, other, map[string]string{
			ListLabelVPC(vpc1): ListLabelValue,
			ListLabelVPC(vpc2): ListLabelValue,
		})
		if err != nil && !apierrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "failed to list VPC peerings") // TODO replace with some internal error to not expose to the user
		}

		ipv4Namespaces := []string{}
		vlanNamespaces := []string{}
		for _, vpcName := range []string{vpc1, vpc2} {
			vpc := &VPC{}
			err := client.Get(ctx, types.NamespacedName{Name: vpcName, Namespace: peering.Namespace}, vpc)
			if err != nil {
				if apierrors.IsNotFound(err) {
					return nil, errors.Errorf("vpc %s not found", vpcName)
				}

				return nil, errors.Wrapf(err, "failed to list VPCs") // TODO replace with some internal error to not expose to the user
			}

			ipv4Namespaces = append(ipv4Namespaces, vpc.Spec.IPv4Namespace)
			vlanNamespaces = append(vlanNamespaces, vpc.Spec.VLANNamespace)
		}

		if len(ipv4Namespaces) != 2 {
			return nil, errors.Errorf("failed to find IPv4 namespaces for VPCs")
		}
		if ipv4Namespaces[0] != ipv4Namespaces[1] {
			return nil, errors.Errorf("VPCs must be in the same IPv4 namespace")
		}

		if len(vlanNamespaces) != 2 {
			return nil, errors.Errorf("failed to find VLAN namespaces for VPCs")
		}
		if vlanNamespaces[0] != vlanNamespaces[1] {
			return nil, errors.Errorf("VPCs must be in the same VLAN namespace")
		}

		if peering.Spec.Remote != "" {
			switchGroup := &wiringapi.SwitchGroup{}
			err := client.Get(ctx, types.NamespacedName{Name: peering.Spec.Remote, Namespace: peering.Namespace}, switchGroup)
			if err != nil {
				if apierrors.IsNotFound(err) {
					return nil, errors.Errorf("switch group %s not found", peering.Spec.Remote)
				}

				return nil, errors.Wrapf(err, "failed to list switch groups") // TODO replace with some internal error to not expose to the user
			}
		}
	}

	return nil, nil
}
