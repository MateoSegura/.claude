# Extension Anatomy

Loaded during Phase 2 (Architecture). Every extension type, its file structure, and when to use it.

---

## Skill

**Use when**: Multi-step workflow with phases, file references, and possibly agent orchestration.

**Context cost**: Medium-High (SKILL.md loads on invocation, supporting files load on demand).

**Structure**:
```
skills/<name>/
├── SKILL.md              # Entry point — the ROUTER (not a knowledge dump)
├── rules/                # Constraints, one file per topic
│   ├── 01-<topic>.md     # Numbered for load order
│   └── 02-<topic>.md
├── reference/            # Lookup tables, specs, cheat sheets
│   └── <topic>.md
├── scaffolds/            # Code templates with customization markers
│   └── <name>.<ext>
└── examples/             # Worked examples
    └── <scenario>.md
```

**SKILL.md anatomy**:
```yaml
---
name: <name>
description: <SPECIFIC trigger sentence — Claude uses this to decide auto-invocation>
# Optional controls:
# disable-model-invocation: true   # Only manual /invoke
# user-invocable: false            # Only Claude can invoke
# context: fork                    # Run in isolated subagent (protects main context)
# model: claude-sonnet-4-5         # Override model for this skill
# argument-hint: "[args]"          # Shown in autocomplete
# allowed-tools: Read, Grep, Glob  # Restrict available tools
---

# Title

> One-line purpose.

## When to Use
- Trigger condition 1
- Trigger condition 2

## Files (load on demand)

| File | Purpose | Load When |
|------|---------|-----------|
| [rules/01-x.md](rules/01-x.md) | X constraints | Phase N |

## Workflow

### Phase 1: NAME
1. Step (read specific file if needed)
2. Step
```

**Critical rules for skills**:
- SKILL.md is a router. It tells Claude what files exist and when to load them.
- Never paste the content of reference files into SKILL.md.
- Every file in the table must actually exist.
- Description must be specific: "Apply C/Zephyr RTOS coding standard when writing or reviewing C code for Zephyr-based firmware projects" NOT "coding standard helper".
- Workflow phases have clear entry conditions and exit criteria.

---

## Agent

**Use when**: Focused, delegated task that should run in isolated context with limited tools.

**Context cost**: Isolated (runs in subagent, does not consume main conversation context).

**Structure**:
```
agents/<name>.md          # Single file defining the agent
```

**Agent file anatomy**:
```markdown
---
name: <name>
description: <what this agent does>
tools: Read, Grep, Glob
model: sonnet
---

You are a <role> agent. Your task is <specific task>.

## Input

You will receive:
- <data item 1>
- <data item 2>

## Process

1. <Step>
2. <Step>

## Output

Return a structured response:

### Findings
- <finding format>

### Recommendation
- <recommendation format>

## Constraints

- <what NOT to do>
- <scope limits>
- <time/context budget>
```

**Critical rules for agents**:
- Minimum viable tool set. Don't give Write access to an analysis agent.
- Clear input/output contract. The skill invoking this agent must know exactly what it sends and what it gets back.
- Use `model: haiku` for simple classification/scanning tasks. Use `sonnet` for analysis. Use `opus` only for complex reasoning.
- Constraints are as important as instructions. Explicitly state what the agent must NOT do.

---

## Hook

**Use when**: Automation that should run on tool events. Formatting, linting, validation, safety checks.

**Context cost**: ZERO. Hooks run as shell commands, not in Claude's context.

**Structure**:
```
hooks/
└── hooks.json
```

**hooks.json anatomy**:
```json
{
  "hooks": {
    "<Event>": [
      {
        "matcher": "<tool name regex>",
        "hooks": [
          {
            "type": "command",
            "command": "<shell command>"
          }
        ]
      }
    ]
  }
}
```

**Events**:
| Event | Fires | Use For |
|-------|-------|---------|
| `PreToolUse` | Before tool runs | Block dangerous operations, validation |
| `PostToolUse` | After tool succeeds | Format files, run linters, trigger builds |
| `PostToolUseFailure` | After tool fails | Log failures, suggest fixes |
| `UserPromptSubmit` | Before prompt processing | Input validation, context injection |
| `SessionStart` | Session begins | Setup, context loading |
| `Stop` | Claude finishes responding | Verification, cleanup |
| `Notification` | Notification sent | External alerts |
| `SubagentStart` | Subagent launches | Logging |
| `SubagentStop` | Subagent completes | Logging |
| `PreCompact` | Before context compaction | Save important state |

**Hook commands**:
- Exit 0 = allow/success
- Exit 2 = block the operation (PreToolUse only)
- Stdout = message shown to Claude
- Stderr = message shown to user as warning
- Stdin = JSON with tool name, input, and file path

**Critical rules for hooks**:
- Hooks are the highest-leverage extension type. Every repetitive action that doesn't need Claude's reasoning should be a hook.
- Keep commands fast. Hooks run synchronously — a slow hook blocks Claude.
- Use `matcher` to narrow scope. Don't run a Python formatter on Rust files.
- Test hooks independently before integrating (`echo '{"tool": "Write"}' | your-command`).

---

## MCP Server Config

**Use when**: Connecting Claude to an external service, API, or tool that runs as a separate process.

**Context cost**: Low (tool signatures load, but tool content only loads on call).

**Structure**:
```
.mcp.json                 # At plugin root
```

**Anatomy**:
```json
{
  "mcpServers": {
    "<server-name>": {
      "type": "http",
      "url": "https://api.service.com/mcp"
    }
  }
}
```

**Transport types**:
| Type | Use For | Example |
|------|---------|---------|
| `http` | Remote services with MCP endpoint | SaaS APIs |
| `stdio` | Local CLI tools | `npx`, `uvx`, compiled binaries |
| `sse` | Legacy streaming servers | Older MCP implementations |

**stdio example**:
```json
{
  "mcpServers": {
    "my-tool": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "my-mcp-server"],
      "env": {
        "API_KEY": "${MY_API_KEY}"
      }
    }
  }
}
```

**Critical rules for MCP configs**:
- Use environment variable expansion (`${VAR}`) for secrets. Never hardcode keys.
- Prefer `http` transport for remote services (it's the current standard).
- Test that the MCP server starts before including it: `npx -y <package> --help`.

---

## Command

**Use when**: Simple user-triggered action. No multi-step workflow, no file references, no agents.

**Context cost**: Low (single markdown file loads on invocation).

**Structure**:
```
commands/<name>.md        # Single file
```

**Anatomy**:
```yaml
---
description: <shown in autocomplete — keep under 80 chars>
---

# /<name>

<What this command does.>

## Usage

/<name> [arguments]

## Instructions

1. <Step>
2. <Step>
```

**Critical rules for commands**:
- If you need more than ~50 lines of instructions, it should be a skill, not a command.
- If it needs to reference multiple files, it should be a skill.
- Commands are for simple, single-purpose actions.

---

## Choosing Between Types

```
                          Does it need multi-step workflow?
                         /                                \
                       YES                                NO
                       |                                   |
                    SKILL                    Does it run on tool events?
                                            /                          \
                                          YES                          NO
                                           |                            |
                                         HOOK                Does it delegate a task?
                                                            /                        \
                                                          YES                        NO
                                                           |                          |
                                                        AGENT              Does it connect external API?
                                                                          /                            \
                                                                        YES                            NO
                                                                         |                              |
                                                                    MCP CONFIG                      COMMAND
```

Most real extensions are **combinations**. A Zephyr development plugin might have:
- 1 skill (coding standard with rules/)
- 3 commands (/build, /flash, /test)
- 2 hooks (auto-format on write, lint on edit)
- 1 agent (code reviewer specialized for embedded)
- 1 MCP config (debug server connection)

All bundled in one plugin.
