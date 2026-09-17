package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/amir-saunders0802/Muster/internal/api"
	"github.com/amir-saunders0802/Muster/internal/catalog"
)

func main() {
	// Load the service information before accepting requests so every endpoint
	// can use the same catalog.
	c, err := catalog.Load("config/services.yaml")
	if err != nil {
		// The API cannot provide useful responses without its catalog, so startup
		// stops instead of running with incomplete data.
		slog.Error("failed to load catalog", "error", err)
		os.Exit(1)
	}

	// Give the loaded catalog to the handlers that serve API requests.
	h := api.NewHandler(c)

	// Register the public routes and connect each one to the code that handles it.
	mux := http.NewServeMux()
	// The router matches the GET path to a handler and runs it.
    // Example: curl localhost:8080/healthz → GET /healthz → api.Healthz
	mux.HandleFunc("GET /healthz", api.Healthz)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /services/{name}", h.ServiceByName)
	
    // Configure the server on port 8080, using mux to connect requests to handler function
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("server starting", "port", server.Addr)

	// Timeouts prevent slow or stuck clients from holding connections open
	// indefinitely while the API waits to read or write data.
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
