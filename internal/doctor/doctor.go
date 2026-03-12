// Package doctor provides health checks for the mkt36z CLI environment.
package doctor

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/config"
	"github.com/mkt36z/cli/internal/local"
	"github.com/mkt36z/cli/internal/version"
)

// CheckStatus represents the result of a single health check.
type CheckStatus string

const (
	StatusOK   CheckStatus = "ok"
	StatusWarn CheckStatus = "warn"
	StatusFail CheckStatus = "fail"
	StatusSkip CheckStatus = "skip"
)

// Check represents a single diagnostic check result.
type Check struct {
	Name    string      `json:"name"`
	Status  CheckStatus `json:"status"`
	Message string      `json:"message"`
}

// Result holds all diagnostic check results.
type Result struct {
	Checks  []Check `json:"checks"`
	Summary string  `json:"summary"`
}

// RunAll executes all health checks and returns the results.
func RunAll(apiURL string) Result {
	var checks []Check

	checks = append(checks, checkVersion())
	checks = append(checks, checkAuth())
	checks = append(checks, checkConfig())
	checks = append(checks, checkConnectivity(apiURL))
	checks = append(checks, checkContext())
	checks = append(checks, checkAssets())
	checks = append(checks, checkRuntime())

	// Build summary
	ok, warn, fail := 0, 0, 0
	for _, c := range checks {
		switch c.Status {
		case StatusOK:
			ok++
		case StatusWarn:
			warn++
		case StatusFail:
			fail++
		}
	}

	summary := fmt.Sprintf("%d passed, %d warnings, %d failed", ok, warn, fail)

	return Result{Checks: checks, Summary: summary}
}

func checkVersion() Check {
	info := version.Get()
	return Check{
		Name:    "Version",
		Status:  StatusOK,
		Message: fmt.Sprintf("%s (%s) built %s", info.Version, info.Commit, info.Date),
	}
}

func checkAuth() Check {
	key, err := api.LoadAPIKey()
	if err != nil {
		return Check{Name: "Auth", Status: StatusFail, Message: fmt.Sprintf("Error reading auth: %v", err)}
	}
	if key == "" {
		// Also check env var
		if os.Getenv("MKT36Z_API_KEY") != "" {
			return Check{Name: "Auth", Status: StatusOK, Message: "Authenticated via MKT36Z_API_KEY env var"}
		}
		return Check{Name: "Auth", Status: StatusWarn, Message: "Not authenticated. Run `mkt36z auth login`"}
	}
	masked := key[:7] + "..." + key[len(key)-4:]
	return Check{Name: "Auth", Status: StatusOK, Message: fmt.Sprintf("Authenticated (%s)", masked)}
}

func checkConfig() Check {
	if !config.Exists() {
		return Check{Name: "Config", Status: StatusWarn, Message: "No config file. Run `mkt36z init` to create one"}
	}
	cfg, err := config.Load("")
	if err != nil {
		return Check{Name: "Config", Status: StatusFail, Message: fmt.Sprintf("Error loading config: %v", err)}
	}
	return Check{Name: "Config", Status: StatusOK, Message: fmt.Sprintf("API URL: %s", cfg.APIURL)}
}

func checkConnectivity(apiURL string) Check {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL+"/api/v1/health", nil)
	if err != nil {
		return Check{Name: "Connectivity", Status: StatusFail, Message: fmt.Sprintf("Invalid URL: %v", err)}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Check{Name: "Connectivity", Status: StatusWarn, Message: "Cannot reach API (offline mode available)"}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return Check{Name: "Connectivity", Status: StatusOK, Message: "API reachable"}
	}
	return Check{Name: "Connectivity", Status: StatusWarn, Message: fmt.Sprintf("API returned %d", resp.StatusCode)}
}

func checkContext() Check {
	ctxDir := config.ContextDir()
	entries, err := os.ReadDir(ctxDir)
	if err != nil {
		if os.IsNotExist(err) {
			return Check{Name: "Context", Status: StatusWarn, Message: "No context files. Run `mkt36z init` to set up brand context"}
		}
		return Check{Name: "Context", Status: StatusFail, Message: fmt.Sprintf("Error reading context: %v", err)}
	}

	count := 0
	for _, e := range entries {
		if !e.IsDir() {
			count++
		}
	}

	if count == 0 {
		return Check{Name: "Context", Status: StatusWarn, Message: "Context directory is empty"}
	}

	score := config.ContextScore()
	return Check{
		Name:    "Context",
		Status:  StatusOK,
		Message: fmt.Sprintf("%d context files (score: %d/100)", count, score),
	}
}

func checkAssets() Check {
	playbooks, _ := local.ListPlaybooks()
	templates, _ := local.ListTemplates()
	workflows, _ := local.ListWorkflows()

	total := len(playbooks) + len(templates) + len(workflows)
	if total == 0 {
		return Check{Name: "Assets", Status: StatusWarn, Message: "No cached assets. Run `mkt36z playbook list` to fetch from API"}
	}

	// Report cache age
	age := local.CacheAge("playbooks")
	ageStr := "unknown"
	if age >= 0 {
		if age < time.Hour {
			ageStr = fmt.Sprintf("%dm ago", int(age.Minutes()))
		} else if age < 24*time.Hour {
			ageStr = fmt.Sprintf("%dh ago", int(age.Hours()))
		} else {
			ageStr = fmt.Sprintf("%dd ago", int(age.Hours()/24))
		}
	}

	return Check{
		Name:   "Assets",
		Status: StatusOK,
		Message: fmt.Sprintf("%d playbooks, %d templates, %d workflows (cached %s)",
			len(playbooks), len(templates), len(workflows), ageStr),
	}
}

func checkRuntime() Check {
	return Check{
		Name:   "Runtime",
		Status: StatusOK,
		Message: fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
}
