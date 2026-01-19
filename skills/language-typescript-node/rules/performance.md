# Performance Rules (TS-PRF-*)

TypeScript's type system can impact both compile-time and runtime performance. These rules ensure efficient type usage and runtime patterns.

## Performance Strategy

- Avoid overly complex type expressions
- Use const assertions for immutable data
- Prefer simple types for better compiler performance
- Consider runtime implications of patterns

---

## TS-PRF-001: Avoid Overly Complex Type Expressions :green_circle:

**Tier**: Recommended

**Rationale**: Complex conditional and mapped types can significantly slow down TypeScript compilation and IDE responsiveness. Prefer simpler types or split complex types into smaller pieces.

```typescript
// Correct - Simple, composable types
interface BaseEntity {
  id: string;
  createdAt: Date;
  updatedAt: Date;
}

interface User extends BaseEntity {
  name: string;
  email: string;
}

type CreateInput<T extends BaseEntity> = Omit<T, keyof BaseEntity>;
type UpdateInput<T extends BaseEntity> = Partial<Omit<T, 'id' | 'createdAt'>>;

// Usage
type CreateUserInput = CreateInput<User>;  // { name: string; email: string }
type UpdateUserInput = UpdateInput<User>;  // { name?: string; email?: string; updatedAt?: Date }

// Correct - Split complex types into steps
type Step1<T> = { [K in keyof T]: T[K] extends object ? Step2<T[K]> : T[K] };
type Step2<T> = Readonly<T>;
type DeepReadonly<T> = Step1<T>;

// Incorrect - Deeply nested conditional types
type DeepPartial<T> = T extends object
  ? T extends Array<infer U>
    ? Array<DeepPartial<U>>
    : T extends Map<infer K, infer V>
      ? Map<K, DeepPartial<V>>
      : T extends Set<infer U>
        ? Set<DeepPartial<U>>
        : T extends Date
          ? T
          : T extends RegExp
            ? T
            : { [K in keyof T]?: DeepPartial<T[K]> }
  : T;

// This can be very slow for large types
type ComplexNested = DeepPartial<VeryLargeType>;
```

---

## TS-PRF-002: Use const Assertions for Immutable Data :green_circle:

**Tier**: Recommended

**Rationale**: `as const` assertions create literal types and readonly properties, enabling better type inference without runtime overhead. They also signal intent that data should not be mutated.

```typescript
// Correct - const assertion for static data
const HTTP_METHODS = ['GET', 'POST', 'PUT', 'DELETE'] as const;
type HttpMethod = typeof HTTP_METHODS[number];  // 'GET' | 'POST' | 'PUT' | 'DELETE'

const STATUS_CODES = {
  OK: 200,
  CREATED: 201,
  BAD_REQUEST: 400,
  NOT_FOUND: 404,
  INTERNAL_ERROR: 500,
} as const;

type StatusCode = typeof STATUS_CODES[keyof typeof STATUS_CODES];
// 200 | 201 | 400 | 404 | 500

// Correct - const assertion in function returns
function getConfig() {
  return {
    apiUrl: 'https://api.example.com',
    timeout: 5000,
    retries: 3,
  } as const;
}

type Config = ReturnType<typeof getConfig>;
// { readonly apiUrl: "https://api.example.com"; readonly timeout: 5000; readonly retries: 3 }

// Correct - Enum-like object with const
const UserRole = {
  ADMIN: 'admin',
  EDITOR: 'editor',
  VIEWER: 'viewer',
} as const;

type UserRole = typeof UserRole[keyof typeof UserRole];  // 'admin' | 'editor' | 'viewer'

// Incorrect - Loses literal types
const HTTP_METHODS = ['GET', 'POST', 'PUT', 'DELETE'];
type HttpMethod = typeof HTTP_METHODS[number];  // string (too wide)

const STATUS_CODES = {
  OK: 200,
  CREATED: 201,
};
type StatusCode = typeof STATUS_CODES[keyof typeof STATUS_CODES];  // number (too wide)
```

---

## Additional Performance Patterns

### Interface vs Type Performance

```typescript
// Interfaces are generally faster for object shapes
// due to caching and simpler equality checks
interface User {
  id: string;
  name: string;
  email: string;
}

// Types are fine for unions and mapped types
type UserRole = 'admin' | 'editor' | 'viewer';
type UserWithRole = User & { role: UserRole };
```

### Lazy Type Evaluation

```typescript
// Correct - Lazy evaluation with conditional types
type LazyEval<T> = T extends infer U ? U : never;

// This delays evaluation until the type is actually used
type ComplexType = LazyEval<ComputeExpensiveType>;
```

### Avoiding Type Recursion Limits

```typescript
// Correct - Bounded recursion with counter
type Decrement = [never, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

type DeepReadonly<T, Depth extends number = 5> = Depth extends 0
  ? T
  : T extends object
    ? { readonly [K in keyof T]: DeepReadonly<T[K], Decrement[Depth]> }
    : T;

// Incorrect - Unbounded recursion
type InfiniteDeepReadonly<T> = T extends object
  ? { readonly [K in keyof T]: InfiniteDeepReadonly<T[K]> }  // May hit recursion limit
  : T;
```

### Runtime Array Methods

```typescript
// Prefer for-of for simple iteration (no intermediate array)
for (const item of items) {
  process(item);
}

// Use reduce only when building a result
const total = items.reduce((sum, item) => sum + item.value, 0);

// Avoid chaining methods when a single pass suffices
// Incorrect - Multiple passes
const result = items
  .filter(x => x.active)
  .map(x => x.value)
  .reduce((a, b) => a + b, 0);

// Correct - Single pass
let result = 0;
for (const item of items) {
  if (item.active) {
    result += item.value;
  }
}
```

### Preallocation Patterns

```typescript
// Correct - Preallocate when size is known
function transformItems<T, U>(items: T[], fn: (item: T) => U): U[] {
  const result: U[] = new Array(items.length);
  for (let i = 0; i < items.length; i++) {
    result[i] = fn(items[i]);
  }
  return result;
}

// Correct - Use map for simple transformations (optimized internally)
const transformed = items.map(item => transformItem(item));
```

### String Building

```typescript
// Correct - Array join for multiple strings
function buildQuery(params: Record<string, string>): string {
  const parts: string[] = [];
  for (const [key, value] of Object.entries(params)) {
    parts.push(`${encodeURIComponent(key)}=${encodeURIComponent(value)}`);
  }
  return parts.join('&');
}

// Correct - Template literals for small concatenations
const greeting = `Hello, ${name}!`;

// Avoid - Repeated string concatenation in loops
let query = '';
for (const [key, value] of Object.entries(params)) {
  query += `${key}=${value}&`;  // Creates new string each iteration
}
```

### Type Inference Helpers

```typescript
// Help TypeScript infer types correctly
function createMap<K extends string, V>(
  entries: readonly (readonly [K, V])[]
): Map<K, V> {
  return new Map(entries);
}

// Usage - types are inferred correctly
const statusMap = createMap([
  ['pending', 0],
  ['active', 1],
  ['completed', 2],
] as const);
```
