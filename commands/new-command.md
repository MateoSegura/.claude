---
description: Create a new slash command
---

# /new-command

Create a user-invokable slash command.

## Usage

```
/new-command [name] [description]
```

## When to Create a Command

Commands are for **user-initiated actions**, not workflows. Use a skill instead if:
- It teaches a multi-step process
- It needs rules, scaffolds, or references
- It should sometimes trigger automatically

## Process

1. **Choose name** - verb-noun format: `build-fix`, `update-docs`, `code-review`

2. **Create file**: `commands/{name}.md`

3. **Write command** with:
   - Description frontmatter
   - What it does
   - How to use it
   - What happens after

## Template

```yaml
---
description: {Brief description shown in /help}
---

# /{name}

{What this command does}

## Usage

```
/{name} [args]
```

## What Happens

1. {Step 1}
2. {Step 2}
3. {Step 3}

## Examples

```
/{name} feature-auth
```

{Expected behavior}
```

## Validation

- [ ] Name is descriptive verb-noun
- [ ] Description is under 80 chars
- [ ] Instructions are clear and actionable
