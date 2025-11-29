package jwk

import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
)

func TestNewECDSAJWKCreator(t *testing.T) {
	creator := NewECDSAJWKCreator("P256")
	if creator == nil {
		t.Fatal("expected non-nil ECDSAJWKCreator")
	}
}

func TestECDSAJWKCreator_Create_P256(t *testing.T) {
	creator := NewECDSAJWKCreator("P256")

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

	// Verify key type is EC
	if output.JWK.KeyType() != jwa.EC {
		t.Errorf("expected key type EC, got %v", output.JWK.KeyType())
	}

	// Verify key ID is set
	if output.JWK.KeyID() == "" {
		t.Error("expected non-empty key ID")
	}
}

func TestECDSAJWKCreator_Create_P384(t *testing.T) {
	creator := NewECDSAJWKCreator("P384")

	output, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == nil {
		t.Fatal("expected non-nil output")
	}

	// Verify key type is EC
	if output.JWK.KeyType() != jwa.EC {
		t.Errorf("expected key type EC, got %v", output.JWK.KeyType())
	}
}

func TestECDSAJWKCreator_Create_P521(t *testing.T) {
	creator := NewECDSAJWKCreator("P521")

	output, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == nil {
		t.Fatal("expected non-nil output")
	}

	// Verify key type is EC
	if output.JWK.KeyType() != jwa.EC {
		t.Errorf("expected key type EC, got %v", output.JWK.KeyType())
	}
}

func TestECDSAJWKCreator_Create_DefaultCurve(t *testing.T) {
	// Unknown curve type should default to P256
	creator := NewECDSAJWKCreator("unknown")

	output, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == nil {
		t.Fatal("expected non-nil output")
	}

	// Verify key type is EC
	if output.JWK.KeyType() != jwa.EC {
		t.Errorf("expected key type EC, got %v", output.JWK.KeyType())
	}
}

func TestECDSAJWKCreator_Create_Uniqueness(t *testing.T) {
	creator := NewECDSAJWKCreator("P256")
	seen := make(map[string]bool)

	for i := 0; i < 5; i++ {
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
