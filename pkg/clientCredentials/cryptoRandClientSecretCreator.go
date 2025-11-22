package clientCredentials

import (
	"encoding/base64"
	rand "math/rand"
)

type CryptoRandClientSecretCreator struct {
	src rand.Source
}

func NewCryptoRandClientSecretCreator() *CryptoRandClientSecretCreator {
	cryptoSrc := &cryptoSource{}
	return &CryptoRandClientSecretCreator{
		src: cryptoSrc,
	}
}

func NewCryptoRandClientSecretCreatorWithSource(src rand.Source) *CryptoRandClientSecretCreator {
	return &CryptoRandClientSecretCreator{
		src: src,
	}
}

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
