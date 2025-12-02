# plugs

Self-updating plugin system for Go binaries.

---

## I'm an End User

*I just want to control my printer.*

Download `x1ctl` for your platform:

| Linux | macOS | Windows |
|-------|-------|--------|
| [amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_windows_arm64.exe) |


```sh
./x1ctl status --ip 192.168.1.x --access-code XXX
./x1ctl update    # self-update
```

**Docs:** [x1ctl](plugins/x1ctl.md) · [Bambu X1 Guide](bambu_x1_user.md)

---

## I'm an Operator

*I manage tools and run automation.*

Download `plugctl` to manage everything:

| Linux | macOS | Windows |
|-------|-------|--------|
| [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_arm64.exe) |


```sh
# Install plugins
plugctl install x1ctl us-task us-task-ui
plugctl list --installed
plugctl update

# Run Taskfiles
us-task --list
us-task-ui    # web GUI
```

**Install location:** `~/.local/bin/ubuntusoftware/` (Linux/macOS) or `%LOCALAPPDATA%\ubuntusoftware\bin\` (Windows). Add to PATH. See [plugctl user guide](plugins/plugctl_user.md) for details.

**Available plugins:**

| Plugin | Description |
|--------|-------------|
| `x1ctl` | CLI for Bambu Lab X1 printers (LAN mode) |
| `fakeprinter` | Mock printer server for testing |
| `us-task` | From github.com/go-task/task |
| `us-task-ui` | From github.com/titpetric/task-ui |
| `us-conduit` | From github.com/ConduitIO/conduit |
| `us-benthos` | From github.com/redpanda-data/benthos |


---

## I'm a Developer

*I want to build or extend plugins.*

```sh
git clone https://github.com/joeblew999/plugs && cd plugs
task build:local      # build for current platform
task ext:build:all    # build external plugins (us-*)
task docs:generate    # regenerate docs
```

**Docs:** [MAINTAINERS.md](MAINTAINERS.md) · [GitHub](https://github.com/joeblew999/plugs)

---

## All Downloads

| | Linux | macOS | Windows |
|---|---|---|---|
| **x1ctl** | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/x1ctl_windows_arm64.exe) |
| **fakeprinter** | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/fakeprinter_windows_arm64.exe) |
| **plugctl** | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_amd64) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_linux_arm64) | [Intel](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_amd64) / [Apple Silicon](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_darwin_arm64) | [amd64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/plugs/releases/latest/download/plugctl_windows_arm64.exe) |


[GitHub Releases](https://github.com/joeblew999/plugs/releases/latest) · [File an Issue](https://github.com/joeblew999/plugs/issues)
