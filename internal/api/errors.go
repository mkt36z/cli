package api

import "fmt"

// APIError represents a classified error from the Workers API.
type APIError struct {
	StatusCode int
	RequestID  string
	Message    string
	Hint       string
}

func (e *APIError) Error() string {
	msg := e.Message
	if e.Hint != "" {
		msg += "\n  " + e.Hint
	}
	if e.RequestID != "" {
		msg += fmt.Sprintf("\n  Request ID: %s — share this with support", e.RequestID)
	}
	return msg
}

// NetworkError wraps a transport-level error.
type NetworkError struct {
	Err       error
	RequestID string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("Network error: %v\n  Check your internet connection. Retry the same command.\n  Request ID: %s", e.Err, e.RequestID)
}

func (e *NetworkError) Unwrap() error { return e.Err }

// classifyError maps HTTP status codes to user-friendly errors.
func classifyError(status int, reqID string) *APIError {
	switch status {
	case 401:
		return &APIError{
			StatusCode: status,
			RequestID:  reqID,
			Message:    "Authentication required.",
			Hint:       "Run `mkt36z auth login` to authenticate.",
		}
	case 403:
		return &APIError{
			StatusCode: status,
			RequestID:  reqID,
			Message:    "Access denied.",
			Hint:       "Check your API key permissions.",
		}
	case 422:
		return &APIError{
			StatusCode: status,
			RequestID:  reqID,
			Message:    "PII detected in request.",
			Hint:       "Remove personal data before submitting.",
		}
	case 426:
		return &APIError{
			StatusCode: status,
			RequestID:  reqID,
			Message:    "CLI version too old.",
			Hint:       "Update: brew upgrade mkt36z",
		}
	case 429:
		return &APIError{
			StatusCode: status,
			RequestID:  reqID,
			Message:    "Rate limited or generation quota reached.",
			Hint:       "Upgrade your plan: `mkt36z usage upgrade` — or wait and retry.",
		}
	default:
		if status >= 500 {
			return &APIError{
				StatusCode: status,
				RequestID:  reqID,
				Message:    "Server error.",
				Hint:       "Report at github.com/mkt36z/cli/issues",
			}
		}
		return &APIError{
			StatusCode: status,
			RequestID:  reqID,
			Message:    fmt.Sprintf("Request failed with status %d.", status),
		}
	}
}
