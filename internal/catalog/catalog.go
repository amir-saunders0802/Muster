package catalog


type Service struct {
    Name        string `yaml:"name" json:"name"`
    Owner       string `yaml:"owner" json:"owner"`
    Repo        string `yaml:"repo" json:"repo"`
    Environment string `yaml:"environment" json:"environment"`
    Tier        int    `yaml:"tier" json:"tier"`
}


type Catalog struct {
    Services []Service `yaml:"services" json:"services"`
}