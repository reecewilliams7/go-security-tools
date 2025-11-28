package jwk

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

// RSAJSONWebKeyCreator creates RSA-based JSON Web Keys.
type RSAJSONWebKeyCreator struct {
	bits int
}

// NewRSAJSONWebKeyCreator creates a new RSAJSONWebKeyCreator with the specified key size in bits.
func NewRSAJSONWebKeyCreator(bits int) *RSAJSONWebKeyCreator {
	c := RSAJSONWebKeyCreator{bits: bits}
	return &c
}

// Create generates a new RSA JSON Web Key.
func (c *RSAJSONWebKeyCreator) Create() (*JSONWebKeyOutput, error) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, c.bits)
	if err != nil {
		return nil, err
	}

	key, err := jwk.FromRaw(rsaKey)
	if err != nil {
		return nil, err
	}

	if err := jwk.AssignKeyID(key); err != nil {
		return nil, err
	}

	pemPrivateKey, err := jwk.EncodePEM(key)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwk.PublicKeyOf(key)
	if err != nil {
		return nil, err
	}

	pemPublicKey, err := jwk.EncodePEM(publicKey)
	if err != nil {
		return nil, err
	}

	jo, err := NewJSONWebKeyOutput(key, string(pemPrivateKey), string(pemPublicKey))
	if err != nil {
		return nil, err
	}

	return jo, nil
}
