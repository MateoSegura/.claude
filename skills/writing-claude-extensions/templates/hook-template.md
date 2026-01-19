# Hook Template

Hooks are shell commands triggered by Claude's tool usage.

## Configuration Location

Hooks are configured in `.claude/settings.json`:

```json
{
  "hooks": {
    "PreToolUse": [...],
    "PostToolUse": [...],
    "Notification": [...],
    "Stop": [...]
  }
}
```

## Hook Types

| Type | When | Use For |
|------|------|---------|
| `PreToolUse` | Before a tool runs | Validation, blocking |
| `PostToolUse` | After a tool runs | Formatting, logging |
| `Notification` | On notifications | Alerts, external integrations |
| `Stop` | When Claude stops | Cleanup, summaries |

## Hook Structure

```json
{
  "matcher": "[tool name or pattern]",
  "command": "[shell command to run]"
}
```

## Environment Variables

Hooks receive context via environment variables:

| Variable | Description | Available In |
|----------|-------------|--------------|
| `$TOOL_NAME` | Name of the tool | All |
| `$TOOL_INPUT` | JSON of tool input | PreToolUse, PostToolUse |
| `$TOOL_OUTPUT` | Tool result | PostToolUse |
| `$FILE_PATH` | File being operated on | Write, Edit, Read |
| `$SESSION_ID` | Current session ID | All |

## Hook Templates

### Lint After Write

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "command": "if [[ \"$FILE_PATH\" == *.ts ]]; then eslint --fix \"$FILE_PATH\"; fi"
      }
    ]
  }
}
```

### Format After Edit

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Edit",
        "command": "prettier --write \"$FILE_PATH\" 2>/dev/null || true"
      }
    ]
  }
}
```

### Validate Before Bash

```json
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Bash",
        "command": "echo \"$TOOL_INPUT\" | jq -e '.command | test(\"rm -rf\") | not' || (echo 'Blocked dangerous command' && exit 1)"
      }
    ]
  }
}
```

### Log All Tool Usage

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "*",
        "command": "echo \"$(date): $TOOL_NAME\" >> ~/.claude/tool.log"
      }
    ]
  }
}
```

### Notify on Session End

```json
{
  "hooks": {
    "Stop": [
      {
        "command": "terminal-notifier -message 'Claude session ended' -title 'Claude Code'"
      }
    ]
  }
}
```

### Type Check After TypeScript Changes

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "command": "if [[ \"$FILE_PATH\" == *.ts ]] || [[ \"$FILE_PATH\" == *.tsx ]]; then npx tsc --noEmit 2>&1 | head -20; fi"
      },
      {
        "matcher": "Edit",
        "command": "if [[ \"$FILE_PATH\" == *.ts ]] || [[ \"$FILE_PATH\" == *.tsx ]]; then npx tsc --noEmit 2>&1 | head -20; fi"
      }
    ]
  }
}
```

### Run Tests for Changed Files

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "command": "if [[ \"$FILE_PATH\" == *.test.* ]]; then npm test -- --findRelatedTests \"$FILE_PATH\"; fi"
      }
    ]
  }
}
```

## Hook Patterns

### Matcher Patterns

| Pattern | Matches |
|---------|---------|
| `"Write"` | Exactly the Write tool |
| `"*"` | All tools |
| `"Edit"` | Exactly the Edit tool |
| `"Bash"` | Exactly the Bash tool |

### Command Patterns

| Pattern | Purpose |
|---------|---------|
| `cmd || true` | Don't fail on error |
| `cmd 2>/dev/null` | Suppress stderr |
| `cmd \| head -N` | Limit output |
| `if [[ cond ]]; then cmd; fi` | Conditional execution |
| `exit 1` | Block the operation (PreToolUse) |

## Best Practices

1. **Keep hooks fast** - They run synchronously
2. **Handle errors gracefully** - Use `|| true` for non-critical hooks
3. **Limit output** - Use `head` to avoid flooding
4. **Be specific with matchers** - Avoid `*` unless necessary
5. **Test hooks manually first** - Ensure commands work

## Debugging Hooks

```bash
# Test a hook command manually
FILE_PATH="test.ts" bash -c 'if [[ "$FILE_PATH" == *.ts ]]; then echo "Would lint"; fi'

# Check hook configuration
cat .claude/settings.json | jq '.hooks'
```

## Common Issues

| Issue | Cause | Fix |
|-------|-------|-----|
| Hook doesn't run | Wrong matcher | Check tool name spelling |
| Hook blocks everything | Exit code non-zero | Add `\|\| true` |
| Too much output | Verbose command | Add `\| head -20` |
| Hook too slow | Heavy command | Make async or skip |
