package config

// Config loads and stores the configuration
type Config struct {
	Hostname string `json:"hostname"`
	Database string `json:"database"`
	Port     string `json:"port"`
}
