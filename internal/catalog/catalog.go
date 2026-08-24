package catalog

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Service defines what one service looks like.
// The tags map YAML/JSON fields to the Go fields.
type Service struct {
	Name        string `yaml:"name" json:"name"`
	Owner       string `yaml:"owner" json:"owner"`
	Repo        string `yaml:"repo" json:"repo"`
	Environment string `yaml:"environment" json:"environment"`
	Tier        int    `yaml:"tier" json:"tier"`
}

// Catalog holds the list of all services.
type Catalog struct {
	Services []Service `yaml:"services" json:"services"`
}

// Load reads the YAML file and returns a Catalog or an error.
func Load(path string) (Catalog, error) {

	// Read the file. data = file contents, err = any error.
	data, err := os.ReadFile(path)
	if err != nil {
		return Catalog{}, err
	}

	// Create an empty Catalog to put the YAML data into.
	var catalog Catalog

	// Convert YAML data into catalog.
	// &catalog lets Unmarshal modify/fill the actual catalog.
	// If it fails, err will NOT be nil.
	if err := yaml.Unmarshal(data, &catalog); err != nil {
		return Catalog{}, err
	}

	// Success: return the populated catalog and no error.
	return catalog, nil
}