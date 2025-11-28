package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the root command for the gst CLI.
var RootCmd = &cobra.Command{
	Use:   "gst",
	Short: "A CLI tool containing various security related functions",
	Long:  "gst (go-security-tools) is a CLI tool containing various security related functions for OAuth2.0, OpenID Connect, and JWK management.",
}

// Execute runs the root command.
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	viper.SetEnvPrefix("GST")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	RootCmd.PersistentFlags().StringP(LogLevelFlag, "l", "Info", "The logging level to use - 'Info', 'Debug', 'Warn', 'Error'")

	if err := viper.BindPFlag(LogLevelFlag, RootCmd.PersistentFlags().Lookup(LogLevelFlag)); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindEnv(LogLevelFlag); err != nil {
		log.Fatal(err)
	}
}
