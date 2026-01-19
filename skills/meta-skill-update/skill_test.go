package meta_skill_update_test

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/MateoSegura/.claude/tests"
)

// TestUpdatingClaudeExtension tests the meta-skill-update skill.
func TestUpdatingClaudeExtension(t *testing.T) {
	if os.Getenv("SKILL_TEST") == "" {
		t.Skip("Set SKILL_TEST=1 to run skill tests (requires Claude CLI)")
	}

	runner := tests.NewTestRunner()
	runner.Verbose = testing.Verbose()

	suite := &tests.Suite{
		Name:  "meta-skill-update",
		Skill: "meta-skill-update",
		Cases: []*tests.TestCase{
			{
				Name:   "discovery-extensions",
				Skill:  "meta-skill-update",
				Prompt: "List all extensions in this knowledge base. What's available?",
				Validators: []tests.Validator{
					tests.CustomValidator("finds-meta-skills", func(output string) (bool, string) {
						output = strings.ToLower(output)
						hasCreate := strings.Contains(output, "meta-skill-create") || strings.Contains(output, "create")
						hasUpdate := strings.Contains(output, "meta-skill-update") || strings.Contains(output, "update")
						if hasCreate && hasUpdate {
							return true, "Found meta skills"
						}
						return false, "Did not find meta skills"
					}),
					tests.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "update-workflow",
				Skill:  "meta-skill-update",
				Prompt: "How would I update a skill to add new patterns?",
				Validators: []tests.Validator{
					tests.CustomValidator("mentions-steps", func(output string) (bool, string) {
						output = strings.ToLower(output)
						hasSteps := (strings.Contains(output, "1.") || strings.Contains(output, "first")) &&
							(strings.Contains(output, "2.") || strings.Contains(output, "then") || strings.Contains(output, "next"))
						if hasSteps {
							return true, "Update workflow steps found"
						}
						return false, "No clear update workflow"
					}),
					tests.CustomValidator("mentions-research", func(output string) (bool, string) {
						output = strings.ToLower(output)
						hasResearch := strings.Contains(output, "search") ||
							strings.Contains(output, "research") ||
							strings.Contains(output, "web") ||
							strings.Contains(output, "latest")
						if hasResearch {
							return true, "Research step mentioned"
						}
						return false, "No research step"
					}),
					tests.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "hook-discovery",
				Skill:  "meta-skill-update",
				Prompt: "How do I find and update hooks in this project?",
				Validators: []tests.Validator{
					tests.ContainsText("settings.json"),
					tests.MatchesRegex(`hooks`),
					tests.NoErrors(),
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
	t.Logf("Score: %.2f%% (Grade: %s)", result.Score*100, tests.DefaultGradeScale().Grade(result.Score))
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
