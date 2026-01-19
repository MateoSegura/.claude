# Code Review Checklist

Use this checklist during code reviews to ensure compliance with the Bash coding standard.

## Structure and Organization

- [ ] Shebang is `#!/usr/bin/env bash` or `#!/bin/bash`
- [ ] `set -euo pipefail` is present immediately after shebang
- [ ] File follows standard organization (constants, functions, main)
- [ ] Script length is under 200 lines (consider refactoring if larger)
- [ ] No executable code between function definitions

**Related Rules**: SH-STR-001, SH-STR-002, SH-STR-003

---

## Naming

- [ ] Constants are `UPPER_SNAKE_CASE` and `readonly`
- [ ] Local variables are `lower_snake_case`
- [ ] Function names are descriptive and `lower_snake_case`
- [ ] No single-letter variables except loop counters in small scope

**Related Rules**: SH-NAM-001, SH-NAM-002

---

## Quoting and Variables

- [ ] ALL variable expansions are quoted: `"${variable}"`
- [ ] Arrays are used for lists (not space-separated strings)
- [ ] Local variables declared with `local` in functions
- [ ] No unquoted `$@` or `$*` (use `"$@"`)
- [ ] Braces used consistently: `${var}` not `$var`

**Related Rules**: SH-VAR-001, SH-VAR-002, SH-VAR-003, SH-VAR-004

---

## Error Handling

- [ ] `trap` is used for cleanup on EXIT/ERR
- [ ] Input validation at function boundaries
- [ ] Meaningful error messages written to stderr
- [ ] Appropriate exit codes used
- [ ] Required commands checked with `command -v`

**Related Rules**: SH-ERR-001, SH-ERR-002, SH-ERR-003, SH-ERR-004

---

## Control Flow

- [ ] `[[ ]]` used instead of `[ ]` for conditionals
- [ ] `$(command)` used instead of backticks
- [ ] `(( ))` used for arithmetic
- [ ] No `eval` with external input
- [ ] Process substitution used instead of pipes when variables modified

**Related Rules**: SH-CTL-001, SH-CTL-002, SH-CTL-003, SH-CMD-001, SH-CMD-003

---

## Functions

- [ ] Functions use `function name()` or `name()` consistently
- [ ] Functions have documentation comments
- [ ] Functions return meaningful exit codes
- [ ] Local variables declared at function start

**Related Rules**: SH-FUN-001, SH-FUN-002, SH-FUN-003

---

## Security

- [ ] No hardcoded secrets or credentials
- [ ] `mktemp` used for temporary files
- [ ] External input validated and sanitized
- [ ] PATH is explicit or restricted
- [ ] No SUID/SGID scripts
- [ ] No `eval` with user input

**Related Rules**: SH-SEC-001, SH-SEC-002, SH-SEC-003, SH-SEC-004, SH-SEC-005

---

## Documentation

- [ ] File header with description and usage
- [ ] Functions have documentation comments
- [ ] Complex logic has explanatory comments
- [ ] Comments explain "why" not "what"

**Related Rules**: SH-DOC-001, SH-DOC-002

---

## Testing

- [ ] Tests exist for non-trivial scripts
- [ ] Edge cases tested (empty input, special characters, spaces)
- [ ] Error conditions tested

**Related Rules**: SH-TST-001, SH-TST-002

---

## Tooling

- [ ] ShellCheck passes with no warnings
- [ ] shfmt formatting applied
- [ ] No disabled ShellCheck rules without justification

---

## Quick Checks

### Critical (Must Fix)

```
[ ] set -euo pipefail present
[ ] All variables quoted
[ ] Local variables in functions
[ ] trap cleanup on EXIT
[ ] Input validated
[ ] No eval with external input
[ ] mktemp for temp files
[ ] No hardcoded secrets
```

### Required (Fix Before Merge)

```
[ ] [[ ]] for conditionals
[ ] $() for command substitution
[ ] (( )) for arithmetic
[ ] Functions documented
[ ] Errors to stderr
[ ] Commands checked with command -v
[ ] Arrays for lists
```

### Recommended (Consider Fixing)

```
[ ] Comments explain why
[ ] Bash version documented
[ ] File header present
```

---

## Common Issues to Watch For

### Word Splitting Bugs

```bash
# Bug: Breaks on filenames with spaces
for file in $(find . -name "*.txt"); do
  process "${file}"
done

# Fix: Use null-separated output
while IFS= read -r -d '' file; do
  process "${file}"
done < <(find . -name "*.txt" -print0)
```

### Variable Scope Leaks

```bash
# Bug: Variables leak to global scope
function bad() {
  result="value"  # Global!
}

# Fix: Use local
function good() {
  local result="value"
}
```

### Subshell Variable Loss

```bash
# Bug: Variable lost in subshell
count=0
some_cmd | while read -r line; do
  ((count++))
done
echo "${count}"  # Still 0!

# Fix: Use process substitution
count=0
while read -r line; do
  ((count++))
done < <(some_cmd)
echo "${count}"  # Correct value
```

### Missing Cleanup

```bash
# Bug: Temp file left behind on error
temp=$(mktemp)
process "${temp}"
rm "${temp}"  # Never reached on error

# Fix: Use trap
temp=$(mktemp)
trap 'rm -f "${temp}"' EXIT
process "${temp}"
```

### Numeric vs String Comparison

```bash
# Bug: String comparison for numbers
if [[ "${count}" > "10" ]]; then  # Wrong! "9" > "10" lexically

# Fix: Use numeric comparison
if (( count > 10 )); then
# or
if [[ "${count}" -gt 10 ]]; then
```
