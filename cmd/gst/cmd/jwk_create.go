package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/reecewilliams7/go-security-tools/pkg/jwk"
)

// JSONWebKeyCreator is an interface for creating JSON Web Keys.
type JSONWebKeyCreator interface {
	Create() (*jwk.JSONWebKeyOutput, error)
}

func init() {
	createJwkCmd.Flags().BoolP(OutputBase64Flag, "b", false, "Whether to output the JWK as a Base64 string.")
	createJwkCmd.Flags().BoolP(OutputPemKeysFlag, "p", false, "Whether to output the JWK RSA private/public key pair as PEM encoded strings.")
	createJwkCmd.Flags().StringP(OutputPathFlag, "o", "", "The path to write the JWK output to. Will withhold output from the console when specified.")
	createJwkCmd.Flags().StringP(OutputFileNameFlag, "f", OutputFileDefault, "The name of the file(s) to write to.")
	createJwkCmd.Flags().IntP(CountFlag, "c", 1, "The count to create.")
	createJwkCmd.Flags().StringP(KeyTypeFlag, "k", JwkAlgorithmRsa2048, "The key type to use. Options are 'RSA-2048', 'RSA-4096', 'ECDSA-P256', 'ECDSA-P384', 'ECDSA-P521'.")
	jwkCmd.AddCommand(createJwkCmd)
}

var createJwkCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a JSON Web Key",
	Long:  "Creates a JSON Web Key using an RSA Private Key",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		if err := viper.BindPFlag(OutputBase64Flag, cmd.Flags().Lookup(OutputBase64Flag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(OutputPathFlag, cmd.Flags().Lookup(OutputPathFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(OutputPemKeysFlag, cmd.Flags().Lookup(OutputPemKeysFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(OutputFileNameFlag, cmd.Flags().Lookup(OutputFileNameFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(CountFlag, cmd.Flags().Lookup(CountFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(KeyTypeFlag, cmd.Flags().Lookup(KeyTypeFlag)); err != nil {
			return err
		}
		return nil
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		outputBase64 := viper.GetBool(OutputBase64Flag)
		outputPemKeys := viper.GetBool(OutputPemKeysFlag)
		outputPath := viper.GetString(OutputPathFlag)
		outputFile := viper.GetString(OutputFileNameFlag)
		count := viper.GetInt(CountFlag)
		jwkAlgorithm := viper.GetString(KeyTypeFlag)

		logger := buildLogger("jwk create")

		writeToFile := false

		if len(outputPath) > 0 {
			if _, err := os.Stat(outputPath); os.IsNotExist(err) {
				return err
			}
			logger.Info(fmt.Sprintf("Output path specified so will write JWK files to following location: %s", outputPath))
			writeToFile = true
		}

		jwkc, err := buildJwkCreator(jwkAlgorithm)
		if err != nil {
			return err
		}

		for range count {
			o, err := jwkc.Create()
			if err != nil {
				return err
			}
			logger.Info("New JWK created successfully. Will now write to output.")

			if writeToFile {
				jwkFilePath := getOutputFilePath(outputPath, outputFile, "jwk", count)
				if err := os.WriteFile(jwkFilePath, []byte(o.JSONWebKeyString), 0600); err != nil {
					return err
				}
				logger.Info(fmt.Sprintf("JWK Private key file written to: %s", jwkFilePath))

				jwkPubFilePath := getOutputFilePath(outputPath, fmt.Sprintf("%s-pub", outputFile), "jwk", count)
				if err := os.WriteFile(jwkPubFilePath, []byte(o.JSONWebKeyPublicString), 0644); err != nil {
					return err
				}
				logger.Info(fmt.Sprintf("JWK Public key file written to: %s", jwkPubFilePath))

				if outputBase64 {
					base64FilePath := getOutputFilePath(outputPath, fmt.Sprintf("%s-base64", outputFile), "jwk", count)
					if err := os.WriteFile(base64FilePath, []byte(o.Base64JSONWebKey), 0600); err != nil {
						return err
					}
					logger.Info(fmt.Sprintf("Base64 JWK Private key file written to: %s", base64FilePath))
				}
				if outputPemKeys {
					rsaPubFilePath := getOutputFilePath(outputPath, outputFile, "pub", count)
					if err := os.WriteFile(rsaPubFilePath, []byte(o.AlgPublicKey), 0644); err != nil {
						return err
					}
					logger.Info(fmt.Sprintf("RSA Public key file written to: %s", rsaPubFilePath))

					rsaPrivateFilePath := getOutputFilePath(outputPath, outputFile, "key", count)
					if err := os.WriteFile(rsaPrivateFilePath, []byte(o.AlgPrivateKey), 0600); err != nil {
						return err
					}
					logger.Info(fmt.Sprintf("RSA Private key file written to: %s", rsaPrivateFilePath))
				}
			} else {
				fmt.Println("**********************************************************")
				fmt.Println("JWK Private Key:")
				fmt.Println(o.JSONWebKeyString)
				fmt.Println("")
				fmt.Println("---------------------------")
				fmt.Println("")
				fmt.Println("JWK Public Key:")
				fmt.Println(o.JSONWebKeyPublicString)

				if outputBase64 {
					fmt.Println("")
					fmt.Println("---------------------------")
					fmt.Println("")
					fmt.Println("Base64 Encoded JWK Private Key:")
					fmt.Println(o.Base64JSONWebKey)
					fmt.Println("")
				}
				if outputPemKeys {
					fmt.Println("---------------------------")
					fmt.Println("")
					fmt.Printf("PEM Encoded %s Private Key:\n", o.JSONWebKey.KeyType())
					fmt.Println(o.AlgPrivateKey)
					fmt.Println("")
					fmt.Println("---------------------------")
					fmt.Println("")
					fmt.Printf("PEM Encoded %s Public Key:\n", o.JSONWebKey.KeyType())
					fmt.Println(o.AlgPublicKey)
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
