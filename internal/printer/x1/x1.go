package x1

import (
	"context"
	"fmt"
	"time"

	"github.com/joeblew999/plugs/internal/lan"
	"github.com/joeblew999/plugs/internal/printer"
)

// Connect dials a Bambu X1 in LAN Mode and performs login.
func Connect(ctx context.Context, opts printer.Options) (*lan.Client, error) {
	if opts.IP == "" || opts.AccessCode == "" {
		return nil, fmt.Errorf("ip and access code required")
	}
	ctx, cancel := withTimeout(ctx, opts.Timeout)
	defer cancel()
	return lan.Dial(ctx, opts.IP, opts.AccessCode, opts.Insecure)
}

func withTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	if d > 0 {
		return context.WithTimeout(ctx, d)
	}
	return context.WithCancel(ctx)
}
