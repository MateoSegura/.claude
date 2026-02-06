---
description: "Create a production-grade Claude Code extension (skill, agent, hook, MCP config, or command) through guided design, research, implementation, and adversarial review."
---

# /system-create

Create a Claude Code extension that actually works.

## Usage

```
/system-create [description of what you want to build]
```

## Instructions

Follow the `system-create` skill workflow. Read the skill definition first:

1. Read `skills/system-create/SKILL.md` from this plugin's directory
2. Execute all 6 phases in order: AUDIT → DESIGN → RESEARCH → IMPLEMENT → REVIEW → INTEGRATE
3. Load reference files only when the workflow specifies (not upfront)
4. Use the specialized agents (system-auditor, system-researcher, system-reviewer) as defined in the skill

If the user provided a description as an argument, use it as the starting input for Phase 1 (AUDIT). If no argument was provided, ask the user what they want to build.
