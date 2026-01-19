# Error Handling Rules (GO-ERR-*)

Error handling is fundamental to Go's design philosophy. These rules ensure errors are handled explicitly and provide sufficient context for debugging.

## Error Handling Strategy

- Errors are values; treat them as such
- Handle errors explicitly at each call site
- Wrap errors with context using `fmt.Errorf` and `%w`
- Use sentinel errors or custom types for errors that callers need to match

## Logging Requirements

| Level | When to Use | Example |
|-------|-------------|---------|
| ERROR | Unrecoverable failures | Failed database connection |
| WARN | Recoverable issues | Retry succeeded after failure |
| INFO | Significant events | Server started, request completed |
| DEBUG | Development details | SQL queries, cache hits |

---

## GO-ERR-001: Never Ignore Errors :red_circle:

**Tier**: Critical

**Rationale**: Silently ignoring errors leads to hard-to-debug issues. Every error must be handled, returned, or explicitly acknowledged.

```go
// Correct - Handle or return the error
data, err := json.Marshal(user)
if err != nil {
    return fmt.Errorf("marshaling user: %w", err)
}

// Correct - Explicit acknowledgment when truly ignorable
_ = conn.Close()  // Best-effort cleanup, error logged elsewhere

// Incorrect - Silently ignoring error
data, _ := json.Marshal(user)

// Incorrect - Error not checked
json.Marshal(user)
```

---

## GO-ERR-002: Wrap Errors with Context :yellow_circle:

**Tier**: Required

**Rationale**: Error messages should provide enough context to diagnose issues without reading the code. Use `%w` to preserve the error chain for `errors.Is` and `errors.As`.

```go
// Correct - Wrap with context using %w
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.Find(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("getting user %s: %w", id, err)
    }
    return user, nil
}

// Incorrect - No context
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.Find(ctx, id)
    if err != nil {
        return nil, err  // Loses context
    }
    return user, nil
}

// Incorrect - Using %v loses error chain
return nil, fmt.Errorf("getting user: %v", err)  // Can't use errors.Is
```

---

## GO-ERR-003: Error Strings Should Not Be Capitalized :yellow_circle:

**Tier**: Required

**Rationale**: Error strings are often printed following other context. Capitalization creates awkward output.

```go
// Correct - Lowercase, no punctuation
return fmt.Errorf("reading config file: %w", err)
return errors.New("connection refused")

// Incorrect - Capitalized and/or punctuated
return fmt.Errorf("Reading config file: %w", err)
return errors.New("Connection refused.")
```

---

## GO-ERR-004: Use Sentinel Errors for Expected Conditions :yellow_circle:

**Tier**: Required

**Rationale**: When callers need to handle specific error conditions, export sentinel errors or custom error types.

```go
// Correct - Sentinel errors for expected conditions
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
```

---

## GO-ERR-005: Don't Panic for Normal Error Handling :red_circle:

**Tier**: Critical

**Rationale**: `panic` should only be used for truly unrecoverable errors or programming bugs. Normal error conditions should return errors.

```go
// Correct - Return error for expected failures
func ParseConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("reading config: %w", err)
    }
    return cfg, nil
}

// Correct - Panic for programmer errors in init
func MustCompileRegex(pattern string) *regexp.Regexp {
    re, err := regexp.Compile(pattern)
    if err != nil {
        panic(fmt.Sprintf("invalid regex %q: %v", pattern, err))
    }
    return re
}

var emailRegex = MustCompileRegex(`^[a-z]+@[a-z]+\.[a-z]+$`)

// Incorrect - Panic for normal errors
func ParseConfig(path string) *Config {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(err)  // Don't panic for expected failures
    }
}
```
