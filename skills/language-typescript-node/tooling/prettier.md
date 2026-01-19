# Prettier Configuration

Prettier is the standard code formatter for TypeScript projects. It handles all formatting concerns, allowing ESLint to focus on code quality rules.

## Installation

```bash
# Install Prettier and ESLint integration
npm install --save-dev prettier eslint-config-prettier

# Or with pnpm
pnpm add -D prettier eslint-config-prettier
```

## Required Version

- **Prettier**: >=3.0

## Configuration File

Create `.prettierrc` at the project root:

```json
{
  "printWidth": 100,
  "tabWidth": 2,
  "useTabs": false,
  "semi": true,
  "singleQuote": true,
  "quoteProps": "as-needed",
  "trailingComma": "es5",
  "bracketSpacing": true,
  "bracketSameLine": false,
  "arrowParens": "avoid",
  "endOfLine": "lf"
}
```

## Configuration Options Explained

| Option | Value | Purpose |
|--------|-------|---------|
| `printWidth` | `100` | Maximum line length |
| `tabWidth` | `2` | Spaces per indentation level |
| `useTabs` | `false` | Use spaces instead of tabs |
| `semi` | `true` | Add semicolons at end of statements |
| `singleQuote` | `true` | Use single quotes for strings |
| `quoteProps` | `as-needed` | Only quote object keys when required |
| `trailingComma` | `es5` | Trailing commas where valid in ES5 |
| `bracketSpacing` | `true` | Spaces inside object braces |
| `bracketSameLine` | `false` | Put `>` on its own line in JSX |
| `arrowParens` | `avoid` | Omit parens for single arrow function params |
| `endOfLine` | `lf` | Use Unix line endings |

## Alternative Configuration (JavaScript)

For more control, create `prettier.config.mjs`:

```javascript
// prettier.config.mjs
/** @type {import("prettier").Config} */
export default {
  printWidth: 100,
  tabWidth: 2,
  useTabs: false,
  semi: true,
  singleQuote: true,
  quoteProps: 'as-needed',
  trailingComma: 'es5',
  bracketSpacing: true,
  bracketSameLine: false,
  arrowParens: 'avoid',
  endOfLine: 'lf',

  // Plugin-specific options
  plugins: ['prettier-plugin-organize-imports'],
};
```

## Ignore File

Create `.prettierignore` to exclude files:

```
# Build outputs
dist/
build/
coverage/

# Dependencies
node_modules/

# Lock files
package-lock.json
pnpm-lock.yaml
yarn.lock

# Generated files
*.d.ts
*.min.js

# Config files (optional)
*.config.js
*.config.mjs
```

## Running Prettier

```bash
# Check formatting
npx prettier --check .

# Format all files
npx prettier --write .

# Format specific files
npx prettier --write "src/**/*.ts"

# Check single file
npx prettier --check src/index.ts
```

## ESLint Integration

Ensure Prettier rules are disabled in ESLint to avoid conflicts:

```javascript
// eslint.config.mjs
import prettierConfig from 'eslint-config-prettier';

export default tseslint.config(
  // ... other configs
  prettierConfig,  // Must be last
);
```

## VS Code Integration

Install the Prettier extension and add to `.vscode/settings.json`:

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
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  }
}
```

## Package.json Scripts

Add formatting scripts:

```json
{
  "scripts": {
    "format": "prettier --write .",
    "format:check": "prettier --check .",
    "lint": "eslint . && prettier --check .",
    "lint:fix": "eslint --fix . && prettier --write ."
  }
}
```

## CI Integration

### GitHub Actions

```yaml
- name: Check formatting
  run: npx prettier --check .
```

### GitLab CI

```yaml
format:
  image: node:20
  script:
    - npm ci
    - npx prettier --check .
```

## Plugins

### Import Sorting

```bash
npm install --save-dev prettier-plugin-organize-imports
```

```json
{
  "plugins": ["prettier-plugin-organize-imports"]
}
```

### Tailwind CSS (if applicable)

```bash
npm install --save-dev prettier-plugin-tailwindcss
```

```json
{
  "plugins": ["prettier-plugin-tailwindcss"]
}
```

## Editor Configuration

Create `.editorconfig` for editor-agnostic settings:

```ini
# .editorconfig
root = true

[*]
indent_style = space
indent_size = 2
end_of_line = lf
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true

[*.md]
trim_trailing_whitespace = false
```

## Troubleshooting

### Conflicts with ESLint

If ESLint and Prettier conflict:

1. Ensure `eslint-config-prettier` is installed
2. Add it last in your ESLint config
3. Use `prettier --check` separately from ESLint

### Different formatting locally vs CI

Ensure:
1. Same Prettier version in `package.json`
2. Same config files are committed
3. Run `npm ci` (not `npm install`) in CI

### Slow formatting

Use `.prettierignore` to skip:
- `node_modules/`
- `dist/`
- Large generated files
