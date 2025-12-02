# Bambu Lab X1 (user)

[Improve this page on GitHub](https://github.com/joeblew999/plugs/blob/main/docs/bambu_x1_user.md) | [File an issue](https://github.com/joeblew999/plugs/issues)

## Downloads (x1ctl)

- Latest release: [releases/latest](https://github.com/joeblew999/plugs/releases/latest)
- Direct binaries:
  - [Linux amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_linux_amd64)
  - [Linux arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_linux_arm64)
  - [macOS amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_darwin_amd64)
  - [macOS arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_darwin_arm64)
  - [Windows amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_windows_amd64.exe)
  - [Windows arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_windows_arm64.exe)

## Quick Connect (LAN Mode)

1. On the printer: enable **LAN Mode**, note the IP and access code.
2. On your PC, run x1ctl:

```sh
# Show build version
x1ctl version

# Basic status (reads first message, surfaces firmware info)
x1ctl status --ip 192.168.1.50 --access-code ABCD

# Echo test (send/receive JSON)
x1ctl echo --ip 192.168.1.50 --access-code ABCD --message "hi"

# Read-only connection
x1ctl read --ip 192.168.1.50 --access-code ABCD
```

3. If it fails: ensure same subnet, port 8883 open, LAN Mode still on, access code correct.

## Self-Update

```sh
# Check for updates
x1ctl version --check

# Update to latest release
x1ctl update
```

## Notes

- `x1ctl version` prints the embedded build version.
- Keep LAN access on trusted networks; the printer uses a self-signed cert.
- For protocol details and scripting, see the technical doc: [bambu_x1_tech.md](bambu_x1_tech.md).
