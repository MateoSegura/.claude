---
name: meta-skill-update
description: Update existing Claude Code extensions
---

# Updating Extensions

This skill helps discover and update existing extensions.

## Quick Start

```
/update-extension [name]
```

Or run `/list-extensions` first to see what's available.

## Workflow

```
1. DISCOVER    → Find all extensions
2. SELECT      → Choose which to update
3. ANALYZE     → Read and understand current state
4. RESEARCH    → Web search for latest patterns
5. PLAN        → Propose changes
6. IMPLEMENT   → Make changes with approval
7. TEST        → Verify nothing broke
```

## Discovery

| Type | Location | Command |
|------|----------|---------|
| Skills | `skills/*/SKILL.md` | `ls skills/` |
| Rules | `rules/**/*.md` | `find rules -name "*.md"` |
| Commands | `commands/*.md` | `ls commands/` |
| Agents | `agents/*.md` | `ls agents/` |
| Hooks | `settings.json` | Read hooks section |
| MCPs | `.mcp.json` | Read mcpServers section |

## Update Scenarios

### Library/Framework Update

1. Check version in skill vs installed version
2. Web search: `"{library} changelog {version}"`
3. Identify breaking changes
4. Update affected patterns and examples
5. Test with new version

### Add Missing Patterns

1. Search codebase for patterns not covered
2. Web search: `"{technology} best practices 2024"`
3. Add new rules/scaffolds
4. Update quick reference

### Fix Issues

1. Identify what's wrong
2. Find correct pattern (docs, working code)
3. Update and verify

### Expand Scope

1. Identify gaps
2. Research new areas
3. Add sections while maintaining consistency

## Research Queries

| Extension Type | Search Queries |
|----------------|----------------|
| Language skill | `"{lang} best practices"`, `"{lang} style guide"` |
| Framework skill | `"{framework} patterns"`, `"{framework} v{version} migration"` |
| Tool skill | `"{tool} advanced usage"`, `"{tool} tips"` |
| Practice skill | `"{practice} checklist"`, `"{practice} examples"` |

## Validation

After update:
- [ ] All existing tests still pass
- [ ] New functionality has tests
- [ ] No regressions
- [ ] Documentation updated
- [ ] Version bumped if appropriate

## Self-Maintenance

This repo can maintain itself via CI:

```yaml
# Weekly update check
- name: Check for updates
  run: |
    claude --skill meta-skill-update \
      --prompt "Check all extensions for outdated patterns"
```

The CI workflow:
1. Scans all extensions
2. Web searches for updates
3. Proposes changes via PR
4. Runs tests to verify
