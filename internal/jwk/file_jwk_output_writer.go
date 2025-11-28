package jwk

import (
	"fmt"
	"os"
)

type FileJwkOutputWriter struct {
	outputPath    string
	outputFile    string
	outputBase64  bool
	outputPemKeys bool
}

func NewFileJwkOutputWriter(outputPath string, outputFile string, outputBase64 bool, outputPemKeys bool) *FileJwkOutputWriter {
	return &FileJwkOutputWriter{
		outputPath:    outputPath,
		outputFile:    outputFile,
		outputBase64:  outputBase64,
		outputPemKeys: outputPemKeys,
	}
}

func (w *FileJwkOutputWriter) Write(o *JWKOutput, i int) error {
	jwkFilePath := getOutputFilePath(w.outputPath, w.outputFile, "jwk", i)
	if err := os.WriteFile(jwkFilePath, []byte(o.JWKString), 0600); err != nil {
		return err
	}
	fmt.Printf("JWK Private key file written to: %s\n", jwkFilePath)

	jwkPubFilePath := getOutputFilePath(w.outputPath, fmt.Sprintf("%s-pub", w.outputFile), "jwk", i)
	if err := os.WriteFile(jwkPubFilePath, []byte(o.JWKPublicString), 0644); err != nil {
		return err
	}
	fmt.Printf("JWK Public key file written to: %s\n", jwkPubFilePath)

	if w.outputBase64 {
		base64FilePath := getOutputFilePath(w.outputPath, fmt.Sprintf("%s-base64", w.outputFile), "jwk", i)
		if err := os.WriteFile(base64FilePath, []byte(o.Base64JWK), 0600); err != nil {
			return err
		}
		fmt.Printf("Base64 JWK Private key file written to: %s\n", base64FilePath)
	}
	if w.outputPemKeys {
		rsaPubFilePath := getOutputFilePath(w.outputPath, w.outputFile, "pub", i)
		if err := os.WriteFile(rsaPubFilePath, []byte(o.PEMPublicKey), 0644); err != nil {
			return err
		}
		fmt.Printf("RSA Public key file written to: %s\n", rsaPubFilePath)

		rsaPrivateFilePath := getOutputFilePath(w.outputPath, w.outputFile, "key", i)
		if err := os.WriteFile(rsaPrivateFilePath, []byte(o.PEMPrivateKey), 0600); err != nil {
			return err
		}
		fmt.Printf("RSA Private key file written to: %s\n", rsaPrivateFilePath)
	}
	return nil
}

func getOutputFilePath(outputPath string, outputFile string, ext string, i int) string {
	return fmt.Sprintf("%s/%s-%d.%s", outputPath, outputFile, i, ext)
}
