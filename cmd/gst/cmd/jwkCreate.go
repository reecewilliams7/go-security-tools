package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jwks "github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
)

func init() {
	createJwkCmd.Flags().BoolP(OutputBase64FlagName, "b", false, "Whether to output the JWK as a Base64 string.")
	viper.BindPFlag(OutputBase64FlagName, createJwkCmd.Flags().Lookup(OutputBase64FlagName))
	viper.BindEnv(OutputBase64FlagName)

	createJwkCmd.Flags().BoolP(OutputRsaKeysFlagName, "r", false, "Whether to output the JWK RSA private/public key pair as PEM encoded strings.")
	viper.BindPFlag(OutputRsaKeysFlagName, createJwkCmd.Flags().Lookup(OutputRsaKeysFlagName))
	viper.BindEnv(OutputRsaKeysFlagName)

	createJwkCmd.Flags().StringP(OutputPathFlagName, "p", "", "The path to write the JWK output to. Will withhold output from the console when specified.")
	viper.BindPFlag(OutputPathFlagName, createJwkCmd.Flags().Lookup(OutputPathFlagName))
	viper.BindEnv(OutputPathFlagName)

	createJwkCmd.Flags().StringP(OutputFileNameFlagName, "f", OutputFileDefaultName, "The name of the file(s) to write to.")
	viper.BindPFlag(OutputFileNameFlagName, createJwkCmd.Flags().Lookup(OutputFileNameFlagName))
	viper.BindEnv(OutputFileNameFlagName)

	jwkCmd.AddCommand(createJwkCmd)
}

var createJwkCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a JSON Web Key",
	Long:  "Creates a JSON Web Key using an RSA Private Key",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputBase64 := viper.GetBool(OutputBase64FlagName)
		outputRsaKeys := viper.GetBool(OutputRsaKeysFlagName)
		outputPath := viper.GetString(OutputPathFlagName)
		outputFile := viper.GetString(OutputFileNameFlagName)

		logger := buildLogger("jwk create")

		var writeToFile = false

		if len(outputPath) > 0 {
			if _, err := os.Stat(outputPath); os.IsNotExist(err) {
				return err
			}
			logger.Info(fmt.Sprintf("Output path specified so will write JWK files to following location: %s", outputPath))
			writeToFile = true
		}

		jwkc := jwks.NewRsaJsonWebKeyCreator()
		o, err := jwkc.Create()
		logger.Info("New JWK created successfully. Will now write to output.")
		if err != nil {
			return err
		}

		if writeToFile {
			jwkFilePath := getOutputFilePath(outputPath, outputFile, "jwk")
			os.WriteFile(jwkFilePath, []byte(o.JsonWebKeyString), os.ModePerm)
			logger.Info(fmt.Sprintf("JWK Private key file written to: %s", jwkFilePath))

			jwkPubFilePath := getOutputFilePath(outputPath, fmt.Sprintf("%s-pub", outputFile), "jwk")
			os.WriteFile(jwkPubFilePath, []byte(o.JsonWebKeyPublicString), os.ModePerm)
			logger.Info(fmt.Sprintf("JWK Public key file written to: %s", jwkPubFilePath))

			if outputBase64 {
				base64FilePath := getOutputFilePath(outputPath, fmt.Sprintf("%s-base64", outputFile), "jwk")
				os.WriteFile(base64FilePath, []byte(o.Base64JsonWebKey), os.ModePerm)
				logger.Info(fmt.Sprintf("Base64 JWK Private key file written to: %s", base64FilePath))
			}
			if outputRsaKeys {
				rsaPubFilePath := getOutputFilePath(outputPath, outputFile, "pub")
				os.WriteFile(rsaPubFilePath, []byte(o.RsaPublicKey), os.ModePerm)
				logger.Info(fmt.Sprintf("RSA Public key file written to: %s", rsaPubFilePath))

				rsaPrivateFilePath := getOutputFilePath(outputPath, outputFile, "key")
				os.WriteFile(rsaPrivateFilePath, []byte(o.RsaPrivateKey), os.ModePerm)
				logger.Info(fmt.Sprintf("RSA Private key file written to: %s", rsaPrivateFilePath))
			}
		} else {
			fmt.Println("JWK Private Key:")
			fmt.Println(o.JsonWebKeyString)
			fmt.Println("")
			fmt.Println("---------------------------")
			fmt.Println("")
			fmt.Println("JWK Public Key:")
			fmt.Println(o.JsonWebKeyPublicString)

			if outputBase64 {
				fmt.Println("")
				fmt.Println("---------------------------")
				fmt.Println("")
				fmt.Println("Base64 Encoded JWK Private Key:")
				fmt.Println(o.Base64JsonWebKey)
				fmt.Println("")
			}
			if outputRsaKeys {
				fmt.Println("---------------------------")
				fmt.Println("")
				fmt.Println("PEM Encoded RSA Private Key:")
				fmt.Println(o.RsaPrivateKey)
				fmt.Println("")
				fmt.Println("---------------------------")
				fmt.Println("")
				fmt.Println("PEM Encoded RSA Public Key:")
				fmt.Println(o.RsaPublicKey)
				fmt.Println("")
			}
		}

		return nil
	},
}

func getOutputFilePath(outputPath string, outputFile string, ext string) string {
	return fmt.Sprintf("%s/%s.%s", outputPath, outputFile, ext)
}
