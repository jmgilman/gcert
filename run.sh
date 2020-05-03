#!/usr/bin/env bash

USAGE="Usage:
  ./run.sh [ADDRESS] [PORT]
"

# Check arguments
if [ -z "$1" ]
then
    echo "$USAGE"
    exit 1
elif [ -z "$2" ]
then
    echo "$USAGE"
    exit 1
fi

# Build image
docker build -t gcert .

# Run server container
docker run -d -p $2:$2 \
    -e GCERT_CF_TOKEN="$(vault kv get --field=api_token secret/cloudflare)" \
    -e GCERT_USER_EMAIL="$(vault kv get --field=email secret/acme/staging)" \
    -e GCERT_USER_KEY="$(vault kv get --field=key secret/acme/staging)" \
    -e GCERT_USER_URI="$(vault kv get --field=uri secret/acme/staging)" \
    -e GCERT_VAULT_ADDRESS="$VAULT_ADDR" \
    -e GCERT_VAULT_ROLE_ID="$(vault kv get --field=role_id secret/approle/gcert)" \
    -e GCERT_VAULT_SECRET_ID="$(vault kv get --field=secret_id secret/approle/gcert)" \
    gcert --address $1 --port $2