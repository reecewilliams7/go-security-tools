package clientcredentials

import "github.com/google/uuid"

// UUIDv7ClientIDCreator creates client IDs using UUID v7.
type UUIDv7ClientIDCreator struct{}

// NewUUIDv7ClientIDCreator creates a new UUIDv7ClientIDCreator.
func NewUUIDv7ClientIDCreator() *UUIDv7ClientIDCreator {
	return &UUIDv7ClientIDCreator{}
}

// Create generates a new UUID v7 client ID.
func (c *UUIDv7ClientIDCreator) Create() (string, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return uuidV7.String(), nil
}
