package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
)

func init() {
	createJwkCmd.Flags().BoolP(OutputBase64Flag, "b", false, "Whether to output the JWK as a Base64 string.")
	createJwkCmd.Flags().BoolP(OutputPemKeysFlag, "p", false, "Whether to output the JWK RSA private/public key pair as PEM encoded strings.")
	createJwkCmd.Flags().BoolP(OutputJWKSFlag, "j", false, "Whether to output all public keys combined as a JWKS (JSON Web Key Set).")
	createJwkCmd.Flags().StringP(OutputPathFlag, "o", "", "The path to write the JWK output to. Will withhold output from the console when specified.")
	createJwkCmd.Flags().StringP(OutputFileFlag, "f", OutputFileDefault, "The name of the file(s) to write to.")
	createJwkCmd.Flags().IntP(CountFlag, "c", 1, "The count to create.")
	createJwkCmd.Flags().StringP(KeyTypeFlag, "k", JwkAlgorithmRsa2048, "The key type to use. Options are 'RSA-2048', 'RSA-4096', 'ECDSA-P256', 'ECDSA-P384', 'ECDSA-P521', 'OKP-Ed25519', 'OKP-X25519', 'HS256', 'HS384', 'HS512'.")
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
		if err := viper.BindPFlag(OutputJWKSFlag, cmd.Flags().Lookup(OutputJWKSFlag)); err != nil {
			return err
		}
		if err := viper.BindPFlag(OutputFileFlag, cmd.Flags().Lookup(OutputFileFlag)); err != nil {
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
		outputJWKS := viper.GetBool(OutputJWKSFlag)
		outputPath := viper.GetString(OutputPathFlag)
		outputFile := viper.GetString(OutputFileFlag)
		count := viper.GetInt(CountFlag)
		jwkAlgorithm := viper.GetString(KeyTypeFlag)

		jwkc, err := buildJWKCreator(jwkAlgorithm)
		if err != nil {
			return err
		}

		jwkWriter, err := buildJWKWriter(outputPath, outputFile, outputBase64, outputPemKeys)
		if err != nil {
			return err
		}

		var outputs []*internaljwk.JWKOutput
		for i := 1; i <= count; i++ {
			o, err := jwkc.Create()
			if err != nil {
				return err
			}
			fmt.Println("New JWK created successfully. Will now write to output.")

			if err := jwkWriter.Write(o, i); err != nil {
				return err
			}
			outputs = append(outputs, o)
		}

		if outputJWKS {
			jwksJSON, err := buildJWKSJSON(outputs)
			if err != nil {
				return fmt.Errorf("failed to build JWKS: %w", err)
			}
			if len(outputPath) > 0 {
				dest := filepath.Join(outputPath, outputFile+".jwks")
				if err := os.WriteFile(dest, jwksJSON, 0644); err != nil {
					return fmt.Errorf("failed to write JWKS file: %w", err)
				}
				fmt.Printf("JWKS written to: %s\n", dest)
			} else {
				fmt.Printf("\nJWKS (Public Keys):\n%s\n", jwksJSON)
			}
		}

		return nil
	},
}

// buildJWKSJSON constructs a JSON Web Key Set from a slice of JWKOutput.
// For asymmetric keys the public key is included; for symmetric keys the full key is used.
func buildJWKSJSON(outputs []*internaljwk.JWKOutput) ([]byte, error) {
	set := jwk.NewSet()
	for _, o := range outputs {
		if o.JWKPublic != nil {
			if err := set.AddKey(o.JWKPublic); err != nil {
				return nil, err
			}
		} else if o.JWK != nil {
			if err := set.AddKey(o.JWK); err != nil {
				return nil, err
			}
		}
	}
	return json.MarshalIndent(set, "", "  ")
}

