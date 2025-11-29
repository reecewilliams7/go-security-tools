package clientcredentials

import "testing"

func TestCryptoSource_Int63(t *testing.T) {
	src := &cryptoSource{}

	for i := 0; i < 100; i++ {
		v := src.Int63()
		// Int63 should always return non-negative values
		if v < 0 {
			t.Errorf("Int63 returned negative value: %d", v)
		}
	}
}

func TestCryptoSource_Uint64(t *testing.T) {
	src := &cryptoSource{}

	// Just verify it doesn't panic and returns something
	for i := 0; i < 100; i++ {
		_ = src.Uint64()
	}
}

func TestCryptoSource_Seed(t *testing.T) {
	src := &cryptoSource{}
	// Seed should be a no-op - just verify it doesn't panic
	src.Seed(12345)
}

func TestCryptoSource_Randomness(t *testing.T) {
	src := &cryptoSource{}
	seen := make(map[uint64]bool)

	// Generate 100 random numbers and check they're unique
	for i := 0; i < 100; i++ {
		v := src.Uint64()
		if seen[v] {
			t.Errorf("duplicate value generated: %d", v)
		}
		seen[v] = true
	}
}
