# Null Handling Rules (TS-NUL-*)

Proper null handling prevents runtime errors and makes code intent explicit. TypeScript's strict null checks force explicit handling of nullable types.

## Null Handling Strategy

- Enable `strictNullChecks` via strict mode
- Use optional chaining (`?.`) for safe property access
- Use nullish coalescing (`??`) for defaults
- Avoid non-null assertions (`!`) unless absolutely necessary
- Use type guards for explicit narrowing

---

## TS-NUL-001: Enable strictNullChecks :red_circle:

**Tier**: Critical

**Rationale**: Without `strictNullChecks`, `null` and `undefined` are assignable to any type, leading to runtime errors. This option forces explicit handling of nullable types.

```json
// tsconfig.json - Correct
{
  "compilerOptions": {
    "strict": true
    // Enables strictNullChecks automatically
  }
}

// Or explicitly
{
  "compilerOptions": {
    "strictNullChecks": true
  }
}
```

---

## TS-NUL-002: Use Optional Chaining for Nullable Access :yellow_circle:

**Tier**: Required

**Rationale**: Optional chaining (`?.`) provides concise, safe property access on potentially null/undefined values, avoiding verbose null checks.

```typescript
// Correct - Optional chaining
interface User {
  name: string;
  address?: {
    street: string;
    city: string;
  };
}

function getCityName(user: User): string | undefined {
  return user.address?.city;
}

function getFirstItemName(items?: Item[]): string | undefined {
  return items?.[0]?.name;
}

// Correct - Optional method call
const result = obj.method?.();

// Incorrect - Verbose null checking
function getCityName(user: User): string | undefined {
  if (user.address && user.address.city) {
    return user.address.city;
  }
  return undefined;
}

// Incorrect - Non-null assertion without validation
function getCityName(user: User): string {
  return user.address!.city;  // Unsafe: may throw at runtime
}
```

---

## TS-NUL-003: Use Nullish Coalescing for Defaults :yellow_circle:

**Tier**: Required

**Rationale**: Nullish coalescing (`??`) only falls back for `null` or `undefined`, preserving falsy values like `0`, `''`, and `false`. This is usually the intended behavior.

```typescript
// Correct - Nullish coalescing preserves falsy values
interface Config {
  timeout?: number;
  retries?: number;
  enableCache?: boolean;
}

function applyDefaults(config: Config): Required<Config> {
  return {
    timeout: config.timeout ?? 5000,
    retries: config.retries ?? 3,        // 0 retries is valid
    enableCache: config.enableCache ?? true,  // false is valid
  };
}

// Correct - Nullish assignment
let value: string | undefined;
value ??= 'default';

// Incorrect - Logical OR treats 0 and false as falsy
function applyDefaults(config: Config): Required<Config> {
  return {
    timeout: config.timeout || 5000,
    retries: config.retries || 3,        // Bug: 0 becomes 3
    enableCache: config.enableCache || true,  // Bug: false becomes true
  };
}
```

---

## TS-NUL-004: Avoid Non-Null Assertions :yellow_circle:

**Tier**: Required

**Rationale**: The non-null assertion operator (`!`) bypasses TypeScript's null checking without runtime validation. Use it only when you can prove the value is non-null through context TypeScript cannot analyze.

```typescript
// Correct - Type guards and narrowing
function processUser(user: User | null): void {
  if (user === null) {
    throw new Error('User is required');
  }
  // TypeScript knows user is non-null here
  console.log(user.name);
}

// Correct - Early return pattern
function getUserName(user: User | null): string {
  if (!user) {
    return 'Unknown';
  }
  return user.name;
}

// Acceptable - Non-null assertion with documented invariant
class UserCache {
  private cache = new Map<string, User>();

  has(id: string): boolean {
    return this.cache.has(id);
  }

  // Caller must ensure user exists via has() before calling
  getExisting(id: string): User {
    // Invariant: Only called after has(id) returns true
    return this.cache.get(id)!;
  }
}

// Incorrect - Blind non-null assertion
function processUser(user: User | null): void {
  console.log(user!.name);  // Unsafe: crashes if null
}

// Incorrect - Assertion in DOM queries without check
const button = document.querySelector('#submit')!;  // May be null
```

---

## Additional Patterns

### Type Narrowing with Type Guards

```typescript
// Type predicate for custom narrowing
function isDefined<T>(value: T | null | undefined): value is T {
  return value !== null && value !== undefined;
}

// Usage
const items: (string | null)[] = ['a', null, 'b'];
const defined = items.filter(isDefined);  // string[]
```

### Discriminated Unions for Null States

```typescript
// Explicit state handling
type LoadState<T> =
  | { status: 'loading' }
  | { status: 'success'; data: T }
  | { status: 'error'; error: Error };

function handleState(state: LoadState<User>): void {
  switch (state.status) {
    case 'loading':
      showSpinner();
      break;
    case 'success':
      showUser(state.data);  // data is available
      break;
    case 'error':
      showError(state.error);  // error is available
      break;
  }
}
```

### Required vs Optional Properties

```typescript
// Clear intent with optional properties
interface UserProfile {
  id: string;           // Always required
  name: string;         // Always required
  email?: string;       // Optional - may be undefined
  bio?: string | null;  // Optional - may be undefined or null
}

// Using Required/Partial for variations
type CreateUserInput = Required<Pick<UserProfile, 'name' | 'email'>>;
type UpdateUserInput = Partial<Omit<UserProfile, 'id'>>;
```
