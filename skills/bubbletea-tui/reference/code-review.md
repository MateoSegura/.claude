# Bubble Tea Code Review Checklist

Use this checklist when reviewing Bubble Tea TUI code.

---

## Architecture

### Model Structure
- [ ] Model uses value receivers for Init, Update, View
- [ ] Model fields are organized (state, sub-models, dimensions)
- [ ] Complex UIs use composition with sub-models
- [ ] Screen/focus states use typed enums, not raw strings/ints

### Immutability
- [ ] Update returns new model, never mutates in place
- [ ] No pointer receivers on Model interface methods
- [ ] State changes happen through returned model

---

## Commands

### Command Implementation
- [ ] Commands are `func() tea.Msg`, not called immediately
- [ ] All I/O operations use timeouts
- [ ] Long operations have context cancellation
- [ ] `tea.Batch` used for concurrent commands
- [ ] `tea.Sequence` used when order matters

### Init Commands
- [ ] Init returns commands for animated components
- [ ] `spinner.Tick` returned for spinners
- [ ] `textinput.Blink` returned for text inputs

---

## Message Handling

### Message Types
- [ ] Custom message types defined, not using primitives
- [ ] Error messages use dedicated `errMsg` type
- [ ] Message types are documented

### Required Handlers
- [ ] `tea.WindowSizeMsg` handled for responsive layout
- [ ] `tea.KeyMsg` handled for user input
- [ ] Quit keys (q, ctrl+c) lead to `tea.Quit`
- [ ] Component tick messages forwarded (spinner.TickMsg, etc.)

---

## Components

### Bubbles Integration
- [ ] Components initialized in constructor, not Init
- [ ] All messages forwarded to child components
- [ ] Components resized on WindowSizeMsg
- [ ] Focus-based input routing implemented

### Component Initialization
- [ ] textinput: Placeholder, CharLimit, Width set
- [ ] spinner: Style and Spinner type set
- [ ] list: Items implement list.Item interface
- [ ] viewport: Width and Height set

---

## Styling (Lipgloss)

### Style Organization
- [ ] Styles defined as package-level variables
- [ ] No inline style definitions in View
- [ ] Consistent color palette used
- [ ] Theme package for shared styles

### Layout
- [ ] `lipgloss.Place` for centering
- [ ] `JoinVertical/JoinHorizontal` for composition
- [ ] Width/Height constraints for predictable sizing
- [ ] Padding/Margin used appropriately

### Borders & Colors
- [ ] Consistent border style throughout
- [ ] Colors specified as hex for consistency
- [ ] Focused/unfocused states visually distinct

---

## Program Options

- [ ] `tea.WithAltScreen()` for full-screen apps
- [ ] `tea.WithMouseCellMotion()` if mouse support needed
- [ ] External message injection uses `p.Send()`

---

## Testing

- [ ] Update function tested with various messages
- [ ] View output tested for expected content
- [ ] Commands tested for correct message return
- [ ] State transitions tested
- [ ] Key navigation tested with table-driven tests

---

## Common Issues to Check

### Critical
- [ ] No blocking operations in commands
- [ ] Messages forwarded to all child components
- [ ] WindowSizeMsg updates all component sizes
- [ ] Quit command accessible from all screens

### Required
- [ ] Error states displayed to user
- [ ] Loading states shown during async operations
- [ ] Help/hint text visible for key bindings
- [ ] Focus indicators visible

### Recommended
- [ ] Graceful degradation for small terminals
- [ ] Consistent key binding patterns (j/k, up/down)
- [ ] Clear visual hierarchy
- [ ] Status bar with context-appropriate hints

---

## Security Considerations

- [ ] User input validated before use
- [ ] No sensitive data in View output
- [ ] File paths validated/sanitized
- [ ] External commands use exec.Command properly

---

## Performance

- [ ] View doesn't perform expensive calculations
- [ ] Large lists use pagination or virtualization
- [ ] Tick intervals appropriate (not too frequent)
- [ ] Batch multiple state updates when possible

---

## Reviewer Notes

```
Architecture:      [ ] Pass  [ ] Needs Work
Commands:          [ ] Pass  [ ] Needs Work
Message Handling:  [ ] Pass  [ ] Needs Work
Components:        [ ] Pass  [ ] Needs Work
Styling:           [ ] Pass  [ ] Needs Work
Testing:           [ ] Pass  [ ] Needs Work

Overall:           [ ] Approve  [ ] Request Changes
```
