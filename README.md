# go-security-tools

![on_commit](https://github.com/reecewilliams7/go-security-tools/actions/workflows/on_commit.yml/badge.svg?branch=main)

A collection of security utilities for OAuth 2.0, OpenID Connect, and JWK management. The project ships three apps built from the same Go library packages.

## Apps

### `gst` — CLI

A command-line tool for generating security artifacts directly in the terminal.

| Command | Description |
|---|---|
| `gst jwk create` | Generate JSON Web Keys |
| `gst client-credentials create` | Generate OAuth 2.0 client ID and secret pairs |
| `gst pkce create` | Generate PKCE code verifier and challenge pairs (RFC 7636) |
| `gst jwt decode` | Decode and inspect a JWT (no signature verification) |
| `gst version` | Print the current version |

**`jwk create` flags**

| Flag | Short | Description | Default |
|---|---|---|---|
| `--kty` | `-k` | Key type: `RSA-2048`, `RSA-4096`, `ECDSA-P256`, `ECDSA-P384`, `ECDSA-P521`, `OKP-Ed25519`, `OKP-X25519`, `HS256`, `HS384`, `HS512` | `RSA-2048` |
| `--count` | `-c` | Number of keys to generate | `1` |
| `--output-base64` | `-b` | Also output the private key as Base64 | `false` |
| `--output-pem-keys` | `-p` | Also output PEM-encoded keys (asymmetric only) | `false` |
| `--output-jwks` | `-j` | Also output all public keys as a JWKS JSON set | `false` |
| `--output-path` | `-o` | Directory to write key files (omit for console output) | |

**`client-credentials create` flags**

| Flag | Short | Description | Default |
|---|---|---|---|
| `--client-id-type` | `-t` | ID format: `uuidv7`, `short-uuid`, `nanoid` | `uuidv7` |
| `--client-secret-type` | `-s` | Secret generator: `crypto-rand` | `crypto-rand` |
| `--secret-length` | `-l` | Secret length in bytes (16–64) | `32` |
| `--secret-encoding` | `-e` | Secret encoding: `base64`, `base64url`, `hex` | `base64` |
| `--count` | `-c` | Number of credential pairs to generate | `1` |

**`pkce create` flags**

| Flag | Short | Description | Default |
|---|---|---|---|
| `--method` | `-m` | Challenge method (currently only `S256`) | `S256` |
| `--count` | `-c` | Number of pairs to generate | `1` |

**`jwt decode` flags**

| Flag | Short | Description |
|---|---|---|
| `--token` | `-t` | JWT string to decode (reads from stdin if omitted) |

All flags can also be set via environment variables with a `GST_` prefix (e.g. `GST_KTY=RSA-4096`).

→ [Full CLI reference](docs/gst.md)

### `gst-tui` — Terminal UI

An interactive terminal UI (built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)) providing all features through a guided menu-driven interface.

```
gst-tui
```

Menu options:
- **Create JWK** — all key types (RSA, ECDSA, OKP, HMAC), Base64/PEM/JWKS toggles, optional file output
- **Create Client Credentials** — uuidv7 / short-uuid / nanoid IDs, configurable secret length and encoding
- **Create PKCE** — generates S256 code verifier and challenge pairs
- **Decode JWT** — paste a token to inspect header, payload, and expiry status

### `gst-web` — Web UI

A lightweight web application (plain Go `net/http` + HTML templates, no JavaScript framework) running on port `8080`.

```
gst-web          # starts server on http://localhost:8080
```

Pages:
- `/jwk` — JWK creator (all key types, Base64/PEM/JWKS output options)
- `/client-credentials` — client credentials generator (nanoid support, secret length/encoding)
- `/pkce` — PKCE pair generator
- `/jwt` — JWT decoder

## Library packages

The three apps are thin wrappers around importable Go packages:

| Package | Description |
|---|---|
| `jwk` | `JWKCreator` interface with RSA, ECDSA, OKP (Ed25519/X25519), and HMAC implementations |
| `clientcredentials` | `ClientIDCreator` and `ClientSecretCreator` interfaces; UUIDv7, ShortUUID, Nanoid ID creators; configurable crypto-rand secret creator |
| `pkce` | `PKCECreator` interface; S256 implementation |
| `internal/jwt` | JWT decode (header + payload parsing, expiry check) |

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