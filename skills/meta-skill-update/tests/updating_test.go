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
				Name:   "discovery-extensions",
				Skill:  "meta-skill-update",
				Prompt: "List all extensions in this knowledge base. What's available?",
				Validators: []skilltest.Validator{
					skilltest.CustomValidator("finds-meta-skills", func(output string) (bool, string) {
						output = strings.ToLower(output)
						hasCreate := strings.Contains(output, "meta-skill-create") || strings.Contains(output, "create")
						hasUpdate := strings.Contains(output, "meta-skill-update") || strings.Contains(output, "update")
						if hasCreate && hasUpdate {
							return true, "Found meta skills"
						}
						return false, "Did not find meta skills"
					}),
					skilltest.NoErrors(),
				},
				Iterations: 2,
			},
			{
				Name:   "update-workflow",
				Skill:  "meta-skill-update",
				Prompt: "How would I update a skill to add new patterns?",
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
					skilltest.CustomValidator("mentions-research", func(output string) (bool, string) {
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
					skilltest.MatchesRegex(`hooks`),
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
