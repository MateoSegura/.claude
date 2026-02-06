# Two-Column Comparison
> Use for before/after, current/proposed, problem/solution, or any side-by-side comparison. Each column is a card with a heading and a structured list.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h2` | 1.618rem | Slide title (h2) — gold via base styles |
| `--font-h2` | 1.618rem | Card heading (h3) — gold via base styles |
| `--c-gold` | #D4A853 | Card headings, bold text |
| `--c-text` | #E8E4DD | List item text |
| `--c-card` | #2A2A2A | Card background |
| `--c-border` | #3A3A3A | Card border, list item separators |
| `--card-radius` | 0.5rem | Card corner rounding |
| `--space-lg` | 2rem | Card padding, grid gap |
| `--space-sm` | 1rem | Grid top margin |
| `--space-xs` | 0.5rem | List item vertical padding |

## Template

```html
---

## Slide Title

<div class="two-col">
  <div class="card">
    <h3>Left Column Title</h3>
    <ul class="card-list">
      <li><strong>Bold lead</strong> — supporting detail</li>
      <li>Second point with specifics</li>
      <li>Third point</li>
    </ul>
  </div>
  <div class="card">
    <h3>Right Column Title</h3>
    <ul class="card-list">
      <li><strong>Bold lead</strong> — supporting detail</li>
      <li>Second point with specifics</li>
      <li>Third point</li>
    </ul>
  </div>
</div>

<!--
Speaker notes here.
-->
```

### Classes Used

| Class | Purpose |
|-------|---------|
| `.two-col` | 2-column grid container with `--space-lg` gap and `--space-sm` top margin |
| `.card` | Card with background, border, radius, and padding |
| `.card-list` | Unstyled list with bottom-border separators between items |

## Sizing Notes

- **3-5 items per column**. Fewer is better — match the count across both columns when possible for visual balance.
- **Use HTML lists only**: Because these lists are inside `<div>` elements, you must use `<ul class="card-list"><li>` tags. Markdown bullet syntax (`- item`) inside HTML divs breaks Vue compilation.
- **Bold the lead phrase**: Use `<strong>` for the first 2-3 words of each item. These render in gold and create a scannable left edge.
- **Parallel structure**: Both columns should use the same grammatical pattern. If the left column uses noun phrases, the right column should too.
- **Column headings**: Use `<h3>` inside each card. The base h3 style applies gold color and `--font-h2` size automatically.
- **Blockquote after**: You can add a `> "takeaway"` blockquote after the closing `</div>` of `.two-col` (at root level, where markdown is safe).

## Example

### Before / After Comparison

```html
---

## Developer Experience: Before vs. After

<div class="two-col">
  <div class="card">
    <h3>Today</h3>
    <ul class="card-list">
      <li><strong>2-week deploys</strong> — manual release train with change board</li>
      <li><strong>Shared staging</strong> — one environment, constant conflicts</li>
      <li><strong>4h incident MTTR</strong> — SSH into production to diagnose</li>
    </ul>
  </div>
  <div class="card">
    <h3>Proposed</h3>
    <ul class="card-list">
      <li><strong>Daily deploys</strong> — automated CI/CD with canary rollouts</li>
      <li><strong>Preview environments</strong> — per-PR ephemeral namespaces</li>
      <li><strong>15min MTTR</strong> — automated rollback on health-check failure</li>
    </ul>
  </div>
</div>

> "Same team, same product — just better tooling."

<!--
"Let me show you what changes for the team day-to-day."

Walk through left column first (pain), then right column (solution).
Let the contrast do the persuading.

Transition: "Now let's look at what this costs..."
-->
```

### Problem / Solution Framing

```html
---

## Support Escalation Pipeline

<div class="two-col">
  <div class="card">
    <h3>The Problem</h3>
    <ul class="card-list">
      <li><strong>340 tickets/week</strong> — 60% are repeat questions</li>
      <li><strong>No self-service</strong> — every issue requires a human</li>
      <li><strong>$42 cost per ticket</strong> — fully loaded support agent time</li>
    </ul>
  </div>
  <div class="card">
    <h3>The Solution</h3>
    <ul class="card-list">
      <li><strong>AI triage layer</strong> — resolves 60% of repeats automatically</li>
      <li><strong>Knowledge base</strong> — searchable docs reduce inbound volume</li>
      <li><strong>$8 cost per ticket</strong> — blended human + AI resolution</li>
    </ul>
  </div>
</div>

<!--
"Here's the support funnel today versus what we're proposing."

Emphasize the cost-per-ticket contrast — $42 down to $8.
That's the number the CFO will remember.

Transition: "Let me show you the ROI model..."
-->
```
