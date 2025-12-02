package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// PluginInfo holds information about an installed plugin.
type PluginInfo struct {
	Name    string
	Path    string
	Version string
}

// PluginDir returns the plugin install directory.
// Respects PLUGCTL_BIN env var, defaults to ~/.plugctl/bin
func PluginDir() string {
	if dir := os.Getenv("PLUGCTL_BIN"); dir != "" {
		return dir
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".", ".plugctl", "bin")
	}
	return filepath.Join(home, ".plugctl", "bin")
}

// EnsurePluginDir creates the plugin directory if it doesn't exist.
func EnsurePluginDir() error {
	dir := PluginDir()
	return os.MkdirAll(dir, 0755)
}

// GitHubAsset represents an asset in a GitHub release.
type GitHubAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

// GitHubReleaseDetail represents a detailed GitHub release API response.
type GitHubReleaseDetail struct {
	TagName string        `json:"tag_name"`
	HTMLURL string        `json:"html_url"`
	Assets  []GitHubAsset `json:"assets"`
}

// ListAvailable queries GitHub releases for available plugin binaries.
// Returns a list of plugin names available for the current platform.
func ListAvailable() ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", GitHubUser, GitHubRepo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("no releases found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api: %s", resp.Status)
	}

	var rel GitHubReleaseDetail
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, err
	}

	// Filter assets for current platform
	suffix := fmt.Sprintf("_%s_%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		suffix += ".exe"
	}

	var plugins []string
	seen := make(map[string]bool)
	for _, asset := range rel.Assets {
		if strings.HasSuffix(asset.Name, suffix) {
			// Extract plugin name (everything before _os_arch)
			name := strings.TrimSuffix(asset.Name, suffix)
			if !seen[name] {
				plugins = append(plugins, name)
				seen[name] = true
			}
		}
	}

	return plugins, nil
}

// ListInstalled scans PluginDir for installed plugins.
func ListInstalled() ([]PluginInfo, error) {
	dir := PluginDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var plugins []PluginInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		// Skip hidden files
		if strings.HasPrefix(name, ".") {
			continue
		}
		plugins = append(plugins, PluginInfo{
			Name: name,
			Path: filepath.Join(dir, name),
		})
	}

	return plugins, nil
}

// InstallPlugin downloads a plugin from GitHub releases to PluginDir.
func InstallPlugin(name string) error {
	if err := EnsurePluginDir(); err != nil {
		return fmt.Errorf("create plugin dir: %w", err)
	}

	url := DownloadURL(name)
	destPath := filepath.Join(PluginDir(), name)

	// Download
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("plugin %q not found for %s/%s", name, runtime.GOOS, runtime.GOARCH)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	// Write to temp file first
	tmpFile, err := os.CreateTemp(PluginDir(), name+"-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		return fmt.Errorf("write: %w", err)
	}
	tmpFile.Close()

	// Make executable
	if err := os.Chmod(tmpPath, 0755); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	// Move to final location
	if err := os.Rename(tmpPath, destPath); err != nil {
		return fmt.Errorf("install: %w", err)
	}

	return nil
}

// InstallLocal copies a local binary to PluginDir.
func InstallLocal(srcPath string) error {
	if err := EnsurePluginDir(); err != nil {
		return fmt.Errorf("create plugin dir: %w", err)
	}

	// Get plugin name from filename (strip platform suffix if present)
	name := filepath.Base(srcPath)
	// Remove common suffixes like _darwin_arm64, _linux_amd64, etc.
	for _, suffix := range []string{
		"_darwin_amd64", "_darwin_arm64",
		"_linux_amd64", "_linux_arm64",
		"_windows_amd64.exe", "_windows_arm64.exe",
	} {
		name = strings.TrimSuffix(name, suffix)
	}

	destPath := filepath.Join(PluginDir(), name)

	// Open source
	src, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("open source: %w", err)
	}
	defer src.Close()

	// Create destination
	dst, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("create dest: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}

// UpdatePlugin downloads the latest version of an installed plugin.
func UpdatePlugin(name string) error {
	return InstallPlugin(name)
}

// UpdateAll updates all installed plugins.
func UpdateAll() error {
	plugins, err := ListInstalled()
	if err != nil {
		return err
	}

	var errs []string
	for _, p := range plugins {
		if err := UpdatePlugin(p.Name); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", p.Name, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("some updates failed:\n  %s", strings.Join(errs, "\n  "))
	}
	return nil
}
