# Text Slide
> Use for narrative points, key arguments, or context-setting where a visual layout is unnecessary. The workhorse slide for verbal-heavy moments.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h2` | 1.618rem | Slide title (h2) — gold color via base styles |
| `--font-body` | 1rem | Bullet text |
| `--c-text` | #E8E4DD | Bullet text color |
| `--c-gold` | #D4A853 | Title color, bold text, blockquote border |
| `--c-card` | #2A2A2A | Blockquote background |

## Template

```markdown
---

## Slide Title

- First point — the most important claim (max 20 words)
- Second point with **bold emphasis** on the key phrase
- Third point with a specific number or outcome

> "One-sentence takeaway or quotation that anchors the message."
```

### Elements Used

| Element | Styling Source |
|---------|---------------|
| `## Title` | Base h2 style — gold, 1.618rem, 600 weight |
| `- Bullet` | Base li style — body size, primary text color |
| `**bold**` | Base strong style — gold color |
| `> Blockquote` | Base blockquote style — gold left border, card background |

## Sizing Notes

- **3 bullets maximum**. This is a hard limit from the content budget. If you have 4+ points, split across two slides or use a two-column layout instead.
- **20 words per bullet**. Decision-makers scan; they do not read. Front-load the key phrase.
- **40 total words** across the slide (excluding title and blockquote).
- **Bold sparingly**: One bold phrase per bullet at most. Bold renders in gold, so overuse dilutes emphasis.
- **Blockquote is optional**: Use it for a memorable stat, a direct quote, or a one-line summary. Omit if the bullets are self-sufficient.
- **No HTML required**: This is a pure-markdown slide. All styling comes from the design system's base typography rules.

## Example

```markdown
---

## Why Kubernetes Over Bare Metal

- **3x deployment frequency** — teams ship daily instead of weekly release trains
- Infrastructure costs drop 40% through bin-packing and autoscaling
- On-call incidents reduced from 12/month to 3 with self-healing pods

> "We spend more time firefighting servers than building product."
```

```markdown
---

## What Customers Are Telling Us

- **"I can't find anything"** — search NPS dropped 18 points in Q3
- Support tickets about navigation up 34% quarter-over-quarter
- Three enterprise renewals cited UX as the deciding factor

> "The product works. The problem is people can't reach it."
```
