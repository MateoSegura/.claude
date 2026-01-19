---
description: Create a new skill in this knowledge base
---

# /new-skill

Create a new skill for teaching Claude complex workflows.

## Usage

```
/new-skill [name] [description]
```

## Process

1. **Validate name** follows convention: `{layer}-{category}-{specific}`
   - Foundation: `language-`, `framework-`, `platform-`, `tool-`, `practice-`, `domain-`
   - Role: `role-{role}-{area}`
   - Meta: `meta-{type}-{action}`

2. **Check for duplicates** - search existing skills

3. **Create structure**:
   ```
   skills/{name}/
   ├── SKILL.md          # Main definition
   ├── rules/            # Skill-specific rules (optional)
   ├── reference/        # Quick refs (optional)
   ├── scaffolds/        # Code templates (optional)
   └── tests/            # Skill tests (required)
   ```

4. **Generate SKILL.md** with:
   - Frontmatter (name, description)
   - Purpose and when to use
   - Step-by-step workflow
   - Examples

5. **Create basic test** in `tests/`

## Template

```yaml
---
name: {name}
description: {description}
---

# {Title}

> **Purpose**: {purpose}

## When to Use

- {use case 1}
- {use case 2}

## Workflow

1. **Step 1**: {description}
2. **Step 2**: {description}

## Examples

{examples}
```

## Validation

After creation:
- [ ] Name follows naming convention
- [ ] SKILL.md has valid frontmatter
- [ ] At least one test case exists
- [ ] Skill is discoverable via `/list-extensions`
