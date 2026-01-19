# Rule Template

Rules are simple always-on guidelines. Use this template for standalone rules.

## File Location

Rules can live in:
- `.claude/rules/[rule-name].md` - Standalone rule file
- `.claude/CLAUDE.md` - Inline in the main Claude file
- `.claude/skills/[skill]/rules/` - Part of a skill

## Standalone Rule Template

```markdown
# [Rule Name]

> **Scope**: [Project | Universal]
> **Enforcement**: [Always | When X]

## Rule

[Clear, concise statement of what Claude should do or avoid]

## Rationale

[Why this rule exists - helps Claude understand when to apply it]

## Examples

### Do This
```[language]
[Good example]
```

### Don't Do This
```[language]
[Bad example]
```

## Exceptions

[When this rule doesn't apply, if any. If no exceptions, state "None."]
```

## Inline Rule (for CLAUDE.md)

```markdown
## [Rule Category]

- **[Rule Name]**: [Brief rule statement]
  - Example: `[code example]`
  - Avoid: `[anti-pattern]`
```

## Examples of Good Rules

### Security Rule
```markdown
# No Hardcoded Secrets

> **Scope**: Universal
> **Enforcement**: Always

## Rule

Never include API keys, passwords, tokens, or other secrets directly in code.
Always use environment variables or secret management systems.

## Rationale

Hardcoded secrets in code can be accidentally committed to version control,
exposing them to anyone with repository access.

## Examples

### Do This
```python
api_key = os.environ.get("API_KEY")
```

### Don't Do This
```python
api_key = "sk-1234567890abcdef"  # NEVER do this
```

## Exceptions

None. This rule has no exceptions.
```

### Style Rule
```markdown
# Use Descriptive Variable Names

> **Scope**: Project
> **Enforcement**: Always

## Rule

Variables should have descriptive names that explain their purpose.
Single-letter names are only acceptable for loop indices (i, j, k) or
well-established conventions (e, err for errors).

## Rationale

Descriptive names make code self-documenting and reduce the need for comments.

## Examples

### Do This
```go
userCount := len(users)
maxRetryAttempts := 3
```

### Don't Do This
```go
x := len(users)
n := 3
```

## Exceptions

- Loop indices: `for i := 0; i < n; i++`
- Error variables: `if err != nil`
- Context: `ctx context.Context`
```

### Process Rule
```markdown
# Always Run Tests Before Committing

> **Scope**: Project
> **Enforcement**: When committing

## Rule

Before creating any commit, run the project's test suite and ensure all tests pass.

## Rationale

Catching test failures before commit prevents broken code from entering the
repository and makes debugging easier.

## Examples

### Do This
```bash
npm test && git commit -m "Add feature"
```

### Don't Do This
```bash
git commit -m "Add feature"  # Without running tests
```

## Exceptions

- Documentation-only changes (no code modified)
- CI configuration changes (tests may not apply)
```

## Rule Naming Conventions

| Pattern | Use For |
|---------|---------|
| `no-[thing]` | Prohibitions (no-hardcoded-secrets) |
| `always-[action]` | Mandatory actions (always-handle-errors) |
| `prefer-[approach]` | Recommendations (prefer-composition) |
| `use-[thing]` | Tool/pattern requirements (use-typescript) |
| `avoid-[thing]` | Discouragements (avoid-any-type) |

## Rule Tiers

| Tier | Marker | Meaning |
|------|--------|---------|
| Critical | :red_circle: | Must follow, no exceptions |
| Required | :yellow_circle: | Should follow, rare exceptions |
| Recommended | :green_circle: | Best practice, flexible |
