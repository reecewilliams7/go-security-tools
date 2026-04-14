# Requirements Specification

## go-security-tools (gst)

**Version**: 1.0  
**Last Updated**: January 2026

---

## 1. Overview

### 1.1 Purpose

The `go-security-tools` (gst) CLI application provides security-related utilities for generating cryptographic keys and OAuth 2.0 client credentials. The tool is designed for developers, security engineers, and DevOps teams who need to create secure authentication and authorization artifacts.

### 1.2 Scope

This document defines the functional and non-functional requirements for the gst CLI tool, covering:
- JSON Web Key (JWK) generation
- OAuth 2.0 Client Credentials generation
- Output formatting and file handling
- CLI interface and configuration

---

## 2. Functional Requirements

### 2.1 JSON Web Key (JWK) Creation

#### 2.1.1 RSA Key Generation

| ID | Requirement | Priority |
|----|-------------|----------|
| JWK-RSA-001 | The system SHALL generate RSA-2048 key pairs | Must Have |
| JWK-RSA-002 | The system SHALL generate RSA-4096 key pairs | Must Have |
| JWK-RSA-003 | RSA keys SHALL be generated using Go's `crypto/rand` for cryptographic security | Must Have |
| JWK-RSA-004 | Each RSA key SHALL be assigned a unique Key ID (kid) | Must Have |

#### 2.1.2 ECDSA Key Generation

| ID | Requirement | Priority |
|----|-------------|----------|
| JWK-EC-001 | The system SHALL generate ECDSA P-256 (secp256r1) key pairs | Must Have |
| JWK-EC-002 | The system SHALL generate ECDSA P-384 (secp384r1) key pairs | Must Have |
| JWK-EC-003 | The system SHALL generate ECDSA P-521 (secp521r1) key pairs | Must Have |
| JWK-EC-004 | ECDSA keys SHALL be generated using Go's `crypto/rand` for cryptographic security | Must Have |
| JWK-EC-005 | Each ECDSA key SHALL be assigned a unique Key ID (kid) | Must Have |

#### 2.1.3 JWK Output Formats

| ID | Requirement | Priority |
|----|-------------|----------|
| JWK-OUT-001 | The system SHALL output JWK in JSON format (RFC 7517 compliant) | Must Have |
| JWK-OUT-002 | The system SHALL output both private and public JWK representations | Must Have |
| JWK-OUT-003 | The system SHALL support Base64 encoding of JWK output | Should Have |
| JWK-OUT-004 | The system SHALL support PEM format export for public keys | Should Have |
| JWK-OUT-005 | The system SHALL support PEM format export for private keys | Should Have |
| JWK-OUT-006 | JSON output SHALL be pretty-printed with proper indentation | Should Have |

#### 2.1.4 JWK Batch Operations

| ID | Requirement | Priority |
|----|-------------|----------|
| JWK-BATCH-001 | The system SHALL support generating multiple JWKs in a single operation | Should Have |
| JWK-BATCH-002 | Each generated JWK SHALL have a unique Key ID (kid) | Must Have |
| JWK-BATCH-003 | The count of keys to generate SHALL be configurable via CLI flag | Should Have |

#### 2.1.5 JWK File Output

| ID | Requirement | Priority |
|----|-------------|----------|
| JWK-FILE-001 | The system SHALL support writing JWK output to files | Should Have |
| JWK-FILE-002 | Private key files SHALL be created with restricted permissions (0600) | Must Have |
| JWK-FILE-003 | Public key files SHALL be created with standard permissions (0644) | Must Have |
| JWK-FILE-004 | The output directory path SHALL be configurable | Should Have |
| JWK-FILE-005 | The output filename base SHALL be configurable | Should Have |
| JWK-FILE-006 | When batch generating, files SHALL be numbered sequentially (e.g., key-1.jwk, key-2.jwk) | Should Have |
| JWK-FILE-007 | The system SHALL validate that the output directory exists before writing | Must Have |

---

### 2.2 Client Credentials Generation

#### 2.2.1 Client ID Generation

| ID | Requirement | Priority |
|----|-------------|----------|
| CC-ID-001 | The system SHALL generate client IDs using UUIDv7 format | Must Have |
| CC-ID-002 | The system SHALL generate client IDs using ShortUUID format | Should Have |
| CC-ID-003 | UUIDv7 client IDs SHALL be time-sortable per RFC 9562 | Must Have |
| CC-ID-004 | ShortUUID client IDs SHALL be URL-safe | Must Have |
| CC-ID-005 | The client ID type SHALL be selectable via CLI flag | Must Have |

#### 2.2.2 Client Secret Generation

| ID | Requirement | Priority |
|----|-------------|----------|
| CC-SEC-001 | The system SHALL generate client secrets using cryptographically secure random bytes | Must Have |
| CC-SEC-002 | Client secrets SHALL be at least 256 bits (32 bytes) of entropy | Must Have |
| CC-SEC-003 | Client secrets SHALL be Base64 encoded for safe storage and transmission | Must Have |
| CC-SEC-004 | The system SHALL use `crypto/rand` as the entropy source | Must Have |

#### 2.2.3 Client Credentials Batch Operations

| ID | Requirement | Priority |
|----|-------------|----------|
| CC-BATCH-001 | The system SHALL support generating multiple credential pairs in a single operation | Should Have |
| CC-BATCH-002 | The count of credential pairs to generate SHALL be configurable via CLI flag | Should Have |

#### 2.2.4 Client Credentials Output

| ID | Requirement | Priority |
|----|-------------|----------|
| CC-OUT-001 | The system SHALL output Client ID and Client Secret to the console | Must Have |
| CC-OUT-002 | Output SHALL clearly separate each credential pair | Should Have |
| CC-OUT-003 | Output SHALL be formatted for easy copy/paste | Should Have |

---

### 2.3 CLI Interface

#### 2.3.1 Command Structure

| ID | Requirement | Priority |
|----|-------------|----------|
| CLI-CMD-001 | The CLI SHALL be invoked using the `gst` command | Must Have |
| CLI-CMD-002 | The CLI SHALL support a `jwk` parent command for JWK operations | Must Have |
| CLI-CMD-003 | The CLI SHALL support a `jwk create` subcommand for JWK generation | Must Have |
| CLI-CMD-004 | The CLI SHALL support a `client-credentials` parent command | Must Have |
| CLI-CMD-005 | The CLI SHALL support a `client-credentials create` subcommand | Must Have |
| CLI-CMD-006 | The CLI SHALL support a `version` command | Should Have |
| CLI-CMD-007 | Each command SHALL display help information with `-h` or `--help` | Must Have |

#### 2.3.2 JWK Create Flags

| ID | Requirement | Flag | Short | Default | Priority |
|----|-------------|------|-------|---------|----------|
| CLI-JWK-001 | Key type selection | `--kty` | `-k` | `RSA-2048` | Must Have |
| CLI-JWK-002 | Base64 output toggle | `--output-base64` | `-b` | `false` | Should Have |
| CLI-JWK-003 | PEM output toggle | `--output-pem-keys` | `-p` | `false` | Should Have |
| CLI-JWK-004 | Output directory path | `--output-path` | `-o` | (empty) | Should Have |
| CLI-JWK-005 | Output filename base | `--output-file` | `-f` | `create-jwk` | Should Have |
| CLI-JWK-006 | Count of keys to generate | `--count` | `-c` | `1` | Should Have |

#### 2.3.3 Client Credentials Create Flags

| ID | Requirement | Flag | Short | Default | Priority |
|----|-------------|------|-------|---------|----------|
| CLI-CC-001 | Client ID type selection | `--client-id-type` | `-t` | `uuidv7` | Must Have |
| CLI-CC-002 | Client secret type selection | `--client-secret-type` | `-s` | `crypto-rand` | Must Have |
| CLI-CC-003 | Count of pairs to generate | `--count` | `-c` | `1` | Should Have |

#### 2.3.4 Environment Variable Configuration

| ID | Requirement | Priority |
|----|-------------|----------|
| CLI-ENV-001 | All CLI flags SHALL be configurable via environment variables | Should Have |
| CLI-ENV-002 | Environment variables SHALL use the `GST_` prefix | Must Have |
| CLI-ENV-003 | Hyphens in flag names SHALL be converted to underscores in env var names | Must Have |
| CLI-ENV-004 | CLI flags SHALL take precedence over environment variables | Must Have |

---

## 3. Non-Functional Requirements

### 3.1 Security

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-SEC-001 | All cryptographic operations SHALL use cryptographically secure random number generation | Must Have |
| NFR-SEC-002 | Private key files SHALL be created with restrictive permissions (0600) | Must Have |
| NFR-SEC-003 | The system SHALL NOT log or persist sensitive key material beyond intended output | Must Have |
| NFR-SEC-004 | Minimum key sizes SHALL meet current industry standards (RSA-2048, P-256) | Must Have |

### 3.2 Performance

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-PERF-001 | Single key generation SHALL complete within 5 seconds | Should Have |
| NFR-PERF-002 | Batch operations SHALL process keys sequentially to avoid memory issues | Should Have |

### 3.3 Reliability

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-REL-001 | The system SHALL return appropriate error messages for invalid inputs | Must Have |
| NFR-REL-002 | The system SHALL validate output directory existence before file operations | Must Have |
| NFR-REL-003 | The system SHALL exit with non-zero status code on errors | Must Have |

### 3.4 Maintainability

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-MAIN-001 | The codebase SHALL follow standard Go conventions and formatting | Should Have |
| NFR-MAIN-002 | Public packages SHALL be importable by external projects | Should Have |
| NFR-MAIN-003 | Internal implementation details SHALL use Go's `internal/` package convention | Should Have |
| NFR-MAIN-004 | The system SHALL use interface-based design for extensibility | Should Have |

### 3.5 Documentation

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-DOC-001 | CLI commands SHALL have auto-generated markdown documentation | Should Have |
| NFR-DOC-002 | Each command SHALL have short and long descriptions | Should Have |
| NFR-DOC-003 | All public interfaces SHALL have GoDoc comments | Should Have |

### 3.6 Portability

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-PORT-001 | The tool SHALL be compilable on Linux, macOS, and Windows | Should Have |
| NFR-PORT-002 | The tool SHALL have no platform-specific dependencies | Should Have |

---

## 4. Technical Constraints

### 4.1 Technology Stack

| Constraint | Description |
|------------|-------------|
| TC-001 | Implementation language SHALL be Go 1.25 or higher |
| TC-002 | CLI framework SHALL be Cobra with Viper for configuration |
| TC-003 | JWK library SHALL be lestrrat-go/jwx/v2 |
| TC-004 | UUID generation SHALL use google/uuid for UUIDv7 |
| TC-005 | ShortUUID generation SHALL use lithammer/shortuuid/v4 |

### 4.2 Standards Compliance

| Standard | Description |
|----------|-------------|
| RFC 7517 | JSON Web Key (JWK) format compliance |
| RFC 7518 | JSON Web Algorithms (JWA) for RSA and ECDSA |
| RFC 6749 | OAuth 2.0 client credentials format |
| RFC 9562 | UUIDv7 time-sortable identifier format |

---

## 5. Interface Definitions

### 5.1 JWKCreator Interface

```go
type JWKCreator interface {
    Create() (*JWKOutput, error)
}
```

**Implementations Required:**
- `RSAJSONWebKeyCreator` - RSA key generation (2048, 4096 bit)
- `ECDSAJWKCreator` - ECDSA key generation (P-256, P-384, P-521)

### 5.2 ClientIDCreator Interface

```go
type ClientIDCreator interface {
    Create() (string, error)
}
```

**Implementations Required:**
- `UUIDv7ClientIDCreator` - UUIDv7 format client IDs
- `ShortUUIDClientIDCreator` - ShortUUID format client IDs

### 5.3 ClientSecretCreator Interface

```go
type ClientSecretCreator interface {
    Create() (string, error)
}
```

**Implementations Required:**
- `CryptoRandClientSecretCreator` - Cryptographically secure secrets

### 5.4 JWKOutputWriter Interface

```go
type JWKOutputWriter interface {
    Write(output *JWKOutput, i int) error
}
```

**Implementations Required:**
- `FileJwkOutputWriter` - Write JWK output to files
- `FmtJWKOutputWriter` - Write JWK output to console

---

## 6. Data Structures

### 6.1 JWKOutput

| Field | Type | Description |
|-------|------|-------------|
| JWK | jwk.Key | The full JWK including private key material |
| JWKPublic | jwk.Key | The public key portion only |
| JWKString | string | JSON-formatted JWK private key |
| JWKPublicString | string | JSON-formatted JWK public key |
| Base64JWK | string | Base64-encoded JWK |
| PEMPublicKey | string | PEM-encoded public key |
| PEMPrivateKey | string | PEM-encoded private key |

### 6.2 ClientCredentials

| Field | Type | Description |
|-------|------|-------------|
| clientID | string | The OAuth 2.0 client identifier |
| clientSecret | string | The OAuth 2.0 client secret |

---

## 7. Acceptance Criteria

### 7.1 JWK Generation

- [ ] RSA-2048 keys can be generated and output in JWK format
- [ ] RSA-4096 keys can be generated and output in JWK format
- [ ] ECDSA P-256 keys can be generated and output in JWK format
- [ ] ECDSA P-384 keys can be generated and output in JWK format
- [ ] ECDSA P-521 keys can be generated and output in JWK format
- [ ] Generated JWKs contain unique Key IDs (kid)
- [ ] Base64 encoding option works correctly
- [ ] PEM export option works correctly
- [ ] File output creates files with correct permissions
- [ ] Batch generation creates multiple unique keys

### 7.2 Client Credentials Generation

- [ ] UUIDv7 client IDs are generated correctly
- [ ] ShortUUID client IDs are generated correctly
- [ ] Client secrets are at least 32 bytes of entropy
- [ ] Client secrets are Base64 encoded
- [ ] Batch generation creates multiple unique credential pairs

### 7.3 CLI Interface

- [ ] All commands display help with `-h` flag
- [ ] All flags work with both short and long forms
- [ ] Environment variables override defaults
- [ ] CLI flags override environment variables
- [ ] Version command displays version information
- [ ] Invalid inputs produce helpful error messages

---

## 8. Glossary

| Term | Definition |
|------|------------|
| **JWK** | JSON Web Key - A JSON data structure representing a cryptographic key |
| **ECDSA** | Elliptic Curve Digital Signature Algorithm |
| **RSA** | Rivest–Shamir–Adleman asymmetric cryptographic algorithm |
| **PEM** | Privacy-Enhanced Mail - Base64 encoded format for cryptographic keys |
| **OAuth 2.0** | Authorization framework for secure delegated access |
| **OIDC** | OpenID Connect - Authentication layer built on OAuth 2.0 |
| **UUID** | Universally Unique Identifier |
| **UUIDv7** | Time-sortable UUID format defined in RFC 9562 |
| **kid** | Key ID - Unique identifier for a JWK |
