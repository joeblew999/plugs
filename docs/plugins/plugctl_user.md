# plugctl User Guide

## Installation

Download from [Releases](https://github.com/joeblew999/plugs/releases/latest):

| Linux | macOS | Windows |
|-------|-------|---------|
| [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_arm64.exe) |

## Quick Start

```sh
# See available plugins
plugctl list

# Install a plugin
plugctl install x1ctl

# Run the installed plugin
x1ctl --help

# Update all plugins
plugctl update
```

## Install Locations

All Ubuntu Software tools install to a dedicated directory to avoid conflicts with other software.

### Binaries

| Platform | Default Location |
|----------|------------------|
| **Linux** | `~/.local/bin/ubuntusoftware/` |
| **macOS** | `~/.local/bin/ubuntusoftware/` |
| **Windows** | `%LOCALAPPDATA%\ubuntusoftware\bin\` |

### Data, Config, Cache

Plugin data is stored separately from binaries so updates/uninstalls don't lose your data:

| Type | Linux | macOS | Windows |
|------|-------|-------|---------|
| **Data** | `~/.local/share/ubuntusoftware/` | `~/Library/Application Support/ubuntusoftware/` | `%LOCALAPPDATA%\ubuntusoftware\` |
| **Config** | `~/.config/ubuntusoftware/` | `~/Library/Application Support/ubuntusoftware/` | `%LOCALAPPDATA%\ubuntusoftware\` |
| **Cache** | `~/.cache/ubuntusoftware/` | `~/Library/Caches/ubuntusoftware/` | `%LOCALAPPDATA%\ubuntusoftware\cache\` |

### Environment Overrides

Override any location with environment variables:

| Variable | What it controls |
|----------|------------------|
| `US_BIN` | Binary install directory |
| `US_DATA` | Data directory |
| `US_CONFIG` | Config directory |
| `US_CACHE` | Cache directory |

### Add to PATH

Add the binary directory to your shell profile:

```sh
# Linux/macOS (bash/zsh)
echo 'export PATH="$HOME/.local/bin/ubuntusoftware:$PATH"' >> ~/.bashrc

# Linux/macOS (fish)
fish_add_path ~/.local/bin/ubuntusoftware

# Windows (PowerShell - run as admin)
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$env:LOCALAPPDATA\ubuntusoftware\bin", "User")
```

## Available Plugins

| Plugin | Description |
|--------|-------------|
| `x1ctl` | Bambu Lab X1 printer control |
| `fakeprinter` | Mock printer for testing |
| `us-task` | Task runner (Taskfile.yml) |
| `us-task-ui` | Web GUI for Taskfiles |
| `us-conduit` | Data pipelines |
| `us-benthos` | Stream processing |

## Commands

```sh
plugctl list              # list available plugins
plugctl list --installed  # list installed plugins
plugctl install NAME      # install a plugin
plugctl update            # update all plugins
plugctl update NAME       # update specific plugin
plugctl update --self     # update plugctl itself
plugctl version           # show version
plugctl docs              # open documentation
```
