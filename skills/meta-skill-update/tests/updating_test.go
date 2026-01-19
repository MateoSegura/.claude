package tests

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/MateoSegura/.claude/skilltest"
)

// TestUpdatingClaudeExtension tests the meta-skill-update skill.
func TestUpdatingClaudeExtension(t *testing.T) {
	if os.Getenv("SKILL_TEST") == "" {
		t.Skip("Set SKILL_TEST=1 to run skill tests (requires Claude CLI)")
	}

	runner := skilltest.NewTestRunner()
	runner.Verbose = testing.Verbose()

	suite := &skilltest.Suite{
		Name:  "meta-skill-update",
		Skill: "meta-skill-update",
		Cases: []*skilltest.TestCase{
			{
				Name:   "discovery-skills",
				Skill:  "meta-skill-update",
				Prompt: "List all skills in this project. What skills are available?",
				Validators: []skilltest.Validator{
					skilltest.CustomValidator("finds-skills", func(output string) (bool, string) {
						output = strings.ToLower(output)
						found := 0
						skills := []string{"bubbletea", "k9s", "extension", "updating"}
						for _, s := range skills {
							if strings.Contains(output, s) {
								found++
							}
						}
						if found >= 2 {
							return true, "Found multiple skills"
						}
						return false, "Did not find skills"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "update-workflow",
				Skill:  "meta-skill-update",
				Prompt: "How would I update the bubbletea-tui skill to support a new version of the bubbletea library?",
				Validators: []skilltest.Validator{
					skilltest.CustomValidator("mentions-steps", func(output string) (bool, string) {
						output = strings.ToLower(output)
						hasSteps := (strings.Contains(output, "1.") || strings.Contains(output, "first")) &&
							(strings.Contains(output, "2.") || strings.Contains(output, "then") || strings.Contains(output, "next"))
						if hasSteps {
							return true, "Update workflow steps found"
						}
						return false, "No clear update workflow"
					}),
					skilltest.CustomValidator("mentions-version", func(output string) (bool, string) {
						output = strings.ToLower(output)
						hasVersion := strings.Contains(output, "version") ||
							strings.Contains(output, "changelog") ||
							strings.Contains(output, "breaking change")
						if hasVersion {
							return true, "Version consideration mentioned"
						}
						return false, "No version consideration"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "hook-discovery",
				Skill:  "meta-skill-update",
				Prompt: "How do I find and update hooks in this project?",
				Validators: []skilltest.Validator{
					skilltest.ContainsText("settings.json"),
					skilltest.MatchesRegex(`\.claude|hooks`),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	result, err := runner.RunSuite(ctx, suite)
	if err != nil {
		t.Fatalf("Suite execution failed: %v", err)
	}

	if err := runner.SaveSuiteResults(result, "updating-results.json"); err != nil {
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
