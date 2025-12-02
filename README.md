# 3d-printers

Late-bound plugin system with self-update from GitHub Releases.

**Docs & Downloads:** https://joeblew999.github.io/3d-printers/

## What This Is

A template for building Go binaries that:
- Self-update from GitHub Releases
- Install to user's home directory (`~/.plugctl/bin/`)
- Can be managed by a central client tool

## Plugins

| Plugin | Description |
|--------|-------------|
| `x1ctl` | CLI for Bambu Lab X1 printers |
| `fakeprinter` | Mock printer server for testing |

## Client Tool

`plugctl` manages plugins without requiring sudo or system-wide installs:

```sh
plugctl list              # list available plugins
plugctl install x1ctl     # install to ~/.plugctl/bin/
plugctl update            # update all installed plugins
plugctl install --local ./dist/x1ctl_darwin_arm64  # install local build
```

Add `~/.plugctl/bin` to your PATH to use installed plugins.

## Self-Update

Every plugin can update itself:

```sh
x1ctl update              # update x1ctl
x1ctl version --check     # check for updates
fakeprinter --update      # update fakeprinter
```

## Fork This Template

Create your own late-bound plugin system:

1. Fork this repo
2. Update `Taskfile.yml`:
   ```yaml
   GITHUB_USER: your-username
   GITHUB_REPO: your-repo
   ```
3. Update `internal/version/version.go` with your repo
4. Add plugins in `cmd/plugins/your-plugin/`
5. Add `README.md` in each plugin folder (merged into generated docs)
6. Run `task docs:generate`
7. Tag `v0.1.0` to trigger first release

See [MAINTAINERS.md](docs/MAINTAINERS.md) for details.

## Development

```sh
task build:local        # build for host platform
task build:all          # build all platforms
task test:all           # run all tests
task docs:generate      # regenerate docs from plugin READMEs
task run:plugctl -- list
```

## License

MIT
