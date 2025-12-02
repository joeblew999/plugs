package lan

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	defaultPort    = "8883"
	defaultTimeout = 10 * time.Second
)

// Client manages a LAN Mode websocket session to a Bambu printer.
type Client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

// Dial connects to the printer, performs a simple login with the LAN access code,
// and returns a ready client. The protocol is subject to firmware changes; this
// focuses on establishing a session for experimentation.
func Dial(ctx context.Context, ip, accessCode string, insecureTLS bool) (*Client, error) {
	u := url.URL{
		Scheme: "wss",
		Host:   net.JoinHostPort(ip, defaultPort),
		Path:   "/",
	}

	d := websocket.Dialer{
		HandshakeTimeout: defaultTimeout,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureTLS, // printer uses self-signed cert
		},
	}

	conn, _, err := d.DialContext(ctx, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dial printer: %w", err)
	}

	c := &Client{conn: conn}
	if err := c.sendLogin(ctx, accessCode); err != nil {
		_ = c.Close()
		return nil, err
	}

	return c, nil
}

// Close terminates the websocket connection.
func (c *Client) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

// SendJSON writes a JSON message to the printer.
func (c *Client) SendJSON(ctx context.Context, payload any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.setWriteDeadline(ctx); err != nil {
		return err
	}
	if err := c.conn.WriteJSON(payload); err != nil {
		return fmt.Errorf("write json: %w", err)
	}
	return nil
}

// ReadRaw reads the next message as raw bytes (useful for debugging protocol changes).
func (c *Client) ReadRaw(ctx context.Context) ([]byte, error) {
	if err := c.setReadDeadline(ctx); err != nil {
		return nil, err
	}
	_, data, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("read message: %w", err)
	}
	return data, nil
}

// ReadJSON unmarshals the next message into v.
func (c *Client) ReadJSON(ctx context.Context, v any) error {
	data, err := c.ReadRaw(ctx)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	return nil
}

func (c *Client) sendLogin(ctx context.Context, accessCode string) error {
	login := map[string]any{
		"cmd":      "login",
		"password": accessCode,
	}
	if err := c.SendJSON(ctx, login); err != nil {
		return fmt.Errorf("send login: %w", err)
	}
	return nil
}

func (c *Client) setReadDeadline(ctx context.Context) error {
	return c.conn.SetReadDeadline(deadlineFromCtx(ctx))
}

func (c *Client) setWriteDeadline(ctx context.Context) error {
	return c.conn.SetWriteDeadline(deadlineFromCtx(ctx))
}

func deadlineFromCtx(ctx context.Context) time.Time {
	if deadline, ok := ctx.Deadline(); ok {
		return deadline
	}
	return time.Now().Add(defaultTimeout)
}
