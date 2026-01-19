# Interface Rules (GO-INT-*)

Interfaces are central to Go's design. These rules ensure interfaces are used effectively for abstraction and testability.

## Interface Principles

- Keep interfaces small (1-3 methods preferred)
- Define interfaces where they are used, not where they are implemented
- Don't export interfaces for mocking; let consumers define their own
- Avoid interface pollution (premature abstraction)

---

## GO-INT-001: Keep Interfaces Small :yellow_circle:

**Tier**: Required

**Rationale**: Small interfaces are easier to implement, mock, and compose.

```go
// Correct - Small, focused interfaces
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

// Incorrect - Large interface is hard to implement and mock
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

---

## GO-INT-002: Define Interfaces at Consumer Site :yellow_circle:

**Tier**: Required

**Rationale**: Consumers should define the interfaces they need. This inverts the dependency and prevents interface pollution in library code.

```go
// Correct - Consumer defines the interface they need
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

// Incorrect - Producer exports interface
// In package service/
type UserServiceInterface interface {  // Over-engineering
    GetUser(ctx context.Context, id string) (*User, error)
    CreateUser(ctx context.Context, user *User) error
}
```

---

## GO-INT-003: Verify Interface Compliance at Compile Time :green_circle:

**Tier**: Recommended

**Rationale**: Use compile-time assertions to verify a type implements an interface.

```go
// Correct - Compile-time interface check
var _ http.Handler = (*Server)(nil)
var _ io.ReadWriteCloser = (*Connection)(nil)

type Server struct { ... }

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // ...
}
```
