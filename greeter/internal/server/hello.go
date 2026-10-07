package server

import (
	"fmt"
	"net/http"
	"unicode"

	"greeter/internal/response"
)

// maxNameLength bounds the name query parameter. The openapi.yaml contract
// sets no upper bound, so this is a narrow sanity check against pathological
// input (not a documented business rule): anything longer than a plausible
// display name is rejected rather than echoed verbatim into a 200 response.
const maxNameLength = 100

// Greeting is the response body for a successful GET /hello, matching the
// Greeting schema in specs/design/components/greeter/openapi.yaml exactly.
type Greeting struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// apiError is the response body for a failed GET /hello, matching the Error
// schema in specs/design/components/greeter/openapi.yaml exactly. It is
// deliberately distinct from response.WriteError's {"error": "..."} shape:
// that shape is this service's internal fallback for routes with no public
// contract (the catch-all 404 in NewRouter), while /hello is the component's
// own documented API and its error body must match that contract, which
// consumers code against.
type apiError struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

// writeAPIError writes an apiError body in the openapi.yaml Error shape.
func writeAPIError(w http.ResponseWriter, status int, message, description string) {
	response.WriteJSON(w, status, apiError{
		Code:        status,
		Message:     message,
		Description: description,
	})
}

// validateName applies one narrow, intentionally minimal rule to the
// optional name query parameter: it must not exceed maxNameLength
// characters and must not contain control characters (e.g. a newline or
// tab), which would be nonsensical in a greeting and could smuggle
// unexpected formatting into the response body. Nothing more elaborate is
// warranted for a single-field greeting input. Omitting the parameter
// entirely never reaches this function — handleHello defaults it to
// "World" before validation runs, so a missing name is never "invalid".
func validateName(name string) error {
	if len(name) > maxNameLength {
		return fmt.Errorf("name must be at most %d characters", maxNameLength)
	}
	for _, r := range name {
		if unicode.IsControl(r) {
			return fmt.Errorf("name must not contain control characters")
		}
	}
	return nil
}

// handleHello implements GET /hello from the component's openapi.yaml: an
// optional name query parameter, defaulting to "World" when omitted, and a
// JSON Greeting echoing the (possibly defaulted) name back to the caller.
//
// Greeter is stateless per its domain model: the greeting is computed fresh
// from only this request's own query parameter on every call. Nothing is
// persisted, cached or logged across requests.
func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	} else if err := validateName(name); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid name parameter", err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, Greeting{
		Message: fmt.Sprintf("Hello, %s!", name),
		Name:    name,
	})
}
