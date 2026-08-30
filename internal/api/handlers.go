package api

import (
	"encoding/json"
	"net/http"

	// Import the catalog package so this file can use
	// the Catalog type we created in internal/catalog/catalog.go.
	"github.com/amir-saunders0802/Muster/internal/catalog"
)

// Handler contains a Catalog so our HTTP handlers
// can access the service data later.
type Handler struct {
	// catalog = the name of this field.
	// catalog.Catalog = the Catalog type from the catalog.go file.
	// This is how we store the catalog inside the Handler struct.
	catalog catalog.Catalog
	// In plain English:
    // "A Handler has a field named catalog that holds a Catalog."
}

// Healthz handles GET /healthz requests.
func Healthz(w http.ResponseWriter, r *http.Request) {
	// w is the response writer: this is how we send data back.
	// r is the incoming HTTP request.
	// *http.Request means r is a pointer to the request.

	// Tell the client that the response body will be JSON.
	w.Header().Set("Content-Type", "application/json")

	// NewEncoder(w) creates a JSON encoder that writes to the response.
	// Encode(...) is like pressing start: it converts the Go map into JSON.
	json.NewEncoder(w).Encode(map[string]string{
		// This is hard-coded because /healthz only needs to prove
		// that the server is alive and able to respond.
		"status": "ok",
	})
}

// Services handles requests for GET /services.
// (h *Handler) makes Services a method on Handler.
// This lets us access the catalog stored inside the Handler using h.catalog.
func (h *Handler) Services(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
    // h.catalog.Services gets the []Service stored inside the Handler's Catalog.
    //
    // json.NewEncoder(w) creates a JSON encoder that writes to the
    // HTTP response writer (w).
    //
    // Encode(h.catalog.Services) converts our Go []Service data into JSON
    // and writes that JSON to w, which sends it back in the HTTP response.
    json.NewEncoder(w).Encode(h.catalog.Services)
}

// ServiceByName handles GET /services/{name} requests.
func (h *Handler) ServiceByName(w http.ResponseWriter, r *http.Request) {

    // Get the {name} value from the URL.
    // Example: /services/payments-api → name = "payments-api"
    name := r.PathValue("name")

    // Loop through every service in the Handler's catalog.
    // _ ignores the index; service is the current service being checked.
    for _, service := range h.catalog.Services {

        // Check if the current service's Name matches the name from the URL.
        if service.Name == name {

            // Tell the client that we are sending JSON.
            w.Header().Set("Content-Type", "application/json")

            // Convert the matching service to JSON and write it to the response.
            json.NewEncoder(w).Encode(service)

            // We found the service and sent it, so stop the function.
            return
        }
    }

    // If the loop finishes without finding the service, return a 404 error.
    http.Error(w, "service not found", http.StatusNotFound)
}

// NewHandler creates a new Handler and gives it a Catalog.
//
// c is the Catalog passed into this function.
//
// &Handler creates a new Handler and returns a pointer to it.
//
// catalog: c means:
// "Put the Catalog stored in c into the Handler's catalog field."

func NewHandler(c catalog.Catalog) *Handler {
	return &Handler{
		catalog: c,
	}
}
// NewHandler takes the Catalog we already loaded from YAML
// and gives it to our Handler so /services and /services/{name}
// can access the service data.