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

  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.1
    hooks:
      - id: go-fmt
      - id: go-imports
        args: ['-local', 'myapp']  # Replace with your module name
      - id: go-vet
      - id: go-mod-tidy
      - id: golangci-lint
        args: ['--timeout=5m']

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
| end-of-file-fixer | Ensure newline at EOF | GO-FMT-* |
| check-yaml | Validate YAML syntax | Config files |
| check-added-large-files | Prevent large file commits | Repository health |
| go-fmt | Format Go code | Formatting |
| go-imports | Organize imports | GO-FMT-001 |
| go-vet | Static analysis | Multiple |
| go-mod-tidy | Clean go.mod/go.sum | Module hygiene |
| golangci-lint | Comprehensive linting | All rules |
| commitizen | Conventional commits | Git history |

## Bypassing Hooks (Emergency Only)

```bash
# Skip all hooks (use sparingly!)
git commit --no-verify -m "emergency fix"

# Skip specific hook
SKIP=golangci-lint git commit -m "wip: incomplete feature"
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
# Run only changed files
pre-commit run --files $(git diff --name-only --cached)

# Specify timeout in config
# Add to .pre-commit-config.yaml under the hook:
#   timeout: 120
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
```
