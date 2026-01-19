# TypeScript Configuration

The TypeScript configuration enforces strict type checking and ensures consistent compilation settings across the project.

## Required Version

- **TypeScript**: >=5.0

## Base Configuration

Create `tsconfig.json` at the project root:

```json
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

## Configuration Options Explained

### Type Checking Options

| Option | Value | Related Rule | Purpose |
|--------|-------|--------------|---------|
| `strict` | `true` | TS-TYP-001, TS-NUL-001 | Enable all strict mode checks |
| `noUncheckedIndexedAccess` | `true` | Type safety | Add undefined to index signatures |
| `noImplicitReturns` | `true` | Type safety | Error on missing returns |
| `noFallthroughCasesInSwitch` | `true` | Type safety | Error on switch fallthrough |
| `noUnusedLocals` | `true` | Code quality | Error on unused variables |
| `noUnusedParameters` | `true` | Code quality | Error on unused parameters |
| `exactOptionalPropertyTypes` | `true` | Type safety | Strict optional properties |
| `useUnknownInCatchVariables` | `true` | TS-ERR-003 | Catch variables are unknown |

### What `strict: true` Enables

```json
{
  "strictNullChecks": true,        // TS-NUL-001
  "strictFunctionTypes": true,     // Type safety
  "strictBindCallApply": true,     // Type safety
  "strictPropertyInitialization": true,  // Type safety
  "noImplicitAny": true,           // TS-TYP-002
  "noImplicitThis": true,          // Type safety
  "alwaysStrict": true             // ES5 strict mode
}
```

## Project-Specific Configurations

### Node.js Application

```json
{
  "compilerOptions": {
    "strict": true,
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "outDir": "./dist",
    "rootDir": "./src",
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,
    "useUnknownInCatchVariables": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}
```

### Library Package

```json
{
  "compilerOptions": {
    "strict": true,
    "target": "ES2020",
    "module": "ESNext",
    "moduleResolution": "Bundler",
    "outDir": "./dist",
    "rootDir": "./src",
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist", "**/*.test.ts"]
}
```

### React Application

```json
{
  "compilerOptions": {
    "strict": true,
    "target": "ES2020",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "moduleResolution": "Bundler",
    "jsx": "react-jsx",
    "outDir": "./dist",
    "declaration": true,
    "sourceMap": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    }
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}
```

## Separate Test Configuration

Create `tsconfig.test.json`:

```json
{
  "extends": "./tsconfig.json",
  "compilerOptions": {
    "noEmit": true,
    "types": ["vitest/globals", "node"]
  },
  "include": ["src/**/*", "tests/**/*"]
}
```

## Common Patterns

### Path Aliases

```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"],
      "@/types/*": ["./src/types/*"],
      "@/utils/*": ["./src/utils/*"],
      "@/services/*": ["./src/services/*"]
    }
  }
}
```

### Monorepo Base Configuration

```json
// packages/tsconfig.base.json
{
  "compilerOptions": {
    "strict": true,
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,
    "composite": true
  }
}
```

```json
// packages/my-package/tsconfig.json
{
  "extends": "../tsconfig.base.json",
  "compilerOptions": {
    "outDir": "./dist",
    "rootDir": "./src"
  },
  "include": ["src/**/*"],
  "references": [
    { "path": "../shared" }
  ]
}
```

## Troubleshooting

### Cannot find module with path alias

Ensure path aliases are also configured in your bundler (Vite, webpack) or use `tsconfig-paths` for Node.js:

```bash
npm install --save-dev tsconfig-paths
```

```bash
node -r tsconfig-paths/register dist/index.js
```

### Index signatures return undefined

With `noUncheckedIndexedAccess`, indexed access returns `T | undefined`:

```typescript
const obj: Record<string, string> = {};
const value = obj['key'];  // string | undefined

// Handle the undefined case
if (value !== undefined) {
  console.log(value.toUpperCase());
}
```

### Optional properties require undefined assignment

With `exactOptionalPropertyTypes`:

```typescript
interface Config {
  timeout?: number;  // Can be missing, but not undefined
}

// Incorrect
const config: Config = { timeout: undefined };

// Correct
const config: Config = {};
```
