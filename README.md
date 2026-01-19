# .claude

Shared Claude Code configuration - skills, rules, hooks, and more.

## Usage

Add as a git submodule to any project:

```bash
git submodule add https://github.com/MateoSegura/.claude.git .claude
```

## Structure

```
.claude/
├── skills/           # Multi-step workflows with rules, scaffolds, and tests
├── rules/            # Simple always-on guidelines
├── hooks/            # Shell commands triggered by tool events
├── agents/           # Specialized subagents for delegated tasks
├── mcp-servers/      # External API/service integrations
├── commands/         # Slash commands
├── skilltest/        # Shared testing framework for skills
├── settings.json     # Global Claude Code settings
└── go.mod           # Go module for tests
```

## Skills

Skills are comprehensive guides that teach Claude specific patterns and practices:

| Skill | Description |
|-------|-------------|
| `bubbletea-tui` | Bubble Tea TUI development patterns |
| `k9s-tui-style` | K9s-style terminal UI design |
| `coding-standard-bash` | Bash scripting standards |
| `coding-standard-c-zephyr` | C/Zephyr embedded development |
| `coding-standard-go-cloud` | Go cloud services standards |
| `coding-standard-react` | React component standards |
| `coding-standard-typescript` | TypeScript coding standards |
| `devops-standard` | DevOps practices and CI/CD |
| `writing-claude-extensions` | Creating Claude Code extensions |
| `updating-claude-extension` | Updating existing extensions |

## Running Tests

Skills include tests to verify Claude follows the patterns correctly:

```bash
# Run all skill tests
SKILL_TEST=1 go test ./...

# Run tests for a specific skill
SKILL_TEST=1 go test ./skills/bubbletea-tui/tests/...

# Run with verbose output
SKILL_TEST=1 go test -v ./skills/k9s-tui-style/tests/...
```

Tests require:
- Claude CLI installed
- `ANTHROPIC_API_KEY` environment variable set

Without the API key, tests run in dry-run mode for structure validation.

## Extension Types

| Type | Use Case |
|------|----------|
| **Skill** | Complex multi-step workflows with rules, examples, and scaffolds |
| **Rule** | Simple always-on constraints (e.g., "never use `any` type") |
| **Hook** | Automated actions after tool use (e.g., run formatter after write) |
| **Agent** | Specialized sub-agents for delegated tasks |
| **MCP Server** | External API/database integrations |
| **Command** | Slash commands for quick actions |

## License

Private configuration repository.
