package clientcredentials

// ClientCredentials represents OAuth 2.0 client credentials.
type ClientCredentials struct {
	clientID     string
	clientSecret string
}

func NewClientCredentials(clientID, clientSecret string) *ClientCredentials {
	return &ClientCredentials{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (cc *ClientCredentials) ClientID() string {
	return cc.clientID
}

func (cc *ClientCredentials) ClientSecret() string {
	return cc.clientSecret
}
