# Data Table
> Use when presenting structured comparisons, feature matrices, pricing tiers, or any data best read as rows and columns. Max 5 rows, 4 columns.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h2` | 1.618rem | Slide title (h2, gold) |
| `--font-small` | 0.85rem | Table cell and header text via `font-size` wrapper |
| `--c-gold` | #D4A853 | Table header text color (via `th` base style) |
| `--c-card` | #2A2A2A | Table header background (via `th` base style) |
| `--c-text` | #E8E4DD | Table cell text (via `td` base style) |
| `--c-border` | #3A3A3A | Cell borders (via `th`/`td` base style) |
| `--space-xs` | 0.5rem | Cell padding (via base table styles) |
| `--space-sm` | 1rem | Cell padding, blockquote padding |

## Template

```html
---

## Slide Title

<div style="font-size: var(--font-small);">

| Column A | Column B | Column C | Column D |
|----------|----------|----------|----------|
| Data 1   | Data 2   | Data 3   | Data 4   |
| Data 5   | Data 6   | Data 7   | Data 8   |
| Data 9   | Data 10  | Data 11  | Data 12  |

</div>

> "One-sentence takeaway that tells the audience what to see in the data."

<!--
"Here's how the numbers break down."

Walk through the key column or row. Highlight the standout cell.

Transition: "And what this means in practice is..."
-->
```

### Classes Used

| Class / Style | Purpose |
|---------------|---------|
| `<div style="font-size: var(--font-small);">` | Wraps the markdown table to set cell font size; the design system base `th`/`td` styles handle colors, padding, and borders automatically |
| `blockquote` (base style) | Gold left-border callout for the takeaway line |

## Sizing Notes

- **Rows**: Maximum 5 data rows (excluding header). If more rows are needed, split across two slides or consolidate into categories.
- **Columns**: Maximum 4 columns. Wider tables overflow the safe zone. If a fifth column is essential, abbreviate headers to 1-2 words and test rendering.
- **Cell content**: Keep each cell under 4 words. Use numbers, short labels, or status indicators (checkmarks, dashes) rather than sentences.
- **Header row**: Column names should be 1-3 words. The design system styles headers with gold text on card background automatically.
- **Alternating rows**: Even rows receive a subtle `rgba(42,42,42,0.4)` background via the `tr:nth-child(even)` base style. No extra class needed.
- **Takeaway**: Always include a blockquote below the table. The audience needs to know what conclusion to draw — never present a table without interpretation.
- **Font sizing**: The `<div>` wrapper with `font-size: var(--font-small)` is required. Without it, table text renders at body size and risks overflow on data-dense slides.

## Example

```html
---

## Vendor Comparison

<div style="font-size: var(--font-small);">

| Criteria | AWS EKS | GCP GKE | Azure AKS |
|----------|---------|---------|-----------|
| Setup time | 2-3 days | 1-2 days | 2-4 days |
| Monthly cost | $1,840 | $1,620 | $1,750 |
| SLA | 99.95% | 99.95% | 99.95% |
| Team expertise | High | Medium | Low |

</div>

> "GKE delivers the fastest setup at the lowest cost — and our team has moderate experience already."

<!--
"Let me walk you through how the three leading platforms compare
on the criteria that matter most for our timeline."

Point to the cost row first — that's where eyes go.
Then highlight setup time as the differentiator.

Transition: "Based on this, here's our recommendation..."
-->
```
