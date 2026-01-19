---
description: Security principles - secret management, input validation, OWASP awareness
---

# Security Principles

> Universal security principles that apply across all platforms. For language-specific implementations, see your platform's security rules.

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## 1. Secret Management

### SEC-001: Never Commit Secrets :red_circle:

**Rationale**: Secrets in version control are permanently exposed. Anyone with repo access (including future access) can extract them.

Secrets include:
- API keys and tokens
- Passwords and connection strings
- Private keys and certificates
- Service account credentials
- Encryption keys

**Requirements**:
- Load secrets from environment variables or secret managers
- Never hardcode credentials in source code
- Use `.env` files for local development (never committed)
- Use CI/CD secrets for pipelines
- Use secret managers (Vault, AWS Secrets Manager) for production

**Gitignore patterns**:
```gitignore
.env
.env.*
!.env.example
*.pem
*.key
*.p12
*-credentials.json
secrets.*
```

### SEC-002: Use Secret Managers for Production :red_circle:

**Rationale**: Environment variables alone are insufficient for production - they may be logged, exposed in process listings, or leaked through debugging.

| Environment | Approach |
|-------------|----------|
| Local Development | `.env` files (never committed) |
| CI/CD | Pipeline secrets (GitHub Secrets, GitLab CI Variables) |
| Production | Secret managers (HashiCorp Vault, AWS Secrets Manager, GCP Secret Manager) |

### SEC-003: Respond to Leaked Secrets Immediately :red_circle:

**Rationale**: Leaked secrets must be treated as compromised regardless of exposure duration.

**Response procedure**:
1. **Revoke immediately** - Generate new credentials before cleanup
2. **Remove from history** - Use BFG Repo-Cleaner or git filter-branch
3. **Force push** - Update remote after cleaning
4. **Audit access** - Check logs for unauthorized use
5. **Report** - Notify security team per incident response

**Note**: The secret remains compromised even after removal - anyone who cloned the repo may have it.

---

## 2. Input Validation

### SEC-004: Validate All External Input :red_circle:

**Rationale**: Unvalidated input is the root cause of most security vulnerabilities (injection, XSS, buffer overflows).

External input includes:
- HTTP requests (headers, body, query params)
- File uploads
- Database reads (data may have been corrupted)
- API responses from external services
- Command-line arguments
- Environment variables
- Message queues

**Validation strategy**:
- Define explicit schemas for expected input
- Validate type, format, length, and range
- Reject invalid input early (fail fast)
- Use validation libraries appropriate to your platform

### SEC-005: Use Allowlists Over Denylists :yellow_circle:

**Rationale**: Denylists are incomplete by definition - you can't anticipate all malicious inputs.

**Principle**:
- Define what IS allowed (allowlist)
- Reject everything else by default
- Example: For file uploads, explicitly list allowed extensions rather than blocking known-bad ones

### SEC-006: Sanitize for Context :red_circle:

**Rationale**: Different output contexts require different sanitization.

| Context | Sanitization |
|---------|-------------|
| HTML output | HTML entity encoding |
| SQL queries | Parameterized queries (never string concatenation) |
| Shell commands | Avoid shell; use direct execution with argument arrays |
| File paths | Validate against allowlist, resolve canonical path |
| URLs | Parse and validate scheme and host |

---

## 3. Dependency Security

### SEC-007: Enable Dependency Scanning :red_circle:

**Rationale**: Third-party dependencies are a major attack vector - most codebases have more dependency code than first-party code.

**Requirements**:
- Enable automated scanning in CI/CD (Dependabot, Snyk, etc.)
- Fail builds on critical/high severity vulnerabilities
- Block merge until vulnerabilities addressed
- Review security advisories for dependencies

### SEC-008: Vulnerability Response Policy :red_circle:

**Rationale**: Clear SLAs ensure vulnerabilities are addressed promptly.

| Severity | CVSS | Response Time |
|----------|------|---------------|
| Critical | 9.0-10.0 | 24 hours |
| High | 7.0-8.9 | 7 days |
| Medium | 4.0-6.9 | 30 days |
| Low | 0.1-3.9 | 90 days |

### SEC-009: Lock File Requirements :red_circle:

**Rationale**: Lock files ensure reproducible builds and prevent supply chain attacks through dependency confusion.

**Requirements**:
- Always commit lock files
- Install from lock files in CI/CD (not resolving fresh)
- Review lock file changes in PRs

---

## 4. Authentication & Authorization

### SEC-010: Secure Password Handling :red_circle:

**Rationale**: Password breaches have cascading effects - users reuse passwords across services.

**Requirements**:
- Never store plaintext passwords
- Use memory-hard hashing (Argon2id, bcrypt, scrypt)
- Never use MD5, SHA1, or unsalted hashes
- Use appropriate cost factors (bcrypt: 12+, Argon2: tune for ~1s)

| Algorithm | Use Case |
|-----------|----------|
| Argon2id | Preferred for new systems |
| bcrypt | General password hashing |
| scrypt | Alternative to Argon2 |
| PBKDF2 | Legacy/compliance (600k+ iterations) |

### SEC-011: Token Management :red_circle:

**Rationale**: Tokens are bearer credentials - possession equals access.

**Requirements**:
- Short-lived access tokens (15 minutes typical)
- Refresh tokens for session extension
- Validate all claims on verification
- Use HttpOnly, Secure, SameSite cookies for web
- Never store tokens in localStorage (XSS vulnerable)

### SEC-012: Principle of Least Privilege :red_circle:

**Rationale**: Limiting permissions reduces blast radius of compromises.

**Requirements**:
- Database accounts: minimal required permissions per service
- API keys: scoped to specific operations
- Service accounts: narrowest IAM policies possible
- Review and audit permissions regularly

---

## 5. OWASP Top 10 Awareness

### SEC-013: Injection Prevention :red_circle:

**Rationale**: Injection flaws (SQL, NoSQL, Command, LDAP) remain the most critical web vulnerabilities.

**Principles**:
- Use parameterized queries for all database operations
- Use ORMs that enforce parameterization
- Avoid shell execution; use direct command execution
- Validate and sanitize all inputs

### SEC-014: Broken Authentication Prevention :red_circle:

**Rationale**: Authentication bypasses grant unauthorized access to all user data.

**Mitigations**:
- Rate limiting on login attempts
- Account lockout after failed attempts
- Multi-factor authentication for sensitive operations
- Session ID regeneration on login
- Secure password recovery flows

### SEC-015: Sensitive Data Protection :red_circle:

**Rationale**: Data breaches cause regulatory, legal, and reputational damage.

| Classification | Protection |
|---------------|------------|
| Critical (passwords, payment cards) | Encryption at rest and transit, strict access |
| Confidential (PII) | Encryption, access logging |
| Internal | Standard access controls |

**Requirements**:
- Encrypt sensitive data at rest
- Use TLS for all data in transit
- Never log sensitive data
- Explicit response schemas (don't expose full objects)

### SEC-016: XSS Prevention :red_circle:

**Rationale**: XSS enables session hijacking, credential theft, and defacement.

**Prevention**:
- Auto-escape all template output
- Use Content Security Policy headers
- Use textContent/innerText instead of innerHTML
- Sanitize rich text with allowlist-based libraries

### SEC-017: CSRF Protection :yellow_circle:

**Rationale**: CSRF exploits trust a site has in a user's browser.

**Prevention**:
- Use anti-CSRF tokens for state-changing operations
- Use SameSite cookie attribute (Lax or Strict)
- Verify Origin/Referer headers

---

## 6. Secure Development Practices

### SEC-018: Static Application Security Testing (SAST) :yellow_circle:

**Rationale**: Automated scanning catches common vulnerabilities before review.

**Requirements**:
- Run SAST on all PRs
- Triage and address findings before merge
- Track false positive suppressions
- Use platform-appropriate tools (see platform-specific rules)

### SEC-019: Security-Focused Code Review :yellow_circle:

**Rationale**: Security issues are easier to fix before merge than after deployment.

**Review checklist**:
- [ ] Authentication required for protected endpoints
- [ ] Authorization checks at function entry points
- [ ] All external input validated
- [ ] SQL queries parameterized
- [ ] Sensitive data not logged
- [ ] Secrets not hardcoded
- [ ] Errors don't leak sensitive information

### SEC-020: Incident Response Process :yellow_circle:

**Rationale**: Prepared response minimizes breach impact.

**When discovering a vulnerability**:
1. Do not disclose publicly
2. Document steps to reproduce
3. Assess severity (use CVSS)
4. Report to security team
5. Preserve evidence

| Severity | Response |
|----------|----------|
| P1 - Critical | Immediate, all hands |
| P2 - High | Same day |
| P3 - Medium | Within one week |
| P4 - Low | Scheduled maintenance |

---

## Quick Reference

| ID | Rule | Tier |
|----|------|------|
| SEC-001 | Never Commit Secrets | Critical |
| SEC-002 | Use Secret Managers for Production | Critical |
| SEC-003 | Respond to Leaked Secrets Immediately | Critical |
| SEC-004 | Validate All External Input | Critical |
| SEC-005 | Use Allowlists Over Denylists | Required |
| SEC-006 | Sanitize for Context | Critical |
| SEC-007 | Enable Dependency Scanning | Critical |
| SEC-008 | Vulnerability Response Policy | Critical |
| SEC-009 | Lock File Requirements | Critical |
| SEC-010 | Secure Password Handling | Critical |
| SEC-011 | Token Management | Critical |
| SEC-012 | Principle of Least Privilege | Critical |
| SEC-013 | Injection Prevention | Critical |
| SEC-014 | Broken Authentication Prevention | Critical |
| SEC-015 | Sensitive Data Protection | Critical |
| SEC-016 | XSS Prevention | Critical |
| SEC-017 | CSRF Protection | Required |
| SEC-018 | Static Application Security Testing | Required |
| SEC-019 | Security-Focused Code Review | Required |
| SEC-020 | Incident Response Process | Required |

---

## Platform-Specific Implementations

For language-specific security implementations, see:
- **Go**: [coding-standard-go-cloud/rules/security.md](../../coding-standard-go-cloud/rules/security.md)
- **C/Zephyr**: [coding-standard-c-zephyr/rules/security.md](../../coding-standard-c-zephyr/rules/security.md)

---

## References

- [OWASP Top 10](https://owasp.org/Top10/)
- [OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org/)
- [CWE/SANS Top 25](https://cwe.mitre.org/top25/)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)
