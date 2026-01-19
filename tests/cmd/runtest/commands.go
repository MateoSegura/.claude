// +build ignore

// This file can be run directly: go run commands.go
// It tests all commands in the knowledge base.
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MateoSegura/.claude/tests"
)

func main() {
	runner := tests.NewTestRunner()
	if runner.DryRun {
		fmt.Println("Warning: Running in dry-run mode (Claude CLI not available)")
		os.Exit(1)
	}
	runner.Verbose = true

	// Set WorkDir to parent of .claude
	workDir, _ := os.Getwd()
	if _, err := os.Stat(filepath.Join(workDir, "commands")); err == nil {
		workDir = filepath.Dir(workDir)
	} else if _, err := os.Stat(filepath.Join(workDir, ".claude", "commands")); err != nil {
		// Try going up from tests/cmd/runtest
		workDir = filepath.Join(workDir, "..", "..", "..")
	}
	runner.WorkDir = workDir

	suite := &tests.Suite{
		Name:          "commands-full",
		ExtensionType: tests.ExtensionCommand,
		Cases: []*tests.TestCase{
			// /new-skill
			{
				Name:          "new-skill-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-skill",
				Prompt:        "Show me how to create a skill called 'language-rust-embedded'. What's the directory structure and SKILL.md template?",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-structure", "Response shows skill directory structure with SKILL.md and explains frontmatter format"),
					tests.ContainsText("SKILL.md"),
				},
			},
			// /new-rule
			{
				Name:          "new-rule-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-rule",
				Prompt:        "Create a rule to prevent console.log in production code. Show me the template.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-rule-format", "Response shows rule markdown format with good/bad examples section"),
					tests.MatchesRegex(`rules/|\.md`),
				},
			},
			// /new-command
			{
				Name:          "new-command-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-command",
				Prompt:        "Show me the template for creating a /deploy command.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-command-format", "Response shows command markdown with frontmatter description field"),
					tests.ContainsText("description"),
				},
			},
			// /new-agent
			{
				Name:          "new-agent-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-agent",
				Prompt:        "Create a code review agent. What template and tools should I use?",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-agent-format", "Response shows agent template with tools and model fields in frontmatter"),
					tests.MatchesRegex(`tools:|Read|Grep`),
				},
			},
			// /new-hook
			{
				Name:          "new-hook-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-hook",
				Prompt:        "Create a hook to run prettier after writing JavaScript files.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-hook-json", "Response shows JSON hook config with PostToolUse event and matcher for JS files"),
					tests.ContainsText("PostToolUse"),
				},
			},
			// /list-extensions
			{
				Name:          "list-extensions-all",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "list-extensions",
				Prompt:        "List all extensions in this knowledge base.",
				Validators: []tests.Validator{
					tests.LLMValidator("lists-extensions", "Response lists extensions by type with meta-skill-create and meta-skill-update mentioned"),
					tests.MatchesRegex(`meta-skill|Commands|Skills`),
				},
			},
			// /update-extension
			{
				Name:          "update-extension-workflow",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "update-extension",
				Prompt:        "Explain the process for updating an existing skill when a library releases a new version.",
				Validators: []tests.Validator{
					tests.LLMValidator("explains-workflow", "Response explains update workflow: analyze current state, research changes, update patterns, test"),
				},
			},
		},
	}

	fmt.Printf("Running command tests: %s\n", suite.Name)
	fmt.Printf("WorkDir: %s\n\n", runner.WorkDir)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	result, err := runner.RunSuite(ctx, suite)
	if err != nil {
		fmt.Printf("Suite execution failed: %v\n", err)
		os.Exit(1)
	}

	// Print results
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Suite: %s\n", result.Name)
	fmt.Printf("Tests: %d total, %d passed, %d failed\n", result.TotalTests, result.Passed, result.Failed)
	fmt.Printf("Score: %.0f%% (Grade: %s)\n", result.Score*100, tests.DefaultGradeScale().Grade(result.Score))
	fmt.Printf("Duration: %v\n", result.Duration)
	fmt.Println(strings.Repeat("=", 60))

	for _, r := range result.Results {
		status := "✓ PASS"
		if !r.Passed {
			status = "✗ FAIL"
		}
		fmt.Printf("\n%s: %s (%.0f%%)\n", status, r.Name, r.Score*100)
		for _, v := range r.Validations {
			vStatus := "  ✓"
			if !v.Passed {
				vStatus = "  ✗"
			}
			fmt.Printf("%s %s: %s\n", vStatus, v.Name, truncate(v.Message, 100))
		}
	}

	if err := runner.SaveSuiteResults(result, "commands-results.json"); err != nil {
		fmt.Printf("\nWarning: couldn't save results: %v\n", err)
	}

	if result.Score < 0.70 {
		fmt.Printf("\n❌ Suite failed: %.0f%% < 70%% threshold\n", result.Score*100)
		os.Exit(1)
	}
	fmt.Printf("\n✅ Suite passed!\n")
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
