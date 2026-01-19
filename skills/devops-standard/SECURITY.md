# Security Standard

> Building secure software requires vigilance at every stage of development. This standard provides actionable guidance for protecting our systems, data, and users.

## Tier Classification

| Tier | Meaning |
|------|---------|
| Critical | Must be followed. Violations block deployment. |
| Required | Expected in all code. Exceptions need documented approval. |
| Recommended | Best practice. Should be followed unless there is good reason not to. |

---

## 1. Secret Management

### 1.1 Never Commit Secrets

**Tier: Critical**

Secrets include: API keys, passwords, tokens, private keys, connection strings, and any credential that grants access to systems or data.

#### What NOT to Do

```python
# VULNERABLE: Hardcoded API key
API_KEY = "sk-live-abc123xyz789"
database_url = "postgresql://admin:secretpass@prod-db.example.com/users"
```

```javascript
// VULNERABLE: Hardcoded token
const STRIPE_KEY = "sk_live_51ABC123";
fetch(url, { headers: { Authorization: "Bearer ghp_xxxxxxxxxxxx" } });
```

#### What TO Do

```python
# SECURE: Load from environment
import os

API_KEY = os.environ.get("API_KEY")
if not API_KEY:
    raise ValueError("API_KEY environment variable is required")
```

```javascript
// SECURE: Load from environment
const API_KEY = process.env.API_KEY;
if (!API_KEY) {
  throw new Error("API_KEY environment variable is required");
}
```

### 1.2 Use Environment Variables or Secret Managers

**Tier: Critical**

| Environment | Recommended Approach |
|-------------|---------------------|
| Local Development | `.env` files (never committed) |
| CI/CD | Pipeline secrets (GitHub Secrets, GitLab CI Variables) |
| Production | Secret managers (AWS Secrets Manager, HashiCorp Vault, GCP Secret Manager) |

#### Secret Manager Example (AWS)

```python
import boto3
import json

def get_secret(secret_name: str) -> dict:
    client = boto3.client("secretsmanager")
    response = client.get_secret_value(SecretId=secret_name)
    return json.loads(response["SecretString"])

# Usage
db_creds = get_secret("prod/database/credentials")
```

### 1.3 Gitignore Patterns for Secrets

**Tier: Critical**

Add these patterns to `.gitignore`:

```gitignore
# Environment files
.env
.env.*
!.env.example

# Private keys
*.pem
*.key
*.p12
*.pfx

# AWS
.aws/credentials

# Google Cloud
*-credentials.json
service-account*.json

# IDE secrets
.idea/**/secrets.xml

# Local configuration with secrets
config.local.*
secrets.*
```

### 1.4 Responding to Accidentally Committed Secrets

**Tier: Critical**

If secrets are committed, act immediately:

1. **Revoke the secret** - Generate a new credential before doing anything else
2. **Remove from history** - Use `git filter-branch` or BFG Repo-Cleaner
3. **Force push** - Update remote after cleaning history
4. **Audit access** - Check logs for unauthorized use
5. **Report** - Notify security team per incident response process

```bash
# Using BFG Repo-Cleaner (faster than filter-branch)
bfg --delete-files secrets.json
bfg --replace-text passwords.txt

git reflog expire --expire=now --all
git gc --prune=now --aggressive
git push --force
```

**Note:** Consider the secret compromised even after removal. Anyone who cloned the repo may have it.

---

## 2. Input Validation

### 2.1 Validate All External Input

**Tier: Critical**

External input includes: HTTP requests, file uploads, database reads, API responses, command-line arguments, environment variables, and message queues.

#### Validation Strategy

```python
from pydantic import BaseModel, validator, EmailStr
from typing import Optional
import re

class UserRegistration(BaseModel):
    email: EmailStr
    username: str
    age: Optional[int] = None

    @validator("username")
    def username_valid(cls, v):
        if not re.match(r"^[a-zA-Z0-9_]{3,30}$", v):
            raise ValueError(
                "Username must be 3-30 characters, alphanumeric and underscores only"
            )
        return v

    @validator("age")
    def age_reasonable(cls, v):
        if v is not None and (v < 0 or v > 150):
            raise ValueError("Age must be between 0 and 150")
        return v
```

```go
import (
    "regexp"
    "errors"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`)

func ValidateUsername(username string) error {
    if !usernameRegex.MatchString(username) {
        return errors.New("invalid username format")
    }
    return nil
}
```

### 2.2 Sanitize Before Use

**Tier: Critical**

Different contexts require different sanitization:

| Context | Sanitization |
|---------|-------------|
| HTML output | HTML entity encoding |
| SQL queries | Parameterized queries |
| Shell commands | Avoid shell; use direct execution |
| File paths | Validate against allowlist, resolve canonical path |
| URLs | Parse and validate scheme and host |

#### HTML Sanitization

```python
import html

def safe_html_output(user_input: str) -> str:
    return html.escape(user_input)

# Input: <script>alert('xss')</script>
# Output: &lt;script&gt;alert('xss')&lt;/script&gt;
```

```javascript
// Using DOMPurify for HTML that needs to preserve formatting
import DOMPurify from "dompurify";

const cleanHTML = DOMPurify.sanitize(userHTML, {
  ALLOWED_TAGS: ["b", "i", "em", "strong", "a"],
  ALLOWED_ATTR: ["href"],
});
```

### 2.3 Use Allowlists Over Denylists

**Tier: Required**

```python
# BAD: Denylist approach (incomplete, bypassable)
BLOCKED_EXTENSIONS = [".exe", ".bat", ".sh"]

def check_file_bad(filename):
    for ext in BLOCKED_EXTENSIONS:
        if filename.endswith(ext):
            return False
    return True  # .cmd, .ps1, .scr slip through

# GOOD: Allowlist approach (explicit, complete)
ALLOWED_EXTENSIONS = {".jpg", ".jpeg", ".png", ".gif", ".pdf"}

def check_file_good(filename):
    import os
    _, ext = os.path.splitext(filename.lower())
    return ext in ALLOWED_EXTENSIONS
```

### 2.4 Language-Specific Validation Libraries

**Tier: Recommended**

| Language | Libraries |
|----------|-----------|
| Python | `pydantic`, `marshmallow`, `cerberus` |
| JavaScript/TypeScript | `zod`, `joi`, `yup` |
| Go | `go-playground/validator`, `ozzo-validation` |
| Java | Bean Validation (JSR 380), `hibernate-validator` |
| Rust | `validator` crate |

#### TypeScript Example with Zod

```typescript
import { z } from "zod";

const UserSchema = z.object({
  email: z.string().email(),
  username: z.string().min(3).max(30).regex(/^[a-zA-Z0-9_]+$/),
  age: z.number().int().min(0).max(150).optional(),
});

type User = z.infer<typeof UserSchema>;

function createUser(input: unknown): User {
  return UserSchema.parse(input); // Throws on invalid input
}
```

---

## 3. Dependency Security

### 3.1 Dependency Scanning

**Tier: Critical**

Enable automated scanning in CI/CD:

#### GitHub Dependabot

```yaml
# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "npm"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10

  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"

  - package-ecosystem: "pip"
    directory: "/"
    schedule:
      interval: "weekly"
```

#### Snyk CI Integration

```yaml
# GitHub Actions
- name: Run Snyk Security Scan
  uses: snyk/actions/node@master
  env:
    SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
  with:
    args: --severity-threshold=high
```

### 3.2 Vulnerability Response Policy

**Tier: Critical**

| Severity | Response Time | Action |
|----------|--------------|--------|
| Critical (CVSS 9.0-10.0) | 24 hours | Immediate patch or mitigation |
| High (CVSS 7.0-8.9) | 7 days | Prioritize in current sprint |
| Medium (CVSS 4.0-6.9) | 30 days | Schedule in backlog |
| Low (CVSS 0.1-3.9) | 90 days | Address when convenient |

### 3.3 Vetting New Dependencies

**Tier: Required**

Before adding a dependency, evaluate:

1. **Maintenance Status**
   - When was the last release?
   - Are issues being addressed?
   - Is there more than one maintainer?

2. **Security History**
   - Check for past CVEs
   - Review security advisories
   - Look at security-related issues

3. **Scope and Size**
   - Does it do more than you need?
   - How many transitive dependencies?
   - Could you implement the needed functionality yourself?

4. **Trust Indicators**
   - Download/star counts
   - Used by reputable projects
   - Organizational backing

```bash
# Check npm package info
npm info <package> --json | jq '{maintainers, time, repository}'

# Check for known vulnerabilities
npm audit --package-lock-only

# View dependency tree
npm ls --all
```

### 3.4 Lock File Requirements

**Tier: Critical**

Always commit lock files:

| Package Manager | Lock File |
|-----------------|-----------|
| npm | `package-lock.json` |
| yarn | `yarn.lock` |
| pnpm | `pnpm-lock.yaml` |
| pip | `requirements.txt` with pinned versions or `Pipfile.lock` |
| Go | `go.sum` |
| Rust | `Cargo.lock` |

```bash
# Install from lock file only (CI/CD)
npm ci           # npm
yarn --frozen-lockfile  # yarn
pip install -r requirements.txt --require-hashes  # pip
```

---

## 4. Authentication & Authorization

### 4.1 Secure Password Handling

**Tier: Critical**

#### Never Store Plaintext Passwords

```python
# BAD: Storing plaintext
user.password = request.form["password"]

# BAD: Using weak hashing
import hashlib
user.password_hash = hashlib.md5(password.encode()).hexdigest()

# GOOD: Using bcrypt with appropriate cost
import bcrypt

def hash_password(password: str) -> bytes:
    return bcrypt.hashpw(password.encode(), bcrypt.gensalt(rounds=12))

def verify_password(password: str, hash: bytes) -> bool:
    return bcrypt.checkpw(password.encode(), hash)
```

#### Recommended Algorithms

| Algorithm | Use Case | Notes |
|-----------|----------|-------|
| bcrypt | General password hashing | Cost factor 12+ |
| Argon2id | Preferred for new systems | Memory-hard, GPU resistant |
| scrypt | Alternative to Argon2 | Memory-hard |
| PBKDF2 | Legacy/compliance needs | 600,000+ iterations |

### 4.2 Token Management

**Tier: Critical**

#### JWT Best Practices

```python
import jwt
from datetime import datetime, timedelta

# GOOD: Short-lived tokens with proper claims
def create_access_token(user_id: str, secret: str) -> str:
    now = datetime.utcnow()
    payload = {
        "sub": user_id,
        "iat": now,
        "exp": now + timedelta(minutes=15),  # Short expiration
        "jti": generate_unique_id(),          # Unique token ID for revocation
    }
    return jwt.encode(payload, secret, algorithm="HS256")

# GOOD: Validate all claims
def verify_token(token: str, secret: str) -> dict:
    return jwt.decode(
        token,
        secret,
        algorithms=["HS256"],  # Explicit algorithm
        options={
            "require": ["sub", "iat", "exp", "jti"],
            "verify_exp": True,
        }
    )
```

#### Token Storage (Client-Side)

| Storage | Security | XSS Vulnerable | CSRF Vulnerable |
|---------|----------|----------------|-----------------|
| HttpOnly Cookie | Highest | No | Yes (mitigate with SameSite) |
| LocalStorage | Low | Yes | No |
| SessionStorage | Low | Yes | No |
| Memory | High | Limited | No |

```javascript
// GOOD: HttpOnly cookie with security flags
res.cookie("token", token, {
  httpOnly: true, // Not accessible via JavaScript
  secure: true, // HTTPS only
  sameSite: "strict", // CSRF protection
  maxAge: 15 * 60 * 1000, // 15 minutes
});
```

### 4.3 Session Security

**Tier: Required**

```python
from flask import session
import secrets

app.config.update(
    SESSION_COOKIE_SECURE=True,       # HTTPS only
    SESSION_COOKIE_HTTPONLY=True,     # No JavaScript access
    SESSION_COOKIE_SAMESITE="Lax",    # CSRF protection
    PERMANENT_SESSION_LIFETIME=1800,   # 30 minutes
)

# Regenerate session ID after authentication
def login_user(user):
    session.clear()
    session["user_id"] = user.id
    session["_fresh"] = True
    session.regenerate()  # New session ID to prevent fixation
```

### 4.4 Principle of Least Privilege

**Tier: Critical**

#### Database Access

```sql
-- BAD: Application uses admin account
GRANT ALL PRIVILEGES ON database.* TO 'app'@'%';

-- GOOD: Minimal required permissions
CREATE USER 'app_reader'@'%' IDENTIFIED BY '...';
GRANT SELECT ON database.users TO 'app_reader'@'%';
GRANT SELECT ON database.products TO 'app_reader'@'%';

CREATE USER 'app_writer'@'%' IDENTIFIED BY '...';
GRANT SELECT, INSERT, UPDATE ON database.orders TO 'app_writer'@'%';
```

#### API Keys and Service Accounts

```yaml
# AWS IAM Policy - Minimal S3 access
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource": "arn:aws:s3:::my-bucket/uploads/*"
    }
  ]
}
```

---

## 5. OWASP Top 10 Awareness

### 5.1 Injection Prevention

**Tier: Critical**

#### SQL Injection

```python
# VULNERABLE
query = f"SELECT * FROM users WHERE username = '{username}'"
cursor.execute(query)

# SECURE: Parameterized query
cursor.execute(
    "SELECT * FROM users WHERE username = %s",
    (username,)
)

# SECURE: ORM (SQLAlchemy)
user = session.query(User).filter(User.username == username).first()
```

#### Command Injection

```python
import subprocess

# VULNERABLE
os.system(f"convert {input_file} {output_file}")

# SECURE: Use list arguments, avoid shell
subprocess.run(
    ["convert", input_file, output_file],
    shell=False,
    check=True
)
```

#### NoSQL Injection

```javascript
// VULNERABLE: MongoDB
db.users.find({ username: req.body.username, password: req.body.password });
// Attack: { "username": "admin", "password": { "$ne": "" } }

// SECURE: Validate types
const username = String(req.body.username);
const password = String(req.body.password);
db.users.find({ username, password: hashPassword(password) });
```

### 5.2 Broken Authentication

**Tier: Critical**

Common vulnerabilities and mitigations:

| Vulnerability | Mitigation |
|--------------|------------|
| Credential stuffing | Rate limiting, account lockout, MFA |
| Weak passwords | Password strength requirements, breach checking |
| Session fixation | Regenerate session ID on login |
| Insecure password recovery | Time-limited tokens, verify identity |

```python
# Rate limiting login attempts
from functools import wraps
from collections import defaultdict
import time

login_attempts = defaultdict(list)

def rate_limit_login(f):
    @wraps(f)
    def decorated(username, password):
        now = time.time()
        attempts = login_attempts[username]

        # Remove attempts older than 15 minutes
        attempts[:] = [t for t in attempts if now - t < 900]

        if len(attempts) >= 5:
            raise TooManyAttemptsError("Too many login attempts")

        attempts.append(now)
        return f(username, password)
    return decorated
```

### 5.3 Sensitive Data Exposure

**Tier: Critical**

#### Data Classification

| Classification | Examples | Protection |
|---------------|----------|------------|
| Critical | Passwords, payment cards, health records | Encryption at rest and transit, strict access control |
| Confidential | PII, internal documents | Encryption, access logging |
| Internal | Non-sensitive business data | Standard access controls |
| Public | Marketing materials | No special protection |

#### Encryption Requirements

```python
# At rest: Use strong encryption
from cryptography.fernet import Fernet

key = Fernet.generate_key()  # Store securely
cipher = Fernet(key)

encrypted = cipher.encrypt(sensitive_data.encode())
decrypted = cipher.decrypt(encrypted).decode()

# In transit: Always HTTPS
# Configure in web server/load balancer
# Enforce with HSTS header
```

#### Prevent Data Leakage

```python
# BAD: Exposing sensitive data in logs
logger.info(f"User login: {user.email}, password: {password}")

# GOOD: Sanitize logs
logger.info(f"User login: {user.email}")

# BAD: Returning full objects
return jsonify(user.__dict__)

# GOOD: Explicit response schema
return jsonify({
    "id": user.id,
    "email": user.email,
    "name": user.name
    # Excludes: password_hash, ssn, etc.
})
```

### 5.4 XSS Prevention

**Tier: Critical**

#### Types of XSS

| Type | Vector | Prevention |
|------|--------|------------|
| Reflected | URL parameters | Output encoding |
| Stored | Database content | Input validation + output encoding |
| DOM-based | Client-side JavaScript | Safe DOM APIs |

#### Prevention Strategies

```javascript
// BAD: innerHTML with user content
element.innerHTML = userInput;

// GOOD: textContent for plain text
element.textContent = userInput;

// GOOD: Safe DOM APIs
const link = document.createElement("a");
link.href = sanitizeUrl(userUrl);
link.textContent = userText;
parent.appendChild(link);
```

```python
# Template auto-escaping (Jinja2)
from markupsafe import Markup, escape

# Automatic escaping in templates
# {{ user_input }} is automatically escaped

# Manual escaping when needed
safe_content = escape(user_input)

# Only mark as safe when you've validated
if is_safe_html(content):
    return Markup(content)
```

#### Content Security Policy

```http
Content-Security-Policy:
    default-src 'self';
    script-src 'self' 'nonce-abc123';
    style-src 'self' 'unsafe-inline';
    img-src 'self' data: https:;
    connect-src 'self' https://api.example.com;
    frame-ancestors 'none';
```

### 5.5 CSRF Protection

**Tier: Required**

```python
# Flask-WTF CSRF protection
from flask_wtf.csrf import CSRFProtect

csrf = CSRFProtect(app)

# In templates
<form method="post">
    {{ csrf_token() }}
    <!-- form fields -->
</form>

# For AJAX requests
fetch("/api/action", {
    method: "POST",
    headers: {
        "X-CSRFToken": document.querySelector("meta[name='csrf-token']").content
    },
    body: JSON.stringify(data)
});
```

```javascript
// SameSite cookies as defense-in-depth
res.cookie("session", sessionId, {
  sameSite: "strict", // Or 'lax' for better UX
  secure: true,
  httpOnly: true,
});
```

---

## 6. Secure Development Practices

### 6.1 Static Application Security Testing (SAST)

**Tier: Required**

#### Recommended Tools

| Language | Tools |
|----------|-------|
| Multi-language | Semgrep, SonarQube, CodeQL |
| Python | Bandit, Safety |
| JavaScript/TypeScript | ESLint security plugins, npm audit |
| Go | gosec, staticcheck |
| Java | SpotBugs, Find Security Bugs |

#### CI Integration Example

```yaml
# GitHub Actions
name: Security Scan

on: [push, pull_request]

jobs:
  sast:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run Semgrep
        uses: returntocorp/semgrep-action@v1
        with:
          config: >-
            p/security-audit
            p/secrets
            p/owasp-top-ten

      - name: Run Bandit (Python)
        run: |
          pip install bandit
          bandit -r src/ -f json -o bandit-report.json

      - name: Upload Results
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: bandit-report.json
```

### 6.2 Security-Focused Code Review

**Tier: Required**

#### Review Checklist

When reviewing code, verify:

**Authentication & Authorization**
- [ ] Authentication required for protected endpoints
- [ ] Authorization checks at function entry points
- [ ] No privilege escalation paths
- [ ] Secure session handling

**Input Handling**
- [ ] All external input validated
- [ ] Output properly encoded for context
- [ ] SQL queries parameterized
- [ ] File operations use validated paths

**Data Protection**
- [ ] Sensitive data not logged
- [ ] Secrets not hardcoded
- [ ] PII handled according to policy
- [ ] Encryption used where required

**Error Handling**
- [ ] Errors don't leak sensitive information
- [ ] Failed operations handled securely
- [ ] Rate limiting on sensitive operations

### 6.3 Threat Modeling Basics

**Tier: Recommended**

Use STRIDE to identify threats:

| Threat | Description | Example Questions |
|--------|-------------|-------------------|
| **S**poofing | Pretending to be someone else | Can users impersonate others? |
| **T**ampering | Modifying data or code | Can data be modified in transit? |
| **R**epudiation | Denying actions | Are actions logged with integrity? |
| **I**nformation Disclosure | Exposing data | What data could leak? |
| **D**enial of Service | Making system unavailable | Can resources be exhausted? |
| **E**levation of Privilege | Gaining unauthorized access | Can users gain admin access? |

#### Simple Threat Model Template

```markdown
## Feature: User File Upload

### Assets
- Uploaded files
- User data
- Server resources

### Trust Boundaries
- Client to Server (untrusted input)
- Server to Storage (internal)

### Threats
1. Malicious file upload (Tampering)
   - Risk: Code execution
   - Mitigation: File type validation, sandboxed processing

2. Path traversal (Information Disclosure)
   - Risk: Access to other users' files
   - Mitigation: Generated filenames, validated paths

3. Large file DoS (Denial of Service)
   - Risk: Storage exhaustion
   - Mitigation: File size limits, quotas
```

### 6.4 Incident Response Process

**Tier: Required**

#### When You Discover a Vulnerability

1. **Do not disclose publicly** - Report through proper channels
2. **Document the issue** - Steps to reproduce, potential impact
3. **Assess severity** - Use CVSS or internal scale
4. **Report immediately** - Contact security team
5. **Preserve evidence** - Don't modify logs or systems

#### Incident Severity Levels

| Level | Description | Response |
|-------|-------------|----------|
| P1 - Critical | Active breach, data exposure | Immediate. All hands. |
| P2 - High | Exploitable vulnerability in production | Same day response |
| P3 - Medium | Vulnerability with limited exposure | Within one week |
| P4 - Low | Minor security improvements | Scheduled maintenance |

---

## Security Review Checklist

Use this checklist for security-focused code reviews:

### Pre-Merge Security Checklist

```markdown
## Security Review

### Secrets & Configuration
- [ ] No hardcoded secrets, API keys, or passwords
- [ ] Environment variables used for sensitive configuration
- [ ] No secrets in logs or error messages

### Input Validation
- [ ] All external input validated and sanitized
- [ ] Allowlists used instead of denylists where applicable
- [ ] File uploads validated (type, size, content)

### Authentication & Authorization
- [ ] Endpoints require appropriate authentication
- [ ] Authorization checked before operations
- [ ] Passwords hashed with bcrypt/Argon2

### Data Protection
- [ ] Sensitive data encrypted at rest and in transit
- [ ] PII handling follows data retention policies
- [ ] API responses don't leak sensitive fields

### Injection Prevention
- [ ] SQL queries use parameterized statements
- [ ] No dynamic shell command construction
- [ ] HTML output properly escaped

### Dependencies
- [ ] New dependencies vetted for security
- [ ] No known vulnerable dependencies
- [ ] Lock files updated

### Logging & Monitoring
- [ ] Security-relevant events logged
- [ ] No sensitive data in logs
- [ ] Error messages don't reveal system details
```

---

## Tool Recommendations

### Essential Security Tools

| Category | Tool | Tier |
|----------|------|------|
| Secret Scanning | GitLeaks, TruffleHog | Critical |
| SAST | Semgrep, CodeQL | Required |
| Dependency Scanning | Dependabot, Snyk | Critical |
| Container Scanning | Trivy, Grype | Required |
| DAST | OWASP ZAP, Burp Suite | Recommended |

### IDE Security Plugins

| IDE | Plugins |
|-----|---------|
| VS Code | Snyk, SonarLint, GitLens |
| JetBrains | Snyk, SonarLint, Security Analyzer |
| Vim/Neovim | ALE with security linters |

### Pre-commit Hooks

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/gitleaks/gitleaks
    rev: v8.18.0
    hooks:
      - id: gitleaks

  - repo: https://github.com/PyCQA/bandit
    rev: 1.7.5
    hooks:
      - id: bandit
        args: ["-c", "pyproject.toml"]

  - repo: https://github.com/returntocorp/semgrep
    rev: v1.45.0
    hooks:
      - id: semgrep
        args: ["--config", "auto"]
```

---

## Quick Reference

### Secure Defaults Checklist

| Setting | Secure Default |
|---------|---------------|
| Cookies | `HttpOnly`, `Secure`, `SameSite=Lax` |
| CORS | Specific origins, not `*` |
| Headers | HSTS, CSP, X-Content-Type-Options |
| Passwords | Bcrypt/Argon2, cost factor 12+ |
| Tokens | Short-lived, cryptographically random |
| Logging | No secrets, no full PII |
| Errors | Generic messages to users |
| Files | Validated paths, restricted permissions |

### Security Headers

```http
Strict-Transport-Security: max-age=31536000; includeSubDomains
Content-Security-Policy: default-src 'self'
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: geolocation=(), microphone=(), camera=()
```

---

## Further Reading

- [OWASP Top 10](https://owasp.org/Top10/)
- [OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org/)
- [CWE/SANS Top 25](https://cwe.mitre.org/top25/)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)
