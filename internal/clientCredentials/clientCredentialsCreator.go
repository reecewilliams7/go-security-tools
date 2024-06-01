package clientCredentials

type ClientCredentialsCreator struct {
	clientIdCreator     ClientIdCreator
	clientSecretCreator ClientSecretCreator
}

func NewClientCredentialsCreator(clientIdCreator ClientIdCreator, clientSecretCreator ClientSecretCreator) *ClientCredentialsCreator {
	return &ClientCredentialsCreator{clientIdCreator: clientIdCreator, clientSecretCreator: clientSecretCreator}
}

func (c *ClientCredentialsCreator) Create() (*clientCredentialsOutput, error) {
	clientId, err := c.clientIdCreator.Create()
	if err != nil {
		return nil, err
	}
	clientSecret, err := c.clientSecretCreator.Create()
	if err != nil {
		return nil, err
	}

	o := NewClientCredentialsOutput(clientId, clientSecret)
	return o, nil
}
