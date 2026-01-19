# Pre-commit Setup

Pre-commit hooks ensure code quality before commits reach the repository.

## Installation

```bash
# Install pre-commit (Python-based)
pip install pre-commit

# Or use Homebrew on macOS
brew install pre-commit
```

## Configuration

Create `.pre-commit-config.yaml` at the project root:

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.5.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
        args: ['--maxkb=1000']
      - id: check-merge-conflict
      - id: detect-private-key

  - repo: local
    hooks:
      - id: typecheck
        name: TypeScript Type Check
        entry: npx tsc --noEmit
        language: system
        types: [typescript]
        pass_filenames: false

      - id: lint
        name: ESLint
        entry: npx eslint --fix
        language: system
        types: [typescript]

      - id: format
        name: Prettier
        entry: npx prettier --write --ignore-unknown
        language: system
        types: [typescript, json]

  - repo: https://github.com/commitizen-tools/commitizen
    rev: v3.13.0
    hooks:
      - id: commitizen
        stages: [commit-msg]
```

## Setup

```bash
# Install hooks in the repository
pre-commit install

# Install commit-msg hook for conventional commits
pre-commit install --hook-type commit-msg

# Run against all files (first time or after config changes)
pre-commit run --all-files

# Update hook versions
pre-commit autoupdate
```

## Hook Descriptions

| Hook | Purpose | Related Rules |
|------|---------|---------------|
| trailing-whitespace | Remove trailing spaces | Clean files |
| end-of-file-fixer | Ensure newline at EOF | Clean files |
| check-yaml | Validate YAML syntax | Config files |
| check-added-large-files | Prevent large file commits | Repository health |
| detect-private-key | Prevent secret commits | TS-SEC-* |
| typecheck | TypeScript compilation | TS-TYP-001 |
| lint | ESLint checking | All rules |
| format | Prettier formatting | Formatting |
| commitizen | Conventional commits | Git history |

## Alternative: Husky + lint-staged

For a Node.js-only solution:

```bash
# Install Husky and lint-staged
npm install --save-dev husky lint-staged

# Initialize Husky
npx husky init
```

Create `.husky/pre-commit`:

```bash
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx lint-staged
```

Create `lint-staged.config.mjs`:

```javascript
// lint-staged.config.mjs
export default {
  '*.ts': ['eslint --fix', 'prettier --write'],
  '*.tsx': ['eslint --fix', 'prettier --write'],
  '*.json': ['prettier --write'],
  '*.md': ['prettier --write'],
};
```

Add to `package.json`:

```json
{
  "scripts": {
    "prepare": "husky"
  }
}
```

## Commit Message Hook

Create `.husky/commit-msg` for conventional commits:

```bash
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx --no -- commitlint --edit "$1"
```

Install commitlint:

```bash
npm install --save-dev @commitlint/cli @commitlint/config-conventional
```

Create `commitlint.config.mjs`:

```javascript
// commitlint.config.mjs
export default {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [
      2,
      'always',
      [
        'feat',     // New feature
        'fix',      // Bug fix
        'docs',     // Documentation
        'style',    // Formatting
        'refactor', // Code change without feature/fix
        'perf',     // Performance improvement
        'test',     // Adding tests
        'chore',    // Maintenance
        'revert',   // Revert previous commit
        'ci',       // CI configuration
      ],
    ],
    'subject-case': [2, 'always', 'lower-case'],
    'subject-empty': [2, 'never'],
    'type-empty': [2, 'never'],
  },
};
```

## Makefile Integration

```makefile
.PHONY: setup-hooks
setup-hooks:
	pre-commit install
	pre-commit install --hook-type commit-msg

.PHONY: lint
lint:
	pre-commit run --all-files

.PHONY: lint-staged
lint-staged:
	pre-commit run
```

## Package.json Scripts

```json
{
  "scripts": {
    "prepare": "husky",
    "lint": "eslint . && prettier --check .",
    "lint:fix": "eslint --fix . && prettier --write .",
    "typecheck": "tsc --noEmit",
    "precommit": "lint-staged"
  }
}
```

## Bypassing Hooks (Emergency Only)

```bash
# Skip all hooks (use sparingly!)
git commit --no-verify -m "emergency fix"

# Skip specific pre-commit hook
SKIP=typecheck git commit -m "wip: incomplete feature"

# Skip with lint-staged
git commit --no-verify -m "wip"
```

## CI Validation

Add pre-commit checks to CI:

```yaml
# GitHub Actions
- name: Run pre-commit
  run: |
    pip install pre-commit
    pre-commit run --all-files
```

## Troubleshooting

### Hook Installation Issues

```bash
# Clean and reinstall
pre-commit clean
pre-commit install

# Check installed hooks
ls -la .git/hooks/
```

### Slow Hook Execution

```bash
# Run only on changed files
pre-commit run --files $(git diff --name-only --cached)

# Skip slow hooks for WIP commits
SKIP=typecheck git commit -m "wip"
```

### TypeScript Not Found

Ensure TypeScript is installed locally:

```bash
npm install --save-dev typescript
```

### ESLint Config Not Found

Ensure `eslint.config.mjs` exists and is valid:

```bash
npx eslint --print-config src/index.ts
```
