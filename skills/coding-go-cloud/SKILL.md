---
name: coding-go-cloud
description: Go coding standard for cloud services and APIs
---

# Go Coding Standard

> **Version**: 1.1.0 | **Status**: Active
> **Base Standards**: Effective Go, Go Code Review Comments, Uber Go Style Guide

This standard establishes coding conventions for Go development, ensuring consistency, quality, maintainability, and security across all Go projects.

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [Error Handling](rules/error-handling.md) | GO-ERR-* | Never ignore errors, wrap with context |
| [Concurrency](rules/concurrency.md) | GO-CON-* | Context handling, goroutine safety |
| [Security](rules/security.md) | GO-SEC-* | SQL injection, XSS, command injection |
| [Naming](rules/naming.md) | GO-NAM-* | MixedCaps, interface naming |
| [Testing](rules/testing.md) | GO-TST-* | Table-driven tests, mocking |
| [Structure](rules/structure.md) | GO-STR-* | Project layout, internal packages |
| [Interfaces](rules/interfaces.md) | GO-INT-* | Small interfaces, consumer-defined |
| [Idioms](rules/idioms.md) | GO-IDM-* | Go patterns and anti-patterns |
| [Documentation](rules/documentation.md) | GO-DOC-* | Doc comments, package docs |
| [Performance](rules/performance.md) | GO-PRF-* | Preallocation, string building |
| [CI/CD](rules/cicd.md) | GO-CICD-* | Build, test, lint, security scanning |

### Tooling

| Tool | File |
|------|------|
| [golangci-lint](tooling/golangci-lint.md) | Linter configuration |
| [Pre-commit](tooling/pre-commit.md) | Git hooks setup |
| [IDE Setup](tooling/ide-setup.md) | VS Code, GoLand, Neovim |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table with tiers |
| [Code Review Checklist](reference/code-review.md) | Review checklist by category |

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Critical Rules (Always Apply)

These rules are non-negotiable and must be followed in all code.

### GO-ERR-001: Never Ignore Errors :red_circle:

Every error must be handled, returned, or explicitly acknowledged.

```go
// Correct
data, err := json.Marshal(user)
if err != nil {
    return fmt.Errorf("marshaling user: %w", err)
}

// Correct - Explicit acknowledgment
_ = conn.Close()  // Best-effort cleanup

// Incorrect
data, _ := json.Marshal(user)
```

### GO-ERR-005: Don't Panic for Normal Error Handling :red_circle:

Use `panic` only for truly unrecoverable errors or programming bugs.

```go
// Correct
func ParseConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("reading config: %w", err)
    }
    return cfg, nil
}

// Incorrect
func ParseConfig(path string) *Config {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(err)  // Don't panic for expected failures
    }
}
```

### GO-CON-001: Always Pass Context to Long-Running Operations :red_circle:

Context enables cancellation, timeouts, and deadline propagation.

```go
// Correct
func (s *Service) Process(ctx context.Context, items []Item) error {
    for _, item := range items {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        if err := s.processItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}
```

### GO-CON-002: Never Store Context in Structs :red_circle:

Context is request-scoped. Pass it as the first parameter.

```go
// Correct
func (s *Server) GetUser(ctx context.Context, id string) (*User, error)

// Incorrect
type Server struct {
    ctx context.Context  // Don't do this
}
```

### GO-CON-003: Ensure Goroutines Can Exit :red_circle:

Every goroutine must have a clear exit path.

```go
// Correct
func (s *Service) StartWorker(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return  // Clean exit
            case <-ticker.C:
                s.doWork()
            }
        }
    }()
}
```

### GO-SEC-001: Never Hardcode Secrets :red_circle:

Use environment variables or secret management services.

```go
// Correct
apiKey: os.Getenv("API_KEY")

// Incorrect
apiKey: "sk-1234567890abcdef"  // NEVER do this
```

### GO-SEC-002: Use Parameterized Queries :red_circle:

Prevent SQL injection with parameterized queries.

```go
// Correct
row := r.db.QueryRowContext(ctx,
    "SELECT id, name FROM users WHERE id = $1", id)

// Incorrect - SQL injection vulnerability!
query := "SELECT id, name FROM users WHERE id = '" + id + "'"
```

### GO-SEC-003: Validate File Paths :red_circle:

Prevent directory traversal attacks.

```go
// Correct
cleanPath := filepath.Clean(userPath)
fullPath := filepath.Join(baseDir, cleanPath)
if !strings.HasPrefix(fullPath, filepath.Clean(baseDir)+string(filepath.Separator)) {
    return nil, errors.New("path traversal detected")
}
```

### GO-SEC-004: Use html/template for HTML Output :red_circle:

Prevent XSS attacks.

```go
// Correct
import "html/template"

// Incorrect - XSS vulnerability!
import "text/template"  // Does NOT escape HTML
```

### GO-SEC-005: Sanitize Inputs to exec.Command :red_circle:

Avoid shell execution; pass arguments directly.

```go
// Correct
cmd := exec.Command("git", "clone", repoURL, destDir)

// Incorrect - Command injection!
cmd := exec.Command("sh", "-c", "git clone "+repoURL)
```

### GO-STR-001: Use internal/ for Private Packages :red_circle:

Packages in `internal/` cannot be imported by external modules.

### GO-NAM-001: Use MixedCaps for All Names :red_circle:

Never use underscores in Go names (except generated code and tests).

```go
// Correct
var maxRetryCount = 3
type HTTPClient struct{}

// Incorrect
var max_retry_count = 3
type HTTP_Client struct{}
```

---

## Quick Rule Lookup

| ID | Rule | Tier |
|----|------|------|
| GO-STR-001 | Use internal/ for Private Packages | Critical |
| GO-STR-002 | Keep main.go Minimal | Required |
| GO-STR-003 | Avoid Generic Package Names | Required |
| GO-NAM-001 | Use MixedCaps for All Names | Critical |
| GO-NAM-002 | Use Short Variable Names in Limited Scope | Recommended |
| GO-NAM-003 | Avoid Getters with "Get" Prefix | Required |
| GO-NAM-004 | Name Interfaces by Method + "er" | Required |
| GO-FMT-001 | Group Imports in Standard Order | Required |
| GO-FMT-002 | Avoid Import Renaming Unless Necessary | Recommended |
| GO-IDM-001 | Accept Interfaces, Return Structs | Required |
| GO-IDM-002 | Use Composite Literals for Initialization | Recommended |
| GO-IDM-003 | Use defer for Resource Cleanup | Required |
| GO-IDM-004 | Use Generics for Type-Safe Collections | Required |
| GO-IDM-005 | Prefer Interfaces Over Generics for Behavior | Required |
| GO-IDM-006 | Use Appropriate Type Constraints | Required |
| GO-ERR-001 | Never Ignore Errors | Critical |
| GO-ERR-002 | Wrap Errors with Context | Required |
| GO-ERR-003 | Error Strings Should Not Be Capitalized | Required |
| GO-ERR-004 | Use Sentinel Errors for Expected Conditions | Required |
| GO-ERR-005 | Don't Panic for Normal Error Handling | Critical |
| GO-INT-001 | Keep Interfaces Small | Required |
| GO-INT-002 | Define Interfaces at Consumer Site | Required |
| GO-INT-003 | Verify Interface Compliance at Compile Time | Recommended |
| GO-CON-001 | Always Pass Context to Long-Running Operations | Critical |
| GO-CON-002 | Never Store Context in Structs | Critical |
| GO-CON-003 | Ensure Goroutines Can Exit | Critical |
| GO-CON-004 | Use sync.WaitGroup for Goroutine Coordination | Required |
| GO-CON-005 | Prefer Mutex for Simple State Protection | Recommended |
| GO-CON-006 | Use errgroup for Concurrent Operations | Required |
| GO-TST-001 | Use Table-Driven Tests | Required |
| GO-TST-002 | Use t.Helper() for Test Helpers | Required |
| GO-TST-003 | Use t.Cleanup() for Test Teardown | Recommended |
| GO-TST-004 | Use Subtests for Parallel Execution | Recommended |
| GO-TST-005 | Use Interfaces for External Dependencies | Required |
| GO-TST-006 | Prefer Fakes Over Mocks for Complex Behavior | Recommended |
| GO-DOC-001 | Document All Exported Declarations | Required |
| GO-DOC-002 | Use Package Comments | Required |
| GO-SEC-001 | Never Hardcode Secrets | Critical |
| GO-SEC-002 | Use Parameterized Queries | Critical |
| GO-SEC-003 | Validate File Paths | Critical |
| GO-SEC-004 | Use html/template for HTML Output | Critical |
| GO-SEC-005 | Sanitize Inputs to exec.Command | Critical |
| GO-PRF-001 | Preallocate Slices When Size is Known | Recommended |
| GO-PRF-002 | Use strings.Builder for String Concatenation | Recommended |

---

## References

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
