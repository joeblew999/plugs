# Bambu Lab X1 (technical)

## Control summary
- LAN Mode: TLS WebSocket on `wss://<ip>:8883` (self-signed cert).
- Auth: send JSON `{"cmd":"login","password":"<access code>"}`.
- Messages: MQTT-like JSON; includes status/telemetry, keepalives, and job upload (G-code frames/metadata). Schemas can change with firmware—verify on your device.
- Discovery: printer broadcasts UDP; static IP works too.
- Security: keep on trusted LAN; do not expose 8883 publicly; pin the cert if you can.

## x1ctl modes (quick)
- `-mode read`: connect and print the first message.
- `-mode status`: same as read, then try to surface firmware/version fields from the first JSON.
- `-mode echo`: send a test JSON and print the reply.
- `-version`: print the embedded build version.

## Minimal Go sketch (adapt to your firmware)
```go
d := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
conn, _, _ := d.Dial("wss://192.168.1.50:8883", nil)
defer conn.Close()
conn.WriteJSON(map[string]any{"cmd": "login", "password": "ABCD"})
_, msg, _ := conn.ReadMessage()
log.Printf("printer said: %s", msg)
```
Use the OrcaSlicer LAN sender as the reference for full flows (auth, keepalives, job upload).

## References
- OrcaSlicer (LAN implementation, MIT): https://github.com/SoftFever/OrcaSlicer
- Bambu Studio (vendor source): https://github.com/bambulab/BambuStudio
- Protocol notes (community): https://github.com/bambulab/BambuStudio/blob/main/resources/protocol/README.md
- Home Assistant integration (practical LAN client): https://github.com/greghesp/ha-bambulab

## What to verify when scripting
- Firmware vs. schema: capture traffic on your firmware to confirm fields and commands.
- LAN Mode on, correct access code, same subnet, 8883 reachable.
- Only use `-insecure` on trusted networks; otherwise pin the cert.
