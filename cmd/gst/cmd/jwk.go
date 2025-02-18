package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(jwkCmd)
}

var jwkCmd = &cobra.Command{
	Use:   "jwk",
	Short: "Tools for working with JSON Web Keys",
	Long:  "Commands for working with JSON Web Keys, such as creating and returnning the public key for a given private key",
}
