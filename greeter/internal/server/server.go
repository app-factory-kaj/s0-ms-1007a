// Package server builds the HTTP router for the greeter service.
package server

import (
	"net/http"

	"greeter/internal/response"
)

// NewRouter builds the HTTP mux for the greeter service.
//
// It currently exposes no application routes: the GET /hello handler from
// the component's openapi.yaml lands in a follow-up issue, added to this
// mux without restructuring it. For now, "/" answers a trivial liveness
// response and every other path falls through to the shared JSON error
// shape, so that convention is already exercised end to end.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.WriteError(w, http.StatusNotFound, "not found")
	})

	return mux
}
