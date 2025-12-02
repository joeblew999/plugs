# us-task User Guide

Task is a task runner / build tool that aims to be simpler and easier to use than GNU Make.

## Quick Start

```sh
# Install via plugctl
plugctl install us-task

# List tasks in current directory
us-task --list

# Run default task
us-task

# Run specific task
us-task build
```

## Taskfile.yml

Create a `Taskfile.yml` in your project:

```yaml
version: "3"

tasks:
  default:
    cmds:
      - echo "Hello, World!"

  build:
    desc: Build the project
    cmds:
      - go build ./...

  test:
    desc: Run tests
    cmds:
      - go test ./...
```

## Common Options

| Flag | Description |
|------|-------------|
| `--list` | List available tasks |
| `--dry` | Show commands without running |
| `--force` | Force run even if up-to-date |
| `--parallel` | Run tasks in parallel |

## Self-Update

```sh
us-task --update    # update to latest version
```

Or use plugctl:

```sh
plugctl update us-task
```

## More Info

- [Taskfile.dev](https://taskfile.dev) - Official documentation
- [GitHub](https://github.com/go-task/task) - Source code
