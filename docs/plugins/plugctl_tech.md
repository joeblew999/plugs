# plugctl Technical Documentation

## Architecture

plugctl manages plugins by:
1. Fetching binary info from GitHub Releases API
2. Downloading platform-specific binaries
3. Installing to XDG-compliant directories (see [User Guide](plugctl_user.md#install-locations))

## Plugin Registry

The registry is `plugins.json` in the repo root. It defines:
- Local plugins (built from this repo)
- External plugins (fetched from other repos)

```json
{
  "plugins": [
    {"name": "x1ctl", "path": "cmd/plugins/x1ctl", "style": "cobra"}
  ],
  "external": [
    {"name": "task", "alias": "us-task", "repo": "github.com/go-task/task", "tag": "v3.40.1"}
  ]
}
```

## Download URLs

Binaries follow the naming convention:
```
{name}_{os}_{arch}[.exe]
```

Example: `x1ctl_darwin_arm64`, `plugctl_windows_amd64.exe`

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `US_BIN` | XDG bin home + ubuntusoftware | Binary install directory |
| `US_DATA` | XDG data home + ubuntusoftware | Data directory |
| `US_CONFIG` | XDG config home + ubuntusoftware | Config directory |
| `US_CACHE` | XDG cache home + ubuntusoftware | Cache directory |
| `PLUGCTL_BIN` | (legacy) | Legacy support, prefer `US_BIN` |

## Source Code

- [cmd/plugctl/](https://github.com/joeblew999/plugs/tree/main/cmd/plugctl) - Main binary
- [internal/version/](https://github.com/joeblew999/plugs/tree/main/internal/version) - Update/install logic

## Adding a Plugin

1. Add entry to `plugins.json`
2. Implement in `cmd/plugins/{name}/`
3. Tag a release

CI builds all plugins for all platforms.
