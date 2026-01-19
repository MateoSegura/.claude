// Command runtest executes extension tests against Claude CLI.
// Usage: go run ./tests/cmd/runtest [suite]
// Suites: skills (default), commands, all
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
	workDir, _ := os.Getwd()
	if _, err := os.Stat(filepath.Join(workDir, "skills")); err == nil {
		workDir = filepath.Dir(workDir)
	}
	runner.WorkDir = workDir

	// Determine which suite to run
	suiteName := "skills"
	if len(os.Args) > 1 {
		suiteName = os.Args[1]
	}

	var suites []*tests.Suite

	switch suiteName {
	case "skills":
		suites = append(suites, skillsSuite())
	case "commands":
		suites = append(suites, commandsSuite())
	case "all":
		suites = append(suites, skillsSuite(), commandsSuite())
	default:
		fmt.Printf("Unknown suite: %s\n", suiteName)
		fmt.Println("Available: skills, commands, all")
		os.Exit(1)
	}

	for _, suite := range suites {
		runSuite(runner, suite)
	}
}

func skillsSuite() *tests.Suite {
	return &tests.Suite{
		Name:          "meta-skill-create",
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
			},
		},
	}
}

func commandsSuite() *tests.Suite {
	return &tests.Suite{
		Name:          "commands",
		ExtensionType: tests.ExtensionCommand,
		Cases: []*tests.TestCase{
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
			{
				Name:          "new-rule-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-rule",
				Prompt:        "Create a rule to prevent console.log in production code. Show me the template.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-rule-format", "Response shows rule markdown format with good/bad examples section"),
				},
			},
			{
				Name:          "new-command-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-command",
				Prompt:        "Show me the template for creating a /deploy command.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-command-format", "Response shows command markdown with frontmatter description field"),
				},
			},
			{
				Name:          "new-agent-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-agent",
				Prompt:        "Show me the markdown template for creating a code review agent. I want to see the frontmatter format with name, description, tools, and model fields.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-agent-format", "Response shows the agent markdown template structure with frontmatter containing tools field"),
				},
			},
			{
				Name:          "new-hook-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-hook",
				Prompt:        "Create a hook to run prettier after writing JavaScript files.",
				Validators: []tests.Validator{
					tests.LLMValidator("shows-hook-json", "Response shows JSON hook config with PostToolUse event"),
					tests.ContainsText("PostToolUse"),
				},
			},
			{
				Name:          "list-extensions-all",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "list-extensions",
				Prompt:        "List all extensions in this knowledge base.",
				Validators: []tests.Validator{
					tests.LLMValidator("lists-extensions", "Response lists extensions by type including skills and commands"),
				},
			},
			{
				Name:          "update-extension-workflow",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "update-extension",
				Prompt:        "Explain the process for updating an existing skill.",
				Validators: []tests.Validator{
					tests.LLMValidator("explains-workflow", "Response explains update workflow with steps like analyze, research, update, test"),
				},
			},
		},
	}
}

func runSuite(runner *tests.TestRunner, suite *tests.Suite) {
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("Running test suite: %s\n", suite.Name)
	fmt.Printf("Extension type: %s\n", suite.ExtensionType)
	fmt.Printf("%s\n\n", strings.Repeat("=", 60))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	result, err := runner.RunSuite(ctx, suite)
	if err != nil {
		fmt.Printf("Suite execution failed: %v\n", err)
		os.Exit(1)
	}

	// Print results
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Suite: %s\n", result.Name)
	fmt.Printf("Tests: %d total, %d passed, %d failed\n", result.TotalTests, result.Passed, result.Failed)
	fmt.Printf("Score: %.0f%% (Grade: %s)\n", result.Score*100, tests.DefaultGradeScale().Grade(result.Score))
	fmt.Printf("Duration: %v\n", result.Duration)
	fmt.Println(strings.Repeat("-", 60))

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
			// Truncate long messages
			msg := v.Message
			if len(msg) > 150 {
				msg = msg[:147] + "..."
			}
			fmt.Printf("%s %s: %s\n", vStatus, v.Name, msg)
		}
	}

	// Save results
	filename := fmt.Sprintf("%s-results.json", suite.Name)
	if err := runner.SaveSuiteResults(result, filename); err != nil {
		fmt.Printf("\nWarning: couldn't save results: %v\n", err)
	} else {
		fmt.Printf("\nResults saved to %s/%s\n", runner.OutputDir, filename)
	}

	if result.Score < 0.70 {
		fmt.Printf("\n❌ Suite failed: %.0f%% < 70%% threshold\n", result.Score*100)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Suite passed!\n")
}
