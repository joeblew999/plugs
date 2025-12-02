package cmd

import (
	"fmt"
	"runtime"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var pathsCmd = &cobra.Command{
	Use:   "paths",
	Short: "Show install and data directories",
	Long: `Show all directories used by plugctl and plugins.

These paths follow XDG Base Directory Specification:
  - Binaries are installed separately from data
  - Updating/removing plugins won't affect your data

Override with environment variables:
  US_BIN     - Binary install directory
  US_DATA    - Data directory
  US_CONFIG  - Config directory
  US_CACHE   - Cache directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Platform: %s/%s\n\n", runtime.GOOS, runtime.GOARCH)
		fmt.Println("Directories:")
		fmt.Printf("  Binaries:  %s\n", version.PluginDir())
		fmt.Printf("  Data:      %s\n", version.DataDir())
		fmt.Printf("  Config:    %s\n", version.ConfigDir())
		fmt.Printf("  Cache:     %s\n", version.CacheDir())
	},
}

func init() {
	rootCmd.AddCommand(pathsCmd)
}
