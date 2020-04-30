package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Endpoint string
	Email string
	Key []byte
	APIKey string
	APIToken string
}

func NewConfigFromEnv() (*Config, error) {
	// Make sure all environment variables have values
	envVars := []string{"USER_EMAIL", "USER_KEY_FILE", "API_KEY", "API_TOKEN"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return &Config{}, fmt.Errorf("%s environment variable must be set and not empty", envVar)
		}
	}

	keyPath := os.Getenv("USER_KEY_FILE")
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return &Config{}, fmt.Errorf("cannot find user key at %s: %s", keyPath, err)
	}

	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading user key at %s: %s", keyPath, err)
	}

	return &Config {
		Email: os.Getenv("USER_EMAIL"),
		Key: keyBytes,
		APIKey: os.Getenv("API_KEY"),
		APIToken: os.Getenv("API_TOKEN"),
	}, nil
}