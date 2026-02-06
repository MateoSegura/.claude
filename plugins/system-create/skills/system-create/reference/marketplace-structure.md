# Marketplace Structure

Loaded during Phase 4 (Implementation). How to package extensions into plugins and marketplaces.

---

## Plugin File Layout

Every plugin lives in its own directory with this structure:

```
plugins/<plugin-name>/
├── .claude-plugin/
│   └── plugin.json           # REQUIRED — plugin manifest
├── skills/
│   └── <skill-name>/
│       ├── SKILL.md           # Skill entry point
│       ├── rules/             # Constraints (optional)
│       ├── reference/         # Lookups (optional)
│       ├── scaffolds/         # Templates (optional)
│       └── examples/          # Examples (optional)
├── commands/
│   └── <command-name>.md      # Slash commands (optional)
├── agents/
│   └── <agent-name>.md        # Custom subagents (optional)
├── hooks/
│   └── hooks.json             # Event hooks (optional)
└── .mcp.json                  # MCP server configs (optional)
```

## plugin.json

```json
{
  "name": "<plugin-name>",
  "description": "<what this plugin does>",
  "version": "<semver>",
  "author": {
    "name": "<author name>",
    "email": "<optional email>"
  },
  "license": "MIT"
}
```

- `name` must match the directory name
- `description` should be a single sentence
- `version` follows semver (1.0.0, 1.1.0, 2.0.0)

## marketplace.json

The catalog file lives at `.claude-plugin/marketplace.json` in the marketplace repo root.

```json
{
  "$schema": "https://anthropic.com/claude-code/marketplace.schema.json",
  "name": "<marketplace-name>",
  "description": "<what this marketplace offers>",
  "owner": {
    "name": "<owner>",
    "email": "<email>"
  },
  "plugins": [
    {
      "name": "<plugin-name>",
      "description": "<description>",
      "version": "<semver>",
      "author": { "name": "<author>" },
      "source": "./plugins/<plugin-name>",
      "category": "<category>",
      "keywords": ["tag1", "tag2"]
    }
  ]
}
```

**Categories**: development, documentation, productivity, testing, security, monitoring, deployment, database, design, learning

**Source types**:
- Local: `"source": "./plugins/<name>"` (bundled in marketplace)
- GitHub: `"source": {"source": "github", "repo": "owner/repo", "ref": "v1.0"}`
- Git URL: `"source": {"source": "url", "url": "https://host/repo.git"}`

## Naming Conventions

### Plugin names
Format: `<domain>-<function>` or `<domain>-<phase>`

Examples:
- `zephyr-dev` (domain: zephyr, phase: development)
- `k8s-debug` (domain: kubernetes, phase: debugging)
- `go-api` (domain: go, function: API development)
- `system-create` (domain: system, function: creation)

### Skill names
Match the plugin name or use: `<plugin>-<specific>`

### Agent names
Format: `system-<role>` or `<plugin>-<role>`

Examples:
- `system-auditor` (system-level, audits extensions)
- `system-reviewer` (system-level, reviews quality)
- `zephyr-fault-analyzer` (zephyr-specific, analyzes faults)

### Command names
Format: `<verb>-<noun>` or `<noun>` for simple actions

Examples:
- `build` (simple action)
- `flash-device` (verb-noun)
- `new-module` (verb-noun)

## Taxonomy

Organize plugins by workflow phase, not by file type:

```
GOOD (by phase):                  BAD (by type):
├── zephyr-dev/                   ├── all-rules/
├── zephyr-debug/                 ├── all-scaffolds/
├── zephyr-test/                  ├── all-commands/
└── zephyr-docs/                  └── all-agents/
```

Phase-based organization means users install what they need:
- Always: `zephyr-dev`
- When debugging: `zephyr-debug`
- When testing: `zephyr-test`

## Placement Rules

| Target | Where to Write | When |
|--------|---------------|------|
| Marketplace plugin | `plugins/<name>/` in marketplace repo | Creating a distributable extension |
| Project-local skill | `.claude/skills/<name>/` | Extension only for this project |
| Personal global skill | `~/.claude/skills/<name>/` | Extension for all your projects |
| Project-local hook | `.claude/settings.json` → hooks | Hook only for this project |
| Personal global hook | `~/.claude/settings.json` → hooks | Hook for all your projects |
