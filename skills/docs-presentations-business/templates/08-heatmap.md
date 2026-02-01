# Heatmap
> When to use: Readiness assessments, scoring matrices, capability grids, risk maps, or any evaluation where rows are categories and columns are criteria. Cell shading communicates intensity at a glance.

## Design Tokens Used

| Token | Value | Role |
|-------|-------|------|
| `--c-gold` | `#D4A853` | Column headers, high-intensity cell text |
| `--c-sub` | `#C8C4BD` | Row headers |
| `--c-muted` | `#6B6B6B` | Corner cell text |
| `--c-border` | `#3A3A3A` | Grid gap color (via 2px gap on background) |
| `--heatmap-header-width` | `11rem` | Row header column width |
| `--font-small` | `0.85rem` | All heatmap text |
| `--space-md` | `1.5rem` | Top margin |
| `rgba(212,168,83,0.4)` | -- | High-intensity cell background |
| `rgba(212,168,83,0.15)` | -- | Medium-intensity cell background |
| `rgba(107,107,107,0.2)` | -- | Low-intensity cell background |

## Template (3-column)

```html
---

## Slide Title

<div class="heatmap heatmap--3col">
  <div class="heatmap-corner"></div>
  <div class="heatmap-colhead">Column A</div>
  <div class="heatmap-colhead">Column B</div>
  <div class="heatmap-colhead">Column C</div>

  <div class="heatmap-rowhead">Row 1</div>
  <div class="heatmap-cell heatmap-cell--high">High</div>
  <div class="heatmap-cell heatmap-cell--med">Med</div>
  <div class="heatmap-cell heatmap-cell--low">Low</div>

  <div class="heatmap-rowhead">Row 2</div>
  <div class="heatmap-cell heatmap-cell--med">Med</div>
  <div class="heatmap-cell heatmap-cell--high">High</div>
  <div class="heatmap-cell heatmap-cell--med">Med</div>

  <div class="heatmap-rowhead">Row 3</div>
  <div class="heatmap-cell heatmap-cell--low">Low</div>
  <div class="heatmap-cell heatmap-cell--med">Med</div>
  <div class="heatmap-cell heatmap-cell--high">High</div>
</div>

> "One-sentence takeaway about the pattern"

---
```

## Template (2-column)

```html
---

## Slide Title

<div class="heatmap heatmap--2col">
  <div class="heatmap-corner"></div>
  <div class="heatmap-colhead">Column A</div>
  <div class="heatmap-colhead">Column B</div>

  <div class="heatmap-rowhead">Row 1</div>
  <div class="heatmap-cell heatmap-cell--high">High</div>
  <div class="heatmap-cell heatmap-cell--low">Low</div>

  <div class="heatmap-rowhead">Row 2</div>
  <div class="heatmap-cell heatmap-cell--med">Med</div>
  <div class="heatmap-cell heatmap-cell--high">High</div>

  <div class="heatmap-rowhead">Row 3</div>
  <div class="heatmap-cell heatmap-cell--low">Low</div>
  <div class="heatmap-cell heatmap-cell--med">Med</div>
</div>

> "One-sentence takeaway about the pattern"

---
```

## Sizing Notes

- **Grid variants**: Use `.heatmap--3col` for 3 data columns, `.heatmap--2col` for 2 data columns. The first column is always the row header at `--heatmap-header-width` (11rem).
- **Max rows**: 5 data rows. Beyond 5, the cells become too small to read at presentation scale.
- **Max columns**: 3 data columns (plus row header). For wider matrices, split into multiple slides or consolidate columns.
- **Cell text**: Keep cell content to 1-3 words. Scores (`9/10`), labels (`High`, `Ready`), or short phrases (`On Track`) work best.
- **Intensity levels**: Exactly 3 levels. Map your data to high/med/low before building the grid.
  - `.heatmap-cell--high`: Gold-tinted, bold text. Use for top scores, strong readiness, high priority.
  - `.heatmap-cell--med`: Subtle gold tint. Use for moderate/acceptable scores.
  - `.heatmap-cell--low`: Grey tint. Use for gaps, low scores, not-ready areas.
- **Corner cell**: The `.heatmap-corner` top-left cell is intentionally blank. Do not put text in it.
- **Row headers**: Left-aligned, semi-bold. Keep under 15 characters.
- **Column headers**: Center-aligned, gold colored. Keep under 12 characters.
- **No inline styles**: All styling comes from CSS classes. No inline styles needed.

## Example

```html
---

## AI Readiness Assessment

<div class="heatmap heatmap--3col">
  <div class="heatmap-corner"></div>
  <div class="heatmap-colhead">Data Quality</div>
  <div class="heatmap-colhead">Team Skills</div>
  <div class="heatmap-colhead">Infrastructure</div>

  <div class="heatmap-rowhead">Code Review</div>
  <div class="heatmap-cell heatmap-cell--high">Strong</div>
  <div class="heatmap-cell heatmap-cell--high">Strong</div>
  <div class="heatmap-cell heatmap-cell--med">Adequate</div>

  <div class="heatmap-rowhead">Test Generation</div>
  <div class="heatmap-cell heatmap-cell--med">Adequate</div>
  <div class="heatmap-cell heatmap-cell--med">Adequate</div>
  <div class="heatmap-cell heatmap-cell--high">Strong</div>

  <div class="heatmap-rowhead">Documentation</div>
  <div class="heatmap-cell heatmap-cell--high">Strong</div>
  <div class="heatmap-cell heatmap-cell--low">Gap</div>
  <div class="heatmap-cell heatmap-cell--med">Adequate</div>

  <div class="heatmap-rowhead">Architecture</div>
  <div class="heatmap-cell heatmap-cell--low">Gap</div>
  <div class="heatmap-cell heatmap-cell--med">Adequate</div>
  <div class="heatmap-cell heatmap-cell--low">Gap</div>
</div>

> "Code review and test generation are ready now; architecture needs 2-month ramp"

---
```
