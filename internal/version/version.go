package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Build-time variables set via -ldflags
var (
	Version = "dev"
	// These can be overridden at build time if needed
	GitHubUser = "joeblew999"
	GitHubRepo = "plugs"
	// DocsHost is where plugin docs are hosted
	DocsHost = "plugs.ubuntusoftware.net"
)

// Info returns version and build info as a formatted string.
func Info() string {
	return fmt.Sprintf("%s %s/%s", Version, runtime.GOOS, runtime.GOARCH)
}

// GitHubRelease represents a GitHub release API response.
type GitHubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

// LatestRelease fetches the latest release tag from GitHub.
func LatestRelease() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", GitHubUser, GitHubRepo)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("no releases found")
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github api: %s", resp.Status)
	}

	var rel GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return "", err
	}
	return rel.TagName, nil
}

// CheckUpdate compares current version with latest release.
// Returns (latestVersion, needsUpdate, error).
func CheckUpdate() (string, bool, error) {
	latest, err := LatestRelease()
	if err != nil {
		return "", false, err
	}
	// Normalize versions for comparison (strip 'v' prefix)
	current := strings.TrimPrefix(Version, "v")
	latestNorm := strings.TrimPrefix(latest, "v")
	// If current contains git info (e.g., "v0.1.4-11-g81f6555"), it's a dev build
	if strings.Contains(current, "-") {
		return latest, true, nil
	}
	return latest, current != latestNorm, nil
}

// DownloadURL returns the download URL for a specific binary and platform.
func DownloadURL(binaryName string) string {
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	asset := fmt.Sprintf("%s_%s_%s%s", binaryName, runtime.GOOS, runtime.GOARCH, ext)
	return fmt.Sprintf("https://github.com/%s/%s/releases/latest/download/%s", GitHubUser, GitHubRepo, asset)
}

// SelfUpdate downloads and replaces the current binary with the latest release.
// Only works if the binary is in PluginDir() (installed via plugctl).
// For binaries outside PluginDir, use plugctl update instead.
func SelfUpdate(binaryName string) error {
	// Get current executable path
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("get executable path: %w", err)
	}

	// Resolve symlinks to get real path
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}

	// Check if binary is in PluginDir (installer-enforced location)
	expectedDir := PluginDir()
	actualDir := filepath.Dir(execPath)

	if actualDir != expectedDir {
		return fmt.Errorf(
			"self-update only works for managed binaries\n"+
				"  Current location: %s\n"+
				"  Expected location: %s\n\n"+
				"To update, either:\n"+
				"  1. Run: plugctl update %s\n"+
				"  2. Move binary to %s and try again",
			execPath, expectedDir, binaryName, expectedDir)
	}

	// Download new binary
	url := DownloadURL(binaryName)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("release not found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	// Write to temp file in same directory (for atomic rename)
	tmpFile, err := os.CreateTemp(expectedDir, binaryName+"-update-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		return fmt.Errorf("write temp file: %w", err)
	}
	tmpFile.Close()

	// Make executable
	if err := os.Chmod(tmpPath, 0755); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	// Replace current binary
	if err := os.Rename(tmpPath, execPath); err != nil {
		// If rename fails (cross-device), try copy
		return copyFile(tmpPath, execPath)
	}

	return nil
}

// DocType represents the type of documentation.
type DocType string

const (
	DocMain DocType = ""     // Main plugin page
	DocUser DocType = "user" // End user guide
	DocTech DocType = "tech" // Developer/technical docs
)

// DocsURL returns the documentation URL for a plugin.
// If pluginName is empty, returns the main docs page.
// docType specifies user/tech docs, empty for main page.
func DocsURL(pluginName string, docType ...DocType) string {
	if pluginName == "" {
		return fmt.Sprintf("https://%s", DocsHost)
	}
	dt := DocMain
	if len(docType) > 0 {
		dt = docType[0]
	}
	if dt == "" {
		return fmt.Sprintf("https://%s/plugins/%s.html", DocsHost, pluginName)
	}
	return fmt.Sprintf("https://%s/plugins/%s_%s.html", DocsHost, pluginName, dt)
}

// TaskfileURL returns the remote taskfile URL for a plugin.
func TaskfileURL(pluginName string) string {
	return fmt.Sprintf("https://%s/taskfiles/plugins/%s.yml", DocsHost, pluginName)
}

// OpenDocs opens the plugin's documentation in the default browser.
func OpenDocs(pluginName string, docType ...DocType) error {
	url := DocsURL(pluginName, docType...)
	return openBrowser(url)
}

// openBrowser opens a URL in the system's default browser.
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	default: // linux, etc.
		cmd = "xdg-open"
		args = []string{url}
	}

	return execCommand(cmd, args...)
}

// execCommand runs a command.
func execCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Start()
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
