# fakeprinter

Mock printer server for testing and development.

## Usage

```sh
# Start mock server on default port
fakeprinter -addr :8883
```

Connect any LAN client to `wss://localhost:8883` (self-signed TLS, any access code).

## Self-Update

```sh
fakeprinter --update         # update to latest
fakeprinter --check-update   # check for updates
fakeprinter --version        # show version
```

## Notes

- Uses self-signed TLS certificate
- Access code is ignored (for testing only)
- Sends hello JSON on connect, then echoes your JSON
