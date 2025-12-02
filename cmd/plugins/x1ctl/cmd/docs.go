package cmd

import (
	"fmt"

	"github.com/joeblew999/plugs/internal/version"
	"github.com/spf13/cobra"
)

var devDocs bool

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open documentation in browser",
	Long: `Open documentation in the default browser.

Use --dev to open developer/technical documentation instead of user guide.`,
	Run: func(cmd *cobra.Command, args []string) {
		docType := version.DocMain
		if devDocs {
			docType = version.DocTech
		}
		if err := version.OpenDocs(binaryName, docType); err != nil {
			fmt.Printf("Failed to open docs: %v\n", err)
		}
	},
}

func init() {
	docsCmd.Flags().BoolVar(&devDocs, "dev", false, "Open developer/technical docs")
	rootCmd.AddCommand(docsCmd)
}
