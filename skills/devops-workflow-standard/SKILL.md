---
name: devops-workflow-standard
description: Universal DevOps practices for git, code review, CI/CD, and security
---

# DevOps Standard

> **Version**: 1.0.0 | **Status**: Active
> **Scope**: Platform-agnostic DevOps practices

This skill establishes universal DevOps practices that apply across all platforms and languages. Platform-specific implementations (tooling, security code, CI/CD configurations) are defined in the respective coding standard skills (Go, Zephyr, etc.).

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [Git Workflow](rules/git-workflow.md) | GIT-* | Branching, commits, PRs, releases |
| [Code Review](rules/code-review.md) | REV-* | Review process, criteria, feedback |
| [Documentation](rules/documentation.md) | DOC-* | READMEs, API docs, ADRs, runbooks |
| [Security Principles](rules/security-principles.md) | SEC-* | Secret management, input validation, OWASP |
| [CI/CD Principles](rules/cicd-principles.md) | CICD-* | Pipeline architecture, quality gates, deployment |

### Tooling

| Tool | File |
|------|------|
| [Pre-commit](tooling/pre-commit.md) | Git hooks configuration |
| [Commitlint](tooling/commitlint.md) | Commit message validation |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table with tiers |
| [Code Review Checklist](reference/code-review-checklist.md) | Review checklist by category |

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Platform-Specific Extensions

This skill covers universal principles. For platform-specific implementations, see:

| Platform | Skill | Security | CI/CD |
|----------|-------|----------|-------|
| **Go** | [coding-standard-go-cloud](../coding-standard-go-cloud/SKILL.md) | [security.md](../coding-standard-go-cloud/rules/security.md) | [cicd.md](../coding-standard-go-cloud/rules/cicd.md) |
| **C/Zephyr** | [coding-standard-c-zephyr](../coding-standard-c-zephyr/SKILL.md) | [security.md](../coding-standard-c-zephyr/rules/security.md) | [cicd.md](../coding-standard-c-zephyr/rules/cicd.md) |

---

## Critical Rules (Always Apply)

These rules are non-negotiable and apply to all platforms.

### GIT-001: Branch Names Must Follow Convention :red_circle:

Consistent naming enables automation and traceability.

```bash
# Correct
feature/PROJ-123-add-user-authentication
fix/PROJ-456-resolve-null-pointer
hotfix/PROJ-789-patch-xss-vulnerability

# Incorrect
feature/add-auth                    # Missing ticket number
PROJ-123-new-feature                # Missing type prefix
```

### GIT-006: All Changes Must Go Through Pull Requests :red_circle:

PRs provide checkpoints for code review, automated testing, and documentation.

### REV-006: Maintain Constructive Tone :red_circle:

Critique the code, not the person. Assume good intent.

### REV-007: Verify Correctness :red_circle:

Logic handles all requirements, edge cases considered, errors handled.

### DOC-001: Every Repository Must Have a README :red_circle:

A README is the entry point. Developers should clone and run within 15 minutes.

### DOC-003: All Public APIs Must Be Documented :red_circle:

Undocumented APIs are unusable and lead to incorrect integrations.

### SEC-001: Never Commit Secrets :red_circle:

Secrets include API keys, passwords, tokens, private keys, connection strings.

### SEC-002: Validate All External Input :red_circle:

External input includes HTTP requests, file uploads, database reads, API responses.

### CICD-BUILD-001: Reproducible Builds :red_circle:

Builds must produce identical artifacts given the same inputs.

### CICD-SEC-001: Dependency Vulnerability Scanning :red_circle:

Third-party dependencies are a major attack vector.

### CICD-DEPLOY-001: Environment Progression :red_circle:

Staged deployments (dev -> staging -> production) reduce incident risk.

---

## Quick Rule Lookup

| ID | Rule | Tier |
|----|------|------|
| GIT-001 | Branch names must follow naming convention | Critical |
| GIT-003 | Subject line must follow format requirements | Critical |
| GIT-006 | All changes to protected branches must go through PRs | Critical |
| GIT-008 | PRs require minimum number of approvals | Critical |
| GIT-009 | All CI checks must pass before merge | Critical |
| GIT-014 | Main branch must be protected | Critical |
| REV-006 | Maintain constructive tone | Critical |
| REV-007 | Verify correctness | Critical |
| DOC-001 | Every repository must have a README | Critical |
| DOC-003 | All public APIs must be documented | Critical |
| DOC-009 | All public APIs must have documentation comments | Critical |
| DOC-012 | Document deployment procedures | Critical |
| SEC-001 | Never commit secrets | Critical |
| SEC-002 | Validate all external input | Critical |
| CICD-BUILD-001 | Reproducible builds | Critical |
| CICD-TEST-002 | Test timeouts | Critical |
| CICD-ANALYZE-001 | Linting enforcement | Critical |
| CICD-SEC-001 | Dependency vulnerability scanning | Critical |
| CICD-SEC-002 | Secret scanning | Critical |
| CICD-DEPLOY-001 | Environment progression | Critical |
| CICD-DEPLOY-003 | Rollback procedures | Critical |
| CICD-MON-001 | Pipeline failure notifications | Critical |

---

## References

- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
- [Semantic Versioning](https://semver.org/)
- [OWASP Top 10](https://owasp.org/Top10/)
- [The Twelve-Factor App](https://12factor.net/)
