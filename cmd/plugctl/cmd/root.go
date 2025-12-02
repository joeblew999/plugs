package cmd

import (
	"fmt"
	"os"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plugctl",
	Short: "Plugin manager for late-bound binary plugins",
	Long: `plugctl manages plugins that self-update from GitHub Releases.

Plugins install to ~/.local/bin/ubuntusoftware/ (override with $US_BIN).
Add this directory to your PATH to use installed plugins.

Example:
  plugctl list              # list available plugins
  plugctl install x1ctl     # install a plugin
  plugctl update            # update all plugins
  plugctl uninstall x1ctl   # remove a plugin
  plugctl paths             # show install locations
  plugctl clean             # remove all plugins`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(versionCmd)

	// Show plugin dir in help
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Ensure plugin dir exists on any command
		_ = version.EnsurePluginDir()
	}
}
