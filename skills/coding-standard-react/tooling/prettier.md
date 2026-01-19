# Prettier Configuration

Prettier is an opinionated code formatter that ensures consistent code style across the project.

## Installation

```bash
# Install Prettier
npm install --save-dev prettier

# Install ESLint integration (optional but recommended)
npm install --save-dev eslint-config-prettier eslint-plugin-prettier
```

## Required Version

- **Prettier**: ^3.0.0

---

## Configuration

Create `.prettierrc` at the project root:

```json
{
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "es5",
  "printWidth": 100,
  "jsxSingleQuote": false,
  "bracketSpacing": true,
  "bracketSameLine": false,
  "arrowParens": "avoid",
  "endOfLine": "lf"
}
```

### Alternative: JavaScript Config

Create `prettier.config.js`:

```javascript
/** @type {import("prettier").Config} */
const config = {
  semi: true,
  singleQuote: true,
  tabWidth: 2,
  trailingComma: 'es5',
  printWidth: 100,
  jsxSingleQuote: false,
  bracketSpacing: true,
  bracketSameLine: false,
  arrowParens: 'avoid',
  endOfLine: 'lf',
};

export default config;
```

---

## Configuration Options Explained

| Option | Value | Rationale |
|--------|-------|-----------|
| `semi` | `true` | Explicit statement endings |
| `singleQuote` | `true` | Consistent with JS conventions |
| `tabWidth` | `2` | Standard React/JS indentation |
| `trailingComma` | `es5` | Cleaner diffs, valid ES5 |
| `printWidth` | `100` | Reasonable line length |
| `jsxSingleQuote` | `false` | HTML convention for attributes |
| `bracketSpacing` | `true` | Readable object literals |
| `bracketSameLine` | `false` | Better readability for JSX |
| `arrowParens` | `avoid` | Cleaner single-param arrows |

---

## Ignore Files

Create `.prettierignore`:

```
# Build outputs
dist/
build/
.next/
out/

# Dependencies
node_modules/

# Generated files
coverage/
*.min.js
*.bundle.js

# Lock files
package-lock.json
yarn.lock
pnpm-lock.yaml

# Config files that shouldn't be formatted
.env*
```

---

## ESLint Integration

To use Prettier with ESLint, add to your ESLint config:

### ESLint 9+ Flat Config

```javascript
// eslint.config.js
import prettier from 'eslint-config-prettier';

export default [
  // ... other configs
  prettier, // Must be last to override other formatting rules
];
```

### Legacy ESLint 8

```json
{
  "extends": [
    "eslint:recommended",
    "plugin:react/recommended",
    "prettier"
  ]
}
```

---

## Running Prettier

```bash
# Check if files are formatted
npx prettier --check .

# Format all files
npx prettier --write .

# Format specific file types
npx prettier --write "**/*.{ts,tsx,js,jsx,json,css,md}"

# Check single file
npx prettier --check src/App.tsx
```

---

## Editor Integration

### VS Code

Install the "Prettier - Code formatter" extension and add to `.vscode/settings.json`:

```json
{
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.formatOnSave": true,
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  }
}
```

### WebStorm/IntelliJ

1. Go to Settings > Languages & Frameworks > JavaScript > Prettier
2. Enable "On save" and "On 'Reformat Code' action"
3. Set Prettier package path

---

## CI Integration

### GitHub Actions

```yaml
- name: Check formatting
  run: npx prettier --check .
```

### Pre-commit Hook

See [Pre-commit](pre-commit.md) for full setup.

```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: prettier
        name: Prettier
        entry: npx prettier --write
        language: system
        types: [javascript, jsx, typescript, tsx, json, css, scss]
        pass_filenames: true
```

---

## Package.json Scripts

```json
{
  "scripts": {
    "format": "prettier --write .",
    "format:check": "prettier --check .",
    "lint": "eslint . && prettier --check ."
  }
}
```
