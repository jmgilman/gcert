package main

import (
	"crypto"
	"github.com/go-acme/lego/v3/certcrypto"
	"github.com/go-acme/lego/v3/certificate"
	"github.com/go-acme/lego/v3/lego"
	"github.com/go-acme/lego/v3/providers/dns/cloudflare"
	"github.com/go-acme/lego/v3/registration"
	cservice "github.com/jmgilman/gcert/proto"
	"log"
)

type User struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

type KeyParser func([]byte) (crypto.PrivateKey, error)

func (u *User) GetEmail() string {
	return u.Email
}
func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func NewUserFromConfig(c *AppConfig, keyParser KeyParser) (*User, error) {
	key, err := keyParser(c.PrivateKey)
	if err != nil {
		return &User{}, err
	}

	return &User{
		Email: c.Email,
		key:   key,
		Registration: &registration.Resource{
			URI: c.URI,
		},
	}, nil
}

func NewClientConfigFromUser(u *User, endpoint string) *lego.Config {
	config := lego.NewConfig(u)
	config.CADirURL = endpoint
	return config
}

func NewProviderConfigFromConfig(c *AppConfig) (providerConfig *cloudflare.Config) {
	providerConfig = cloudflare.NewDefaultConfig()
	providerConfig.AuthEmail = c.Email
	providerConfig.AuthToken = c.CFToken

	return
}

func GetRequest(domains []string) certificate.ObtainRequest {
	return certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
}

func RequestCert(endpoint string, domains []string, c *AppConfig) (*cservice.CertificateResponse, error) {
	// Setup a new Lego client
	user, err := NewUserFromConfig(c, certcrypto.ParsePEMPrivateKey)
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}
	client, err := lego.NewClient(NewClientConfigFromUser(user, endpoint))
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}

	// Configure the CloudFlare provider
	provider, err := cloudflare.NewDNSProviderConfig(NewProviderConfigFromConfig(c))
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}
	if err := client.Challenge.SetDNS01Provider(provider); err != nil {
		return &cservice.CertificateResponse{}, err
	}

	log.Printf("Sending certificate request to %s\n", endpoint)
	certificates, err := client.Certificate.Obtain(GetRequest(domains))
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}

	log.Println("Authenticating against vault at ", c.VaultCfg.Address)
	vaultClient, err := NewDefaultClient(c)
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}
	if err := AuthenticateClient(vaultClient, c); err != nil {
		return &cservice.CertificateResponse{}, err
	}

	log.Println("Writing certificates to Vault")
	paths, err := WriteCertificate(vaultClient, certificates, domains)
	if err != nil {
		return &cservice.CertificateResponse{
			VaultPaths: paths,
			Success:    false,
		}, err
	}

	log.Println("Done!")
	return &cservice.CertificateResponse{
		VaultPaths: paths,
		Success:    true,
	}, nil
}
