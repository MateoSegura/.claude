# List Pattern

K9s-style lists are information-dense, keyboard-navigable, and status-aware.

## K9S-LST-001: List Item Structure :red_circle:

Each item has a consistent structure:

```
▸ ● item-name              /path/to/item                    2h ago
↑ ↑ ↑                      ↑                                ↑
│ │ └── Name (bold when selected)                           └── Timestamp
│ └──── State icon (colored)                    Path (truncated, gray)
└────── Selector (▸ when selected, space otherwise)
```

```go
func renderListItem(item Item, selected bool, width int) string {
    // Selector
    var selector string
    if selected {
        selector = lipgloss.NewStyle().
            Foreground(theme.Gold).
            Render("▸ ")
    } else {
        selector = "  "
    }

    // State icon
    icon := lipgloss.NewStyle().
        Foreground(theme.GetStateColor(item.State)).
        Render(theme.GetStateIcon(item.State))

    // Name
    var nameStyle lipgloss.Style
    if selected {
        nameStyle = lipgloss.NewStyle().
            Foreground(theme.Gold).
            Bold(true)
    } else {
        nameStyle = lipgloss.NewStyle().
            Foreground(theme.White)
    }
    name := nameStyle.Width(20).Render(item.Name)

    // Path (truncated from left if too long)
    path := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Width(35).
        Render(truncatePath(item.Path, 35))

    // Time
    time := lipgloss.NewStyle().
        Foreground(theme.GrayDark).
        Render(formatRelativeTime(item.UpdatedAt))

    return selector + icon + " " + name + path + time
}
```

## K9S-LST-002: Two-Line Item Variant :yellow_circle:

For items with more metadata, use two lines:

```
▸ ● item-name
    /full/path/to/item  •  2h ago  •  running
```

```go
func renderListItemTwoLine(item Item, selected bool) string {
    // Line 1: Selector + Icon + Name
    var line1Parts []string

    if selected {
        line1Parts = append(line1Parts,
            lipgloss.NewStyle().Foreground(theme.Gold).Render("▸ "))
    } else {
        line1Parts = append(line1Parts, "  ")
    }

    line1Parts = append(line1Parts,
        lipgloss.NewStyle().
            Foreground(theme.GetStateColor(item.State)).
            Render(theme.GetStateIcon(item.State)+" "))

    nameStyle := lipgloss.NewStyle().Foreground(theme.White)
    if selected {
        nameStyle = nameStyle.Foreground(theme.Gold).Bold(true)
    }
    line1Parts = append(line1Parts, nameStyle.Render(item.Name))

    line1 := strings.Join(line1Parts, "")

    // Line 2: Indented metadata
    meta := []string{
        lipgloss.NewStyle().Foreground(theme.Gray).Render(item.Path),
        lipgloss.NewStyle().Foreground(theme.GrayDark).Render(formatRelativeTime(item.UpdatedAt)),
        lipgloss.NewStyle().Foreground(theme.GetStateColor(item.State)).Render(item.State),
    }
    line2 := "    " + strings.Join(meta, "  "+theme.IconBullet+"  ")

    return lipgloss.JoinVertical(lipgloss.Left, line1, line2, "")
}
```

## K9S-LST-003: Selection Highlighting :red_circle:

The selected item must be visually distinct:

1. **Triangle indicator** (▸) on the left
2. **Gold color** for the name
3. **Bold text** for emphasis

```go
// Selected styling
if selected {
    selector = theme.IconTriangleR + " "
    nameStyle = lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true)
} else {
    selector = "  "  // 2 spaces to align with "▸ "
    nameStyle = lipgloss.NewStyle().
        Foreground(theme.White)
}
```

## K9S-LST-004: State Icons and Colors :yellow_circle:

Use semantic icons and colors for item states:

```go
var StateIcons = map[string]string{
    "new":       "✦",  // Sparkle
    "pending":   "○",  // Empty circle
    "running":   "●",  // Filled circle
    "success":   "✓",  // Check
    "completed": "✓",
    "failed":    "✗",  // Cross
    "error":     "✗",
    "warning":   "⚠",
}

var StateColors = map[string]lipgloss.Color{
    "new":       lipgloss.Color("#00D9FF"), // Cyan
    "pending":   lipgloss.Color("#6B6B6B"), // Gray
    "running":   lipgloss.Color("#39FF14"), // Lime
    "success":   lipgloss.Color("#FFD700"), // Gold
    "completed": lipgloss.Color("#FFD700"),
    "failed":    lipgloss.Color("#FF073A"), // Red
    "error":     lipgloss.Color("#FF073A"),
    "warning":   lipgloss.Color("#FF6B00"), // Orange
}
```

## K9S-LST-005: List Navigation :red_circle:

Standard vim-style navigation:

```go
case tea.KeyMsg:
    switch {
    case key.Matches(msg, keys.Up):
        if m.cursor > 0 {
            m.cursor--
        }
    case key.Matches(msg, keys.Down):
        if m.cursor < len(m.items)-1 {
            m.cursor++
        }
    case key.Matches(msg, keys.Select):
        if len(m.items) > 0 {
            return m, func() tea.Msg {
                return ItemSelectedMsg{Item: m.items[m.cursor]}
            }
        }
    }
```

## K9S-LST-006: Scrolling for Long Lists :yellow_circle:

When list exceeds visible area, implement viewport scrolling:

```go
func (m Model) visibleItems() []Item {
    visibleHeight := m.chrome.ContentHeight() - 2 // title + padding
    itemHeight := 3 // lines per item (for two-line variant)
    maxVisible := visibleHeight / itemHeight

    start := 0
    if m.cursor >= maxVisible {
        start = m.cursor - maxVisible + 1
    }
    end := min(start+maxVisible, len(m.items))

    return m.items[start:end]
}
```

## K9S-LST-007: Column Alignment :yellow_circle:

Use fixed widths for columns to maintain alignment:

```go
const (
    ColName  = 20
    ColPath  = 35
    ColTime  = 10
    ColState = 10
)

func renderRow(name, path, time, state string) string {
    return lipgloss.JoinHorizontal(lipgloss.Left,
        lipgloss.NewStyle().Width(ColName).Render(name),
        lipgloss.NewStyle().Width(ColPath).Foreground(theme.Gray).Render(path),
        lipgloss.NewStyle().Width(ColTime).Foreground(theme.GrayDark).Render(time),
        lipgloss.NewStyle().Width(ColState).Render(state),
    )
}
```

## K9S-LST-008: Path Truncation :yellow_circle:

Truncate long paths from the left with ellipsis:

```go
func truncatePath(path string, maxLen int) string {
    if len(path) <= maxLen {
        return path
    }
    return "…" + path[len(path)-maxLen+1:]
}

// /very/long/path/to/some/directory
// becomes
// …th/to/some/directory
```

## K9S-LST-009: Relative Time :yellow_circle:

Display timestamps as relative time:

```go
func formatRelativeTime(t time.Time) string {
    if t.IsZero() {
        return "never"
    }
    dur := time.Since(t)
    switch {
    case dur < time.Minute:
        return "just now"
    case dur < time.Hour:
        return fmt.Sprintf("%dm ago", int(dur.Minutes()))
    case dur < 24*time.Hour:
        return fmt.Sprintf("%dh ago", int(dur.Hours()))
    case dur < 7*24*time.Hour:
        return fmt.Sprintf("%dd ago", int(dur.Hours()/24))
    default:
        return t.Format("Jan 2")
    }
}
```
