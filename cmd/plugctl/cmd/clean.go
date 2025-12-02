package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var (
	cleanAll   bool
	cleanForce bool
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove all installed plugins",
	Long: `Remove all installed plugins from the plugin directory.

By default, only removes binaries. Use --all to also remove data/config/cache.

WARNING: This is destructive. Use --force to skip confirmation.

Examples:
  plugctl clean           # remove all plugins (keeps data)
  plugctl clean --all     # remove plugins AND all data
  plugctl clean --force   # skip confirmation prompt`,
	Run: func(cmd *cobra.Command, args []string) {
		dirs := []struct {
			name string
			path string
			bin  bool
		}{
			{"Binaries", version.PluginDir(), true},
		}

		if cleanAll {
			dirs = append(dirs,
				struct {
					name string
					path string
					bin  bool
				}{"Data", version.DataDir(), false},
				struct {
					name string
					path string
					bin  bool
				}{"Config", version.ConfigDir(), false},
				struct {
					name string
					path string
					bin  bool
				}{"Cache", version.CacheDir(), false},
			)
		}

		// Show what will be deleted
		fmt.Println("The following directories will be removed:")
		for _, d := range dirs {
			if _, err := os.Stat(d.path); err == nil {
				fmt.Printf("  %s: %s\n", d.name, d.path)
			} else {
				fmt.Printf("  %s: %s (not found)\n", d.name, d.path)
			}
		}

		// Confirm unless --force
		if !cleanForce {
			fmt.Print("\nAre you sure? [y/N]: ")
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))
			if response != "y" && response != "yes" {
				fmt.Println("Aborted")
				return
			}
		}

		// Delete directories
		var errors []string
		for _, d := range dirs {
			if _, err := os.Stat(d.path); os.IsNotExist(err) {
				continue
			}
			if err := os.RemoveAll(d.path); err != nil {
				errors = append(errors, fmt.Sprintf("%s: %v", d.name, err))
			} else {
				fmt.Printf("Removed: %s\n", d.path)
			}
		}

		if len(errors) > 0 {
			fmt.Println("\nErrors:")
			for _, e := range errors {
				fmt.Printf("  %s\n", e)
			}
			os.Exit(1)
		}

		fmt.Println("\nClean complete")
		if !cleanAll {
			fmt.Println("Note: Data/config preserved. Use --all to remove everything.")
		}
	},
}

func init() {
	cleanCmd.Flags().BoolVar(&cleanAll, "all", false, "Also remove data, config, and cache directories")
	cleanCmd.Flags().BoolVarP(&cleanForce, "force", "f", false, "Skip confirmation prompt")
	rootCmd.AddCommand(cleanCmd)
}
