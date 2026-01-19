---
description: CI/CD principles - pipeline architecture, quality gates, deployment strategies
---

# CI/CD Principles

> Universal CI/CD principles that apply across all platforms. For language-specific configurations, see your platform's CI/CD rules.

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## 1. Pipeline Architecture

### Pipeline Stages Overview

Every pipeline MUST implement these core stages in order:

```
BUILD ──> TEST ──> ANALYZE ──> SECURITY ──> DEPLOY
  │         │         │           │           │
  ▼         ▼         ▼           ▼           ▼
Compile   Unit     Static     Vuln Scan   Environment
Deps      Integ    Coverage   Secret Scan Promotion
Artifacts E2E      Lint       SAST/DAST   Rollback
```

### Stage Definitions

| Stage | Purpose | Blocking | Timeout |
|-------|---------|----------|---------|
| **Build** | Compile code, resolve dependencies, create artifacts | Yes | 10 min |
| **Test** | Execute automated tests | Yes | 20 min |
| **Analyze** | Static analysis, linting, coverage | Yes | 10 min |
| **Security** | Vulnerability and secret scanning | Yes | 15 min |
| **Deploy** | Promote artifacts to environments | Yes | 30 min |

---

## 2. Build Stage

### CICD-BUILD-001: Reproducible Builds :red_circle:

**Rationale**: Builds must produce identical artifacts given the same inputs to ensure reliability and auditability.

**Requirements**:
- Pin all dependency versions exactly (no ranges)
- Use lock files (go.sum, package-lock.json, Cargo.lock, etc.)
- Specify exact base images with SHA256 digests for containers
- Document and version build tools

```yaml
# Correct: Pinned versions
FROM node:20.11.0-alpine@sha256:abc123...
RUN npm ci  # Uses lock file

# Incorrect: Floating tags
FROM node:latest
RUN npm install  # May resolve different versions
```

### CICD-BUILD-002: Build Artifacts :yellow_circle:

**Rationale**: Artifacts must be traceable and identifiable throughout the deployment pipeline.

**Requirements**:
- Generate unique artifact identifiers (commit SHA + build number)
- Store artifacts in versioned registry
- Include build metadata (git commit, timestamp, builder)
- Retain artifacts for minimum 90 days

**Artifact naming format**:
```
{app}-{version}-{commit_sha:8}-{build_number}
Example: api-server-1.2.3-abc12345-42
```

### CICD-BUILD-003: Version Tagging :yellow_circle:

**Rationale**: Consistent versioning enables traceability and rollback capabilities.

**Requirements**:
- Use semantic versioning (MAJOR.MINOR.PATCH)
- Tag releases in git with `v` prefix (v1.2.3)
- Pre-release versions use suffixes (v1.2.3-rc.1, v1.2.3-beta.2)
- Never reuse or overwrite version tags

### CICD-BUILD-004: Build Caching :green_circle:

**Rationale**: Effective caching significantly reduces build times and CI costs.

**Strategy**:
- Cache dependency downloads
- Cache build outputs where safe
- Use content-addressable cache keys (hash of lock files)
- Invalidate caches on tool version changes

---

## 3. Test Stage

### CICD-TEST-001: Test Parallelization :yellow_circle:

**Rationale**: Parallel test execution reduces feedback time and improves developer productivity.

**Requirements**:
- Run independent test suites in parallel
- Use test sharding for large test suites
- Balance shard sizes for consistent timing
- Report aggregated results

### CICD-TEST-002: Test Timeouts :red_circle:

**Rationale**: Unbounded tests can hang pipelines and waste resources.

| Test Type | Default Timeout | Maximum Allowed |
|-----------|-----------------|-----------------|
| Unit test (per test) | 5 seconds | 30 seconds |
| Integration test (per test) | 30 seconds | 2 minutes |
| E2E test (per test) | 2 minutes | 10 minutes |
| Test suite (total) | 10 minutes | 30 minutes |

### CICD-TEST-003: Flaky Test Handling :yellow_circle:

**Rationale**: Flaky tests erode trust in CI and slow development velocity.

**Policy**:
1. **Detection**: Track test pass rates over time
2. **Quarantine**: Move tests with <95% pass rate to quarantine suite
3. **Fix or Remove**: Quarantined tests must be fixed within 2 weeks or removed
4. **No Retry Masking**: Do not use automatic retries to hide flakiness

### CICD-TEST-004: Test Result Reporting :yellow_circle:

**Rationale**: Clear test reporting enables rapid diagnosis and trend analysis.

**Requirements**:
- Output results in standard format (JUnit XML, JSON)
- Include failure messages and stack traces
- Report test duration for performance tracking
- Upload results to CI system for visibility

---

## 4. Static Analysis Stage

### CICD-ANALYZE-001: Linting Enforcement :red_circle:

**Rationale**: Automated linting catches style violations and common errors before review.

**Requirements**:
- Run language-appropriate linters on all changed files
- Fail pipeline on linter errors
- Configure linters via committed config files
- No inline disables without justification comment

### CICD-ANALYZE-002: Code Coverage Requirements :yellow_circle:

**Rationale**: Coverage metrics ensure tests exercise critical code paths.

| Metric | Minimum | Target |
|--------|---------|--------|
| Line coverage | 70% | 85% |
| Branch coverage | 60% | 75% |
| New code coverage | 80% | 90% |

### CICD-ANALYZE-003: Code Formatting :red_circle:

**Rationale**: Consistent formatting eliminates style debates and reduces diff noise.

**Requirements**:
- Run formatter check (not auto-fix) in CI
- Fail pipeline on formatting violations
- Use committed formatter configuration

---

## 5. Security Stage

### CICD-SEC-001: Dependency Vulnerability Scanning :red_circle:

**Rationale**: Third-party dependencies are a major attack vector.

**Requirements**:
- Scan all dependencies on every build
- Fail on critical/high severity vulnerabilities
- Block merge until vulnerabilities addressed
- Allow temporary exceptions with documented timeline

### CICD-SEC-002: Secret Scanning :red_circle:

**Rationale**: Secrets in code lead to security breaches.

**Requirements**:
- Scan all commits for secrets before merge
- Block commits containing detected secrets
- Scan git history for existing leaks (periodic)
- Rotate any exposed secrets immediately

### CICD-SEC-003: Static Application Security Testing (SAST) :yellow_circle:

**Rationale**: SAST identifies security vulnerabilities in source code.

**Requirements**:
- Run SAST on all PRs
- Review and triage findings before merge
- Track false positive suppressions
- Address findings within SLA

| Severity | Resolution SLA |
|----------|---------------|
| Critical | 24 hours |
| High | 7 days |
| Medium | 30 days |
| Low | 90 days |

---

## 6. Quality Gates

### Pre-Merge Quality Gates

All PRs MUST pass these gates before merge:

| Gate | Criteria | Blocking |
|------|----------|----------|
| **Build** | Successful compilation | Yes |
| **Unit Tests** | 100% pass rate | Yes |
| **Integration Tests** | 100% pass rate | Yes |
| **Coverage** | >= 70% lines, >= 80% new code | Yes |
| **Linting** | Zero errors | Yes |
| **Formatting** | Zero violations | Yes |
| **Security Scan** | No critical/high vulnerabilities | Yes |
| **Secret Scan** | No secrets detected | Yes |
| **Code Review** | >= 1 approval | Yes |

---

## 7. Deployment

### CICD-DEPLOY-001: Environment Progression :red_circle:

**Rationale**: Staged deployments reduce production incident risk.

```
DEV ────> STAGING ────> UAT ────> PROD
 │           │           │          │
Auto      Auto        Manual     Manual
on merge  on success  approval   approval
```

| Environment | Trigger | Approval | Testing |
|-------------|---------|----------|---------|
| Development | Auto on PR merge | None | Automated |
| Staging | Auto on dev success | None | Automated + Manual QA |
| UAT | Manual | Product Owner | Manual acceptance |
| Production | Manual | Release Manager + Ops | Smoke tests |

### CICD-DEPLOY-002: Deployment Strategies :yellow_circle:

**Rationale**: Safe deployment strategies minimize user impact during releases.

**Blue-Green Deployment** (Default for stateless services):
1. Deploy to green environment
2. Run smoke tests
3. Switch traffic
4. Wait for stability
5. Teardown blue

**Canary Deployment** (For high-traffic services):
1. Deploy to 5% of traffic
2. Monitor for 15 minutes
3. Increment to 25%, 50%
4. Full rollout to 100%

**Rolling Deployment** (For stateful services):
- Max unavailable: 25%
- Max surge: 25%
- Min ready seconds: 30

### CICD-DEPLOY-003: Rollback Procedures :red_circle:

**Rationale**: Fast rollback capability minimizes incident impact.

**Requirements**:
- All deployments MUST support instant rollback
- Rollback MUST NOT require a new build
- Previous version artifacts retained for 30 days minimum
- Database migrations MUST be backwards-compatible

**Automatic rollback triggers**:
- Health check failure
- Error rate > 5%
- Latency p99 > 2x baseline

### CICD-DEPLOY-004: Feature Flags :green_circle:

**Rationale**: Feature flags enable safe deployment of incomplete features and rapid incident response.

**Guidelines**:
- Use feature flags for all user-facing changes
- Flags should have owners and expiration dates
- Remove flags within 30 days of full rollout
- Emergency kill switches for critical features

---

## 8. Monitoring

### CICD-MON-001: Pipeline Failure Notifications :red_circle:

**Rationale**: Fast notification of failures enables rapid response.

**Requirements**:
- Notify on all pipeline failures
- Include failure reason and link to logs
- Route to appropriate channel (PR author, team)
- Escalate on repeated failures

### CICD-MON-002: Deployment Notifications :yellow_circle:

**Rationale**: Deployment visibility enables coordination and incident correlation.

**Events to notify**:
- Deployment started
- Deployment succeeded
- Deployment failed
- Rollback initiated
- Rollback completed

### CICD-MON-003: Performance Tracking :green_circle:

**Rationale**: Pipeline performance metrics identify optimization opportunities.

**Metrics to track**:
- Build duration (by stage)
- Test duration (by suite)
- Queue time (time waiting for runner)
- Success rate (by pipeline, stage)
- Flaky test rate
- Cache hit rate

---

## Quick Reference

| ID | Rule | Tier |
|----|------|------|
| CICD-BUILD-001 | Reproducible Builds | Critical |
| CICD-BUILD-002 | Build Artifacts | Required |
| CICD-BUILD-003 | Version Tagging | Required |
| CICD-BUILD-004 | Build Caching | Recommended |
| CICD-TEST-001 | Test Parallelization | Required |
| CICD-TEST-002 | Test Timeouts | Critical |
| CICD-TEST-003 | Flaky Test Handling | Required |
| CICD-TEST-004 | Test Result Reporting | Required |
| CICD-ANALYZE-001 | Linting Enforcement | Critical |
| CICD-ANALYZE-002 | Code Coverage Requirements | Required |
| CICD-ANALYZE-003 | Code Formatting | Critical |
| CICD-SEC-001 | Dependency Vulnerability Scanning | Critical |
| CICD-SEC-002 | Secret Scanning | Critical |
| CICD-SEC-003 | Static Application Security Testing | Required |
| CICD-DEPLOY-001 | Environment Progression | Critical |
| CICD-DEPLOY-002 | Deployment Strategies | Required |
| CICD-DEPLOY-003 | Rollback Procedures | Critical |
| CICD-DEPLOY-004 | Feature Flags | Recommended |
| CICD-MON-001 | Pipeline Failure Notifications | Critical |
| CICD-MON-002 | Deployment Notifications | Required |
| CICD-MON-003 | Performance Tracking | Recommended |

---

## Platform-Specific Implementations

For language-specific CI/CD configurations, see:
- **Go**: [coding-standard-go-cloud/rules/cicd.md](../../coding-standard-go-cloud/rules/cicd.md)
- **C/Zephyr**: [coding-standard-c-zephyr/rules/cicd.md](../../coding-standard-c-zephyr/rules/cicd.md)

---

## References

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GitLab CI/CD Documentation](https://docs.gitlab.com/ee/ci/)
- [Trunk Based Development](https://trunkbaseddevelopment.com/)
- [The Twelve-Factor App](https://12factor.net/)
- [DORA Metrics](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-performance)
