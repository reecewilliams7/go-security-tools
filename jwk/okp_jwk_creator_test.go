package jwk

import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
)

func TestNewOKPJWKCreator(t *testing.T) {
	creator := NewOKPJWKCreator("Ed25519")
	if creator == nil {
		t.Fatal("expected non-nil OKPJWKCreator")
	}
}

func TestOKPJWKCreator_Create_Ed25519(t *testing.T) {
	creator := NewOKPJWKCreator("Ed25519")

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
	if output.JWKPublic == nil {
		t.Error("expected non-nil JWKPublic")
	}
	if output.PEMPrivateKey == "" {
		t.Error("expected non-empty PEMPrivateKey")
	}
	if output.PEMPublicKey == "" {
		t.Error("expected non-empty PEMPublicKey")
	}
	if output.JWK.KeyType() != jwa.OKP {
		t.Errorf("expected key type OKP, got %v", output.JWK.KeyType())
	}
	if output.JWK.KeyID() == "" {
		t.Error("expected non-empty key ID")
	}
}

func TestOKPJWKCreator_Create_X25519(t *testing.T) {
	creator := NewOKPJWKCreator("X25519")

	output, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == nil {
		t.Fatal("expected non-nil output")
	}
	if output.JWK.KeyType() != jwa.OKP {
		t.Errorf("expected key type OKP, got %v", output.JWK.KeyType())
	}
	if output.JWK.KeyID() == "" {
		t.Error("expected non-empty key ID")
	}
}

func TestOKPJWKCreator_Create_UnknownCurve(t *testing.T) {
	creator := NewOKPJWKCreator("UnknownCurve")

	_, err := creator.Create()
	if err == nil {
		t.Fatal("expected error for unknown curve")
	}
}

func TestOKPJWKCreator_Create_Ed25519_Uniqueness(t *testing.T) {
	creator := NewOKPJWKCreator("Ed25519")
	seen := make(map[string]bool)

	for i := range 20 {
		output, err := creator.Create()
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		kid := output.JWK.KeyID()
		if seen[kid] {
			t.Errorf("duplicate key ID generated: %s", kid)
		}
		seen[kid] = true
	}
}
