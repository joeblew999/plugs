# plugs

Self-updating plugin system for Go binaries.

---

## I'm an End User

*I just want to control my printer.*

Download `x1ctl` for your platform:

{{.EndUserDownloads}}

```sh
./x1ctl status --ip 192.168.1.x --access-code XXX
./x1ctl update    # self-update
```

**Docs:** [x1ctl](plugins/x1ctl.md) · [Bambu X1 Guide](bambu_x1_user.md)

---

## I'm an Operator

*I manage tools and run automation.*

Download `plugctl` to manage everything:

{{.OperatorDownloads}}

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

{{.PluginTable}}

---

## I'm a Developer

*I want to build or extend plugins.*

```sh
git clone {{.RepoURL}} && cd plugs
task build:local      # build for current platform
task ext:build:all    # build external plugins (us-*)
task docs:generate    # regenerate docs
```

**Docs:** [MAINTAINERS.md](MAINTAINERS.md) · [GitHub]({{.RepoURL}})

---

## All Downloads

{{.DownloadTable}}

[GitHub Releases]({{.ReleasesURL}}/latest) · [File an Issue]({{.RepoURL}}/issues)
