package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var echoMessage string

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Send a test JSON echo message",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := connect()
		defer client.Close()

		log.Printf("connected to %s; sending echo payload...", ip)
		payload := map[string]any{
			"cmd": "echo",
			"msg": echoMessage,
		}
		if err := client.SendJSON(ctx, payload); err != nil {
			log.Fatalf("send: %v", err)
		}
		resp, err := client.ReadRaw(ctx)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Printf("printer replied:\n%s\n", string(resp))
	},
}

func init() {
	echoCmd.Flags().StringVar(&echoMessage, "message", "hello from x1ctl", "Echo message to send")
}
