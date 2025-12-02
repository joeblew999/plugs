# us-benthos Technical Documentation

## Architecture

us-benthos is a repackaged build of [redpanda-data/benthos](https://github.com/redpanda-data/benthos), a stream processor.

## Build Configuration

From `plugins.json`:

```json
{
  "name": "benthos",
  "alias": "us-benthos",
  "repo": "github.com/redpanda-data/benthos",
  "tag": "v4.35.0",
  "path": "cmd/benthos"
}
```

## Source Code

- [redpanda-data/benthos](https://github.com/redpanda-data/benthos) - Upstream repository
- [plugins.json](https://github.com/joeblew999/plugs/blob/main/plugins.json) - Plugin registry
