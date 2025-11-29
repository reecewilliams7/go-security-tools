package clientcredentials

import (
	"errors"
	"testing"
)

// mockClientIDCreator is a mock implementation of ClientIDCreator.
type mockClientIDCreator struct {
	id  string
	err error
}

func (m *mockClientIDCreator) Create() (string, error) {
	return m.id, m.err
}

// mockClientSecretCreator is a mock implementation of ClientSecretCreator.
type mockClientSecretCreator struct {
	secret string
	err    error
}

func (m *mockClientSecretCreator) Create() (string, error) {
	return m.secret, m.err
}

func TestNewClientCredentialsCreator(t *testing.T) {
	idCreator := &mockClientIDCreator{}
	secretCreator := &mockClientSecretCreator{}

	ccc := NewClientCredentialsCreator(idCreator, secretCreator)
	if ccc == nil {
		t.Fatal("expected non-nil ClientCredentialsCreator")
	}
}

func TestCreateClientCredentials_Success(t *testing.T) {
	expectedID := "test-client-id"
	expectedSecret := "test-client-secret"

	idCreator := &mockClientIDCreator{id: expectedID}
	secretCreator := &mockClientSecretCreator{secret: expectedSecret}
	ccc := NewClientCredentialsCreator(idCreator, secretCreator)

	cc, err := ccc.CreateClientCredentials()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cc.ClientID() != expectedID {
		t.Errorf("expected ClientID %q, got %q", expectedID, cc.ClientID())
	}
	if cc.ClientSecret() != expectedSecret {
		t.Errorf("expected ClientSecret %q, got %q", expectedSecret, cc.ClientSecret())
	}
}

func TestCreateClientCredentials_ClientIDError(t *testing.T) {
	expectedErr := errors.New("client id creation failed")
	idCreator := &mockClientIDCreator{err: expectedErr}
	secretCreator := &mockClientSecretCreator{secret: "secret"}
	ccc := NewClientCredentialsCreator(idCreator, secretCreator)

	cc, err := ccc.CreateClientCredentials()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
	if cc != nil {
		t.Errorf("expected nil ClientCredentials, got %v", cc)
	}
}

func TestCreateClientCredentials_ClientSecretError(t *testing.T) {
	expectedErr := errors.New("client secret creation failed")
	idCreator := &mockClientIDCreator{id: "test-id"}
	secretCreator := &mockClientSecretCreator{err: expectedErr}
	ccc := NewClientCredentialsCreator(idCreator, secretCreator)

	cc, err := ccc.CreateClientCredentials()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
	if cc != nil {
		t.Errorf("expected nil ClientCredentials, got %v", cc)
	}
}
