package clientCredentials

import "github.com/google/uuid"

type UUIDv7ClientIdCreator struct{}

func NewUUIDv7ClientIdCreator() *UUIDv7ClientIdCreator {
	return &UUIDv7ClientIdCreator{}
}

func (c *UUIDv7ClientIdCreator) Create() (string, error) {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return uuidv7.String(), nil
}
