package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(clientCredentialsCmd)
}

var clientCredentialsCmd = &cobra.Command{
	Use:   "client-credentials",
	Short: "Tools for working with Client Credentials",
	Long:  "Tools for working with Client Credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
