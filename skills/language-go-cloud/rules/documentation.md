# Documentation Rules (GO-DOC-*)

Documentation is essential for API usability and maintenance. These rules ensure consistent and helpful documentation.

## Comment Style

- Use `//` for all comments (avoid `/* */`)
- Doc comments start with the name of the thing they describe
- Complete sentences with proper punctuation

## Documentation Requirements

| Element | Documentation Required |
|---------|----------------------|
| Exported functions | Yes - purpose, parameters, return values |
| Exported types | Yes - purpose, usage |
| Exported constants | Yes if not self-explanatory |
| Packages | Yes - package comment in doc.go |
| Complex algorithms | Yes - explain the approach |

---

## GO-DOC-001: Document All Exported Declarations :yellow_circle:

**Tier**: Required

**Rationale**: Doc comments appear in godoc and IDE tooltips. They are essential for API usability.

```go
// Correct - Complete doc comments
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

// Incorrect - Missing or incomplete docs
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

---

## GO-DOC-002: Use Package Comments :yellow_circle:

**Tier**: Required

**Rationale**: Package comments provide an overview that appears in godoc.

```go
// Correct - Package comment in doc.go
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
```
