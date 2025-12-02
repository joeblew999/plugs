package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Connect and print the first message",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := connect()
		defer client.Close()

		log.Printf("connected to %s; waiting for first message...", ip)
		msg, err := client.ReadRaw(ctx)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Printf("printer said:\n%s\n", string(msg))
	},
}
