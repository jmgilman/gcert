package main

import (
	"context"
	"flag"
	"fmt"
	cservice "github.com/jmgilman/gcert/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type certServer struct {

}

func (c *certServer) GetCertificate(ctx context.Context, in *cservice.CertificateRequest) (*cservice.CertificateResponse, error) {
	cert := &cservice.CertificateResponse{Domain: in.Domain}
	return cert, nil
}

func main() {
	fmt.Println("Hello grpc!")

	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	cservice.RegisterCertificateServiceServer(grpcServer, &certServer{})
	grpcServer.Serve(lis)
}
