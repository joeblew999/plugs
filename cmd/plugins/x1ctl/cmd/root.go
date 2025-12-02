package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	ip         string
	accessCode string
	insecure   bool
	timeout    time.Duration
)

var rootCmd = &cobra.Command{
	Use:   "x1ctl",
	Short: "CLI for Bambu Lab X1 LAN control",
}

// Execute runs the CLI.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ip, "ip", "", "Printer LAN IP address")
	rootCmd.PersistentFlags().StringVar(&accessCode, "access-code", "", "Printer LAN access code (from device screen)")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", true, "Allow self-signed TLS from printer")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", 15*time.Second, "Dial/read timeout")
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(versionCmd)
}
