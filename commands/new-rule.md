---
description: Create a new always-on rule
---

# /new-rule

Create a rule that Claude always follows.

## Usage

```
/new-rule [name] [description]
```

## Process

1. **Determine scope**:
   - Global: `rules/{name}.md`
   - Category: `rules/{category}/{name}.md`

2. **Write rule** with:
   - Clear statement of what to do/not do
   - Why it matters
   - Good and bad examples

3. **Test rule** by asking Claude to do something that should trigger it

## Template

```markdown
# {Rule Name}

{Clear statement of the rule}

## Why

{Explanation of why this rule exists}

## Examples

### Good

```{lang}
{good example}
```

### Bad

```{lang}
{bad example}
```

## Exceptions

{When this rule doesn't apply, if any}
```

## Common Categories

| Category | Examples |
|----------|----------|
| `security/` | No hardcoded secrets, input validation |
| `coding/` | Error handling, naming conventions |
| `workflow/` | Commit messages, PR process |
| `testing/` | Coverage requirements, test patterns |
