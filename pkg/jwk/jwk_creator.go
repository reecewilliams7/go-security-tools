package jwk

import (
	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
)

// JWKCreator is an interface for creating JSON Web Keys.
type JWKCreator interface {
	Create() (*internaljwk.JWKOutput, error)
}
