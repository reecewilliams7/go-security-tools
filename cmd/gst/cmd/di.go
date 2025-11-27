package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/reecewilliams7/go-security-tools/pkg/clientCredentials"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
)

func buildLogger(prefix string) hclog.Logger {
	configLogLevel := viper.GetString(LogLevelFlagName)

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

func buildClientCredentialsCreator(clientIdType string, clientSecretType string) (*clientCredentials.ClientCredentialsCreator, error) {
	var clientIdCreator clientCredentials.ClientIdCreator
	var clientSecretCreator clientCredentials.ClientSecretCreator
	var err error

	switch clientIdType {
	case ClientIdTypeUUIDv7:
		clientIdCreator = clientCredentials.NewUUIDv7ClientIdCreator()
	case ClientIdTypeShort:
		clientIdCreator = clientCredentials.NewShortUuidClientIdCreator()
	default:
		return &clientCredentials.ClientCredentialsCreator{}, fmt.Errorf("unknown client id type: %s", clientIdType)
	}

	switch clientSecretType {
	case ClientSecretTypeCryptoRand:
		clientSecretCreator = clientCredentials.NewCryptoRandClientSecretCreator()
	default:
		return &clientCredentials.ClientCredentialsCreator{}, fmt.Errorf("unknown client secret type: %s", clientSecretType)
	}

	ccc := clientCredentials.NewClientCredentialsCreator(clientIdCreator, clientSecretCreator)

	return ccc, err
}
