// config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// PortService defines a service on a specific port.
type PortService struct {
	Port    int    `json:"port"`
	Service string `json:"service"` // "ssh", "http", "mysql", "generic"
	Banner  string `json:"banner"`  // Custom banner for the service.
}

// Config holds the entire configuration.
type Config struct {
	Services            []PortService `json:"services"`
	TLSCertFile         string        `json:"tlsCertFile"`
	TLSKeyFile          string        `json:"tlsKeyFile"`
	DashboardPort       int           `json:"dashboardPort"`
	SuspiciousThreshold int           `json:"suspiciousThreshold"`
}

// LoadConfig reads the configuration from a JSON file.
func LoadConfig(path string) (Config, error) {
	var cfg Config

	file, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("could not open config file %s: %v", path, err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return cfg, fmt.Errorf("error reading config file: %v", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("error parsing config file: %v", err)
	}
	return cfg, nil
}
