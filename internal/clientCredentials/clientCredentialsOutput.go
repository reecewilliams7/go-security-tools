package clientCredentials

type clientCredentialsOutput struct {
	ClientId     string
	ClientSecret string
}

func NewClientCredentialsOutput(clientId string, clientSecret string) *clientCredentialsOutput {
	return &clientCredentialsOutput{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}
