package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/amir-saunders0802/Muster/internal/api"
	"github.com/amir-saunders0802/Muster/internal/catalog"
)

func main() {
	// Load the YAML catalog.
	c, err := catalog.Load("config/services.yaml")
	if err != nil {
		slog.Error("failed to load catalog", "error", err)
		os.Exit(1)
	}

	// Create a handler that contains the loaded catalog.
	h := api.NewHandler(c)

	// Create the HTTP router.
	mux := http.NewServeMux()

	// Connect URLs to handler functions.
	mux.HandleFunc("GET /healthz", api.Healthz)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /services/{name}", h.ServiceByName)

	// Log that the server is starting.
	slog.Info("server starting", "port", 8080)

	// Start the server and listen on port 8080.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}