package jsonWebKeys

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type RsaJsonWebKeyCreator struct {
	bits int
}

func NewRsaJsonWebKeyCreator(bits int) *RsaJsonWebKeyCreator {
	c := RsaJsonWebKeyCreator{bits: bits}
	return &c
}

func (c *RsaJsonWebKeyCreator) Create() (*JsonWebKeyOutput, error) {
	var rawKey interface{}

	rsaKey, err := rsa.GenerateKey(rand.Reader, c.bits)
	if err != nil {
		return nil, err
	}
	rawKey = rsaKey

	key, err := jwk.FromRaw(rawKey)

	if err != nil {
		return nil, err
	}

	jwk.AssignKeyID(key)

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

	jo, err := NewJsonWebKeyOutput(key, string(pemPrivateKey), string(pemPublicKey))
	if err != nil {
		return nil, err
	}

	return jo, nil
}
