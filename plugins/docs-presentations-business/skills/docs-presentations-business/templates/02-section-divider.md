# Section Divider
> Use between major sections of the deck to signal a topic shift. Creates a centered, full-height pause slide.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h1` | 2.618rem | Divider title size |
| `--font-small` | 0.85rem | Act/section label size |
| `--font-body` | 1rem | Subtitle size |
| `--c-text` | #E8E4DD | Title color |
| `--c-muted` | #6B6B6B | Act label + subtitle color |
| `--space-xs` | 0.5rem | Gap between label and title, title and subtitle |

## Template

```html
---

<div class="divider-slide">
  <div class="divider-act">Section Label</div>
  <div class="divider-title">Section Title</div>
  <div class="divider-sub">One-line subtitle or framing question</div>
</div>

<!--
"Transition line that bridges from the previous section."

Pause briefly. Let the audience reset.

Transition: "So let's look at..."
-->
```

### Classes Used

| Class | Purpose |
|-------|---------|
| `.divider-slide` | Full-height flex column, centered both axes |
| `.divider-act` | Uppercase label above the title — small, muted, wide letter-spacing (0.3em) |
| `.divider-title` | Large title text — uses `--font-h1`, bold, primary text color |
| `.divider-sub` | Supporting line below the title — body size, muted color |

## Sizing Notes

- **No markdown headings**: This slide uses only `<div>` elements with divider classes. Do not add `#` or `##` headings — the classes handle all typography.
- **Act label**: Short — typically "Part I", "Act 2", "Section 3", or a category name like "The Problem" or "Financials". All-caps is applied via CSS `text-transform`.
- **Title**: 2-5 words. This is the section name the audience anchors on.
- **Subtitle**: Optional. One sentence max. If omitted, remove the `.divider-sub` div entirely.
- **Vertical centering**: The `.divider-slide` class handles full centering — no layout frontmatter needed.

## Example

```html
---

<div class="divider-slide">
  <div class="divider-act">Part II</div>
  <div class="divider-title">The Cost of Inaction</div>
  <div class="divider-sub">What happens if we stay on the current path</div>
</div>

<!--
"We've seen the opportunity. Now let's talk about what
it costs us to do nothing."

Pause here — let the title sit for 2-3 seconds.

Transition: "The first cost is one we're already paying..."
-->
```
