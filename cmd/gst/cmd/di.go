package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/reecewilliams7/go-security-tools/pkg/clientcredentials"
	"github.com/reecewilliams7/go-security-tools/pkg/jwk"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
)

func buildLogger(prefix string) hclog.Logger {
	configLogLevel := viper.GetString(LogLevelFlag)

	var writer io.Writer = os.Stderr

	logLevel := hclog.LevelFromString(configLogLevel)

	logger := hclog.New(&hclog.LoggerOptions{
		Level:      logLevel,
		TimeFormat: "2006/01/02 15:04:05",
		Name:       "gst",
		Output:     writer,
	})

	sublogger := logger.Named(prefix)
	return sublogger
}

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

func buildJwkCreator(jwkAlgorithm string) (JSONWebKeyCreator, error) {
	switch jwkAlgorithm {
	case JwkAlgorithmRsa2048:
		return jwk.NewRSAJSONWebKeyCreator(2048), nil
	case JwkAlgorithmRsa4096:
		return jwk.NewRSAJSONWebKeyCreator(4096), nil
	case JwkAlgorithmEcdsaP256:
		return jwk.NewECDSAJSONWebKeyCreator("P256"), nil
	case JwkAlgorithmEcdsaP384:
		return jwk.NewECDSAJSONWebKeyCreator("P384"), nil
	case JwkAlgorithmEcdsaP521:
		return jwk.NewECDSAJSONWebKeyCreator("P521"), nil
	default:
		return nil, fmt.Errorf("unknown JWK algorithm: %s", jwkAlgorithm)
	}
}
