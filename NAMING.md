# Extension Naming Convention

This document defines the naming taxonomy for all Claude Code extensions in this knowledge base.

## Constraints

| Extension | Structure | Discovery |
|-----------|-----------|-----------|
| **Skills** | Flat: `skills/[name]/SKILL.md` | By folder name |
| **Agents** | Flat: `agents/[name].md` | By filename |
| **Commands** | Flat-ish: `commands/[name].md` | By filename |
| **Rules** | Nested OK: `rules/[path]/[name].md` | Recursive |
| **Hooks** | Config only: `settings.json` | N/A |
| **MCP** | Config only: `.mcp.json` | N/A |

## Naming Pattern

```
<domain>-<subdomain>-<specifics>
```

### Format Rules

1. **Lowercase with hyphens** - `coding-go-cloud` not `Coding_Go_Cloud`
2. **Domain first** - enables alphabetical grouping
3. **Increasing specificity** - broad → narrow
4. **No abbreviations** - `typescript` not `ts`, `architecture` not `arch`
5. **Max 3 segments** - `domain-subdomain-specific`

---

## Domain Taxonomy

### `coding-*` — Programming Standards

Standards for writing code in specific languages and contexts.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `coding-go-*` | Go language | `coding-go-cloud`, `coding-go-cli` |
| `coding-typescript-*` | TypeScript | `coding-typescript-node`, `coding-typescript-browser` |
| `coding-python-*` | Python | `coding-python-ml`, `coding-python-api` |
| `coding-c-*` | C language | `coding-c-embedded`, `coding-c-zephyr` |
| `coding-rust-*` | Rust language | `coding-rust-embedded`, `coding-rust-wasm` |
| `coding-bash-*` | Shell scripting | `coding-bash-scripts`, `coding-bash-ci` |
| `coding-react-*` | React framework | `coding-react-components`, `coding-react-native` |

### `design-*` — Design Standards

UI/UX design patterns and visual standards.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `design-tui-*` | Terminal UIs | `design-tui-bubbletea`, `design-tui-k9s` |
| `design-web-*` | Web design | `design-web-tailwind`, `design-web-accessibility` |
| `design-mobile-*` | Mobile design | `design-mobile-ios`, `design-mobile-android` |
| `design-system-*` | Design systems | `design-system-tokens`, `design-system-components` |

### `docs-*` — Documentation Standards

Technical writing and documentation patterns.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `docs-api-*` | API docs | `docs-api-openapi`, `docs-api-graphql` |
| `docs-technical-*` | Technical writing | `docs-technical-adr`, `docs-technical-runbooks` |
| `docs-user-*` | User documentation | `docs-user-guides`, `docs-user-tutorials` |
| `docs-code-*` | Code documentation | `docs-code-comments`, `docs-code-readme` |

### `architecture-*` — System Architecture

System design and architectural patterns.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `architecture-cloud-*` | Cloud patterns | `architecture-cloud-aws`, `architecture-cloud-gcp` |
| `architecture-distributed-*` | Distributed systems | `architecture-distributed-events`, `architecture-distributed-cqrs` |
| `architecture-embedded-*` | Embedded systems | `architecture-embedded-rtos`, `architecture-embedded-firmware` |
| `architecture-data-*` | Data architecture | `architecture-data-pipeline`, `architecture-data-lake` |

### `devops-*` — Operations & Infrastructure

CI/CD, infrastructure, and operational patterns.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `devops-cicd-*` | CI/CD pipelines | `devops-cicd-github`, `devops-cicd-gitlab` |
| `devops-infra-*` | Infrastructure | `devops-infra-terraform`, `devops-infra-kubernetes` |
| `devops-monitoring-*` | Observability | `devops-monitoring-prometheus`, `devops-monitoring-datadog` |
| `devops-security-*` | Security ops | `devops-security-scanning`, `devops-security-secrets` |

### `testing-*` — Testing Standards

Testing methodologies and frameworks.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `testing-unit-*` | Unit testing | `testing-unit-go`, `testing-unit-jest` |
| `testing-integration-*` | Integration tests | `testing-integration-api`, `testing-integration-db` |
| `testing-e2e-*` | End-to-end | `testing-e2e-playwright`, `testing-e2e-cypress` |
| `testing-performance-*` | Performance | `testing-performance-load`, `testing-performance-benchmark` |

### `workflow-*` — Development Workflows

Process and workflow automation.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `workflow-git-*` | Git workflows | `workflow-git-conventional`, `workflow-git-trunk` |
| `workflow-review-*` | Code review | `workflow-review-pr`, `workflow-review-security` |
| `workflow-release-*` | Release process | `workflow-release-semver`, `workflow-release-changelog` |

### `meta-*` — Claude Extensions

Extensions about building Claude extensions.

| Subdomain | Description | Examples |
|-----------|-------------|----------|
| `meta-skills-*` | Writing skills | `meta-skills-create`, `meta-skills-update` |
| `meta-agents-*` | Writing agents | `meta-agents-create`, `meta-agents-patterns` |
| `meta-mcp-*` | MCP servers | `meta-mcp-create`, `meta-mcp-patterns` |

---

## Extension-Specific Conventions

### Skills

```
skills/<domain>-<subdomain>-<specific>/
├── SKILL.md           # name: <domain>-<subdomain>-<specific>
├── rules/
├── reference/
├── scaffolds/
└── tests/
```

**Example:** `skills/coding-go-cloud/SKILL.md`
```yaml
---
name: coding-go-cloud
description: Go programming standards for cloud services
---
```

### Agents

```
agents/<domain>-<subdomain>-<capability>.md
```

**Example:** `agents/coding-go-reviewer.md`
```yaml
---
name: coding-go-reviewer
description: Reviews Go code for cloud service patterns
tools: Read, Glob, Grep
---
```

### Rules

Rules can use nested directories for organization:

```
rules/
├── coding/
│   ├── go/
│   │   ├── error-handling.md
│   │   └── concurrency.md
│   └── typescript/
│       └── no-any.md
└── security/
    └── secrets.md
```

Rule files should use descriptive names without domain prefix (folder provides context).

### Commands

```
commands/<domain>-<action>.md
```

**Example:** `commands/coding-lint.md`, `commands/docs-generate.md`

---

## Migration Map

Current → New naming:

| Current | New |
|---------|-----|
| `bubbletea-tui` | `design-tui-bubbletea` |
| `k9s-tui-style` | `design-tui-k9s` |
| `coding-standard-go-cloud` | `coding-go-cloud` |
| `coding-standard-typescript` | `coding-typescript-node` |
| `coding-standard-react` | `coding-react-components` |
| `coding-standard-bash` | `coding-bash-scripts` |
| `coding-standard-c-zephyr` | `coding-c-zephyr` |
| `devops-standard` | `devops-workflow-standard` |
| `writing-claude-extensions` | `meta-skills-create` |
| `updating-claude-extension` | `meta-skills-update` |

---

## Validation Checklist

Before adding a new extension:

- [ ] Uses lowercase with hyphens
- [ ] Follows `domain-subdomain-specific` pattern
- [ ] Domain exists in taxonomy (or needs addition)
- [ ] Name is descriptive without abbreviations
- [ ] Max 3 segments
- [ ] No collision with existing names
- [ ] Folder/file matches `name` field in frontmatter
