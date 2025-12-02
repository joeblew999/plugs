package cmd

import (
	"fmt"
	"os"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var listInstalled bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available or installed plugins",
	Long: `List plugins available from GitHub releases or installed locally.

By default, shows plugins available for download.
Use --installed to show plugins in the local plugin directory.`,
	RunE: runList,
}

func init() {
	listCmd.Flags().BoolVarP(&listInstalled, "installed", "i", false, "List installed plugins instead of available")
}

func runList(cmd *cobra.Command, args []string) error {
	if listInstalled {
		return listInstalledPlugins()
	}
	return listAvailablePlugins()
}

func listAvailablePlugins() error {
	plugins, err := version.ListAvailable()
	if err != nil {
		return fmt.Errorf("fetch available plugins: %w", err)
	}

	if len(plugins) == 0 {
		fmt.Println("No plugins available for this platform")
		return nil
	}

	fmt.Println("Available plugins:")
	for _, name := range plugins {
		fmt.Printf("  %s\n", name)
	}
	fmt.Printf("\nInstall with: plugctl install <name>\n")
	return nil
}

func listInstalledPlugins() error {
	plugins, err := version.ListInstalled()
	if err != nil {
		return fmt.Errorf("scan plugin dir: %w", err)
	}

	if len(plugins) == 0 {
		fmt.Printf("No plugins installed in %s\n", version.PluginDir())
		fmt.Println("\nInstall with: plugctl install <name>")
		return nil
	}

	fmt.Printf("Installed plugins (%s):\n", version.PluginDir())
	for _, p := range plugins {
		fmt.Printf("  %s\n", p.Name)
	}

	// Check if plugin dir is in PATH
	pluginDir := version.PluginDir()
	path := os.Getenv("PATH")
	if path != "" {
		inPath := false
		for _, dir := range splitPath(path) {
			if dir == pluginDir {
				inPath = true
				break
			}
		}
		if !inPath {
			fmt.Printf("\nTip: Add %s to your PATH\n", pluginDir)
		}
	}

	return nil
}

func splitPath(path string) []string {
	// Handle both Unix (:) and Windows (;) path separators
	if len(path) > 0 && path[0] == '/' {
		return splitOn(path, ':')
	}
	// Could be Windows or just a single path element
	if idx := indexOf(path, ';'); idx >= 0 {
		return splitOn(path, ';')
	}
	return splitOn(path, ':')
}

func splitOn(s string, sep byte) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	result = append(result, s[start:])
	return result
}

func indexOf(s string, b byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i
		}
	}
	return -1
}
