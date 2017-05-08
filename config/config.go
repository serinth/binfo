package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config for application
type Config struct {
	BuildServer         string   `json:"buildServer"`
	Projects            []string `json:"projects"`
	RefreshIntervalSecs uint64   `json:"refreshIntervalSecs"`
}

// GetConfig returns system config
func GetConfig(filePath string) (Config, error) {
	raw, err := ioutil.ReadFile(filePath)

	var config Config
	err = json.Unmarshal(raw, &config)

	return config, err
}
