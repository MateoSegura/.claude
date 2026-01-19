# .claude

Shared Claude Code configuration - a scalable knowledge base for AI factory agents.

## Usage

Add as a git submodule to any project:

```bash
git submodule add https://github.com/MateoSegura/.claude.git .claude
```

## Structure

```
.claude/
├── skills/           # Multi-step workflows with rules, scaffolds, and tests
├── rules/            # Simple always-on guidelines (supports nesting)
├── hooks/            # Shell commands triggered by tool events
├── agents/           # Specialized subagents for delegated tasks
├── mcp-servers/      # External API/service integrations
├── commands/         # Slash commands
├── skilltest/        # Shared Go testing framework
├── NAMING.md         # Extension naming convention
├── settings.json     # Global Claude Code settings
└── go.mod            # Go module for tests
```

## Naming Convention

All extensions follow the pattern: `<domain>-<subdomain>-<specific>`

See [NAMING.md](NAMING.md) for the complete taxonomy.

### Domains

| Domain | Description |
|--------|-------------|
| `coding-*` | Programming standards by language/context |
| `design-*` | UI/UX design patterns |
| `docs-*` | Documentation standards |
| `architecture-*` | System architecture patterns |
| `devops-*` | Operations and infrastructure |
| `testing-*` | Testing methodologies |
| `workflow-*` | Development workflows |
| `meta-*` | Claude extensions about Claude |

## Skills

| Skill | Domain | Description |
|-------|--------|-------------|
| `coding-go-cloud` | coding | Go standards for cloud services |
| `coding-typescript-node` | coding | TypeScript for Node.js |
| `coding-react-components` | coding | React component patterns |
| `coding-bash-scripts` | coding | Bash scripting standards |
| `coding-c-zephyr` | coding | C for Zephyr RTOS embedded |
| `design-tui-bubbletea` | design | Bubble Tea TUI development |
| `design-tui-k9s` | design | K9s-style terminal UI |
| `devops-workflow-standard` | devops | Git, CI/CD, code review practices |
| `meta-skills-create` | meta | Creating Claude extensions |
| `meta-skills-update` | meta | Updating Claude extensions |

## Running Tests

Skills include tests to verify Claude follows the patterns correctly:

```bash
# Run all skill tests
SKILL_TEST=1 go test ./...

# Run tests for a specific skill
SKILL_TEST=1 go test ./skills/design-tui-bubbletea/tests/...

# Run with verbose output
SKILL_TEST=1 go test -v ./skills/coding-go-cloud/tests/...
```

Tests require:
- Claude CLI installed
- `ANTHROPIC_API_KEY` environment variable set

Without the API key, tests run in dry-run mode for structure validation.

## Extension Types

| Type | Storage | Discovery | Nesting |
|------|---------|-----------|---------|
| **Skill** | `skills/[name]/SKILL.md` | By folder | Flat only |
| **Rule** | `rules/**/*.md` | Recursive | Supported |
| **Agent** | `agents/[name].md` | By file | Flat only |
| **Command** | `commands/[name].md` | By file | Subdirs for namespacing |
| **Hook** | `settings.json` | Config | N/A |
| **MCP** | `.mcp.json` | Config | N/A |

## License

Private configuration repository.
