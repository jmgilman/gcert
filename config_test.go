package main

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewConfigFromEnv(t *testing.T) {
	// Setup environment
	for _, envVar := range environmentVars {
		if err := os.Setenv(envVar, "testing"); err != nil {
			t.Fatal(err)
		}
	}

	// Setup fake key
	fakeKey := base64.StdEncoding.EncodeToString([]byte("testing"))
	if err := os.Setenv("GCERT_USER_KEY", fakeKey); err != nil {
		t.Fatal(err)
	}

	result, err := NewConfigFromEnv()
	assert.Empty(t, err)
	assert.Equal(t, result.CFToken, "testing")
	assert.Equal(t, result.PrivateKey, []byte("testing"))
	assert.Equal(t, result.Email, "testing")
	assert.Equal(t, result.URI, "testing")
	assert.Equal(t, result.VaultCfg.Address, "testing")
	assert.Equal(t, result.VaultCfg.Path, "testing")
	assert.Equal(t, result.VaultCfg.RoleID, "testing")
	assert.Equal(t, result.VaultCfg.SecretID, "testing")
}
