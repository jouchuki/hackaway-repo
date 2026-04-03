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

const openClawImage = "clawbernetes/openclaw:latest"

// OpenClawHarness is the default harness backed by the OpenClaw runtime.
type OpenClawHarness struct{}

func (h *OpenClawHarness) Name() string  { return "openclaw" }
func (h *OpenClawHarness) Image() string { return openClawImage }

// SupportsA2A returns true — OpenClaw has first-class A2A gateway support.
func (h *OpenClawHarness) SupportsA2A() bool { return true }
