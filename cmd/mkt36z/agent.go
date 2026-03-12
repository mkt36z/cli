package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/local"
	"github.com/mkt36z/cli/internal/ui"
)

var agentCmd = &cobra.Command{
	Use:   "agent <command>",
	Short: "Run and manage individual AI agents",
	Long: `Run individual agents directly or view agent information.

  Commands:
    run       Run a specific agent directly
    list      List all available agents
    status    Check agent run status
    history   View agent run history

  Examples:
    mkt36z agent list
    mkt36z agent run strategist "AI-powered CRM for startups"
    mkt36z agent run social "product launch" --channel linkedin
    mkt36z agent run qa "review this content"
    mkt36z agent history`,
}

var (
	agentChannel string
	agentType    string
)

var agentRunCmd = &cobra.Command{
	Use:   "run <agent-name> <topic>",
	Short: "Run a specific agent directly",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		agentName := args[0]
		topic := strings.Join(args[1:], " ")

		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		userCtx := loadUserContext()

		reqBody := map[string]interface{}{
			"topic": topic,
			"options": map[string]interface{}{
				"channel": agentChannel,
				"type":    agentType,
			},
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/agents/%s/run", agentName)

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner(fmt.Sprintf("Running %s agent...", agentName))
		}

		resp, err := client.Do(context.Background(), "POST", path, bytes.NewReader(bodyJSON))
		if err != nil {
			if sp != nil {
				sp.Stop("")
			}
			return err
		}

		var content strings.Builder
		var qaScore map[string]interface{}

		for event := range api.StreamSSE(context.Background(), resp) {
			switch event.Type {
			case api.EventContentChunk:
				var chunk struct {
					Content string `json:"content"`
				}
				if json.Unmarshal([]byte(event.Data), &chunk) == nil {
					content.WriteString(chunk.Content)
				}
			case api.EventQAScore:
				if json.Unmarshal([]byte(event.Data), &qaScore) != nil {
					qaScore = nil
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
					return fmt.Errorf("agent error: %s", errData.Error)
				}
			case api.EventDone:
				// Complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Agent complete")
		}

		_ = local.AppendHistory(local.HistoryEntry{
			Command: "agent run " + agentName,
			Input:   topic,
			Output:  content.String(),
		})

		if flagJSON {
			result := map[string]interface{}{
				"agent":   agentName,
				"topic":   topic,
				"content": content.String(),
			}
			if qaScore != nil {
				result["qa_score"] = qaScore
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

var agentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available agents",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		resp, err := client.Do(context.Background(), "GET", "/api/v1/agents", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result struct {
			Agents []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Label       string `json:"label"`
				Tier        string `json:"tier"`
				Phase       int    `json:"phase"`
				Description string `json:"description"`
			} `json:"agents"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tLABEL\tTIER\tPHASE")
		fmt.Fprintln(w, "--\t----\t-----\t----\t-----")
		for _, a := range result.Agents {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\n", a.ID, a.Name, a.Label, a.Tier, a.Phase)
		}
		w.Flush()

		return nil
	},
}

var agentStatusCmd = &cobra.Command{
	Use:   "status <run-id>",
	Short: "Check agent run status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/agents/runs/%s", args[0])
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

var agentHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "View agent run history",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := local.ReadHistory(20)
		if err != nil {
			return fmt.Errorf("failed to read history: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(entries, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(entries) == 0 {
			fmt.Println("No agent history found.")
			return nil
		}

		for _, e := range entries {
			cmd := e.Command
			input := e.Input
			if len(input) > 60 {
				input = input[:60] + "..."
			}
			qa := ""
			if e.QAScore > 0 {
				qa = fmt.Sprintf("  [QA: %d]", e.QAScore)
			}
			fmt.Printf("  %-30s  %s%s\n", cmd, input, qa)
		}

		return nil
	},
}

func init() {
	agentRunCmd.Flags().StringVar(&agentChannel, "channel", "", "Target channel for agent")
	agentRunCmd.Flags().StringVar(&agentType, "type", "", "Content/analysis type")

	agentCmd.AddCommand(agentRunCmd)
	agentCmd.AddCommand(agentListCmd)
	agentCmd.AddCommand(agentStatusCmd)
	agentCmd.AddCommand(agentHistoryCmd)
	rootCmd.AddCommand(agentCmd)
}
