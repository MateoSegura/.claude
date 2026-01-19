# Command Rules (SH-CMD-*)

Command substitution and arithmetic operations should follow consistent patterns for readability and safety.

## Command Strategy

- Use `$()` for command substitution
- Handle substitution failures explicitly
- Use `(( ))` for arithmetic operations

---

## SH-CMD-001: Use $() for Command Substitution :red_circle:

**Tier**: Critical

**Rationale**: `$()` is more readable, can be nested easily, and handles complex quoting better than backticks.

```bash
# Correct
local current_date
current_date="$(date +%Y-%m-%d)"

local git_root
git_root="$(git rev-parse --show-toplevel)"

# Nested substitution
local user_home
user_home="$(dirname "$(getent passwd "${USER}" | cut -d: -f6)")"

# Multi-line command
local output
output="$(
  find . -name "*.txt" -print0 |
    xargs -0 grep -l "pattern"
)"
```

```bash
# Incorrect
local current_date
current_date=`date +%Y-%m-%d`
# ^ Backticks are harder to read and nest

local git_root
git_root=`git rev-parse --show-toplevel`
```

**ShellCheck**: SC2006 (use $(...) notation instead of backticks)

---

## SH-CMD-002: Handle Command Substitution Failures :yellow_circle:

**Tier**: Required

**Rationale**: Command substitutions can fail. With `set -e`, failures in assignments are not caught. Use explicit checks.

```bash
# Correct
local output
if ! output="$(some_command 2>&1)"; then
  die "Command failed: ${output}"
fi

# Or check afterward
local result
result="$(some_command)" || die "some_command failed"

# For commands that may legitimately return empty
local value
value="$(get_optional_value)" || true
if [[ -z "${value}" ]]; then
  value="default"
fi

# Capture both output and exit code
local output
local exit_code=0
output="$(some_command 2>&1)" || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
  err "Command failed (exit ${exit_code}): ${output}"
fi
```

```bash
# Incorrect - failure may not be caught
local output
output="$(failing_command)"
# With 'set -e', this might not exit as expected in all bash versions
```

---

## SH-CMD-003: Use (( )) for Arithmetic :yellow_circle:

**Tier**: Required

**Rationale**: Arithmetic context `(( ))` is cleaner than `expr` or `let`, supports all arithmetic operations, and does not require `$` for variables.

```bash
# Correct
((count++))
((total = count * 2))
((remaining = total - used))

if (( count > max_count )); then
  die "Count exceeded maximum"
fi

# Arithmetic expansion for assignment
local result=$(( (a + b) * c ))

# Pre/post increment
((++count))   # Pre-increment
((count++))   # Post-increment

# Compound assignment
((total += value))
((count -= 1))
((product *= factor))
```

```bash
# Incorrect
count=`expr $count + 1`   # Antiquated, slow
let count=count+1         # Less readable
count=$[$count + 1]       # Deprecated syntax
```

---

## Additional Patterns

### Command Output to Array

```bash
# Using mapfile (bash 4+)
mapfile -t lines < "${file}"

# Or using read loop
local -a lines=()
while IFS= read -r line; do
  lines+=("${line}")
done < "${file}"
```

### Suppress Output

```bash
# Suppress stdout
command > /dev/null

# Suppress stderr
command 2> /dev/null

# Suppress both
command &> /dev/null
# or
command > /dev/null 2>&1
```

### Here Documents

```bash
# Multi-line string
cat << 'EOF'
This is a multi-line
string that preserves
all formatting.
EOF

# With variable expansion
cat << EOF
User: ${USER}
Home: ${HOME}
EOF

# Indent-tolerant (<<-)
function usage() {
  cat <<- EOF
		Usage: ${SCRIPT_NAME} [options]

		Options:
		  -h    Show help
		  -v    Verbose mode
	EOF
}
```

### Here Strings

```bash
# Single line input
read -r name <<< "default value"

# To a command
grep "pattern" <<< "${variable}"
```

### Process Substitution

```bash
# Use command output as file
diff <(command1) <(command2)

# Read from command (no subshell)
while IFS= read -r line; do
  process "${line}"
done < <(generate_lines)
```

### Redirecting File Descriptors

```bash
# Save and restore stdout
exec 3>&1                    # Save stdout to fd 3
exec > "${log_file}"         # Redirect stdout to file
echo "This goes to file"
exec 1>&3                    # Restore stdout from fd 3
exec 3>&-                    # Close fd 3

# Redirect stderr to stdout
command 2>&1

# Swap stdout and stderr
command 3>&1 1>&2 2>&3 3>&-
```

### Checking Command Existence

```bash
# Correct way to check if command exists
if command -v jq &> /dev/null; then
  echo "jq is available"
fi

# Check and use
if ! command -v docker &> /dev/null; then
  die "Docker is required but not installed"
fi
```
