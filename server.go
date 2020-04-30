package main

import (
	"context"
	cservice "github.com/jmgilman/gcert/proto"
	"log"
)

type certServer struct {
	config *Config
}

func (c *certServer) GetCertificate(ctx context.Context, in *cservice.CertificateRequest) (*cservice.CertificateResponse, error) {
	log.Printf("Received request for %s.%s\n", in.Subdomain, in.Domain)
	cert := &cservice.CertificateResponse{Domain: in.Domain}
	return cert, nil
}
