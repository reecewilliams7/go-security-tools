package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	ccs "github.com/reecewilliams7/go-security-tools/internal/clientCredentials"
)

func init() {
	clientCredentialsCmd.AddCommand(createClientCredentialsCmd)
}

var createClientCredentialsCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a Client ID and Secret that can be used as Client Credentials in OAuth2.0 and OpenID Connect",
	Long:  "TODO",
	RunE: func(cmd *cobra.Command, args []string) error {
		clientIdCreator := ccs.NewGuidClientIdCreator()
		clientSecretCreator := ccs.NewCryptoRandClientSecretCreator()
		cc := ccs.NewClientCredentialsCreator(clientIdCreator, clientSecretCreator)
		o, err := cc.Create()
		if err != nil {
			return err
		}

		fmt.Println("Client Id:")
		fmt.Printf("%s\n", o.ClientId)
		fmt.Println("Client Secret:")
		fmt.Printf("%s\n", o.ClientSecret)

		return nil
	},
}
