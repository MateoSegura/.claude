---
name: k9s-tui-style
description: K9s-inspired terminal UI design system for professional CLI tools
---

# K9s TUI Style Guide

> **Version**: 1.0.0 | **Status**: Active
> **Companion**: Use with `bubbletea-tui` skill for implementation
> **Reference**: [K9s - Kubernetes CLI](https://k9scli.io/)

This design system captures the K9s aesthetic: professional, keyboard-driven, information-dense terminal UIs that feel like premium developer tools.

---

## Design Philosophy

K9s represents a specific TUI aesthetic:

1. **Professional Dark Theme** - Rich blacks with gold/cyan accents
2. **Information Density** - Show data efficiently, no wasted space
3. **Keyboard-First** - Every action has a keyboard shortcut
4. **Chrome Pattern** - Consistent header/footer on every screen
5. **Status-Aware** - Color-coded states, context breadcrumbs
6. **Responsive** - Graceful handling of small terminals

---

## Navigation

### Patterns

| Pattern | File | Purpose |
|---------|------|---------|
| [Chrome](patterns/chrome.md) | K9S-CHR-* | Header/footer wrapper |
| [Lists](patterns/lists.md) | K9S-LST-* | Selectable item lists |
| [Forms](patterns/forms.md) | K9S-FRM-* | Input forms |
| [Modals](patterns/modals.md) | K9S-MOD-* | Dialogs and overlays |
| [Empty States](patterns/empty-states.md) | K9S-EMP-* | No data views |

### Scaffolds

| Scaffold | File | Purpose |
|----------|------|---------|
| [Chrome Component](scaffolds/chrome.go) | Header/footer wrapper |
| [List Screen](scaffolds/list-screen.go) | K9s-style list view |
| [Form Screen](scaffolds/form-screen.go) | K9s-style input form |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Design cheatsheet |

---

## Core Principles

### K9S-CORE-001: Chrome Pattern :red_circle:

Every screen has identical chrome (header + footer). Content changes, chrome stays consistent.

```
┌─────────────────────────────────────────────────────────────────────┐
│ ◆ APPNAME                                          Context [count] │  <- HEADER
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│                         CONTENT AREA                                │
│                    (varies by screen)                               │
│                                                                     │
├─────────────────────────────────────────────────────────────────────┤
│ <↑↓>Navigate  <Enter>Select  <n>New  <q>Quit           <Ctrl+S>... │  <- FOOTER
└─────────────────────────────────────────────────────────────────────┘
```

**Header Structure**:
- Left: App logo + name (gold, bold)
- Right: Context breadcrumb (gray)

**Footer Structure**:
- Left: Context-specific shortcuts (gold keys, gray descriptions)
- Right: Persistent shortcuts (cyan keys)

### K9S-CORE-002: Color Hierarchy :red_circle:

Colors have semantic meaning:

| Color | Hex | Usage |
|-------|-----|-------|
| **Gold** | `#FFD700` | Primary accent, selected items, interactive elements |
| **Black** | `#0A0A0A` | Main background |
| **BlackLight** | `#1A1A1A` | Header/footer background, elevated surfaces |
| **White** | `#FAFAFA` | Primary text |
| **Gray** | `#6B6B6B` | Secondary text, descriptions |
| **GrayDark** | `#3D3D3D` | Muted text, timestamps |
| **Cyan** | `#00D9FF` | Info, links, secondary actions |
| **Lime** | `#39FF14` | Success, running states |
| **Red** | `#FF073A` | Errors, failed states |
| **Orange** | `#FF6B00` | Warnings, pending states |

### K9S-CORE-003: Keyboard-First Design :red_circle:

Every action must be keyboard accessible. Common patterns:

| Key | Action | Universal |
|-----|--------|-----------|
| `j`/`↓` | Move down | Yes |
| `k`/`↑` | Move up | Yes |
| `Enter` | Select/confirm | Yes |
| `Esc` | Back/cancel | Yes |
| `q` | Quit | Yes |
| `?` | Help | Yes |
| `/` | Search/filter | When applicable |
| `Tab` | Next field/pane | When applicable |

### K9S-CORE-004: Shortcut Hint Format :yellow_circle:

Shortcuts displayed as `<key>description` with specific styling:

```go
// Key in gold, bold
key := lipgloss.NewStyle().
    Foreground(theme.Gold).
    Bold(true).
    Render("<" + "Enter" + ">")

// Description in gray
desc := lipgloss.NewStyle().
    Foreground(theme.Gray).
    Render("Select")

// Combined: <Enter>Select
hint := key + desc
```

### K9S-CORE-005: Selection Indicator :yellow_circle:

Selected items use the triangle indicator `▸` plus gold highlight:

```go
// Selected item
if selected {
    indicator := theme.IconTriangleR + " "  // "▸ "
    style := lipgloss.NewStyle().
        Foreground(theme.Black).
        Background(theme.Gold).
        Bold(true).
        Padding(0, 2)
    return style.Render(indicator + text)
}

// Unselected item
style := lipgloss.NewStyle().
    Foreground(theme.Gray).
    Padding(0, 2)
return style.Render("  " + text)  // Align with selected items
```

### K9S-CORE-006: Minimum Terminal Size :yellow_circle:

K9s-style apps require minimum 80x24. Handle smaller terminals gracefully:

```go
const (
    MinWidth  = 80
    MinHeight = 24
)

func (m Model) View() string {
    if m.width < MinWidth || m.height < MinHeight {
        return renderTooSmall(m.width, m.height)
    }
    return m.renderNormal()
}

func renderTooSmall(w, h int) string {
    msg := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true).
        Render("Terminal too small")

    hint := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render("Minimum: 80x24")

    return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center,
        lipgloss.JoinVertical(lipgloss.Center, msg, hint))
}
```

---

## Screen Types

### K9S-SCR-001: Welcome Screen :yellow_circle:

First-run experience with logo and clear CTA.

```
◆ APPNAME                                                        v1.0.0
───────────────────────────────────────────────────────────────────────

                    ██╗     ██████╗  ██████╗  ██████╗
                    ██║    ██╔═══██╗██╔════╝ ██╔═══██╗
                    ██║    ██║   ██║██║  ███╗██║   ██║
                    ██║    ██║   ██║██║   ██║██║   ██║
                    ███████╗╚██████╔╝╚██████╔╝╚██████╔╝
                    ╚══════╝ ╚═════╝  ╚═════╝  ╚═════╝

                     Your tagline goes here

                     ▸ Get Started
                       I know what I'm doing

───────────────────────────────────────────────────────────────────────
<↑↓>Navigate  <Enter>Select  <q>Quit
```

### K9S-SCR-002: List Screen :yellow_circle:

Selectable list with metadata columns.

```
◆ APPNAME                                                   Items [42]
───────────────────────────────────────────────────────────────────────
▸ ● item-name-1           /path/to/item              2h ago
  ○ item-name-2           /path/to/another           1d ago
  ✓ item-name-3           /path/to/third             3d ago
  ✗ item-name-4           /path/to/fourth            1w ago
───────────────────────────────────────────────────────────────────────
<↑↓>Navigate  <Enter>Select  <n>New  <d>Delete  <q>Quit
```

### K9S-SCR-003: Form Screen :yellow_circle:

Clean input form with labeled fields.

```
◆ APPNAME                                                    New Item
───────────────────────────────────────────────────────────────────────

                           Create New Item

                    Name         ┌────────────────────────┐
                                 │ my-item                │
                                 └────────────────────────┘

                    Path         ╭────────────────────────╮
                                 │ /path/to/workspace     │
                                 ╰────────────────────────╯

                    Workspace should be an existing directory

───────────────────────────────────────────────────────────────────────
<Tab>Next  <Enter>Submit  <Esc>Cancel
```

### K9S-SCR-004: Empty State :yellow_circle:

Clear message with actionable CTA when no data.

```
◆ APPNAME                                                      Items
───────────────────────────────────────────────────────────────────────


                           ✦ No Items Yet

                 Create your first item to get started.


                      ▸ Press N to create an item


───────────────────────────────────────────────────────────────────────
<n>New  <q>Quit
```

---

## Component Patterns

### Header Bar

```go
func renderHeader(title, context string, width int) string {
    left := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true).
        Render(title)  // "◆ APPNAME"

    right := lipgloss.NewStyle().
        Foreground(theme.GrayDark).
        Render(context)  // "Items [42]"

    gap := width - lipgloss.Width(left) - lipgloss.Width(right) - 2

    return lipgloss.NewStyle().
        Background(theme.BlackLight).
        Width(width).
        Padding(0, 1).
        Render(left + strings.Repeat(" ", gap) + right)
}
```

### Footer Bar

```go
type Shortcut struct {
    Key  string
    Desc string
}

func renderFooter(shortcuts []Shortcut, width int) string {
    var parts []string
    for _, s := range shortcuts {
        key := lipgloss.NewStyle().
            Foreground(theme.Gold).
            Bold(true).
            Render("<" + s.Key + ">")
        desc := lipgloss.NewStyle().
            Foreground(theme.Gray).
            Render(s.Desc)
        parts = append(parts, key+desc)
    }

    return lipgloss.NewStyle().
        Background(theme.BlackLight).
        Width(width).
        Padding(0, 1).
        Render(strings.Join(parts, "  "))
}
```

### List Item

```go
func renderListItem(name, path, time string, state string, selected bool) string {
    // State icon with color
    stateIcon := theme.GetStateIcon(state)
    stateColor := theme.GetStateColor(state)
    icon := lipgloss.NewStyle().Foreground(stateColor).Render(stateIcon)

    // Selection
    var selector string
    var nameStyle lipgloss.Style
    if selected {
        selector = lipgloss.NewStyle().Foreground(theme.Gold).Render("▸ ")
        nameStyle = lipgloss.NewStyle().Foreground(theme.Gold).Bold(true)
    } else {
        selector = "  "
        nameStyle = lipgloss.NewStyle().Foreground(theme.White)
    }

    // Compose line
    namePart := nameStyle.Width(20).Render(name)
    pathPart := lipgloss.NewStyle().Foreground(theme.Gray).Width(30).Render(path)
    timePart := lipgloss.NewStyle().Foreground(theme.GrayDark).Render(time)

    return selector + icon + " " + namePart + pathPart + timePart
}
```

---

## Icons Reference

| Icon | Const | Usage |
|------|-------|-------|
| ◆ | `IconDiamond` | Logo, premium |
| ▸ | `IconTriangleR` | Selected, play |
| ● | `IconCircleFill` | Active, running |
| ○ | `IconCircleEmpty` | Pending, inactive |
| ✓ | `IconCheck` | Success, complete |
| ✗ | `IconCross` | Error, failed |
| ✦ | `IconSparkle` | New, special |
| ⚙ | `IconSettings` | Settings |
| • | `IconBullet` | Separator |

---

## State Colors

| State | Color | Icon |
|-------|-------|------|
| `new` | Cyan | ✦ |
| `pending` | Gray | ○ |
| `running` | Lime | ● |
| `success` | Gold | ✓ |
| `failed` | Red | ✗ |
| `warning` | Orange | ⚠ |

---

## Integration with bubbletea-tui

This skill defines the **design system**. Use `bubbletea-tui` for **implementation patterns**.

```go
// Use K9s-style chrome (this skill)
chrome := NewChrome(width, height).
    SetTitle("◆ MYAPP").
    SetContext("Items [42]").
    SetShortcuts(shortcuts)

// Use bubbletea patterns (bubbletea-tui skill)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // BTT-ARC-001: Return new model, don't mutate
    // BTT-MSG-002: Handle WindowSizeMsg
    // BTT-CMP-002: Forward to children
}
```
