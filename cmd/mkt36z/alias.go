package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/config"
)

var aliasCmd = &cobra.Command{
	Use:   "alias <command>",
	Short: "Manage command aliases",
	Long: `Create shortcuts for frequently used commands.

  Commands:
    set       Create or update an alias
    list      Show all aliases
    remove    Delete an alias

  Examples:
    mkt36z alias set h "generate headline"
    mkt36z alias set brief "generate brief"
    mkt36z alias list
    mkt36z alias remove h

  Then use:
    mkt36z h "AI CRM"    ->    mkt36z generate headline "AI CRM"`,
}

var aliasSetCmd = &cobra.Command{
	Use:   "set <name> <expansion>",
	Short: "Create or update an alias",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		expansion := args[1]

		aliases, err := loadAliases()
		if err != nil {
			return err
		}

		aliases[name] = expansion
		if err := saveAliases(aliases); err != nil {
			return err
		}

		fmt.Printf("Alias set: %s -> %s\n", name, expansion)
		return nil
	},
}

var aliasListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all aliases",
	RunE: func(cmd *cobra.Command, args []string) error {
		aliases, err := loadAliases()
		if err != nil {
			return err
		}

		if flagJSON {
			out, _ := json.MarshalIndent(aliases, "", "  ")
			fmt.Println(string(out))
			return nil
		}

		if len(aliases) == 0 {
			fmt.Println("No aliases configured. Create one with: mkt36z alias set <name> <command>")
			return nil
		}

		names := make([]string, 0, len(aliases))
		for name := range aliases {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			fmt.Printf("  %-12s -> %s\n", name, aliases[name])
		}

		return nil
	},
}

var aliasRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Delete an alias",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		aliases, err := loadAliases()
		if err != nil {
			return err
		}

		if _, ok := aliases[name]; !ok {
			return fmt.Errorf("alias %q not found", name)
		}

		delete(aliases, name)
		if err := saveAliases(aliases); err != nil {
			return err
		}

		fmt.Printf("Alias removed: %s\n", name)
		return nil
	},
}

// aliasFilePath returns the path to the aliases file.
func aliasFilePath() string {
	return filepath.Join(config.Dir(), "aliases.json")
}

func loadAliases() (map[string]string, error) {
	data, err := os.ReadFile(aliasFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, fmt.Errorf("reading aliases: %w", err)
	}

	var aliases map[string]string
	if err := json.Unmarshal(data, &aliases); err != nil {
		return nil, fmt.Errorf("parsing aliases: %w", err)
	}
	return aliases, nil
}

func saveAliases(aliases map[string]string) error {
	if err := config.EnsureDir(config.Dir()); err != nil {
		return err
	}
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(aliasFilePath(), data, 0600)
}

// expandAlias checks if the first arg is a known alias and expands it.
// Called from main before cobra parses args.
func expandAlias(args []string) []string {
	if len(args) < 2 {
		return args
	}

	aliases, err := loadAliases()
	if err != nil || len(aliases) == 0 {
		return args
	}

	candidate := args[1]
	expansion, ok := aliases[candidate]
	if !ok {
		return args
	}

	// Replace the alias with its expansion
	expanded := strings.Fields(expansion)
	result := make([]string, 0, len(args)+len(expanded))
	result = append(result, args[0])
	result = append(result, expanded...)
	result = append(result, args[2:]...)
	return result
}

func init() {
	aliasCmd.AddCommand(aliasSetCmd)
	aliasCmd.AddCommand(aliasListCmd)
	aliasCmd.AddCommand(aliasRemoveCmd)
	rootCmd.AddCommand(aliasCmd)
}
