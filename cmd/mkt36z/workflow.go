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

var workflowCmd = &cobra.Command{
	Use:   "workflow <command>",
	Short: "Access and run marketing workflows",
	Long: `Browse and execute marketing workflows.

  Commands:
    list   List available workflows
    show   Show workflow contents
    run    Execute a workflow with AI agents

  Examples:
    mkt36z workflow list
    mkt36z workflow show content-calendar
    mkt36z workflow run content-calendar "Q1 2025 blog strategy"`,
}

var workflowListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available workflows",
	RunE: func(cmd *cobra.Command, args []string) error {
		names, err := local.ListWorkflows()
		if err != nil {
			return fmt.Errorf("failed to list workflows: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(map[string]interface{}{
				"workflows": names,
				"count":     len(names),
			}, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(names) == 0 {
			fmt.Println("No workflows found.")
			return nil
		}

		fmt.Println("Available workflows:")
		for _, name := range names {
			fmt.Printf("  - %s\n", name)
		}
		return nil
	},
}

var workflowShowCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show workflow contents",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := local.GetWorkflow(args[0])
		if err != nil {
			return err
		}

		if flagJSON {
			out, _ := json.MarshalIndent(map[string]interface{}{
				"name":    args[0],
				"content": content,
			}, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		fmt.Println(content)
		return nil
	},
}

var workflowRunCmd = &cobra.Command{
	Use:   "run <name> <topic>",
	Short: "Execute a workflow with AI agents",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		workflowName := args[0]
		topic := strings.Join(args[1:], " ")

		content, err := local.GetWorkflow(workflowName)
		if err != nil {
			return err
		}

		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		userCtx := loadUserContext()

		reqBody := map[string]interface{}{
			"topic":    topic,
			"workflow": content,
			"type":     "brief",
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner(fmt.Sprintf("Running workflow: %s...", workflowName))
		}

		resp, err := client.Do(context.Background(), "POST", "/api/v1/workflows/run", bytes.NewReader(bodyJSON))
		if err != nil {
			if sp != nil {
				sp.Stop("")
			}
			return err
		}

		var output strings.Builder
		for event := range api.StreamSSE(context.Background(), resp) {
			switch event.Type {
			case api.EventContentChunk:
				var chunk struct {
					Content string `json:"content"`
				}
				if json.Unmarshal([]byte(event.Data), &chunk) == nil {
					output.WriteString(chunk.Content)
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
					return fmt.Errorf("workflow error: %s", errData.Error)
				}
			case api.EventDone:
				// Complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Workflow complete")
		}

		_ = local.AppendHistory(local.HistoryEntry{
			Command: "workflow run " + workflowName,
			Input:   topic,
			Output:  output.String(),
		})

		if flagJSON {
			result := map[string]interface{}{
				"workflow": workflowName,
				"topic":    topic,
				"content":  output.String(),
			}
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
		} else {
			rendered, err := ui.RenderMarkdown(output.String())
			if err != nil {
				fmt.Print(output.String())
			} else {
				fmt.Print(rendered)
			}
		}

		return nil
	},
}

func init() {
	workflowCmd.AddCommand(workflowListCmd)
	workflowCmd.AddCommand(workflowShowCmd)
	workflowCmd.AddCommand(workflowRunCmd)
	rootCmd.AddCommand(workflowCmd)
}
