// Package main - Basic Bubble Tea Application Scaffold
//
// USAGE: Copy this file and modify for your needs.
// This provides the minimal working structure for a Bubble Tea app.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// === STYLES ===

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFD700"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000")).
			Background(lipgloss.Color("#FFD700")).
			Padding(0, 1)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA"))

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B6B6B"))
)

// === MODEL ===

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	width    int
	height   int
}

func initialModel() model {
	return model{
		choices:  []string{"Option 1", "Option 2", "Option 3"},
		selected: make(map[int]struct{}),
	}
}

// === MESSAGES ===

// Add custom message types here:
// type dataLoadedMsg struct { data []string }
// type errMsg struct { err error }

// === INIT ===

func (m model) Init() tea.Cmd {
	// Return initial commands here
	// Example: return tea.Batch(loadDataCmd, startTimerCmd)
	return nil
}

// === UPDATE ===

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

// === VIEW ===

func (m model) View() string {
	s := titleStyle.Render("Select Options") + "\n\n"

	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		line := fmt.Sprintf("%s[%s] %s", cursor, checked, choice)

		if m.cursor == i {
			s += selectedStyle.Render(line) + "\n"
		} else {
			s += normalStyle.Render(line) + "\n"
		}
	}

	s += "\n" + mutedStyle.Render("j/k: move  space: select  q: quit")

	return s
}

// === MAIN ===

func main() {
	p := tea.NewProgram(
		initialModel(),
		// Uncomment for full-screen apps:
		// tea.WithAltScreen(),
		// tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
