# plugs

A plugin system for Go binaries that talk to hardware. Self-updates from GitHub Releases.

**Docs:** https://plugs.ubuntusoftware.net

## Hosting

- **Docs & Binaries:** GitHub Pages + GitHub Releases
- **DNS:** Cloudflare (`plugs.ubuntusoftware.net` → GitHub Pages)

## Overview

Plugs provides a framework for building CLI tools that communicate with hardware devices (printers, IoT, embedded systems) and distribute them as self-updating binaries.

Features:
- Self-update from GitHub Releases (no package manager required)
- User-local installs to `~/.plugctl/bin/` (no sudo)
- Central management via `plugctl`
- Cross-platform: Linux, macOS, Windows (amd64/arm64)

## Plugins

| Plugin | Description |
|--------|-------------|
| `x1ctl` | CLI for Bambu Lab X1 printers (LAN mode) |
| `fakeprinter` | Mock printer server for testing |

## Install

Download `plugctl` from [Releases](https://github.com/joeblew999/plugs/releases/latest), then:

```sh
plugctl install x1ctl       # install a plugin
plugctl list                # list available plugins
plugctl list --installed    # list installed plugins
```

Add `~/.plugctl/bin` to your PATH.

## Usage

### Managing Plugins

```sh
plugctl install x1ctl                # install from GitHub releases
plugctl install --local ./my-binary  # install local build
plugctl update                       # update all plugins
plugctl update x1ctl                 # update specific plugin
plugctl update --self                # update plugctl itself
plugctl version --check              # check for updates
plugctl version --all                # show all installed versions
```

### Plugin Self-Update

Every plugin can update itself:

```sh
x1ctl update              # update x1ctl
x1ctl version --check     # check for updates
fakeprinter --update      # update fakeprinter (flag-based CLI)
fakeprinter --version     # show version
```

## Building Your Own Plugin

1. Fork this repo
2. Update `Taskfile.yml` with your GitHub user/repo
3. Update `internal/version/version.go` with your repo
4. Add your plugin in `cmd/plugins/your-plugin/`
5. Tag a release to trigger CI

See [MAINTAINERS.md](docs/MAINTAINERS.md) for details.

## Development

```sh
task build:local          # build for current platform
task build:all            # build all platforms
task test:all             # run tests
task run:plugctl -- list  # run plugctl
```

## License

MIT
