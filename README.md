# gcert
> A small gRPC server which requests new Let's Encrypt certificates for the given domains and writes them to Vault

The gcert service is used for automating the retrieval of Let's Encrypt SSL certificates for the Gilman Lab (glab). It
receives its configuration from the environment and will respond to incoming RPC invocations by requesting certificates
for the given list of domains and then writing the resulting certificates to a static location in the configured Vault
secret store. 

This ultimately results in the automation of providing SSL certs to the various glab services. A cofigured cron job
will periodically request new certificates to glab domains and downstream apps will retrieve the certificates using
the Vault API - thus decoupling apps that need certificates from the actual certificate retrieval process.

## Usage

The gcert service expects the following environment variables to be defined:

* `GCERT_CF_TOKEN`: A CloudFlare API token with DNS:Edit and Zone:Read permissions for all glab domains
* `GCERT_USER_EMAIL`: The email address of a registered user
* `GCERT_USER_KEY`: The private key (encoded with base64) of the registered user
* `GCERT_USER_URI`: The unique URI of the registered user
* `VAULT_ADDR`: The URL address of the Vault server to use
* `VAULT_PATH`: The mount name for the AppRole auth method (i.e approle)
* `VAULT_ROLE_ID`: The role ID to use for authenticating to the Vault server
* `VAULT_SECRET_ID`: The secret ID to use for authenticating to the Vault server 

Once the above has been configured, the server can be started with:

```
$> ./gcert --address 0.0.0.0 --port 8080
```

The server will begin listening for calls to the RPC methods defined in `proto/certificate_service.proto`.

## Retrieving Certificates

The default base path for the gcert server is `secret/ssl/[domain]`. For example, if a request is received to generate
certificates for test.example.com, then the certificates will be written to `secret/ssl/test.example.com`. The following
certificate details are written:

* `cert_url`: The URL for the certificate
* `cert_stable_url`: The stable URL for the certificate
* `private_key`: The private key for the certificate (base64 encoded)
* `certificate`: The contents of the certificate (base64 encoded)
* `issuer_certificate`: The certificate of the issuing CA (base64 encoded)

It's possible to request a certificate which covers multiple domains, in which case the above details are written to
Vault using each domain. This results in the same certificates being written multiple times, but helps remove ambiguity 
as to which domains the certificates cover. For example, a request for test1.example.com and test2.example.com will
result in the certificate details being written to `secret/ssl/test1.example.com` and `secret/ssl/test2.example.com`.

## Docker

The repo comes with all the necessary files to build and start the gcert process using Docker. All that is required is
a valid Vault token and an associated server. To build and start the container:

```
$> ./run.sh
```

The run script assumes the default `$VAULT_TOKEN` and `$VAULT_ADDR` environment variables are already set. The service 
will automatically be brought up in a container (named `gcert`) and begin listening on `0.0.0.0:8080`.