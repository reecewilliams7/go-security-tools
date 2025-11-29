package cmd

import (
	"fmt"
	"os"

	"github.com/reecewilliams7/go-security-tools/clientcredentials"
	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
	"github.com/reecewilliams7/go-security-tools/jwk"
)

func buildClientCredentialsCreator(clientIDType string, clientSecretType string) (*clientcredentials.ClientCredentialsCreator, error) {
	var clientIDCreator clientcredentials.ClientIDCreator
	var clientSecretCreator clientcredentials.ClientSecretCreator

	switch clientIDType {
	case ClientIdTypeUUIDv7:
		clientIDCreator = clientcredentials.NewUUIDv7ClientIDCreator()
	case ClientIdTypeShort:
		clientIDCreator = clientcredentials.NewShortUUIDClientIDCreator()
	default:
		return &clientcredentials.ClientCredentialsCreator{}, fmt.Errorf("unknown client id type: %s", clientIDType)
	}

	switch clientSecretType {
	case ClientSecretTypeCryptoRand:
		clientSecretCreator = clientcredentials.NewCryptoRandClientSecretCreator()
	default:
		return &clientcredentials.ClientCredentialsCreator{}, fmt.Errorf("unknown client secret type: %s", clientSecretType)
	}

	ccc := clientcredentials.NewClientCredentialsCreator(clientIDCreator, clientSecretCreator)

	return ccc, nil
}

func buildJWKCreator(jwkAlgorithm string) (jwk.JWKCreator, error) {
	switch jwkAlgorithm {
	case JwkAlgorithmRsa2048:
		return jwk.NewRSAJSONWebKeyCreator(2048), nil
	case JwkAlgorithmRsa4096:
		return jwk.NewRSAJSONWebKeyCreator(4096), nil
	case JwkAlgorithmEcdsaP256:
		return jwk.NewECDSAJWKCreator("P256"), nil
	case JwkAlgorithmEcdsaP384:
		return jwk.NewECDSAJWKCreator("P384"), nil
	case JwkAlgorithmEcdsaP521:
		return jwk.NewECDSAJWKCreator("P521"), nil
	default:
		return nil, fmt.Errorf("unknown JWK algorithm: %s", jwkAlgorithm)
	}
}

func buildJWKWriter(outputPath string, outputFile string, outputBase64 bool, outputPemKeys bool) (internaljwk.JWKOutputWriter, error) {
	writeToFile := false

	if len(outputPath) > 0 {
		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			return nil, err
		}
		fmt.Printf("Output path specified so will write JWK files to following location: %s\n", outputPath)
		writeToFile = true
	}

	if writeToFile {
		return internaljwk.NewFileJwkOutputWriter(outputPath, outputFile, outputBase64, outputPemKeys), nil
	}
	return internaljwk.NewFmtJWKOutputWriter(outputBase64, outputPemKeys), nil
}
