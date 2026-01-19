package skilltests

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestBubbleTeaTUI tests the bubbletea-tui skill.
func TestBubbleTeaTUI(t *testing.T) {
	if os.Getenv("SKILL_TEST") == "" {
		t.Skip("Set SKILL_TEST=1 to run skill tests (requires Claude CLI)")
	}

	runner := NewTestRunner()
	runner.WorkDir = findProjectRoot()
	runner.Verbose = testing.Verbose()

	suite := &Suite{
		Name:  "bubbletea-tui",
		Skill: "bubbletea-tui",
		Cases: []*TestCase{
			{
				Name:   "basic-model-creation",
				Skill:  "bubbletea-tui",
				Prompt: "Create a simple Bubble Tea model for a counter that can increment and decrement. Just show me the Go code.",
				Validators: []Validator{
					ContainsCode("go"),
					ContainsText("tea.Model"),
					ContainsText("Init()"),
					ContainsText("Update("),
					ContainsText("View()"),
					MatchesRegex(`func.*Update.*tea\.Msg`),
					NoErrors(),
					CustomValidator("immutable-update", checkImmutableUpdate),
				},
				Iterations: 3, // Run multiple times for consistency
			},
			{
				Name:   "spinner-component",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea app that shows a spinner while loading data. Use the bubbles spinner component.",
				Validators: []Validator{
					ContainsCode("go"),
					ContainsText("spinner"),
					ContainsText("spinner.Model"),
					MatchesRegex(`spinner\.New\(\)`),
					ContainsText("spinner.Tick"),
					NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "keyboard-handling",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea model that handles keyboard input: q to quit, enter to confirm, escape to cancel.",
				Validators: []Validator{
					ContainsCode("go"),
					ContainsText("tea.KeyMsg"),
					MatchesRegex(`case\s+"q".*tea\.Quit`),
					ContainsText("key.Type"),
					NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "lipgloss-styling",
				Skill:  "bubbletea-tui",
				Prompt: "Create a styled Bubble Tea component using lipgloss with a border, padding, and colors.",
				Validators: []Validator{
					ContainsCode("go"),
					ContainsText("lipgloss"),
					MatchesRegex(`lipgloss\.NewStyle\(\)`),
					MatchesRegex(`Border|Padding|Foreground|Background`),
					NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "command-pattern",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea model that fetches data asynchronously using tea.Cmd.",
				Validators: []Validator{
					ContainsCode("go"),
					ContainsText("tea.Cmd"),
					MatchesRegex(`func\s+\w+\(\)\s+tea\.Cmd`),
					NoErrors(),
					CustomValidator("returns-cmd", checkReturnsCmd),
				},
				Iterations: 2,
			},
			{
				Name:   "list-component",
				Skill:  "bubbletea-tui",
				Prompt: "Create a Bubble Tea app with a selectable list using the bubbles list component. Show items and handle selection.",
				Validators: []Validator{
					ContainsCode("go"),
					ContainsText("list.Model"),
					ContainsText("list.New"),
					ContainsText("list.Item"),
					NoErrors(),
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
	t.Logf("Score: %.2f%% (Grade: %s)", result.Score*100, DefaultGradeScale().Grade(result.Score))
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

	// Fail if below threshold
	if result.Score < 0.70 {
		t.Errorf("Suite score %.2f%% is below 70%% threshold", result.Score*100)
	}
}

// checkImmutableUpdate validates that Update returns new state, not mutates.
func checkImmutableUpdate(output string) (bool, string) {
	// Look for return m, cmd pattern (good) vs m.field = value without return (bad)
	hasReturn := strings.Contains(output, "return m,") || strings.Contains(output, "return m ,")
	hasDirectMutation := strings.Contains(output, "m.") && !strings.Contains(output, "return")

	if hasReturn && !hasDirectMutation {
		return true, "Update pattern looks immutable"
	}
	return false, "Update pattern may have mutation issues"
}

// checkReturnsCmd validates that async functions return tea.Cmd.
func checkReturnsCmd(output string) (bool, string) {
	// Look for function returning tea.Cmd
	if strings.Contains(output, "tea.Cmd") && strings.Contains(output, "return func()") {
		return true, "Proper tea.Cmd return pattern found"
	}
	if strings.Contains(output, "tea.Cmd") {
		return true, "tea.Cmd usage found"
	}
	return false, "No clear tea.Cmd return pattern"
}

// findProjectRoot finds the project root by looking for .claude/skills directory.
func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		// Look for .claude/skills directory - that's where skills are
		skillsDir := filepath.Join(dir, ".claude", "skills")
		if info, err := os.Stat(skillsDir); err == nil && info.IsDir() {
			return dir
		}
		// Also check if we're inside .claude already
		if filepath.Base(dir) == ".claude" {
			parent := filepath.Dir(dir)
			skillsDir = filepath.Join(parent, ".claude", "skills")
			if info, err := os.Stat(skillsDir); err == nil && info.IsDir() {
				return parent
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	// Fallback - try relative path from skill-tests
	cwd, _ := os.Getwd()
	if strings.Contains(cwd, "skill-tests") {
		// We're in .claude/skill-tests, go up to project root
		parts := strings.Split(cwd, ".claude")
		if len(parts) > 0 {
			return strings.TrimSuffix(parts[0], string(filepath.Separator))
		}
	}
	return cwd
}
