#!/usr/bin/env bash

export VAULT_TOKEN="$1"
export VAULT_ADDR="$2"

export GCERT_CF_TOKEN="$(./vault kv get --field=api_token secret/cloudflare)"
export GCERT_USER_EMAIL="$(./vault kv get --field=email secret/acme/staging)"
export GCERT_USER_KEY="$(./vault kv get --field=key secret/acme/staging)"
export GCERT_USER_URI="$(./vault kv get --field=uri secret/acme/staging)"
export VAULT_ADDR="$2"
export VAULT_PATH="approle"
export VAULT_ROLE_ID="$(./vault kv get --field=role_id secret/approle/gcert)"
export VAULT_SECRET_ID="$(./vault kv get --field=secret_id secret/approle/gcert)"

./gcert --address $3 --port $4