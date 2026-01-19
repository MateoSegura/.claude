// Command runtest executes extension tests against Claude CLI.
// Usage: go run ./tests/cmd/runtest
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
		fmt.Println("Results will be simulated, not real.")
		fmt.Println()
	}
	runner.Verbose = true

	// WorkDir should be the project root (parent of .claude)
	// If we're in .claude directory, go up one level
	workDir, _ := os.Getwd()
	if _, err := os.Stat(filepath.Join(workDir, "skills")); err == nil {
		// We're in the .claude directory, go up one level
		workDir = filepath.Dir(workDir)
	}
	runner.WorkDir = workDir

	// Simple focused test suite
	suite := &tests.Suite{
		Name:          "meta-skill-create-basic",
		ExtensionType: tests.ExtensionSkill,
		Extension:     "meta-skill-create",
		Cases: []*tests.TestCase{
			{
				Name:      "recommends-hook-for-formatting",
				Extension: "meta-skill-create",
				Prompt:    "I want to automatically run gofmt after Claude writes Go files. What type of Claude Code extension should I create?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"identifies-hook",
						"The response recommends using a HOOK (specifically mentioning PostToolUse or post-tool-use event) for automatically running commands after file writes",
					),
				},
				Iterations: 1,
			},
			{
				Name:      "recommends-mcp-for-database",
				Extension: "meta-skill-create",
				Prompt:    "I want Claude to be able to query my PostgreSQL database directly. What extension type should I use?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"identifies-mcp",
						"The response recommends using an MCP server (Model Context Protocol) for database access",
					),
				},
				Iterations: 1,
			},
			{
				Name:      "recommends-rule-for-constraint",
				Extension: "meta-skill-create",
				Prompt:    "I want Claude to never use the 'any' type when writing TypeScript. What's the simplest extension for this?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"identifies-rule",
						"The response recommends using a RULE (not a skill) for simple always-on constraints",
					),
				},
				Iterations: 1,
			},
		},
	}

	fmt.Printf("Running test suite: %s\n", suite.Name)
	fmt.Printf("Extension: %s (%s)\n\n", suite.Extension, suite.ExtensionType)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
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
			fmt.Printf("%s %s: %s\n", vStatus, v.Name, v.Message)
		}
	}

	// Save results
	if err := runner.SaveSuiteResults(result, "test-results.json"); err != nil {
		fmt.Printf("\nWarning: couldn't save results: %v\n", err)
	} else {
		fmt.Printf("\nResults saved to %s/test-results.json\n", runner.OutputDir)
	}

	if result.Score < 0.70 {
		fmt.Printf("\n❌ Suite failed: %.0f%% < 70%% threshold\n", result.Score*100)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Suite passed!\n")
}
