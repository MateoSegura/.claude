---
name: docs-presentations-business
description: Generate Slidev presentations from existing business presentation markdown files. Reads section files (speaker notes, data, visual descriptions) and produces branded, presentable Slidev decks with charts, layouts, and design — focused on rendering, not content strategy.
---

# Business Presentations — Slidev Generator

Convert existing business presentation markdown into **branded Slidev decks** ready to present.

```
Phase 1: SCAN              Phase 2: PROPOSE           Phase 3: RENDER
(read existing files)  →   (approval loop)        →   (parallel agents)

Read all section .md       Show deck plan             One agent per section
Extract slide content      User edits/approves        Slidev + charts + layout
Map charts & visuals       Iterate until approved     Assembly + run.sh
Detect structure
```

**Core philosophy**: The content already exists. This skill's job is to turn slide blueprints (titles, bullets, data tables, visual descriptions, speaker notes) into polished Slidev decks with real charts, two-column layouts, metric grids, and Mermaid diagrams. No content generation — only rendering.

## Supporting files

| File | Purpose | When to load |
|------|---------|-------------|
| [reference.md](reference.md) | Brand, rendering rules, content budgets, Mermaid rules | Phase 3 (rendering) |
| [templates/_design-system.md](templates/_design-system.md) | CSS custom properties, full `<style>` block, all CSS classes | Phase 3 (rendering) |
| [templates/00-example-deck.md](templates/00-example-deck.md) | Complete example deck using design system classes | Phase 3 (formatting reference) |
| Template files 01-14 | Individual slide type templates with sizing notes | Phase 3 (load per slide type) |
| [templates/example-strategy.md](templates/example-strategy.md) | Strategy document structure | Phase 1 (source detection) |

### Template Index

| # | Template | File |
|---|----------|------|
| 01 | Title Card | [templates/01-title-card.md](templates/01-title-card.md) |
| 02 | Section Divider | [templates/02-section-divider.md](templates/02-section-divider.md) |
| 03 | Text Slide | [templates/03-text-slide.md](templates/03-text-slide.md) |
| 04 | Metric Grid | [templates/04-metric-grid.md](templates/04-metric-grid.md) |
| 05 | Two-Column | [templates/05-two-column.md](templates/05-two-column.md) |
| 06 | Bar Chart | [templates/06-bar-chart.md](templates/06-bar-chart.md) |
| 07 | Scenario Cards | [templates/07-scenario-cards.md](templates/07-scenario-cards.md) |
| 08 | Heatmap | [templates/08-heatmap.md](templates/08-heatmap.md) |
| 09 | Timeline | [templates/09-timeline.md](templates/09-timeline.md) |
| 10 | Decision Ask | [templates/10-decision-ask.md](templates/10-decision-ask.md) |
| 11 | Data Table | [templates/11-data-table.md](templates/11-data-table.md) |
| 12 | Mermaid Diagram | [templates/12-mermaid-diagram.md](templates/12-mermaid-diagram.md) |
| 13 | Center Pause | [templates/13-center-pause.md](templates/13-center-pause.md) |
| 14 | Closing Card | [templates/14-closing-card.md](templates/14-closing-card.md) |

---

## Phase 1: SCAN

**Goal**: Read existing presentation files and build a rendering plan.

### Step 1.1 — Find and read source files

Look for presentation content in the directory the user points to (or the project root). Expected input formats:

| Format | How to detect |
|--------|--------------|
| **Section files** | `NN-name.md` files with `## Slide N:` headings |
| **Overview file** | `00-overview.md` with deck flow table |
| **Strategy doc** | `strategy.md` with presentation architecture |
| **Single file** | One `.md` with slide-level headings |

Read ALL source files. Extract from each slide:

- **Title** — the `## Slide N:` heading or `**Title:**` field
- **On-screen text** — bullets, quotes, numbered items
- **Data** — tables with numbers, chart data, financial figures
- **Visual description** — `**Visual:**`, `**Chart type:**`, `**Design notes:**` fields
- **Speaker notes** — `**Speaker notes:**` blocks (become Slidev presenter notes)
- **Preparation notes** — `**Preparation notes:**` blocks (omitted from slides)

### Step 1.2 — Build the rendering plan

For each slide, determine the best Slidev rendering using the template index:

| Source content | Template |
|---------------|----------|
| Bullets + title only | 03 — Text Slide |
| Data table + chart type | 11 — Data Table or 12 — Mermaid Diagram |
| Before/After columns | 05 — Two-Column |
| Metric numbers | 04 — Metric Grid |
| Timeline / roadmap | 09 — Timeline |
| Architecture / flow | 12 — Mermaid Diagram |
| Heatmap / matrix | 08 — Heatmap |
| Demo slide | 01 — Title Card (variant) |
| "The Pause" / black screen | 13 — Center Pause |
| Ask / decision slide | 10 — Decision Ask |
| Bar chart | 06 — Bar Chart |
| Stacked comparison | 07 — Scenario Cards |
| Section break | 02 — Section Divider |
| Final slide | 14 — Closing Card |

### Step 1.3 — Synthesize rendering manifest

```yaml
source_dir: <path>
source_files: <count>
total_slides: <count>
sections:
  - file: "01-demos.md"
    title: "<title>"
    slides: <count>
    render_types: [demo, metric-grid, text, chart]
  - file: "02-landscape.md"
    ...
backup_slides: <count>
output: <single slides.md | multi-file>
```

**Output mode**:
- **≤15 slides** → single `slides.md`
- **>15 slides** → one `.md` per section + `index.md` navigation page

---

## Phase 2: PROPOSE (Approval Loop)

**Goal**: Show the user what will be rendered and get approval before generating.

### Step 2.1 — Present the rendering plan

Use `AskUserQuestion` to show:
- Each section with slide count and rendering types
- Which slides get charts vs. text vs. layouts
- Total deck size
- Any slides where the visual description is ambiguous (offer choices)

Options:
- **Approve** → proceed to Phase 3
- **Edit** → change rendering choices for specific slides
- **Split/merge** → combine sections or split long ones
- **Skip** → exclude specific slides or sections

### Step 2.2 — Iterate

Loop until the user approves. If the source content is ambiguous:
- Propose the best rendering and explain why
- Offer alternatives: "This ROI data could be a bar chart or a metric grid — which do you prefer?"

---

## Phase 3: RENDER (Parallel Generation)

**Goal**: Produce Slidev files from the approved rendering plan.

### Step 3.1 — Load reference material

Before launching agents, read these files:
- [reference.md](reference.md) — brand, rendering rules, content budgets
- [templates/_design-system.md](templates/_design-system.md) — full CSS `<style>` block, all class definitions
- [templates/00-example-deck.md](templates/00-example-deck.md) — formatting reference

Also identify which template files (01-14) are needed based on the rendering plan. Load only the ones required for each agent.

### Step 3.2 — Launch parallel rendering agents

Launch **one Task agent per section** (`subagent_type: general-purpose`). Each agent receives:

1. The source section file content (the existing markdown)
2. The rendering plan for that section (which slides get which treatment)
3. The full content of `reference.md`
4. The full content of `templates/_design-system.md`
5. The relevant template files for the slide types in this section
6. Instructions to follow the Slidev format exactly

**Agent instructions template:**

```
You are rendering business presentation slides into Slidev format.

INPUT: The source markdown below contains slide content with titles, on-screen text,
speaker notes, visual descriptions, and data tables.

OUTPUT: A Slidev-formatted .md file. For each slide:
1. Convert the title and on-screen text into Slidev slide content
2. Use the CSS classes from _design-system.md — NEVER use inline styles for values
   that have CSS classes (bar charts, timelines, scenarios, metrics, etc.)
3. Follow the template files for each slide type's structure
4. Convert speaker notes into Slidev presenter notes (<!-- speaker notes --> blocks)
5. OMIT preparation notes entirely — they don't go in the deck
6. Follow the brand palette, typography, and layout rules exactly

CRITICAL RULES:
- Never use markdown bullets inside HTML divs — use <ul><li> tags
- Every Mermaid diagram needs <Transform> wrapper and legend
- Speaker notes go in HTML comment blocks after the slide content
- Maximum content budgets from reference.md are hard limits
- Use CSS classes from the design system — NOT inline styles
- The ONLY inline styles allowed are:
  - width percentages on .bar-fill elements (style="width: XX%")
  - font-size: var(--font-small) on table wrapper divs
  - margin-top: var(--space-*) for spacing adjustments
```

### Step 3.3 — Assemble output

#### Single-deck output (≤15 slides)

```
<output_dir>/
├── package.json
├── run.sh
├── setup/
│   └── mermaid.ts
└── slides.md
```

#### Multi-deck output (>15 slides)

```
<output_dir>/
├── package.json
├── run.sh
├── setup/
│   └── mermaid.ts
├── index.md           # Navigation page
├── 01-<section>.md
├── 02-<section>.md
├── ...
└── 0N-backup.md
```

The first section file includes the full `<style>` block from `_design-system.md` (with footer variables customized). All subsequent sections reference it.

### Step 3.4 — Generate support files

Write `package.json`, `setup/mermaid.ts`, and `run.sh` using the templates from `reference.md`.

### Step 3.5 — Verify each deck

```
[ ] Every slide renders valid Slidev markdown
[ ] No markdown syntax inside HTML divs
[ ] Every Mermaid diagram has <Transform> + dark theme + legend
[ ] Speaker notes are in HTML comment blocks
[ ] Preparation notes are NOT in the output
[ ] Content budgets respected (bullets, words, table rows)
[ ] Brand palette used consistently
[ ] Charts render the actual data from source tables
[ ] Slide separators (---) are correct
[ ] CSS classes from design system used (no redundant inline styles)
[ ] Safe zone padding applied via .slidev-layout
[ ] run.sh is executable
```

### Step 3.6 — Export guidance

Detect WSL: `grep -qi microsoft /proc/version 2>/dev/null`

- **Non-WSL**: `npm install && npm i -D playwright-chromium`, then `npx slidev export --output <name>.pdf --timeout 60000`
- **WSL/fallback**: `npm install`, start slidev, tell user to open `http://localhost:3030/export`

---

## Content Budgets (Hard Limits)

Business slides follow tighter budgets than code presentations since the audience is non-technical decision-makers.

| Element | Max Per Slide |
|---------|---------------|
| Bullet points | 3 |
| Words per bullet | 20 |
| Table rows | 5 |
| Table columns | 4 |
| Metric cards | 6 |
| Diagram nodes | 10 |
| Bar chart items | 6 |
| Mermaid edges | 12 |

---

## Adaptation Rules

The rendering is **content-driven**, not structure-driven. The skill adapts to whatever the source files contain:

- If the source has 5 sections → render 5 sections
- If the source has 50 slides → split across multiple decks
- If the source has no charts → render text and layout slides
- If the source has financial data → render metric grids and comparison cards
- If the source has timelines → render timeline visuals
- If the source has before/after → render two-column cards

There is no required phase sequence (HOOK/CONTEXT/etc.) imposed by this skill. The content structure comes from the source files.
