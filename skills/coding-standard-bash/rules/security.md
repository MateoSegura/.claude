# Security Rules (SH-SEC-*)

Security in shell scripts requires careful handling of external input, secrets, and system resources.

## Security Principles

1. **Never trust external input** - Validate and sanitize all input
2. **Quote everything** - Prevent injection attacks
3. **Use safe temporary files** - Prevent race conditions
4. **Avoid eval** - Prevents code injection
5. **Principle of least privilege** - Drop permissions when possible

---

## SH-SEC-001: Never Use eval with External Input :red_circle:

**Tier**: Critical

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

# Correct - Use indirect expansion for dynamic variable names
local var_name="PATH"
echo "${!var_name}"

# Correct - Use declare for dynamic assignment
local var_name="my_var"
declare "${var_name}=value"
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

## SH-SEC-002: Use mktemp for Temporary Files :red_circle:

**Tier**: Critical

**Rationale**: Manual temp file creation is vulnerable to race conditions and symlink attacks. `mktemp` creates files atomically with secure permissions.

```bash
# Correct - Temporary file
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

---

## SH-SEC-003: Validate and Sanitize Input :red_circle:

**Tier**: Critical

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

# Correct - Sanitize for display/logging
function sanitize_for_log() {
  local input="$1"
  # Remove control characters
  printf '%s' "${input}" | tr -d '[:cntrl:]'
}

# Correct - Validate numeric input
function validate_port() {
  local port="$1"

  if [[ ! "${port}" =~ ^[0-9]+$ ]]; then
    die "Port must be numeric: ${port}"
  fi

  if (( port < 1 || port > 65535 )); then
    die "Port must be 1-65535: ${port}"
  fi
}
```

```bash
# Incorrect - Using input directly
local filename="$1"
cat "${filename}"
# What if filename is "/etc/shadow" or "../../etc/passwd"?
```

---

## SH-SEC-004: Never Store Secrets in Scripts :red_circle:

**Tier**: Critical

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

## SH-SEC-005: Use Restricted PATH :yellow_circle:

**Tier**: Required

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

## Additional Security Patterns

### Secure File Operations

```bash
# Create file with secure permissions
function create_secure_file() {
  local file="$1"

  # Set umask before creating
  local old_umask
  old_umask="$(umask)"
  umask 077

  touch "${file}"

  umask "${old_umask}"
}

# Or explicitly set permissions
touch "${file}"
chmod 600 "${file}"
```

### Locking to Prevent Race Conditions

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

### Escaping for SQL/Commands

```bash
# If you must construct commands (try to avoid)
function escape_for_shell() {
  local input="$1"
  printf '%q' "${input}"
}

# Better: Use arrays to avoid shell interpretation
local -a cmd=("grep" "-r" "${pattern}" "${directory}")
"${cmd[@]}"
```

### Dropping Privileges

```bash
# Run command as different user
function run_as_user() {
  local user="$1"
  shift

  sudo -u "${user}" "$@"
}

# Check we're not running as root
function require_non_root() {
  if [[ "${EUID}" -eq 0 ]]; then
    die "This script should not be run as root"
  fi
}
```

### Timeout for Network Operations

```bash
# Prevent hanging on network issues
function fetch_with_timeout() {
  local url="$1"
  local timeout="${2:-30}"

  curl --max-time "${timeout}" --fail --silent "${url}"
}
```
