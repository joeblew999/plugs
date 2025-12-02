# Demo Printer (technical)

## Behavior
- Listens on TLS WebSocket `wss://localhost:8883` with a self-signed cert.
- Sends an initial JSON: `{"hello":"fake-printer","ts":<unix>}`.
- Echoes any JSON you send, adding a timestamp.
- Ignores access code (testing only).

## Running
- `task fakeprinter` or `./fakeprinter_<os>_<arch> -addr :8883`
- Keep `-insecure` true on x1ctl when talking to the mock (self-signed cert).

## Extending the mock
- Code lives at `cmd/fakeprinter`. You can add handlers for specific commands or simulate status/telemetry shapes to mirror the X1 protocol.
- For real protocol shapes, inspect OrcaSlicer/HA integrations and mirror the JSON fields you care about.
