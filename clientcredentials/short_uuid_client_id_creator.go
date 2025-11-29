package clientcredentials

import "github.com/lithammer/shortuuid/v4"

// ShortUUIDClientIDCreator creates client IDs using short UUIDs.
type ShortUUIDClientIDCreator struct{}

// NewShortUUIDClientIDCreator creates a new ShortUUIDClientIDCreator.
func NewShortUUIDClientIDCreator() *ShortUUIDClientIDCreator {
	return &ShortUUIDClientIDCreator{}
}

// Create generates a new short UUID client ID.
func (c *ShortUUIDClientIDCreator) Create() (string, error) {
	id := shortuuid.New()
	return id, nil
}
