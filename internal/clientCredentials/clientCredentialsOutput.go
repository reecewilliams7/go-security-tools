package clientCredentials

type clientCredentialsOutput struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func NewClientCredentialsOutput(clientId string, clientSecret string) *clientCredentialsOutput {
	return &clientCredentialsOutput{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}
