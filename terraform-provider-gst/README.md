# Terraform Provider GST (Go Security Tools)

A Terraform provider for generating security credentials using the Go Security Tools library.

## Features

This provider enables you to generate:

- **JSON Web Keys (JWK)**: RSA and ECDSA keys for cryptographic operations
- **OAuth2 Client Credentials**: Client ID and Client Secret pairs for authentication

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21 (to build the provider)

## Building The Provider

```bash
cd terraform-provider-gst
go build -o terraform-provider-gst
```

## Installing The Provider Locally

For local testing, you can use Terraform's development overrides. Create or modify `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "reecewilliams7/gst" = "/workspaces/go-security-tools/terraform-provider-gst"
  }
  direct {}
}
```

Then build the provider:

```bash
go build -o terraform-provider-gst
```

## Using the Provider

```terraform
terraform {
  required_providers {
    gst = {
      source = "reecewilliams7/gst"
    }
  }
}

provider "gst" {}

# Generate an RSA JWK
resource "gst_jwk" "rsa_key" {
  algorithm = "rsa"
  key_size  = 2048
}

# Generate an ECDSA JWK
resource "gst_jwk" "ecdsa_key" {
  algorithm  = "ecdsa"
  curve_type = "P256"
}

# Generate OAuth2 client credentials
resource "gst_client_credentials" "oauth_client" {}

# Output the values (be careful with sensitive data!)
output "client_id" {
  value = gst_client_credentials.oauth_client.client_id
}

output "jwk_public_key" {
  value = gst_jwk.rsa_key.jwk_public_string
}
```

## Resources

### gst_jwk

Generates a JSON Web Key (JWK) for cryptographic operations.

**Arguments:**

- `algorithm` - (Optional) The algorithm type. Supported values: `rsa`, `ecdsa`. Default: `rsa`
- `key_size` - (Optional) Key size in bits for RSA keys. Default: `2048`. Common values: 2048, 3072, 4096
- `curve_type` - (Optional) Elliptic curve type for ECDSA keys. Supported values: `P256`, `P384`, `P521`. Default: `P256`

**Attributes:**

- `id` - The Key ID (kid) from the JWK
- `jwk_string` - Complete JWK as JSON string (sensitive)
- `jwk_public_string` - Public portion of the JWK as JSON string
- `base64_jwk` - Base64-encoded JWK (sensitive)
- `pem_public_key` - Public key in PEM format
- `pem_private_key` - Private key in PEM format (sensitive)

### gst_client_credentials

Generates OAuth2 client credentials.

**Attributes:**

- `id` - Unique identifier (same as client_id)
- `client_id` - Generated client ID (UUID format)
- `client_secret` - Generated client secret (sensitive)

## Development

### Building

```bash
go build -o terraform-provider-gst
```

### Testing

```bash
go test ./...
```

### Generating Documentation

```bash
go generate ./...
```

## Examples

See the [examples](examples/) directory for complete usage examples.

## License

See the LICENSE file in the repository root.
