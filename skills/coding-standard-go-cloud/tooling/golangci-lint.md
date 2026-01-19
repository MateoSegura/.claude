# golangci-lint Configuration

golangci-lint is the standard linter aggregator for Go projects. This configuration enforces the rules defined in this coding standard.

## Installation

```bash
# Install latest version
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Or use Homebrew on macOS
brew install golangci-lint
```

## Required Version

- **Minimum**: golangci-lint v1.55+
- **Go**: 1.21+

## Configuration File

Place this configuration in `.golangci.yml` at the project root:

```yaml
# .golangci.yml
version: "2"

run:
  timeout: 5m

linters:
  default: standard
  enable:
    - staticcheck
    - govet
    - errcheck
    - gosec
    - revive
    - gocyclo
    - gocognit
    - dupl
    - funlen
    - gofmt
    - goimports
    - misspell
    - unconvert
    - unparam
    - prealloc
    - bodyclose

linters-settings:
  gocyclo:
    min-complexity: 15
  gocognit:
    min-complexity: 20
  funlen:
    lines: 100
    statements: 50
  dupl:
    threshold: 100
```

## Linter Descriptions

| Linter | Purpose | Related Rules |
|--------|---------|---------------|
| staticcheck | Advanced static analysis | Multiple |
| govet | Reports suspicious constructs | GO-ERR-001 |
| errcheck | Unchecked error returns | GO-ERR-001 |
| gosec | Security issues | GO-SEC-* |
| revive | Style and best practices | GO-NAM-*, GO-DOC-* |
| gocyclo | Cyclomatic complexity | Maintainability |
| gocognit | Cognitive complexity | Maintainability |
| dupl | Code duplication | DRY principle |
| funlen | Function length | Maintainability |
| gofmt | Formatting | GO-FMT-* |
| goimports | Import ordering | GO-FMT-001 |
| misspell | Spelling mistakes | Documentation |
| unconvert | Unnecessary conversions | Performance |
| unparam | Unused parameters | Clean code |
| prealloc | Slice preallocation | GO-PRF-001 |
| bodyclose | HTTP response body close | Resource cleanup |

## Running the Linter

```bash
# Run all linters
golangci-lint run

# Run on specific directories
golangci-lint run ./internal/...

# Run specific linters only
golangci-lint run --enable=errcheck,gosec

# Fix auto-fixable issues
golangci-lint run --fix

# Output in different formats
golangci-lint run --out-format=json
golangci-lint run --out-format=github-actions  # For CI
```

## CI Integration

### GitHub Actions

```yaml
- name: golangci-lint
  uses: golangci/golangci-lint-action@v3
  with:
    version: v1.55
    args: --timeout=5m
```

### GitLab CI

```yaml
lint:
  image: golangci/golangci-lint:v1.55
  script:
    - golangci-lint run --timeout=5m
```

## Excluding False Positives

```yaml
# In .golangci.yml
issues:
  exclude-rules:
    # Exclude some linters for test files
    - path: _test\.go
      linters:
        - dupl
        - funlen

    # Exclude specific error messages
    - linters:
        - gosec
      text: "G104: Errors unhandled"
      path: "internal/test/"
```
