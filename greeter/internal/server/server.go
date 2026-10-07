// Package server builds the HTTP router for the greeter service.
package server

import (
	"net/http"

	"greeter/internal/response"
)

// NewRouter builds the HTTP mux for the greeter service.
//
// "/" answers a trivial liveness response, "GET /hello" implements the
// component's openapi.yaml contract (see hello.go), and every other path
// falls through to the shared internal JSON error shape.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /hello", handleHello)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.WriteError(w, http.StatusNotFound, "not found")
	})

	return mux
}
