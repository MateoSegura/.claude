# Empty States Pattern

Empty states appear when there's no data to display. They should be helpful, not dead-ends.

## K9S-EMP-001: Empty State Structure :red_circle:

Empty states are centered with icon, message, and clear CTA:

```
◆ APPNAME                                                      Items
───────────────────────────────────────────────────────────────────────


                           ✦ No Items Yet

                 Create your first item to get started.


                      ▸ Press N to create an item


───────────────────────────────────────────────────────────────────────
<n>New  <q>Quit
```

```go
func renderEmptyState(icon, title, message, ctaKey, ctaText string) string {
    // Icon + Title (gold, bold)
    iconTitle := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Bold(true).
        Render(icon + " " + title)

    // Message (gray)
    msg := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render(message)

    // CTA (highlighted button style)
    cta := lipgloss.NewStyle().
        Foreground(theme.Black).
        Background(theme.Gold).
        Bold(true).
        Padding(0, 2).
        Render(theme.IconTriangleR + " Press " + ctaKey + " to " + ctaText)

    return lipgloss.JoinVertical(lipgloss.Center,
        iconTitle,
        "",
        msg,
        "",
        "",
        cta,
    )
}

// Usage:
emptyView := renderEmptyState(
    theme.IconSparkle,
    "No Projects Yet",
    "Create your first project to start orchestrating AI agents.",
    "N",
    "create a project",
)
```

## K9S-EMP-002: Contextual Empty States :yellow_circle:

Different empty states for different contexts:

```go
// No items at all
func renderNoItems() string {
    return renderEmptyState(
        "✦", "No Items",
        "Create your first item to get started.",
        "N", "create an item",
    )
}

// Search returned no results
func renderNoResults(query string) string {
    icon := lipgloss.NewStyle().Foreground(theme.Gray).Render("⌕")
    title := lipgloss.NewStyle().Foreground(theme.White).Render("No Results")

    msg := lipgloss.NewStyle().Foreground(theme.Gray).Render(
        fmt.Sprintf("No items match \"%s\"", query))

    hint := lipgloss.NewStyle().Foreground(theme.GrayDark).Render(
        "Try a different search term or clear the filter")

    return lipgloss.JoinVertical(lipgloss.Center,
        icon + " " + title,
        "",
        msg,
        "",
        hint,
    )
}

// Error loading data
func renderLoadError(err error) string {
    icon := lipgloss.NewStyle().Foreground(theme.Red).Bold(true).Render("✗")
    title := lipgloss.NewStyle().Foreground(theme.Red).Render("Failed to Load")

    msg := lipgloss.NewStyle().Foreground(theme.Gray).Render(err.Error())

    cta := lipgloss.NewStyle().Foreground(theme.Gold).Render(
        "Press R to retry")

    return lipgloss.JoinVertical(lipgloss.Center,
        icon + " " + title,
        "",
        msg,
        "",
        cta,
    )
}
```

## K9S-EMP-003: Loading State :yellow_circle:

While data is loading, show a spinner:

```go
func renderLoading(spinnerView string, message string) string {
    // Spinner (animated)
    spinner := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Render(spinnerView)

    // Message
    msg := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render(message)

    return lipgloss.JoinHorizontal(lipgloss.Center, spinner, " ", msg)
}

// In View:
if m.loading {
    loadingView := renderLoading(m.spinner.View(), "Loading items...")
    return m.chrome.RenderCentered(loadingView)
}
```

## K9S-EMP-004: First-Run Welcome :yellow_circle:

Special empty state for first-time users:

```go
func renderFirstRun(logo string) string {
    // Logo (gold)
    logoView := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Render(logo)

    // Tagline
    tagline := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render("Your helpful tagline here.")

    // Options
    option1 := lipgloss.NewStyle().
        Foreground(theme.Black).
        Background(theme.Gold).
        Bold(true).
        Padding(0, 2).
        Render("▸ Get Started")

    option2 := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Padding(0, 2).
        Render("  Skip, I know what I'm doing")

    return lipgloss.JoinVertical(lipgloss.Center,
        logoView,
        tagline,
        "",
        "",
        option1,
        option2,
    )
}
```

## K9S-EMP-005: Permission/Access Denied :yellow_circle:

When user lacks permission:

```go
func renderAccessDenied(resource string) string {
    icon := lipgloss.NewStyle().
        Foreground(theme.Orange).
        Bold(true).
        Render("⚠")

    title := lipgloss.NewStyle().
        Foreground(theme.Orange).
        Render("Access Denied")

    msg := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render(fmt.Sprintf("You don't have permission to access %s.", resource))

    hint := lipgloss.NewStyle().
        Foreground(theme.GrayDark).
        Render("Contact your administrator for access.")

    return lipgloss.JoinVertical(lipgloss.Center,
        icon + " " + title,
        "",
        msg,
        "",
        hint,
    )
}
```

## K9S-EMP-006: Connection Lost :yellow_circle:

When connection to backend is lost:

```go
func renderDisconnected() string {
    icon := lipgloss.NewStyle().
        Foreground(theme.Red).
        Bold(true).
        Render("●")

    title := lipgloss.NewStyle().
        Foreground(theme.Red).
        Render("Disconnected")

    msg := lipgloss.NewStyle().
        Foreground(theme.Gray).
        Render("Lost connection to the server.")

    cta := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Render("Press R to reconnect")

    status := lipgloss.NewStyle().
        Foreground(theme.GrayDark).
        Render("Retrying in 5s...")

    return lipgloss.JoinVertical(lipgloss.Center,
        icon + " " + title,
        "",
        msg,
        "",
        cta,
        "",
        status,
    )
}
```

## K9S-EMP-007: Filtered Empty :yellow_circle:

When filter/search excludes all items:

```go
func renderFilteredEmpty(activeFilters []string) string {
    icon := "⌕"
    title := "No Matching Items"

    // Show active filters
    filters := lipgloss.NewStyle().
        Foreground(theme.Cyan).
        Render(strings.Join(activeFilters, ", "))

    msg := lipgloss.JoinHorizontal(lipgloss.Left,
        lipgloss.NewStyle().Foreground(theme.Gray).Render("Active filters: "),
        filters,
    )

    cta := lipgloss.NewStyle().
        Foreground(theme.Gold).
        Render("Press Esc to clear filters")

    return lipgloss.JoinVertical(lipgloss.Center,
        lipgloss.NewStyle().Foreground(theme.Gray).Render(icon + " " + title),
        "",
        msg,
        "",
        cta,
    )
}
```
