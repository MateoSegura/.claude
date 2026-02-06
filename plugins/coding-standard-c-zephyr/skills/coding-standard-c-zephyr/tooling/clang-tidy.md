# clang-tidy Configuration

Save as `.clang-tidy` in project root.

## Configuration

```yaml
---
Checks: >
  -*,
  bugprone-*,
  -bugprone-easily-swappable-parameters,
  cert-*,
  clang-analyzer-*,
  -clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling,
  concurrency-*,
  misc-*,
  -misc-unused-parameters,
  modernize-use-bool-literals,
  performance-*,
  portability-*,
  readability-braces-around-statements,
  readability-else-after-return,
  readability-function-size,
  readability-identifier-naming,
  readability-implicit-bool-conversion,
  readability-isolate-declaration,
  readability-misleading-indentation,
  readability-misplaced-array-index,
  readability-redundant-*,
  readability-simplify-*

WarningsAsErrors: >
  bugprone-use-after-move,
  cert-err33-c,
  cert-mem30-c,
  cert-mem31-c,
  clang-analyzer-core.*,
  clang-analyzer-deadcode.*,
  concurrency-mt-unsafe,
  readability-braces-around-statements

CheckOptions:
  - key: readability-identifier-naming.FunctionCase
    value: lower_case
  - key: readability-identifier-naming.VariableCase
    value: lower_case
  - key: readability-identifier-naming.GlobalConstantCase
    value: UPPER_CASE
  - key: readability-identifier-naming.MacroDefinitionCase
    value: UPPER_CASE
  - key: readability-identifier-naming.EnumConstantCase
    value: UPPER_CASE
  - key: readability-identifier-naming.StructCase
    value: lower_case
  - key: readability-function-size.LineThreshold
    value: '100'
  - key: readability-function-size.StatementThreshold
    value: '50'
  - key: readability-function-size.BranchThreshold
    value: '10'
  - key: readability-function-size.ParameterThreshold
    value: '6'
  - key: readability-function-size.NestingThreshold
    value: '4'

HeaderFilterRegex: '.*'
...
```

## Key Checks

| Check | Purpose |
|-------|---------|
| `bugprone-*` | Common bug patterns |
| `cert-*` | SEI CERT C rules |
| `clang-analyzer-*` | Static analysis |
| `concurrency-*` | Thread safety issues |
| `readability-braces-around-statements` | Enforce braces |

## Usage

```bash
# Run on a single file
clang-tidy src/main.c -- -I include

# Run with compile_commands.json
clang-tidy -p build src/main.c
```
