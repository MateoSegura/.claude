# Async Patterns Rules (TS-ASY-*)

Asynchronous code is fundamental to modern TypeScript applications. These rules ensure proper error handling, performance, and type safety in async code.

## Async Strategy

- Always handle promise rejections
- Prefer async/await over promise chains
- Use Promise.all for independent operations
- Type async functions correctly with Promise<T>
- Avoid floating promises

---

## TS-ASY-001: Always Handle Promise Rejections :red_circle:

**Tier**: Critical

**Rationale**: Unhandled promise rejections can crash Node.js applications and cause silent failures in browsers. Every promise must have error handling.

```typescript
// Correct - async/await with try-catch
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

// Correct - Promise chain with .catch()
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

// Incorrect - Unhandled promise (floating promise)
async function saveUser(user: User): Promise<void> {
  fetch('/api/users', {
    method: 'POST',
    body: JSON.stringify(user),
  });  // Promise not awaited, errors lost
}

// Incorrect - Missing catch
function loadData(): void {
  fetchData().then(data => process(data));  // No .catch()
}
```

---

## TS-ASY-002: Prefer async/await Over Promise Chains :green_circle:

**Tier**: Recommended

**Rationale**: `async/await` syntax is more readable, easier to debug (better stack traces), and less prone to errors than nested `.then()` chains.

```typescript
// Correct - async/await with sequential operations
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

// Correct - Early return with async/await
async function getUser(id: string): Promise<User | null> {
  const cached = await cache.get(id);
  if (cached) {
    return cached;
  }

  const user = await database.findUser(id);
  if (user) {
    await cache.set(id, user);
  }
  return user;
}

// Incorrect - Nested promise chains
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

---

## TS-ASY-003: Use Promise.all for Independent Operations :yellow_circle:

**Tier**: Required

**Rationale**: When multiple async operations don't depend on each other, running them concurrently with `Promise.all` improves performance significantly.

```typescript
// Correct - Concurrent independent operations
async function getDashboardData(userId: string): Promise<Dashboard> {
  const [user, orders, notifications] = await Promise.all([
    fetchUser(userId),
    fetchOrders(userId),
    fetchNotifications(userId),
  ]);

  return { user, orders, notifications };
}

// Correct - Promise.allSettled for operations that may fail independently
async function sendNotifications(userIds: string[]): Promise<NotificationResults> {
  const results = await Promise.allSettled(
    userIds.map(id => sendNotification(id))
  );

  return {
    succeeded: results.filter(r => r.status === 'fulfilled').length,
    failed: results.filter(r => r.status === 'rejected').length,
  };
}

// Correct - Promise.race for timeouts
async function fetchWithTimeout<T>(
  promise: Promise<T>,
  timeoutMs: number
): Promise<T> {
  const timeout = new Promise<never>((_, reject) => {
    setTimeout(() => reject(new Error('Timeout')), timeoutMs);
  });
  return Promise.race([promise, timeout]);
}

// Incorrect - Sequential operations that could be parallel
async function getDashboardData(userId: string): Promise<Dashboard> {
  const user = await fetchUser(userId);          // Waits
  const orders = await fetchOrders(userId);      // Then waits
  const notifications = await fetchNotifications(userId);  // Then waits

  return { user, orders, notifications };
}
```

---

## TS-ASY-004: Type Async Functions Correctly :yellow_circle:

**Tier**: Required

**Rationale**: Async functions always return a Promise. The return type annotation should use `Promise<T>` with the unwrapped value type.

```typescript
// Correct - Properly typed async function
async function fetchUser(id: string): Promise<User | null> {
  const response = await fetch(`/api/users/${id}`);
  if (response.status === 404) {
    return null;
  }
  return response.json();
}

// Correct - Async arrow function
const fetchUser = async (id: string): Promise<User | null> => {
  const response = await fetch(`/api/users/${id}`);
  return response.ok ? response.json() : null;
};

// Correct - Interface with async method
interface UserRepository {
  findById(id: string): Promise<User | null>;
  save(user: User): Promise<User>;
  delete(id: string): Promise<void>;
}

// Correct - Async method in class
class UserService {
  async getUser(id: string): Promise<User | null> {
    return this.repository.findById(id);
  }
}

// Incorrect - Missing Promise wrapper
async function fetchUser(id: string): User | null {  // Error
  // ...
}

// Incorrect - Using Promise in async function body
async function fetchUser(id: string): Promise<User> {
  return new Promise((resolve) => {  // Unnecessary wrapping
    resolve(this.repository.findById(id));
  });
}
```

---

## Additional Patterns

### Async Iteration

```typescript
// Correct - for-await-of for async iterables
async function processStream(stream: AsyncIterable<Data>): Promise<void> {
  for await (const chunk of stream) {
    await processChunk(chunk);
  }
}

// Correct - Async generator
async function* paginate<T>(
  fetchPage: (cursor: string) => Promise<Page<T>>
): AsyncGenerator<T> {
  let cursor = '';
  while (true) {
    const page = await fetchPage(cursor);
    for (const item of page.items) {
      yield item;
    }
    if (!page.nextCursor) break;
    cursor = page.nextCursor;
  }
}
```

### Avoiding Common Pitfalls

```typescript
// Incorrect - forEach with async (doesn't wait)
async function processItems(items: Item[]): Promise<void> {
  items.forEach(async item => {
    await processItem(item);  // Runs in parallel, errors not caught
  });
}

// Correct - Sequential processing
async function processItems(items: Item[]): Promise<void> {
  for (const item of items) {
    await processItem(item);
  }
}

// Correct - Parallel processing with error handling
async function processItems(items: Item[]): Promise<void> {
  await Promise.all(items.map(item => processItem(item)));
}
```

### Cleanup with finally

```typescript
async function withConnection<T>(
  fn: (conn: Connection) => Promise<T>
): Promise<T> {
  const conn = await getConnection();
  try {
    return await fn(conn);
  } finally {
    await conn.close();
  }
}
```
