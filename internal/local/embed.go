// Package local provides content assets (playbooks, templates, workflows)
// fetched from the mkt36z API and cached locally for fast offline access.
// Also discovers user custom assets from ~/.mkt36z/{playbooks,templates}/ and
// project-local .mkt36z/{playbooks,templates}/ directories.
package local

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mkt36z/cli/internal/api"
	"github.com/mkt36z/cli/internal/config"
)

// apiClient is the package-level client used for fetching assets.
// Set via SetAPIClient from root.go's PersistentPreRun.
var apiClient *api.Client

// SetAPIClient configures the API client used for asset fetching.
// Called once during CLI initialization.
func SetAPIClient(c *api.Client) {
	apiClient = c
}

// Asset represents a content asset.
type Asset struct {
	Name    string // filename without extension
	Path    string // cache path or custom path
	Content string // raw markdown content
}

// ListPlaybooks returns the names of all available playbooks.
// Custom playbooks from ~/.mkt36z/playbooks/ and .mkt36z/playbooks/ are
// included with a " [custom]" suffix.
func ListPlaybooks() ([]string, error) {
	return listAssetsWithCustom("playbooks")
}

// GetPlaybook returns the content of a playbook by name.
// Checks custom directories first, then cache, then API.
func GetPlaybook(name string) (string, error) {
	cleanName := strings.TrimSuffix(name, " [custom]")
	if content, err := getCustomAsset("playbooks", cleanName); err == nil {
		return content, nil
	}
	return getAssetCached("playbooks", cleanName)
}

// ListTemplates returns the names of all available templates.
// Custom templates from ~/.mkt36z/templates/ and .mkt36z/templates/ are
// included with a " [custom]" suffix.
func ListTemplates() ([]string, error) {
	return listAssetsWithCustom("templates")
}

// GetTemplate returns the content of a template by name.
func GetTemplate(name string) (string, error) {
	cleanName := strings.TrimSuffix(name, " [custom]")
	if content, err := getCustomAsset("templates", cleanName); err == nil {
		return content, nil
	}
	return getAssetCached("templates", cleanName)
}

// ListWorkflows returns the names of all available workflows.
func ListWorkflows() ([]string, error) {
	return listAssetsWithCustom("workflows")
}

// GetWorkflow returns the content of a workflow by name.
func GetWorkflow(name string) (string, error) {
	cleanName := strings.TrimSuffix(name, " [custom]")
	if content, err := getCustomAsset("workflows", cleanName); err == nil {
		return content, nil
	}
	return getAssetCached("workflows", cleanName)
}

// getAssetCached tries cache, then API, then stale cache (offline fallback).
func getAssetCached(assetType, name string) (string, error) {
	if err := validateAssetName(name); err != nil {
		return "", err
	}

	// 1. Try fresh cache
	if content, err := cachedContent(assetType, name); err == nil {
		return content, nil
	}

	// 2. Try API
	content, err := fetchAssetContent(apiClient, assetType, name)
	if err == nil {
		// Cache for next time (best-effort)
		_ = cacheContent(assetType, name, content)
		return content, nil
	}

	// 3. Offline fallback: serve stale cache
	if content, staleErr := staleCachedContent(assetType, name); staleErr == nil {
		return content, nil
	}

	return "", fmt.Errorf("%s %q not found: %w\n  Run `mkt36z auth login` if not authenticated", assetType, name, err)
}

// listAssetsCached tries cache, then API, then stale cache (offline fallback).
func listAssetsCached(assetType string) ([]string, error) {
	// 1. Try fresh cache
	if names, err := cachedList(assetType, false); err == nil {
		return names, nil
	}

	// 2. Try API
	names, err := fetchAssetList(apiClient, assetType)
	if err == nil {
		// Cache for next time (best-effort)
		_ = cacheList(assetType, names)
		return names, nil
	}

	// 3. Offline fallback: serve stale cache
	if names, staleErr := cachedList(assetType, true); staleErr == nil {
		return names, nil
	}

	return nil, fmt.Errorf("could not list %s: %w\n  Run `mkt36z auth login` if not authenticated", assetType, err)
}

// listAssetsWithCustom returns API/cached assets + custom assets merged together.
// Custom assets are marked with " [custom]" suffix.
// Returns the API error when the final result is empty so callers can show
// actionable feedback (e.g. "not authenticated" or "API unreachable").
func listAssetsWithCustom(assetType string) ([]string, error) {
	// Fetch from API/cache
	fetched, apiErr := listAssetsCached(assetType)
	if apiErr != nil {
		// If API fails, still try custom assets below
		fetched = nil
	}

	// Collect custom assets
	customNames := discoverCustomAssets(assetType)

	// Merge: fetched first, then custom (deduplicated)
	seen := make(map[string]bool)
	for _, name := range fetched {
		seen[name] = true
	}

	var result []string
	result = append(result, fetched...)
	for _, name := range customNames {
		if !seen[name] {
			result = append(result, name+" [custom]")
		}
	}

	sort.Strings(result)

	// Surface the API error when the result is empty so callers can
	// distinguish "no assets exist" from "couldn't reach the API".
	if len(result) == 0 && apiErr != nil {
		return nil, apiErr
	}

	return result, nil
}

// discoverCustomAssets finds .md files in user custom directories.
// Checks: ~/.mkt36z/{assetType}/ and .mkt36z/{assetType}/ (project-local).
func discoverCustomAssets(assetType string) []string {
	var names []string
	seen := make(map[string]bool)

	// Global: ~/.mkt36z/playbooks/ (or templates/ or workflows/)
	globalDir := filepath.Join(config.Dir(), assetType)
	names = append(names, scanDirForMD(globalDir, seen)...)

	// Project-local: .mkt36z/playbooks/
	if cwd, err := os.Getwd(); err == nil {
		localDir := filepath.Join(cwd, ".mkt36z", assetType)
		names = append(names, scanDirForMD(localDir, seen)...)
	}

	sort.Strings(names)
	return names
}

func scanDirForMD(dir string, seen map[string]bool) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var names []string
	for _, e := range entries {
		// SECURITY (VULN-13): Skip symlinks to prevent symlink-based file reads
		if e.Type()&os.ModeSymlink != 0 {
			continue
		}
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		name := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
		if !seen[name] {
			seen[name] = true
			names = append(names, name)
		}
	}
	return names
}

// validateAssetName rejects names containing path traversal sequences or separators.
// SECURITY (VULN-05): Prevents reading arbitrary files via crafted asset names
// like "../../etc/passwd".
func validateAssetName(name string) error {
	if strings.Contains(name, "..") {
		return fmt.Errorf("invalid asset name: %q", name)
	}
	if strings.ContainsAny(name, `/\`) {
		return fmt.Errorf("invalid asset name: %q", name)
	}
	if name == "" || name == "." {
		return fmt.Errorf("invalid asset name: %q", name)
	}
	return nil
}

// getCustomAsset reads a custom asset from user directories.
func getCustomAsset(assetType, name string) (string, error) {
	if err := validateAssetName(name); err != nil {
		return "", err
	}

	// Check project-local first
	if cwd, err := os.Getwd(); err == nil {
		localPath := filepath.Join(cwd, ".mkt36z", assetType, name+".md")
		if data, err := os.ReadFile(localPath); err == nil {
			return string(data), nil
		}
	}

	// Then global
	globalPath := filepath.Join(config.Dir(), assetType, name+".md")
	data, err := os.ReadFile(globalPath)
	if err != nil {
		return "", fmt.Errorf("custom asset %q not found", name)
	}
	return string(data), nil
}
