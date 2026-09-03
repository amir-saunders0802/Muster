package catalog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	// Use a temporary fixture so this test checks the loader independently of
	// how many services happen to be in the application's real configuration.
	path := filepath.Join(t.TempDir(), "services.yaml")
	data := []byte(`services:
  - name: test-api
    owner: test-team
    repo: github.com/example/test-api
    environment: test
    tier: 2
`)
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("could not create test catalog: %v", err)
	}

	catalog, err := Load(path)
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	if len(catalog.Services) != 1 {
		t.Fatalf("expected 1 service, got %d", len(catalog.Services))
	}

	service := catalog.Services[0]
	if service.Name != "test-api" || service.Owner != "test-team" ||
		service.Repo != "github.com/example/test-api" ||
		service.Environment != "test" || service.Tier != 2 {
		errorMessage := "loaded service does not match the test catalog: %+v"
		t.Errorf(errorMessage, service)
	}
}

func TestLoadMissingFile(t *testing.T) {
	// A missing file should be reported instead of being treated as an empty
	// catalog.
	catalog, err := Load(filepath.Join(t.TempDir(), "missing.yaml"))
	if err == nil {
		t.Fatal("Load() expected an error for a missing file")
	}
	if len(catalog.Services) != 0 {
		t.Errorf("expected an empty catalog, got %d services", len(catalog.Services))
	}
}

func TestLoadInvalidYAML(t *testing.T) {
	// A file that cannot be parsed should return an error and no partial data.
	path := filepath.Join(t.TempDir(), "invalid.yaml")
	if err := os.WriteFile(path, []byte("services: [invalid"), 0o600); err != nil {
		t.Fatalf("could not create invalid test catalog: %v", err)
	}

	catalog, err := Load(path)
	if err == nil {
		t.Fatal("Load() expected an error for invalid YAML")
	}
	if len(catalog.Services) != 0 {
		t.Errorf("expected an empty catalog, got %d services", len(catalog.Services))
	}
}
