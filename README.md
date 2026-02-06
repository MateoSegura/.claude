# Mateo's Claude Code Marketplace

A curated collection of Claude Code plugins for embedded development, documentation generation, and developer productivity.

## Install

```bash
/plugin marketplace add MateoSegura/.claude
```

Then install any plugin:

```bash
/plugin install <plugin-name>@mateo-marketplace
```

## Plugins

| Plugin | Category | Description |
|--------|----------|-------------|
| `coding-standard-c-zephyr` | Development | Complete C/Zephyr RTOS coding standard with 50+ rules, clang-format/clang-tidy/cppcheck configs, and review checklists |
| `docs-diagrams-code` | Documentation | Generate reproducible Mermaid diagrams (architecture, sequence, dependency, type) from Go source code |
| `docs-presentations-business` | Documentation | Generate branded Slidev decks from business strategy docs. 14 slide templates + full design system |
| `docs-presentations-code` | Documentation | Generate Slidev presentations for libraries/tools. Three-phase pipeline with parallel agents |
| `extension-toolkit` | Productivity | Create and manage Claude Code extensions: skills, rules, commands, agents, hooks |

## Structure

```
.claude-plugin/
  marketplace.json          # Marketplace catalog
plugins/
  coding-standard-c-zephyr/ # C/Zephyr RTOS coding standard
  docs-diagrams-code/       # Mermaid diagram generation
  docs-presentations-business/ # Business presentation generator
  docs-presentations-code/  # Code presentation generator
  extension-toolkit/        # Extension creation & management
```

## Contributing

Each plugin lives in `plugins/<name>/` with:
- `.claude-plugin/plugin.json` - Plugin manifest
- `skills/<name>/SKILL.md` - Skills (multi-step workflows)
- `commands/<name>.md` - Slash commands
- `hooks/hooks.json` - Event hooks (optional)
- `agents/<name>.md` - Custom subagents (optional)
- `.mcp.json` - MCP server configs (optional)

## License

MIT
