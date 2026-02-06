# Bar Chart
> When to use: Comparing quantities across categories — revenue by segment, time allocation, feature adoption, cost breakdown. Best for 3-6 items where relative magnitude matters more than exact values.

## Design Tokens Used

| Token | Value | Role |
|-------|-------|------|
| `--c-gold` | `#D4A853` | Primary bar fill, primary value text |
| `--c-muted` | `#6B6B6B` | Secondary bar fill (`.bar-fill--muted`) |
| `--c-sub` | `#C8C4BD` | Bar labels, muted value text |
| `--c-card` | `#2A2A2A` | Bar track background |
| `--bar-height` | `2rem` | Bar thickness |
| `--bar-label-width` | `10rem` | Left label column width |
| `--bar-value-width` | `6.5rem` | Right value column width |
| `--space-sm` | `1rem` | Row spacing (margin-bottom) |
| `--font-small` | `0.85rem` | Label and value text size |

## Template

```html
---

## Slide Title

<div class="bar-row">
  <div class="bar-label">Category A</div>
  <div class="bar-track">
    <div class="bar-fill" style="width: 85%"></div>
  </div>
  <div class="bar-value">$850K</div>
</div>

<div class="bar-row">
  <div class="bar-label">Category B</div>
  <div class="bar-track">
    <div class="bar-fill" style="width: 62%"></div>
  </div>
  <div class="bar-value">$620K</div>
</div>

<div class="bar-row">
  <div class="bar-label">Category C</div>
  <div class="bar-track">
    <div class="bar-fill bar-fill--muted" style="width: 45%"></div>
  </div>
  <div class="bar-value bar-value--muted">$450K</div>
</div>

<div class="bar-row">
  <div class="bar-label">Category D</div>
  <div class="bar-track">
    <div class="bar-fill bar-fill--muted" style="width: 30%"></div>
  </div>
  <div class="bar-value bar-value--muted">$300K</div>
</div>

> "One-sentence takeaway highlighting the key insight"

---
```

### Variant: All-primary bars (no muted)

Use when all bars represent the same type — no need to visually demote any row.

```html
<div class="bar-row">
  <div class="bar-label">Label</div>
  <div class="bar-track">
    <div class="bar-fill" style="width: XX%"></div>
  </div>
  <div class="bar-value">Value</div>
</div>
```

### Variant: Mixed primary + muted

Use when top items are emphasized and bottom items are supporting context. Add `bar-fill--muted` to the fill and `bar-value--muted` to the value.

```html
<div class="bar-row">
  <div class="bar-label">Label</div>
  <div class="bar-track">
    <div class="bar-fill bar-fill--muted" style="width: XX%"></div>
  </div>
  <div class="bar-value bar-value--muted">Value</div>
</div>
```

## Sizing Notes

- **Max bars**: 6. Beyond 6, the slide becomes crowded and bars lose visual weight.
- **Width percentages**: Calculate relative to the largest value. Largest bar = `style="width: 100%"`, others proportional.
- **Primary vs muted**: Use gold (`.bar-fill`) for the top 2-3 items you want the audience to focus on. Use muted (`.bar-fill--muted`) for context/comparison items.
- **Label length**: Keep labels under 20 characters. Longer labels push into bar track space.
- **Value formatting**: Use consistent units (`$XXK`, `XX%`, `XX hrs`). Right-aligned by default.
- **Only inline style allowed**: `style="width: XX%"` on `.bar-fill` to set bar length. All other styling comes from CSS classes.
- **Blockquote**: Include a one-line takeaway at the bottom when the chart supports a clear narrative.
- **4 bars**: Comfortable default. Leaves room for the blockquote summary.
- **5-6 bars**: Still works but reduce or omit the blockquote to avoid overflow.

## Example

```html
---

## Engineering Time Allocation

<div class="bar-row">
  <div class="bar-label">Feature Dev</div>
  <div class="bar-track">
    <div class="bar-fill" style="width: 100%"></div>
  </div>
  <div class="bar-value">42%</div>
</div>

<div class="bar-row">
  <div class="bar-label">Code Review</div>
  <div class="bar-track">
    <div class="bar-fill" style="width: 57%"></div>
  </div>
  <div class="bar-value">24%</div>
</div>

<div class="bar-row">
  <div class="bar-label">Bug Fixes</div>
  <div class="bar-track">
    <div class="bar-fill bar-fill--muted" style="width: 40%"></div>
  </div>
  <div class="bar-value bar-value--muted">17%</div>
</div>

<div class="bar-row">
  <div class="bar-label">Meetings</div>
  <div class="bar-track">
    <div class="bar-fill bar-fill--muted" style="width: 26%"></div>
  </div>
  <div class="bar-value bar-value--muted">11%</div>
</div>

<div class="bar-row">
  <div class="bar-label">Documentation</div>
  <div class="bar-track">
    <div class="bar-fill bar-fill--muted" style="width: 14%"></div>
  </div>
  <div class="bar-value bar-value--muted">6%</div>
</div>

> "66% of engineering time goes to feature development and code review"

---
```
