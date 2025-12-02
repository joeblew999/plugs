# 3D Printer Tools

Control your Bambu Lab X1 printer over LAN. No cloud required.

## What's Here

**x1ctl** - Command-line tool to connect to your Bambu Lab X1 printer over your local network. Check status, send commands, and monitor prints without going through Bambu's cloud.

**fakeprinter** - A mock printer server for testing and development. Useful if you want to build tools without a real printer connected.

**plugctl** - Plugin manager for installing and updating tools. Manages plugins in `~/.plugctl/bin/`.

## Quick Start

1. Download the binary for your system from the table below
2. Enable LAN Mode on your printer and note the IP address and access code
3. Run: `x1ctl status --ip YOUR_PRINTER_IP --access-code YOUR_CODE`

## Downloads

Pick your platform:

{{.DownloadTable}}
Or get all binaries from [GitHub Releases]({{.ReleasesURL}}/latest).

## Keeping Up to Date

All tools can update themselves. No need to re-download manually:

```sh
x1ctl update           # updates x1ctl to latest
fakeprinter --update   # updates fakeprinter to latest
plugctl update         # updates all installed plugins
```

## Plugins

{{.PluginLinks}}
See individual plugin pages for detailed documentation.

## Documentation

- **[Bambu X1 User Guide](bambu_x1_user.md)** - How to connect and use x1ctl with your printer
- **[Bambu X1 Technical](bambu_x1_tech.md)** - Protocol details for developers
- **[Fake Printer Guide](fake_printer_user.md)** - Running the mock server for testing

## Contributing

Found a bug? Want to add a feature?

- [File an issue]({{.RepoURL}}/issues)
- [View source on GitHub]({{.RepoURL}})
- [Fork this template](MAINTAINERS.md) for your own Go binary projects
