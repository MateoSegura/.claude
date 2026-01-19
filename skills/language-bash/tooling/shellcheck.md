# ShellCheck Configuration

ShellCheck is the primary static analysis tool for shell scripts. It catches common bugs, security issues, and style violations.

## Installation

```bash
# Using apt (Debian/Ubuntu)
sudo apt install shellcheck

# Using Homebrew (macOS)
brew install shellcheck

# Using dnf (Fedora)
sudo dnf install ShellCheck

# Using pacman (Arch)
sudo pacman -S shellcheck

# Using cabal (Haskell)
cabal update && cabal install ShellCheck
```

## Required Version

- **ShellCheck**: 0.9.0+

---

## Configuration File

Create `.shellcheckrc` at the project root:

```ini
# .shellcheckrc
# ShellCheck configuration

# Target shell
shell=bash

# Enable all optional checks
enable=all

# Severity level (error, warning, info, style)
severity=style

# Specific rule configuration:
# SC1090: Can't follow non-constant source
# SC1091: Not following sourced file
# These are handled by our CI environment
disable=SC1090,SC1091

# Source path for sourced files
source-path=SCRIPTDIR
```

---

## Running ShellCheck

```bash
# Check single file
shellcheck script.sh

# Check multiple files
shellcheck scripts/*.sh

# Check with specific shell
shellcheck --shell=bash script.sh

# Output formats
shellcheck --format=tty script.sh      # Human readable (default)
shellcheck --format=gcc script.sh      # GCC-style
shellcheck --format=json script.sh     # JSON
shellcheck --format=checkstyle script.sh  # Checkstyle XML
shellcheck --format=diff script.sh     # Diff for auto-fix

# Set severity level
shellcheck --severity=error script.sh   # Only errors
shellcheck --severity=warning script.sh # Errors + warnings
shellcheck --severity=info script.sh    # All except style
shellcheck --severity=style script.sh   # All (default)

# Exclude specific checks
shellcheck --exclude=SC2086,SC2046 script.sh

# Include specific checks
shellcheck --include=SC2086 script.sh
```

---

## Important ShellCheck Codes

### Critical (Must Fix)

| Code | Description | Rule |
|------|-------------|------|
| SC2086 | Double quote to prevent globbing and word splitting | SH-VAR-001 |
| SC2046 | Quote this to prevent word splitting | SH-VAR-001 |
| SC2006 | Use `$()` instead of legacy backticks | SH-CMD-001 |
| SC2091 | Remove surrounding `$()` to avoid executing output | SH-SEC-001 |
| SC2059 | Don't use variables in printf format string | Security |
| SC2064 | Use single quotes for trap to avoid expansion | SH-ERR-001 |

### Required (Should Fix)

| Code | Description | Rule |
|------|-------------|------|
| SC2034 | Variable appears unused (may be scope issue) | SH-VAR-004 |
| SC2155 | Declare and assign separately to avoid masking return values | Error handling |
| SC2164 | Use `cd ... || exit` in case cd fails | SH-ERR-003 |
| SC2181 | Check exit code directly, not via `$?` | Style |
| SC2207 | Prefer mapfile or read -a for arrays | SH-VAR-003 |
| SC2230 | Use `command -v` instead of `which` | SH-ERR-004 |

### Style (Consider Fixing)

| Code | Description | Rule |
|------|-------------|------|
| SC2039 | Bashism in sh script | SH-PRT-001 |
| SC2044 | For loop over find output is fragile | SH-CTL-003 |
| SC2096 | Shebang with arguments is not portable | SH-STR-001 |
| SC2236 | Use `-n` instead of `! -z` | Style |
| SC2166 | Prefer `[[ ]]` over `[ ]` for tests | SH-CTL-001 |

---

## Inline Directives

### Disable for Line

```bash
# shellcheck disable=SC2086
echo $unquoted_variable
```

### Disable for Block

```bash
# shellcheck disable=SC2086,SC2046
process $var $(cmd)
next_line $also_unquoted
# shellcheck enable=SC2086,SC2046
```

### Disable for File

```bash
#!/usr/bin/env bash
# shellcheck disable=SC2034  # Unused variables are intentional

readonly VAR1="value1"
readonly VAR2="value2"
```

### Source Directive

```bash
# Tell ShellCheck about sourced file
# shellcheck source=./lib/utils.sh
source "${SCRIPT_DIR}/lib/utils.sh"

# Or specify source path
# shellcheck source-path=SCRIPTDIR
source "${SCRIPT_DIR}/lib/utils.sh"
```

---

## CI Integration

### GitHub Actions

```yaml
name: ShellCheck

on: [push, pull_request]

jobs:
  shellcheck:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run ShellCheck
        uses: ludeeus/action-shellcheck@master
        with:
          severity: warning
          scandir: './scripts'
          format: tty
```

### GitLab CI

```yaml
shellcheck:
  image: koalaman/shellcheck-alpine:stable
  script:
    - find . -name "*.sh" -print0 | xargs -0 shellcheck --severity=warning
  only:
    changes:
      - "**/*.sh"
```

### Pre-commit Hook

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/koalaman/shellcheck-precommit
    rev: v0.9.0
    hooks:
      - id: shellcheck
        args: ["--severity=warning"]
```

---

## VS Code Integration

Install the [ShellCheck extension](https://marketplace.visualstudio.com/items?itemName=timonwong.shellcheck).

`.vscode/settings.json`:

```json
{
  "shellcheck.enable": true,
  "shellcheck.run": "onSave",
  "shellcheck.executablePath": "shellcheck",
  "shellcheck.customArgs": ["--severity=warning"],
  "shellcheck.ignorePatterns": {
    "**/*.zsh": true,
    "**/.git/**": true
  }
}
```

---

## Troubleshooting

### "Can't follow non-constant source"

```bash
# Problem
source "${SCRIPT_DIR}/lib/${module}.sh"

# Solution 1: Disable check
# shellcheck disable=SC1090
source "${SCRIPT_DIR}/lib/${module}.sh"

# Solution 2: Provide hint
# shellcheck source-path=lib/
source "${SCRIPT_DIR}/lib/${module}.sh"
```

### "Not following sourced file"

```bash
# Problem: ShellCheck can't find file
source ./lib/utils.sh

# Solution: Provide explicit source
# shellcheck source=./lib/utils.sh
source ./lib/utils.sh
```

### Checking Files Without .sh Extension

```bash
# Use -a to auto-detect or specify shell
shellcheck -s bash script_without_extension

# Or in file
#!/usr/bin/env bash
# shellcheck shell=bash
```
