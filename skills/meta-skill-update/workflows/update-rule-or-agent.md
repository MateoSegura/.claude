# Updating Rules and Agents

Workflows for updating standalone rules and agent definitions.

---

## Part 1: Updating Rules

### Step 1: Analyze Current Rule

Read the rule file from `.claude/rules/[rule-name].md`:

```markdown
# Rule Name

> **Scope**: Project | Universal
> **Enforcement**: Always | When X

## Rule
[Current rule statement]

## Rationale
[Why this rule exists]

## Examples
### Do This
[Good example]

### Don't Do This
[Bad example]

## Exceptions
[When rule doesn't apply]
```

### Step 2: Identify Update Needs

| Update Type | Signs |
|-------------|-------|
| Scope change | Rule applies more/less broadly |
| Tier change | Rule importance has changed |
| Example update | Examples are outdated/wrong |
| Exception addition | New valid exception case |
| Rationale update | Better explanation needed |
| Complete rewrite | Rule fundamentally wrong |

### Step 3: Ask Clarifying Questions

```markdown
Current rule: [rule-name]
Scope: [scope] | Tier: [tier]

What needs updating?

[ ] Change scope (Project ↔ Universal)
[ ] Change tier (Critical/Required/Recommended)
[ ] Update examples
[ ] Add exceptions
[ ] Improve rationale
[ ] Complete rewrite
```

### Step 4: Plan Rule Update

```markdown
# Rule Update Plan: [rule-name]

## Current
- Scope: [scope]
- Tier: [tier]
- Summary: [brief]

## Proposed Changes
1. [Change 1]
2. [Change 2]

## New Content Preview
[Show key changed sections]

## Impact
- Affects: [what code/patterns]
- Breaking: [yes/no]
```

### Step 5: Common Rule Updates

#### Change Tier

```markdown
# Before
> **Enforcement**: Always

## Rule
Avoid using `any` type in TypeScript.

# After
> **Enforcement**: Always

## Rule
:red_circle: **Critical**: Never use `any` type in TypeScript.
```

#### Add Exception

```markdown
# Before
## Exceptions
None.

# After
## Exceptions
- Generated code (auto-generated types may use any)
- Third-party library wrappers when types unavailable
```

#### Update Examples

```markdown
# Before
### Do This
```typescript
const x: string = getValue()
```

# After
### Do This
```typescript
const x: string = getValue()
// Or with type inference when obvious
const items = ["a", "b", "c"] // string[] inferred
```
```

---

## Part 2: Updating Agents

### Step 1: Analyze Current Agent

Agent definitions vary by project. Common structure:

```yaml
name: code-reviewer
description: Reviews code for bugs and style
tools:
  - Read
  - Glob
  - Grep
model: sonnet
prompt_prefix: |
  You are a code reviewer...
```

### Step 2: Identify Update Needs

| Update Type | When |
|-------------|------|
| Tool access | Agent needs more/fewer tools |
| Model change | Task complexity changed |
| Prompt update | Agent behavior needs refinement |
| Description | Better explain when to use |

### Step 3: Ask Clarifying Questions

```markdown
Current agent: [name]
Model: [model]
Tools: [list]

What needs updating?

[ ] Add tools (which?)
[ ] Remove tools (which?)
[ ] Change model (haiku/sonnet/opus)
[ ] Update prompt/behavior
[ ] Improve description
```

### Step 4: Plan Agent Update

```markdown
# Agent Update Plan: [name]

## Current Configuration
- Model: [model]
- Tools: [list]
- Prompt: [first line...]

## Proposed Changes

### Model
[old] → [new]
Reason: [why]

### Tools
Add: [list]
Remove: [list]
Reason: [why]

### Prompt Changes
[Describe changes to prompt_prefix]

## Verification
- [ ] Test with sample task
- [ ] Verify tool access works
- [ ] Check output quality
```

### Step 5: Common Agent Updates

#### Add Web Search

```yaml
# Before
tools:
  - Read
  - Glob
  - Grep

# After
tools:
  - Read
  - Glob
  - Grep
  - WebSearch
  - WebFetch
```

#### Upgrade Model

```yaml
# Before
model: haiku

# After
model: sonnet  # Task requires more reasoning
```

#### Refine Prompt

```yaml
# Before
prompt_prefix: |
  You are a code reviewer.

# After
prompt_prefix: |
  You are a code reviewer focused on finding bugs and security issues.

  Priority order:
  1. Security vulnerabilities
  2. Logic errors
  3. Performance issues
  4. Style/conventions

  For each issue, provide:
  - File and line number
  - Severity (CRITICAL/HIGH/MEDIUM/LOW)
  - Description
  - Suggested fix
```

---

## Verification

### For Rules

- [ ] Rule statement is clear
- [ ] Examples compile/run
- [ ] Exceptions are documented
- [ ] Tier marker is correct (:red_circle:/:yellow_circle:/:green_circle:)

### For Agents

- [ ] Tools list is valid
- [ ] Model is appropriate for task
- [ ] Prompt is clear and focused
- [ ] Description helps users know when to use

---

## Update Checklist

### Rule Update

- [ ] Read current rule
- [ ] Identified what needs changing
- [ ] Asked user for confirmation
- [ ] Made changes
- [ ] Verified examples work
- [ ] Updated any cross-references

### Agent Update

- [ ] Read current agent definition
- [ ] Identified what needs changing
- [ ] Asked user for confirmation
- [ ] Made changes
- [ ] Tested with sample task
- [ ] Updated documentation
