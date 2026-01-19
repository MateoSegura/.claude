# Form Pattern

K9s-style forms are clean, focused, and keyboard-navigable.

## K9S-FRM-001: Form Layout :red_circle:

Forms are centered with labeled fields in a consistent layout:

```
                        Form Title

                Name         ╭────────────────────────╮
                             │ value here             │
                             ╰────────────────────────╯

                Path         ╭────────────────────────╮
                             │ another value          │
                             ╰────────────────────────╯

                Helper text goes here
```

```go
func renderForm(title string, fields []FormField) string {
    // Title
    titleView := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true).
        Render(title)

    // Fields
    var fieldViews []string
    labelStyle := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Width(12).
        Align(lipgloss.Right)

    for _, f := range fields {
        label := labelStyle.Render(f.Label)

        inputStyle := theme.Input
        if f.Focused {
            inputStyle = theme.InputFocused
        }
        input := inputStyle.Width(40).Render(f.Input.View())

        row := lipgloss.JoinHorizontal(lipgloss.Left, label, "  ", input)
        fieldViews = append(fieldViews, row, "")
    }

    return lipgloss.JoinVertical(lipgloss.Center,
        titleView,
        "",
        lipgloss.JoinVertical(lipgloss.Left, fieldViews...),
    )
}
```

## K9S-FRM-002: Input Field Styling :red_circle:

Inputs have different styles for focused/unfocused states:

```go
var (
    // Unfocused input
    Input = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("#252525")). // Charcoal
        Padding(0, 1)

    // Focused input (gold border)
    InputFocused = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#FFD700")). // Gold
        Padding(0, 1)
)
```

## K9S-FRM-003: Tab Navigation :red_circle:

Tab moves between fields, Enter submits:

```go
case tea.KeyMsg:
    switch msg.String() {
    case "tab", "shift+tab":
        // Move to next/prev field
        if msg.String() == "tab" {
            m.focusIndex = (m.focusIndex + 1) % len(m.fields)
        } else {
            m.focusIndex = (m.focusIndex - 1 + len(m.fields)) % len(m.fields)
        }
        // Update focus
        for i := range m.fields {
            if i == m.focusIndex {
                m.fields[i].Focus()
            } else {
                m.fields[i].Blur()
            }
        }
        return m, textinput.Blink

    case "enter":
        // Validate and submit
        if m.validate() {
            return m, m.submitCmd()
        }

    case "esc":
        // Cancel form
        return m, func() tea.Msg { return FormCancelledMsg{} }
    }
```

## K9S-FRM-004: Input Configuration :yellow_circle:

Configure bubbles textinput for K9s style:

```go
func newStyledInput(placeholder string, width int) textinput.Model {
    ti := textinput.New()
    ti.Placeholder = placeholder
    ti.Prompt = ""  // No prompt character
    ti.Width = width
    ti.CharLimit = 256

    // Text styling
    ti.TextStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FAFAFA"))  // White

    // Placeholder styling
    ti.PlaceholderStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#3D3D3D"))  // GrayDark

    // Cursor styling
    ti.Cursor.Style = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFD700"))  // Gold

    return ti
}
```

## K9S-FRM-005: Helper Text :yellow_circle:

Show contextual help below forms:

```go
func renderHelperText(text string) string {
    return lipgloss.NewStyle().
        Foreground(lipgloss.Color("#3D3D3D")).  // GrayDark
        Render(text)
}

// Usage:
helperText := renderHelperText("Workspace should be an existing directory")
```

## K9S-FRM-006: Validation Feedback :yellow_circle:

Show validation errors inline:

```go
func renderFieldWithError(label, inputView, errorMsg string, focused bool) string {
    labelStyle := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Width(12)

    inputStyle := theme.Input
    if focused {
        inputStyle = theme.InputFocused
    }
    if errorMsg != "" {
        inputStyle = inputStyle.BorderForeground(theme.Red)
    }

    row := lipgloss.JoinHorizontal(lipgloss.Left,
        labelStyle.Render(label),
        "  ",
        inputStyle.Width(40).Render(inputView),
    )

    if errorMsg != "" {
        errStyle := lipgloss.NewStyle().
            Foreground(theme.Red).
            MarginLeft(14)
        row = lipgloss.JoinVertical(lipgloss.Left, row, errStyle.Render(errorMsg))
    }

    return row
}
```

## K9S-FRM-007: Password Fields :yellow_circle:

Use EchoMode for sensitive inputs:

```go
passwordInput := textinput.New()
passwordInput.EchoMode = textinput.EchoPassword
passwordInput.EchoCharacter = '•'
```

## K9S-FRM-008: Form Shortcuts :red_circle:

Forms have specific shortcuts (update footer):

```go
var ShortcutsForm = []Shortcut{
    {Key: "Tab", Desc: "Next"},
    {Key: "Enter", Desc: "Submit"},
    {Key: "Esc", Desc: "Cancel"},
}
```

## K9S-FRM-009: Multi-line Input :yellow_circle:

For larger text, use textarea:

```go
import "github.com/charmbracelet/bubbles/textarea"

func newStyledTextArea(placeholder string, width, height int) textarea.Model {
    ta := textarea.New()
    ta.Placeholder = placeholder
    ta.SetWidth(width)
    ta.SetHeight(height)
    ta.ShowLineNumbers = false
    ta.FocusedStyle.CursorLine = lipgloss.NewStyle()  // No highlight
    ta.FocusedStyle.Base = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(theme.Gold)
    ta.BlurredStyle.Base = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(theme.Charcoal)
    return ta
}
```

## K9S-FRM-010: Dropdown/Select :yellow_circle:

For selection from options, render as a mini-list:

```go
func renderSelect(options []string, selected int, focused bool) string {
    var lines []string
    for i, opt := range options {
        var line string
        if i == selected {
            if focused {
                line = lipgloss.NewStyle().
                    Foreground(theme.Black).
                    Background(theme.Gold).
                    Padding(0, 1).
                    Render("▸ " + opt)
            } else {
                line = lipgloss.NewStyle().
                    Foreground(theme.Gold).
                    Padding(0, 1).
                    Render("▸ " + opt)
            }
        } else {
            line = lipgloss.NewStyle().
                Foreground(theme.Gray).
                Padding(0, 1).
                Render("  " + opt)
        }
        lines = append(lines, line)
    }
    return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
```
