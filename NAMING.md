# Extension Naming Convention

This document defines the naming taxonomy for all Claude Code extensions in this knowledge base.

> **Quick Reference:** See [GUIDE.md](GUIDE.md) for loading instructions and composition examples.

## Constraints

| Extension | Structure | Discovery |
|-----------|-----------|-----------|
| **Skills** | Flat: `skills/{name}/SKILL.md` | By folder name |
| **Agents** | Flat: `agents/{name}.md` | By filename |
| **Commands** | Flat-ish: `commands/{name}.md` | By filename |
| **Rules** | Nested OK: `rules/{path}/{name}.md` | Recursive |
| **Hooks** | Config only: `settings.json` | N/A |
| **MCP** | Config only: `.mcp.json` | N/A |

## Naming Philosophy

Names are designed for **LLM parseability** - an agent should be able to self-select relevant skills based on task context without explicit mappings.

### Format Rules

1. **Lowercase with hyphens** - `language-go-cloud` not `Language_Go_Cloud`
2. **Layer prefix first** - enables categorical discovery
3. **Increasing specificity** - broad → narrow
4. **No abbreviations** - `typescript` not `ts`, `security` not `sec`
5. **Descriptive over brief** - clarity for machine parsing

---

## Two-Layer Taxonomy

### Layer 1: Foundations

Shared knowledge that multiple roles can load. These provide **what to know**.

#### Technical Foundations

| Prefix | Purpose | Pattern | Examples |
|--------|---------|---------|----------|
| `language-` | Programming language standards | `language-{lang}-{context}` | `language-go-cloud`, `language-typescript-node`, `language-c-embedded` |
| `framework-` | Framework-specific patterns | `framework-{name}` | `framework-react`, `framework-bubbletea`, `framework-gin`, `framework-nextjs` |
| `platform-` | Platform/infrastructure | `platform-{name}` | `platform-kubernetes`, `platform-aws`, `platform-zephyr`, `platform-vercel` |
| `tool-` | Tool-specific workflows | `tool-{name}` | `tool-git`, `tool-docker`, `tool-terraform`, `tool-nix` |

#### Practice Foundations

Cross-cutting concerns that apply regardless of technology stack:

| Prefix | Purpose | Pattern | Examples |
|--------|---------|---------|----------|
| `practice-security-` | Security patterns | `practice-security-{area}` | `practice-security-auth`, `practice-security-input-validation`, `practice-security-secrets` |
| `practice-documentation-` | Documentation standards | `practice-documentation-{type}` | `practice-documentation-api`, `practice-documentation-architecture`, `practice-documentation-runbook` |
| `practice-testing-` | Testing methodologies | `practice-testing-{method}` | `practice-testing-tdd`, `practice-testing-property-based`, `practice-testing-mutation` |
| `practice-performance-` | Performance patterns | `practice-performance-{area}` | `practice-performance-profiling`, `practice-performance-caching`, `practice-performance-database` |
| `practice-accessibility-` | Accessibility standards | `practice-accessibility-{platform}` | `practice-accessibility-web`, `practice-accessibility-mobile` |
| `practice-observability-` | Logging/metrics/tracing | `practice-observability-{aspect}` | `practice-observability-logging`, `practice-observability-metrics`, `practice-observability-tracing` |

#### Domain Foundations

Industry and business domain expertise:

| Prefix | Purpose | Pattern | Examples |
|--------|---------|---------|----------|
| `domain-` | Domain-specific knowledge | `domain-{industry}` | `domain-fintech`, `domain-healthcare`, `domain-iot`, `domain-gaming`, `domain-ecommerce` |

### Layer 2: Roles

Behavioral overlays that define how an agent acts. These provide **how to behave**.

| Prefix | Role | Responsibilities | Examples |
|--------|------|------------------|----------|
| `role-manager-` | Manager | Task decomposition, orchestration, long-running coordination, delegation | `role-manager-planning`, `role-manager-decomposition`, `role-manager-delegation` |
| `role-architect-` | Architect | System design, technical decisions, API design, data modeling | `role-architect-system`, `role-architect-api`, `role-architect-data` |
| `role-developer-` | Developer | Code implementation, following standards, writing clean code | `role-developer-backend`, `role-developer-frontend`, `role-developer-fullstack` |
| `role-reviewer-` | Reviewer | Code review, security audits, quality gates | `role-reviewer-code`, `role-reviewer-security`, `role-reviewer-architecture` |
| `role-tester-` | Tester | Test strategy, test implementation, coverage analysis | `role-tester-unit`, `role-tester-integration`, `role-tester-e2e` |
| `role-devops-` | DevOps | CI/CD pipelines, deployment, infrastructure as code | `role-devops-pipeline`, `role-devops-release`, `role-devops-infrastructure` |
| `role-sre-` | SRE | Monitoring, debugging, incident response, reliability | `role-sre-debugging`, `role-sre-monitoring`, `role-sre-incident` |
| `role-writer-` | Writer | Technical documentation, specs, guides, tutorials | `role-writer-technical`, `role-writer-api`, `role-writer-tutorial` |
| `role-pm-` | Product Manager | Requirements gathering, prioritization, roadmapping | `role-pm-requirements`, `role-pm-prioritization`, `role-pm-roadmap` |
| `role-designer-` | Designer | UI/UX design, prototyping, design systems | `role-designer-ui`, `role-designer-ux`, `role-designer-system` |

### Meta Layer

Skills for managing the extension system itself:

| Prefix | Purpose | Examples |
|--------|---------|----------|
| `meta-skill-` | Skill management | `meta-skill-create`, `meta-skill-update` |
| `meta-agent-` | Agent configuration | `meta-agent-compose`, `meta-agent-test` |
| `meta-rule-` | Rule management | `meta-rule-create` |

---

## Extension-Specific Conventions

### Skills

```
skills/{prefix}-{name}/
├── SKILL.md           # name: {prefix}-{name}
├── rules/             # Skill-specific rules
├── reference/         # Reference materials
├── scaffolds/         # Code templates
└── tests/             # Skill tests
```

**Example:** `skills/language-go-cloud/SKILL.md`
```yaml
---
name: language-go-cloud
description: Go programming standards for cloud services
---
```

### Agents

Pre-composed configurations that combine multiple skills:

```
agents/{role}-{specialization}.md
```

**Example:** `agents/developer-go-backend.md`
```yaml
---
name: developer-go-backend
description: Go backend developer agent
skills:
  - language-go-cloud
  - framework-gin
  - platform-kubernetes
  - practice-security-auth
  - role-developer-backend
---
```

### Rules

Rules can use nested directories (hierarchy in folders, not names):

```
rules/
├── security/
│   ├── input-validation.md
│   └── secrets.md
├── coding/
│   ├── go/
│   │   └── error-handling.md
│   └── typescript/
│       └── strict-types.md
└── workflow/
    └── commit-messages.md
```

### Commands

```
commands/{action}-{target}.md
```

**Example:** `commands/generate-api-docs.md`, `commands/review-security.md`

---

## Migration Map

Previous → New naming:

| Previous | New | Layer |
|----------|-----|-------|
| `coding-go-cloud` | `language-go-cloud` | Foundation |
| `coding-typescript-node` | `language-typescript-node` | Foundation |
| `coding-react-components` | `framework-react` | Foundation |
| `coding-bash-scripts` | `language-bash` | Foundation |
| `coding-c-zephyr` | `language-c-zephyr` | Foundation |
| `design-tui-bubbletea` | `framework-bubbletea` | Foundation |
| `design-tui-k9s` | `framework-k9s-style` | Foundation |
| `devops-workflow-standard` | `tool-git-workflow` | Foundation |
| `meta-skills-create` | `meta-skill-create` | Meta |
| `meta-skills-update` | `meta-skill-update` | Meta |

---

## Validation Checklist

Before adding a new extension:

- [ ] Uses lowercase with hyphens
- [ ] Has correct layer prefix (`language-`, `framework-`, `platform-`, `tool-`, `practice-`, `domain-`, `role-`, `meta-`)
- [ ] Descriptive without abbreviations
- [ ] No collision with existing names
- [ ] Folder/file matches `name` field in frontmatter
- [ ] Added to GUIDE.md if new category

---

## Quick Reference

**Foundation prefixes (shared knowledge):**
- `language-{lang}-{context}` - Programming languages
- `framework-{name}` - Frameworks and libraries
- `platform-{name}` - Platforms and infrastructure
- `tool-{name}` - Development tools
- `practice-{area}-{specific}` - Cross-cutting practices
- `domain-{industry}` - Business domains

**Role prefixes (behavioral):**
- `role-manager-{area}` - Orchestration and planning
- `role-architect-{area}` - System design
- `role-developer-{area}` - Implementation
- `role-reviewer-{area}` - Code review
- `role-tester-{area}` - Testing
- `role-devops-{area}` - CI/CD and deployment
- `role-sre-{area}` - Operations
- `role-writer-{area}` - Documentation
- `role-pm-{area}` - Product management
- `role-designer-{area}` - UI/UX design

**Meta prefix:**
- `meta-{type}-{action}` - Extension management
