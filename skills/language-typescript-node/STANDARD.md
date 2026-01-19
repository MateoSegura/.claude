# TypeScript Coding Standard

> **Version**: 1.0.0
> **Status**: Active
> **Base Standard**: Google TypeScript Style Guide, typescript-eslint Recommended
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes coding conventions for TypeScript development. It ensures:

- **Consistency**: Uniform code style across all TypeScript projects
- **Type Safety**: Maximum utilization of TypeScript's type system to catch errors at compile time
- **Quality**: Reduced defects through proven patterns and strict compiler settings
- **Maintainability**: Code that is easy to read, modify, and extend

### 1.2 Scope

This standard applies to:
- [x] All TypeScript source files in production repositories
- [x] Test code (with documented exceptions)
- [x] Build scripts and tooling written in TypeScript
- [x] Configuration files (tsconfig.json, eslint.config.mjs)

### 1.3 Audience

- Software engineers writing TypeScript code
- Code reviewers evaluating TypeScript changes
- Tech leads establishing project standards
- DevOps engineers configuring CI/CD pipelines

### 1.4 Relationship to Industry Standards

| Standard | Relationship |
|----------|--------------|
| [Google TypeScript Style Guide](https://google.github.io/styleguide/tsguide.html) | **Base** - Core rules apply unless documented otherwise |
| [typescript-eslint Recommended](https://typescript-eslint.io/rules/) | **Enforcement** - Rules enforced via ESLint configuration |
| [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/) | **Reference** - Official language documentation |

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

Each rule includes:
- **Rule ID**: Unique identifier (e.g., `TS-TYP-001`)
- **Tier**: Enforcement level
- **Rationale**: Why this rule exists
- **Example**: Correct and incorrect code

### Rule ID Format

Rule IDs follow the pattern `TS-XXX-NNN` where:
- `TS` = TypeScript
- `XXX` = Category code (TYP=Types, NUL=Null Handling, MOD=Modules, ASY=Async, ERR=Errors, TST=Testing, FMT=Formatting, NAM=Naming, SEC=Security)
- `NNN` = Numeric identifier

---

## 3. Project Structure

### 3.1 Directory Layout

```
project-root/
├── src/
│   ├── index.ts              # Main entry point
│   ├── types/                # Shared type definitions
│   │   └── index.ts          # Type barrel file
│   ├── utils/                # Utility functions
│   ├── services/             # Business logic services
│   ├── models/               # Data models and entities
│   └── errors/               # Custom error classes
├── tests/
│   ├── unit/                 # Unit tests
│   ├── integration/          # Integration tests
│   └── fixtures/             # Test data and mocks
├── dist/                     # Compiled output (gitignored)
├── tsconfig.json             # TypeScript configuration
├── eslint.config.mjs         # ESLint flat config
└── package.json
```

### 3.2 File Naming

| Type | Convention | Example |
|------|------------|---------|
| Source files | kebab-case | `user-service.ts` |
| Test files | kebab-case with .test suffix | `user-service.test.ts` |
| Type definition files | kebab-case | `api-types.ts` |
| Constants | kebab-case | `http-constants.ts` |
| React components | PascalCase | `UserProfile.tsx` |

### 3.3 Module Organization

#### TS-MOD-001: Barrel File Usage :yellow_circle:

**Rationale**: Barrel files (index.ts re-exports) simplify imports but can cause circular dependencies and tree-shaking issues. Use them judiciously for public APIs only.

```typescript
// ✅ Correct - Barrel file for public API
// src/types/index.ts
export type { User, UserRole } from './user.js';
export type { Order, OrderStatus } from './order.js';

// Consumer
import type { User, Order } from './types/index.js';

// ❌ Incorrect - Deep barrel nesting causing circular imports
// src/index.ts
export * from './services/index.js';
export * from './models/index.js';  // models imports services
export * from './utils/index.js';   // utils imports models
```

---

## 4. Type System Rules

### 4.1 Strict Mode Configuration

#### TS-TYP-001: Enable Strict Mode :red_circle:

**Rationale**: Strict mode enables a family of compiler options that provide the strongest type safety guarantees. It catches null/undefined errors, implicit any types, and other common mistakes at compile time rather than runtime.

```json
// ✅ Correct - tsconfig.json
{
  "compilerOptions": {
    "strict": true
  }
}

// ❌ Incorrect - Partial strict settings or disabled
{
  "compilerOptions": {
    "strict": false,
    "noImplicitAny": true  // Incomplete coverage
  }
}
```

#### TS-TYP-002: No Implicit Any :red_circle:

**Rationale**: Implicit `any` defeats the purpose of TypeScript by creating type-unsafe code. Every variable and parameter must have an explicit or inferred type. The `any` type disables all type checking for that value.

```typescript
// ✅ Correct - Explicit types
function processItems(items: string[]): number {
  return items.length;
}

function parseData(input: unknown): ParsedData {
  // Use unknown for truly unknown types, then narrow
  if (typeof input === 'string') {
    return JSON.parse(input) as ParsedData;
  }
  throw new Error('Invalid input type');
}

// ❌ Incorrect - Implicit any
function processItems(items) {  // Parameter 'items' implicitly has 'any' type
  return items.length;
}

function parseData(input) {  // Implicit any
  return JSON.parse(input);
}
```

#### TS-TYP-003: Prefer Unknown Over Any :red_circle:

**Rationale**: `unknown` is the type-safe counterpart to `any`. While both accept any value, `unknown` requires type narrowing before use, preventing accidental unsafe operations.

```typescript
// ✅ Correct - Using unknown with type guards
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

// ❌ Incorrect - Using any bypasses type checking
function safeJsonParse(text: string): any {
  return JSON.parse(text);
}

function processApiResponse(response: any): User {
  return response;  // No validation, unsafe
}
```

### 4.2 Type vs Interface

#### TS-TYP-004: Use Interface for Object Shapes :yellow_circle:

**Rationale**: Interfaces provide better error messages, support declaration merging for library extension, and have better performance for complex type hierarchies with `extends`. Use `type` for unions, intersections, mapped types, and primitives.

```typescript
// ✅ Correct - Interface for object shapes
interface User {
  id: string;
  name: string;
  email: string;
}

interface AdminUser extends User {
  permissions: string[];
}

// ✅ Correct - Type for unions, intersections, utilities
type UserRole = 'admin' | 'editor' | 'viewer';
type UserWithRole = User & { role: UserRole };
type ReadonlyUser = Readonly<User>;
type UserKeys = keyof User;

// ❌ Incorrect - Type for simple object shapes
type User = {
  id: string;
  name: string;
  email: string;
};

// ❌ Incorrect - Interface for union types (not possible)
// interface UserRole = 'admin' | 'editor';  // Syntax error
```

#### TS-TYP-005: Use Explicit Return Types for Public Functions :yellow_circle:

**Rationale**: Explicit return types serve as documentation, catch unintended return type changes, and improve TypeScript compiler performance by reducing inference work.

```typescript
// ✅ Correct - Explicit return type on exported function
export function calculateTotal(items: CartItem[]): number {
  return items.reduce((sum, item) => sum + item.price * item.quantity, 0);
}

export async function fetchUser(id: string): Promise<User | null> {
  const response = await api.get(`/users/${id}`);
  return response.data;
}

// ✅ Acceptable - Inferred return type for private/internal functions
function formatPrice(amount: number) {
  return `$${amount.toFixed(2)}`;
}

// ❌ Incorrect - No return type on exported function
export function calculateTotal(items: CartItem[]) {
  return items.reduce((sum, item) => sum + item.price * item.quantity, 0);
}
```

### 4.3 Generic Types

#### TS-TYP-006: Use Meaningful Generic Names :green_circle:

**Rationale**: While single letters like `T` are conventional for simple generics, complex generic functions benefit from descriptive names that communicate intent.

```typescript
// ✅ Correct - Descriptive generic names for complex types
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

// ✅ Acceptable - Single letter for simple, obvious cases
function identity<T>(value: T): T {
  return value;
}

// ❌ Incorrect - Cryptic names in complex scenarios
function mergeConfigs<A, B>(base: A, override: B): A & B {
  return { ...base, ...override };
}
```

#### TS-TYP-007: Constrain Generic Types Appropriately :yellow_circle:

**Rationale**: Generic constraints prevent misuse and provide better autocomplete. Unconstrained generics accept any type, which may not be the intent.

```typescript
// ✅ Correct - Constrained generic
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

// ❌ Incorrect - Unconstrained generic allows invalid usage
function getProperty<T, K>(obj: T, key: K) {
  return obj[key];  // Error: Type 'K' cannot be used to index type 'T'
}
```

### 4.4 Utility Types

#### TS-TYP-008: Prefer Built-in Utility Types :green_circle:

**Rationale**: TypeScript provides well-tested utility types that handle edge cases correctly. Custom implementations often miss subtleties like readonly modifiers or optional properties.

```typescript
// ✅ Correct - Using built-in utility types
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

// Function parameters
function updateUser(id: string, updates: Partial<User>): Promise<User> {
  // ...
}

// ❌ Incorrect - Manually recreating utility types
type UserUpdate = {
  id?: string;
  name?: string;
  email?: string;
  createdAt?: Date;
};
```

---

## 5. Null Handling

### 5.1 Strict Null Checks

#### TS-NUL-001: Enable strictNullChecks :red_circle:

**Rationale**: Without `strictNullChecks`, `null` and `undefined` are assignable to any type, leading to runtime errors. This option forces explicit handling of nullable types.

```json
// ✅ Correct - Enabled via strict: true or explicitly
{
  "compilerOptions": {
    "strict": true
    // or explicitly: "strictNullChecks": true
  }
}
```

#### TS-NUL-002: Use Optional Chaining for Nullable Access :yellow_circle:

**Rationale**: Optional chaining (`?.`) provides concise, safe property access on potentially null/undefined values, avoiding verbose null checks.

```typescript
// ✅ Correct - Optional chaining
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

// ❌ Incorrect - Verbose null checking
function getCityName(user: User): string | undefined {
  if (user.address && user.address.city) {
    return user.address.city;
  }
  return undefined;
}

// ❌ Incorrect - Non-null assertion without validation
function getCityName(user: User): string {
  return user.address!.city;  // Unsafe: may throw at runtime
}
```

#### TS-NUL-003: Use Nullish Coalescing for Defaults :yellow_circle:

**Rationale**: Nullish coalescing (`??`) only falls back for `null` or `undefined`, preserving falsy values like `0`, `''`, and `false`. This is usually the intended behavior for defaults.

```typescript
// ✅ Correct - Nullish coalescing preserves falsy values
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

// ❌ Incorrect - Logical OR treats 0 and false as falsy
function applyDefaults(config: Config): Required<Config> {
  return {
    timeout: config.timeout || 5000,
    retries: config.retries || 3,        // Bug: 0 becomes 3
    enableCache: config.enableCache || true,  // Bug: false becomes true
  };
}
```

#### TS-NUL-004: Avoid Non-Null Assertions :yellow_circle:

**Rationale**: The non-null assertion operator (`!`) bypasses TypeScript's null checking without runtime validation. It should only be used when you can prove the value is non-null through context TypeScript cannot analyze.

```typescript
// ✅ Correct - Type guards and narrowing
function processUser(user: User | null): void {
  if (user === null) {
    throw new Error('User is required');
  }
  // TypeScript knows user is non-null here
  console.log(user.name);
}

// ✅ Acceptable - Non-null assertion with documented invariant
class UserCache {
  private cache = new Map<string, User>();

  // Caller must ensure user exists via has() before calling
  getExisting(id: string): User {
    // Invariant: Only called after has(id) returns true
    return this.cache.get(id)!;
  }
}

// ❌ Incorrect - Blind non-null assertion
function processUser(user: User | null): void {
  console.log(user!.name);  // Unsafe: crashes if null
}
```

---

## 6. Module Organization

### 6.1 Import Organization

#### TS-MOD-002: Order Imports Consistently :yellow_circle:

**Rationale**: Consistent import ordering improves readability and reduces merge conflicts. Group imports by external dependencies, internal modules, and types.

```typescript
// ✅ Correct - Organized imports
// 1. Node.js built-ins
import { readFile } from 'node:fs/promises';
import path from 'node:path';

// 2. External dependencies
import express from 'express';
import { z } from 'zod';

// 3. Internal modules (absolute paths)
import { UserService } from '@/services/user-service.js';
import { formatDate } from '@/utils/date.js';

// 4. Internal modules (relative paths)
import { validateInput } from './validation.js';

// 5. Type-only imports
import type { User, UserRole } from '@/types/user.js';
import type { Request, Response } from 'express';

// ❌ Incorrect - Unorganized imports
import type { User } from '@/types/user.js';
import express from 'express';
import { readFile } from 'node:fs/promises';
import { UserService } from '@/services/user-service.js';
import type { Request } from 'express';
import { z } from 'zod';
```

#### TS-MOD-003: Use Type-Only Imports :yellow_circle:

**Rationale**: Type-only imports are erased at compile time and clearly communicate that an import is used only for type information. This can improve build performance and prevents runtime side effects.

```typescript
// ✅ Correct - Separate type imports
import { UserService } from './user-service.js';
import type { User, UserCreateInput } from './types.js';

// ✅ Correct - Inline type imports (TypeScript 4.5+)
import { UserService, type User, type UserCreateInput } from './user.js';

// ❌ Incorrect - Mixed without type modifier
import { UserService, User, UserCreateInput } from './user.js';
// (If User and UserCreateInput are only used as types)
```

#### TS-MOD-004: Use .js Extensions in Imports :yellow_circle:

**Rationale**: When targeting ESM output, TypeScript requires `.js` extensions in imports even for `.ts` source files. This ensures the compiled output has correct import paths.

```typescript
// ✅ Correct - .js extension for ESM compatibility
import { helper } from './helper.js';
import { User } from '../types/user.js';

// ❌ Incorrect - Missing extension (fails in ESM)
import { helper } from './helper';
import { User } from '../types/user';
```

### 6.2 Export Patterns

#### TS-MOD-005: Prefer Named Exports :green_circle:

**Rationale**: Named exports provide better IDE support (autocomplete, refactoring), prevent inconsistent naming across imports, and enable tree-shaking.

```typescript
// ✅ Correct - Named exports
// user-service.ts
export class UserService {
  // ...
}

export function createUser(data: UserCreateInput): User {
  // ...
}

// Consumer
import { UserService, createUser } from './user-service.js';

// ❌ Incorrect - Default export allows inconsistent naming
// user-service.ts
export default class UserService {
  // ...
}

// Consumer - different names for same thing
import UserSvc from './user-service.js';
import UserManager from './user-service.js';
```

---

## 7. Async Patterns

### 7.1 Promise Handling

#### TS-ASY-001: Always Handle Promise Rejections :red_circle:

**Rationale**: Unhandled promise rejections can crash Node.js applications and cause silent failures in browsers. Every promise must have error handling.

```typescript
// ✅ Correct - async/await with try-catch
async function fetchUserData(id: string): Promise<User> {
  try {
    const response = await fetch(`/api/users/${id}`);
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    logger.error('Failed to fetch user', { id, error });
    throw error;  // Re-throw or handle appropriately
  }
}

// ✅ Correct - Promise chain with .catch()
function fetchUserData(id: string): Promise<User> {
  return fetch(`/api/users/${id}`)
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error: ${response.status}`);
      }
      return response.json();
    })
    .catch(error => {
      logger.error('Failed to fetch user', { id, error });
      throw error;
    });
}

// ❌ Incorrect - Unhandled promise
async function saveUser(user: User): Promise<void> {
  fetch('/api/users', {
    method: 'POST',
    body: JSON.stringify(user),
  });  // Promise not awaited, errors lost
}
```

#### TS-ASY-002: Prefer async/await Over Promise Chains :green_circle:

**Rationale**: `async/await` syntax is more readable, easier to debug (stack traces), and less prone to errors than nested `.then()` chains. It also enables standard try-catch-finally error handling.

```typescript
// ✅ Correct - async/await with sequential operations
async function processOrder(orderId: string): Promise<OrderResult> {
  const order = await fetchOrder(orderId);
  const inventory = await checkInventory(order.items);

  if (!inventory.available) {
    return { success: false, reason: 'Out of stock' };
  }

  const payment = await processPayment(order);
  const shipment = await createShipment(order);

  return { success: true, shipmentId: shipment.id };
}

// ❌ Incorrect - Nested promise chains
function processOrder(orderId: string): Promise<OrderResult> {
  return fetchOrder(orderId)
    .then(order => {
      return checkInventory(order.items)
        .then(inventory => {
          if (!inventory.available) {
            return { success: false, reason: 'Out of stock' };
          }
          return processPayment(order)
            .then(payment => {
              return createShipment(order)
                .then(shipment => {
                  return { success: true, shipmentId: shipment.id };
                });
            });
        });
    });
}
```

#### TS-ASY-003: Use Promise.all for Independent Operations :yellow_circle:

**Rationale**: When multiple async operations don't depend on each other, running them concurrently with `Promise.all` improves performance significantly.

```typescript
// ✅ Correct - Concurrent independent operations
async function getDashboardData(userId: string): Promise<Dashboard> {
  const [user, orders, notifications] = await Promise.all([
    fetchUser(userId),
    fetchOrders(userId),
    fetchNotifications(userId),
  ]);

  return { user, orders, notifications };
}

// ✅ Correct - Promise.allSettled for operations that may fail independently
async function sendNotifications(userIds: string[]): Promise<NotificationResults> {
  const results = await Promise.allSettled(
    userIds.map(id => sendNotification(id))
  );

  return {
    succeeded: results.filter(r => r.status === 'fulfilled').length,
    failed: results.filter(r => r.status === 'rejected').length,
  };
}

// ❌ Incorrect - Sequential operations that could be parallel
async function getDashboardData(userId: string): Promise<Dashboard> {
  const user = await fetchUser(userId);          // Waits
  const orders = await fetchOrders(userId);      // Then waits
  const notifications = await fetchNotifications(userId);  // Then waits

  return { user, orders, notifications };
}
```

### 7.2 Async Function Types

#### TS-ASY-004: Type Async Functions Correctly :yellow_circle:

**Rationale**: Async functions always return a Promise. The return type annotation should reflect the unwrapped value type, and the function signature should use `Promise<T>`.

```typescript
// ✅ Correct - Properly typed async function
async function fetchUser(id: string): Promise<User | null> {
  const response = await fetch(`/api/users/${id}`);
  if (response.status === 404) {
    return null;
  }
  return response.json();
}

// ✅ Correct - Async arrow function
const fetchUser = async (id: string): Promise<User | null> => {
  // ...
};

// ✅ Correct - Interface with async method
interface UserRepository {
  findById(id: string): Promise<User | null>;
  save(user: User): Promise<User>;
  delete(id: string): Promise<void>;
}

// ❌ Incorrect - Missing Promise wrapper
async function fetchUser(id: string): User | null {  // Error
  // ...
}
```

---

## 8. Error Handling

### 8.1 Error Types

#### TS-ERR-001: Create Typed Error Classes :yellow_circle:

**Rationale**: Custom error classes enable type-safe error handling with `instanceof` checks. They can carry additional context and enable exhaustive error handling.

```typescript
// ✅ Correct - Typed error hierarchy
abstract class AppError extends Error {
  abstract readonly code: string;
  abstract readonly statusCode: number;

  constructor(message: string, public readonly cause?: Error) {
    super(message);
    this.name = this.constructor.name;
    Error.captureStackTrace(this, this.constructor);
  }
}

class NotFoundError extends AppError {
  readonly code = 'NOT_FOUND';
  readonly statusCode = 404;

  constructor(resource: string, id: string, cause?: Error) {
    super(`${resource} with id '${id}' not found`, cause);
  }
}

class ValidationError extends AppError {
  readonly code = 'VALIDATION_ERROR';
  readonly statusCode = 400;

  constructor(
    message: string,
    public readonly field: string,
    public readonly value: unknown,
    cause?: Error
  ) {
    super(message, cause);
  }
}

// Usage with type narrowing
function handleError(error: unknown): Response {
  if (error instanceof NotFoundError) {
    return { status: error.statusCode, body: { code: error.code } };
  }
  if (error instanceof ValidationError) {
    return {
      status: error.statusCode,
      body: { code: error.code, field: error.field }
    };
  }
  // Unknown error
  return { status: 500, body: { code: 'INTERNAL_ERROR' } };
}

// ❌ Incorrect - Plain Error with no type information
function fetchUser(id: string): Promise<User> {
  throw new Error('User not found');  // No context, hard to handle
}
```

### 8.2 Result Pattern

#### TS-ERR-002: Consider Result Pattern for Expected Failures :green_circle:

**Rationale**: The Result pattern makes error handling explicit in the type system. Unlike exceptions, Result types cannot be accidentally ignored and enable exhaustive error handling. Use for expected, recoverable failures.

```typescript
// ✅ Correct - Result type implementation
type Result<T, E = Error> =
  | { ok: true; value: T }
  | { ok: false; error: E };

// Helper functions
function ok<T>(value: T): Result<T, never> {
  return { ok: true, value };
}

function err<E>(error: E): Result<never, E> {
  return { ok: false, error };
}

// Usage
interface ParseError {
  message: string;
  line: number;
  column: number;
}

function parseConfig(text: string): Result<Config, ParseError> {
  try {
    const config = JSON.parse(text);
    if (!isValidConfig(config)) {
      return err({ message: 'Invalid config structure', line: 1, column: 1 });
    }
    return ok(config);
  } catch (e) {
    return err({ message: 'Invalid JSON', line: 1, column: 1 });
  }
}

// Consumer is forced to handle both cases
const result = parseConfig(configText);
if (result.ok) {
  console.log('Config loaded:', result.value);
} else {
  console.error('Parse error:', result.error.message);
}

// ❌ Incorrect - Throwing for expected failures
function parseConfig(text: string): Config {
  try {
    const config = JSON.parse(text);
    if (!isValidConfig(config)) {
      throw new Error('Invalid config');  // Caller might forget to catch
    }
    return config;
  } catch (e) {
    throw new Error('Parse failed');  // Type system doesn't track this
  }
}
```

### 8.3 Error Handling Patterns

#### TS-ERR-003: Type Unknown Catch Parameters :red_circle:

**Rationale**: In TypeScript 4.4+, catch clause variables are `unknown` by default with `useUnknownInCatchVariables`. Always validate error types before accessing properties.

```typescript
// ✅ Correct - Type guard for errors
function isError(value: unknown): value is Error {
  return value instanceof Error;
}

function getErrorMessage(error: unknown): string {
  if (isError(error)) {
    return error.message;
  }
  if (typeof error === 'string') {
    return error;
  }
  return 'An unknown error occurred';
}

async function fetchData(): Promise<Data> {
  try {
    const response = await fetch('/api/data');
    return await response.json();
  } catch (error) {
    // Properly handle unknown type
    const message = getErrorMessage(error);
    logger.error('Fetch failed', { error: message });
    throw new FetchError(message, isError(error) ? error : undefined);
  }
}

// ❌ Incorrect - Assuming error is Error type
async function fetchData(): Promise<Data> {
  try {
    const response = await fetch('/api/data');
    return await response.json();
  } catch (error) {
    // error is unknown, accessing .message is unsafe
    logger.error(error.message);  // TypeScript error (or runtime error)
    throw error;
  }
}
```

---

## 9. Testing Requirements

### 9.1 Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% | 85% |
| Integration tests | 50% | 70% |
| Critical paths | 90% | 100% |

### 9.2 Test Organization

```
tests/
├── unit/
│   ├── services/
│   │   └── user-service.test.ts
│   └── utils/
│       └── validation.test.ts
├── integration/
│   └── api/
│       └── users.test.ts
└── fixtures/
    ├── users.ts
    └── orders.ts
```

### 9.3 Test Naming

| Element | Convention | Example |
|---------|------------|---------|
| Test files | `*.test.ts` or `*.spec.ts` | `user-service.test.ts` |
| Test suites | Describe the unit under test | `describe('UserService', ...)` |
| Test cases | Describe behavior | `it('should return null for unknown user', ...)` |
| Test fixtures | Descriptive names | `createMockUser()`, `validOrderData` |

### 9.4 Testing Rules

#### TS-TST-001: Type Test Fixtures and Mocks :yellow_circle:

**Rationale**: Typed test fixtures catch errors when the source types change. Untyped mocks can drift from real implementations, causing false positives.

```typescript
// ✅ Correct - Typed test fixtures
import type { User, Order } from '@/types/index.js';

function createMockUser(overrides?: Partial<User>): User {
  return {
    id: 'user-123',
    name: 'Test User',
    email: 'test@example.com',
    createdAt: new Date('2024-01-01'),
    ...overrides,
  };
}

// Type-safe mock implementation
const mockUserRepository: jest.Mocked<UserRepository> = {
  findById: jest.fn(),
  save: jest.fn(),
  delete: jest.fn(),
};

describe('UserService', () => {
  it('should return user when found', async () => {
    const expectedUser = createMockUser({ name: 'Alice' });
    mockUserRepository.findById.mockResolvedValue(expectedUser);

    const service = new UserService(mockUserRepository);
    const result = await service.getUser('user-123');

    expect(result).toEqual(expectedUser);
  });
});

// ❌ Incorrect - Untyped fixtures
const mockUser = {
  id: 'user-123',
  name: 'Test User',
  // Missing required fields - no TypeScript error
};

const mockRepo = {
  findById: jest.fn(),
  // Missing methods - no TypeScript error
};
```

#### TS-TST-002: Test Type Inference with Explicit Assertions :green_circle:

**Rationale**: Type tests verify that functions return expected types. Use explicit type annotations or type assertions to catch type regressions.

```typescript
// ✅ Correct - Explicit type assertions in tests
import { expectTypeOf } from 'vitest';

describe('UserService types', () => {
  it('should return User or null from findById', async () => {
    const service = new UserService(mockRepository);
    const result = await service.findById('123');

    // Type-level test
    expectTypeOf(result).toEqualTypeOf<User | null>();
  });

  it('should accept partial user for updates', () => {
    const service = new UserService(mockRepository);

    // This should compile - verifies the type accepts Partial<User>
    service.update('123', { name: 'New Name' });
    service.update('123', { email: 'new@example.com' });
    service.update('123', {});  // Empty partial is valid
  });
});
```

---

## 10. Documentation

### 10.1 Comment Style

Use JSDoc comments for documentation. TypeScript types should not be duplicated in JSDoc tags.

### 10.2 Documentation Requirements

| Element | Documentation Required |
|---------|----------------------|
| Public functions | Yes - describe purpose and behavior |
| Public types/interfaces | Yes - describe purpose |
| Complex logic | Yes - explain why, not what |
| Private internals | Optional - only if complex |

### 10.3 Documentation Rules

#### TS-DOC-001: Document Public API Without Type Duplication :green_circle:

**Rationale**: JSDoc comments complement TypeScript types by explaining intent and behavior. Do not duplicate type information already expressed in the code.

```typescript
// ✅ Correct - Describes behavior without duplicating types
/**
 * Retrieves a user by their unique identifier.
 * Returns null if the user does not exist or has been deactivated.
 *
 * @example
 * const user = await userService.findById('user-123');
 * if (user) {
 *   console.log(user.name);
 * }
 */
export async function findById(id: string): Promise<User | null> {
  // ...
}

/**
 * Validates that the email address is properly formatted and not already
 * registered in the system. Used during registration and email change flows.
 */
export function validateEmail(email: string): ValidationResult {
  // ...
}

// ❌ Incorrect - Duplicates type information
/**
 * @param {string} id - The user id (string)
 * @returns {Promise<User | null>} - A promise that resolves to User or null
 */
export async function findById(id: string): Promise<User | null> {
  // ...
}
```

---

## 11. Security Considerations

### 11.1 Input Validation

#### TS-SEC-001: Validate External Input at Boundaries :red_circle:

**Rationale**: TypeScript types only exist at compile time. Runtime data from APIs, user input, and file I/O must be validated. Use runtime validation libraries like Zod for type-safe validation.

```typescript
// ✅ Correct - Runtime validation with Zod
import { z } from 'zod';

const UserInputSchema = z.object({
  name: z.string().min(1).max(100),
  email: z.string().email(),
  age: z.number().int().min(0).max(150).optional(),
});

type UserInput = z.infer<typeof UserInputSchema>;

async function createUser(input: unknown): Promise<User> {
  // Validate and parse at the boundary
  const validated = UserInputSchema.parse(input);

  // validated is now typed as UserInput
  return userRepository.create(validated);
}

// API handler
app.post('/users', async (req, res) => {
  try {
    const user = await createUser(req.body);
    res.json(user);
  } catch (error) {
    if (error instanceof z.ZodError) {
      res.status(400).json({ errors: error.errors });
    } else {
      res.status(500).json({ error: 'Internal server error' });
    }
  }
});

// ❌ Incorrect - Trusting external input types
interface UserInput {
  name: string;
  email: string;
}

async function createUser(input: UserInput): Promise<User> {
  // TypeScript is satisfied but input is not actually validated
  return userRepository.create(input);
}

app.post('/users', async (req, res) => {
  // req.body could be anything at runtime
  const user = await createUser(req.body as UserInput);  // Unsafe cast
  res.json(user);
});
```

### 11.2 Type Assertions

#### TS-SEC-002: Avoid Unsafe Type Assertions :red_circle:

**Rationale**: Type assertions (`as`) override TypeScript's type checking. Using them incorrectly can introduce type-safety bugs that won't be caught until runtime.

```typescript
// ✅ Correct - Type guard with validation
function isUser(value: unknown): value is User {
  return (
    typeof value === 'object' &&
    value !== null &&
    'id' in value &&
    typeof (value as { id: unknown }).id === 'string' &&
    'name' in value &&
    typeof (value as { name: unknown }).name === 'string'
  );
}

function processApiResponse(data: unknown): User {
  if (isUser(data)) {
    return data;  // Type-safe narrowing
  }
  throw new Error('Invalid user data');
}

// ❌ Incorrect - Unsafe type assertion
function processApiResponse(data: unknown): User {
  return data as User;  // No validation, crashes if wrong
}

// ❌ Incorrect - Double assertion to bypass checks
function hackTheType(value: string): number {
  return value as unknown as number;  // Nonsensical, will fail
}
```

---

## 12. Performance Guidelines

### 12.1 Type Performance

#### TS-PRF-001: Avoid Overly Complex Type Expressions :green_circle:

**Rationale**: Complex conditional and mapped types can significantly slow down TypeScript compilation and IDE responsiveness. Prefer simpler types or split complex types into smaller pieces.

```typescript
// ✅ Correct - Simple, composable types
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

// ❌ Incorrect - Deeply nested conditional types
type DeepPartial<T> = T extends object
  ? T extends Array<infer U>
    ? Array<DeepPartial<U>>
    : T extends Map<infer K, infer V>
      ? Map<K, DeepPartial<V>>
      : T extends Set<infer U>
        ? Set<DeepPartial<U>>
        : { [K in keyof T]?: DeepPartial<T[K]> }
  : T;

// This can be slow for large types
type ComplexNested = DeepPartial<VeryLargeType>;
```

### 12.2 Runtime Performance

#### TS-PRF-002: Use const Assertions for Immutable Data :green_circle:

**Rationale**: `as const` assertions create literal types and readonly properties, enabling better type inference without runtime overhead. They also signal intent that data should not be mutated.

```typescript
// ✅ Correct - const assertion for static data
const HTTP_METHODS = ['GET', 'POST', 'PUT', 'DELETE'] as const;
type HttpMethod = typeof HTTP_METHODS[number];  // 'GET' | 'POST' | 'PUT' | 'DELETE'

const STATUS_CODES = {
  OK: 200,
  CREATED: 201,
  BAD_REQUEST: 400,
  NOT_FOUND: 404,
} as const;
type StatusCode = typeof STATUS_CODES[keyof typeof STATUS_CODES];  // 200 | 201 | 400 | 404

// ❌ Incorrect - Loses literal types
const HTTP_METHODS = ['GET', 'POST', 'PUT', 'DELETE'];
type HttpMethod = typeof HTTP_METHODS[number];  // string (too wide)

const STATUS_CODES = {
  OK: 200,
  CREATED: 201,
};
type StatusCode = typeof STATUS_CODES[keyof typeof STATUS_CODES];  // number (too wide)
```

---

## 13. Tooling Configuration

### 13.1 Required Tools

| Tool | Purpose | Version |
|------|---------|---------|
| TypeScript | Type checking and compilation | >=5.0 |
| ESLint | Linting | >=9.0 |
| typescript-eslint | TypeScript ESLint integration | >=8.0 |
| Prettier | Code formatting | >=3.0 |

### 13.2 TypeScript Configuration

```json
// tsconfig.json
{
  "compilerOptions": {
    // Type Checking - Maximum strictness
    "strict": true,
    "noUncheckedIndexedAccess": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "exactOptionalPropertyTypes": true,
    "useUnknownInCatchVariables": true,

    // Modules
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "resolveJsonModule": true,
    "esModuleInterop": true,
    "isolatedModules": true,

    // Emit
    "target": "ES2022",
    "outDir": "./dist",
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true,

    // Path mapping (optional)
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    },

    // Interop
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}
```

### 13.3 ESLint Configuration

```javascript
// eslint.config.mjs
import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';
import prettierConfig from 'eslint-config-prettier';

export default tseslint.config(
  // Base configurations
  eslint.configs.recommended,
  ...tseslint.configs.strictTypeChecked,
  ...tseslint.configs.stylisticTypeChecked,

  // TypeScript parser options
  {
    languageOptions: {
      parserOptions: {
        projectService: true,
        tsconfigRootDir: import.meta.dirname,
      },
    },
  },

  // Custom rules
  {
    rules: {
      // Type safety
      '@typescript-eslint/no-explicit-any': 'error',
      '@typescript-eslint/no-unsafe-assignment': 'error',
      '@typescript-eslint/no-unsafe-call': 'error',
      '@typescript-eslint/no-unsafe-member-access': 'error',
      '@typescript-eslint/no-unsafe-return': 'error',
      '@typescript-eslint/strict-boolean-expressions': 'error',
      '@typescript-eslint/no-non-null-assertion': 'warn',

      // Async
      '@typescript-eslint/no-floating-promises': 'error',
      '@typescript-eslint/no-misused-promises': 'error',
      '@typescript-eslint/await-thenable': 'error',
      '@typescript-eslint/require-await': 'error',

      // Consistency
      '@typescript-eslint/consistent-type-imports': [
        'error',
        { prefer: 'type-imports', fixStyle: 'separate-type-imports' }
      ],
      '@typescript-eslint/consistent-type-exports': 'error',
      '@typescript-eslint/explicit-function-return-type': [
        'warn',
        { allowExpressions: true }
      ],
      '@typescript-eslint/naming-convention': [
        'error',
        {
          selector: 'interface',
          format: ['PascalCase'],
        },
        {
          selector: 'typeAlias',
          format: ['PascalCase'],
        },
        {
          selector: 'enum',
          format: ['PascalCase'],
        },
        {
          selector: 'enumMember',
          format: ['PascalCase', 'UPPER_CASE'],
        },
      ],

      // Code quality
      '@typescript-eslint/no-unused-vars': [
        'error',
        { argsIgnorePattern: '^_', varsIgnorePattern: '^_' }
      ],
      '@typescript-eslint/prefer-nullish-coalescing': 'error',
      '@typescript-eslint/prefer-optional-chain': 'error',

      // Disabled rules that conflict with Prettier
      '@typescript-eslint/indent': 'off',
      '@typescript-eslint/quotes': 'off',
    },
  },

  // Test files - relaxed rules
  {
    files: ['**/*.test.ts', '**/*.spec.ts'],
    rules: {
      '@typescript-eslint/no-explicit-any': 'off',
      '@typescript-eslint/no-unsafe-assignment': 'off',
      '@typescript-eslint/no-non-null-assertion': 'off',
    },
  },

  // Disable formatting rules (handled by Prettier)
  prettierConfig,
);
```

### 13.4 Prettier Configuration

```json
// .prettierrc
{
  "printWidth": 100,
  "tabWidth": 2,
  "useTabs": false,
  "semi": true,
  "singleQuote": true,
  "quoteProps": "as-needed",
  "trailingComma": "es5",
  "bracketSpacing": true,
  "arrowParens": "avoid",
  "endOfLine": "lf"
}
```

### 13.5 Pre-commit Hooks

```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: typecheck
        name: TypeScript Type Check
        entry: npx tsc --noEmit
        language: system
        types: [typescript]
        pass_filenames: false

      - id: lint
        name: ESLint
        entry: npx eslint --fix
        language: system
        types: [typescript]

      - id: format
        name: Prettier
        entry: npx prettier --write
        language: system
        types: [typescript]
```

---

## 14. Code Review Checklist

Quick reference for code reviewers:

### Type Safety
- [ ] No `any` types (or documented exception)
- [ ] No unsafe type assertions (`as`)
- [ ] Unknown external data is validated at boundaries
- [ ] Null/undefined handled with narrowing, not assertions
- [ ] Error handling covers both success and failure cases

### Types and Interfaces
- [ ] Interfaces used for object shapes
- [ ] Types used for unions, intersections, utilities
- [ ] Generic constraints are appropriate
- [ ] Public functions have explicit return types

### Async Code
- [ ] All promises are awaited or handled
- [ ] Independent operations use Promise.all
- [ ] Error handling in async functions is complete

### Module Organization
- [ ] Imports are properly ordered
- [ ] Type-only imports use `type` keyword
- [ ] File extensions are included (.js for ESM)
- [ ] No circular dependencies

### Testing
- [ ] Tests cover new/changed functionality
- [ ] Mocks and fixtures are properly typed
- [ ] Error cases are tested
- [ ] Edge cases are considered

### Documentation
- [ ] Public APIs have JSDoc comments
- [ ] Complex logic is explained
- [ ] No duplicate type information in comments

### Security
- [ ] External input is validated
- [ ] No hardcoded secrets
- [ ] Sensitive data is not logged

### Performance
- [ ] No overly complex type expressions
- [ ] Appropriate use of const assertions
- [ ] Efficient async patterns

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| TS-TYP-001 | Enable Strict Mode | :red_circle: |
| TS-TYP-002 | No Implicit Any | :red_circle: |
| TS-TYP-003 | Prefer Unknown Over Any | :red_circle: |
| TS-TYP-004 | Use Interface for Object Shapes | :yellow_circle: |
| TS-TYP-005 | Use Explicit Return Types for Public Functions | :yellow_circle: |
| TS-TYP-006 | Use Meaningful Generic Names | :green_circle: |
| TS-TYP-007 | Constrain Generic Types Appropriately | :yellow_circle: |
| TS-TYP-008 | Prefer Built-in Utility Types | :green_circle: |
| TS-NUL-001 | Enable strictNullChecks | :red_circle: |
| TS-NUL-002 | Use Optional Chaining for Nullable Access | :yellow_circle: |
| TS-NUL-003 | Use Nullish Coalescing for Defaults | :yellow_circle: |
| TS-NUL-004 | Avoid Non-Null Assertions | :yellow_circle: |
| TS-MOD-001 | Barrel File Usage | :yellow_circle: |
| TS-MOD-002 | Order Imports Consistently | :yellow_circle: |
| TS-MOD-003 | Use Type-Only Imports | :yellow_circle: |
| TS-MOD-004 | Use .js Extensions in Imports | :yellow_circle: |
| TS-MOD-005 | Prefer Named Exports | :green_circle: |
| TS-ASY-001 | Always Handle Promise Rejections | :red_circle: |
| TS-ASY-002 | Prefer async/await Over Promise Chains | :green_circle: |
| TS-ASY-003 | Use Promise.all for Independent Operations | :yellow_circle: |
| TS-ASY-004 | Type Async Functions Correctly | :yellow_circle: |
| TS-ERR-001 | Create Typed Error Classes | :yellow_circle: |
| TS-ERR-002 | Consider Result Pattern for Expected Failures | :green_circle: |
| TS-ERR-003 | Type Unknown Catch Parameters | :red_circle: |
| TS-TST-001 | Type Test Fixtures and Mocks | :yellow_circle: |
| TS-TST-002 | Test Type Inference with Explicit Assertions | :green_circle: |
| TS-DOC-001 | Document Public API Without Type Duplication | :green_circle: |
| TS-SEC-001 | Validate External Input at Boundaries | :red_circle: |
| TS-SEC-002 | Avoid Unsafe Type Assertions | :red_circle: |
| TS-PRF-001 | Avoid Overly Complex Type Expressions | :green_circle: |
| TS-PRF-002 | Use const Assertions for Immutable Data | :green_circle: |

---

## Appendix B: Glossary

| Term | Definition |
|------|------------|
| **Type narrowing** | Using type guards to refine a type to a more specific type |
| **Type guard** | A function that returns a type predicate (`value is T`) |
| **Discriminated union** | A union type where each member has a common property with literal types |
| **Utility types** | Built-in generic types like `Partial`, `Pick`, `Omit` |
| **Branded type** | A primitive type made unique via intersection with a symbol property |
| **Result pattern** | A pattern returning `{ ok: true, value } | { ok: false, error }` instead of throwing |
| **Barrel file** | An `index.ts` file that re-exports from other modules |

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- [Google TypeScript Style Guide](https://google.github.io/styleguide/tsguide.html)
- [typescript-eslint Rules](https://typescript-eslint.io/rules/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/)
- [TypeScript TSConfig Reference](https://www.typescriptlang.org/tsconfig/)
- [Types vs Interfaces in TypeScript](https://blog.logrocket.com/types-vs-interfaces-typescript/)
- [Error Handling with Result Types](https://typescript.tv/best-practices/error-handling-with-result-types/)
- [Strict Null Checks Best Practice](https://www.tsmean.com/articles/learn-typescript/strict-null-checks-best-practice/)
