// package main contains the core logic for the gcert server
//go:generate protoc -I proto/ proto/certificate_service.proto --go_out=plugins=grpc:proto --go_opt=paths=source_relative
package main

import (
	"flag"
	"fmt"
	cservice "github.com/jmgilman/gcert/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	address := flag.String("address", "localhost", "Address to bind to")
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	log.Println("Loading environment...")

	config, err := NewConfigFromEnv()
	if err != nil {
		log.Fatal("Could not load environment:", err)
	}

	server := &certServer{config: config}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	log.Printf("Listening on %s:%d...\n", *address, *port)
	grpcServer := grpc.NewServer()
	cservice.RegisterCertificateServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed starting server:", err)
	}
}
