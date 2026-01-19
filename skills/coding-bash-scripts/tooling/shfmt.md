# shfmt Configuration

shfmt is a shell script formatter that ensures consistent code style.

## Installation

```bash
# Using Go
go install mvdan.cc/sh/v3/cmd/shfmt@latest

# Using Homebrew (macOS)
brew install shfmt

# Using apt (Debian/Ubuntu with snap)
sudo snap install shfmt

# Using Docker
docker run --rm -v "$(pwd):/mnt" mvdan/shfmt -l -w /mnt
```

## Required Version

- **shfmt**: 3.7.0+

---

## Configuration

shfmt reads settings from `.editorconfig`:

```ini
# .editorconfig

root = true

[*.sh]
# Use spaces for indentation
indent_style = space
indent_size = 2

# Shell variant (bash, posix, mksh, bats)
shell_variant = bash

# Binary operators may start a line
binary_next_line = true

# Indent switch cases
switch_case_indent = true

# Space after redirect operators
space_redirects = false

# Keep column alignment
keep_padding = false

# Function opening brace on same line
function_next_line = false

# End files with newline
insert_final_newline = true

# Trim trailing whitespace
trim_trailing_whitespace = true
```

---

## Command Line Options

```bash
# Format file in place
shfmt -w script.sh

# Format all .sh files
shfmt -w .

# Check formatting (exit 1 if changes needed)
shfmt -d script.sh

# List files that need formatting
shfmt -l .

# Specify shell dialect
shfmt -ln bash script.sh     # Bash
shfmt -ln posix script.sh    # POSIX
shfmt -ln mksh script.sh     # MirBSD ksh
shfmt -ln bats script.sh     # Bats

# Indentation
shfmt -i 2 script.sh         # 2 spaces
shfmt -i 4 script.sh         # 4 spaces
shfmt -i 0 script.sh         # Tabs

# Other options
shfmt -bn script.sh          # Binary ops may start line
shfmt -ci script.sh          # Indent switch cases
shfmt -sr script.sh          # Space after redirects
shfmt -kp script.sh          # Keep column alignment
shfmt -fn script.sh          # Function opening brace on next line
```

---

## Standard Configuration

For this coding standard, use:

```bash
shfmt -ln bash -i 2 -bn -ci -sr=false -w script.sh
```

| Option | Value | Description |
|--------|-------|-------------|
| `-ln bash` | Bash | Target Bash dialect |
| `-i 2` | 2 | Two-space indentation |
| `-bn` | true | Binary operators can start lines |
| `-ci` | true | Indent switch cases |
| `-sr=false` | false | No space after redirects |

---

## Before and After Examples

### Binary Operators

```bash
# Before
if [[ -f "${file}" && -r "${file}" && -s "${file}" ]]; then

# After (-bn)
if [[ -f "${file}" \
  && -r "${file}" \
  && -s "${file}" ]]; then
```

### Switch Case Indentation

```bash
# Before
case "${opt}" in
start)
  start_service
  ;;
stop)
  stop_service
  ;;
esac

# After (-ci)
case "${opt}" in
  start)
    start_service
    ;;
  stop)
    stop_service
    ;;
esac
```

### Redirect Spacing

```bash
# With -sr (space after redirect)
echo "output" > file.txt
cat < input.txt

# Without -sr (no space - default)
echo "output" >file.txt
cat <input.txt

# Standard preference: No space for consistency
echo "output" > file.txt  # We keep the space for readability
```

---

## CI Integration

### GitHub Actions

```yaml
name: shfmt

on: [push, pull_request]

jobs:
  format:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run shfmt
        uses: mvdan/sh-action@v1
        with:
          check: true
```

### GitLab CI

```yaml
shfmt:
  image: mvdan/shfmt:latest
  script:
    - shfmt -d .
  only:
    changes:
      - "**/*.sh"
```

### Pre-commit Hook

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/scop/pre-commit-shfmt
    rev: v3.7.0-1
    hooks:
      - id: shfmt
        args: ["-i", "2", "-bn", "-ci", "-w"]
```

---

## VS Code Integration

Install the [shell-format extension](https://marketplace.visualstudio.com/items?itemName=foxundermoon.shell-format).

`.vscode/settings.json`:

```json
{
  "[shellscript]": {
    "editor.defaultFormatter": "foxundermoon.shell-format",
    "editor.formatOnSave": true,
    "editor.tabSize": 2
  },
  "shellformat.path": "shfmt",
  "shellformat.flag": "-i 2 -bn -ci"
}
```

---

## Makefile Integration

```makefile
.PHONY: fmt fmt-check

SHELL_FILES := $(shell find . -name "*.sh" -type f)

fmt:
	shfmt -ln bash -i 2 -bn -ci -w $(SHELL_FILES)

fmt-check:
	shfmt -ln bash -i 2 -bn -ci -d $(SHELL_FILES)
```

---

## Ignoring Files

shfmt doesn't have a built-in ignore file, but you can:

```bash
# Use find with exclusions
find . -name "*.sh" -not -path "./vendor/*" -print0 | xargs -0 shfmt -d

# Or use .editorconfig to exclude
[vendor/**/*.sh]
indent_style = unset
```

---

## Troubleshooting

### Parser Errors

```bash
# Get more details on parse errors
shfmt -ln bash script.sh 2>&1

# Common issues:
# - Missing quotes
# - Unclosed brackets
# - Heredoc issues
```

### Formatting Conflict with ShellCheck

Some shfmt formatting may trigger ShellCheck warnings. Generally, ShellCheck rules take precedence for correctness.

### Preserving Alignment

```bash
# shfmt may break intentional alignment
readonly FOO="value"
readonly LONGER="another"

# Use -kp to preserve
shfmt -kp script.sh
```
