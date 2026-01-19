// Package screens - K9s-Style List Screen Scaffold
//
// USAGE: Copy this file and customize for your list view.
// Rename "Item" to your domain type, update rendering as needed.
package screens

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// =============================================================================
// THEME (inline for scaffold - use your theme package in real code)
// =============================================================================

var (
	colorGold      = lipgloss.Color("#FFD700")
	colorBlack     = lipgloss.Color("#0A0A0A")
	colorBlackLight = lipgloss.Color("#1A1A1A")
	colorWhite     = lipgloss.Color("#FAFAFA")
	colorGray      = lipgloss.Color("#6B6B6B")
	colorGrayDark  = lipgloss.Color("#3D3D3D")
	colorLime      = lipgloss.Color("#39FF14")
	colorRed       = lipgloss.Color("#FF073A")
	colorCyan      = lipgloss.Color("#00D9FF")
)

// =============================================================================
// ICONS
// =============================================================================

const (
	iconDiamond   = "◆"
	iconTriangleR = "▸"
	iconCheck     = "✓"
	iconCross     = "✗"
	iconCircle    = "●"
	iconCircleO   = "○"
	iconSparkle   = "✦"
	iconBullet    = "•"
)

// =============================================================================
// ITEM TYPE (customize this for your domain)
// =============================================================================

// Item represents a list item. Customize fields for your use case.
type Item struct {
	ID        string
	Name      string
	Path      string
	State     string // "running", "success", "failed", "pending"
	UpdatedAt time.Time
}

// =============================================================================
// MESSAGES
// =============================================================================

type ItemSelectedMsg struct {
	Item Item
}

type ItemDeleteMsg struct {
	Item Item
}

type ItemCreateMsg struct{}

// =============================================================================
// KEY BINDINGS
// =============================================================================

type listKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	New    key.Binding
	Delete key.Binding
	Quit   key.Binding
}

var listKeys = listKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	New: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// =============================================================================
// MODEL
// =============================================================================

// ListScreenModel is a K9s-style list screen.
type ListScreenModel struct {
	items    []Item
	cursor   int
	width    int
	height   int
	title    string
}

// NewListScreen creates a new list screen.
func NewListScreen(title string) ListScreenModel {
	return ListScreenModel{
		title: title,
	}
}

// SetItems updates the items list.
func (m *ListScreenModel) SetItems(items []Item) {
	m.items = items
	if m.cursor >= len(items) {
		m.cursor = max(0, len(items)-1)
	}
}

// SetSize updates dimensions.
func (m *ListScreenModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// =============================================================================
// TEA.MODEL IMPLEMENTATION
// =============================================================================

func (m ListScreenModel) Init() tea.Cmd {
	return nil
}

func (m ListScreenModel) Update(msg tea.Msg) (ListScreenModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, listKeys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, listKeys.Down):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case key.Matches(msg, listKeys.Select):
			if len(m.items) > 0 {
				return m, func() tea.Msg {
					return ItemSelectedMsg{Item: m.items[m.cursor]}
				}
			}

		case key.Matches(msg, listKeys.New):
			return m, func() tea.Msg {
				return ItemCreateMsg{}
			}

		case key.Matches(msg, listKeys.Delete):
			if len(m.items) > 0 {
				return m, func() tea.Msg {
					return ItemDeleteMsg{Item: m.items[m.cursor]}
				}
			}

		case key.Matches(msg, listKeys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ListScreenModel) View() string {
	// Header
	header := m.renderHeader()

	// Content
	var content string
	if len(m.items) == 0 {
		content = m.renderEmpty()
	} else {
		content = m.renderList()
	}

	// Footer
	footer := m.renderFooter()

	// Compose
	contentHeight := m.height - 2 // header + footer
	contentArea := lipgloss.NewStyle().
		Width(m.width).
		Height(contentHeight).
		Render(content)

	return lipgloss.JoinVertical(lipgloss.Left, header, contentArea, footer)
}

// =============================================================================
// RENDER HELPERS
// =============================================================================

func (m ListScreenModel) renderHeader() string {
	left := lipgloss.NewStyle().
		Foreground(colorGold).
		Bold(true).
		Render(iconDiamond + " " + m.title)

	right := lipgloss.NewStyle().
		Foreground(colorGrayDark).
		Render(fmt.Sprintf("Items [%d]", len(m.items)))

	gap := m.width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}

	return lipgloss.NewStyle().
		Background(colorBlackLight).
		Width(m.width).
		Padding(0, 1).
		Render(left + strings.Repeat(" ", gap) + right)
}

func (m ListScreenModel) renderFooter() string {
	var shortcuts []string

	if len(m.items) > 0 {
		shortcuts = []string{
			renderShortcut("↑↓", "Navigate"),
			renderShortcut("Enter", "Select"),
			renderShortcut("n", "New"),
			renderShortcut("d", "Delete"),
			renderShortcut("q", "Quit"),
		}
	} else {
		shortcuts = []string{
			renderShortcut("n", "New"),
			renderShortcut("q", "Quit"),
		}
	}

	return lipgloss.NewStyle().
		Background(colorBlackLight).
		Width(m.width).
		Padding(0, 1).
		Render(strings.Join(shortcuts, "  "))
}

func renderShortcut(key, desc string) string {
	k := lipgloss.NewStyle().Foreground(colorGold).Bold(true).Render("<" + key + ">")
	d := lipgloss.NewStyle().Foreground(colorGray).Render(desc)
	return k + d
}

func (m ListScreenModel) renderEmpty() string {
	icon := lipgloss.NewStyle().
		Foreground(colorGold).
		Bold(true).
		Render(iconSparkle + " No Items Yet")

	msg := lipgloss.NewStyle().
		Foreground(colorGray).
		Render("Create your first item to get started.")

	cta := lipgloss.NewStyle().
		Foreground(colorBlack).
		Background(colorGold).
		Bold(true).
		Padding(0, 2).
		Render(iconTriangleR + " Press N to create an item")

	content := lipgloss.JoinVertical(lipgloss.Center, icon, "", msg, "", "", cta)

	return lipgloss.Place(
		m.width,
		m.height-2,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m ListScreenModel) renderList() string {
	var items []string
	for i, item := range m.items {
		items = append(items, m.renderItem(item, i == m.cursor))
	}
	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

func (m ListScreenModel) renderItem(item Item, selected bool) string {
	// State icon and color
	stateIcon, stateColor := getStateIconColor(item.State)
	icon := lipgloss.NewStyle().Foreground(stateColor).Render(stateIcon)

	// Selector
	var selector string
	var nameStyle lipgloss.Style
	if selected {
		selector = lipgloss.NewStyle().Foreground(colorGold).Render(iconTriangleR + " ")
		nameStyle = lipgloss.NewStyle().Foreground(colorGold).Bold(true)
	} else {
		selector = "  "
		nameStyle = lipgloss.NewStyle().Foreground(colorWhite)
	}

	// Name
	name := nameStyle.Width(20).Render(item.Name)

	// Path (truncated)
	path := lipgloss.NewStyle().
		Foreground(colorGray).
		Width(35).
		Render(truncatePath(item.Path, 35))

	// Time
	timeStr := lipgloss.NewStyle().
		Foreground(colorGrayDark).
		Render(formatRelativeTime(item.UpdatedAt))

	// Line 1
	line1 := selector + icon + " " + name + path + timeStr

	return line1 + "\n"
}

func getStateIconColor(state string) (string, lipgloss.Color) {
	switch state {
	case "running":
		return iconCircle, colorLime
	case "success", "completed":
		return iconCheck, colorGold
	case "failed", "error":
		return iconCross, colorRed
	case "pending":
		return iconCircleO, colorGray
	case "new":
		return iconSparkle, colorCyan
	default:
		return iconCircleO, colorGray
	}
}

func truncatePath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}
	return "…" + path[len(path)-maxLen+1:]
}

func formatRelativeTime(t time.Time) string {
	if t.IsZero() {
		return "never"
	}
	dur := time.Since(t)
	switch {
	case dur < time.Minute:
		return "just now"
	case dur < time.Hour:
		return fmt.Sprintf("%dm ago", int(dur.Minutes()))
	case dur < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(dur.Hours()))
	case dur < 7*24*time.Hour:
		return fmt.Sprintf("%dd ago", int(dur.Hours()/24))
	default:
		return t.Format("Jan 2")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
