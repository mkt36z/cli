// Package local provides content assets (playbooks, templates, workflows)
// fetched from the API and cached locally for fast offline access.
package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mkt36z/cli/internal/config"
)

// DefaultCacheTTL is how long cached assets are considered fresh.
const DefaultCacheTTL = 24 * time.Hour

// manifestEntry represents a cached asset list item.
type manifestEntry struct {
	Name        string `json:"name"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// manifestCache is the on-disk format for cached asset lists.
type manifestCache struct {
	CachedAt time.Time       `json:"cached_at"`
	Assets   []manifestEntry `json:"assets"`
}

// assetsCacheDir returns the cache directory for a given asset type.
// e.g. ~/.cache/mkt36z/assets/playbooks/
func assetsCacheDir(assetType string) string {
	return filepath.Join(config.CacheDir(), "assets", assetType)
}

// cachedList reads the cached manifest for an asset type.
// Returns nil if the cache is missing or stale (beyond TTL).
// If forceStale is true, returns the cache regardless of age (offline fallback).
func cachedList(assetType string, forceStale bool) ([]string, error) {
	path := filepath.Join(assetsCacheDir(assetType), ".manifest.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m manifestCache
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	if !forceStale && time.Since(m.CachedAt) > DefaultCacheTTL {
		return nil, fmt.Errorf("cache expired")
	}

	names := make([]string, len(m.Assets))
	for i, a := range m.Assets {
		names[i] = a.Name
	}
	return names, nil
}

// cacheList writes an asset list to the manifest cache.
func cacheList(assetType string, names []string) error {
	dir := assetsCacheDir(assetType)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	entries := make([]manifestEntry, len(names))
	for i, n := range names {
		entries[i] = manifestEntry{Name: n}
	}

	m := manifestCache{
		CachedAt: time.Now(),
		Assets:   entries,
	}

	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, ".manifest.json"), data, 0600)
}

// cachedContent reads a single asset's content from cache.
// Returns empty string and error if not cached or stale.
func cachedContent(assetType, name string) (string, error) {
	path := filepath.Join(assetsCacheDir(assetType), name+".md")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if time.Since(info.ModTime()) > DefaultCacheTTL {
		return "", fmt.Errorf("cache expired")
	}

	return string(data), nil
}

// staleCachedContent reads a cached asset regardless of age (offline fallback).
func staleCachedContent(assetType, name string) (string, error) {
	path := filepath.Join(assetsCacheDir(assetType), name+".md")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// cacheContent writes a single asset's content to cache.
func cacheContent(assetType, name, content string) error {
	dir := assetsCacheDir(assetType)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, name+".md"), []byte(content), 0600)
}

// ClearCache removes all cached assets.
func ClearCache() error {
	return os.RemoveAll(filepath.Join(config.CacheDir(), "assets"))
}

// CacheAge returns how long ago the manifest was cached for an asset type.
// Returns -1 if no cache exists.
func CacheAge(assetType string) time.Duration {
	path := filepath.Join(assetsCacheDir(assetType), ".manifest.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return -1
	}

	var m manifestCache
	if err := json.Unmarshal(data, &m); err != nil {
		return -1
	}

	return time.Since(m.CachedAt)
}
