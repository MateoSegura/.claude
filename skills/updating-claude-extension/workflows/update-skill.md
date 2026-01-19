# Updating a Skill

Step-by-step workflow for updating an existing skill.

## Overview

```
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│   ANALYZE    │────▶│   RESEARCH   │────▶│    PLAN      │
│  Current     │     │  Updates     │     │   Changes    │
└──────────────┘     └──────────────┘     └──────────────┘
       │                    │                    │
       ▼                    ▼                    ▼
  Read all files      Web search for       Draft update
  Count rules         latest docs          plan with
  Check versions      Check codebase       specific changes
  Note gaps           Find patterns
```

---

## Step 1: Analyze Current State

### Files to Read

```
.claude/skills/[skill-name]/
├── SKILL.md              # Main definition
├── rules/*.md            # All rule files
├── reference/*.md        # Reference docs
└── scaffolds/*           # Code templates
```

### Extract Key Information

From SKILL.md:
```markdown
---
name: [skill-name]
description: [description]
---

> **Version**: X.Y.Z | **Status**: Active/Deprecated
> **Dependencies**: [list]
```

### Create State Summary

```markdown
## Current State: [skill-name]

**Version**: X.Y.Z
**Description**: [description]
**Dependencies**: [list]

### Rules
| ID | Name | Tier |
|----|------|------|
| BTT-ARC-001 | Immutable Updates | Critical |
| ... | ... | ... |

Total: N rules across M categories

### Scaffolds
| File | Purpose |
|------|---------|
| basic-app.go | Minimal working app |
| ... | ... |

### Reference Docs
- quick-reference.md
- code-review.md
```

---

## Step 2: Research Updates

### Web Search Queries

| Purpose | Query |
|---------|-------|
| Latest version | "[library] latest release [year]" |
| Changelog | "[library] changelog breaking changes" |
| New features | "[library] new features [version]" |
| Best practices | "[library] best practices [year]" |
| Migration | "[library] migration guide v[old] to v[new]" |

### Codebase Exploration

Search for:
1. **Usage patterns** - How is this library used in the project?
2. **Rule violations** - Are current rules being followed?
3. **Missing patterns** - What patterns exist that aren't documented?
4. **Deprecated usage** - What patterns should be removed?

```bash
# Example: Find bubbletea usage
grep -r "tea\." --include="*.go" .
grep -r "bubbletea" --include="*.go" .
```

### Version Comparison

```bash
# Check installed version
cat go.mod | grep [library]
cat package.json | jq '.dependencies["[library]"]'

# Compare with skill documentation
grep -i "version" .claude/skills/[skill]/SKILL.md
```

---

## Step 3: Identify Update Types

### A. Version Update

The skill references an old library version:

```markdown
## Changes Needed

### Version References
- [ ] Update version in SKILL.md header
- [ ] Update any version-specific examples
- [ ] Review breaking changes

### API Changes
- [ ] [Old API] → [New API]
- [ ] [Deprecated feature] → [Replacement]

### New Features to Document
- [ ] [Feature 1] - added in vX.Y
- [ ] [Feature 2] - added in vX.Y
```

### B. Pattern Addition

The codebase uses patterns not covered by the skill:

```markdown
## Missing Patterns

### From Codebase Analysis
- [ ] Pattern: [description] - found in [file]
- [ ] Pattern: [description] - found in [file]

### From Documentation Review
- [ ] Pattern: [description] - in official docs
- [ ] Pattern: [description] - common community pattern

### Proposed Rules
- [ ] [RULE-ID]: [Rule name] - covers [pattern]
```

### C. Rule Correction

Existing rules have issues:

```markdown
## Rule Issues

### Incorrect Examples
- [ ] [RULE-ID]: Example does not compile
- [ ] [RULE-ID]: Example uses deprecated API

### Outdated Recommendations
- [ ] [RULE-ID]: [Old way] → [New way]

### Missing Exceptions
- [ ] [RULE-ID]: Should allow [exception case]
```

### D. Scaffold Updates

Code templates need updating:

```markdown
## Scaffold Updates

### Compilation Issues
- [ ] [scaffold.go]: Fails with [error]

### Outdated Patterns
- [ ] [scaffold.go]: Uses [old pattern]

### Missing Scaffolds
- [ ] [New scaffold]: For [use case]
```

---

## Step 4: Ask Clarifying Questions

Use `AskUserQuestion`:

```markdown
I've analyzed the [skill-name] skill and found potential updates.

**Current State:**
- Version: X.Y.Z
- Rules: N
- Scaffolds: M

**Potential Updates Found:**

1. **Version Update**
   Installed library is v[new], skill documents v[old]

2. **Missing Patterns**
   Found [N] patterns in codebase not covered by skill

3. **Outdated Rules**
   [N] rules reference deprecated patterns

Which updates should we focus on?
[ ] All of the above
[ ] Version update only
[ ] Add missing patterns only
[ ] Fix outdated rules only
[ ] Something else (describe)
```

---

## Step 5: Plan the Update

### Update Plan Template

```markdown
# Update Plan: [skill-name] Skill

## Summary
[One paragraph describing what will be updated and why]

## Changes by File

### SKILL.md
- [ ] Update version header to X.Y.Z
- [ ] Add section on [new feature]
- [ ] Update rule [RULE-ID] example

### rules/[category].md
- [ ] Add rule [NEW-ID]: [description]
- [ ] Update rule [RULE-ID]: [change]
- [ ] Remove rule [OLD-ID]: [reason]

### scaffolds/[file]
- [ ] Update to use [new pattern]
- [ ] Fix compilation error in [function]
- [ ] Add new scaffold: [name]

### reference/quick-reference.md
- [ ] Add new rules to table
- [ ] Update cheatsheet section

## Breaking Changes
- [List any changes that might affect users]

## Verification
- [ ] All scaffolds compile
- [ ] Examples run correctly
- [ ] No rule ID conflicts
- [ ] Version numbers consistent
```

---

## Step 6: Implement Updates

### Order of Operations

1. **Update scaffolds first** - Verify they compile
2. **Update rules** - Add/modify/remove rules
3. **Update SKILL.md** - Main documentation
4. **Update reference docs** - Quick reference, checklists
5. **Verify everything** - Final checks

### Verification Steps

```bash
# For Go scaffolds
go build ./...

# For TypeScript scaffolds
npx tsc --noEmit

# For general verification
# - Read each scaffold and ensure it's valid
# - Check rule IDs are unique
# - Verify cross-references work
```

---

## Common Update Scenarios

### Scenario: Bubble Tea Update

```markdown
## bubbletea-tui Update: v1.3 → v2.0

### API Changes
- [ ] `tea.WindowSizeMsg` → `tea.WindowSizeMsg` (unchanged)
- [ ] New: `tea.Suspend`, `tea.Resume` commands
- [ ] Deprecated: [old feature]

### New Rules to Add
- [ ] BTT-CMD-008: Use Suspend for shell commands
- [ ] BTT-ARC-006: New lifecycle hooks

### Scaffolds to Update
- [ ] basic-app.go: Add suspend handling example
```

### Scenario: Adding Missing Pattern

```markdown
## Adding Table Component Pattern

### Found In
- tools/cody/internal/tui/screens/projects.go (line 45-89)

### New Content
1. Add to rules/components.md:
   - BTT-CMP-006: Table Component Pattern

2. Add to scaffolds/:
   - table-view.go: Table component scaffold

3. Update reference/quick-reference.md:
   - Add table to component cheatsheet
```

### Scenario: Fixing Compilation Error

```markdown
## Fixing Scaffold: component.go

### Issue
Line 23: `undefined: tea.KeyType`

### Fix
Replace:
```go
if msg.Type == tea.KeyEnter {
```

With:
```go
if key, ok := msg.(tea.KeyMsg); ok && key.Type == tea.KeyEnter {
```

### Verification
```bash
go build .claude/skills/bubbletea-tui/scaffolds/component.go
```
```
