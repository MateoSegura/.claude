# Center Pause
> Use as a dramatic pause, rhetorical question, or breathing moment between dense sections. Black screen with a single centered statement. Minimal content — let the audience absorb what came before.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h1` | 2.618rem | Centered quote/statement (h1) |
| `--c-text` | #E8E4DD | Quote text color (via h1 base style) |
| `--c-muted` | #6B6B6B | Supporting context text |
| `--font-small` | 0.85rem | Context text size |
| `--space-lg` | 2rem | Top margin on context text |

## Template

```markdown
---
layout: center
class: text-center
---

# "Quote or key statement"

<div class="pause-context">
  One line of supporting context
</div>

<!--
Pause. Let this land.

"Read the quote aloud or paraphrase it."

Hold for 2-3 seconds of silence before advancing.

Transition: "With that in mind, let's look at..."
-->
```

### Classes Used

| Class | Purpose |
|-------|---------|
| `layout: center` | Slidev built-in layout — vertically and horizontally centers all content |
| `class: text-center` | Slidev utility — applies `text-align: center` to the slide |
| `.pause-context` | Supporting text below the quote — muted color, small font, top margin via design system |

## Sizing Notes

- **Title**: One sentence maximum. Wrap in quotation marks for rhetorical statements or direct quotes. Keep under 10 words for impact.
- **Context line**: Optional. One short line only (under 10 words). If the quote speaks for itself, omit the context div entirely.
- **No other content**: This slide type should contain nothing else — no bullets, no images, no charts. Its purpose is negative space.
- **Speaker notes**: The presenter should pause for 2-3 seconds of silence. This slide is a pacing tool, not an information slide.
- **Placement**: Use between major sections, after a surprising data point, or before a key decision slide. Never use two pause slides in sequence.
- **Frontmatter**: Both `layout: center` and `class: text-center` are required in the YAML block. The layout handles vertical centering; the class handles horizontal centering.

## Example

```markdown
---
layout: center
class: text-center
---

# "Every hour of downtime costs us $14,000"

<div class="pause-context">
  Based on Q3 incident data across all production services
</div>

<!--
Pause. Let the number sink in.

"Every hour of downtime costs us fourteen thousand dollars."

Hold for a full beat. Make eye contact with the CFO.

Transition: "So the question isn't whether we can afford
to modernize — it's whether we can afford not to."
-->
```
