# Claude Notes

## Task/CI Sync

GitHub Actions (`.github/workflows/ci.yml`) calls Taskfile tasks:
- `task build:all` - build job + release job

Keep both in sync when renaming tasks.

## Task Namespaces

All tasks are namespaced:
- `build:*` - clean, list, one, local, all, version
- `install:*` - local, remote
- `release:*` - tag, create, list, latest
- `docs:*` - serve, open, check, pages:*
- `gh:*` - auth, open
- `run:*` - fakeprinter, x1ctl
- `test:*` - unit, version, update, all

Top-level tasks are workflow aliases only: `default`, `ci`

## Binaries

Two binaries in `cmd/`:
- `x1ctl` - Cobra CLI, uses subcommands (`version`, `update`)
- `fakeprinter` - flag-based, uses flags (`--version`, `--update`)

Both use `internal/version` for self-update from GitHub releases.
