package clientCredentials

import (
	"testing"
)

func TestCryptoSource_Int63(t *testing.T) {
	source := cryptoSource{}
	result := source.Int63()
	if result < 0 {
		t.Errorf("Expected non-negative result, got %d", result)
	}
}

func TestCryptoSource_Uint64(t *testing.T) {
	source := cryptoSource{}
	result := source.Uint64()
	if result == 0 {
		t.Errorf("Expected non-zero result, got %d", result)
	}
}
