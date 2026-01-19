// Package components - Reusable Bubble Tea Component Scaffold
//
// USAGE: Copy this file, rename the struct/methods, and customize.
// Components follow the same MVU pattern but return their own type from Update.
package components

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// === STYLES ===

var (
	componentStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#6B6B6B")).
			Padding(1, 2)

	componentFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FFD700")).
				Padding(1, 2)
)

// === KEY BINDINGS ===

type ComponentKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
}

var componentKeys = ComponentKeyMap{
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
}

// === MESSAGES ===

// ComponentSelectedMsg is sent when an item is selected.
// Export this so parent models can handle it.
type ComponentSelectedMsg struct {
	Index int
	Value string
}

// === MODEL ===

// ComponentModel is a reusable component.
// Rename this to match your component's purpose (e.g., SidebarModel, ListModel).
type ComponentModel struct {
	items    []string
	cursor   int
	focused  bool
	width    int
	height   int
}

// NewComponentModel creates a new component.
func NewComponentModel(items []string) ComponentModel {
	return ComponentModel{
		items:   items,
		focused: false,
	}
}

// === PUBLIC METHODS ===

// SetFocused sets the focus state.
func (m *ComponentModel) SetFocused(focused bool) {
	m.focused = focused
}

// SetSize sets the component dimensions.
func (m *ComponentModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// SelectedIndex returns the current cursor position.
func (m ComponentModel) SelectedIndex() int {
	return m.cursor
}

// SelectedItem returns the currently selected item.
func (m ComponentModel) SelectedItem() string {
	if m.cursor >= 0 && m.cursor < len(m.items) {
		return m.items[m.cursor]
	}
	return ""
}

// SetItems updates the items list.
func (m *ComponentModel) SetItems(items []string) {
	m.items = items
	if m.cursor >= len(items) {
		m.cursor = len(items) - 1
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
}

// === TEA.MODEL INTERFACE ===

// Init implements tea.Model.
// Components typically don't need initialization commands.
func (m ComponentModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
// NOTE: Returns ComponentModel, not tea.Model - this is the component pattern.
func (m ComponentModel) Update(msg tea.Msg) (ComponentModel, tea.Cmd) {
	// Only process input when focused
	if !m.focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, componentKeys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, componentKeys.Down):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case key.Matches(msg, componentKeys.Select):
			if len(m.items) > 0 {
				return m, func() tea.Msg {
					return ComponentSelectedMsg{
						Index: m.cursor,
						Value: m.items[m.cursor],
					}
				}
			}
		}
	}

	return m, nil
}

// View implements tea.Model.
func (m ComponentModel) View() string {
	if len(m.items) == 0 {
		return componentStyle.Render("No items")
	}

	var content string
	for i, item := range m.items {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}

		line := cursor + item

		if i == m.cursor && m.focused {
			line = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFD700")).
				Render(line)
		}

		content += line + "\n"
	}

	// Choose style based on focus
	style := componentStyle
	if m.focused {
		style = componentFocusedStyle
	}

	return style.
		Width(m.width).
		Height(m.height).
		Render(content)
}

// Help returns key binding help text.
func (m ComponentModel) Help() string {
	return "↑/k: up  ↓/j: down  enter: select"
}
