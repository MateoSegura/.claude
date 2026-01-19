# Pre-commit Hooks Configuration

Universal pre-commit hook configuration for enforcing standards before commit.

---

## Installation

```bash
# Install pre-commit
pip install pre-commit

# Or with pipx (recommended)
pipx install pre-commit

# Install hooks in repository
pre-commit install
pre-commit install --hook-type commit-msg
```

---

## Base Configuration

Create `.pre-commit-config.yaml` in repository root:

```yaml
# .pre-commit-config.yaml
repos:
  # General checks
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.5.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-json
      - id: check-toml
      - id: check-added-large-files
        args: ['--maxkb=1000']
      - id: check-merge-conflict
      - id: detect-private-key
      - id: check-case-conflict
      - id: mixed-line-ending
        args: ['--fix=lf']

  # Commit message validation
  - repo: https://github.com/compilerla/conventional-pre-commit
    rev: v3.0.0
    hooks:
      - id: conventional-pre-commit
        stages: [commit-msg]

  # Secret detection
  - repo: https://github.com/gitleaks/gitleaks
    rev: v8.18.0
    hooks:
      - id: gitleaks

  # Markdown linting
  - repo: https://github.com/igorshubovych/markdownlint-cli
    rev: v0.38.0
    hooks:
      - id: markdownlint
        args: ['--fix']
```

---

## Language-Specific Additions

### Go

```yaml
  # Go formatting and linting
  - repo: https://github.com/golangci/golangci-lint
    rev: v1.55.0
    hooks:
      - id: golangci-lint

  - repo: local
    hooks:
      - id: go-fmt
        name: go fmt
        entry: gofmt -w
        language: system
        types: [go]
```

### Python

```yaml
  # Python formatting
  - repo: https://github.com/astral-sh/ruff-pre-commit
    rev: v0.1.6
    hooks:
      - id: ruff
        args: ['--fix']
      - id: ruff-format

  # Security scanning
  - repo: https://github.com/PyCQA/bandit
    rev: 1.7.5
    hooks:
      - id: bandit
        args: ['-c', 'pyproject.toml']
```

### JavaScript/TypeScript

```yaml
  # ESLint
  - repo: https://github.com/pre-commit/mirrors-eslint
    rev: v8.56.0
    hooks:
      - id: eslint
        additional_dependencies:
          - eslint@8.56.0
          - typescript
          - '@typescript-eslint/parser'
          - '@typescript-eslint/eslint-plugin'

  # Prettier
  - repo: https://github.com/pre-commit/mirrors-prettier
    rev: v4.0.0
    hooks:
      - id: prettier
```

### C/Zephyr

```yaml
  # clang-format
  - repo: https://github.com/pre-commit/mirrors-clang-format
    rev: v17.0.6
    hooks:
      - id: clang-format
        types_or: [c, c++]

  # cppcheck
  - repo: local
    hooks:
      - id: cppcheck
        name: cppcheck
        entry: cppcheck --error-exitcode=1 --enable=warning,style
        language: system
        types: [c, c++]
```

---

## Running Hooks

```bash
# Run on staged files (default)
pre-commit run

# Run on all files
pre-commit run --all-files

# Run specific hook
pre-commit run <hook-id>

# Skip hooks temporarily (use sparingly)
git commit --no-verify -m "message"
```

---

## CI Integration

```yaml
# GitHub Actions
- name: Run pre-commit
  uses: pre-commit/action@v3.0.0

# GitLab CI
pre-commit:
  stage: lint
  script:
    - pip install pre-commit
    - pre-commit run --all-files
```

---

## Troubleshooting

### Hook Installation Issues

```bash
# Clear cache and reinstall
pre-commit clean
pre-commit install --install-hooks
```

### Performance

```bash
# Run only on changed files in CI
pre-commit run --from-ref origin/main --to-ref HEAD
```

### Skipping Specific Files

Add to `.pre-commit-config.yaml`:

```yaml
exclude: |
  (?x)^(
    vendor/|
    generated/|
    \.min\.js$
  )
```
