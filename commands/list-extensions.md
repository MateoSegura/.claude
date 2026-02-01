---
description: List all extensions in this knowledge base
---

# /list-extensions

Show all skills, rules, commands, agents, hooks, and MCPs.

## Usage

```
/list-extensions [type]
```

Types: `skills`, `rules`, `commands`, `agents`, `hooks`, `mcps`, `all` (default)

## Output Format

```
## Skills (2)
- meta-skill-create: Guide for creating extensions
- meta-skill-update: Guide for updating extensions

## Commands (7)
- /new-skill: Create a new skill
- /new-rule: Create a new rule
- /new-command: Create a new command
- /new-agent: Create a new agent
- /new-hook: Create a new hook
- /update-extension: Update any extension
- /list-extensions: List all extensions

## Rules (0)
(none)

## Agents (0)
(none)

## Hooks (0)
(none)

## MCPs (0)
(none)
```

## Discovery Locations

| Type | Location | Pattern |
|------|----------|---------|
| Skills | `skills/` | `*/SKILL.md` |
| Commands | `commands/` | `*.md` |
| Rules | `rules/` | `**/*.md` |
| Agents | `agents/` | `*.md` |
| Hooks | `settings.json` | `hooks.*` |
| MCPs | `.mcp.json` | `mcpServers.*` |
