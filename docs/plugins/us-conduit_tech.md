# us-conduit Technical Documentation

## Architecture

us-conduit is a repackaged build of [ConduitIO/conduit](https://github.com/ConduitIO/conduit), a data integration platform.

## Build Configuration

From `plugins.json`:

```json
{
  "name": "conduit",
  "alias": "us-conduit",
  "repo": "github.com/ConduitIO/conduit",
  "tag": "v0.12.2",
  "path": "cmd/conduit"
}
```

## Source Code

- [ConduitIO/conduit](https://github.com/ConduitIO/conduit) - Upstream repository
- [plugins.json](https://github.com/joeblew999/plugs/blob/main/plugins.json) - Plugin registry
