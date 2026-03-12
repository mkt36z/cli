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

var templateCmd = &cobra.Command{
	Use:   "template <command>",
	Short: "Access and run content templates",
	Long: `Browse and execute content templates.

  Commands:
    list   List available templates
    show   Show template contents
    run    Execute a template with AI agents

  Examples:
    mkt36z template list
    mkt36z template show cold-email
    mkt36z template run cold-email "SaaS sales outreach"`,
}

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		names, err := local.ListTemplates()
		if err != nil {
			return fmt.Errorf("failed to list templates: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(map[string]interface{}{
				"templates": names,
				"count":     len(names),
			}, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(names) == 0 {
			fmt.Println("No templates found.")
			fmt.Println("\n  Add custom templates to ~/.mkt36z/templates/")
			return nil
		}

		fmt.Println("Available templates:")
		for _, name := range names {
			fmt.Printf("  - %s\n", name)
		}
		return nil
	},
}

var templateShowCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show template contents",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := local.GetTemplate(args[0])
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

var templateRunCmd = &cobra.Command{
	Use:   "run <name> <topic>",
	Short: "Execute a template with AI agents",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateName := args[0]
		topic := strings.Join(args[1:], " ")

		content, err := local.GetTemplate(templateName)
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
			"template": content,
			"type":     "copy",
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner(fmt.Sprintf("Running template: %s...", templateName))
		}

		resp, err := client.Do(context.Background(), "POST", "/api/v1/templates/run", bytes.NewReader(bodyJSON))
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
			case api.EventError:
				if sp != nil {
					sp.Stop("")
				}
				var errData struct {
					Error string `json:"error"`
				}
				if json.Unmarshal([]byte(event.Data), &errData) == nil {
					return fmt.Errorf("template error: %s", errData.Error)
				}
			case api.EventDone:
				// Complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Template complete")
		}

		_ = local.AppendHistory(local.HistoryEntry{
			Command: "template run " + templateName,
			Input:   topic,
			Output:  output.String(),
		})

		if flagJSON {
			result := map[string]interface{}{
				"template": templateName,
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
	templateCmd.AddCommand(templateListCmd)
	templateCmd.AddCommand(templateShowCmd)
	templateCmd.AddCommand(templateRunCmd)
	rootCmd.AddCommand(templateCmd)
}
