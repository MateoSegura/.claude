# Commands Rules

## BTT-CMD-001: Commands Are Functions Returning Messages :red_circle:

**Tier**: Critical

A `tea.Cmd` is defined as `func() tea.Msg`. Commands perform side effects (I/O, timers, etc.) and return messages that the Update function can handle.

```go
// Definition
type Cmd func() Msg

// CORRECT - Command returns a message
func loadDataCmd() tea.Msg {
    data, err := fetchFromAPI()
    if err != nil {
        return errMsg{err}
    }
    return dataLoadedMsg{data}
}

// In Update, pass the function (don't call it):
return m, loadDataCmd  // CORRECT: pass function
return m, loadDataCmd() // WRONG: calls immediately, returns Msg not Cmd
```

**Common pattern with closures**:
```go
func (m Model) loadUserCmd(id string) tea.Cmd {
    return func() tea.Msg {
        user, err := m.client.GetUser(id)
        if err != nil {
            return errMsg{err}
        }
        return userLoadedMsg{user}
    }
}

// Usage:
return m, m.loadUserCmd("user-123")
```

---

## BTT-CMD-002: Use tea.Batch for Concurrent Commands :yellow_circle:

**Tier**: Required

`tea.Batch` runs multiple commands concurrently. Use it when commands are independent.

```go
func (m Model) Init() tea.Cmd {
    return tea.Batch(
        m.loadConfigCmd(),
        m.loadUserCmd(),
        m.spinner.Tick,
        textinput.Blink,
    )
}

// After multiple operations complete:
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case submitMsg:
        return m, tea.Batch(
            m.saveDataCmd(),
            m.logAnalyticsCmd(),
            m.showNotificationCmd(),
        )
    }
}
```

**Note**: Results from batched commands arrive in any order.

---

## BTT-CMD-003: Use tea.Sequence for Ordered Commands :yellow_circle:

**Tier**: Required

`tea.Sequence` runs commands one after another. Use when order matters.

```go
// Save data, THEN quit
cmd := tea.Sequence(
    m.saveDataCmd,
    tea.Quit,
)

// Validate, THEN submit, THEN show success
cmd := tea.Sequence(
    m.validateCmd,
    m.submitCmd,
    m.showSuccessCmd,
)
```

**Note**: If any command in the sequence returns `nil`, the sequence continues. If it returns a message, that message is processed before the next command runs.

---

## BTT-CMD-004: Never Block in Commands :red_circle:

**Tier**: Critical

Commands run in goroutines but should complete in reasonable time. Always use timeouts for I/O.

```go
// CORRECT - HTTP with timeout
func fetchDataCmd() tea.Msg {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return errMsg{err}
    }
    defer resp.Body.Close()

    var data Data
    json.NewDecoder(resp.Body).Decode(&data)
    return dataLoadedMsg{data}
}

// CORRECT - Channel with timeout
func waitForEventCmd(ch <-chan Event) tea.Cmd {
    return func() tea.Msg {
        select {
        case event := <-ch:
            return eventMsg{event}
        case <-time.After(30 * time.Second):
            return timeoutMsg{}
        }
    }
}

// INCORRECT - Blocks forever
func badCmd() tea.Msg {
    <-make(chan struct{})  // Never returns!
    return nil
}
```

---

## BTT-CMD-005: Use tea.Tick for Timers :yellow_circle:

**Tier**: Required

Use `tea.Tick` for recurring updates (animations, polling, etc.).

```go
type tickMsg time.Time

func tickCmd() tea.Cmd {
    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tickMsg:
        m.elapsed++
        return m, tickCmd()  // Schedule next tick
    }
    return m, nil
}
```

**For one-time delays**:
```go
func delayCmd(d time.Duration) tea.Cmd {
    return tea.Tick(d, func(t time.Time) tea.Msg {
        return delayCompleteMsg{}
    })
}
```

---

## BTT-CMD-006: tea.Quit and tea.ClearScreen :yellow_circle:

**Tier**: Required

Special built-in commands:

```go
// Quit the program
return m, tea.Quit

// Clear screen (useful before quitting from alt screen)
return m, tea.ClearScreen

// Sequence: clear then quit
return m, tea.Sequence(tea.ClearScreen, tea.Quit)
```

---

## BTT-CMD-007: Use tea.Printf for Debug Output :green_circle:

**Tier**: Recommended

Print debug messages without corrupting the TUI.

```go
// Correct - Uses tea.Printf
return m, tea.Printf("Debug: loaded %d items", len(items))

// Incorrect - fmt.Println corrupts alt screen
fmt.Println("Debug info")  // DON'T do this in TUI
```
