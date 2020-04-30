package main

import (
	"context"
	cservice "github.com/jmgilman/gcert/proto"
	"log"
)

type certServer struct {
	config *Config
}

var EndPoints = [...]string{
	"https://acme-v02.api.letsencrypt.org/directory",
	"https://acme-staging-v02.api.letsencrypt.org/directory",
}

func (c *certServer) GetCertificate(ctx context.Context, in *cservice.CertificateRequest) (*cservice.CertificateResponse, error) {
	log.Printf("Received certificate request for %s\n", in.Domains)
	return RequestCert(EndPoints[in.Endpoint], in.Domains, c.config)
}
