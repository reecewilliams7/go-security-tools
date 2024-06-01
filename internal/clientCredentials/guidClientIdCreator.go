package clientCredentials

import "github.com/google/uuid"

type GuidClientIdCreator struct{}

func NewGuidClientIdCreator() *GuidClientIdCreator {
	return &GuidClientIdCreator{}
}

func (c *GuidClientIdCreator) Create() (string, error) {
	return uuid.New().String(), nil
}
