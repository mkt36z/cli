// mkt36z is a CLI for AI-powered marketing intelligence.
//
// Architecture: thin Go client → Cloudflare Workers → OpenRouter LLM
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/mkt36z/cli/internal/config"
	"github.com/mkt36z/cli/internal/version"
)

func main() {
	// Expand aliases before cobra parses args
	os.Args = expandAlias(os.Args)

	// Non-blocking version check (cached 24h, stderr only)
	go checkForUpdate()

	// Graceful Ctrl+C handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Fprintln(os.Stderr, "\nInterrupted. Cleaning up...")
		cancel()
		// Give commands a moment to save partial results
		time.Sleep(500 * time.Millisecond)
		os.Exit(130)
	}()

	// Make context available to commands
	rootCmd.SetContext(ctx)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// checkForUpdate checks GitHub releases for a newer CLI version.
// Results are cached for 24 hours. Output is on stderr only.
func checkForUpdate() {
	cacheFile := filepath.Join(config.CacheDir(), "version-check.json")

	// Check cache freshness
	if info, err := os.Stat(cacheFile); err == nil {
		if time.Since(info.ModTime()) < 24*time.Hour {
			return
		}
	}

	current := version.Get()
	if current.Version == "dev" {
		return // Skip check for dev builds
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/mkt36z/cli/releases/latest")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return
	}

	// Cache the result
	_ = config.EnsureDir(config.CacheDir())
	cacheData, _ := json.Marshal(map[string]string{
		"latest":  release.TagName,
		"current": current.Version,
	})
	_ = os.WriteFile(cacheFile, cacheData, 0600)

	// Compare versions (simple string comparison — tag is like "v0.2.0")
	latest := release.TagName
	if latest != "" && latest != "v"+current.Version && latest != current.Version {
		fmt.Fprintf(os.Stderr, "A new version of mkt36z is available: %s → %s. Run: brew upgrade mkt36z\n",
			current.Version, latest)
	}
}
