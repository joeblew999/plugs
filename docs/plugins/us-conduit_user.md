# us-conduit User Guide

Conduit is a data integration tool for building real-time data pipelines.

## Quick Start

```sh
# Install via plugctl
plugctl install us-conduit

# Start Conduit
us-conduit
```

## Features

- Connect to databases, message queues, APIs
- Transform data in-flight
- Real-time streaming
- Web UI for pipeline management

## Configuration

Create a `pipeline.yaml`:

```yaml
version: 2.2
pipelines:
  - id: my-pipeline
    connectors:
      - id: source
        type: source
        plugin: builtin:file
        settings:
          path: ./input.json
      - id: destination
        type: destination
        plugin: builtin:file
        settings:
          path: ./output.json
```

Run with:

```sh
us-conduit --config pipeline.yaml
```

## Self-Update

```sh
plugctl update us-conduit
```

## More Info

- [Conduit.io](https://conduit.io) - Official documentation
- [GitHub](https://github.com/ConduitIO/conduit) - Source code
