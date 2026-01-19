# Components Rules

## BTT-CMP-001: Initialize Components in Model Constructor :yellow_circle:

**Tier**: Required

Set up bubbles components when creating your model, not in `Init()`.

```go
import (
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbles/spinner"
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/viewport"
)

func NewModel() Model {
    // Text input
    ti := textinput.New()
    ti.Placeholder = "Enter your name..."
    ti.Focus()
    ti.CharLimit = 50
    ti.Width = 30

    // Spinner
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    // List
    items := []list.Item{
        item{title: "Item 1", desc: "Description 1"},
    }
    l := list.New(items, list.NewDefaultDelegate(), 0, 0)
    l.Title = "My List"

    // Viewport (for scrollable content)
    vp := viewport.New(80, 20)
    vp.SetContent("Scrollable content here...")

    return Model{
        textInput: ti,
        spinner:   s,
        list:      l,
        viewport:  vp,
    }
}
```

---

## BTT-CMP-002: Forward Messages to Child Components :red_circle:

**Tier**: Critical

Components need to receive messages to function. Spinners need tick messages, text inputs need key events, etc.

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Always forward to components that need updates
    var cmd tea.Cmd

    // Spinner needs TickMsg
    m.spinner, cmd = m.spinner.Update(msg)
    cmds = append(cmds, cmd)

    // Text input needs KeyMsg
    m.textInput, cmd = m.textInput.Update(msg)
    cmds = append(cmds, cmd)

    // List needs KeyMsg and WindowSizeMsg
    m.list, cmd = m.list.Update(msg)
    cmds = append(cmds, cmd)

    // Viewport needs KeyMsg for scrolling
    m.viewport, cmd = m.viewport.Update(msg)
    cmds = append(cmds, cmd)

    // Your custom logic...
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }

    return m, tea.Batch(cmds...)
}
```

---

## BTT-CMP-003: Return Component Init Commands :red_circle:

**Tier**: Critical

Animated components (spinner, textinput cursor) need their tick/blink commands started.

```go
func (m Model) Init() tea.Cmd {
    return tea.Batch(
        m.spinner.Tick,       // Start spinner animation
        textinput.Blink,      // Start cursor blink
        m.loadDataCmd(),      // Your init command
    )
}
```

**Component Init commands**:
- `spinner.Tick` - Start spinner
- `textinput.Blink` - Start cursor blink
- `textarea.Blink` - Start cursor blink
- `list.StartSpinner()` - If list has spinner

---

## BTT-CMP-004: Conditionally Forward Based on Focus :yellow_circle:

**Tier**: Required

Only forward key events to the focused component to avoid conflicts.

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Tab switches focus
        if msg.String() == "tab" {
            m.focus = (m.focus + 1) % 2
            return m, nil
        }

        // Forward keys only to focused component
        var cmd tea.Cmd
        switch m.focus {
        case FocusList:
            m.list, cmd = m.list.Update(msg)
        case FocusInput:
            m.textInput, cmd = m.textInput.Update(msg)
        }
        cmds = append(cmds, cmd)

    default:
        // Non-key messages go to all components
        var cmd tea.Cmd
        m.spinner, cmd = m.spinner.Update(msg)
        cmds = append(cmds, cmd)
    }

    return m, tea.Batch(cmds...)
}
```

---

## BTT-CMP-005: Resize Components on WindowSizeMsg :yellow_circle:

**Tier**: Required

Update component dimensions when the window resizes.

```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height

    // Update list dimensions
    m.list.SetWidth(msg.Width)
    m.list.SetHeight(msg.Height - 4)  // Leave room for header/footer

    // Update viewport dimensions
    m.viewport.Width = msg.Width - 2
    m.viewport.Height = msg.Height - 6

    // Update text input width
    m.textInput.Width = msg.Width - 10

    return m, nil
```

---

## Common Bubbles Components Reference

### Spinner
```go
import "github.com/charmbracelet/bubbles/spinner"

s := spinner.New()
s.Spinner = spinner.Dot  // or Line, MiniDot, Jump, Pulse, Points, Globe, Moon, Monkey
s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

// Init: return s.Tick
// Update: s, cmd = s.Update(msg)
// View: s.View()
```

### TextInput
```go
import "github.com/charmbracelet/bubbles/textinput"

ti := textinput.New()
ti.Placeholder = "Type here..."
ti.Focus()
ti.CharLimit = 100
ti.Width = 40
ti.EchoMode = textinput.EchoPassword  // For passwords
ti.SetValue("initial value")

// Value: ti.Value()
// Init: return textinput.Blink
```

### TextArea
```go
import "github.com/charmbracelet/bubbles/textarea"

ta := textarea.New()
ta.Placeholder = "Enter description..."
ta.SetWidth(60)
ta.SetHeight(10)
ta.Focus()

// Init: return textarea.Blink
```

### List
```go
import "github.com/charmbracelet/bubbles/list"

// Items must implement list.Item interface
type item struct {
    title, desc string
}
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

items := []list.Item{item{"Go", "Programming language"}}
l := list.New(items, list.NewDefaultDelegate(), 40, 20)
l.Title = "Languages"
l.SetFilteringEnabled(true)

// Selected: l.SelectedItem()
```

### Viewport (Scrollable)
```go
import "github.com/charmbracelet/bubbles/viewport"

vp := viewport.New(80, 20)
vp.SetContent(longString)
vp.GotoTop()  // or GotoBottom()

// Scrolls with up/down/pgup/pgdn automatically
```

### Table
```go
import "github.com/charmbracelet/bubbles/table"

columns := []table.Column{
    {Title: "ID", Width: 4},
    {Title: "Name", Width: 20},
}
rows := []table.Row{
    {"1", "Alice"},
    {"2", "Bob"},
}

t := table.New(
    table.WithColumns(columns),
    table.WithRows(rows),
    table.WithFocused(true),
    table.WithHeight(10),
)

// Selected: t.SelectedRow()
```

### Progress
```go
import "github.com/charmbracelet/bubbles/progress"

p := progress.New(progress.WithDefaultGradient())
p.Width = 40

// View with percentage: p.ViewAs(0.5)  // 50%
```
