# User Quick Start

## Download

Get the binary for your platform from [Releases](https://github.com/joeblew999/plugs/releases/latest) or use the direct links on the [main page](index.md).

## x1ctl (Bambu X1 printer control)

```sh
# Check version
x1ctl version

# Connect to printer (LAN mode)
x1ctl status --ip 192.168.1.50 --access-code ABCD

# Echo test
x1ctl echo --ip 192.168.1.50 --access-code ABCD --message "test"

# Update to latest
x1ctl update
```

See [Bambu X1 User Guide](bambu_x1_user.md) for full details.

## fakeprinter (demo/test server)

```sh
# Run fake printer
fakeprinter -addr :8883

# Check version
fakeprinter --version

# Update to latest
fakeprinter --update
```

See [Fake Printer Guide](fake_printer_user.md) for details.

## Self-Update

All binaries can update themselves from GitHub Releases:

```sh
# x1ctl
x1ctl version --check  # check for updates
x1ctl update           # download latest

# fakeprinter
fakeprinter --check-update
fakeprinter --update
```

## Troubleshooting

- Ensure firewall allows port 8883 (printer uses TLS websocket)
- Check printer is in LAN mode with correct access code
- Run with verbose flags if available

## Contributing

See [MAINTAINERS.md](MAINTAINERS.md) for development setup and how to fork this template for your own binaries.
