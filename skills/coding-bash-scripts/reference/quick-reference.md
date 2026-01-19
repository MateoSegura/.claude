# Quick Reference - All Rules

Complete table of all rules with IDs, descriptions, and tiers.

## Rule Tiers

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Rules by Category

### File Structure (SH-STR-*)

| ID | Rule | Tier |
|----|------|------|
| SH-STR-001 | Use consistent shebang (`#!/usr/bin/env bash`) | Critical |
| SH-STR-002 | Enable strict mode (`set -euo pipefail`) | Critical |
| SH-STR-003 | Organize code in standard order | Required |

### Naming (SH-NAM-*)

| ID | Rule | Tier |
|----|------|------|
| SH-NAM-001 | Use descriptive variable names | Required |
| SH-NAM-002 | Use readonly for constants | Required |

### Variables (SH-VAR-*)

| ID | Rule | Tier |
|----|------|------|
| SH-VAR-001 | Always quote variable expansions | Critical |
| SH-VAR-002 | Use braces for variable clarity | Required |
| SH-VAR-003 | Use arrays for lists | Required |
| SH-VAR-004 | Declare local variables in functions | Critical |

### Error Handling (SH-ERR-*)

| ID | Rule | Tier |
|----|------|------|
| SH-ERR-001 | Implement cleanup with trap | Critical |
| SH-ERR-002 | Use error output function | Required |
| SH-ERR-003 | Validate inputs at boundaries | Critical |
| SH-ERR-004 | Check command availability | Required |

### Control Flow (SH-CTL-*)

| ID | Rule | Tier |
|----|------|------|
| SH-CTL-001 | Use `[[ ]]` for conditionals | Critical |
| SH-CTL-002 | Use explicit comparison operators | Required |
| SH-CTL-003 | Use proper loop constructs | Required |

### Commands (SH-CMD-*)

| ID | Rule | Tier |
|----|------|------|
| SH-CMD-001 | Use `$()` for command substitution | Critical |
| SH-CMD-002 | Handle command substitution failures | Required |
| SH-CMD-003 | Use `(( ))` for arithmetic | Required |

### Functions (SH-FUN-*)

| ID | Rule | Tier |
|----|------|------|
| SH-FUN-001 | Use function keyword with braces | Required |
| SH-FUN-002 | Document functions with comments | Required |
| SH-FUN-003 | Return meaningful exit codes | Required |

### Security (SH-SEC-*)

| ID | Rule | Tier |
|----|------|------|
| SH-SEC-001 | Never use eval with external input | Critical |
| SH-SEC-002 | Use mktemp for temporary files | Critical |
| SH-SEC-003 | Validate and sanitize input | Critical |
| SH-SEC-004 | Never store secrets in scripts | Critical |
| SH-SEC-005 | Use restricted PATH | Required |

### Testing (SH-TST-*)

| ID | Rule | Tier |
|----|------|------|
| SH-TST-001 | Write tests for shell scripts | Required |
| SH-TST-002 | Test edge cases | Required |

### Documentation (SH-DOC-*)

| ID | Rule | Tier |
|----|------|------|
| SH-DOC-001 | Write meaningful comments | Recommended |
| SH-DOC-002 | Include file header | Required |

### Portability (SH-PRT-*)

| ID | Rule | Tier |
|----|------|------|
| SH-PRT-001 | Document Bash version requirements | Recommended |

---

## Summary by Tier

### Critical Rules (12 total)

| ID | Rule |
|----|------|
| SH-STR-001 | Use consistent shebang |
| SH-STR-002 | Enable strict mode |
| SH-VAR-001 | Always quote variable expansions |
| SH-VAR-004 | Declare local variables in functions |
| SH-ERR-001 | Implement cleanup with trap |
| SH-ERR-003 | Validate inputs at boundaries |
| SH-CTL-001 | Use [[ ]] for conditionals |
| SH-CMD-001 | Use $() for command substitution |
| SH-SEC-001 | Never use eval with external input |
| SH-SEC-002 | Use mktemp for temporary files |
| SH-SEC-003 | Validate and sanitize input |
| SH-SEC-004 | Never store secrets in scripts |

### Required Rules (17 total)

| ID | Rule |
|----|------|
| SH-STR-003 | Organize code in standard order |
| SH-NAM-001 | Use descriptive variable names |
| SH-NAM-002 | Use readonly for constants |
| SH-VAR-002 | Use braces for variable clarity |
| SH-VAR-003 | Use arrays for lists |
| SH-ERR-002 | Use error output function |
| SH-ERR-004 | Check command availability |
| SH-CTL-002 | Use explicit comparison operators |
| SH-CTL-003 | Use proper loop constructs |
| SH-CMD-002 | Handle command substitution failures |
| SH-CMD-003 | Use (( )) for arithmetic |
| SH-FUN-001 | Use function keyword with braces |
| SH-FUN-002 | Document functions with comments |
| SH-FUN-003 | Return meaningful exit codes |
| SH-SEC-005 | Use restricted PATH |
| SH-TST-001 | Write tests for shell scripts |
| SH-TST-002 | Test edge cases |
| SH-DOC-002 | Include file header |

### Recommended Rules (2 total)

| ID | Rule |
|----|------|
| SH-DOC-001 | Write meaningful comments |
| SH-PRT-001 | Document Bash version requirements |

---

## ShellCheck Mapping

| Rule ID | ShellCheck Code | Description |
|---------|-----------------|-------------|
| SH-STR-001 | SC2096 | Shebang with arguments not portable |
| SH-VAR-001 | SC2086 | Double quote to prevent globbing |
| SH-VAR-003 | SC2207 | Prefer mapfile or read -a |
| SH-VAR-004 | SC2034 | Variable appears unused (scope issue) |
| SH-ERR-004 | SC2230 | Use command -v instead of which |
| SH-CTL-001 | SC2039 | Bashism in sh script |
| SH-CTL-003 | SC2044 | For loops over find are fragile |
| SH-CMD-001 | SC2006 | Use $() instead of backticks |
| SH-SEC-001 | SC2091 | Remove $() to avoid running output |

---

## Essential Patterns

### Script Template

```bash
#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_NAME="$(basename "${BASH_SOURCE[0]}")"

function main() {
  # Main logic
  :
}

main "$@"
```

### Error Functions

```bash
function err() { echo "[ERROR] $*" >&2; }
function die() { err "$@"; exit 1; }
```

### Cleanup Pattern

```bash
trap cleanup EXIT ERR INT TERM
function cleanup() {
  local exit_code=$?
  # cleanup logic
  exit "${exit_code}"
}
```
