# Function Rules (SH-FUN-*)

Well-designed functions improve code organization, reusability, and testability.

## Function Strategy

- Use `function` keyword with braces
- Document functions with comments
- Return meaningful exit codes
- Declare local variables

---

## SH-FUN-001: Use function Keyword with Braces :yellow_circle:

**Tier**: Required

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

## SH-FUN-002: Document Functions with Comments :yellow_circle:

**Tier**: Required

**Rationale**: Function documentation enables understanding without reading implementation details.

```bash
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

  if [[ ! -f "${log_file}" ]]; then
    return 1
  fi

  if [[ ! -r "${log_file}" ]]; then
    return 2
  fi

  grep -E "${ERROR_PATTERN}" "${log_file}" > "${output_file}"
}
```

```bash
# Incorrect - no documentation
function extract_errors() {
  grep -E "$ERROR_PATTERN" "$1" > "${2:-/dev/stdout}"
}
```

---

## SH-FUN-003: Return Meaningful Exit Codes :yellow_circle:

**Tier**: Required

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

## Additional Patterns

### Returning Data from Functions

```bash
# Method 1: Output to stdout (preferred)
function get_user_name() {
  local user_id="$1"
  # Query and output
  echo "${name}"
}

# Usage
local name
name="$(get_user_name "${id}")"

# Method 2: Output variable (for complex data)
function get_user_info() {
  local user_id="$1"
  local -n result_ref="$2"  # nameref (bash 4.3+)

  result_ref[name]="John"
  result_ref[email]="john@example.com"
}

# Usage
declare -A user
get_user_info "123" user
echo "${user[name]}"

# Method 3: Global variable (last resort)
function get_user_data() {
  USER_DATA_NAME="John"
  USER_DATA_EMAIL="john@example.com"
}
```

### Default Arguments

```bash
function process() {
  local input="${1:?Input file required}"
  local output="${2:-output.txt}"
  local verbose="${3:-false}"

  # Implementation
}
```

### Variable Number of Arguments

```bash
function log_all() {
  local level="$1"
  shift  # Remove first argument

  for message in "$@"; do
    echo "[${level}] ${message}"
  done
}

# Usage
log_all "INFO" "Starting process" "Loading config" "Ready"
```

### Function with Options

```bash
function deploy() {
  local environment=""
  local dry_run=false
  local verbose=false

  while [[ $# -gt 0 ]]; do
    case "$1" in
      -e|--environment)
        environment="$2"
        shift 2
        ;;
      -n|--dry-run)
        dry_run=true
        shift
        ;;
      -v|--verbose)
        verbose=true
        shift
        ;;
      *)
        err "Unknown option: $1"
        return 1
        ;;
    esac
  done

  if [[ -z "${environment}" ]]; then
    err "Environment is required"
    return 1
  fi

  # Implementation using local variables
}

# Usage
deploy --environment prod --verbose
```

### Private Helper Functions

```bash
# Prefix with underscore for internal functions
function _validate_internal() {
  local value="$1"
  [[ "${value}" =~ ^[a-z]+$ ]]
}

function _format_output() {
  local data="$1"
  printf "%s\n" "${data}"
}

# Public function
function process_data() {
  local input="$1"

  if ! _validate_internal "${input}"; then
    return 1
  fi

  _format_output "${input}"
}
```

### Callback Pattern

```bash
function for_each_file() {
  local directory="$1"
  local callback="$2"

  local file
  for file in "${directory}"/*; do
    if [[ -f "${file}" ]]; then
      "${callback}" "${file}"
    fi
  done
}

function process_file() {
  local file="$1"
  echo "Processing: ${file}"
}

# Usage
for_each_file "/path/to/dir" process_file
```
