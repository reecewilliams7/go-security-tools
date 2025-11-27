package clientCredentials

type ClientSecretCreator interface {
	Create() (string, error)
}
