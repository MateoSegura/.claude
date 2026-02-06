# Pre-commit Configuration

Save as `.pre-commit-config.yaml` in project root.

## Configuration

```yaml
repos:
  - repo: local
    hooks:
      - id: clang-format
        name: clang-format
        entry: clang-format
        args: [-i, --style=file, --Werror]
        language: system
        types: [c]

      - id: clang-tidy
        name: clang-tidy
        entry: clang-tidy
        args: [--config-file=.clang-tidy]
        language: system
        types: [c]

      - id: cppcheck
        name: cppcheck
        entry: cppcheck
        args: [--error-exitcode=1, --enable=warning,style]
        language: system
        types: [c]

      - id: trailing-whitespace
        name: Trim Trailing Whitespace
        entry: trailing-whitespace-fixer
        language: system
        types: [c]

      - id: check-added-large-files
        name: Check for large files
        entry: check-added-large-files
        args: [--maxkb=100]
        language: system
```

## Installation

```bash
# Install pre-commit
pip install pre-commit

# Install hooks
pre-commit install

# Run manually on all files
pre-commit run --all-files
```

## Required Tools

| Tool | Version | Installation |
|------|---------|--------------|
| clang-format | 14.0+ | `apt install clang-format` |
| clang-tidy | 14.0+ | `apt install clang-tidy` |
| cppcheck | 2.7+ | `apt install cppcheck` |
| pre-commit | latest | `pip install pre-commit` |
