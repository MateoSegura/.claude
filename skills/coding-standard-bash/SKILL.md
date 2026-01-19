---
name: bash-standard
description: Complete Bash/Shell scripting coding standard reference (project)
---

# Bash/Shell Scripting Coding Standard

> **Version**: 1.0.0 | **Status**: Active
> **Base Standards**: Google Shell Style Guide, ShellCheck

This standard establishes coding conventions for Bash/Shell script development, ensuring consistency, robustness, security, and maintainability across all shell scripts.

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [File Structure](rules/structure.md) | SH-STR-* | Shebang, strict mode, organization |
| [Naming](rules/naming.md) | SH-NAM-* | Variables, functions, constants |
| [Variables](rules/variables.md) | SH-VAR-* | Quoting, arrays, local scope |
| [Error Handling](rules/error-handling.md) | SH-ERR-* | Trap, validation, exit codes |
| [Control Flow](rules/control-flow.md) | SH-CTL-* | Conditionals, loops |
| [Commands](rules/commands.md) | SH-CMD-* | Substitution, arithmetic |
| [Functions](rules/functions.md) | SH-FUN-* | Definition, documentation, returns |
| [Security](rules/security.md) | SH-SEC-* | Eval, temp files, secrets |
| [Testing](rules/testing.md) | SH-TST-* | Bats tests, edge cases |
| [Documentation](rules/documentation.md) | SH-DOC-* | Comments, file headers |
| [Portability](rules/portability.md) | SH-PRT-* | Bash version requirements |

### Tooling

| Tool | File |
|------|------|
| [ShellCheck](tooling/shellcheck.md) | Static analysis configuration |
| [shfmt](tooling/shfmt.md) | Code formatting |
| [Bats](tooling/bats.md) | Testing framework |
| [Pre-commit](tooling/pre-commit.md) | Git hooks setup |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table with tiers |
| [Code Review Checklist](reference/code-review.md) | Review checklist by category |

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## When to Use Shell Scripts

Shell scripts are appropriate for:
- Small utilities and wrapper scripts (under 100 lines)
- Simple automation tasks
- System initialization and configuration
- Glue code between other programs

**Do not use shell scripts when:**
- The script exceeds 200 lines of actual code
- Complex data structures are needed
- Performance is critical
- Cross-platform portability beyond Unix is required

In these cases, use Python, Go, or another structured language.

---

## Critical Rules (Always Apply)

These rules are non-negotiable and must be followed in all scripts.

### SH-STR-001: Use Consistent Shebang :red_circle:

```bash
# Correct
#!/usr/bin/env bash

# Also acceptable for system scripts
#!/bin/bash
```

### SH-STR-002: Enable Strict Mode :red_circle:

Strict mode catches common errors early and prevents scripts from continuing after failures.

```bash
#!/usr/bin/env bash
set -euo pipefail

# For debugging (add -x)
set -euxo pipefail
```

| Option | Effect |
|--------|--------|
| `-e` (errexit) | Exit immediately if any command fails |
| `-u` (nounset) | Treat unset variables as errors |
| `-o pipefail` | Pipeline fails if any command fails |

### SH-VAR-001: Always Quote Variable Expansions :red_circle:

Unquoted variables undergo word splitting and glob expansion, leading to bugs and injection vulnerabilities.

```bash
# Correct
local filename="$1"
cp "${source_file}" "${dest_dir}/"
for file in "${files[@]}"; do
  process "${file}"
done

# Incorrect - word splitting breaks filenames with spaces
cp $source_file $dest_dir/
```

### SH-VAR-004: Declare Local Variables in Functions :red_circle:

Without `local`, variables leak into global scope.

```bash
# Correct
function process_file() {
  local filename="$1"
  local line_count=0
  # ...
}

# Incorrect - pollutes global scope
function process_file() {
  filename="$1"
  line_count=0
}
```

### SH-ERR-001: Implement Cleanup with trap :red_circle:

Scripts that create temporary files or modify system state must clean up on exit.

```bash
TEMP_DIR=""

function cleanup() {
  local exit_code=$?
  if [[ -n "${TEMP_DIR}" && -d "${TEMP_DIR}" ]]; then
    rm -rf "${TEMP_DIR}"
  fi
  exit "${exit_code}"
}

trap cleanup EXIT ERR INT TERM

function main() {
  TEMP_DIR="$(mktemp -d)"
  # Script logic...
}
```

### SH-ERR-003: Validate Inputs at Boundaries :red_circle:

Early validation prevents cascading failures and provides clear error messages.

```bash
function validate_arguments() {
  if [[ $# -lt 1 ]]; then
    die "Usage: ${SCRIPT_NAME} <input_file>"
  fi

  local input_file="$1"
  if [[ ! -f "${input_file}" ]]; then
    die "Input file does not exist: ${input_file}"
  fi
}
```

### SH-CTL-001: Use [[ ]] for Conditionals :red_circle:

`[[ ]]` is a bash builtin with better features: no word splitting, regex support, logical operators.

```bash
# Correct
if [[ -f "${file}" ]]; then
  echo "File exists"
fi

if [[ "${string}" =~ ^[0-9]+$ ]]; then
  echo "Is numeric"
fi

# Incorrect - [ ] requires more quoting, no regex
if [ -f "${file}" ]; then
  echo "File exists"
fi
```

### SH-CMD-001: Use $() for Command Substitution :red_circle:

`$()` is more readable, can be nested easily, and handles complex quoting better than backticks.

```bash
# Correct
current_date="$(date +%Y-%m-%d)"

# Incorrect
current_date=`date +%Y-%m-%d`
```

### SH-SEC-001: Never Use eval with External Input :red_circle:

`eval` executes arbitrary code. Combined with external input, it creates command injection vulnerabilities.

```bash
# Correct - Use arrays for dynamic commands
local -a cmd=("$@")
"${cmd[@]}"

# Incorrect - SECURITY VULNERABILITY
eval "echo ${user_input}"
```

### SH-SEC-002: Use mktemp for Temporary Files :red_circle:

Manual temp file creation is vulnerable to race conditions and symlink attacks.

```bash
# Correct
temp_file="$(mktemp)" || die "Failed to create temp file"
trap 'rm -f "${temp_file}"' EXIT

# Incorrect - Predictable names
temp_file="/tmp/myapp.$$"
```

### SH-SEC-003: Validate and Sanitize Input :red_circle:

External input can contain shell metacharacters or unexpected values.

```bash
# Correct - Validate expected format
function validate_username() {
  local username="$1"
  if [[ ! "${username}" =~ ^[a-zA-Z][a-zA-Z0-9_-]{2,31}$ ]]; then
    die "Invalid username format: ${username}"
  fi
}
```

### SH-SEC-004: Never Store Secrets in Scripts :red_circle:

Scripts are often stored in version control and are readable.

```bash
# Correct - Use environment variables
readonly DB_PASSWORD="${DB_PASSWORD:?DB_PASSWORD required}"

# Incorrect - NEVER DO THIS
readonly API_KEY="sk-1234567890abcdef"
```

---

## Quick Rule Lookup

| ID | Rule | Tier |
|----|------|------|
| SH-STR-001 | Use consistent shebang | Critical |
| SH-STR-002 | Enable strict mode (set -euo pipefail) | Critical |
| SH-STR-003 | Organize code in standard order | Required |
| SH-NAM-001 | Use descriptive variable names | Required |
| SH-NAM-002 | Use readonly for constants | Required |
| SH-VAR-001 | Always quote variable expansions | Critical |
| SH-VAR-002 | Use braces for variable clarity | Required |
| SH-VAR-003 | Use arrays for lists | Required |
| SH-VAR-004 | Declare local variables in functions | Critical |
| SH-ERR-001 | Implement cleanup with trap | Critical |
| SH-ERR-002 | Use error output function | Required |
| SH-ERR-003 | Validate inputs at boundaries | Critical |
| SH-ERR-004 | Check command availability | Required |
| SH-CTL-001 | Use [[ ]] for conditionals | Critical |
| SH-CTL-002 | Use explicit comparison operators | Required |
| SH-CTL-003 | Use proper loop constructs | Required |
| SH-CMD-001 | Use $() for command substitution | Critical |
| SH-CMD-002 | Handle command substitution failures | Required |
| SH-CMD-003 | Use (( )) for arithmetic | Required |
| SH-FUN-001 | Use function keyword with braces | Required |
| SH-FUN-002 | Document functions with comments | Required |
| SH-FUN-003 | Return meaningful exit codes | Required |
| SH-SEC-001 | Never use eval with external input | Critical |
| SH-SEC-002 | Use mktemp for temporary files | Critical |
| SH-SEC-003 | Validate and sanitize input | Critical |
| SH-SEC-004 | Never store secrets in scripts | Critical |
| SH-SEC-005 | Use restricted PATH | Required |
| SH-TST-001 | Write tests for shell scripts | Required |
| SH-TST-002 | Test edge cases | Required |
| SH-DOC-001 | Write meaningful comments | Recommended |
| SH-DOC-002 | Include file header | Required |
| SH-PRT-001 | Document Bash version requirements | Recommended |

---

## References

- [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html)
- [ShellCheck Wiki](https://github.com/koalaman/shellcheck/wiki)
- [shfmt - Shell Formatter](https://github.com/mvdan/sh)
- [Bats - Bash Automated Testing System](https://github.com/bats-core/bats-core)
- [Bash Reference Manual](https://www.gnu.org/software/bash/manual/bash.html)
