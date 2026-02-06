---
name: docs-presentations-code
description: Generate purpose-driven Slidev presentations for libraries and tools. Three-phase pipeline — discover the codebase, propose a deck plan to the user, then generate presentations in parallel. Produces branded decks that help audiences understand when, why, and how to use a project.
---

# Project Presentations

Generate Slidev presentations through a **3-phase orchestrated pipeline**:

```
Phase 1: DISCOVER          Phase 2: PROPOSE           Phase 3: GENERATE
(parallel agents)    →     (human approval loop)   →  (parallel agents)

Explore subsystems         "Here's what I'd build"    One agent per deck
Rate complexity            User edits/approves         All decks + index
Map boundaries             Iterate until "approved"    run.sh with nav
```

**Core philosophy**: In an era of AI-assisted development, nobody reads implementation code in a presentation. Developers need to understand the *purpose*, *integration points*, and *real-world scenarios* — the code writes itself once they know what to build.

## Supporting files

- For brand, styling, rendering rules, and slide structure, see [reference.md](reference.md)
- For a complete example deck with correct diagram/code formatting, see [templates/example-deck.md](templates/example-deck.md)

Load `reference.md` when generating slides in Phase 3. Load `templates/example-deck.md` as a formatting reference during Phase 3.

---

## Phase 1: DISCOVER

**Goal**: Deeply understand the codebase before proposing anything. Never generate slides without discovery.

### Step 1.1 — Launch parallel exploration agents

Launch **3 agents in parallel** using the Task tool (`subagent_type: Explore`):

| Agent | Prompt Focus | Collects |
|-------|-------------|----------|
| **Architecture Explorer** | "Map the top-level modules, entry points, dependencies, and external integrations of this project. Identify separable subsystems. Report boundaries between components." | Subsystem list, dependency graph, integration points |
| **Purpose Analyzer** | "Read READMEs, docs, doc comments, and commit history. Determine: What problem does this solve? Who is the target audience? What are the key use cases? What would someone Google that leads them here?" | Problem statement, audience, use cases |
| **Metrics Collector** | "Count: source LOC, test LOC, external dependencies, file count, release tags/versions. Check for CI, test commands, and coverage reports." | Numbers for Vital Signs and Trust Signals |

### Step 1.2 — Synthesize into a Project Profile

After all agents return, synthesize their findings into a **Project Profile**:

```
Project: <name>
Rating: <POC | MVP | Product>
  - POC: < 1K LOC, single-file or trivial structure, proof of concept
  - MVP: 1K-10K LOC, clear API surface, focused purpose
  - Product: 10K+ LOC, multiple subsystems, separable components
Subsystems: <count> (<list of names>)
Recommended decks: <count>
  - "<id>" — <title> (<depth: lighter | full>, <slide count>)
  - ...
Key use cases found:
  1. <use case>
  2. <use case>
  3. <use case>
```

**Depth scaling**:
- **POC**: 1 lighter deck (12-16 slides, fewer scenarios, simpler diagrams, max 5 nodes)
- **MVP**: 1 full deck (15-24 slides, standard format)
- **Product**: N decks — one `00-overview` (lighter) + one full deck per major subsystem

---

## Phase 2: PROPOSE (Human Approval Loop)

**Goal**: Get the user's explicit approval before generating anything. This is an iterative conversation, not a single yes/no.

### Step 2.1 — Present the Project Profile

Use `AskUserQuestion` to present the profile and proposed deck plan. Show:
- The project rating and reasoning
- Each proposed deck: id, title, focus, depth, estimated slide count
- The key use cases discovered

Offer the user options:
- **Approve as-is** → proceed to Phase 3
- **Edit** → change titles, focus, depth, or slide emphasis for any deck
- **Add deck** → user proposes a new deck topic
- **Remove deck** → drop a proposed deck
- **Ask questions** → user wants more info about what was discovered

### Step 2.2 — Iterate

If the user asks for changes or more information:
- If Claude has the info from Phase 1 → answer and re-propose
- If Claude doesn't have the info → say so honestly, offer to explore further or skip
- **Claude can push back**: "I didn't find deployment infrastructure in this codebase — the project is a library consumed via `go get`. Want me to explore further, or skip that deck?"

**Loop continues until the user says "approved"** (or equivalent affirmation).

### Step 2.3 — Lock the manifest

Once approved, the final deck manifest is fixed:

```yaml
project: <name>
build_tag: <version or YYYY-Mon-DD>
rating: <POC | MVP | Product>
decks:
  - id: "00-overview"
    title: "<title>"
    focus: "<focus description>"
    depth: "lighter"
    slides: 15
  - id: "01-<subsystem>"
    title: "<title>"
    focus: "<focus description>"
    depth: "full"
    slides: 20
```

For single-deck projects, the manifest has one entry and `id` is omitted (output is `slides.md`).

---

## Phase 3: GENERATE (Parallel Deck Creation)

**Goal**: Produce all approved decks in parallel, with a shared nav index for multi-deck projects.

### Step 3.1 — Load reference material

Before launching generation agents, read:
- [reference.md](reference.md) — brand, rendering rules, slide structure
- [templates/example-deck.md](templates/example-deck.md) — formatting reference

### Step 3.2 — Launch parallel generation agents

Launch **one Task agent per deck** (`subagent_type: general-purpose`). Each agent receives:

1. The approved deck spec (title, focus, depth, slide count)
2. The relevant discovery data for its subsystem
3. The full content of `reference.md` (brand, rendering, slide rules)
4. The full content of `templates/example-deck.md` (formatting reference)

Each agent follows the Five Acts structure from `reference.md` and writes its deck's `.md` file.

### Step 3.3 — Assemble output

#### Single-deck output (POC/MVP, 1 deck)

```
docs/presentations/<BUILD_TAG>/
├── package.json
├── run.sh
├── setup/
│   └── mermaid.ts
└── slides.md
```

#### Multi-deck output (Product, N decks)

```
docs/presentations/<BUILD_TAG>/
├── package.json
├── run.sh
├── setup/
│   └── mermaid.ts
├── index.md
├── 00-overview.md
├── 01-<subsystem>.md
├── 02-<subsystem>.md
└── ...
```

### `run.sh` — Multi-deck aware

```bash
#!/usr/bin/env bash
set -euo pipefail
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
info()  { printf "${GREEN}[+]${NC} %s\n" "$1"; }
warn()  { printf "${YELLOW}[!]${NC} %s\n" "$1"; }
error() { printf "${RED}[x]${NC} %s\n" "$1"; }

if ! command -v node &>/dev/null; then error "Node.js not installed."; exit 1; fi
NODE_VERSION=$(node -v | sed 's/v//' | cut -d. -f1)
if [ "$NODE_VERSION" -lt 18 ]; then error "Node.js v18+ required (found $(node -v))."; exit 1; fi
info "Node.js $(node -v)"

if [ ! -d "node_modules" ]; then warn "Installing dependencies..."; npm install; else info "Dependencies ready"; fi

PORT="${2:-3030}"

# Resolve deck file
if [ $# -eq 0 ]; then
  # No args: if index.md exists, use it; otherwise use slides.md
  if [ -f "index.md" ]; then
    DECK="index.md"
  else
    DECK="slides.md"
  fi
else
  # Arg provided: resolve to .md file
  DECK="$1"
  [[ "$DECK" != *.md ]] && DECK="${DECK}.md"
  if [ ! -f "$DECK" ]; then
    error "Deck not found: $DECK"
    echo ""
    echo "Available decks:"
    ls -1 *.md 2>/dev/null | grep -v 'node_modules' | sed 's/^/  /'
    exit 1
  fi
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
info "Starting Slidev → http://localhost:${PORT}"
echo "  Deck: $DECK"
echo "  Ctrl+C to stop · /export for PDF"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
npx slidev "$DECK" --port "$PORT" --open
```

### `index.md` — Navigation landing page (multi-deck only)

A Slidev deck with 2 slides:
- **Slide 1**: Project title, rating badge, build tag, count of presentations
- **Slide 2**: Card grid — one card per deck showing title, focus summary, slide count. Each card tells the user to run `./run.sh <deck-name>`.

Uses the same brand styling as all other decks.

### Step 3.4 — Verify each deck

Run the generation checklist per deck:

```
[ ] Five Acts: ORIENT, MAP, SCENARIOS, PROVE, LAND
[ ] Slide count in range (lighter: 12-16, full: 15-24)
[ ] No slide exceeds content budget
[ ] Max 2 code slides, max 10 lines each
[ ] Every diagram has dark theme + legend
[ ] Every sequenceDiagram has <Transform> + zoom: 0.9 + ≤6 messages
[ ] Content ratios in range (diagrams+scenarios 35-50%, code ≤10%)
[ ] At least 2 real-world use cases in SCENARIOS
[ ] No internal implementation code (caller-perspective only)
[ ] Ordering rules satisfied
[ ] No forward references
[ ] No markdown syntax inside HTML elements (use <ul><li> not "- item" inside divs)
```

### Step 3.5 — Write files and set permissions

Write all files. Run `chmod +x run.sh`. Also copy each deck to `docs/templates/<project>-<deck-id>.md` for version control.

### Step 3.6 — Export

**Detect WSL**: `grep -qi microsoft /proc/version 2>/dev/null`

- **Non-WSL**: `npm install && npm i -D playwright-chromium`, then `npx slidev export --output <name>.pdf --timeout 60000`
- **WSL/fallback**: `npm install`, start slidev, tell user to open `http://localhost:3030/export`

---

## Build Tag Resolution

1. Check in order: Go `const Version` / `go.mod`, Node `package.json`, Python `__version__` / `pyproject.toml`, Rust `Cargo.toml`, `git describe --tags --abbrev=0`, `VERSION` file
2. Version found → tag is `v0.2.0`
3. No version → tag is `YYYY-Mon-DD` (e.g. `2026-Jan-28`)

---

## Project Type Adaptations

| Project Type | ORIENT | MAP | SCENARIOS | PROVE |
|-------------|--------|-----|-----------|-------|
| **Library/SDK** | What it does, install | Integration boundaries | "When you need X...", quick start | Trust signals, project map |
| **CLI Tool** | Commands overview, install | System context | Workflow automation scenarios | Adoption stats, error recovery |
| **Web Service** | What it serves, deploy | Request flow (caller view) | Integration patterns for clients | Uptime, scaling characteristics |
| **Embedded/RTOS** | Hardware, constraints | System block diagram | Deployment scenarios, use cases | Certification, test coverage |
| **Data Pipeline** | I/O formats, throughput | Pipeline topology | "When your data looks like X..." | Validation, performance numbers |
| **Monorepo** | Package overview | Dependency graph | "Use package X when...", composition | Cross-package compatibility |
