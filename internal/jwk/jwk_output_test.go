package jwk

import (
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func TestNewJWKOutput(t *testing.T) {
	// Create an RSA key for testing
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	key, err := jwk.FromRaw(rsaKey)
	if err != nil {
		t.Fatalf("failed to create JWK from RSA key: %v", err)
	}

	if err := jwk.AssignKeyID(key); err != nil {
		t.Fatalf("failed to assign key ID: %v", err)
	}

	output, err := NewJWKOutput(key)
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

	// Verify JWKString is valid JSON
	if !strings.HasPrefix(output.JWKString, "{") {
		t.Error("JWKString should start with '{'")
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

	// Verify PEM format
	if !strings.Contains(output.PEMPrivateKey, "-----BEGIN") {
		t.Error("PEMPrivateKey should contain PEM header")
	}

	if output.PEMPublicKey == "" {
		t.Error("expected non-empty PEMPublicKey")
	}

	if !strings.Contains(output.PEMPublicKey, "-----BEGIN") {
		t.Error("PEMPublicKey should contain PEM header")
	}
}
