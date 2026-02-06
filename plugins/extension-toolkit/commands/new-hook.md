---
description: Create a new tool event hook
---

# /new-hook

Create a hook that triggers on tool events.

## Usage

```
/new-hook [event] [matcher] [command]
```

## Hook Events

| Event | When |
|-------|------|
| `PreToolUse` | Before a tool runs |
| `PostToolUse` | After a tool completes |
| `Notification` | When Claude notifies |
| `Stop` | When Claude stops |

## Process

1. **Choose event** - when should it trigger?

2. **Define matcher** - which tools?
   - `Write` - file writes
   - `Edit` - file edits
   - `Bash` - shell commands
   - `tool == "Write" && tool_input.file_path matches "\\.ts$"` - conditional

3. **Write command** - shell command to run
   - Use `$FILE_PATH` for the affected file
   - Return non-zero to block (PreToolUse)
   - Output to stderr for warnings

4. **Add to settings.json**

## Template

```json
{
  "hooks": {
    "{Event}": [
      {
        "matcher": "{matcher}",
        "command": "{command}"
      }
    ]
  }
}
```

## Examples

### Lint after TypeScript write

```json
{
  "matcher": "tool == \"Write\" && tool_input.file_path matches \"\\\\.tsx?$\"",
  "command": "eslint --fix \"$FILE_PATH\""
}
```

### Block dangerous commands

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"rm -rf /\"",
  "command": "echo 'Blocked dangerous command' && exit 1"
}
```

### Format Go files

```json
{
  "matcher": "Write && .go$",
  "command": "gofmt -w \"$FILE_PATH\""
}
```
