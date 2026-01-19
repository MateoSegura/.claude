# Bubble Tea Quick Reference

## Rule Summary

| ID | Rule | Tier |
|----|------|------|
| BTT-ARC-001 | Model Must Be Immutable | :red_circle: Critical |
| BTT-ARC-002 | Use Value Receivers | :red_circle: Critical |
| BTT-ARC-003 | Compose Models for Complex UIs | :yellow_circle: Required |
| BTT-ARC-004 | Use Enums for Screen/State | :yellow_circle: Required |
| BTT-ARC-005 | Separate View from Update | :green_circle: Recommended |
| BTT-CMD-001 | Commands Return Messages | :red_circle: Critical |
| BTT-CMD-002 | Use tea.Batch for Concurrent | :yellow_circle: Required |
| BTT-CMD-003 | Use tea.Sequence for Ordered | :yellow_circle: Required |
| BTT-CMD-004 | Never Block in Commands | :red_circle: Critical |
| BTT-CMD-005 | Use tea.Tick for Timers | :yellow_circle: Required |
| BTT-CMD-006 | tea.Quit and tea.ClearScreen | :yellow_circle: Required |
| BTT-CMD-007 | Use tea.Printf for Debug | :green_circle: Recommended |
| BTT-MSG-001 | Define Custom Message Types | :yellow_circle: Required |
| BTT-MSG-002 | Handle WindowSizeMsg | :red_circle: Critical |
| BTT-MSG-003 | Handle KeyMsg Correctly | :red_circle: Critical |
| BTT-PRG-001 | Use WithAltScreen | :yellow_circle: Required |
| BTT-PRG-002 | Use Send for External | :green_circle: Recommended |
| BTT-KEY-001 | Use bubbles/key | :yellow_circle: Required |
| BTT-CMP-001 | Initialize Components | :yellow_circle: Required |
| BTT-CMP-002 | Forward Messages to Children | :red_circle: Critical |
| BTT-CMP-003 | Return Component Init | :red_circle: Critical |
| BTT-CMP-004 | Conditionally Forward by Focus | :yellow_circle: Required |
| BTT-CMP-005 | Resize on WindowSizeMsg | :yellow_circle: Required |
| BTT-STY-001 | Reusable Style Variables | :yellow_circle: Required |
| BTT-STY-002 | Use lipgloss.Place | :yellow_circle: Required |
| BTT-STY-003 | JoinHorizontal/JoinVertical | :yellow_circle: Required |
| BTT-STY-004 | Set Width/Height | :yellow_circle: Required |
| BTT-STY-005 | Use Borders Consistently | :yellow_circle: Required |
| BTT-STY-006 | Padding and Margin | :yellow_circle: Required |
| BTT-STY-007 | Color Specification | :yellow_circle: Required |
| BTT-TST-001 | Test Update Logic | :yellow_circle: Required |
| BTT-TST-002 | Test View Output | :yellow_circle: Required |
| BTT-TST-003 | Test Command Messages | :yellow_circle: Required |
| BTT-TST-004 | Test State Transitions | :yellow_circle: Required |

---

## Essential Patterns

### Minimal Bubble Tea App
```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type model struct {
    width, height int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width, m.height = msg.Width, msg.Height
    case tea.KeyMsg:
        if msg.String() == "q" { return m, tea.Quit }
    }
    return m, nil
}

func (m model) View() string {
    return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, "Hello!")
}

func main() {
    tea.NewProgram(model{}, tea.WithAltScreen()).Run()
}
```

### Message Pattern
```go
// Define message types
type dataLoadedMsg struct { data []Item }
type errMsg struct { err error }

// Command that returns message
func loadCmd() tea.Cmd {
    return func() tea.Msg {
        data, err := fetch()
        if err != nil { return errMsg{err} }
        return dataLoadedMsg{data}
    }
}

// Handle in Update
case dataLoadedMsg:
    m.items = msg.data
    return m, nil
case errMsg:
    m.err = msg.err
    return m, nil
```

### Component Integration
```go
// Init
func (m model) Init() tea.Cmd {
    return tea.Batch(m.spinner.Tick, textinput.Blink)
}

// Update - forward to children
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
    var cmd tea.Cmd
    m.spinner, cmd = m.spinner.Update(msg)
    cmds = append(cmds, cmd)
    m.input, cmd = m.input.Update(msg)
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}
```

### Key Bindings
```go
import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
    Up   key.Binding
    Down key.Binding
    Quit key.Binding
}

var keys = keyMap{
    Up:   key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
    Down: key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
    Quit: key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

// In Update:
case tea.KeyMsg:
    switch {
    case key.Matches(msg, keys.Up):   m.cursor--
    case key.Matches(msg, keys.Down): m.cursor++
    case key.Matches(msg, keys.Quit): return m, tea.Quit
    }
```

---

## Lipgloss Cheatsheet

### Style Basics
```go
style := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FFD700")).
    Background(lipgloss.Color("#000")).
    Padding(1, 2).
    Margin(1, 2).
    Width(40).
    Height(10).
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#6B6B6B"))

output := style.Render("Hello")
```

### Layout
```go
// Vertical stack
lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

// Horizontal columns
lipgloss.JoinHorizontal(lipgloss.Top, left, middle, right)

// Centering
lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
```

### Colors
```go
lipgloss.Color("#FFD700")     // Hex
lipgloss.Color("205")         // ANSI 256
lipgloss.AdaptiveColor{Light: "#000", Dark: "#FFF"}
```

---

## Bubbles Components

| Component | Import | Init Cmd | Notes |
|-----------|--------|----------|-------|
| spinner | `bubbles/spinner` | `spinner.Tick` | Animated loading |
| textinput | `bubbles/textinput` | `textinput.Blink` | Single line input |
| textarea | `bubbles/textarea` | `textarea.Blink` | Multi-line input |
| list | `bubbles/list` | None | Filterable list |
| viewport | `bubbles/viewport` | None | Scrollable content |
| table | `bubbles/table` | None | Data table |
| progress | `bubbles/progress` | None | Progress bar |
| paginator | `bubbles/paginator` | None | Page navigation |

---

## Common Issues

| Issue | Cause | Fix |
|-------|-------|-----|
| Spinner not animating | Missing Init command | Return `spinner.Tick` from Init |
| Input not responding | Not forwarding messages | Forward all messages to textinput |
| Screen corrupted | Using fmt.Println | Use tea.Printf instead |
| View not updating | Returning same model | Ensure you modify and return model |
| Commands not running | Calling instead of passing | Pass `cmd` not `cmd()` |
| No window size | Missing handler | Handle `tea.WindowSizeMsg` |
