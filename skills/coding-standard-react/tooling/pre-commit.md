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
      - id: check-json
      - id: check-added-large-files
        args: ['--maxkb=1000']
      - id: check-merge-conflict
      - id: detect-private-key

  - repo: local
    hooks:
      - id: eslint
        name: ESLint
        entry: npx eslint --fix --max-warnings=0
        language: system
        types: [javascript, jsx, ts, tsx]
        pass_filenames: true

      - id: prettier
        name: Prettier
        entry: npx prettier --write --ignore-unknown
        language: system
        types: [javascript, jsx, ts, tsx, json, css, scss, markdown]
        pass_filenames: true

      - id: typescript
        name: TypeScript
        entry: npx tsc --noEmit
        language: system
        types: [ts, tsx]
        pass_filenames: false

      - id: test
        name: Tests
        entry: npx jest --bail --findRelatedTests
        language: system
        types: [ts, tsx]
        pass_filenames: true
```

---

## Setup

```bash
# Install hooks in the repository
pre-commit install

# Run against all files (first time or after config changes)
pre-commit run --all-files

# Update hook versions
pre-commit autoupdate
```

---

## Alternative: lint-staged with Husky

For a Node.js-only solution, use Husky and lint-staged:

### Installation

```bash
npm install --save-dev husky lint-staged
npx husky init
```

### Configuration

Create `.lintstagedrc.json`:

```json
{
  "*.{ts,tsx}": [
    "eslint --fix --max-warnings=0",
    "prettier --write",
    "jest --bail --findRelatedTests"
  ],
  "*.{js,jsx}": [
    "eslint --fix --max-warnings=0",
    "prettier --write"
  ],
  "*.{json,css,scss,md}": [
    "prettier --write"
  ]
}
```

Create `.husky/pre-commit`:

```bash
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx lint-staged
```

---

## Hook Descriptions

| Hook | Purpose | Related Rules |
|------|---------|---------------|
| trailing-whitespace | Remove trailing spaces | Clean files |
| end-of-file-fixer | Ensure newline at EOF | Formatting |
| check-yaml | Validate YAML syntax | Config files |
| check-json | Validate JSON syntax | Config files |
| check-added-large-files | Prevent large file commits | Repository health |
| check-merge-conflict | Detect conflict markers | Git hygiene |
| detect-private-key | Prevent secret commits | Security |
| eslint | Lint JavaScript/TypeScript | All RCT-* rules |
| prettier | Format code | Formatting |
| typescript | Type check | Type safety |
| test | Run related tests | RCT-TST-* |

---

## Bypassing Hooks (Emergency Only)

```bash
# Skip all hooks (use sparingly!)
git commit --no-verify -m "emergency fix"

# Skip specific hook (pre-commit)
SKIP=eslint git commit -m "wip: incomplete feature"

# Skip specific hook (lint-staged)
SKIP_PREFLIGHT_CHECK=true git commit -m "wip"
```

---

## Troubleshooting

### Hook Installation Issues

```bash
# Clean and reinstall (pre-commit)
pre-commit clean
pre-commit install

# Clean and reinstall (husky)
rm -rf .husky
npx husky init
```

### Slow Hook Execution

```bash
# Run only changed files (pre-commit)
pre-commit run --files $(git diff --name-only --cached)

# Use --bail for faster test failure
npx jest --bail --findRelatedTests
```

### TypeScript Errors on Unrelated Files

```yaml
# In .pre-commit-config.yaml, run tsc without filenames
- id: typescript
  name: TypeScript
  entry: npx tsc --noEmit
  language: system
  types: [ts, tsx]
  pass_filenames: false  # Important: checks entire project
```

---

## Makefile Integration

```makefile
.PHONY: setup-hooks
setup-hooks:
	pre-commit install

.PHONY: lint
lint:
	pre-commit run --all-files

.PHONY: lint-staged
lint-staged:
	pre-commit run

.PHONY: clean-hooks
clean-hooks:
	pre-commit clean
	pre-commit uninstall
```

---

## CI Integration

Run the same checks in CI:

```yaml
# GitHub Actions
name: Lint

on: [push, pull_request]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
          cache: 'npm'
      - run: npm ci
      - run: npx eslint .
      - run: npx prettier --check .
      - run: npx tsc --noEmit
      - run: npm test -- --coverage
```
