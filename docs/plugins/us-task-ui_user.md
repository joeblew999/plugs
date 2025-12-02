# us-task-ui User Guide

A web-based graphical interface for running Taskfile tasks.

## Quick Start

```sh
# Install via plugctl
plugctl install us-task-ui

# Start the web server
us-task-ui

# Open browser to http://localhost:8080
```

## Features

- Visual task list from your Taskfile.yml
- Click to run tasks
- See output in real-time
- No terminal required

## Options

| Flag | Description |
|------|-------------|
| `--port` | HTTP port (default: 8080) |
| `--dir` | Working directory with Taskfile.yml |

## Self-Update

```sh
plugctl update us-task-ui
```

## More Info

- [GitHub](https://github.com/titpetric/task-ui) - Source code
