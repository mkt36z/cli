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
)

var playbookCmd = &cobra.Command{
	Use:   "playbook <command>",
	Short: "Access and run marketing playbooks",
	Long: `Browse and execute marketing playbooks.

  Commands:
    list   List available playbooks
    show   Show playbook contents
    run    Execute a playbook with AI agents

  Examples:
    mkt36z playbook list
    mkt36z playbook show product-launch
    mkt36z playbook run product-launch "AI CRM for startups"`,
}

var playbookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available playbooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		names, err := local.ListPlaybooks()
		if err != nil {
			return fmt.Errorf("failed to list playbooks: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(map[string]interface{}{
				"playbooks": names,
				"count":     len(names),
			}, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(names) == 0 {
			fmt.Println("No playbooks found.")
			fmt.Println("\n  Add custom playbooks to ~/.mkt36z/playbooks/")
			return nil
		}

		fmt.Println("Available playbooks:")
		for _, name := range names {
			fmt.Printf("  - %s\n", name)
		}
		return nil
	},
}

var playbookShowCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show playbook contents",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := local.GetPlaybook(args[0])
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

var playbookRunCmd = &cobra.Command{
	Use:   "run <name> <topic>",
	Short: "Execute a playbook with AI agents",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		playbookName := args[0]
		topic := strings.Join(args[1:], " ")

		// Load playbook content
		content, err := local.GetPlaybook(playbookName)
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
			"playbook": content,
			"type":     "brief",
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)
		fmt.Fprintf(os.Stderr, "Running playbook: %s...\n", playbookName)

		resp, err := client.Do(context.Background(), "POST", "/api/v1/playbooks/run", bytes.NewReader(bodyJSON))
		if err != nil {
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
					if !flagJSON {
						fmt.Print(chunk.Content)
					}
				}
			case api.EventAgentProgress:
				if !flagJSON {
					var progress struct {
						Message string `json:"message"`
					}
					if json.Unmarshal([]byte(event.Data), &progress) == nil {
						fmt.Fprintf(os.Stderr, "\r%s", progress.Message)
					}
				}
			case api.EventError:
				var errData struct {
					Error string `json:"error"`
				}
				if json.Unmarshal([]byte(event.Data), &errData) == nil {
					return fmt.Errorf("playbook error: %s", errData.Error)
				}
			case api.EventDone:
				// Complete
			}
		}

		fmt.Fprintln(os.Stderr)

		_ = local.AppendHistory(local.HistoryEntry{
			Command: "playbook run " + playbookName,
			Input:   topic,
			Output:  output.String(),
		})

		if flagJSON {
			result := map[string]interface{}{
				"playbook": playbookName,
				"topic":    topic,
				"content":  output.String(),
			}
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
		}

		return nil
	},
}

func init() {
	playbookCmd.AddCommand(playbookListCmd)
	playbookCmd.AddCommand(playbookShowCmd)
	playbookCmd.AddCommand(playbookRunCmd)
	rootCmd.AddCommand(playbookCmd)
}
