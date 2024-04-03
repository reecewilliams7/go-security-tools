package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jwks "github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
)

func init() {
	createJwkCmd.Flags().BoolP(OutputBase64FlagName, "b", false, "Whether to output the JWK as a Base64 string")
	viper.BindPFlag(OutputBase64FlagName, createJwkCmd.Flags().Lookup(OutputBase64FlagName))
	viper.BindEnv(OutputBase64FlagName)

	createJwkCmd.Flags().BoolP(OutputRsaKeysFlagName, "r", false, "Whether to output the JWK RSA private/public key pair as PEM encoded strings")
	viper.BindPFlag(OutputRsaKeysFlagName, createJwkCmd.Flags().Lookup(OutputRsaKeysFlagName))
	viper.BindEnv(OutputRsaKeysFlagName)

	RootCmd.AddCommand(createJwkCmd)
}

var createJwkCmd = &cobra.Command{
	Use:   "create-jwk",
	Short: "Creates a JsonWebKey",
	Long:  "TODO",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputBase64 := viper.GetBool(OutputBase64FlagName)
		outputRsaKeys := viper.GetBool(OutputRsaKeysFlagName)

		//logger := buildLogger("create-jwk")

		jwkc := jwks.NewJsonWebKeyCreator()
		o, err := jwkc.Create()
		if err != nil {
			return err
		}

		fmt.Println(o.JsonWebKeyString)
		if outputBase64 {
			fmt.Println("")
			fmt.Println("---------------------------")
			fmt.Println("")
			fmt.Println(o.Base64JsonWebKey)
			fmt.Println("")
		}
		if outputRsaKeys {
			fmt.Println("---------------------------")
			fmt.Println("")
			fmt.Println(o.RsaPrivateKey)
			fmt.Println("")
			fmt.Println("---------------------------")
			fmt.Println("")
			fmt.Println(o.RsaPublicKey)
			fmt.Println("")
		}
		return nil
	},
}
