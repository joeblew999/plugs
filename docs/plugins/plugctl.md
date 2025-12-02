# plugctl

Plugin manager for late-bound binary plugins.

**Docs:** [User Guide](plugctl_user.md) · [Technical](plugctl_tech.md)

## Usage

```sh
# List available plugins
plugctl list

# Install plugins
plugctl install x1ctl us-task us-task-ui

# Update all plugins
plugctl update

# Show versions
plugctl version --all
```

## Plugin Directory

See [User Guide](plugctl_user.md#install-locations) for platform-specific paths.

Default: `~/.local/bin/ubuntusoftware/` (Linux/macOS), `%LOCALAPPDATA%\ubuntusoftware\bin\` (Windows)

Override with `$US_BIN` env var.

## Self-Update

```sh
plugctl update --self     # update plugctl itself
plugctl version --check   # check for updates
```
