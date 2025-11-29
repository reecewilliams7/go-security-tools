package jwk

import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
)

func TestNewRSAJSONWebKeyCreator(t *testing.T) {
	creator := NewRSAJSONWebKeyCreator(2048)
	if creator == nil {
		t.Fatal("expected non-nil RSAJSONWebKeyCreator")
	}
}

func TestRSAJSONWebKeyCreator_Create_2048(t *testing.T) {
	creator := NewRSAJSONWebKeyCreator(2048)

	output, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == nil {
		t.Fatal("expected non-nil output")
	}
	if output.JWK == nil {
		t.Error("expected non-nil JWK")
	}
	if output.JWKString == "" {
		t.Error("expected non-empty JWKString")
	}
	if output.JWKPublicString == "" {
		t.Error("expected non-empty JWKPublicString")
	}
	if output.Base64JWK == "" {
		t.Error("expected non-empty Base64JWK")
	}
	if output.PEMPrivateKey == "" {
		t.Error("expected non-empty PEMPrivateKey")
	}
	if output.PEMPublicKey == "" {
		t.Error("expected non-empty PEMPublicKey")
	}

	// Verify key type is RSA
	if output.JWK.KeyType() != jwa.RSA {
		t.Errorf("expected key type RSA, got %v", output.JWK.KeyType())
	}

	// Verify key ID is set
	if output.JWK.KeyID() == "" {
		t.Error("expected non-empty key ID")
	}
}

func TestRSAJSONWebKeyCreator_Create_4096(t *testing.T) {
	creator := NewRSAJSONWebKeyCreator(4096)

	output, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == nil {
		t.Fatal("expected non-nil output")
	}

	// Verify key type is RSA
	if output.JWK.KeyType() != jwa.RSA {
		t.Errorf("expected key type RSA, got %v", output.JWK.KeyType())
	}
}

func TestRSAJSONWebKeyCreator_Create_Uniqueness(t *testing.T) {
	creator := NewRSAJSONWebKeyCreator(2048)
	seen := make(map[string]bool)

	for i := 0; i < 3; i++ {
		output, err := creator.Create()
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		keyID := output.JWK.KeyID()
		if seen[keyID] {
			t.Errorf("duplicate key ID generated: %s", keyID)
		}
		seen[keyID] = true
	}
}
