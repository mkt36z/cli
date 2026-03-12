package local

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
)

// Embed all bundled content at compile time. ~1 MB of markdown.
// This is the primary source — zero latency, always available offline.

//go:embed content/playbooks/*.md
var embeddedPlaybooks embed.FS

//go:embed content/templates/*.md
var embeddedTemplates embed.FS

//go:embed content/workflows/*.md
var embeddedWorkflows embed.FS

// embeddedFS returns the correct embedded filesystem for an asset type.
func embeddedFS(assetType string) embed.FS {
	switch assetType {
	case "playbooks":
		return embeddedPlaybooks
	case "templates":
		return embeddedTemplates
	case "workflows":
		return embeddedWorkflows
	default:
		return embed.FS{}
	}
}

// listEmbedded returns the names (without .md extension) of all embedded assets.
func listEmbedded(assetType string) []string {
	efs := embeddedFS(assetType)
	dir := filepath.Join("content", assetType)

	entries, err := fs.ReadDir(efs, dir)
	if err != nil {
		return nil
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		names = append(names, strings.TrimSuffix(e.Name(), ".md"))
	}
	return names
}

// getEmbedded reads a single embedded asset's content.
func getEmbedded(assetType, name string) (string, error) {
	efs := embeddedFS(assetType)
	path := filepath.Join("content", assetType, name+".md")

	data, err := fs.ReadFile(efs, path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
