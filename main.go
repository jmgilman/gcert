package main

import (
	"context"
	"flag"
	"fmt"
	cservice "github.com/jmgilman/gcert/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const letsEncryptURL = "https://acme-staging-v02.api.letsencrypt.org/directory"

type certServer struct {
	config *Config
}

type Config struct {
	Endpoint string
	Email string
	Key []byte
	APIKey string
	APIToken string
}

func (c *certServer) GetCertificate(ctx context.Context, in *cservice.CertificateRequest) (*cservice.CertificateResponse, error) {
	cert := &cservice.CertificateResponse{Domain: in.Domain}
	return cert, nil
}

func NewConfigFromEnv() (*Config, error) {
	// Make sure all environment variables have values
	envVars := []string{"USER_EMAIL", "USER_KEY_FILE", "API_KEY", "API_TOKEN"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return &Config{}, fmt.Errorf("%s environment variable must be set", envVar)
		}
	}

	keyPath := os.Getenv("USER_KEY_FILE")
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return &Config{}, fmt.Errorf("cannot find user key at %s: %s", keyPath, err)
	}

	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading user key at %s: %s", keyPath, err)
	}

	return &Config {
		Email: os.Getenv("USER_EMAIL"),
		Key: keyBytes,
		APIKey: os.Getenv("API_KEY"),
		APIToken: os.Getenv("API_TOKEN"),
	}, nil
}

func main() {
	address := flag.String("address", "localhost", "Address to bind to")
	port := flag.Int("port", 8080, "Port to listen on")
	endPoint := flag.String("endpoint", letsEncryptURL, "Endpoint to request a certificate from")
	flag.Parse()

	log.Println("Loading environment...")

	config, err := NewConfigFromEnv()
	if err != nil {
		log.Fatal("Could not load environment:", err)
	}

	config.Endpoint = *endPoint
	server := &certServer{config: config}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	grpcServer := grpc.NewServer()
	cservice.RegisterCertificateServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed starting server:", err)
	}
}