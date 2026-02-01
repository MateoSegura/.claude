# Scenario Cards
> When to use: Presenting ROI projections, pricing tiers, risk scenarios, or any grouped comparison where 2-3 options differ by magnitude. The primary card draws the eye to the recommended or most likely scenario.

## Design Tokens Used

| Token | Value | Role |
|-------|-------|------|
| `--c-gold` | `#D4A853` | Primary card border, primary label, primary value |
| `--c-card` | `#2A2A2A` | Card background |
| `--c-border` | `#3A3A3A` | Non-primary card border |
| `--c-sub` | `#C8C4BD` | Non-primary values, detail text |
| `--c-muted` | `#6B6B6B` | Non-primary labels, sub-detail text |
| `--font-metric-lg` | `3.25rem` | Hero number in each card |
| `--font-caption` | `0.75rem` | Uppercase scenario label |
| `--font-small` | `0.85rem` | Detail and sub-detail text |
| `--space-md` | `1.5rem` | Gap between cards |
| `--space-sm` | `1rem` | Card internal padding, row top margin |
| `--space-xs` | `0.5rem` | Vertical spacing within card |
| `--card-radius` | `0.5rem` | Card corner radius |

## Template

```html
---

## Slide Title

<div class="scenario-row">
  <div class="scenario-card">
    <div class="scenario-label">Scenario A</div>
    <div class="scenario-value">$XXX</div>
    <div class="scenario-detail">Primary metric description</div>
    <div class="scenario-sub">Secondary detail</div>
  </div>
  <div class="scenario-card scenario-card--primary">
    <div class="scenario-label">Scenario B</div>
    <div class="scenario-value">$XXX</div>
    <div class="scenario-detail">Primary metric description</div>
    <div class="scenario-sub">Secondary detail</div>
  </div>
  <div class="scenario-card">
    <div class="scenario-label">Scenario C</div>
    <div class="scenario-value">$XXX</div>
    <div class="scenario-detail">Primary metric description</div>
    <div class="scenario-sub">Secondary detail</div>
  </div>
</div>

> "One-sentence takeaway framing the recommended scenario"

---
```

## Sizing Notes

- **Card count**: 3 cards is the standard layout. 2 cards work for binary comparisons (use `flex: 1` sizing, cards stretch evenly). Never exceed 3 cards -- split into multiple slides instead.
- **Primary card**: Exactly one card gets `.scenario-card--primary` for the gold border. This is the recommended, expected, or focal scenario. Place it in the center for visual balance.
- **Hero number**: The `.scenario-value` dominates each card. Keep it short: `$240K`, `3.2x`, `18mo`. Avoid long strings that break the mono-width alignment.
- **Detail line**: One line max for `.scenario-detail`. Describes what the hero number represents (e.g., "Annual savings", "Return on investment").
- **Sub-detail**: Optional. Use `.scenario-sub` for a secondary metric or assumption (e.g., "Based on 12-month adoption"). Omit if not needed.
- **Label**: Uppercase by default via CSS. Keep to 1-2 words: "Conservative", "Moderate", "Aggressive" or "Year 1", "Year 2", "Year 3".
- **Blockquote**: Strongly recommended. Frames the decision for the audience after they scan the three numbers.

## Example

```html
---

## Projected ROI — Year 1

<div class="scenario-row">
  <div class="scenario-card">
    <div class="scenario-label">Conservative</div>
    <div class="scenario-value">1.8x</div>
    <div class="scenario-detail">$180K annual savings</div>
    <div class="scenario-sub">30% adoption rate</div>
  </div>
  <div class="scenario-card scenario-card--primary">
    <div class="scenario-label">Moderate</div>
    <div class="scenario-value">3.2x</div>
    <div class="scenario-detail">$320K annual savings</div>
    <div class="scenario-sub">60% adoption rate</div>
  </div>
  <div class="scenario-card">
    <div class="scenario-label">Aggressive</div>
    <div class="scenario-value">5.1x</div>
    <div class="scenario-detail">$510K annual savings</div>
    <div class="scenario-sub">90% adoption rate</div>
  </div>
</div>

> "Even the conservative estimate pays back the investment in under 7 months"

---
```
