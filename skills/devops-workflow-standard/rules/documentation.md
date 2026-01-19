---
description: Documentation rules - READMEs, API docs, ADRs, runbooks
---

# Documentation Standard

## Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | CRITICAL | CI blocking | Build fails, merge blocked |
| **Required** | REQUIRED | PR review | Must fix before merge |
| **Recommended** | RECOMMENDED | Advisory | Fix encouraged |

---

## README Requirements

### DOC-001: Every Repository Must Have a README CRITICAL

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

### DOC-002: README Must Enable Quick Start REQUIRED

**Rationale**: A developer should be able to clone and run the project within 15 minutes using only the README.

Requirements:

- Installation steps must be copy-pasteable commands
- All prerequisites must be listed with version requirements
- Common setup issues and solutions should be documented
- Example output should be shown where helpful

---

## API Documentation

### DOC-003: All Public APIs Must Be Documented CRITICAL

**Rationale**: Undocumented APIs are effectively unusable and lead to incorrect integrations, support burden, and breaking changes.

All public APIs (REST, GraphQL, gRPC) MUST have:

1. Complete endpoint/operation documentation
2. Request and response schemas with examples
3. Error codes and their meanings
4. Authentication and authorization requirements
5. Rate limiting information
6. Versioning policy

### DOC-004: Use OpenAPI 3.0+ Specification REQUIRED

**Rationale**: OpenAPI provides a machine-readable format that enables automated tooling, client generation, and validation.

All REST APIs MUST provide an OpenAPI specification that includes:

- API info with title, description, version, and contact
- Server definitions for all environments
- Complete path documentation with operationId and tags
- Parameter descriptions with examples
- Response schemas for all status codes
- Security scheme definitions
- Reusable components for schemas, responses, and parameters

### DOC-005: Document All Error Codes REQUIRED

**Rationale**: Comprehensive error documentation enables clients to handle failures gracefully and reduces support burden.

All APIs MUST document:

| Element | Description |
|---------|-------------|
| Error code | Machine-readable identifier (e.g., `USER_NOT_FOUND`) |
| HTTP status | Appropriate HTTP status code |
| Message | Human-readable description |
| Resolution | How to resolve the error (if applicable) |

### DOC-006: Document API Versioning Policy REQUIRED

**Rationale**: Clear versioning policy sets expectations for clients and prevents breaking changes.

API documentation MUST include:

- Version lifecycle phases (Current, Deprecated, Retired)
- What constitutes breaking vs. non-breaking changes
- Deprecation process and timeline
- Migration guidance for version upgrades

---

## Code Documentation

### DOC-007: Comment Why, Not What REQUIRED

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

### DOC-008: Required Comment Scenarios REQUIRED

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

### DOC-009: All Public APIs Must Have Documentation Comments CRITICAL

**Rationale**: Public APIs form contracts with consumers. Documentation prevents misuse and reduces support burden.

All exported functions, types, and constants MUST have documentation comments.

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

## Architecture Decision Records (ADRs)

### DOC-010: Write ADR for Significant Decisions REQUIRED

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

ADRs are stored in `docs/adr/` with sequential numbering:

```
docs/
└── adr/
    ├── 0001-record-architecture-decisions.md
    ├── 0002-use-postgresql-for-persistence.md
    ├── 0003-adopt-event-sourcing-for-orders.md
    └── README.md  # Index of all ADRs
```

### DOC-011: Maintain ADR Status REQUIRED

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

## Runbooks and Operations Documentation

### DOC-012: Document Deployment Procedures CRITICAL

**Rationale**: Deployment documentation ensures consistent, repeatable deployments and reduces risk of human error.

Every deployable service MUST have:

- Prerequisites checklist (access, credentials, dependencies)
- Pre-deployment checklist (tests, reviews, migrations, feature flags)
- Step-by-step deployment instructions with commands
- Verification steps with specific commands and expected outputs
- Rollback procedure with commands and verification
- Configuration change procedures
- Contact information for escalation

### DOC-013: Provide Troubleshooting Documentation REQUIRED

**Rationale**: Troubleshooting guides reduce mean time to resolution and enable on-call engineers to resolve issues without escalation.

Troubleshooting guides MUST include:

- Quick diagnostic commands
- Common issues with symptoms, diagnostic steps, causes, and resolutions
- Escalation procedures
- Links to dashboards, logs, and traces

### DOC-014: Document On-Call Procedures REQUIRED

**Rationale**: Clear on-call documentation ensures consistent incident response and reduces stress on responders.

On-call documentation MUST include:

- Responsibilities and expectations
- Severity level definitions with response times
- Initial response checklist
- Incident management procedures (starting, during, closing)
- Handoff procedures
- Contact information for escalation

---

## Documentation Maintenance

### DOC-015: Update Documentation with Code Changes REQUIRED

**Rationale**: Outdated documentation is worse than no documentation - it actively misleads developers.

Documentation updates MUST be included in the same PR as code changes when:

- Public API signature changes
- Configuration options change
- New features are added
- Behavior changes
- Error codes or messages change

### DOC-016: Review Documentation in PRs REQUIRED

**Rationale**: Documentation should receive the same review scrutiny as code.

Code review checklist for documentation:

- [ ] README updated if setup process changed
- [ ] API documentation reflects endpoint changes
- [ ] Code comments explain non-obvious logic
- [ ] Public APIs have documentation comments
- [ ] ADR written for architectural decisions
- [ ] Runbooks updated for operational changes
- [ ] No outdated comments remain

### DOC-017: Document Deprecations Clearly REQUIRED

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

3. **Migration guide**: Provide clear migration path with before/after examples and timeline

4. **Changelog entry**: Note deprecation in release notes

### DOC-018: Conduct Regular Documentation Audits RECOMMENDED

**Rationale**: Periodic audits catch documentation drift and ensure completeness.

Quarterly audit checklist:

- [ ] All repositories have current READMEs
- [ ] API documentation matches implementation
- [ ] Runbooks have been tested/validated
- [ ] Dead links identified and fixed
- [ ] Outdated ADRs marked as superseded
- [ ] New team members can onboard using docs alone

---

## Rule Quick Reference

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

## Documentation Types Summary

| Type | Location | Update Frequency | Owner |
|------|----------|------------------|-------|
| README | Repository root | With significant changes | Development team |
| API Docs | `docs/api/` or OpenAPI spec | With API changes | API team |
| Code Comments | In source files | With code changes | Individual developers |
| ADRs | `docs/adr/` | New decisions | Tech leads/architects |
| Runbooks | `docs/runbooks/` | Operational changes | SRE/DevOps team |
| Architecture | `docs/architecture/` | Major design changes | Tech leads |
