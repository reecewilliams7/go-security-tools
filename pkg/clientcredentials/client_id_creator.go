package clientcredentials

// ClientIDCreator is an interface for creating client IDs.
type ClientIDCreator interface {
	Create() (string, error)
}
