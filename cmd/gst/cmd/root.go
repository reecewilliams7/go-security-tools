package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	RootCmd = &cobra.Command{
		Use:   "gst",
		Short: "A CLI tool containing various security related functions",
		Long:  "TODO",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	RootCmd.PersistentFlags().StringP(LogLevelFlagName, "l", "Info", "The logging level to use - 'Info', 'Debug', 'Warn', 'Error'")

	viper.BindPFlag(LogLevelFlagName, RootCmd.PersistentFlags().Lookup(LogLevelFlagName))

	viper.SetEnvPrefix("GST")
	viper.BindEnv(LogLevelFlagName)
}
