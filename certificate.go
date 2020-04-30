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

func (u *User) GetEmail() string {
	return u.Email
}
func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func RequestCert(endPoint string, domains []string, c *Config) (*cservice.CertificateResponse, error) {
	key, err := certcrypto.ParsePEMPrivateKey(c.PrivateKey)
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}
	user := User{
		Email: c.Email,
		key:   key,
		Registration: &registration.Resource{
			URI: c.URI,
		},
	}

	// Setup a new Lego client
	config := lego.NewConfig(&user)
	config.CADirURL = endPoint
	client, err := lego.NewClient(config)
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}

	// Configure the CloudFlare provider
	providerConfig := cloudflare.NewDefaultConfig()
	providerConfig.AuthEmail = c.Email
	providerConfig.AuthToken = c.CFToken

	provider, err := cloudflare.NewDNSProviderConfig(providerConfig)
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}

	if err := client.Challenge.SetDNS01Provider(provider); err != nil {
		return &cservice.CertificateResponse{}, err
	}

	// Setup and send certificate request
	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}

	log.Printf("Sending certificate request to %s\n...", endPoint)
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		return &cservice.CertificateResponse{}, err
	}

	return &cservice.CertificateResponse{
		Domains:           domains,
		CertURL:           certificates.CertURL,
		CertStableURL:     certificates.CertStableURL,
		PrivateKey:        certificates.PrivateKey,
		Certificate:       certificates.Certificate,
		IssuerCertificate: certificates.IssuerCertificate,
		CSR:               certificates.CSR,
	}, nil
}
