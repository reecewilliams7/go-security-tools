package jwk

import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
)

func TestNewHMACJWKCreator(t *testing.T) {
	creator := NewHMACJWKCreator("HS256")
	if creator == nil {
		t.Fatal("expected non-nil HMACJWKCreator")
	}
}

func TestHMACJWKCreator_Create(t *testing.T) {
	tests := []struct {
		algorithm string
		wantSize  int
	}{
		{"HS256", 32},
		{"HS384", 48},
		{"HS512", 64},
	}

	for _, tc := range tests {
		t.Run(tc.algorithm, func(t *testing.T) {
			creator := NewHMACJWKCreator(tc.algorithm)

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
			if output.Base64JWK == "" {
				t.Error("expected non-empty Base64JWK")
			}
			// Symmetric keys have no public key
			if output.JWKPublic != nil {
				t.Error("expected nil JWKPublic for symmetric key")
			}
			if output.JWKPublicString != "" {
				t.Error("expected empty JWKPublicString for symmetric key")
			}
			if output.PEMPrivateKey != "" {
				t.Error("expected empty PEMPrivateKey for symmetric key")
			}
			if output.PEMPublicKey != "" {
				t.Error("expected empty PEMPublicKey for symmetric key")
			}
			if output.JWK.KeyType() != jwa.OctetSeq {
				t.Errorf("expected key type OctSeq, got %v", output.JWK.KeyType())
			}
			if output.JWK.KeyID() == "" {
				t.Error("expected non-empty key ID")
			}
		})
	}
}

func TestHMACJWKCreator_Create_UnknownAlgorithm(t *testing.T) {
	creator := NewHMACJWKCreator("HS999")

	_, err := creator.Create()
	if err == nil {
		t.Fatal("expected error for unknown algorithm")
	}
}

func TestHMACJWKCreator_Create_Uniqueness(t *testing.T) {
	creator := NewHMACJWKCreator("HS256")
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
