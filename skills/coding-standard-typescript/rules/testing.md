# Testing Rules (TS-TST-*)

Type-safe testing catches errors when source types change. These rules ensure test code maintains the same quality standards as production code.

## Testing Strategy

- Type test fixtures and mocks
- Use typed test utilities
- Test type inference with explicit assertions
- Mock external dependencies with typed interfaces
- Maintain test coverage requirements

---

## TS-TST-001: Type Test Fixtures and Mocks :yellow_circle:

**Tier**: Required

**Rationale**: Typed test fixtures catch errors when source types change. Untyped mocks can drift from real implementations, causing false positives.

```typescript
// Correct - Typed test fixtures
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

function createMockOrder(overrides?: Partial<Order>): Order {
  return {
    id: 'order-456',
    userId: 'user-123',
    items: [],
    total: 0,
    status: 'pending',
    ...overrides,
  };
}

// Correct - Type-safe mock implementation
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

// Incorrect - Untyped fixtures
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

---

## TS-TST-002: Test Type Inference with Explicit Assertions :green_circle:

**Tier**: Recommended

**Rationale**: Type tests verify that functions return expected types. Use explicit type annotations or type assertions to catch type regressions.

```typescript
// Correct - Explicit type assertions in tests (Vitest)
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

// Correct - Type assertion using satisfies
const config = {
  timeout: 5000,
  retries: 3,
} satisfies Config;

// Correct - Compile-time type test
function assertType<T>(_value: T): void {}

// This line will fail to compile if getUser doesn't return User | null
assertType<User | null>(await service.getUser('123'));
```

---

## Additional Patterns

### Factory Functions for Test Data

```typescript
// Builder pattern for complex test objects
class UserBuilder {
  private user: User = {
    id: 'default-id',
    name: 'Default Name',
    email: 'default@example.com',
    createdAt: new Date(),
  };

  withId(id: string): this {
    this.user.id = id;
    return this;
  }

  withName(name: string): this {
    this.user.name = name;
    return this;
  }

  withEmail(email: string): this {
    this.user.email = email;
    return this;
  }

  build(): User {
    return { ...this.user };
  }
}

// Usage in tests
const user = new UserBuilder()
  .withId('user-123')
  .withName('Alice')
  .build();
```

### Typed Test Helpers

```typescript
// Helper with proper typing
function setupTest(): {
  service: UserService;
  repository: jest.Mocked<UserRepository>;
} {
  const repository: jest.Mocked<UserRepository> = {
    findById: jest.fn(),
    save: jest.fn(),
    delete: jest.fn(),
  };
  const service = new UserService(repository);
  return { service, repository };
}

describe('UserService', () => {
  it('should save user', async () => {
    const { service, repository } = setupTest();
    const user = createMockUser();

    repository.save.mockResolvedValue(user);
    const result = await service.createUser({ name: 'Test', email: 'test@example.com' });

    expect(result).toEqual(user);
    expect(repository.save).toHaveBeenCalledWith(expect.objectContaining({
      name: 'Test',
      email: 'test@example.com',
    }));
  });
});
```

### Testing Async Code

```typescript
// Correct - Testing async errors
describe('UserService', () => {
  it('should throw NotFoundError for unknown user', async () => {
    const { service, repository } = setupTest();
    repository.findById.mockResolvedValue(null);

    await expect(service.getUser('unknown')).rejects.toThrow(NotFoundError);
  });

  it('should handle concurrent operations', async () => {
    const { service, repository } = setupTest();
    const users = [createMockUser({ id: '1' }), createMockUser({ id: '2' })];

    repository.findById
      .mockResolvedValueOnce(users[0])
      .mockResolvedValueOnce(users[1]);

    const results = await Promise.all([
      service.getUser('1'),
      service.getUser('2'),
    ]);

    expect(results).toEqual(users);
  });
});
```

### Mocking External Dependencies

```typescript
// Define interface at consumer site for mocking
interface HttpClient {
  get<T>(url: string): Promise<T>;
  post<T>(url: string, data: unknown): Promise<T>;
}

// Production implementation
class FetchHttpClient implements HttpClient {
  async get<T>(url: string): Promise<T> {
    const response = await fetch(url);
    return response.json();
  }

  async post<T>(url: string, data: unknown): Promise<T> {
    const response = await fetch(url, {
      method: 'POST',
      body: JSON.stringify(data),
    });
    return response.json();
  }
}

// Test mock
function createMockHttpClient(): jest.Mocked<HttpClient> {
  return {
    get: jest.fn(),
    post: jest.fn(),
  };
}
```

### Coverage Requirements

| Type | Minimum | Target |
|------|---------|--------|
| Unit tests | 70% | 85% |
| Integration tests | 50% | 70% |
| Critical paths | 90% | 100% |

### Test File Organization

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
