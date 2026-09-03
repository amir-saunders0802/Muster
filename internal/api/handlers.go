package api

import (
	"encoding/json"
	"net/http"

	"github.com/amir-saunders0802/Muster/internal/catalog"
)

// Handler provides HTTP handlers with access to the service catalog.
type Handler struct {
	catalog catalog.Catalog
}

// Healthz handles GET /healthz requests.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// A fixed response gives monitoring tools a lightweight way to confirm that
	// the API is running and able to answer requests.
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// Services handles requests for GET /services.
func (h *Handler) Services(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.catalog.Services)
}

// ServiceByName handles GET /services/{name} requests.
func (h *Handler) ServiceByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	for _, service := range h.catalog.Services {
		if service.Name == name {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(service)
			return
		}
	}

	// Unknown service names should produce a clear client error.
	http.Error(w, "service not found", http.StatusNotFound)
}

// NewHandler connects the loaded catalog to the API handlers.
func NewHandler(c catalog.Catalog) *Handler {
	return &Handler{
		catalog: c,
	}
}
