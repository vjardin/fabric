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

	"go.githedgehog.com/fabric/pkg/manager/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalSpec defines the desired state of External
type ExternalSpec struct{}

// ExternalStatus defines the observed state of External
type ExternalStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// External is the Schema for the externals API
type External struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalSpec   `json:"spec,omitempty"`
	Status ExternalStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExternalList contains a list of External
type ExternalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []External `json:"items"`
}

func init() {
	SchemeBuilder.Register(&External{}, &ExternalList{})
}

func (external *External) Default() {
	// TODO
}

func (external *External) Validate(ctx context.Context, client validation.Client) (admission.Warnings, error) {
	return nil, nil // TODO
}
