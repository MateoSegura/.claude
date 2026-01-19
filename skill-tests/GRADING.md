# Skill Test Grading Scale

## Overview

Skills are tested by running prompts through Claude CLI with the skill loaded, then validating the output meets expectations. Multiple iterations are run to assess consistency.

## Grading Scale

| Grade | Score Range | Meaning |
|-------|-------------|---------|
| **A** | 90-100% | Excellent: Skill consistently works as expected |
| **B** | 80-89% | Good: Skill mostly works with minor inconsistencies |
| **C** | 70-79% | Acceptable: Skill works but has noticeable gaps |
| **D** | 60-69% | Poor: Skill has significant issues |
| **F** | Below 60% | Failing: Skill needs major revision |

## Why This Scale?

### A (90%+) - Excellent
- Consistent behavior across multiple runs
- All core validators pass
- Output matches expected patterns
- No errors or unexpected behavior

### B (80-89%) - Good
- Minor inconsistencies in output format
- Most validators pass
- May occasionally miss non-critical expectations
- Still useful and reliable for its purpose

### C (70-79%) - Acceptable
- Some runs fail validators
- Output is correct but may vary in format
- Core functionality works but edge cases may fail
- Needs improvement but is usable

### D (60-69%) - Poor
- Frequent validator failures
- Output often misses expectations
- Inconsistent behavior
- Significant revision needed

### F (Below 60%) - Failing
- Core functionality broken
- Most validators fail
- Skill is not achieving its purpose
- Complete rewrite may be needed

## Test Categories

### 1. Behavioral Tests
Test that the skill changes Claude's behavior as expected.

**Validators:**
- `ContainsText`: Output contains expected text
- `MatchesRegex`: Output matches patterns
- `ContainsCode`: Appropriate code blocks present

### 2. Consistency Tests
Run the same prompt multiple times to check determinism.

**Metrics:**
- Pattern appearance rate across iterations
- Output variance
- Rule adherence rate

### 3. Edge Case Tests
Test boundary conditions and unusual inputs.

**Validators:**
- Error handling
- Rejection of inappropriate requests
- Graceful degradation

## Running Tests

```bash
# Run all skill tests
cd .claude/skill-tests
SKILL_TEST=1 go test -v ./...

# Run specific skill test
SKILL_TEST=1 go test -v -run TestBubbleTeaTUI

# Run with multiple iterations
SKILL_TEST=1 go test -v -count=5 ./...

# Generate detailed output
SKILL_TEST=1 go test -v ./... 2>&1 | tee test-output.txt
```

## Test Output Location

All test outputs are saved to `/tmp/skill-tests/`:
- `*-output.txt`: Raw Claude output for each test
- `*-results.json`: Structured test results

## Interpreting Results

### Example Output
```
Suite: bubbletea-tui
Tests: 18 total, 16 passed, 2 failed
Score: 88.89% (Grade: B)
Duration: 12m34s

FAILED: keyboard-handling (iteration 2) - Score: 50.00%
  - regex: case\s+"q".*tea\.Quit: Pattern match: false
```

### What the Results Mean

1. **Total Tests**: Number of test × iterations
2. **Score**: (Passed validations) / (Total validations)
3. **Grade**: Letter grade based on score
4. **Failed Tests**: Lists which validations failed and why

## Improving Scores

### Common Issues

| Issue | Cause | Fix |
|-------|-------|-----|
| Inconsistent output | Skill prompts too vague | Make rules more specific |
| Missing patterns | Skill doesn't cover case | Add rules/examples |
| Wrong structure | Template issues | Fix scaffolds |
| High variance | Non-deterministic prompts | Add concrete examples |

### Improvement Process

1. Run tests with `-v` flag to see details
2. Check `/tmp/skill-tests/*.txt` for actual outputs
3. Compare failing outputs against expected patterns
4. Update skill rules/examples to be more explicit
5. Re-run tests to verify improvement

## Adding New Tests

```go
{
    Name:   "test-name",
    Skill:  "skill-name",
    Prompt: "Specific task for Claude",
    Validators: []Validator{
        ContainsText("expected output"),
        MatchesRegex(`pattern`),
        CustomValidator("name", func(output string) (bool, string) {
            // Custom validation logic
            return passed, "message"
        }),
    },
    Iterations: 3, // Run multiple times
},
```

## CI Integration

Tests can be run in CI by setting `SKILL_TEST=1`:

```yaml
- name: Run Skill Tests
  run: |
    cd .claude/skill-tests
    SKILL_TEST=1 go test -v ./...
  env:
    ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
```

Note: Tests require Claude CLI to be installed and API key configured.
