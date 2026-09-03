package catalog

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Service contains the ownership, repository, environment, and importance
// information recorded for one service in the catalog.
type Service struct {
	Name        string `yaml:"name" json:"name"`
	Owner       string `yaml:"owner" json:"owner"`
	Repo        string `yaml:"repo" json:"repo"`
	Environment string `yaml:"environment" json:"environment"`
	Tier        int    `yaml:"tier" json:"tier"`
}

// Catalog is the top-level shape expected in the services YAML file.
type Catalog struct {
	Services []Service `yaml:"services" json:"services"`
}

// Load reads a YAML catalog from path and converts it into the application's
// service data. It returns an error when the file cannot be read or parsed.
func Load(path string) (Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Catalog{}, err
	}

	var catalog Catalog

	if err := yaml.Unmarshal(data, &catalog); err != nil {
		return Catalog{}, err
	}

	return catalog, nil
}
