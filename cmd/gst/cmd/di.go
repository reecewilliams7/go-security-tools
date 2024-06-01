package cmd

import (
	"io"
	"os"

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
