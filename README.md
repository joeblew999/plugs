# plugs

Self-updating plugin system for Go binaries. **Docs:** https://plugs.ubuntusoftware.net

## Quick Start

| You Are | Download | Then Run |
|---------|----------|----------|
| **End User** | `x1ctl` | `./x1ctl status --ip ... --access-code ...` |
| **Operator** | `plugctl` + `us-task` | `plugctl install x1ctl` / `us-task --list` |
| **Developer** | Clone repo | `task build:local` |

## Binaries

| Binary | Who | What |
|--------|-----|------|
| `x1ctl` | End User | Control Bambu Lab X1 printers |
| `fakeprinter` | End User | Mock printer for testing |
| `plugctl` | Operator | Install/update plugins |
| `us-task` | Operator | Run Taskfiles (our Task build) |
| `us-task-ui` | Operator | Web GUI for Taskfiles |
| `us-conduit` | Operator | Data pipelines |
| `us-benthos` | Operator | Stream processing |

All binaries self-update: `x1ctl update` or `plugctl update`

## End Users

Download from [Releases](https://github.com/joeblew999/plugs/releases/latest):

```sh
./x1ctl status --ip 192.168.1.x --access-code XXX
./x1ctl update    # self-update
./x1ctl docs      # open documentation
```

## Operators

Manage plugins and run Taskfiles:

```sh
# Install plugins
plugctl install x1ctl us-task us-task-ui
plugctl list --installed
plugctl update

# Run tasks
us-task --list
us-task-ui        # opens web GUI
```

### Install Locations

Plugins install to dedicated directories (XDG-compliant):

| Platform | Binaries | Data |
|----------|----------|------|
| **Linux** | `~/.local/bin/ubuntusoftware/` | `~/.local/share/ubuntusoftware/` |
| **macOS** | `~/.local/bin/ubuntusoftware/` | `~/Library/Application Support/ubuntusoftware/` |
| **Windows** | `%LOCALAPPDATA%\ubuntusoftware\bin\` | `%LOCALAPPDATA%\ubuntusoftware\` |

Add to PATH, or override with `US_BIN` env var. See [plugctl docs](https://plugs.ubuntusoftware.net/plugins/plugctl_user.html) for details.

## Developers

```sh
git clone https://github.com/joeblew999/plugs && cd plugs
task build:local      # build for current platform
task ext:build:all    # build external plugins (us-*)
task docs:generate    # regenerate docs
```

### Adding Plugins

1. Add code in `cmd/plugins/your-plugin/`
2. Add entry to `plugins.json`
3. `task build:local && task docs:generate`
4. Tag release → CI builds all platforms

### Project Structure

```
cmd/plugins/          # End user binaries
cmd/plugctl/          # Operator binary
internal/             # Dev tools (not released)
plugins.json          # Source of truth
taskfiles/external.yml # External plugin builds
```

## Naming

- Local plugins: `x1ctl`, `fakeprinter`, `plugctl`
- External plugins: `us-*` prefix (avoids conflicts)

## Registry

`plugins.json` is the central registry. To add a plugin:

1. Fork this repo
2. Add your plugin to `plugins.json`
3. Submit a PR

We build and release all registered plugins.

## License

MIT
