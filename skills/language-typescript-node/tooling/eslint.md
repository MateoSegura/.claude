# ESLint Configuration

ESLint with typescript-eslint is the standard linting solution for TypeScript projects. This configuration enforces the rules defined in this coding standard.

## Installation

```bash
# Install ESLint and typescript-eslint
npm install --save-dev eslint @eslint/js typescript-eslint eslint-config-prettier

# Or with pnpm
pnpm add -D eslint @eslint/js typescript-eslint eslint-config-prettier
```

## Required Versions

- **ESLint**: >=9.0 (flat config)
- **typescript-eslint**: >=8.0
- **TypeScript**: >=5.0

## Configuration File

Create `eslint.config.mjs` at the project root:

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
      // Type safety (TS-TYP-*, TS-SEC-*)
      '@typescript-eslint/no-explicit-any': 'error',
      '@typescript-eslint/no-unsafe-assignment': 'error',
      '@typescript-eslint/no-unsafe-call': 'error',
      '@typescript-eslint/no-unsafe-member-access': 'error',
      '@typescript-eslint/no-unsafe-return': 'error',
      '@typescript-eslint/strict-boolean-expressions': 'error',
      '@typescript-eslint/no-non-null-assertion': 'warn',

      // Async (TS-ASY-*)
      '@typescript-eslint/no-floating-promises': 'error',
      '@typescript-eslint/no-misused-promises': 'error',
      '@typescript-eslint/await-thenable': 'error',
      '@typescript-eslint/require-await': 'error',

      // Consistency (TS-MOD-*, TS-TYP-*)
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
    files: ['**/*.test.ts', '**/*.spec.ts', '**/tests/**/*.ts'],
    rules: {
      '@typescript-eslint/no-explicit-any': 'off',
      '@typescript-eslint/no-unsafe-assignment': 'off',
      '@typescript-eslint/no-non-null-assertion': 'off',
      '@typescript-eslint/no-unsafe-member-access': 'off',
    },
  },

  // Disable formatting rules (handled by Prettier)
  prettierConfig,
);
```

## Rule Descriptions

| ESLint Rule | Purpose | Standard Rule |
|-------------|---------|---------------|
| `no-explicit-any` | Disallow any types | TS-TYP-003 |
| `no-unsafe-*` | Prevent unsafe operations on any | TS-SEC-002 |
| `strict-boolean-expressions` | Require strict boolean conditions | Type safety |
| `no-non-null-assertion` | Warn on ! operator | TS-NUL-004 |
| `no-floating-promises` | Require promise handling | TS-ASY-001 |
| `no-misused-promises` | Prevent promise misuse | TS-ASY-001 |
| `consistent-type-imports` | Enforce type imports | TS-MOD-003 |
| `explicit-function-return-type` | Require return types | TS-TYP-005 |
| `prefer-nullish-coalescing` | Prefer ?? over \|\| | TS-NUL-003 |
| `prefer-optional-chain` | Prefer ?. over && | TS-NUL-002 |

## Running the Linter

```bash
# Run ESLint on all TypeScript files
npx eslint .

# Run on specific directories
npx eslint src/ tests/

# Fix auto-fixable issues
npx eslint --fix .

# Output in different formats
npx eslint --format stylish .
npx eslint --format json . > eslint-report.json
```

## CI Integration

### GitHub Actions

```yaml
- name: ESLint
  run: npx eslint . --format @microsoft/eslint-formatter-sarif --output-file eslint-results.sarif
  continue-on-error: true

- name: Upload ESLint results
  uses: github/codeql-action/upload-sarif@v2
  with:
    sarif_file: eslint-results.sarif
```

### GitLab CI

```yaml
lint:
  image: node:20
  script:
    - npm ci
    - npx eslint . --format json > eslint-report.json
  artifacts:
    reports:
      codequality: eslint-report.json
```

## VS Code Integration

Add to `.vscode/settings.json`:

```json
{
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": "explicit"
  },
  "eslint.validate": ["typescript", "typescriptreact"]
}
```

## Excluding Files

Add to `eslint.config.mjs`:

```javascript
export default tseslint.config(
  // Ignore patterns
  {
    ignores: [
      'dist/**',
      'node_modules/**',
      '*.config.js',
      '*.config.mjs',
      'coverage/**',
    ],
  },
  // ... rest of config
);
```

## Troubleshooting

### Type-aware linting is slow

```javascript
// Use projectService for better performance
{
  languageOptions: {
    parserOptions: {
      projectService: true,  // Faster than project array
    },
  },
}
```

### Cannot find tsconfig

```javascript
{
  languageOptions: {
    parserOptions: {
      projectService: true,
      tsconfigRootDir: import.meta.dirname,
    },
  },
}
```

### Rules conflict with Prettier

Always include `eslint-config-prettier` last to disable conflicting rules.
