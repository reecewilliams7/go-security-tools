package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// verifierBytes is the number of random bytes used to generate the code verifier.
// 64 bytes → 86 base64url characters, within the RFC 7636 range of 43–128.
const verifierBytes = 64

// S256Creator creates PKCE pairs using the S256 (SHA-256) method.
type S256Creator struct{}

// NewS256Creator returns a new S256Creator.
func NewS256Creator() *S256Creator {
	return &S256Creator{}
}

// Create generates a new PKCE code verifier and its SHA-256 code challenge.
func (c *S256Creator) Create() (*PKCEPair, error) {
	raw := make([]byte, verifierBytes)
	if _, err := rand.Read(raw); err != nil {
		return nil, err
	}

	verifier := base64.RawURLEncoding.EncodeToString(raw)

	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])

	return &PKCEPair{
		CodeVerifier:  verifier,
		CodeChallenge: challenge,
		Method:        "S256",
	}, nil
}
