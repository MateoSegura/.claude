# Go Coding Standard

> **Version**: 1.1.0
> **Status**: Active
> **Base Standards**: [Effective Go](https://go.dev/doc/effective_go), [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments), [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes coding conventions for Go development. It ensures:

- **Consistency**: Uniform code style across all Go projects
- **Quality**: Reduced defects through proven patterns and idioms
- **Maintainability**: Code that is easy to read, modify, and extend
- **Security**: Prevention of common vulnerabilities through safe patterns

### 1.2 Scope

This standard applies to:
- [x] All Go source files in production repositories
- [x] Test code (with documented exceptions)
- [x] Build scripts and tooling written in Go

### 1.3 Audience

- Software engineers writing Go code
- Code reviewers evaluating Go changes
- Tech leads establishing project standards

### 1.4 Relationship to Industry Standards

| Standard | Relationship |
|----------|--------------|
| [Effective Go](https://go.dev/doc/effective_go) | **Base** - Foundational Go idioms and patterns |
| [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) | **Base** - Code review best practices |
| [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) | **Supplementary** - Additional patterns and conventions |
| [golang-standards/project-layout](https://github.com/golang-standards/project-layout) | **Reference** - Project structure guidance |

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

Each rule includes:
- **Rule ID**: Unique identifier (e.g., `GO-ERR-001`)
- **Tier**: Enforcement level
- **Rationale**: Why this rule exists
- **Example**: Correct and incorrect code

---

## 3. Project Structure

### 3.1 Directory Layout

```
project-root/
├── cmd/
│   └── myapp/
│       └── main.go           # Application entry point
├── internal/
│   ├── app/                  # Application-specific code
│   │   └── myapp/
│   ├── pkg/                  # Shared internal packages
│   │   ├── config/
│   │   └── middleware/
│   ├── domain/               # Business logic and entities
│   ├── repository/           # Data access layer
│   └── service/              # Service layer
├── pkg/                      # Public library code (if applicable)
│   └── client/
├── api/                      # API definitions (OpenAPI, protobuf)
├── configs/                  # Configuration file templates
├── scripts/                  # Build and utility scripts
├── testdata/                 # Test fixtures
├── go.mod
├── go.sum
└── Makefile
```

### 3.2 File Naming

| Type | Convention | Example |
|------|------------|---------|
| Source files | lowercase, underscores | `user_service.go` |
| Test files | `*_test.go` | `user_service_test.go` |
| Platform-specific | `*_os.go` | `file_linux.go` |
| Build constraints | `*_arch.go` | `asm_amd64.go` |

### 3.3 Module Organization

#### GO-STR-001: Use internal/ for Private Packages :red_circle:

**Rationale**: The `internal/` directory has special compiler treatment. Packages inside cannot be imported by external modules, enforcing encapsulation at the module level.

```go
// ✅ Correct - Internal packages are protected
// project/internal/auth/token.go
package auth

func GenerateToken(userID string) string { ... }

// ❌ Incorrect - Exposing implementation details in pkg/
// project/pkg/auth/token.go  <- Can be imported externally
package auth

func GenerateToken(userID string) string { ... }
```

#### GO-STR-002: Keep main.go Minimal :yellow_circle:

**Rationale**: The main package should only handle bootstrapping. Business logic belongs in internal packages for testability and reuse.

```go
// ✅ Correct - Minimal main.go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"

    "myapp/internal/app"
    "myapp/internal/config"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
    defer cancel()

    if err := app.Run(ctx, cfg); err != nil {
        log.Fatal(err)
    }
}

// ❌ Incorrect - Business logic in main
package main

import (
    "database/sql"
    "net/http"
)

func main() {
    db, _ := sql.Open("postgres", "...")

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        rows, _ := db.Query("SELECT * FROM users")
        // ... 100 lines of handler logic
    })

    http.ListenAndServe(":8080", nil)
}
```

#### GO-STR-003: Avoid Generic Package Names :yellow_circle:

**Rationale**: Package names like `util`, `common`, `helpers`, `misc`, and `base` do not convey specific functionality and become dumping grounds for unrelated code.

```go
// ✅ Correct - Descriptive package names
import (
    "myapp/internal/validator"
    "myapp/internal/stringutil"
    "myapp/internal/httputil"
)

// ❌ Incorrect - Generic package names
import (
    "myapp/internal/util"
    "myapp/internal/common"
    "myapp/internal/helpers"
)
```

---

## 4. Naming Conventions

### 4.1 General Principles

- Names should be descriptive and reveal intent
- Use MixedCaps or mixedCaps, never underscores
- Short names for limited scope; longer names for wider scope
- Package names qualify exported names (avoid `chubby.ChubbyFile`, prefer `chubby.File`)

### 4.2 Specific Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Variables (local) | mixedCaps, short | `i`, `buf`, `conn` |
| Variables (package) | mixedCaps, descriptive | `defaultTimeout` |
| Constants | MixedCaps | `MaxRetries`, `DefaultPort` |
| Functions | MixedCaps | `ParseConfig`, `handleRequest` |
| Types | MixedCaps | `UserService`, `Config` |
| Interfaces | Method + "er" suffix | `Reader`, `Closer`, `Handler` |
| Packages | lowercase, single word | `http`, `json`, `config` |
| Receiver | 1-2 letters, consistent | `(s *Server)`, `(c *Client)` |

### 4.3 Naming Rules

#### GO-NAM-001: Use MixedCaps for All Names :red_circle:

**Rationale**: Go convention uses MixedCaps (exported) and mixedCaps (unexported). Underscores break the visual flow and are reserved for generated code and test functions.

```go
// ✅ Correct
var maxRetryCount = 3
type HTTPClient struct{}
func ParseJSON(data []byte) error { ... }

// ❌ Incorrect
var max_retry_count = 3
type HTTP_Client struct{}
func Parse_JSON(data []byte) error { ... }
```

#### GO-NAM-002: Use Short Variable Names in Limited Scope :green_circle:

**Rationale**: Per Go Code Review Comments, the basic rule is: the further from its declaration that a name is used, the more descriptive the name must be. Single-letter names are appropriate for loop indices and short-lived variables.

```go
// ✅ Correct
for i, v := range items {
    process(v)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // w and r are idiomatic for HTTP handlers
}

// ❌ Incorrect - Overly verbose for limited scope
for index, value := range items {
    process(value)
}

func (server *Server) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
) {
    // Too verbose for standard handler signature
}
```

#### GO-NAM-003: Avoid Getters with "Get" Prefix :yellow_circle:

**Rationale**: Go does not provide automatic getter/setter support. If you have a field called `owner`, the getter should be `Owner()`, not `GetOwner()`. The setter can use `SetOwner()`.

```go
// ✅ Correct
type User struct {
    name string
}

func (u *User) Name() string     { return u.name }
func (u *User) SetName(n string) { u.name = n }

// ❌ Incorrect
func (u *User) GetName() string  { return u.name }
```

#### GO-NAM-004: Name Interfaces by Method + "er" :yellow_circle:

**Rationale**: Single-method interfaces should use the method name plus an "-er" suffix. This is idiomatic and makes the interface's purpose immediately clear.

```go
// ✅ Correct
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Stringer interface {
    String() string
}

type UserRepository interface {
    // Multi-method interfaces use descriptive names
    Find(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, user *User) error
}

// ❌ Incorrect
type IReader interface {  // Don't use I prefix
    Read(p []byte) (n int, err error)
}

type ReadInterface interface {  // Avoid "Interface" suffix
    Read(p []byte) (n int, err error)
}
```

---

## 5. Formatting Rules

### 5.1 Indentation and Spacing

| Aspect | Rule |
|--------|------|
| Indentation | Tabs (enforced by gofmt) |
| Line length | No hard limit; use judgment |
| Trailing whitespace | Forbidden |
| Final newline | Required |

### 5.2 Braces and Blocks

Go enforces brace style via `gofmt`. Opening braces must be on the same line as the statement.

### 5.3 Import Ordering

#### GO-FMT-001: Group Imports in Standard Order :yellow_circle:

**Rationale**: Consistent import ordering improves readability and makes it easy to locate dependencies. Use `goimports` to automate this.

```go
// ✅ Correct - Three groups separated by blank lines
import (
    // Standard library
    "context"
    "fmt"
    "net/http"

    // Third-party packages
    "github.com/gorilla/mux"
    "go.uber.org/zap"

    // Internal packages
    "myapp/internal/config"
    "myapp/internal/service"
)

// ❌ Incorrect - Mixed or unsorted imports
import (
    "myapp/internal/config"
    "fmt"
    "github.com/gorilla/mux"
    "context"
    "myapp/internal/service"
)
```

#### GO-FMT-002: Avoid Import Renaming Unless Necessary :green_circle:

**Rationale**: Renamed imports require readers to mentally map names. Only rename to avoid collisions or for extremely long package names.

```go
// ✅ Correct - Rename only when necessary
import (
    "crypto/rand"
    mrand "math/rand"  // Collision with crypto/rand
)

// ✅ Correct - Local package shadows standard library
import (
    "encoding/json"

    localjson "myapp/internal/json"  // Collision with encoding/json
)

// ❌ Incorrect - Unnecessary renaming
import (
    j "encoding/json"  // No collision, unnecessary
)
```

---

## 6. Language Idioms

### 6.1 Preferred Patterns

- Use composite literals for struct initialization
- Prefer `make` for slices, maps, and channels
- Use `defer` for cleanup operations
- Accept interfaces, return structs
- Keep interfaces small (1-3 methods)

### 6.2 Anti-Patterns

- Global mutable state
- Naked returns in long functions
- Empty interface (`interface{}`) without type assertions
- Ignoring error returns
- Premature optimization

### 6.3 Idiom Rules

#### GO-IDM-001: Accept Interfaces, Return Structs :yellow_circle:

**Rationale**: Functions should accept interfaces to be flexible about inputs but return concrete types to be explicit about outputs. This enables callers to use dependency injection while providing clear return types.

```go
// ✅ Correct - Accept interface, return concrete type
type UserRepository interface {
    Find(ctx context.Context, id string) (*User, error)
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// ❌ Incorrect - Returning interface hides implementation
func NewUserService(repo UserRepository) Service {
    return &UserService{repo: repo}
}
```

#### GO-IDM-002: Use Composite Literals for Initialization :green_circle:

**Rationale**: Composite literals are clearer than calling `new()` and then assigning fields. They allow partial initialization and make field assignments explicit.

```go
// ✅ Correct - Composite literal with named fields
cfg := &Config{
    Host:    "localhost",
    Port:    8080,
    Timeout: 30 * time.Second,
}

// ✅ Correct - Zero value initialization when appropriate
var buf bytes.Buffer  // Zero value is ready to use

// ❌ Incorrect - Using new() then assigning fields
cfg := new(Config)
cfg.Host = "localhost"
cfg.Port = 8080
cfg.Timeout = 30 * time.Second
```

#### GO-IDM-003: Use defer for Resource Cleanup :yellow_circle:

**Rationale**: `defer` ensures cleanup code runs when the function returns, regardless of how it exits. This prevents resource leaks and makes code cleaner.

```go
// ✅ Correct - Defer cleanup immediately after acquiring resource
func ReadFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()  // Guaranteed to run

    return io.ReadAll(f)
}

// ✅ Correct - Defer mutex unlock
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    val, ok := c.data[key]
    return val, ok
}

// ❌ Incorrect - Manual cleanup prone to errors
func ReadFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    data, err := io.ReadAll(f)
    if err != nil {
        f.Close()  // Easy to forget
        return nil, err
    }

    f.Close()
    return data, nil
}
```

### 6.4 Generics (Go 1.18+)

Go 1.18 introduced generics (type parameters), enabling type-safe reusable code. Use generics judiciously to reduce duplication without sacrificing readability.

#### GO-IDM-004: Use Generics for Type-Safe Collections and Algorithms :yellow_circle:

**Rationale**: Generics eliminate the need for `interface{}` and type assertions in collection types and algorithms. They provide compile-time type safety and better performance by avoiding boxing.

```go
// ✅ Correct - Generic function for finding element in slice
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// ✅ Correct - Generic data structure
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage - type-safe at compile time
intStack := &Stack[int]{}
intStack.Push(42)
val, ok := intStack.Pop()  // val is int, not interface{}

// ❌ Incorrect - Using interface{} when generics are appropriate
type Stack struct {
    items []interface{}
}

func (s *Stack) Pop() interface{} {
    // Requires type assertion by caller
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

val := stack.Pop().(int)  // Runtime panic if type is wrong
```

#### GO-IDM-005: Prefer Interfaces Over Generics for Behavior Abstraction :yellow_circle:

**Rationale**: Interfaces define behavior contracts and enable polymorphism. Generics are for type parameterization. Use interfaces when you need to abstract behavior; use generics when you need the same logic for different types.

```go
// ✅ Correct - Interface for behavior abstraction
type Processor interface {
    Process(ctx context.Context, data []byte) error
}

func RunPipeline(ctx context.Context, processors []Processor, data []byte) error {
    for _, p := range processors {
        if err := p.Process(ctx, data); err != nil {
            return err
        }
    }
    return nil
}

// ✅ Correct - Generics for type-parameterized operations
func Map[T, U any](items []T, fn func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}

// ❌ Incorrect - Generics where interface is more appropriate
func RunPipeline[T Processor](ctx context.Context, processors []T, data []byte) error {
    // Generic adds no value here; interface is cleaner
    for _, p := range processors {
        if err := p.Process(ctx, data); err != nil {
            return err
        }
    }
    return nil
}

// ❌ Incorrect - Interface where generics provide type safety
func Map(items []interface{}, fn func(interface{}) interface{}) []interface{} {
    // Loses type safety, requires casting everywhere
    result := make([]interface{}, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}
```

#### GO-IDM-006: Use Appropriate Type Constraints :yellow_circle:

**Rationale**: Type constraints specify what operations are available on type parameters. Use the most restrictive constraint that satisfies your needs. Prefer standard constraints from the `constraints` package or define custom constraints for domain-specific requirements.

```go
// ✅ Correct - Using standard constraints
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// ✅ Correct - Custom constraint for specific behavior
type Number interface {
    ~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](values []T) T {
    var total T
    for _, v := range values {
        total += v
    }
    return total
}

// ✅ Correct - Constraint with method requirement
type Validator interface {
    Validate() error
}

func ValidateAll[T Validator](items []T) error {
    for _, item := range items {
        if err := item.Validate(); err != nil {
            return err
        }
    }
    return nil
}

// ❌ Incorrect - Using 'any' when more specific constraint is needed
func Min[T any](a, b T) T {  // Compile error: cannot compare T
    if a < b {
        return a
    }
    return b
}

// ❌ Incorrect - Overly permissive constraint
func ProcessNumbers[T any](values []T) T {  // T should be constrained to numbers
    var total T
    for _, v := range values {
        total += v  // Compile error: cannot use += on T
    }
    return total
}
```

---

## 7. Error Handling

### 7.1 Error Handling Strategy

- Errors are values; treat them as such
- Handle errors explicitly at each call site
- Wrap errors with context using `fmt.Errorf` and `%w`
- Use sentinel errors or custom types for errors that callers need to match

### 7.2 Error Propagation

Errors should bubble up with additional context at each level. Use `%w` verb for wrapping to preserve the error chain.

### 7.3 Logging Requirements

| Level | When to Use | Example |
|-------|-------------|---------|
| ERROR | Unrecoverable failures | Failed database connection |
| WARN | Recoverable issues | Retry succeeded after failure |
| INFO | Significant events | Server started, request completed |
| DEBUG | Development details | SQL queries, cache hits |

### 7.4 Error Handling Rules

#### GO-ERR-001: Never Ignore Errors :red_circle:

**Rationale**: Silently ignoring errors leads to hard-to-debug issues. Every error must be handled, returned, or explicitly acknowledged.

```go
// ✅ Correct - Handle or return the error
data, err := json.Marshal(user)
if err != nil {
    return fmt.Errorf("marshaling user: %w", err)
}

// ✅ Correct - Explicit acknowledgment when truly ignorable
_ = conn.Close()  // Best-effort cleanup, error logged elsewhere

// ❌ Incorrect - Silently ignoring error
data, _ := json.Marshal(user)

// ❌ Incorrect - Error not checked
json.Marshal(user)
```

#### GO-ERR-002: Wrap Errors with Context :yellow_circle:

**Rationale**: Error messages should provide enough context to diagnose issues without reading the code. Use `%w` to preserve the error chain for `errors.Is` and `errors.As`.

```go
// ✅ Correct - Wrap with context using %w
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.Find(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("getting user %s: %w", id, err)
    }
    return user, nil
}

// ✅ Correct - Chain of wrapped errors
// Output: "processing order 123: getting user abc: sql: no rows"

// ❌ Incorrect - No context
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.Find(ctx, id)
    if err != nil {
        return nil, err  // Loses context
    }
    return user, nil
}

// ❌ Incorrect - Using %v loses error chain
return nil, fmt.Errorf("getting user: %v", err)  // Can't use errors.Is
```

#### GO-ERR-003: Error Strings Should Not Be Capitalized :yellow_circle:

**Rationale**: Error strings are often printed following other context. Capitalization creates awkward output like `"Reading config: Something bad"`.

```go
// ✅ Correct - Lowercase, no punctuation
return fmt.Errorf("reading config file: %w", err)
return errors.New("connection refused")

// ❌ Incorrect - Capitalized and/or punctuated
return fmt.Errorf("Reading config file: %w", err)
return errors.New("Connection refused.")
```

#### GO-ERR-004: Use Sentinel Errors for Expected Conditions :yellow_circle:

**Rationale**: When callers need to handle specific error conditions, export sentinel errors or custom error types. This enables type-safe error handling with `errors.Is` and `errors.As`.

```go
// ✅ Correct - Sentinel errors for expected conditions
var (
    ErrNotFound      = errors.New("not found")
    ErrAlreadyExists = errors.New("already exists")
    ErrUnauthorized  = errors.New("unauthorized")
)

func (r *UserRepo) Find(ctx context.Context, id string) (*User, error) {
    user, err := r.db.Get(ctx, id)
    if err == sql.ErrNoRows {
        return nil, ErrNotFound
    }
    if err != nil {
        return nil, fmt.Errorf("querying user: %w", err)
    }
    return user, nil
}

// Caller can check specific conditions
user, err := repo.Find(ctx, id)
if errors.Is(err, ErrNotFound) {
    return nil, status.Error(codes.NotFound, "user not found")
}

// ❌ Incorrect - String matching is fragile
if err.Error() == "not found" {  // Don't do this
    // ...
}
```

#### GO-ERR-005: Don't Panic for Normal Error Handling :red_circle:

**Rationale**: `panic` should only be used for truly unrecoverable errors or programming bugs. Normal error conditions should return errors.

```go
// ✅ Correct - Return error for expected failures
func ParseConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("reading config: %w", err)
    }
    // ...
    return cfg, nil
}

// ✅ Correct - Panic for programmer errors in init
func MustCompileRegex(pattern string) *regexp.Regexp {
    re, err := regexp.Compile(pattern)
    if err != nil {
        panic(fmt.Sprintf("invalid regex %q: %v", pattern, err))
    }
    return re
}

var emailRegex = MustCompileRegex(`^[a-z]+@[a-z]+\.[a-z]+$`)

// ❌ Incorrect - Panic for normal errors
func ParseConfig(path string) *Config {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(err)  // Don't panic for expected failures
    }
    // ...
}
```

---

## 8. Interface Design

### 8.1 Interface Principles

- Keep interfaces small (1-3 methods preferred)
- Define interfaces where they are used, not where they are implemented
- Don't export interfaces for mocking; let consumers define their own
- Avoid interface pollution (premature abstraction)

### 8.2 Interface Rules

#### GO-INT-001: Keep Interfaces Small :yellow_circle:

**Rationale**: Small interfaces are easier to implement, mock, and compose. The standard library's `io.Reader` and `io.Writer` are exemplary.

```go
// ✅ Correct - Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Compose when needed
type ReadWriter interface {
    Reader
    Writer
}

// ❌ Incorrect - Large interface is hard to implement and mock
type Storage interface {
    Get(key string) ([]byte, error)
    Set(key string, value []byte) error
    Delete(key string) error
    List(prefix string) ([]string, error)
    Watch(prefix string) (<-chan Event, error)
    Transaction(fn func(Tx) error) error
    // ... 10 more methods
}
```

#### GO-INT-002: Define Interfaces at Consumer Site :yellow_circle:

**Rationale**: Consumers should define the interfaces they need. This inverts the dependency and prevents interface pollution in library code.

```go
// ✅ Correct - Consumer defines the interface they need
// In package handler/
type UserGetter interface {
    GetUser(ctx context.Context, id string) (*User, error)
}

type Handler struct {
    users UserGetter
}

func NewHandler(users UserGetter) *Handler {
    return &Handler{users: users}
}

// Producer just implements the methods
// In package service/
type UserService struct { ... }

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // ...
}

// ❌ Incorrect - Producer exports interface
// In package service/
type UserServiceInterface interface {  // Over-engineering
    GetUser(ctx context.Context, id string) (*User, error)
    CreateUser(ctx context.Context, user *User) error
    // ...
}
```

#### GO-INT-003: Verify Interface Compliance at Compile Time :green_circle:

**Rationale**: Use compile-time assertions to verify a type implements an interface. This catches implementation drift early.

```go
// ✅ Correct - Compile-time interface check
var _ http.Handler = (*Server)(nil)
var _ io.ReadWriteCloser = (*Connection)(nil)

type Server struct { ... }

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // ...
}

// Fails to compile if Server doesn't implement http.Handler
```

---

## 9. Concurrency Patterns

### 9.1 Philosophy

> **Do not communicate by sharing memory; instead, share memory by communicating.**

### 9.2 Channels vs Mutexes

| Use Channels When | Use Mutexes When |
|-------------------|------------------|
| Passing ownership of data | Protecting internal state |
| Distributing work | Caching |
| Communicating async results | Simple counters |

### 9.3 Concurrency Rules

#### GO-CON-001: Always Pass Context to Long-Running Operations :red_circle:

**Rationale**: Context enables cancellation, timeouts, and deadline propagation. Without it, goroutines can't be stopped gracefully.

```go
// ✅ Correct - Accept and check context
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

// ❌ Incorrect - No context, can't cancel
func (s *Service) Process(items []Item) error {
    for _, item := range items {
        if err := s.processItem(item); err != nil {
            return err
        }
    }
    return nil
}
```

#### GO-CON-002: Never Store Context in Structs :red_circle:

**Rationale**: Context is request-scoped. Storing it in structs leads to stale contexts and unclear ownership. Always pass as the first parameter.

```go
// ✅ Correct - Context as first parameter
type Server struct {
    db *sql.DB
}

func (s *Server) GetUser(ctx context.Context, id string) (*User, error) {
    row := s.db.QueryRowContext(ctx, "SELECT ...", id)
    // ...
}

// ❌ Incorrect - Context stored in struct
type Server struct {
    ctx context.Context  // Don't do this
    db  *sql.DB
}

func (s *Server) GetUser(id string) (*User, error) {
    row := s.db.QueryRowContext(s.ctx, "SELECT ...", id)
    // ...
}
```

#### GO-CON-003: Ensure Goroutines Can Exit :red_circle:

**Rationale**: Leaked goroutines consume resources indefinitely. Every goroutine must have a clear exit path, typically via context cancellation or channel close.

```go
// ✅ Correct - Goroutine exits on context cancellation
func (s *Service) StartWorker(ctx context.Context) {
    go func() {
        ticker := time.NewTicker(time.Second)
        defer ticker.Stop()

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

// ❌ Incorrect - Goroutine can never exit
func (s *Service) StartWorker() {
    go func() {
        for {
            time.Sleep(time.Second)
            s.doWork()
        }
    }()
}
```

#### GO-CON-004: Use sync.WaitGroup for Goroutine Coordination :yellow_circle:

**Rationale**: WaitGroup provides a clean way to wait for multiple goroutines. Always call `Add` before starting the goroutine.

```go
// ✅ Correct - WaitGroup with Add before goroutine
func ProcessAll(ctx context.Context, items []Item) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(items))

    for _, item := range items {
        wg.Add(1)  // Add before starting goroutine
        go func(item Item) {
            defer wg.Done()
            if err := process(ctx, item); err != nil {
                errCh <- err
            }
        }(item)
    }

    wg.Wait()
    close(errCh)

    for err := range errCh {
        if err != nil {
            return err
        }
    }
    return nil
}

// ❌ Incorrect - Add inside goroutine (race condition)
for _, item := range items {
    go func(item Item) {
        wg.Add(1)  // Race condition!
        defer wg.Done()
        process(item)
    }(item)
}
```

#### GO-CON-005: Prefer Mutex for Simple State Protection :green_circle:

**Rationale**: Channels are for communication; mutexes are for protecting state. Using channels for simple state protection adds unnecessary complexity.

```go
// ✅ Correct - Mutex for state protection
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// ❌ Incorrect - Channel overkill for simple counter
type Counter struct {
    inc   chan struct{}
    value chan int
}

func NewCounter() *Counter {
    c := &Counter{
        inc:   make(chan struct{}),
        value: make(chan int),
    }
    go c.run()  // Extra goroutine for simple counter
    return c
}
```

#### GO-CON-006: Use errgroup for Concurrent Operations with Error Handling :yellow_circle:

**Rationale**: `errgroup` from `golang.org/x/sync/errgroup` provides a cleaner pattern than manual `sync.WaitGroup` with error channels. It automatically cancels remaining goroutines when one fails and returns the first error. Use `errgroup` when you need concurrent execution with error handling; use `WaitGroup` when errors aren't applicable or need custom handling.

```go
// ✅ Correct - errgroup for concurrent operations with errors
import "golang.org/x/sync/errgroup"

func FetchAllData(ctx context.Context, urls []string) ([][]byte, error) {
    g, ctx := errgroup.WithContext(ctx)
    results := make([][]byte, len(urls))

    for i, url := range urls {
        i, url := i, url  // Capture loop variables (not needed in Go 1.22+)
        g.Go(func() error {
            data, err := fetchURL(ctx, url)
            if err != nil {
                return fmt.Errorf("fetching %s: %w", url, err)
            }
            results[i] = data
            return nil
        })
    }

    if err := g.Wait(); err != nil {
        return nil, err  // Returns first error encountered
    }
    return results, nil
}

// ✅ Correct - errgroup with context cancellation
func ProcessItems(ctx context.Context, items []Item) error {
    g, ctx := errgroup.WithContext(ctx)

    // When one goroutine fails, ctx is canceled, signaling others to stop
    for _, item := range items {
        item := item
        g.Go(func() error {
            select {
            case <-ctx.Done():
                return ctx.Err()  // Exit early if another goroutine failed
            default:
            }
            return processItem(ctx, item)
        })
    }

    return g.Wait()
}

// ✅ Correct - errgroup with concurrency limit
func ProcessWithLimit(ctx context.Context, items []Item) error {
    g, ctx := errgroup.WithContext(ctx)
    g.SetLimit(10)  // Max 10 concurrent goroutines

    for _, item := range items {
        item := item
        g.Go(func() error {
            return processItem(ctx, item)
        })
    }

    return g.Wait()
}

// When to use WaitGroup vs errgroup:
//
// Use WaitGroup when:
// - Operations don't return errors (fire-and-forget)
// - You need custom error aggregation (collect all errors, not just first)
// - You need to track success/failure counts
//
// Use errgroup when:
// - You want to fail fast on first error
// - You need automatic context cancellation
// - You want cleaner code without manual error channels

// ❌ Incorrect - Manual WaitGroup with error channel (verbose)
func FetchAllDataManual(ctx context.Context, urls []string) ([][]byte, error) {
    var wg sync.WaitGroup
    results := make([][]byte, len(urls))
    errCh := make(chan error, len(urls))

    for i, url := range urls {
        wg.Add(1)
        go func(i int, url string) {
            defer wg.Done()
            data, err := fetchURL(ctx, url)
            if err != nil {
                errCh <- fmt.Errorf("fetching %s: %w", url, err)
                return
            }
            results[i] = data
        }(i, url)
    }

    wg.Wait()
    close(errCh)

    // Must manually check for errors
    for err := range errCh {
        if err != nil {
            return nil, err
        }
    }
    return results, nil
}
```

---

## 10. Testing Requirements

### 10.1 Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% | 85% |
| Integration tests | 50% | 70% |

### 10.2 Test Organization

- Test files live alongside source files (`foo_test.go` next to `foo.go`)
- Use `testdata/` directory for test fixtures
- Integration tests can use build tags: `//go:build integration`

### 10.3 Test Naming

| Element | Convention | Example |
|---------|------------|---------|
| Test files | `*_test.go` | `handler_test.go` |
| Test functions | `Test<Function>_<Scenario>` | `TestGetUser_NotFound` |
| Benchmark functions | `Benchmark<Function>` | `BenchmarkParse` |
| Example functions | `Example<Function>` | `ExampleNewClient` |
| Test helpers | Unexported | `setupTestDB()` |

### 10.4 Testing Rules

#### GO-TST-001: Use Table-Driven Tests :yellow_circle:

**Rationale**: Table-driven tests reduce duplication, make it easy to add cases, and clearly show the inputs and expected outputs for each scenario.

```go
// ✅ Correct - Table-driven test
func TestParsePort(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    int
        wantErr bool
    }{
        {
            name:  "valid port",
            input: "8080",
            want:  8080,
        },
        {
            name:    "negative port",
            input:   "-1",
            wantErr: true,
        },
        {
            name:    "port too high",
            input:   "65536",
            wantErr: true,
        },
        {
            name:    "non-numeric",
            input:   "abc",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParsePort(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("ParsePort(%q) error = %v, wantErr %v",
                    tt.input, err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("ParsePort(%q) = %v, want %v",
                    tt.input, got, tt.want)
            }
        })
    }
}

// ❌ Incorrect - Repetitive test functions
func TestParsePort_Valid(t *testing.T) {
    got, err := ParsePort("8080")
    if err != nil {
        t.Fatal(err)
    }
    if got != 8080 {
        t.Errorf("got %d, want 8080", got)
    }
}

func TestParsePort_Negative(t *testing.T) {
    _, err := ParsePort("-1")
    if err == nil {
        t.Error("expected error")
    }
}
// ... more duplicate functions
```

#### GO-TST-002: Use t.Helper() for Test Helpers :yellow_circle:

**Rationale**: `t.Helper()` marks a function as a test helper. When tests fail, the stack trace will point to the test that called the helper, not the helper itself.

```go
// ✅ Correct - Helper function marked with t.Helper()
func assertNoError(t *testing.T, err error) {
    t.Helper()  // Points failures to caller
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func createTestUser(t *testing.T, db *sql.DB) *User {
    t.Helper()
    user := &User{Name: "test"}
    err := db.Create(user)
    if err != nil {
        t.Fatalf("creating test user: %v", err)
    }
    return user
}

func TestGetUser(t *testing.T) {
    db := setupTestDB(t)
    user := createTestUser(t, db)  // Failure points here
    // ...
}

// ❌ Incorrect - Missing t.Helper()
func assertNoError(t *testing.T, err error) {
    // Missing t.Helper() - failure points here instead of caller
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}
```

#### GO-TST-003: Use t.Cleanup() for Test Teardown :green_circle:

**Rationale**: `t.Cleanup()` ensures cleanup functions run after the test completes, even if it panics. This is cleaner than defer and works with parallel tests.

```go
// ✅ Correct - Using t.Cleanup()
func setupTestDB(t *testing.T) *sql.DB {
    t.Helper()

    db, err := sql.Open("postgres", testDSN)
    if err != nil {
        t.Fatalf("opening db: %v", err)
    }

    t.Cleanup(func() {
        db.Close()
    })

    return db
}

func TestUser(t *testing.T) {
    db := setupTestDB(t)  // Cleanup is automatic
    // ... test code
}

// ❌ Incorrect - Manual cleanup in every test
func TestUser(t *testing.T) {
    db, err := sql.Open("postgres", testDSN)
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()  // Have to remember this every time
    // ...
}
```

#### GO-TST-004: Use Subtests for Parallel Execution :green_circle:

**Rationale**: Subtests with `t.Run()` enable parallel execution, better organization, and selective test running with `-run` flag.

```go
// ✅ Correct - Parallel subtests
func TestUserService(t *testing.T) {
    t.Parallel()  // Mark parent as parallel

    t.Run("Create", func(t *testing.T) {
        t.Parallel()  // Run this subtest in parallel
        // ...
    })

    t.Run("Update", func(t *testing.T) {
        t.Parallel()
        // ...
    })

    t.Run("Delete", func(t *testing.T) {
        t.Parallel()
        // ...
    })
}
```

### 10.5 Mocking and Test Doubles

Effective testing requires isolating units from their dependencies. Go's interface-based design makes it easy to substitute real implementations with test doubles.

#### Recommended Mocking Libraries

| Library | Use Case | Notes |
|---------|----------|-------|
| [testify/mock](https://github.com/stretchr/testify) | General-purpose mocking | Widely adopted, assertion helpers included |
| [go.uber.org/mock](https://github.com/uber-go/mock) | Generated mocks | Formerly golang/mock, type-safe generated mocks |
| Manual fakes | Simple interfaces | No external dependency, full control |

#### GO-TST-005: Use Interfaces for External Dependencies to Enable Mocking :yellow_circle:

**Rationale**: Interfaces allow substituting real implementations with test doubles. Define interfaces at the consumer site and inject dependencies to enable isolated unit testing without external services.

```go
// ✅ Correct - Interface-based dependency injection
// Define interface at consumer site (in the service package)
type UserRepository interface {
    Find(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, user *User) error
}

type EmailSender interface {
    Send(ctx context.Context, to, subject, body string) error
}

// Service accepts interfaces, enabling test doubles
type UserService struct {
    repo   UserRepository
    email  EmailSender
}

func NewUserService(repo UserRepository, email EmailSender) *UserService {
    return &UserService{repo: repo, email: email}
}

// In test file - manual fake implementation
type fakeUserRepo struct {
    users map[string]*User
    err   error
}

func (f *fakeUserRepo) Find(ctx context.Context, id string) (*User, error) {
    if f.err != nil {
        return nil, f.err
    }
    user, ok := f.users[id]
    if !ok {
        return nil, ErrNotFound
    }
    return user, nil
}

func (f *fakeUserRepo) Save(ctx context.Context, user *User) error {
    if f.err != nil {
        return f.err
    }
    f.users[user.ID] = user
    return nil
}

func TestUserService_GetUser(t *testing.T) {
    // Arrange - create fake with test data
    repo := &fakeUserRepo{
        users: map[string]*User{
            "123": {ID: "123", Name: "Alice"},
        },
    }
    email := &fakeEmailSender{}
    svc := NewUserService(repo, email)

    // Act
    user, err := svc.GetUser(context.Background(), "123")

    // Assert
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("got name %q, want %q", user.Name, "Alice")
    }
}

// ✅ Correct - Using testify/mock
import "github.com/stretchr/testify/mock"

type MockUserRepo struct {
    mock.Mock
}

func (m *MockUserRepo) Find(ctx context.Context, id string) (*User, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepo) Save(ctx context.Context, user *User) error {
    args := m.Called(ctx, user)
    return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
    mockRepo := new(MockUserRepo)
    mockEmail := new(MockEmailSender)

    // Set expectations
    mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*User")).Return(nil)
    mockEmail.On("Send", mock.Anything, "alice@example.com", mock.Anything, mock.Anything).Return(nil)

    svc := NewUserService(mockRepo, mockEmail)
    err := svc.CreateUser(context.Background(), &User{Email: "alice@example.com"})

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Verify expectations were met
    mockRepo.AssertExpectations(t)
    mockEmail.AssertExpectations(t)
}

// ❌ Incorrect - Concrete dependency, untestable without real database
type UserService struct {
    db *sql.DB  // Concrete type, can't substitute in tests
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // Requires real database connection to test
    row := s.db.QueryRowContext(ctx, "SELECT ...", id)
    // ...
}

// ❌ Incorrect - Global dependency, hard to test
var globalDB *sql.DB

func GetUser(ctx context.Context, id string) (*User, error) {
    // Global state makes testing difficult
    row := globalDB.QueryRowContext(ctx, "SELECT ...", id)
    // ...
}
```

#### GO-TST-006: Prefer Fakes Over Mocks for Complex Behavior :green_circle:

**Rationale**: Fakes provide working implementations with simplified behavior (e.g., in-memory storage), while mocks verify interactions. Fakes are better for complex dependencies where verifying exact call sequences would make tests brittle.

```go
// ✅ Correct - Fake implementation for complex behavior
// fake_storage.go (can be in a testutil package)
type FakeStorage struct {
    mu    sync.RWMutex
    items map[string][]byte
}

func NewFakeStorage() *FakeStorage {
    return &FakeStorage{items: make(map[string][]byte)}
}

func (f *FakeStorage) Get(ctx context.Context, key string) ([]byte, error) {
    f.mu.RLock()
    defer f.mu.RUnlock()

    data, ok := f.items[key]
    if !ok {
        return nil, ErrNotFound
    }
    return data, nil
}

func (f *FakeStorage) Put(ctx context.Context, key string, data []byte) error {
    f.mu.Lock()
    defer f.mu.Unlock()

    f.items[key] = data
    return nil
}

func (f *FakeStorage) Delete(ctx context.Context, key string) error {
    f.mu.Lock()
    defer f.mu.Unlock()

    delete(f.items, key)
    return nil
}

// Test can exercise realistic scenarios
func TestCacheService_Integration(t *testing.T) {
    storage := NewFakeStorage()
    cache := NewCacheService(storage)

    // Test realistic workflow
    err := cache.Set(context.Background(), "key1", []byte("value1"))
    if err != nil {
        t.Fatal(err)
    }

    val, err := cache.Get(context.Background(), "key1")
    if err != nil {
        t.Fatal(err)
    }
    if string(val) != "value1" {
        t.Errorf("got %q, want %q", val, "value1")
    }

    // Can test edge cases naturally
    _, err = cache.Get(context.Background(), "nonexistent")
    if !errors.Is(err, ErrNotFound) {
        t.Errorf("got error %v, want ErrNotFound", err)
    }
}

// ❌ Incorrect - Overly specific mock expectations
func TestCacheService_TooSpecific(t *testing.T) {
    mockStorage := new(MockStorage)

    // Brittle: tests implementation details, not behavior
    mockStorage.On("Get", mock.Anything, "key1").Return(nil, ErrNotFound).Once()
    mockStorage.On("Put", mock.Anything, "key1", []byte("value1")).Return(nil).Once()
    mockStorage.On("Get", mock.Anything, "key1").Return([]byte("value1"), nil).Once()

    // If implementation changes (e.g., checks cache before storage),
    // this test breaks even though behavior is correct
}
```

---

## 11. Documentation

### 11.1 Comment Style

- Use `//` for all comments (avoid `/* */`)
- Doc comments start with the name of the thing they describe
- Complete sentences with proper punctuation

### 11.2 Documentation Requirements

| Element | Documentation Required |
|---------|----------------------|
| Exported functions | Yes - purpose, parameters, return values |
| Exported types | Yes - purpose, usage |
| Exported constants | Yes if not self-explanatory |
| Packages | Yes - package comment in doc.go |
| Complex algorithms | Yes - explain the approach |

### 11.3 API Documentation

Go generates documentation from comments using godoc format. All exported declarations should have doc comments.

### 11.4 Documentation Rules

#### GO-DOC-001: Document All Exported Declarations :yellow_circle:

**Rationale**: Doc comments appear in godoc and IDE tooltips. They are essential for API usability. Comments should be complete sentences starting with the name.

```go
// ✅ Correct - Complete doc comments
// Config holds the application configuration.
// It is typically loaded from environment variables or a config file.
type Config struct {
    // Host is the server hostname to bind to.
    Host string

    // Port is the TCP port number. Must be between 1 and 65535.
    Port int

    // Timeout is the request timeout duration.
    Timeout time.Duration
}

// NewConfig creates a Config with default values.
// The defaults are suitable for local development.
func NewConfig() *Config {
    return &Config{
        Host:    "localhost",
        Port:    8080,
        Timeout: 30 * time.Second,
    }
}

// ❌ Incorrect - Missing or incomplete docs
type Config struct {
    Host    string
    Port    int
    Timeout time.Duration
}

// Creates config
func NewConfig() *Config {  // Doesn't start with function name
    // ...
}
```

#### GO-DOC-002: Use Package Comments :yellow_circle:

**Rationale**: Package comments provide an overview that appears in godoc. For multi-file packages, use a `doc.go` file.

```go
// ✅ Correct - Package comment in doc.go
// Package auth provides authentication and authorization utilities.
//
// It supports multiple authentication methods including JWT tokens,
// API keys, and OAuth2. The package is designed to be used with
// standard net/http handlers.
//
// Basic usage:
//
//	middleware := auth.NewMiddleware(auth.Config{
//	    Secret: os.Getenv("JWT_SECRET"),
//	})
//	http.Handle("/api/", middleware.Wrap(apiHandler))
package auth

// ❌ Incorrect - Missing package comment
package auth

// Or too terse:
// auth package
package auth
```

---

## 12. Security Considerations

### 12.1 Input Validation

- Validate all external input at system boundaries
- Use allowlists over denylists
- Sanitize data before use in SQL, HTML, or commands

### 12.2 Secret Management

- Never hardcode secrets
- Use environment variables or secret management services
- Don't log sensitive data

### 12.3 Dependency Management

- Keep dependencies updated
- Audit dependencies for vulnerabilities (`govulncheck`)
- Prefer standard library when possible

### 12.4 Security Rules

#### GO-SEC-001: Never Hardcode Secrets :red_circle:

**Rationale**: Hardcoded secrets end up in version control and logs. They cannot be rotated without code changes.

```go
// ✅ Correct - Secrets from environment or config
func NewClient() *Client {
    return &Client{
        apiKey: os.Getenv("API_KEY"),
    }
}

// ✅ Correct - Secret from injected config
func NewClient(cfg Config) *Client {
    return &Client{
        apiKey: cfg.APIKey,
    }
}

// ❌ Incorrect - Hardcoded secret
func NewClient() *Client {
    return &Client{
        apiKey: "sk-1234567890abcdef",  // NEVER do this
    }
}
```

#### GO-SEC-002: Use Parameterized Queries :red_circle:

**Rationale**: SQL injection is a critical vulnerability. Always use parameterized queries, never string concatenation.

```go
// ✅ Correct - Parameterized query
func (r *UserRepo) Find(ctx context.Context, id string) (*User, error) {
    row := r.db.QueryRowContext(ctx,
        "SELECT id, name, email FROM users WHERE id = $1", id)
    // ...
}

// ✅ Correct - Using query builder with parameters
func (r *UserRepo) Search(ctx context.Context, name string) ([]User, error) {
    query := sq.Select("id", "name", "email").
        From("users").
        Where(sq.Eq{"name": name})
    // ...
}

// ❌ Incorrect - String concatenation (SQL injection!)
func (r *UserRepo) Find(ctx context.Context, id string) (*User, error) {
    query := "SELECT id, name, email FROM users WHERE id = '" + id + "'"
    row := r.db.QueryRowContext(ctx, query)
    // ...
}
```

#### GO-SEC-003: Validate File Paths to Prevent Directory Traversal :red_circle:

**Rationale**: Directory traversal attacks exploit unchecked file paths to access files outside intended directories. Always clean paths with `filepath.Clean` and validate they remain within the expected base directory.

```go
// ✅ Correct - Validate path stays within base directory
func SafeReadFile(baseDir, userPath string) ([]byte, error) {
    // Clean the path to resolve . and ..
    cleanPath := filepath.Clean(userPath)

    // Ensure no absolute path escape
    if filepath.IsAbs(cleanPath) {
        return nil, errors.New("absolute paths not allowed")
    }

    // Join with base and clean again
    fullPath := filepath.Join(baseDir, cleanPath)

    // Verify the result is still under baseDir
    if !strings.HasPrefix(fullPath, filepath.Clean(baseDir)+string(filepath.Separator)) {
        return nil, errors.New("path traversal detected")
    }

    return os.ReadFile(fullPath)
}

// ✅ Correct - Using filepath.Rel to validate
func SafeServePath(baseDir, requestPath string) (string, error) {
    fullPath := filepath.Join(baseDir, filepath.Clean(requestPath))

    // Get relative path from base - fails if fullPath escapes baseDir
    rel, err := filepath.Rel(baseDir, fullPath)
    if err != nil || strings.HasPrefix(rel, "..") {
        return "", errors.New("invalid path")
    }

    return fullPath, nil
}

// ❌ Incorrect - Direct path concatenation
func UnsafeReadFile(baseDir, userPath string) ([]byte, error) {
    // Vulnerable: userPath = "../../../etc/passwd"
    fullPath := baseDir + "/" + userPath
    return os.ReadFile(fullPath)
}

// ❌ Incorrect - filepath.Join alone is not sufficient
func StillUnsafe(baseDir, userPath string) ([]byte, error) {
    // filepath.Join cleans but doesn't validate containment
    // userPath = "../../../etc/passwd" still escapes baseDir
    fullPath := filepath.Join(baseDir, userPath)
    return os.ReadFile(fullPath)  // Can read /etc/passwd!
}
```

#### GO-SEC-004: Use html/template for HTML Output :red_circle:

**Rationale**: `text/template` does not escape HTML, making it vulnerable to Cross-Site Scripting (XSS) attacks. Always use `html/template` when generating HTML output, as it automatically escapes content based on context.

```go
// ✅ Correct - Using html/template for HTML
import "html/template"

func RenderPage(w http.ResponseWriter, data PageData) error {
    tmpl, err := template.ParseFiles("page.html")
    if err != nil {
        return err
    }
    // html/template automatically escapes data.UserName
    // <script>alert('xss')</script> becomes safe text
    return tmpl.Execute(w, data)
}

// ✅ Correct - Parsing HTML templates from embedded files
//go:embed templates/*.html
var templateFS embed.FS

var templates = template.Must(template.ParseFS(templateFS, "templates/*.html"))

func RenderUser(w http.ResponseWriter, user User) error {
    return templates.ExecuteTemplate(w, "user.html", user)
}

// ❌ Incorrect - Using text/template for HTML (XSS vulnerability!)
import "text/template"  // WRONG for HTML!

func RenderPage(w http.ResponseWriter, data PageData) error {
    tmpl, err := template.ParseFiles("page.html")
    if err != nil {
        return err
    }
    // text/template does NOT escape HTML
    // data.UserName = "<script>alert('xss')</script>" executes!
    return tmpl.Execute(w, data)
}

// ❌ Incorrect - String concatenation for HTML
func RenderUser(w http.ResponseWriter, name string) {
    // Direct injection vulnerability
    fmt.Fprintf(w, "<h1>Welcome, %s</h1>", name)
}
```

#### GO-SEC-005: Sanitize Inputs to exec.Command :red_circle:

**Rationale**: Command injection occurs when untrusted input is passed to shell commands. Avoid shell execution when possible; use `exec.Command` with explicit arguments instead of shell interpolation. Validate and sanitize all command inputs.

```go
// ✅ Correct - exec.Command with separate arguments (no shell)
func GitClone(repoURL, destDir string) error {
    // Validate URL format
    if !isValidGitURL(repoURL) {
        return errors.New("invalid repository URL")
    }

    // Arguments are passed directly, not through shell
    cmd := exec.Command("git", "clone", "--depth", "1", repoURL, destDir)
    return cmd.Run()
}

// ✅ Correct - Allowlist validation for dynamic commands
var allowedCommands = map[string]bool{
    "status": true,
    "logs":   true,
    "info":   true,
}

func RunDockerCommand(container, command string) ([]byte, error) {
    // Validate against allowlist
    if !allowedCommands[command] {
        return nil, fmt.Errorf("command %q not allowed", command)
    }

    // Validate container name format
    if !isValidContainerName(container) {
        return nil, errors.New("invalid container name")
    }

    cmd := exec.Command("docker", command, container)
    return cmd.Output()
}

func isValidContainerName(name string) bool {
    // Only allow alphanumeric, underscore, hyphen
    matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, name)
    return matched
}

// ❌ Incorrect - Shell execution with user input (command injection!)
func UnsafeGitClone(repoURL string) error {
    // repoURL = "https://evil.com; rm -rf /"
    cmd := exec.Command("sh", "-c", "git clone "+repoURL)
    return cmd.Run()
}

// ❌ Incorrect - Unvalidated input to command
func UnsafeDockerLogs(container string) ([]byte, error) {
    // container = "foo; cat /etc/passwd"
    cmd := exec.Command("sh", "-c", "docker logs "+container)
    return cmd.Output()
}
```

---

## 13. Performance Guidelines

### 13.1 Memory Management

- Preallocate slices when size is known
- Reuse buffers with `sync.Pool` for hot paths
- Be mindful of string/byte conversions

### 13.2 Optimization Policy

1. Write clear, correct code first
2. Measure with benchmarks and profiling
3. Optimize only proven bottlenecks
4. Document why optimizations exist

### 13.3 Performance Rules

#### GO-PRF-001: Preallocate Slices When Size is Known :green_circle:

**Rationale**: Appending to slices may cause reallocation. Preallocating avoids unnecessary allocations and copies.

```go
// ✅ Correct - Preallocate slice
func ProcessUsers(users []User) []Result {
    results := make([]Result, 0, len(users))  // Preallocate capacity
    for _, u := range users {
        results = append(results, process(u))
    }
    return results
}

// ✅ Correct - Direct indexing when length is known
func ProcessUsers(users []User) []Result {
    results := make([]Result, len(users))  // Preallocate with length
    for i, u := range users {
        results[i] = process(u)
    }
    return results
}

// ❌ Incorrect - Growing slice dynamically
func ProcessUsers(users []User) []Result {
    var results []Result  // Will reallocate multiple times
    for _, u := range users {
        results = append(results, process(u))
    }
    return results
}
```

#### GO-PRF-002: Use strings.Builder for String Concatenation :green_circle:

**Rationale**: String concatenation with `+` creates a new string each time. `strings.Builder` minimizes allocations.

```go
// ✅ Correct - strings.Builder for multiple concatenations
func BuildQuery(fields []string) string {
    var b strings.Builder
    b.WriteString("SELECT ")
    for i, f := range fields {
        if i > 0 {
            b.WriteString(", ")
        }
        b.WriteString(f)
    }
    b.WriteString(" FROM users")
    return b.String()
}

// ❌ Incorrect - String concatenation in loop
func BuildQuery(fields []string) string {
    query := "SELECT "
    for i, f := range fields {
        if i > 0 {
            query += ", "
        }
        query += f  // New allocation each iteration
    }
    query += " FROM users"
    return query
}
```

---

## 14. Tooling Configuration

### 14.1 Required Tools

| Tool | Purpose | Version |
|------|---------|---------|
| go | Compiler and tools | 1.21+ |
| gofmt | Code formatting | (bundled) |
| goimports | Import management | latest |
| golangci-lint | Static analysis | 1.55+ |
| govulncheck | Vulnerability scanning | latest |

### 14.2 golangci-lint Configuration

```yaml
# .golangci.yml
version: "2"

run:
  timeout: 5m

linters:
  default: standard
  enable:
    # Bug detection
    - staticcheck
    - govet
    - errcheck
    - gosec

    # Code quality
    - revive
    - gocyclo
    - gocognit
    - dupl
    - funlen

    # Style
    - gofmt
    - goimports
    - misspell
    - unconvert
    - unparam

    # Performance
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

  revive:
    rules:
      - name: exported
        arguments:
          - checkPrivateReceivers
          - disableStutteringCheck
      - name: var-naming
      - name: package-comments
      - name: dot-imports
      - name: blank-imports
      - name: context-as-argument
      - name: context-keys-type
      - name: error-return
      - name: error-strings
      - name: error-naming
      - name: increment-decrement
      - name: range
      - name: receiver-naming
      - name: time-naming
      - name: unexported-return
      - name: indent-error-flow
      - name: errorf
      - name: empty-block
      - name: superfluous-else
      - name: unused-parameter
      - name: unreachable-code
      - name: redefines-builtin-id

  gosec:
    excludes:
      - G104  # Audit errors not checked (handled by errcheck)

issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - dupl
        - funlen
        - gocyclo
        - gocognit

    - path: cmd/
      linters:
        - gochecknoinits

  max-issues-per-linter: 0
  max-same-issues: 0

formatters:
  enable:
    - gofmt
    - goimports

  settings:
    goimports:
      local-prefixes: "mycompany.com"
```

### 14.3 Pre-commit Hooks

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.1
    hooks:
      - id: go-fmt
      - id: go-imports
        args: ["-local", "mycompany.com"]
      - id: go-vet
      - id: go-build
      - id: go-mod-tidy

  - repo: https://github.com/golangci/golangci-lint
    rev: v1.55.2
    hooks:
      - id: golangci-lint
```

### 14.4 Makefile Targets

```makefile
# Makefile
.PHONY: all build test lint fmt clean

GO := go
GOLANGCI_LINT := golangci-lint

all: lint test build

build:
	$(GO) build -o bin/ ./cmd/...

test:
	$(GO) test -race -coverprofile=coverage.out ./...

test-integration:
	$(GO) test -race -tags=integration ./...

lint:
	$(GOLANGCI_LINT) run ./...

fmt:
	$(GO) fmt ./...
	goimports -w -local mycompany.com .

clean:
	rm -rf bin/ coverage.out

tidy:
	$(GO) mod tidy
	$(GO) mod verify

vulncheck:
	govulncheck ./...
```

### 14.5 IDE Configuration

Configure your IDE to align with this standard's linting and formatting rules.

#### VS Code Settings

Add to `.vscode/settings.json` in your project:

```json
{
    // Go extension settings
    "go.useLanguageServer": true,
    "go.lintTool": "golangci-lint",
    "go.lintFlags": ["--fast"],
    "go.lintOnSave": "package",

    // Format on save
    "editor.formatOnSave": true,
    "[go]": {
        "editor.defaultFormatter": "golang.go",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": "explicit"
        }
    },
    "[go.mod]": {
        "editor.formatOnSave": true
    },

    // gopls settings
    "gopls": {
        "ui.semanticTokens": true,
        "ui.completion.usePlaceholders": true,
        "ui.diagnostic.analyses": {
            "unusedparams": true,
            "shadow": true,
            "fieldalignment": false,
            "nilness": true,
            "unusedwrite": true,
            "useany": true
        },
        "ui.diagnostic.staticcheck": true,
        "formatting.gofumpt": false,
        "ui.completion.completeFunctionCalls": true
    },

    // Test settings
    "go.testFlags": ["-v", "-race"],
    "go.coverOnSave": false,
    "go.coverOnSingleTest": true,
    "go.coverOnSingleTestFile": true,

    // Build settings
    "go.buildFlags": ["-v"],
    "go.buildOnSave": "off",

    // Import settings aligned with standard
    "go.formatTool": "goimports",
    "go.alternateTools": {
        "goimports": "goimports"
    }
}
```

#### VS Code Recommended Extensions

Create `.vscode/extensions.json`:

```json
{
    "recommendations": [
        "golang.go",
        "zxh404.vscode-proto3",
        "redhat.vscode-yaml",
        "EditorConfig.EditorConfig"
    ]
}
```

#### GoLand Configuration

For JetBrains GoLand, configure the following:

**File Watchers** (Settings > Tools > File Watchers):

| Watcher | Program | Arguments | Output paths |
|---------|---------|-----------|--------------|
| goimports | `goimports` | `-w -local mycompany.com $FilePath$` | `$FilePath$` |
| golangci-lint | `golangci-lint` | `run --fix $FilePath$` | `$FilePath$` |

**Code Style** (Settings > Editor > Code Style > Go):
- Tabs and Indents: Use tab character (Go default)
- Imports: Enable "Group stdlib imports", "Move all imports in a single declaration"
- Other: Enable "Add leading space to comments"

**Inspections** (Settings > Editor > Inspections > Go):
- Enable: Unhandled error, Unreachable code, Unused parameter, Shadow declaration
- Configure severity levels to match golangci-lint settings

**Run/Debug Configuration Template**:
```
Go Build:
  - Run kind: Package
  - Package path: mycompany.com/myapp/cmd/myapp
  - Output directory: $PROJECT_DIR$/bin
  - Go tool arguments: -race

Go Test:
  - Test kind: Package
  - Package path: mycompany.com/myapp/...
  - Go tool arguments: -race -v
```

#### EditorConfig

Add `.editorconfig` to project root for consistent formatting across editors:

```ini
# .editorconfig
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true

[*.go]
indent_style = tab
indent_size = 4

[*.{json,yaml,yml}]
indent_style = space
indent_size = 2

[Makefile]
indent_style = tab

[*.md]
trim_trailing_whitespace = false
```

---

## 15. Code Review Checklist

Quick reference for code reviewers:

### Structure and Organization
- [ ] Code is in the correct package (`internal/` for private, `pkg/` for public)
- [ ] No circular dependencies
- [ ] main.go is minimal (bootstrapping only)
- [ ] Package names are descriptive (not `util`, `common`, `helpers`)

### Naming
- [ ] Uses MixedCaps, no underscores
- [ ] Variable name length proportional to scope
- [ ] No stuttering (`user.UserName` should be `user.Name`)
- [ ] Interfaces use "-er" suffix for single methods

### Error Handling
- [ ] All errors are handled (not ignored with `_`)
- [ ] Errors are wrapped with context using `%w`
- [ ] Error strings are lowercase without punctuation
- [ ] Sentinel errors used for conditions callers need to match
- [ ] No panic for expected error conditions

### Concurrency
- [ ] Context passed as first parameter
- [ ] Context not stored in structs
- [ ] All goroutines have exit paths
- [ ] sync.WaitGroup.Add() called before starting goroutine
- [ ] Mutex used for state protection, channels for communication
- [ ] errgroup used for concurrent operations with error handling

### Interfaces
- [ ] Interfaces are small (1-3 methods)
- [ ] Interfaces defined at consumer site
- [ ] Functions accept interfaces, return structs
- [ ] Compile-time interface checks for key types

### Testing
- [ ] Table-driven tests for multiple scenarios
- [ ] t.Helper() used in test helpers
- [ ] t.Cleanup() used for teardown
- [ ] Tests can run in parallel where appropriate
- [ ] Edge cases covered
- [ ] External dependencies use interfaces for mocking
- [ ] Fakes used for complex dependency behavior

### Documentation
- [ ] All exported declarations have doc comments
- [ ] Comments start with the name being documented
- [ ] Package has a doc comment (doc.go for multi-file packages)
- [ ] Complex logic has explanatory comments

### Security
- [ ] No hardcoded secrets
- [ ] SQL uses parameterized queries
- [ ] External input is validated
- [ ] File paths validated against directory traversal
- [ ] html/template used for HTML output (not text/template)
- [ ] exec.Command inputs sanitized and validated

### Performance
- [ ] Slices preallocated when size is known
- [ ] strings.Builder used for string concatenation
- [ ] No premature optimization (benchmarks exist for optimized code)

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| GO-STR-001 | Use internal/ for Private Packages | :red_circle: |
| GO-STR-002 | Keep main.go Minimal | :yellow_circle: |
| GO-STR-003 | Avoid Generic Package Names | :yellow_circle: |
| GO-NAM-001 | Use MixedCaps for All Names | :red_circle: |
| GO-NAM-002 | Use Short Variable Names in Limited Scope | :green_circle: |
| GO-NAM-003 | Avoid Getters with "Get" Prefix | :yellow_circle: |
| GO-NAM-004 | Name Interfaces by Method + "er" | :yellow_circle: |
| GO-FMT-001 | Group Imports in Standard Order | :yellow_circle: |
| GO-FMT-002 | Avoid Import Renaming Unless Necessary | :green_circle: |
| GO-IDM-001 | Accept Interfaces, Return Structs | :yellow_circle: |
| GO-IDM-002 | Use Composite Literals for Initialization | :green_circle: |
| GO-IDM-003 | Use defer for Resource Cleanup | :yellow_circle: |
| GO-IDM-004 | Use Generics for Type-Safe Collections and Algorithms | :yellow_circle: |
| GO-IDM-005 | Prefer Interfaces Over Generics for Behavior Abstraction | :yellow_circle: |
| GO-IDM-006 | Use Appropriate Type Constraints | :yellow_circle: |
| GO-ERR-001 | Never Ignore Errors | :red_circle: |
| GO-ERR-002 | Wrap Errors with Context | :yellow_circle: |
| GO-ERR-003 | Error Strings Should Not Be Capitalized | :yellow_circle: |
| GO-ERR-004 | Use Sentinel Errors for Expected Conditions | :yellow_circle: |
| GO-ERR-005 | Don't Panic for Normal Error Handling | :red_circle: |
| GO-INT-001 | Keep Interfaces Small | :yellow_circle: |
| GO-INT-002 | Define Interfaces at Consumer Site | :yellow_circle: |
| GO-INT-003 | Verify Interface Compliance at Compile Time | :green_circle: |
| GO-CON-001 | Always Pass Context to Long-Running Operations | :red_circle: |
| GO-CON-002 | Never Store Context in Structs | :red_circle: |
| GO-CON-003 | Ensure Goroutines Can Exit | :red_circle: |
| GO-CON-004 | Use sync.WaitGroup for Goroutine Coordination | :yellow_circle: |
| GO-CON-005 | Prefer Mutex for Simple State Protection | :green_circle: |
| GO-CON-006 | Use errgroup for Concurrent Operations with Error Handling | :yellow_circle: |
| GO-TST-001 | Use Table-Driven Tests | :yellow_circle: |
| GO-TST-002 | Use t.Helper() for Test Helpers | :yellow_circle: |
| GO-TST-003 | Use t.Cleanup() for Test Teardown | :green_circle: |
| GO-TST-004 | Use Subtests for Parallel Execution | :green_circle: |
| GO-TST-005 | Use Interfaces for External Dependencies to Enable Mocking | :yellow_circle: |
| GO-TST-006 | Prefer Fakes Over Mocks for Complex Behavior | :green_circle: |
| GO-DOC-001 | Document All Exported Declarations | :yellow_circle: |
| GO-DOC-002 | Use Package Comments | :yellow_circle: |
| GO-SEC-001 | Never Hardcode Secrets | :red_circle: |
| GO-SEC-002 | Use Parameterized Queries | :red_circle: |
| GO-SEC-003 | Validate File Paths to Prevent Directory Traversal | :red_circle: |
| GO-SEC-004 | Use html/template for HTML Output | :red_circle: |
| GO-SEC-005 | Sanitize Inputs to exec.Command | :red_circle: |
| GO-PRF-001 | Preallocate Slices When Size is Known | :green_circle: |
| GO-PRF-002 | Use strings.Builder for String Concatenation | :green_circle: |

---

## Appendix B: Glossary

| Term | Definition |
|------|------------|
| Composite Literal | Expression that creates a value for struct, array, slice, or map |
| Defer | Statement that schedules a function call to run after the surrounding function returns |
| Exported | Identifier starting with uppercase letter, visible outside the package |
| Goroutine | Lightweight thread managed by the Go runtime |
| MixedCaps | Naming convention using uppercase for word boundaries (e.g., `UserService`) |
| Receiver | Parameter that associates a method with a type |
| Sentinel Error | Predeclared error value that callers can compare against |
| Zero Value | Default value for a type when not explicitly initialized |

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.1.0 | 2026-01-04 | Added: Generics section (6.4) with GO-IDM-004/005/006; Security rules GO-SEC-003/004/005; Mocking section (10.5) with GO-TST-005/006; GO-CON-006 for errgroup; IDE configuration section (14.5) |
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- [Effective Go](https://go.dev/doc/effective_go) - Official Go best practices
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) - Code review guidelines
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) - Uber's conventions
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout) - Project structure
- [Go Wiki: Table Driven Tests](https://go.dev/wiki/TableDrivenTests) - Testing patterns
- [Go Blog: Context](https://go.dev/blog/context) - Context usage patterns
- [Go Blog: Contexts and Structs](https://go.dev/blog/context-and-structs) - Context best practices
- [golangci-lint](https://golangci-lint.run/) - Linter documentation
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) - Vulnerability scanning
