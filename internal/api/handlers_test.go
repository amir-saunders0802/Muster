package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amir-saunders0802/Muster/internal/catalog"
)

// TestHandlers checks the expected response for each API route, including the
// error returned when a requested service does not exist.
func TestHandlers(t *testing.T) {

	// Use a small in-memory catalog so the handler tests do not depend on the
	// production configuration file.
	c := catalog.Catalog{
		Services: []catalog.Service{
			{
				Name:        "payments-api",
				Owner:       "payments-team",
				Repo:        "github.com/example/payments-api",
				Environment: "production",
				Tier:        1,
			},
		},
	}

	h := NewHandler(c)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", Healthz)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /services/{name}", h.ServiceByName)

	// Keep the route checks together so each endpoint has a clear expected
	// status code and response value.
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "health check",
			path:       "/healthz",
			wantStatus: http.StatusOK,
			wantBody:   `"status":"ok"`,
		},
		{
			name:       "list services",
			path:       "/services",
			wantStatus: http.StatusOK,
			wantBody:   `"payments-api"`,
		},
		{
			name:       "find service by name",
			path:       "/services/payments-api",
			wantStatus: http.StatusOK,
			wantBody:   `"payments-api"`,
		},
		{
			name:       "service not found",
			path:       "/services/does-not-exist",
			wantStatus: http.StatusNotFound,
			wantBody:   "service not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			if !strings.Contains(rec.Body.String(), tt.wantBody) {
				t.Errorf(
					"expected body to contain %q, got %q",
					tt.wantBody,
					rec.Body.String(),
				)
			}
		})
	}
}
