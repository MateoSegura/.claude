# Styling Rules (Lipgloss)

## BTT-STY-001: Create Reusable Style Variables :yellow_circle:

**Tier**: Required

Define styles as package-level variables or in a dedicated theme package. Never inline complex styles.

```go
// CORRECT - Package-level styles
var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FFD700")).
        MarginBottom(1)

    selectedStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#000")).
        Background(lipgloss.Color("#FFD700")).
        Bold(true).
        Padding(0, 1)

    borderStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#6B6B6B"))
)

// INCORRECT - Inline styles (repetitive, hard to maintain)
func (m Model) View() string {
    return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700")).Render("Title")
}
```

---

## BTT-STY-002: Use lipgloss.Place for Centering :yellow_circle:

**Tier**: Required

Center content horizontally and/or vertically within a space.

```go
func (m Model) View() string {
    content := m.renderContent()

    // Center both horizontally and vertically
    return lipgloss.Place(
        m.width,            // total width
        m.height,           // total height
        lipgloss.Center,    // horizontal position
        lipgloss.Center,    // vertical position
        content,
    )
}

// Just horizontal centering
centered := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, "Title")

// With background color
placed := lipgloss.Place(
    m.width,
    m.height,
    lipgloss.Center,
    lipgloss.Center,
    content,
    lipgloss.WithWhitespaceChars(" "),
    lipgloss.WithWhitespaceBackground(lipgloss.Color("#0A0A0A")),
)
```

---

## BTT-STY-003: Use JoinHorizontal/JoinVertical for Layout :yellow_circle:

**Tier**: Required

Compose views by joining rendered strings.

```go
// Vertical layout (stacking)
view := lipgloss.JoinVertical(lipgloss.Left,
    header,
    content,
    footer,
)

// Horizontal layout (columns)
main := lipgloss.JoinHorizontal(lipgloss.Top,
    sidebar,    // Left column
    divider,    // Separator
    content,    // Right column
)

// Alignment options:
// lipgloss.Left, lipgloss.Center, lipgloss.Right (for JoinVertical)
// lipgloss.Top, lipgloss.Center, lipgloss.Bottom (for JoinHorizontal)
```

---

## BTT-STY-004: Set Width/Height for Consistent Sizing :yellow_circle:

**Tier**: Required

Constrain elements to specific dimensions for predictable layouts.

```go
// Fixed width box
box := lipgloss.NewStyle().
    Width(40).
    Border(lipgloss.RoundedBorder()).
    Render(content)

// Fixed height (adds padding to fill)
panel := lipgloss.NewStyle().
    Width(60).
    Height(20).
    Render(content)

// MaxWidth - truncates if exceeded
limited := lipgloss.NewStyle().
    MaxWidth(30).
    Render(longText)
```

---

## BTT-STY-005: Use Borders Consistently :yellow_circle:

**Tier**: Required

Lipgloss provides several border styles.

```go
// Border types
lipgloss.NormalBorder()      // ┌─┐│ │└─┘
lipgloss.RoundedBorder()     // ╭─╮│ │╰─╯
lipgloss.DoubleBorder()      // ╔═╗║ ║╚═╝
lipgloss.ThickBorder()       // ┏━┓┃ ┃┗━┛
lipgloss.HiddenBorder()      // No visible border (preserves spacing)

// Applying borders
style := lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#FFD700")).
    BorderTop(true).
    BorderBottom(true).
    BorderLeft(true).
    BorderRight(true)  // All sides (default)

// Selective borders
style := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder(), true, false, true, false)  // top, right, bottom, left
```

---

## BTT-STY-006: Use Padding and Margin Correctly :yellow_circle:

**Tier**: Required

- **Padding**: Space inside the border
- **Margin**: Space outside the border

```go
// Padding (inside)
style := lipgloss.NewStyle().
    Padding(1, 2)       // vertical, horizontal
    Padding(1, 2, 1, 2) // top, right, bottom, left
    PaddingTop(1).
    PaddingRight(2).
    PaddingBottom(1).
    PaddingLeft(2)

// Margin (outside)
style := lipgloss.NewStyle().
    Margin(1, 2)        // vertical, horizontal
    MarginTop(1).
    MarginBottom(1)
```

---

## BTT-STY-007: Color Specification :yellow_circle:

**Tier**: Required

Lipgloss supports multiple color formats.

```go
// Hex colors (recommended for consistency)
lipgloss.Color("#FFD700")
lipgloss.Color("#ff0")  // Short form

// ANSI 256 colors
lipgloss.Color("205")   // Hot pink
lipgloss.Color("39")    // Cyan

// ANSI 16 basic colors
lipgloss.Color("9")     // Bright red

// Adaptive colors (light/dark mode)
lipgloss.AdaptiveColor{
    Light: "#000000",
    Dark:  "#FFFFFF",
}

// Complete color (all formats)
lipgloss.CompleteColor{
    TrueColor: "#FFD700",
    ANSI256:   "220",
    ANSI:      "11",
}
```

---

## BTT-STY-008: Text Formatting :green_circle:

**Tier**: Recommended

```go
style := lipgloss.NewStyle().
    Bold(true).
    Italic(true).
    Underline(true).
    Strikethrough(true).
    Blink(true).         // Not widely supported
    Faint(true).         // Dimmed text
    Reverse(true)        // Swap fg/bg
```

---

## BTT-STY-009: Text Alignment :yellow_circle:

**Tier**: Required

```go
// Horizontal alignment (requires width)
style := lipgloss.NewStyle().
    Width(40).
    Align(lipgloss.Left)    // or Center, Right

// Vertical alignment (requires height)
style := lipgloss.NewStyle().
    Width(40).
    Height(10).
    Align(lipgloss.Center, lipgloss.Center)  // horizontal, vertical
```

---

## BTT-STY-010: Copy and Inherit Styles :green_circle:

**Tier**: Recommended

```go
// Base style
baseStyle := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FAFAFA"))

// Inherit and extend
boldStyle := baseStyle.Bold(true)
errorStyle := baseStyle.Foreground(lipgloss.Color("#FF0000"))

// Copy explicitly (styles are immutable, methods return new styles)
newStyle := baseStyle.Copy().Background(lipgloss.Color("#000"))
```

---

## BTT-STY-011: Use GetWidth/GetHeight for Calculations :green_circle:

**Tier**: Recommended

Get rendered dimensions for layout calculations.

```go
rendered := style.Render(content)
width := lipgloss.Width(rendered)
height := lipgloss.Height(rendered)

// Calculate remaining space
remaining := m.width - lipgloss.Width(sidebar) - lipgloss.Width(divider)
contentStyle := lipgloss.NewStyle().Width(remaining)
```

---

## Common Color Palettes

### Dark Theme (Premium)
```go
var (
    Gold       = lipgloss.Color("#FFD700")
    Black      = lipgloss.Color("#0A0A0A")
    Charcoal   = lipgloss.Color("#252525")
    White      = lipgloss.Color("#FAFAFA")
    Gray       = lipgloss.Color("#6B6B6B")
    Cyan       = lipgloss.Color("#00D9FF")
    Lime       = lipgloss.Color("#39FF14")
    Red        = lipgloss.Color("#FF073A")
    Orange     = lipgloss.Color("#FF6B00")
)
```

### ANSI 256 Safe Colors
```go
var (
    Red     = lipgloss.Color("196")
    Green   = lipgloss.Color("46")
    Yellow  = lipgloss.Color("226")
    Blue    = lipgloss.Color("21")
    Magenta = lipgloss.Color("201")
    Cyan    = lipgloss.Color("51")
    White   = lipgloss.Color("231")
    Gray    = lipgloss.Color("245")
)
```
