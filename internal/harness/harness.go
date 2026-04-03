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

package harness

// Harness abstracts the agent runtime that wraps an LLM.
// Different harnesses (OpenClaw, Hermes, etc.) have different capabilities;
// the controller uses this interface to decide which features to configure.
type Harness interface {
	// Name returns the harness identifier (e.g. "openclaw", "hermes").
	Name() string

	// Image returns the container image for this harness.
	Image() string

	// SupportsA2A reports whether the harness supports the Agent-to-Agent
	// communication protocol. Harnesses that return false will have A2A
	// port, env vars, and plugin configuration skipped by the controller.
	SupportsA2A() bool
}

// ForName returns the Harness implementation for the given name.
// Unrecognised names default to OpenClaw.
func ForName(name string) Harness {
	switch name {
	case "hermes":
		return &HermesHarness{}
	case "observeclaw":
		return &ObserveClawHarness{}
	default:
		return &OpenClawHarness{}
	}
}
