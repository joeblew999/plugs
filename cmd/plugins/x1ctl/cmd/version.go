package cmd

import (
	"fmt"
	"log"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

const binaryName = "x1ctl"

var checkUpdate bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build version and check for updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Info())
		if checkUpdate {
			latest, needsUpdate, err := version.CheckUpdate()
			if err != nil {
				log.Printf("check update: %v", err)
				return
			}
			if needsUpdate {
				fmt.Printf("Update available: %s (run '%s update' to upgrade)\n", latest, binaryName)
			} else {
				fmt.Println("You are running the latest version.")
			}
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update to the latest release from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		latest, needsUpdate, err := version.CheckUpdate()
		if err != nil {
			log.Fatalf("check update: %v", err)
		}
		if !needsUpdate {
			fmt.Println("Already at latest version:", version.Version)
			return
		}
		fmt.Printf("Updating from %s to %s...\n", version.Version, latest)
		if err := version.SelfUpdate(binaryName); err != nil {
			log.Fatalf("update failed: %v", err)
		}
		fmt.Println("Update complete. Restart the program to use the new version.")
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&checkUpdate, "check", "c", false, "Check for available updates")
	rootCmd.AddCommand(updateCmd)
}
