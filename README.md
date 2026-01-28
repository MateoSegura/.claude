# .claude

Claude Code configuration for enhanced AI-assisted development.

## Structure

```
.claude/
├── commands/           # User-invocable slash commands
│   ├── new-skill.md
│   ├── new-rule.md
│   ├── new-command.md
│   ├── new-agent.md
│   ├── new-hook.md
│   ├── list-extensions.md
│   └── update-extension.md
├── skills/             # Multi-step workflow instructions
│   ├── meta-skill-create/
│   └── meta-skill-update/
└── settings.json       # Claude Code settings
```

## Commands

| Command | Purpose |
|---------|---------|
| `/new-skill` | Create a new skill extension |
| `/new-rule` | Create a simple always-on rule |
| `/new-command` | Create a user-invocable command |
| `/new-agent` | Create a specialized subagent |
| `/new-hook` | Create a tool event hook |
| `/list-extensions` | List all extensions |
| `/update-extension` | Update an existing extension |

## Skills

### meta-skill-create
Guides Claude through selecting the right extension type:
- SKILL: Multi-step workflows
- RULE: Simple always-on constraints
- HOOK: Automated tool event responses
- MCP: External service integration
- AGENT: Specialized subagents

### meta-skill-update
Guides Claude through updating existing extensions when libraries or requirements change.

## Testing

See [.claude-test](https://github.com/MateoSegura/.claude-test) for the testing framework.

## Usage

This directory can be placed in:
- `~/.claude/` - Global config for all projects
- `<project>/.claude/` - Project-specific config

## License

MIT
