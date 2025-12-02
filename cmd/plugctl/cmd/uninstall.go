package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall <plugin>",
	Short: "Remove an installed plugin",
	Long: `Remove an installed plugin from the plugin directory.

This only removes the binary, not any data/config the plugin may have stored.
Use 'plugctl paths' to see where data is stored.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		pluginPath := filepath.Join(version.PluginDir(), name)

		if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
			fmt.Printf("Plugin %q not installed\n", name)
			os.Exit(1)
		}

		if err := os.Remove(pluginPath); err != nil {
			fmt.Printf("Failed to uninstall %q: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("Uninstalled: %s\n", name)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
