package version

import (
	"encoding/json"
	"runtime"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	info := Get()

	if info.GoVersion != runtime.Version() {
		t.Errorf("GoVersion = %q, want %q", info.GoVersion, runtime.Version())
	}
	if info.OS != runtime.GOOS {
		t.Errorf("OS = %q, want %q", info.OS, runtime.GOOS)
	}
	if info.Arch != runtime.GOARCH {
		t.Errorf("Arch = %q, want %q", info.Arch, runtime.GOARCH)
	}
}

func TestString(t *testing.T) {
	info := Get()
	s := info.String()

	if !strings.Contains(s, "mkt36z") {
		t.Errorf("String() missing 'mkt36z': %q", s)
	}
	if !strings.Contains(s, runtime.Version()) {
		t.Errorf("String() missing Go version: %q", s)
	}
}

func TestJSON(t *testing.T) {
	info := Get()
	j := info.JSON()

	var parsed Info
	if err := json.Unmarshal([]byte(j), &parsed); err != nil {
		t.Fatalf("JSON() output is not valid JSON: %v\n%s", err, j)
	}

	// Verify all required fields are present.
	if parsed.Version == "" {
		t.Error("JSON missing version field")
	}
	if parsed.GoVersion == "" {
		t.Error("JSON missing go_version field")
	}
	if parsed.OS == "" {
		t.Error("JSON missing os field")
	}
	if parsed.Arch == "" {
		t.Error("JSON missing arch field")
	}
}
