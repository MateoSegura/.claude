# Timeline
> When to use: Phase-based roadmaps, implementation plans, project milestones, or any sequential progression. Shows where you are now and what comes next. The active phase draws immediate focus.

## Design Tokens Used

| Token | Value | Role |
|-------|-------|------|
| `--c-gold` | `#D4A853` | Active dot fill, active phase title, future dot border |
| `--c-border` | `#3A3A3A` | Connecting line, future dot background |
| `--c-sub` | `#C8C4BD` | Phase date text |
| `--c-muted` | `#6B6B6B` | Phase description text |
| `--timeline-dot` | `2.25rem` | Dot diameter |
| `--space-lg` | `2rem` | Container top margin |
| `--space-xs` | `0.5rem` | Vertical spacing within phase |
| `--font-small` | `0.85rem` | Phase title text |
| `--font-caption` | `0.75rem` | Phase date text |
| `--font-micro` | `0.618rem` | Phase description text |

## Template

```html
---

## Slide Title

<div class="timeline-container">
  <div class="timeline-line"></div>
  <div class="timeline-phases">

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--active"></div>
      <div class="timeline-phase-title">Phase 1</div>
      <div class="timeline-phase-date">Month Year</div>
      <div class="timeline-phase-desc">1-line description</div>
    </div>

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--future"></div>
      <div class="timeline-phase-title">Phase 2</div>
      <div class="timeline-phase-date">Month Year</div>
      <div class="timeline-phase-desc">1-line description</div>
    </div>

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--future"></div>
      <div class="timeline-phase-title">Phase 3</div>
      <div class="timeline-phase-date">Month Year</div>
      <div class="timeline-phase-desc">1-line description</div>
    </div>

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--future"></div>
      <div class="timeline-phase-title">Phase 4</div>
      <div class="timeline-phase-date">Month Year</div>
      <div class="timeline-phase-desc">1-line description</div>
    </div>

  </div>
</div>

> "One-sentence framing of the overall timeline"

---
```

## Sizing Notes

- **Max phases**: 4. At 4 phases the dots, titles, dates, and descriptions all fit comfortably. Beyond 4, text overlaps and phases become unreadable.
- **Min phases**: 2. Fewer than 2 does not warrant a timeline -- use a simple text slide instead.
- **Active phase**: Exactly one phase uses `.timeline-dot--active` (solid gold fill). This represents "where we are now" or the current/starting phase. Always the first phase in a forward-looking roadmap.
- **Future phases**: All subsequent phases use `.timeline-dot--future` (grey fill with gold border outline). These represent upcoming milestones.
- **Completed phases**: If showing past phases, use `.timeline-dot--active` for the filled look. For a multi-completed scenario, the design system does not define a separate "completed" class -- use `.timeline-dot--active` for the current phase only and future for all others.
- **Phase title**: 1-3 words. Keep it a noun phrase: "Pilot Launch", "Full Rollout", "Optimization".
- **Phase date**: Month + Year or Quarter notation: "Mar 2026", "Q2 2026". Keep consistent across all phases.
- **Phase description**: Optional. One short line (under 8 words). Omit if the title and date are self-explanatory.
- **Connecting line**: The `.timeline-line` is a 2px horizontal rule positioned at the vertical center of the dots. It spans the full width automatically.
- **No inline styles**: All styling comes from CSS classes.
- **Vertical space**: The timeline container sits below the slide title with `--space-lg` margin. Leave room for a blockquote takeaway underneath if needed.

## Example

```html
---

## Implementation Roadmap

<div class="timeline-container">
  <div class="timeline-line"></div>
  <div class="timeline-phases">

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--active"></div>
      <div class="timeline-phase-title">Pilot</div>
      <div class="timeline-phase-date">Mar 2026</div>
      <div class="timeline-phase-desc">5-person team trial</div>
    </div>

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--future"></div>
      <div class="timeline-phase-title">Expand</div>
      <div class="timeline-phase-date">Jun 2026</div>
      <div class="timeline-phase-desc">Engineering-wide rollout</div>
    </div>

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--future"></div>
      <div class="timeline-phase-title">Optimize</div>
      <div class="timeline-phase-date">Sep 2026</div>
      <div class="timeline-phase-desc">Custom workflows + metrics</div>
    </div>

    <div class="timeline-phase">
      <div class="timeline-dot timeline-dot--future"></div>
      <div class="timeline-phase-title">Scale</div>
      <div class="timeline-phase-date">Dec 2026</div>
      <div class="timeline-phase-desc">Cross-department adoption</div>
    </div>

  </div>
</div>

> "Pilot to full scale in 9 months with quarterly checkpoints"

---
```
