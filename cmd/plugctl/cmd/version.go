package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var (
	checkVersion bool
	showAll      bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: `Show plugctl version and optionally plugin versions.

Use --check to check for updates.
Use --all to show versions of all installed plugins.`,
	RunE: runVersion,
}

func init() {
	versionCmd.Flags().BoolVarP(&checkVersion, "check", "c", false, "Check for updates")
	versionCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all installed plugin versions")
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("plugctl %s\n", version.Info())

	if checkVersion {
		latest, needsUpdate, err := version.CheckUpdate()
		if err != nil {
			fmt.Printf("  Update check failed: %v\n", err)
		} else if needsUpdate {
			fmt.Printf("  Update available: %s (run 'plugctl update --self')\n", latest)
		} else {
			fmt.Println("  Up to date")
		}
	}

	if showAll {
		fmt.Println()
		return showPluginVersions()
	}

	return nil
}

func showPluginVersions() error {
	plugins, err := version.ListInstalled()
	if err != nil {
		return fmt.Errorf("list installed: %w", err)
	}

	if len(plugins) == 0 {
		fmt.Println("No plugins installed")
		return nil
	}

	fmt.Println("Installed plugins:")
	for _, p := range plugins {
		ver := getPluginVersion(p.Path)
		fmt.Printf("  %s %s\n", p.Name, ver)
	}

	return nil
}

func getPluginVersion(path string) string {
	// Try running the plugin with --version or version
	for _, arg := range []string{"--version", "version"} {
		cmd := exec.Command(path, arg)
		cmd.Stderr = nil
		out, err := cmd.Output()
		if err == nil {
			// Extract first line
			s := strings.TrimSpace(string(out))
			if idx := strings.Index(s, "\n"); idx > 0 {
				s = s[:idx]
			}
			return s
		}
	}

	// Fallback: check if file exists and is executable
	if info, err := os.Stat(path); err == nil {
		if info.Mode()&0111 != 0 {
			return "(installed)"
		}
	}

	return "(unknown)"
}

func init() {
	// Also support running plugctl with no command to show version
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("plugctl %s\n", version.Info())
			fmt.Printf("Plugin directory: %s\n", version.PluginDir())
			fmt.Println()
			fmt.Println("Run 'plugctl --help' for usage")
		}
	}

	// Add PATH hint to root help
	rootCmd.SetUsageTemplate(rootCmd.UsageTemplate() + fmt.Sprintf(`
Plugin Directory:
  %s
  Add this to your PATH to use installed plugins.
`, filepath.Join("~", ".plugctl", "bin")))
}
