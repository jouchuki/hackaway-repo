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

// TempoStorageSpec defines storage configuration for Tempo.
type TempoStorageSpec struct {
	// size is the PVC size for trace storage.
	// +optional
	Size string `json:"size,omitempty"`

	// storageClass is the Kubernetes StorageClass to use. Empty string uses cluster default.
	// +optional
	StorageClass string `json:"storageClass,omitempty"`
}

// TempoSpec defines the Grafana Tempo (OTLP collector) configuration.
type TempoSpec struct {
	// enabled controls whether the operator deploys Tempo.
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// retentionDays is how long to retain trace data.
	// +optional
	RetentionDays int `json:"retentionDays,omitempty"`

	// storage defines the PVC configuration for Tempo.
	// +optional
	Storage TempoStorageSpec `json:"storage,omitempty"`
}

// GrafanaSpec defines the Grafana visualization layer configuration.
type GrafanaSpec struct {
	// enabled controls whether the operator deploys Grafana.
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// dashboards is the list of pre-built dashboard names to install.
	// +optional
	Dashboards []string `json:"dashboards,omitempty"`

	// adminCredentialsSecret is the name of the Secret containing Grafana admin credentials.
	// +optional
	AdminCredentialsSecret string `json:"adminCredentialsSecret,omitempty"`

	// expose defines how Grafana is exposed outside the cluster.
	// Valid values: "port-forward", "loadbalancer", "ingress".
	// +optional
	// +kubebuilder:validation:Enum=port-forward;loadbalancer;ingress
	Expose string `json:"expose,omitempty"`
}

// ClawObservabilitySpec defines the desired state of ClawObservability.
type ClawObservabilitySpec struct {
	// tempo defines the Grafana Tempo OTLP collector configuration.
	// +optional
	Tempo TempoSpec `json:"tempo,omitempty"`

	// grafana defines the Grafana visualization configuration.
	// +optional
	Grafana GrafanaSpec `json:"grafana,omitempty"`

	// otlpEndpoint is the cluster-internal OTLP endpoint that agents write traces to.
	// +optional
	OTLPEndpoint string `json:"otlpEndpoint,omitempty"`

	// otlpProtocol is the OTLP transport protocol (e.g. "http/protobuf").
	// +optional
	OTLPProtocol string `json:"otlpProtocol,omitempty"`
}

// ClawObservabilityStatus defines the observed state of ClawObservability.
type ClawObservabilityStatus struct {
	// conditions represent the current state of the ClawObservability resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// tempoReady indicates whether the Tempo deployment is ready.
	// +optional
	TempoReady bool `json:"tempoReady,omitempty"`

	// grafanaReady indicates whether the Grafana deployment is ready.
	// +optional
	GrafanaReady bool `json:"grafanaReady,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Tempo",type=boolean,JSONPath=`.status.tempoReady`
// +kubebuilder:printcolumn:name="Grafana",type=boolean,JSONPath=`.status.grafanaReady`

// ClawObservability is the Schema for the clawobservabilities API.
// It declares the telemetry stack (Tempo + Grafana) as a first-class Kubernetes primitive.
type ClawObservability struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ClawObservability
	// +required
	Spec ClawObservabilitySpec `json:"spec"`

	// status defines the observed state of ClawObservability
	// +optional
	Status ClawObservabilityStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ClawObservabilityList contains a list of ClawObservability
type ClawObservabilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ClawObservability `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClawObservability{}, &ClawObservabilityList{})
}
