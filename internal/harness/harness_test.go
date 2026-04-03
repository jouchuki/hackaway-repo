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

import "testing"

func TestHarnessProperties(t *testing.T) {
	tests := []struct {
		name      string
		harness   Harness
		wantName  string
		wantImage string
		wantA2A   bool
	}{
		{
			name:      "OpenClaw supports A2A",
			harness:   &OpenClawHarness{},
			wantName:  "openclaw",
			wantImage: "clawbernetes/openclaw:latest",
			wantA2A:   true,
		},
		{
			name:      "ObserveClaw inherits A2A support from OpenClaw",
			harness:   &ObserveClawHarness{},
			wantName:  "observeclaw",
			wantImage: "clawbernetes/openclaw:latest",
			wantA2A:   true,
		},
		{
			name:      "Hermes does not support A2A",
			harness:   &HermesHarness{},
			wantName:  "hermes",
			wantImage: "clawbernetes/hermes:latest",
			wantA2A:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.harness.Name(); got != tt.wantName {
				t.Errorf("Name() = %q, want %q", got, tt.wantName)
			}
			if got := tt.harness.Image(); got != tt.wantImage {
				t.Errorf("Image() = %q, want %q", got, tt.wantImage)
			}
			if got := tt.harness.SupportsA2A(); got != tt.wantA2A {
				t.Errorf("SupportsA2A() = %v, want %v", got, tt.wantA2A)
			}
		})
	}
}

func TestForName(t *testing.T) {
	tests := []struct {
		input    string
		wantName string
		wantA2A  bool
	}{
		{"openclaw", "openclaw", true},
		{"observeclaw", "observeclaw", true},
		{"hermes", "hermes", false},
		{"unknown", "openclaw", true}, // defaults to OpenClaw
		{"", "openclaw", true},        // empty defaults to OpenClaw
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			h := ForName(tt.input)
			if got := h.Name(); got != tt.wantName {
				t.Errorf("ForName(%q).Name() = %q, want %q", tt.input, got, tt.wantName)
			}
			if got := h.SupportsA2A(); got != tt.wantA2A {
				t.Errorf("ForName(%q).SupportsA2A() = %v, want %v", tt.input, got, tt.wantA2A)
			}
		})
	}
}
