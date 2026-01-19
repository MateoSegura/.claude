# Testing Rules

## BTT-TST-001: Test Update Function Logic :yellow_circle:

**Tier**: Required

The `Update` function is pure (given state + message, returns new state + commands). This makes it easy to test.

```go
func TestModel_Update_KeyPress(t *testing.T) {
    m := initialModel()

    // Simulate key press
    msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}

    newModel, cmd := m.Update(msg)
    result := newModel.(model)

    // Assert state changed correctly
    if result.cursor != 1 {
        t.Errorf("expected cursor=1, got %d", result.cursor)
    }

    // Assert no command returned
    if cmd != nil {
        t.Error("expected nil command")
    }
}

func TestModel_Update_Quit(t *testing.T) {
    m := initialModel()

    msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}
    _, cmd := m.Update(msg)

    // Check that quit command was returned
    // Note: tea.Quit is a func, compare behavior
    if cmd == nil {
        t.Error("expected quit command")
    }
}
```

---

## BTT-TST-002: Test View Output :yellow_circle:

**Tier**: Required

Test that View produces expected output for given state.

```go
func TestModel_View_ShowsTitle(t *testing.T) {
    m := model{
        title: "Test Title",
        width: 80,
        height: 24,
    }

    view := m.View()

    if !strings.Contains(view, "Test Title") {
        t.Error("view should contain title")
    }
}

func TestModel_View_HighlightsSelected(t *testing.T) {
    m := model{
        items:  []string{"One", "Two", "Three"},
        cursor: 1,
    }

    view := m.View()

    // Check cursor indicator is on correct line
    lines := strings.Split(view, "\n")
    for i, line := range lines {
        if strings.Contains(line, "Two") {
            if !strings.Contains(line, ">") {
                t.Error("selected item should have cursor")
            }
        }
    }
}
```

---

## BTT-TST-003: Test Commands Return Expected Messages :yellow_circle:

**Tier**: Required

Commands are functions that return messages. Test that they return the right message type.

```go
func TestLoadDataCmd_Success(t *testing.T) {
    // Setup mock/fake data source
    m := model{
        client: &mockClient{data: []string{"a", "b"}},
    }

    cmd := m.loadDataCmd()
    msg := cmd()  // Execute command

    // Assert message type
    dataMsg, ok := msg.(dataLoadedMsg)
    if !ok {
        t.Fatalf("expected dataLoadedMsg, got %T", msg)
    }

    if len(dataMsg.data) != 2 {
        t.Errorf("expected 2 items, got %d", len(dataMsg.data))
    }
}

func TestLoadDataCmd_Error(t *testing.T) {
    m := model{
        client: &mockClient{err: errors.New("network error")},
    }

    cmd := m.loadDataCmd()
    msg := cmd()

    errMsg, ok := msg.(errMsg)
    if !ok {
        t.Fatalf("expected errMsg, got %T", msg)
    }

    if errMsg.err.Error() != "network error" {
        t.Errorf("unexpected error: %v", errMsg.err)
    }
}
```

---

## BTT-TST-004: Test Message Handling Transitions :yellow_circle:

**Tier**: Required

Test that specific messages cause correct state transitions.

```go
func TestModel_HandleDataLoaded(t *testing.T) {
    m := model{loading: true}

    msg := dataLoadedMsg{data: []string{"item1", "item2"}}
    newModel, _ := m.Update(msg)
    result := newModel.(model)

    if result.loading {
        t.Error("loading should be false after data loaded")
    }
    if len(result.items) != 2 {
        t.Errorf("expected 2 items, got %d", len(result.items))
    }
}

func TestModel_HandleWindowResize(t *testing.T) {
    m := model{}

    msg := tea.WindowSizeMsg{Width: 120, Height: 40}
    newModel, _ := m.Update(msg)
    result := newModel.(model)

    if result.width != 120 || result.height != 40 {
        t.Errorf("expected 120x40, got %dx%d", result.width, result.height)
    }
}
```

---

## BTT-TST-005: Use Table-Driven Tests for Key Handling :green_circle:

**Tier**: Recommended

```go
func TestModel_KeyNavigation(t *testing.T) {
    tests := []struct {
        name       string
        initial    int
        key        string
        expected   int
        itemCount  int
    }{
        {"down from top", 0, "j", 1, 5},
        {"down from middle", 2, "j", 3, 5},
        {"down at bottom", 4, "j", 4, 5},  // Should not go past last
        {"up from bottom", 4, "k", 3, 5},
        {"up from middle", 2, "k", 1, 5},
        {"up at top", 0, "k", 0, 5},       // Should not go negative
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            m := model{
                cursor: tt.initial,
                items:  make([]string, tt.itemCount),
            }

            msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.key)}
            newModel, _ := m.Update(msg)
            result := newModel.(model)

            if result.cursor != tt.expected {
                t.Errorf("expected cursor=%d, got %d", tt.expected, result.cursor)
            }
        })
    }
}
```

---

## BTT-TST-006: Test Component Integration :green_circle:

**Tier**: Recommended

Test that child components receive and process messages correctly.

```go
func TestModel_SpinnerUpdates(t *testing.T) {
    m := NewModel()  // Creates model with spinner

    // Simulate spinner tick
    msg := spinner.TickMsg{}
    newModel, cmd := m.Update(msg)
    result := newModel.(Model)

    // Spinner should return another tick command
    if cmd == nil {
        t.Error("spinner should return tick command")
    }

    // View should include spinner
    view := result.View()
    // Spinner character should be present (depends on frame)
}
```

---

## BTT-TST-007: Test Init Returns Required Commands :yellow_circle:

**Tier**: Required

```go
func TestModel_Init_ReturnsCommands(t *testing.T) {
    m := NewModel()
    cmd := m.Init()

    if cmd == nil {
        t.Error("Init should return commands for animated components")
    }

    // If using Batch, you can't easily inspect individual commands
    // But you can test the effect by running them
}
```

---

## BTT-TST-008: Snapshot Testing for Complex Views :green_circle:

**Tier**: Recommended

For complex views, use golden files or snapshot testing.

```go
func TestModel_View_Snapshot(t *testing.T) {
    m := model{
        title:  "Test",
        items:  []string{"One", "Two", "Three"},
        cursor: 1,
        width:  40,
        height: 10,
    }

    got := m.View()

    // Compare against golden file
    golden := filepath.Join("testdata", "view_snapshot.golden")

    if *update {
        os.WriteFile(golden, []byte(got), 0644)
        return
    }

    want, _ := os.ReadFile(golden)
    if got != string(want) {
        t.Errorf("view mismatch:\ngot:\n%s\nwant:\n%s", got, want)
    }
}
```

---

## Testing Utilities

### Creating Key Messages
```go
// Character key
tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}

// Special keys
tea.KeyMsg{Type: tea.KeyEnter}
tea.KeyMsg{Type: tea.KeyEsc}
tea.KeyMsg{Type: tea.KeyUp}
tea.KeyMsg{Type: tea.KeyDown}
tea.KeyMsg{Type: tea.KeyCtrlC}

// Ctrl+key combinations
tea.KeyMsg{Type: tea.KeyCtrlC}
tea.KeyMsg{Type: tea.KeyCtrlD}
```

### Mock Client Pattern
```go
type mockClient struct {
    data []string
    err  error
}

func (m *mockClient) FetchData() ([]string, error) {
    return m.data, m.err
}
```
