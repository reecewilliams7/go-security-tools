package jwk

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/x25519"

	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
)

// OKPJWKCreator creates OKP (Octet Key Pair) JSON Web Keys.
// Supported curves are "Ed25519" and "X25519".
type OKPJWKCreator struct {
	curve string
}

// NewOKPJWKCreator creates a new OKPJWKCreator with the specified curve.
// Supported curve values are "Ed25519" and "X25519".
func NewOKPJWKCreator(curve string) *OKPJWKCreator {
	return &OKPJWKCreator{curve: curve}
}

// Create generates a new OKP JSON Web Key.
func (c *OKPJWKCreator) Create() (*internaljwk.JWKOutput, error) {
	var rawKey any

	switch c.curve {
	case "Ed25519":
		_, privKey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return nil, err
		}
		rawKey = privKey
	case "X25519":
		_, privKey, err := x25519.GenerateKey(rand.Reader)
		if err != nil {
			return nil, err
		}
		rawKey = privKey
	default:
		return nil, fmt.Errorf("unsupported OKP curve: %s", c.curve)
	}

	key, err := jwk.FromRaw(rawKey)
	if err != nil {
		return nil, err
	}

	if err := jwk.AssignKeyID(key); err != nil {
		return nil, err
	}

	return internaljwk.NewJWKOutput(key)
}
