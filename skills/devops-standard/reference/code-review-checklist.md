# Code Review Checklist

Quick reference for code reviewers evaluating changes.

---

## Pre-Review Checks

- [ ] PR description is complete and clear
- [ ] CI checks are passing
- [ ] PR is appropriately sized (<400 lines preferred)
- [ ] Linked issue/ticket exists

---

## Correctness

- [ ] Code does what the description claims
- [ ] Logic handles all documented requirements
- [ ] Edge cases are handled
- [ ] Error handling is appropriate
- [ ] No obvious bugs or logic errors
- [ ] Change doesn't break existing functionality

---

## Design

- [ ] Code is in the right location
- [ ] Abstractions are appropriate (not over/under-engineered)
- [ ] Minimal dependencies between components
- [ ] Consistent with existing patterns
- [ ] Will be maintainable in 6 months

---

## Complexity

- [ ] Functions are reasonably sized (<40 lines)
- [ ] Nesting depth is manageable (<=3 levels)
- [ ] No unnecessary complexity
- [ ] Could a simpler approach work?

---

## Tests

- [ ] New code has tests
- [ ] Tests cover edge cases
- [ ] Tests are readable and maintainable
- [ ] Tests actually test the change
- [ ] Bug fixes include regression tests

---

## Naming

- [ ] Names clearly communicate intent
- [ ] No abbreviations (except ubiquitous ones)
- [ ] Boolean names are questions (isEnabled, hasPermission)
- [ ] Names match domain terminology

---

## Comments

- [ ] Comments explain "why" not "what"
- [ ] Complex logic is explained
- [ ] No outdated comments
- [ ] No commented-out code

---

## Style

- [ ] Code passes all linter checks
- [ ] Formatting matches project standards
- [ ] Import ordering follows conventions
- [ ] Consistent with surrounding code

---

## Security

- [ ] No hardcoded secrets
- [ ] Input validation present
- [ ] SQL queries parameterized
- [ ] No injection vulnerabilities
- [ ] Sensitive data not logged
- [ ] Dependencies are trusted

---

## Documentation

- [ ] Public APIs documented
- [ ] README updated if needed
- [ ] Breaking changes documented
- [ ] Migration guide for breaking changes

---

## Extended Checklist (Complex Changes)

### Architecture

- [ ] Fits within system architecture
- [ ] Scaling implications considered
- [ ] Backwards compatibility maintained
- [ ] Migration path documented

### Performance

- [ ] No N+1 queries
- [ ] Appropriate caching
- [ ] Resource limits considered
- [ ] No memory leaks

### Observability

- [ ] Logging is appropriate
- [ ] Metrics added if needed
- [ ] Errors are traceable
- [ ] Health checks updated

---

## Comment Prefixes

Use these prefixes to clarify intent:

| Prefix | Meaning | Blocking? |
|--------|---------|-----------|
| `blocking:` | Must be addressed before merge | Yes |
| `question:` | Seeking clarification | Depends |
| `suggestion:` | Consider this alternative | No |
| `nit:` | Minor style/preference issue | No |
| `note:` | FYI, no action needed | No |
| `praise:` | Positive feedback | No |

---

## Approval Guidance

### Approve (LGTM) when:
- All CRITICAL and REQUIRED rules met
- All automated checks pass
- Adequate test coverage
- Documentation updated
- No blocking concerns

### Request Changes when:
- Correctness issues or bugs
- Security vulnerabilities
- Missing required tests
- Incomplete error handling
- Critical standard violations

### Comment (Non-blocking) when:
- Style suggestions
- Alternative approaches
- Minor improvements for future
- Educational observations
