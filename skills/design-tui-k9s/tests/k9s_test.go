package tests

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/MateoSegura/.claude/skilltest"
)

// TestK9sTUIStyle tests the design-tui-k9s skill.
func TestK9sTUIStyle(t *testing.T) {
	if os.Getenv("SKILL_TEST") == "" {
		t.Skip("Set SKILL_TEST=1 to run skill tests (requires Claude CLI)")
	}

	runner := skilltest.NewTestRunner()
	runner.Verbose = testing.Verbose()

	suite := &skilltest.Suite{
		Name:  "design-tui-k9s",
		Skill: "design-tui-k9s",
		Cases: []*skilltest.TestCase{
			{
				Name:   "chrome-component",
				Skill:  "design-tui-k9s",
				Prompt: "Create a K9s-style chrome component with a header showing app name and version, and a footer with keyboard hints.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("header"),
					skilltest.ContainsText("footer"),
					skilltest.MatchesRegex(`lipgloss|NewStyle`),
					skilltest.CustomValidator("has-hints", func(output string) (bool, string) {
						hasHints := strings.Contains(output, "hint") || strings.Contains(output, "Hint") ||
							strings.Contains(output, "shortcut") || strings.Contains(output, "key")
						if hasHints {
							return true, "Keyboard hints pattern found"
						}
						return false, "No keyboard hints pattern found"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "list-screen",
				Skill:  "design-tui-k9s",
				Prompt: "Create a K9s-style list screen that displays items with selection highlighting. Include vim-style navigation (j/k for up/down).",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.MatchesRegex(`case\s+["']j["']|case\s+["']k["']`),
					skilltest.ContainsText("cursor"),
					skilltest.ContainsText("selected"),
					skilltest.CustomValidator("highlighting", func(output string) (bool, string) {
						hasHighlight := strings.Contains(output, "highlight") ||
							strings.Contains(output, "selected") ||
							strings.Contains(output, "Background")
						if hasHighlight {
							return true, "Selection highlighting found"
						}
						return false, "No selection highlighting pattern"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "form-screen",
				Skill:  "design-tui-k9s",
				Prompt: "Create a K9s-style form screen with text inputs and action buttons. Use the K9s dark theme colors.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("textinput"),
					skilltest.MatchesRegex(`button|submit|save|confirm`),
					skilltest.CustomValidator("dark-theme", func(output string) (bool, string) {
						hasDark := strings.Contains(output, "#") ||
							strings.Contains(output, "Background") ||
							strings.Contains(output, "Color(")
						if hasDark {
							return true, "Color styling found"
						}
						return false, "No color styling found"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "modal-dialog",
				Skill:  "design-tui-k9s",
				Prompt: "Create a K9s-style confirmation modal that overlays the current screen with yes/no options.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.MatchesRegex(`modal|dialog|overlay|confirm`),
					skilltest.ContainsText("yes"),
					skilltest.ContainsText("no"),
					skilltest.CustomValidator("centered", func(output string) (bool, string) {
						hasCentering := strings.Contains(output, "center") ||
							strings.Contains(output, "Place") ||
							strings.Contains(output, "width/2") ||
							strings.Contains(output, "height/2")
						if hasCentering {
							return true, "Modal centering logic found"
						}
						return false, "No modal centering logic"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "empty-state",
				Skill:  "design-tui-k9s",
				Prompt: "Create a K9s-style empty state view for when there's no data to display. Include an icon and helpful message.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.CustomValidator("icon-or-emoji", func(output string) (bool, string) {
						hasIcon := strings.Contains(output, "icon") ||
							strings.Contains(output, "Icon") ||
							strings.Contains(output, "📭") ||
							strings.Contains(output, "∅") ||
							strings.Contains(output, "●") ||
							strings.Contains(output, "○")
						if hasIcon {
							return true, "Icon or visual indicator found"
						}
						return false, "No icon or visual indicator"
					}),
					skilltest.MatchesRegex(`empty|no.*data|nothing`),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "color-palette",
				Skill:  "design-tui-k9s",
				Prompt: "Define a K9s-style color palette in Go using lipgloss. Include colors for primary, secondary, success, warning, error states.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("lipgloss"),
					skilltest.MatchesRegex(`Color\(|AdaptiveColor|#[0-9a-fA-F]{6}`),
					skilltest.CustomValidator("semantic-colors", func(output string) (bool, string) {
						colors := []string{"primary", "secondary", "success", "warning", "error"}
						found := 0
						for _, c := range colors {
							if strings.Contains(strings.ToLower(output), c) {
								found++
							}
						}
						if found >= 3 {
							return true, "Semantic color names found"
						}
						return false, "Missing semantic color names"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	result, err := runner.RunSuite(ctx, suite)
	if err != nil {
		t.Fatalf("Suite execution failed: %v", err)
	}

	if err := runner.SaveSuiteResults(result, "k9s-results.json"); err != nil {
		t.Logf("Warning: couldn't save results: %v", err)
	}

	t.Logf("Suite: %s", result.Name)
	t.Logf("Tests: %d total, %d passed, %d failed", result.TotalTests, result.Passed, result.Failed)
	t.Logf("Score: %.2f%% (Grade: %s)", result.Score*100, skilltest.DefaultGradeScale().Grade(result.Score))
	t.Logf("Duration: %v", result.Duration)

	for _, r := range result.Results {
		if !r.Passed {
			t.Logf("FAILED: %s (iteration %d) - Score: %.2f%%", r.Name, r.Iteration, r.Score*100)
			for _, v := range r.Validations {
				if !v.Passed {
					t.Logf("  - %s: %s", v.Name, v.Message)
				}
			}
		}
	}

	if result.Score < 0.70 {
		t.Errorf("Suite score %.2f%% is below 70%% threshold", result.Score*100)
	}
}
