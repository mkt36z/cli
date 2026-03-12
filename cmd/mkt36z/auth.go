package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage API authentication",
	Long:  "Log in, log out, check status, and retrieve your API token.",
}

// --- auth login ---

var authLoginAPIKey string

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with the mkt36z API",
	Long: `Authenticate by providing an API key.

  Interactive:
    mkt36z auth login

  Non-interactive:
    mkt36z auth login --api-key sk_live_xxx`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key := authLoginAPIKey

		// SECURITY (VULN-14): Warn if API key was passed via CLI flag (visible in ps/history)
		if key != "" {
			fmt.Fprintln(os.Stderr, "Warning: passing API keys via --api-key exposes them in process lists and shell history.")
			fmt.Fprintln(os.Stderr, "Prefer: mkt36z auth login (interactive) or MKT36Z_API_KEY env var.")
		}

		// Check env var
		if key == "" {
			key = os.Getenv("MKT36Z_API_KEY")
		}

		// Prompt interactively
		if key == "" {
			fmt.Print("Enter your API key: ")
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("reading input: %w", err)
			}
			key = strings.TrimSpace(line)
		}

		if key == "" {
			return fmt.Errorf("API key cannot be empty")
		}

		if err := api.SaveAPIKey(key); err != nil {
			return fmt.Errorf("saving API key: %w", err)
		}

		fmt.Println("Authenticated successfully. Key saved.")
		return nil
	},
}

// --- auth logout ---

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored API credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := api.RemoveAPIKey(); err != nil {
			return fmt.Errorf("removing credentials: %w", err)
		}
		fmt.Println("Logged out. Credentials removed.")
		return nil
	},
}

// --- auth status ---

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show authentication status and plan info",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil {
			return fmt.Errorf("loading credentials: %w", err)
		}

		if key == "" {
			if flagJSON {
				fmt.Println(`{"authenticated":false}`)
			} else {
				fmt.Println("Not authenticated. Run `mkt36z auth login` to get started.")
			}
			return nil
		}

		// Try to fetch status from the API
		client := api.NewClient(flagAPIURL, key)
		resp, err := client.Do(context.Background(), "GET", "/api/v1/auth/status", nil)
		if err != nil {
			// Offline fallback: show that we have a key but can't verify
			masked := maskKey(key)
			if flagJSON {
				data, _ := json.Marshal(map[string]interface{}{
					"authenticated": true,
					"key":           masked,
					"verified":      false,
				})
				fmt.Println(string(data))
			} else {
				fmt.Printf("Authenticated (offline)\n  Key: %s\n  Could not verify with API: %v\n", masked, err)
			}
			return nil
		}
		defer resp.Body.Close()

		// Online: decode and display status
		var status map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
			return fmt.Errorf("parsing status response: %w", err)
		}

		if flagJSON {
			status["authenticated"] = true
			status["key"] = maskKey(key)
			out, _ := json.MarshalIndent(status, "", "  ")
			fmt.Println(string(out))
		} else {
			fmt.Printf("Authenticated\n  Key: %s\n", maskKey(key))
			if plan, ok := status["plan"].(string); ok {
				fmt.Printf("  Plan: %s\n", plan)
			}
			if org, ok := status["org"].(string); ok {
				fmt.Printf("  Org: %s\n", org)
			}
		}
		return nil
	},
}

// --- auth token ---

var authTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Print the raw API key (for scripting)",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil {
			return fmt.Errorf("loading credentials: %w", err)
		}
		if key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login`")
		}
		fmt.Print(key)
		return nil
	},
}

// --- auth rotate ---

var authRotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Rotate your API key",
	Long: `Request a new API key, invalidating the old one.

  Example:
    mkt36z auth rotate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := api.LoadAPIKey()
		if err != nil || key == "" {
			return fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
		}

		client := api.NewClient(flagAPIURL, key)
		resp, err := client.Do(context.Background(), "POST", "/api/v1/auth/rotate", strings.NewReader("{}"))
		if err != nil {
			return fmt.Errorf("key rotation failed: %w", err)
		}
		defer resp.Body.Close()

		var result struct {
			NewKey string `json:"new_key"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("parsing rotation response: %w", err)
		}

		if result.NewKey == "" {
			return fmt.Errorf("server did not return a new key")
		}

		if err := api.SaveAPIKey(result.NewKey); err != nil {
			return fmt.Errorf("saving new key: %w", err)
		}

		fmt.Fprintf(os.Stderr, "API key rotated successfully.\n  New key: %s\n", maskKey(result.NewKey))
		return nil
	},
}

func init() {
	authLoginCmd.Flags().StringVar(&authLoginAPIKey, "api-key", "", "API key (non-interactive)")

	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authLogoutCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authTokenCmd)
	authCmd.AddCommand(authRotateCmd)
	rootCmd.AddCommand(authCmd)
}

// maskKey shows only the first 7 and last 4 characters of an API key.
func maskKey(key string) string {
	if len(key) <= 11 {
		return "sk_***"
	}
	return key[:7] + "..." + key[len(key)-4:]
}
