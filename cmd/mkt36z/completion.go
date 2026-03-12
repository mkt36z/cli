package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion <shell>",
	Short: "Generate shell completions",
	Long: `Generate shell completion scripts for bash, zsh, fish, or powershell.

  Install completions:
    bash:       mkt36z completion bash > /etc/bash_completion.d/mkt36z
    zsh:        mkt36z completion zsh > "${fpath[1]}/_mkt36z"
    fish:       mkt36z completion fish > ~/.config/fish/completions/mkt36z.fish
    powershell: mkt36z completion powershell | Out-String | Invoke-Expression

  Examples:
    mkt36z completion bash
    mkt36z completion zsh > ~/.zfunc/_mkt36z && compinit`,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return rootCmd.GenBashCompletionV2(os.Stdout, true)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			return rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return fmt.Errorf("unsupported shell: %s", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
