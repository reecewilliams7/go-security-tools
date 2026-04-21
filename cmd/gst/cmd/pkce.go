package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(pkceCmd)
}

var pkceCmd = &cobra.Command{
	Use:   "pkce",
	Short: "Tools for working with PKCE (Proof Key for Code Exchange)",
	Long:  "Commands for working with PKCE as defined in RFC 7636.",
}
