package clientcredentials

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
)

const (
	// SecretEncodingBase64 encodes the secret as standard Base64.
	SecretEncodingBase64 = "base64"
	// SecretEncodingBase64URL encodes the secret as URL-safe Base64 (no padding).
	SecretEncodingBase64URL = "base64url"
	// SecretEncodingHex encodes the secret as a lowercase hex string.
	SecretEncodingHex = "hex"

	defaultSecretLength   = 32
	defaultSecretEncoding = SecretEncodingBase64
)

// CryptoRandClientSecretCreator creates client secrets using cryptographically
// secure random number generation.
type CryptoRandClientSecretCreator struct {
	src      rand.Source
	length   int
	encoding string
}

// NewCryptoRandClientSecretCreator creates a new CryptoRandClientSecretCreator
// using the default crypto/rand source, 32-byte length, and Base64 encoding.
func NewCryptoRandClientSecretCreator() *CryptoRandClientSecretCreator {
	return &CryptoRandClientSecretCreator{
		src:      &cryptoSource{},
		length:   defaultSecretLength,
		encoding: defaultSecretEncoding,
	}
}

// NewCryptoRandClientSecretCreatorWithSource creates a new CryptoRandClientSecretCreator
// with a custom random source. Uses default length and encoding.
func NewCryptoRandClientSecretCreatorWithSource(src rand.Source) *CryptoRandClientSecretCreator {
	return &CryptoRandClientSecretCreator{
		src:      src,
		length:   defaultSecretLength,
		encoding: defaultSecretEncoding,
	}
}

// NewCryptoRandClientSecretCreatorWithConfig creates a new CryptoRandClientSecretCreator
// with configurable byte length and output encoding.
// length must be between 16 and 64; values outside this range are clamped.
// encoding must be one of "base64", "base64url", or "hex".
func NewCryptoRandClientSecretCreatorWithConfig(length int, encoding string) *CryptoRandClientSecretCreator {
	if length < 16 {
		length = 16
	}
	if length > 64 {
		length = 64
	}
	if encoding != SecretEncodingBase64URL && encoding != SecretEncodingHex {
		encoding = SecretEncodingBase64
	}
	return &CryptoRandClientSecretCreator{
		src:      &cryptoSource{},
		length:   length,
		encoding: encoding,
	}
}

// Create generates a new cryptographically secure client secret.
func (c *CryptoRandClientSecretCreator) Create() (string, error) {
	rnd := rand.New(c.src)
	buff := make([]byte, c.length)
	if _, err := rnd.Read(buff); err != nil {
		return "", err
	}

	switch c.encoding {
	case SecretEncodingBase64URL:
		return base64.RawURLEncoding.EncodeToString(buff), nil
	case SecretEncodingHex:
		return hex.EncodeToString(buff), nil
	default:
		return base64.StdEncoding.EncodeToString(buff), nil
	}
}
