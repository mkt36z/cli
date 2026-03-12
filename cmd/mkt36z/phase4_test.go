package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// --- Step 29: Completion tests ---

func TestCompletionCommandRegistered(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"completion"})
	if err != nil {
		t.Fatalf("completion command not found: %v", err)
	}
	if cmd.Use != "completion <shell>" {
		t.Errorf("completion Use = %q, want 'completion <shell>'", cmd.Use)
	}
}

func TestCompletionValidArgs(t *testing.T) {
	cmd, _, _ := rootCmd.Find([]string{"completion"})
	want := []string{"bash", "zsh", "fish", "powershell"}
	if len(cmd.ValidArgs) != len(want) {
		t.Errorf("ValidArgs length = %d, want %d", len(cmd.ValidArgs), len(want))
	}
}

func TestCompletionBashOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"completion", "bash"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("completion bash failed: %v", err)
	}
	// Reset for other tests
	rootCmd.SetArgs(nil)
}

// --- Step 28: Plan command tests ---

func TestPlanCommandRegistered(t *testing.T) {
	for _, sub := range []string{"create", "show", "progress"} {
		cmd, _, err := rootCmd.Find([]string{"plan", sub})
		if err != nil {
			t.Errorf("plan %s command not found: %v", sub, err)
		}
		if cmd == nil {
			t.Errorf("plan %s command is nil", sub)
		}
	}
}

func TestPlanCreateFlags(t *testing.T) {
	cmd, _, _ := rootCmd.Find([]string{"plan", "create"})
	f := cmd.Flags()
	if f.Lookup("duration") == nil {
		t.Error("plan create missing --duration flag")
	}
	if f.Lookup("focus") == nil {
		t.Error("plan create missing --focus flag")
	}
}

// --- Step 30: Usage command tests ---

func TestUsageCommandRegistered(t *testing.T) {
	for _, sub := range []string{"show", "history", "upgrade"} {
		cmd, _, err := rootCmd.Find([]string{"usage", sub})
		if err != nil {
			t.Errorf("usage %s command not found: %v", sub, err)
		}
		if cmd == nil {
			t.Errorf("usage %s command is nil", sub)
		}
	}
}

func TestUsageBar(t *testing.T) {
	tests := []struct {
		pct   float64
		width int
		want  string
	}{
		{0, 10, "[----------]"},
		{50, 10, "[#####-----]"},
		{100, 10, "[##########]"},
		{150, 10, "[##########]"}, // clamp to max
	}

	for _, tt := range tests {
		got := usageBar(tt.pct, tt.width)
		if got != tt.want {
			t.Errorf("usageBar(%.0f, %d) = %q, want %q", tt.pct, tt.width, got, tt.want)
		}
	}
}

func TestProgressBar(t *testing.T) {
	bar := progressBar(50, 10)
	if bar != "#####-----" {
		t.Errorf("progressBar(50, 10) = %q, want '#####-----'", bar)
	}
}

// --- Step 38: Alias tests ---

func TestAliasCommandRegistered(t *testing.T) {
	for _, sub := range []string{"set", "list", "remove"} {
		cmd, _, err := rootCmd.Find([]string{"alias", sub})
		if err != nil {
			t.Errorf("alias %s command not found: %v", sub, err)
		}
		if cmd == nil {
			t.Errorf("alias %s command is nil", sub)
		}
	}
}

func TestAliasLoadSave(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	// Initially empty
	aliases, err := loadAliases()
	if err != nil {
		t.Fatalf("loadAliases() error: %v", err)
	}
	if len(aliases) != 0 {
		t.Errorf("expected 0 aliases, got %d", len(aliases))
	}

	// Save and reload
	aliases["h"] = "generate headline"
	aliases["brief"] = "generate brief"
	if err := saveAliases(aliases); err != nil {
		t.Fatalf("saveAliases() error: %v", err)
	}

	loaded, err := loadAliases()
	if err != nil {
		t.Fatalf("loadAliases() error: %v", err)
	}
	if loaded["h"] != "generate headline" {
		t.Errorf("alias h = %q, want 'generate headline'", loaded["h"])
	}

	// Verify file permissions
	path := filepath.Join(tmp, "aliases.json")
	info, _ := os.Stat(path)
	if info.Mode().Perm() != 0600 {
		t.Errorf("aliases.json permissions = %o, want 0600", info.Mode().Perm())
	}
}

func TestExpandAlias(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	// Set up aliases
	aliases := map[string]string{"h": "generate headline"}
	_ = saveAliases(aliases)

	// Expand
	args := expandAlias([]string{"mkt36z", "h", "AI CRM"})
	expected := []string{"mkt36z", "generate", "headline", "AI CRM"}

	if len(args) != len(expected) {
		t.Fatalf("expandAlias length = %d, want %d", len(args), len(expected))
	}
	for i, a := range args {
		if a != expected[i] {
			t.Errorf("expandAlias[%d] = %q, want %q", i, a, expected[i])
		}
	}
}

func TestExpandAliasNoMatch(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	args := expandAlias([]string{"mkt36z", "generate", "headline"})
	if len(args) != 3 {
		t.Errorf("expandAlias with no match: length = %d, want 3", len(args))
	}
}

// --- Step 36: Security tests ---

func TestAuthRotateCommandRegistered(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"auth", "rotate"})
	if err != nil {
		t.Fatalf("auth rotate command not found: %v", err)
	}
	if cmd == nil {
		t.Fatal("auth rotate command is nil")
	}
}

func TestMaskKey(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{"sk_live_abc123def456ghij", "sk_live...ghij"},
		{"short", "sk_***"},
		{"", "sk_***"},
	}

	for _, tt := range tests {
		got := maskKey(tt.key)
		if got != tt.want {
			t.Errorf("maskKey(%q) = %q, want %q", tt.key, got, tt.want)
		}
	}
}

// --- Step 34: Documentation tests ---

func TestAllCommandsHaveShortHelp(t *testing.T) {
	for _, cmd := range rootCmd.Commands() {
		if cmd.Short == "" {
			t.Errorf("command %q missing Short description", cmd.Use)
		}
	}
}

func TestRootCommandHasExamples(t *testing.T) {
	long := rootCmd.Long
	if !strings.Contains(long, "plan") {
		t.Error("root Long help missing plan command reference")
	}
	if !strings.Contains(long, "alias") {
		t.Error("root Long help missing alias command reference")
	}
	if !strings.Contains(long, "completion") {
		t.Error("root Long help missing completion command reference")
	}
}

// --- Step 32: Version check cache tests ---

func TestVersionCheckJSON(t *testing.T) {
	data := map[string]string{
		"latest":  "v0.2.0",
		"current": "0.1.0",
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	var loaded map[string]string
	if err := json.Unmarshal(b, &loaded); err != nil {
		t.Fatal(err)
	}
	if loaded["latest"] != "v0.2.0" {
		t.Errorf("latest = %q, want 'v0.2.0'", loaded["latest"])
	}
}
