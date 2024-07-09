package clientCredentials

import (
	"encoding/base64"
	"testing"
)

type mockCryptoRandSource struct{}

func (m *mockCryptoRandSource) Int63() int64 {
	return 0
}

func (m *mockCryptoRandSource) Seed(seed int64) {}

func TestCryptoRandClientSecretCreator_Create(t *testing.T) {
	mockSrc := &mockCryptoRandSource{}
	creator := NewCryptoRandClientSecretCreatorWithSource(mockSrc)
	secret, err := creator.Create()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedSecret := base64.StdEncoding.EncodeToString([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	if secret != expectedSecret {
		t.Errorf("Expected secret to be %q, got %q", expectedSecret, secret)
	}
}
