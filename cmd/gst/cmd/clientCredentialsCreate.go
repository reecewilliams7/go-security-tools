package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createClientCredentialsCmd.Flags().IntP(CountFlagName, "c", 1, "The count to create.")
	createClientCredentialsCmd.Flags().StringP(ClientIdTypeFlag, "t", ClientIdTypeUUIDv7, "The type of Client ID to create. Options are 'uuidv7' and 'short'.")
	createClientCredentialsCmd.Flags().StringP(ClientSecretTypeFlag, "s", ClientSecretTypeCryptoRand, "The type of Client Secret to create. Options are 'crypto-rand'.")
	clientCredentialsCmd.AddCommand(createClientCredentialsCmd)
}

var createClientCredentialsCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a Client ID and Secret that can be used as Client Credentials in OAuth2.0 and OpenID Connect",
	Long:  "TODO",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag(CountFlagName, cmd.Flags().Lookup(CountFlagName))
		viper.BindPFlag(ClientIdTypeFlag, cmd.Flags().Lookup(ClientIdTypeFlag))
		viper.BindPFlag(ClientSecretTypeFlag, cmd.Flags().Lookup(ClientSecretTypeFlag))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		count := viper.GetInt(CountFlagName)
		clientIdType := viper.GetString(ClientIdTypeFlag)
		clientSecretType := viper.GetString(ClientSecretTypeFlag)

		ccc, err := buildClientCredentialsCreator(clientIdType, clientSecretType)
		if err != nil {
			return err
		}

		for range count {
			fmt.Println("**********************************************************")

			cc, err := ccc.CreateClientCredentials()
			if err != nil {
				return err
			}

			fmt.Println("Client Id:")
			fmt.Printf("%s\n", cc.ClientID)
			fmt.Println("Client Secret:")
			fmt.Printf("%s\n", cc.ClientSecret)
			fmt.Println("**********************************************************")
		}

		return nil
	},
}
