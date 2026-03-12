// Package testing provides DX testing utilities: golden files, benchmarks, and accessibility checks.
package testing

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GoldenFile manages golden file test fixtures.
type GoldenFile struct {
	Dir string // directory for golden files
}

// NewGoldenFile creates a golden file manager for the given testdata directory.
func NewGoldenFile(dir string) *GoldenFile {
	return &GoldenFile{Dir: dir}
}

// Update writes the golden file content. Used with -update flag in tests.
func (g *GoldenFile) Update(name, content string) error {
	path := filepath.Join(g.Dir, name+".golden")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

// Read returns the golden file content.
func (g *GoldenFile) Read(name string) (string, error) {
	path := filepath.Join(g.Dir, name+".golden")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("golden file %q not found: %w", name, err)
	}
	return string(data), nil
}

// Compare checks if actual matches the golden file. Returns a diff description on mismatch.
func (g *GoldenFile) Compare(name, actual string) (bool, string) {
	expected, err := g.Read(name)
	if err != nil {
		return false, fmt.Sprintf("golden file missing: %s", err)
	}

	if expected == actual {
		return true, ""
	}

	// Simple line-by-line diff
	expLines := strings.Split(expected, "\n")
	actLines := strings.Split(actual, "\n")

	var diff strings.Builder
	maxLines := len(expLines)
	if len(actLines) > maxLines {
		maxLines = len(actLines)
	}

	for i := 0; i < maxLines; i++ {
		exp := ""
		act := ""
		if i < len(expLines) {
			exp = expLines[i]
		}
		if i < len(actLines) {
			act = actLines[i]
		}
		if exp != act {
			diff.WriteString(fmt.Sprintf("line %d:\n  expected: %q\n  actual:   %q\n", i+1, exp, act))
		}
	}

	return false, diff.String()
}

// AccessibilityCheck verifies CLI output meets accessibility standards.
type AccessibilityCheck struct {
	Errors   []string
	Warnings []string
}

// CheckOutput validates CLI output for accessibility.
func CheckOutput(output string) AccessibilityCheck {
	var check AccessibilityCheck

	// Check for color-only information (requires --no-color support)
	if strings.Contains(output, "\033[") {
		check.Warnings = append(check.Warnings, "Output contains ANSI escape codes — ensure --no-color works")
	}

	// Check for unicode-only status indicators
	unicodeOnly := []string{"✓", "✗", "◎", "○"}
	for _, u := range unicodeOnly {
		if strings.Contains(output, u) {
			check.Warnings = append(check.Warnings, fmt.Sprintf("Uses unicode character %q — ensure ASCII fallback exists", u))
		}
	}

	// Check line length for terminal readability
	for i, line := range strings.Split(output, "\n") {
		if len(line) > 120 {
			check.Warnings = append(check.Warnings, fmt.Sprintf("Line %d exceeds 120 chars (%d)", i+1, len(line)))
		}
	}

	return check
}

// TimeToValue measures the number of commands to reach first value.
type TimeToValue struct {
	Steps       []string // command sequence
	Description string   // what value is delivered
}

// DefaultTimeToValue returns the expected fast path to first value.
func DefaultTimeToValue() TimeToValue {
	return TimeToValue{
		Steps: []string{
			"mkt36z auth login",
			"mkt36z init",
			"mkt36z generate headline \"my product\"",
		},
		Description: "3 commands to first generated content",
	}
}
