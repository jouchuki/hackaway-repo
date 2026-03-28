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

// GatewayWebhookSpec defines a webhook target for alerts.
type GatewayWebhookSpec struct {
	// url is the webhook endpoint.
	// +optional
	URL string `json:"url,omitempty"`

	// minSeverity is the minimum severity to trigger the webhook.
	// +optional
	MinSeverity string `json:"minSeverity,omitempty"`

	// headers are additional HTTP headers to send.
	// +optional
	Headers map[string]string `json:"headers,omitempty"`
}

// EvaluatorRouteSpec defines a model route for a complexity class.
type EvaluatorRouteSpec struct {
	// provider is the model provider.
	// +optional
	Provider string `json:"provider,omitempty"`

	// model is the model name.
	// +optional
	Model string `json:"model,omitempty"`
}

// GatewayEvaluatorSpec defines a single routing evaluator.
type GatewayEvaluatorSpec struct {
	// name is the evaluator name.
	// +optional
	Name string `json:"name,omitempty"`

	// type is the evaluator type (regex, classifier).
	// +optional
	Type string `json:"type,omitempty"`

	// priority determines evaluation order (highest first).
	// +optional
	Priority int `json:"priority,omitempty"`

	// patterns is a list of regex patterns for regex-type evaluators.
	// +optional
	Patterns []string `json:"patterns,omitempty"`

	// action is what to do on match (block, proxy, route).
	// +optional
	Action string `json:"action,omitempty"`

	// blockReply is the message returned when action is block.
	// +optional
	BlockReply string `json:"blockReply,omitempty"`

	// emitEvent controls whether a Kubernetes event is fired.
	// +optional
	EmitEvent bool `json:"emitEvent,omitempty"`

	// classifierModel is the model used for classifier-type evaluators.
	// +optional
	ClassifierModel string `json:"classifierModel,omitempty"`

	// ollamaEndpoint is the Ollama endpoint for local classifier models.
	// +optional
	OllamaEndpoint string `json:"ollamaEndpoint,omitempty"`

	// timeoutMs is the timeout in milliseconds for the evaluator.
	// +optional
	TimeoutMs int `json:"timeoutMs,omitempty"`

	// proxyUrl is the proxy endpoint for proxy-type actions.
	// +optional
	ProxyURL string `json:"proxyUrl,omitempty"`

	// redactReplacement is the replacement text for redacted content.
	// +optional
	RedactReplacement string `json:"redactReplacement,omitempty"`

	// webhooks are webhook targets for this evaluator.
	// +optional
	Webhooks []GatewayWebhookSpec `json:"webhooks,omitempty"`

	// classifierEndpoint is the endpoint for the classifier model.
	// +optional
	ClassifierEndpoint string `json:"classifierEndpoint,omitempty"`

	// routes maps complexity classes to model routes.
	// +optional
	Routes map[string]EvaluatorRouteSpec `json:"routes,omitempty"`
}

// GatewayRoutingSpec defines the model routing pipeline.
type GatewayRoutingSpec struct {
	// enabled controls whether routing is active.
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// logEveryDecision controls whether every routing decision is logged.
	// +optional
	LogEveryDecision bool `json:"logEveryDecision,omitempty"`

	// evaluators is the ordered list of routing evaluators.
	// +optional
	Evaluators []GatewayEvaluatorSpec `json:"evaluators,omitempty"`
}

// GatewayAnomalySpec defines anomaly detection thresholds.
type GatewayAnomalySpec struct {
	// spendSpikeMultiplier triggers an alert when hourly spend exceeds this multiple of the 7-day average.
	// +optional
	SpendSpikeMultiplier int `json:"spendSpikeMultiplier,omitempty"`

	// idleBurnMinutes triggers an alert after this many minutes of continuous LLM calls with no useful output.
	// +optional
	IdleBurnMinutes int `json:"idleBurnMinutes,omitempty"`

	// errorLoopThreshold auto-pauses the agent after this many consecutive errors.
	// +optional
	ErrorLoopThreshold int `json:"errorLoopThreshold,omitempty"`

	// tokenInflationMultiplier alerts when average input tokens exceed this multiple.
	// +optional
	TokenInflationMultiplier int `json:"tokenInflationMultiplier,omitempty"`

	// checkIntervalSeconds is how often anomaly detectors run.
	// +optional
	CheckIntervalSeconds int `json:"checkIntervalSeconds,omitempty"`
}

// ClawGatewaySpec defines the desired state of ClawGateway.
type ClawGatewaySpec struct {
	// topology is the gateway deployment model ("centralized" or "sidecar").
	// +optional
	// +kubebuilder:validation:Enum=centralized;sidecar
	Topology string `json:"topology,omitempty"`

	// port is the port the gateway listens on.
	// +optional
	Port int `json:"port,omitempty"`

	// routing defines the model routing pipeline.
	// +optional
	Routing GatewayRoutingSpec `json:"routing,omitempty"`

	// anomaly defines anomaly detection thresholds.
	// +optional
	Anomaly GatewayAnomalySpec `json:"anomaly,omitempty"`

	// webhooks defines alerting webhook targets.
	// +optional
	Webhooks []GatewayWebhookSpec `json:"webhooks,omitempty"`
}

// ClawGatewayStatus defines the observed state of ClawGateway.
type ClawGatewayStatus struct {
	// conditions represent the current state of the ClawGateway resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClawGateway is the Schema for the clawgateways API.
type ClawGateway struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ClawGateway
	// +required
	Spec ClawGatewaySpec `json:"spec"`

	// status defines the observed state of ClawGateway
	// +optional
	Status ClawGatewayStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ClawGatewayList contains a list of ClawGateway
type ClawGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ClawGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClawGateway{}, &ClawGatewayList{})
}
