# AGENTS.md

This document provides guidance for AI agents working with the `go-security-tools` repository.

## Project Overview

`go-security-tools` (gst) is a Go-based CLI tool providing security-related utilities for OAuth 2.0, OpenID Connect, and JWK (JSON Web Key) management. The project follows idiomatic Go patterns with a clean separation of concerns and interface-based dependency injection.

## Technology Stack

- **Language**: Go 1.25+
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) with [Viper](https://github.com/spf13/viper) for configuration
- **JWK Library**: [lestrrat-go/jwx/v2](https://github.com/lestrrat-go/jwx)
- **UUID Libraries**: `google/uuid`, `lithammer/shortuuid`

## Repository Structure

```
go-security-tools/
├── cmd/                          # Application entry points
│   ├── gst/                      # Main CLI application
│   │   ├── main.go               # Entry point
│   │   └── cmd/                  # Cobra command definitions
│   │       ├── root.go           # Root command setup
│   │       ├── di.go             # Dependency injection/factory functions
│   │       ├── constants.go      # CLI flags and constants
│   │       ├── jwk.go            # JWK parent command
│   │       ├── jwk_create.go     # JWK creation command
│   │       ├── client_credentials.go       # Client credentials parent command
│   │       ├── client_credentials_create.go # Client credentials creation command
│   │       └── version.go        # Version command
│   └── create-docs/              # Documentation generator
│       └── main.go
├── clientcredentials/            # OAuth 2.0 / OIDC client credentials package
│   ├── client_credentials.go     # ClientCredentials struct
│   ├── client_credentials_creator.go  # Creator orchestrator
│   ├── client_id_creator.go      # ClientIDCreator interface
│   ├── client_secret_creator.go  # ClientSecretCreator interface
│   ├── uuidv7_client_id_creator.go    # UUIDv7 implementation
│   ├── short_uuid_client_id_creator.go # ShortUUID implementation
│   ├── crypto_rand_client_secret_creator.go # Crypto-secure secret generator
│   └── crypto_source.go          # Cryptographic random source
├── jwk/                          # Public JWK creation package
│   ├── jwk_creator.go            # JWKCreator interface
│   ├── rsa_jwk_creator.go        # RSA key implementation
│   └── ecdsa_jwk_creator.go      # ECDSA key implementation
├── internal/                     # Internal packages (not exported)
│   └── jwk/
│       ├── jwk_output.go         # JWKOutput data structure
│       ├── jwk_output_writer.go  # JWKOutputWriter interface
│       ├── file_jwk_output_writer.go  # File output implementation
│       └── fmt_jwk_output_writer.go   # Console output implementation
└── docs/                         # Auto-generated CLI documentation
```

## Architecture Patterns

### Interface-Based Design

The codebase uses interfaces to enable flexibility and testability:

- **`ClientIDCreator`**: Interface for generating client IDs
  - Implementations: `UUIDv7ClientIDCreator`, `ShortUUIDClientIDCreator`
- **`ClientSecretCreator`**: Interface for generating client secrets
  - Implementations: `CryptoRandClientSecretCreator`
- **`JWKCreator`**: Interface for creating JSON Web Keys
  - Implementations: `RSAJSONWebKeyCreator`, `ECDSAJWKCreator`
- **`JWKOutputWriter`**: Interface for writing JWK output
  - Implementations: `FileJwkOutputWriter`, `FmtJWKOutputWriter`

### Dependency Injection

The `cmd/gst/cmd/di.go` file contains factory functions that wire up implementations based on CLI flags:
- `buildClientCredentialsCreator()` - Creates client credentials creator with appropriate ID/secret generators
- `buildJWKCreator()` - Creates JWK creator based on algorithm selection
- `buildJWKWriter()` - Creates appropriate output writer (file or console)

### Package Visibility

- **Public packages** (`clientcredentials/`, `jwk/`): Can be imported by external projects
- **Internal packages** (`internal/jwk/`): Private implementation details, not importable externally

## CLI Commands

The tool is invoked as `gst` with the following command structure:

```
gst
├── client-credentials
│   └── create        # Create OAuth 2.0 client credentials
├── jwk
│   └── create        # Create JSON Web Keys
└── version           # Display version information
```

### Key CLI Flags

**JWK Creation:**
- `--kty` / `-k`: Key type (`RSA-2048`, `RSA-4096`, `ECDSA-P256`, `ECDSA-P384`, `ECDSA-P521`)
- `--output-base64` / `-b`: Output JWK as Base64
- `--output-pem-keys` / `-p`: Output PEM-encoded keys
- `--output-path` / `-o`: Write to file instead of console
- `--count` / `-c`: Number of keys to create

**Client Credentials Creation:**
- `--client-id-type` / `-t`: ID type (`uuidv7`, `short-uuid`)
- `--client-secret-type` / `-s`: Secret type (`crypto-rand`)
- `--count` / `-c`: Number of credential pairs to create

### Environment Variables

All flags can be set via environment variables with `GST_` prefix (hyphens become underscores):
- `GST_KTY`, `GST_OUTPUT_BASE64`, `GST_CLIENT_ID_TYPE`, etc.

## Development Guidelines

### Running Tests

```bash
go test ./...
```

### Building the CLI

```bash
go build -o gst ./cmd/gst
```

### Generating Documentation

```bash
./create-docs.sh
# or
go run ./cmd/create-docs
```

### Adding New Features

1. **New Client ID Type**: Implement `ClientIDCreator` interface in `clientcredentials/`
2. **New Secret Type**: Implement `ClientSecretCreator` interface in `clientcredentials/`
3. **New JWK Algorithm**: Implement `JWKCreator` interface in `jwk/`
4. **New Output Format**: Implement `JWKOutputWriter` interface in `internal/jwk/`

Then update `cmd/gst/cmd/di.go` factory functions and `constants.go` with new options.

### Code Style

- Follow standard Go conventions and `gofmt`
- Use descriptive names for interfaces (suffix with `-er` pattern: `Creator`, `Writer`)
- Keep packages focused on single responsibility
- Write unit tests alongside implementations (`*_test.go` files)

## File Naming Conventions

- Interface definitions: `<concept>.go` (e.g., `client_id_creator.go`)
- Implementations: `<implementation>_<concept>.go` (e.g., `uuidv7_client_id_creator.go`)
- Tests: `<source_file>_test.go`

## Key Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/spf13/cobra` | CLI command framework |
| `github.com/spf13/viper` | Configuration management |
| `github.com/lestrrat-go/jwx/v2` | JWK creation and manipulation |
| `github.com/google/uuid` | UUID generation |
| `github.com/lithammer/shortuuid/v4` | Short UUID generation |

## Notes for AI Agents

1. **When modifying CLI commands**: Update both the command file and `constants.go` if adding new flags
2. **When adding implementations**: Follow the existing interface pattern and add corresponding factory logic in `di.go`
3. **When writing tests**: Use table-driven tests where appropriate, following existing test patterns
4. **Documentation**: The `docs/` folder is auto-generated - don't edit manually; run `create-docs.sh` instead
5. **Internal packages**: Respect Go's `internal/` package convention - these are not public API
