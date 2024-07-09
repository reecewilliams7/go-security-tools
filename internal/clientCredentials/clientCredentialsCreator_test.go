package clientCredentials

import (
	"testing"
)

type mockClientIdCreator struct{}

func (m *mockClientIdCreator) Create() (string, error) {
	return "mock-client-id", nil
}

type mockClientSecretCreator struct{}

func (m *mockClientSecretCreator) Create() (string, error) {
	return "mock-client-secret", nil
}

func TestClientCredentialsCreator_Create(t *testing.T) {
	clientIdCreator := &mockClientIdCreator{}
	clientSecretCreator := &mockClientSecretCreator{}
	creator := NewClientCredentialsCreator(clientIdCreator, clientSecretCreator)

	output, err := creator.Create()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedClientId := "mock-client-id"
	if output.ClientId != expectedClientId {
		t.Errorf("Expected ClientId to be %q, got %q", expectedClientId, output.ClientId)
	}

	expectedClientSecret := "mock-client-secret"
	if output.ClientSecret != expectedClientSecret {
		t.Errorf("Expected ClientSecret to be %q, got %q", expectedClientSecret, output.ClientSecret)
	}
}
