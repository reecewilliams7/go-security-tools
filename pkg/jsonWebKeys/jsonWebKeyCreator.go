package jsonWebKeys

type JsonWebKeyCreator interface {
	Create(length int) (*JsonWebKeyOutput, error)
}
