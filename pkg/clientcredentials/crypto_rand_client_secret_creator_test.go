package clientcredentials

import (
	"encoding/base64"
	"math/rand"
	"testing"
)

func TestNewCryptoRandClientSecretCreator(t *testing.T) {
	creator := NewCryptoRandClientSecretCreator()
	if creator == nil {
		t.Fatal("expected non-nil CryptoRandClientSecretCreator")
	}
}

func TestNewCryptoRandClientSecretCreatorWithSource(t *testing.T) {
	// Use a deterministic source for testing
	src := rand.NewSource(12345)
	creator := NewCryptoRandClientSecretCreatorWithSource(src)
	if creator == nil {
		t.Fatal("expected non-nil CryptoRandClientSecretCreator")
	}
}

func TestCryptoRandClientSecretCreator_Create(t *testing.T) {
	creator := NewCryptoRandClientSecretCreator()

	secret, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if secret == "" {
		t.Error("expected non-empty secret")
	}

	// Verify it's valid base64
	decoded, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		t.Errorf("secret is not valid base64: %v", err)
	}
	// Should decode to 32 bytes
	if len(decoded) != 32 {
		t.Errorf("expected 32 decoded bytes, got %d", len(decoded))
	}
}

func TestCryptoRandClientSecretCreator_Create_Uniqueness(t *testing.T) {
	creator := NewCryptoRandClientSecretCreator()
	seen := make(map[string]bool)

	for i := 0; i < 100; i++ {
		secret, err := creator.Create()
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		if seen[secret] {
			t.Errorf("duplicate secret generated: %s", secret)
		}
		seen[secret] = true
	}
}

func TestCryptoRandClientSecretCreator_CreateWithDeterministicSource(t *testing.T) {
	// Using a deterministic source should produce deterministic output
	src := rand.NewSource(12345)
	creator := NewCryptoRandClientSecretCreatorWithSource(src)

	secret1, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Reset with same seed
	src = rand.NewSource(12345)
	creator = NewCryptoRandClientSecretCreatorWithSource(src)

	secret2, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if secret1 != secret2 {
		t.Errorf("expected same secrets with same seed, got %q and %q", secret1, secret2)
	}
}
