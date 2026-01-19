# Security Rules (GO-SEC-*)

Security is paramount in production code. These rules prevent common vulnerabilities and enforce secure coding practices.

## Input Validation

- Validate all external input at system boundaries
- Use allowlists over denylists
- Sanitize data before use in SQL, HTML, or commands

## Secret Management

- Never hardcode secrets
- Use environment variables or secret management services
- Don't log sensitive data

---

## GO-SEC-001: Never Hardcode Secrets :red_circle:

**Tier**: Critical

**Rationale**: Hardcoded secrets end up in version control and logs. They cannot be rotated without code changes.

```go
// Correct - Secrets from environment or config
func NewClient() *Client {
    return &Client{
        apiKey: os.Getenv("API_KEY"),
    }
}

// Correct - Secret from injected config
func NewClient(cfg Config) *Client {
    return &Client{
        apiKey: cfg.APIKey,
    }
}

// Incorrect - Hardcoded secret
func NewClient() *Client {
    return &Client{
        apiKey: "sk-1234567890abcdef",  // NEVER do this
    }
}
```

---

## GO-SEC-002: Use Parameterized Queries :red_circle:

**Tier**: Critical

**Rationale**: SQL injection is a critical vulnerability. Always use parameterized queries, never string concatenation.

```go
// Correct - Parameterized query
func (r *UserRepo) Find(ctx context.Context, id string) (*User, error) {
    row := r.db.QueryRowContext(ctx,
        "SELECT id, name, email FROM users WHERE id = $1", id)
    // ...
}

// Incorrect - String concatenation (SQL injection!)
func (r *UserRepo) Find(ctx context.Context, id string) (*User, error) {
    query := "SELECT id, name, email FROM users WHERE id = '" + id + "'"
    row := r.db.QueryRowContext(ctx, query)
    // ...
}
```

---

## GO-SEC-003: Validate File Paths to Prevent Directory Traversal :red_circle:

**Tier**: Critical

**Rationale**: Directory traversal attacks exploit unchecked file paths to access files outside intended directories.

```go
// Correct - Validate path stays within base directory
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

// Incorrect - Direct path concatenation
func UnsafeReadFile(baseDir, userPath string) ([]byte, error) {
    // Vulnerable: userPath = "../../../etc/passwd"
    fullPath := baseDir + "/" + userPath
    return os.ReadFile(fullPath)
}
```

---

## GO-SEC-004: Use html/template for HTML Output :red_circle:

**Tier**: Critical

**Rationale**: `text/template` does not escape HTML, making it vulnerable to Cross-Site Scripting (XSS) attacks.

```go
// Correct - Using html/template for HTML
import "html/template"

func RenderPage(w http.ResponseWriter, data PageData) error {
    tmpl, err := template.ParseFiles("page.html")
    if err != nil {
        return err
    }
    // html/template automatically escapes data.UserName
    return tmpl.Execute(w, data)
}

// Incorrect - Using text/template for HTML (XSS vulnerability!)
import "text/template"  // WRONG for HTML!

func RenderPage(w http.ResponseWriter, data PageData) error {
    tmpl, err := template.ParseFiles("page.html")
    if err != nil {
        return err
    }
    // text/template does NOT escape HTML
    return tmpl.Execute(w, data)
}
```

---

## GO-SEC-005: Sanitize Inputs to exec.Command :red_circle:

**Tier**: Critical

**Rationale**: Command injection occurs when untrusted input is passed to shell commands. Avoid shell execution when possible.

```go
// Correct - exec.Command with separate arguments (no shell)
func GitClone(repoURL, destDir string) error {
    // Validate URL format
    if !isValidGitURL(repoURL) {
        return errors.New("invalid repository URL")
    }

    // Arguments are passed directly, not through shell
    cmd := exec.Command("git", "clone", "--depth", "1", repoURL, destDir)
    return cmd.Run()
}

// Correct - Allowlist validation for dynamic commands
var allowedCommands = map[string]bool{
    "status": true,
    "logs":   true,
    "info":   true,
}

func RunDockerCommand(container, command string) ([]byte, error) {
    if !allowedCommands[command] {
        return nil, fmt.Errorf("command %q not allowed", command)
    }

    if !isValidContainerName(container) {
        return nil, errors.New("invalid container name")
    }

    cmd := exec.Command("docker", command, container)
    return cmd.Output()
}

// Incorrect - Shell execution with user input (command injection!)
func UnsafeGitClone(repoURL string) error {
    // repoURL = "https://evil.com; rm -rf /"
    cmd := exec.Command("sh", "-c", "git clone "+repoURL)
    return cmd.Run()
}
```
