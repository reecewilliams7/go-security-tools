package clientcredentials

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
)

// cryptoSource implements rand.Source using crypto/rand for cryptographically
// secure random number generation.
type cryptoSource struct{}

// Seed is a no-op for cryptoSource since crypto/rand is self-seeding.
func (s cryptoSource) Seed(_ int64) {}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

// Uint64 returns a pseudo-random 64-bit unsigned integer.
func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
