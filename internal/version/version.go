// Package version provides build-time version information injected via ldflags.
//
// Build with:
//
//	go build -ldflags "-X github.com/mkt36z/cli/internal/version.Version=0.1.0
//	  -X github.com/mkt36z/cli/internal/version.Commit=$(git rev-parse --short HEAD)
//	  -X github.com/mkt36z/cli/internal/version.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// Injected at build time via ldflags. Defaults for development builds.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// Info holds structured version information.
type Info struct {
	Version   string `json:"version" yaml:"version"`
	Commit    string `json:"commit" yaml:"commit"`
	Date      string `json:"date" yaml:"date"`
	GoVersion string `json:"go_version" yaml:"go_version"`
	OS        string `json:"os" yaml:"os"`
	Arch      string `json:"arch" yaml:"arch"`
}

// Get returns the current build information.
func Get() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		Date:      Date,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns a human-readable version string.
func (i Info) String() string {
	return fmt.Sprintf("mkt36z %s (%s) built %s\n%s %s/%s",
		i.Version, i.Commit, i.Date, i.GoVersion, i.OS, i.Arch)
}

// JSON returns the version info as a JSON string.
func (i Info) JSON() string {
	b, _ := json.MarshalIndent(i, "", "  ")
	return string(b)
}
