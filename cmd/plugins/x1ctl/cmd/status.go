package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/joeblew999/plugs/internal/printer"
	"github.com/joeblew999/plugs/internal/printer/x1"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Connect and read the first message (tries to surface firmware info)",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := connect()
		defer client.Close()

		log.Printf("connected to %s; waiting for first message (status)...", ip)
		msg, err := client.ReadRaw(ctx)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Printf("printer said:\n%s\n", string(msg))
		if ver := extractVersion(msg); ver != "" {
			fmt.Printf("detected firmware/version: %s\n", ver)
		}
	},
}

type lanSession interface {
	Close() error
	ReadRaw(ctx context.Context) ([]byte, error)
	SendJSON(ctx context.Context, payload any) error
}

func connect() (lanSession, context.Context) {
	if ip == "" || accessCode == "" {
		log.Fatalf("ip and access-code are required")
	}
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	client, err := x1.Connect(ctx, printer.Options{
		IP:         ip,
		AccessCode: accessCode,
		Insecure:   insecure,
		Timeout:    timeout,
	})
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	return client, ctx
}
