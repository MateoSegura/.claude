# Skill Template

Copy this template when creating a new skill.

## Directory Structure

```
.claude/skills/[skill-name]/
├── SKILL.md              # Main skill definition (required)
├── rules/                # Specific rules/guidelines (optional)
│   ├── rule-category-1.md
│   └── rule-category-2.md
├── reference/            # Quick references (optional)
│   ├── quick-reference.md
│   └── code-review.md
└── scaffolds/            # Copy-paste templates (optional)
    ├── template-1.go
    └── template-2.go
```

## SKILL.md Template

```markdown
---
name: [skill-name]
description: [One-line description for skill discovery]
---

# [Skill Title]

> **Version**: 1.0.0 | **Status**: Active
> **Dependencies**: [Other skills this builds on, if any]

[2-3 sentence overview of what this skill teaches Claude to do.]

---

## When to Use This Skill

Use this skill when:
- [Condition 1]
- [Condition 2]
- [Condition 3]

Do NOT use this skill when:
- [Anti-condition 1]
- [Anti-condition 2]

---

## Navigation

### Rules

| Category | File | Description |
|----------|------|-------------|
| [Category] | [rules/file.md] | [Brief description] |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | [Purpose] |

### Scaffolds

| Template | File | Purpose |
|----------|------|---------|
| [Template Name] | [scaffolds/file.ext] | [When to use] |

---

## Core Workflow

### Step 1: [First Step]

[Explanation of what to do]

```[language]
// Example code or pattern
```

### Step 2: [Second Step]

[Explanation]

### Step 3: [Third Step]

[Explanation]

---

## Key Rules

### [RULE-ID-001]: [Rule Name] :red_circle:

**Tier**: Critical | Required | Recommended

[Explanation of the rule]

```[language]
// CORRECT
[good example]

// INCORRECT
[bad example]
```

### [RULE-ID-002]: [Rule Name] :yellow_circle:

[Continue pattern...]

---

## Quick Reference

| ID | Rule | Tier |
|----|------|------|
| [RULE-ID-001] | [Brief description] | Critical |
| [RULE-ID-002] | [Brief description] | Required |

---

## Common Patterns

### Pattern 1: [Pattern Name]

[Explanation and example]

### Pattern 2: [Pattern Name]

[Explanation and example]

---

## Troubleshooting

| Problem | Cause | Solution |
|---------|-------|----------|
| [Issue] | [Why it happens] | [How to fix] |

---

## References

- [External doc 1](url)
- [External doc 2](url)
```

## Rule File Template (rules/*.md)

```markdown
# [Category] Rules

## [RULE-ID-001]: [Rule Name] :red_circle:

**Tier**: Critical

[Full explanation of the rule]

### Why This Matters

[Rationale]

### Examples

```[language]
// CORRECT
[good code]

// INCORRECT
[bad code]
```

### Exceptions

[When this rule doesn't apply, if any]

---

## [RULE-ID-002]: [Rule Name] :yellow_circle:

[Continue pattern...]
```

## Reference File Template (reference/quick-reference.md)

```markdown
# [Skill Name] Quick Reference

## Rule Summary

| ID | Rule | Tier |
|----|------|------|
| [ID] | [Description] | [Tier] |

## Common Patterns

### [Pattern Name]
```[language]
[minimal example]
```

## Cheatsheet

| Task | How |
|------|-----|
| [Task] | [Quick answer] |

## Checklist

- [ ] [Item 1]
- [ ] [Item 2]
```

## Scaffold Template (scaffolds/*.ext)

```[language]
// Package [name] - [Brief description]
//
// USAGE: Copy this file and modify for your needs.
// [Additional usage instructions]
package [name]

// [Code with clear comments explaining customization points]
```
