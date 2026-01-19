# Language Idioms Rules (GO-IDM-*)

Go has established idioms that make code more readable and maintainable. These rules capture the most important patterns.

## Preferred Patterns

- Use composite literals for struct initialization
- Prefer `make` for slices, maps, and channels
- Use `defer` for cleanup operations
- Accept interfaces, return structs
- Keep interfaces small (1-3 methods)

## Anti-Patterns

- Global mutable state
- Naked returns in long functions
- Empty interface (`interface{}`) without type assertions
- Ignoring error returns
- Premature optimization

---

## GO-IDM-001: Accept Interfaces, Return Structs :yellow_circle:

**Tier**: Required

**Rationale**: Functions should accept interfaces to be flexible about inputs but return concrete types to be explicit about outputs.

```go
// Correct - Accept interface, return concrete type
type UserRepository interface {
    Find(ctx context.Context, id string) (*User, error)
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// Incorrect - Returning interface hides implementation
func NewUserService(repo UserRepository) Service {
    return &UserService{repo: repo}
}
```

---

## GO-IDM-002: Use Composite Literals for Initialization :green_circle:

**Tier**: Recommended

**Rationale**: Composite literals are clearer than calling `new()` and then assigning fields.

```go
// Correct - Composite literal with named fields
cfg := &Config{
    Host:    "localhost",
    Port:    8080,
    Timeout: 30 * time.Second,
}

// Correct - Zero value initialization when appropriate
var buf bytes.Buffer  // Zero value is ready to use

// Incorrect - Using new() then assigning fields
cfg := new(Config)
cfg.Host = "localhost"
cfg.Port = 8080
cfg.Timeout = 30 * time.Second
```

---

## GO-IDM-003: Use defer for Resource Cleanup :yellow_circle:

**Tier**: Required

**Rationale**: `defer` ensures cleanup code runs when the function returns, regardless of how it exits.

```go
// Correct - Defer cleanup immediately after acquiring resource
func ReadFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()  // Guaranteed to run

    return io.ReadAll(f)
}

// Correct - Defer mutex unlock
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    val, ok := c.data[key]
    return val, ok
}
```

---

## GO-IDM-004: Use Generics for Type-Safe Collections and Algorithms :yellow_circle:

**Tier**: Required

**Rationale**: Generics eliminate the need for `interface{}` and type assertions in collection types and algorithms.

```go
// Correct - Generic function for finding element in slice
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Correct - Generic data structure
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

// Incorrect - Using interface{} when generics are appropriate
type Stack struct {
    items []interface{}
}

func (s *Stack) Pop() interface{} {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

val := stack.Pop().(int)  // Runtime panic if type is wrong
```

---

## GO-IDM-005: Prefer Interfaces Over Generics for Behavior Abstraction :yellow_circle:

**Tier**: Required

**Rationale**: Interfaces define behavior contracts and enable polymorphism. Generics are for type parameterization. Use interfaces when you need to abstract behavior; use generics when you need the same logic for different types.

```go
// Correct - Interface for behavior abstraction
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

// Correct - Generics for type-parameterized operations
func Map[T, U any](items []T, fn func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}
```

---

## GO-IDM-006: Use Appropriate Type Constraints :yellow_circle:

**Tier**: Required

**Rationale**: Type constraints specify what operations are available on type parameters. Use the most restrictive constraint that satisfies your needs.

```go
// Correct - Using standard constraints
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Correct - Custom constraint for specific behavior
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
```
