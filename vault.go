package main

import (
	"encoding/base64"
	"fmt"
	"github.com/go-acme/lego/v3/certificate"
	"github.com/hashicorp/vault/api"
)

const basePath = "secret/ssl/"

func NewDefaultClient(c *AppConfig) (*api.Client, error) {
	return api.NewClient(nil)
}

func AuthenticateClient(c *api.Client, config *AppConfig) error {
	data := map[string]interface{}{
		"role_id":   config.VaultCfg.RoleID,
		"secret_id": config.VaultCfg.SecretID,
	}
	path := fmt.Sprintf("auth/%s/login", config.VaultCfg.Path)

	secret, err := c.Logical().Write(path, data)
	if err != nil {
		return err
	}
	if secret.Auth == nil {
		return fmt.Errorf("authentication returned an empty token")
	}

	c.SetToken(secret.Auth.ClientToken)
	return nil
}

func WriteCertificate(c *api.Client, r *certificate.Resource, domains []string) ([]string, error) {
	var paths []string
	for _, domain := range domains {
		path := basePath + domain
		data := map[string]interface{}{
			"cert_url":           r.CertURL,
			"cert_stable_url":    r.CertStableURL,
			"private_key":        base64.StdEncoding.EncodeToString(r.PrivateKey),
			"certificate":        base64.StdEncoding.EncodeToString(r.Certificate),
			"issuer_certificate": base64.StdEncoding.EncodeToString(r.IssuerCertificate),
		}

		_, err := c.Logical().Write(path, data)
		if err != nil {
			return paths, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}
