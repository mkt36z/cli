package api

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoadAPIKey(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	key := "sk_live_test123456"
	if err := SaveAPIKey(key); err != nil {
		t.Fatalf("SaveAPIKey() error: %v", err)
	}

	loaded, err := LoadAPIKey()
	if err != nil {
		t.Fatalf("LoadAPIKey() error: %v", err)
	}
	if loaded != key {
		t.Errorf("LoadAPIKey() = %q, want %q", loaded, key)
	}
}

func TestAuthFilePermissions(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	if err := SaveAPIKey("sk_live_test"); err != nil {
		t.Fatalf("SaveAPIKey() error: %v", err)
	}

	path := filepath.Join(tmp, "auth.json")
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat() error: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("auth.json permissions = %o, want 0600", info.Mode().Perm())
	}
}

func TestLoadAPIKeyMissing(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	key, err := LoadAPIKey()
	if err != nil {
		t.Fatalf("LoadAPIKey() error: %v", err)
	}
	if key != "" {
		t.Errorf("LoadAPIKey() = %q, want empty string", key)
	}
}

func TestRemoveAPIKey(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	_ = SaveAPIKey("sk_live_test")
	if err := RemoveAPIKey(); err != nil {
		t.Fatalf("RemoveAPIKey() error: %v", err)
	}

	key, _ := LoadAPIKey()
	if key != "" {
		t.Errorf("after RemoveAPIKey(), LoadAPIKey() = %q, want empty", key)
	}
}

func TestRemoveAPIKeyNonExistent(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	// Should not error
	if err := RemoveAPIKey(); err != nil {
		t.Errorf("RemoveAPIKey() on non-existent file: %v", err)
	}
}
