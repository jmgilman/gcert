package main

import (
	"crypto"
	"crypto/rsa"
	"github.com/go-acme/lego/v3/registration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MockParser(key []byte) (crypto.PrivateKey, error) {
	return rsa.PrivateKey{}, nil
}

func GetTestConfig(t *testing.T) *Config {
	t.Helper()
	return &Config{
		Email:      "testing",
		PrivateKey: []byte{},
		URI:        "testing",
		CFToken:    "testing",
	}
}

func GetTestUser(t *testing.T) *User {
	config := GetTestConfig(t)
	return &User{
		Email: config.Email,
		key:   rsa.PrivateKey{},
		Registration: &registration.Resource{
			URI: config.URI,
		},
	}
}

func TestNewUserFromConfig(t *testing.T) {
	config := GetTestConfig(t)

	user, err := NewUserFromConfig(config, MockParser)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, "testing")
	assert.Equal(t, user.key, rsa.PrivateKey{})
	assert.Equal(t, user.Registration.URI, "testing")
}

func TestNewClientConfigFromUser(t *testing.T) {
	config := NewClientConfigFromUser(GetTestUser(t), "testing")
	assert.Equal(t, config.CADirURL, "testing")
}

func TestNewProviderConfigFromConfig(t *testing.T) {
	config := NewProviderConfigFromConfig(GetTestConfig(t))
	assert.Equal(t, config.AuthEmail, "testing")
	assert.Equal(t, config.AuthToken, "testing")
}
