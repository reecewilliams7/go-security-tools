package cmd

import (
	"fmt"

	"github.com/reecewilliams7/go-security-tools/pkce"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createPKCECmd.Flags().IntP(CountFlag, "c", 1, "The number of PKCE pairs to create.")
	pkceCmd.AddCommand(createPKCECmd)
}

var createPKCECmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a PKCE code verifier and challenge pair",
	Long:  "Creates a PKCE code_verifier and code_challenge pair using the S256 method (RFC 7636).",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		return viper.BindPFlag(CountFlag, cmd.Flags().Lookup(CountFlag))
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		count := viper.GetInt(CountFlag)
		creator := pkce.NewS256Creator()

		for i := range count {
			fmt.Println("**********************************************************")
			fmt.Printf("PKCE Pair %d\n", i+1)

			pair, err := creator.Create()
			if err != nil {
				return err
			}

			fmt.Println("Code Verifier:")
			fmt.Println(pair.CodeVerifier)
			fmt.Println("Code Challenge:")
			fmt.Println(pair.CodeChallenge)
			fmt.Printf("Method: %s\n", pair.Method)
			fmt.Println("**********************************************************")
		}

		return nil
	},
}
