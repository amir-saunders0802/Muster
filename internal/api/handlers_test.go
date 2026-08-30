package api

// Import packages needed for HTTP testing.
import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amir-saunders0802/Muster/internal/catalog"
)

// TestHandlers tests all of our HTTP endpoints.
func TestHandlers(t *testing.T) {

	// Create a fake Catalog with one service for testing.
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

	// Put the test Catalog into a Handler.
	h := NewHandler(c)

	// Create a test HTTP router.
	mux := http.NewServeMux()

	// Connect each URL to its handler.
	mux.HandleFunc("GET /healthz", Healthz)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /services/{name}", h.ServiceByName)

	// Create a table/list of test cases.
	tests := []struct {
		name       string // Name of the test.
		path       string // URL we want to test.
		wantStatus int    // HTTP status we expect.
		wantBody   string // Text we expect in the response.
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

	// Loop through every test case.
	for _, tt := range tests {

		// Run each test using its name.
		t.Run(tt.name, func(t *testing.T) {

			// Create a fake GET request using the test path.
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)

			// Create something to capture the HTTP response.
			rec := httptest.NewRecorder()

			// Send the request through our router.
			mux.ServeHTTP(rec, req)

			// Fail if the status code is not what we expected.
			if rec.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			// Fail if the response body doesn't contain what we expected.
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