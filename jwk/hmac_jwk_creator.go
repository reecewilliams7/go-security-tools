package jwk

import (
	"crypto/rand"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"

	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
)

// HMACJWKCreator creates symmetric (oct) JSON Web Keys for HMAC algorithms.
// Supported algorithms are "HS256", "HS384", and "HS512".
type HMACJWKCreator struct {
	algorithm string
}

// NewHMACJWKCreator creates a new HMACJWKCreator with the specified algorithm.
func NewHMACJWKCreator(algorithm string) *HMACJWKCreator {
	return &HMACJWKCreator{algorithm: algorithm}
}

// Create generates a new HMAC JSON Web Key.
func (c *HMACJWKCreator) Create() (*internaljwk.JWKOutput, error) {
	var keySize int
	switch c.algorithm {
	case "HS256":
		keySize = 32
	case "HS384":
		keySize = 48
	case "HS512":
		keySize = 64
	default:
		return nil, fmt.Errorf("unsupported HMAC algorithm: %s", c.algorithm)
	}

	keyBytes := make([]byte, keySize)
	if _, err := rand.Read(keyBytes); err != nil {
		return nil, err
	}

	key, err := jwk.FromRaw(keyBytes)
	if err != nil {
		return nil, err
	}

	if err := jwk.AssignKeyID(key); err != nil {
		return nil, err
	}

	return internaljwk.NewJWKOutput(key)
}
