# x1ctl

CLI for Bambu Lab X1 printers over LAN. No cloud required.

## Usage

```sh
# Check printer status
x1ctl status --ip 192.168.1.x --access-code ABC123

# Send a command
x1ctl echo --ip 192.168.1.x --access-code ABC123 '{"test": true}'

# Read values
x1ctl read --ip 192.168.1.x --access-code ABC123 nozzle_temp
```

## Self-Update

```sh
x1ctl update              # update to latest
x1ctl version --check     # check for updates
```

## Setup

1. Enable LAN Mode on your Bambu Lab X1 printer
2. Note the IP address and access code from the printer's network settings
3. Download from [releases](https://github.com/joeblew999/plugs/releases/latest)
