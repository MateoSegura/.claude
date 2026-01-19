# Pre-commit Setup

Pre-commit hooks ensure shell scripts meet quality standards before commits.

## Installation

```bash
# Install pre-commit (Python)
pip install pre-commit

# Or using Homebrew (macOS)
brew install pre-commit

# Or using pipx
pipx install pre-commit
```

---

## Configuration

Create `.pre-commit-config.yaml` at the project root:

```yaml
# .pre-commit-config.yaml

repos:
  # Standard pre-commit hooks
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
      - id: check-executables-have-shebangs
      - id: check-shebang-scripts-are-executable

  # ShellCheck
  - repo: https://github.com/koalaman/shellcheck-precommit
    rev: v0.9.0
    hooks:
      - id: shellcheck
        args: ["--severity=warning"]

  # shfmt
  - repo: https://github.com/scop/pre-commit-shfmt
    rev: v3.7.0-1
    hooks:
      - id: shfmt
        args: ["-i", "2", "-bn", "-ci", "-w"]

  # Local hooks
  - repo: local
    hooks:
      # Run Bats tests on changed test files
      - id: bats
        name: Run Bats tests
        entry: bats
        language: system
        files: _test\.bats$
        types: [file]

      # Custom script validation
      - id: validate-scripts
        name: Validate shell scripts
        entry: ./scripts/validate-scripts.sh
        language: script
        files: \.sh$
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

# Run specific hook
pre-commit run shellcheck --all-files
pre-commit run shfmt --all-files
```

---

## Hook Descriptions

| Hook | Purpose | Related Rules |
|------|---------|---------------|
| trailing-whitespace | Remove trailing spaces | Clean files |
| end-of-file-fixer | Ensure newline at EOF | Clean files |
| check-executables-have-shebangs | Scripts have shebang | SH-STR-001 |
| check-shebang-scripts-are-executable | Shebang files are +x | SH-STR-001 |
| detect-private-key | Prevent secret commits | SH-SEC-004 |
| shellcheck | Static analysis | All rules |
| shfmt | Code formatting | Formatting |
| bats | Run tests | SH-TST-* |

---

## Custom Validation Script

Create `scripts/validate-scripts.sh`:

```bash
#!/usr/bin/env bash
set -euo pipefail

# Validate shell scripts
# Used as pre-commit hook

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

err() {
  echo "[ERROR] $*" >&2
}

validate_shebang() {
  local file="$1"
  local first_line
  first_line=$(head -1 "${file}")

  if [[ "${first_line}" != "#!/usr/bin/env bash" ]] && \
     [[ "${first_line}" != "#!/bin/bash" ]]; then
    err "${file}: Invalid shebang: ${first_line}"
    return 1
  fi
}

validate_strict_mode() {
  local file="$1"

  if ! grep -q "^set -.*e.*u.*o pipefail" "${file}" && \
     ! grep -q "^set -euo pipefail" "${file}"; then
    err "${file}: Missing strict mode (set -euo pipefail)"
    return 1
  fi
}

main() {
  local exit_code=0

  for file in "$@"; do
    if [[ ! -f "${file}" ]]; then
      continue
    fi

    validate_shebang "${file}" || exit_code=1
    validate_strict_mode "${file}" || exit_code=1
  done

  exit "${exit_code}"
}

main "$@"
```

---

## Alternative: Git Hooks (No pre-commit)

If you prefer native Git hooks:

### .git/hooks/pre-commit

```bash
#!/usr/bin/env bash
set -euo pipefail

# Get list of staged shell scripts
staged_files=$(git diff --cached --name-only --diff-filter=ACM | grep '\.sh$' || true)

if [[ -z "${staged_files}" ]]; then
  exit 0
fi

echo "Running ShellCheck..."
echo "${staged_files}" | xargs shellcheck --severity=warning

echo "Running shfmt..."
echo "${staged_files}" | xargs shfmt -d

echo "All checks passed!"
```

Make executable:

```bash
chmod +x .git/hooks/pre-commit
```

### Share with Team

```bash
# Create hooks directory in repo
mkdir -p .githooks

# Move hook
mv .git/hooks/pre-commit .githooks/

# Configure git to use repo hooks
git config core.hooksPath .githooks
```

---

## Makefile Integration

```makefile
.PHONY: install-hooks lint lint-fix test

# Install pre-commit hooks
install-hooks:
	pre-commit install

# Run all linters
lint:
	pre-commit run --all-files

# Run linters and fix issues
lint-fix:
	pre-commit run shfmt --all-files
	pre-commit run trailing-whitespace --all-files --hook-stage manual

# Run tests
test:
	bats tests/

# Full check (lint + test)
check: lint test
```

---

## CI Validation

Ensure hooks pass in CI:

```yaml
# GitHub Actions
- name: Run pre-commit
  run: |
    pip install pre-commit
    pre-commit run --all-files
```

---

## Bypassing Hooks (Emergency Only)

```bash
# Skip all hooks (use sparingly!)
git commit --no-verify -m "emergency fix"

# Skip specific hook
SKIP=shellcheck git commit -m "wip: work in progress"

# Skip multiple hooks
SKIP=shellcheck,shfmt git commit -m "wip"
```

---

## Troubleshooting

### Hook Not Running

```bash
# Verify hooks are installed
ls -la .git/hooks/pre-commit

# Reinstall
pre-commit install --force
```

### shfmt Changes Files

```bash
# Run shfmt to see changes
shfmt -d script.sh

# Auto-fix
shfmt -w script.sh
```

### ShellCheck Fails

```bash
# Run shellcheck directly for details
shellcheck script.sh

# Temporarily disable specific check
# shellcheck disable=SC2086
echo $unquoted_var
```

### Slow Hooks

```bash
# Run only on changed files
pre-commit run --files $(git diff --name-only)

# Skip slow hooks for WIP
SKIP=bats git commit -m "wip"
```
