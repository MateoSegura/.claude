# Naming Rules (GO-NAM-*)

Good naming is essential for code readability. These rules follow Go's established conventions.

## General Principles

- Names should be descriptive and reveal intent
- Use MixedCaps or mixedCaps, never underscores
- Short names for limited scope; longer names for wider scope
- Package names qualify exported names (avoid `chubby.ChubbyFile`, prefer `chubby.File`)

## Specific Conventions

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

---

## GO-NAM-001: Use MixedCaps for All Names :red_circle:

**Tier**: Critical

**Rationale**: Go convention uses MixedCaps (exported) and mixedCaps (unexported). Underscores break the visual flow and are reserved for generated code and test functions.

```go
// Correct
var maxRetryCount = 3
type HTTPClient struct{}
func ParseJSON(data []byte) error { ... }

// Incorrect
var max_retry_count = 3
type HTTP_Client struct{}
func Parse_JSON(data []byte) error { ... }
```

---

## GO-NAM-002: Use Short Variable Names in Limited Scope :green_circle:

**Tier**: Recommended

**Rationale**: Per Go Code Review Comments, the basic rule is: the further from its declaration that a name is used, the more descriptive the name must be.

```go
// Correct
for i, v := range items {
    process(v)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // w and r are idiomatic for HTTP handlers
}

// Incorrect - Overly verbose for limited scope
for index, value := range items {
    process(value)
}
```

---

## GO-NAM-003: Avoid Getters with "Get" Prefix :yellow_circle:

**Tier**: Required

**Rationale**: Go does not provide automatic getter/setter support. If you have a field called `owner`, the getter should be `Owner()`, not `GetOwner()`.

```go
// Correct
type User struct {
    name string
}

func (u *User) Name() string     { return u.name }
func (u *User) SetName(n string) { u.name = n }

// Incorrect
func (u *User) GetName() string  { return u.name }
```

---

## GO-NAM-004: Name Interfaces by Method + "er" :yellow_circle:

**Tier**: Required

**Rationale**: Single-method interfaces should use the method name plus an "-er" suffix. This is idiomatic and makes the interface's purpose immediately clear.

```go
// Correct
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

// Incorrect
type IReader interface {  // Don't use I prefix
    Read(p []byte) (n int, err error)
}

type ReadInterface interface {  // Avoid "Interface" suffix
    Read(p []byte) (n int, err error)
}
```
