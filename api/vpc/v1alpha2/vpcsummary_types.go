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
	wiringapi "go.githedgehog.com/fabric/api/wiring/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VPCSummarySpec defines the desired state of VPCSummary
type VPCSummarySpec struct {
	Name        string   `json:"name"`
	VPC         VPCSpec  `json:"vpc"`
	VLAN        uint16   `json:"vlan"`
	Peers       []string `json:"peers,omitempty"`
	Connections []string `json:"connections"`
	// TODO Connection NS
}

// VPCSummaryStatus defines the observed state of VPCSummary
type VPCSummaryStatus struct {
	Applied wiringapi.ApplyStatus `json:"applied,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=hedgehog;fabric,shortName=vpcsum;vs
// +kubebuilder:printcolumn:name="VPC",type=string,JSONPath=`.spec.vpc`,priority=0
// +kubebuilder:printcolumn:name="VLAN",type=string,JSONPath=`.spec.vlan`,priority=0
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,priority=0
// VPCSummary is the Schema for the vpcsummaries API
type VPCSummary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPCSummarySpec   `json:"spec,omitempty"`
	Status VPCSummaryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VPCSummaryList contains a list of VPCSummary
type VPCSummaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCSummary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCSummary{}, &VPCSummaryList{})
}
