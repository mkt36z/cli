// Package local provides content assets (playbooks, templates, workflows)
// bundled into the binary via go:embed. User custom assets from
// ~/.mkt36z/{playbooks,templates}/ and .mkt36z/{playbooks,templates}/
// are merged in with a " [custom]" suffix.
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
// Retained for future use (e.g. fetching user-shared assets from the API).
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
func ListPlaybooks() ([]string, error) {
	return listAssets("playbooks")
}

// GetPlaybook returns the content of a playbook by name.
// Priority: custom (project-local, then global) → embedded.
func GetPlaybook(name string) (string, error) {
	return getAsset("playbooks", name)
}

// ListTemplates returns the names of all available templates.
func ListTemplates() ([]string, error) {
	return listAssets("templates")
}

// GetTemplate returns the content of a template by name.
func GetTemplate(name string) (string, error) {
	return getAsset("templates", name)
}

// ListWorkflows returns the names of all available workflows.
func ListWorkflows() ([]string, error) {
	return listAssets("workflows")
}

// GetWorkflow returns the content of a workflow by name.
func GetWorkflow(name string) (string, error) {
	return getAsset("workflows", name)
}

// getAsset returns content for a named asset.
// Priority: custom directories → embedded binary content.
func getAsset(assetType, name string) (string, error) {
	cleanName := strings.TrimSuffix(name, " [custom]")
	if err := validateAssetName(cleanName); err != nil {
		return "", err
	}

	// 1. Custom directories (user overrides)
	if content, err := getCustomAsset(assetType, cleanName); err == nil {
		return content, nil
	}

	// 2. Embedded in binary (always available)
	if content, err := getEmbedded(assetType, cleanName); err == nil {
		return content, nil
	}

	return "", fmt.Errorf("%s %q not found", assetType, cleanName)
}

// listAssets returns embedded assets + custom assets merged together.
// Custom assets that don't exist in embedded are marked with " [custom]" suffix.
func listAssets(assetType string) ([]string, error) {
	// Embedded content: always available, zero latency
	embedded := listEmbedded(assetType)

	// Custom assets from user directories
	customNames := discoverCustomAssets(assetType)

	// Merge: embedded first, then custom (deduplicated)
	seen := make(map[string]bool, len(embedded))
	for _, name := range embedded {
		seen[name] = true
	}

	var result []string
	result = append(result, embedded...)
	for _, name := range customNames {
		if !seen[name] {
			result = append(result, name+" [custom]")
		}
	}

	sort.Strings(result)
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
