package api

import (
	"fmt"
	"strings"
	"testing"
)

func TestClassifyError(t *testing.T) {
	tests := []struct {
		status  int
		wantMsg string
		wantHint string
	}{
		{401, "Authentication required", "auth login"},
		{403, "Access denied", "permissions"},
		{422, "PII detected", "personal data"},
		{426, "CLI version too old", "upgrade"},
		{429, "Rate limited", "usage upgrade"},
		{500, "Server error", "github.com"},
		{418, "status 418", ""},
	}

	for _, tt := range tests {
		err := classifyError(tt.status, "req_test123")
		if !strings.Contains(err.Error(), tt.wantMsg) {
			t.Errorf("classifyError(%d) message = %q, want to contain %q", tt.status, err.Error(), tt.wantMsg)
		}
		if tt.wantHint != "" && !strings.Contains(err.Error(), tt.wantHint) {
			t.Errorf("classifyError(%d) hint = %q, want to contain %q", tt.status, err.Error(), tt.wantHint)
		}
		if err.RequestID != "req_test123" {
			t.Errorf("classifyError(%d) RequestID = %q, want 'req_test123'", tt.status, err.RequestID)
		}
	}
}

func TestNetworkError(t *testing.T) {
	err := &NetworkError{
		Err:       fmt.Errorf("connection refused"),
		RequestID: "req_abc",
	}
	msg := err.Error()
	if !strings.Contains(msg, "Network error") {
		t.Errorf("NetworkError message = %q, want to contain 'Network error'", msg)
	}
	if !strings.Contains(msg, "req_abc") {
		t.Errorf("NetworkError message = %q, want to contain request ID", msg)
	}
}
