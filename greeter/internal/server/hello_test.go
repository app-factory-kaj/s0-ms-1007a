package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloGreeting(t *testing.T) {
	tests := []struct {
		name           string
		query          string // raw query string, e.g. "name=Ada"
		wantStatus     int
		wantNameInBody string // checked against Greeting.Name and Greeting.Message when wantStatus is 200
	}{
		{
			name:           "name provided",
			query:          "name=Ada",
			wantStatus:     http.StatusOK,
			wantNameInBody: "Ada",
		},
		{
			name:           "name omitted defaults to World",
			query:          "",
			wantStatus:     http.StatusOK,
			wantNameInBody: "World",
		},
		{
			name:       "name too long is invalid",
			query:      "name=" + strings.Repeat("a", maxNameLength+1),
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "name with control character is invalid",
			query:      "name=Ada%0A",
			wantStatus: http.StatusBadRequest,
		},
	}

	mux := NewRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/hello"
			if tt.query != "" {
				url += "?" + tt.query
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d (body: %s)", rec.Code, tt.wantStatus, rec.Body.String())
			}

			if rec.Header().Get("Content-Type") != "application/json" {
				t.Fatalf("Content-Type = %q, want application/json", rec.Header().Get("Content-Type"))
			}

			switch tt.wantStatus {
			case http.StatusOK:
				var got Greeting
				if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
					t.Fatalf("decode Greeting: %v (body: %s)", err, rec.Body.String())
				}
				if got.Message == "" {
					t.Fatalf("Greeting.Message is empty")
				}
				if !strings.Contains(got.Message, tt.wantNameInBody) {
					t.Errorf("Greeting.Message = %q, want it to contain %q", got.Message, tt.wantNameInBody)
				}
				if got.Name != tt.wantNameInBody {
					t.Errorf("Greeting.Name = %q, want %q", got.Name, tt.wantNameInBody)
				}
			case http.StatusBadRequest:
				var got apiError
				if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
					t.Fatalf("decode apiError: %v (body: %s)", err, rec.Body.String())
				}
				if got.Code != http.StatusBadRequest {
					t.Errorf("apiError.Code = %d, want %d", got.Code, http.StatusBadRequest)
				}
				if got.Message == "" {
					t.Errorf("apiError.Message is empty")
				}
			}
		})
	}
}

func TestHelloIsStatelessAcrossRequests(t *testing.T) {
	mux := NewRouter()

	// A request with a name must not influence a later request that omits
	// one: Greeter computes each response fresh and keeps no history.
	req1 := httptest.NewRequest(http.MethodGet, "/hello?name=Ada", nil)
	rec1 := httptest.NewRecorder()
	mux.ServeHTTP(rec1, req1)

	req2 := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec2 := httptest.NewRecorder()
	mux.ServeHTTP(rec2, req2)

	var got Greeting
	if err := json.Unmarshal(rec2.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode Greeting: %v (body: %s)", err, rec2.Body.String())
	}
	if got.Name != "World" {
		t.Fatalf("second request Name = %q, want %q (greeting leaked state across requests)", got.Name, "World")
	}
}
