// Package theme - Lipgloss Theme System Scaffold
//
// USAGE: Copy this file to your project and customize colors/styles.
// Import as: import "yourproject/internal/tui/theme"
// Use as: theme.Title.Render("Hello")
package theme

import "github.com/charmbracelet/lipgloss"

// =============================================================================
// COLOR PALETTE
// =============================================================================

// Primary colors - your brand identity
var (
	ColorPrimary    = lipgloss.Color("#FFD700") // Gold - primary accent
	ColorSecondary  = lipgloss.Color("#00D9FF") // Cyan - secondary accent
	ColorTertiary   = lipgloss.Color("#BF00FF") // Purple - tertiary accent
)

// Semantic colors - convey meaning
var (
	ColorSuccess = lipgloss.Color("#39FF14") // Neon lime
	ColorWarning = lipgloss.Color("#FF6B00") // Neon orange
	ColorError   = lipgloss.Color("#FF073A") // Neon red
	ColorInfo    = lipgloss.Color("#00D9FF") // Cyan
)

// Background colors - layered surfaces
var (
	ColorBgBase     = lipgloss.Color("#0A0A0A") // Darkest - main background
	ColorBgSurface  = lipgloss.Color("#121212") // Elevated surfaces
	ColorBgOverlay  = lipgloss.Color("#1A1A1A") // Cards, panels
	ColorBgBorder   = lipgloss.Color("#252525") // Borders, dividers
)

// Text colors - hierarchy
var (
	ColorTextPrimary   = lipgloss.Color("#FAFAFA") // Primary text
	ColorTextSecondary = lipgloss.Color("#C0C0C0") // Secondary text
	ColorTextMuted     = lipgloss.Color("#6B6B6B") // Muted/disabled
	ColorTextDisabled  = lipgloss.Color("#3D3D3D") // Fully disabled
)

// =============================================================================
// SEMANTIC ALIASES (easier to remember)
// =============================================================================

var (
	// Quick access
	Gold    = ColorPrimary
	Cyan    = ColorSecondary
	Purple  = ColorTertiary
	Lime    = ColorSuccess
	Orange  = ColorWarning
	Red     = ColorError
	White   = ColorTextPrimary
	Gray    = ColorTextMuted
	Black   = ColorBgBase
)

// =============================================================================
// TYPOGRAPHY STYLES
// =============================================================================

var (
	// Title - large, bold headings
	Title = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)

	// Subtitle - supporting text under titles
	Subtitle = lipgloss.NewStyle().
			Foreground(ColorTextSecondary)

	// Heading - section headers
	Heading = lipgloss.NewStyle().
		Foreground(ColorTextPrimary).
		Bold(true)

	// Body - normal text
	Body = lipgloss.NewStyle().
		Foreground(ColorTextPrimary)

	// Muted - de-emphasized text
	Muted = lipgloss.NewStyle().
		Foreground(ColorTextMuted)

	// Code - monospace/technical
	Code = lipgloss.NewStyle().
		Foreground(ColorSuccess).
		Background(ColorBgOverlay).
		Padding(0, 1)
)

// =============================================================================
// SEMANTIC STYLES
// =============================================================================

var (
	Success = lipgloss.NewStyle().Foreground(ColorSuccess)
	Warning = lipgloss.NewStyle().Foreground(ColorWarning)
	Error   = lipgloss.NewStyle().Foreground(ColorError)
	Info    = lipgloss.NewStyle().Foreground(ColorInfo)

	SuccessBold = Success.Bold(true)
	WarningBold = Warning.Bold(true)
	ErrorBold   = Error.Bold(true)
	InfoBold    = Info.Bold(true)
)

// =============================================================================
// CONTAINER STYLES
// =============================================================================

var (
	// Card - elevated container with border
	Card = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorBgBorder).
		Padding(1, 2)

	// CardFocused - card with primary color border
	CardFocused = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 2)

	// Panel - larger container
	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorBgBorder).
		Padding(1, 3)

	// PanelFocused - panel with accent border
	PanelFocused = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 3)
)

// =============================================================================
// INTERACTIVE STYLES
// =============================================================================

var (
	// Button - primary action
	Button = lipgloss.NewStyle().
		Foreground(ColorBgBase).
		Background(ColorPrimary).
		Padding(0, 2).
		Bold(true)

	// ButtonSecondary - secondary action
	ButtonSecondary = lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Background(ColorBgBorder).
			Padding(0, 2)

	// Selected - highlighted/selected item
	Selected = lipgloss.NewStyle().
			Foreground(ColorBgBase).
			Background(ColorPrimary).
			Bold(true).
			Padding(0, 1)

	// Cursor - current cursor position
	Cursor = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)
)

// =============================================================================
// INPUT STYLES
// =============================================================================

var (
	// Input - text input field
	Input = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(ColorBgBorder).
		Padding(0, 1)

	// InputFocused - focused text input
	InputFocused = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(ColorPrimary).
			Padding(0, 1)

	// Placeholder - placeholder text
	Placeholder = lipgloss.NewStyle().
			Foreground(ColorTextDisabled)
)

// =============================================================================
// STATUS BAR STYLES
// =============================================================================

var (
	// StatusBar - bottom status bar
	StatusBar = lipgloss.NewStyle().
			Background(ColorBgSurface).
			Foreground(ColorTextMuted).
			Padding(0, 1)

	// HintKey - keyboard shortcut key
	HintKey = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)

	// HintAction - keyboard shortcut action
	HintAction = lipgloss.NewStyle().
			Foreground(ColorTextMuted)
)

// =============================================================================
// ICONS (Unicode)
// =============================================================================

var (
	IconCheck      = "✓"
	IconCross      = "✗"
	IconArrowRight = "→"
	IconArrowLeft  = "←"
	IconArrowUp    = "↑"
	IconArrowDown  = "↓"
	IconDot        = "•"
	IconCircle     = "○"
	IconCircleFill = "●"
	IconTriangleR  = "▸"
	IconTriangleD  = "▾"
	IconDiamond    = "◆"
	IconStar       = "★"
	IconSpinner    = "⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏"
)

// Spinner frames for animation
var SpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
var DotsFrames = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

// Center returns a style that centers content in the given width.
func Center(width int) lipgloss.Style {
	return lipgloss.NewStyle().Width(width).Align(lipgloss.Center)
}

// FullWidth returns a style that fills the given width.
func FullWidth(width int) lipgloss.Style {
	return lipgloss.NewStyle().Width(width)
}

// StateColor returns the appropriate color for a state string.
func StateColor(state string) lipgloss.Color {
	switch state {
	case "success", "completed", "done":
		return ColorSuccess
	case "warning", "pending", "waiting":
		return ColorWarning
	case "error", "failed":
		return ColorError
	case "info", "running", "active":
		return ColorInfo
	default:
		return ColorTextMuted
	}
}

// StateStyle returns a style for the given state.
func StateStyle(state string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(StateColor(state))
}
