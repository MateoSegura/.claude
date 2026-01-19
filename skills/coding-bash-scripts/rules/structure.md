# File Structure Rules (SH-STR-*)

Consistent file structure makes scripts easier to navigate, understand, and maintain.

## Structure Strategy

- Use consistent shebang (`#!/usr/bin/env bash`)
- Enable strict mode immediately after shebang
- Organize code in standard order
- Keep scripts under 200 lines

---

## SH-STR-001: Use Consistent Shebang :red_circle:

**Tier**: Critical

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

## SH-STR-002: Enable Strict Mode :red_circle:

**Tier**: Critical

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

## SH-STR-003: Organize Code in Standard Order :yellow_circle:

**Tier**: Required

**Rationale**: Consistent organization makes scripts easier to navigate and understand.

```bash
#!/usr/bin/env bash
#
# Script description: One-line summary.
#
# Usage: script_name [options] <arguments>
#

# 1. Strict mode settings
set -euo pipefail

# 2. Constants (readonly, uppercase)
readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_NAME="$(basename "${BASH_SOURCE[0]}")"
readonly VERSION="1.0.0"

# 3. Global variables (uppercase)
VERBOSE=false
DRY_RUN=false

# 4. Source dependencies
source "${SCRIPT_DIR}/lib/utils.sh"

# 5. Function definitions
function example_function() {
  local arg1="$1"
  # Function body
}

# 6. Main function
function main() {
  parse_arguments "$@"
  # Main logic here
}

# 7. Script entry point - must be at bottom
main "$@"
```

```bash
# Incorrect - executable code between functions
function setup() { ... }

echo "Starting..."  # Don't put code here!

function cleanup() { ... }
```

---

## Standard File Layout Template

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

set -euo pipefail

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_NAME="$(basename "${BASH_SOURCE[0]}")"

VERBOSE=false

#######################################
# Function description.
# Globals:
#   VERBOSE
# Arguments:
#   $1 - Description
# Outputs:
#   Writes to stdout
# Returns:
#   0 on success, non-zero on error
#######################################
function example_function() {
  local arg1="$1"
  # Implementation
}

function main() {
  # Main logic
  :
}

main "$@"
```

---

## File Extensions

| Type | Extension | Example |
|------|-----------|---------|
| Executable scripts (in PATH) | No extension | `deploy` |
| Library/source files | `.sh` | `utils.sh` |
| Test files | `_test.sh` or `.bats` | `utils_test.sh` |
