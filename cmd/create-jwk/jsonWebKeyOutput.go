package main

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JsonWebKeyOutput struct {
	JsonWebKey       *jwk.Key
	JsonWebKeyString string
	Base64JsonWebKey string
	RsaPublicKey     string
	RsaPrivateKey    string
}
