# Demo Printer (user)

[Improve this page on GitHub](https://github.com/joeblew999/plugs/blob/main/docs/fake_printer_user.md) | [File an issue](https://github.com/joeblew999/plugs/issues)

This mock printer lets you try a printer driver without real hardware.

## Downloads

- Latest release: [releases/latest](https://github.com/joeblew999/plugs/releases/latest)
- Fake printer binaries (mock server):
  - [Linux amd64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_linux_amd64)
  - [Linux arm64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_linux_arm64)
  - [macOS amd64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_darwin_amd64)
  - [macOS arm64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_darwin_arm64)
  - [Windows amd64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_windows_amd64.exe)
  - [Windows arm64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_windows_arm64.exe)

## Quick Start

1. Start the mock server:
   ```sh
   ./fakeprinter -addr :8883
   ```

2. Connect with your LAN client/driver to `wss://localhost:8883` (self-signed TLS, any access code). The mock sends a hello JSON then echoes your JSON.

3. Stop the mock with Ctrl+C.

## Self-Update

```sh
# Check version
fakeprinter --version

# Check for updates
fakeprinter --check-update

# Update to latest release
fakeprinter --update
```

## Notes

- Self-signed TLS; keep `-insecure` true for the mock.
- Access code is ignored by the mock (for testing only).
- For implementation details, see [fake_printer_tech.md](fake_printer_tech.md).
