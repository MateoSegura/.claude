# Benchmark Framework

A meta-evaluation system for testing `.claude` configurations against real-world problems.

## What This Does

Instead of testing "does Claude follow my skill instructions?", this tests:
**"Does my .claude configuration actually make Claude more effective at solving real problems?"**

```
┌─────────────────────────────────────────────────────────────────┐
│   Your .claude config  →  Real GitHub Issues  →  Success Rate  │
│                                                                 │
│   baseline (no config):  54% success                            │
│   your-golang-config:    73% success  (+19 points)              │
└─────────────────────────────────────────────────────────────────┘
```

## Quick Start

```bash
# Dry run - validate corpus without running Claude
go run ./benchmark/cmd/bench --dry-run

# Compare your config against baseline
go run ./benchmark/cmd/bench --config ~/.claude

# Run only easy Go issues
go run ./benchmark/cmd/bench --difficulty easy --language go

# Test a single issue
go run ./benchmark/cmd/bench --issue easy-go-nil-check-001

# Verbose output
go run ./benchmark/cmd/bench --config ~/.claude --verbose
```

## How It Works

### 1. Corpus Definition

Issues are defined in YAML files in `benchmark/corpus/`:

```yaml
issues:
  - id: "easy-go-nil-check-001"
    title: "Fix nil pointer dereference"
    difficulty: easy
    task_type: bug_fix
    language: go
    repo_url: "https://github.com/example/repo"
    repo_ref: "main"
    prompt: |
      Fix the nil pointer dereference in this code...
    eval_method: llm_judge
    success_criteria: |
      The solution must check for nil before access...
```

### 2. Benchmark Execution

For each issue:
1. Clone the repository
2. Apply your `.claude` config (or run without for baseline)
3. Run Claude with the issue prompt
4. Evaluate the result

### 3. Evaluation Methods

| Method | Description |
|--------|-------------|
| `test_suite` | Run repo's test suite (go test, npm test, etc.) |
| `llm_judge` | Claude evaluates if solution meets criteria |
| `custom_check` | Run custom validation script |
| `hybrid` | 60% test suite + 40% LLM judgment |

### 4. Results

Results are saved to `/tmp/claude-benchmark/results/` as JSON:

```json
{
  "config_results": {
    "baseline": {
      "success_rate": 0.54,
      "by_difficulty": {
        "easy": { "success_rate": 0.80 },
        "medium": { "success_rate": 0.50 },
        "hard": { "success_rate": 0.20 }
      }
    },
    "your-config": {
      "success_rate": 0.73,
      ...
    }
  }
}
```

## Creating Your Own Corpus

### Structure

```
benchmark/corpus/
├── sample.yaml      # Example issues
├── golang.yaml      # Go-specific issues
├── typescript.yaml  # TypeScript issues
└── your-domain.yaml # Your custom issues
```

### Issue Fields

| Field | Required | Description |
|-------|----------|-------------|
| `id` | Yes | Unique identifier |
| `title` | Yes | Human-readable title |
| `description` | Yes | What needs to be done |
| `difficulty` | Yes | easy, medium, hard |
| `task_type` | Yes | bug_fix, feature, refactor, test, documentation |
| `language` | Yes | Primary language |
| `repo_url` | Yes | GitHub URL to clone |
| `repo_ref` | No | Branch/tag/commit (default: main) |
| `prompt` | Yes | The actual prompt given to Claude |
| `eval_method` | Yes | How to evaluate success |
| `test_command` | No | For test_suite eval |
| `success_criteria` | No | For llm_judge eval |
| `check_script` | No | For custom_check eval |
| `context_files` | No | Files to include as context |
| `expected_files` | No | Files that should be modified |
| `tags` | No | For filtering |

### Finding Good Issues

Good benchmark issues should be:

1. **Self-contained** - Solvable without external context
2. **Verifiable** - Clear success/failure criteria
3. **Reproducible** - Same starting state each time
4. **Diverse** - Mix of difficulty levels and task types

Sources:
- GitHub issues labeled "good first issue"
- Known bugs with clear repro steps
- Features from "help wanted" issues
- Refactoring tasks with test coverage

## Interpreting Results

### Success Rate Comparison

```
baseline:     54%
your-config:  73%
improvement:  +19 percentage points
```

This means your config helped Claude solve 19% more issues than raw Claude.

### Statistical Significance

For meaningful results:
- Run at least 20-30 issues per difficulty level
- Run multiple iterations of the same benchmark
- Look for consistent improvements across categories

### What to Optimize

If your config helps with:
- **Easy issues only** → Skills are too verbose, adding noise
- **Hard issues only** → Skills provide valuable context
- **Specific languages** → Language-specific skills are working
- **Specific task types** → Role-based skills are working

## Limitations

1. **Corpus quality** - Results are only as good as your test cases
2. **LLM-as-judge variance** - LLM evaluation has inherent variability
3. **Time cost** - Each issue requires cloning + Claude execution
4. **API costs** - Running benchmarks costs API credits

## Future Improvements

- [ ] Parallel issue execution
- [ ] Caching cloned repositories
- [ ] Statistical significance testing
- [ ] Historical result tracking
- [ ] Integration with CI/CD
- [ ] Community corpus sharing
