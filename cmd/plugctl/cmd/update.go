package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var updateSelf bool

var updateCmd = &cobra.Command{
	Use:   "update [plugin]",
	Short: "Update plugins to latest version",
	Long: `Update one or all installed plugins from GitHub releases.

Without arguments, updates all installed plugins.
With a plugin name, updates only that plugin.
Use --self to update plugctl itself.

Examples:
  plugctl update           # update all installed plugins
  plugctl update x1ctl     # update specific plugin
  plugctl update --self    # update plugctl`,
	Args: cobra.MaximumNArgs(1),
	RunE: runUpdate,
}

func init() {
	updateCmd.Flags().BoolVar(&updateSelf, "self", false, "Update plugctl itself")
}

func runUpdate(cmd *cobra.Command, args []string) error {
	if updateSelf {
		return updateSelfBinary()
	}

	if len(args) > 0 {
		return updateOne(args[0])
	}

	return updateAll()
}

func updateSelfBinary() error {
	fmt.Println("Updating plugctl...")

	latest, needsUpdate, err := version.CheckUpdate()
	if err != nil {
		return fmt.Errorf("check update: %w", err)
	}

	if !needsUpdate {
		fmt.Printf("Already at latest version (%s)\n", version.Version)
		return nil
	}

	fmt.Printf("Updating to %s...\n", latest)
	if err := version.SelfUpdate("plugctl"); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Println("Updated successfully. Restart plugctl to use new version.")
	return nil
}

func updateOne(name string) error {
	// Check if plugin is installed
	destPath := filepath.Join(version.PluginDir(), name)
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		return fmt.Errorf("plugin %q not installed (use 'plugctl install %s' first)", name, name)
	}

	fmt.Printf("Updating %s...\n", name)
	if err := version.UpdatePlugin(name); err != nil {
		return err
	}

	fmt.Printf("Updated: %s\n", destPath)
	return nil
}

func updateAll() error {
	plugins, err := version.ListInstalled()
	if err != nil {
		return fmt.Errorf("list installed: %w", err)
	}

	if len(plugins) == 0 {
		fmt.Println("No plugins installed")
		return nil
	}

	fmt.Printf("Updating %d plugins...\n", len(plugins))

	var failed []string
	for _, p := range plugins {
		fmt.Printf("  %s... ", p.Name)
		if err := version.UpdatePlugin(p.Name); err != nil {
			fmt.Printf("FAILED: %v\n", err)
			failed = append(failed, p.Name)
		} else {
			fmt.Println("ok")
		}
	}

	if len(failed) > 0 {
		return fmt.Errorf("%d plugins failed to update", len(failed))
	}

	fmt.Println("All plugins updated")
	return nil
}
