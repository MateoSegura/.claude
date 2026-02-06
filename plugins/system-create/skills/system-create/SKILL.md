---
name: system-create
description: Create production-grade Claude Code extensions (skills, agents, hooks, MCP configs, commands) through guided architecture, research, implementation, and adversarial review. Invoke when the user wants to build a new skill, agent, hook, MCP integration, or command for their Claude Code setup or marketplace.
user-invocable: true
argument-hint: "[description of what you want to build]"
---

# System Create

Build Claude Code extensions that actually work. Not templates -- engineered systems with proper architecture, lazy context loading, and verified prompt quality.

## When to Use

- User wants to create a new skill, agent, hook, MCP config, or command
- User describes a workflow they want automated
- User asks "can Claude do X?" and the answer requires a new extension
- User wants to improve an existing extension

## Files (load per-phase, never all at once)

| File | Purpose | Load When |
|------|---------|-----------|
| [reference/extension-anatomy.md](reference/extension-anatomy.md) | Structure and file layout for every extension type | Phase 2 (Architecture) |
| [reference/marketplace-structure.md](reference/marketplace-structure.md) | Plugin packaging, marketplace.json, naming conventions | Phase 4 (Implementation) |
| [reference/quality-standards.md](reference/quality-standards.md) | What separates good extensions from garbage | Phase 5 (Review) |
| [reference/prompt-engineering.md](reference/prompt-engineering.md) | Techniques that increase Claude's success rate | Phase 4-5 (Implementation + Review) |

## Workflow

### Phase 1: AUDIT

**Goal**: Determine if this extension should exist.

Launch the `system-auditor` agent (Task tool, `subagent_type: general-purpose`) with the user's request. The agent:

1. Scans the current project's `.claude/` directory (if it exists)
2. Scans `~/.claude/` for globally installed skills, commands, hooks
3. Checks installed marketplace plugins via `~/.claude/plugins/installed_plugins.json`
4. Checks the official marketplace catalog via `~/.claude/plugins/marketplaces/*/`
5. Evaluates whether the user's request is already covered

The agent returns one of:
- **EXISTS** — name the existing extension and explain how to use it. Ask the user if they want to extend it instead.
- **PARTIAL** — something similar exists but doesn't fully cover the request. Recommend extending vs. creating new.
- **NEW** — nothing covers this. Proceed to Phase 2.

If EXISTS or PARTIAL, present findings to the user with `AskUserQuestion`. Only proceed if the user confirms they want a new extension.

### Phase 2: DESIGN (Conversational)

**Goal**: Architect the extension through conversation. Do not rush this phase.

Read [reference/extension-anatomy.md](reference/extension-anatomy.md) now.

Work with the user to determine:

**2.1 — What type of extension?**

Ask yourself these questions (do not dump them on the user — synthesize into a clear recommendation):

| If the need is... | Then build a... | Why |
|---|---|---|
| A multi-step workflow with phases | **Skill** | Skills orchestrate complex work with lazy file loading |
| A focused delegated task | **Agent** | Agents run in isolated context with limited tools |
| Automation on tool events (format on save, lint on edit) | **Hook** | Hooks cost zero context tokens |
| Connecting an external API or service | **MCP config** | MCP provides tools without custom code |
| A simple user-triggered action | **Command** | Commands are lightweight slash-invocable prompts |
| A combination | **Plugin** with multiple components | Most real extensions combine 2-3 types |

**2.2 — Scope definition**

Through conversation, establish:
- **What it does** (be specific — "generates Zephyr device tree overlays" not "helps with device trees")
- **What it does NOT do** (explicit exclusions prevent scope creep)
- **Trigger conditions** (when should Claude auto-invoke this? what keywords/patterns?)
- **Context cost budget** (how many files will Claude need to read? can anything be pushed to hooks?)
- **External dependencies** (CLI tools, APIs, language runtimes, MCP servers)

**2.3 — Architecture proposal**

Present the user with:
- Extension type(s) and why
- File structure (directories and files that will be created)
- Which files are entry points vs. loaded-on-demand
- What goes into hooks (zero-cost automation) vs. skills (context-consuming workflows)
- Agent breakdown if applicable (what each agent does, what tools it needs, what model)

Use `AskUserQuestion` to get approval before proceeding. Iterate if the user wants changes.

### Phase 3: RESEARCH

**Goal**: Gather the implementation knowledge Claude doesn't already have.

This phase is conditional — skip if Claude already has sufficient knowledge.

Launch the `system-researcher` agent (Task tool, `subagent_type: general-purpose`) when:
- The extension involves an external CLI tool (need exact flags, output formats)
- The extension involves an API (need endpoints, auth, response schemas)
- The extension involves a framework pattern (need current best practices)
- The extension involves a standard or specification (need exact rules)

The researcher agent uses `WebSearch` and `WebFetch` to gather:
- Official documentation
- API references
- CLI help output
- Best practice guides
- Common pitfalls

The agent returns structured findings that feed directly into Phase 4.

If no research is needed (e.g., the extension is about a workflow Claude already knows), skip to Phase 4 and note that research was skipped.

### Phase 4: IMPLEMENT

**Goal**: Write every file for the extension. No placeholders, no TODOs, no stubs.

Read [reference/marketplace-structure.md](reference/marketplace-structure.md) and [reference/prompt-engineering.md](reference/prompt-engineering.md) now.

**4.1 — Write the entry point**

For skills, write SKILL.md following these rules:
- **Frontmatter**: name, description (specific enough for auto-routing), any control flags
- **File reference table**: every supporting file with "Load When" column
- **Workflow phases**: numbered, with clear entry/exit criteria per phase
- **Agent specifications**: which agents to launch, what data they receive, what they return
- SKILL.md is a ROUTER — it tells Claude what exists and when to load it. It never contains the knowledge itself.

For agents, write the agent .md file with:
- Clear role definition (who is this agent?)
- Specific task description (what does it do?)
- Tool list (minimum required)
- Output format (exactly what it returns)
- Constraints (what it must NOT do)

For hooks, write hooks.json with:
- Event type and matcher
- Command that runs
- Clear comments explaining the hook's purpose

For commands, write the command .md with:
- Description frontmatter
- Usage with examples
- Step-by-step instructions

**4.2 — Write supporting files**

For each file referenced in SKILL.md:
- `rules/` — constraints and guidelines, one file per topic, numbered for load order
- `reference/` — lookup tables, cheat sheets, specifications
- `scaffolds/` — code templates with clear markers for customization points
- `examples/` — worked examples showing the extension in action

**4.3 — Write the plugin manifest**

Create `.claude-plugin/plugin.json` with name, description, version, author.

**4.4 — Self-check before review**

Before moving to Phase 5, verify:
- [ ] Every file referenced in SKILL.md actually exists
- [ ] No file contains placeholder text (TODO, FIXME, TBD, "add here")
- [ ] SKILL.md frontmatter description is specific (not "helps with X", but "does X when Y")
- [ ] Hooks are used for anything that can be automated without context
- [ ] The cheapest appropriate model is specified for agents (`haiku` for simple, `sonnet` for moderate)
- [ ] File paths in references are correct relative paths

### Phase 5: REVIEW

**Goal**: Adversarial quality review. Improve until it meets the bar.

Read [reference/quality-standards.md](reference/quality-standards.md) and [reference/prompt-engineering.md](reference/prompt-engineering.md) now.

Launch the `system-reviewer` agent (Task tool, `subagent_type: general-purpose`). Give it:
1. All files that were created in Phase 4 (full content, not just paths)
2. The original user request
3. The architecture from Phase 2
4. The quality standards and prompt engineering references

The reviewer evaluates across 5 dimensions and scores each 1-5:

| Dimension | 1 (Fail) | 5 (Ship) |
|-----------|----------|----------|
| **Routing** | Generic description, won't auto-invoke correctly | Specific trigger conditions, Claude will pick this up reliably |
| **Architecture** | Everything dumped in one file, no lazy loading | Thin router, files loaded per-phase, hooks for automation |
| **Completeness** | Stubs, TODOs, missing files | Every file exists, every workflow step is actionable |
| **Prompt Quality** | Vague instructions, no examples, no constraints | Clear role, structured output, verification steps, few-shot examples |
| **Engineering** | Wasteful context, wrong model choices, no error handling | Minimal context cost, appropriate models, graceful failures |

The reviewer returns:
- Scores per dimension
- Specific issues found (file, line, problem, fix)
- Rewritten sections where the fix isn't obvious

**Passing bar**: All dimensions >= 4. If any dimension < 4, apply the reviewer's fixes and re-review. Maximum 2 review cycles (if still failing after 2 rounds, present the issues to the user for guidance).

### Phase 6: INTEGRATE

**Goal**: Place the extension in the correct location and verify it loads.

1. Determine target location:
   - If inside a marketplace repo → `plugins/<name>/`
   - If for current project → `.claude/skills/<name>/` or `.claude/commands/`
   - If global personal → `~/.claude/skills/<name>/`

2. Write all files to the target location

3. If inside a marketplace repo, update `marketplace.json` to include the new plugin entry

4. Verify the extension is discoverable:
   - For skills: confirm SKILL.md exists at the expected path
   - For hooks: confirm hooks.json is valid JSON
   - For agents: confirm agent .md exists at the expected path
   - For commands: confirm command .md exists with valid frontmatter

5. Present the user with a summary:
   - What was created (list of files)
   - How to invoke it (command, auto-trigger description, or hook event)
   - What it does (one sentence)
