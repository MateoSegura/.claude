---
name: meta-skills-create
description: Guide for creating Claude Code extensions (skills, commands, hooks, agents, rules, MCPs)
---

# Writing Claude Extensions

> **Purpose**: Help users create the right type of Claude Code extension
> **Mode**: Always enter planning mode before implementation

This skill guides users through creating Claude Code extensions. It determines the right component type, validates the request, and ensures proper planning before any implementation.

---

## STEP 1: Understand the Request

When a user wants to create something for Claude, first understand their intent:

**Ask yourself:**
1. What problem are they trying to solve?
2. When should this trigger? (always, on command, on event, externally?)
3. Who benefits? (Claude, the user, external systems?)
4. What does it need access to? (files, tools, APIs?)

---

## STEP 2: Component Selection Decision Tree

Use this decision tree to determine the right component type:

```
START: What is the user trying to create?
│
├─► "I want Claude to always follow certain rules/patterns"
│   └─► **RULE** - Guidelines Claude always follows
│
├─► "I want Claude to know how to do X properly"
│   └─► Is it a multi-step workflow or just guidelines?
│       ├─► Multi-step workflow with procedures → **SKILL**
│       └─► Simple guidelines/standards → **RULE**
│
├─► "I want a shortcut command like /foo"
│   └─► Does it just trigger a workflow?
│       ├─► Yes → **SKILL** (skills can be invoked as commands)
│       └─► Needs custom argument parsing? → **COMMAND** (rare)
│
├─► "I want something to happen automatically when X occurs"
│   └─► **HOOK** - Triggered by tool events
│       Examples: pre-commit checks, post-edit validation
│
├─► "I want Claude to delegate complex tasks to specialized workers"
│   └─► **AGENT** - Subagent with specific tools/focus
│
├─► "I want Claude to access external APIs/services"
│   └─► **MCP SERVER** - External tool integration
│
└─► "I want to teach Claude about a library/framework"
    └─► **SKILL** with reference docs and scaffolds
```

---

## STEP 3: Validate the Request

### REJECT if:

| Situation | Why Reject | Suggest Instead |
|-----------|------------|-----------------|
| User wants Claude to do something unethical | Against usage policy | Explain limitations |
| Feature already exists in Claude Code | Redundant | Point to existing feature |
| Request is for a one-time task | Overkill | Just do the task directly |
| User doesn't understand what they want | Need clarity | Ask clarifying questions |
| Request requires capabilities Claude doesn't have | Impossible | Explain limitations |

### Example rejections:

```
❌ "Make a skill that hacks into systems"
   → Reject: Against usage policy

❌ "Make a command that commits code"
   → Reject: /commit already exists, use that

❌ "Make a skill for this one PR review"
   → Reject: Just do the review, no need for a skill

❌ "Make something that makes my code better"
   → Needs clarity: What specific aspect? Style? Performance? Security?
```

### ACCEPT if:

- Clear, specific use case
- Reusable (will be used multiple times)
- Appropriate scope for the component type
- Doesn't duplicate existing functionality

---

## STEP 4: Enter Planning Mode

**ALWAYS enter planning mode before implementation.**

Once the component type is determined and the request is validated:

1. **Use EnterPlanMode tool** to switch to planning mode
2. **Explore the codebase** to understand existing patterns
3. **Ask clarifying questions** using AskUserQuestion:
   - What's the primary use case?
   - What triggers this? (user invokes, always-on, event-based?)
   - What should the output/behavior be?
   - Are there edge cases to handle?
   - What's the minimum viable version?

4. **Write a plan** that includes:
   - Component type and rationale
   - File structure
   - Key behaviors
   - Integration points
   - Testing strategy

5. **Get user approval** before implementing

---

## Component Type Reference

### SKILL
**What**: Workflow definitions with procedures, rules, and scaffolds
**When**: Teaching Claude how to do something complex
**Structure**:
```
.claude/skills/my-skill/
├── SKILL.md          # Main definition
├── rules/            # Specific guidelines
├── reference/        # Quick refs, checklists
└── scaffolds/        # Copy-paste templates
```
**Invocation**: Automatically when relevant, or via `/my-skill`

### RULE
**What**: Simple always-follow guidelines
**When**: Enforcing standards without complex workflows
**Structure**: Single `.md` file in `.claude/rules/` or inline in CLAUDE.md
**Invocation**: Always active

### HOOK
**What**: Shell commands triggered by tool events
**When**: Automating responses to Claude's actions
**Events**: `PreToolUse`, `PostToolUse`, `Notification`, `Stop`
**Structure**: Configured in `.claude/settings.json`
**Example**:
```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "command": "eslint --fix $FILE_PATH"
      }
    ]
  }
}
```

### AGENT
**What**: Specialized subagent for delegated tasks
**When**: Complex task requires focused expertise with limited tools
**Structure**: Defined in plugin system or settings
**Example uses**: Code reviewer, security scanner, test runner

### MCP SERVER
**What**: External tool/API integration via Model Context Protocol
**When**: Claude needs to interact with external services
**Structure**: Separate server process with tool definitions
**Example uses**: GitHub API, database access, cloud services

### COMMAND
**What**: Custom slash command with argument parsing
**When**: Need custom CLI-like interface (rare - skills usually suffice)
**Note**: Most "commands" should just be skills

---

## Planning Questions by Component Type

### For SKILLS:
1. What workflow or process does this teach?
2. What are the key steps?
3. Are there decision points or branches?
4. What scaffolds/templates would help?
5. How should it integrate with other skills?

### For RULES:
1. What specific behavior should always apply?
2. Are there exceptions?
3. How will violations be detected?
4. Is this project-specific or universal?

### For HOOKS:
1. What event triggers this?
2. What tool(s) should it match?
3. What should the hook do?
4. What happens if the hook fails?
5. Should it block the operation or just notify?

### For AGENTS:
1. What specialized task does this handle?
2. What tools does it need access to?
3. What tools should be restricted?
4. How does it report back?
5. What's the failure mode?

### For MCPs:
1. What external service does this connect to?
2. What operations are needed?
3. How is authentication handled?
4. What are the rate limits/costs?
5. How should errors be handled?

---

## Implementation Checklist

Before writing any code, ensure:

- [ ] Component type is confirmed with user
- [ ] Use case is clear and specific
- [ ] Request has been validated (not rejected)
- [ ] Planning mode has been entered
- [ ] Clarifying questions have been asked
- [ ] User has approved the plan
- [ ] Existing patterns in codebase have been reviewed

---

## Example Interaction Flow

```
User: "I want Claude to always use our company's coding style"

Claude: [Analyzes request]
- Problem: Enforce coding standards
- Trigger: Always (not on-demand)
- Scope: Style guidelines

Claude: "This sounds like a RULE - simple guidelines Claude should always follow.

Before I create this, I have a few questions:
1. What language(s) does this cover?
2. Do you have existing style documentation I should reference?
3. Should this include formatting rules, or just patterns/practices?
4. Are there any exceptions where the rules shouldn't apply?"

User: [Provides answers]

Claude: [Enters planning mode, writes plan, gets approval, implements]
```

---

## STEP 5: Write Tests (Required for Skills)

After implementing a skill, **always create tests** to verify it works correctly.

### Test Location

Tests are placed in `.claude/skill-tests/` and are invisible to Claude during normal operation.

### Test Structure

```go
{
    Name:   "test-case-name",
    Skill:  "your-skill-name",
    Prompt: "A specific task that exercises the skill",
    Validators: []Validator{
        ContainsText("expected output text"),
        MatchesRegex(`pattern to find`),
        ContainsCode("go"), // Check for code blocks
        NoErrors(),         // No error indicators
        CustomValidator("name", func(output string) (bool, string) {
            // Custom validation logic
            passed := strings.Contains(output, "expected")
            return passed, "Explanation of result"
        }),
    },
    Iterations: 3, // Run multiple times for consistency
},
```

### Test Categories to Include

1. **Core Functionality**: Does the skill change behavior as expected?
2. **Consistency**: Does it produce consistent results across runs?
3. **Edge Cases**: Does it handle unusual inputs gracefully?
4. **Integration**: Does it work with other skills/rules?

### Running Tests

```bash
# Run all skill tests
cd .claude/skill-tests
SKILL_TEST=1 go test -v ./...

# Run specific test
SKILL_TEST=1 go test -v -run TestYourSkill

# Run with multiple iterations
SKILL_TEST=1 go test -v -count=5 ./...
```

### Grading Scale

| Grade | Score | Meaning |
|-------|-------|---------|
| A | 90%+ | Excellent - consistent, correct behavior |
| B | 80-89% | Good - minor inconsistencies |
| C | 70-79% | Acceptable - works but has gaps |
| D | 60-69% | Poor - significant issues |
| F | <60% | Failing - needs major revision |

### Test Output

Results are saved to `/tmp/skill-tests/`:
- `*-output.txt`: Raw Claude output
- `*-results.json`: Structured results

### Test Checklist

Before marking a skill complete:

- [ ] At least 3 test cases created
- [ ] Each test runs 2-3 iterations
- [ ] Score is 70% or higher (Grade C+)
- [ ] No critical validators failing

---

## Related Skills

- `bubbletea-tui`: For skills that teach TUI development
- `k9s-tui-style`: For skills with K9s design patterns
- `updating-claude-extension`: For updating existing extensions
- Existing coding standards: `coding-standard-*`
