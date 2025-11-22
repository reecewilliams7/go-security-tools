package clientCredentials

type ClientCredentialsCreator struct {
	ClientIdCreator     ClientIdCreator
	ClientSecretCreator ClientSecretCreator
}

func NewClientCredentialsCreator(clientIdCreator ClientIdCreator, clientSecretCreator ClientSecretCreator) *ClientCredentialsCreator {
	return &ClientCredentialsCreator{
		ClientIdCreator:     clientIdCreator,
		ClientSecretCreator: clientSecretCreator,
	}
}

func (ccc *ClientCredentialsCreator) CreateClientCredentials() (*ClientCredentials, error) {
	clientId, err := ccc.ClientIdCreator.Create()
	if err != nil {
		return nil, err
	}
	clientSecret, err := ccc.ClientSecretCreator.Create()
	if err != nil {
		return nil, err
	}

	return &ClientCredentials{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}, nil
}
