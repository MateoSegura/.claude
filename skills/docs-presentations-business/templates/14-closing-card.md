# Closing Card
> Use as the final slide of every deck. Restates the presentation title, delivers a closing tagline or call to action, and includes the generation attribution footer.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h1` | 2.618rem | Presentation title (h1) |
| `--font-subtitle` | 1.25rem | Closing tagline text size |
| `--font-small` | 0.85rem | Attribution footer text size |
| `--c-text` | #E8E4DD | Title color (via h1 base style) |
| `--c-sub` | #C8C4BD | Tagline text color |
| `--c-muted` | #6B6B6B | Attribution text color |
| `--space-lg` | 2rem | Tagline top margin |
| `--space-xl` | 3rem | Attribution top margin |

## Template

```html
---

# Presentation Title

<div class="closing-tagline">
  Closing tagline or call to action
</div>

<div class="closing-attribution">
  Generated with Claude Code
</div>

<!--
"Thank you. I'd love to take your questions."

If no questions come immediately, restate the single most
important number or decision from the deck.

Fallback: "The key takeaway is [one sentence]."
-->
```

### Classes Used

| Class | Purpose |
|-------|---------|
| `.closing-tagline` | Subtitle-sized text in muted tone — delivers the final message at `--font-subtitle` size with `--c-sub` color and `--space-lg` top margin |
| `.closing-attribution` | Small muted footer — "Generated with Claude Code" at `--font-small` size with `--c-muted` color and `--space-xl` top margin |

## Sizing Notes

- **Title**: Must match the title from the opening title card slide exactly. Uses standard `# h1` markdown — the design system handles font-size and color.
- **Tagline**: One line, maximum 15 words. This is a call to action, a summary statement, or a forward-looking message. Not a repeat of the subtitle.
- **Attribution**: Always "Generated with Claude Code" — do not modify this text. The `.closing-attribution` class positions it with generous top margin to separate it visually from the tagline.
- **No badges**: Unlike the title card, the closing card omits badges. The audience already has the context; this slide is about the final impression.
- **No blockquotes or data**: This is a clean, minimal ending. If there is a final data point to land, put it on the preceding slide or in a center-pause slide before this one.
- **Speaker notes**: Prepare a one-sentence restatement of the deck's core message as a fallback if questions don't come immediately.

## Example

```html
---

# Cloud Infrastructure Proposal

<div class="closing-tagline">
  Ready to cut deployment time by 80% — pending your approval to proceed
</div>

<div class="closing-attribution">
  Generated with Claude Code
</div>

<!--
"Thank you for your time. We're ready to begin the migration
as soon as we have the green light."

If asked about risk: "We've built in rollback at every phase —
slide 9 has the full timeline."

Fallback: "The bottom line: $240K investment, $890K annual savings,
paid back in four months."
-->
```
