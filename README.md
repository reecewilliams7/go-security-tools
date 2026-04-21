# go-security-tools

![on_commit](https://github.com/reecewilliams7/go-security-tools/actions/workflows/on_commit.yml/badge.svg?branch=main)

A collection of security utilities for OAuth 2.0, OpenID Connect, and JWK management. The project ships three apps built from the same Go library packages.

## Apps

### `gst` — CLI

A command-line tool for generating security artifacts directly in the terminal.

| Command | Description |
|---|---|
| `gst jwk create` | Generate JSON Web Keys (RSA-2048, RSA-4096, ECDSA-P256/P384/P521) |
| `gst client-credentials create` | Generate OAuth 2.0 client ID and secret pairs |
| `gst version` | Print the current version |

Key flags for `jwk create`: `--kty`, `--count`, `--output-base64`, `--output-pem-keys`, `--output-path`.  
Key flags for `client-credentials create`: `--client-id-type` (`uuidv7`, `short-uuid`), `--client-secret-type` (`crypto-rand`), `--count`.

All flags can also be set via environment variables with a `GST_` prefix (e.g. `GST_KTY=RSA-4096`).

→ [Full CLI reference](docs/gst.md)

### `gst-tui` — Terminal UI

An interactive terminal UI (built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)) providing the same JWK and client credentials functionality through a guided menu-driven interface.

```
gst-tui
```

### `gst-web` — Web UI

A lightweight web application (plain Go `net/http` + HTML templates, no JavaScript framework) running on port `8080`. Provides the same JWK and client credentials generation through a browser-based form interface.

```
gst-web          # starts server on http://localhost:8080
```

## Installation

Download the latest release archive for your platform from the [Releases](../../releases) page. Each app is released as a separate archive:

- `gst_<OS>_<arch>.tar.gz`
- `gst-tui_<OS>_<arch>.tar.gz`
- `gst-web_<OS>_<arch>.tar.gz`

## Building from source

```bash
go build ./cmd/gst        # CLI
go build ./cmd/gst-tui    # Terminal UI
go build ./cmd/gst-web    # Web UI
```

Requires Go 1.25+.

## Testing

```bash
go test ./...
```