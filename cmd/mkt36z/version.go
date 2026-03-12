package main

import (
	"fmt"

	"github.com/mkt36z/cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show CLI version, build info, and runtime details",
	Long: `Display version information for the mkt36z CLI.

Examples:
  mkt36z version            Human-readable version info
  mkt36z version --json     Machine-readable JSON output`,
	Run: func(cmd *cobra.Command, args []string) {
		info := version.Get()
		if flagJSON {
			fmt.Println(info.JSON())
		} else {
			fmt.Println(info.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
