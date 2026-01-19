# Finding Claude Extensions

Detailed guide for discovering all extension types in a project.

## Skills

### Location
```
.claude/skills/*/SKILL.md
```

### Discovery Command
```bash
# List all skills with their descriptions
for skill in .claude/skills/*/SKILL.md; do
  name=$(basename $(dirname "$skill"))
  desc=$(grep -A1 "^description:" "$skill" 2>/dev/null | tail -1 | sed 's/^[ -]*//')
  echo "- $name: $desc"
done
```

### What to Extract

From SKILL.md frontmatter:
```yaml
---
name: skill-name
description: One-line description
---
```

From content:
- Version (look for `> **Version**:`)
- Status (Active/Deprecated)
- Dependencies
- Rule count (count `###.*:.*:` patterns)

### Output Format

```
Skills Found:
┌─────────────────────────────────────────────────────────────┐
│ 1. bubbletea-tui                                            │
│    Complete Bubble Tea TUI development standard             │
│    Rules: 34 | Scaffolds: 4 | Version: 1.0.0               │
├─────────────────────────────────────────────────────────────┤
│ 2. k9s-tui-style                                            │
│    K9s-inspired terminal UI design system                   │
│    Rules: 36 | Scaffolds: 3 | Version: 1.0.0               │
├─────────────────────────────────────────────────────────────┤
│ 3. writing-claude-extensions                                │
│    Guide for creating new skills, commands, hooks           │
│    Templates: 5 | Version: 1.0.0                           │
└─────────────────────────────────────────────────────────────┘
```

---

## Rules

### Location
```
.claude/rules/*.md
```

Also check:
```
.claude/CLAUDE.md  # May contain inline rules
```

### Discovery Command
```bash
# List standalone rules
for rule in .claude/rules/*.md; do
  name=$(basename "$rule" .md)
  tier=$(grep -o ':red_circle:\|:yellow_circle:\|:green_circle:' "$rule" | head -1)
  echo "- $name ($tier)"
done
```

### What to Extract

- Rule name (from filename or `# Title`)
- Scope (Project/Universal)
- Tier (Critical :red_circle: / Required :yellow_circle: / Recommended :green_circle:)
- Enforcement (Always/When X)

### Output Format

```
Rules Found:
┌─────────────────────────────────────────────────────────────┐
│ 1. no-hardcoded-secrets                                     │
│    Scope: Universal | Tier: Critical 🔴                     │
├─────────────────────────────────────────────────────────────┤
│ 2. prefer-composition                                       │
│    Scope: Project | Tier: Required 🟡                       │
└─────────────────────────────────────────────────────────────┘
```

---

## Hooks

### Location
```
.claude/settings.json -> hooks
```

### Discovery Command
```bash
# Extract hooks from settings
cat .claude/settings.json 2>/dev/null | jq '.hooks // empty'
```

### What to Extract

```json
{
  "hooks": {
    "PreToolUse": [
      { "matcher": "Bash", "command": "..." }
    ],
    "PostToolUse": [
      { "matcher": "Write", "command": "..." },
      { "matcher": "Edit", "command": "..." }
    ],
    "Notification": [],
    "Stop": []
  }
}
```

For each hook:
- Type (PreToolUse/PostToolUse/Notification/Stop)
- Matcher (tool name or `*`)
- Command (what it runs)

### Output Format

```
Hooks Found:
┌─────────────────────────────────────────────────────────────┐
│ PreToolUse Hooks:                                           │
│ 1. Bash → Validate dangerous commands                       │
├─────────────────────────────────────────────────────────────┤
│ PostToolUse Hooks:                                          │
│ 2. Write → Lint TypeScript files                            │
│ 3. Edit → Format with Prettier                              │
├─────────────────────────────────────────────────────────────┤
│ Stop Hooks:                                                 │
│ 4. * → Send notification on session end                     │
└─────────────────────────────────────────────────────────────┘
```

---

## Agents

### Location

Varies by project. Common patterns:
```
.claude/agents.yaml
.claude/plugins/*/agents.yaml
plugin.yaml -> agents section
```

### What to Extract

```yaml
agents:
  - name: code-reviewer
    description: Reviews code for bugs and style
    tools: [Read, Glob, Grep]
    model: sonnet
    prompt_prefix: |
      You are a code reviewer...
```

For each agent:
- Name
- Description
- Available tools
- Model (haiku/sonnet/opus)
- Prompt prefix (first line or summary)

### Output Format

```
Agents Found:
┌─────────────────────────────────────────────────────────────┐
│ 1. code-reviewer                                            │
│    Reviews code for bugs and style                          │
│    Model: sonnet | Tools: Read, Glob, Grep                  │
├─────────────────────────────────────────────────────────────┤
│ 2. security-scanner                                         │
│    Scans for security vulnerabilities                       │
│    Model: opus | Tools: Read, Glob, Grep, WebSearch         │
└─────────────────────────────────────────────────────────────┘
```

---

## MCP Servers

### Location

```
.claude/settings.json -> mcpServers
~/.config/claude/claude_desktop_config.json -> mcpServers
```

### Discovery Command
```bash
# From project settings
cat .claude/settings.json 2>/dev/null | jq '.mcpServers // empty'

# From user config
cat ~/.config/claude/claude_desktop_config.json 2>/dev/null | jq '.mcpServers // empty'
```

### What to Extract

```json
{
  "mcpServers": {
    "github": {
      "command": "node",
      "args": ["/path/to/server.js"],
      "env": { "GITHUB_TOKEN": "..." }
    }
  }
}
```

For each server:
- Name (key)
- Command
- Whether it has env vars (don't show values!)

### Output Format

```
MCP Servers Found:
┌─────────────────────────────────────────────────────────────┐
│ 1. github                                                   │
│    Command: node /path/to/github-server.js                  │
│    Env vars: GITHUB_TOKEN                                   │
├─────────────────────────────────────────────────────────────┤
│ 2. database                                                 │
│    Command: python /path/to/db-server.py                    │
│    Env vars: DB_CONNECTION_STRING                           │
└─────────────────────────────────────────────────────────────┘
```

---

## Combined Discovery

### Full Scan Script

```bash
echo "=== Claude Extensions Discovery ==="
echo ""

echo "📚 SKILLS"
for skill in .claude/skills/*/SKILL.md 2>/dev/null; do
  [ -f "$skill" ] && echo "  - $(basename $(dirname $skill))"
done

echo ""
echo "📋 RULES"
for rule in .claude/rules/*.md 2>/dev/null; do
  [ -f "$rule" ] && echo "  - $(basename $rule .md)"
done

echo ""
echo "🪝 HOOKS"
if [ -f .claude/settings.json ]; then
  jq -r '.hooks | to_entries[] | select(.value | length > 0) | "  - \(.key): \(.value | length) hook(s)"' .claude/settings.json 2>/dev/null
fi

echo ""
echo "🤖 MCP SERVERS"
if [ -f .claude/settings.json ]; then
  jq -r '.mcpServers | keys[]? | "  - \(.)"' .claude/settings.json 2>/dev/null
fi
```

---

## Presenting to User

After discovery, use `AskUserQuestion` with clear grouping:

```
I found the following Claude extensions in this project:

**Skills (3)**
- bubbletea-tui: Complete Bubble Tea TUI development standard
- k9s-tui-style: K9s-inspired terminal UI design system
- writing-claude-extensions: Guide for creating new extensions

**Rules (2)**
- no-hardcoded-secrets: Never include secrets in code
- prefer-composition: Favor composition over inheritance

**Hooks (3)**
- PreToolUse:Bash - Block dangerous commands
- PostToolUse:Write - Lint after writing
- PostToolUse:Edit - Format after editing

Which extension would you like to update?
```
