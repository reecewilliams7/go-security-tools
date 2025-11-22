package clientCredentials

type ClientIdCreator interface {
	Create() (string, error)
}
