// Package screens - Full-Screen View Scaffold
//
// USAGE: Copy this file for each major screen in your application.
// Screens are typically composed of multiple components and handle navigation.
package screens

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// === STYLES ===

var (
	headerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1A1A1A")).
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			Padding(0, 1)

	contentStyle = lipgloss.NewStyle().
			Padding(1, 2)

	footerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1A1A1A")).
			Foreground(lipgloss.Color("#6B6B6B")).
			Padding(0, 1)
)

// === KEY BINDINGS ===

type ScreenKeyMap struct {
	Back key.Binding
	Help key.Binding
	Quit key.Binding
}

var screenKeys = ScreenKeyMap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// === MESSAGES ===

// ScreenBackMsg signals navigation back to previous screen.
type ScreenBackMsg struct{}

// ScreenDataLoadedMsg carries loaded data.
type ScreenDataLoadedMsg struct {
	Data interface{}
}

// ScreenErrorMsg carries an error.
type ScreenErrorMsg struct {
	Err error
}

// === MODEL ===

// ScreenModel represents a full-screen view.
type ScreenModel struct {
	title     string
	showHelp  bool
	loading   bool
	err       error
	width     int
	height    int

	// Add your sub-components here:
	// list    ListModel
	// detail  DetailModel
	// input   textinput.Model
}

// NewScreenModel creates a new screen.
func NewScreenModel(title string) ScreenModel {
	return ScreenModel{
		title: title,
	}
}

// === PUBLIC METHODS ===

// SetSize updates screen dimensions and propagates to children.
func (m *ScreenModel) SetSize(width, height int) {
	m.width = width
	m.height = height

	// Propagate to children:
	// contentHeight := height - 4 // header + footer
	// m.list.SetSize(width/3, contentHeight)
	// m.detail.SetSize(width*2/3, contentHeight)
}

// SetLoading sets the loading state.
func (m *ScreenModel) SetLoading(loading bool) {
	m.loading = loading
}

// SetError sets an error to display.
func (m *ScreenModel) SetError(err error) {
	m.err = err
}

// === TEA.MODEL INTERFACE ===

// Init implements tea.Model.
func (m ScreenModel) Init() tea.Cmd {
	// Return initial commands, e.g., load data
	// return tea.Batch(
	//     m.loadDataCmd(),
	//     m.spinner.Tick,
	// )
	return nil
}

// Update implements tea.Model.
// NOTE: Returns ScreenModel, not tea.Model - consistent with component pattern.
func (m ScreenModel) Update(msg tea.Msg) (ScreenModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		// Global screen keys
		switch {
		case key.Matches(msg, screenKeys.Back):
			return m, func() tea.Msg { return ScreenBackMsg{} }

		case key.Matches(msg, screenKeys.Help):
			m.showHelp = !m.showHelp
			return m, nil

		case key.Matches(msg, screenKeys.Quit):
			return m, tea.Quit
		}

	case ScreenDataLoadedMsg:
		m.loading = false
		// Process loaded data
		return m, nil

	case ScreenErrorMsg:
		m.loading = false
		m.err = msg.Err
		return m, nil
	}

	// Forward to child components:
	// var cmd tea.Cmd
	// m.list, cmd = m.list.Update(msg)
	// cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View implements tea.Model.
func (m ScreenModel) View() string {
	// Build header
	header := m.renderHeader()

	// Build content
	var content string
	if m.err != nil {
		content = m.renderError()
	} else if m.loading {
		content = m.renderLoading()
	} else if m.showHelp {
		content = m.renderHelp()
	} else {
		content = m.renderContent()
	}

	// Build footer
	footer := m.renderFooter()

	// Calculate content height
	contentHeight := m.height - 2 // header + footer lines

	// Render content area
	contentArea := contentStyle.
		Width(m.width).
		Height(contentHeight).
		Render(content)

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		contentArea,
		footer,
	)
}

// === RENDER HELPERS ===

func (m ScreenModel) renderHeader() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true).
		Render(m.title)

	return headerStyle.Width(m.width).Render(title)
}

func (m ScreenModel) renderContent() string {
	// Compose your child views:
	// sidebar := m.list.View()
	// detail := m.detail.View()
	// return lipgloss.JoinHorizontal(lipgloss.Top, sidebar, " ", detail)

	return "Screen content goes here"
}

func (m ScreenModel) renderLoading() string {
	return lipgloss.Place(
		m.width,
		m.height-4,
		lipgloss.Center,
		lipgloss.Center,
		"Loading...",
	)
}

func (m ScreenModel) renderError() string {
	errorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF073A")).
		Bold(true)

	return lipgloss.Place(
		m.width,
		m.height-4,
		lipgloss.Center,
		lipgloss.Center,
		errorStyle.Render("Error: "+m.err.Error()),
	)
}

func (m ScreenModel) renderHelp() string {
	helpText := `
Keyboard Shortcuts:
  ↑/k        Move up
  ↓/j        Move down
  Enter      Select
  Tab        Switch pane
  Esc        Go back
  ?          Toggle help
  q          Quit
`
	return helpText
}

func (m ScreenModel) renderFooter() string {
	hints := []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Render("esc") +
			lipgloss.NewStyle().Foreground(lipgloss.Color("#6B6B6B")).Render(":back"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Render("?") +
			lipgloss.NewStyle().Foreground(lipgloss.Color("#6B6B6B")).Render(":help"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Render("q") +
			lipgloss.NewStyle().Foreground(lipgloss.Color("#6B6B6B")).Render(":quit"),
	}

	return footerStyle.Width(m.width).Render(
		lipgloss.JoinHorizontal(lipgloss.Left, hints[0], "  ", hints[1], "  ", hints[2]),
	)
}
