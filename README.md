# .claude

Scalable knowledge base for AI factory agents - a shared Claude Code configuration.

## Quick Start

Add as a git submodule to any project:

```bash
git submodule add https://github.com/MateoSegura/.claude.git .claude
```

**New here?** Start with [GUIDE.md](GUIDE.md) for loading instructions and composition examples.

## Structure

```
.claude/
├── GUIDE.md              # Start here - loading instructions for LLMs
├── NAMING.md             # Naming convention reference
├── README.md             # This file
├── skills/               # Multi-step workflows with rules, scaffolds, and tests
│   ├── language-*/       # Programming language foundations
│   ├── framework-*/      # Framework-specific patterns
│   ├── platform-*/       # Platform/infrastructure knowledge
│   ├── tool-*/           # Tool-specific workflows
│   ├── practice-*/       # Cross-cutting practices
│   ├── domain-*/         # Business domain knowledge
│   ├── role-*/           # Role behavioral patterns
│   └── meta-*/           # Extension management
├── rules/                # Simple always-on guidelines (supports nesting)
├── agents/               # Pre-composed agent configurations
├── commands/             # User-invocable slash commands
├── hooks/                # Shell commands triggered by tool events
├── mcp-servers/          # External API/service integrations
├── skilltest/            # Shared Go testing framework
├── settings.json         # Global Claude Code settings
└── go.mod                # Go module for tests
```

## Two-Layer Taxonomy

Extensions use a Foundation + Role architecture. See [NAMING.md](NAMING.md) for details.

### Foundation Layer (What to Know)

| Prefix | Purpose | Examples |
|--------|---------|----------|
| `language-` | Programming language standards | `language-go-cloud`, `language-typescript-node` |
| `framework-` | Framework-specific patterns | `framework-react`, `framework-bubbletea` |
| `platform-` | Platform/infrastructure | `platform-kubernetes`, `platform-zephyr` |
| `tool-` | Tool-specific workflows | `tool-git-workflow`, `tool-docker` |
| `practice-` | Cross-cutting concerns | `practice-security-auth`, `practice-testing-tdd` |
| `domain-` | Business domain knowledge | `domain-fintech`, `domain-healthcare` |

### Role Layer (How to Behave)

| Prefix | Role | Examples |
|--------|------|----------|
| `role-manager-` | Task decomposition, orchestration | `role-manager-planning` |
| `role-architect-` | System design, decisions | `role-architect-api` |
| `role-developer-` | Code implementation | `role-developer-backend` |
| `role-reviewer-` | Code review, audits | `role-reviewer-security` |
| `role-tester-` | Testing strategies | `role-tester-e2e` |
| `role-devops-` | CI/CD, deployment | `role-devops-pipeline` |
| `role-sre-` | Operations, incidents | `role-sre-debugging` |
| `role-writer-` | Documentation | `role-writer-technical` |
| `role-pm-` | Requirements, prioritization | `role-pm-requirements` |
| `role-designer-` | UI/UX design | `role-designer-ui` |

### Meta Layer

| Prefix | Purpose | Examples |
|--------|---------|----------|
| `meta-` | Extension management | `meta-skill-create`, `meta-skill-update` |

## Current Skills

| Skill | Layer | Description |
|-------|-------|-------------|
| `language-go-cloud` | Foundation | Go standards for cloud services |
| `language-typescript-node` | Foundation | TypeScript for Node.js |
| `language-bash` | Foundation | Bash scripting standards |
| `language-c-zephyr` | Foundation | C for Zephyr RTOS embedded |
| `framework-react` | Foundation | React component patterns |
| `framework-bubbletea` | Foundation | Bubble Tea TUI development |
| `framework-k9s-style` | Foundation | K9s-style terminal UI |
| `tool-git-workflow` | Foundation | Git, CI/CD, code review practices |
| `meta-skill-create` | Meta | Creating Claude extensions |
| `meta-skill-update` | Meta | Updating Claude extensions |

## Running Tests

Skills include tests to verify Claude follows the patterns correctly:

```bash
# Run all skill tests
SKILL_TEST=1 go test ./...

# Run tests for a specific skill
SKILL_TEST=1 go test ./skills/framework-bubbletea/tests/...

# Run with verbose output
SKILL_TEST=1 go test -v ./skills/language-go-cloud/tests/...
```

Tests require:
- Claude CLI installed
- `ANTHROPIC_API_KEY` environment variable set

Without the API key, tests run in dry-run mode for structure validation.

## Extension Types

| Type | Storage | Discovery | Nesting |
|------|---------|-----------|---------|
| **Skill** | `skills/{name}/SKILL.md` | By folder | Flat only |
| **Rule** | `rules/**/*.md` | Recursive | Supported |
| **Agent** | `agents/{name}.md` | By file | Flat only |
| **Command** | `commands/{name}.md` | By file | Subdirs for namespacing |
| **Hook** | `settings.json` | Config | N/A |
| **MCP** | `.mcp.json` | Config | N/A |

## License

Private configuration repository.
