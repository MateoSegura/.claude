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
├── tests/                # Go testing framework
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

---

## Testing Framework

The `tests/` package provides a Go framework for testing Claude Code extensions by invoking the Claude CLI and validating outputs.

### Extension Test Strategies

| Extension Type | What to Test | Validation Approach |
|----------------|--------------|---------------------|
| **Skill** | Does Claude follow the skill's guidance? | LLM-as-judge or pattern matching |
| **Rule** | Does output comply with the rule? | LLM-as-judge for nuanced rules |
| **Command** | Does it produce correct structure? | Pattern matching, file checks |
| **Agent** | Can it complete specialized tasks? | LLM-as-judge on output quality |
| **Hook** | Does command run with expected effects? | File system / side effect checks |
| **MCP** | Do tools return valid data? | Response validation |

### Quick Start

```bash
# Run the test suite
go run ./tests/cmd/runtest

# Run Go test files (requires SKILL_TEST=1)
SKILL_TEST=1 go test ./skills/meta-skill-create/...
```

Requirements:
- Claude CLI installed and authenticated
- Go 1.23+

### Writing Tests

Tests live alongside extensions as `skill_test.go`:

```
skills/my-skill/
├── SKILL.md
└── skill_test.go    # Tests for this skill
```

**Example test file:**

```go
package my_skill_test

import (
    "context"
    "testing"
    "time"
    "github.com/MateoSegura/.claude/tests"
)

func TestMySkill(t *testing.T) {
    runner := tests.NewTestRunner()

    suite := &tests.Suite{
        Name:          "my-skill",
        ExtensionType: tests.ExtensionSkill,
        Extension:     "my-skill",
        Cases: []*tests.TestCase{
            {
                Name:      "recommends-correct-approach",
                Extension: "my-skill",
                Prompt:    "How should I handle X?",
                Validators: []tests.Validator{
                    tests.LLMValidator(
                        "correct-recommendation",
                        "Response recommends Y approach with proper justification",
                    ),
                },
            },
        },
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
    defer cancel()

    result, _ := runner.RunSuite(ctx, suite)
    if result.Score < 0.70 {
        t.Errorf("Score %.0f%% below 70%% threshold", result.Score*100)
    }
}
```

### Validators

| Validator | Use Case | Example |
|-----------|----------|---------|
| `LLMValidator(name, criteria)` | Nuanced judgment | `LLMValidator("identifies-hook", "Response recommends using a hook")` |
| `ContainsText(text)` | Exact text match | `ContainsText("PostToolUse")` |
| `MatchesRegex(pattern)` | Pattern match | `MatchesRegex(`eslint\|ESLint`)` |
| `ContainsCode(lang)` | Code block exists | `ContainsCode("json")` |
| `NoErrors()` | No error indicators | `NoErrors()` |
| `CustomValidator(name, fn)` | Custom logic | See below |

**Custom validator example:**

```go
tests.CustomValidator("has-steps", func(output string) (bool, string) {
    hasSteps := strings.Contains(output, "1.") && strings.Contains(output, "2.")
    if hasSteps {
        return true, "Found numbered steps"
    }
    return false, "Missing numbered steps"
})
```

### LLM-as-Judge

The `LLMValidator` invokes Claude to evaluate outputs:

```go
tests.LLMValidator(
    "identifies-mcp",
    "The response recommends using an MCP server for database access",
)
```

This is powerful for validating:
- Semantic correctness (not just keywords)
- Quality of explanations
- Appropriate recommendations
- Following nuanced guidelines

### Grading Scale

| Grade | Score | Meaning |
|-------|-------|---------|
| A | 90%+ | Excellent - consistently works |
| B | 80-89% | Good - mostly works |
| C | 70-79% | Acceptable - works with gaps |
| D | 60-69% | Poor - significant issues |
| F | <60% | Failing - needs major work |

### Test Output

Results are saved as JSON to `/tmp/extension-tests/`:

```json
{
  "name": "meta-skill-create-basic",
  "extension_type": "skill",
  "extension": "meta-skill-create",
  "total_tests": 3,
  "passed": 3,
  "failed": 0,
  "score": 1,
  "results": [...]
}
```

---

## Validation Checklist

Before adding an extension:

- [ ] Uses correct naming convention
- [ ] Has frontmatter with name/description
- [ ] Folder/file matches name field
- [ ] Includes usage examples
- [ ] Has tests (for skills)
