package api

import (
	"encoding/json"
	"net/http"
)

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