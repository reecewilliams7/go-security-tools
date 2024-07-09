package clientCredentials

import (
	"testing"
)

func TestNewClientCredentialsOutput(t *testing.T) {
	clientId := "test-client-id"
	clientSecret := "test-client-secret"
	output := NewClientCredentialsOutput(clientId, clientSecret)

	if output.ClientId != clientId {
		t.Errorf("Expected ClientId to be %q, got %q", clientId, output.ClientId)
	}

	if output.ClientSecret != clientSecret {
		t.Errorf("Expected ClientSecret to be %q, got %q", clientSecret, output.ClientSecret)
	}
}
