package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "development"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the gst version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gst version: %s\n", version)
	},
}
