// Package registry loads and provides access to the plugins.json configuration.
package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Plugin represents a local plugin definition.
type Plugin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Style       string `json:"style"`   // "cobra" or "flags"
	Manager     bool   `json:"manager"` // true for plugctl itself
}

// External represents an external plugin to fetch and build.
type External struct {
	Name   string `json:"name"`             // upstream name (used for fetch dir)
	Alias  string `json:"alias,omitempty"`  // our binary name (if different from upstream)
	Repo   string `json:"repo"`             // e.g. "github.com/ConduitIO/conduit"
	Tag    string `json:"tag"`              // e.g. "v0.11.0"
	Path   string `json:"path"`             // build path within repo, e.g. "cmd/conduit"
}

// BinaryName returns the name we use for the built binary.
func (e *External) BinaryName() string {
	if e.Alias != "" {
		return e.Alias
	}
	return e.Name
}

// Registry holds the full plugin configuration.
type Registry struct {
	Version  string     `json:"version"`
	Plugins  []Plugin   `json:"plugins"`
	External []External `json:"external"`
}

// Load reads plugins.json from the given directory.
func Load(dir string) (*Registry, error) {
	path := filepath.Join(dir, "plugins.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read plugins.json: %w", err)
	}

	var reg Registry
	if err := json.Unmarshal(data, &reg); err != nil {
		return nil, fmt.Errorf("parse plugins.json: %w", err)
	}

	return &reg, nil
}

// LoadFromRoot finds and loads plugins.json from the repo root.
func LoadFromRoot() (*Registry, error) {
	// Try current directory first
	if reg, err := Load("."); err == nil {
		return reg, nil
	}

	// Try to find repo root by looking for go.mod
	dir, err := findRepoRoot()
	if err != nil {
		return nil, err
	}

	return Load(dir)
}

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find repo root (no go.mod found)")
		}
		dir = parent
	}
}

// Binaries returns names of all plugins that should be built.
func (r *Registry) Binaries() []string {
	var names []string
	for _, p := range r.Plugins {
		names = append(names, p.Name)
	}
	return names
}

// Installable returns plugins that can be installed (excludes manager).
func (r *Registry) Installable() []Plugin {
	var plugins []Plugin
	for _, p := range r.Plugins {
		if !p.Manager {
			plugins = append(plugins, p)
		}
	}
	return plugins
}

// Get returns a plugin by name.
func (r *Registry) Get(name string) (*Plugin, bool) {
	for _, p := range r.Plugins {
		if p.Name == name {
			return &p, true
		}
	}
	return nil, false
}

// AllBinaries returns names of all binaries (local + external).
func (r *Registry) AllBinaries() []string {
	names := r.Binaries()
	for _, e := range r.External {
		names = append(names, e.BinaryName())
	}
	return names
}

// GetExternal returns an external plugin by name.
func (r *Registry) GetExternal(name string) (*External, bool) {
	for _, e := range r.External {
		if e.Name == name {
			return &e, true
		}
	}
	return nil, false
}
