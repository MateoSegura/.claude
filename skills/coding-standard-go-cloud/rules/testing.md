# Testing Rules (GO-TST-*)

Testing is essential for maintaining code quality. These rules ensure tests are effective, maintainable, and follow Go conventions.

## Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% | 85% |
| Integration tests | 50% | 70% |

## Test Organization

- Test files live alongside source files (`foo_test.go` next to `foo.go`)
- Use `testdata/` directory for test fixtures
- Integration tests can use build tags: `//go:build integration`

## Test Naming

| Element | Convention | Example |
|---------|------------|---------|
| Test files | `*_test.go` | `handler_test.go` |
| Test functions | `Test<Function>_<Scenario>` | `TestGetUser_NotFound` |
| Benchmark functions | `Benchmark<Function>` | `BenchmarkParse` |
| Example functions | `Example<Function>` | `ExampleNewClient` |
| Test helpers | Unexported | `setupTestDB()` |

## Mocking Libraries

| Library | Use Case | Notes |
|---------|----------|-------|
| testify/mock | General-purpose mocking | Widely adopted, assertion helpers included |
| go.uber.org/mock | Generated mocks | Formerly golang/mock, type-safe generated mocks |
| Manual fakes | Simple interfaces | No external dependency, full control |

---

## GO-TST-001: Use Table-Driven Tests :yellow_circle:

**Tier**: Required

**Rationale**: Table-driven tests reduce duplication, make it easy to add cases, and clearly show the inputs and expected outputs.

```go
// Correct - Table-driven test
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
```

---

## GO-TST-002: Use t.Helper() for Test Helpers :yellow_circle:

**Tier**: Required

**Rationale**: `t.Helper()` marks a function as a test helper. When tests fail, the stack trace will point to the test that called the helper.

```go
// Correct - Helper function marked with t.Helper()
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
```

---

## GO-TST-003: Use t.Cleanup() for Test Teardown :green_circle:

**Tier**: Recommended

**Rationale**: `t.Cleanup()` ensures cleanup functions run after the test completes, even if it panics.

```go
// Correct - Using t.Cleanup()
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
```

---

## GO-TST-004: Use Subtests for Parallel Execution :green_circle:

**Tier**: Recommended

**Rationale**: Subtests with `t.Run()` enable parallel execution, better organization, and selective test running.

```go
// Correct - Parallel subtests
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

---

## GO-TST-005: Use Interfaces for External Dependencies to Enable Mocking :yellow_circle:

**Tier**: Required

**Rationale**: Interfaces allow substituting real implementations with test doubles.

```go
// Correct - Interface-based dependency injection
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

func TestUserService_GetUser(t *testing.T) {
    repo := &fakeUserRepo{
        users: map[string]*User{
            "123": {ID: "123", Name: "Alice"},
        },
    }
    email := &fakeEmailSender{}
    svc := NewUserService(repo, email)

    user, err := svc.GetUser(context.Background(), "123")

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("got name %q, want %q", user.Name, "Alice")
    }
}
```

---

## GO-TST-006: Prefer Fakes Over Mocks for Complex Behavior :green_circle:

**Tier**: Recommended

**Rationale**: Fakes provide working implementations with simplified behavior, while mocks verify interactions. Fakes are better for complex dependencies where verifying exact call sequences would make tests brittle.

```go
// Correct - Fake implementation for complex behavior
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
```
