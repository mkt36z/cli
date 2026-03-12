package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// WorkspaceConfig represents a project-level .mkt36z.yaml file.
type WorkspaceConfig struct {
	Name         string            `yaml:"name"`
	Description  string            `yaml:"description,omitempty"`
	DefaultAgent string            `yaml:"default_agent,omitempty"`
	Agents       []string          `yaml:"agents,omitempty"`
	Context      map[string]string `yaml:"context,omitempty"` // context overrides
	Pipeline     []PipelineStep    `yaml:"pipeline,omitempty"`
}

// PipelineStep represents a step in a composable pipeline.
type PipelineStep struct {
	Agent   string            `yaml:"agent"`
	Type    string            `yaml:"type,omitempty"`
	Options map[string]string `yaml:"options,omitempty"`
}

// FindWorkspaceConfig walks up the directory tree to find .mkt36z.yaml.
// Returns the config and its directory path, or nil if not found.
func FindWorkspaceConfig() (*WorkspaceConfig, string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	for {
		path := filepath.Join(dir, ".mkt36z.yaml")
		if _, err := os.Stat(path); err == nil {
			cfg, err := LoadWorkspaceConfig(path)
			if err != nil {
				return nil, "", err
			}
			return cfg, dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return nil, "", nil
}

// LoadWorkspaceConfig loads a workspace config from a file path.
func LoadWorkspaceConfig(path string) (*WorkspaceConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading workspace config: %w", err)
	}

	var cfg WorkspaceConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing workspace config: %w", err)
	}

	return &cfg, nil
}

// SaveWorkspaceConfig writes a workspace config to .mkt36z.yaml in the given directory.
func SaveWorkspaceConfig(dir string, cfg *WorkspaceConfig) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshaling workspace config: %w", err)
	}

	path := filepath.Join(dir, ".mkt36z.yaml")
	return os.WriteFile(path, data, 0644)
}

// HasWorkspaceConfig checks if a .mkt36z.yaml exists in the current directory tree.
func HasWorkspaceConfig() bool {
	cfg, _, _ := FindWorkspaceConfig()
	return cfg != nil
}
