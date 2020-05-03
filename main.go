// package main contains the core logic for the gcert server
//go:generate protoc -I proto/ proto/certificate_service.proto --go_out=plugins=grpc:proto --go_opt=paths=source_relative
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	cservice "github.com/jmgilman/gcert/proto"
	"github.com/peterbourgon/ff/v3"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

// main reads configuration data from the environment and starts the RPC server
func main() {
	log.Println("Initializing...")
	fs := flag.NewFlagSet("gcert", flag.ExitOnError)
	var (
		address       = fs.String("address", "localhost", "Address to bind to")
		port          = fs.Int("port", 8080, "Port to listen on")
		cfToken       = fs.String("cf-token", "", "CloudFlare API token")
		userEmail     = fs.String("user-email", "", "Let's Encrypt account email")
		userKey       = fs.String("user-key", "", "Let's Encrypt account private key (base64 encoded)")
		userUri       = fs.String("user-uri", "", "Let's Encrypt account URI")
		vaultAddress  = fs.String("vault-address", os.Getenv("VAULT_ADDR"), "Address to Vault server")
		vaultPath     = fs.String("vault-path", "approle", "Path to AppRole backend")
		vaultRoleID   = fs.String("vault-role-id", "", "Vault AppRole role ID")
		vaultSecretID = fs.String("vault-secret-id", "", "Vault AppRole secret ID")
	)
	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("GCERT")); err != nil {
		log.Fatal("Could not parse flags: ", err)
	}

	flagMap := map[string]string{
		"cf-token":        *cfToken,
		"user-email":      *userEmail,
		"user-key":        *userKey,
		"user-uri":        *userUri,
		"vault-address":   *vaultAddress,
		"vault-path":      *vaultPath,
		"vault-role-id":   *vaultRoleID,
		"vault-secret-id": *vaultSecretID,
	}

	if err := notEmpty(flagMap); err != nil {
		log.Fatal(err)
	}

	key, err := base64.StdEncoding.DecodeString(*userKey)
	if err != nil {
		log.Fatal("Could not parse user key as base64: ", err)
	}

	config := &AppConfig{
		Email:      *userEmail,
		PrivateKey: key,
		URI:        *userUri,
		CFToken:    *cfToken,
		VaultCfg: VaultConfig{
			Address:  *vaultAddress,
			Path:     *vaultPath,
			RoleID:   *vaultRoleID,
			SecretID: *vaultSecretID,
		},
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

func notEmpty(flags map[string]string) error {
	for k, v := range flags {
		if v == "" {
			return fmt.Errorf("%s must be defined and not empty", k)
		}
	}
	return nil
}
