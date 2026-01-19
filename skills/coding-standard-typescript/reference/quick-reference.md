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

### Type System (TS-TYP-*)

| ID | Rule | Tier |
|----|------|------|
| TS-TYP-001 | Enable Strict Mode | Critical |
| TS-TYP-002 | No Implicit Any | Critical |
| TS-TYP-003 | Prefer Unknown Over Any | Critical |
| TS-TYP-004 | Use Interface for Object Shapes | Required |
| TS-TYP-005 | Use Explicit Return Types for Public Functions | Required |
| TS-TYP-006 | Use Meaningful Generic Names | Recommended |
| TS-TYP-007 | Constrain Generic Types Appropriately | Required |
| TS-TYP-008 | Prefer Built-in Utility Types | Recommended |

### Null Handling (TS-NUL-*)

| ID | Rule | Tier |
|----|------|------|
| TS-NUL-001 | Enable strictNullChecks | Critical |
| TS-NUL-002 | Use Optional Chaining for Nullable Access | Required |
| TS-NUL-003 | Use Nullish Coalescing for Defaults | Required |
| TS-NUL-004 | Avoid Non-Null Assertions | Required |

### Modules (TS-MOD-*)

| ID | Rule | Tier |
|----|------|------|
| TS-MOD-001 | Barrel File Usage | Required |
| TS-MOD-002 | Order Imports Consistently | Required |
| TS-MOD-003 | Use Type-Only Imports | Required |
| TS-MOD-004 | Use .js Extensions in Imports | Required |
| TS-MOD-005 | Prefer Named Exports | Recommended |

### Async Patterns (TS-ASY-*)

| ID | Rule | Tier |
|----|------|------|
| TS-ASY-001 | Always Handle Promise Rejections | Critical |
| TS-ASY-002 | Prefer async/await Over Promise Chains | Recommended |
| TS-ASY-003 | Use Promise.all for Independent Operations | Required |
| TS-ASY-004 | Type Async Functions Correctly | Required |

### Error Handling (TS-ERR-*)

| ID | Rule | Tier |
|----|------|------|
| TS-ERR-001 | Create Typed Error Classes | Required |
| TS-ERR-002 | Consider Result Pattern for Expected Failures | Recommended |
| TS-ERR-003 | Type Unknown Catch Parameters | Critical |

### Testing (TS-TST-*)

| ID | Rule | Tier |
|----|------|------|
| TS-TST-001 | Type Test Fixtures and Mocks | Required |
| TS-TST-002 | Test Type Inference with Explicit Assertions | Recommended |

### Documentation (TS-DOC-*)

| ID | Rule | Tier |
|----|------|------|
| TS-DOC-001 | Document Public API Without Type Duplication | Recommended |

### Security (TS-SEC-*)

| ID | Rule | Tier |
|----|------|------|
| TS-SEC-001 | Validate External Input at Boundaries | Critical |
| TS-SEC-002 | Avoid Unsafe Type Assertions | Critical |

### Performance (TS-PRF-*)

| ID | Rule | Tier |
|----|------|------|
| TS-PRF-001 | Avoid Overly Complex Type Expressions | Recommended |
| TS-PRF-002 | Use const Assertions for Immutable Data | Recommended |

---

## Summary by Tier

### Critical Rules (9 total)

| ID | Rule |
|----|------|
| TS-TYP-001 | Enable Strict Mode |
| TS-TYP-002 | No Implicit Any |
| TS-TYP-003 | Prefer Unknown Over Any |
| TS-NUL-001 | Enable strictNullChecks |
| TS-ASY-001 | Always Handle Promise Rejections |
| TS-ERR-003 | Type Unknown Catch Parameters |
| TS-SEC-001 | Validate External Input at Boundaries |
| TS-SEC-002 | Avoid Unsafe Type Assertions |

### Required Rules (14 total)

| ID | Rule |
|----|------|
| TS-TYP-004 | Use Interface for Object Shapes |
| TS-TYP-005 | Use Explicit Return Types for Public Functions |
| TS-TYP-007 | Constrain Generic Types Appropriately |
| TS-NUL-002 | Use Optional Chaining for Nullable Access |
| TS-NUL-003 | Use Nullish Coalescing for Defaults |
| TS-NUL-004 | Avoid Non-Null Assertions |
| TS-MOD-001 | Barrel File Usage |
| TS-MOD-002 | Order Imports Consistently |
| TS-MOD-003 | Use Type-Only Imports |
| TS-MOD-004 | Use .js Extensions in Imports |
| TS-ASY-003 | Use Promise.all for Independent Operations |
| TS-ASY-004 | Type Async Functions Correctly |
| TS-ERR-001 | Create Typed Error Classes |
| TS-TST-001 | Type Test Fixtures and Mocks |

### Recommended Rules (8 total)

| ID | Rule |
|----|------|
| TS-TYP-006 | Use Meaningful Generic Names |
| TS-TYP-008 | Prefer Built-in Utility Types |
| TS-MOD-005 | Prefer Named Exports |
| TS-ASY-002 | Prefer async/await Over Promise Chains |
| TS-ERR-002 | Consider Result Pattern for Expected Failures |
| TS-TST-002 | Test Type Inference with Explicit Assertions |
| TS-DOC-001 | Document Public API Without Type Duplication |
| TS-PRF-001 | Avoid Overly Complex Type Expressions |
| TS-PRF-002 | Use const Assertions for Immutable Data |

---

## ESLint Rule Mapping

| Rule ID | ESLint Rule |
|---------|-------------|
| TS-TYP-001 | tsconfig: `strict: true` |
| TS-TYP-002 | `@typescript-eslint/no-implicit-any` |
| TS-TYP-003 | `@typescript-eslint/no-explicit-any` |
| TS-TYP-005 | `@typescript-eslint/explicit-function-return-type` |
| TS-NUL-004 | `@typescript-eslint/no-non-null-assertion` |
| TS-MOD-002 | `import/order` |
| TS-MOD-003 | `@typescript-eslint/consistent-type-imports` |
| TS-ASY-001 | `@typescript-eslint/no-floating-promises` |
| TS-ASY-001 | `@typescript-eslint/no-misused-promises` |
| TS-SEC-002 | `@typescript-eslint/no-unsafe-assignment` |
| TS-SEC-002 | `@typescript-eslint/no-unsafe-call` |
