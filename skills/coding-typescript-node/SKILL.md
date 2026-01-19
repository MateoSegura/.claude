---
name: coding-typescript-node
description: TypeScript coding standard for Node.js applications
---

# TypeScript Coding Standard

> **Version**: 1.0.0 | **Status**: Active
> **Base Standards**: Google TypeScript Style Guide, typescript-eslint Recommended

This standard establishes coding conventions for TypeScript development, ensuring consistency, type safety, quality, and maintainability across all TypeScript projects.

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [Type System](rules/types.md) | TS-TYP-* | Strict mode, generics, utility types |
| [Null Handling](rules/null-handling.md) | TS-NUL-* | Optional chaining, nullish coalescing |
| [Modules](rules/modules.md) | TS-MOD-* | Imports, exports, barrel files |
| [Async Patterns](rules/async.md) | TS-ASY-* | Promises, async/await |
| [Error Handling](rules/error-handling.md) | TS-ERR-* | Typed errors, Result pattern |
| [Testing](rules/testing.md) | TS-TST-* | Typed fixtures, mocks |
| [Documentation](rules/documentation.md) | TS-DOC-* | JSDoc without type duplication |
| [Security](rules/security.md) | TS-SEC-* | Input validation, type assertions |
| [Performance](rules/performance.md) | TS-PRF-* | Type complexity, const assertions |

### Tooling

| Tool | File |
|------|------|
| [ESLint](tooling/eslint.md) | ESLint flat config with typescript-eslint |
| [TypeScript Config](tooling/tsconfig.md) | tsconfig.json settings |
| [Prettier](tooling/prettier.md) | Code formatting |
| [Pre-commit](tooling/pre-commit.md) | Git hooks setup |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table with tiers |
| [Code Review Checklist](reference/code-review.md) | Review checklist by category |

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Critical Rules (Always Apply)

These rules are non-negotiable and must be followed in all code.

### TS-TYP-001: Enable Strict Mode :red_circle:

Strict mode provides the strongest type safety guarantees.

```json
// tsconfig.json
{
  "compilerOptions": {
    "strict": true
  }
}
```

### TS-TYP-002: No Implicit Any :red_circle:

Every variable and parameter must have an explicit or inferred type.

```typescript
// Correct
function processItems(items: string[]): number {
  return items.length;
}

// Incorrect
function processItems(items) {  // Parameter 'items' implicitly has 'any' type
  return items.length;
}
```

### TS-TYP-003: Prefer Unknown Over Any :red_circle:

Use `unknown` instead of `any` for truly unknown types, then narrow.

```typescript
// Correct
function safeJsonParse(text: string): unknown {
  return JSON.parse(text);
}

function isUser(value: unknown): value is User {
  return typeof value === 'object' && value !== null && 'id' in value;
}

// Incorrect
function safeJsonParse(text: string): any {
  return JSON.parse(text);  // Bypasses type checking
}
```

### TS-NUL-001: Enable strictNullChecks :red_circle:

Without `strictNullChecks`, `null` and `undefined` are assignable to any type.

```json
{
  "compilerOptions": {
    "strict": true  // Enables strictNullChecks
  }
}
```

### TS-ASY-001: Always Handle Promise Rejections :red_circle:

Every promise must have error handling.

```typescript
// Correct
async function fetchUserData(id: string): Promise<User> {
  try {
    const response = await fetch(`/api/users/${id}`);
    return await response.json();
  } catch (error) {
    logger.error('Failed to fetch user', { id, error });
    throw error;
  }
}

// Incorrect - Unhandled promise
async function saveUser(user: User): Promise<void> {
  fetch('/api/users', { method: 'POST', body: JSON.stringify(user) });
}
```

### TS-ERR-003: Type Unknown Catch Parameters :red_circle:

Always validate error types before accessing properties.

```typescript
// Correct
function getErrorMessage(error: unknown): string {
  if (error instanceof Error) {
    return error.message;
  }
  return 'An unknown error occurred';
}

// Incorrect
try {
  // ...
} catch (error) {
  logger.error(error.message);  // error is unknown
}
```

### TS-SEC-001: Validate External Input at Boundaries :red_circle:

Runtime data from APIs must be validated. Use Zod or similar.

```typescript
// Correct
import { z } from 'zod';

const UserInputSchema = z.object({
  name: z.string().min(1).max(100),
  email: z.string().email(),
});

async function createUser(input: unknown): Promise<User> {
  const validated = UserInputSchema.parse(input);
  return userRepository.create(validated);
}

// Incorrect - Trusting external input
async function createUser(input: UserInput): Promise<User> {
  return userRepository.create(input);  // No runtime validation
}
```

### TS-SEC-002: Avoid Unsafe Type Assertions :red_circle:

Type assertions (`as`) override TypeScript's type checking unsafely.

```typescript
// Correct - Type guard with validation
function isUser(value: unknown): value is User {
  return typeof value === 'object' && value !== null && 'id' in value;
}

function processApiResponse(data: unknown): User {
  if (isUser(data)) {
    return data;
  }
  throw new Error('Invalid user data');
}

// Incorrect
function processApiResponse(data: unknown): User {
  return data as User;  // No validation
}
```

---

## Quick Rule Lookup

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
| TS-NUL-001 | Enable strictNullChecks | Critical |
| TS-NUL-002 | Use Optional Chaining for Nullable Access | Required |
| TS-NUL-003 | Use Nullish Coalescing for Defaults | Required |
| TS-NUL-004 | Avoid Non-Null Assertions | Required |
| TS-MOD-001 | Barrel File Usage | Required |
| TS-MOD-002 | Order Imports Consistently | Required |
| TS-MOD-003 | Use Type-Only Imports | Required |
| TS-MOD-004 | Use .js Extensions in Imports | Required |
| TS-MOD-005 | Prefer Named Exports | Recommended |
| TS-ASY-001 | Always Handle Promise Rejections | Critical |
| TS-ASY-002 | Prefer async/await Over Promise Chains | Recommended |
| TS-ASY-003 | Use Promise.all for Independent Operations | Required |
| TS-ASY-004 | Type Async Functions Correctly | Required |
| TS-ERR-001 | Create Typed Error Classes | Required |
| TS-ERR-002 | Consider Result Pattern for Expected Failures | Recommended |
| TS-ERR-003 | Type Unknown Catch Parameters | Critical |
| TS-TST-001 | Type Test Fixtures and Mocks | Required |
| TS-TST-002 | Test Type Inference with Explicit Assertions | Recommended |
| TS-DOC-001 | Document Public API Without Type Duplication | Recommended |
| TS-SEC-001 | Validate External Input at Boundaries | Critical |
| TS-SEC-002 | Avoid Unsafe Type Assertions | Critical |
| TS-PRF-001 | Avoid Overly Complex Type Expressions | Recommended |
| TS-PRF-002 | Use const Assertions for Immutable Data | Recommended |

---

## References

- [Google TypeScript Style Guide](https://google.github.io/styleguide/tsguide.html)
- [typescript-eslint Rules](https://typescript-eslint.io/rules/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/)
- [TypeScript TSConfig Reference](https://www.typescriptlang.org/tsconfig/)
