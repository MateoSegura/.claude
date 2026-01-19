# Control Flow Rules (SH-CTL-*)

Proper control flow structures make scripts readable and prevent subtle bugs.

## Control Flow Strategy

- Use `[[ ]]` for all conditionals
- Use explicit comparison operators
- Choose the right loop for the task
- Avoid subshells when variables need modification

---

## SH-CTL-001: Use [[ ]] for Conditionals :red_circle:

**Tier**: Critical

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

## SH-CTL-002: Use Explicit Comparison Operators :yellow_circle:

**Tier**: Required

**Rationale**: Explicit operators make code readable and prevent subtle bugs.

### String Comparisons

```bash
[[ "${str}" == "value" ]]   # Equal
[[ "${str}" != "value" ]]   # Not equal
[[ "${str}" < "value" ]]    # Less than (lexicographic)
[[ "${str}" > "value" ]]    # Greater than (lexicographic)
[[ -z "${str}" ]]           # Is empty
[[ -n "${str}" ]]           # Is not empty
```

### Numeric Comparisons

```bash
[[ "${num}" -eq 10 ]]       # Equal
[[ "${num}" -ne 10 ]]       # Not equal
[[ "${num}" -lt 10 ]]       # Less than
[[ "${num}" -le 10 ]]       # Less than or equal
[[ "${num}" -gt 10 ]]       # Greater than
[[ "${num}" -ge 10 ]]       # Greater than or equal

# Arithmetic context for numbers (preferred)
if (( num > 10 )); then
  echo "Greater"
fi
```

### File Tests

```bash
[[ -e "${path}" ]]    # Exists
[[ -f "${path}" ]]    # Is regular file
[[ -d "${path}" ]]    # Is directory
[[ -r "${path}" ]]    # Is readable
[[ -w "${path}" ]]    # Is writable
[[ -x "${path}" ]]    # Is executable
[[ -s "${path}" ]]    # Has size > 0
[[ -L "${path}" ]]    # Is symlink
```

```bash
# Incorrect
[[ "${str}" ]]              # Unclear - testing for non-empty?
[[ ! "${str}" ]]            # Confusing negation
[[ "${num}" > 10 ]]         # Wrong! This is string comparison
```

---

## SH-CTL-003: Use Proper Loop Constructs :yellow_circle:

**Tier**: Required

**Rationale**: Choose the right loop for the task. Avoid subshells in pipelines when variable modification is needed.

### Iterate Over Array

```bash
# Correct
for item in "${array[@]}"; do
  process "${item}"
done
```

### C-Style For Loop

```bash
# Correct
for ((i = 0; i < count; i++)); do
  process "${i}"
done
```

### Read Lines from File

```bash
# Correct - preserves whitespace
while IFS= read -r line; do
  process "${line}"
done < "${input_file}"

# Handle file without trailing newline
while IFS= read -r line || [[ -n "${line}" ]]; do
  process "${line}"
done < "${input_file}"
```

### Read with Process Substitution

```bash
# Correct - preserves variables (no subshell)
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
```

### Read Null-Separated Data

```bash
# Correct - handles filenames with spaces/newlines
while IFS= read -r -d '' file; do
  process "${file}"
done < <(find . -name "*.txt" -print0)
```

```bash
# Incorrect - Word splitting on find output
for file in $(find . -name "*.txt"); do
  # Breaks on filenames with spaces
  process "${file}"
done
```

**ShellCheck**: SC2044 (for loops over find output are fragile)

---

## Additional Patterns

### Case Statement

```bash
case "${command}" in
  start|run)
    start_service
    ;;
  stop)
    stop_service
    ;;
  restart)
    stop_service
    start_service
    ;;
  status)
    show_status
    ;;
  *)
    die "Unknown command: ${command}"
    ;;
esac
```

### Pattern Matching in Case

```bash
case "${filename}" in
  *.tar.gz|*.tgz)
    tar -xzf "${filename}"
    ;;
  *.tar.bz2)
    tar -xjf "${filename}"
    ;;
  *.zip)
    unzip "${filename}"
    ;;
  *)
    die "Unknown archive format: ${filename}"
    ;;
esac
```

### Continue and Break

```bash
for file in "${files[@]}"; do
  # Skip hidden files
  if [[ "${file}" == .* ]]; then
    continue
  fi

  # Stop on error file
  if [[ "${file}" == "STOP" ]]; then
    break
  fi

  process "${file}"
done
```

### Select Menu

```bash
PS3="Select an option: "
select opt in "Option 1" "Option 2" "Quit"; do
  case "${opt}" in
    "Option 1")
      handle_option_1
      ;;
    "Option 2")
      handle_option_2
      ;;
    "Quit")
      break
      ;;
    *)
      echo "Invalid option"
      ;;
  esac
done
```

### Until Loop

```bash
# Wait for condition
until [[ -f "${ready_file}" ]]; do
  info "Waiting for ${ready_file}..."
  sleep 1
done
```
