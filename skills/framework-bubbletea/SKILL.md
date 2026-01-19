---
name: framework-bubbletea
description: Complete Bubble Tea TUI development standard for Go terminal applications
---

# Bubble Tea TUI Development Standard

> **Version**: 1.0.0 | **Status**: Active
> **Libraries**: bubbletea v1.3.10+, bubbles v0.21.0+, lipgloss v1.1.0+
> **Architecture**: Model-View-Update (MVU) / Elm Architecture

This standard establishes patterns for building terminal user interfaces with Charm's Bubble Tea ecosystem.

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [Architecture](rules/architecture.md) | BTT-ARC-* | MVU pattern, model composition |
| [Commands](rules/commands.md) | BTT-CMD-* | tea.Cmd, Batch, Sequence |
| [Components](rules/components.md) | BTT-CMP-* | Bubbles components integration |
| [Styling](rules/styling.md) | BTT-STY-* | Lipgloss styling patterns |
| [Testing](rules/testing.md) | BTT-TST-* | TUI testing strategies |

### Scaffolds (Copy-Paste Templates)

| Scaffold | File | Purpose |
|----------|------|---------|
| [Basic App](scaffolds/basic-app.go) | Minimal working app |
| [Component](scaffolds/component.go) | Reusable component |
| [Screen](scaffolds/screen.go) | Full-screen view |
| [Theme](scaffolds/theme.go) | Color/style system |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table |
| [Code Review](reference/code-review.md) | Review checklist |

---

## Rule Classification

| Tier | Marker | Response |
|------|--------|----------|
| **Critical** | :red_circle: | Must fix |
| **Required** | :yellow_circle: | Should fix |
| **Recommended** | :green_circle: | Consider |

---

## Core Architecture: Model-View-Update (MVU)

Every Bubble Tea application follows the MVU pattern:

```go
type Model interface {
    Init() tea.Cmd           // Initial command (called once)
    Update(tea.Msg) (Model, tea.Cmd)  // Handle messages, return new state + commands
    View() string            // Render UI as string
}
```

### BTT-ARC-001: Model Must Be Immutable :red_circle:

The Update function returns a NEW model, not a mutated one. This ensures predictable state transitions.

```go
// CORRECT - Return new model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }
    return m, nil  // Return m (or modified copy), never modify in place
}

// INCORRECT - Mutating state without returning
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.counter++  // DON'T: pointer receiver mutation
    return m, nil
}
```

### BTT-ARC-002: Use Value Receivers for Model Methods :red_circle:

Model methods use value receivers to maintain immutability semantics.

```go
// CORRECT
func (m Model) Init() tea.Cmd { return nil }
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }
func (m Model) View() string { ... }

// INCORRECT
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }
```

### BTT-ARC-003: Compose Models for Complex UIs :yellow_circle:

Break complex UIs into sub-models, each managing its own state.

```go
type AppModel struct {
    screen    Screen          // Current screen enum
    sidebar   SidebarModel    // Sub-model
    content   ContentModel    // Sub-model
    statusbar StatusbarModel  // Sub-model
    width     int
    height    int
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Route to sub-models
    var cmd tea.Cmd
    m.sidebar, cmd = m.sidebar.Update(msg)
    cmds = append(cmds, cmd)

    m.content, cmd = m.content.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}
```

---

## Commands (tea.Cmd)

### BTT-CMD-001: Commands Are Functions Returning Messages :red_circle:

A `tea.Cmd` is `func() tea.Msg`. Commands perform side effects and return messages.

```go
// CORRECT - Command returns a message
func loadDataCmd() tea.Msg {
    data, err := fetchData()
    if err != nil {
        return errMsg{err}
    }
    return dataLoadedMsg{data}
}

// Usage in Update:
return m, loadDataCmd  // Pass function, don't call it
```

### BTT-CMD-002: Use tea.Batch for Concurrent Commands :yellow_circle:

Run multiple independent commands concurrently.

```go
func (m Model) Init() tea.Cmd {
    return tea.Batch(
        loadConfigCmd,
        startSpinnerCmd,
        subscribeEventsCmd,
    )
}
```

### BTT-CMD-003: Use tea.Sequence for Ordered Commands :yellow_circle:

Run commands in sequence when order matters.

```go
// Save config, THEN quit
cmd := tea.Sequence(saveConfigCmd, tea.Quit)
```

### BTT-CMD-004: Never Block in Commands :red_circle:

Commands run in goroutines but should complete reasonably quickly.

```go
// CORRECT - HTTP with timeout
func fetchCmd() tea.Msg {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    resp, err := http.DefaultClient.Do(req.WithContext(ctx))
    // ...
}

// INCORRECT - Blocking forever
func badCmd() tea.Msg {
    <-make(chan struct{})  // Blocks forever!
    return nil
}
```

---

## Messages

### BTT-MSG-001: Define Custom Message Types :yellow_circle:

Use specific types for each message, not primitive types.

```go
// CORRECT - Specific message types
type dataLoadedMsg struct {
    items []Item
}

type errMsg struct {
    err error
}

type tickMsg time.Time

// INCORRECT - Using primitives
case string:  // Ambiguous!
case error:   // Could come from anywhere
```

### BTT-MSG-002: Handle tea.WindowSizeMsg for Responsive Layout :red_circle:

Always handle window resize to make your TUI responsive.

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.updateChildSizes()  // Propagate to sub-models
        return m, nil
    }
    // ...
}
```

### BTT-MSG-003: Handle tea.KeyMsg Correctly :red_circle:

Use type switch and check key strings or key types.

```go
case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
        return m, tea.Quit
    case "up", "k":
        m.cursor--
    case "down", "j":
        m.cursor++
    case "enter", " ":
        m.selected[m.cursor] = true
    }
```

---

## Program Options

### BTT-PRG-001: Use WithAltScreen for Full-Screen Apps :yellow_circle:

Alt screen prevents scrollback buffer pollution.

```go
p := tea.NewProgram(
    model,
    tea.WithAltScreen(),           // Full-screen mode
    tea.WithMouseCellMotion(),     // Mouse support
)
```

### BTT-PRG-002: Use Send for External Messages :green_circle:

Inject messages from outside the program (e.g., from goroutines).

```go
p := tea.NewProgram(model)

go func() {
    for event := range externalEvents {
        p.Send(externalEventMsg{event})
    }
}()

_, err := p.Run()
```

---

## Key Bindings

### BTT-KEY-001: Use bubbles/key for Structured Key Bindings :yellow_circle:

Define key bindings with help text for documentation.

```go
import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
    Up     key.Binding
    Down   key.Binding
    Select key.Binding
    Quit   key.Binding
}

var keys = KeyMap{
    Up: key.NewBinding(
        key.WithKeys("up", "k"),
        key.WithHelp("↑/k", "move up"),
    ),
    Down: key.NewBinding(
        key.WithKeys("down", "j"),
        key.WithHelp("↓/j", "move down"),
    ),
    Select: key.NewBinding(
        key.WithKeys("enter", " "),
        key.WithHelp("enter", "select"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("q", "ctrl+c"),
        key.WithHelp("q", "quit"),
    ),
}

// In Update:
case tea.KeyMsg:
    switch {
    case key.Matches(msg, keys.Up):
        m.cursor--
    case key.Matches(msg, keys.Down):
        m.cursor++
    }
```

---

## Bubbles Components

### BTT-CMP-001: Initialize Components in Model Constructor :yellow_circle:

Set up bubbles components when creating your model.

```go
func NewModel() Model {
    ti := textinput.New()
    ti.Placeholder = "Enter your name..."
    ti.Focus()
    ti.CharLimit = 50
    ti.Width = 30

    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    return Model{
        textInput: ti,
        spinner:   s,
    }
}
```

### BTT-CMP-002: Forward Messages to Child Components :red_circle:

Components need their messages to function (ticks, blinks, etc.).

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Forward to spinner (needs TickMsg)
    var spinnerCmd tea.Cmd
    m.spinner, spinnerCmd = m.spinner.Update(msg)
    cmds = append(cmds, spinnerCmd)

    // Forward to textinput (needs key events)
    var inputCmd tea.Cmd
    m.textInput, inputCmd = m.textInput.Update(msg)
    cmds = append(cmds, inputCmd)

    // Your update logic...

    return m, tea.Batch(cmds...)
}
```

### BTT-CMP-003: Return Component Init Commands :red_circle:

Animated components need their Init command to start.

```go
func (m Model) Init() tea.Cmd {
    return tea.Batch(
        m.spinner.Tick,         // Start spinner animation
        textinput.Blink,        // Start cursor blink
    )
}
```

---

## Styling with Lipgloss

### BTT-STY-001: Create Reusable Style Variables :yellow_circle:

Define styles as package-level variables or in a theme package.

```go
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

    mutedStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#6B6B6B"))
)
```

### BTT-STY-002: Use lipgloss.Place for Centering :yellow_circle:

Center content within the terminal.

```go
func (m Model) View() string {
    content := renderContent()

    return lipgloss.Place(
        m.width,
        m.height,
        lipgloss.Center,
        lipgloss.Center,
        content,
    )
}
```

### BTT-STY-003: Use JoinHorizontal/JoinVertical for Layout :yellow_circle:

Compose views by joining rendered strings.

```go
// Vertical layout
view := lipgloss.JoinVertical(lipgloss.Left,
    header,
    content,
    footer,
)

// Horizontal layout (e.g., sidebar + content)
main := lipgloss.JoinHorizontal(lipgloss.Top,
    sidebar,
    divider,
    content,
)
```

### BTT-STY-004: Set Width/Height for Consistent Sizing :yellow_circle:

Constrain elements to specific dimensions.

```go
box := lipgloss.NewStyle().
    Width(40).
    Height(10).
    Border(lipgloss.RoundedBorder()).
    Render(content)
```

---

## Quick Rule Lookup

| ID | Rule | Tier |
|----|------|------|
| BTT-ARC-001 | Model Must Be Immutable | Critical |
| BTT-ARC-002 | Use Value Receivers | Critical |
| BTT-ARC-003 | Compose Models for Complex UIs | Required |
| BTT-CMD-001 | Commands Return Messages | Critical |
| BTT-CMD-002 | Use tea.Batch for Concurrent | Required |
| BTT-CMD-003 | Use tea.Sequence for Ordered | Required |
| BTT-CMD-004 | Never Block in Commands | Critical |
| BTT-MSG-001 | Define Custom Message Types | Required |
| BTT-MSG-002 | Handle WindowSizeMsg | Critical |
| BTT-MSG-003 | Handle KeyMsg Correctly | Critical |
| BTT-PRG-001 | Use WithAltScreen | Required |
| BTT-PRG-002 | Use Send for External | Recommended |
| BTT-KEY-001 | Use bubbles/key | Required |
| BTT-CMP-001 | Initialize Components | Required |
| BTT-CMP-002 | Forward Messages to Children | Critical |
| BTT-CMP-003 | Return Component Init | Critical |
| BTT-STY-001 | Reusable Style Variables | Required |
| BTT-STY-002 | Use lipgloss.Place | Required |
| BTT-STY-003 | JoinHorizontal/JoinVertical | Required |
| BTT-STY-004 | Set Width/Height | Required |

---

## References

- [Bubble Tea README](https://github.com/charmbracelet/bubbletea)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)
- [Lip Gloss Styling](https://github.com/charmbracelet/lipgloss)
- [The Elm Architecture](https://guide.elm-lang.org/architecture/)
