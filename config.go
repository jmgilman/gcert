package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

// AppConfig contains the configuration data needed for the gcert server to operate
type AppConfig struct {
	Email      string
	PrivateKey []byte
	URI        string
	CFToken    string
	VaultCfg   VaultConfig
}

// VaultConfig contains the configuration data for authenticating and writing against a Vault server
type VaultConfig struct {
	Address  string
	Path     string
	RoleID   string
	SecretID string
}

// environmentVars contains a list of all required environment variables for the gcert service to operate
var environmentVars = [...]string{
	"GCERT_CF_TOKEN",
	"GCERT_USER_EMAIL",
	"GCERT_USER_KEY",
	"GCERT_USER_URI",
	"VAULT_ADDR",
	"VAULT_PATH",
	"VAULT_ROLE_ID",
	"VAULT_SECRET_ID",
}

// NewConfigFromEnv validates the required environment variables are present (and not empty) and then returns an
// AppConfig configured using those environment variables.
func NewConfigFromEnv() (*AppConfig, error) {
	// Make sure all environment variables have values
	for _, envVar := range environmentVars {
		if os.Getenv(envVar) == "" {
			return &AppConfig{}, fmt.Errorf("%s environment variable must be set and not empty", envVar)
		}
	}

	keyBytes, err := base64.StdEncoding.DecodeString(os.Getenv("GCERT_USER_KEY"))
	if err != nil {
		return &AppConfig{}, fmt.Errorf("error decoding user private key: %s\n", err)
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
