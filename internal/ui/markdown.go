package ui

import (
	"github.com/charmbracelet/glamour"
)

// RenderMarkdown renders markdown content with terminal-appropriate styling.
// Falls back to raw text when colors are disabled.
func RenderMarkdown(content string) (string, error) {
	if IsColorDisabled() {
		return content, nil
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		return content, err
	}

	return renderer.Render(content)
}
