package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ErrorBox renders a formatted error with a suggested fix.
func ErrorBox(err error, suggestion string) string {
	if IsColorDisabled() {
		return errorPlain(err, suggestion)
	}

	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(ColorRed).
		Padding(0, 1)

	title := Error.Render("Error")
	msg := fmt.Sprintf("%s %s", title, err.Error())

	if suggestion != "" {
		msg += "\n\n" + Dim.Render("Fix: ") + suggestion
	}

	return border.Render(msg)
}

func errorPlain(err error, suggestion string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Error: %s", err.Error())
	if suggestion != "" {
		fmt.Fprintf(&b, "\nFix: %s", suggestion)
	}
	return b.String()
}

// RenderError prints a formatted error to stderr and returns the error.
func RenderError(err error, suggestion string) error {
	fmt.Fprintln(os.Stderr, ErrorBox(err, suggestion))
	return err
}

// CommonErrors maps HTTP status codes to user-friendly messages with fix suggestions.
var CommonErrors = map[int]struct {
	Message    string
	Suggestion string
}{
	401: {
		Message:    "Authentication required",
		Suggestion: "Run `mkt36z auth login` to authenticate",
	},
	403: {
		Message:    "Access denied",
		Suggestion: "Check your API key permissions or upgrade your plan",
	},
	429: {
		Message:    "Rate limited",
		Suggestion: "Wait a moment or upgrade: `mkt36z usage upgrade`",
	},
	500: {
		Message:    "Server error",
		Suggestion: "Try again. If the issue persists, report at github.com/mkt36z/cli/issues",
	},
}

// ErrorForStatus returns a formatted error for an HTTP status code.
// requestID is included for support reference.
func ErrorForStatus(status int, requestID string) string {
	detail, ok := CommonErrors[status]
	if !ok {
		detail = struct {
			Message    string
			Suggestion string
		}{
			Message:    fmt.Sprintf("Unexpected error (HTTP %d)", status),
			Suggestion: "Try again or report at github.com/mkt36z/cli/issues",
		}
	}

	msg := detail.Message
	if requestID != "" {
		msg += fmt.Sprintf(" (Request ID: %s)", requestID)
	}

	return ErrorBox(fmt.Errorf("%s", msg), detail.Suggestion)
}
