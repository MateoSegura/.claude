# ESLint Configuration

ESLint is the standard linter for React projects. This configuration enforces the rules defined in this coding standard.

## Installation

```bash
# Install ESLint and React plugins
npm install --save-dev eslint @eslint/js typescript-eslint \
  eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-jsx-a11y

# For TypeScript projects
npm install --save-dev @typescript-eslint/eslint-plugin @typescript-eslint/parser
```

## Required Versions

- **ESLint**: ^9.0.0 (flat config) or ^8.0.0 (legacy)
- **eslint-plugin-react**: ^7.33.0
- **eslint-plugin-react-hooks**: ^4.6.0
- **eslint-plugin-jsx-a11y**: ^6.7.0
- **TypeScript**: ^5.0.0

---

## ESLint 9+ Flat Config

Place this configuration in `eslint.config.js` at the project root:

```javascript
// eslint.config.js
import js from '@eslint/js';
import tseslint from 'typescript-eslint';
import react from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import jsxA11y from 'eslint-plugin-jsx-a11y';

export default tseslint.config(
  js.configs.recommended,
  ...tseslint.configs.recommended,
  {
    files: ['**/*.{ts,tsx}'],
    plugins: {
      react,
      'react-hooks': reactHooks,
      'jsx-a11y': jsxA11y,
    },
    languageOptions: {
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
      },
    },
    settings: {
      react: {
        version: 'detect',
      },
    },
    rules: {
      // React rules
      'react/jsx-uses-react': 'off', // Not needed with new JSX transform
      'react/react-in-jsx-scope': 'off', // Not needed with new JSX transform
      'react/prop-types': 'off', // Using TypeScript
      'react/jsx-no-target-blank': 'error',
      'react/jsx-key': 'error',
      'react/jsx-no-duplicate-props': 'error',
      'react/jsx-no-undef': 'error',
      'react/no-children-prop': 'error',
      'react/no-danger-with-children': 'error',
      'react/no-deprecated': 'error',
      'react/no-direct-mutation-state': 'error',
      'react/no-unescaped-entities': 'error',
      'react/no-unknown-property': 'error',
      'react/self-closing-comp': 'error',
      'react/jsx-boolean-value': ['error', 'never'],
      'react/jsx-curly-brace-presence': ['error', { props: 'never', children: 'never' }],
      'react/jsx-fragments': ['error', 'syntax'],
      'react/jsx-no-useless-fragment': 'error',
      'react/jsx-pascal-case': 'error',
      'react/no-array-index-key': 'warn',
      'react/no-unstable-nested-components': 'error',
      'react/function-component-definition': ['error', {
        namedComponents: 'function-declaration',
        unnamedComponents: 'arrow-function',
      }],

      // React Hooks rules
      'react-hooks/rules-of-hooks': 'error',
      'react-hooks/exhaustive-deps': 'warn',

      // Accessibility rules
      'jsx-a11y/alt-text': 'error',
      'jsx-a11y/anchor-has-content': 'error',
      'jsx-a11y/anchor-is-valid': 'error',
      'jsx-a11y/aria-props': 'error',
      'jsx-a11y/aria-proptypes': 'error',
      'jsx-a11y/aria-role': 'error',
      'jsx-a11y/aria-unsupported-elements': 'error',
      'jsx-a11y/click-events-have-key-events': 'error',
      'jsx-a11y/heading-has-content': 'error',
      'jsx-a11y/html-has-lang': 'error',
      'jsx-a11y/img-redundant-alt': 'error',
      'jsx-a11y/interactive-supports-focus': 'error',
      'jsx-a11y/label-has-associated-control': 'error',
      'jsx-a11y/media-has-caption': 'error',
      'jsx-a11y/mouse-events-have-key-events': 'error',
      'jsx-a11y/no-access-key': 'error',
      'jsx-a11y/no-autofocus': 'warn',
      'jsx-a11y/no-distracting-elements': 'error',
      'jsx-a11y/no-interactive-element-to-noninteractive-role': 'error',
      'jsx-a11y/no-noninteractive-element-interactions': 'error',
      'jsx-a11y/no-noninteractive-element-to-interactive-role': 'error',
      'jsx-a11y/no-noninteractive-tabindex': 'error',
      'jsx-a11y/no-redundant-roles': 'error',
      'jsx-a11y/no-static-element-interactions': 'error',
      'jsx-a11y/role-has-required-aria-props': 'error',
      'jsx-a11y/role-supports-aria-props': 'error',
      'jsx-a11y/scope': 'error',
      'jsx-a11y/tabindex-no-positive': 'error',
    },
  }
);
```

---

## Legacy ESLint 8 Configuration

For projects using ESLint 8 or below, use `.eslintrc.json`:

```json
{
  "env": {
    "browser": true,
    "es2021": true
  },
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:jsx-a11y/recommended"
  ],
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "ecmaFeatures": {
      "jsx": true
    },
    "ecmaVersion": "latest",
    "sourceType": "module"
  },
  "plugins": [
    "react",
    "react-hooks",
    "@typescript-eslint",
    "jsx-a11y"
  ],
  "settings": {
    "react": {
      "version": "detect"
    }
  },
  "rules": {
    "react/react-in-jsx-scope": "off",
    "react/prop-types": "off",
    "react-hooks/rules-of-hooks": "error",
    "react-hooks/exhaustive-deps": "warn",
    "react/jsx-pascal-case": "error",
    "react/jsx-key": "error",
    "react/no-array-index-key": "warn",
    "jsx-a11y/alt-text": "error",
    "jsx-a11y/click-events-have-key-events": "error"
  }
}
```

---

## Rule Mapping to Standard

| ESLint Rule | Standard Rule | Purpose |
|-------------|---------------|---------|
| `react-hooks/rules-of-hooks` | RCT-HKS-001 | Hooks at top level |
| `react-hooks/exhaustive-deps` | RCT-HKS-002 | Effect dependencies |
| `react/jsx-pascal-case` | RCT-STY-001 | PascalCase components |
| `react/jsx-key` | RCT-RND-003 | Stable keys |
| `react/no-array-index-key` | RCT-RND-003 | No index keys |
| `jsx-a11y/alt-text` | RCT-A11-002 | Image alt text |
| `jsx-a11y/click-events-have-key-events` | RCT-A11-003 | Keyboard accessibility |
| `jsx-a11y/aria-props` | RCT-A11-004 | ARIA correctness |

---

## Running ESLint

```bash
# Run on all files
npx eslint .

# Run on specific directories
npx eslint src/

# Fix auto-fixable issues
npx eslint . --fix

# Output in different formats
npx eslint . --format=json
npx eslint . --format=stylish

# Check specific file types
npx eslint "**/*.{ts,tsx}"
```

---

## CI Integration

### GitHub Actions

```yaml
- name: ESLint
  run: npx eslint . --format=stylish
```

### With Caching

```yaml
- name: ESLint
  uses: actions/setup-node@v4
  with:
    node-version: '20'
    cache: 'npm'
- run: npm ci
- run: npx eslint . --cache --cache-location .eslintcache
```

---

## Excluding Files

```javascript
// eslint.config.js
export default [
  {
    ignores: [
      'dist/**',
      'build/**',
      'node_modules/**',
      '*.config.js',
      'coverage/**',
    ],
  },
  // ... rest of config
];
```

For legacy config, create `.eslintignore`:

```
dist/
build/
node_modules/
coverage/
```
