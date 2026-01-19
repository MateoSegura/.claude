# Component Selection Decision Tree

Use this detailed decision tree to determine the correct Claude Code extension type.

## Primary Decision: What's the trigger?

```
Q1: When should this activate?
│
├─► ALWAYS (every conversation/action)
│   │
│   └─► Q2: Is it behavioral guidelines or a process?
│       ├─► Guidelines/standards → RULE
│       └─► Process/workflow → SKILL (with auto-invocation)
│
├─► ON USER REQUEST (explicit invocation)
│   │
│   └─► Q3: How complex is the workflow?
│       ├─► Simple (few steps) → SKILL
│       ├─► Complex (many steps, decisions) → SKILL
│       └─► Just alias for existing behavior → May not need anything
│
├─► ON TOOL EVENT (before/after Claude uses a tool)
│   │
│   └─► HOOK
│       └─► Q4: Should it block or just observe?
│           ├─► Block if validation fails → PreToolUse hook
│           └─► Just run after → PostToolUse hook
│
├─► DELEGATED (Claude spawns it for focused work)
│   │
│   └─► AGENT
│       └─► Q5: What tools does it need?
│           ├─► All tools → May not need custom agent
│           ├─► Subset of tools → Custom agent definition
│           └─► External tools → Need MCP + Agent
│
└─► EXTERNAL TRIGGER (API call, external system)
    │
    └─► MCP SERVER
        └─► Q6: Is it read-only or read-write?
            ├─► Read-only (fetching data) → Simple MCP
            └─► Read-write (mutations) → MCP with careful permissions
```

## Secondary Decision: What's the scope?

```
Q7: How broad is this?
│
├─► UNIVERSAL (applies to all projects)
│   └─► Consider: Should this be in global settings?
│
├─► PROJECT-SPECIFIC (this repo only)
│   └─► Put in: .claude/ directory
│
└─► TEAM-SPECIFIC (shared across team repos)
    └─► Consider: Shared repository of extensions
```

## Tertiary Decision: What's the complexity?

```
Q8: How much does Claude need to "know"?
│
├─► Simple facts/rules (do X, don't do Y)
│   └─► RULE is sufficient
│
├─► Procedures (step 1, then step 2, then...)
│   └─► SKILL with clear steps
│
├─► Decision trees (if X then Y, else Z)
│   └─► SKILL with explicit branching
│
├─► Deep domain knowledge (library/framework expertise)
│   └─► SKILL with reference docs + scaffolds
│
└─► External integrations (APIs, services)
    └─► MCP SERVER
```

## Quick Reference Table

| User Says... | Component | Rationale |
|--------------|-----------|-----------|
| "Always check for X before committing" | Hook | Triggered by git operations |
| "Follow our style guide" | Rule | Simple always-on guidelines |
| "Help me write tests properly" | Skill | Multi-step workflow |
| "Create a command /deploy" | Skill | Skill can be invoked as command |
| "Run linter after every file edit" | Hook | PostToolUse on Write/Edit |
| "Let Claude access our internal API" | MCP | External service integration |
| "Have a specialized code reviewer" | Agent | Delegated focused task |
| "Teach Claude about React patterns" | Skill | Domain knowledge + examples |
| "Warn me about security issues" | Rule or Hook | Depends on when (always vs on-action) |
| "Format code automatically" | Hook | PostToolUse automation |

## Edge Cases

### "I want both X and Y"
Sometimes users need multiple components working together:

```
Example: "I want Claude to follow security rules AND have a security reviewer"

Solution:
1. RULE: Security guidelines (always active)
2. AGENT: Security reviewer (for deep audits)
3. SKILL: Security review workflow (orchestrates when to use agent)
```

### "I'm not sure what I want"
Ask clarifying questions:

1. "Can you give me a specific example of when you'd use this?"
2. "What's the problem you're trying to solve?"
3. "How often would this be needed?"
4. "What should the end result look like?"

### "The existing feature doesn't work how I want"
Consider:
- Is it a bug? → Report/fix the existing feature
- Is it a preference? → Rule to modify behavior
- Is it a different workflow? → New skill

### "I want to modify Claude's core behavior"
Usually not possible via extensions. Explain:
- What CAN be customized (context, guidelines, tools)
- What CANNOT be changed (model capabilities, safety)
