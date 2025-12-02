# us-task Technical Documentation

## Architecture

us-task is a repackaged build of [go-task/task](https://github.com/go-task/task), the popular Taskfile runner.

## Why Repackage?

- **Consistent naming**: `us-task` fits our plugin naming convention
- **Managed updates**: Update via `plugctl update` alongside other plugins
- **Single location**: All tools in `~/.local/bin/ubuntusoftware/`

## Build Configuration

From `plugins.json`:

```json
{
  "name": "task",
  "alias": "us-task",
  "repo": "github.com/go-task/task",
  "tag": "v3.40.1",
  "path": "cmd/task"
}
```

## Version Pinning

The external plugin system pins to specific tags for reproducibility. Update the tag in `plugins.json` to upgrade.

## Source Code

- [go-task/task](https://github.com/go-task/task) - Upstream repository
- [plugins.json](https://github.com/joeblew999/plugs/blob/main/plugins.json) - Plugin registry
