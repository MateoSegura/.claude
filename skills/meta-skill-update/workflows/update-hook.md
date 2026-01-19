# Updating a Hook

Step-by-step workflow for updating an existing hook.

## Overview

Hooks are configured in `.claude/settings.json`. Updates involve modifying the JSON configuration.

```
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│   ANALYZE    │────▶│   CLARIFY    │────▶│   UPDATE     │
│  Current     │     │  Goals       │     │   Config     │
└──────────────┘     └──────────────┘     └──────────────┘
       │                    │                    │
       ▼                    ▼                    ▼
  Read settings.json   What needs to      Edit settings.json
  Understand command   change? Timing?    Verify command
  Test current         Command? Matcher?  Test new behavior
```

---

## Step 1: Analyze Current Hook

### Read Hook Configuration

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

### Understand Hook Components

| Component | Current Value | Purpose |
|-----------|---------------|---------|
| Type | PostToolUse | When it runs |
| Matcher | Write | What triggers it |
| Command | eslint --fix | What it does |

### Test Current Behavior

```bash
# Simulate the hook
FILE_PATH="test.ts" bash -c 'if [[ "$FILE_PATH" == *.ts ]]; then eslint --fix "$FILE_PATH"; fi'
```

---

## Step 2: Identify Update Needs

### Common Update Reasons

| Reason | Example |
|--------|---------|
| Change timing | PreToolUse → PostToolUse |
| Expand scope | Match more file types |
| Update command | New linter version/flags |
| Add error handling | Prevent hook failures |
| Improve performance | Make command faster |
| Fix bugs | Command doesn't work |

### Ask Clarifying Questions

```markdown
Current hook: PostToolUse:Write
Command: `eslint --fix "$FILE_PATH"` for .ts files

What would you like to change?

[ ] Change when it runs (Pre/Post/Stop)
[ ] Expand to more file types
[ ] Update the command
[ ] Add error handling
[ ] Make it faster
[ ] Something else
```

---

## Step 3: Plan Hook Update

### Update Plan Template

```markdown
# Hook Update Plan

## Current
- Type: PostToolUse
- Matcher: Write
- Command: `if [[ "$FILE_PATH" == *.ts ]]; then eslint --fix "$FILE_PATH"; fi`

## Proposed Change
- Type: PostToolUse (unchanged)
- Matcher: Write (unchanged)
- Command: `if [[ "$FILE_PATH" == *.ts ]] || [[ "$FILE_PATH" == *.tsx ]]; then eslint --fix "$FILE_PATH" 2>/dev/null || true; fi`

## Changes Made
1. Added .tsx file support
2. Added error suppression (|| true)
3. Redirected stderr to prevent noise

## Testing
```bash
FILE_PATH="test.tsx" bash -c '[new command]'
```
```

---

## Step 4: Common Hook Updates

### A. Expand File Type Support

**Before:**
```json
{
  "matcher": "Write",
  "command": "if [[ \"$FILE_PATH\" == *.ts ]]; then eslint --fix \"$FILE_PATH\"; fi"
}
```

**After:**
```json
{
  "matcher": "Write",
  "command": "if [[ \"$FILE_PATH\" =~ \\.(ts|tsx|js|jsx)$ ]]; then eslint --fix \"$FILE_PATH\"; fi"
}
```

### B. Add Error Handling

**Before:**
```json
{
  "matcher": "Write",
  "command": "prettier --write \"$FILE_PATH\""
}
```

**After:**
```json
{
  "matcher": "Write",
  "command": "prettier --write \"$FILE_PATH\" 2>/dev/null || true"
}
```

### C. Change Timing

**Before (Post):**
```json
{
  "hooks": {
    "PostToolUse": [
      { "matcher": "Bash", "command": "echo 'Command ran'" }
    ]
  }
}
```

**After (Pre - for validation):**
```json
{
  "hooks": {
    "PreToolUse": [
      { "matcher": "Bash", "command": "echo \"$TOOL_INPUT\" | jq -e '.command | test(\"rm -rf\") | not' || exit 1" }
    ]
  }
}
```

### D. Add Output Limiting

**Before:**
```json
{
  "matcher": "Write",
  "command": "npx tsc --noEmit"
}
```

**After:**
```json
{
  "matcher": "Write",
  "command": "npx tsc --noEmit 2>&1 | head -20"
}
```

### E. Make Conditional on Tool Input

**Before:**
```json
{
  "matcher": "Bash",
  "command": "echo 'Bash used'"
}
```

**After:**
```json
{
  "matcher": "Bash",
  "command": "if echo \"$TOOL_INPUT\" | jq -e '.command | contains(\"git\")' > /dev/null; then echo 'Git command'; fi"
}
```

---

## Step 5: Update settings.json

### Safe Update Process

1. **Read current settings**
```bash
cat .claude/settings.json | jq '.'
```

2. **Backup (optional)**
```bash
cp .claude/settings.json .claude/settings.json.bak
```

3. **Edit the hook**
Use Edit tool to modify the specific hook in settings.json

4. **Verify JSON is valid**
```bash
cat .claude/settings.json | jq '.'
```

5. **Test the command manually**
```bash
FILE_PATH="test.ts" TOOL_INPUT='{"command":"test"}' bash -c '[command]'
```

---

## Step 6: Verify Update

### Verification Checklist

- [ ] JSON is valid (jq parses without error)
- [ ] Command runs without error
- [ ] Command handles edge cases (missing file, wrong type)
- [ ] Error handling works (|| true if needed)
- [ ] Output is limited (| head -N if verbose)
- [ ] Command is fast enough (< 1s for most hooks)

### Test Scenarios

| Scenario | Test |
|----------|------|
| Matching file | Trigger Write on .ts file |
| Non-matching file | Trigger Write on .md file |
| Error case | Trigger with non-existent file |
| Performance | Time the command execution |

---

## Troubleshooting

| Issue | Cause | Fix |
|-------|-------|-----|
| Hook never runs | Wrong matcher | Check tool name spelling |
| Hook blocks everything | Non-zero exit | Add `\|\| true` |
| Too much output | Verbose command | Add `\| head -20` |
| JSON parse error | Invalid JSON | Check quotes, commas |
| Command fails | Missing tool | Check tool is installed |

---

## Hook Update Examples

### Example 1: Add Biome Support

**Goal:** Replace ESLint with Biome

**Before:**
```json
{
  "matcher": "Write",
  "command": "if [[ \"$FILE_PATH\" == *.ts ]]; then eslint --fix \"$FILE_PATH\"; fi"
}
```

**After:**
```json
{
  "matcher": "Write",
  "command": "if [[ \"$FILE_PATH\" =~ \\.(ts|tsx|js|jsx)$ ]]; then biome check --apply \"$FILE_PATH\" 2>/dev/null || true; fi"
}
```

### Example 2: Add Go Formatting

**Goal:** Format Go files on write

**New Hook:**
```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "command": "if [[ \"$FILE_PATH\" == *.go ]]; then gofmt -w \"$FILE_PATH\" && goimports -w \"$FILE_PATH\"; fi"
      }
    ]
  }
}
```

### Example 3: Security Validation

**Goal:** Block writes to sensitive files

**New Hook:**
```json
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Write",
        "command": "if [[ \"$FILE_PATH\" =~ (\\.env|credentials|secrets) ]]; then echo 'Blocked: sensitive file' && exit 1; fi"
      }
    ]
  }
}
```
