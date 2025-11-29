package jwk

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"path/filepath"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func createTestJWKOutput(t *testing.T) *JWKOutput {
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

func TestNewFileJwkOutputWriter(t *testing.T) {
	writer := NewFileJwkOutputWriter("/tmp", "test", false, false)
	if writer == nil {
		t.Fatal("expected non-nil FileJwkOutputWriter")
	}
}

func TestFileJwkOutputWriter_Write_Basic(t *testing.T) {
	tempDir := t.TempDir()
	writer := NewFileJwkOutputWriter(tempDir, "test-jwk", false, false)
	output := createTestJWKOutput(t)

	err := writer.Write(output, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that files were created
	jwkFile := filepath.Join(tempDir, "test-jwk-1.jwk")
	if _, err := os.Stat(jwkFile); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", jwkFile)
	}

	pubFile := filepath.Join(tempDir, "test-jwk-pub-1.jwk")
	if _, err := os.Stat(pubFile); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", pubFile)
	}
}

func TestFileJwkOutputWriter_Write_WithBase64(t *testing.T) {
	tempDir := t.TempDir()
	writer := NewFileJwkOutputWriter(tempDir, "test-jwk", true, false)
	output := createTestJWKOutput(t)

	err := writer.Write(output, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that base64 file was created
	base64File := filepath.Join(tempDir, "test-jwk-base64-1.jwk")
	if _, err := os.Stat(base64File); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", base64File)
	}
}

func TestFileJwkOutputWriter_Write_WithPemKeys(t *testing.T) {
	tempDir := t.TempDir()
	writer := NewFileJwkOutputWriter(tempDir, "test-jwk", false, true)
	output := createTestJWKOutput(t)

	err := writer.Write(output, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that PEM files were created
	pubFile := filepath.Join(tempDir, "test-jwk-1.pub")
	if _, err := os.Stat(pubFile); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", pubFile)
	}

	keyFile := filepath.Join(tempDir, "test-jwk-1.key")
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", keyFile)
	}
}

func TestFileJwkOutputWriter_Write_MultipleIterations(t *testing.T) {
	tempDir := t.TempDir()
	writer := NewFileJwkOutputWriter(tempDir, "test-jwk", false, false)

	for i := 1; i <= 3; i++ {
		output := createTestJWKOutput(t)
		err := writer.Write(output, i)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		jwkFile := filepath.Join(tempDir, "test-jwk-"+string(rune('0'+i))+".jwk")
		// Use proper integer formatting
		expectedFile := filepath.Join(tempDir, "test-jwk-"+itoa(i)+".jwk")
		if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", expectedFile)
		}
		_ = jwkFile // Suppress unused warning
	}
}

func TestGetOutputFilePath(t *testing.T) {
	tests := []struct {
		outputPath string
		outputFile string
		ext        string
		i          int
		expected   string
	}{
		{"/tmp", "test", "jwk", 1, "/tmp/test-1.jwk"},
		{"/home/user", "mykey", "pem", 5, "/home/user/mykey-5.pem"},
		{".", "output", "key", 10, "./output-10.key"},
	}

	for _, tt := range tests {
		result := getOutputFilePath(tt.outputPath, tt.outputFile, tt.ext, tt.i)
		if result != tt.expected {
			t.Errorf("getOutputFilePath(%q, %q, %q, %d) = %q, expected %q",
				tt.outputPath, tt.outputFile, tt.ext, tt.i, result, tt.expected)
		}
	}
}

// Helper function to convert int to string
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	result := ""
	for i > 0 {
		result = string(rune('0'+i%10)) + result
		i /= 10
	}
	return result
}
