# Commitlint Configuration

Commit message validation following Conventional Commits specification.

---

## Installation

### Node.js (Husky)

```bash
# Install commitlint
npm install -D @commitlint/cli @commitlint/config-conventional

# Install husky
npm install -D husky
npx husky install

# Add commit-msg hook
npx husky add .husky/commit-msg 'npx --no -- commitlint --edit "$1"'
```

### Pre-commit (Python)

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/compilerla/conventional-pre-commit
    rev: v3.0.0
    hooks:
      - id: conventional-pre-commit
        stages: [commit-msg]
```

---

## Configuration

Create `commitlint.config.js` or `.commitlintrc.yaml`:

### JavaScript Config

```javascript
// commitlint.config.js
module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    // Type must be one of these
    'type-enum': [
      2,
      'always',
      [
        'feat',     // New feature
        'fix',      // Bug fix
        'docs',     // Documentation only
        'style',    // Formatting, no code change
        'refactor', // Code change without feature/fix
        'test',     // Adding/updating tests
        'chore',    // Maintenance tasks
        'perf',     // Performance improvement
        'ci',       // CI/CD changes
        'build',    // Build system changes
        'revert',   // Revert previous commit
      ],
    ],
    // Type must be lowercase
    'type-case': [2, 'always', 'lower-case'],
    // Subject must be lowercase
    'subject-case': [2, 'always', 'lower-case'],
    // Subject max length
    'subject-max-length': [2, 'always', 50],
    // No period at end of subject
    'subject-full-stop': [2, 'never', '.'],
    // Body lines max length
    'body-max-line-length': [2, 'always', 72],
    // Header max length
    'header-max-length': [2, 'always', 72],
  },
};
```

### YAML Config

```yaml
# .commitlintrc.yaml
extends:
  - '@commitlint/config-conventional'

rules:
  type-enum:
    - 2
    - always
    - [feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert]
  type-case:
    - 2
    - always
    - lower-case
  subject-case:
    - 2
    - always
    - lower-case
  subject-max-length:
    - 2
    - always
    - 50
  body-max-line-length:
    - 2
    - always
    - 72
  header-max-length:
    - 2
    - always
    - 72
```

---

## Commit Message Format

```
<type>(<scope>): <subject>

[optional body]

[optional footer(s)]
```

### Examples

```bash
# Simple feature
feat(auth): add password reset functionality

# Bug fix with context
fix(api): handle empty response gracefully

# Breaking change
feat(api)!: change authentication response format

BREAKING CHANGE: The /auth/token endpoint now returns a JSON object
instead of a plain token string.

# With issue reference
fix(upload): resolve file corruption on large uploads

Fixes: PROJ-456
```

---

## Type Descriptions

| Type | Description | SemVer Impact |
|------|-------------|---------------|
| `feat` | New feature | MINOR |
| `fix` | Bug fix | PATCH |
| `docs` | Documentation only | None |
| `style` | Formatting, no code change | None |
| `refactor` | Code change without feature/fix | None |
| `test` | Adding/updating tests | None |
| `chore` | Maintenance tasks | None |
| `perf` | Performance improvement | PATCH |
| `ci` | CI/CD changes | None |
| `build` | Build system changes | None |
| `revert` | Revert previous commit | Depends |

---

## CI Integration

### GitHub Actions

```yaml
- name: Validate PR title
  uses: amannn/action-semantic-pull-request@v5
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
  with:
    types: |
      feat
      fix
      docs
      style
      refactor
      test
      chore
      perf
      ci
      build
      revert
```

### GitLab CI

```yaml
commitlint:
  stage: lint
  script:
    - npm install @commitlint/cli @commitlint/config-conventional
    - npx commitlint --from origin/main --to HEAD
```

---

## Shell Hook (Alternative)

For environments without Node.js:

```bash
#!/bin/bash
# .git/hooks/commit-msg

commit_msg_file=$1
commit_msg=$(cat "$commit_msg_file")

# Pattern: type(scope): subject
pattern='^(feat|fix|docs|style|refactor|test|chore|perf|ci|build|revert)(\([a-z0-9-]+\))?(!)?: .{1,50}$'

if ! echo "$commit_msg" | head -1 | grep -qE "$pattern"; then
    echo "ERROR: Invalid commit message format"
    echo ""
    echo "Expected: type(scope): subject (max 50 chars)"
    echo "Types: feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert"
    echo ""
    echo "Examples:"
    echo "  feat(auth): add login functionality"
    echo "  fix(api): handle null response"
    echo "  docs(readme): update installation steps"
    exit 1
fi
```

Make executable:

```bash
chmod +x .git/hooks/commit-msg
```
