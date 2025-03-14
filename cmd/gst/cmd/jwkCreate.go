package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jwks "github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
)

type JsonWebKeyCreator interface {
	Create() (*jwks.JsonWebKeyOutput, error)
}

func init() {
	createJwkCmd.Flags().BoolP(OutputBase64FlagName, "b", false, "Whether to output the JWK as a Base64 string.")
	createJwkCmd.Flags().BoolP(OutputRsaKeysFlagName, "r", false, "Whether to output the JWK RSA private/public key pair as PEM encoded strings.")
	createJwkCmd.Flags().StringP(OutputPathFlagName, "p", "", "The path to write the JWK output to. Will withhold output from the console when specified.")
	createJwkCmd.Flags().StringP(OutputFileNameFlagName, "f", OutputFileDefaultName, "The name of the file(s) to write to.")
	createJwkCmd.Flags().IntP(CountFlagName, "c", 1, "The count to create.")

	jwkCmd.AddCommand(createJwkCmd)
}

var createJwkCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a JSON Web Key",
	Long:  "Creates a JSON Web Key using an RSA Private Key",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag(OutputBase64FlagName, cmd.Flags().Lookup(OutputBase64FlagName))
		viper.BindPFlag(OutputPathFlagName, cmd.Flags().Lookup(OutputPathFlagName))
		viper.BindPFlag(OutputRsaKeysFlagName, cmd.Flags().Lookup(OutputRsaKeysFlagName))
		viper.BindPFlag(OutputFileNameFlagName, cmd.Flags().Lookup(OutputFileNameFlagName))
		viper.BindPFlag(CountFlagName, cmd.Flags().Lookup(CountFlagName))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		outputBase64 := viper.GetBool(OutputBase64FlagName)
		outputRsaKeys := viper.GetBool(OutputRsaKeysFlagName)
		outputPath := viper.GetString(OutputPathFlagName)
		outputFile := viper.GetString(OutputFileNameFlagName)
		count := viper.GetInt(CountFlagName)

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

		for range count {
			o, err := jwkc.Create()
			logger.Info("New JWK created successfully. Will now write to output.")
			if err != nil {
				return err
			}

			if writeToFile {
				jwkFilePath := getOutputFilePath(outputPath, outputFile, "jwk", count)
				os.WriteFile(jwkFilePath, []byte(o.JsonWebKeyString), os.ModePerm)
				logger.Info(fmt.Sprintf("JWK Private key file written to: %s", jwkFilePath))

				jwkPubFilePath := getOutputFilePath(outputPath, fmt.Sprintf("%s-pub", outputFile), "jwk", count)
				os.WriteFile(jwkPubFilePath, []byte(o.JsonWebKeyPublicString), os.ModePerm)
				logger.Info(fmt.Sprintf("JWK Public key file written to: %s", jwkPubFilePath))

				if outputBase64 {
					base64FilePath := getOutputFilePath(outputPath, fmt.Sprintf("%s-base64", outputFile), "jwk", count)
					os.WriteFile(base64FilePath, []byte(o.Base64JsonWebKey), os.ModePerm)
					logger.Info(fmt.Sprintf("Base64 JWK Private key file written to: %s", base64FilePath))
				}
				if outputRsaKeys {
					rsaPubFilePath := getOutputFilePath(outputPath, outputFile, "pub", count)
					os.WriteFile(rsaPubFilePath, []byte(o.RsaPublicKey), os.ModePerm)
					logger.Info(fmt.Sprintf("RSA Public key file written to: %s", rsaPubFilePath))

					rsaPrivateFilePath := getOutputFilePath(outputPath, outputFile, "key", count)
					os.WriteFile(rsaPrivateFilePath, []byte(o.RsaPrivateKey), os.ModePerm)
					logger.Info(fmt.Sprintf("RSA Private key file written to: %s", rsaPrivateFilePath))
				}
			} else {
				fmt.Println("**********************************************************")
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
				fmt.Println("**********************************************************")
			}
		}

		return nil
	},
}

func getOutputFilePath(outputPath string, outputFile string, ext string, count int) string {
	if count > 1 {
		return fmt.Sprintf("%s/%s-%d.%s", outputPath, outputFile, count, ext)
	}

	return fmt.Sprintf("%s/%s.%s", outputPath, outputFile, ext)
}
