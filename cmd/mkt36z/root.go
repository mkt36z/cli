package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/local"
)

// Global flag values — accessed by subcommands.
var (
	flagJSON    bool
	flagAPIURL  string
	flagConfig  string
	flagVerbose int
	flagNoColor bool
	flagTiming  bool
)

// Default API URL for production Workers.
const defaultAPIURL = "https://api.mkt36z.com"

var rootCmd = &cobra.Command{
	Use:   "mkt36z",
	Short: "AI-powered marketing intelligence CLI",
	Long: `mkt36z — Marketing intelligence that compounds.

  Generate high-converting copy, analyze positioning, run multi-agent
  campaigns, and build a 90-day marketing plan — all from your terminal.

  Get started:
    mkt36z init                         Set up your brand context
    mkt36z generate headline "AI CRM"   Generate headlines instantly
    mkt36z analyze positioning           Diagnose your market position
    mkt36z doctor                        Check your setup

  Plan & execute:
    mkt36z plan create "Launch AI CRM"  Build a 90-day marketing plan
    mkt36z campaign create "Q1 launch"  Run a multi-agent campaign
    mkt36z usage show                   Check your quota and billing

  Customize:
    mkt36z alias set h "generate headline"   Create shortcuts
    mkt36z completion bash                   Shell completions

  Documentation: https://docs.mkt36z.com
  Source:        https://github.com/mkt36z/cli`,
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Inject API client for asset fetching (no network call here — lazy)
		key, _ := api.LoadAPIKey()
		if envKey := os.Getenv("MKT36Z_API_KEY"); envKey != "" {
			key = envKey
		}
		if key != "" {
			local.SetAPIClient(api.NewClient(flagAPIURL, key))
		}
	},
}

func init() {
	pf := rootCmd.PersistentFlags()

	pf.BoolVar(&flagJSON, "json", false, "Output in machine-readable JSON")
	pf.StringVar(&flagAPIURL, "api-url", defaultAPIURL, "Override Workers API URL")
	pf.StringVar(&flagConfig, "config", "", "Override config file path")
	pf.CountVarP(&flagVerbose, "verbose", "v", "Verbose output (-v info, -vv debug, -vvv trace)")
	pf.BoolVar(&flagNoColor, "no-color", false, "Disable ANSI color output")
	pf.BoolVar(&flagTiming, "timing", false, "Show request timing breakdown")

	// Respect NO_COLOR environment variable (https://no-color.org/)
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		flagNoColor = true
	}
}
