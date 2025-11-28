package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createJwkCmd.Flags().BoolP(OutputBase64Flag, "b", false, "Whether to output the JWK as a Base64 string.")
	createJwkCmd.Flags().BoolP(OutputPemKeysFlag, "p", false, "Whether to output the JWK RSA private/public key pair as PEM encoded strings.")
	createJwkCmd.Flags().StringP(OutputPathFlag, "o", "", "The path to write the JWK output to. Will withhold output from the console when specified.")
	createJwkCmd.Flags().StringP(OutputFileFlag, "f", OutputFileDefault, "The name of the file(s) to write to.")
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

		for i := 1; i <= count; i++ {
			o, err := jwkc.Create()
			if err != nil {
				return err
			}
			fmt.Println("New JWK created successfully. Will now write to output.")

			err = jwkWriter.Write(o, i)
			if err != nil {
				return err
			}
		}

		return nil
	},
}
