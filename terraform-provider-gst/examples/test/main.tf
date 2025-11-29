terraform {
  required_providers {
    gst = {
      source = "reecewilliams7/gst"
    }
  }
}

provider "gst" {}

# Generate a simple RSA JWK
resource "gst_jwk" "test" {
  algorithm = "rsa"
  key_size  = 2048
}

# Generate client credentials
resource "gst_client_credentials" "test" {}

output "jwk_kid" {
  value = gst_jwk.test.id
}

output "client_id" {
  value = gst_client_credentials.test.client_id
}
