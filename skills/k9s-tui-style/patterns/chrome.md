# Chrome Pattern

The Chrome pattern wraps every screen with consistent header and footer bars.

## K9S-CHR-001: Header Structure :red_circle:

```
┌─────────────────────────────────────────────────────────────────────┐
│ ◆ APPNAME                                          Context [count] │
└─────────────────────────────────────────────────────────────────────┘
  ↑ Left: Logo                                       ↑ Right: Breadcrumb
```

**Left Side (Title)**:
- App icon (◆ diamond)
- App name in CAPS
- Gold color, bold

**Right Side (Context)**:
- Current location/state
- Item count in brackets
- Gray/muted color

```go
func renderHeader(title, context string, width int) string {
    left := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFD700")). // Gold
        Bold(true).
        Render(title)

    right := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#3D3D3D")). // GrayDark
        Render(context)

    gap := width - lipgloss.Width(left) - lipgloss.Width(right) - 2
    if gap < 1 { gap = 1 }

    return lipgloss.NewStyle().
        Background(lipgloss.Color("#1A1A1A")). // BlackLight
        Width(width).
        Padding(0, 1).
        Render(left + strings.Repeat(" ", gap) + right)
}
```

## K9S-CHR-002: Footer Structure :red_circle:

```
┌─────────────────────────────────────────────────────────────────────┐
│ <↑↓>Navigate  <Enter>Select  <n>New  <q>Quit           <Ctrl+S>... │
└─────────────────────────────────────────────────────────────────────┘
  ↑ Left: Context shortcuts                          ↑ Right: Persistent
```

**Left Side (Context Shortcuts)**:
- Change based on current screen/mode
- Keys in gold with angle brackets
- Descriptions in gray

**Right Side (Persistent)**:
- Always-available shortcuts
- Keys in cyan (secondary accent)
- Optional, can be empty

```go
type Shortcut struct {
    Key  string
    Desc string
}

func renderFooter(shortcuts []Shortcut, persistent *Shortcut, width int) string {
    var parts []string
    for _, s := range shortcuts {
        key := lipgloss.NewStyle().
            Foreground(lipgloss.Color("#FFD700")).
            Bold(true).
            Render("<" + s.Key + ">")
        desc := lipgloss.NewStyle().
            Foreground(lipgloss.Color("#6B6B6B")).
            Render(s.Desc)
        parts = append(parts, key+desc)
    }
    left := strings.Join(parts, "  ")

    var right string
    if persistent != nil {
        key := lipgloss.NewStyle().
            Foreground(lipgloss.Color("#00D9FF")).
            Bold(true).
            Render("<" + persistent.Key + ">")
        desc := lipgloss.NewStyle().
            Foreground(lipgloss.Color("#3D3D3D")).
            Render(persistent.Desc)
        right = key + desc
    }

    gap := width - lipgloss.Width(left) - lipgloss.Width(right) - 2
    if gap < 1 { gap = 1 }

    return lipgloss.NewStyle().
        Background(lipgloss.Color("#1A1A1A")).
        Width(width).
        Padding(0, 1).
        Render(left + strings.Repeat(" ", gap) + right)
}
```

## K9S-CHR-003: Content Area :yellow_circle:

Content fills the space between header and footer.

```go
func (c Chrome) ContentHeight() int {
    height := c.height
    height -= 1 // header
    height -= 1 // footer
    return max(height, 1)
}

func (c Chrome) Render(content string) string {
    header := c.renderHeader()
    footer := c.renderFooter()

    contentStyle := lipgloss.NewStyle().
        Width(c.width).
        Height(c.ContentHeight())

    return lipgloss.JoinVertical(lipgloss.Left,
        header,
        contentStyle.Render(content),
        footer,
    )
}
```

## K9S-CHR-004: Centered Content Variant :yellow_circle:

For welcome screens, empty states, and modals, center the content.

```go
func (c Chrome) RenderCentered(content string) string {
    header := c.renderHeader()
    footer := c.renderFooter()

    centered := lipgloss.Place(
        c.width,
        c.ContentHeight(),
        lipgloss.Center,
        lipgloss.Center,
        content,
    )

    return lipgloss.JoinVertical(lipgloss.Left, header, centered, footer)
}
```

## K9S-CHR-005: Dynamic Shortcuts :yellow_circle:

Shortcuts change based on context. Define common sets:

```go
var (
    ShortcutsNavigation = []Shortcut{
        {Key: "↑↓", Desc: "Navigate"},
        {Key: "Enter", Desc: "Select"},
        {Key: "q", Desc: "Quit"},
    }

    ShortcutsForm = []Shortcut{
        {Key: "Tab", Desc: "Next"},
        {Key: "Enter", Desc: "Submit"},
        {Key: "Esc", Desc: "Cancel"},
    }

    ShortcutsList = []Shortcut{
        {Key: "↑↓", Desc: "Navigate"},
        {Key: "Enter", Desc: "Select"},
        {Key: "n", Desc: "New"},
        {Key: "d", Desc: "Delete"},
        {Key: "q", Desc: "Quit"},
    }

    ShortcutsEmpty = []Shortcut{
        {Key: "n", Desc: "New"},
        {Key: "q", Desc: "Quit"},
    }
)
```

## K9S-CHR-006: Too Small Handling :yellow_circle:

When terminal is below minimum size (80x24), show a helpful message.

```go
func (c Chrome) renderTooSmall() string {
    msg := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFD700")).
        Bold(true).
        Render("Terminal too small")

    hint := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#6B6B6B")).
        Render("Minimum: 80x24")

    content := lipgloss.JoinVertical(lipgloss.Center, msg, hint)

    return lipgloss.Place(
        c.width,
        c.height,
        lipgloss.Center,
        lipgloss.Center,
        content,
    )
}
```
