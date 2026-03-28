/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ToolPolicySpec defines which tools an agent is allowed or denied.
type ToolPolicySpec struct {
	// allow is a whitelist of permitted tool names.
	// +optional
	Allow []string `json:"allow,omitempty"`

	// deny is a blacklist of denied tool names.
	// +optional
	Deny []string `json:"deny,omitempty"`
}

// BudgetSpec defines per-agent spend limits.
type BudgetSpec struct {
	// daily is the USD daily spend limit.
	// +optional
	Daily int `json:"daily,omitempty"`

	// monthly is the USD monthly spend limit.
	// +optional
	Monthly int `json:"monthly,omitempty"`

	// warnAt is the fraction of budget at which to downgrade model (e.g. "0.9").
	// +optional
	WarnAt string `json:"warnAt,omitempty"`

	// downgradeModel is the model to switch to when warnAt threshold is hit.
	// +optional
	DowngradeModel string `json:"downgradeModel,omitempty"`

	// downgradeProvider is the provider for the downgrade model.
	// +optional
	DowngradeProvider string `json:"downgradeProvider,omitempty"`
}

// ClawPolicySpec defines the desired state of ClawPolicy.
type ClawPolicySpec struct {
	// toolPolicy controls which tools are allowed or denied.
	// +optional
	ToolPolicy ToolPolicySpec `json:"toolPolicy,omitempty"`

	// budget defines spend limits enforced by the gateway.
	// +optional
	Budget BudgetSpec `json:"budget,omitempty"`
}

// ClawPolicyStatus defines the observed state of ClawPolicy.
type ClawPolicyStatus struct {
	// conditions represent the current state of the ClawPolicy resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClawPolicy is the Schema for the clawpolicies API.
type ClawPolicy struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ClawPolicy
	// +required
	Spec ClawPolicySpec `json:"spec"`

	// status defines the observed state of ClawPolicy
	// +optional
	Status ClawPolicyStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ClawPolicyList contains a list of ClawPolicy
type ClawPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ClawPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClawPolicy{}, &ClawPolicyList{})
}
