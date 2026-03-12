package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
)

var usageCmd = &cobra.Command{
	Use:   "usage <command>",
	Short: "View usage, billing, and plan information",
	Long: `Track your generation usage, view history, and manage your plan.

  Plans:
    Free              100 generations/mo
    Starter ($29/mo)  500 generations/mo
    Growth  ($59/mo)  2,000 generations/mo
    Scale   ($149/mo) 10,000 generations/mo + $0.01/overage

  Commands:
    show      Current period usage and plan details
    history   Monthly usage history
    upgrade   Open the upgrade page in your browser

  Examples:
    mkt36z usage show
    mkt36z usage history
    mkt36z usage upgrade`,
}

var usageShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Current period usage and plan details",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		resp, err := client.Do(context.Background(), "GET", "/api/v1/usage", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result struct {
			Plan                string `json:"plan"`
			PlanLabel           string `json:"plan_label"`
			Price               int    `json:"price"`
			GenerationsUsed     int    `json:"generations_used"`
			GenerationsLimit    int    `json:"generations_limit"`
			GenerationsRemaining int   `json:"generations_remaining"`
			OverageCents        int    `json:"overage_cents"`
			ResetDate           string `json:"reset_date"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		fmt.Printf("Plan: %s\n\n", result.PlanLabel)

		// Generations usage bar
		pct := 0.0
		if result.GenerationsLimit > 0 {
			pct = float64(result.GenerationsUsed) / float64(result.GenerationsLimit) * 100
		}
		fmt.Printf("  Generations:  %s %d/%d (%.0f%%)\n",
			usageBar(pct, 20), result.GenerationsUsed, result.GenerationsLimit, pct)
		fmt.Printf("  Remaining:    %d generations\n", result.GenerationsRemaining)

		if result.OverageCents > 0 {
			fmt.Printf("  Overage:      $0.%02d per generation beyond limit\n", result.OverageCents)
		}

		fmt.Printf("  Resets:       %s\n", result.ResetDate)

		// Proactive warning at 80%
		if pct >= 80 {
			fmt.Fprintf(os.Stderr, "\n  You've used %.0f%% of your monthly generations.\n  Upgrade your plan: mkt36z usage upgrade\n", pct)
		}

		return nil
	},
}

var usageHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Monthly usage history",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		resp, err := client.Do(context.Background(), "GET", "/api/v1/usage/history", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result struct {
			History []struct {
				Month string `json:"month"`
				Count int    `json:"count"`
				Plan  string `json:"plan"`
			} `json:"history"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(result.History) == 0 {
			fmt.Println("No usage history yet.")
			return nil
		}

		fmt.Printf("%-12s  %-8s  %s\n", "Month", "Used", "Plan")
		fmt.Printf("%-12s  %-8s  %s\n", "-----", "----", "----")
		for _, h := range result.History {
			fmt.Printf("%-12s  %-8d  %s\n", h.Month, h.Count, h.Plan)
		}

		return nil
	},
}

var usageUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Open the upgrade page in your browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := "https://mkt36z.com/upgrade"

		var openCmd string
		var openArgs []string

		switch runtime.GOOS {
		case "darwin":
			openCmd = "open"
			openArgs = []string{url}
		case "windows":
			openCmd = "cmd"
			openArgs = []string{"/c", "start", url}
		default:
			openCmd = "xdg-open"
			openArgs = []string{url}
		}

		if err := exec.Command(openCmd, openArgs...).Start(); err != nil {
			// Fallback: print the URL
			fmt.Printf("Visit: %s\n", url)
			return nil
		}

		fmt.Fprintf(os.Stderr, "Opening upgrade page...\n")
		return nil
	},
}

func usageBar(pct float64, width int) string {
	filled := int(pct / 100 * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	return "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
}

func init() {
	usageCmd.AddCommand(usageShowCmd)
	usageCmd.AddCommand(usageHistoryCmd)
	usageCmd.AddCommand(usageUpgradeCmd)
	rootCmd.AddCommand(usageCmd)
}
