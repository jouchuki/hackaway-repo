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

package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Server is the dashboard API server that exposes log streaming and chat relay
// endpoints for ClawAgent resources.
type Server struct {
	// Client is the controller-runtime client for reading CRDs.
	Client client.Client

	// Clientset is the typed Kubernetes client for pod log access.
	Clientset kubernetes.Interface

	// Namespace is the namespace where agents run.
	Namespace string
}

// Start launches the API server on the given address (e.g. ":9090").
// It blocks until the context is cancelled.
func (s *Server) Start(ctx context.Context, addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/agents/", s.routeAgentSubresource)

	srv := &http.Server{Addr: addr, Handler: mux}

	go func() {
		<-ctx.Done()
		_ = srv.Shutdown(context.Background())
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("api server: %w", err)
	}
	return nil
}

// routeAgentSubresource dispatches /api/agents/{name}/logs and
// /api/agents/{name}/chat to the appropriate handler.
func (s *Server) routeAgentSubresource(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/api/agents/")
	rest = strings.TrimSuffix(rest, "/")

	parts := strings.SplitN(rest, "/", 2)
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "agent name required", http.StatusBadRequest)
		return
	}

	name := parts[0]
	sub := ""
	if len(parts) == 2 {
		sub = parts[1]
	}

	switch sub {
	case "logs":
		s.handleAgentLogs(w, r, name)
	case "chat":
		s.handleAgentChat(w, r, name)
	default:
		http.Error(w, "unknown sub-resource", http.StatusNotFound)
	}
}
