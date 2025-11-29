package clientcredentials

// ClientSecretCreator is an interface for creating client secrets.
type ClientSecretCreator interface {
	Create() (string, error)
}
