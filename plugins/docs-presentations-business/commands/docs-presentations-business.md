# Generate Business Presentation

Generate a strategy-backed business presentation with speaker notes, data verification, and blindspot analysis.

## Usage

```
/docs-presentations-business <context-file-or-description>
```

**Arguments:**
- A path to a file containing your business context (situation, goals, constraints), OR
- A natural language description of what you're proposing

**Examples:**
```
/docs-presentations-business initial-prompt
/docs-presentations-business "Propose adopting AI coding tools for a 15-person IoT company"
/docs-presentations-business strategy-brief.md
```

---

## Instructions

This command triggers the `docs-presentations-business` skill pipeline. Follow the skill's four-phase process exactly.

### Pre-flight

1. Read the skill definition: `.claude/skills/docs-presentations-business/SKILL.md`
2. Read the reference guide: `.claude/skills/docs-presentations-business/reference.md`
3. Read the example template: `.claude/skills/docs-presentations-business/templates/example-strategy.md`
4. If a context file was provided, read it

### Execute the pipeline

Follow the four phases defined in `SKILL.md`:

1. **DISCOVER** — Gather context from the user, then launch 3 parallel research agents (Market Researcher, Financial Modeler, Risk Analyst)
2. **STRATEGIZE** — Generate the strategy document (situation assessment, audience analysis, presentation architecture, financial model, blindspots, audience playbook, preparation checklist)
3. **PROPOSE** — Present the deck plan to the user for iterative approval
4. **GENERATE** — Launch parallel agents per section, then run a cohesion reviewer

### Output structure

```
<project>/
├── strategy.md
├── presentation/
│   ├── 00-overview.md
│   ├── 01-<section>.md
│   ├── 02-<section>.md
│   ├── ...
│   └── 0N-backup.md
```

### Quality checks

Before delivering, verify:
- [ ] Every data claim has a source
- [ ] Every calculation is reproducible from stated assumptions
- [ ] Slide numbering is sequential across all files
- [ ] Data points match everywhere they appear
- [ ] Speaker notes are in first person, directive style
- [ ] Every blindspot has a prepared counterargument
- [ ] The ask is specific: what decisions, by when
