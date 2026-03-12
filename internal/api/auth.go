package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mkt36z/cli/internal/config"
)

// AuthData is the structure persisted in auth.json.
type AuthData struct {
	APIKey string `json:"api_key"`
}

// SaveAPIKey persists the API key to ~/.mkt36z/auth.json with 0600 permissions.
func SaveAPIKey(key string) error {
	dir := config.Dir()
	if err := config.EnsureDir(dir); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	data := AuthData{APIKey: key}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "auth.json")
	return os.WriteFile(path, b, 0600)
}

// LoadAPIKey reads the API key from auth.json.
// Returns empty string if the file doesn't exist.
//
// SECURITY (VULN-21): Verifies file permissions are 0600 to prevent
// other users from reading the API key.
func LoadAPIKey() (string, error) {
	path := config.AuthFilePath()

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", fmt.Errorf("reading auth file: %w", err)
	}

	// Check file permissions — warn and fix if too permissive
	if perm := info.Mode().Perm(); perm&0077 != 0 {
		fmt.Fprintf(os.Stderr, "Warning: %s has insecure permissions %o. Fixing to 0600.\n", path, perm)
		if err := os.Chmod(path, 0600); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not fix permissions: %v\n", err)
		}
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading auth file: %w", err)
	}

	var data AuthData
	if err := json.Unmarshal(b, &data); err != nil {
		return "", fmt.Errorf("parsing auth file: %w", err)
	}

	return data.APIKey, nil
}

// RemoveAPIKey deletes the auth.json file (logout).
func RemoveAPIKey() error {
	path := config.AuthFilePath()
	err := os.Remove(path)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
