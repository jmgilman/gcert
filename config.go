package main

import (
	"fmt"
	"os"
)

type Config struct {
	Email      string
	PrivateKey []byte
	URI        string
	CFToken    string
}

type KeyReader func(filename string) ([]byte, error)

var environmentVars = [...]string{
	"GCERT_CF_TOKEN",
	"GCERT_USER_EMAIL",
	"GCERT_USER_KEY_FILE",
	"GCERT_USER_URI",
}

func NewConfigFromEnv(reader KeyReader) (*Config, error) {
	// Make sure all environment variables have values
	for _, envVar := range environmentVars {
		if os.Getenv(envVar) == "" {
			return &Config{}, fmt.Errorf("%s environment variable must be set and not empty", envVar)
		}
	}

	keyBytes, err := reader(os.Getenv("GCERT_USER_KEY_FILE"))
	if err != nil {
		return &Config{}, fmt.Errorf("error reading user key at %s: %s", os.Getenv("GCERT_USER_KEY_FILE"), err)
	}

	return &Config{
		Email:      os.Getenv("GCERT_USER_EMAIL"),
		PrivateKey: keyBytes,
		URI:        os.Getenv("GCERT_USER_URI"),
		CFToken:    os.Getenv("GCERT_CF_TOKEN"),
	}, nil
}
