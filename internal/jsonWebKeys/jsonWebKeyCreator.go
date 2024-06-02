package jsonWebKeys

type JsonWebKeyCreator interface {
	Create() (*jsonWebKeyOutput, error)
}
