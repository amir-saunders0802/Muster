package main

import (
	"log/slog"
	"net/http"
	"os"

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
	mux.HandleFunc("GET /healthz", api.Healthz)
	mux.HandleFunc("GET /services", h.Services)
	mux.HandleFunc("GET /services/{name}", h.ServiceByName)

	slog.Info("server starting", "port", 8080)

	// Keep the API listening for requests until the server stops or encounters
	// an error that prevents it from continuing.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
