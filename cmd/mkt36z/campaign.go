package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/local"
	"github.com/mkt36z/cli/internal/ui"
)

var campaignCmd = &cobra.Command{
	Use:   "campaign <command>",
	Short: "Manage multi-agent marketing campaigns",
	Long: `Create and manage multi-agent campaigns that chain multiple agents together.

  Commands:
    create    Create a new campaign from a brief
    list      List recent campaigns
    status    Check campaign status
    approve   Approve a campaign draft
    reject    Reject and request revisions
    export    Export campaign to files
    resume    Resume a paused campaign

  Examples:
    mkt36z campaign create "Q1 product launch for AI CRM"
    mkt36z campaign list
    mkt36z campaign status abc123
    mkt36z campaign approve abc123
    mkt36z campaign export abc123 --format markdown`,
}

var campaignCreateCmd = &cobra.Command{
	Use:   "create <brief>",
	Short: "Create a new multi-agent campaign",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		brief := strings.Join(args, " ")

		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		userCtx := loadUserContext()

		agents := []string{"strategist", "writer", "channel", "qa"}
		reqBody := map[string]interface{}{
			"brief":  brief,
			"agents": agents,
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner("Creating campaign...")
		}

		resp, err := client.Do(context.Background(), "POST", "/api/v1/campaigns", bytes.NewReader(bodyJSON))
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
					return fmt.Errorf("campaign error: %s", errData.Error)
				}
			case api.EventDone:
				// Complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Campaign complete")
		}

		_ = local.AppendHistory(local.HistoryEntry{
			Command: "campaign create",
			Input:   brief,
			Output:  content.String(),
		})

		if flagJSON {
			result := map[string]interface{}{
				"brief":   brief,
				"content": content.String(),
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

var campaignListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent campaigns",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		resp, err := client.Do(context.Background(), "GET", "/api/v1/campaigns", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result struct {
			Campaigns []struct {
				ID        string `json:"id"`
				Brief     string `json:"brief"`
				Status    string `json:"status"`
				CreatedAt string `json:"created_at"`
			} `json:"campaigns"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(result.Campaigns) == 0 {
			fmt.Println("No campaigns found. Create one with: mkt36z campaign create \"your brief\"")
			return nil
		}

		for _, c := range result.Campaigns {
			fmt.Printf("  %s  %-12s  %s\n", c.ID[:8], c.Status, c.Brief)
		}

		return nil
	},
}

var campaignStatusCmd = &cobra.Command{
	Use:   "status <campaign-id>",
	Short: "Check campaign status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/campaigns/%s", args[0])
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

var (
	campaignExportFormat string
)

var campaignApproveCmd = &cobra.Command{
	Use:   "approve <campaign-id>",
	Short: "Approve a campaign draft",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/campaigns/%s/approve", args[0])
		body := bytes.NewReader([]byte(`{"action":"approve"}`))
		resp, err := client.Do(context.Background(), "POST", path, body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		fmt.Fprintf(os.Stderr, "Campaign %s approved.\n", args[0])
		return nil
	},
}

var campaignRejectCmd = &cobra.Command{
	Use:   "reject <campaign-id>",
	Short: "Reject and request revisions",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/campaigns/%s/reject", args[0])
		body := bytes.NewReader([]byte(`{"action":"reject"}`))
		resp, err := client.Do(context.Background(), "POST", path, body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		fmt.Fprintf(os.Stderr, "Campaign %s rejected. Revisions requested.\n", args[0])
		return nil
	},
}

var campaignExportCmd = &cobra.Command{
	Use:   "export <campaign-id>",
	Short: "Export campaign to files",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/campaigns/%s/export?format=%s", args[0], campaignExportFormat)
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

var campaignResumeCmd = &cobra.Command{
	Use:   "resume <campaign-id>",
	Short: "Resume a paused campaign",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewStreamingClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/campaigns/%s/resume", args[0])

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner(fmt.Sprintf("Resuming campaign %s...", args[0]))
		}

		resp, err := client.Do(context.Background(), "POST", path, bytes.NewReader([]byte("{}")))
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
			case api.EventDone:
				// Complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Campaign resumed")
		}
		if !flagJSON {
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

func init() {
	campaignExportCmd.Flags().StringVar(&campaignExportFormat, "format", "markdown", "Export format (markdown, json, html)")

	campaignCmd.AddCommand(campaignCreateCmd)
	campaignCmd.AddCommand(campaignListCmd)
	campaignCmd.AddCommand(campaignStatusCmd)
	campaignCmd.AddCommand(campaignApproveCmd)
	campaignCmd.AddCommand(campaignRejectCmd)
	campaignCmd.AddCommand(campaignExportCmd)
	campaignCmd.AddCommand(campaignResumeCmd)
	rootCmd.AddCommand(campaignCmd)
}
