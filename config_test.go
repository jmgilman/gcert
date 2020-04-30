package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func MockReader(filepath string) ([]byte, error) {
	return []byte{}, nil
}

func TestNewConfigFromEnv(t *testing.T) {
	// Setup environment
	for _, envVar := range environmentVars {
		if err := os.Setenv(envVar, "testing"); err != nil {
			t.Fatal(err)
		}
	}

	result, err := NewConfigFromEnv(MockReader)
	assert.Empty(t, err)
	assert.Equal(t, result.CFToken, "testing")
	assert.Equal(t, result.PrivateKey, []byte{})
	assert.Equal(t, result.Email, "testing")
	assert.Equal(t, result.URI, "testing")
	assert.Equal(t, result.VaultCfg.Address, "testing")
	assert.Equal(t, result.VaultCfg.Path, "testing")
	assert.Equal(t, result.VaultCfg.RoleID, "testing")
	assert.Equal(t, result.VaultCfg.SecretID, "testing")
}
