package internal

import (
	"encoding/json"
	"os"
)

// Config represents the mapping of file patterns to actions.
type Config map[string]map[string][]string

// LoadConfig loads the configuration from a JSON file.
func LoadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
