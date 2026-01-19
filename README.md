# .claude

Scalable knowledge base for AI factory agents. Self-maintaining via CI-driven Claude updates.

## Quick Start

```bash
# Add as submodule
git submodule add https://github.com/MateoSegura/.claude.git .claude
```

## Structure

```
.claude/
├── skills/               # Complex workflows (meta-* for self-management)
├── commands/             # Slash commands (/new-skill, /new-rule, etc.)
├── rules/                # Always-on guidelines
├── agents/               # Specialized subagents
├── hooks/                # Tool event triggers
├── mcp-servers/          # External service integrations
├── skilltest/            # Go testing framework
├── settings.json         # Global settings
└── go.mod                # Go module
```

---

## Extension Types

| Type | Location | Discovery | When to Use |
|------|----------|-----------|-------------|
| **Skill** | `skills/{name}/SKILL.md` | Flat | Complex multi-step workflows |
| **Command** | `commands/{name}.md` | Flat | User-invokable actions |
| **Rule** | `rules/**/*.md` | Recursive | Always-follow guidelines |
| **Agent** | `agents/{name}.md` | Flat | Delegated specialized tasks |
| **Hook** | `settings.json` | Config | Automated tool triggers |
| **MCP** | `.mcp.json` | Config | External API access |

---

## Naming Convention

Names are **LLM-parseable** - agents self-select skills based on context.

### Format

```
{layer}-{category}-{specific}
```

**Rules:**
- Lowercase with hyphens
- No abbreviations (`typescript` not `ts`)
- Descriptive over brief

### Layers

#### Foundation (What to Know)

| Prefix | Purpose | Examples |
|--------|---------|----------|
| `language-` | Programming standards | `language-go-cloud`, `language-typescript-node` |
| `framework-` | Framework patterns | `framework-react`, `framework-nextjs` |
| `platform-` | Infrastructure | `platform-kubernetes`, `platform-aws` |
| `tool-` | Tool workflows | `tool-git`, `tool-docker` |
| `practice-` | Cross-cutting concerns | `practice-security-auth`, `practice-testing-tdd` |
| `domain-` | Business domains | `domain-fintech`, `domain-healthcare` |

#### Role (How to Behave)

| Prefix | Role |
|--------|------|
| `role-manager-` | Task decomposition, orchestration |
| `role-architect-` | System design, decisions |
| `role-developer-` | Code implementation |
| `role-reviewer-` | Code review, audits |
| `role-tester-` | Testing strategies |
| `role-devops-` | CI/CD, deployment |
| `role-sre-` | Operations, debugging |
| `role-writer-` | Documentation |
| `role-pm-` | Requirements, prioritization |
| `role-designer-` | UI/UX design |

#### Meta (Self-Management)

| Prefix | Purpose |
|--------|---------|
| `meta-skill-` | Skill management |
| `meta-rule-` | Rule management |
| `meta-command-` | Command management |
| `meta-agent-` | Agent management |
| `meta-hook-` | Hook management |
| `meta-mcp-` | MCP management |

---

## Loading Strategy

When given a task, an LLM should parse for:

1. **Technical signals** → `language-`, `framework-`, `platform-`
2. **Practice signals** → `practice-security-*`, `practice-testing-*`
3. **Role signals** → `role-developer-*`, `role-reviewer-*`

**Example:** "Build a REST API in Go for Kubernetes"
```
language-go-cloud
platform-kubernetes
practice-security-auth
role-developer-backend
```

---

## Meta Commands

Create and maintain extensions:

| Command | Description |
|---------|-------------|
| `/new-skill` | Create a new skill |
| `/new-rule` | Create a new rule |
| `/new-command` | Create a new command |
| `/new-agent` | Create a new agent |
| `/new-hook` | Create a new hook |
| `/update-extension` | Update any extension |
| `/list-extensions` | List all extensions |

---

## Self-Maintenance

This repo maintains itself via weekly CI:

1. **Scan** for outdated patterns
2. **Web search** for latest best practices
3. **Propose** updates via PR
4. **Test** skill changes

---

## Extension Patterns

### Skills

```
skills/{name}/
├── SKILL.md          # Main definition (required)
├── rules/            # Skill-specific rules
├── reference/        # Quick refs, checklists
├── scaffolds/        # Code templates
└── tests/            # Skill tests
```

**SKILL.md format:**
```yaml
---
name: skill-name
description: What it does
---

# Skill Name

Instructions here...
```

### Commands

```yaml
---
description: Brief description
---

# Command Name

What this command does and how to use it.
```

### Rules

```markdown
# Rule Name

What this rule enforces and why.

## Examples

Good and bad examples.
```

### Agents

```yaml
---
name: agent-name
description: What it does
tools: Read, Grep, Glob, Bash
model: sonnet
---

You are a specialized agent for...
```

### Hooks

```json
{
  "hooks": {
    "PostToolUse": [{
      "matcher": "Write",
      "command": "eslint --fix $FILE_PATH"
    }]
  }
}
```

---

## Running Tests

```bash
# Run all skill tests
SKILL_TEST=1 go test ./...

# Run specific skill tests
SKILL_TEST=1 go test ./skills/meta-skill-create/tests/...
```

Tests require:
- Claude CLI installed
- `ANTHROPIC_API_KEY` set

---

## Validation Checklist

Before adding an extension:

- [ ] Uses correct naming convention
- [ ] Has frontmatter with name/description
- [ ] Folder/file matches name field
- [ ] Includes usage examples
- [ ] Has tests (for skills)
