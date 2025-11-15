# Contributing to DLG

Thank you for your interest in contributing to DLG! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Documentation](#documentation)
- [Pull Request Process](#pull-request-process)
- [Adding New Protocols](#adding-new-protocols)
- [Style Guide](#style-guide)

## Code of Conduct

### Our Pledge

We are committed to providing a welcoming and inclusive environment for all contributors. We pledge to:

- Be respectful and considerate
- Welcome diverse perspectives
- Accept constructive criticism gracefully
- Focus on what's best for the community
- Show empathy towards others

### Our Standards

Examples of behavior that contributes to a positive environment:

- Using welcoming and inclusive language
- Being respectful of differing viewpoints
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

Examples of unacceptable behavior:

- Trolling, insulting/derogatory comments, and personal attacks
- Public or private harassment
- Publishing others' private information without permission
- Other conduct which could reasonably be considered inappropriate

## Getting Started

### Prerequisites

- Go 1.24.0 or later
- Git
- Make
- Docker (optional, for integration tests)

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork:

```bash
git clone https://github.com/YOUR_USERNAME/dlg.git
cd dlg
```

3. Add upstream remote:

```bash
git remote add upstream https://github.com/hodgesds/dlg.git
```

## Development Setup

### Install Dependencies

```bash
go mod download
```

### Build

```bash
make build
```

The binary will be available at `./dlg`.

### Run Tests

```bash
make test
```

### Run Linters

```bash
make lint
```

## Making Changes

### Branching Strategy

1. Create a feature branch from `main`:

```bash
git checkout -b feature/my-feature
```

2. Make your changes
3. Commit with clear messages
4. Push to your fork
5. Open a pull request

### Commit Messages

Write clear, concise commit messages:

```
Add support for PostgreSQL prepared statements

- Implement prepared statement caching
- Add configuration options
- Update documentation
- Add tests
```

Format:
- First line: Brief summary (50 chars or less)
- Blank line
- Detailed description (if needed)
- List of changes

### Branch Naming

Use descriptive branch names:

- `feature/add-kafka-support` - New features
- `fix/http-timeout-bug` - Bug fixes
- `docs/update-api-guide` - Documentation
- `refactor/executor-interface` - Refactoring
- `test/add-redis-tests` - Tests

## Testing

### Test Coverage

All code changes should include tests. We aim for high test coverage:

- Unit tests for all new functions
- Integration tests for protocol implementations
- End-to-end tests for complex features

### Writing Tests

Use the standard Go testing package with testify for assertions:

```go
package mypackage

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMyFunction(t *testing.T) {
    result, err := MyFunction("input")
    require.NoError(t, err)
    assert.Equal(t, "expected", result)
}
```

### Test Structure

Organize tests by:

1. **Unit Tests**: Test individual functions in isolation
2. **Integration Tests**: Test interactions with external systems
3. **End-to-End Tests**: Test complete workflows

### Running Specific Tests

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./executor/http

# Run specific test
go test -run TestExecuteSimpleGET ./executor/http

# Run with coverage
go test -cover ./...

# Run with race detection
go test -race ./...
```

### Test Best Practices

- Use table-driven tests for multiple scenarios
- Test error cases
- Test edge cases
- Mock external dependencies
- Use test helpers to reduce duplication
- Clean up resources in tests

Example table-driven test:

```go
func TestValidation(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
    }{
        {"valid URL", "https://example.com", false},
        {"invalid URL", "not-a-url", true},
        {"empty URL", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := Validate(tt.input)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

## Documentation

### Code Documentation

Document all exported types, functions, and constants:

```go
// Config represents the configuration for HTTP load testing.
// It contains settings for connection pooling, request details,
// and execution parameters.
type Config struct {
    // Count is the number of requests to execute
    Count int `yaml:"count"`

    // MaxIdleConns sets the maximum idle connections in the pool
    MaxIdleConns *int `yaml:"maxIdleConns,omitempty"`
}

// Execute runs the HTTP load test according to the provided configuration.
// It returns an error if any requests fail or if the configuration is invalid.
func (e *httpExecutor) Execute(ctx context.Context, conf *Config) error {
    // Implementation
}
```

### User Documentation

Update relevant documentation when making changes:

- **README.md**: Overview and quick start
- **docs/USER_GUIDE.md**: User-facing features
- **docs/API.md**: API changes
- **docs/ARCHITECTURE.md**: Architectural changes

### Example Code

Provide examples for new features:

```go
// Example usage in tests
func ExampleExecutor() {
    executor := http.New(prometheus.NewRegistry())
    config := &http.Config{
        Count: 100,
        Payload: http.Payload{
            URL:    "https://api.example.com",
            Method: "GET",
        },
    }
    err := executor.Execute(context.Background(), config)
    if err != nil {
        log.Fatal(err)
    }
    // Output: Load test completed
}
```

## Pull Request Process

### Before Submitting

1. **Run tests**: Ensure all tests pass
2. **Run linters**: Fix any linting issues
3. **Update documentation**: Add/update relevant docs
4. **Add tests**: Include tests for new functionality
5. **Update CHANGELOG**: Add entry for your changes

### PR Checklist

- [ ] Tests pass locally
- [ ] Code follows style guide
- [ ] Documentation updated
- [ ] CHANGELOG updated
- [ ] Commits are clean and well-described
- [ ] No merge conflicts
- [ ] PR description explains changes

### PR Description Template

```markdown
## Description
Brief description of changes

## Motivation
Why are these changes needed?

## Changes
- Change 1
- Change 2

## Testing
How were these changes tested?

## Screenshots (if applicable)
Add screenshots for UI changes

## Checklist
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] CHANGELOG updated
```

### Review Process

1. Submit PR
2. Maintainers review
3. Address feedback
4. Approval required from at least one maintainer
5. Merge when approved

## Adding New Protocols

### Protocol Implementation Checklist

1. **Configuration** (`/config/{protocol}/`)
   - Create config structure
   - Add YAML tags
   - Document fields
   - Add validation

2. **Executor** (`/executor/{protocol}/`)
   - Implement executor interface
   - Add Prometheus metrics
   - Handle errors properly
   - Support context cancellation

3. **CLI Command** (`/cmd/{protocol}/`)
   - Create Cobra command
   - Add flags
   - Wire up executor

4. **MCP Tools** (`/mcpserver/tools.go`)
   - Add tool definition
   - Implement handler
   - Document tool

5. **Tests**
   - Unit tests for config
   - Unit tests for executor
   - Integration tests (if possible)

6. **Documentation**
   - Add to README
   - Add to USER_GUIDE
   - Add examples

### Protocol Implementation Example

#### 1. Configuration

```go
// config/myprotocol/myprotocol.go
package myprotocol

type Config struct {
    Address string `yaml:"address"`
    Count   int    `yaml:"count"`
    Timeout *time.Duration `yaml:"timeout,omitempty"`
}
```

#### 2. Executor

```go
// executor/myprotocol/myprotocol.go
package myprotocol

import (
    "context"
    "github.com/hodgesds/dlg/config/myprotocol"
    "github.com/hodgesds/dlg/executor"
    "github.com/prometheus/client_golang/prometheus"
)

type myprotocolExecutor struct {
    reg *prometheus.Registry
}

func New(reg *prometheus.Registry) executor.MyProtocol {
    return &myprotocolExecutor{reg: reg}
}

func (e *myprotocolExecutor) Execute(ctx context.Context, conf *myprotocol.Config) error {
    // Implementation
    return nil
}
```

#### 3. CLI Command

```go
// cmd/myprotocol/myprotocol.go
package myprotocol

import (
    "github.com/spf13/cobra"
    "github.com/hodgesds/dlg/config/myprotocol"
)

var Cmd = &cobra.Command{
    Use:   "myprotocol",
    Short: "Run MyProtocol load test",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Parse flags and execute
        return nil
    },
}

func init() {
    Cmd.Flags().String("address", "localhost:1234", "Server address")
    Cmd.Flags().Int("count", 100, "Number of operations")
}
```

#### 4. Tests

```go
// config/myprotocol/myprotocol_test.go
package myprotocol

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
    config := &Config{
        Address: "localhost:1234",
        Count:   100,
    }
    assert.Equal(t, "localhost:1234", config.Address)
    assert.Equal(t, 100, config.Count)
}
```

## Style Guide

### Go Style

Follow the official Go style guidelines:

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

### Key Points

1. **Formatting**: Use `gofmt` or `goimports`
2. **Naming**:
   - Exported: `ExecuteLoadTest`
   - Unexported: `executeLoadTest`
   - Interfaces: `Executor` (not `IExecutor`)
   - Packages: Short, lowercase, no underscores

3. **Error Handling**:
   - Return errors, don't panic
   - Wrap errors with context: `fmt.Errorf("failed to connect: %w", err)`
   - Check all errors

4. **Comments**:
   - Document all exported items
   - Use complete sentences
   - Start with the name being documented

5. **Structure**:
   - Small, focused functions
   - Clear separation of concerns
   - Limit package dependencies

### Example

```go
// Good
package http

// Config represents HTTP load test configuration.
type Config struct {
    URL    string `yaml:"url"`
    Method string `yaml:"method"`
}

// Execute runs the HTTP load test.
func Execute(ctx context.Context, conf *Config) error {
    if conf.URL == "" {
        return fmt.Errorf("URL is required")
    }
    // Implementation
    return nil
}

// Bad
package http

type config struct {  // Should be Config (exported)
    url string        // Should be URL
    method string
}

func execute(conf *config) {  // Should return error
    // Implementation with no error handling
}
```

## Questions?

If you have questions:

1. Check existing documentation
2. Search existing issues
3. Ask in GitHub Discussions
4. Open an issue

## License

By contributing, you agree that your contributions will be licensed under the Apache License 2.0.

## Recognition

Contributors will be recognized in:

- CONTRIBUTORS file
- Release notes
- Project README (for significant contributions)

Thank you for contributing to DLG!
