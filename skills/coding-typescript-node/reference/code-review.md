# Code Review Checklist

Use this checklist during code reviews to ensure compliance with the TypeScript coding standard.

## Type Safety

- [ ] No `any` types (or documented exception)
- [ ] No unsafe type assertions (`as`)
- [ ] Unknown external data is validated at boundaries
- [ ] Null/undefined handled with narrowing, not assertions
- [ ] Error handling covers both success and failure cases
- [ ] Strict mode enabled in tsconfig.json

**Related Rules**: TS-TYP-001, TS-TYP-002, TS-TYP-003, TS-SEC-001, TS-SEC-002

---

## Types and Interfaces

- [ ] Interfaces used for object shapes
- [ ] Types used for unions, intersections, utilities
- [ ] Generic constraints are appropriate
- [ ] Public functions have explicit return types
- [ ] Built-in utility types preferred over custom

**Related Rules**: TS-TYP-004, TS-TYP-005, TS-TYP-006, TS-TYP-007, TS-TYP-008

---

## Null Handling

- [ ] Optional chaining used for nullable access
- [ ] Nullish coalescing used for defaults
- [ ] No blind non-null assertions
- [ ] Type guards used for narrowing
- [ ] Falsy values handled correctly (0, '', false)

**Related Rules**: TS-NUL-001, TS-NUL-002, TS-NUL-003, TS-NUL-004

---

## Async Code

- [ ] All promises are awaited or handled
- [ ] Independent operations use Promise.all
- [ ] Error handling in async functions is complete
- [ ] No floating promises
- [ ] Async functions return Promise<T>

**Related Rules**: TS-ASY-001, TS-ASY-002, TS-ASY-003, TS-ASY-004

---

## Module Organization

- [ ] Imports are properly ordered (stdlib, external, internal, types)
- [ ] Type-only imports use `type` keyword
- [ ] File extensions are included (.js for ESM)
- [ ] No circular dependencies
- [ ] Barrel files used judiciously

**Related Rules**: TS-MOD-001, TS-MOD-002, TS-MOD-003, TS-MOD-004, TS-MOD-005

---

## Error Handling

- [ ] Custom error classes for domain errors
- [ ] Catch parameters typed as unknown
- [ ] Error messages provide context
- [ ] Expected failures use Result pattern (where appropriate)
- [ ] Re-throws preserve original error

**Related Rules**: TS-ERR-001, TS-ERR-002, TS-ERR-003

---

## Testing

- [ ] Tests cover new/changed functionality
- [ ] Mocks and fixtures are properly typed
- [ ] Error cases are tested
- [ ] Edge cases are considered
- [ ] Type assertions verify expected types

**Related Rules**: TS-TST-001, TS-TST-002

---

## Documentation

- [ ] Public APIs have JSDoc comments
- [ ] Comments describe behavior, not types
- [ ] Complex logic is explained
- [ ] No duplicate type information in comments
- [ ] Examples provided for complex functions

**Related Rules**: TS-DOC-001

---

## Security

- [ ] External input is validated with Zod/io-ts
- [ ] No hardcoded secrets
- [ ] Sensitive data is not logged
- [ ] SQL uses parameterized queries
- [ ] User input is sanitized before output

**Related Rules**: TS-SEC-001, TS-SEC-002

---

## Performance

- [ ] No overly complex type expressions
- [ ] Appropriate use of const assertions
- [ ] Efficient async patterns
- [ ] No unnecessary type computations

**Related Rules**: TS-PRF-001, TS-PRF-002

---

## Configuration

- [ ] tsconfig.json has strict mode enabled
- [ ] ESLint configured with typescript-eslint
- [ ] Prettier configured for consistent formatting
- [ ] Pre-commit hooks installed

---

## Quick Checks

### Critical (Must Fix)

```
[ ] strict: true in tsconfig.json
[ ] No any types without justification
[ ] All promises handled
[ ] Unknown catch parameters narrowed
[ ] External input validated
[ ] No unsafe type assertions
```

### Required (Fix Before Merge)

```
[ ] Explicit return types on public functions
[ ] Type-only imports for types
[ ] Import order consistent
[ ] Typed test fixtures
[ ] Error classes for domain errors
```

### Recommended (Consider Fixing)

```
[ ] Meaningful generic names
[ ] Built-in utility types used
[ ] Named exports preferred
[ ] Result pattern for expected failures
[ ] Public API documented
```
