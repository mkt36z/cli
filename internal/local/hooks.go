package local

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mkt36z/cli/internal/config"
)

// RunHooks executes hooks for the given event type.
// Environment variables like $OUTPUT_FILE and $CAMPAIGN_ID are expanded.
//
// SECURITY: Only hooks from the global config (~/.mkt36z/config.yaml) are executed.
// Project-local .mkt36z.yaml hooks are ignored to prevent supply-chain attacks
// (VULN-01/02).
func RunHooks(hookType string, env map[string]string) {
	// Load ONLY global config — never trust project-local hooks
	cfg, err := config.LoadGlobal()
	if err != nil || cfg.Hooks == nil {
		return
	}

	var hooks []config.HookEntry
	switch hookType {
	case "post_generate":
		hooks = cfg.Hooks.PostGenerate
	case "post_analyze":
		hooks = cfg.Hooks.PostAnalyze
	case "post_campaign_approve":
		hooks = cfg.Hooks.PostCampaignApprove
	default:
		return
	}

	for _, hook := range hooks {
		expanded := expandEnv(hook.Command, env)
		runHookCommand(expanded)
	}
}

func expandEnv(command string, env map[string]string) string {
	result := command
	for k, v := range env {
		// Shell-escape the value to prevent injection via variable content
		result = strings.ReplaceAll(result, "$"+k, shellEscape(v))
	}
	return result
}

// shellEscape wraps a value in single quotes, escaping any embedded single quotes.
// This prevents shell metacharacter injection when values are interpolated into commands.
func shellEscape(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

func runHookCommand(command string) {
	// We keep sh -c because hooks legitimately need shell features (pipes, redirects),
	// but all interpolated values are now single-quote escaped (see shellEscape).
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stderr // Hooks output goes to stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Hook failed: %s: %v\n", command, err)
	}
}
