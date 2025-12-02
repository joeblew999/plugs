# us-task-ui Technical Documentation

## Architecture

us-task-ui is a repackaged build of [titpetric/task-ui](https://github.com/titpetric/task-ui), providing a web interface for Taskfile.yml.

## Build Configuration

From `plugins.json`:

```json
{
  "name": "task-ui",
  "alias": "us-task-ui",
  "repo": "github.com/titpetric/task-ui",
  "tag": "v0.0.3",
  "path": "."
}
```

## Source Code

- [titpetric/task-ui](https://github.com/titpetric/task-ui) - Upstream repository
- [plugins.json](https://github.com/joeblew999/plugs/blob/main/plugins.json) - Plugin registry
