package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/config"
)

var (
	learnFromURL  string
	learnFromFile string
)

var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "Extract context from existing content",
	Long: `Use AI to extract brand context from URLs or files.

  From a URL:
    mkt36z learn --from-url https://example.com/about

  From a file:
    mkt36z learn --from-file pitch-deck.md

  From multiple files:
    mkt36z learn --from-file "campaigns/*.md"

  Extracted context is merged into your existing context files.
  Check your score before and after: mkt36z doctor`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if learnFromURL == "" && learnFromFile == "" {
			return fmt.Errorf("provide --from-url or --from-file")
		}

		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		scoreBefore := config.ContextScore()

		if learnFromURL != "" {
			return learnURL(key, scoreBefore)
		}
		return learnFile(key, scoreBefore)
	},
}

func init() {
	learnCmd.Flags().StringVar(&learnFromURL, "from-url", "", "URL to extract context from")
	learnCmd.Flags().StringVar(&learnFromFile, "from-file", "", "File path or glob to extract context from")
	rootCmd.AddCommand(learnCmd)
}

func learnURL(apiKey string, scoreBefore int) error {
	fmt.Printf("Learning from URL: %s\n", learnFromURL)

	client := api.NewClient(flagAPIURL, apiKey)
	body, _ := json.Marshal(map[string]string{
		"url": learnFromURL,
	})

	resp, err := client.Do(context.Background(), "POST", "/api/v1/learn", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("parsing response: %w", err)
	}

	scoreAfter := config.ContextScore()
	printLearnResult(scoreBefore, scoreAfter, result)
	return nil
}

func learnFile(apiKey string, scoreBefore int) error {
	// Expand glob
	matches, err := filepath.Glob(learnFromFile)
	if err != nil {
		return fmt.Errorf("invalid glob pattern: %w", err)
	}
	if len(matches) == 0 {
		// Try as a literal path
		if _, err := os.Stat(learnFromFile); err != nil {
			return fmt.Errorf("no files found matching: %s", learnFromFile)
		}
		matches = []string{learnFromFile}
	}

	fmt.Printf("Learning from %d file(s)...\n", len(matches))

	client := api.NewClient(flagAPIURL, apiKey)

	for _, path := range matches {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("  Skipping %s: %v\n", path, err)
			continue
		}

		fmt.Printf("  Processing: %s\n", path)

		body, _ := json.Marshal(map[string]string{
			"content":  string(data),
			"filename": filepath.Base(path),
		})

		resp, err := client.Do(context.Background(), "POST", "/api/v1/learn", bytes.NewReader(body))
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
			continue
		}
		resp.Body.Close()
	}

	scoreAfter := config.ContextScore()
	printLearnResult(scoreBefore, scoreAfter, nil)
	return nil
}

func printLearnResult(before, after int, result map[string]interface{}) {
	fmt.Println()
	if before != after {
		fmt.Printf("Context score: %d → %d (+%d)\n", before, after, after-before)
	} else {
		fmt.Printf("Context score: %d/100 (no change)\n", after)
	}

	if flagJSON && result != nil {
		result["score_before"] = before
		result["score_after"] = after
		out, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(out))
	}
}
