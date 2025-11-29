package jwk

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func createTestJWKOutputForFmt(t *testing.T) *JWKOutput {
	t.Helper()

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
		t.Fatalf("failed to create JWK output: %v", err)
	}

	return output
}

func TestNewFmtJWKOutputWriter(t *testing.T) {
	writer := NewFmtJWKOutputWriter(false, false)
	if writer == nil {
		t.Fatal("expected non-nil FmtJWKOutputWriter")
	}
}

func TestFmtJWKOutputWriter_Write_Basic(t *testing.T) {
	writer := NewFmtJWKOutputWriter(false, false)
	output := createTestJWKOutputForFmt(t)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := writer.Write(output, 1)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	stdout := buf.String()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check output contains expected elements
	if !strings.Contains(stdout, "JWK Private Key:") {
		t.Error("expected output to contain 'JWK Private Key:'")
	}
	if !strings.Contains(stdout, "JWK Public Key:") {
		t.Error("expected output to contain 'JWK Public Key:'")
	}
	if !strings.Contains(stdout, asteriskLine) {
		t.Error("expected output to contain asterisk separator line")
	}
}

func TestFmtJWKOutputWriter_Write_WithBase64(t *testing.T) {
	writer := NewFmtJWKOutputWriter(true, false)
	output := createTestJWKOutputForFmt(t)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := writer.Write(output, 1)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	stdout := buf.String()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check output contains base64 section
	if !strings.Contains(stdout, "Base64 Encoded JWK Private Key:") {
		t.Error("expected output to contain 'Base64 Encoded JWK Private Key:'")
	}
}

func TestFmtJWKOutputWriter_Write_WithPemKeys(t *testing.T) {
	writer := NewFmtJWKOutputWriter(false, true)
	output := createTestJWKOutputForFmt(t)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := writer.Write(output, 1)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	stdout := buf.String()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check output contains PEM sections
	if !strings.Contains(stdout, "Private Key:") {
		t.Error("expected output to contain 'Private Key:'")
	}
	if !strings.Contains(stdout, "Public Key:") {
		t.Error("expected output to contain 'Public Key:'")
	}
}
