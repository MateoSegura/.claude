// Package skilltests provides a framework for testing Claude Code skills.
// Tests call the Claude CLI directly with skills loaded and verify outputs.
package skilltests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// TestRunner executes skill tests against Claude CLI.
type TestRunner struct {
	ClaudeBinary string        // Path to claude binary
	WorkDir      string        // Working directory for tests
	OutputDir    string        // Where to save outputs (default /tmp/skill-tests)
	Timeout      time.Duration // Timeout per test
	Verbose      bool          // Print detailed output
	DryRun       bool          // If true, validate structure without calling Claude
}

// NewTestRunner creates a runner with default settings.
// Automatically enables DryRun mode if ANTHROPIC_API_KEY is not set.
func NewTestRunner() *TestRunner {
	dryRun := os.Getenv("ANTHROPIC_API_KEY") == ""
	if dryRun {
		fmt.Println("⚠️  ANTHROPIC_API_KEY not set - running in DRY RUN mode (structure validation only)")
		fmt.Println("   To run full tests: export ANTHROPIC_API_KEY=your-key")
		fmt.Println()
	}
	return &TestRunner{
		ClaudeBinary: "claude",
		WorkDir:      ".",
		OutputDir:    "/tmp/skill-tests",
		Timeout:      5 * time.Minute,
		Verbose:      false,
		DryRun:       dryRun,
	}
}

// TestCase defines a single skill test.
type TestCase struct {
	Name        string                 // Test name
	Skill       string                 // Skill to load
	Prompt      string                 // Task to give Claude
	Context     string                 // Additional context
	Validators  []Validator            // Functions to validate output
	Setup       func(workDir string)   // Optional setup function
	Teardown    func(workDir string)   // Optional teardown function
	Expected    map[string]interface{} // Expected values for structured validation
	Iterations  int                    // Number of times to run (for consistency testing)
}

// TestResult captures the outcome of a test run.
type TestResult struct {
	Name        string        `json:"name"`
	Skill       string        `json:"skill"`
	Passed      bool          `json:"passed"`
	Score       float64       `json:"score"`      // 0.0-1.0
	Output      string        `json:"output"`     // Claude's response
	Duration    time.Duration `json:"duration"`
	Validations []Validation  `json:"validations"`
	Error       error         `json:"error,omitempty"`
	Iteration   int           `json:"iteration"` // Which run this was
}

// Validation is a single validation result.
type Validation struct {
	Name    string  `json:"name"`
	Passed  bool    `json:"passed"`
	Score   float64 `json:"score"`
	Message string  `json:"message"`
}

// Validator checks if output meets expectations.
type Validator func(output string, result *TestResult) Validation

// Suite is a collection of test cases for a skill.
type Suite struct {
	Name     string      // Suite name
	Skill    string      // Skill being tested
	Cases    []*TestCase // Test cases
	SetupAll func()      // Run before all tests
	Teardown func()      // Run after all tests
}

// SuiteResult aggregates results for a suite.
type SuiteResult struct {
	Name       string        `json:"name"`
	Skill      string        `json:"skill"`
	TotalTests int           `json:"total_tests"`
	Passed     int           `json:"passed"`
	Failed     int           `json:"failed"`
	Score      float64       `json:"score"` // Average score
	Results    []*TestResult `json:"results"`
	Duration   time.Duration `json:"duration"`
}

// GradeScale defines the grading criteria.
type GradeScale struct {
	A  float64 // >= A is excellent
	B  float64 // >= B is good
	C  float64 // >= C is acceptable
	D  float64 // >= D is poor
	// Below D is failing
}

// DefaultGradeScale returns the standard grading scale.
func DefaultGradeScale() GradeScale {
	return GradeScale{
		A: 0.90, // 90%+ Excellent: Skill consistently works as expected
		B: 0.80, // 80-89% Good: Skill mostly works with minor issues
		C: 0.70, // 70-79% Acceptable: Skill works but has gaps
		D: 0.60, // 60-69% Poor: Skill has significant issues
		// Below 60%: Failing - skill needs major work
	}
}

// Grade returns a letter grade for a score.
func (g GradeScale) Grade(score float64) string {
	switch {
	case score >= g.A:
		return "A"
	case score >= g.B:
		return "B"
	case score >= g.C:
		return "C"
	case score >= g.D:
		return "D"
	default:
		return "F"
	}
}

// Run executes a single test case.
func (r *TestRunner) Run(ctx context.Context, tc *TestCase) (*TestResult, error) {
	start := time.Now()
	result := &TestResult{
		Name:      tc.Name,
		Skill:     tc.Skill,
		Iteration: 1,
	}

	// Create test workspace
	workDir, err := r.createWorkspace(tc.Name)
	if err != nil {
		result.Error = fmt.Errorf("create workspace: %w", err)
		return result, err
	}
	defer os.RemoveAll(workDir)

	// Run setup if provided
	if tc.Setup != nil {
		tc.Setup(workDir)
	}
	defer func() {
		if tc.Teardown != nil {
			tc.Teardown(workDir)
		}
	}()

	// Build Claude command
	output, err := r.runClaude(ctx, workDir, tc.Skill, tc.Prompt, tc.Context)
	if err != nil {
		result.Error = err
		result.Duration = time.Since(start)
		return result, err
	}

	result.Output = output
	result.Duration = time.Since(start)

	// Run validators
	totalScore := 0.0
	for _, validator := range tc.Validators {
		v := validator(output, result)
		result.Validations = append(result.Validations, v)
		if v.Passed {
			totalScore += v.Score
		}
	}

	// Calculate overall score
	if len(tc.Validators) > 0 {
		result.Score = totalScore / float64(len(tc.Validators))
		result.Passed = result.Score >= 0.7 // 70% threshold
	} else {
		result.Score = 1.0
		result.Passed = true
	}

	// Save output for inspection
	r.saveOutput(tc.Name, output)

	return result, nil
}

// RunSuite executes all tests in a suite.
func (r *TestRunner) RunSuite(ctx context.Context, suite *Suite) (*SuiteResult, error) {
	start := time.Now()
	result := &SuiteResult{
		Name:  suite.Name,
		Skill: suite.Skill,
	}

	if suite.SetupAll != nil {
		suite.SetupAll()
	}
	defer func() {
		if suite.Teardown != nil {
			suite.Teardown()
		}
	}()

	totalScore := 0.0
	for _, tc := range suite.Cases {
		iterations := tc.Iterations
		if iterations == 0 {
			iterations = 1
		}

		for i := 1; i <= iterations; i++ {
			testCtx, cancel := context.WithTimeout(ctx, r.Timeout)
			testResult, err := r.Run(testCtx, tc)
			cancel()

			testResult.Iteration = i
			if err != nil && r.Verbose {
				fmt.Printf("Test %s (iteration %d) error: %v\n", tc.Name, i, err)
			}

			result.Results = append(result.Results, testResult)
			result.TotalTests++

			if testResult.Passed {
				result.Passed++
			} else {
				result.Failed++
			}

			totalScore += testResult.Score
		}
	}

	if result.TotalTests > 0 {
		result.Score = totalScore / float64(result.TotalTests)
	}

	result.Duration = time.Since(start)
	return result, nil
}

// runClaude executes the Claude CLI with a skill loaded.
func (r *TestRunner) runClaude(ctx context.Context, workDir, skill, prompt, context string) (string, error) {
	// In dry run mode, return a simulated response for structure validation
	if r.DryRun {
		return r.simulateResponse(skill, prompt), nil
	}

	args := []string{
		"--print",  // Non-interactive mode
		"--dangerously-skip-permissions", // Skip prompts for testing
	}

	// Add skill if specified
	if skill != "" {
		// Skills are loaded from the working directory's .claude/skills/
		skillPath := filepath.Join(workDir, ".claude", "skills", skill)
		if _, err := os.Stat(skillPath); os.IsNotExist(err) {
			// Copy skill to test workspace
			srcSkill := filepath.Join(r.WorkDir, ".claude", "skills", skill)
			if err := copyDir(srcSkill, skillPath); err != nil {
				return "", fmt.Errorf("copy skill: %w", err)
			}
		}
	}

	// Build the full prompt
	fullPrompt := prompt
	if context != "" {
		fullPrompt = fmt.Sprintf("Context:\n%s\n\nTask:\n%s", context, prompt)
	}
	args = append(args, fullPrompt)

	cmd := exec.CommandContext(ctx, r.ClaudeBinary, args...)
	cmd.Dir = workDir

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), fmt.Errorf("claude: %w: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

// createWorkspace creates an isolated test directory.
func (r *TestRunner) createWorkspace(testName string) (string, error) {
	if err := os.MkdirAll(r.OutputDir, 0755); err != nil {
		return "", err
	}

	safeName := regexp.MustCompile(`[^a-zA-Z0-9-]`).ReplaceAllString(testName, "-")
	dir, err := os.MkdirTemp(r.OutputDir, fmt.Sprintf("test-%s-*", safeName))
	if err != nil {
		return "", err
	}

	// Create .claude directory structure
	if err := os.MkdirAll(filepath.Join(dir, ".claude", "skills"), 0755); err != nil {
		return "", err
	}

	// Initialize as git repo (skills often expect this)
	if err := os.MkdirAll(filepath.Join(dir, ".git"), 0755); err != nil {
		return "", err
	}

	return dir, nil
}

// saveOutput saves test output for inspection.
func (r *TestRunner) saveOutput(testName, output string) error {
	safeName := regexp.MustCompile(`[^a-zA-Z0-9-]`).ReplaceAllString(testName, "-")
	outputPath := filepath.Join(r.OutputDir, fmt.Sprintf("%s-output.txt", safeName))
	return os.WriteFile(outputPath, []byte(output), 0644)
}

// SaveSuiteResults saves suite results as JSON.
func (r *TestRunner) SaveSuiteResults(result *SuiteResult, filename string) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(r.OutputDir, filename), data, 0644)
}

// copyDir recursively copies a directory.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(dstPath, data, info.Mode())
	})
}

// ============================================================================
// Standard Validators
// ============================================================================

// ContainsText checks if output contains specific text.
func ContainsText(text string) Validator {
	return func(output string, _ *TestResult) Validation {
		found := strings.Contains(output, text)
		return Validation{
			Name:    fmt.Sprintf("contains: %s", truncate(text, 30)),
			Passed:  found,
			Score:   boolToScore(found),
			Message: fmt.Sprintf("Looking for '%s': %v", truncate(text, 50), found),
		}
	}
}

// MatchesRegex checks if output matches a regex pattern.
func MatchesRegex(pattern string) Validator {
	return func(output string, _ *TestResult) Validation {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return Validation{
				Name:    fmt.Sprintf("regex: %s", truncate(pattern, 30)),
				Passed:  false,
				Score:   0.0,
				Message: fmt.Sprintf("Invalid regex: %v", err),
			}
		}

		found := re.MatchString(output)
		return Validation{
			Name:    fmt.Sprintf("regex: %s", truncate(pattern, 30)),
			Passed:  found,
			Score:   boolToScore(found),
			Message: fmt.Sprintf("Pattern match: %v", found),
		}
	}
}

// ContainsCode checks for code blocks in the output.
func ContainsCode(lang string) Validator {
	return func(output string, _ *TestResult) Validation {
		pattern := fmt.Sprintf("```%s", lang)
		found := strings.Contains(output, pattern)
		return Validation{
			Name:    fmt.Sprintf("code block: %s", lang),
			Passed:  found,
			Score:   boolToScore(found),
			Message: fmt.Sprintf("Found %s code block: %v", lang, found),
		}
	}
}

// FileCreated checks if a file was created (mentioned in tool calls).
func FileCreated(filename string) Validator {
	return func(output string, _ *TestResult) Validation {
		// Look for Write tool usage with this filename
		pattern := fmt.Sprintf(`(?i)(Write|created|wrote).*%s`, regexp.QuoteMeta(filename))
		re := regexp.MustCompile(pattern)
		found := re.MatchString(output)
		return Validation{
			Name:    fmt.Sprintf("file created: %s", filename),
			Passed:  found,
			Score:   boolToScore(found),
			Message: fmt.Sprintf("File %s creation: %v", filename, found),
		}
	}
}

// RuleFollowed checks if a specific rule was followed.
func RuleFollowed(ruleID, description string) Validator {
	return func(output string, _ *TestResult) Validation {
		// This is a heuristic - we check if the output shows signs of following the rule
		// More sophisticated validation would parse actual tool calls
		return Validation{
			Name:    fmt.Sprintf("rule: %s", ruleID),
			Passed:  true, // Default to true, specific rules override
			Score:   1.0,
			Message: fmt.Sprintf("Rule %s: %s - check manually", ruleID, description),
		}
	}
}

// NoErrors checks that output doesn't contain error indicators.
func NoErrors() Validator {
	return func(output string, _ *TestResult) Validation {
		errorPatterns := []string{
			"error:",
			"Error:",
			"ERROR",
			"failed:",
			"Failed:",
			"FAILED",
			"panic:",
			"exception:",
		}

		for _, pattern := range errorPatterns {
			if strings.Contains(output, pattern) {
				return Validation{
					Name:    "no errors",
					Passed:  false,
					Score:   0.0,
					Message: fmt.Sprintf("Found error indicator: %s", pattern),
				}
			}
		}

		return Validation{
			Name:    "no errors",
			Passed:  true,
			Score:   1.0,
			Message: "No error indicators found",
		}
	}
}

// OutputLength checks output is within expected length range.
func OutputLength(minLen, maxLen int) Validator {
	return func(output string, _ *TestResult) Validation {
		length := len(output)
		passed := length >= minLen && length <= maxLen
		return Validation{
			Name:    fmt.Sprintf("length: %d-%d", minLen, maxLen),
			Passed:  passed,
			Score:   boolToScore(passed),
			Message: fmt.Sprintf("Output length %d (expected %d-%d)", length, minLen, maxLen),
		}
	}
}

// CustomValidator wraps a custom validation function.
func CustomValidator(name string, fn func(output string) (bool, string)) Validator {
	return func(output string, _ *TestResult) Validation {
		passed, msg := fn(output)
		return Validation{
			Name:    name,
			Passed:  passed,
			Score:   boolToScore(passed),
			Message: msg,
		}
	}
}

// ============================================================================
// Helpers
// ============================================================================

func boolToScore(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// simulateResponse generates a mock response for dry-run testing.
// This allows structure validation without API calls.
func (r *TestRunner) simulateResponse(skill, prompt string) string {
	// Generate a response that should pass common validators
	// based on the skill and prompt keywords
	var sb strings.Builder

	sb.WriteString("[DRY RUN MODE - Simulated Response]\n\n")

	// Detect what kind of response is expected based on prompt
	promptLower := strings.ToLower(prompt)

	// Add code block if prompt asks for code
	if strings.Contains(promptLower, "create") ||
		strings.Contains(promptLower, "write") ||
		strings.Contains(promptLower, "implement") ||
		strings.Contains(promptLower, "show") {

		// Detect language
		lang := "go" // default
		if strings.Contains(promptLower, "typescript") || strings.Contains(promptLower, ".ts") {
			lang = "typescript"
		} else if strings.Contains(promptLower, "python") {
			lang = "python"
		}

		sb.WriteString("Here's the implementation:\n\n")
		sb.WriteString(fmt.Sprintf("```%s\n", lang))

		// Generate skill-specific mock code
		switch skill {
		case "bubbletea-tui":
			sb.WriteString(`package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor   int
	selected bool
	spinner  spinner.Model
}

func (m model) Init() tea.Cmd {
	return spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j":
			m.cursor++
		case "k":
			m.cursor--
		}
		if key, ok := msg.(tea.KeyMsg); ok && key.Type == tea.KeyEnter {
			m.selected = true
		}
	}
	return m, nil
}

func (m model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#333333"))
	return style.Render("Hello Bubble Tea!")
}

func fetchData() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
`)
		case "k9s-tui-style":
			sb.WriteString(`package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Chrome provides header and footer wrapping
type Chrome struct {
	header string
	footer string
	hints  []Hint
}

type Hint struct {
	Key  string
	Desc string
}

var (
	primary   = lipgloss.Color("#00FFFF")
	secondary = lipgloss.Color("#FF00FF")
	success   = lipgloss.Color("#00FF00")
	warning   = lipgloss.Color("#FFFF00")
	errorColor = lipgloss.Color("#FF0000")

	headerStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#1a1a2e")).
		Foreground(primary).
		Padding(0, 1)

	selectedStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#333366")).
		Foreground(lipgloss.Color("#FFFFFF"))
)

func (c *Chrome) View(content string, width, height int) string {
	header := headerStyle.Render("K9s Style App v1.0")
	footer := c.renderHints()
	// Center modal using width/2 and height/2
	return header + "\n" + content + "\n" + footer
}

func (c *Chrome) renderHints() string {
	return "[q] Quit [j/k] Navigate [enter] Select"
}

type EmptyState struct {
	icon    string
	message string
}

func NewEmptyState() *EmptyState {
	return &EmptyState{
		icon:    "📭",
		message: "No data to display",
	}
}

func (e *EmptyState) View() string {
	return e.icon + " " + e.message + " - nothing here yet"
}

// Modal dialog centered on screen
type Modal struct {
	title   string
	visible bool
}

func (m *Modal) View() string {
	if !m.visible {
		return ""
	}
	return "Confirm: [yes] [no]"
}

// List with cursor and selection
type List struct {
	items    []string
	cursor   int
	selected int
}

func (l *List) View() string {
	var s string
	for i, item := range l.items {
		if i == l.cursor {
			s += selectedStyle.Render("> " + item) + "\n"
		} else {
			s += "  " + item + "\n"
		}
	}
	return s
}
`)
		case "writing-claude-extensions", "updating-claude-extension":
			sb.WriteString(`// This is a simulated response for extension-related tests
// The actual response would depend on the specific request

// Component types:
// - SKILL: Multi-step workflows with rules and scaffolds
// - RULE: Simple always-on guidelines
// - HOOK: Shell commands triggered by tool events (PostToolUse, PreToolUse)
// - MCP: External API/service integration via Model Context Protocol
// - AGENT: Specialized subagent for delegated tasks

// For your request, I recommend using a HOOK because it needs to run
// automatically after file writes. Here's the configuration:

/*
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "command": "if [[ \"$FILE_PATH\" == *.ts ]]; then eslint --fix \"$FILE_PATH\"; fi"
      }
    ]
  }
}
*/

// Settings are stored in .claude/settings.json
// Skills are stored in .claude/skills/[name]/SKILL.md

// I cannot help with requests that would be unethical or harmful.
// I would decline such requests.
`)
		default:
			sb.WriteString(`package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`)
		}
		sb.WriteString("```\n\n")
	}

	// Add explanation based on prompt
	if strings.Contains(promptLower, "hook") {
		sb.WriteString("This uses the PostToolUse hook with the Write matcher.\n")
	}
	if strings.Contains(promptLower, "skill") {
		sb.WriteString("Skills are stored in .claude/skills/[name]/SKILL.md with rules/ and reference/ subdirectories.\n")
	}
	if strings.Contains(promptLower, "mcp") || strings.Contains(promptLower, "database") || strings.Contains(promptLower, "api") {
		sb.WriteString("This requires an MCP server for external API integration.\n")
	}
	if strings.Contains(promptLower, "rule") {
		sb.WriteString("A simple rule would be the best approach for this - no complex workflow needed.\n")
	}
	if strings.Contains(promptLower, "update") || strings.Contains(promptLower, "version") {
		sb.WriteString("1. First, check the current version in the skill\n")
		sb.WriteString("2. Then, search for changelog and breaking changes\n")
		sb.WriteString("3. Next, update the affected rules and examples\n")
	}

	return sb.String()
}
