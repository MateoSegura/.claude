# Modal Pattern

K9s-style modals are centered overlays for confirmations and focused interactions.

## K9S-MOD-001: Modal Structure :red_circle:

Modals float above content with a bordered container:

```
┌─────────────────────────────────────────────────────────────────────┐
│ ◆ APPNAME                                              Context     │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│                    ╭───────────────────────────────╮                │
│                    │                               │                │
│                    │       Modal Title             │                │
│                    │                               │                │
│                    │       Modal content here      │                │
│                    │                               │                │
│                    │   <y>Yes  <n>No  <Esc>Cancel │                │
│                    │                               │                │
│                    ╰───────────────────────────────╯                │
│                                                                     │
├─────────────────────────────────────────────────────────────────────┤
│ (shortcuts hidden or modal-specific)                                │
└─────────────────────────────────────────────────────────────────────┘
```

## K9S-MOD-002: Confirmation Dialog :red_circle:

Simple yes/no confirmation with destructive action warning:

```go
func renderConfirmDialog(title, message string, width int) string {
    // Title
    titleStyle := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true)

    // Message
    msgStyle := lipgloss.NewStyle().
        Foreground(theme.White)

    // Warning for destructive actions
    warnStyle := lipgloss.NewStyle().
        Foreground(theme.Red)

    // Shortcuts
    yesKey := lipgloss.NewStyle().Foreground(theme.Lime).Bold(true).Render("<y>")
    noKey := lipgloss.NewStyle().Foreground(theme.Red).Bold(true).Render("<n>")
    escKey := lipgloss.NewStyle().Foreground(theme.Gray).Render("<Esc>")

    shortcuts := yesKey + "Yes  " + noKey + "No  " + escKey + "Cancel"

    content := lipgloss.JoinVertical(lipgloss.Center,
        titleStyle.Render(title),
        "",
        msgStyle.Render(message),
        "",
        shortcuts,
    )

    // Modal box
    return lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(theme.Gold).
        Padding(1, 3).
        Render(content)
}
```

## K9S-MOD-003: Delete Confirmation :yellow_circle:

Deletion dialogs show the item being deleted:

```go
func renderDeleteConfirm(itemName string) string {
    title := lipgloss.NewStyle().
        Foreground(theme.Red).
        Bold(true).
        Render("Delete Item")

    itemStyle := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true)

    message := lipgloss.JoinVertical(lipgloss.Center,
        "Are you sure you want to delete",
        itemStyle.Render(itemName),
        "",
        lipgloss.NewStyle().Foreground(theme.GrayDark).
            Render("This action cannot be undone."),
    )

    return renderConfirmDialog("Delete Item", message, 50)
}
```

## K9S-MOD-004: Info Modal :yellow_circle:

Display detailed information in a modal:

```go
func renderInfoModal(title string, info map[string]string) string {
    // Title
    titleView := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true).
        Render(title)

    // Key-value pairs
    labelStyle := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Width(15).
        Align(lipgloss.Right)

    valueStyle := lipgloss.NewStyle().
        Foreground(theme.White)

    var rows []string
    for label, value := range info {
        row := labelStyle.Render(label+":") + "  " + valueStyle.Render(value)
        rows = append(rows, row)
    }

    // Close hint
    hint := lipgloss.NewStyle().
        Foreground(theme.GrayDark).
        Render("Press Esc to close")

    content := lipgloss.JoinVertical(lipgloss.Left,
        titleView,
        "",
        lipgloss.JoinVertical(lipgloss.Left, rows...),
        "",
        hint,
    )

    return lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(theme.Charcoal).
        Padding(1, 3).
        Render(content)
}
```

## K9S-MOD-005: Modal Rendering in View :red_circle:

Render modal centered over the chrome content:

```go
func (m Model) View() string {
    // Base view with chrome
    base := m.chrome.Render(m.renderContent())

    // If modal is active, overlay it
    if m.showModal {
        modal := m.renderModal()

        // Center modal over base
        return lipgloss.Place(
            m.width,
            m.height,
            lipgloss.Center,
            lipgloss.Center,
            modal,
            lipgloss.WithWhitespaceChars(" "),
            lipgloss.WithWhitespaceForeground(lipgloss.Color("#000000")),
        )
    }

    return base
}
```

## K9S-MOD-006: Modal Key Handling :red_circle:

Modal captures all input when active:

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // If modal is active, route all input to modal
    if m.showModal {
        return m.updateModal(msg)
    }

    // Normal update logic
    return m.updateNormal(msg)
}

func (m Model) updateModal(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "y", "Y":
            m.showModal = false
            return m, m.confirmAction()
        case "n", "N", "esc":
            m.showModal = false
            return m, nil
        }
    }
    return m, nil
}
```

## K9S-MOD-007: Input Modal :yellow_circle:

Modal with text input for rename, create, etc:

```go
type InputModal struct {
    title       string
    label       string
    input       textinput.Model
    width       int
    onSubmit    func(string) tea.Cmd
    onCancel    func() tea.Cmd
}

func (m InputModal) View() string {
    title := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true).
        Render(m.title)

    label := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render(m.label)

    inputBox := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(theme.Gold).
        Padding(0, 1).
        Width(40).
        Render(m.input.View())

    hints := lipgloss.NewStyle().
        Foreground(theme.GrayDark).
        Render("<Enter>Submit  <Esc>Cancel")

    content := lipgloss.JoinVertical(lipgloss.Center,
        title,
        "",
        label,
        inputBox,
        "",
        hints,
    )

    return lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).
        BorderForeground(theme.Gold).
        Padding(1, 3).
        Render(content)
}
```

## K9S-MOD-008: Toast/Notification :yellow_circle:

Brief messages that appear and auto-dismiss:

```go
type Toast struct {
    message  string
    severity string // "success", "error", "info"
    duration time.Duration
}

func renderToast(t Toast) string {
    var style lipgloss.Style
    var icon string

    switch t.severity {
    case "success":
        style = lipgloss.NewStyle().
            Foreground(theme.Black).
            Background(theme.Lime)
        icon = "✓ "
    case "error":
        style = lipgloss.NewStyle().
            Foreground(theme.White).
            Background(theme.Red)
        icon = "✗ "
    default:
        style = lipgloss.NewStyle().
            Foreground(theme.Black).
            Background(theme.Cyan)
        icon = "ℹ "
    }

    return style.Padding(0, 2).Render(icon + t.message)
}

// Position toast at bottom of screen
func (m Model) View() string {
    base := m.chrome.Render(m.content())

    if m.toast != nil {
        toast := renderToast(*m.toast)
        // Position at bottom-right
        base = lipgloss.JoinVertical(lipgloss.Right, base, toast)
    }

    return base
}
```
