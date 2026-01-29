# cppcheck Configuration

## Configuration File

Save as `cppcheck.cfg`:

```
--enable=warning,style,performance,portability
--std=c11
--platform=unix32
--suppress=missingIncludeSystem
--suppress=unusedFunction
--inline-suppr
--error-exitcode=1
--force
```

## Key Options

| Option | Purpose |
|--------|---------|
| `--enable=warning` | Enable warning messages |
| `--enable=style` | Enable style checks |
| `--enable=performance` | Enable performance suggestions |
| `--enable=portability` | Enable portability warnings |
| `--std=c11` | Use C11 standard |
| `--error-exitcode=1` | Return error code on issues |

## Usage

```bash
# Check single file
cppcheck --cppcheck-build-dir=.cppcheck src/main.c

# Check entire project
cppcheck --cppcheck-build-dir=.cppcheck src/

# With configuration file
cppcheck --project=compile_commands.json

# Inline suppression
/* cppcheck-suppress nullPointer */
ptr->field = value;
```

## Common Suppressions

```c
/* Suppress specific check */
/* cppcheck-suppress uninitvar */

/* Suppress for specific line */
x = y; /* cppcheck-suppress unreadVariable */
```
