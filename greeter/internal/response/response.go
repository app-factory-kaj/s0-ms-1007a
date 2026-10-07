// Package response holds the shared JSON response helpers for the greeter
// service, so every handler — present and future — answers errors with the
// same shape.
package response

import (
	"encoding/json"
	"net/http"
)

// errorBody is the JSON error shape used by every error response in this
// service: {"error": "message"}.
type errorBody struct {
	Error string `json:"error"`
}

// WriteError writes a JSON error body with the given HTTP status code, in
// the service-wide {"error": "message"} shape.
func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorBody{Error: message})
}

// WriteJSON writes a successful JSON response with the given HTTP status
// code and body.
func WriteJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
