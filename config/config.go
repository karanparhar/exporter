package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Urls []string `json:"urls"`
}

func ReadConfigFile(f string) (*Config, error) {

	file, err := os.Open(f)

	if err != nil {
		return nil, err
	}

	var config *Config

	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
