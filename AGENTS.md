# AGENTS.md - AI Coding Agent Instructions

This document provides instructions for AI coding agents working on the confluence-cli project.

## Project Overview

confluence-cli is a Go CLI tool for interacting with the Confluence Cloud API. It uses the Atlassian Confluence Cloud REST API v2 (OpenAPI spec in `swagger.json`).

- **CLI Framework**: Cobra (github.com/spf13/cobra)
- **Configuration**: Viper (github.com/spf13/viper)
- **API Client**: Auto-generated Swagger client for Confluence Cloud REST API v2

## Build System

This project uses [Task](https://taskfile.dev/) as the build runner (alternative to Make).

### Available Commands

| Command | Description |
|---------|-------------|
| `task build` | Build the CLI binary |
| `task install` | Install the CLI binary to $GOPATH/bin |
| `task test` | Run all tests |
| `task coverage` | Run tests with coverage summary |
| `task coverage-html` | Generate HTML coverage report |
| `task open-coverage-html` | Generate and open HTML coverage report in browser |
| `task pull-swagger` | Download latest Confluence API OpenAPI spec |
| `task generate-swagger` | Generate Go client from OpenAPI spec |

### Running a Single Test

```bash
# Run a specific test function
go test -v -run TestFunctionName ./path/to/package

# Run tests in a specific package
go test -v ./cmd/...

# Run tests matching a pattern
go test -v -run "TestCreate.*" ./...

# Run a single test with verbose output
go test -v -run "^TestExactName$" ./internal/client
```

## Code Style Guidelines

### Go Version

Use the Go version specified in `go.mod`. Follow the latest stable Go conventions.

### Formatting

- Use `gofmt` or `goimports` for formatting (enforced automatically)
- Line length: No hard limit, but prefer readable line lengths (~100-120 chars)
- Use tabs for indentation (Go standard)

### Imports

Organize imports in three groups, separated by blank lines:
1. Standard library packages
2. Third-party packages
3. Local/internal packages

```go
import (
    "context"
    "fmt"
    "net/http"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    "github.com/alexhokl/confluence-cli/internal/client"
    "github.com/alexhokl/confluence-cli/internal/config"
)
```

### Naming Conventions

- **Packages**: Short, lowercase, single-word names (e.g., `client`, `config`, `cmd`)
- **Files**: Lowercase with underscores (e.g., `page_service.go`, `api_client.go`)
- **Exported identifiers**: PascalCase (e.g., `CreatePage`, `ClientConfig`)
- **Unexported identifiers**: camelCase (e.g., `parseResponse`, `httpClient`)
- **Constants**: PascalCase for exported, camelCase for unexported
- **Interfaces**: Name after behavior, often ending in `-er` (e.g., `Reader`, `PageCreator`)
- **Test files**: `*_test.go` (e.g., `client_test.go`)

### Error Handling

- Always check errors explicitly; never ignore them
- Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Return errors to callers; let main/cmd layer handle user-facing messages
- Use sentinel errors for expected error conditions (e.g., `var ErrPageNotFound = errors.New("page not found")`)

### Context Usage

- Always pass `context.Context` as the first parameter
- Use context for cancellation, timeouts, and request-scoped values
- Never store context in structs

### CLI Structure (Cobra)

This project uses [Cobra](https://github.com/spf13/cobra) for CLI commands. Place commands in `cmd/` directory (one file per command group) and use `internal/` for non-exported packages.

### Testing

- Write table-driven tests when testing multiple cases
- Use `t.Helper()` in test helper functions
- Use `testify/assert` or `testify/require` for assertions (if added as dependency)
- Mock external dependencies for unit tests

### Documentation

- Write doc comments for all exported functions, types, and packages
- Start comments with the identifier name
- Keep comments concise but informative

### Dependencies

- Minimize external dependencies
- Prefer standard library when reasonable
- Pin dependency versions in `go.mod`

### Configuration

- Use environment variables for sensitive data (API tokens, credentials)
- Support configuration files via Viper (if applicable)
- Never commit `.env` files or credentials

## Project-Specific Notes

### API Client Generation

The project can generate a Go API client from the OpenAPI spec:

```bash
task generate-swagger
```

This generates client code in the `swagger/` directory using openapi-generator.

### Authentication

Confluence Cloud API uses API tokens. Store credentials securely:
- Use environment variables: `CONFLUENCE_API_TOKEN`, `CONFLUENCE_USER`, `CONFLUENCE_URL`
- Support config file in `~/.config/confluence-cli/config.yaml`

## Common Tasks for Agents

1. **Adding a new command**: Create file in `cmd/`, register with root command
2. **Adding API functionality**: Implement in `internal/client/`
3. **Writing tests**: Create `*_test.go` alongside source files
4. **Updating dependencies**: Run `go mod tidy` after changes
