terraform {
  required_providers {
    gst = {
      source = "reecewilliams7/gst"
    }
  }
}

provider "gst" {}

# Generate an RSA JWK with 2048-bit key
resource "gst_jwk" "rsa_2048" {
  algorithm = "rsa"
  key_size  = 2048
}

# Generate an RSA JWK with 4096-bit key
resource "gst_jwk" "rsa_4096" {
  algorithm = "rsa"
  key_size  = 4096
}

# Generate an ECDSA JWK with P-256 curve
resource "gst_jwk" "ecdsa_p256" {
  algorithm  = "ecdsa"
  curve_type = "P256"
}

# Generate an ECDSA JWK with P-384 curve
resource "gst_jwk" "ecdsa_p384" {
  algorithm  = "ecdsa"
  curve_type = "P384"
}

# Generate OAuth2 client credentials
resource "gst_client_credentials" "oauth_client" {}

# Outputs
output "rsa_2048_kid" {
  description = "Key ID of the RSA 2048 JWK"
  value       = gst_jwk.rsa_2048.id
}

output "rsa_2048_public_key" {
  description = "Public key portion of the RSA 2048 JWK"
  value       = gst_jwk.rsa_2048.jwk_public_string
}

output "ecdsa_p256_public_key" {
  description = "Public key portion of the ECDSA P-256 JWK"
  value       = gst_jwk.ecdsa_p256.jwk_public_string
}

output "client_id" {
  description = "OAuth2 Client ID"
  value       = gst_client_credentials.oauth_client.client_id
}

output "client_secret" {
  description = "OAuth2 Client Secret"
  value       = gst_client_credentials.oauth_client.client_secret
  sensitive   = true
}
