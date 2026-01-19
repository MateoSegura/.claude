# Structure Rules (GO-STR-*)

Project structure is fundamental to maintainability. These rules establish conventions for organizing Go code.

## Directory Layout

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

## File Naming

| Type | Convention | Example |
|------|------------|---------|
| Source files | lowercase, underscores | `user_service.go` |
| Test files | `*_test.go` | `user_service_test.go` |
| Platform-specific | `*_os.go` | `file_linux.go` |
| Build constraints | `*_arch.go` | `asm_amd64.go` |

---

## GO-STR-001: Use internal/ for Private Packages :red_circle:

**Tier**: Critical

**Rationale**: The `internal/` directory has special compiler treatment. Packages inside cannot be imported by external modules, enforcing encapsulation at the module level.

```go
// Correct - Internal packages are protected
// project/internal/auth/token.go
package auth

func GenerateToken(userID string) string { ... }

// Incorrect - Exposing implementation details in pkg/
// project/pkg/auth/token.go  <- Can be imported externally
package auth

func GenerateToken(userID string) string { ... }
```

---

## GO-STR-002: Keep main.go Minimal :yellow_circle:

**Tier**: Required

**Rationale**: The main package should only handle bootstrapping. Business logic belongs in internal packages for testability and reuse.

```go
// Correct - Minimal main.go
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

// Incorrect - Business logic in main
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

---

## GO-STR-003: Avoid Generic Package Names :yellow_circle:

**Tier**: Required

**Rationale**: Package names like `util`, `common`, `helpers`, `misc`, and `base` do not convey specific functionality and become dumping grounds for unrelated code.

```go
// Correct - Descriptive package names
import (
    "myapp/internal/validator"
    "myapp/internal/stringutil"
    "myapp/internal/httputil"
)

// Incorrect - Generic package names
import (
    "myapp/internal/util"
    "myapp/internal/common"
    "myapp/internal/helpers"
)
```
