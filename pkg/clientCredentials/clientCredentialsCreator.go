package clientCredentials

type ClientCredentialsCreator struct {
	clientIdCreator     ClientIdCreator
	clientSecretCreator ClientSecretCreator
}

func NewClientCredentialsCreator(clientIdCreator ClientIdCreator, clientSecretCreator ClientSecretCreator) *ClientCredentialsCreator {
	return &ClientCredentialsCreator{
		clientIdCreator:     clientIdCreator,
		clientSecretCreator: clientSecretCreator,
	}
}

func (ccc *ClientCredentialsCreator) CreateClientCredentials() (*ClientCredentials, error) {
	clientId, err := ccc.clientIdCreator.Create()
	if err != nil {
		return nil, err
	}
	clientSecret, err := ccc.clientSecretCreator.Create()
	if err != nil {
		return nil, err
	}

	return &ClientCredentials{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}, nil
}
