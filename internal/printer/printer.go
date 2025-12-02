package printer

import (
	"context"
	"time"
)

// Options for connecting to a printer over LAN.
type Options struct {
	IP         string
	AccessCode string
	Insecure   bool
	Timeout    time.Duration
}

// Session is the minimal interface shared by printer drivers.
type Session interface {
	Close() error
	ReadRaw(ctx context.Context) ([]byte, error)
	SendJSON(ctx context.Context, payload any) error
}
