package jwk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

// ECDSAJSONWebKeyCreator creates ECDSA-based JSON Web Keys.
type ECDSAJSONWebKeyCreator struct {
	curveType string
}

// NewECDSAJSONWebKeyCreator creates a new ECDSAJSONWebKeyCreator with the specified curve type.
// Supported curve types are "P256", "P384", and "P521".
func NewECDSAJSONWebKeyCreator(curveType string) *ECDSAJSONWebKeyCreator {
	c := ECDSAJSONWebKeyCreator{curveType: curveType}
	return &c
}

// Create generates a new ECDSA JSON Web Key.
func (c *ECDSAJSONWebKeyCreator) Create() (*JSONWebKeyOutput, error) {
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
