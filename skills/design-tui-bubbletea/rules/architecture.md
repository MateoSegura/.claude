# Architecture Rules

## BTT-ARC-001: Model Must Be Immutable :red_circle:

**Tier**: Critical

The Model-View-Update pattern requires immutable state. The `Update` function receives the current model and returns a new model (or the same one unchanged). Never mutate the model in place.

```go
// CORRECT - Value receiver, returns model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case incrementMsg:
        m.count++  // OK: m is a copy
        return m, nil
    }
    return m, nil
}

// INCORRECT - Pointer receiver mutation
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.count++  // BAD: mutating shared state
    return m, nil
}
```

**Why**: Immutability ensures:
- Predictable state transitions
- Easy debugging (each state is a snapshot)
- Potential for time-travel debugging
- No race conditions from shared mutable state

---

## BTT-ARC-002: Use Value Receivers for Model Methods :red_circle:

**Tier**: Critical

All three Model interface methods (`Init`, `Update`, `View`) must use value receivers.

```go
// CORRECT
func (m Model) Init() tea.Cmd { ... }
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }
func (m Model) View() string { ... }

// INCORRECT
func (m *Model) Init() tea.Cmd { ... }
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }
func (m *Model) View() string { ... }
```

**Exception**: Helper methods that don't implement the interface can use pointer receivers if they need to modify state, but Update should still copy and return.

---

## BTT-ARC-003: Compose Models for Complex UIs :yellow_circle:

**Tier**: Required

Break complex UIs into sub-models. Each sub-model manages its own state and responds to messages.

```go
type AppModel struct {
    // State
    screen    Screen
    focus     Focus
    width     int
    height    int

    // Sub-models (components)
    sidebar   SidebarModel
    content   ContentModel
    statusbar StatusbarModel
    input     textinput.Model
    spinner   spinner.Model
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Global message handling first
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        // Propagate to children
        m.sidebar.SetSize(m.width/3, m.height-2)
        m.content.SetSize(m.width*2/3, m.height-2)
    }

    // Forward to children (they filter what they need)
    var cmd tea.Cmd

    m.sidebar, cmd = m.sidebar.Update(msg)
    cmds = append(cmds, cmd)

    m.content, cmd = m.content.Update(msg)
    cmds = append(cmds, cmd)

    m.spinner, cmd = m.spinner.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}
```

**Benefits**:
- Separation of concerns
- Reusable components
- Easier testing
- Clearer code organization

---

## BTT-ARC-004: Use Enums for Screen/State Management :yellow_circle:

**Tier**: Required

Use typed constants for screen states and focus management.

```go
type Screen int

const (
    ScreenWelcome Screen = iota
    ScreenMain
    ScreenSettings
    ScreenHelp
)

type Focus int

const (
    FocusSidebar Focus = iota
    FocusContent
    FocusInput
)

type Model struct {
    screen Screen
    focus  Focus
    // ...
}

func (m Model) View() string {
    switch m.screen {
    case ScreenWelcome:
        return m.renderWelcome()
    case ScreenMain:
        return m.renderMain()
    case ScreenSettings:
        return m.renderSettings()
    default:
        return "Unknown screen"
    }
}
```

---

## BTT-ARC-005: Separate View Logic from Update Logic :green_circle:

**Tier**: Recommended

Keep `View()` focused on rendering. Complex rendering logic should be extracted to helper methods.

```go
func (m Model) View() string {
    return lipgloss.JoinVertical(lipgloss.Left,
        m.renderHeader(),
        m.renderContent(),
        m.renderFooter(),
    )
}

func (m Model) renderHeader() string {
    // Header rendering logic
}

func (m Model) renderContent() string {
    switch m.screen {
    case ScreenMain:
        return m.renderMainContent()
    case ScreenSettings:
        return m.renderSettingsContent()
    }
    return ""
}

func (m Model) renderFooter() string {
    // Footer with hints
}
```
