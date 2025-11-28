package jwk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/lestrrat-go/jwx/v2/jwk"

	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
)

// ECDSAJSONWebKeyCreator creates ECDSA-based JSON Web Keys.
type ECDSAJWKCreator struct {
	curveType string
}

// NewECDSAJWKCreator creates a new ECDSAJWKCreator with the specified curve type.
// Supported curve types are "P256", "P384", and "P521".
func NewECDSAJWKCreator(curveType string) *ECDSAJWKCreator {
	c := ECDSAJWKCreator{curveType: curveType}
	return &c
}

// Create generates a new ECDSA JSON Web Key.
func (c *ECDSAJWKCreator) Create() (*internaljwk.JWKOutput, error) {
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

	jo, err := internaljwk.NewJWKOutput(key)
	if err != nil {
		return nil, err
	}

	return jo, nil
}
