package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config <command>",
	Short: "Manage CLI configuration and context",
	Long: `View and modify configuration, context dimensions, and brand voice.

  Commands:
    show      Show current configuration
    set       Set a configuration value
    get       Get a configuration value
    edit      Open context file in editor
    brand     Show/set brand voice
    context   Show context score and breakdown
    workspace List/use workspace configurations

  Exit codes:
    0  Success
    1  General error
    2  Configuration error
    3  Authentication error

  Examples:
    mkt36z config show
    mkt36z config set api_url https://api.mkt36z.com
    mkt36z config get api_url
    mkt36z config brand
    mkt36z config context`,
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := map[string]interface{}{
			"api_url":  flagAPIURL,
			"verbose":  flagVerbose,
			"json":     flagJSON,
			"no_color": flagNoColor,
		}

		// Add config file path
		if flagConfig != "" {
			cfg["config_file"] = flagConfig
		}

		// Context score
		score := config.ContextScore()
		cfg["context_score"] = score

		if flagJSON {
			out, _ := json.MarshalIndent(cfg, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		fmt.Println("Current configuration:")
		fmt.Printf("  api_url:       %s\n", flagAPIURL)
		fmt.Printf("  verbose:       %d\n", flagVerbose)
		fmt.Printf("  json:          %t\n", flagJSON)
		fmt.Printf("  no_color:      %t\n", flagNoColor)
		if flagConfig != "" {
			fmt.Printf("  config_file:   %s\n", flagConfig)
		}
		fmt.Printf("  context_score: %d/100\n", score)

		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value := args[1]

		// Validate known keys
		validKeys := []string{"api_url", "verbose", "json", "no_color", "config_file"}
		found := false
		for _, k := range validKeys {
			if k == key {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("unknown config key: %s\n  Valid keys: %s", key, strings.Join(validKeys, ", "))
		}

		// For now, persist via environment variables guidance
		fmt.Fprintf(os.Stderr, "To persist config, set environment variable:\n")
		fmt.Fprintf(os.Stderr, "  export MKT36Z_%s=%s\n", strings.ToUpper(key), value)
		fmt.Fprintf(os.Stderr, "\nOr add to your shell profile.\n")

		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		values := map[string]interface{}{
			"api_url":  flagAPIURL,
			"verbose":  flagVerbose,
			"json":     flagJSON,
			"no_color": flagNoColor,
		}

		val, ok := values[key]
		if !ok {
			return fmt.Errorf("unknown config key: %s", key)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(map[string]interface{}{key: val}, "", "  ")
			fmt.Println(string(out))
		} else {
			fmt.Println(val)
		}

		return nil
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit <context-type>",
	Short: "Open context file in editor",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctxType := config.ContextType(args[0])

		// Validate context type
		validTypes := []string{
			"business", "product", "icp", "competitive", "brand-voice",
			"channels", "content-seo", "sales", "goals",
			"language-bank", "proof-stack", "anti-patterns",
		}

		found := false
		for _, t := range validTypes {
			if string(ctxType) == t {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("unknown context type: %s\n  Valid types: %s", args[0], strings.Join(validTypes, ", "))
		}

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vi"
		}

		dir := config.ContextDir()
		path := fmt.Sprintf("%s/%s.yaml", dir, args[0])

		fmt.Fprintf(os.Stderr, "Open %s in your editor:\n  %s %s\n", args[0], editor, path)
		return nil
	},
}

var configBrandCmd = &cobra.Command{
	Use:   "brand",
	Short: "Show brand voice configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.ContextFileExists(config.ContextBrandVoice) {
			fmt.Println("No brand voice configured. Run `mkt36z init` to set up.")
			return nil
		}

		var voice config.BrandVoice
		if err := config.LoadContext(config.ContextBrandVoice, &voice); err != nil {
			return fmt.Errorf("failed to load brand voice: %w", err)
		}

		if flagJSON {
			out, _ := json.MarshalIndent(voice, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		fmt.Println("Brand Voice:")
		if voice.CompanyName != "" {
			fmt.Printf("  Company:     %s\n", voice.CompanyName)
		}
		if len(voice.Tone) > 0 {
			fmt.Printf("  Tone:        %s\n", strings.Join(voice.Tone, ", "))
		}
		if len(voice.Personality) > 0 {
			fmt.Printf("  Personality: %s\n", strings.Join(voice.Personality, ", "))
		}
		if voice.Archetype != "" {
			fmt.Printf("  Archetype:   %s\n", voice.Archetype)
		}
		if len(voice.ForbiddenWords) > 0 {
			fmt.Printf("  Forbidden:   %s\n", strings.Join(voice.ForbiddenWords, ", "))
		}

		return nil
	},
}

var configContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Show context score and breakdown",
	RunE: func(cmd *cobra.Command, args []string) error {
		score := config.ContextScore()
		breakdown := config.Breakdown()

		if flagJSON {
			out, _ := json.MarshalIndent(map[string]interface{}{
				"total":     score,
				"breakdown": breakdown,
			}, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		fmt.Printf("Context Score: %d/100\n\n", score)

		for _, b := range breakdown {
			status := "✗"
			if b.Filled {
				status = "✓"
			}
			fmt.Printf("  %s %-18s  (weight: %d)\n", status, b.Label, b.MaxPoints)
		}

		missing := config.HighestImpactMissing()
		if missing != "" {
			fmt.Printf("\nHighest impact missing: %s\n", missing.Label())
		}

		return nil
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configEditCmd)
	configCmd.AddCommand(configBrandCmd)
	configCmd.AddCommand(configContextCmd)
	rootCmd.AddCommand(configCmd)
}
