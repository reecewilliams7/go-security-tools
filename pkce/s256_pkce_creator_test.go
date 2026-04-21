package pkce

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
)

func TestNewS256Creator(t *testing.T) {
	creator := NewS256Creator()
	if creator == nil {
		t.Fatal("expected non-nil S256Creator")
	}
}

func TestS256Creator_Create(t *testing.T) {
	creator := NewS256Creator()

	pair, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pair == nil {
		t.Fatal("expected non-nil PKCEPair")
	}
	if pair.CodeVerifier == "" {
		t.Error("expected non-empty CodeVerifier")
	}
	if pair.CodeChallenge == "" {
		t.Error("expected non-empty CodeChallenge")
	}
	if pair.Method != "S256" {
		t.Errorf("expected method S256, got %s", pair.Method)
	}
}

func TestS256Creator_Create_VerifierLength(t *testing.T) {
	creator := NewS256Creator()

	pair, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// RFC 7636 requires 43–128 characters
	length := len(pair.CodeVerifier)
	if length < 43 || length > 128 {
		t.Errorf("verifier length %d is outside the RFC 7636 range 43–128", length)
	}
}

func TestS256Creator_Create_VerifierURLSafe(t *testing.T) {
	creator := NewS256Creator()

	pair, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// base64url characters: A-Z a-z 0-9 - _
	for _, ch := range pair.CodeVerifier {
		if !strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_", ch) {
			t.Errorf("verifier contains non-URL-safe character: %q", ch)
		}
	}
}

func TestS256Creator_Create_ChallengeDerivation(t *testing.T) {
	creator := NewS256Creator()

	pair, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify: challenge == BASE64URL(SHA256(verifier))
	sum := sha256.Sum256([]byte(pair.CodeVerifier))
	expected := base64.RawURLEncoding.EncodeToString(sum[:])

	if pair.CodeChallenge != expected {
		t.Errorf("code challenge does not match SHA256 of verifier\ngot:  %s\nwant: %s", pair.CodeChallenge, expected)
	}
}

func TestS256Creator_Create_Uniqueness(t *testing.T) {
	creator := NewS256Creator()
	seen := make(map[string]bool)

	for i := range 50 {
		pair, err := creator.Create()
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		if seen[pair.CodeVerifier] {
			t.Errorf("duplicate verifier generated: %s", pair.CodeVerifier)
		}
		seen[pair.CodeVerifier] = true
	}
}
