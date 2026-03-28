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

// ConnectorRateLimitSpec defines rate limiting for the connector.
type ConnectorRateLimitSpec struct {
	// requestsPerMinute is the maximum number of requests per minute.
	// +optional
	RequestsPerMinute int `json:"requestsPerMinute,omitempty"`
}

// ClawConnectorSpec defines the desired state of ClawConnector.
type ClawConnectorSpec struct {
	// type is the connector type (e.g. "database", "http").
	// +optional
	Type string `json:"type,omitempty"`

	// driver is the database driver (e.g. "postgresql"). Only for database type.
	// +optional
	Driver string `json:"driver,omitempty"`

	// host is the database host. Only for database type.
	// +optional
	Host string `json:"host,omitempty"`

	// port is the database port. Only for database type.
	// +optional
	Port int `json:"port,omitempty"`

	// database is the database name. Only for database type.
	// +optional
	Database string `json:"database,omitempty"`

	// baseUrl is the base URL. Only for http type.
	// +optional
	BaseURL string `json:"baseUrl,omitempty"`

	// credentialsSecret is the name of the Kubernetes Secret containing credentials.
	// +optional
	CredentialsSecret string `json:"credentialsSecret,omitempty"`

	// permissions is a list of allowed operations (e.g. SELECT, INSERT).
	// +optional
	Permissions []string `json:"permissions,omitempty"`

	// rateLimit defines rate limiting for the connector.
	// +optional
	RateLimit ConnectorRateLimitSpec `json:"rateLimit,omitempty"`
}

// ClawConnectorStatus defines the observed state of ClawConnector.
type ClawConnectorStatus struct {
	// conditions represent the current state of the ClawConnector resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClawConnector is the Schema for the clawconnectors API.
type ClawConnector struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ClawConnector
	// +required
	Spec ClawConnectorSpec `json:"spec"`

	// status defines the observed state of ClawConnector
	// +optional
	Status ClawConnectorStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ClawConnectorList contains a list of ClawConnector
type ClawConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ClawConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClawConnector{}, &ClawConnectorList{})
}
