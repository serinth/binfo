package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	cwd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	_, err2 := GetConfig(cwd + "/config.json")
	if err2 != nil {
		t.Error(err2)
	}
}
