# Example: Writing JWK to a file using local_file

terraform {
  required_providers {
    gst = {
      source = "reecewilliams7/gst"
    }
    local = {
      source  = "hashicorp/local"
      version = "~> 2.0"
    }
  }
}

provider "gst" {}
provider "local" {}

# Generate an RSA JWK
resource "gst_jwk" "signing_key" {
  algorithm = "rsa"
  key_size  = 2048
}

# Write the private key to a file
resource "local_file" "private_key" {
  content         = gst_jwk.signing_key.pem_private_key
  filename        = "${path.module}/private_key.pem"
  file_permission = "0600"
}

# Write the public key to a file
resource "local_file" "public_key" {
  content         = gst_jwk.signing_key.pem_public_key
  filename        = "${path.module}/public_key.pem"
  file_permission = "0644"
}

# Write the JWK to a file
resource "local_file" "jwk" {
  content         = gst_jwk.signing_key.jwk_public_string
  filename        = "${path.module}/jwk.json"
  file_permission = "0644"
}

output "key_id" {
  value = gst_jwk.signing_key.id
}
