# Module Rules (TS-MOD-*)

Module organization affects code discoverability, bundle size, and build performance. These rules ensure consistent, maintainable module structure.

## Module Strategy

- Order imports consistently by source
- Use type-only imports for types
- Include file extensions for ESM compatibility
- Use barrel files judiciously
- Prefer named exports over default exports

---

## TS-MOD-001: Barrel File Usage :yellow_circle:

**Tier**: Required

**Rationale**: Barrel files (index.ts re-exports) simplify imports but can cause circular dependencies and tree-shaking issues. Use them judiciously for public APIs only.

```typescript
// Correct - Barrel file for public API
// src/types/index.ts
export type { User, UserRole } from './user.js';
export type { Order, OrderStatus } from './order.js';

// Consumer
import type { User, Order } from './types/index.js';

// Correct - Direct imports for internal modules
import { UserService } from './services/user-service.js';
import { OrderService } from './services/order-service.js';

// Incorrect - Deep barrel nesting causing circular imports
// src/index.ts
export * from './services/index.js';
export * from './models/index.js';  // models imports services
export * from './utils/index.js';   // utils imports models
```

---

## TS-MOD-002: Order Imports Consistently :yellow_circle:

**Tier**: Required

**Rationale**: Consistent import ordering improves readability and reduces merge conflicts. Group imports by external dependencies, internal modules, and types.

```typescript
// Correct - Organized imports
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

// Incorrect - Unorganized imports
import type { User } from '@/types/user.js';
import express from 'express';
import { readFile } from 'node:fs/promises';
import { UserService } from '@/services/user-service.js';
import type { Request } from 'express';
import { z } from 'zod';
```

---

## TS-MOD-003: Use Type-Only Imports :yellow_circle:

**Tier**: Required

**Rationale**: Type-only imports are erased at compile time and clearly communicate that an import is used only for type information. This can improve build performance.

```typescript
// Correct - Separate type imports
import { UserService } from './user-service.js';
import type { User, UserCreateInput } from './types.js';

// Correct - Inline type imports (TypeScript 4.5+)
import { UserService, type User, type UserCreateInput } from './user.js';

// Correct - Type-only re-exports
export type { User, UserRole } from './user.js';

// Incorrect - Mixed without type modifier
import { UserService, User, UserCreateInput } from './user.js';
// (If User and UserCreateInput are only used as types)
```

---

## TS-MOD-004: Use .js Extensions in Imports :yellow_circle:

**Tier**: Required

**Rationale**: When targeting ESM output, TypeScript requires `.js` extensions in imports even for `.ts` source files. This ensures compiled output has correct import paths.

```typescript
// Correct - .js extension for ESM compatibility
import { helper } from './helper.js';
import { User } from '../types/user.js';
import { UserService } from '@/services/user-service.js';

// Incorrect - Missing extension (fails in ESM)
import { helper } from './helper';
import { User } from '../types/user';

// Note: Some bundlers (Vite, webpack) handle this automatically
// but native ESM requires extensions
```

---

## TS-MOD-005: Prefer Named Exports :green_circle:

**Tier**: Recommended

**Rationale**: Named exports provide better IDE support (autocomplete, refactoring), prevent inconsistent naming across imports, and enable tree-shaking.

```typescript
// Correct - Named exports
// user-service.ts
export class UserService {
  // ...
}

export function createUser(data: UserCreateInput): User {
  // ...
}

export const DEFAULT_ROLE = 'viewer';

// Consumer - consistent naming
import { UserService, createUser, DEFAULT_ROLE } from './user-service.js';

// Incorrect - Default export allows inconsistent naming
// user-service.ts
export default class UserService {
  // ...
}

// Consumer - different names for same thing
import UserSvc from './user-service.js';
import UserManager from './user-service.js';
import Service from './user-service.js';
```

---

## Additional Patterns

### Path Aliases

```json
// tsconfig.json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"],
      "@/types/*": ["./src/types/*"],
      "@/utils/*": ["./src/utils/*"]
    }
  }
}
```

```typescript
// Usage with path aliases
import { UserService } from '@/services/user-service.js';
import type { User } from '@/types/user.js';
```

### Re-exporting Patterns

```typescript
// Correct - Selective re-export for public API
// src/index.ts
export { UserService } from './services/user-service.js';
export { OrderService } from './services/order-service.js';
export type { User, UserRole } from './types/user.js';
export type { Order, OrderStatus } from './types/order.js';

// Incorrect - Export everything
export * from './services/index.js';  // Exposes internal implementation
```

### Dynamic Imports

```typescript
// Correct - Dynamic import for code splitting
async function loadEditor(): Promise<Editor> {
  const { Editor } = await import('./components/editor.js');
  return new Editor();
}

// Correct - Conditional dynamic import
async function loadAnalytics(): Promise<void> {
  if (process.env.NODE_ENV === 'production') {
    const { initAnalytics } = await import('./analytics.js');
    await initAnalytics();
  }
}
```
