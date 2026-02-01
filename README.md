# .claude

Claude Code configuration for enhanced AI-assisted development.

## Structure

```
.claude/
├── commands/                       # User-invocable slash commands
│   ├── docs-diagrams-code.md
│   ├── list-extensions.md
│   ├── new-agent.md
│   ├── new-command.md
│   ├── new-hook.md
│   ├── new-rule.md
│   ├── new-skill.md
│   └── update-extension.md
├── skills/                         # Multi-step workflow instructions
│   ├── coding-standard-c-zephyr/
│   ├── docs-presentations-code/
│   ├── meta-skill-create/
│   └── meta-skill-update/
├── settings.json                   # Claude Code settings
```

## Commands

| Command | Purpose |
|---------|---------|
| `/docs-diagrams-code` | Generate Mermaid diagrams from Go source code |
| `/list-extensions` | List all extensions |
| `/new-agent` | Create a specialized subagent |
| `/new-command` | Create a user-invocable command |
| `/new-hook` | Create a tool event hook |
| `/new-rule` | Create a simple always-on rule |
| `/new-skill` | Create a new skill extension |
| `/update-extension` | Update an existing extension |

## Skills

### coding-standard-c-zephyr
Complete C/Zephyr RTOS coding standard with rules for naming, types, memory, concurrency, error handling, security, and more. Includes tooling configs (clang-format, clang-tidy, cppcheck) and a quick-reference checklist.

### docs-presentations-code
Three-phase pipeline (Discover, Propose, Generate) for creating branded Slidev presentations from codebases. Produces purpose-driven decks following a Five Acts structure with dark-theme Mermaid diagrams.

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
