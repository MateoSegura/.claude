---
description: Go CI/CD rules - build, test, lint, security scanning for Go projects
---

# Go CI/CD Rules

Go-specific CI/CD configurations and tooling. For universal CI/CD principles, see [devops-standard/rules/cicd-principles.md](../../devops-standard/rules/cicd-principles.md).

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## 1. Build Stage

### GO-CICD-001: Use `go build` with Proper Flags :red_circle:

**Rationale**: Proper build flags ensure reproducible, optimized, and secure binaries.

```bash
# Production build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-s -w -X main.version=${VERSION}" \
  -o bin/app ./cmd/app

# Development build (with race detector)
go build -race -o bin/app ./cmd/app
```

**Required flags for production**:
- `CGO_ENABLED=0` - Static binary (no C dependencies)
- `-ldflags="-s -w"` - Strip debug symbols (smaller binary)
- `-ldflags="-X main.version=..."` - Embed version info

### GO-CICD-002: Use `go mod` for Dependency Management :red_circle:

**Rationale**: Go modules ensure reproducible builds and dependency security.

```bash
# Verify dependencies match go.sum
go mod verify

# Tidy dependencies before commit
go mod tidy

# Download dependencies (CI)
go mod download
```

**Requirements**:
- Always commit `go.mod` and `go.sum`
- Run `go mod tidy` before commits
- Verify modules in CI with `go mod verify`

---

## 2. Test Stage

### GO-CICD-003: Use `go test` with Coverage :red_circle:

**Rationale**: Coverage metrics ensure tests exercise critical code paths.

```bash
# Run tests with coverage
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html

# Check coverage threshold
go tool cover -func=coverage.out | grep total | awk '{print $3}'
```

**GitHub Actions example**:

```yaml
- name: Run tests
  run: |
    go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

- name: Upload coverage
  uses: codecov/codecov-action@v4
  with:
    files: coverage.out
    fail_ci_if_error: true
```

### GO-CICD-004: Use Table-Driven Test Output :yellow_circle:

**Rationale**: Structured test output enables better CI reporting.

```bash
# JSON output for CI parsing
go test -v -json ./... > test-results.json

# Convert to JUnit XML (for CI systems)
go install github.com/jstemmer/go-junit-report/v2@latest
go test -v ./... 2>&1 | go-junit-report > report.xml
```

### GO-CICD-005: Enable Race Detector in CI :red_circle:

**Rationale**: The race detector catches concurrency bugs that are hard to reproduce.

```bash
# Always use -race in CI
go test -race ./...
go build -race ./cmd/app  # For integration tests
```

**Note**: Race detector increases memory usage 5-10x and slows execution 2-20x. Use only in CI, not production.

---

## 3. Static Analysis Stage

### GO-CICD-006: Use golangci-lint :red_circle:

**Rationale**: golangci-lint runs many linters in parallel efficiently.

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run ./...

# Run with specific config
golangci-lint run --config=.golangci.yml ./...
```

**Recommended `.golangci.yml`**:

```yaml
run:
  timeout: 5m
  go: '1.22'

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused
    - gofmt
    - goimports
    - misspell
    - unconvert
    - unparam
    - gosec
    - prealloc
    - revive
    - bodyclose
    - noctx
    - sqlclosecheck
    - exportloopref

linters-settings:
  errcheck:
    check-blank: true
  govet:
    enable-all: true
  gosec:
    severity: medium
    confidence: medium

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
```

### GO-CICD-007: Use `go vet` :red_circle:

**Rationale**: `go vet` catches bugs that compile but are likely mistakes.

```bash
# Run vet
go vet ./...

# Note: golangci-lint includes govet by default
```

### GO-CICD-008: Enforce `gofmt` Formatting :red_circle:

**Rationale**: Consistent formatting eliminates style debates.

```bash
# Check formatting (CI)
gofmt -l . | grep . && exit 1 || true

# Or with golangci-lint
golangci-lint run --enable gofmt
```

---

## 4. Security Stage

### GO-CICD-009: Use `govulncheck` for Vulnerability Scanning :red_circle:

**Rationale**: govulncheck identifies known vulnerabilities in dependencies and your code.

```bash
# Install
go install golang.org/x/vuln/cmd/govulncheck@latest

# Scan dependencies and code
govulncheck ./...

# JSON output for CI
govulncheck -json ./... > vulns.json
```

**GitHub Actions example**:

```yaml
- name: Run govulncheck
  uses: golang/govulncheck-action@v1
  with:
    go-version-input: '1.22'
    go-package: './...'
```

### GO-CICD-010: Use `gosec` for Security Linting :red_circle:

**Rationale**: gosec detects security issues in Go source code.

```bash
# Install
go install github.com/securego/gosec/v2/cmd/gosec@latest

# Run
gosec ./...

# With SARIF output for GitHub
gosec -fmt sarif -out results.sarif ./...
```

**Note**: gosec is included in golangci-lint as `gosec` linter.

### GO-CICD-011: Dependency Audit :red_circle:

**Rationale**: Audit dependencies for known vulnerabilities.

```bash
# Using go mod
go list -m -json all | nancy sleuth

# Or with govulncheck (preferred)
govulncheck ./...
```

---

## 5. Complete GitHub Actions Pipeline

```yaml
name: Go CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  GO_VERSION: '1.22'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: Verify dependencies
        run: |
          go mod verify
          go mod tidy
          git diff --exit-code go.mod go.sum

      - name: Build
        run: go build -v ./...

  test:
    runs-on: ubuntu-latest
    needs: build
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: Run tests
        run: |
          go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          files: coverage.out
          fail_ci_if_error: true

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: golangci-lint
        uses: golangci/golangci-lint-action@v4
        with:
          version: latest
          args: --timeout=5m

  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: Run govulncheck
        uses: golang/govulncheck-action@v1
        with:
          go-version-input: ${{ env.GO_VERSION }}
          go-package: './...'

      - name: Run gosec
        uses: securego/gosec@master
        with:
          args: '-fmt sarif -out results.sarif ./...'

      - name: Upload SARIF
        uses: github/codeql-action/upload-sarif@v3
        with:
          sarif_file: results.sarif
```

---

## 6. Makefile Targets

```makefile
.PHONY: build test lint security all

GO_VERSION := 1.22
BINARY_NAME := app
VERSION := $(shell git describe --tags --always --dirty)

all: lint test build

build:
	CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=$(VERSION)" \
		-o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run --timeout=5m ./...

security:
	govulncheck ./...
	gosec ./...

deps:
	go mod download
	go mod verify
	go mod tidy

clean:
	rm -rf bin/ coverage.out
```

---

## Quick Reference

| ID | Rule | Tier |
|----|------|------|
| GO-CICD-001 | Use `go build` with proper flags | Critical |
| GO-CICD-002 | Use `go mod` for dependency management | Critical |
| GO-CICD-003 | Use `go test` with coverage | Critical |
| GO-CICD-004 | Use table-driven test output | Required |
| GO-CICD-005 | Enable race detector in CI | Critical |
| GO-CICD-006 | Use golangci-lint | Critical |
| GO-CICD-007 | Use `go vet` | Critical |
| GO-CICD-008 | Enforce `gofmt` formatting | Critical |
| GO-CICD-009 | Use `govulncheck` for vulnerability scanning | Critical |
| GO-CICD-010 | Use `gosec` for security linting | Critical |
| GO-CICD-011 | Dependency audit | Critical |

---

## References

- [Go Testing](https://golang.org/pkg/testing/)
- [golangci-lint](https://golangci-lint.run/)
- [govulncheck](https://go.dev/security/vulncheck)
- [gosec](https://github.com/securego/gosec)
