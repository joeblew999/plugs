[![Ubuntu Software](https://www.ubuntusoftware.net/images/logo.svg)](https://www.ubuntusoftware.net)

*AI Infrastructure for Design and Documents*

---

# plugs

A plugin system for hardware control tools. Self-updating binaries distributed via GitHub Releases.

## What's Here

**plugctl** - The plugin manager. Use it to install, update, and manage plugins.

**Plugins** - The actual tools that talk to hardware:
- **x1ctl** - Control Bambu Lab X1 printers over LAN
- **fakeprinter** - Mock printer server for testing

## Quick Start

### Option 1: Use plugctl (recommended)

Download `plugctl` from the table below, then:

```sh
plugctl install x1ctl       # install x1ctl
plugctl list --installed    # see what's installed
```

Add `~/.plugctl/bin` to your PATH.

### Option 2: Direct download

Download the binary for your platform from the table below and run it directly.

## Download plugctl

| Linux | macOS | Windows |
|---|---|---|
| [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_arm64.exe) |

Then use `plugctl install <plugin>` to install plugins.

Or download plugins directly from [GitHub Releases](https://github.com/joeblew999/plugs/releases/latest).

## Using plugctl

```sh
# List and install
plugctl list                # available plugins
plugctl list --installed    # installed plugins
plugctl install x1ctl       # install from GitHub

# Updates
plugctl update              # update all plugins
plugctl update x1ctl        # update one plugin
plugctl update --self       # update plugctl itself

# Version info
plugctl version             # show plugctl version
plugctl version --check     # check for updates
plugctl version --all       # show all installed versions

# Local development
plugctl install --local ./dist/x1ctl_darwin_arm64
```

## Self-Update

All plugins can update themselves:

```sh
# Cobra-style plugins (subcommands)
x1ctl update              # update to latest
x1ctl version --check     # check for updates

# Flag-style plugins
fakeprinter --update      # update to latest
fakeprinter --version     # show version
```

## Plugin Documentation

- [x1ctl](plugins/x1ctl.md) - Bambu Lab X1 printer control
- [fakeprinter](plugins/fakeprinter.md) - Mock printer for testing

## Technical Docs

- [Bambu X1 User Guide](bambu_x1_user.md) - Connect to your printer
- [Bambu X1 Protocol](bambu_x1_tech.md) - MQTT/TLS protocol details
- [Fake Printer Setup](fake_printer_user.md) - Run mock server

## For Developers

- [MAINTAINERS.md](MAINTAINERS.md) - Fork this template for your own plugins
- [GitHub Repository](https://github.com/joeblew999/plugs)
- [File an Issue](https://github.com/joeblew999/plugs/issues)
