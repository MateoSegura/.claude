# Claude Extensions Guide

This is your entry point for understanding and loading extensions from this knowledge base.

## Quick Start

**To configure an agent**, combine:
1. **Foundation skills** - Technical knowledge (languages, frameworks, platforms, practices)
2. **Role skills** - Behavioral patterns (how to act as developer, reviewer, architect, etc.)

Example: A Go backend developer working on cloud services would load:
- `language-go-cloud` (Go patterns for cloud)
- `framework-gin` (if using Gin)
- `platform-kubernetes` (if deploying to k8s)
- `practice-security-auth` (if handling auth)
- `role-developer-backend` (backend developer behaviors)

## Taxonomy

### Foundation Layer

These provide **knowledge** that can be shared across roles.

#### Technical Foundations

| Prefix | What it provides | Examples |
|--------|------------------|----------|
| `language-{lang}-{context}` | Language standards and idioms | `language-go-cloud`, `language-typescript-node`, `language-c-embedded` |
| `framework-{name}` | Framework-specific patterns | `framework-react`, `framework-bubbletea`, `framework-gin` |
| `platform-{name}` | Platform/infrastructure knowledge | `platform-kubernetes`, `platform-aws`, `platform-zephyr` |
| `tool-{name}` | Tool-specific workflows | `tool-git`, `tool-docker`, `tool-terraform` |

#### Practice Foundations

Cross-cutting concerns that apply regardless of stack:

| Prefix | What it provides | Examples |
|--------|------------------|----------|
| `practice-security-{area}` | Security patterns | `practice-security-auth`, `practice-security-input-validation`, `practice-security-secrets` |
| `practice-documentation-{type}` | Documentation standards | `practice-documentation-api`, `practice-documentation-architecture`, `practice-documentation-runbook` |
| `practice-testing-{method}` | Testing methodologies | `practice-testing-tdd`, `practice-testing-property-based`, `practice-testing-mutation` |
| `practice-performance-{area}` | Performance patterns | `practice-performance-profiling`, `practice-performance-caching`, `practice-performance-database` |
| `practice-accessibility-{platform}` | Accessibility standards | `practice-accessibility-web`, `practice-accessibility-mobile` |
| `practice-observability-{aspect}` | Logging/metrics/tracing | `practice-observability-logging`, `practice-observability-metrics`, `practice-observability-tracing` |

#### Domain Foundations

Industry and business domain expertise:

| Prefix | What it provides | Examples |
|--------|------------------|----------|
| `domain-{industry}` | Domain-specific knowledge | `domain-fintech`, `domain-healthcare`, `domain-iot`, `domain-gaming` |

### Role Layer

These define **behaviors** - how an agent should act when performing a role.

| Prefix | Role | Responsibilities |
|--------|------|------------------|
| `role-manager-{area}` | Manager | Task decomposition, orchestration, long-running coordination, delegation |
| `role-architect-{area}` | Architect | System design, technical decisions, API design, data modeling |
| `role-developer-{area}` | Developer | Code implementation, following standards, writing clean code |
| `role-reviewer-{area}` | Reviewer | Code review, security audits, quality gates |
| `role-tester-{area}` | Tester | Test strategy, test implementation, coverage analysis |
| `role-devops-{area}` | DevOps | CI/CD pipelines, deployment, infrastructure as code |
| `role-sre-{area}` | SRE | Monitoring, debugging, incident response, reliability |
| `role-writer-{area}` | Writer | Technical documentation, specs, guides, tutorials |
| `role-pm-{area}` | Product Manager | Requirements gathering, prioritization, roadmapping |
| `role-designer-{area}` | Designer | UI/UX design, prototyping, design systems |

### Meta Layer

Skills for managing this extension system itself:

| Prefix | Purpose | Examples |
|--------|---------|----------|
| `meta-skill-{action}` | Skill management | `meta-skill-create`, `meta-skill-update` |
| `meta-agent-{action}` | Agent management | `meta-agent-compose`, `meta-agent-test` |

## Loading Strategy

### For an LLM Self-Selecting Skills

When given a task, parse for:

1. **Technical signals** → Load matching `language-`, `framework-`, `platform-` skills
   - "Go API" → `language-go-cloud`
   - "React component" → `framework-react`
   - "deploy to Kubernetes" → `platform-kubernetes`

2. **Practice signals** → Load matching `practice-` skills
   - "secure", "auth" → `practice-security-*`
   - "document", "API docs" → `practice-documentation-*`
   - "test", "TDD" → `practice-testing-*`

3. **Role signals** → Load matching `role-` skills
   - "implement", "build", "create" → `role-developer-*`
   - "review", "audit" → `role-reviewer-*`
   - "design system", "architect" → `role-architect-*`
   - "break down", "plan", "decompose" → `role-manager-*`
   - "debug", "incident" → `role-sre-*`

### Composition Examples

**Task: "Build a REST API in Go with authentication for Kubernetes"**
```
language-go-cloud
framework-gin (or framework-echo)
platform-kubernetes
practice-security-auth
role-developer-backend
```

**Task: "Review this PR for security issues"**
```
practice-security-auth
practice-security-input-validation
role-reviewer-security
```

**Task: "Break down this feature into implementable tasks"**
```
role-manager-decomposition
role-manager-planning
```

**Task: "Debug why pods are crashing in production"**
```
platform-kubernetes
practice-observability-logging
role-sre-debugging
```

**Task: "Write API documentation for our service"**
```
practice-documentation-api
role-writer-technical
```

## Directory Structure

```
.claude/
├── GUIDE.md              # This file - start here
├── NAMING.md             # Detailed naming conventions
├── README.md             # Repository overview
├── skills/               # Complex multi-step workflows
│   ├── language-*/       # Language foundations
│   ├── framework-*/      # Framework foundations
│   ├── platform-*/       # Platform foundations
│   ├── tool-*/           # Tool foundations
│   ├── practice-*/       # Practice foundations
│   ├── domain-*/         # Domain foundations
│   ├── role-*/           # Role behaviors
│   └── meta-*/           # Meta skills
├── rules/                # Simple always-on guidelines (supports nesting)
├── agents/               # Pre-composed agent configurations
├── commands/             # User-invocable slash commands
├── hooks/                # Tool event triggers
└── mcp-servers/          # External service integrations
```

## Extension Types Reference

| Type | Location | Discovery | Use Case |
|------|----------|-----------|----------|
| **Skill** | `skills/{name}/SKILL.md` | Flat only | Complex workflows with rules and references |
| **Rule** | `rules/**/*.md` | Recursive | Simple always-on guidelines |
| **Agent** | `agents/{name}.md` | Flat only | Pre-composed configurations |
| **Command** | `commands/{name}.md` | Flat | User-invocable actions |
| **Hook** | `settings.json` | Config | Automated tool triggers |
| **MCP** | `.mcp.json` | Config | External service access |

## Best Practices

1. **Start with foundations** - Load language/framework skills before role skills
2. **Layer practices as needed** - Only load practice skills relevant to the task
3. **One primary role** - Usually one role-* skill drives behavior, others supplement
4. **Domain when relevant** - Load domain-* skills for industry-specific work
5. **Check GUIDE.md first** - This file is your index to everything available
