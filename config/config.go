package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config loads and stores the configuration
type Config struct {
	Hostname string `json:"hostname"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

// ReadConfig function helps to get the config back
func ReadConfig(configFile string) Config {
	var cfg Config
	fileData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Not able to read file")
	}
	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		log.Fatalf("Not able to un marshal the config")
	}
	return cfg
}
