package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/local"
	"github.com/mkt36z/cli/internal/ui"
)

var planCmd = &cobra.Command{
	Use:   "plan <command>",
	Short: "Create and manage 90-day marketing plans",
	Long: `Generate AI-powered marketing plans with milestones and track progress.

  Commands:
    create    Generate a new 90-day marketing plan
    show      Display an existing plan
    progress  Show execution progress against a plan

  Examples:
    mkt36z plan create "Launch AI CRM to mid-market SaaS"
    mkt36z plan show abc123
    mkt36z plan progress abc123`,
}

var (
	planDuration   int
	planFocusAreas []string
)

var planCreateCmd = &cobra.Command{
	Use:   "create <objective>",
	Short: "Generate a new marketing plan",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		objective := strings.Join(args, " ")

		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		userCtx := loadUserContext()

		reqBody := map[string]interface{}{
			"objective":    objective,
			"duration_days": planDuration,
		}
		if len(planFocusAreas) > 0 {
			reqBody["focus_areas"] = planFocusAreas
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner(fmt.Sprintf("Creating %d-day marketing plan...", planDuration))
		}

		resp, err := client.Do(context.Background(), "POST", "/api/v1/plan/create", bytes.NewReader(bodyJSON))
		if err != nil {
			if sp != nil {
				sp.Stop("")
			}
			return err
		}

		var content strings.Builder
		for event := range api.StreamSSE(context.Background(), resp) {
			switch event.Type {
			case api.EventContentChunk:
				var chunk struct {
					Content string `json:"content"`
				}
				if json.Unmarshal([]byte(event.Data), &chunk) == nil {
					content.WriteString(chunk.Content)
				}
			case api.EventAgentProgress:
				if sp != nil {
					var progress struct {
						Message string `json:"message"`
					}
					if json.Unmarshal([]byte(event.Data), &progress) == nil {
						sp.UpdateMessage(progress.Message)
					}
				}
			case api.EventError:
				if sp != nil {
					sp.Stop("")
				}
				var errData struct {
					Error string `json:"error"`
				}
				if json.Unmarshal([]byte(event.Data), &errData) == nil {
					return fmt.Errorf("plan error: %s", errData.Error)
				}
			case api.EventDone:
				// Complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Plan complete")
		}

		_ = local.AppendHistory(local.HistoryEntry{
			Command: "plan create",
			Input:   objective,
			Output:  content.String(),
		})

		if flagJSON {
			result := map[string]interface{}{
				"objective": objective,
				"duration":  planDuration,
				"content":   content.String(),
			}
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
		} else {
			rendered, err := ui.RenderMarkdown(content.String())
			if err != nil {
				fmt.Print(content.String())
			} else {
				fmt.Print(rendered)
			}
		}

		return nil
	},
}

var planShowCmd = &cobra.Command{
	Use:   "show <plan-id>",
	Short: "Display an existing plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/plan/%s", args[0])
		resp, err := client.Do(context.Background(), "GET", path, nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		out, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(out))
		return nil
	},
}

var planProgressCmd = &cobra.Command{
	Use:   "progress <plan-id>",
	Short: "Show execution progress against a plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/plan/%s/progress", args[0])
		resp, err := client.Do(context.Background(), "GET", path, nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		// Human-readable progress display
		if title, ok := result["title"].(string); ok {
			fmt.Printf("Plan: %s\n", title)
		}
		if pct, ok := result["progress_pct"].(float64); ok {
			bar := progressBar(pct, 30)
			fmt.Printf("Progress: %s %.0f%%\n", bar, pct)
		}
		if milestones, ok := result["milestones"].([]interface{}); ok {
			fmt.Printf("\nMilestones (%d):\n", len(milestones))
			for _, m := range milestones {
				if ms, ok := m.(map[string]interface{}); ok {
					status := "  "
					if done, ok := ms["completed"].(bool); ok && done {
						status = "x"
					}
					fmt.Printf("  [%s] %s\n", status, ms["title"])
				}
			}
		}
		if alerts, ok := result["alerts"].([]interface{}); ok && len(alerts) > 0 {
			fmt.Printf("\nAlerts:\n")
			for _, a := range alerts {
				fmt.Printf("  ! %s\n", a)
			}
		}

		return nil
	},
}

func progressBar(pct float64, width int) string {
	filled := int(pct / 100 * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	return strings.Repeat("#", filled) + strings.Repeat("-", width-filled)
}

func init() {
	planCreateCmd.Flags().IntVar(&planDuration, "duration", 90, "Plan duration in days (30, 60, or 90)")
	planCreateCmd.Flags().StringSliceVar(&planFocusAreas, "focus", nil, "Focus areas (e.g. content,seo,paid)")

	planCmd.AddCommand(planCreateCmd)
	planCmd.AddCommand(planShowCmd)
	planCmd.AddCommand(planProgressCmd)
	rootCmd.AddCommand(planCmd)
}
