// Package ui provides terminal rendering primitives for the mkt36z CLI.
// All output respects NO_COLOR, TERM=dumb, and pipe detection.
package ui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
)

// Brand color palette.
var (
	ColorBrandBlue = lipgloss.Color("#3B82F6")
	ColorGreen     = lipgloss.Color("#22C55E")
	ColorAmber     = lipgloss.Color("#F59E0B")
	ColorRed       = lipgloss.Color("#EF4444")
	ColorGray      = lipgloss.Color("#6B7280")
	ColorDimGray   = lipgloss.Color("#9CA3AF")
	ColorWhite     = lipgloss.Color("#F9FAFB")
)

// Shared styles used across the CLI.
var (
	// Text styles
	Bold      = lipgloss.NewStyle().Bold(true)
	Dim       = lipgloss.NewStyle().Foreground(ColorDimGray)
	Success   = lipgloss.NewStyle().Foreground(ColorGreen)
	Warning   = lipgloss.NewStyle().Foreground(ColorAmber)
	Error     = lipgloss.NewStyle().Foreground(ColorRed).Bold(true)
	Info      = lipgloss.NewStyle().Foreground(ColorBrandBlue)
	Subtle    = lipgloss.NewStyle().Foreground(ColorGray)

	// Structural styles
	Title     = lipgloss.NewStyle().Bold(true).Foreground(ColorBrandBlue)
	Heading   = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	Code      = lipgloss.NewStyle().Foreground(ColorAmber)
	URL       = lipgloss.NewStyle().Foreground(ColorBrandBlue).Underline(true)

	// Status indicators
	CheckMark = Success.Render("✓")
	CrossMark = Error.Render("✗")
	WarnMark  = Warning.Render("⚠")
	InfoMark  = Info.Render("●")
)

// IsTTY returns true if stdout is a terminal (not piped).
func IsTTY() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

// IsColorDisabled returns true if color should be suppressed.
func IsColorDisabled() bool {
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return true
	}
	if os.Getenv("TERM") == "dumb" {
		return true
	}
	return !IsTTY()
}
