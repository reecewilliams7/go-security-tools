package clientcredentials

// ClientCredentialsCreator creates client credentials using the provided
// client ID and client secret creators.
type ClientCredentialsCreator struct {
	clientIDCreator     ClientIDCreator
	clientSecretCreator ClientSecretCreator
}

// NewClientCredentialsCreator creates a new ClientCredentialsCreator.
func NewClientCredentialsCreator(clientIDCreator ClientIDCreator, clientSecretCreator ClientSecretCreator) *ClientCredentialsCreator {
	return &ClientCredentialsCreator{
		clientIDCreator:     clientIDCreator,
		clientSecretCreator: clientSecretCreator,
	}
}

// CreateClientCredentials generates a new ClientCredentials with a client ID and secret.
func (ccc *ClientCredentialsCreator) CreateClientCredentials() (*ClientCredentials, error) {
	clientID, err := ccc.clientIDCreator.Create()
	if err != nil {
		return nil, err
	}
	clientSecret, err := ccc.clientSecretCreator.Create()
	if err != nil {
		return nil, err
	}

	return &ClientCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}, nil
}
