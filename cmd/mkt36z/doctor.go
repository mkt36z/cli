package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/doctor"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your mkt36z setup",
	Long: `Run diagnostic checks on your mkt36z environment.

  Checks: version, authentication, config, API connectivity,
  context files, embedded assets, and runtime info.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result := doctor.RunAll(flagAPIURL)

		if flagJSON {
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		fmt.Println("mkt36z doctor")
		fmt.Println()

		for _, check := range result.Checks {
			var icon string
			switch check.Status {
			case doctor.StatusOK:
				icon = "✓"
			case doctor.StatusWarn:
				icon = "!"
			case doctor.StatusFail:
				icon = "✗"
			case doctor.StatusSkip:
				icon = "-"
			}
			fmt.Printf("  %s %-14s %s\n", icon, check.Name, check.Message)
		}

		fmt.Println()
		fmt.Println(result.Summary)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
