# Metric Grid
> Use to display 3, 4, or 6 key metrics as a card grid. Best for KPI summaries, financial snapshots, or before/after comparisons where each data point stands alone.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h2` | 1.618rem | Slide title (h2) |
| `--font-metric` | 2.618rem | Metric value size |
| `--font-caption` | 0.75rem | Metric label — uppercase |
| `--c-gold` | #D4A853 | Metric value color |
| `--c-muted` | #6B6B6B | Metric label color |
| `--c-card` | #2A2A2A | Card background |
| `--c-border` | #3A3A3A | Card border |
| `--card-radius` | 0.5rem | Card corner rounding |
| `--space-md` | 1.5rem | Grid gap and top margin |
| `--space-lg` | 2rem | Card internal padding |
| `--space-xs` | 0.5rem | Label top margin |

## Template — 3 Cards (Default)

```html
---

## Slide Title

<div class="metric-grid">
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
</div>

<!--
Speaker notes here.
-->
```

## Template — 6 Cards

```html
---

## Slide Title

<div class="metric-grid metric-grid--6">
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">Value</div>
    <div class="metric-label">Label</div>
  </div>
</div>

<!--
Speaker notes here.
-->
```

### Classes Used

| Class | Purpose |
|-------|---------|
| `.metric-grid` | 3-column grid container with `--space-md` gap and top margin |
| `.metric-grid--4` | Override to 4-column layout (add alongside `.metric-grid`) |
| `.metric-grid--6` | Override to 3-column layout for 6 cards / 2 rows (add alongside `.metric-grid`) |
| `.metric-card` | Card with background, border, padding, centered text |
| `.metric-value` | Large gold mono number — `--font-metric` size |
| `.metric-label` | Small uppercase muted label below the value |

## Sizing Notes

- **3 cards**: Use `.metric-grid` alone. Default 3-column layout. Best for hero stats.
- **4 cards**: Use `.metric-grid .metric-grid--4`. Fits tighter — keep values short (4 characters max).
- **6 cards**: Use `.metric-grid .metric-grid--6`. Renders as 3 columns x 2 rows. This is the hard maximum from the content budget.
- **Values**: Keep to 1-5 characters. Use abbreviations: `$2.4M`, `340ms`, `99.9%`, `12x`. The mono font and gold color make these the focal point.
- **Labels**: 1-3 words, all uppercase (applied by CSS). Describe what the number means: `REVENUE`, `P99 LATENCY`, `UPTIME SLA`.
- **No mixing**: Every card in a grid should be the same type of data (all financial, all performance, all counts). Mix types across slides, not within a grid.

## Example

### 3-Card Financial Summary

```html
---

## Investment at a Glance

<div class="metric-grid">
  <div class="metric-card">
    <div class="metric-value">$240K</div>
    <div class="metric-label">Total Investment</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">14mo</div>
    <div class="metric-label">Payback Period</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">3.2x</div>
    <div class="metric-label">3-Year ROI</div>
  </div>
</div>

<!--
"Three numbers tell the whole story."

Point to each card left-to-right as you narrate.

Transition: "Let me break down how we get to that ROI..."
-->
```

### 6-Card Operational Snapshot

```html
---

## Current Platform Health

<div class="metric-grid metric-grid--6">
  <div class="metric-card">
    <div class="metric-value">99.2%</div>
    <div class="metric-label">Uptime (90d)</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">340ms</div>
    <div class="metric-label">P95 Latency</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">12</div>
    <div class="metric-label">Incidents / Month</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">4.2h</div>
    <div class="metric-label">Mean Time to Recover</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">23%</div>
    <div class="metric-label">CPU Utilization</div>
  </div>
  <div class="metric-card">
    <div class="metric-value">$18K</div>
    <div class="metric-label">Monthly Infra Spend</div>
  </div>
</div>

<!--
"Here's a snapshot of where we stand today."

Top row is reliability — bottom row is efficiency.
Let the audience absorb before narrating.

Transition: "Two numbers should jump out at you..."
-->
```
