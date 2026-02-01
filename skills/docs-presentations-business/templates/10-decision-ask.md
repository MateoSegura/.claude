# Decision Ask
> When to use: The final "ask" slide where you present specific decisions for the audience to approve. Large-format numbered items make each decision unmissable. Use at the end of a business case or proposal deck.

## Design Tokens Used

| Token | Value | Role |
|-------|-------|------|
| `--c-gold` | `#D4A853` | Decision number, blockquote border |
| `--c-text` | `#E8E4DD` | Decision title |
| `--c-sub` | `#C8C4BD` | Decision description |
| `--c-card` | `#2A2A2A` | Blockquote background |
| `--font-h1` | `2.618rem` | Decision number size |
| `--font-subtitle` | `1.25rem` | Decision title size |
| `--font-body` | `1rem` | Decision description size |
| `--space-lg` | `2rem` | Spacing between decision items |
| `--space-xs` | `0.5rem` | Gap between number and text, desc top margin |

## Template

```html
---

## Slide Title

<div class="decision-item">
  <span class="decision-number">1.</span>
  <div>
    <div class="decision-title">First decision title</div>
    <div class="decision-desc">One-line description with a specific number or scope</div>
  </div>
</div>

<div class="decision-item">
  <span class="decision-number">2.</span>
  <div>
    <div class="decision-title">Second decision title</div>
    <div class="decision-desc">One-line description with a specific number or scope</div>
  </div>
</div>

> "Summary statement that makes the decision feel obvious"

---
```

## Sizing Notes

- **Decision count**: 1-3 decisions per slide. 2 is the sweet spot -- clear binary or paired ask. 3 is the maximum before the slide becomes cluttered. For 4+ decisions, split across slides or consolidate.
- **Number format**: Use `1.` `2.` `3.` with the trailing period. The number is rendered in `--font-h1` (2.618rem) JetBrains Mono bold gold -- it acts as a visual anchor.
- **Title**: One line, under 8 words. Action-oriented: "Approve the pilot budget", "Assign a project lead", "Set the Q2 launch date". Uses `--font-subtitle` (1.25rem), semi-bold, primary text color.
- **Description**: One line, under 20 words. Include a specific number, dollar amount, timeline, or scope to make it concrete: "$50K for 3-month pilot with 5-person team". Uses `--font-body` (1rem), subtitle color.
- **Blockquote**: Required. This is the "make it easy" line -- a single stat or framing that makes approval feel obvious. Positioned at the bottom of the slide via standard blockquote styling (gold left border, card background).
- **Alignment**: `.decision-item` uses flexbox with `align-items: baseline` so the number baseline aligns with the title text.
- **No inline styles**: All styling comes from CSS classes. The old reference.md pattern used inline styles for font-size, color, font-family, margin, display, and gap -- all of these are now handled by the CSS classes.
- **Vertical rhythm**: Each `.decision-item` has `--space-lg` (2rem) bottom margin. With 2 items + blockquote, this fills the slide comfortably without overflow.

## Example

```html
---

## Two Decisions

<div class="decision-item">
  <span class="decision-number">1.</span>
  <div>
    <div class="decision-title">Approve the AI tooling pilot</div>
    <div class="decision-desc">$48K budget for a 3-month trial across the platform engineering team</div>
  </div>
</div>

<div class="decision-item">
  <span class="decision-number">2.</span>
  <div>
    <div class="decision-title">Designate a pilot lead</div>
    <div class="decision-desc">Senior engineer to own metrics, feedback loops, and the go/no-go report</div>
  </div>
</div>

> "At 3.2x projected ROI, the pilot pays for itself before month 4"

---
```
