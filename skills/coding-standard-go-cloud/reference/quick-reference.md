# Quick Reference - All Rules

Complete table of all rules with IDs, descriptions, and tiers.

## Rule Tiers

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Rules by Category

### Structure (GO-STR-*)

| ID | Rule | Tier |
|----|------|------|
| GO-STR-001 | Use internal/ for Private Packages | Critical |
| GO-STR-002 | Keep main.go Minimal | Required |
| GO-STR-003 | Avoid Generic Package Names | Required |

### Naming (GO-NAM-*)

| ID | Rule | Tier |
|----|------|------|
| GO-NAM-001 | Use MixedCaps for All Names | Critical |
| GO-NAM-002 | Use Short Variable Names in Limited Scope | Recommended |
| GO-NAM-003 | Avoid Getters with "Get" Prefix | Required |
| GO-NAM-004 | Name Interfaces by Method + "er" | Required |

### Formatting (GO-FMT-*)

| ID | Rule | Tier |
|----|------|------|
| GO-FMT-001 | Group Imports in Standard Order | Required |
| GO-FMT-002 | Avoid Import Renaming Unless Necessary | Recommended |

### Idioms (GO-IDM-*)

| ID | Rule | Tier |
|----|------|------|
| GO-IDM-001 | Accept Interfaces, Return Structs | Required |
| GO-IDM-002 | Use Composite Literals for Initialization | Recommended |
| GO-IDM-003 | Use defer for Resource Cleanup | Required |
| GO-IDM-004 | Use Generics for Type-Safe Collections and Algorithms | Required |
| GO-IDM-005 | Prefer Interfaces Over Generics for Behavior Abstraction | Required |
| GO-IDM-006 | Use Appropriate Type Constraints | Required |

### Error Handling (GO-ERR-*)

| ID | Rule | Tier |
|----|------|------|
| GO-ERR-001 | Never Ignore Errors | Critical |
| GO-ERR-002 | Wrap Errors with Context | Required |
| GO-ERR-003 | Error Strings Should Not Be Capitalized | Required |
| GO-ERR-004 | Use Sentinel Errors for Expected Conditions | Required |
| GO-ERR-005 | Don't Panic for Normal Error Handling | Critical |

### Interfaces (GO-INT-*)

| ID | Rule | Tier |
|----|------|------|
| GO-INT-001 | Keep Interfaces Small | Required |
| GO-INT-002 | Define Interfaces at Consumer Site | Required |
| GO-INT-003 | Verify Interface Compliance at Compile Time | Recommended |

### Concurrency (GO-CON-*)

| ID | Rule | Tier |
|----|------|------|
| GO-CON-001 | Always Pass Context to Long-Running Operations | Critical |
| GO-CON-002 | Never Store Context in Structs | Critical |
| GO-CON-003 | Ensure Goroutines Can Exit | Critical |
| GO-CON-004 | Use sync.WaitGroup for Goroutine Coordination | Required |
| GO-CON-005 | Prefer Mutex for Simple State Protection | Recommended |
| GO-CON-006 | Use errgroup for Concurrent Operations with Error Handling | Required |

### Testing (GO-TST-*)

| ID | Rule | Tier |
|----|------|------|
| GO-TST-001 | Use Table-Driven Tests | Required |
| GO-TST-002 | Use t.Helper() for Test Helpers | Required |
| GO-TST-003 | Use t.Cleanup() for Test Teardown | Recommended |
| GO-TST-004 | Use Subtests for Parallel Execution | Recommended |
| GO-TST-005 | Use Interfaces for External Dependencies to Enable Mocking | Required |
| GO-TST-006 | Prefer Fakes Over Mocks for Complex Behavior | Recommended |

### Documentation (GO-DOC-*)

| ID | Rule | Tier |
|----|------|------|
| GO-DOC-001 | Document All Exported Declarations | Required |
| GO-DOC-002 | Use Package Comments | Required |

### Security (GO-SEC-*)

| ID | Rule | Tier |
|----|------|------|
| GO-SEC-001 | Never Hardcode Secrets | Critical |
| GO-SEC-002 | Use Parameterized Queries | Critical |
| GO-SEC-003 | Validate File Paths to Prevent Directory Traversal | Critical |
| GO-SEC-004 | Use html/template for HTML Output | Critical |
| GO-SEC-005 | Sanitize Inputs to exec.Command | Critical |

### Performance (GO-PRF-*)

| ID | Rule | Tier |
|----|------|------|
| GO-PRF-001 | Preallocate Slices When Size is Known | Recommended |
| GO-PRF-002 | Use strings.Builder for String Concatenation | Recommended |

---

## Summary by Tier

### Critical Rules (14 total)

- GO-STR-001, GO-NAM-001
- GO-ERR-001, GO-ERR-005
- GO-CON-001, GO-CON-002, GO-CON-003
- GO-SEC-001, GO-SEC-002, GO-SEC-003, GO-SEC-004, GO-SEC-005

### Required Rules (21 total)

- GO-STR-002, GO-STR-003
- GO-NAM-003, GO-NAM-004
- GO-FMT-001
- GO-IDM-001, GO-IDM-003, GO-IDM-004, GO-IDM-005, GO-IDM-006
- GO-ERR-002, GO-ERR-003, GO-ERR-004
- GO-INT-001, GO-INT-002
- GO-CON-004, GO-CON-006
- GO-TST-001, GO-TST-002, GO-TST-005
- GO-DOC-001, GO-DOC-002

### Recommended Rules (10 total)

- GO-NAM-002
- GO-FMT-002
- GO-IDM-002
- GO-INT-003
- GO-CON-005
- GO-TST-003, GO-TST-004, GO-TST-006
- GO-PRF-001, GO-PRF-002
