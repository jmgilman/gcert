package main

import (
	"context"
	cservice "github.com/jmgilman/gcert/proto"
	"log"
)

type certServer struct {
	config *AppConfig
}

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

var EndPoints = [...]string{
	"https://acme-v02.api.letsencrypt.org/directory",
	"https://acme-staging-v02.api.letsencrypt.org/directory",
}

// GetCertificate implements the RPC interface for the GetCertificate function
func (c *certServer) GetCertificate(ctx context.Context, in *cservice.CertificateRequest) (resp *cservice.CertificateResponse, err error) {
	log.Printf("Received certificate request for %s\n", in.Domains)
	resp, err = RequestCert(EndPoints[in.Endpoint], in.Domains, c.config)
	if err != nil {
		log.Println("Error requesting new certificates: ", err)
	}
	return
}
