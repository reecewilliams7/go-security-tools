package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ccs "github.com/reecewilliams7/go-security-tools/internal/clientCredentials"
)

func init() {
	createClientCredentialsCmd.Flags().IntP(CountFlagName, "c", 1, "The count to create.")

	clientCredentialsCmd.AddCommand(createClientCredentialsCmd)
}

var createClientCredentialsCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a Client ID and Secret that can be used as Client Credentials in OAuth2.0 and OpenID Connect",
	Long:  "TODO",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag(CountFlagName, cmd.Flags().Lookup(CountFlagName))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		count := viper.GetInt(CountFlagName)

		clientIdCreator := ccs.NewGuidClientIdCreator()
		clientSecretCreator := ccs.NewCryptoRandClientSecretCreator()

		for range count {
			fmt.Println("**********************************************************")
			ci, err := clientIdCreator.Create()
			if err != nil {
				return err
			}

			cs, err := clientSecretCreator.Create()
			if err != nil {
				return err
			}

			fmt.Println("Client Id:")
			fmt.Printf("%s\n", ci)
			fmt.Println("Client Secret:")
			fmt.Printf("%s\n", cs)
			fmt.Println("**********************************************************")
		}

		return nil
	},
}
