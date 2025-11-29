package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createClientCredentialsCmd.Flags().IntP(CountFlag, "c", 1, "The count to create.")
	createClientCredentialsCmd.Flags().StringP(ClientIdTypeFlag, "t", ClientIdTypeUUIDv7, "The type of Client ID to create. Options are 'uuidv7' and 'short-uuid'.")
	createClientCredentialsCmd.Flags().StringP(ClientSecretTypeFlag, "s", ClientSecretTypeCryptoRand, "The type of Client Secret to create. Options are 'crypto-rand'.")
	clientCredentialsCmd.AddCommand(createClientCredentialsCmd)
}

var createClientCredentialsCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a Client ID and Secret that can be used as Client Credentials in OAuth2.0 and OpenID Connect",
	Long:  "Creates a Client ID and Secret pair for use with OAuth2.0 and OpenID Connect flows.",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		if err := viper.BindPFlag(CountFlag, cmd.Flags().Lookup(CountFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(ClientIdTypeFlag, cmd.Flags().Lookup(ClientIdTypeFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(ClientSecretTypeFlag, cmd.Flags().Lookup(ClientSecretTypeFlag)); err != nil {
			return err
		}
		return nil
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		count := viper.GetInt(CountFlag)
		clientIDType := viper.GetString(ClientIdTypeFlag)
		clientSecretType := viper.GetString(ClientSecretTypeFlag)

		ccc, err := buildClientCredentialsCreator(clientIDType, clientSecretType)
		if err != nil {
			return err
		}

		for range count {
			fmt.Println("**********************************************************")

			cc, err := ccc.CreateClientCredentials()
			if err != nil {
				return err
			}

			fmt.Println("Client ID:")
			fmt.Printf("%s\n", cc.ClientID())
			fmt.Println("Client Secret:")
			fmt.Printf("%s\n", cc.ClientSecret())
			fmt.Println("**********************************************************")
		}

		return nil
	},
}
