package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var localPath string

var installCmd = &cobra.Command{
	Use:   "install <plugin>",
	Short: "Install a plugin from GitHub releases or local file",
	Long: `Install a plugin to the local plugin directory.

By default, downloads from GitHub releases.
Use --local to install from a local file (useful for development).

Examples:
  plugctl install x1ctl                          # from GitHub
  plugctl install --local ./dist/x1ctl_darwin_arm64  # from local file`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInstall,
}

func init() {
	installCmd.Flags().StringVarP(&localPath, "local", "l", "", "Install from local file instead of GitHub")
}

func runInstall(cmd *cobra.Command, args []string) error {
	if localPath != "" {
		return installLocal()
	}

	if len(args) == 0 {
		return fmt.Errorf("plugin name required (or use --local)")
	}

	return installRemote(args[0])
}

func installRemote(name string) error {
	fmt.Printf("Installing %s from GitHub...\n", name)

	if err := version.InstallPlugin(name); err != nil {
		return err
	}

	destPath := filepath.Join(version.PluginDir(), name)
	fmt.Printf("Installed: %s\n", destPath)
	return nil
}

func installLocal() error {
	// Verify file exists
	info, err := os.Stat(localPath)
	if err != nil {
		return fmt.Errorf("file not found: %s", localPath)
	}
	if info.IsDir() {
		return fmt.Errorf("expected file, got directory: %s", localPath)
	}

	fmt.Printf("Installing from %s...\n", localPath)

	if err := version.InstallLocal(localPath); err != nil {
		return err
	}

	// Get installed name
	name := filepath.Base(localPath)
	for _, suffix := range []string{
		"_darwin_amd64", "_darwin_arm64",
		"_linux_amd64", "_linux_arm64",
		"_windows_amd64.exe", "_windows_arm64.exe",
	} {
		if len(name) > len(suffix) {
			name = name[:len(name)-len(suffix)]
			break
		}
	}

	destPath := filepath.Join(version.PluginDir(), name)
	fmt.Printf("Installed: %s\n", destPath)
	return nil
}
