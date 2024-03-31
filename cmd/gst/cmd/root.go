package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//"github.com/spf13/cobra/doc"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gst",
		Short: "A CLI tool containing various security related functions",
		Long:  "TODO",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	rootCmd.PersistentFlags().StringP(LogLevelFlagName, "l", "Info", "The logging level to use - 'Info', 'Debug', 'Warn', 'Error'")

	viper.BindPFlag(LogLevelFlagName, rootCmd.PersistentFlags().Lookup(LogLevelFlagName))

	viper.SetEnvPrefix("GST")
	viper.BindEnv(LogLevelFlagName)
}
