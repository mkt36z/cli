// Package config provides XDG-compliant configuration management for the mkt36z CLI.
package config

import (
	"os"
	"path/filepath"
	"runtime"
)

const appName = "mkt36z"

// Dir returns the primary config directory (~/.mkt36z/).
// Respects XDG_CONFIG_HOME on Linux, uses standard locations on macOS/Windows.
func Dir() string {
	if v := os.Getenv("MKT36Z_CONFIG_DIR"); v != "" {
		return v
	}

	switch runtime.GOOS {
	case "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "."+appName)
	case "windows":
		if v := os.Getenv("APPDATA"); v != "" {
			return filepath.Join(v, appName)
		}
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "."+appName)
	default: // linux, freebsd, etc.
		if v := os.Getenv("XDG_CONFIG_HOME"); v != "" {
			return filepath.Join(v, appName)
		}
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "."+appName)
	}
}

// CacheDir returns the cache directory (~/.cache/mkt36z/).
func CacheDir() string {
	if v := os.Getenv("MKT36Z_CACHE_DIR"); v != "" {
		return v
	}

	switch runtime.GOOS {
	case "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "Library", "Caches", appName)
	case "windows":
		if v := os.Getenv("LOCALAPPDATA"); v != "" {
			return filepath.Join(v, appName, "cache")
		}
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cache", appName)
	default:
		if v := os.Getenv("XDG_CACHE_HOME"); v != "" {
			return filepath.Join(v, appName)
		}
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cache", appName)
	}
}

// ConfigFilePath returns the path to the global config file.
func ConfigFilePath() string {
	return filepath.Join(Dir(), "config.yaml")
}

// AuthFilePath returns the path to the auth credentials file.
func AuthFilePath() string {
	return filepath.Join(Dir(), "auth.json")
}

// ContextDir returns the path to the context files directory.
func ContextDir() string {
	return filepath.Join(Dir(), "context")
}

// HistoryFilePath returns the path to the generation history file.
func HistoryFilePath() string {
	return filepath.Join(CacheDir(), "history.jsonl")
}

// EnsureDir creates a directory if it doesn't exist, with 0700 permissions.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0700)
}

// FindProjectConfig walks up from dir looking for .mkt36z.yaml.
// Returns the path if found, empty string otherwise.
func FindProjectConfig(dir string) string {
	for {
		candidate := filepath.Join(dir, ".mkt36z.yaml")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}
