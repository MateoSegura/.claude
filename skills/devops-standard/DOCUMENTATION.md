# Documentation Standard

> **Version**: 1.0.0
> **Status**: Active
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes documentation requirements for all projects. It ensures:

- **Discoverability**: Developers can find and understand codebases quickly
- **Maintainability**: Documentation stays current with code changes
- **Onboarding**: New team members can become productive faster
- **Operational Readiness**: Systems can be operated and debugged reliably

### 1.2 Scope

This standard applies to:

- All production repositories
- Internal tools and libraries
- APIs (internal and external)
- Infrastructure and deployment configurations
- Operational runbooks

### 1.3 Audience

- Software engineers creating and maintaining code
- Technical writers
- DevOps and SRE teams
- Code reviewers evaluating documentation quality

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | CRITICAL | CI blocking | Build fails, merge blocked |
| **Required** | REQUIRED | PR review | Must fix before merge |
| **Recommended** | RECOMMENDED | Advisory | Fix encouraged |

---

## 3. README Requirements

### 3.1 Required README Elements

#### DOC-001: Every Repository Must Have a README CRITICAL

**Rationale**: A README is the entry point for any repository. Without it, developers waste time figuring out basic project information.

Every repository MUST contain a `README.md` at the root with the following sections:

| Section | Required | Description |
|---------|----------|-------------|
| Project Title | Yes | Clear, descriptive name |
| Description | Yes | 2-3 sentences explaining what this project does |
| Status Badges | Recommended | Build status, coverage, version |
| Installation | Yes | How to set up the development environment |
| Usage | Yes | Basic usage examples |
| Configuration | If applicable | Environment variables, config files |
| Contributing | Yes | Link to CONTRIBUTING.md or inline guidelines |
| License | Yes | License type and link to LICENSE file |

### 3.2 README Quality Standards

#### DOC-002: README Must Enable Quick Start REQUIRED

**Rationale**: A developer should be able to clone and run the project within 15 minutes using only the README.

Requirements:

- Installation steps must be copy-pasteable commands
- All prerequisites must be listed with version requirements
- Common setup issues and solutions should be documented
- Example output should be shown where helpful

### 3.3 README Template

```markdown
# Project Name

Brief description of what this project does and its primary use case.

[![Build Status](https://example.com/build-badge.svg)](https://example.com/builds)
[![Coverage](https://example.com/coverage-badge.svg)](https://example.com/coverage)
[![Version](https://example.com/version-badge.svg)](https://example.com/releases)

## Overview

A more detailed explanation of the project (2-3 paragraphs):
- What problem does it solve?
- Who is it for?
- Key features and capabilities

## Prerequisites

- Go 1.21 or later
- PostgreSQL 14 or later
- Node.js 18 LTS (for frontend development)

## Installation

### Quick Start

```bash
# Clone the repository
git clone https://github.com/org/project.git
cd project

# Install dependencies
make deps

# Set up environment
cp .env.example .env
# Edit .env with your configuration

# Run the application
make run
```

### Development Setup

Detailed instructions for setting up a full development environment.

## Usage

### Basic Example

```bash
# Start the server
./project serve --port 8080

# Make a request
curl http://localhost:8080/api/health
```

### Common Operations

Document the most common operations users will perform.

## Configuration

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DATABASE_URL` | PostgreSQL connection string | - | Yes |
| `PORT` | Server port | `8080` | No |
| `LOG_LEVEL` | Logging verbosity | `info` | No |

See [Configuration Guide](docs/configuration.md) for detailed options.

## Architecture

Brief overview of the system architecture. Link to detailed documentation.

```
┌─────────┐     ┌─────────┐     ┌─────────┐
│  Client │────>│   API   │────>│   DB    │
└─────────┘     └─────────┘     └─────────┘
```

## API Reference

See [API Documentation](docs/api.md) for complete endpoint reference.

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for:
- Code of conduct
- Development workflow
- Pull request process
- Coding standards

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.

## Support

- Documentation: [docs/](docs/)
- Issues: [GitHub Issues](https://github.com/org/project/issues)
- Discussions: [GitHub Discussions](https://github.com/org/project/discussions)
```

---

## 4. API Documentation

### 4.1 API Documentation Requirements

#### DOC-003: All Public APIs Must Be Documented CRITICAL

**Rationale**: Undocumented APIs are effectively unusable and lead to incorrect integrations, support burden, and breaking changes.

All public APIs (REST, GraphQL, gRPC) MUST have:

1. Complete endpoint/operation documentation
2. Request and response schemas with examples
3. Error codes and their meanings
4. Authentication and authorization requirements
5. Rate limiting information
6. Versioning policy

### 4.2 REST API Documentation Format

#### DOC-004: Use OpenAPI 3.0+ Specification REQUIRED

**Rationale**: OpenAPI provides a machine-readable format that enables automated tooling, client generation, and validation.

All REST APIs MUST provide an OpenAPI specification that includes:

```yaml
openapi: 3.0.3
info:
  title: Service Name API
  description: |
    Detailed description of the API, its purpose, and key concepts.

    ## Authentication
    This API uses Bearer token authentication. Include the token in the
    Authorization header: `Authorization: Bearer <token>`

    ## Rate Limiting
    - 1000 requests per minute for authenticated users
    - 100 requests per minute for unauthenticated requests

    ## Versioning
    API version is included in the URL path: `/api/v1/...`
  version: 1.0.0
  contact:
    name: API Support
    email: api-support@example.com

servers:
  - url: https://api.example.com/v1
    description: Production
  - url: https://api.staging.example.com/v1
    description: Staging

paths:
  /users/{userId}:
    get:
      summary: Get user by ID
      description: |
        Retrieves a user's profile information by their unique identifier.

        Requires `users:read` scope.
      operationId: getUserById
      tags:
        - Users
      parameters:
        - name: userId
          in: path
          required: true
          description: Unique user identifier (UUID format)
          schema:
            type: string
            format: uuid
          example: "123e4567-e89b-12d3-a456-426614174000"
      responses:
        '200':
          description: User found
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User'
              example:
                id: "123e4567-e89b-12d3-a456-426614174000"
                email: "user@example.com"
                name: "John Doe"
                createdAt: "2024-01-15T10:30:00Z"
        '404':
          description: User not found
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
              example:
                code: "USER_NOT_FOUND"
                message: "No user found with the specified ID"
        '401':
          $ref: '#/components/responses/Unauthorized'
        '500':
          $ref: '#/components/responses/InternalError'
      security:
        - bearerAuth: [users:read]

components:
  schemas:
    User:
      type: object
      required:
        - id
        - email
        - name
      properties:
        id:
          type: string
          format: uuid
          description: Unique user identifier
        email:
          type: string
          format: email
          description: User's email address
        name:
          type: string
          description: User's display name
          minLength: 1
          maxLength: 100
        createdAt:
          type: string
          format: date-time
          description: Account creation timestamp

    Error:
      type: object
      required:
        - code
        - message
      properties:
        code:
          type: string
          description: Machine-readable error code
        message:
          type: string
          description: Human-readable error description
        details:
          type: object
          description: Additional error context
          additionalProperties: true

  responses:
    Unauthorized:
      description: Authentication required or invalid
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
          example:
            code: "UNAUTHORIZED"
            message: "Invalid or missing authentication token"

    InternalError:
      description: Internal server error
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
          example:
            code: "INTERNAL_ERROR"
            message: "An unexpected error occurred"

  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
```

### 4.3 Error Documentation

#### DOC-005: Document All Error Codes REQUIRED

**Rationale**: Comprehensive error documentation enables clients to handle failures gracefully and reduces support burden.

All APIs MUST document:

| Element | Description |
|---------|-------------|
| Error code | Machine-readable identifier (e.g., `USER_NOT_FOUND`) |
| HTTP status | Appropriate HTTP status code |
| Message | Human-readable description |
| Resolution | How to resolve the error (if applicable) |

Example error code documentation:

```markdown
## Error Codes

### Authentication Errors (401)

| Code | Description | Resolution |
|------|-------------|------------|
| `AUTH_TOKEN_MISSING` | No authentication token provided | Include Bearer token in Authorization header |
| `AUTH_TOKEN_EXPIRED` | Authentication token has expired | Request a new token from /auth/token |
| `AUTH_TOKEN_INVALID` | Token format or signature invalid | Verify token is correctly formatted JWT |

### Authorization Errors (403)

| Code | Description | Resolution |
|------|-------------|------------|
| `FORBIDDEN` | User lacks required permissions | Request appropriate role/permissions |
| `SCOPE_INSUFFICIENT` | Token missing required scope | Request token with necessary scopes |

### Resource Errors (404)

| Code | Description | Resolution |
|------|-------------|------------|
| `USER_NOT_FOUND` | Requested user does not exist | Verify user ID is correct |
| `RESOURCE_NOT_FOUND` | Generic resource not found | Check resource identifier |

### Validation Errors (400)

| Code | Description | Resolution |
|------|-------------|------------|
| `VALIDATION_ERROR` | Request body validation failed | Check `details` field for specific issues |
| `INVALID_FORMAT` | Field format is incorrect | Follow documented format requirements |

### Server Errors (500)

| Code | Description | Resolution |
|------|-------------|------------|
| `INTERNAL_ERROR` | Unexpected server error | Retry request; contact support if persistent |
| `SERVICE_UNAVAILABLE` | Dependent service unavailable | Retry with exponential backoff |
```

### 4.4 API Versioning Documentation

#### DOC-006: Document API Versioning Policy REQUIRED

**Rationale**: Clear versioning policy sets expectations for clients and prevents breaking changes.

API documentation MUST include:

```markdown
## Versioning Policy

This API follows semantic versioning. The major version is included in the URL path.

### Version Lifecycle

| Phase | Duration | Support Level |
|-------|----------|---------------|
| Current | Ongoing | Full support, new features |
| Deprecated | 12 months | Security fixes only |
| Retired | - | No support, may be removed |

### Breaking Changes

The following are considered breaking changes and require a major version bump:
- Removing an endpoint
- Removing a required field from responses
- Adding a required field to requests
- Changing field types
- Changing authentication requirements

### Non-Breaking Changes

The following changes may occur within a version:
- Adding new endpoints
- Adding optional request fields
- Adding response fields
- Adding new error codes
- Performance improvements

### Deprecation Process

1. Deprecation announced in changelog and API response headers
2. `Deprecation` header added to affected endpoints
3. 12-month migration period
4. Endpoint removed in next major version
```

---

## 5. Code Documentation

### 5.1 Comment Philosophy

#### DOC-007: Comment Why, Not What REQUIRED

**Rationale**: Code should be self-documenting for "what" it does. Comments add value by explaining "why" decisions were made.

```go
// BAD: Comments that describe what code does
// Loop through users
for _, user := range users {
    // Check if user is active
    if user.IsActive {
        // Add to result
        result = append(result, user)
    }
}

// GOOD: Comments that explain why
// Filter to active users only - inactive users are soft-deleted
// and should not appear in any user-facing lists per GDPR requirements
for _, user := range users {
    if user.IsActive {
        result = append(result, user)
    }
}
```

### 5.2 When to Comment

#### DOC-008: Required Comment Scenarios REQUIRED

**Rationale**: Certain situations always benefit from explanatory comments.

MUST add comments for:

| Scenario | Example |
|----------|---------|
| Non-obvious business logic | Regulatory requirements, domain rules |
| Performance optimizations | Why a less readable approach was chosen |
| Workarounds | Bug fixes, library limitations |
| Magic numbers | Configuration values, thresholds |
| Complex algorithms | High-level explanation of approach |
| Security considerations | Why certain checks exist |
| Concurrency reasoning | Lock ordering, synchronization rationale |

```go
// GOOD: Explaining business logic
// Accounts dormant for 2+ years require reactivation per banking regulations
// (12 CFR 1005.20). Send notification before archiving.
if time.Since(account.LastActivity) > 2*365*24*time.Hour {
    notifyDormantAccount(account)
}

// GOOD: Explaining a workaround
// Using string concatenation instead of fmt.Sprintf due to performance
// critical path - benchmarks showed 3x improvement (see issue #1234)
key := prefix + ":" + id

// GOOD: Explaining magic number
const maxRetries = 3 // Based on p99 latency SLO - 3 retries stays under 500ms
```

### 5.3 Public API Documentation

#### DOC-009: All Public APIs Must Have Documentation Comments CRITICAL

**Rationale**: Public APIs form contracts with consumers. Documentation prevents misuse and reduces support burden.

All exported functions, types, and constants MUST have documentation comments.

### 5.4 Language-Specific Documentation Formats

#### Go (godoc)

```go
// UserService provides operations for managing user accounts.
// It handles user creation, authentication, and profile management.
//
// UserService is safe for concurrent use.
type UserService struct {
    // ...
}

// CreateUser registers a new user account with the provided details.
// It validates the email format, checks for duplicates, and sends
// a verification email upon successful creation.
//
// Returns ErrEmailExists if a user with the email already exists.
// Returns ErrInvalidEmail if the email format is invalid.
func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    // ...
}
```

#### TypeScript (JSDoc/TSDoc)

```typescript
/**
 * Service for managing user accounts.
 * Handles user creation, authentication, and profile management.
 *
 * @example
 * ```typescript
 * const service = new UserService(config);
 * const user = await service.createUser({
 *   email: 'user@example.com',
 *   name: 'John Doe'
 * });
 * ```
 */
class UserService {
  /**
   * Creates a new user account.
   *
   * @param request - User creation parameters
   * @returns The created user object
   * @throws {EmailExistsError} If email is already registered
   * @throws {ValidationError} If request parameters are invalid
   */
  async createUser(request: CreateUserRequest): Promise<User> {
    // ...
  }
}
```

#### C (Doxygen)

```c
/**
 * @brief Initialize the sensor subsystem.
 *
 * Configures and calibrates all connected sensors. Must be called
 * before any sensor read operations. This function blocks until
 * calibration is complete (typically 2-3 seconds).
 *
 * @param[in] config Sensor configuration parameters
 * @param[out] status Initialization status for each sensor
 *
 * @return 0 on success, negative errno on failure
 * @retval -EINVAL Invalid configuration parameters
 * @retval -ETIMEDOUT Sensor calibration timeout
 * @retval -EIO Hardware communication failure
 *
 * @note Thread-safe. Can be called from any context.
 * @warning Do not call from interrupt context.
 */
int sensor_init(const struct sensor_config *config,
                struct sensor_status *status);
```

---

## 6. Architecture Decision Records (ADRs)

### 6.1 Purpose and Usage

#### DOC-010: Write ADR for Significant Decisions REQUIRED

**Rationale**: ADRs capture the context and reasoning behind architectural decisions, enabling future maintainers to understand why the system is built a certain way.

Write an ADR when:

| Trigger | Example |
|---------|---------|
| Technology selection | Choosing a database, framework, or library |
| Architectural pattern | Adopting microservices, event sourcing, CQRS |
| Integration approach | How systems will communicate |
| Significant trade-off | Choosing consistency over availability |
| Policy decision | Error handling strategy, logging approach |
| Breaking from convention | Why you diverged from standard practice |

### 6.2 ADR Storage and Numbering

ADRs are stored in `docs/adr/` with sequential numbering:

```
docs/
└── adr/
    ├── 0001-record-architecture-decisions.md
    ├── 0002-use-postgresql-for-persistence.md
    ├── 0003-adopt-event-sourcing-for-orders.md
    └── README.md  # Index of all ADRs
```

### 6.3 ADR Template

```markdown
# ADR-NNNN: Title

**Status**: Proposed | Accepted | Deprecated | Superseded by [ADR-XXXX](XXXX-title.md)

**Date**: YYYY-MM-DD

**Deciders**: [List of people involved in the decision]

**Technical Story**: [Link to ticket/issue if applicable]

## Context

Describe the situation that requires a decision. Include:
- What is the current state?
- What problem are we trying to solve?
- What constraints exist (technical, business, timeline)?
- What forces are at play?

Be factual and objective. This section should make the decision understandable
even to someone unfamiliar with the project.

## Decision

State the decision clearly and concisely.

"We will use [technology/pattern/approach] for [purpose]."

Explain the key aspects of the decision:
- What specifically are we doing?
- What are the boundaries of this decision?
- What is explicitly NOT covered by this decision?

## Consequences

### Positive

- Benefit 1
- Benefit 2
- Benefit 3

### Negative

- Drawback 1 and how we will mitigate it
- Drawback 2 and how we will mitigate it

### Neutral

- Change that is neither positive nor negative

## Alternatives Considered

### Alternative 1: [Name]

Description of the alternative.

**Pros**:
- Advantage 1
- Advantage 2

**Cons**:
- Disadvantage 1
- Disadvantage 2

**Why rejected**: Brief explanation.

### Alternative 2: [Name]

[Same structure as above]

## References

- [Link to relevant documentation]
- [Link to benchmark results]
- [Link to related ADRs]
```

### 6.4 ADR Lifecycle

#### DOC-011: Maintain ADR Status REQUIRED

**Rationale**: Outdated ADRs can mislead future developers. Status tracking ensures ADRs remain useful references.

| Status | Meaning |
|--------|---------|
| **Proposed** | Under discussion, not yet accepted |
| **Accepted** | Decision has been made and is in effect |
| **Deprecated** | Decision is no longer relevant but kept for history |
| **Superseded** | Replaced by a newer ADR (link to successor) |

When superseding an ADR:

1. Create the new ADR
2. Update the old ADR's status to "Superseded by ADR-XXXX"
3. Add link to old ADR in new ADR's references

---

## 7. Runbooks and Operations Documentation

### 7.1 Deployment Documentation

#### DOC-012: Document Deployment Procedures CRITICAL

**Rationale**: Deployment documentation ensures consistent, repeatable deployments and reduces risk of human error.

Every deployable service MUST have:

```markdown
# Deployment Guide: [Service Name]

## Prerequisites

- [ ] Access to deployment environment
- [ ] Required credentials configured
- [ ] Dependent services available

## Deployment Checklist

### Pre-Deployment

- [ ] All tests passing in CI
- [ ] Code review approved
- [ ] Database migrations reviewed (if applicable)
- [ ] Feature flags configured
- [ ] Monitoring dashboards ready

### Deployment Steps

1. **Notify team**
   ```
   Post in #deployments: "Starting [service] deployment to [env]"
   ```

2. **Run database migrations** (if applicable)
   ```bash
   make migrate ENV=production
   ```

3. **Deploy application**
   ```bash
   make deploy ENV=production VERSION=v1.2.3
   ```

4. **Verify deployment**
   ```bash
   # Check health endpoint
   curl https://api.example.com/health

   # Verify version
   curl https://api.example.com/version
   ```

5. **Monitor for issues**
   - Check error rates in [Dashboard Link]
   - Monitor latency in [Dashboard Link]
   - Watch logs in [Logging Link]

### Post-Deployment

- [ ] Verify core functionality
- [ ] Check for error rate increase
- [ ] Confirm metrics are reporting
- [ ] Update deployment log

## Rollback Procedure

If issues are detected:

1. **Initiate rollback**
   ```bash
   make rollback ENV=production VERSION=v1.2.2
   ```

2. **Verify rollback**
   ```bash
   curl https://api.example.com/version
   ```

3. **Notify team**
   ```
   Post in #deployments: "Rolled back [service] to [version] due to [reason]"
   ```

4. **Create incident ticket**

## Configuration Changes

| Change Type | Process |
|-------------|---------|
| Environment variable | Update in [config management system] |
| Feature flag | Toggle in [feature flag system] |
| Secrets | Update via [secrets management] |

## Contacts

| Role | Contact |
|------|---------|
| Service owner | @team-name |
| On-call | [PagerDuty link] |
```

### 7.2 Troubleshooting Guides

#### DOC-013: Provide Troubleshooting Documentation REQUIRED

**Rationale**: Troubleshooting guides reduce mean time to resolution and enable on-call engineers to resolve issues without escalation.

```markdown
# Troubleshooting Guide: [Service Name]

## Quick Diagnostics

```bash
# Check service health
curl -s https://api.example.com/health | jq

# Check recent logs
kubectl logs -l app=service-name --since=10m

# Check resource usage
kubectl top pods -l app=service-name
```

## Common Issues

### Issue: High Error Rate

**Symptoms**:
- Error rate above 1% on dashboard
- 5xx responses in logs
- Alerts firing for error SLO

**Diagnostic Steps**:
1. Check error logs for stack traces
   ```bash
   kubectl logs -l app=service-name | grep -i error | tail -50
   ```
2. Check dependent service health
3. Review recent deployments

**Common Causes and Resolutions**:

| Cause | Evidence | Resolution |
|-------|----------|------------|
| Database connection exhausted | "connection pool exhausted" in logs | Scale up database or restart pods |
| Downstream service unavailable | Timeout errors to specific service | Check downstream service status |
| Memory pressure | OOM kills in events | Increase memory limits or fix leak |

### Issue: High Latency

**Symptoms**:
- P99 latency above SLO threshold
- Slow response alerts
- User complaints about performance

**Diagnostic Steps**:
1. Check database query performance
2. Review distributed traces
3. Check for resource contention

**Common Causes and Resolutions**:

| Cause | Evidence | Resolution |
|-------|----------|------------|
| Slow database queries | High DB latency in traces | Optimize queries, add indexes |
| GC pressure | GC pause time in metrics | Tune GC, reduce allocations |
| Network issues | High network latency in traces | Check network path, DNS |

### Issue: Service Unavailable

**Symptoms**:
- Health checks failing
- No pods running
- Connection refused errors

**Diagnostic Steps**:
1. Check pod status
   ```bash
   kubectl get pods -l app=service-name
   kubectl describe pod <pod-name>
   ```
2. Check events
   ```bash
   kubectl get events --sort-by='.lastTimestamp' | tail -20
   ```
3. Check node status

**Common Causes and Resolutions**:

| Cause | Evidence | Resolution |
|-------|----------|------------|
| Image pull failure | "ImagePullBackOff" status | Verify image exists and credentials |
| Resource quota exceeded | "Insufficient" in events | Request quota increase or optimize |
| Crash loop | "CrashLoopBackOff" status | Check logs for startup errors |

## Escalation

If unable to resolve after 15 minutes:

1. Page secondary on-call
2. Create incident channel: #inc-[date]-[service]
3. Begin incident documentation

## Useful Links

- [Service Dashboard](link)
- [Logs](link)
- [Traces](link)
- [Runbook Index](link)
```

### 7.3 On-Call Procedures

#### DOC-014: Document On-Call Procedures REQUIRED

**Rationale**: Clear on-call documentation ensures consistent incident response and reduces stress on responders.

```markdown
# On-Call Guide: [Team/Service]

## On-Call Responsibilities

- Respond to alerts within 15 minutes
- Assess severity and escalate as needed
- Document incidents and resolutions
- Hand off active incidents at shift end

## Alert Response

### Severity Levels

| Severity | Response Time | Examples |
|----------|--------------|----------|
| SEV-1 | Immediate | Complete outage, data loss |
| SEV-2 | 15 minutes | Partial outage, degraded performance |
| SEV-3 | 1 hour | Non-critical functionality affected |
| SEV-4 | Next business day | Minor issues, maintenance items |

### Initial Response Checklist

- [ ] Acknowledge alert
- [ ] Open incident channel (SEV-1/2)
- [ ] Check relevant dashboards
- [ ] Review recent changes/deployments
- [ ] Begin troubleshooting or escalate

## Incident Management

### Starting an Incident

1. Create Slack channel: `#inc-YYYYMMDD-brief-description`
2. Post incident template:
   ```
   **Incident Started**: [time]
   **Severity**: [SEV-N]
   **Impact**: [description]
   **Status**: Investigating
   **Incident Commander**: @[name]
   ```
3. Page additional responders if needed

### During the Incident

- Post updates every 15 minutes (SEV-1/2) or 30 minutes (SEV-3)
- Document all actions taken
- Coordinate with stakeholders

### Closing an Incident

1. Confirm service is restored
2. Post resolution summary
3. Create follow-up ticket for post-mortem
4. Update incident tracker

## Handoff Procedure

At shift end:

1. Review active alerts and incidents
2. Update incoming on-call on any ongoing issues
3. Transfer incident commander role if applicable
4. Log handoff in on-call channel

## Contacts

| Role | Contact | When to Use |
|------|---------|-------------|
| Secondary on-call | [PagerDuty] | Escalation, coverage |
| Engineering manager | @[name] | SEV-1, external communication |
| Security team | @security-oncall | Security incidents |
| Database team | @dba-oncall | Database issues |
```

---

## 8. Documentation Maintenance

### 8.1 Keeping Documentation Current

#### DOC-015: Update Documentation with Code Changes REQUIRED

**Rationale**: Outdated documentation is worse than no documentation - it actively misleads developers.

Documentation updates MUST be included in the same PR as code changes when:

- Public API signature changes
- Configuration options change
- New features are added
- Behavior changes
- Error codes or messages change

### 8.2 Documentation Review in Pull Requests

#### DOC-016: Review Documentation in PRs REQUIRED

**Rationale**: Documentation should receive the same review scrutiny as code.

Code review checklist for documentation:

- [ ] README updated if setup process changed
- [ ] API documentation reflects endpoint changes
- [ ] Code comments explain non-obvious logic
- [ ] Public APIs have documentation comments
- [ ] ADR written for architectural decisions
- [ ] Runbooks updated for operational changes
- [ ] No outdated comments remain

### 8.3 Deprecation Documentation

#### DOC-017: Document Deprecations Clearly REQUIRED

**Rationale**: Clear deprecation notices give consumers time to migrate and prevent surprises.

When deprecating APIs or features:

1. **Code annotation**: Use language-specific deprecation markers
   ```go
   // Deprecated: Use NewFunction instead. Will be removed in v3.0.
   func OldFunction() {}
   ```

2. **Documentation update**: Mark as deprecated in API docs
   ```yaml
   /old-endpoint:
     get:
       deprecated: true
       description: |
         **Deprecated**: Use `/new-endpoint` instead.
         This endpoint will be removed in API v3.
   ```

3. **Migration guide**: Provide clear migration path
   ```markdown
   ## Migration Guide: OldFunction to NewFunction

   ### What's Changing
   `OldFunction` is deprecated and will be removed in v3.0.

   ### How to Migrate
   Replace calls to `OldFunction()` with `NewFunction()`.

   | Before | After |
   |--------|-------|
   | `OldFunction(a, b)` | `NewFunction(ctx, a, b)` |

   ### Timeline
   - v2.5: Deprecation warning added
   - v2.8: Warning escalated to error in strict mode
   - v3.0: Function removed
   ```

4. **Changelog entry**: Note deprecation in release notes

### 8.4 Documentation Audits

#### DOC-018: Conduct Regular Documentation Audits RECOMMENDED

**Rationale**: Periodic audits catch documentation drift and ensure completeness.

Quarterly audit checklist:

- [ ] All repositories have current READMEs
- [ ] API documentation matches implementation
- [ ] Runbooks have been tested/validated
- [ ] Dead links identified and fixed
- [ ] Outdated ADRs marked as superseded
- [ ] New team members can onboard using docs alone

---

## 9. Documentation Types Summary

| Type | Location | Update Frequency | Owner |
|------|----------|------------------|-------|
| README | Repository root | With significant changes | Development team |
| API Docs | `docs/api/` or OpenAPI spec | With API changes | API team |
| Code Comments | In source files | With code changes | Individual developers |
| ADRs | `docs/adr/` | New decisions | Tech leads/architects |
| Runbooks | `docs/runbooks/` | Operational changes | SRE/DevOps team |
| Architecture | `docs/architecture/` | Major design changes | Tech leads |

---

## 10. Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| DOC-001 | Every repository must have a README | CRITICAL |
| DOC-002 | README must enable quick start | REQUIRED |
| DOC-003 | All public APIs must be documented | CRITICAL |
| DOC-004 | Use OpenAPI 3.0+ specification | REQUIRED |
| DOC-005 | Document all error codes | REQUIRED |
| DOC-006 | Document API versioning policy | REQUIRED |
| DOC-007 | Comment why, not what | REQUIRED |
| DOC-008 | Required comment scenarios | REQUIRED |
| DOC-009 | All public APIs must have documentation comments | CRITICAL |
| DOC-010 | Write ADR for significant decisions | REQUIRED |
| DOC-011 | Maintain ADR status | REQUIRED |
| DOC-012 | Document deployment procedures | CRITICAL |
| DOC-013 | Provide troubleshooting documentation | REQUIRED |
| DOC-014 | Document on-call procedures | REQUIRED |
| DOC-015 | Update documentation with code changes | REQUIRED |
| DOC-016 | Review documentation in PRs | REQUIRED |
| DOC-017 | Document deprecations clearly | REQUIRED |
| DOC-018 | Conduct regular documentation audits | RECOMMENDED |

---

## Appendix A: Documentation Review Checklist

Quick reference for reviewers:

### README Quality
- [ ] Can clone and run in 15 minutes
- [ ] All prerequisites listed with versions
- [ ] Commands are copy-pasteable
- [ ] Configuration options documented

### API Documentation
- [ ] All endpoints documented
- [ ] Request/response examples included
- [ ] Error codes explained
- [ ] Authentication requirements clear

### Code Comments
- [ ] Explain "why", not "what"
- [ ] Complex logic annotated
- [ ] Public APIs documented
- [ ] No outdated comments

### Operational Docs
- [ ] Deployment steps complete
- [ ] Rollback procedure documented
- [ ] Troubleshooting guide current
- [ ] On-call procedures clear

---

## Appendix B: Tooling Recommendations

| Purpose | Tool | Notes |
|---------|------|-------|
| API documentation | OpenAPI/Swagger | Machine-readable spec |
| API documentation UI | Swagger UI, Redoc | Auto-generated from OpenAPI |
| Static site generation | MkDocs, Docusaurus | For comprehensive docs |
| Diagram creation | Mermaid, PlantUML | Diagrams as code |
| Link checking | markdown-link-check | CI integration |
| Spell checking | cspell, aspell | Prevent typos |

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- [Write the Docs - Documentation Guide](https://www.writethedocs.org/guide/)
- [OpenAPI Specification](https://spec.openapis.org/oas/latest.html)
- [Architectural Decision Records](https://adr.github.io/)
- [Google Developer Documentation Style Guide](https://developers.google.com/style)
- [Microsoft Writing Style Guide](https://docs.microsoft.com/en-us/style-guide/)
