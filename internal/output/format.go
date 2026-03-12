// Package output provides format detection and rendering for CLI output.
// Supports human, JSON, and table output modes with automatic pipe detection.
package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mkt36z/cli/internal/ui"
)

// Format determines how output is rendered.
type Format int

const (
	FormatHuman Format = iota // Rich terminal output (default for TTY)
	FormatJSON                // Machine-readable JSON (--json or piped)
	FormatTable               // Structured table output
)

// Formatter controls output rendering based on format and TTY state.
type Formatter struct {
	format Format
	isTTY  bool
}

// NewFormatter creates a formatter based on the --json flag.
// Automatically detects pipe vs TTY.
func NewFormatter(jsonFlag bool) *Formatter {
	f := &Formatter{
		isTTY: ui.IsTTY(),
	}
	if jsonFlag {
		f.format = FormatJSON
	} else if !f.isTTY {
		// When piped, default to plain text (no ANSI).
		f.format = FormatHuman
	}
	return f
}

// IsJSON returns true if output should be JSON.
func (f *Formatter) IsJSON() bool {
	return f.format == FormatJSON
}

// IsTTY returns true if output is going to a terminal.
func (f *Formatter) IsTTY() bool {
	return f.isTTY
}

// PrintJSON writes a value as indented JSON to stdout.
func (f *Formatter) PrintJSON(v interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// PrintText writes plain text to stdout.
func (f *Formatter) PrintText(s string) {
	fmt.Println(s)
}

// PrintResult writes either JSON or human-readable output.
func (f *Formatter) PrintResult(jsonData interface{}, humanText string) {
	if f.IsJSON() {
		f.PrintJSON(jsonData) //nolint:errcheck
	} else {
		f.PrintText(humanText)
	}
}
