# Title Card
> Use as the first slide of every deck. Sets the presentation title, subtitle, and context badges.

## Design Tokens Used

| Token | Value | Purpose |
|-------|-------|---------|
| `--font-h1` | 2.618rem | Presentation title (h1) |
| `--font-subtitle` | 1.25rem | Subtitle text size |
| `--c-text` | #E8E4DD | Title color (via h1 base style) |
| `--c-sub` | #C8C4BD | Subtitle color |
| `--c-gold` | #D4A853 | Primary badge border + text |
| `--c-card` | #2A2A2A | Badge background |
| `--c-border` | #3A3A3A | Default badge border |
| `--space-lg` | 2rem | Subtitle top margin |
| `--space-xl` | 3rem | Badge row top margin |
| `--space-sm` | 1rem | Gap between badges |

## Template

```html
---

# Presentation Title

<div style="margin-top: var(--space-lg);">
  <p style="font-size: var(--font-subtitle); color: var(--c-sub); font-weight: 400;">
    Subtitle or one-line value proposition
  </p>
</div>

<div class="badge-row">
  <span class="badge badge--primary">Primary Badge</span>
  <span class="badge">Secondary Badge</span>
  <span class="badge">Secondary Badge</span>
</div>

<!--
"Opening line — what you say as the slide appears."

Set the stage. One sentence of context.

Transition: "Let me start by showing you..."
-->
```

### Classes Used

| Class | Purpose |
|-------|---------|
| `.badge-row` | Flex container for badges; applies `gap: var(--space-sm)` and `margin-top: var(--space-xl)` |
| `.badge` | Default badge styling — card background, border, mono font, `--font-small` size |
| `.badge--primary` | Gold border + gold text variant for the lead badge |

## Sizing Notes

- **Title**: Standard `# h1` markdown — the design system base styles handle font-size and color automatically.
- **Subtitle**: One line only. Keep under 12 words. Uses an inline `<p>` with token-based styles because this element appears only on title cards.
- **Badges**: 2-4 badges. The first badge should use `.badge--primary` (gold accent). Remaining badges use plain `.badge` (muted border).
- **Badge text**: Short labels — 1-3 words each, typically in monospace. Examples: version numbers, date ranges, audience labels.
- **Spacing**: The subtitle sits at `--space-lg` (32px) below the title. The badge row sits at `--space-xl` (48px) below the subtitle via the `.badge-row` class.

## Example

```html
---

# Cloud Infrastructure Proposal

<div style="margin-top: var(--space-lg);">
  <p style="font-size: var(--font-subtitle); color: var(--c-sub); font-weight: 400;">
    Migrating on-premise workloads to a managed Kubernetes platform
  </p>
</div>

<div class="badge-row">
  <span class="badge badge--primary">Q2 2026</span>
  <span class="badge">$240K budget</span>
  <span class="badge">Engineering</span>
</div>

<!--
"Good afternoon. Today I want to walk you through our proposal
for moving off bare metal and onto managed Kubernetes."

This is a board-level audience — lead with the business case,
not the technology. Keep it confident and concise.

Transition: "Let's start with where we are today."
-->
```
