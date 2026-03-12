package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigMigration(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	// Write a v0 config (no version field)
	oldConfig := `api_url: https://api.mkt36z.com
default_org: testorg
`
	configPath := filepath.Join(tmp, "config.yaml")
	if err := os.WriteFile(configPath, []byte(oldConfig), 0600); err != nil {
		t.Fatal(err)
	}

	// Load should trigger migration
	cfg, err := Load("")
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Version != ConfigVersion {
		t.Errorf("Version = %d, want %d", cfg.Version, ConfigVersion)
	}

	// Backup should exist
	entries, _ := os.ReadDir(tmp)
	var files []string
	for _, e := range entries {
		files = append(files, e.Name())
	}
	t.Logf("Files in config dir: %v", files)
	backupPath := filepath.Join(tmp, "config.yaml.v0.bak")
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		t.Errorf("migration backup not created, files: %v", files)
	}
}

func TestConfigVersionCurrent(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	// Write a current-version config (Save sets Version to ConfigVersion)
	cfg := DefaultConfig()
	if err := Save(cfg); err != nil {
		t.Fatal(err)
	}

	loaded, err := Load("")
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Save writes ConfigVersion, so loading should get it back
	if loaded.Version != ConfigVersion {
		t.Errorf("Version = %d, want %d", loaded.Version, ConfigVersion)
	}
}

func TestHooksConfig(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	cfg := DefaultConfig()
	cfg.Hooks = &HooksConfig{
		PostGenerate: []HookEntry{{Command: "echo done"}},
	}
	if err := Save(cfg); err != nil {
		t.Fatal(err)
	}

	loaded, err := Load("")
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if loaded.Hooks == nil {
		t.Fatal("Hooks is nil")
	}
	if len(loaded.Hooks.PostGenerate) != 1 {
		t.Fatalf("PostGenerate length = %d, want 1", len(loaded.Hooks.PostGenerate))
	}
	if loaded.Hooks.PostGenerate[0].Command != "echo done" {
		t.Errorf("PostGenerate[0].Command = %q, want 'echo done'", loaded.Hooks.PostGenerate[0].Command)
	}
}
