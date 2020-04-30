package main

import (
	"flag"
	"fmt"
	cservice "github.com/jmgilman/gcert/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const letsEncryptURL = "https://acme-staging-v02.api.letsencrypt.org/directory"

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