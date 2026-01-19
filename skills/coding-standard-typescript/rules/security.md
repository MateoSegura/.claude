# Security Rules (TS-SEC-*)

TypeScript types exist only at compile time. Runtime security requires explicit validation and careful handling of external data.

## Security Strategy

- Validate external input at boundaries
- Avoid unsafe type assertions
- Use runtime validation libraries (Zod, io-ts)
- Never trust user input
- Sanitize data before output

---

## TS-SEC-001: Validate External Input at Boundaries :red_circle:

**Tier**: Critical

**Rationale**: TypeScript types only exist at compile time. Runtime data from APIs, user input, and file I/O must be validated. Use runtime validation libraries like Zod for type-safe validation.

```typescript
// Correct - Runtime validation with Zod
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

// Correct - Safe parsing with result
async function createUserSafe(input: unknown): Promise<Result<User, ValidationError[]>> {
  const result = UserInputSchema.safeParse(input);
  if (!result.success) {
    return err(result.error.errors.map(e => ({
      field: e.path.join('.'),
      message: e.message,
    })));
  }
  return ok(await userRepository.create(result.data));
}

// Incorrect - Trusting external input types
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

---

## TS-SEC-002: Avoid Unsafe Type Assertions :red_circle:

**Tier**: Critical

**Rationale**: Type assertions (`as`) override TypeScript's type checking. Using them incorrectly can introduce type-safety bugs that won't be caught until runtime.

```typescript
// Correct - Type guard with validation
function isUser(value: unknown): value is User {
  return (
    typeof value === 'object' &&
    value !== null &&
    'id' in value &&
    typeof (value as { id: unknown }).id === 'string' &&
    'name' in value &&
    typeof (value as { name: unknown }).name === 'string' &&
    'email' in value &&
    typeof (value as { email: unknown }).email === 'string'
  );
}

function processApiResponse(data: unknown): User {
  if (isUser(data)) {
    return data;  // Type-safe narrowing
  }
  throw new Error('Invalid user data');
}

// Correct - Using Zod for validation
const UserSchema = z.object({
  id: z.string(),
  name: z.string(),
  email: z.string().email(),
});

function processApiResponse(data: unknown): User {
  return UserSchema.parse(data);
}

// Incorrect - Unsafe type assertion
function processApiResponse(data: unknown): User {
  return data as User;  // No validation, crashes if wrong
}

// Incorrect - Double assertion to bypass checks
function hackTheType(value: string): number {
  return value as unknown as number;  // Nonsensical, will fail
}

// Incorrect - Assertion on nullable values
function getLength(value: string | null): number {
  return (value as string).length;  // Crashes if null
}
```

---

## Additional Security Patterns

### Sanitizing User Input

```typescript
// HTML escaping for XSS prevention
function escapeHtml(text: string): string {
  const map: Record<string, string> = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#039;',
  };
  return text.replace(/[&<>"']/g, char => map[char]);
}

// SQL parameterization (using query builders)
async function findUserByEmail(email: string): Promise<User | null> {
  // Correct - Parameterized query
  const result = await db.query(
    'SELECT * FROM users WHERE email = $1',
    [email]
  );
  return result.rows[0] ?? null;
}

// Incorrect - String interpolation (SQL injection risk)
async function findUserByEmail(email: string): Promise<User | null> {
  const result = await db.query(
    `SELECT * FROM users WHERE email = '${email}'`  // NEVER do this
  );
  return result.rows[0] ?? null;
}
```

### Path Validation

```typescript
import path from 'node:path';

function validateFilePath(userPath: string, baseDir: string): string {
  // Normalize and resolve the path
  const resolved = path.resolve(baseDir, userPath);

  // Ensure the resolved path is within the base directory
  if (!resolved.startsWith(path.resolve(baseDir) + path.sep)) {
    throw new Error('Path traversal detected');
  }

  return resolved;
}

// Usage
const safePath = validateFilePath(req.params.filename, '/uploads');
const content = await fs.readFile(safePath);
```

### Environment Variable Validation

```typescript
const EnvSchema = z.object({
  NODE_ENV: z.enum(['development', 'production', 'test']),
  DATABASE_URL: z.string().url(),
  API_KEY: z.string().min(32),
  PORT: z.string().transform(Number).pipe(z.number().int().positive()),
});

type Env = z.infer<typeof EnvSchema>;

function validateEnv(): Env {
  const result = EnvSchema.safeParse(process.env);
  if (!result.success) {
    console.error('Invalid environment variables:', result.error.format());
    process.exit(1);
  }
  return result.data;
}

export const env = validateEnv();
```

### Secrets Management

```typescript
// Correct - Environment variables for secrets
const config = {
  apiKey: process.env.API_KEY,
  dbPassword: process.env.DB_PASSWORD,
};

if (!config.apiKey || !config.dbPassword) {
  throw new Error('Missing required environment variables');
}

// Incorrect - Hardcoded secrets
const config = {
  apiKey: 'sk-1234567890abcdef',  // NEVER do this
  dbPassword: 'supersecret123',   // NEVER do this
};
```

### JWT Validation

```typescript
import { z } from 'zod';
import jwt from 'jsonwebtoken';

const JwtPayloadSchema = z.object({
  sub: z.string(),
  email: z.string().email(),
  role: z.enum(['user', 'admin']),
  exp: z.number(),
  iat: z.number(),
});

type JwtPayload = z.infer<typeof JwtPayloadSchema>;

function verifyToken(token: string): JwtPayload {
  // Verify signature first
  const decoded = jwt.verify(token, process.env.JWT_SECRET!);

  // Then validate payload structure
  return JwtPayloadSchema.parse(decoded);
}
```

### Rate Limiting Types

```typescript
interface RateLimitConfig {
  windowMs: number;      // Time window in milliseconds
  maxRequests: number;   // Max requests per window
  keyGenerator: (req: Request) => string;
}

const rateLimitSchema = z.object({
  windowMs: z.number().int().positive().max(3600000),
  maxRequests: z.number().int().positive().max(10000),
});

function createRateLimiter(config: RateLimitConfig): Middleware {
  // Validate config at startup
  rateLimitSchema.parse({
    windowMs: config.windowMs,
    maxRequests: config.maxRequests,
  });

  // Return rate limiting middleware
  return async (req, res, next) => {
    const key = config.keyGenerator(req);
    // ... rate limiting logic
  };
}
```
