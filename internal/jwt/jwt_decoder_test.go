package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func makeToken(header, payload map[string]any) string {
	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)
	return fmt.Sprintf("%s.%s.fakesig",
		base64.RawURLEncoding.EncodeToString(h),
		base64.RawURLEncoding.EncodeToString(p),
	)
}

func TestDecode_Valid(t *testing.T) {
	token := makeToken(
		map[string]any{"alg": "RS256", "typ": "JWT"},
		map[string]any{"sub": "user123", "iss": "https://example.com"},
	)

	out, err := Decode(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Header["alg"] != "RS256" {
		t.Errorf("expected alg RS256, got %v", out.Header["alg"])
	}
	if out.Payload["sub"] != "user123" {
		t.Errorf("expected sub user123, got %v", out.Payload["sub"])
	}
}

func TestDecode_Expiry_NotExpired(t *testing.T) {
	future := time.Now().Add(time.Hour).Unix()
	token := makeToken(
		map[string]any{"alg": "RS256"},
		map[string]any{"exp": float64(future)},
	)

	out, err := Decode(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ExpiresAt == nil {
		t.Fatal("expected non-nil ExpiresAt")
	}
	if out.IsExpired {
		t.Error("expected token to not be expired")
	}
}

func TestDecode_Expiry_Expired(t *testing.T) {
	past := time.Now().Add(-time.Hour).Unix()
	token := makeToken(
		map[string]any{"alg": "RS256"},
		map[string]any{"exp": float64(past)},
	)

	out, err := Decode(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.IsExpired {
		t.Error("expected token to be expired")
	}
}

func TestDecode_NoExpiry(t *testing.T) {
	token := makeToken(
		map[string]any{"alg": "RS256"},
		map[string]any{"sub": "user"},
	)

	out, err := Decode(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ExpiresAt != nil {
		t.Error("expected nil ExpiresAt")
	}
	if out.IsExpired {
		t.Error("expected IsExpired to be false")
	}
}

func TestDecode_InvalidParts(t *testing.T) {
	_, err := Decode("only.two")
	if err == nil {
		t.Fatal("expected error for token with wrong number of parts")
	}
	if !strings.Contains(err.Error(), "3") {
		t.Errorf("error should mention expected 3 parts, got: %v", err)
	}
}

func TestDecode_InvalidBase64Header(t *testing.T) {
	_, err := Decode("!!!.payload.sig")
	if err == nil {
		t.Fatal("expected error for invalid base64 header")
	}
}

func TestDecode_InvalidJSONPayload(t *testing.T) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256"}`))
	badPayload := base64.RawURLEncoding.EncodeToString([]byte(`not-json`))
	token := header + "." + badPayload + ".sig"

	_, err := Decode(token)
	if err == nil {
		t.Fatal("expected error for invalid JSON payload")
	}
}

func TestDecode_TrimsWhitespace(t *testing.T) {
	token := makeToken(
		map[string]any{"alg": "HS256"},
		map[string]any{"sub": "trimmed"},
	)

	out, err := Decode("  " + token + "\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Payload["sub"] != "trimmed" {
		t.Error("expected whitespace to be trimmed from token")
	}
}
