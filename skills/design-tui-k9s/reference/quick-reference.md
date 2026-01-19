# K9s TUI Style - Quick Reference

## Color Palette

| Name | Hex | Usage |
|------|-----|-------|
| Gold | `#FFD700` | Primary accent, selected, interactive |
| Black | `#0A0A0A` | Main background |
| BlackLight | `#1A1A1A` | Header/footer background |
| White | `#FAFAFA` | Primary text |
| Gray | `#6B6B6B` | Secondary text |
| GrayDark | `#3D3D3D` | Muted text, timestamps |
| Charcoal | `#252525` | Borders |
| Cyan | `#00D9FF` | Info, secondary actions |
| Lime | `#39FF14` | Success, running |
| Red | `#FF073A` | Errors, failed |
| Orange | `#FF6B00` | Warnings |

## Icons

| Icon | Const | Usage |
|------|-------|-------|
| ◆ | `iconDiamond` | Logo, brand |
| ▸ | `iconTriangleR` | Selected item |
| ● | `iconCircle` | Running, active |
| ○ | `iconCircleO` | Pending |
| ✓ | `iconCheck` | Success |
| ✗ | `iconCross` | Error, failed |
| ✦ | `iconSparkle` | New, special |
| • | `iconBullet` | Separator |
| ⚠ | `iconWarning` | Warning |

## Screen Layout

```
┌─────────────────────────────────────────────────────────────────────┐
│ ◆ APPNAME                                          Context [count] │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│                           CONTENT                                   │
│                                                                     │
├─────────────────────────────────────────────────────────────────────┤
│ <key>action  <key>action  <key>action                   <key>...   │
└─────────────────────────────────────────────────────────────────────┘
```

## Key Bindings

| Key | Action | Context |
|-----|--------|---------|
| `j`/`↓` | Down | Lists |
| `k`/`↑` | Up | Lists |
| `Enter` | Select/Submit | All |
| `Esc` | Back/Cancel | All |
| `Tab` | Next field | Forms |
| `q` | Quit | All |
| `n` | New | Lists |
| `d` | Delete | Lists |
| `/` | Search | Lists (optional) |
| `?` | Help | All (optional) |

## Shortcut Format

```go
// Key in gold, bold with angle brackets
key := lipgloss.NewStyle().
    Foreground(colorGold).
    Bold(true).
    Render("<Enter>")

// Description in gray
desc := lipgloss.NewStyle().
    Foreground(colorGray).
    Render("Select")

// Result: <Enter>Select
```

## Selection Indicator

```go
// Selected
"▸ " + goldBoldText

// Unselected
"  " + normalText
```

## Minimum Size

```go
const (
    MinWidth  = 80
    MinHeight = 24
)
```

## State Mapping

| State | Icon | Color |
|-------|------|-------|
| new | ✦ | Cyan |
| pending | ○ | Gray |
| running | ● | Lime |
| success | ✓ | Gold |
| failed | ✗ | Red |
| warning | ⚠ | Orange |

## Header Pattern

```go
left := goldBold(iconDiamond + " APPNAME")
right := grayDark("Context [42]")
gap := width - len(left) - len(right)
header := blackLightBg(left + spaces(gap) + right)
```

## Footer Pattern

```go
shortcuts := []string{
    goldBold("<↑↓>") + gray("Navigate"),
    goldBold("<Enter>") + gray("Select"),
    goldBold("<q>") + gray("Quit"),
}
footer := blackLightBg(join(shortcuts, "  "))
```

## List Item Pattern

```go
// Line format: selector icon name path time
"▸ ● item-name              /path/to/item          2h ago"
"  ○ another-item           /path/to/another       1d ago"
```

## Form Field Pattern

```go
Label:         ╭────────────────────────╮
               │ value here             │
               ╰────────────────────────╯
```

## Empty State Pattern

```go
✦ No Items Yet
Create your first item to get started.

▸ Press N to create an item
```

## Modal Pattern

```go
╭───────────────────────────────╮
│                               │
│       Modal Title             │
│                               │
│       Content here            │
│                               │
│   <y>Yes  <n>No  <Esc>Cancel  │
│                               │
╰───────────────────────────────╯
```

## Checklist

- [ ] Header with logo + context
- [ ] Footer with context shortcuts
- [ ] Minimum size handling (80x24)
- [ ] Selected item has ▸ and gold
- [ ] States have semantic colors
- [ ] Vim navigation (j/k or arrows)
- [ ] Esc goes back
- [ ] q quits
- [ ] Empty state with CTA
- [ ] Form has Tab navigation
