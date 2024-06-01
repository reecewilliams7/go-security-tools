package clientCredentials

import (
	"encoding/base64"
	rand "math/rand"
)

type CryptoRandClientSecretCreator struct{}

func NewCryptoRandClientSecretCreator() *CryptoRandClientSecretCreator {
	return &CryptoRandClientSecretCreator{}
}

func (c *CryptoRandClientSecretCreator) Create() (string, error) {
	var src cryptoSource
	rnd := rand.New(src)
	buff := make([]byte, 32)
	_, err := rnd.Read(buff)
	if err != nil {
		return "", err
	}

	secret := base64.StdEncoding.EncodeToString(buff)
	return secret, nil
}
