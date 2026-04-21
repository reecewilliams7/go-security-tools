package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(jwtCmd)
}

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Tools for working with JSON Web Tokens",
	Long:  "Commands for working with JSON Web Tokens (JWTs).",
}
