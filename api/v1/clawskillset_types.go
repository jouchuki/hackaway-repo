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

// SkillEntry defines a single skill with its name and content.
type SkillEntry struct {
	// name is the skill identifier (used as the directory name).
	// +required
	Name string `json:"name"`

	// content is the skill's SKILL.md content.
	// +required
	Content string `json:"content"`
}

// ClawSkillSetSpec defines the desired state of ClawSkillSet.
type ClawSkillSetSpec struct {
	// skills is the list of skills in this set.
	// +optional
	Skills []SkillEntry `json:"skills,omitempty"`
}

// ClawSkillSetStatus defines the observed state of ClawSkillSet.
type ClawSkillSetStatus struct {
	// conditions represent the current state of the ClawSkillSet resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClawSkillSet is the Schema for the clawskillsets API.
type ClawSkillSet struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ClawSkillSet
	// +required
	Spec ClawSkillSetSpec `json:"spec"`

	// status defines the observed state of ClawSkillSet
	// +optional
	Status ClawSkillSetStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ClawSkillSetList contains a list of ClawSkillSet
type ClawSkillSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ClawSkillSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClawSkillSet{}, &ClawSkillSetList{})
}
