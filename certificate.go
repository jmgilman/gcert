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

// User represents a registered ACME user
type User struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

// KeyParser is a function which parses a byte slice and returns a private key
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

// NewUserFromConfig returns a User configured from the given AppConfig. The KeyParser is used to parse the []byte key
// provided by the AppConfig.
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

// NewClientConfigFromUser returns a lego.Config configured with the given user and HTTP(S) endpoint
func NewClientConfigFromUser(u *User, endpoint string) *lego.Config {
	config := lego.NewConfig(u)
	config.CADirURL = endpoint
	return config
}

// NewCloudFlareConfig returns a CloudFlare provider configuration using the given AppConfig
func NewCloudFlareConfig(c *AppConfig) (providerConfig *cloudflare.Config) {
	providerConfig = cloudflare.NewDefaultConfig()
	providerConfig.AuthEmail = c.Email
	providerConfig.AuthToken = c.CFToken

	return
}

// GetRequest returns a certificate request for the given list of domains
func GetRequest(domains []string) certificate.ObtainRequest {
	return certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
}

// RequestCert configures a Lego client, provider, and request which it uses to send a request to the given endpoint for
// the given list of domains. It automatically performs a DNS01 authentication test using CloudFlare as the provider. If
// the test is successful, it writes the resulting certificates to the configured Vault instance and returns whether the
// attempt was successful and any Vault paths written to.
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
	provider, err := cloudflare.NewDNSProviderConfig(NewCloudFlareConfig(c))
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
