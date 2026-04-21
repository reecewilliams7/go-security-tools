package clientcredentials

import (
	"crypto/rand"
	"fmt"
)

// nanoidAlphabet is a URL-safe 64-character alphabet.
// Because len(alphabet) = 64 and 64 divides 256 evenly, a simple byte-mod
// mapping introduces no statistical bias.
const (
	nanoidAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
	nanoidSize     = 21
)

// NanoidClientIDCreator generates compact, URL-safe, collision-resistant
// client IDs using a crypto/rand source and a custom 64-character alphabet.
type NanoidClientIDCreator struct{}

// NewNanoidClientIDCreator returns a new NanoidClientIDCreator.
func NewNanoidClientIDCreator() *NanoidClientIDCreator {
	return &NanoidClientIDCreator{}
}

// Create generates a new nanoid-style client ID.
func (c *NanoidClientIDCreator) Create() (string, error) {
	b := make([]byte, nanoidSize)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate nanoid: %w", err)
	}

	result := make([]byte, nanoidSize)
	for i, v := range b {
		result[i] = nanoidAlphabet[v&63]
	}
	return string(result), nil
}
