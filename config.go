package main

import (
	"fmt"
	"os"
)

type AppConfig struct {
	Email      string
	PrivateKey []byte
	URI        string
	CFToken    string
	VaultCfg   VaultConfig
}

type VaultConfig struct {
	Address  string
	Path     string
	RoleID   string
	SecretID string
}

type KeyReader func(filename string) ([]byte, error)

var environmentVars = [...]string{
	"GCERT_CF_TOKEN",
	"GCERT_USER_EMAIL",
	"GCERT_USER_KEY_FILE",
	"GCERT_USER_URI",
	"VAULT_ADDR",
	"VAULT_PATH",
	"VAULT_ROLE_ID",
	"VAULT_SECRET_ID",
}

func NewConfigFromEnv(reader KeyReader) (*AppConfig, error) {
	// Make sure all environment variables have values
	for _, envVar := range environmentVars {
		if os.Getenv(envVar) == "" {
			return &AppConfig{}, fmt.Errorf("%s environment variable must be set and not empty", envVar)
		}
	}

	keyBytes, err := reader(os.Getenv("GCERT_USER_KEY_FILE"))
	if err != nil {
		return &AppConfig{}, fmt.Errorf("error reading user key at %s: %s", os.Getenv("GCERT_USER_KEY_FILE"), err)
	}

	return &AppConfig{
		Email:      os.Getenv("GCERT_USER_EMAIL"),
		PrivateKey: keyBytes,
		URI:        os.Getenv("GCERT_USER_URI"),
		CFToken:    os.Getenv("GCERT_CF_TOKEN"),
		VaultCfg: VaultConfig{
			Address:  os.Getenv("VAULT_ADDR"),
			Path:     os.Getenv("VAULT_PATH"),
			RoleID:   os.Getenv("VAULT_ROLE_ID"),
			SecretID: os.Getenv("VAULT_SECRET_ID"),
		},
	}, nil
}
