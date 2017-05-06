package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config for application
type Config struct {
	BuildServer string      `json:"buildServer"`
	Credentials credentials `json:"credentials"`
	Projects    []string    `json:"projects"`
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetConfig returns system config
func GetConfig(filePath string) (Config, error) {
	raw, err := ioutil.ReadFile(filePath)

	var config Config
	err = json.Unmarshal(raw, &config)

	return config, err
}
