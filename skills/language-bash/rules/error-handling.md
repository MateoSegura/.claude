# Error Handling Rules (SH-ERR-*)

Robust error handling prevents cascading failures and aids debugging. Scripts should fail safely and predictably.

## Error Handling Strategy

1. Enable strict mode (`set -euo pipefail`)
2. Use `trap` for cleanup on exit
3. Validate inputs at boundaries
4. Provide meaningful error messages
5. Use appropriate exit codes

---

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Misuse of command/invalid arguments |
| 126 | Command not executable |
| 127 | Command not found |
| 128+N | Fatal error signal N |

---

## SH-ERR-001: Implement Cleanup with trap :red_circle:

**Tier**: Critical

**Rationale**: Scripts that create temporary files, acquire locks, or modify system state must clean up on exit, including when terminated by signals.

```bash
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
TEMP_FILE=$(mktemp)
# If script fails, temp file is never removed

echo "data" > "${TEMP_FILE}"
process "${TEMP_FILE}"
rm "${TEMP_FILE}"  # Never reached on error!
```

---

## SH-ERR-002: Use Error Output Function :yellow_circle:

**Tier**: Required

**Rationale**: Consistent error output format aids debugging and log analysis. Errors should go to stderr, not stdout.

```bash
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

## SH-ERR-003: Validate Inputs at Boundaries :red_circle:

**Tier**: Critical

**Rationale**: Early validation prevents cascading failures and provides clear error messages. Never trust external input.

```bash
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

## SH-ERR-004: Check Command Availability :yellow_circle:

**Tier**: Required

**Rationale**: Scripts should fail fast with helpful messages if required commands are missing.

```bash
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

## Additional Patterns

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

### Error Context

```bash
function with_context() {
  local context="$1"
  shift

  if ! "$@"; then
    err "${context}: command failed"
    return 1
  fi
}

# Usage
with_context "Downloading artifact" curl -fsSL "${url}" -o "${output}"
```

### Capturing Exit Codes

```bash
# Capture exit code without triggering errexit
local exit_code=0
some_command || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
  err "Command failed with exit code: ${exit_code}"
fi
```

### Assert Functions

```bash
function assert_file_exists() {
  local file="$1"
  local message="${2:-File does not exist: ${file}}"

  if [[ ! -f "${file}" ]]; then
    die "${message}"
  fi
}

function assert_not_empty() {
  local var="$1"
  local name="${2:-Variable}"

  if [[ -z "${var}" ]]; then
    die "${name} cannot be empty"
  fi
}

# Usage
assert_file_exists "${config_file}" "Config file missing"
assert_not_empty "${user_id}" "User ID"
```
