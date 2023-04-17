package jsonWebKeys

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type jsonWebKeyCreator struct{}

func NewJsonWebKeyCreator() *jsonWebKeyCreator {
	c := jsonWebKeyCreator{}
	return &c
}

func (*jsonWebKeyCreator) Create() (*jsonWebKeyOutput, error) {
	var rawKey interface{}
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	rawKey = rsaKey

	key, err := jwk.FromRaw(rawKey)

	if err != nil {
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

	jo, err := NewJsonWebKeyOutput(&key, string(pemPrivateKey), string(pemPublicKey))
	if err != nil {
		return nil, err
	}

	return jo, nil
}
