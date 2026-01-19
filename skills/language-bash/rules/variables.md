# Variable Rules (SH-VAR-*)

Proper variable handling is essential for robust, secure shell scripts. Quoting prevents word splitting and glob expansion bugs.

## Variable Strategy

- Quote ALL variable expansions
- Use braces for clarity and parameter expansion
- Use arrays for lists of items
- Declare local variables in functions

---

## SH-VAR-001: Always Quote Variable Expansions :red_circle:

**Tier**: Critical

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

# Correct - all arguments
process_args "$@"
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

## SH-VAR-002: Use Braces for Variable Clarity :yellow_circle:

**Tier**: Required

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

# Error if unset
local required="${REQUIRED_VAR:?Variable required}"
```

```bash
# Incorrect
echo "$variable"suffix    # Ambiguous - is it $variable or $variablesuffix?
echo "$user"_home         # Works but inconsistent
```

### Parameter Expansion Reference

| Syntax | Description |
|--------|-------------|
| `${var:-default}` | Use default if unset or empty |
| `${var:=default}` | Assign default if unset or empty |
| `${var:?error}` | Exit with error if unset or empty |
| `${var:+replacement}` | Use replacement if set and non-empty |
| `${#var}` | Length of variable |
| `${var%pattern}` | Remove shortest suffix match |
| `${var%%pattern}` | Remove longest suffix match |
| `${var#pattern}` | Remove shortest prefix match |
| `${var##pattern}` | Remove longest prefix match |
| `${var/pattern/replacement}` | Replace first match |
| `${var//pattern/replacement}` | Replace all matches |

---

## SH-VAR-003: Use Arrays for Lists :yellow_circle:

**Tier**: Required

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

# Array length
echo "Found ${#files[@]} files"

# Check if array is empty
if [[ ${#files[@]} -eq 0 ]]; then
  echo "No files found"
fi
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

### Array Operations

```bash
# Declare array
local -a items=("one" "two" "three")
declare -a more_items

# Append to array
items+=("four")

# Access elements
echo "${items[0]}"    # First element
echo "${items[-1]}"   # Last element (bash 4.2+)

# All elements
echo "${items[@]}"    # Each as separate word
echo "${items[*]}"    # All as single word

# Slice
echo "${items[@]:1:2}"  # Elements 1 and 2

# Array indices
echo "${!items[@]}"   # All indices
```

---

## SH-VAR-004: Declare Local Variables in Functions :red_circle:

**Tier**: Critical

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

### Local Variable Modifiers

```bash
function example() {
  local name="value"       # Regular local
  local -r constant="val"  # Local readonly
  local -i number=42       # Integer type
  local -a array=()        # Indexed array
  local -A map=()          # Associative array (bash 4+)
  local -l lowercase       # Lowercase value
  local -u uppercase       # Uppercase value
}
```

---

## Additional Patterns

### Safe Read from File

```bash
# Read file line by line preserving whitespace
while IFS= read -r line || [[ -n "${line}" ]]; do
  process "${line}"
done < "${input_file}"

# The || [[ -n "${line}" ]] handles files without trailing newline
```

### Null-Separated Input (for filenames)

```bash
# Safe handling of filenames with special characters
while IFS= read -r -d '' file; do
  process "${file}"
done < <(find . -type f -print0)
```

### Associative Arrays (Bash 4+)

```bash
declare -A config

# Populate
config["host"]="localhost"
config["port"]="8080"
config["debug"]="true"

# Access
echo "Host: ${config[host]}"

# Check if key exists
if [[ -v config[host] ]]; then
  echo "Host is set"
fi

# Iterate
for key in "${!config[@]}"; do
  echo "${key}=${config[${key}]}"
done
```
