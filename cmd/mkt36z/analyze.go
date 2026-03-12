package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/output"
	"github.com/mkt36z/cli/internal/ui"
)

var validAnalyzeTypes = []string{
	"positioning", "offer", "awareness", "competitors",
	"pmf", "seo", "brand-health", "revenue", "attribution",
	"email-health", "growth-loops", "ad-performance",
	"content-decay", "category",
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze <type> [topic]",
	Short: "Analyze marketing strategy and performance",
	Long: fmt.Sprintf(`Run AI-powered analysis on your marketing strategy.

  Types: %s

  Examples:
    mkt36z analyze positioning
    mkt36z analyze offer "SaaS subscription model"
    mkt36z analyze awareness "https://example.com/landing-page"
    mkt36z analyze competitors
    mkt36z analyze revenue "SaaS metrics"
    mkt36z analyze seo "https://example.com"
    mkt36z analyze brand-health "our brand perception"
    mkt36z analyze pmf "product-market fit signals"

  Use --json for machine-readable output.`,
		strings.Join(validAnalyzeTypes, ", ")),
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		analysisType := args[0]

		if !isValidAnalyzeType(analysisType) {
			return fmt.Errorf("invalid analysis type: %s\n  Valid types: %s", analysisType, strings.Join(validAnalyzeTypes, ", "))
		}

		// Build topic
		topic := ""
		if len(args) > 1 {
			topic = strings.Join(args[1:], " ")
		}

		// Check stdin
		if topic == "" {
			if data, err := output.ReadStdin(); err == nil && data != "" {
				topic = data
			}
		}

		// For some analysis types, topic is optional (uses context)
		if topic == "" {
			topic = buildTopicFromContext(analysisType)
		}

		if topic == "" {
			return fmt.Errorf("provide a topic or set up context first: mkt36z init")
		}

		// Load API key
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		// Load user context
		userCtx := loadUserContext()

		reqBody := map[string]interface{}{
			"topic": topic,
		}
		if userCtx != nil {
			reqBody["context"] = userCtx
		}

		bodyJSON, _ := json.Marshal(reqBody)

		client := api.NewStreamingClient(flagAPIURL, key)
		path := fmt.Sprintf("/api/v1/analyze/%s", analysisType)

		var sp *ui.Spinner
		if !flagJSON {
			sp = ui.NewSpinner(fmt.Sprintf("Analyzing %s...", analysisType))
		}

		resp, err := client.Do(context.Background(), "POST", path, bytes.NewReader(bodyJSON))
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
					return fmt.Errorf("analysis error: %s", errData.Error)
				}

			case api.EventDone:
				// Stream complete
			}
		}

		if sp != nil {
			sp.Stop(ui.CheckMark + " Analysis complete")
		}

		if flagJSON {
			result := map[string]interface{}{
				"type":     analysisType,
				"topic":    topic,
				"analysis": content.String(),
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

func init() {
	rootCmd.AddCommand(analyzeCmd)
}

func isValidAnalyzeType(t string) bool {
	for _, v := range validAnalyzeTypes {
		if v == t {
			return true
		}
	}
	return false
}

// buildTopicFromContext creates a topic string from user context for contextual analysis.
func buildTopicFromContext(analysisType string) string {
	ctx := loadUserContext()
	if ctx == nil {
		return ""
	}

	// Build a summary from available context
	var parts []string

	if biz, ok := ctx["business"]; ok {
		data, _ := json.Marshal(biz)
		parts = append(parts, "Business: "+string(data))
	}
	if prod, ok := ctx["product"]; ok {
		data, _ := json.Marshal(prod)
		parts = append(parts, "Product: "+string(data))
	}
	if icp, ok := ctx["icp"]; ok {
		data, _ := json.Marshal(icp)
		parts = append(parts, "ICP: "+string(data))
	}
	if comp, ok := ctx["competitive"]; ok {
		data, _ := json.Marshal(comp)
		parts = append(parts, "Competitive: "+string(data))
	}

	if len(parts) == 0 {
		return ""
	}

	return fmt.Sprintf("Analyze %s for my business:\n\n%s", analysisType, strings.Join(parts, "\n"))
}
