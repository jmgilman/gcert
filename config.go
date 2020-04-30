package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Email      string
	PrivateKey []byte
	URI        string
	CFToken    string
}

var environmentVars = [...]string{
	"GCERT_CF_TOKEN",
	"GCERT_USER_EMAIL",
	"GCERT_USER_KEY_FILE",
	"GCERT_USER_URI",
}

func NewConfigFromEnv() (*Config, error) {
	// Make sure all environment variables have values
	for _, envVar := range environmentVars {
		if os.Getenv(envVar) == "" {
			return &Config{}, fmt.Errorf("%s environment variable must be set and not empty", envVar)
		}
	}

	keyPath := os.Getenv("GCERT_USER_KEY_FILE")
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return &Config{}, fmt.Errorf("cannot find user key at %s: %s", keyPath, err)
	}

	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading user key at %s: %s", keyPath, err)
	}

	return &Config{
		Email:      os.Getenv("GCERT_USER_EMAIL"),
		PrivateKey: keyBytes,
		URI:        os.Getenv("GCERT_USER_URI"),
		CFToken:    os.Getenv("GCERT_CF_TOKEN"),
	}, nil
}
