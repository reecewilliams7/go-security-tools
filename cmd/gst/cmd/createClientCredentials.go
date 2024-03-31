package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	ccs "github.com/reecewilliams7/go-security-tools/internal/clientCredentials"
)

func init() {
	rootCmd.AddCommand(createClientCredentialsCmd)
}

var createClientCredentialsCmd = &cobra.Command{
	Use:   "create-client-credentials",
	Short: "Creates a Client ID and Secret",
	Long:  "TODO",
	RunE: func(cmd *cobra.Command, args []string) error {
		cc := ccs.NewClientCredentialsCreator()
		o := cc.Create()

		fmt.Printf("ClientId: %s \n", o.ClientId)
		fmt.Printf("ClientSecret %s \n", o.ClientSecret)

		return nil
	},
}
