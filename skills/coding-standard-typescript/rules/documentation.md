# Documentation Rules (TS-DOC-*)

Documentation in TypeScript should complement the type system, not duplicate it. JSDoc comments explain intent and behavior while types express structure.

## Documentation Strategy

- Document public API behavior, not types
- Use JSDoc for exported declarations
- Provide examples for complex functions
- Avoid duplicating type information
- Keep comments in sync with code

---

## TS-DOC-001: Document Public API Without Type Duplication :green_circle:

**Tier**: Recommended

**Rationale**: JSDoc comments complement TypeScript types by explaining intent and behavior. Do not duplicate type information already expressed in the code.

```typescript
// Correct - Describes behavior without duplicating types
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

/**
 * Calculates the total price including tax and any applicable discounts.
 * Discounts are applied before tax calculation.
 *
 * @throws {InvalidCartError} If the cart is empty or contains invalid items
 */
export function calculateTotal(cart: Cart): number {
  // ...
}

// Incorrect - Duplicates type information
/**
 * @param {string} id - The user id (string)
 * @returns {Promise<User | null>} - A promise that resolves to User or null
 */
export async function findById(id: string): Promise<User | null> {
  // ...
}

// Incorrect - States the obvious
/**
 * Gets the user name.
 * @returns The name of the user.
 */
export function getName(): string {
  return this.name;
}
```

---

## Documentation Requirements

| Element | Documentation Required |
|---------|----------------------|
| Public functions | Yes - describe purpose and behavior |
| Public types/interfaces | Yes - describe purpose |
| Complex logic | Yes - explain why, not what |
| Private internals | Optional - only if complex |
| Re-exports | No - documented at source |

---

## Additional Patterns

### Interface Documentation

```typescript
/**
 * Configuration options for the HTTP client.
 * All timeout values are in milliseconds.
 */
export interface HttpClientConfig {
  /** Base URL for all requests. Must include protocol. */
  baseUrl: string;

  /** Request timeout. Defaults to 30000ms. */
  timeout?: number;

  /** Number of retry attempts for failed requests. Defaults to 3. */
  retries?: number;

  /** Custom headers to include in all requests. */
  headers?: Record<string, string>;
}
```

### Type Alias Documentation

```typescript
/**
 * Represents the possible states of an async operation.
 * Use discriminated union pattern to handle each state.
 */
export type AsyncState<T, E = Error> =
  | { status: 'idle' }
  | { status: 'loading' }
  | { status: 'success'; data: T }
  | { status: 'error'; error: E };
```

### Function Overload Documentation

```typescript
/**
 * Formats a date for display.
 */
export function formatDate(date: Date): string;
/**
 * Formats a date with a custom format string.
 * @param format - Format tokens: YYYY, MM, DD, HH, mm, ss
 */
export function formatDate(date: Date, format: string): string;
export function formatDate(date: Date, format?: string): string {
  // Implementation
}
```

### Class Documentation

```typescript
/**
 * Service for managing user accounts and authentication.
 *
 * @example
 * const userService = new UserService(repository, authProvider);
 * const user = await userService.authenticate(credentials);
 */
export class UserService {
  /**
   * Creates a new UserService instance.
   *
   * @param repository - Data access layer for user storage
   * @param authProvider - Authentication provider for credential verification
   */
  constructor(
    private readonly repository: UserRepository,
    private readonly authProvider: AuthProvider
  ) {}

  /**
   * Authenticates a user with the provided credentials.
   * On success, returns the authenticated user with a session token.
   *
   * @throws {InvalidCredentialsError} If credentials are invalid
   * @throws {AccountLockedError} If the account is locked due to failed attempts
   */
  async authenticate(credentials: Credentials): Promise<AuthenticatedUser> {
    // ...
  }
}
```

### Module-Level Documentation

```typescript
/**
 * @module validation
 *
 * Provides runtime validation utilities for user input.
 * All validators return a Result type for type-safe error handling.
 *
 * @example
 * import { validateEmail, validatePassword } from './validation.js';
 *
 * const emailResult = validateEmail(input.email);
 * if (!emailResult.ok) {
 *   showError(emailResult.error);
 * }
 */

export function validateEmail(email: string): Result<string, ValidationError> {
  // ...
}
```

### Deprecation Documentation

```typescript
/**
 * @deprecated Use `findById` instead. Will be removed in v3.0.
 */
export function getUser(id: string): Promise<User | null> {
  return findById(id);
}

/**
 * @deprecated Use {@link UserService.authenticate} instead.
 * This function doesn't support MFA.
 */
export function login(email: string, password: string): Promise<User> {
  // ...
}
```

### Link References

```typescript
/**
 * Processes orders according to the fulfillment workflow.
 * See {@link OrderStatus} for possible status transitions.
 * Related: {@link ShippingService.createShipment}
 */
export function processOrder(order: Order): Promise<ProcessedOrder> {
  // ...
}
```

### Inline Comments for Complex Logic

```typescript
function calculateDiscount(cart: Cart, user: User): number {
  // Apply loyalty discount first (stacks with promotions)
  let discount = user.loyaltyTier * 0.05;

  // Promotional discounts are capped at 30%
  const promoDiscount = Math.min(cart.promoDiscount, 0.3);
  discount += promoDiscount;

  // Total discount cannot exceed 50% per business rules
  return Math.min(discount, 0.5);
}
```
