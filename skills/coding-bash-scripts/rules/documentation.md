# Documentation Rules (SH-DOC-*)

Documentation in shell scripts explains intent, usage, and non-obvious behavior.

## Documentation Strategy

- Include file headers with usage information
- Document functions with comments
- Explain "why" not "what"
- Keep comments up to date

---

## SH-DOC-001: Write Meaningful Comments :green_circle:

**Tier**: Recommended

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

# Correct - Explains business logic
# Users with admin role can bypass rate limiting
# per security team decision (JIRA-1234)
if [[ "${user_role}" == "admin" ]]; then
  rate_limit=0
fi
```

```bash
# Incorrect - States the obvious
# Increment counter
((counter++))

# Loop through files
for file in "${files[@]}"; do
  process "${file}"
done

# Check if file exists
if [[ -f "${file}" ]]; then
  # ...
fi
```

---

## SH-DOC-002: Include File Header :yellow_circle:

**Tier**: Required

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
# Dependencies:
#   - curl
#   - jq
#   - aws-cli (v2)
#
# Author: Team Name
# Created: 2024-01-15
#

set -euo pipefail
```

---

## Function Documentation Format

```bash
#######################################
# Brief description of what the function does.
#
# Extended description if needed, explaining
# the purpose and any important details.
#
# Globals:
#   GLOBAL_VAR - Description of how it's used (read/write)
#   CONFIG_PATH - Path to configuration (read)
#
# Arguments:
#   $1 - Description of first argument (required)
#   $2 - Description of second argument (optional, default: value)
#
# Outputs:
#   Writes result to stdout
#   Writes errors to stderr
#
# Returns:
#   0 - Success
#   1 - File not found
#   2 - Permission denied
#
# Example:
#   result=$(process_data "input.txt" "output.txt")
#######################################
function process_data() {
  local input_file="$1"
  local output_file="${2:-/dev/stdout}"
  # Implementation
}
```

---

## Additional Documentation Patterns

### TODO Comments

```bash
# TODO(username): Description of what needs to be done
# Include ticket reference if available (JIRA-1234)

# TODO(johndoe): Refactor to use new API endpoint (JIRA-5678)
# The current endpoint is deprecated and will be removed Q2 2024
```

### FIXME Comments

```bash
# FIXME: Brief description of the bug
# More details about the issue and potential fix

# FIXME: Race condition when multiple instances run
# Need to implement proper file locking (see SH-SEC-002)
```

### Deprecation Notices

```bash
#######################################
# DEPRECATED: Use new_function() instead
#
# This function will be removed in v2.0.
# Migration guide: https://docs.example.com/migration
#######################################
function old_function() {
  warn "old_function is deprecated, use new_function instead"
  new_function "$@"
}
```

### Section Headers

```bash
###############################################################################
# Configuration
###############################################################################

readonly CONFIG_FILE="/etc/myapp/config"
readonly LOG_DIR="/var/log/myapp"

###############################################################################
# Helper Functions
###############################################################################

function log() {
  # ...
}

###############################################################################
# Main Logic
###############################################################################

function main() {
  # ...
}
```

### Inline Documentation for Complex Logic

```bash
function calculate_backoff() {
  local attempt="$1"
  local base_delay="${2:-1}"
  local max_delay="${3:-300}"

  # Exponential backoff with jitter
  # Formula: min(max_delay, base_delay * 2^attempt + random(0, 1000ms))
  # This prevents thundering herd when multiple clients retry simultaneously

  local exponential_delay=$(( base_delay * (2 ** attempt) ))
  local jitter=$(( RANDOM % 1000 ))

  # Cap at maximum delay to prevent excessive wait times
  local delay=$(( exponential_delay + jitter / 1000 ))
  if (( delay > max_delay )); then
    delay="${max_delay}"
  fi

  echo "${delay}"
}
```

### Help Function

```bash
function show_help() {
  cat << EOF
Usage: ${SCRIPT_NAME} [OPTIONS] <command>

Deploy application to target environment.

Commands:
  deploy      Deploy the application
  rollback    Rollback to previous version
  status      Show deployment status

Options:
  -e, --environment ENV   Target environment (dev|staging|prod)
  -v, --verbose           Enable verbose output
  -n, --dry-run           Show what would be done
  -h, --help              Show this help message

Environment Variables:
  DEPLOY_TOKEN      Authentication token (required)
  DEPLOY_TIMEOUT    Timeout in seconds (default: 300)

Examples:
  ${SCRIPT_NAME} deploy -e staging
  ${SCRIPT_NAME} rollback -e prod --verbose
  ${SCRIPT_NAME} status -e dev

For more information, see: https://docs.example.com/deploy
EOF
}
```

### Version Information

```bash
function show_version() {
  cat << EOF
${SCRIPT_NAME} ${VERSION}
Copyright (c) 2024 Company Name
License: MIT
EOF
}
```
