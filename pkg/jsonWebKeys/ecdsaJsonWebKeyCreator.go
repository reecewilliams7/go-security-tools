package jsonWebKeys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type ECDSAJsonWebKeyCreator struct {
	curveType string
}

func NewECDSAJsonWebKeyCreator(curveType string) *ECDSAJsonWebKeyCreator {
	c := ECDSAJsonWebKeyCreator{curveType: curveType}
	return &c
}

func (c *ECDSAJsonWebKeyCreator) Create() (*JsonWebKeyOutput, error) {
	var curve elliptic.Curve
	switch c.curveType {
	case "P256":
		curve = elliptic.P256()
	case "P384":
		curve = elliptic.P384()
	case "P521":
		curve = elliptic.P521()
	default:
		curve = elliptic.P256()
	}

	ecdsaKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	key, err := jwk.FromRaw(ecdsaKey)

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
