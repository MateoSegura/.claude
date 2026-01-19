package tests

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/MateoSegura/.claude/skilltest"
)

// TestBubbleTeaTUI tests the bubbletea-tui skill.
func TestBubbleTeaTUI(t *testing.T) {
	if os.Getenv("SKILL_TEST") == "" {
		t.Skip("Set SKILL_TEST=1 to run skill tests (requires Claude CLI)")
	}

	runner := skilltest.NewTestRunner()
	runner.Verbose = testing.Verbose()

	suite := &skilltest.Suite{
		Name:  "bubbletea-tui",
		Skill: "bubbletea-tui",
		Cases: []*skilltest.TestCase{
			{
				Name:   "basic-model-creation",
				Skill:  "bubbletea-tui",
				Prompt: "Create a simple Bubble Tea model for a counter that can increment and decrement. Just show me the Go code.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("tea.Model"),
					skilltest.ContainsText("Init()"),
					skilltest.ContainsText("Update("),
					skilltest.ContainsText("View()"),
					skilltest.MatchesRegex(`func.*Update.*tea\.Msg`),
					skilltest.NoErrors(),
					skilltest.CustomValidator("immutable-update", checkImmutableUpdate),
				},
				Iterations: 3,
			},
			{
				Name:   "spinner-component",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea app that shows a spinner while loading data. Use the bubbles spinner component.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("spinner"),
					skilltest.ContainsText("spinner.Model"),
					skilltest.MatchesRegex(`spinner\.New\(\)`),
					skilltest.ContainsText("spinner.Tick"),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "keyboard-handling",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea model that handles keyboard input: q to quit, enter to confirm, escape to cancel.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("tea.KeyMsg"),
					skilltest.MatchesRegex(`case\s+"q".*tea\.Quit`),
					skilltest.ContainsText("key.Type"),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "lipgloss-styling",
				Skill:  "bubbletea-tui",
				Prompt: "Create a styled Bubble Tea component using lipgloss with a border, padding, and colors.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("lipgloss"),
					skilltest.MatchesRegex(`lipgloss\.NewStyle\(\)`),
					skilltest.MatchesRegex(`Border|Padding|Foreground|Background`),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "command-pattern",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea model that fetches data asynchronously using tea.Cmd.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("tea.Cmd"),
					skilltest.MatchesRegex(`func\s+\w+\(\)\s+tea\.Cmd`),
					skilltest.NoErrors(),
					skilltest.CustomValidator("returns-cmd", checkReturnsCmd),
				},
				Iterations: 2,
			},
			{
				Name:   "list-component",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea app with a selectable list using the bubbles list component. Show items and handle selection.",
				Validators: []skilltest.Validator{
					skilltest.ContainsCode("go"),
					skilltest.ContainsText("list.Model"),
					skilltest.ContainsText("list.New"),
					skilltest.ContainsText("list.Item"),
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

	// Save results
	if err := runner.SaveSuiteResults(result, "bubbletea-results.json"); err != nil {
		t.Logf("Warning: couldn't save results: %v", err)
	}

	// Report
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

func checkImmutableUpdate(output string) (bool, string) {
	hasReturn := strings.Contains(output, "return m,") || strings.Contains(output, "return m ,")
	hasDirectMutation := strings.Contains(output, "m.") && !strings.Contains(output, "return")

	if hasReturn && !hasDirectMutation {
		return true, "Update pattern looks immutable"
	}
	return false, "Update pattern may have mutation issues"
}

func checkReturnsCmd(output string) (bool, string) {
	if strings.Contains(output, "tea.Cmd") && strings.Contains(output, "return func()") {
		return true, "Proper tea.Cmd return pattern found"
	}
	if strings.Contains(output, "tea.Cmd") {
		return true, "tea.Cmd usage found"
	}
	return false, "No clear tea.Cmd return pattern"
}
