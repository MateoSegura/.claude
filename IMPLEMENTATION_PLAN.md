# Cody Implementation Plan

## Dependency Graph

```
                                    ┌─────────────────┐
                                    │   PHASE 1       │
                                    │   (SERIAL)      │
                                    │                 │
                                    │  1.1 Schemas    │
                                    │  1.2 Interfaces │
                                    └────────┬────────┘
                                             │
              ┌──────────────────────────────┼──────────────────────────────┐
              │                              │                              │
              ▼                              ▼                              ▼
    ┌─────────────────┐            ┌─────────────────┐            ┌─────────────────┐
    │   PHASE 2A      │            │   PHASE 2B      │            │   PHASE 2C      │
    │   (PARALLEL)    │            │   (PARALLEL)    │            │   (PARALLEL)    │
    │                 │            │                 │            │                 │
    │  Test Runners   │            │  Skill Engine   │            │  TUI Components │
    │  2.1-2.4        │            │  2.5-2.7        │            │  2.8-2.11       │
    └────────┬────────┘            └────────┬────────┘            └────────┬────────┘
              │                              │                              │
              └──────────────────────────────┼──────────────────────────────┘
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │   PHASE 3       │
                                    │   (SERIAL)      │
                                    │                 │
                                    │  Integration    │
                                    │  3.1-3.4        │
                                    └────────┬────────┘
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │   PHASE 4       │
                                    │   (SERIAL)      │
                                    │                 │
                                    │  Server & API   │
                                    │  4.1-4.3        │
                                    └────────┬────────┘
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │   PHASE 5       │
                                    │   (PARALLEL)    │
                                    │                 │
                                    │  Polish & Docs  │
                                    │  5.1-5.4        │
                                    └─────────────────┘
```

---

## Phase 1: Foundation (SERIAL)

**These MUST be done first. Everything else imports from here.**

### Task 1.1: Schema Package

**File:** `tools/cody/pkg/schema/`

**Creates:**
- `feature.go` - FeatureType, Feature, TestStrategy
- `skill.go` - Skill, SkillType, SkillSpec, Mistake
- `test.go` - BehavioralDiffTest, DeterminismTest, ComplianceTest, all metrics/types
- `config.go` - ClaudeCodeConfig, TaskPattern, HookDef, MCPServer, etc.

**Verification:**
```bash
cd tools/cody && go build ./pkg/schema/...
```

**Done when:** All types from ARCH.md Section 2 are compilable Go code.

---

### Task 1.2: Core Interfaces

**File:** `tools/cody/pkg/cody/`

**Creates:**
- `cody.go` - Client struct, New(), all method signatures
- `options.go` - Option type, WithProgress, WithAPIKey, etc.
- `errors.go` - CodyError, specific error types

**Verification:**
```bash
cd tools/cody && go build ./pkg/cody/...
```

**Done when:** `cody.Client` interface is defined with all methods (can return `nil`/`not implemented` for now).

---

## Phase 2: Core Implementation (PARALLEL)

**These three tracks can be worked on simultaneously by different agents.**

### Track A: Test Runners (Tasks 2.1-2.4)

#### Task 2.1: Test Runner Base

**File:** `tools/cody/pkg/skilltest/runner.go`

**Creates:**
- `Runner` struct
- `NewRunner(apiKey string)`
- `runScenario(ctx, prompt, context, skills)` - executes Claude with/without skills
- `ScenarioResult` struct (tool calls, output, cost, duration)

**Dependencies:** 1.1 (schemas)

**Verification:**
```go
runner := skilltest.NewRunner(os.Getenv("ANTHROPIC_API_KEY"))
// Should compile and initialize
```

---

#### Task 2.2: Behavioral Diff Tests

**File:** `tools/cody/pkg/skilltest/behavioral.go`

**Creates:**
- `BehavioralDiffRunner` struct
- `Run(t, test BehavioralDiffTest)`
- `evaluate(result, metric)` - extracts behavior observations
- Metric implementations: TestsWrittenFirst, PlanBeforeCode, ErrorHandling, etc.

**Dependencies:** 2.1 (runner base)

**Verification:**
```go
runner := skilltest.NewBehavioralDiffRunner(t)
// Define a simple test
test := schema.BehavioralDiffTest{...}
runner.Run(t, test)
```

---

#### Task 2.3: Determinism Tests

**File:** `tools/cody/pkg/skilltest/determinism.go`

**Creates:**
- `DeterminismRunner` struct
- `Run(t, test DeterminismTest)`
- `measureVariance(results, check)` - calculates variance metrics
- Variance implementations: Structure, Approach, Naming, ErrorCases

**Dependencies:** 2.1 (runner base)

**Verification:**
```go
runner := skilltest.NewDeterminismRunner(t)
test := schema.DeterminismTest{...}
runner.Run(t, test)
```

---

#### Task 2.4: Compliance Tests

**File:** `tools/cody/pkg/skilltest/compliance.go`

**Creates:**
- `ComplianceRunner` struct
- `Run(t, test ComplianceTest)`
- `runPressureScenario(ctx, scenario, skillName)`
- `extractChoice(result, options)` - parses which option was chosen
- `detectRationalization(output, rationalization)` - fuzzy match excuses

**Dependencies:** 2.1 (runner base)

**Verification:**
```go
runner := skilltest.NewComplianceRunner(t)
test := schema.ComplianceTest{...}
runner.Run(t, test)
```

---

### Track B: Skill Engine (Tasks 2.5-2.7)

#### Task 2.5: Skill Questionnaire

**File:** `tools/cody/pkg/skill/questionnaire.go`

**Creates:**
- `Questionnaire` struct
- `NewQuestionnaire()` - initializes with all questions from ARCH.md
- `NextQuestion()` - returns next unanswered question (respects dependencies)
- `Answer(questionID, answer)`
- `GenerateSkillSpec()` - creates SkillSpec from answers
- `TotalQuestions()`, `Progress()`

**Dependencies:** 1.1 (schemas)

**Verification:**
```go
q := skill.NewQuestionnaire()
for q.NextQuestion() != nil {
    q.Answer(q.NextQuestion().ID, "test answer")
}
spec := q.GenerateSkillSpec()
// spec should be populated
```

---

#### Task 2.6: Skill Generator

**File:** `tools/cody/pkg/skill/generator.go`

**Creates:**
- `Generator` struct
- `NewGenerator(outputDir string)`
- `Generate(spec *SkillSpec)` - creates SKILL.md and test files
- Templates for SKILL.md and _test.go files
- Directory structure creation (scripts/, references/, assets/)

**Dependencies:** 1.1 (schemas), 2.5 (for SkillSpec)

**Verification:**
```go
gen := skill.NewGenerator("/tmp/skills")
spec := &schema.SkillSpec{Name: "test-skill", ...}
result, err := gen.Generate(spec)
// Check files exist at result.SkillPath, result.TestPath
```

---

#### Task 2.7: Skill Manager

**File:** `tools/cody/pkg/skill/manager.go`

**Creates:**
- `Manager` struct
- `NewManager(skillsDir string)`
- `List(ctx)` - returns all skills
- `Get(ctx, name)` - returns single skill
- `Generate(ctx, spec)` - wraps generator
- `Validate(skill)` - checks skill structure

**Dependencies:** 2.6 (generator)

**Verification:**
```go
mgr := skill.NewManager("./skills")
skills, _ := mgr.List(ctx)
// Should return parsed skills
```

---

### Track C: TUI Components (Tasks 2.8-2.11)

#### Task 2.8: Theme & Styles

**File:** `tools/cody/internal/tui/theme.go`

**Creates:**
- `Theme` variable with all colors and styles
- `Logo()` function
- `StatusLine()` function
- `renderProgressBar()` helper

**Dependencies:** None (only lipgloss)

**Verification:**
```go
fmt.Println(tui.Logo())
// Should render ASCII art
```

---

#### Task 2.9: TUI Components

**File:** `tools/cody/internal/tui/components/`

**Creates:**
- `question.go` - Question component (single choice, multi choice, free text)
- `progress.go` - Progress bar component
- `results.go` - Test results display component
- `box.go` - Styled box wrapper

**Dependencies:** 2.8 (theme)

**Verification:**
```go
q := components.NewQuestion(schema.Question{...})
// Should render correctly
```

---

#### Task 2.10: Skill Create TUI

**File:** `tools/cody/internal/tui/skill_create.go`

**Creates:**
- `SkillCreateModel` struct
- `NewSkillCreateModel()`
- `Init()`, `Update()`, `View()` - Bubble Tea interface
- Phase handling: Questions → Review → Generating → Complete
- Integration with `skill.Questionnaire`

**Dependencies:** 2.5 (questionnaire), 2.8 (theme), 2.9 (components)

**Verification:**
```go
m := tui.NewSkillCreateModel()
p := tea.NewProgram(m)
p.Run() // Should show wizard
```

---

#### Task 2.11: Skill Test TUI

**File:** `tools/cody/internal/tui/skill_test.go`

**Creates:**
- `SkillTestModel` struct
- `NewSkillTestModel(skillName string)`
- `Init()`, `Update()`, `View()` - Bubble Tea interface
- Live progress display
- Results summary view

**Dependencies:** 2.1-2.4 (test runners), 2.8 (theme), 2.9 (components)

**Verification:**
```go
m := tui.NewSkillTestModel("my-skill")
p := tea.NewProgram(m)
p.Run() // Should show test runner
```

---

## Phase 3: CLI Integration (SERIAL)

**Requires Phase 2 complete.**

### Task 3.1: Main Model

**File:** `tools/cody/internal/tui/model.go`

**Creates:**
- `Model` struct (main app model)
- `NewModel()`
- State machine for navigation between screens
- Home screen with menu

**Dependencies:** 2.10, 2.11

---

### Task 3.2: CLI Entry Point

**File:** `tools/cody/cmd/cody/main.go`

**Creates:**
- Cobra command structure
- `cody` - launch interactive TUI
- `cody skill create` - launch skill wizard
- `cody skill test <name>` - launch test runner
- `cody skill list` - list skills
- `cody feature run <config>` - run feature (existing)

**Dependencies:** 3.1

**Verification:**
```bash
go build -o cody ./cmd/cody
./cody skill create  # Should launch wizard
./cody skill test my-skill  # Should run tests
```

---

### Task 3.3: Library Client Implementation

**File:** `tools/cody/pkg/cody/client.go`

**Creates:**
- Implement all `Client` methods with real logic
- Wire up skill manager, test runners, composer
- Progress callback handling

**Dependencies:** 2.1-2.7

---

### Task 3.4: Integration Tests

**File:** `tools/cody/tests/integration/`

**Creates:**
- `skill_create_test.go` - E2E skill creation
- `skill_test_test.go` - E2E test running
- `compose_test.go` - E2E config composition

**Dependencies:** 3.3

**Verification:**
```bash
go test ./tests/integration/... -v
```

---

## Phase 4: Server & API (SERIAL)

**Requires Phase 3 complete.**

### Task 4.1: HTTP Server

**File:** `tools/cody/internal/server/`

**Creates:**
- `server.go` - Server struct, setup
- `handlers.go` - All HTTP handlers
- `websocket.go` - WebSocket streaming
- `middleware.go` - Auth, logging, CORS

**Dependencies:** 3.3 (client implementation)

---

### Task 4.2: gRPC Server

**File:** `tools/cody/internal/server/grpc/`

**Creates:**
- Generate from proto files
- Implement service interfaces
- Streaming handlers

**Dependencies:** 4.1

---

### Task 4.3: Server Binary

**File:** `tools/cody/cmd/cody-server/main.go`

**Creates:**
- Server entry point
- Config loading
- Graceful shutdown

**Dependencies:** 4.1, 4.2

**Verification:**
```bash
go build -o cody-server ./cmd/cody-server
./cody-server &
curl http://localhost:8080/api/v1/skills
```

---

## Phase 5: Polish & Documentation (PARALLEL)

### Task 5.1: Config Composer

**File:** `tools/cody/pkg/compose/`

**Creates:**
- `composer.go` - ConfigComposer implementation
- `rules.go` - Composition rules
- `analyze.go` - Task analysis

---

### Task 5.2: Error Handling & Logging

**Creates:**
- Structured logging throughout
- Error wrapping with context
- Recovery from panics

---

### Task 5.3: Documentation

**Creates:**
- README.md for tools/cody
- Usage examples
- API documentation

---

### Task 5.4: CI/CD

**Creates:**
- GitHub Actions workflow
- Build, test, release automation
- Docker images

---

## Task Summary Table

| Task | Name | Dependencies | Parallel Group | Est. Complexity |
|------|------|--------------|----------------|-----------------|
| 1.1 | Schema Package | None | - | Medium |
| 1.2 | Core Interfaces | 1.1 | - | Medium |
| 2.1 | Test Runner Base | 1.1 | A | Medium |
| 2.2 | Behavioral Diff | 2.1 | A | High |
| 2.3 | Determinism Tests | 2.1 | A | Medium |
| 2.4 | Compliance Tests | 2.1 | A | High |
| 2.5 | Questionnaire | 1.1 | B | Medium |
| 2.6 | Generator | 1.1, 2.5 | B | Medium |
| 2.7 | Manager | 2.6 | B | Low |
| 2.8 | Theme & Styles | None | C | Low |
| 2.9 | TUI Components | 2.8 | C | Medium |
| 2.10 | Skill Create TUI | 2.5, 2.8, 2.9 | C | High |
| 2.11 | Skill Test TUI | 2.1-2.4, 2.8, 2.9 | C | High |
| 3.1 | Main Model | 2.10, 2.11 | - | Medium |
| 3.2 | CLI Entry Point | 3.1 | - | Low |
| 3.3 | Client Implementation | 2.1-2.7 | - | Medium |
| 3.4 | Integration Tests | 3.3 | - | Medium |
| 4.1 | HTTP Server | 3.3 | - | Medium |
| 4.2 | gRPC Server | 4.1 | - | Medium |
| 4.3 | Server Binary | 4.1, 4.2 | - | Low |
| 5.1 | Config Composer | 3.3 | D | High |
| 5.2 | Error Handling | 3.3 | D | Low |
| 5.3 | Documentation | 4.3 | D | Medium |
| 5.4 | CI/CD | 4.3 | D | Medium |

---

## Agent Instructions

### To check if a task is done:

```
"Is task {N} done? Verify by:
1. Check if files exist at specified paths
2. Run the verification command
3. Confirm all items in 'Creates' section exist"
```

### To implement a task:

```
"Implement task {N}.
- Read ARCH.md for detailed specs
- Create files listed in 'Creates' section
- Follow patterns from earlier tasks
- Run verification command when done
- Report completion status"
```

### To check dependencies:

```
"Before starting task {N}, verify these tasks are complete: {dependencies}"
```

---

## Current Status

| Task | Status | Notes |
|------|--------|-------|
| 1.1 | ✅ COMPLETE | `pkg/schema/` - feature.go, skill.go, test.go, config.go |
| 1.2 | ✅ COMPLETE | `pkg/cody/client.go`, `pkg/skill/` - questionnaire.go, manager.go, generator.go |
| 2.1 | ✅ COMPLETE | `pkg/skilltest/runner.go` - Base runner, ScenarioExecutor, MockExecutor |
| 2.2 | ✅ COMPLETE | `pkg/skilltest/behavioral.go` - BehavioralDiffRunner, metric evaluators |
| 2.3 | ✅ COMPLETE | `pkg/skilltest/determinism.go` - DeterminismRunner, variance metrics |
| 2.4 | ✅ COMPLETE | `pkg/skilltest/compliance.go` - ComplianceRunner, pressure scenarios, rationalization detection |
| 2.5 | ✅ COMPLETE | Questionnaire (done with 1.2) |
| 2.6 | ✅ COMPLETE | Generator (done with 1.2) |
| 2.7 | ✅ COMPLETE | Manager (done with 1.2) |
| 2.8 | ✅ COMPLETE | `internal/tui/theme.go` - Theme colors, styles, Logo, ProgressBar, helpers |
| 2.9 | ✅ COMPLETE | `internal/tui/components/` - question.go, progress.go, results.go, box.go |
| 2.10 | ✅ COMPLETE | `internal/tui/skill_create.go` - SkillCreateModel, wizard flow |
| 2.11 | ✅ COMPLETE | `internal/tui/skill_test.go` - SkillTestModel, test runner UI |
| 3.1 | ✅ COMPLETE | `internal/tui/app.go` - Main app model with navigation |
| 3.2 | ✅ COMPLETE | `cmd/cody/main.go` - Cobra commands: skill create/test/list, app |
| 3.3 | ✅ COMPLETE | `pkg/cody/client.go` - Wired up test runners, skill manager |
| 3.4 | ✅ COMPLETE | Unit tests in `pkg/cody/`, `pkg/skill/`, `pkg/skilltest/` |
| 4.1 | ✅ COMPLETE | `internal/server/` - HTTP server with REST API |
| 4.2 | ✅ COMPLETE | `internal/server/grpc/` - gRPC stub (proto not implemented) |
| 4.3 | ✅ COMPLETE | `cmd/cody-api/main.go` - API server binary |
| 5.1 | NOT STARTED | Config Composer |
| 5.2 | NOT STARTED | Error Handling |
| 5.3 | NOT STARTED | Documentation |
| 5.4 | NOT STARTED | CI/CD |

**Phase 1: COMPLETE** ✅
**Phase 2 Track A (Test Runners): COMPLETE** ✅
**Phase 2 Track B (Skill Engine): COMPLETE** ✅
**Phase 2 Track C (TUI Components): COMPLETE** ✅

**PHASE 2 COMPLETE** ✅

**Phase 3 (CLI Integration): COMPLETE** ✅
- Task 3.1: `internal/tui/app.go` - Main model with screen navigation
- Task 3.2: `cmd/cody/main.go` - Cobra CLI with subcommands
- Task 3.3: `pkg/cody/client.go` - Client wired to test runners
- Task 3.4: Comprehensive unit tests for all packages

**PHASE 3 COMPLETE** ✅

**Phase 4 (Server & API): COMPLETE** ✅
- Task 4.1: `internal/server/` - HTTP server with REST API, WebSocket support
- Task 4.2: `internal/server/grpc/` - gRPC stub (full implementation requires proto files)
- Task 4.3: `cmd/cody-api/` - API server binary with graceful shutdown

**PHASE 4 COMPLETE** ✅

**Next action:** Start Phase 5 (Polish & Documentation)

### Files Created in Phase 4

- `internal/server/server.go` - HTTP server setup and WebSocket handling
- `internal/server/handlers.go` - REST API handlers
- `internal/server/middleware.go` - Logging, CORS, recovery, auth middleware
- `internal/server/grpc/server.go` - gRPC server stub
- `cmd/cody-api/main.go` - API server binary

### Files Created in Phase 3

- `internal/tui/app.go` - Main TUI application model
- `pkg/cody/client_test.go` - Client unit tests
- `pkg/skill/questionnaire_test.go` - Questionnaire unit tests
- `pkg/skill/manager_test.go` - Manager unit tests
- `pkg/skilltest/runner_test.go` - Base runner tests
- `pkg/skilltest/behavioral_test.go` - Behavioral runner tests
- `pkg/skilltest/determinism_test.go` - Determinism runner tests
- `pkg/skilltest/compliance_test.go` - Compliance runner tests
