package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// ConfigVersion tracks the schema version for migration.
const ConfigVersion = 2

// Config is the top-level configuration for the mkt36z CLI.
// 5-layer hierarchy (highest priority first):
//  1. Environment variables (MKT36Z_API_URL, MKT36Z_API_KEY)
//  2. CLI flags (--api-url, --config)
//  3. Project config (.mkt36z.yaml in cwd or parent)
//  4. Global config (~/.mkt36z/config.yaml)
//  5. Built-in defaults
type Config struct {
	Version    int         `yaml:"version,omitempty" json:"version,omitempty"`
	APIURL     string      `yaml:"api_url" json:"api_url"`
	APIKey     string      `yaml:"-" json:"-"` // never persisted in config.yaml
	DefaultOrg string      `yaml:"default_org,omitempty" json:"default_org,omitempty"`
	Brand      *BrandVoice `yaml:"brand,omitempty" json:"brand,omitempty"`
	Analytics  bool        `yaml:"analytics,omitempty" json:"analytics,omitempty"`
	Aliases    map[string]string `yaml:"aliases,omitempty" json:"aliases,omitempty"`
	Hooks      *HooksConfig      `yaml:"hooks,omitempty" json:"hooks,omitempty"`
}

// HooksConfig defines pre/post hooks for operations.
type HooksConfig struct {
	PostGenerate        []HookEntry `yaml:"post_generate,omitempty" json:"post_generate,omitempty"`
	PostAnalyze         []HookEntry `yaml:"post_analyze,omitempty" json:"post_analyze,omitempty"`
	PostCampaignApprove []HookEntry `yaml:"post_campaign_approve,omitempty" json:"post_campaign_approve,omitempty"`
}

// HookEntry defines a single hook command.
type HookEntry struct {
	Command string `yaml:"command" json:"command"`
}

// DefaultConfig returns the built-in defaults (layer 5).
// Note: Version is left at 0 so that migration detects old configs.
// Save() sets it to ConfigVersion before writing.
func DefaultConfig() *Config {
	return &Config{
		APIURL: "https://api.mkt36z.com",
	}
}

// Load reads the config using the 5-layer hierarchy.
// configOverride is the --config flag value (empty = use defaults).
//
// SECURITY: Project-local .mkt36z.yaml is loaded for non-sensitive fields only.
// Hooks from project configs are stripped to prevent supply-chain attacks (VULN-02).
func Load(configOverride string) (*Config, error) {
	cfg := DefaultConfig()

	// Layer 4: Global config (trusted — user's own config)
	globalPath := ConfigFilePath()
	if err := loadYAML(globalPath, cfg); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Remember hooks from global config before project config can overwrite them
	globalHooks := cfg.Hooks

	// Layer 3: Project config (walk up from cwd)
	if cwd, err := os.Getwd(); err == nil {
		if projectPath := FindProjectConfig(cwd); projectPath != "" {
			if err := loadYAML(projectPath, cfg); err != nil {
				return nil, err
			}
			// SECURITY: Restore global hooks — project configs must not set hooks.
			// A malicious .mkt36z.yaml in a cloned repo could execute arbitrary
			// commands via hooks. Only the user's global config is trusted for hooks.
			cfg.Hooks = globalHooks
		}
	}

	// Layer 2: --config override loads on top
	if configOverride != "" {
		if err := loadYAML(configOverride, cfg); err != nil {
			return nil, err
		}
	}

	// Layer 1: Environment variables (highest priority)
	if v := os.Getenv("MKT36Z_API_URL"); v != "" {
		cfg.APIURL = v
	}
	if v := os.Getenv("MKT36Z_API_KEY"); v != "" {
		cfg.APIKey = v
	}

	// SECURITY: Enforce HTTPS for API URL to prevent credential interception (VULN-10)
	if cfg.APIURL != "" && !strings.HasPrefix(cfg.APIURL, "https://") {
		// Allow http://localhost and http://127.0.0.1 for local development
		if !strings.HasPrefix(cfg.APIURL, "http://localhost") && !strings.HasPrefix(cfg.APIURL, "http://127.0.0.1") {
			return nil, fmt.Errorf("API URL must use HTTPS: %s", cfg.APIURL)
		}
	}

	// Auto-migrate old config versions (only if a config file exists)
	if cfg.Version < ConfigVersion && Exists() {
		if err := migrateConfig(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: config migration failed: %v\n", err)
		}
	}

	return cfg, nil
}

// LoadGlobal reads ONLY the global config (~/.mkt36z/config.yaml) + env vars.
// Used by security-sensitive operations like hooks that must not trust project configs.
func LoadGlobal() (*Config, error) {
	cfg := DefaultConfig()

	globalPath := ConfigFilePath()
	if err := loadYAML(globalPath, cfg); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if v := os.Getenv("MKT36Z_API_URL"); v != "" {
		cfg.APIURL = v
	}
	if v := os.Getenv("MKT36Z_API_KEY"); v != "" {
		cfg.APIKey = v
	}

	return cfg, nil
}

// migrateConfig upgrades older config versions to current.
func migrateConfig(cfg *Config) error {
	globalPath := ConfigFilePath()

	// Back up old config before migration
	if _, err := os.Stat(globalPath); err == nil {
		backupPath := fmt.Sprintf("%s.v%d.bak", globalPath, cfg.Version)
		data, err := os.ReadFile(globalPath)
		if err == nil {
			_ = os.WriteFile(backupPath, data, 0600)
			fmt.Fprintf(os.Stderr, "Config migrated from v%d to v%d. Backup saved at %s\n",
				cfg.Version, ConfigVersion, backupPath)
		}
	}

	cfg.Version = ConfigVersion

	// Save migrated config
	return Save(cfg)
}

// Exists returns true if a global config file exists on disk.
func Exists() bool {
	_, err := os.Stat(ConfigFilePath())
	return err == nil
}

// Save writes the config to the global config file.
func Save(cfg *Config) error {
	// Always write the current config version
	cfg.Version = ConfigVersion

	if err := EnsureDir(Dir()); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigFilePath(), data, 0600)
}

func loadYAML(path string, target interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, target)
}
