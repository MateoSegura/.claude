# IDE Setup

Properly configured IDEs catch issues early and improve developer productivity.

## VS Code

### Required Extensions

| Extension | Purpose |
|-----------|---------|
| golang.go | Official Go extension |
| ms-vscode.vscode-json | JSON support |

### Settings

Add to `.vscode/settings.json`:

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "go.formatTool": "goimports",
  "go.testFlags": ["-v"],
  "go.coverOnSave": true,
  "go.coverageDecorator": {
    "type": "highlight",
    "coveredHighlightColor": "rgba(64,128,64,0.2)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.2)"
  },
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.organizeImports": "explicit"
  },
  "[go]": {
    "editor.insertSpaces": false,
    "editor.tabSize": 4,
    "editor.defaultFormatter": "golang.go"
  },
  "[go.mod]": {
    "editor.defaultFormatter": "golang.go"
  },
  "gopls": {
    "ui.semanticTokens": true,
    "ui.completion.usePlaceholders": true,
    "formatting.gofumpt": false
  }
}
```

### Launch Configuration

Add to `.vscode/launch.json`:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd/myapp"
    },
    {
      "name": "Debug Test",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "program": "${fileDirname}"
    },
    {
      "name": "Attach to Process",
      "type": "go",
      "request": "attach",
      "mode": "local",
      "processId": 0
    }
  ]
}
```

### Tasks Configuration

Add to `.vscode/tasks.json`:

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "Build",
      "type": "shell",
      "command": "go build ./...",
      "group": "build",
      "problemMatcher": "$go"
    },
    {
      "label": "Test",
      "type": "shell",
      "command": "go test -v ./...",
      "group": "test",
      "problemMatcher": "$go"
    },
    {
      "label": "Lint",
      "type": "shell",
      "command": "golangci-lint run",
      "problemMatcher": "$go"
    }
  ]
}
```

---

## GoLand / IntelliJ IDEA

### Essential Settings

1. **Go Modules**: Enable in `Settings > Go > Go Modules`

2. **File Watchers** (`Settings > Tools > File Watchers`):
   - Add `goimports` watcher for import organization
   - Add `gofmt` watcher for formatting

3. **Inspections** (`Settings > Editor > Inspections > Go`):
   - Enable all error-level inspections
   - Enable unused code warnings
   - Enable shadowed variable warnings

### Code Style

`Settings > Editor > Code Style > Go`:

- Tabs: Use tab character
- Tab size: 4
- Imports: Group by stdlib, third-party, project

### Live Templates

Add useful snippets in `Settings > Editor > Live Templates > Go`:

**Error Handling (err)**:
```go
if err != nil {
    return $ZERO$, fmt.Errorf("$MSG$: %w", err)
}
```

**Table Test (tdt)**:
```go
tests := []struct {
    name    string
    input   $INPUT$
    want    $WANT$
    wantErr bool
}{
    {
        name:  "$NAME$",
        input: $INPUT_VAL$,
        want:  $WANT_VAL$,
    },
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        $END$
    })
}
```

### Run Configurations

1. **Go Build**: Build the main package
2. **Go Test**: Run tests with `-v` flag
3. **golangci-lint**: External tool configuration

---

## Neovim

### With nvim-lspconfig

```lua
-- init.lua
require('lspconfig').gopls.setup{
  settings = {
    gopls = {
      analyses = {
        unusedparams = true,
        shadow = true,
      },
      staticcheck = true,
      gofumpt = false,
    },
  },
}
```

### With null-ls for Formatting

```lua
local null_ls = require("null-ls")

null_ls.setup({
  sources = {
    null_ls.builtins.formatting.goimports,
    null_ls.builtins.diagnostics.golangci_lint,
  },
})
```

---

## EditorConfig

Add `.editorconfig` to project root:

```ini
# .editorconfig
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true

[*.go]
indent_style = tab
indent_size = 4

[*.{yaml,yml,json}]
indent_style = space
indent_size = 2

[Makefile]
indent_style = tab
```
