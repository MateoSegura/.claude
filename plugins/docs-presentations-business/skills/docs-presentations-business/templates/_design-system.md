# Design System

CSS custom properties, typography scale, spacing tokens, and component classes for all business presentation templates. Include this `<style>` block in the first slide file of every deck.

---

## How to Use

1. Copy the entire `<style>` block below into the first `.md` file of your Slidev deck (after the YAML frontmatter)
2. Set `--footer-left` and `--footer-right` to your presenter name and deck name
3. Use CSS classes from this system instead of inline styles — templates reference these classes

---

## Full Style Block

```html
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&family=JetBrains+Mono:wght@400;500&display=swap');

/* ── Color Tokens ──────────────────────────────────────────── */
:root {
  --c-bg:       #1A1A1A;
  --c-text:     #E8E4DD;
  --c-gold:     #D4A853;
  --c-gold-dim: #B8912A;
  --c-card:     #2A2A2A;
  --c-code:     #1E1E1E;
  --c-border:   #3A3A3A;
  --c-muted:    #6B6B6B;
  --c-dim:      #4A4A4A;
  --c-sub:      #C8C4BD;

  /* ── Typography Scale (Golden Ratio 1.618) ───────────────── */
  --font-display:   4.236rem;   /* Single hero number            */
  --font-h1:        2.618rem;   /* Slide title, divider title    */
  --font-h2:        1.618rem;   /* Section heading (h2/h3)       */
  --font-subtitle:  1.25rem;    /* Subtitle, card headings       */
  --font-body:      1rem;       /* Body text, bullets            */
  --font-small:     0.85rem;    /* Labels, table cells           */
  --font-caption:   0.75rem;    /* Uppercase labels              */
  --font-micro:     0.618rem;   /* Footer                        */

  /* Metric-specific sizes */
  --font-metric:    2.618rem;   /* Metric card values            */
  --font-metric-lg: 3.25rem;    /* Scenario card hero numbers    */

  /* ── Spacing Scale (8px base) ────────────────────────────── */
  --space-xs:  0.5rem;    /*  8px  */
  --space-sm:  1rem;      /* 16px  */
  --space-md:  1.5rem;    /* 24px  */
  --space-lg:  2rem;      /* 32px  */
  --space-xl:  3rem;      /* 48px  */
  --space-2xl: 4rem;      /* 64px  */

  /* ── Component Dimensions ────────────────────────────────── */
  --bar-height:        2rem;     /* 32px — was 28px     */
  --bar-label-width:   10rem;    /* was 120px/180px     */
  --bar-value-width:   6.5rem;   /* was 60px/100px      */
  --timeline-dot:      2.25rem;  /* 36px — was 28px     */
  --card-radius:       0.5rem;   /* 8px                 */
  --heatmap-header-width: 11rem; /* was 140px/180px     */

  /* ── Footer ──────────────────────────────────────────────── */
  --footer-left:  'Presenter · Month Year';
  --footer-right: 'presentation-name';

  /* ── Slidev Overrides ────────────────────────────────────── */
  --slidev-theme-default-background: var(--c-bg);
  --slidev-theme-default-headingColor: var(--c-text);
}

/* ── Safe Zone ─────────────────────────────────────────────── */
/* ~7% top, ~5% sides, ~10% bottom (room for footer)           */
.slidev-layout {
  background: var(--c-bg) !important;
  color: var(--c-text) !important;
  font-family: 'Inter', system-ui, sans-serif !important;
  padding: 48px 60px 68px 60px !important;
}

/* ── Persistent Footer ─────────────────────────────────────── */
.slidev-layout::before {
  content: var(--footer-left);
  position: fixed; bottom: 1em; left: 2em;
  font-size: var(--font-micro); color: var(--c-dim);
  font-family: 'Inter', system-ui;
  letter-spacing: 0.02em; z-index: 100;
}
.slidev-layout::after {
  content: var(--footer-right);
  position: fixed; bottom: 1em; right: 2em;
  font-size: var(--font-micro); color: var(--c-dim);
  font-family: 'JetBrains Mono', monospace;
  z-index: 100;
}

/* ── Base Typography ───────────────────────────────────────── */
h1 {
  color: var(--c-text) !important;
  font-size: var(--font-h1) !important;
  font-weight: 700 !important;
}
h2, h3 {
  color: var(--c-gold) !important;
  font-size: var(--font-h2) !important;
  font-weight: 600 !important;
}
p, li { color: var(--c-text) !important; font-size: var(--font-body); }
strong { color: var(--c-gold) !important; }
a { color: var(--c-gold) !important; }
code:not(pre code) {
  color: var(--c-gold) !important; background: var(--c-card) !important;
  padding: 0.15em 0.4em; border-radius: 4px; font-size: 0.9em;
}
blockquote {
  border-left: 3px solid var(--c-gold) !important;
  background: var(--c-card) !important;
  padding: var(--space-xs) var(--space-sm) !important;
}

/* ── Tables ────────────────────────────────────────────────── */
table { border-collapse: collapse; width: 100%; }
th {
  background: var(--c-card) !important; color: var(--c-gold) !important;
  padding: var(--space-xs) var(--space-sm);
  border: 1px solid var(--c-border); font-size: var(--font-small);
}
td {
  color: var(--c-text) !important;
  padding: var(--space-xs) var(--space-sm);
  border: 1px solid var(--c-border); font-size: var(--font-small);
}
tr:nth-child(even) { background: rgba(42, 42, 42, 0.4); }

/* ── Code Blocks ───────────────────────────────────────────── */
pre {
  background: var(--c-code) !important;
  border: 1px solid #333333 !important;
  border-radius: 6px !important;
  font-family: 'JetBrains Mono', 'Fira Code', monospace !important;
  font-size: 0.82em !important;
  padding: var(--space-sm) !important;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.4) !important;
}

/* ── Mermaid ───────────────────────────────────────────────── */
.mermaid {
  overflow: visible !important;
  display: flex;
  justify-content: center;
  background: transparent !important;
}
.mermaid svg {
  max-width: 95% !important;
  max-height: 440px !important;
  height: auto !important;
}
.mermaid .cluster text { fill: var(--c-gold); }
.mermaid .edgeLabel { background-color: var(--c-bg); color: var(--c-text); }

/* ── Metric Grid ───────────────────────────────────────────── */
.metric-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: var(--space-md);
  margin-top: var(--space-md);
}
.metric-grid--4 { grid-template-columns: repeat(4, 1fr); }
.metric-grid--6 { grid-template-columns: repeat(3, 1fr); }

.metric-card {
  background: var(--c-card);
  border: 1px solid var(--c-border);
  border-radius: var(--card-radius);
  padding: var(--space-lg);
  text-align: center;
}
.metric-value {
  font-size: var(--font-metric);
  font-weight: 700;
  color: var(--c-gold);
  font-family: 'JetBrains Mono', monospace;
  line-height: 1.2;
}
.metric-label {
  font-size: var(--font-caption);
  color: var(--c-muted);
  margin-top: var(--space-xs);
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

/* ── Two-Column Layout ─────────────────────────────────────── */
.two-col {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--space-lg);
  margin-top: var(--space-sm);
}
.card {
  background: var(--c-card);
  border: 1px solid var(--c-border);
  border-radius: var(--card-radius);
  padding: var(--space-lg);
}
.card h3 { margin-top: 0; }
.card-list {
  list-style: none;
  padding: 0;
  margin: 0;
}
.card-list li {
  padding: var(--space-xs) 0;
  border-bottom: 1px solid var(--c-border);
}
.card-list li:last-child { border-bottom: none; }

/* ── Bar Chart ─────────────────────────────────────────────── */
.bar-row {
  display: flex;
  align-items: center;
  margin-bottom: var(--space-sm);
}
.bar-label {
  width: var(--bar-label-width);
  color: var(--c-sub);
  font-size: var(--font-small);
  flex-shrink: 0;
}
.bar-track {
  flex: 1;
  background: var(--c-card);
  border-radius: 4px;
  height: var(--bar-height);
  position: relative;
}
.bar-fill {
  height: 100%;
  border-radius: 4px;
  background: var(--c-gold);
}
.bar-fill--muted { background: var(--c-muted); }
.bar-value {
  width: var(--bar-value-width);
  text-align: right;
  color: var(--c-gold);
  font-family: 'JetBrains Mono', monospace;
  font-size: var(--font-small);
  flex-shrink: 0;
  padding-left: var(--space-xs);
}
.bar-value--muted { color: var(--c-sub); }

/* ── Scenario Cards ────────────────────────────────────────── */
.scenario-row {
  display: flex;
  gap: var(--space-md);
  justify-content: center;
  margin-top: var(--space-sm);
}
.scenario-card {
  text-align: center;
  flex: 1;
  background: var(--c-card);
  border: 1px solid var(--c-border);
  border-radius: var(--card-radius);
  padding: var(--space-sm);
}
.scenario-card--primary {
  border: 2px solid var(--c-gold);
}
.scenario-card .scenario-label {
  font-size: var(--font-caption);
  color: var(--c-muted);
  text-transform: uppercase;
  letter-spacing: 0.08em;
}
.scenario-card--primary .scenario-label {
  color: var(--c-gold);
  font-weight: 600;
}
.scenario-value {
  font-size: var(--font-metric-lg);
  font-family: 'JetBrains Mono', monospace;
  font-weight: 700;
  margin: var(--space-xs) 0;
  color: var(--c-sub);
}
.scenario-card--primary .scenario-value {
  color: var(--c-gold);
}
.scenario-detail {
  font-size: var(--font-small);
  color: var(--c-sub);
}
.scenario-sub {
  font-size: var(--font-small);
  color: var(--c-muted);
  margin-top: var(--space-xs);
}

/* ── Badge Row (Title Card) ────────────────────────────────── */
.badge-row {
  display: flex;
  gap: var(--space-sm);
  margin-top: var(--space-xl);
}
.badge {
  background: var(--c-card);
  border: 1px solid var(--c-border);
  padding: 0.3em 0.8em;
  border-radius: 4px;
  font-family: 'JetBrains Mono', monospace;
  font-size: var(--font-small);
  color: var(--c-sub);
}
.badge--primary {
  border-color: var(--c-gold);
  color: var(--c-gold);
}

/* ── Section Divider ───────────────────────────────────────── */
.divider-slide {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}
.divider-act {
  font-size: var(--font-small);
  color: var(--c-muted);
  text-transform: uppercase;
  letter-spacing: 0.3em;
  margin-bottom: var(--space-xs);
}
.divider-title {
  font-size: var(--font-h1);
  font-weight: 700;
  color: var(--c-text);
}
.divider-sub {
  font-size: var(--font-body);
  color: var(--c-muted);
  margin-top: var(--space-xs);
}

/* ── Timeline ──────────────────────────────────────────────── */
.timeline-container {
  margin-top: var(--space-lg);
  position: relative;
}
.timeline-line {
  position: absolute;
  top: calc(var(--timeline-dot) / 2);
  left: 0; right: 0;
  height: 2px;
  background: var(--c-border);
}
.timeline-phases {
  display: flex;
  justify-content: space-between;
  position: relative;
}
.timeline-phase {
  text-align: center;
  flex: 1;
}
.timeline-dot {
  width: var(--timeline-dot);
  height: var(--timeline-dot);
  border-radius: 50%;
  margin: 0 auto var(--space-xs);
}
.timeline-dot--active {
  background: var(--c-gold);
}
.timeline-dot--future {
  background: var(--c-border);
  border: 2px solid var(--c-gold);
}
.timeline-phase-title {
  font-size: var(--font-small);
  font-weight: 600;
}
.timeline-dot--active + .timeline-phase-title,
.timeline-phase:first-child .timeline-phase-title {
  color: var(--c-gold);
}
.timeline-phase-date {
  font-size: var(--font-caption);
  color: var(--c-sub);
  margin-top: var(--space-xs);
}
.timeline-phase-desc {
  font-size: var(--font-micro);
  color: var(--c-muted);
  margin-top: 0.2rem;
}

/* ── Heatmap / Matrix ──────────────────────────────────────── */
.heatmap {
  display: grid;
  gap: 2px;
  font-size: var(--font-small);
  margin-top: var(--space-md);
}
.heatmap--3col { grid-template-columns: var(--heatmap-header-width) repeat(3, 1fr); }
.heatmap--2col { grid-template-columns: var(--heatmap-header-width) repeat(2, 1fr); }
.heatmap-corner { padding: 0.6em; color: var(--c-muted); }
.heatmap-colhead {
  padding: 0.6em;
  text-align: center;
  color: var(--c-gold);
  font-weight: 600;
}
.heatmap-rowhead {
  padding: 0.6em;
  color: var(--c-sub);
  font-weight: 600;
}
.heatmap-cell {
  padding: 0.6em;
  text-align: center;
  border-radius: 4px;
}
.heatmap-cell--high  { background: rgba(212,168,83,0.4); color: var(--c-gold); font-weight: 600; }
.heatmap-cell--med   { background: rgba(212,168,83,0.15); }
.heatmap-cell--low   { background: rgba(107,107,107,0.2); }

/* ── Decision / Ask ────────────────────────────────────────── */
.decision-item {
  display: flex;
  align-items: baseline;
  gap: var(--space-xs);
  margin-bottom: var(--space-lg);
}
.decision-number {
  font-size: var(--font-h1);
  color: var(--c-gold);
  font-family: 'JetBrains Mono', monospace;
  font-weight: 700;
}
.decision-title {
  font-size: var(--font-subtitle);
  color: var(--c-text);
  font-weight: 600;
}
.decision-desc {
  font-size: var(--font-body);
  color: var(--c-sub);
  margin-top: var(--space-xs);
}

/* ── Center Pause ──────────────────────────────────────────── */
.pause-context {
  margin-top: var(--space-lg);
  color: var(--c-muted);
  font-size: var(--font-small);
}

/* ── Closing Card ──────────────────────────────────────────── */
.closing-tagline {
  font-size: var(--font-subtitle);
  color: var(--c-sub);
  margin-top: var(--space-lg);
}
.closing-attribution {
  margin-top: var(--space-xl);
  color: var(--c-muted);
  font-size: var(--font-small);
}
</style>
```

---

## Token Reference

### Colors

| Token | Hex | Use |
|-------|-----|-----|
| `--c-bg` | `#1A1A1A` | Slide background |
| `--c-text` | `#E8E4DD` | Primary text, h1 |
| `--c-gold` | `#D4A853` | Accents, h2/h3, key metrics, borders |
| `--c-gold-dim` | `#B8912A` | Mermaid stroke on primary nodes |
| `--c-card` | `#2A2A2A` | Cards, table headers, blockquotes |
| `--c-code` | `#1E1E1E` | Code block background |
| `--c-border` | `#3A3A3A` | Borders, dividers |
| `--c-muted` | `#6B6B6B` | Captions, secondary text |
| `--c-dim` | `#4A4A4A` | Footer text |
| `--c-sub` | `#C8C4BD` | Subtitles, supporting text |

### Typography Scale

| Token | Value | Use |
|-------|-------|-----|
| `--font-display` | 4.236rem | Single hero number |
| `--font-h1` | 2.618rem | Slide title, divider title |
| `--font-h2` | 1.618rem | Section heading (h2/h3) |
| `--font-subtitle` | 1.25rem | Subtitle, card headings |
| `--font-body` | 1rem | Body text, bullets |
| `--font-small` | 0.85rem | Labels, table cells |
| `--font-caption` | 0.75rem | Uppercase labels |
| `--font-micro` | 0.618rem | Footer |
| `--font-metric` | 2.618rem | Metric card values |
| `--font-metric-lg` | 3.25rem | Scenario card hero numbers |

### Spacing Scale

| Token | Value |
|-------|-------|
| `--space-xs` | 0.5rem (8px) |
| `--space-sm` | 1rem (16px) |
| `--space-md` | 1.5rem (24px) |
| `--space-lg` | 2rem (32px) |
| `--space-xl` | 3rem (48px) |
| `--space-2xl` | 4rem (64px) |

### Component Dimensions

| Token | Value | Previous |
|-------|-------|----------|
| `--bar-height` | 2rem (32px) | 28px |
| `--bar-label-width` | 10rem | 120px/180px |
| `--bar-value-width` | 6.5rem | 60px/100px |
| `--timeline-dot` | 2.25rem (36px) | 28px |
| `--card-radius` | 0.5rem (8px) | 8px |
| `--heatmap-header-width` | 11rem | 140px/180px |

### CSS Classes

| Class | Component |
|-------|-----------|
| `.metric-grid`, `.metric-grid--4`, `.metric-grid--6` | Metric grid container (3/4/6 columns) |
| `.metric-card`, `.metric-value`, `.metric-label` | Individual metric card |
| `.two-col`, `.card`, `.card-list` | Two-column comparison layout |
| `.bar-row`, `.bar-label`, `.bar-track`, `.bar-fill`, `.bar-value` | Horizontal bar chart |
| `.bar-fill--muted`, `.bar-value--muted` | Non-primary bar variants |
| `.scenario-row`, `.scenario-card`, `.scenario-card--primary` | ROI scenario cards |
| `.scenario-label`, `.scenario-value`, `.scenario-detail`, `.scenario-sub` | Scenario card parts |
| `.badge-row`, `.badge`, `.badge--primary` | Title card badges |
| `.divider-slide`, `.divider-act`, `.divider-title`, `.divider-sub` | Section divider |
| `.timeline-container`, `.timeline-line`, `.timeline-phases`, `.timeline-phase` | Timeline layout |
| `.timeline-dot`, `.timeline-dot--active`, `.timeline-dot--future` | Timeline dots |
| `.timeline-phase-title`, `.timeline-phase-date`, `.timeline-phase-desc` | Timeline text |
| `.heatmap`, `.heatmap--3col`, `.heatmap--2col` | Heatmap container |
| `.heatmap-corner`, `.heatmap-colhead`, `.heatmap-rowhead`, `.heatmap-cell` | Heatmap parts |
| `.heatmap-cell--high`, `.heatmap-cell--med`, `.heatmap-cell--low` | Heatmap intensity |
| `.decision-item`, `.decision-number`, `.decision-title`, `.decision-desc` | Decision/ask layout |
| `.pause-context` | Center pause supporting text |
| `.closing-tagline`, `.closing-attribution` | Closing card text |
