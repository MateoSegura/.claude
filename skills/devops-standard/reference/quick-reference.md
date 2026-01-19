# Quick Reference - All DevOps Rules

Complete table of all rules with IDs, descriptions, and tiers.

## Rule Tiers

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Git Workflow (GIT-*)

| ID | Rule | Tier |
|----|------|------|
| GIT-001 | Branch names must follow naming convention | Critical |
| GIT-002 | Feature branches must be short-lived | Required |
| GIT-003 | Subject line must follow format requirements | Critical |
| GIT-004 | Commit body should explain what and why | Recommended |
| GIT-005 | Use footers for metadata and references | Recommended |
| GIT-006 | All changes to protected branches must go through PRs | Critical |
| GIT-007 | PR titles must match commit message format | Required |
| GIT-008 | PRs require minimum number of approvals | Critical |
| GIT-009 | All CI checks must pass before merge | Critical |
| GIT-010 | Use squash merge for feature branches | Required |
| GIT-011 | Delete branches after merge | Required |
| GIT-012 | Address stale branches within 30 days | Required |
| GIT-013 | Rebase feature branches before PR | Recommended |
| GIT-014 | Main branch must be protected | Critical |
| GIT-015 | Release branches must be protected | Critical |
| GIT-016 | Hotfixes follow expedited process | Critical |
| GIT-017 | Release branches stabilize code for deployment | Required |
| GIT-018 | Use semantic versioning for releases | Required |

---

## Code Review (REV-*)

| ID | Rule | Tier |
|----|------|------|
| REV-001 | Complete self-review before requesting review | Required |
| REV-002 | Provide complete PR description | Required |
| REV-003 | Respond to all review comments | Required |
| REV-004 | Respond within 24 business hours | Required |
| REV-005 | Provide thorough review | Required |
| REV-006 | Maintain constructive tone | Critical |
| REV-007 | Verify correctness | Critical |
| REV-008 | Evaluate design appropriateness | Required |
| REV-009 | Challenge unnecessary complexity | Required |
| REV-010 | Require adequate test coverage | Required |
| REV-011 | Ensure clear naming | Required |
| REV-012 | Review comment quality | Recommended |
| REV-013 | Enforce style consistency | Required |
| REV-014 | Verify documentation updates | Required |
| REV-015 | Limit review cycles | Recommended |
| REV-016 | Resolve disagreements constructively | Required |
| REV-017 | Provide actionable feedback | Required |
| REV-018 | Keep PRs small | Required |
| REV-019 | Split large changes appropriately | Required |

---

## Documentation (DOC-*)

| ID | Rule | Tier |
|----|------|------|
| DOC-001 | Every repository must have a README | Critical |
| DOC-002 | README must enable quick start | Required |
| DOC-003 | All public APIs must be documented | Critical |
| DOC-004 | Use OpenAPI 3.0+ specification | Required |
| DOC-005 | Document all error codes | Required |
| DOC-006 | Document API versioning policy | Required |
| DOC-007 | Comment why, not what | Required |
| DOC-008 | Required comment scenarios | Required |
| DOC-009 | All public APIs must have documentation comments | Critical |
| DOC-010 | Write ADR for significant decisions | Required |
| DOC-011 | Maintain ADR status | Required |
| DOC-012 | Document deployment procedures | Critical |
| DOC-013 | Provide troubleshooting documentation | Required |
| DOC-014 | Document on-call procedures | Required |
| DOC-015 | Update documentation with code changes | Required |
| DOC-016 | Review documentation in PRs | Required |
| DOC-017 | Document deprecations clearly | Required |
| DOC-018 | Conduct regular documentation audits | Recommended |

---

## Security Principles (SEC-*)

| ID | Rule | Tier |
|----|------|------|
| SEC-001 | Never commit secrets | Critical |
| SEC-002 | Use secret managers for production | Critical |
| SEC-003 | Respond to leaked secrets immediately | Critical |
| SEC-004 | Validate all external input | Critical |
| SEC-005 | Use allowlists over denylists | Required |
| SEC-006 | Sanitize for context | Critical |
| SEC-007 | Enable dependency scanning | Critical |
| SEC-008 | Vulnerability response policy | Critical |
| SEC-009 | Lock file requirements | Critical |
| SEC-010 | Secure password handling | Critical |
| SEC-011 | Token management | Critical |
| SEC-012 | Principle of least privilege | Critical |
| SEC-013 | Injection prevention | Critical |
| SEC-014 | Broken authentication prevention | Critical |
| SEC-015 | Sensitive data protection | Critical |
| SEC-016 | XSS prevention | Critical |
| SEC-017 | CSRF protection | Required |
| SEC-018 | Static application security testing | Required |
| SEC-019 | Security-focused code review | Required |
| SEC-020 | Incident response process | Required |

---

## CI/CD Principles (CICD-*)

| ID | Rule | Tier |
|----|------|------|
| CICD-BUILD-001 | Reproducible builds | Critical |
| CICD-BUILD-002 | Build artifacts | Required |
| CICD-BUILD-003 | Version tagging | Required |
| CICD-BUILD-004 | Build caching | Recommended |
| CICD-TEST-001 | Test parallelization | Required |
| CICD-TEST-002 | Test timeouts | Critical |
| CICD-TEST-003 | Flaky test handling | Required |
| CICD-TEST-004 | Test result reporting | Required |
| CICD-ANALYZE-001 | Linting enforcement | Critical |
| CICD-ANALYZE-002 | Code coverage requirements | Required |
| CICD-ANALYZE-003 | Code formatting | Critical |
| CICD-SEC-001 | Dependency vulnerability scanning | Critical |
| CICD-SEC-002 | Secret scanning | Critical |
| CICD-SEC-003 | Static application security testing | Required |
| CICD-DEPLOY-001 | Environment progression | Critical |
| CICD-DEPLOY-002 | Deployment strategies | Required |
| CICD-DEPLOY-003 | Rollback procedures | Critical |
| CICD-DEPLOY-004 | Feature flags | Recommended |
| CICD-MON-001 | Pipeline failure notifications | Critical |
| CICD-MON-002 | Deployment notifications | Required |
| CICD-MON-003 | Performance tracking | Recommended |

---

## Summary by Tier

### Critical Rules (45 total)

**Git**: GIT-001, GIT-003, GIT-006, GIT-008, GIT-009, GIT-014, GIT-015, GIT-016

**Review**: REV-006, REV-007

**Documentation**: DOC-001, DOC-003, DOC-009, DOC-012

**Security**: SEC-001 through SEC-016 (excluding SEC-005, SEC-017 through SEC-020)

**CI/CD**: CICD-BUILD-001, CICD-TEST-002, CICD-ANALYZE-001, CICD-ANALYZE-003, CICD-SEC-001, CICD-SEC-002, CICD-DEPLOY-001, CICD-DEPLOY-003, CICD-MON-001

### Required Rules (43 total)

**Git**: GIT-002, GIT-007, GIT-010, GIT-011, GIT-012, GIT-017, GIT-018

**Review**: REV-001 through REV-005, REV-008 through REV-011, REV-013, REV-014, REV-016 through REV-019

**Documentation**: DOC-002, DOC-004 through DOC-008, DOC-010, DOC-011, DOC-013 through DOC-017

**Security**: SEC-005, SEC-017 through SEC-020

**CI/CD**: CICD-BUILD-002, CICD-BUILD-003, CICD-TEST-001, CICD-TEST-003, CICD-TEST-004, CICD-ANALYZE-002, CICD-SEC-003, CICD-DEPLOY-002, CICD-MON-002

### Recommended Rules (7 total)

**Git**: GIT-004, GIT-005, GIT-013

**Review**: REV-012, REV-015

**Documentation**: DOC-018

**CI/CD**: CICD-BUILD-004, CICD-DEPLOY-004, CICD-MON-003
