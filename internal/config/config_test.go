package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.APIURL != "https://api.mkt36z.com" {
		t.Errorf("DefaultConfig APIURL = %q, want https://api.mkt36z.com", cfg.APIURL)
	}
}

func TestLoadWithEnvOverride(t *testing.T) {
	// Use a temp dir so we don't read real config files.
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)
	t.Setenv("MKT36Z_API_URL", "https://custom.api.com")

	cfg, err := Load("")
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.APIURL != "https://custom.api.com" {
		t.Errorf("APIURL = %q, want https://custom.api.com", cfg.APIURL)
	}
}

func TestSaveAndLoad(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	original := &Config{
		APIURL:     "https://test.api.com",
		DefaultOrg: "testorg",
		Brand: &BrandVoice{
			CompanyName: "TestCo",
			Tone:        []string{"confident", "warm"},
		},
	}

	if err := Save(original); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := Load("")
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if loaded.APIURL != original.APIURL {
		t.Errorf("APIURL = %q, want %q", loaded.APIURL, original.APIURL)
	}
	if loaded.DefaultOrg != original.DefaultOrg {
		t.Errorf("DefaultOrg = %q, want %q", loaded.DefaultOrg, original.DefaultOrg)
	}
	if loaded.Brand == nil || loaded.Brand.CompanyName != "TestCo" {
		t.Error("Brand not loaded correctly")
	}
}

func TestFindProjectConfig(t *testing.T) {
	tmp := t.TempDir()
	nested := filepath.Join(tmp, "a", "b", "c")
	os.MkdirAll(nested, 0755)

	// Place config at root.
	configPath := filepath.Join(tmp, ".mkt36z.yaml")
	os.WriteFile(configPath, []byte("api_url: test"), 0644)

	found := FindProjectConfig(nested)
	if found != configPath {
		t.Errorf("FindProjectConfig() = %q, want %q", found, configPath)
	}
}

func TestContextScore(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	// Empty — score should be 0.
	if score := ContextScore(); score != 0 {
		t.Errorf("ContextScore() = %d, want 0 (no files)", score)
	}

	// Create business context file.
	ctxDir := filepath.Join(tmp, "context")
	os.MkdirAll(ctxDir, 0700)
	os.WriteFile(filepath.Join(ctxDir, "business.yaml"), []byte("company_name: Test"), 0600)

	if score := ContextScore(); score != 10 {
		t.Errorf("ContextScore() = %d, want 10 (business only)", score)
	}
}

func TestHighestImpactMissing(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("MKT36Z_CONFIG_DIR", tmp)

	missing := HighestImpactMissing()
	// Product and ICP both have weight 15, but Product comes first in AllContextTypes.
	if missing != ContextProduct && missing != ContextICP {
		t.Errorf("HighestImpactMissing() = %q, want product or icp", missing)
	}
}
