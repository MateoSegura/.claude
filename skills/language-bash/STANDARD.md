# Bash/Shell Scripting Coding Standard

> **Version**: 1.0.0
> **Status**: Active
> **Base Standards**: [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html), [ShellCheck](https://github.com/koalaman/shellcheck)
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes coding conventions for Bash/Shell script development. It ensures:

- **Consistency**: Uniform shell scripting style across all projects
- **Robustness**: Scripts that fail safely and predictably
- **Security**: Prevention of injection attacks and race conditions
- **Maintainability**: Code that is easy to read, debug, and extend

### 1.2 Scope

This standard applies to:
- [x] All Bash shell scripts in production repositories
- [x] Build and deployment automation scripts
- [x] Development tooling and helper scripts
- [x] CI/CD pipeline scripts
- [x] System administration scripts

### 1.3 Audience

- Software engineers writing shell scripts
- DevOps engineers creating automation
- Code reviewers evaluating shell script changes
- System administrators maintaining infrastructure scripts

### 1.4 Relationship to Industry Standards

| Standard | Relationship |
|----------|--------------|
| [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html) | **Base** - All rules apply unless documented otherwise |
| [ShellCheck](https://github.com/koalaman/shellcheck) | **Required** - All scripts must pass ShellCheck |
| POSIX Shell | **Reference** - Consulted for portability considerations |

### 1.5 When to Use Shell Scripts

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
- Complex error handling or recovery is needed

In these cases, use Python, Go, or another structured language.

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

Each rule includes:
- **Rule ID**: Unique identifier (e.g., `SH-ERR-001`)
- **Tier**: Enforcement level
- **Rationale**: Why this rule exists
- **Example**: Correct and incorrect code
- **ShellCheck**: Related ShellCheck codes (where applicable)

---

## 3. File Structure and Organization

### 3.1 File Extensions

| Type | Extension | Example |
|------|-----------|---------|
| Executable scripts (in PATH) | No extension | `deploy` |
| Library/source files | `.sh` | `utils.sh` |
| Test files | `_test.sh` or `.bats` | `utils_test.sh` |

### 3.2 Standard File Layout

```bash
#!/usr/bin/env bash
#
# Script description: One-line summary of what this script does.
#
# Usage: script_name [options] <arguments>
#
# Options:
#   -h, --help     Show this help message
#   -v, --verbose  Enable verbose output
#
# Examples:
#   script_name -v input.txt
#

# Strict mode settings
set -euo pipefail

# Constants (readonly, uppercase)
readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_NAME="$(basename "${BASH_SOURCE[0]}")"
readonly VERSION="1.0.0"

# Global variables (uppercase)
VERBOSE=false
DRY_RUN=false

# Source dependencies
source "${SCRIPT_DIR}/lib/utils.sh"

#######################################
# Function descriptions use this format.
# Globals:
#   VERBOSE
# Arguments:
#   $1 - Description of first argument
#   $2 - Description of second argument (optional)
# Outputs:
#   Writes to stdout
# Returns:
#   0 on success, non-zero on error
#######################################
function example_function() {
  local arg1="$1"
  local arg2="${2:-default}"
  # Function body
}

#######################################
# Main entry point.
# Globals:
#   All global variables
# Arguments:
#   Command line arguments
#######################################
function main() {
  parse_arguments "$@"
  # Main logic here
}

# Script entry point - must be at bottom
main "$@"
```

### 3.3 File Structure Rules

#### SH-STR-001: Use Consistent Shebang :red_circle:

**Rationale**: A consistent shebang ensures scripts execute with the expected interpreter. Using `#!/usr/bin/env bash` provides better portability across systems where bash may be installed in different locations.

```bash
# Correct
#!/usr/bin/env bash

# Also acceptable for system scripts
#!/bin/bash
```

```bash
# Incorrect
#!/bin/sh
# ^ May not support bash features

#! /usr/bin/env bash
# ^ Space after #! is non-standard

#!/usr/bin/env bash -e
# ^ Options in env shebang are not portable
```

**ShellCheck**: SC2096 (shebang with arguments not portable)

---

#### SH-STR-002: Enable Strict Mode :red_circle:

**Rationale**: Strict mode catches common errors early and prevents scripts from continuing after failures. This is the single most important practice for robust shell scripts.

```bash
# Correct
#!/usr/bin/env bash
set -euo pipefail

# For debugging (add -x)
set -euxo pipefail
```

```bash
# Incorrect
#!/usr/bin/env bash
# Missing strict mode - script continues after errors

# Partial strict mode
set -e
# ^ Missing -u (unset variables) and pipefail
```

| Option | Effect |
|--------|--------|
| `-e` (errexit) | Exit immediately if any command fails |
| `-u` (nounset) | Treat unset variables as errors |
| `-o pipefail` | Pipeline fails if any command fails |
| `-x` (xtrace) | Print commands before execution (debugging) |

**Note**: Place `set` commands immediately after the shebang, before any other code.

---

#### SH-STR-003: Organize Code in Standard Order :yellow_circle:

**Rationale**: Consistent organization makes scripts easier to navigate and understand.

```bash
# Correct order:
# 1. Shebang
# 2. File header comment
# 3. set statements (strict mode)
# 4. Constants (readonly)
# 5. Global variables
# 6. Source/include statements
# 7. Function definitions
# 8. Main function
# 9. main "$@" call at bottom
```

```bash
# Incorrect - executable code between functions
function setup() { ... }

echo "Starting..."  # Don't put code here!

function cleanup() { ... }
```

---

## 4. Naming Conventions

### 4.1 General Principles

- Names should be descriptive and reveal intent
- Use lowercase with underscores for most identifiers
- Use UPPERCASE for constants and environment variables
- Avoid abbreviations except for well-known terms (e.g., `tmp`, `err`, `msg`)

### 4.2 Specific Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Constants | `UPPER_SNAKE_CASE` | `MAX_RETRIES`, `DEFAULT_TIMEOUT` |
| Environment variables | `UPPER_SNAKE_CASE` | `LOG_LEVEL`, `CONFIG_PATH` |
| Local variables | `lower_snake_case` | `file_count`, `user_input` |
| Functions | `lower_snake_case` | `process_file`, `validate_input` |
| Script files | `lower-kebab-case` or `lower_snake_case` | `deploy-app`, `run_tests.sh` |
| Private functions | `_leading_underscore` | `_internal_helper` |

### 4.3 Naming Rules

#### SH-NAM-001: Use Descriptive Variable Names :yellow_circle:

**Rationale**: Clear names reduce cognitive load and make code self-documenting. Single-letter variables obscure meaning and make maintenance difficult.

```bash
# Correct
readonly max_retry_count=3
local input_file="$1"
local line_number=0
local error_message=""

for config_file in "${config_files[@]}"; do
  process_config "${config_file}"
done
```

```bash
# Incorrect
readonly n=3
local f="$1"
local l=0
local e=""

for c in "${arr[@]}"; do
  proc "${c}"
done
```

**Exception**: Loop counters (`i`, `j`, `k`) and standard conventions (`fd` for file descriptor) are acceptable in limited scopes.

---

#### SH-NAM-002: Use readonly for Constants :yellow_circle:

**Rationale**: Declaring constants as `readonly` prevents accidental modification and signals intent to readers.

```bash
# Correct
readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly CONFIG_FILE="/etc/myapp/config.conf"
readonly -a SUPPORTED_FORMATS=("json" "yaml" "toml")
```

```bash
# Incorrect
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONFIG_FILE="/etc/myapp/config.conf"
# ^ Can be accidentally modified later
```

---

## 5. Variables and Quoting

### 5.1 Quoting Rules

**The Golden Rule: Quote Everything**

Unless you specifically need word splitting or glob expansion, quote all variable expansions.

#### SH-VAR-001: Always Quote Variable Expansions :red_circle:

**Rationale**: Unquoted variables undergo word splitting and glob expansion, leading to bugs with filenames containing spaces, injection vulnerabilities, and unexpected behavior.

```bash
# Correct
local filename="$1"
local message="${error_prefix}: ${error_message}"
cp "${source_file}" "${dest_dir}/"
echo "Processing: ${filename}"

# Correct - array expansion
local -a files=("file 1.txt" "file 2.txt")
for file in "${files[@]}"; do
  process "${file}"
done
```

```bash
# Incorrect
local filename=$1
# ^ Breaks with filenames containing spaces

cp $source_file $dest_dir/
# ^ Word splitting + glob expansion = disaster

echo Processing: $filename
# ^ Word splitting breaks output

for file in ${files[@]}; do
  # ^ Array elements with spaces will be split
  process $file
done
```

**ShellCheck**: SC2086 (double quote to prevent globbing and word splitting)

---

#### SH-VAR-002: Use Braces for Variable Clarity :yellow_circle:

**Rationale**: Braces prevent ambiguity in variable names and enable parameter expansion features.

```bash
# Correct
echo "${variable}"
echo "${filename%.txt}.json"  # Parameter expansion
echo "${user}_home"           # Clear boundary
local path="${base_dir}/${sub_dir}"

# Default values
local config="${CONFIG_FILE:-/etc/default.conf}"
local timeout="${TIMEOUT:=30}"  # Also assigns if unset
```

```bash
# Incorrect
echo "$variable"suffix    # Ambiguous - is it $variable or $variablesuffix?
echo "$user"_home         # Works but inconsistent
```

---

#### SH-VAR-003: Use Arrays for Lists :yellow_circle:

**Rationale**: Arrays properly handle elements containing spaces and special characters. String-based lists break on whitespace.

```bash
# Correct
local -a files=()
while IFS= read -r -d '' file; do
  files+=("${file}")
done < <(find . -name "*.txt" -print0)

# Iterate safely
for file in "${files[@]}"; do
  process "${file}"
done

# Pass array to function
process_files "${files[@]}"
```

```bash
# Incorrect
files=$(find . -name "*.txt")
# ^ Filenames with spaces/newlines break

for file in $files; do
  # ^ Word splitting corrupts filenames
  process "$file"
done
```

**ShellCheck**: SC2207 (prefer mapfile or read -a)

---

#### SH-VAR-004: Declare Local Variables in Functions :red_circle:

**Rationale**: Without `local`, variables leak into global scope, causing hard-to-debug issues and potential security vulnerabilities.

```bash
# Correct
function process_file() {
  local filename="$1"
  local line_count=0
  local -a lines=()
  local -r max_lines=1000  # Local readonly
  local -i counter=0       # Integer type

  while IFS= read -r line; do
    ((line_count++))
    lines+=("${line}")
  done < "${filename}"
}
```

```bash
# Incorrect
function process_file() {
  filename="$1"      # Pollutes global scope!
  line_count=0       # Will affect callers!

  while IFS= read -r line; do
    ((line_count++))
  done < "${filename}"
}
```

**ShellCheck**: SC2034 (variable appears unused - may indicate scope issue)

---

## 6. Error Handling

### 6.1 Error Handling Strategy

1. Enable strict mode (`set -euo pipefail`)
2. Use `trap` for cleanup on exit
3. Validate inputs at boundaries
4. Provide meaningful error messages
5. Use appropriate exit codes

### 6.2 Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Misuse of command/invalid arguments |
| 126 | Command not executable |
| 127 | Command not found |
| 128+N | Fatal error signal N |

### 6.3 Error Handling Rules

#### SH-ERR-001: Implement Cleanup with trap :red_circle:

**Rationale**: Scripts that create temporary files, acquire locks, or modify system state must clean up on exit, including when terminated by signals.

```bash
# Correct
#!/usr/bin/env bash
set -euo pipefail

TEMP_DIR=""
LOCK_FILE=""

function cleanup() {
  local exit_code=$?

  # Remove temporary directory
  if [[ -n "${TEMP_DIR}" && -d "${TEMP_DIR}" ]]; then
    rm -rf "${TEMP_DIR}"
  fi

  # Release lock
  if [[ -n "${LOCK_FILE}" && -f "${LOCK_FILE}" ]]; then
    rm -f "${LOCK_FILE}"
  fi

  exit "${exit_code}"
}

# Register cleanup for multiple signals
trap cleanup EXIT ERR INT TERM

function main() {
  TEMP_DIR="$(mktemp -d)"
  LOCK_FILE="/var/lock/myapp.lock"

  # Script logic...
}

main "$@"
```

```bash
# Incorrect - no cleanup
#!/usr/bin/env bash

TEMP_FILE=$(mktemp)
# If script fails, temp file is never removed

echo "data" > "${TEMP_FILE}"
process "${TEMP_FILE}"
rm "${TEMP_FILE}"  # Never reached on error!
```

---

#### SH-ERR-002: Use Error Output Function :yellow_circle:

**Rationale**: Consistent error output format aids debugging and log analysis. Errors should go to stderr, not stdout.

```bash
# Correct
function err() {
  echo "[ERROR] $(date '+%Y-%m-%d %H:%M:%S') $*" >&2
}

function warn() {
  echo "[WARN] $(date '+%Y-%m-%d %H:%M:%S') $*" >&2
}

function info() {
  echo "[INFO] $(date '+%Y-%m-%d %H:%M:%S') $*"
}

function die() {
  err "$@"
  exit 1
}

# Usage
if [[ ! -f "${config_file}" ]]; then
  die "Configuration file not found: ${config_file}"
fi
```

```bash
# Incorrect
if [[ ! -f "${config_file}" ]]; then
  echo "Error: Configuration file not found"
  # ^ Goes to stdout, missing context
  exit 1
fi
```

---

#### SH-ERR-003: Validate Inputs at Boundaries :red_circle:

**Rationale**: Early validation prevents cascading failures and provides clear error messages. Never trust external input.

```bash
# Correct
function validate_arguments() {
  if [[ $# -lt 1 ]]; then
    die "Usage: ${SCRIPT_NAME} <input_file> [output_file]"
  fi

  local input_file="$1"

  if [[ ! -f "${input_file}" ]]; then
    die "Input file does not exist: ${input_file}"
  fi

  if [[ ! -r "${input_file}" ]]; then
    die "Input file is not readable: ${input_file}"
  fi
}

function main() {
  validate_arguments "$@"
  local input_file="$1"
  # Proceed with validated input...
}
```

```bash
# Incorrect - no validation
function main() {
  local input_file="$1"  # May be empty!
  cat "${input_file}"    # Cryptic error if file missing
}
```

---

#### SH-ERR-004: Check Command Availability :yellow_circle:

**Rationale**: Scripts should fail fast with helpful messages if required commands are missing.

```bash
# Correct
function check_dependencies() {
  local -a required_commands=("jq" "curl" "aws")
  local missing=()

  for cmd in "${required_commands[@]}"; do
    if ! command -v "${cmd}" &>/dev/null; then
      missing+=("${cmd}")
    fi
  done

  if [[ ${#missing[@]} -gt 0 ]]; then
    die "Missing required commands: ${missing[*]}"
  fi
}
```

```bash
# Incorrect
jq '.key' data.json
# ^ Cryptic "command not found" if jq missing
```

**ShellCheck**: SC2230 (use command -v instead of which)

---

## 7. Control Structures

### 7.1 Conditionals

#### SH-CTL-001: Use [[ ]] for Conditionals :red_circle:

**Rationale**: `[[ ]]` is a bash builtin with better features: no word splitting, regex support, logical operators, and safer string comparison.

```bash
# Correct
if [[ -f "${file}" ]]; then
  echo "File exists"
fi

if [[ "${string}" == "value" ]]; then
  echo "Match"
fi

if [[ "${string}" =~ ^[0-9]+$ ]]; then
  echo "Is numeric"
fi

if [[ -n "${var}" && "${var}" != "skip" ]]; then
  process "${var}"
fi
```

```bash
# Incorrect
if [ -f "${file}" ]; then
  # ^ [ ] requires more quoting, no regex
  echo "File exists"
fi

if [ "${string}" == "value" ]; then
  # ^ == not POSIX in [ ]
  echo "Match"
fi

if [ -n "${var}" -a "${var}" != "skip" ]; then
  # ^ -a is deprecated
  process "${var}"
fi
```

**ShellCheck**: SC2039 (bashism in sh script)

---

#### SH-CTL-002: Use Explicit Comparison Operators :yellow_circle:

**Rationale**: Explicit operators make code readable and prevent subtle bugs.

```bash
# Correct - String comparisons
[[ "${str}" == "value" ]]   # Equal
[[ "${str}" != "value" ]]   # Not equal
[[ "${str}" < "value" ]]    # Less than (lexicographic)
[[ -z "${str}" ]]           # Is empty
[[ -n "${str}" ]]           # Is not empty

# Correct - Numeric comparisons
[[ "${num}" -eq 10 ]]       # Equal
[[ "${num}" -ne 10 ]]       # Not equal
[[ "${num}" -lt 10 ]]       # Less than
[[ "${num}" -le 10 ]]       # Less than or equal
[[ "${num}" -gt 10 ]]       # Greater than
[[ "${num}" -ge 10 ]]       # Greater than or equal

# Correct - Arithmetic context for numbers
if (( num > 10 )); then
  echo "Greater"
fi
```

```bash
# Incorrect
[[ "${str}" ]]              # Unclear - testing for non-empty?
[[ ! "${str}" ]]            # Confusing negation
[[ "${num}" > 10 ]]         # Wrong! This is string comparison
```

---

### 7.2 Loops

#### SH-CTL-003: Use Proper Loop Constructs :yellow_circle:

**Rationale**: Choose the right loop for the task. Avoid subshells in pipelines when variable modification is needed.

```bash
# Correct - Iterate over array
for item in "${array[@]}"; do
  process "${item}"
done

# Correct - C-style for loop
for ((i = 0; i < count; i++)); do
  process "${i}"
done

# Correct - Read lines from file
while IFS= read -r line; do
  process "${line}"
done < "${input_file}"

# Correct - Read with process substitution (preserves variables)
while IFS= read -r line; do
  ((counter++))
done < <(some_command)
echo "Counter: ${counter}"  # Variable preserved!
```

```bash
# Incorrect - Variable lost in subshell
counter=0
some_command | while IFS= read -r line; do
  ((counter++))  # Modified in subshell
done
echo "Counter: ${counter}"  # Still 0!

# Incorrect - Word splitting on find output
for file in $(find . -name "*.txt"); do
  # Breaks on filenames with spaces
  process "${file}"
done
```

**ShellCheck**: SC2044 (for loops over find output are fragile)

---

## 8. Commands and Substitution

### 8.1 Command Substitution

#### SH-CMD-001: Use $() for Command Substitution :red_circle:

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

#### SH-CMD-002: Handle Command Substitution Failures :yellow_circle:

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
```

```bash
# Incorrect - failure may not be caught
local output
output="$(failing_command)"
# With 'set -e', this might not exit as expected in all bash versions
```

---

### 8.2 Arithmetic

#### SH-CMD-003: Use (( )) for Arithmetic :yellow_circle:

**Rationale**: Arithmetic context `(( ))` is cleaner than `expr` or `let`, supports all arithmetic operations, and does not require `$` for variables.

```bash
# Correct
((count++))
((total = count * 2))
if (( count > max_count )); then
  die "Count exceeded maximum"
fi

# Arithmetic expansion for assignment
local result=$(( (a + b) * c ))
```

```bash
# Incorrect
count=`expr $count + 1`   # Antiquated, slow
let count=count+1         # Less readable
count=$[$count + 1]       # Deprecated syntax
```

---

## 9. Functions

### 9.1 Function Design

#### SH-FUN-001: Use function Keyword with Braces :yellow_circle:

**Rationale**: The `function` keyword clearly identifies function definitions and distinguishes them from command invocations.

```bash
# Correct
function process_file() {
  local filename="$1"
  # Implementation
}

# Also acceptable (POSIX-compatible)
process_file() {
  local filename="$1"
  # Implementation
}
```

```bash
# Incorrect - inconsistent style
function process_file {
  # Missing parentheses - works but non-standard
}
```

---

#### SH-FUN-002: Document Functions with Comments :yellow_circle:

**Rationale**: Function documentation enables understanding without reading implementation details.

```bash
# Correct
#######################################
# Processes a log file and extracts errors.
# Globals:
#   ERROR_PATTERN - regex pattern for errors
#   VERBOSE - if true, print progress
# Arguments:
#   $1 - Path to log file (required)
#   $2 - Output file path (optional, defaults to stdout)
# Outputs:
#   Writes extracted errors to stdout or output file
# Returns:
#   0 - Success
#   1 - Input file not found
#   2 - Input file not readable
#######################################
function extract_errors() {
  local log_file="$1"
  local output_file="${2:-/dev/stdout}"
  # Implementation
}
```

```bash
# Incorrect - no documentation
function extract_errors() {
  grep -E "$ERROR_PATTERN" "$1" > "${2:-/dev/stdout}"
}
```

---

#### SH-FUN-003: Return Meaningful Exit Codes :yellow_circle:

**Rationale**: Functions should communicate success/failure via exit codes. Reserve stdout for data output.

```bash
# Correct
function validate_config() {
  local config_file="$1"

  if [[ ! -f "${config_file}" ]]; then
    err "Config file not found: ${config_file}"
    return 1
  fi

  if ! jq empty "${config_file}" 2>/dev/null; then
    err "Config file is not valid JSON: ${config_file}"
    return 2
  fi

  return 0
}

# Usage
if ! validate_config "${config}"; then
  die "Configuration validation failed"
fi
```

```bash
# Incorrect - returning data as exit code
function get_count() {
  local count
  count=$(wc -l < "$1")
  return "${count}"  # Wrong! Exit codes are 0-255
}
```

---

## 10. Security

### 10.1 Security Principles

1. **Never trust external input** - Validate and sanitize all input
2. **Quote everything** - Prevent injection attacks
3. **Use safe temporary files** - Prevent race conditions
4. **Avoid eval** - Prevents code injection
5. **Principle of least privilege** - Drop permissions when possible

### 10.2 Security Rules

#### SH-SEC-001: Never Use eval with External Input :red_circle:

**Rationale**: `eval` executes arbitrary code. Combined with external input, it creates command injection vulnerabilities.

```bash
# Correct - Use arrays for dynamic commands
function run_with_args() {
  local -a cmd=("$@")
  "${cmd[@]}"
}

# Correct - Use associative arrays for key-value
declare -A config
config["key"]="value"
echo "${config[key]}"

# Correct - Use indirect expansion
local var_name="PATH"
echo "${!var_name}"
```

```bash
# Incorrect - SECURITY VULNERABILITY
local user_input="$1"
eval "echo ${user_input}"
# If user_input is '; rm -rf /' - disaster!

# Incorrect - Dynamic variable names with eval
eval "${var_name}=value"
```

**ShellCheck**: SC2091 (remove surrounding $() to avoid running output)

---

#### SH-SEC-002: Use mktemp for Temporary Files :red_circle:

**Rationale**: Manual temp file creation is vulnerable to race conditions and symlink attacks. `mktemp` creates files atomically with secure permissions.

```bash
# Correct
function main() {
  local temp_file
  temp_file="$(mktemp)" || die "Failed to create temp file"
  trap 'rm -f "${temp_file}"' EXIT

  # Use temp file...
  echo "data" > "${temp_file}"
}

# Correct - Temporary directory
function main() {
  local temp_dir
  temp_dir="$(mktemp -d)" || die "Failed to create temp directory"
  trap 'rm -rf "${temp_dir}"' EXIT

  # Create files inside temp directory
  echo "data" > "${temp_dir}/file1.txt"
}

# Correct - With template
local temp_file
temp_file="$(mktemp -t "myapp.XXXXXX")"
```

```bash
# Incorrect - Predictable names
local temp_file="/tmp/myapp.$$"
echo "data" > "${temp_file}"
# Attacker can predict PID and create symlink

# Incorrect - Insecure location
local temp_file="/tmp/myapp_temp"
# Same file used by all instances - race condition
```

**ShellCheck**: SC2086 applied here prevents issues with space in TMPDIR

---

#### SH-SEC-003: Validate and Sanitize Input :red_circle:

**Rationale**: External input can contain shell metacharacters, null bytes, or unexpected values that cause security issues or crashes.

```bash
# Correct - Validate expected format
function validate_username() {
  local username="$1"

  if [[ ! "${username}" =~ ^[a-zA-Z][a-zA-Z0-9_-]{2,31}$ ]]; then
    die "Invalid username format: ${username}"
  fi
}

# Correct - Validate file path
function validate_path() {
  local path="$1"
  local base_dir="$2"

  # Resolve to absolute path
  local resolved
  resolved="$(realpath -m "${path}")" || die "Invalid path"

  # Check it's under expected directory (prevent traversal)
  if [[ "${resolved}" != "${base_dir}"/* ]]; then
    die "Path traversal detected: ${path}"
  fi
}

# Correct - Sanitize for display
function sanitize_for_log() {
  local input="$1"
  # Remove control characters
  printf '%s' "${input}" | tr -d '[:cntrl:]'
}
```

```bash
# Incorrect - Using input directly
local filename="$1"
cat "${filename}"
# What if filename is "/etc/shadow" or "../../etc/passwd"?
```

---

#### SH-SEC-004: Never Store Secrets in Scripts :red_circle:

**Rationale**: Scripts are often stored in version control and are readable. Secrets in scripts are easily exposed.

```bash
# Correct - Use environment variables
readonly DB_PASSWORD="${DB_PASSWORD:?DB_PASSWORD environment variable required}"

# Correct - Read from secure file
function get_api_key() {
  local key_file="${HOME}/.config/myapp/api_key"

  if [[ ! -f "${key_file}" ]]; then
    die "API key file not found: ${key_file}"
  fi

  # Check permissions
  local perms
  perms="$(stat -c %a "${key_file}")"
  if [[ "${perms}" != "600" ]]; then
    die "API key file has insecure permissions: ${perms} (expected 600)"
  fi

  cat "${key_file}"
}

# Correct - Use secret management
function get_secret() {
  aws secretsmanager get-secret-value \
    --secret-id "$1" \
    --query 'SecretString' \
    --output text
}
```

```bash
# Incorrect - NEVER DO THIS
readonly API_KEY="sk-1234567890abcdef"
readonly DB_PASSWORD="supersecret123"
```

---

#### SH-SEC-005: Use Restricted PATH :yellow_circle:

**Rationale**: An attacker-controlled PATH can cause scripts to execute malicious versions of commands.

```bash
# Correct - Set explicit PATH at script start
#!/usr/bin/env bash
set -euo pipefail

# Restrict PATH to known-safe directories
export PATH="/usr/local/bin:/usr/bin:/bin"

# Or use full paths for critical commands
/usr/bin/rm -f "${temp_file}"
/bin/chmod 600 "${config_file}"
```

```bash
# Incorrect - Relying on inherited PATH
#!/usr/bin/env bash
rm -f "${temp_file}"
# Attacker could have put malicious 'rm' in PATH
```

---

## 11. Testing

### 11.1 Testing Framework

Use [Bats](https://github.com/bats-core/bats-core) (Bash Automated Testing System) for testing shell scripts.

### 11.2 Test File Structure

```
project/
  scripts/
    deploy.sh
    utils.sh
  tests/
    deploy_test.bats
    utils_test.bats
    test_helper.bash
    fixtures/
      sample_config.json
```

### 11.3 Testing Rules

#### SH-TST-001: Write Tests for Shell Scripts :yellow_circle:

**Rationale**: Tests catch regressions and document expected behavior. Shell scripts are prone to subtle bugs that tests can catch.

```bash
# tests/utils_test.bats

#!/usr/bin/env bats

setup() {
  # Load the script being tested
  source "${BATS_TEST_DIRNAME}/../scripts/utils.sh"

  # Create temp directory for test files
  TEST_TEMP_DIR="$(mktemp -d)"
}

teardown() {
  # Clean up
  rm -rf "${TEST_TEMP_DIR}"
}

@test "validate_email accepts valid email" {
  run validate_email "user@example.com"
  [ "$status" -eq 0 ]
}

@test "validate_email rejects invalid email" {
  run validate_email "not-an-email"
  [ "$status" -eq 1 ]
  [[ "$output" =~ "Invalid email" ]]
}

@test "process_file handles spaces in filename" {
  local test_file="${TEST_TEMP_DIR}/file with spaces.txt"
  echo "test content" > "${test_file}"

  run process_file "${test_file}"
  [ "$status" -eq 0 ]
}

@test "process_file fails on missing file" {
  run process_file "/nonexistent/file.txt"
  [ "$status" -eq 1 ]
}
```

---

#### SH-TST-002: Test Edge Cases :yellow_circle:

**Rationale**: Shell scripts are particularly vulnerable to edge cases involving special characters, empty values, and unusual input.

```bash
# Test cases to include:
@test "handles empty string input" {
  run process_input ""
  [ "$status" -eq 1 ]
}

@test "handles input with spaces" {
  run process_input "hello world"
  [ "$status" -eq 0 ]
}

@test "handles input with special characters" {
  run process_input 'test$var`cmd`$(injection)'
  [ "$status" -eq 0 ]
  # Verify no command execution occurred
}

@test "handles input with newlines" {
  run process_input $'line1\nline2'
  [ "$status" -eq 0 ]
}

@test "handles unicode input" {
  run process_input "Hello"
  [ "$status" -eq 0 ]
}

@test "handles extremely long input" {
  local long_input
  long_input="$(printf 'a%.0s' {1..10000})"
  run process_input "${long_input}"
  [ "$status" -eq 0 ]
}
```

---

## 12. Documentation

### 12.1 Comment Style

#### SH-DOC-001: Write Meaningful Comments :green_circle:

**Rationale**: Comments should explain why, not what. The code shows what happens; comments explain the reasoning.

```bash
# Correct - Explains why
# Use process substitution instead of pipe to preserve variable
# modifications in the loop body
while IFS= read -r line; do
  ((count++))
done < <(generate_lines)

# Correct - Explains non-obvious behavior
# Sleep before retry to avoid overwhelming the API
# and to allow transient issues to resolve
sleep $(( 2 ** retry_count ))

# Correct - Documents workaround
# HACK: AWS CLI v2 changed output format; strip quotes manually
# until we can update to v2.5+ which has --no-cli-auto-prompt
result="${result//\"/}"
```

```bash
# Incorrect - States the obvious
# Increment counter
((counter++))

# Loop through files
for file in "${files[@]}"; do
  process "${file}"
done
```

---

#### SH-DOC-002: Include File Header :yellow_circle:

**Rationale**: File headers provide context, usage information, and ownership.

```bash
#!/usr/bin/env bash
#
# deploy.sh - Deploy application to target environment
#
# This script handles the complete deployment workflow including:
# - Building artifacts
# - Running pre-deployment checks
# - Deploying to the specified environment
# - Running post-deployment verification
#
# Usage:
#   deploy.sh [-e|--environment <env>] [-v|--verbose] [-n|--dry-run]
#
# Options:
#   -e, --environment   Target environment (dev|staging|prod)
#   -v, --verbose       Enable verbose output
#   -n, --dry-run       Show what would be done without making changes
#   -h, --help          Show this help message
#
# Environment Variables:
#   DEPLOY_TOKEN        Authentication token (required)
#   DEPLOY_TIMEOUT      Deployment timeout in seconds (default: 300)
#
# Examples:
#   deploy.sh -e staging
#   deploy.sh --environment prod --verbose
#   DEPLOY_TIMEOUT=600 deploy.sh -e prod
#
# Exit Codes:
#   0 - Success
#   1 - General error
#   2 - Invalid arguments
#   3 - Deployment failed
#   4 - Verification failed
#
```

---

## 13. Portability Considerations

### 13.1 Bash-Specific Features

This standard targets Bash 4.0+. When portability to older systems is required:

| Feature | Bash 4.0+ | Portable Alternative |
|---------|-----------|---------------------|
| Associative arrays | `declare -A` | External file or multiple arrays |
| `mapfile` / `readarray` | Yes | `while read` loop |
| `${var,,}` lowercase | Yes | `tr '[:upper:]' '[:lower:]'` |
| `${var^^}` uppercase | Yes | `tr '[:lower:]' '[:upper:]'` |
| `|&` (pipe stderr) | Yes | `2>&1 |` |
| `&>` (redirect both) | Yes | `> file 2>&1` |

### 13.2 Portability Rules

#### SH-PRT-001: Document Bash Version Requirements :green_circle:

**Rationale**: Explicit version requirements prevent mysterious failures on older systems.

```bash
# Correct
#!/usr/bin/env bash
#
# Requires: Bash 4.0+ (for associative arrays)
#

# Check version at runtime
if [[ "${BASH_VERSION}" < "4.0" ]]; then
  echo "Error: This script requires Bash 4.0 or later" >&2
  exit 1
fi
```

---

## 14. Tooling Configuration

### 14.1 Required Tools

| Tool | Purpose | Version |
|------|---------|---------|
| [ShellCheck](https://github.com/koalaman/shellcheck) | Static analysis | 0.9.0+ |
| [shfmt](https://github.com/mvdan/sh) | Code formatting | 3.7.0+ |
| [Bats](https://github.com/bats-core/bats-core) | Testing framework | 1.10.0+ |

### 14.2 ShellCheck Configuration

Create `.shellcheckrc` in the project root:

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
# SC2154: Variable referenced but not assigned
# These are handled by our CI environment
disable=SC1090,SC1091

# Source path for sourced files
source-path=SCRIPTDIR
```

**Common ShellCheck Codes to Know:**

| Code | Description | Severity |
|------|-------------|----------|
| SC2086 | Double quote to prevent globbing | Warning |
| SC2046 | Quote to prevent word splitting | Warning |
| SC2006 | Use `$()` instead of backticks | Style |
| SC2034 | Variable appears unused | Warning |
| SC2155 | Declare and assign separately | Warning |
| SC2164 | Use `cd ... || exit` | Warning |
| SC2181 | Check exit code directly | Style |
| SC2230 | Use `command -v` instead of `which` | Warning |

### 14.3 shfmt Configuration

Create `.editorconfig` in the project root:

```ini
# .editorconfig

root = true

[*.sh]
# Use spaces for indentation
indent_style = space
indent_size = 2

# Shell variant
shell_variant = bash

# Binary operators may start a line
binary_next_line = true

# Indent switch cases
switch_case_indent = true

# Space after redirect operators
space_redirects = false

# Keep column alignment
keep_padding = false

# Function braces on same line
function_next_line = false

# End files with newline
insert_final_newline = true

# Trim trailing whitespace
trim_trailing_whitespace = true
```

**Command-line equivalent:**
```bash
shfmt -ln bash -i 2 -bn -ci -sr=false -w script.sh
```

### 14.4 IDE Setup

#### VS Code

Install extensions:
- [ShellCheck](https://marketplace.visualstudio.com/items?itemName=timonwong.shellcheck)
- [shell-format](https://marketplace.visualstudio.com/items?itemName=foxundermoon.shell-format)
- [Bash IDE](https://marketplace.visualstudio.com/items?itemName=mads-hartmann.bash-ide-vscode)

`.vscode/settings.json`:
```json
{
  "[shellscript]": {
    "editor.defaultFormatter": "foxundermoon.shell-format",
    "editor.formatOnSave": true,
    "editor.tabSize": 2
  },
  "shellcheck.enable": true,
  "shellcheck.run": "onSave",
  "shellcheck.executablePath": "shellcheck",
  "shellformat.path": "shfmt",
  "shellformat.flag": "-i 2 -bn -ci"
}
```

### 14.5 Pre-commit Hooks

```yaml
# .pre-commit-config.yaml

repos:
  - repo: https://github.com/koalaman/shellcheck-precommit
    rev: v0.9.0
    hooks:
      - id: shellcheck
        args: ["--severity=warning"]

  - repo: https://github.com/scop/pre-commit-shfmt
    rev: v3.7.0-1
    hooks:
      - id: shfmt
        args: ["-i", "2", "-bn", "-ci", "-w"]

  - repo: local
    hooks:
      - id: bats
        name: Run Bats tests
        entry: bats
        language: system
        files: _test\.bats$
        types: [file]
```

### 14.6 CI Integration

```yaml
# .github/workflows/shell-lint.yml

name: Shell Script Lint

on:
  push:
    paths:
      - '**/*.sh'
      - '.shellcheckrc'
  pull_request:
    paths:
      - '**/*.sh'
      - '.shellcheckrc'

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

  shfmt:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run shfmt
        uses: mvdan/sh-action@v1
        with:
          check: true

  bats:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Bats
        uses: bats-core/bats-action@1.5.0

      - name: Run tests
        run: bats tests/
```

---

## 15. Code Review Checklist

Quick reference for code reviewers:

### Structure and Organization
- [ ] Shebang is `#!/usr/bin/env bash` or `#!/bin/bash`
- [ ] `set -euo pipefail` is present immediately after shebang
- [ ] File follows standard organization (constants, functions, main)
- [ ] Script length is under 200 lines (consider refactoring if larger)
- [ ] No executable code between function definitions

### Naming
- [ ] Constants are `UPPER_SNAKE_CASE` and `readonly`
- [ ] Local variables are `lower_snake_case`
- [ ] Function names are descriptive and `lower_snake_case`
- [ ] No single-letter variables except loop counters in small scope

### Quoting and Variables
- [ ] ALL variable expansions are quoted: `"${variable}"`
- [ ] Arrays are used for lists (not space-separated strings)
- [ ] Local variables declared with `local` in functions
- [ ] No unquoted `$@` or `$*` (use `"$@"`)

### Error Handling
- [ ] `trap` is used for cleanup on EXIT/ERR
- [ ] Input validation at function boundaries
- [ ] Meaningful error messages written to stderr
- [ ] Appropriate exit codes used
- [ ] Required commands checked with `command -v`

### Control Flow
- [ ] `[[ ]]` used instead of `[ ]` for conditionals
- [ ] `$(command)` used instead of backticks
- [ ] `(( ))` used for arithmetic
- [ ] No `eval` with external input
- [ ] Process substitution used instead of pipes when variables modified

### Security
- [ ] No hardcoded secrets or credentials
- [ ] `mktemp` used for temporary files
- [ ] External input validated and sanitized
- [ ] PATH is explicit or restricted
- [ ] No SUID/SGID scripts

### Documentation
- [ ] File header with description and usage
- [ ] Functions have documentation comments
- [ ] Complex logic has explanatory comments
- [ ] Comments explain "why" not "what"

### Testing
- [ ] Tests exist for non-trivial scripts
- [ ] Edge cases tested (empty input, special characters, spaces)
- [ ] Error conditions tested

### Tooling
- [ ] ShellCheck passes with no warnings
- [ ] shfmt formatting applied
- [ ] No disabled ShellCheck rules without justification

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| SH-STR-001 | Use consistent shebang | :red_circle: |
| SH-STR-002 | Enable strict mode (set -euo pipefail) | :red_circle: |
| SH-STR-003 | Organize code in standard order | :yellow_circle: |
| SH-NAM-001 | Use descriptive variable names | :yellow_circle: |
| SH-NAM-002 | Use readonly for constants | :yellow_circle: |
| SH-VAR-001 | Always quote variable expansions | :red_circle: |
| SH-VAR-002 | Use braces for variable clarity | :yellow_circle: |
| SH-VAR-003 | Use arrays for lists | :yellow_circle: |
| SH-VAR-004 | Declare local variables in functions | :red_circle: |
| SH-ERR-001 | Implement cleanup with trap | :red_circle: |
| SH-ERR-002 | Use error output function | :yellow_circle: |
| SH-ERR-003 | Validate inputs at boundaries | :red_circle: |
| SH-ERR-004 | Check command availability | :yellow_circle: |
| SH-CTL-001 | Use [[ ]] for conditionals | :red_circle: |
| SH-CTL-002 | Use explicit comparison operators | :yellow_circle: |
| SH-CTL-003 | Use proper loop constructs | :yellow_circle: |
| SH-CMD-001 | Use $() for command substitution | :red_circle: |
| SH-CMD-002 | Handle command substitution failures | :yellow_circle: |
| SH-CMD-003 | Use (( )) for arithmetic | :yellow_circle: |
| SH-FUN-001 | Use function keyword with braces | :yellow_circle: |
| SH-FUN-002 | Document functions with comments | :yellow_circle: |
| SH-FUN-003 | Return meaningful exit codes | :yellow_circle: |
| SH-SEC-001 | Never use eval with external input | :red_circle: |
| SH-SEC-002 | Use mktemp for temporary files | :red_circle: |
| SH-SEC-003 | Validate and sanitize input | :red_circle: |
| SH-SEC-004 | Never store secrets in scripts | :red_circle: |
| SH-SEC-005 | Use restricted PATH | :yellow_circle: |
| SH-TST-001 | Write tests for shell scripts | :yellow_circle: |
| SH-TST-002 | Test edge cases | :yellow_circle: |
| SH-DOC-001 | Write meaningful comments | :green_circle: |
| SH-DOC-002 | Include file header | :yellow_circle: |
| SH-PRT-001 | Document Bash version requirements | :green_circle: |

---

## Appendix B: Glossary

| Term | Definition |
|------|------------|
| Bashism | A feature specific to Bash that is not POSIX-compliant |
| Exit code | Numeric value (0-255) returned by a command indicating success or failure |
| Glob | Pattern matching for filenames (e.g., `*.txt`) |
| Parameter expansion | Variable manipulation syntax (e.g., `${var:-default}`) |
| Shebang | The `#!` line at the start of a script specifying the interpreter |
| Strict mode | `set -euo pipefail` settings for safer script execution |
| Subshell | A child shell process that cannot modify parent variables |
| Word splitting | Shell behavior of splitting unquoted variables on whitespace |

---

## Appendix C: Common Patterns

### Argument Parsing

```bash
function parse_arguments() {
  while [[ $# -gt 0 ]]; do
    case "$1" in
      -h|--help)
        show_help
        exit 0
        ;;
      -v|--verbose)
        VERBOSE=true
        shift
        ;;
      -e|--environment)
        ENVIRONMENT="$2"
        shift 2
        ;;
      --)
        shift
        break
        ;;
      -*)
        die "Unknown option: $1"
        ;;
      *)
        POSITIONAL_ARGS+=("$1")
        shift
        ;;
    esac
  done
}
```

### Retry Logic

```bash
function retry() {
  local max_attempts="$1"
  local delay="$2"
  shift 2
  local cmd=("$@")

  local attempt=1
  while true; do
    if "${cmd[@]}"; then
      return 0
    fi

    if (( attempt >= max_attempts )); then
      err "Command failed after ${max_attempts} attempts: ${cmd[*]}"
      return 1
    fi

    warn "Attempt ${attempt} failed, retrying in ${delay}s..."
    sleep "${delay}"
    ((attempt++))
  done
}

# Usage
retry 3 5 curl -f https://api.example.com/health
```

### Locking

```bash
readonly LOCK_FILE="/var/lock/${SCRIPT_NAME}.lock"

function acquire_lock() {
  if ! mkdir "${LOCK_FILE}" 2>/dev/null; then
    local pid
    pid="$(cat "${LOCK_FILE}/pid" 2>/dev/null)" || true
    die "Another instance is running (PID: ${pid:-unknown})"
  fi

  echo $$ > "${LOCK_FILE}/pid"
  trap 'rm -rf "${LOCK_FILE}"' EXIT
}
```

### Progress Indicator

```bash
function progress() {
  local current="$1"
  local total="$2"
  local width=50

  local percent=$(( current * 100 / total ))
  local filled=$(( current * width / total ))
  local empty=$(( width - filled ))

  printf "\r[%s%s] %3d%%" \
    "$(printf '#%.0s' $(seq 1 "${filled}"))" \
    "$(printf ' %.0s' $(seq 1 "${empty}"))" \
    "${percent}"

  if (( current == total )); then
    echo
  fi
}
```

---

## Appendix D: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix E: References

- [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html)
- [ShellCheck Wiki](https://github.com/koalaman/shellcheck/wiki)
- [shfmt - Shell Formatter](https://github.com/mvdan/sh)
- [Bats - Bash Automated Testing System](https://github.com/bats-core/bats-core)
- [Bash Reference Manual](https://www.gnu.org/software/bash/manual/bash.html)
- [Safer Bash Scripts with set -euxo pipefail](https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail/)
- [Bash Pitfalls](https://mywiki.wooledge.org/BashPitfalls)
- [Pure Bash Bible](https://github.com/dylanaraps/pure-bash-bible)
