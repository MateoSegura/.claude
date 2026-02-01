# clang-format Configuration

Save as `.clang-format` in project root.

## Configuration

```yaml
---
Language: Cpp
BasedOnStyle: LLVM

# Indentation
IndentWidth: 8
TabWidth: 8
UseTab: Always
IndentCaseLabels: false
NamespaceIndentation: None

# Braces
BreakBeforeBraces: Linux
BraceWrapping:
  AfterFunction: true
  AfterControlStatement: false
  AfterEnum: false
  AfterStruct: false
  BeforeElse: false
  SplitEmptyFunction: false

# Line handling
ColumnLimit: 100
ReflowComments: true
MaxEmptyLinesToKeep: 1

# Alignment
AlignAfterOpenBracket: Align
AlignConsecutiveAssignments: false
AlignConsecutiveDeclarations: false
AlignEscapedNewlines: Left
AlignOperands: true
AlignTrailingComments: true

# Spacing
SpaceAfterCStyleCast: false
SpaceBeforeAssignmentOperators: true
SpaceBeforeParens: ControlStatements
SpaceInEmptyParentheses: false
SpacesInCStyleCastParentheses: false
SpacesInParentheses: false
SpacesInSquareBrackets: false

# Includes
SortIncludes: true
IncludeBlocks: Preserve
IncludeCategories:
  - Regex: '^".*\.h"'
    Priority: 1
  - Regex: '^<zephyr/.*>'
    Priority: 2
  - Regex: '^<(std|errno|string|limits).*>'
    Priority: 3
  - Regex: '^<.*>'
    Priority: 4

# Pointers and references
DerivePointerAlignment: false
PointerAlignment: Right

# Other
AllowShortBlocksOnASingleLine: false
AllowShortCaseLabelsOnASingleLine: false
AllowShortFunctionsOnASingleLine: None
AllowShortIfStatementsOnASingleLine: false
AllowShortLoopsOnASingleLine: false
BreakStringLiterals: true
Cpp11BracedListStyle: false
KeepEmptyLinesAtTheStartOfBlocks: false
...
```

## Usage

```bash
# Format a single file
clang-format -i src/main.c

# Format all C files
find src -name "*.c" -o -name "*.h" | xargs clang-format -i

# Check formatting without modifying
clang-format --dry-run --Werror src/main.c
```
