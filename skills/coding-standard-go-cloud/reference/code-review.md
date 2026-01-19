# Code Review Checklist

Use this checklist during code reviews to ensure compliance with the Go coding standard.

## Structure and Organization

- [ ] Code is in the correct package (`internal/` for private, `pkg/` for public)
- [ ] No circular dependencies
- [ ] main.go is minimal (bootstrapping only)
- [ ] Package names are descriptive (not `util`, `common`, `helpers`)

**Related Rules**: GO-STR-001, GO-STR-002, GO-STR-003

---

## Naming

- [ ] Uses MixedCaps, no underscores
- [ ] Variable name length proportional to scope
- [ ] No stuttering (`user.UserName` should be `user.Name`)
- [ ] Interfaces use "-er" suffix for single methods
- [ ] No "Get" prefix on getter methods

**Related Rules**: GO-NAM-001, GO-NAM-002, GO-NAM-003, GO-NAM-004

---

## Error Handling

- [ ] All errors are handled (not ignored with `_`)
- [ ] Errors are wrapped with context using `%w`
- [ ] Error strings are lowercase without punctuation
- [ ] Sentinel errors used for conditions callers need to match
- [ ] No panic for expected error conditions

**Related Rules**: GO-ERR-001, GO-ERR-002, GO-ERR-003, GO-ERR-004, GO-ERR-005

---

## Concurrency

- [ ] Context passed as first parameter
- [ ] Context not stored in structs
- [ ] All goroutines have exit paths
- [ ] sync.WaitGroup.Add() called before starting goroutine
- [ ] Mutex used for state protection, channels for communication
- [ ] errgroup used for concurrent operations with error handling

**Related Rules**: GO-CON-001, GO-CON-002, GO-CON-003, GO-CON-004, GO-CON-005, GO-CON-006

---

## Interfaces

- [ ] Interfaces are small (1-3 methods)
- [ ] Interfaces defined at consumer site
- [ ] Functions accept interfaces, return structs
- [ ] Compile-time interface checks for key types

**Related Rules**: GO-INT-001, GO-INT-002, GO-INT-003, GO-IDM-001

---

## Testing

- [ ] Table-driven tests for multiple scenarios
- [ ] t.Helper() used in test helpers
- [ ] t.Cleanup() used for teardown
- [ ] Tests can run in parallel where appropriate
- [ ] External dependencies use interfaces for mocking

**Related Rules**: GO-TST-001, GO-TST-002, GO-TST-003, GO-TST-004, GO-TST-005, GO-TST-006

---

## Documentation

- [ ] All exported declarations have doc comments
- [ ] Comments start with the name being documented
- [ ] Package has a doc comment

**Related Rules**: GO-DOC-001, GO-DOC-002

---

## Security

- [ ] No hardcoded secrets
- [ ] SQL uses parameterized queries
- [ ] External input is validated
- [ ] File paths validated against directory traversal
- [ ] html/template used for HTML output
- [ ] exec.Command inputs sanitized and validated

**Related Rules**: GO-SEC-001, GO-SEC-002, GO-SEC-003, GO-SEC-004, GO-SEC-005

---

## Performance

- [ ] Slices preallocated when size is known
- [ ] strings.Builder used for string concatenation in loops

**Related Rules**: GO-PRF-001, GO-PRF-002

---

## Formatting

- [ ] Imports grouped: stdlib, third-party, internal
- [ ] No unnecessary import renaming
- [ ] Code formatted with gofmt/goimports

**Related Rules**: GO-FMT-001, GO-FMT-002

---

## Language Idioms

- [ ] Composite literals used for struct initialization
- [ ] defer used for resource cleanup
- [ ] Generics used appropriately for type-safe collections
- [ ] Type constraints are appropriately restrictive

**Related Rules**: GO-IDM-001, GO-IDM-002, GO-IDM-003, GO-IDM-004, GO-IDM-005, GO-IDM-006
