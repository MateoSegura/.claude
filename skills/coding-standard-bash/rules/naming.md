# Naming Rules (SH-NAM-*)

Consistent naming conventions make code self-documenting and reduce cognitive load.

## Naming Strategy

- Use lowercase with underscores for most identifiers
- Use UPPERCASE for constants and environment variables
- Names should be descriptive and reveal intent
- Avoid abbreviations except for well-known terms

---

## Naming Conventions Summary

| Element | Convention | Example |
|---------|------------|---------|
| Constants | `UPPER_SNAKE_CASE` | `MAX_RETRIES`, `DEFAULT_TIMEOUT` |
| Environment variables | `UPPER_SNAKE_CASE` | `LOG_LEVEL`, `CONFIG_PATH` |
| Local variables | `lower_snake_case` | `file_count`, `user_input` |
| Functions | `lower_snake_case` | `process_file`, `validate_input` |
| Script files | `lower-kebab-case` or `lower_snake_case` | `deploy-app`, `run_tests.sh` |
| Private functions | `_leading_underscore` | `_internal_helper` |

---

## SH-NAM-001: Use Descriptive Variable Names :yellow_circle:

**Tier**: Required

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

```bash
# Acceptable - short scope, obvious purpose
for ((i = 0; i < count; i++)); do
  echo "Item $i"
done
```

---

## SH-NAM-002: Use readonly for Constants :yellow_circle:

**Tier**: Required

**Rationale**: Declaring constants as `readonly` prevents accidental modification and signals intent to readers.

```bash
# Correct
readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly CONFIG_FILE="/etc/myapp/config.conf"
readonly -a SUPPORTED_FORMATS=("json" "yaml" "toml")
readonly MAX_RETRIES=3
readonly DEFAULT_TIMEOUT=30
```

```bash
# Incorrect
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONFIG_FILE="/etc/myapp/config.conf"
# ^ Can be accidentally modified later
```

---

## Additional Patterns

### Boolean Variables

```bash
# Use true/false for boolean flags
VERBOSE=false
DRY_RUN=false

if [[ "${VERBOSE}" == "true" ]]; then
  echo "Verbose mode enabled"
fi

# Or use empty/non-empty
DEBUG=""

if [[ -n "${DEBUG}" ]]; then
  set -x
fi
```

### Environment Variable Defaults

```bash
# Document expected environment variables at script start
readonly LOG_LEVEL="${LOG_LEVEL:-info}"
readonly CONFIG_PATH="${CONFIG_PATH:-/etc/myapp/config}"
readonly TIMEOUT="${TIMEOUT:-30}"

# Required environment variables
readonly API_KEY="${API_KEY:?API_KEY environment variable required}"
```

### Prefixing Related Variables

```bash
# Group related variables with common prefix
readonly DB_HOST="${DB_HOST:-localhost}"
readonly DB_PORT="${DB_PORT:-5432}"
readonly DB_USER="${DB_USER:-postgres}"
readonly DB_NAME="${DB_NAME:-myapp}"

readonly HTTP_TIMEOUT="${HTTP_TIMEOUT:-30}"
readonly HTTP_RETRIES="${HTTP_RETRIES:-3}"
```

### Private/Internal Functions

```bash
# Use underscore prefix for internal helpers
function _log() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] $*"
}

function _validate_internal() {
  # Not meant to be called directly
  :
}

# Public function
function process_data() {
  _log "Processing started"
  _validate_internal
  # ...
}
```
