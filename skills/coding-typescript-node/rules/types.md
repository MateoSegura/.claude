# Type System Rules (TS-TYP-*)

TypeScript's type system is the foundation of type-safe code. These rules ensure maximum utilization of the type system to catch errors at compile time.

## Type System Strategy

- Enable strict mode for maximum type safety
- Use `unknown` instead of `any` for truly unknown types
- Prefer interfaces for object shapes, types for unions
- Use built-in utility types rather than custom implementations
- Constrain generics appropriately

---

## TS-TYP-001: Enable Strict Mode :red_circle:

**Tier**: Critical

**Rationale**: Strict mode enables a family of compiler options that provide the strongest type safety guarantees. It catches null/undefined errors, implicit any types, and other common mistakes at compile time.

```json
// tsconfig.json - Correct
{
  "compilerOptions": {
    "strict": true
  }
}

// Incorrect - Partial strict settings
{
  "compilerOptions": {
    "strict": false,
    "noImplicitAny": true  // Incomplete coverage
  }
}
```

---

## TS-TYP-002: No Implicit Any :red_circle:

**Tier**: Critical

**Rationale**: Implicit `any` defeats the purpose of TypeScript by creating type-unsafe code. Every variable and parameter must have an explicit or inferred type.

```typescript
// Correct - Explicit types
function processItems(items: string[]): number {
  return items.length;
}

function parseData(input: unknown): ParsedData {
  if (typeof input === 'string') {
    return JSON.parse(input) as ParsedData;
  }
  throw new Error('Invalid input type');
}

// Incorrect - Implicit any
function processItems(items) {  // Parameter 'items' implicitly has 'any' type
  return items.length;
}

function parseData(input) {  // Implicit any
  return JSON.parse(input);
}
```

---

## TS-TYP-003: Prefer Unknown Over Any :red_circle:

**Tier**: Critical

**Rationale**: `unknown` is the type-safe counterpart to `any`. While both accept any value, `unknown` requires type narrowing before use, preventing accidental unsafe operations.

```typescript
// Correct - Using unknown with type guards
function safeJsonParse(text: string): unknown {
  return JSON.parse(text);
}

function processApiResponse(response: unknown): User {
  if (isUser(response)) {
    return response;
  }
  throw new Error('Invalid response format');
}

function isUser(value: unknown): value is User {
  return (
    typeof value === 'object' &&
    value !== null &&
    'id' in value &&
    'name' in value
  );
}

// Incorrect - Using any bypasses type checking
function safeJsonParse(text: string): any {
  return JSON.parse(text);
}

function processApiResponse(response: any): User {
  return response;  // No validation, unsafe
}
```

---

## TS-TYP-004: Use Interface for Object Shapes :yellow_circle:

**Tier**: Required

**Rationale**: Interfaces provide better error messages, support declaration merging, and have better performance with `extends`. Use `type` for unions, intersections, and mapped types.

```typescript
// Correct - Interface for object shapes
interface User {
  id: string;
  name: string;
  email: string;
}

interface AdminUser extends User {
  permissions: string[];
}

// Correct - Type for unions, intersections, utilities
type UserRole = 'admin' | 'editor' | 'viewer';
type UserWithRole = User & { role: UserRole };
type ReadonlyUser = Readonly<User>;
type UserKeys = keyof User;

// Incorrect - Type for simple object shapes
type User = {
  id: string;
  name: string;
  email: string;
};
```

---

## TS-TYP-005: Use Explicit Return Types for Public Functions :yellow_circle:

**Tier**: Required

**Rationale**: Explicit return types serve as documentation, catch unintended return type changes, and improve TypeScript compiler performance.

```typescript
// Correct - Explicit return type on exported function
export function calculateTotal(items: CartItem[]): number {
  return items.reduce((sum, item) => sum + item.price * item.quantity, 0);
}

export async function fetchUser(id: string): Promise<User | null> {
  const response = await api.get(`/users/${id}`);
  return response.data;
}

// Acceptable - Inferred return type for private/internal functions
function formatPrice(amount: number) {
  return `$${amount.toFixed(2)}`;
}

// Incorrect - No return type on exported function
export function calculateTotal(items: CartItem[]) {
  return items.reduce((sum, item) => sum + item.price * item.quantity, 0);
}
```

---

## TS-TYP-006: Use Meaningful Generic Names :green_circle:

**Tier**: Recommended

**Rationale**: Complex generic functions benefit from descriptive names that communicate intent. Single letters are acceptable for simple, obvious cases.

```typescript
// Correct - Descriptive generic names for complex types
function mergeConfigs<TBase extends object, TOverride extends Partial<TBase>>(
  base: TBase,
  override: TOverride
): TBase & TOverride {
  return { ...base, ...override };
}

interface Repository<TEntity, TId = string> {
  findById(id: TId): Promise<TEntity | null>;
  save(entity: TEntity): Promise<TEntity>;
  delete(id: TId): Promise<void>;
}

// Acceptable - Single letter for simple, obvious cases
function identity<T>(value: T): T {
  return value;
}

// Incorrect - Cryptic names in complex scenarios
function mergeConfigs<A, B>(base: A, override: B): A & B {
  return { ...base, ...override };
}
```

---

## TS-TYP-007: Constrain Generic Types Appropriately :yellow_circle:

**Tier**: Required

**Rationale**: Generic constraints prevent misuse and provide better autocomplete. Unconstrained generics accept any type, which may not be the intent.

```typescript
// Correct - Constrained generic
function getProperty<TObj extends object, TKey extends keyof TObj>(
  obj: TObj,
  key: TKey
): TObj[TKey] {
  return obj[key];
}

interface Identifiable {
  id: string;
}

function findById<T extends Identifiable>(items: T[], id: string): T | undefined {
  return items.find(item => item.id === id);
}

// Incorrect - Unconstrained generic allows invalid usage
function getProperty<T, K>(obj: T, key: K) {
  return obj[key];  // Error: Type 'K' cannot be used to index type 'T'
}
```

---

## TS-TYP-008: Prefer Built-in Utility Types :green_circle:

**Tier**: Recommended

**Rationale**: TypeScript provides well-tested utility types that handle edge cases correctly. Custom implementations often miss subtleties.

```typescript
// Correct - Using built-in utility types
interface User {
  id: string;
  name: string;
  email: string;
  createdAt: Date;
}

type UserUpdate = Partial<User>;
type UserCreate = Omit<User, 'id' | 'createdAt'>;
type UserSummary = Pick<User, 'id' | 'name'>;
type ReadonlyUser = Readonly<User>;
type NullableUser = User | null;

function updateUser(id: string, updates: Partial<User>): Promise<User> {
  // ...
}

// Incorrect - Manually recreating utility types
type UserUpdate = {
  id?: string;
  name?: string;
  email?: string;
  createdAt?: Date;
};
```
