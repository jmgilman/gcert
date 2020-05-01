#!/usr/bin/env bash

# Build image
docker build -t gcert .

# Run server
docker run -d -p 8080:8080 --name gcert gcert $(echo $VAULT_TOKEN) $(echo $VAULT_ADDR) 0.0.0.0 8080