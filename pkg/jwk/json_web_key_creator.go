package jwk

// JSONWebKeyCreator is an interface for creating JSON Web Keys.
type JSONWebKeyCreator interface {
	Create() (*JSONWebKeyOutput, error)
}
