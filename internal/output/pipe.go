package output

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// HasStdin returns true if stdin has data available (piped input).
func HasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// ReadStdin reads all available data from stdin.
// Returns empty string if stdin is a terminal (no piped data).
func ReadStdin() (string, error) {
	if !HasStdin() {
		return "", nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// StdinIsTTY returns true if stdin is connected to a terminal.
func StdinIsTTY() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
}
