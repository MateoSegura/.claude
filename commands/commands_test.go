package commands_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/MateoSegura/.claude/tests"
)

// TestCommands tests all slash commands in this knowledge base.
func TestCommands(t *testing.T) {
	if os.Getenv("SKILL_TEST") == "" {
		t.Skip("Set SKILL_TEST=1 to run command tests (requires Claude CLI)")
	}

	runner := tests.NewTestRunner()
	runner.Verbose = testing.Verbose()

	// Set WorkDir to parent of .claude
	workDir, _ := os.Getwd()
	if _, err := os.Stat(filepath.Join(workDir, "commands")); err == nil {
		workDir = filepath.Dir(workDir)
	}
	runner.WorkDir = workDir

	suite := &tests.Suite{
		Name:          "commands",
		ExtensionType: tests.ExtensionCommand,
		Cases: []*tests.TestCase{
			// /new-skill tests
			{
				Name:          "new-skill-shows-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-skill",
				Prompt:        "I want to create a new skill called 'language-rust-embedded' for Rust embedded development. Show me what structure and template to use.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"shows-skill-structure",
						"Response shows the skill directory structure with SKILL.md file and explains the frontmatter format with name/description fields",
					),
					tests.ContainsText("SKILL.md"),
				},
				Iterations: 1,
			},
			{
				Name:          "new-skill-validates-naming",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-skill",
				Prompt:        "Create a skill called 'myskill' - is this name correct?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"explains-naming-convention",
						"Response explains that the name should follow the convention {layer}-{category}-{specific} and suggests a better name format",
					),
				},
				Iterations: 1,
			},

			// /new-rule tests
			{
				Name:          "new-rule-shows-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-rule",
				Prompt:        "I want to create a rule that prevents hardcoded API keys in code. Show me how.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"shows-rule-template",
						"Response shows the rule markdown template with sections for the rule statement, why it matters, and good/bad examples",
					),
					tests.MatchesRegex(`rules/|\.md`),
				},
				Iterations: 1,
			},
			{
				Name:          "new-rule-suggests-category",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-rule",
				Prompt:        "Where should I put a rule about error handling in Go?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"suggests-category",
						"Response suggests an appropriate category/directory for the rule (like coding/ or a language-specific folder)",
					),
				},
				Iterations: 1,
			},

			// /new-command tests
			{
				Name:          "new-command-shows-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-command",
				Prompt:        "I want to create a /deploy command that deploys to staging. Show me the template.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"shows-command-template",
						"Response shows the command markdown template with frontmatter description and usage sections",
					),
					tests.ContainsText("description"),
				},
				Iterations: 1,
			},
			{
				Name:          "new-command-vs-skill",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-command",
				Prompt:        "Should I create a command or a skill for teaching Claude how to write good commit messages?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"recommends-skill-over-command",
						"Response explains that a skill is better for teaching multi-step processes or workflows, while commands are for user-initiated actions",
					),
				},
				Iterations: 1,
			},

			// /new-agent tests
			{
				Name:          "new-agent-shows-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-agent",
				Prompt:        "Show me the markdown template for creating a code review agent. I want to see the frontmatter format with name, description, tools, and model fields.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"shows-agent-template",
						"Response shows the agent markdown template structure with frontmatter containing tools field",
					),
					tests.MatchesRegex(`tools:|Read|Grep|Glob`),
				},
				Iterations: 1,
			},
			{
				Name:          "new-agent-tool-selection",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-agent",
				Prompt:        "What tools should a security scanning agent have?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"recommends-appropriate-tools",
						"Response recommends appropriate tools for security scanning (at minimum Read, Grep, Glob, possibly Bash) and explains why",
					),
				},
				Iterations: 1,
			},

			// /new-hook tests
			{
				Name:          "new-hook-shows-events",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-hook",
				Prompt:        "What hook events are available and when do they trigger?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"explains-hook-events",
						"Response explains the available hook events: PreToolUse, PostToolUse, and when each triggers",
					),
					tests.MatchesRegex(`PreToolUse|PostToolUse`),
				},
				Iterations: 1,
			},
			{
				Name:          "new-hook-json-template",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "new-hook",
				Prompt:        "Show me how to create a hook that runs black formatter after Python files are written.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"shows-hook-json",
						"Response shows valid JSON hook configuration with matcher for Python files and command to run black",
					),
					tests.ContainsText("PostToolUse"),
					tests.MatchesRegex(`black|\.py`),
				},
				Iterations: 1,
			},

			// /list-extensions tests
			{
				Name:          "list-extensions-discovers",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "list-extensions",
				Prompt:        "List all extensions in this knowledge base.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"lists-extension-types",
						"Response lists extensions organized by type (Skills, Commands, Rules, etc.) with at least the meta-skill-create and meta-skill-update skills mentioned",
					),
					tests.MatchesRegex(`meta-skill-create|meta-skill-update`),
				},
				Iterations: 1,
			},
			{
				Name:          "list-extensions-shows-commands",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "list-extensions",
				Prompt:        "Show me just the commands available.",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"lists-commands",
						"Response lists the available commands including new-skill, new-rule, new-command, new-agent, new-hook",
					),
					tests.MatchesRegex(`new-skill|new-rule|new-command`),
				},
				Iterations: 1,
			},

			// /update-extension tests
			{
				Name:          "update-extension-workflow",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "update-extension",
				Prompt:        "What's the process for updating an existing skill?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"explains-update-workflow",
						"Response explains the update workflow including: discovering/selecting the extension, analyzing current state, researching updates, and testing",
					),
				},
				Iterations: 1,
			},
			{
				Name:          "update-extension-version-bump",
				ExtensionType: tests.ExtensionCommand,
				Extension:     "update-extension",
				Prompt:        "A library I have a skill for just released v2.0 with breaking changes. How do I update the skill?",
				Validators: []tests.Validator{
					tests.LLMValidator(
						"explains-version-update",
						"Response explains how to handle library version updates: check changelog, identify breaking changes, update patterns/examples, and test",
					),
				},
				Iterations: 1,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	result, err := runner.RunSuite(ctx, suite)
	if err != nil {
		t.Fatalf("Suite execution failed: %v", err)
	}

	if err := runner.SaveSuiteResults(result, "commands-results.json"); err != nil {
		t.Logf("Warning: couldn't save results: %v", err)
	}

	t.Logf("Suite: %s", result.Name)
	t.Logf("Tests: %d total, %d passed, %d failed", result.TotalTests, result.Passed, result.Failed)
	t.Logf("Score: %.2f%% (Grade: %s)", result.Score*100, tests.DefaultGradeScale().Grade(result.Score))
	t.Logf("Duration: %v", result.Duration)

	for _, r := range result.Results {
		if !r.Passed {
			t.Logf("FAILED: %s - Score: %.2f%%", r.Name, r.Score*100)
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
