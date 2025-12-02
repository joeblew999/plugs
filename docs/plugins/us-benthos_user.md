# us-benthos User Guide

Benthos is a high-performance stream processor for data pipelines.

## Quick Start

```sh
# Install via plugctl
plugctl install us-benthos

# Run with a config
us-benthos -c pipeline.yaml
```

## Features

- Connect to Kafka, NATS, Redis, HTTP, and more
- Transform with Bloblang
- Serverless-ready
- Metrics and tracing built-in

## Configuration

Create a `pipeline.yaml`:

```yaml
input:
  stdin: {}

pipeline:
  processors:
    - mapping: |
        root = this.uppercase()

output:
  stdout: {}
```

Run:

```sh
echo "hello world" | us-benthos -c pipeline.yaml
# Output: HELLO WORLD
```

## Self-Update

```sh
plugctl update us-benthos
```

## More Info

- [Benthos.dev](https://www.benthos.dev) - Official documentation
- [GitHub](https://github.com/redpanda-data/benthos) - Source code
