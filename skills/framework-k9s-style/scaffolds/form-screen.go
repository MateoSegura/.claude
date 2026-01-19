// Package screens - K9s-Style Form Screen Scaffold
//
// USAGE: Copy this file and customize fields for your form.
// Modify fields slice and validation as needed.
package screens

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// =============================================================================
// THEME (inline for scaffold - use your theme package in real code)
// =============================================================================

var (
	formColorGold      = lipgloss.Color("#FFD700")
	formColorBlack     = lipgloss.Color("#0A0A0A")
	formColorBlackLight = lipgloss.Color("#1A1A1A")
	formColorWhite     = lipgloss.Color("#FAFAFA")
	formColorGray      = lipgloss.Color("#6B6B6B")
	formColorGrayDark  = lipgloss.Color("#3D3D3D")
	formColorCharcoal  = lipgloss.Color("#252525")
	formColorRed       = lipgloss.Color("#FF073A")
)

// =============================================================================
// STYLES
// =============================================================================

var (
	formInputStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(formColorCharcoal).
		Padding(0, 1)

	formInputFocusedStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(formColorGold).
		Padding(0, 1)

	formInputErrorStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(formColorRed).
		Padding(0, 1)
)

// =============================================================================
// MESSAGES
// =============================================================================

// FormSubmittedMsg is sent when the form is successfully submitted.
type FormSubmittedMsg struct {
	Values map[string]string
}

// FormCancelledMsg is sent when the form is cancelled.
type FormCancelledMsg struct{}

// =============================================================================
// KEY BINDINGS
// =============================================================================

type formKeyMap struct {
	Next   key.Binding
	Prev   key.Binding
	Submit key.Binding
	Cancel key.Binding
}

var formKeys = formKeyMap{
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "prev"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel"),
	),
}

// =============================================================================
// FIELD DEFINITION
// =============================================================================

// FormField defines a form field.
type FormField struct {
	Key         string // Unique key for this field
	Label       string // Display label
	Placeholder string
	Required    bool
	Width       int
	input       textinput.Model
	error       string
}

// =============================================================================
// MODEL
// =============================================================================

// FormScreenModel is a K9s-style form screen.
type FormScreenModel struct {
	title      string
	fields     []FormField
	focusIndex int
	width      int
	height     int
	helperText string
}

// NewFormScreen creates a new form screen.
func NewFormScreen(title string, fields []FormField) FormScreenModel {
	// Initialize text inputs for each field
	for i := range fields {
		ti := textinput.New()
		ti.Placeholder = fields[i].Placeholder
		ti.Prompt = ""
		ti.Width = fields[i].Width
		if fields[i].Width == 0 {
			ti.Width = 40
		}
		ti.CharLimit = 256
		ti.TextStyle = lipgloss.NewStyle().Foreground(formColorWhite)
		ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(formColorGrayDark)
		ti.Cursor.Style = lipgloss.NewStyle().Foreground(formColorGold)

		if i == 0 {
			ti.Focus()
		}

		fields[i].input = ti
	}

	return FormScreenModel{
		title:      title,
		fields:     fields,
		focusIndex: 0,
	}
}

// SetHelperText sets the helper text shown below the form.
func (m *FormScreenModel) SetHelperText(text string) {
	m.helperText = text
}

// SetSize updates dimensions.
func (m *FormScreenModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// GetValues returns a map of field keys to values.
func (m FormScreenModel) GetValues() map[string]string {
	values := make(map[string]string)
	for _, f := range m.fields {
		values[f.Key] = strings.TrimSpace(f.input.Value())
	}
	return values
}

// =============================================================================
// TEA.MODEL IMPLEMENTATION
// =============================================================================

func (m FormScreenModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FormScreenModel) Update(msg tea.Msg) (FormScreenModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, formKeys.Next):
			m.focusIndex = (m.focusIndex + 1) % len(m.fields)
			m.updateFocus()
			return m, textinput.Blink

		case key.Matches(msg, formKeys.Prev):
			m.focusIndex = (m.focusIndex - 1 + len(m.fields)) % len(m.fields)
			m.updateFocus()
			return m, textinput.Blink

		case key.Matches(msg, formKeys.Submit):
			if m.validate() {
				return m, func() tea.Msg {
					return FormSubmittedMsg{Values: m.GetValues()}
				}
			}
			return m, nil

		case key.Matches(msg, formKeys.Cancel):
			return m, func() tea.Msg {
				return FormCancelledMsg{}
			}
		}
	}

	// Forward to focused input
	var cmd tea.Cmd
	m.fields[m.focusIndex].input, cmd = m.fields[m.focusIndex].input.Update(msg)
	return m, cmd
}

func (m *FormScreenModel) updateFocus() {
	for i := range m.fields {
		if i == m.focusIndex {
			m.fields[i].input.Focus()
		} else {
			m.fields[i].input.Blur()
		}
	}
}

func (m *FormScreenModel) validate() bool {
	valid := true
	for i := range m.fields {
		m.fields[i].error = ""
		value := strings.TrimSpace(m.fields[i].input.Value())
		if m.fields[i].Required && value == "" {
			m.fields[i].error = "This field is required"
			valid = false
		}
	}
	return valid
}

func (m FormScreenModel) View() string {
	// Header
	header := m.renderHeader()

	// Form content
	form := m.renderForm()

	// Footer
	footer := m.renderFooter()

	// Compose
	contentHeight := m.height - 2
	centered := lipgloss.Place(
		m.width,
		contentHeight,
		lipgloss.Center,
		lipgloss.Center,
		form,
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, centered, footer)
}

// =============================================================================
// RENDER HELPERS
// =============================================================================

func (m FormScreenModel) renderHeader() string {
	left := lipgloss.NewStyle().
		Foreground(formColorGold).
		Bold(true).
		Render("◆ " + m.title)

	return lipgloss.NewStyle().
		Background(formColorBlackLight).
		Width(m.width).
		Padding(0, 1).
		Render(left)
}

func (m FormScreenModel) renderFooter() string {
	shortcuts := []string{
		renderFormShortcut("Tab", "Next"),
		renderFormShortcut("Enter", "Submit"),
		renderFormShortcut("Esc", "Cancel"),
	}

	return lipgloss.NewStyle().
		Background(formColorBlackLight).
		Width(m.width).
		Padding(0, 1).
		Render(strings.Join(shortcuts, "  "))
}

func renderFormShortcut(key, desc string) string {
	k := lipgloss.NewStyle().Foreground(formColorGold).Bold(true).Render("<" + key + ">")
	d := lipgloss.NewStyle().Foreground(formColorGray).Render(desc)
	return k + d
}

func (m FormScreenModel) renderForm() string {
	// Title
	title := lipgloss.NewStyle().
		Foreground(formColorGold).
		Bold(true).
		Render(m.title)

	// Fields
	var fieldViews []string
	labelStyle := lipgloss.NewStyle().
		Foreground(formColorGray).
		Width(12).
		Align(lipgloss.Right)

	for i, f := range m.fields {
		label := labelStyle.Render(f.Label)

		// Determine input style
		inputStyle := formInputStyle
		if i == m.focusIndex {
			inputStyle = formInputFocusedStyle
		}
		if f.error != "" {
			inputStyle = formInputErrorStyle
		}

		inputView := inputStyle.Width(f.Width + 4).Render(f.input.View())

		row := lipgloss.JoinHorizontal(lipgloss.Left, label, "  ", inputView)

		// Add error message if present
		if f.error != "" {
			errStyle := lipgloss.NewStyle().
				Foreground(formColorRed).
				MarginLeft(14)
			row = lipgloss.JoinVertical(lipgloss.Left, row, errStyle.Render(f.error))
		}

		fieldViews = append(fieldViews, row, "")
	}

	// Helper text
	var helperView string
	if m.helperText != "" {
		helperView = lipgloss.NewStyle().
			Foreground(formColorGrayDark).
			Render(m.helperText)
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		lipgloss.JoinVertical(lipgloss.Left, fieldViews...),
		helperView,
	)
}
