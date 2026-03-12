package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// RenderTable renders a formatted table with headers and rows.
// Falls back to plain tab-separated output when colors are disabled.
func RenderTable(headers []string, rows [][]string) string {
	if len(headers) == 0 {
		return ""
	}

	if IsColorDisabled() {
		return renderPlainTable(headers, rows)
	}

	// Calculate column widths.
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Add padding.
	for i := range widths {
		widths[i] += 2
	}

	headerStyle := lipgloss.NewStyle().Bold(true).Foreground(ColorBrandBlue)
	cellStyle := lipgloss.NewStyle()
	dimStyle := lipgloss.NewStyle().Foreground(ColorDimGray)

	var b strings.Builder

	// Header row.
	for i, h := range headers {
		b.WriteString(headerStyle.Width(widths[i]).Render(h))
	}
	b.WriteString("\n")

	// Separator.
	for i, w := range widths {
		b.WriteString(dimStyle.Render(strings.Repeat("─", w)))
		if i < len(widths)-1 {
			b.WriteString(dimStyle.Render(" "))
		}
	}
	b.WriteString("\n")

	// Data rows.
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				b.WriteString(cellStyle.Width(widths[i]).Render(cell))
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

func renderPlainTable(headers []string, rows [][]string) string {
	var b strings.Builder
	b.WriteString(strings.Join(headers, "\t"))
	b.WriteString("\n")
	for _, row := range rows {
		b.WriteString(strings.Join(row, "\t"))
		b.WriteString("\n")
	}
	return b.String()
}
