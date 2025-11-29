package clientcredentials

import (
	crand "crypto/rand"
	"encoding/binary"
)

// cryptoSource implements rand.Source using crypto/rand for cryptographically
// secure random number generation.
type cryptoSource struct{}

// Seed is a no-op for cryptoSource since crypto/rand is self-seeding.
func (s cryptoSource) Seed(_ int64) {}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
// Note: This method panics if crypto/rand fails, as the rand.Source interface
// does not support returning errors. In practice, crypto/rand.Read should never
// fail on supported systems.
func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

// Uint64 returns a pseudo-random 64-bit unsigned integer.
// Note: This method panics if crypto/rand fails, as the rand.Source interface
// does not support returning errors. In practice, crypto/rand.Read should never
// fail on supported systems.
func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return v
}
