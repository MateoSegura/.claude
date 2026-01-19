---
name: meta-skill-create
description: Create new Claude Code extensions (skills, rules, commands, agents, hooks, MCPs)
---

# Creating Extensions

This skill helps create the right extension type for any need.

## Decision Tree

```
What are you trying to do?

├─► "Claude should always follow X"
│   └─► RULE → /new-rule
│
├─► "Claude should know how to do X"
│   ├─► Multi-step workflow? → SKILL → /new-skill
│   └─► Simple guideline? → RULE → /new-rule
│
├─► "I want a /command shortcut"
│   └─► COMMAND → /new-command
│
├─► "Automate on tool events"
│   └─► HOOK → /new-hook
│
├─► "Delegate specialized tasks"
│   └─► AGENT → /new-agent
│
└─► "Access external APIs"
    └─► MCP SERVER (manual setup)
```

## Quick Reference

| Type | When | Command |
|------|------|---------|
| **Skill** | Complex multi-step workflows | `/new-skill` |
| **Rule** | Always-follow guidelines | `/new-rule` |
| **Command** | User-invoked actions | `/new-command` |
| **Agent** | Delegated specialized tasks | `/new-agent` |
| **Hook** | Tool event automation | `/new-hook` |
| **MCP** | External API access | Manual |

## Naming Convention

```
{layer}-{category}-{specific}
```

### Prefixes

**Foundation** (what to know):
- `language-` - Programming standards
- `framework-` - Framework patterns
- `platform-` - Infrastructure
- `tool-` - Tool workflows
- `practice-` - Cross-cutting concerns
- `domain-` - Business domains

**Role** (how to behave):
- `role-manager-` - Task decomposition
- `role-developer-` - Code implementation
- `role-reviewer-` - Code review
- etc.

**Meta** (self-management):
- `meta-skill-` - Skill management
- `meta-rule-` - Rule management
- etc.

## Validation Before Creating

**REJECT if:**
- One-time task (just do it directly)
- Feature already exists
- Unethical/against policy
- Unclear what they want (ask questions first)

**ACCEPT if:**
- Clear, specific use case
- Reusable (multiple uses)
- Appropriate scope
- No duplicates

## Workflow

1. **Understand** - What problem? When triggers? Who benefits?
2. **Decide** - Use decision tree above
3. **Validate** - Check reject/accept criteria
4. **Create** - Use appropriate `/new-*` command
5. **Test** - Verify it works

## Templates

Each `/new-*` command provides a template. Use them.
