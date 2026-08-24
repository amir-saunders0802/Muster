package catalog

import "testing"

func TestLoad(t *testing.T) {
	catalog, err := Load("../../config/services.yaml")
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	if len(catalog.Services) != 5 {
		t.Errorf("expected 5 services, got %d", len(catalog.Services))
	}
}