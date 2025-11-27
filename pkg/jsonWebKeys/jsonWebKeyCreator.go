package jsonWebKeys

type JsonWebKeyCreator interface {
	Create() (*JsonWebKeyOutput, error)
}
