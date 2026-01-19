---
description: Create a new specialized subagent
---

# /new-agent

Create a subagent for delegated tasks.

## Usage

```
/new-agent [name] [description]
```

## When to Create an Agent

Agents handle **focused, delegated tasks** with limited tools. Create one when:
- Task requires specialized expertise
- Main agent shouldn't be distracted
- Task can run in parallel
- Limited tool access improves focus

## Process

1. **Define purpose** - what specific task does it handle?

2. **Choose tools** - minimum needed:
   - `Read, Grep, Glob` - for analysis
   - `+ Bash` - for commands
   - `+ Edit, Write` - for modifications

3. **Select model**:
   - `haiku` - fast, simple tasks
   - `sonnet` - balanced (default)
   - `opus` - complex reasoning

4. **Create file**: `agents/{name}.md`

## Template

```yaml
---
name: {name}
description: {What it does}
tools: Read, Grep, Glob, Bash
model: sonnet
---

You are a specialized agent for {purpose}.

## Your Task

{Clear description of what to do}

## Constraints

- {Constraint 1}
- {Constraint 2}

## Output Format

{How to report results}
```

## Common Agent Types

| Agent | Tools | Purpose |
|-------|-------|---------|
| `code-reviewer` | Read, Grep, Glob | Review code quality |
| `security-scanner` | Read, Grep, Glob, Bash | Find vulnerabilities |
| `test-runner` | Read, Bash | Run and analyze tests |
| `doc-updater` | Read, Grep, Edit | Sync documentation |
