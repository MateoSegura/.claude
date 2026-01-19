# CI/CD Standard

> **Version**: 0.1.0
> **Status**: Draft
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes Continuous Integration and Continuous Deployment (CI/CD) practices for all projects. It ensures:

- **Consistency**: Uniform pipeline structure across all repositories
- **Quality**: Automated enforcement of code quality and security standards
- **Reliability**: Reproducible builds and predictable deployments
- **Velocity**: Fast feedback loops for developers

### 1.2 Scope

This standard applies to:
- All production repositories
- All CI/CD pipeline configurations
- Build, test, and deployment automation
- Quality gate enforcement

### 1.3 Audience

- DevOps engineers configuring pipelines
- Software engineers modifying CI/CD workflows
- Tech leads establishing project pipelines
- Release managers coordinating deployments

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | CRITICAL | CI blocking | Build fails, deployment blocked |
| **Required** | REQUIRED | CI warning | Must fix before merge |
| **Recommended** | RECOMMENDED | Advisory | Fix encouraged |

---

## 3. Pipeline Architecture

### 3.1 Pipeline Stages Overview

Every pipeline MUST implement these core stages in order:

```
┌─────────────┐   ┌─────────────┐   ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
│    BUILD    │──▶│    TEST     │──▶│   ANALYZE   │──▶│  SECURITY   │──▶│   DEPLOY    │
└─────────────┘   └─────────────┘   └─────────────┘   └─────────────┘   └─────────────┘
     │                  │                 │                 │                 │
     ▼                  ▼                 ▼                 ▼                 ▼
  Compile           Unit Tests      Static Analysis   Vulnerability      Environment
  Dependencies      Integration     Code Coverage     Secret Scanning    Promotion
  Artifacts         E2E Tests       Lint Checks       SAST/DAST          Rollback
```

### 3.2 Stage Definitions

| Stage | Purpose | Blocking | Timeout |
|-------|---------|----------|---------|
| **Build** | Compile code, resolve dependencies, create artifacts | Yes | 10 min |
| **Test** | Execute automated tests | Yes | 20 min |
| **Analyze** | Static analysis, linting, coverage | Yes | 10 min |
| **Security** | Vulnerability and secret scanning | Yes | 15 min |
| **Deploy** | Promote artifacts to environments | Yes | 30 min |

---

## 4. Build Stage

### 4.1 Build Requirements

#### CICD-BUILD-001: Reproducible Builds CRITICAL

**Rationale**: Builds must produce identical artifacts given the same inputs to ensure reliability and auditability.

Requirements:
- Pin all dependency versions exactly (no ranges)
- Use lock files (`go.sum`, `package-lock.json`, `Cargo.lock`)
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

#### CICD-BUILD-002: Build Artifacts REQUIRED

**Rationale**: Artifacts must be traceable and identifiable throughout the deployment pipeline.

Requirements:
- Generate unique artifact identifiers (commit SHA + build number)
- Store artifacts in versioned registry (container registry, artifact repository)
- Include build metadata (git commit, timestamp, builder)
- Retain artifacts for minimum 90 days

```yaml
artifact-naming:
  format: "{app}-{version}-{commit_sha:8}-{build_number}"
  example: "api-server-1.2.3-abc12345-42"

metadata:
  required:
    - git_commit_sha
    - git_branch
    - build_timestamp
    - build_number
    - builder_version
```

#### CICD-BUILD-003: Version Tagging REQUIRED

**Rationale**: Consistent versioning enables traceability and rollback capabilities.

Version Tag Rules:
- Use semantic versioning (MAJOR.MINOR.PATCH)
- Tag releases in git with `v` prefix (`v1.2.3`)
- Pre-release versions use suffixes (`v1.2.3-rc.1`, `v1.2.3-beta.2`)
- Never reuse or overwrite version tags

```bash
# Release tagging
git tag -a v1.2.3 -m "Release 1.2.3: Feature description"
git push origin v1.2.3

# Pre-release tagging
git tag -a v1.2.3-rc.1 -m "Release candidate 1 for 1.2.3"
```

#### CICD-BUILD-004: Build Caching RECOMMENDED

**Rationale**: Effective caching significantly reduces build times and CI costs.

Caching Strategy:
- Cache dependency downloads (node_modules, go mod cache)
- Cache build outputs where safe (compiled objects, generated code)
- Use content-addressable cache keys (hash of lock files)
- Invalidate caches on tool version changes

```yaml
# GitHub Actions caching example
- name: Cache dependencies
  uses: actions/cache@v4
  with:
    path: |
      ~/.npm
      node_modules
    key: deps-${{ hashFiles('package-lock.json') }}
    restore-keys: |
      deps-

- name: Cache build
  uses: actions/cache@v4
  with:
    path: .next/cache
    key: build-${{ hashFiles('src/**', 'package-lock.json') }}
```

---

## 5. Test Stage

### 5.1 Test Execution Requirements

#### CICD-TEST-001: Test Parallelization REQUIRED

**Rationale**: Parallel test execution reduces feedback time and improves developer productivity.

Requirements:
- Run independent test suites in parallel
- Use test sharding for large test suites
- Balance shard sizes for consistent timing
- Report aggregated results

```yaml
# Parallel test matrix
test:
  strategy:
    matrix:
      shard: [1, 2, 3, 4]
  steps:
    - run: npm test -- --shard=${{ matrix.shard }}/4
```

#### CICD-TEST-002: Test Timeouts CRITICAL

**Rationale**: Unbounded tests can hang pipelines and waste resources.

Timeout Requirements:

| Test Type | Default Timeout | Maximum Allowed |
|-----------|-----------------|-----------------|
| Unit test (per test) | 5 seconds | 30 seconds |
| Integration test (per test) | 30 seconds | 2 minutes |
| E2E test (per test) | 2 minutes | 10 minutes |
| Test suite (total) | 10 minutes | 30 minutes |

```yaml
# Jest timeout configuration
testTimeout: 5000  # 5 seconds per test

# Go test timeout
go test -timeout 30s ./...

# Pytest timeout
pytest --timeout=30
```

#### CICD-TEST-003: Flaky Test Handling REQUIRED

**Rationale**: Flaky tests erode trust in CI and slow development velocity.

Flaky Test Policy:
1. **Detection**: Track test pass rates over time
2. **Quarantine**: Move tests with <95% pass rate to quarantine suite
3. **Fix or Remove**: Quarantined tests must be fixed within 2 weeks or removed
4. **No Retry Masking**: Do not use automatic retries to hide flakiness

```yaml
# Flaky test tracking
- name: Run tests with flaky detection
  run: |
    npm test --json > results.json
    # Upload to test analytics platform

# Quarantine configuration
quarantined_tests:
  - test_name: "flaky_network_test"
    quarantined_date: "2026-01-01"
    owner: "@developer"
    reason: "Intermittent timeout on CI"
```

#### CICD-TEST-004: Test Result Reporting REQUIRED

**Rationale**: Clear test reporting enables rapid diagnosis and trend analysis.

Reporting Requirements:
- Output results in standard format (JUnit XML, JSON)
- Include failure messages and stack traces
- Report test duration for performance tracking
- Upload results to CI system for visibility

```yaml
- name: Run tests
  run: npm test -- --reporter=junit --outputFile=results.xml

- name: Upload test results
  uses: actions/upload-artifact@v4
  if: always()
  with:
    name: test-results
    path: results.xml

- name: Publish test report
  uses: mikepenz/action-junit-report@v4
  if: always()
  with:
    report_paths: results.xml
```

---

## 6. Static Analysis Stage

### 6.1 Analysis Requirements

#### CICD-ANALYZE-001: Linting Enforcement CRITICAL

**Rationale**: Automated linting catches style violations and common errors before review.

Requirements:
- Run language-appropriate linters on all changed files
- Fail pipeline on linter errors
- Configure linters via committed config files
- No inline disables without justification comment

| Language | Linter | Config File |
|----------|--------|-------------|
| Go | golangci-lint | `.golangci.yml` |
| TypeScript | ESLint | `.eslintrc.js` |
| Python | ruff | `pyproject.toml` |
| Bash | ShellCheck | `.shellcheckrc` |
| C | clang-tidy | `.clang-tidy` |

#### CICD-ANALYZE-002: Code Coverage Requirements REQUIRED

**Rationale**: Coverage metrics ensure tests exercise critical code paths.

Coverage Thresholds:

| Metric | Minimum | Target |
|--------|---------|--------|
| Line coverage | 70% | 85% |
| Branch coverage | 60% | 75% |
| New code coverage | 80% | 90% |

```yaml
# Coverage enforcement
- name: Check coverage
  run: |
    npm test -- --coverage
    # Fail if below threshold
    npx nyc check-coverage --lines 70 --branches 60

# Coverage reporting
- name: Upload coverage
  uses: codecov/codecov-action@v4
  with:
    fail_ci_if_error: true
    minimum_coverage: 70
```

#### CICD-ANALYZE-003: Code Formatting CRITICAL

**Rationale**: Consistent formatting eliminates style debates and reduces diff noise.

Requirements:
- Run formatter check (not auto-fix) in CI
- Fail pipeline on formatting violations
- Use committed formatter configuration

```yaml
- name: Check formatting
  run: |
    # Go
    gofmt -l . | grep . && exit 1 || true

    # TypeScript
    npx prettier --check "src/**/*.ts"

    # Python
    ruff format --check .
```

---

## 7. Security Stage

### 7.1 Security Scanning Requirements

#### CICD-SEC-001: Dependency Vulnerability Scanning CRITICAL

**Rationale**: Third-party dependencies are a major attack vector.

Requirements:
- Scan all dependencies on every build
- Fail on critical/high severity vulnerabilities
- Block merge until vulnerabilities addressed
- Allow temporary exceptions with documented timeline

```yaml
- name: Dependency scan
  run: |
    # npm
    npm audit --audit-level=high

    # Go
    govulncheck ./...

    # Python
    pip-audit --strict

- name: Dependency review
  uses: actions/dependency-review-action@v4
  with:
    fail-on-severity: high
    deny-licenses: GPL-3.0, AGPL-3.0
```

#### CICD-SEC-002: Secret Scanning CRITICAL

**Rationale**: Secrets in code lead to security breaches.

Requirements:
- Scan all commits for secrets before merge
- Block commits containing detected secrets
- Scan git history for existing leaks (periodic)
- Rotate any exposed secrets immediately

```yaml
- name: Secret scanning
  uses: trufflesecurity/trufflehog@main
  with:
    extra_args: --only-verified

- name: Gitleaks scan
  uses: gitleaks/gitleaks-action@v2
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

#### CICD-SEC-003: Static Application Security Testing (SAST) REQUIRED

**Rationale**: SAST identifies security vulnerabilities in source code.

Requirements:
- Run SAST on all PRs
- Review and triage findings before merge
- Track false positive suppressions
- Address high/critical findings within SLA

| Severity | Resolution SLA |
|----------|---------------|
| Critical | 24 hours |
| High | 7 days |
| Medium | 30 days |
| Low | 90 days |

```yaml
- name: CodeQL analysis
  uses: github/codeql-action/analyze@v3
  with:
    category: "/language:${{ matrix.language }}"

- name: Semgrep scan
  uses: returntocorp/semgrep-action@v1
  with:
    config: p/default
```

---

## 8. Quality Gates

### 8.1 Pre-Merge Quality Gates

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

### 8.2 Quality Gate Configuration

```yaml
# Branch protection rules
branch_protection:
  main:
    required_status_checks:
      strict: true
      contexts:
        - build
        - test
        - lint
        - security-scan
        - coverage
    required_pull_request_reviews:
      required_approving_review_count: 1
      dismiss_stale_reviews: true
      require_code_owner_reviews: true
    restrictions:
      enforce_admins: true

# Quality gate thresholds
quality_gates:
  coverage:
    lines: 70
    branches: 60
    new_code: 80

  security:
    block_on: [critical, high]
    warn_on: [medium]

  performance:
    build_time_budget: 600  # 10 minutes
    test_time_budget: 1200  # 20 minutes
```

---

## 9. Deployment

### 9.1 Environment Promotion

#### CICD-DEPLOY-001: Environment Progression CRITICAL

**Rationale**: Staged deployments reduce production incident risk.

Environment Progression:

```
┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐
│   DEV   │───▶│ STAGING │───▶│   UAT   │───▶│  PROD   │
└─────────┘    └─────────┘    └─────────┘    └─────────┘
     │              │              │              │
  Automatic     Automatic      Manual         Manual
  on merge     on success    approval       approval
                             required       required
```

| Environment | Trigger | Approval | Testing |
|-------------|---------|----------|---------|
| Development | Auto on PR merge | None | Automated |
| Staging | Auto on dev success | None | Automated + Manual QA |
| UAT | Manual | Product Owner | Manual acceptance |
| Production | Manual | Release Manager + Ops | Smoke tests |

#### CICD-DEPLOY-002: Deployment Strategies REQUIRED

**Rationale**: Safe deployment strategies minimize user impact during releases.

Supported Strategies:

**Blue-Green Deployment** (Default for stateless services)
```yaml
deployment:
  strategy: blue-green
  steps:
    - deploy_to: green
    - run_smoke_tests: green
    - switch_traffic: green
    - wait_for_stability: 5m
    - teardown: blue
```

**Canary Deployment** (For high-traffic services)
```yaml
deployment:
  strategy: canary
  steps:
    - deploy_canary: 5%
    - monitor: 15m
    - increment: 25%
    - monitor: 15m
    - increment: 50%
    - monitor: 15m
    - full_rollout: 100%
```

**Rolling Deployment** (For stateful services)
```yaml
deployment:
  strategy: rolling
  max_unavailable: 25%
  max_surge: 25%
  min_ready_seconds: 30
```

#### CICD-DEPLOY-003: Rollback Procedures CRITICAL

**Rationale**: Fast rollback capability minimizes incident impact.

Rollback Requirements:
- All deployments MUST support instant rollback
- Rollback MUST NOT require a new build
- Previous version artifacts retained for 30 days minimum
- Database migrations MUST be backwards-compatible

```yaml
# Automated rollback on failure
deployment:
  rollback:
    automatic: true
    triggers:
      - health_check_failure
      - error_rate > 5%
      - latency_p99 > 2x baseline
    procedure:
      - revert_traffic
      - notify_team
      - create_incident

# Manual rollback command
rollback:
  command: |
    kubectl rollout undo deployment/$APP_NAME
    # OR
    aws deploy stop-deployment --deployment-id $ID --auto-rollback
```

#### CICD-DEPLOY-004: Feature Flags RECOMMENDED

**Rationale**: Feature flags enable safe deployment of incomplete features and rapid incident response.

Feature Flag Guidelines:
- Use feature flags for all user-facing changes
- Flags should have owners and expiration dates
- Remove flags within 30 days of full rollout
- Emergency kill switches for critical features

```yaml
# Feature flag configuration
feature_flags:
  new_checkout_flow:
    default: false
    environments:
      development: true
      staging: true
      production: false  # Controlled rollout
    rollout:
      percentage: 10
      users: ["beta_testers"]
    owner: "@product-team"
    expires: "2026-02-01"
```

---

## 10. Monitoring and Alerting

### 10.1 Pipeline Monitoring

#### CICD-MON-001: Pipeline Failure Notifications CRITICAL

**Rationale**: Fast notification of failures enables rapid response.

Notification Requirements:
- Notify on all pipeline failures
- Include failure reason and link to logs
- Route to appropriate channel (PR author, team)
- Escalate on repeated failures

```yaml
notifications:
  pipeline_failure:
    channels:
      - slack: "#ci-alerts"
      - email: pr_author
    content:
      - pipeline_name
      - failed_stage
      - failure_reason
      - logs_url
      - commit_info

  repeated_failure:
    threshold: 3  # consecutive failures
    escalate_to:
      - slack: "#engineering-urgent"
      - pagerduty: on-call
```

#### CICD-MON-002: Deployment Notifications REQUIRED

**Rationale**: Deployment visibility enables coordination and incident correlation.

Deployment Notification Events:
- Deployment started
- Deployment succeeded
- Deployment failed
- Rollback initiated
- Rollback completed

```yaml
notifications:
  deployment:
    channels:
      - slack: "#deployments"
    events:
      - started
      - completed
      - failed
      - rolled_back
    content:
      - environment
      - version
      - deployer
      - changelog_url
      - dashboard_url
```

#### CICD-MON-003: Performance Tracking RECOMMENDED

**Rationale**: Pipeline performance metrics identify optimization opportunities.

Metrics to Track:
- Build duration (by stage)
- Test duration (by suite)
- Queue time (time waiting for runner)
- Success rate (by pipeline, stage)
- Flaky test rate
- Cache hit rate

```yaml
metrics:
  pipeline:
    - name: build_duration_seconds
      type: histogram
      labels: [pipeline, stage, status]

    - name: test_duration_seconds
      type: histogram
      labels: [suite, shard]

    - name: pipeline_success_total
      type: counter
      labels: [pipeline, branch]

    - name: cache_hit_ratio
      type: gauge
      labels: [cache_type]

alerts:
  - name: slow_builds
    condition: build_duration_seconds > 600
    action: notify_devops

  - name: high_failure_rate
    condition: success_rate < 0.9 over 1h
    action: notify_team_lead
```

---

## 11. Example Pipeline Configuration

### 11.1 Complete GitHub Actions Pipeline

```yaml
# .github/workflows/ci.yml
name: CI Pipeline

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main, develop]

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

env:
  NODE_VERSION: '20.11.0'
  GO_VERSION: '1.22.0'

jobs:
  # ============================================
  # BUILD STAGE
  # ============================================
  build:
    name: Build
    runs-on: ubuntu-latest
    timeout-minutes: 10
    outputs:
      artifact-name: ${{ steps.build.outputs.artifact-name }}

    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Build
        id: build
        run: |
          npm run build
          ARTIFACT_NAME="${{ github.event.repository.name }}-${{ github.sha }}-${{ github.run_number }}"
          echo "artifact-name=$ARTIFACT_NAME" >> $GITHUB_OUTPUT

      - name: Upload build artifact
        uses: actions/upload-artifact@v4
        with:
          name: ${{ steps.build.outputs.artifact-name }}
          path: dist/
          retention-days: 90

  # ============================================
  # TEST STAGE
  # ============================================
  test-unit:
    name: Unit Tests
    runs-on: ubuntu-latest
    timeout-minutes: 15
    needs: build
    strategy:
      fail-fast: false
      matrix:
        shard: [1, 2, 3, 4]

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Run unit tests
        run: npm test -- --shard=${{ matrix.shard }}/4 --coverage --reporters=default --reporters=jest-junit
        env:
          JEST_JUNIT_OUTPUT_DIR: ./reports

      - name: Upload test results
        uses: actions/upload-artifact@v4
        if: always()
        with:
          name: test-results-${{ matrix.shard }}
          path: reports/

      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          flags: unit-tests-shard-${{ matrix.shard }}

  test-integration:
    name: Integration Tests
    runs-on: ubuntu-latest
    timeout-minutes: 20
    needs: build
    services:
      postgres:
        image: postgres:16
        env:
          POSTGRES_PASSWORD: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Run integration tests
        run: npm run test:integration -- --reporters=default --reporters=jest-junit
        env:
          DATABASE_URL: postgres://postgres:test@localhost:5432/test
          JEST_JUNIT_OUTPUT_DIR: ./reports

      - name: Upload test results
        uses: actions/upload-artifact@v4
        if: always()
        with:
          name: integration-test-results
          path: reports/

  # ============================================
  # ANALYZE STAGE
  # ============================================
  lint:
    name: Lint
    runs-on: ubuntu-latest
    timeout-minutes: 10

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Run ESLint
        run: npm run lint -- --format=@microsoft/eslint-formatter-sarif --output-file=eslint.sarif
        continue-on-error: true

      - name: Upload SARIF
        uses: github/codeql-action/upload-sarif@v3
        with:
          sarif_file: eslint.sarif

      - name: Check lint status
        run: npm run lint

  format-check:
    name: Format Check
    runs-on: ubuntu-latest
    timeout-minutes: 5

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Check formatting
        run: npx prettier --check "src/**/*.{ts,tsx,js,jsx,json,css,md}"

  coverage-check:
    name: Coverage Check
    runs-on: ubuntu-latest
    timeout-minutes: 15
    needs: [test-unit]

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Run tests with coverage
        run: npm test -- --coverage --coverageReporters=json-summary

      - name: Check coverage thresholds
        run: |
          npx nyc check-coverage \
            --lines 70 \
            --branches 60 \
            --functions 70 \
            --statements 70

  # ============================================
  # SECURITY STAGE
  # ============================================
  security-scan:
    name: Security Scan
    runs-on: ubuntu-latest
    timeout-minutes: 15

    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: ${{ env.NODE_VERSION }}
          cache: 'npm'

      - name: Dependency audit
        run: npm audit --audit-level=high

      - name: Secret scanning
        uses: trufflesecurity/trufflehog@main
        with:
          extra_args: --only-verified

  codeql:
    name: CodeQL Analysis
    runs-on: ubuntu-latest
    timeout-minutes: 15
    permissions:
      security-events: write

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Initialize CodeQL
        uses: github/codeql-action/init@v3
        with:
          languages: javascript-typescript

      - name: Autobuild
        uses: github/codeql-action/autobuild@v3

      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@v3

  dependency-review:
    name: Dependency Review
    runs-on: ubuntu-latest
    if: github.event_name == 'pull_request'

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Dependency Review
        uses: actions/dependency-review-action@v4
        with:
          fail-on-severity: high
          deny-licenses: GPL-3.0, AGPL-3.0

  # ============================================
  # QUALITY GATE
  # ============================================
  quality-gate:
    name: Quality Gate
    runs-on: ubuntu-latest
    needs: [build, test-unit, test-integration, lint, format-check, coverage-check, security-scan, codeql]
    if: always()

    steps:
      - name: Check all jobs passed
        run: |
          if [[ "${{ needs.build.result }}" != "success" ]] || \
             [[ "${{ needs.test-unit.result }}" != "success" ]] || \
             [[ "${{ needs.test-integration.result }}" != "success" ]] || \
             [[ "${{ needs.lint.result }}" != "success" ]] || \
             [[ "${{ needs.format-check.result }}" != "success" ]] || \
             [[ "${{ needs.coverage-check.result }}" != "success" ]] || \
             [[ "${{ needs.security-scan.result }}" != "success" ]] || \
             [[ "${{ needs.codeql.result }}" != "success" ]]; then
            echo "Quality gate failed"
            exit 1
          fi
          echo "All quality gates passed"

  # ============================================
  # DEPLOY STAGE
  # ============================================
  deploy-staging:
    name: Deploy to Staging
    runs-on: ubuntu-latest
    needs: [quality-gate]
    if: github.ref == 'refs/heads/main' && github.event_name == 'push'
    environment:
      name: staging
      url: https://staging.example.com

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Download artifact
        uses: actions/download-artifact@v4
        with:
          name: ${{ needs.build.outputs.artifact-name }}
          path: dist/

      - name: Deploy to staging
        run: |
          echo "Deploying to staging..."
          # Add actual deployment commands here

      - name: Run smoke tests
        run: |
          echo "Running smoke tests..."
          # Add smoke test commands here

      - name: Notify deployment
        uses: slackapi/slack-github-action@v1
        with:
          channel-id: 'deployments'
          payload: |
            {
              "text": "Deployed to staging",
              "blocks": [
                {
                  "type": "section",
                  "text": {
                    "type": "mrkdwn",
                    "text": "*Staging Deployment Complete*\nVersion: ${{ github.sha }}\nEnvironment: staging"
                  }
                }
              ]
            }
        env:
          SLACK_BOT_TOKEN: ${{ secrets.SLACK_BOT_TOKEN }}

  deploy-production:
    name: Deploy to Production
    runs-on: ubuntu-latest
    needs: [deploy-staging]
    if: github.ref == 'refs/heads/main' && github.event_name == 'push'
    environment:
      name: production
      url: https://example.com

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Download artifact
        uses: actions/download-artifact@v4
        with:
          name: ${{ needs.build.outputs.artifact-name }}
          path: dist/

      - name: Deploy to production (Blue-Green)
        run: |
          echo "Deploying to production (green)..."
          # 1. Deploy to green environment
          # 2. Run health checks
          # 3. Switch traffic
          # 4. Monitor for issues
          # 5. Teardown blue

      - name: Run smoke tests
        run: |
          echo "Running production smoke tests..."

      - name: Notify deployment
        uses: slackapi/slack-github-action@v1
        with:
          channel-id: 'deployments'
          payload: |
            {
              "text": "Deployed to production",
              "blocks": [
                {
                  "type": "section",
                  "text": {
                    "type": "mrkdwn",
                    "text": "*Production Deployment Complete*\nVersion: ${{ github.sha }}\nEnvironment: production"
                  }
                }
              ]
            }
        env:
          SLACK_BOT_TOKEN: ${{ secrets.SLACK_BOT_TOKEN }}
```

---

## 12. Pipeline Change Checklist

Use this checklist when modifying CI/CD pipelines:

### Pre-Change Review
- [ ] Change has been tested in feature branch
- [ ] No secrets or credentials hardcoded
- [ ] Timeouts are set appropriately
- [ ] Failure notifications configured
- [ ] Rollback procedure documented

### Stage Configuration
- [ ] Build stage produces versioned artifacts
- [ ] Test stage covers all test types
- [ ] Coverage thresholds enforced
- [ ] Security scans enabled
- [ ] Quality gates configured

### Deployment Configuration
- [ ] Environment progression defined
- [ ] Approval gates for production
- [ ] Rollback procedure tested
- [ ] Health checks configured
- [ ] Monitoring and alerts in place

### Documentation
- [ ] Pipeline changes documented
- [ ] Runbook updated if needed
- [ ] Team notified of changes

### Post-Change Verification
- [ ] Pipeline runs successfully
- [ ] All stages execute correctly
- [ ] Notifications working
- [ ] Metrics being collected

---

## 13. Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| CICD-BUILD-001 | Reproducible Builds | CRITICAL |
| CICD-BUILD-002 | Build Artifacts | REQUIRED |
| CICD-BUILD-003 | Version Tagging | REQUIRED |
| CICD-BUILD-004 | Build Caching | RECOMMENDED |
| CICD-TEST-001 | Test Parallelization | REQUIRED |
| CICD-TEST-002 | Test Timeouts | CRITICAL |
| CICD-TEST-003 | Flaky Test Handling | REQUIRED |
| CICD-TEST-004 | Test Result Reporting | REQUIRED |
| CICD-ANALYZE-001 | Linting Enforcement | CRITICAL |
| CICD-ANALYZE-002 | Code Coverage Requirements | REQUIRED |
| CICD-ANALYZE-003 | Code Formatting | CRITICAL |
| CICD-SEC-001 | Dependency Vulnerability Scanning | CRITICAL |
| CICD-SEC-002 | Secret Scanning | CRITICAL |
| CICD-SEC-003 | Static Application Security Testing | REQUIRED |
| CICD-DEPLOY-001 | Environment Progression | CRITICAL |
| CICD-DEPLOY-002 | Deployment Strategies | REQUIRED |
| CICD-DEPLOY-003 | Rollback Procedures | CRITICAL |
| CICD-DEPLOY-004 | Feature Flags | RECOMMENDED |
| CICD-MON-001 | Pipeline Failure Notifications | CRITICAL |
| CICD-MON-002 | Deployment Notifications | REQUIRED |
| CICD-MON-003 | Performance Tracking | RECOMMENDED |

---

## Appendix A: Glossary

| Term | Definition |
|------|------------|
| **Artifact** | A deployable unit produced by a build (binary, container image, package) |
| **Blue-Green** | Deployment strategy maintaining two identical environments for instant switchover |
| **Canary** | Deployment strategy gradually shifting traffic to new version |
| **Quality Gate** | Automated check that must pass before proceeding to next stage |
| **SAST** | Static Application Security Testing - analyzing source code for vulnerabilities |
| **DAST** | Dynamic Application Security Testing - testing running application |
| **Flaky Test** | A test that produces inconsistent results without code changes |

---

## Appendix B: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | 2026-01-04 | Initial draft |

---

## Appendix C: References

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GitLab CI/CD Documentation](https://docs.gitlab.com/ee/ci/)
- [Trunk Based Development](https://trunkbaseddevelopment.com/)
- [The Twelve-Factor App](https://12factor.net/)
- [OWASP DevSecOps Guideline](https://owasp.org/www-project-devsecops-guideline/)
- [DORA Metrics](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-performance)
