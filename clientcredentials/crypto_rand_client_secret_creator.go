package clientcredentials

import (
	"encoding/base64"
	"math/rand"
)

// CryptoRandClientSecretCreator creates client secrets using cryptographically
// secure random number generation.
type CryptoRandClientSecretCreator struct {
	src rand.Source
}

// NewCryptoRandClientSecretCreator creates a new CryptoRandClientSecretCreator
// using the default crypto/rand source.
func NewCryptoRandClientSecretCreator() *CryptoRandClientSecretCreator {
	cryptoSrc := &cryptoSource{}
	return &CryptoRandClientSecretCreator{
		src: cryptoSrc,
	}
}

// NewCryptoRandClientSecretCreatorWithSource creates a new CryptoRandClientSecretCreator
// with a custom random source.
func NewCryptoRandClientSecretCreatorWithSource(src rand.Source) *CryptoRandClientSecretCreator {
	return &CryptoRandClientSecretCreator{
		src: src,
	}
}

// Create generates a new cryptographically secure client secret.
func (c *CryptoRandClientSecretCreator) Create() (string, error) {
	rnd := rand.New(c.src)
	buff := make([]byte, 32)
	_, err := rnd.Read(buff)
	if err != nil {
		return "", err
	}

	secret := base64.StdEncoding.EncodeToString(buff)
	return secret, nil
}
