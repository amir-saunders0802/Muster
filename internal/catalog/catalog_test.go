package catalog

import "testing"

func TestLoad(t *testing.T) {
	// Load the real catalog file to verify that its YAML structure matches the
	// format expected by the application.
	catalog, err := Load("../../config/services.yaml")
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	// The test also catches missing or unexpectedly added entries in the sample
	// catalog.
	if len(catalog.Services) != 5 {
		t.Errorf("expected 5 services, got %d", len(catalog.Services))
	}
}
