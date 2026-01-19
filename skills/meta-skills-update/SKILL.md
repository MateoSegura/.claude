---
name: meta-skills-update
description: Update existing Claude Code extensions with web search and interactive guidance
---

# Updating Claude Extensions

> **Version**: 1.0.0 | **Status**: Active
> **Dependencies**: writing-claude-extensions (for templates)

This skill helps users discover their existing Claude Code extensions and guides them through updating any extension using web search for latest patterns, codebase exploration, and clarifying questions.

---

## When to Use This Skill

Use this skill when:
- User wants to update an existing skill, rule, hook, agent, or MCP
- User says "update my skill" or "improve my extension"
- User wants to refresh an extension with latest library docs
- User wants to add features to an existing extension

Do NOT use this skill when:
- User wants to create a NEW extension (use `writing-claude-extensions`)
- User just wants to read/view an extension (use Read tool)
- User wants to delete an extension

---

## Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                    STEP 1: DISCOVERY                        │
│                                                             │
│  Scan for all Claude extensions in the project:             │
│  • .claude/skills/*/SKILL.md                               │
│  • .claude/rules/*.md                                      │
│  • .claude/settings.json (hooks section)                   │
│  • Plugin agents (if any)                                  │
│  • MCP servers in settings                                 │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                 STEP 2: PRESENT OPTIONS                     │
│                                                             │
│  Use AskUserQuestion to show discovered extensions:         │
│  • Group by type (Skills, Rules, Hooks, etc.)              │
│  • Show name and brief description                         │
│  • Let user select which to update                         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│               STEP 3: ANALYZE CURRENT STATE                 │
│                                                             │
│  Read and understand the selected extension:                │
│  • What does it currently do?                              │
│  • What rules/patterns does it define?                     │
│  • What dependencies does it have?                         │
│  • When was it last updated?                               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│              STEP 4: GATHER UPDATE CONTEXT                  │
│                                                             │
│  Parallel research:                                         │
│  • Web search for latest patterns/docs                     │
│  • Explore codebase for current usage                      │
│  • Check if dependencies have new versions                 │
│  • Look for issues/gaps in current implementation          │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│              STEP 5: CLARIFY UPDATE GOALS                   │
│                                                             │
│  Ask user specific questions:                               │
│  • What aspects need updating?                             │
│  • Are there specific issues to fix?                       │
│  • Should we add new features?                             │
│  • Update to match new library versions?                   │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│              STEP 6: PLAN MODE & UPDATE                     │
│                                                             │
│  Enter plan mode with:                                      │
│  • Summary of current state                                │
│  • Proposed changes                                        │
│  • Impact analysis                                         │
│  • Implementation steps                                    │
└─────────────────────────────────────────────────────────────┘
```

---

## Step 1: Discovery

### Finding Skills

```bash
# Find all skills
find .claude/skills -name "SKILL.md" 2>/dev/null
```

Or use Glob:
```
.claude/skills/*/SKILL.md
```

Parse each SKILL.md for frontmatter:
- `name`: Skill identifier
- `description`: One-line summary

### Finding Rules

```bash
# Find standalone rules
find .claude/rules -name "*.md" 2>/dev/null
```

Or use Glob:
```
.claude/rules/*.md
```

### Finding Hooks

Read `.claude/settings.json` and extract the `hooks` section:

```json
{
  "hooks": {
    "PreToolUse": [...],
    "PostToolUse": [...],
    "Notification": [...],
    "Stop": [...]
  }
}
```

### Finding Agents

Check for plugin agent definitions (varies by project structure).

### Finding MCP Servers

Read `.claude/settings.json` or `claude_desktop_config.json` for `mcpServers`.

---

## Step 2: Present Options

Use `AskUserQuestion` to present discovered extensions:

```markdown
## Available Extensions

### Skills
1. bubbletea-tui - Complete Bubble Tea TUI development standard
2. k9s-tui-style - K9s-inspired terminal UI design system
3. writing-claude-extensions - Guide for creating new extensions

### Rules
4. no-hardcoded-secrets - Universal security rule
5. prefer-composition - Go design pattern

### Hooks
6. PostToolUse:Write - Lint TypeScript files after write
7. PreToolUse:Bash - Block dangerous commands

Which extension would you like to update?
```

---

## Step 3: Analyze Current State

For the selected extension, read and understand:

### For Skills
- Read SKILL.md frontmatter and content
- Read all files in rules/, reference/, scaffolds/
- Count rules, patterns, scaffolds
- Identify key concepts covered

### For Rules
- Read the rule file
- Understand scope (Project/Universal)
- Note tier (Critical/Required/Recommended)
- Check examples

### For Hooks
- Parse the hook configuration
- Understand trigger (PreToolUse, PostToolUse, etc.)
- Analyze the command
- Check matcher pattern

### For Agents
- Read agent definition
- Note tools available
- Understand prompt_prefix
- Check model setting

### For MCPs
- Read server implementation
- List available tools
- Check dependencies
- Review security measures

---

## Step 4: Gather Update Context

### Web Search Queries

| Extension Type | Search Queries |
|----------------|----------------|
| Library skill | "[library] latest features 2024", "[library] best practices" |
| Framework skill | "[framework] changelog", "[framework] migration guide" |
| Tool hook | "[tool] CLI options", "[tool] new flags" |
| API MCP | "[api] v2 changes", "[api] new endpoints" |

### Codebase Exploration

| What to Find | Why |
|--------------|-----|
| Usage of skill patterns | Are the rules being followed? |
| Violations of rules | What's not working? |
| New patterns not covered | What's missing? |
| Deprecated patterns | What should be removed? |

### Version Checking

```bash
# Check package versions
cat go.mod | grep [library]
cat package.json | jq '.dependencies["[library]"]'
```

Compare installed version with:
- Latest stable release
- Version documented in skill
- Breaking changes between versions

---

## Step 5: Clarify Update Goals

Use `AskUserQuestion` with specific options:

### For Skills

```
What aspects of the skill need updating?

[ ] Update for new library version
[ ] Add missing patterns/rules
[ ] Fix incorrect examples
[ ] Add new scaffolds
[ ] Improve documentation
[ ] Other (describe)
```

### For Rules

```
How should the rule be updated?

[ ] Change enforcement (Critical/Required/Recommended)
[ ] Update examples
[ ] Add exceptions
[ ] Modify scope
[ ] Other (describe)
```

### For Hooks

```
What hook changes are needed?

[ ] Change trigger timing (Pre/Post)
[ ] Update command
[ ] Change matcher
[ ] Add error handling
[ ] Other (describe)
```

---

## Step 6: Plan Mode & Update

After gathering all context, enter plan mode with:

### Plan Structure

```markdown
# Update Plan: [Extension Name]

## Current State
- Version: [current]
- Last updated: [date if known]
- Key features: [list]

## Proposed Changes

### Additions
- [New rule/pattern/feature]

### Modifications
- [Changes to existing content]

### Removals
- [Deprecated or incorrect content]

## Impact Analysis
- Files affected: [count]
- Breaking changes: [yes/no]
- Testing required: [what to verify]

## Implementation Steps
1. [First step]
2. [Second step]
...
```

---

## Extension Type Reference

| Type | Location | Key Files |
|------|----------|-----------|
| Skill | `.claude/skills/[name]/` | SKILL.md, rules/, reference/, scaffolds/ |
| Rule | `.claude/rules/` | [rule-name].md |
| Hook | `.claude/settings.json` | hooks.PreToolUse, hooks.PostToolUse |
| Agent | Plugin definition | agents.yaml or similar |
| MCP | Varies | package.json, src/index.ts |

---

## Update Checklist

- [ ] Discovered all extensions in project
- [ ] Presented options to user
- [ ] Read and understood current state
- [ ] Searched web for latest patterns
- [ ] Explored codebase for usage
- [ ] Checked for version updates
- [ ] Asked user for update goals
- [ ] Created update plan
- [ ] Entered plan mode for approval
- [ ] Implemented approved changes
- [ ] Verified changes work

---

## Common Update Scenarios

### Scenario: Library Version Update

1. Check current version in skill
2. Check installed version in project
3. Web search for changelog
4. Identify breaking changes
5. Update affected rules and scaffolds
6. Update version references

### Scenario: Adding Missing Patterns

1. Explore codebase for patterns not covered
2. Web search for common patterns
3. Draft new rules
4. Add scaffolds if helpful
5. Update quick reference

### Scenario: Fixing Incorrect Examples

1. Identify incorrect examples
2. Find correct patterns (docs, working code)
3. Update examples
4. Verify scaffolds compile
5. Update related rules if needed

### Scenario: Expanding Scope

1. Understand current scope
2. Identify gaps
3. Research new areas
4. Draft new sections
5. Maintain consistency with existing content

---

## Step 7: Run Tests (Required for Skills)

After updating a skill, **always run tests** to verify the update didn't break anything.

### Running Existing Tests

```bash
# Run tests for the updated skill
cd .claude/skill-tests
SKILL_TEST=1 go test -v -run Test[SkillName]

# Run with multiple iterations for consistency check
SKILL_TEST=1 go test -v -count=3 -run Test[SkillName]
```

### Update Tests If Needed

If your update changes expected behavior:

1. **Read existing tests** in `.claude/skill-tests/`
2. **Update validators** to match new expectations
3. **Add new test cases** for new functionality
4. **Remove obsolete tests** for removed features

### Minimum Test Coverage

Before marking update complete:

- [ ] All existing tests still pass
- [ ] New functionality has tests
- [ ] Score is 70% or higher (Grade C+)
- [ ] No regressions in previous behavior

### Test Score Tracking

Track scores over time to measure improvement:

```
Before update: 85% (Grade B)
After update: 92% (Grade A)
Improvement: +7%
```

### Test Output Location

Results saved to `/tmp/skill-tests/`:
- `*-output.txt`: Raw Claude output for each test
- `*-results.json`: Structured test results with scores

---

## Related Skills

- `writing-claude-extensions`: For creating new extensions
- `bubbletea-tui`: TUI development patterns
- `k9s-tui-style`: K9s-style UI design
