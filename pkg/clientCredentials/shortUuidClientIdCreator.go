package clientCredentials

import "github.com/lithammer/shortuuid/v4"

type ShortUuidClientIdCreator struct{}

func NewShortUuidClientIdCreator() *ShortUuidClientIdCreator {
	return &ShortUuidClientIdCreator{}
}

func (c *ShortUuidClientIdCreator) Create() (string, error) {
	uuid := shortuuid.New()
	return uuid, nil
}
