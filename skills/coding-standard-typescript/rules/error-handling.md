# Error Handling Rules (TS-ERR-*)

Error handling in TypeScript requires both compile-time type safety and runtime validation. These rules ensure errors are handled explicitly and provide sufficient context for debugging.

## Error Handling Strategy

- Create typed error classes for structured error handling
- Consider Result pattern for expected failures
- Always type unknown catch parameters
- Provide context in error messages
- Use error boundaries appropriately

---

## TS-ERR-001: Create Typed Error Classes :yellow_circle:

**Tier**: Required

**Rationale**: Custom error classes enable type-safe error handling with `instanceof` checks. They can carry additional context and enable exhaustive error handling.

```typescript
// Correct - Typed error hierarchy
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

// Incorrect - Plain Error with no type information
function fetchUser(id: string): Promise<User> {
  throw new Error('User not found');  // No context, hard to handle
}
```

---

## TS-ERR-002: Consider Result Pattern for Expected Failures :green_circle:

**Tier**: Recommended

**Rationale**: The Result pattern makes error handling explicit in the type system. Unlike exceptions, Result types cannot be accidentally ignored and enable exhaustive error handling.

```typescript
// Correct - Result type implementation
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

// Incorrect - Throwing for expected failures
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

---

## TS-ERR-003: Type Unknown Catch Parameters :red_circle:

**Tier**: Critical

**Rationale**: In TypeScript 4.4+, catch clause variables are `unknown` by default. Always validate error types before accessing properties.

```typescript
// Correct - Type guard for errors
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

// Correct - Narrowing in catch block
try {
  await riskyOperation();
} catch (error) {
  if (error instanceof NetworkError) {
    await retry();
  } else if (error instanceof ValidationError) {
    showValidationMessage(error.field, error.message);
  } else {
    throw error;  // Re-throw unknown errors
  }
}

// Incorrect - Assuming error is Error type
async function fetchData(): Promise<Data> {
  try {
    const response = await fetch('/api/data');
    return await response.json();
  } catch (error) {
    // error is unknown, accessing .message is unsafe
    logger.error(error.message);  // TypeScript error
    throw error;
  }
}
```

---

## Additional Patterns

### Error Factory Functions

```typescript
// Create errors with consistent formatting
function createNotFoundError(resource: string, id: string): NotFoundError {
  return new NotFoundError(resource, id);
}

function createValidationError(
  field: string,
  message: string,
  value: unknown
): ValidationError {
  return new ValidationError(message, field, value);
}
```

### Error Aggregation

```typescript
// Collect multiple validation errors
class AggregateValidationError extends AppError {
  readonly code = 'VALIDATION_ERRORS';
  readonly statusCode = 400;

  constructor(public readonly errors: ValidationError[]) {
    super(`${errors.length} validation errors occurred`);
  }
}

function validateUser(input: unknown): User {
  const errors: ValidationError[] = [];

  if (!input || typeof input !== 'object') {
    throw new ValidationError('Input must be an object', 'input', input);
  }

  const obj = input as Record<string, unknown>;

  if (typeof obj.name !== 'string') {
    errors.push(new ValidationError('Name is required', 'name', obj.name));
  }

  if (typeof obj.email !== 'string' || !obj.email.includes('@')) {
    errors.push(new ValidationError('Valid email required', 'email', obj.email));
  }

  if (errors.length > 0) {
    throw new AggregateValidationError(errors);
  }

  return obj as User;
}
```

### Async Error Boundaries

```typescript
// Wrapper for consistent async error handling
async function withErrorHandling<T>(
  operation: () => Promise<T>,
  context: string
): Promise<T> {
  try {
    return await operation();
  } catch (error) {
    const message = getErrorMessage(error);
    logger.error(`${context} failed`, { error: message });

    if (error instanceof AppError) {
      throw error;
    }

    throw new InternalError(`${context}: ${message}`, isError(error) ? error : undefined);
  }
}

// Usage
const user = await withErrorHandling(
  () => userService.findById(id),
  'Fetching user'
);
```

### Retrying Failed Operations

```typescript
interface RetryOptions {
  maxAttempts: number;
  delayMs: number;
  shouldRetry?: (error: unknown) => boolean;
}

async function withRetry<T>(
  operation: () => Promise<T>,
  options: RetryOptions
): Promise<T> {
  const { maxAttempts, delayMs, shouldRetry = () => true } = options;
  let lastError: unknown;

  for (let attempt = 1; attempt <= maxAttempts; attempt++) {
    try {
      return await operation();
    } catch (error) {
      lastError = error;
      if (attempt === maxAttempts || !shouldRetry(error)) {
        throw error;
      }
      await new Promise(resolve => setTimeout(resolve, delayMs * attempt));
    }
  }

  throw lastError;
}
```
