package clientCredentials

import (
	"testing"

	"github.com/google/uuid"
)

func TestGuidClientIdCreator_Create(t *testing.T) {
	creator := NewGuidClientIdCreator()
	clientID, err := creator.Create()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = uuid.Parse(clientID)
	if err != nil {
		t.Errorf("Invalid UUID format: %v", err)
	}
}
