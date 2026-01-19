// Package components - K9s-Style Chrome Component
//
// USAGE: Copy this file to your project's components package.
// Chrome wraps every screen with consistent header and footer bars.
package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// =============================================================================
// THEME COLORS (copy from your theme package or inline)
// =============================================================================

var (
	colorGold      = lipgloss.Color("#FFD700")
	colorBlack     = lipgloss.Color("#0A0A0A")
	colorBlackLight = lipgloss.Color("#1A1A1A")
	colorGray      = lipgloss.Color("#6B6B6B")
	colorGrayDark  = lipgloss.Color("#3D3D3D")
	colorCyan      = lipgloss.Color("#00D9FF")
)

// =============================================================================
// MINIMUM SIZE
// =============================================================================

const (
	MinWidth  = 80
	MinHeight = 24
)

// =============================================================================
// SHORTCUT TYPE
// =============================================================================

// Shortcut represents a keyboard shortcut displayed in the footer.
type Shortcut struct {
	Key  string // The key(s), e.g., "↑↓", "Enter", "q"
	Desc string // Description, e.g., "Navigate", "Select", "Quit"
}

// Common shortcut sets
var (
	ShortcutsNavigation = []Shortcut{
		{Key: "↑↓", Desc: "Navigate"},
		{Key: "Enter", Desc: "Select"},
		{Key: "q", Desc: "Quit"},
	}

	ShortcutsNavigationWithBack = []Shortcut{
		{Key: "↑↓", Desc: "Navigate"},
		{Key: "Enter", Desc: "Select"},
		{Key: "Esc", Desc: "Back"},
		{Key: "q", Desc: "Quit"},
	}

	ShortcutsForm = []Shortcut{
		{Key: "Tab", Desc: "Next"},
		{Key: "Enter", Desc: "Submit"},
		{Key: "Esc", Desc: "Cancel"},
	}

	ShortcutsList = []Shortcut{
		{Key: "↑↓", Desc: "Navigate"},
		{Key: "Enter", Desc: "Select"},
		{Key: "n", Desc: "New"},
		{Key: "d", Desc: "Delete"},
		{Key: "q", Desc: "Quit"},
	}

	ShortcutsEmpty = []Shortcut{
		{Key: "n", Desc: "New"},
		{Key: "q", Desc: "Quit"},
	}
)

// =============================================================================
// CHROME CONFIG
// =============================================================================

// ChromeConfig configures the chrome appearance.
type ChromeConfig struct {
	// Header
	Title      string // Left side, e.g., "◆ MYAPP"
	Context    string // Right side, e.g., "Items [42]"
	ShowHeader bool

	// Footer
	Shortcuts      []Shortcut // Context-specific shortcuts (left side)
	PersistentKey  string     // Always-shown shortcut key (right side)
	PersistentDesc string     // Always-shown shortcut description

	// Dimensions
	Width  int
	Height int
}

// =============================================================================
// CHROME
// =============================================================================

// Chrome renders K9s-style header and footer around content.
type Chrome struct {
	config ChromeConfig
}

// NewChrome creates a new chrome renderer.
func NewChrome(config ChromeConfig) Chrome {
	return Chrome{config: config}
}

// DefaultChrome creates chrome with default app branding.
// Customize the title to your app name.
func DefaultChrome(width, height int) Chrome {
	return NewChrome(ChromeConfig{
		Title:          "◆ MYAPP", // Change this to your app
		ShowHeader:     true,
		PersistentKey:  "Ctrl+S",
		PersistentDesc: "Settings",
		Width:          width,
		Height:         height,
	})
}

// === Setters (return new Chrome for chaining) ===

// SetSize updates the dimensions.
func (c Chrome) SetSize(width, height int) Chrome {
	c.config.Width = width
	c.config.Height = height
	return c
}

// SetTitle updates the header title.
func (c Chrome) SetTitle(title string) Chrome {
	c.config.Title = title
	return c
}

// SetContext updates the header context string.
func (c Chrome) SetContext(context string) Chrome {
	c.config.Context = context
	return c
}

// SetShortcuts updates the footer shortcuts.
func (c Chrome) SetShortcuts(shortcuts []Shortcut) Chrome {
	c.config.Shortcuts = shortcuts
	return c
}

// SetPersistent sets the persistent shortcut on the right side of footer.
func (c Chrome) SetPersistent(key, desc string) Chrome {
	c.config.PersistentKey = key
	c.config.PersistentDesc = desc
	return c
}

// === Getters ===

// IsTooSmall returns true if terminal is below minimum size.
func (c Chrome) IsTooSmall() bool {
	return c.config.Width < MinWidth || c.config.Height < MinHeight
}

// ContentHeight returns available height for content.
func (c Chrome) ContentHeight() int {
	h := c.config.Height
	if c.config.ShowHeader {
		h-- // header takes 1 line
	}
	h-- // footer takes 1 line
	return max(h, 1)
}

// ContentWidth returns available width for content.
func (c Chrome) ContentWidth() int {
	return c.config.Width
}

// === Rendering ===

// Render wraps content with header and footer.
func (c Chrome) Render(content string) string {
	if c.IsTooSmall() {
		return c.renderTooSmall()
	}

	var parts []string

	if c.config.ShowHeader {
		parts = append(parts, c.renderHeader())
	}

	contentStyle := lipgloss.NewStyle().
		Width(c.config.Width).
		Height(c.ContentHeight())
	parts = append(parts, contentStyle.Render(content))

	parts = append(parts, c.renderFooter())

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// RenderCentered wraps centered content with header and footer.
func (c Chrome) RenderCentered(content string) string {
	if c.IsTooSmall() {
		return c.renderTooSmall()
	}

	var parts []string

	if c.config.ShowHeader {
		parts = append(parts, c.renderHeader())
	}

	centered := lipgloss.Place(
		c.config.Width,
		c.ContentHeight(),
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
	parts = append(parts, centered)

	parts = append(parts, c.renderFooter())

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// === Internal Renderers ===

func (c Chrome) renderTooSmall() string {
	msg := lipgloss.NewStyle().
		Foreground(colorGold).
		Bold(true).
		Render("Terminal too small")

	hint := lipgloss.NewStyle().
		Foreground(colorGray).
		Render("Minimum: 80x24")

	content := lipgloss.JoinVertical(lipgloss.Center, msg, hint)

	return lipgloss.Place(
		c.config.Width,
		c.config.Height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (c Chrome) renderHeader() string {
	// Left: Title
	left := lipgloss.NewStyle().
		Foreground(colorGold).
		Bold(true).
		Render(c.config.Title)

	// Right: Context
	right := lipgloss.NewStyle().
		Foreground(colorGrayDark).
		Render(c.config.Context)

	// Calculate gap
	gap := c.config.Width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}

	// Header bar style
	headerStyle := lipgloss.NewStyle().
		Background(colorBlackLight).
		Width(c.config.Width).
		Padding(0, 1)

	return headerStyle.Render(left + strings.Repeat(" ", gap) + right)
}

func (c Chrome) renderFooter() string {
	// Build left side: shortcuts
	var shortcutParts []string
	for _, s := range c.config.Shortcuts {
		key := lipgloss.NewStyle().
			Foreground(colorGold).
			Bold(true).
			Render("<" + s.Key + ">")
		desc := lipgloss.NewStyle().
			Foreground(colorGray).
			Render(s.Desc)
		shortcutParts = append(shortcutParts, key+desc)
	}
	left := strings.Join(shortcutParts, "  ")

	// Build right side: persistent shortcut
	var right string
	if c.config.PersistentKey != "" {
		key := lipgloss.NewStyle().
			Foreground(colorCyan).
			Bold(true).
			Render("<" + c.config.PersistentKey + ">")
		desc := lipgloss.NewStyle().
			Foreground(colorGrayDark).
			Render(c.config.PersistentDesc)
		right = key + desc
	}

	// Calculate gap
	gap := c.config.Width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}

	// Footer bar style
	footerStyle := lipgloss.NewStyle().
		Background(colorBlackLight).
		Width(c.config.Width).
		Padding(0, 1)

	return footerStyle.Render(left + strings.Repeat(" ", gap) + right)
}

// max returns the larger of two ints.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
