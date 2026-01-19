// Command run-tests executes all skill tests and generates a report.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Note: TestReport and SuiteReport types available for future JSON reporting
// Currently using go test output directly

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	iterations := flag.Int("n", 3, "number of iterations per test")
	outputDir := flag.String("o", "/tmp/skill-tests", "output directory")
	flag.Parse()

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║               Claude Code Skill Test Runner                  ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Ensure output directory exists
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output dir: %v\n", err)
		os.Exit(1)
	}

	start := time.Now()

	// Set environment for tests
	os.Setenv("SKILL_TEST", "1")

	// Build test command
	args := []string{"test", "-v", "./..."}
	if *verbose {
		args = append(args, "-v")
	}
	args = append(args, fmt.Sprintf("-count=%d", *iterations))

	// Find test directory
	testDir := findTestDir()
	if testDir == "" {
		fmt.Println("Error: Could not find skill-tests directory")
		os.Exit(1)
	}

	fmt.Printf("Running tests from: %s\n", testDir)
	fmt.Printf("Iterations per test: %d\n", *iterations)
	fmt.Printf("Output directory: %s\n", *outputDir)
	fmt.Println()

	// Run tests
	cmd := exec.Command("go", args...)
	cmd.Dir = testDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "SKILL_TEST=1")

	err := cmd.Run()

	duration := time.Since(start)

	// Generate summary report
	fmt.Println()
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Println("                         TEST SUMMARY                           ")
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Printf("Duration: %v\n", duration)

	if err != nil {
		fmt.Println("Status: SOME TESTS FAILED")
		os.Exit(1)
	} else {
		fmt.Println("Status: ALL TESTS PASSED")
	}
}

func findTestDir() string {
	// Look for skill-tests directory
	candidates := []string{
		".claude/skill-tests",
		"../.claude/skill-tests",
		"../../.claude/skill-tests",
	}

	cwd, _ := os.Getwd()
	for _, c := range candidates {
		path := filepath.Join(cwd, c)
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			return path
		}
	}

	return ""
}
