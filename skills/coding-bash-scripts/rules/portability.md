# Portability Rules (SH-PRT-*)

Portability considerations ensure scripts work across different systems and Bash versions.

## Portability Strategy

- Document Bash version requirements
- Use portable constructs when possible
- Test on target systems
- Understand Bash-specific features

---

## SH-PRT-001: Document Bash Version Requirements :green_circle:

**Tier**: Recommended

**Rationale**: Explicit version requirements prevent mysterious failures on older systems.

```bash
#!/usr/bin/env bash
#
# Requires: Bash 4.0+ (for associative arrays)
#

set -euo pipefail

# Check version at runtime
if [[ "${BASH_VERSION}" < "4.0" ]]; then
  echo "Error: This script requires Bash 4.0 or later" >&2
  echo "Current version: ${BASH_VERSION}" >&2
  exit 1
fi

# Or check for specific feature
if ! declare -A test_array 2>/dev/null; then
  echo "Error: This script requires associative array support (Bash 4.0+)" >&2
  exit 1
fi
```

---

## Bash Version Features

### Bash 4.0+ Features

| Feature | Example | Portable Alternative |
|---------|---------|---------------------|
| Associative arrays | `declare -A map` | External file or multiple arrays |
| `mapfile` / `readarray` | `mapfile -t arr < file` | `while read` loop |
| `${var,,}` lowercase | `${name,,}` | `tr '[:upper:]' '[:lower:]'` |
| `${var^^}` uppercase | `${name^^}` | `tr '[:lower:]' '[:upper:]'` |
| `coproc` | `coproc cmd` | Named pipes |

### Bash 4.2+ Features

| Feature | Example |
|---------|---------|
| Negative array indices | `${array[-1]}` (last element) |
| `test -v` variable set | `[[ -v variable ]]` |
| `lastpipe` option | Execute last pipe segment in current shell |

### Bash 4.3+ Features

| Feature | Example |
|---------|---------|
| Namerefs | `local -n ref="$1"` |
| Negative substring | `${var: -3}` |
| `inherit_errexit` | Subshells inherit errexit |

### Bash 4.4+ Features

| Feature | Example |
|---------|---------|
| `${var@Q}` quoted | Safe quoting for eval |
| `${var@a}` attributes | Get variable attributes |
| `mapfile -d` | Custom delimiter |

### Bash 5.0+ Features

| Feature | Example |
|---------|---------|
| `EPOCHSECONDS` | Seconds since epoch |
| `EPOCHREALTIME` | Microsecond precision |
| `BASH_ARGV0` | Set `$0` |

---

## Portable Alternatives

### Associative Arrays (Bash 4.0+)

```bash
# Bash 4.0+ - Associative array
declare -A config
config["host"]="localhost"
config["port"]="8080"

# Portable alternative - Functions
get_config() {
  local key="$1"
  case "${key}" in
    host) echo "localhost" ;;
    port) echo "8080" ;;
  esac
}
host="$(get_config host)"
```

### Case Conversion

```bash
# Bash 4.0+ - Parameter expansion
lowercase="${string,,}"
uppercase="${string^^}"

# Portable - Using tr
lowercase="$(echo "${string}" | tr '[:upper:]' '[:lower:]')"
uppercase="$(echo "${string}" | tr '[:lower:]' '[:upper:]')"
```

### mapfile/readarray

```bash
# Bash 4.0+ - mapfile
mapfile -t lines < "${file}"

# Portable - while loop
lines=()
while IFS= read -r line; do
  lines+=("${line}")
done < "${file}"
```

### Last Array Element

```bash
# Bash 4.2+ - Negative index
last="${array[-1]}"

# Portable
last="${array[${#array[@]}-1]}"
```

### Namerefs

```bash
# Bash 4.3+ - Nameref
function get_result() {
  local -n result_ref="$1"
  result_ref="value"
}

# Portable - Eval (use carefully!)
function get_result() {
  local var_name="$1"
  eval "${var_name}='value'"
}

# Portable - stdout
function get_result() {
  echo "value"
}
result="$(get_result)"
```

---

## POSIX Compatibility

For maximum portability, avoid these Bashisms:

| Bashism | POSIX Alternative |
|---------|------------------|
| `[[ ]]` | `[ ]` with proper quoting |
| `(( ))` | `expr` or `$(( ))` |
| `$'...'` quotes | `printf` |
| `source` | `.` (dot) |
| `function name()` | `name()` |
| `local` | (not POSIX, but widely supported) |
| `declare` | (not POSIX) |
| `+=` array append | `arr[${#arr[@]}]=value` |
| `<<<` here-string | `echo | command` |
| `&>` redirect | `> file 2>&1` |

---

## Platform Considerations

### macOS vs Linux

```bash
# Date command differences
# Linux
date -d "2024-01-15" +%s

# macOS
date -j -f "%Y-%m-%d" "2024-01-15" +%s

# Portable
if [[ "$(uname)" == "Darwin" ]]; then
  timestamp=$(date -j -f "%Y-%m-%d" "${date_str}" +%s)
else
  timestamp=$(date -d "${date_str}" +%s)
fi
```

### stat Command

```bash
# Linux
stat -c %a "${file}"

# macOS
stat -f %Lp "${file}"

# Portable
if [[ "$(uname)" == "Darwin" ]]; then
  perms=$(stat -f %Lp "${file}")
else
  perms=$(stat -c %a "${file}")
fi
```

### sed Extended Regex

```bash
# Linux
sed -E 's/pattern/replacement/'

# macOS (older)
sed -E 's/pattern/replacement/'

# Most portable
sed 's/pattern/replacement/'  # Basic regex
```

### readlink

```bash
# Linux - readlink -f
readlink -f "${path}"

# macOS - requires coreutils or:
function resolve_path() {
  local path="$1"
  if [[ -L "${path}" ]]; then
    path="$(readlink "${path}")"
  fi
  cd "$(dirname "${path}")" && pwd -P
}
```

---

## Testing Portability

```bash
# Test on multiple Bash versions using Docker
docker run --rm -v "$(pwd):/scripts" bash:4.0 bash /scripts/test.sh
docker run --rm -v "$(pwd):/scripts" bash:4.4 bash /scripts/test.sh
docker run --rm -v "$(pwd):/scripts" bash:5.0 bash /scripts/test.sh

# Test POSIX compatibility
docker run --rm -v "$(pwd):/scripts" busybox sh /scripts/test.sh
```
