# Claude Code Mastery Architecture

## 1. Vision & Goals

### Vision

Build a **verified, testable system** for mastering Claude Code that:
1. Documents every feature with deep understanding of capabilities and constraints
2. Encodes best practices as testable skills
3. Proves skills work through automated Go tests
4. Enables dynamic composition of optimal Claude Code configurations

### Goals

**G1: Complete Feature Mastery**
- Inventory every Claude Code feature (skills, hooks, MCP, plugins, LSP, rules, agents, plan mode)
- Document when to use each, when NOT to use each, and how they interact
- Identify gaps in current understanding

**G2: Testable Skills**
- Every skill has associated Go tests
- Three test types: behavioral diff, determinism, compliance
- Tests run against real Claude instances (not mocks)

**G3: Automation-Ready Foundation**
- Schemas are Go structs that Cody can import directly
- Feature inventory is structured for programmatic consumption
- Decision logic is explicit, not buried in prose

**G4: Sellable Product**
- Different tiers: solo devs, teams, AI-first shops
- Clear value prop at each tier
- Proof that the system works (test results, metrics)

### Success Criteria

| Metric | Target |
|--------|--------|
| Feature coverage | 100% of Claude Code features documented |
| Skill test coverage | Every skill has all 3 test types |
| Behavioral diff accuracy | >90% detection of skill impact |
| Determinism improvement | Skills reduce output variance by >50% |
| Compliance under pressure | >95% rule-following in adversarial scenarios |

---

## 2. Schemas

These Go structs are the foundation. Cody imports these directly.

### 2.1 Feature Schema

```go
package schema

// FeatureType categorizes Claude Code features
type FeatureType string

const (
    FeatureTypeSkill     FeatureType = "skill"
    FeatureTypeHook      FeatureType = "hook"
    FeatureTypeMCP       FeatureType = "mcp_server"
    FeatureTypePlugin    FeatureType = "plugin"
    FeatureTypeLSP       FeatureType = "lsp_server"
    FeatureTypeRule      FeatureType = "rule"
    FeatureTypeAgent     FeatureType = "agent"
    FeatureTypePlanMode  FeatureType = "plan_mode"
)

// Feature represents a Claude Code capability
type Feature struct {
    // Identity
    Name        string      `json:"name" yaml:"name"`
    Type        FeatureType `json:"type" yaml:"type"`
    Description string      `json:"description" yaml:"description"`

    // Capabilities
    WhatItDoes     string   `json:"what_it_does" yaml:"what_it_does"`
    WhenToUse      []string `json:"when_to_use" yaml:"when_to_use"`
    WhenNotToUse   []string `json:"when_not_to_use" yaml:"when_not_to_use"`
    Capabilities   []string `json:"capabilities" yaml:"capabilities"`
    Limitations    []string `json:"limitations" yaml:"limitations"`

    // Integration
    DependsOn      []string `json:"depends_on" yaml:"depends_on"`       // Other features this requires
    ConflictsWith  []string `json:"conflicts_with" yaml:"conflicts_with"` // Features that clash
    EnhancedBy     []string `json:"enhanced_by" yaml:"enhanced_by"`     // Features that make this better

    // Configuration
    ConfigSchema   string   `json:"config_schema" yaml:"config_schema"` // JSON Schema for config
    ExampleConfig  string   `json:"example_config" yaml:"example_config"`

    // Testing
    TestStrategies []TestStrategy `json:"test_strategies" yaml:"test_strategies"`
}

// TestStrategy defines how to test a feature
type TestStrategy struct {
    Type        TestType `json:"type" yaml:"type"`
    Description string   `json:"description" yaml:"description"`
    Template    string   `json:"template" yaml:"template"` // Go test template
}

type TestType string

const (
    TestTypeBehavioralDiff TestType = "behavioral_diff"
    TestTypeDeterminism    TestType = "determinism"
    TestTypeCompliance     TestType = "compliance"
)
```

### 2.2 Skill Schema

```go
package schema

// Skill represents a testable Claude Code skill
type Skill struct {
    // Identity (maps to SKILL.md frontmatter)
    Name        string `json:"name" yaml:"name"`
    Description string `json:"description" yaml:"description"` // "Use when..." format

    // Content
    Overview      string   `json:"overview" yaml:"overview"`
    WhenToUse     []string `json:"when_to_use" yaml:"when_to_use"`
    WhenNotToUse  []string `json:"when_not_to_use" yaml:"when_not_to_use"`
    CorePattern   string   `json:"core_pattern" yaml:"core_pattern"`     // Before/after or key technique
    QuickRef      string   `json:"quick_ref" yaml:"quick_ref"`           // Scannable reference
    CommonMistakes []Mistake `json:"common_mistakes" yaml:"common_mistakes"`

    // Testing (required for all skills)
    Tests         SkillTests `json:"tests" yaml:"tests"`

    // Resources
    Scripts    []string `json:"scripts" yaml:"scripts"`       // Paths to executable scripts
    References []string `json:"references" yaml:"references"` // Paths to reference docs
    Assets     []string `json:"assets" yaml:"assets"`         // Paths to templates/files

    // Metadata
    Keywords   []string `json:"keywords" yaml:"keywords"` // CSO: search optimization
    SkillType  SkillType `json:"skill_type" yaml:"skill_type"`
}

type SkillType string

const (
    SkillTypeTechnique SkillType = "technique" // Concrete method with steps
    SkillTypePattern   SkillType = "pattern"   // Way of thinking about problems
    SkillTypeReference SkillType = "reference" // API docs, syntax guides
    SkillTypeDiscipline SkillType = "discipline" // Rules/requirements to follow
)

type Mistake struct {
    What string `json:"what" yaml:"what"`
    Fix  string `json:"fix" yaml:"fix"`
}

// SkillTests contains all test definitions for a skill
type SkillTests struct {
    BehavioralDiff []BehavioralDiffTest `json:"behavioral_diff" yaml:"behavioral_diff"`
    Determinism    []DeterminismTest    `json:"determinism" yaml:"determinism"`
    Compliance     []ComplianceTest     `json:"compliance" yaml:"compliance"`
}
```

### 2.3 Test Schemas

```go
package schema

// BehavioralDiffTest measures if a skill changes agent behavior
type BehavioralDiffTest struct {
    Name        string `json:"name" yaml:"name"`
    Description string `json:"description" yaml:"description"`

    // The scenario to run
    Prompt      string `json:"prompt" yaml:"prompt"`
    Context     string `json:"context" yaml:"context"` // Additional context/files

    // What to measure
    Metric      BehaviorMetric `json:"metric" yaml:"metric"`

    // Expected outcomes
    WithoutSkill ExpectedBehavior `json:"without_skill" yaml:"without_skill"`
    WithSkill    ExpectedBehavior `json:"with_skill" yaml:"with_skill"`

    // Thresholds
    MinDiffPercent float64 `json:"min_diff_percent" yaml:"min_diff_percent"` // Minimum behavior change required
}

type BehaviorMetric string

const (
    MetricTestsWrittenFirst    BehaviorMetric = "tests_written_first"     // TDD compliance
    MetricPlanBeforeCode       BehaviorMetric = "plan_before_code"        // Planning behavior
    MetricErrorHandling        BehaviorMetric = "error_handling_present"  // Defensive coding
    MetricCodeReviewRequested  BehaviorMetric = "code_review_requested"   // Review discipline
    MetricCommitMessageQuality BehaviorMetric = "commit_message_quality"  // Commit hygiene
    MetricCustom               BehaviorMetric = "custom"                  // User-defined metric
)

type ExpectedBehavior struct {
    Rate       float64 `json:"rate" yaml:"rate"`             // Expected % of time behavior occurs
    Tolerance  float64 `json:"tolerance" yaml:"tolerance"`   // Acceptable variance
    CustomEval string  `json:"custom_eval" yaml:"custom_eval"` // Go code for custom evaluation
}

// DeterminismTest measures if a skill reduces output variance
type DeterminismTest struct {
    Name        string `json:"name" yaml:"name"`
    Description string `json:"description" yaml:"description"`

    // The scenario to run multiple times
    Prompt      string `json:"prompt" yaml:"prompt"`
    Context     string `json:"context" yaml:"context"`
    Iterations  int    `json:"iterations" yaml:"iterations"` // How many times to run (default: 10)

    // What to measure for consistency
    ConsistencyChecks []ConsistencyCheck `json:"consistency_checks" yaml:"consistency_checks"`

    // Thresholds
    MaxVariancePercent float64 `json:"max_variance_percent" yaml:"max_variance_percent"`
}

type ConsistencyCheck struct {
    Name       string          `json:"name" yaml:"name"`
    Type       ConsistencyType `json:"type" yaml:"type"`
    Target     string          `json:"target" yaml:"target"`     // What to check (file, output section, etc.)
    CustomEval string          `json:"custom_eval" yaml:"custom_eval"` // Go code for custom evaluation
}

type ConsistencyType string

const (
    ConsistencyTypeStructure   ConsistencyType = "structure"    // Same files created, same functions defined
    ConsistencyTypeApproach    ConsistencyType = "approach"     // Same algorithm/pattern chosen
    ConsistencyTypeNaming      ConsistencyType = "naming"       // Consistent naming conventions
    ConsistencyTypeErrorCases  ConsistencyType = "error_cases"  // Same edge cases handled
    ConsistencyTypeCustom      ConsistencyType = "custom"
)

// ComplianceTest measures if a skill is followed under pressure
type ComplianceTest struct {
    Name        string `json:"name" yaml:"name"`
    Description string `json:"description" yaml:"description"`

    // The pressure scenario
    Scenario    PressureScenario `json:"scenario" yaml:"scenario"`

    // The rule that should be followed
    Rule        string   `json:"rule" yaml:"rule"`
    RuleCheck   string   `json:"rule_check" yaml:"rule_check"` // Go code to verify compliance

    // Expected rationalizations (to detect violations)
    KnownRationalizations []string `json:"known_rationalizations" yaml:"known_rationalizations"`

    // Thresholds
    MinComplianceRate float64 `json:"min_compliance_rate" yaml:"min_compliance_rate"` // e.g., 0.95 = 95%
}

type PressureScenario struct {
    Setup       string         `json:"setup" yaml:"setup"`           // Initial context
    Prompt      string         `json:"prompt" yaml:"prompt"`         // The task
    Pressures   []PressureType `json:"pressures" yaml:"pressures"`   // Types of pressure applied
    Options     []string       `json:"options" yaml:"options"`       // Explicit choices (A/B/C)
    CorrectOption string       `json:"correct_option" yaml:"correct_option"` // Which option follows the rule
}

type PressureType string

const (
    PressureTypeTime       PressureType = "time"       // Deadline, urgency
    PressureTypeSunkCost   PressureType = "sunk_cost"  // Already invested effort
    PressureTypeAuthority  PressureType = "authority"  // Senior/manager says skip it
    PressureTypeEconomic   PressureType = "economic"   // Money/job at stake
    PressureTypeExhaustion PressureType = "exhaustion" // End of day, tired
    PressureTypeSocial     PressureType = "social"     // Looking dogmatic/inflexible
    PressureTypePragmatic  PressureType = "pragmatic"  // "Being practical"
)
```

### 2.4 Configuration Schema

```go
package schema

// ClaudeCodeConfig represents an optimal Claude Code setup for a task
type ClaudeCodeConfig struct {
    // Identity
    Name        string `json:"name" yaml:"name"`
    Description string `json:"description" yaml:"description"`

    // Task matching
    TaskPatterns []TaskPattern `json:"task_patterns" yaml:"task_patterns"`

    // Components
    Skills      []SkillRef  `json:"skills" yaml:"skills"`
    Hooks       []HookDef   `json:"hooks" yaml:"hooks"`
    MCPServers  []MCPServer `json:"mcp_servers" yaml:"mcp_servers"`
    Plugins     []PluginRef `json:"plugins" yaml:"plugins"`
    LSPServers  []LSPServer `json:"lsp_servers" yaml:"lsp_servers"`
    Rules       []RuleDef   `json:"rules" yaml:"rules"`
    Agents      []AgentDef  `json:"agents" yaml:"agents"`
    PlanMode    *PlanModeConfig `json:"plan_mode" yaml:"plan_mode"`
}

type TaskPattern struct {
    Keywords    []string `json:"keywords" yaml:"keywords"`
    FileTypes   []string `json:"file_types" yaml:"file_types"`   // e.g., ".go", ".ts"
    ProjectType string   `json:"project_type" yaml:"project_type"` // e.g., "web-api", "cli", "library"
    Complexity  string   `json:"complexity" yaml:"complexity"`     // "simple", "medium", "complex"
}

type SkillRef struct {
    Name     string `json:"name" yaml:"name"`
    Required bool   `json:"required" yaml:"required"` // Must load vs optional
    Priority int    `json:"priority" yaml:"priority"` // Load order
}

type HookDef struct {
    Event   HookEvent `json:"event" yaml:"event"`
    Matcher string    `json:"matcher" yaml:"matcher"` // Tool/file pattern to match
    Command string    `json:"command" yaml:"command"`
    Timeout int       `json:"timeout" yaml:"timeout"` // Seconds
}

type HookEvent string

const (
    HookEventPreToolUse   HookEvent = "PreToolUse"
    HookEventPostToolUse  HookEvent = "PostToolUse"
    HookEventNotification HookEvent = "Notification"
    HookEventStop         HookEvent = "Stop"
)

type MCPServer struct {
    Name        string            `json:"name" yaml:"name"`
    Command     string            `json:"command" yaml:"command"`
    Args        []string          `json:"args" yaml:"args"`
    Env         map[string]string `json:"env" yaml:"env"`
    Description string            `json:"description" yaml:"description"`
}

type PluginRef struct {
    Name    string `json:"name" yaml:"name"`
    Version string `json:"version" yaml:"version"`
    Config  string `json:"config" yaml:"config"` // JSON config for plugin
}

type LSPServer struct {
    Name       string   `json:"name" yaml:"name"`
    Language   string   `json:"language" yaml:"language"`
    Command    string   `json:"command" yaml:"command"`
    Args       []string `json:"args" yaml:"args"`
    Extensions []string `json:"extensions" yaml:"extensions"` // File extensions to handle
}

type RuleDef struct {
    Type    RuleType `json:"type" yaml:"type"`
    Content string   `json:"content" yaml:"content"`
    Scope   string   `json:"scope" yaml:"scope"` // "global", "project", "directory"
}

type RuleType string

const (
    RuleTypeInstruction RuleType = "instruction" // Tell Claude how to behave
    RuleTypeConstraint  RuleType = "constraint"  // Forbid certain actions
    RuleTypePreference  RuleType = "preference"  // Prefer X over Y
)

type AgentDef struct {
    Name        string   `json:"name" yaml:"name"`
    Description string   `json:"description" yaml:"description"`
    Tools       []string `json:"tools" yaml:"tools"`     // Tools agent can use
    Model       string   `json:"model" yaml:"model"`     // "sonnet", "opus", "haiku"
    Color       string   `json:"color" yaml:"color"`
}

type PlanModeConfig struct {
    Enabled         bool     `json:"enabled" yaml:"enabled"`
    RequireApproval bool     `json:"require_approval" yaml:"require_approval"`
    PlanFile        string   `json:"plan_file" yaml:"plan_file"` // Where to write plans
    AllowedPrompts  []string `json:"allowed_prompts" yaml:"allowed_prompts"` // Pre-approved actions
}
```

---

## 3. Feature Inventory

### 3.1 Skills

```yaml
name: skills
type: skill
description: "Inject domain knowledge, workflows, and specialized capabilities into Claude's context"

what_it_does: |
  Skills are markdown files (SKILL.md) with YAML frontmatter that get loaded into
  Claude's context when triggered. They provide:
  - Specialized workflows (TDD, debugging, code review)
  - Domain knowledge (APIs, frameworks, company processes)
  - Tool integrations (file formats, external services)
  - Behavioral constraints (rules to follow)

when_to_use:
  - "Encoding repeatable workflows that Claude should follow consistently"
  - "Providing domain-specific knowledge Claude doesn't have"
  - "Enforcing discipline (TDD, review processes, commit hygiene)"
  - "Teaching Claude how to use specific tools or APIs"
  - "Standardizing behavior across team members"

when_not_to_use:
  - "One-off instructions (just put in prompt)"
  - "Project-specific conventions (use CLAUDE.md rules instead)"
  - "Information Claude already knows well"
  - "Mechanical checks that can be automated (use hooks instead)"

capabilities:
  - "Load on-demand based on task detection"
  - "Include bundled resources (scripts, references, assets)"
  - "Progressive disclosure (metadata → body → resources)"
  - "Searchable by keywords and description"
  - "Can reference other skills"

limitations:
  - "Context window cost (every skill consumes tokens)"
  - "Can't enforce behavior (only instruct)"
  - "No runtime verification (skill says X, Claude might do Y)"
  - "Description quality determines discoverability"

depends_on: []

conflicts_with:
  - "Contradictory rules in CLAUDE.md"
  - "Other skills with conflicting instructions"

enhanced_by:
  - "hooks (verify skill compliance)"
  - "mcp_server (provide tools skill references)"
  - "plan_mode (ensure skill is consulted before action)"

test_strategies:
  - type: behavioral_diff
    description: "Measure if skill changes agent behavior in expected direction"
    template: "See Section 4.1"
  - type: determinism
    description: "Measure if skill reduces variance in approach"
    template: "See Section 4.2"
  - type: compliance
    description: "Measure if skill rules are followed under pressure"
    template: "See Section 4.3"
```

### 3.2 Hooks

```yaml
name: hooks
type: hook
description: "Execute shell commands in response to Claude Code events"

what_it_does: |
  Hooks are shell commands that run automatically when specific events occur:
  - PreToolUse: Before Claude uses a tool (can block or modify)
  - PostToolUse: After Claude uses a tool (can validate or react)
  - Notification: When Claude sends a notification
  - Stop: When Claude stops (session end, interrupt)

  Hooks receive JSON input (event details) and can output JSON to influence behavior.

when_to_use:
  - "Enforcing constraints that must be verified (not just instructed)"
  - "Automating repetitive validations (lint, format, security scan)"
  - "Integrating with external systems (notifications, logging)"
  - "Blocking dangerous operations (production deploys, destructive commands)"
  - "Adding guardrails that Claude can't rationalize around"

when_not_to_use:
  - "Teaching Claude how to do something (use skills)"
  - "Complex logic that needs context (hooks are stateless)"
  - "Things that should be optional (hooks always run)"
  - "Performance-critical paths (hooks add latency)"

capabilities:
  - "Block tool calls (exit non-zero with reason)"
  - "Modify tool inputs (output modified JSON)"
  - "Run arbitrary shell commands"
  - "Access environment variables"
  - "Match specific tools or file patterns"

limitations:
  - "No access to conversation context"
  - "Synchronous (blocks until complete)"
  - "Shell-based (complex logic is awkward)"
  - "Can't modify Claude's understanding, only block/allow actions"

depends_on: []

conflicts_with:
  - "Hooks that block each other's requirements"

enhanced_by:
  - "skills (instruct behavior that hooks verify)"
  - "mcp_server (hooks can call MCP tools)"

test_strategies:
  - type: behavioral_diff
    description: "Verify hook blocks/allows expected tool calls"
    template: |
      func TestHookBlocks(t *testing.T) {
          // Run scenario that should trigger hook
          // Verify tool call was blocked/allowed
          // Check hook output matches expected
      }
  - type: compliance
    description: "Verify Claude adapts when hook blocks action"
    template: |
      func TestHookComplianceAdaptation(t *testing.T) {
          // Set up hook that blocks certain action
          // Give Claude task that would trigger blocked action
          // Verify Claude finds alternative approach
      }
```

### 3.3 MCP Servers

```yaml
name: mcp_servers
type: mcp_server
description: "External tool integrations via Model Context Protocol"

what_it_does: |
  MCP servers expose tools and resources to Claude via a standardized protocol:
  - Tools: Functions Claude can call (database queries, API calls, browser automation)
  - Resources: Data Claude can read (file systems, documentation)
  - Prompts: Pre-defined prompt templates

  Claude Code connects to MCP servers and makes their capabilities available as tools.

when_to_use:
  - "Integrating external services (databases, APIs, browsers)"
  - "Providing tools that need runtime execution"
  - "Accessing real-time data (not static context)"
  - "Complex operations that can't be done with bash"
  - "Sandboxed tool execution with defined interfaces"

when_not_to_use:
  - "Static information (just put in skill/context)"
  - "Simple bash commands (overhead not worth it)"
  - "Operations Claude can do natively"
  - "When latency is critical (MCP adds round-trip)"

capabilities:
  - "Arbitrary tool definitions"
  - "Typed inputs and outputs"
  - "Resource access with URI schemes"
  - "Can run as separate process (isolation)"
  - "Standard protocol (reusable across clients)"

limitations:
  - "Requires running server process"
  - "Network/IPC latency"
  - "Tool definitions must be complete (Claude can't infer)"
  - "Error handling across process boundary"

depends_on: []

conflicts_with:
  - "Multiple servers providing same tool name"

enhanced_by:
  - "skills (teach Claude when/how to use MCP tools)"
  - "hooks (validate MCP tool usage)"

test_strategies:
  - type: behavioral_diff
    description: "Verify MCP tools are used when appropriate"
    template: |
      func TestMCPToolSelection(t *testing.T) {
          // Set up MCP server with specific tools
          // Give task that should use those tools
          // Verify correct tool was called with correct args
      }
  - type: determinism
    description: "Verify consistent tool usage patterns"
    template: |
      func TestMCPToolConsistency(t *testing.T) {
          // Run same task N times
          // Verify same MCP tools used each time
          // Check args are consistent
      }
```

### 3.4 Plugins

```yaml
name: plugins
type: plugin
description: "Bundles of skills, agents, hooks, and configuration"

what_it_does: |
  Plugins are installable packages that extend Claude Code with:
  - Skills: Domain knowledge and workflows
  - Agents: Specialized subagent definitions
  - Hooks: Automated validations
  - Configuration: MCP servers, settings

  Plugins can be installed from marketplaces or local directories.

when_to_use:
  - "Distributing a coherent set of capabilities"
  - "Team standardization (everyone uses same plugin)"
  - "Complex integrations requiring multiple components"
  - "Third-party extensions"

when_not_to_use:
  - "Single skill or hook (just create the file directly)"
  - "Project-specific customization (use local config)"
  - "Experimental/unstable capabilities"

capabilities:
  - "Version management"
  - "Marketplace distribution"
  - "Bundled dependencies"
  - "Configuration via JSON"

limitations:
  - "Opaque (harder to customize than individual files)"
  - "Update management"
  - "Potential conflicts between plugins"

depends_on:
  - "skills (plugins contain skills)"
  - "hooks (plugins contain hooks)"
  - "agent (plugins contain agents)"

test_strategies:
  - type: behavioral_diff
    description: "Verify plugin installation changes behavior"
    template: |
      func TestPluginBehavior(t *testing.T) {
          // Run task without plugin
          // Install plugin
          // Run same task
          // Verify behavior changed as expected
      }
```

### 3.5 LSP Servers

```yaml
name: lsp_servers
type: lsp_server
description: "Language intelligence via Language Server Protocol"

what_it_does: |
  LSP servers provide language-aware features:
  - Diagnostics: Errors, warnings, suggestions
  - Completions: Code suggestions
  - Hover: Type information
  - Go to definition, find references

  Claude Code can use LSP diagnostics to understand code issues.

when_to_use:
  - "Getting accurate type errors and diagnostics"
  - "Languages with complex type systems (TypeScript, Go, Rust)"
  - "When Claude needs to understand existing code structure"
  - "Catching errors before running tests"

when_not_to_use:
  - "Dynamic languages with weak LSP support"
  - "Small scripts (overhead not worth it)"
  - "When fast iteration is priority (LSP can be slow)"

capabilities:
  - "Real-time diagnostics"
  - "Type information"
  - "Cross-file analysis"
  - "Refactoring support"

limitations:
  - "Language-specific (need separate server per language)"
  - "Setup complexity"
  - "Memory and CPU usage"
  - "Some languages have weak LSP implementations"

depends_on: []

enhanced_by:
  - "skills (teach Claude how to interpret LSP diagnostics)"
  - "hooks (auto-fix certain diagnostic types)"

test_strategies:
  - type: behavioral_diff
    description: "Verify Claude uses LSP diagnostics"
    template: |
      func TestLSPDiagnosticUsage(t *testing.T) {
          // Create file with type errors
          // Run with LSP enabled
          // Verify Claude addresses LSP-reported issues
      }
```

### 3.6 Rules (CLAUDE.md)

```yaml
name: rules
type: rule
description: "Project-level behavioral instructions in CLAUDE.md files"

what_it_does: |
  CLAUDE.md files contain project-specific instructions:
  - Coding conventions
  - Project structure
  - Forbidden actions
  - Preferred approaches

  Files are loaded based on scope:
  - ~/.claude/CLAUDE.md (global)
  - /project/CLAUDE.md (project root)
  - /project/subdir/CLAUDE.md (directory-specific)

when_to_use:
  - "Project-specific conventions"
  - "Forbidden patterns or practices"
  - "Preferred libraries or approaches"
  - "Team agreements"
  - "Instructions that don't need testing/verification"

when_not_to_use:
  - "Reusable knowledge (use skills instead)"
  - "Enforceable constraints (use hooks instead)"
  - "Complex workflows (use skills instead)"

capabilities:
  - "Hierarchical (global → project → directory)"
  - "Simple markdown format"
  - "Always loaded (no triggering needed)"
  - "Can reference skills to load"

limitations:
  - "No verification (just instructions)"
  - "Can be ignored by Claude"
  - "No structure/schema"
  - "Hard to test effectiveness"

depends_on: []

conflicts_with:
  - "Skills with contradictory instructions"

test_strategies:
  - type: compliance
    description: "Verify CLAUDE.md rules are followed"
    template: |
      func TestClaudeMdCompliance(t *testing.T) {
          // Set up CLAUDE.md with specific rule
          // Give task that could violate rule
          // Verify rule is followed
      }
```

### 3.7 Agents

```yaml
name: agents
type: agent
description: "Specialized subagent definitions with constrained tools and focus"

what_it_does: |
  Agents are subagent configurations that:
  - Have specific tool access (not full toolset)
  - Have focused system prompts
  - Can be spawned via Task tool
  - Run with their own context

  Used for: code review, exploration, testing, specialized tasks

when_to_use:
  - "Tasks needing focused context (not main conversation)"
  - "Parallel work on independent subtasks"
  - "Specialized roles (reviewer, explorer, tester)"
  - "Limiting tool access for safety"

when_not_to_use:
  - "Simple tasks in current context"
  - "Tasks needing full conversation history"
  - "Interactive back-and-forth with user"

capabilities:
  - "Tool restriction"
  - "Custom system prompts"
  - "Parallel execution"
  - "Isolated context"

limitations:
  - "Context handoff overhead"
  - "No direct user interaction"
  - "Results must be summarized"

depends_on:
  - "skills (agents can load skills)"

test_strategies:
  - type: behavioral_diff
    description: "Verify agent constraints are respected"
    template: |
      func TestAgentToolConstraints(t *testing.T) {
          // Define agent with limited tools
          // Give task requiring disallowed tool
          // Verify agent doesn't use forbidden tool
      }
```

### 3.8 Plan Mode

```yaml
name: plan_mode
type: plan_mode
description: "Force planning before execution"

what_it_does: |
  Plan mode requires Claude to:
  1. Research and understand the task
  2. Write a plan document
  3. Get user approval
  4. Only then execute

  Triggered via EnterPlanMode tool or configured to be default.

when_to_use:
  - "Complex multi-step implementations"
  - "Tasks with multiple valid approaches"
  - "When alignment before action is critical"
  - "Architectural decisions"

when_not_to_use:
  - "Simple, obvious tasks"
  - "Research/exploration (no execution needed)"
  - "Emergency fixes"

capabilities:
  - "Forces deliberation"
  - "User approval gate"
  - "Plan document artifact"
  - "Can request permissions upfront"

limitations:
  - "Slower for simple tasks"
  - "Requires user interaction for approval"
  - "Plan != execution (plans can be wrong)"

test_strategies:
  - type: compliance
    description: "Verify plan mode is entered when configured"
    template: |
      func TestPlanModeEnforcement(t *testing.T) {
          // Configure plan mode as required
          // Give complex task
          // Verify Claude enters plan mode before coding
      }
```

---

## 4. Testing Methodology

### 4.1 Behavioral Diff Tests

**Purpose:** Measure if a skill/feature changes agent behavior in the expected direction.

**Method:**
1. Run scenario WITHOUT skill
2. Run scenario WITH skill
3. Measure specific behavior metric
4. Assert difference exceeds threshold

```go
package skilltest

import (
    "context"
    "testing"

    "github.com/anthropics/claude-code-sdk-go"
)

// BehavioralDiffRunner executes behavioral diff tests
type BehavioralDiffRunner struct {
    client  *claude.Client
    workdir string
}

func (r *BehavioralDiffRunner) Run(t *testing.T, test BehavioralDiffTest) {
    // Phase 1: Run WITHOUT skill
    baselineResults := make([]BehaviorObservation, test.Iterations)
    for i := 0; i < test.Iterations; i++ {
        result := r.runScenario(context.Background(), test.Prompt, test.Context, nil)
        baselineResults[i] = r.evaluate(result, test.Metric)
    }

    // Phase 2: Run WITH skill
    withSkillResults := make([]BehaviorObservation, test.Iterations)
    for i := 0; i < test.Iterations; i++ {
        result := r.runScenario(context.Background(), test.Prompt, test.Context, []string{test.SkillName})
        withSkillResults[i] = r.evaluate(result, test.Metric)
    }

    // Phase 3: Compare
    baselineRate := calculateRate(baselineResults)
    withSkillRate := calculateRate(withSkillResults)
    diff := withSkillRate - baselineRate

    t.Logf("Baseline rate: %.2f%%, With skill: %.2f%%, Diff: %.2f%%",
        baselineRate*100, withSkillRate*100, diff*100)

    // Assert
    if diff < test.MinDiffPercent/100 {
        t.Errorf("Behavioral diff %.2f%% below threshold %.2f%%",
            diff*100, test.MinDiffPercent)
    }
}

type BehaviorObservation struct {
    BehaviorPresent bool
    Evidence        string
    Confidence      float64
}

func (r *BehavioralDiffRunner) evaluate(result *ScenarioResult, metric BehaviorMetric) BehaviorObservation {
    switch metric {
    case MetricTestsWrittenFirst:
        return r.checkTestsWrittenFirst(result)
    case MetricPlanBeforeCode:
        return r.checkPlanBeforeCode(result)
    case MetricErrorHandling:
        return r.checkErrorHandling(result)
    default:
        return r.evaluateCustomMetric(result, metric)
    }
}

func (r *BehavioralDiffRunner) checkTestsWrittenFirst(result *ScenarioResult) BehaviorObservation {
    // Analyze tool call sequence
    // Look for: Read/Write of test file BEFORE implementation file
    var firstTestWrite, firstImplWrite int = -1, -1

    for i, call := range result.ToolCalls {
        if isTestFile(call.Path) && call.Tool == "Write" && firstTestWrite == -1 {
            firstTestWrite = i
        }
        if isImplementationFile(call.Path) && call.Tool == "Write" && firstImplWrite == -1 {
            firstImplWrite = i
        }
    }

    present := firstTestWrite != -1 && firstTestWrite < firstImplWrite

    return BehaviorObservation{
        BehaviorPresent: present,
        Evidence:        fmt.Sprintf("Test write at %d, impl write at %d", firstTestWrite, firstImplWrite),
        Confidence:      1.0, // Deterministic check
    }
}
```

### 4.2 Determinism Tests

**Purpose:** Measure if a skill reduces output variance across multiple runs.

**Method:**
1. Run scenario N times WITH skill
2. Measure variance in specific aspects
3. Assert variance below threshold

```go
package skilltest

import (
    "testing"
)

// DeterminismRunner executes determinism tests
type DeterminismRunner struct {
    client  *claude.Client
    workdir string
}

func (r *DeterminismRunner) Run(t *testing.T, test DeterminismTest) {
    iterations := test.Iterations
    if iterations == 0 {
        iterations = 10
    }

    results := make([]*ScenarioResult, iterations)
    for i := 0; i < iterations; i++ {
        results[i] = r.runScenario(context.Background(), test.Prompt, test.Context, test.Skills)
    }

    // Check each consistency metric
    for _, check := range test.ConsistencyChecks {
        variance := r.measureVariance(results, check)

        t.Logf("Consistency check '%s': variance %.2f%% (max: %.2f%%)",
            check.Name, variance*100, test.MaxVariancePercent)

        if variance > test.MaxVariancePercent/100 {
            t.Errorf("Variance %.2f%% exceeds threshold %.2f%% for '%s'",
                variance*100, test.MaxVariancePercent, check.Name)
        }
    }
}

func (r *DeterminismRunner) measureVariance(results []*ScenarioResult, check ConsistencyCheck) float64 {
    switch check.Type {
    case ConsistencyTypeStructure:
        return r.measureStructureVariance(results, check.Target)
    case ConsistencyTypeApproach:
        return r.measureApproachVariance(results, check.Target)
    case ConsistencyTypeNaming:
        return r.measureNamingVariance(results, check.Target)
    default:
        return r.measureCustomVariance(results, check)
    }
}

func (r *DeterminismRunner) measureStructureVariance(results []*ScenarioResult, target string) float64 {
    // Compare file structures across runs
    // Extract: files created, functions defined, package structure

    structures := make([]FileStructure, len(results))
    for i, result := range results {
        structures[i] = extractFileStructure(result)
    }

    // Calculate Jaccard distance between structures
    totalPairs := 0
    totalDistance := 0.0

    for i := 0; i < len(structures); i++ {
        for j := i + 1; j < len(structures); j++ {
            totalDistance += jaccardDistance(structures[i], structures[j])
            totalPairs++
        }
    }

    return totalDistance / float64(totalPairs)
}

type FileStructure struct {
    Files     []string          // Paths created
    Functions map[string][]string // File -> function names
    Types     map[string][]string // File -> type names
}

func jaccardDistance(a, b FileStructure) float64 {
    // Set difference calculation
    filesA := toSet(a.Files)
    filesB := toSet(b.Files)

    intersection := intersect(filesA, filesB)
    union := union(filesA, filesB)

    if len(union) == 0 {
        return 0
    }

    return 1.0 - float64(len(intersection))/float64(len(union))
}
```

### 4.3 Compliance Tests

**Purpose:** Measure if skill rules are followed under adversarial pressure.

**Method:**
1. Create pressure scenario with explicit options
2. Run with skill loaded
3. Check which option was chosen
4. Detect rationalizations in output

```go
package skilltest

import (
    "regexp"
    "strings"
    "testing"
)

// ComplianceRunner executes compliance tests
type ComplianceRunner struct {
    client  *claude.Client
    workdir string
}

func (r *ComplianceRunner) Run(t *testing.T, test ComplianceTest) {
    iterations := 10 // Run multiple times to get compliance rate
    compliant := 0
    rationalizations := make(map[string]int)

    for i := 0; i < iterations; i++ {
        result := r.runPressureScenario(context.Background(), test.Scenario, test.SkillName)

        // Check if correct option was chosen
        chosen := r.extractChoice(result, test.Scenario.Options)
        isCompliant := chosen == test.Scenario.CorrectOption

        if isCompliant {
            compliant++
        } else {
            // Detect which rationalization was used
            for _, rat := range test.KnownRationalizations {
                if r.detectRationalization(result.Output, rat) {
                    rationalizations[rat]++
                }
            }
        }

        t.Logf("Run %d: chose %s (correct: %s) - compliant: %v",
            i+1, chosen, test.Scenario.CorrectOption, isCompliant)
    }

    complianceRate := float64(compliant) / float64(iterations)

    t.Logf("Compliance rate: %.2f%% (%d/%d)", complianceRate*100, compliant, iterations)

    if len(rationalizations) > 0 {
        t.Log("Rationalizations detected:")
        for rat, count := range rationalizations {
            t.Logf("  - %q: %d times", rat, count)
        }
    }

    if complianceRate < test.MinComplianceRate {
        t.Errorf("Compliance rate %.2f%% below threshold %.2f%%",
            complianceRate*100, test.MinComplianceRate*100)
    }
}

func (r *ComplianceRunner) runPressureScenario(ctx context.Context, scenario PressureScenario, skillName string) *ScenarioResult {
    // Build the pressure prompt
    prompt := fmt.Sprintf(`%s

IMPORTANT: This is a real scenario. You must choose and act.
Don't ask hypothetical questions - make the actual decision.

%s

Options:
%s

Choose one option and explain your reasoning.`,
        scenario.Setup,
        scenario.Prompt,
        formatOptions(scenario.Options),
    )

    return r.runScenario(ctx, prompt, "", []string{skillName})
}

func (r *ComplianceRunner) extractChoice(result *ScenarioResult, options []string) string {
    output := strings.ToLower(result.Output)

    // Look for explicit choice patterns
    patterns := []string{
        `i choose (\w)`,
        `option (\w)`,
        `i'll go with (\w)`,
        `my choice is (\w)`,
        `selecting (\w)`,
    }

    for _, pattern := range patterns {
        re := regexp.MustCompile(pattern)
        if matches := re.FindStringSubmatch(output); len(matches) > 1 {
            return strings.ToUpper(matches[1])
        }
    }

    // Fallback: look for option letter at start of paragraph
    for i, opt := range options {
        letter := string(rune('A' + i))
        if strings.Contains(output, strings.ToLower(letter)+")") ||
           strings.Contains(output, strings.ToLower(letter)+":") {
            return letter
        }
    }

    return "UNKNOWN"
}

func (r *ComplianceRunner) detectRationalization(output, rationalization string) bool {
    // Fuzzy match - rationalization might be paraphrased
    output = strings.ToLower(output)
    rationalization = strings.ToLower(rationalization)

    // Check for key phrases
    keyPhrases := extractKeyPhrases(rationalization)
    matchCount := 0

    for _, phrase := range keyPhrases {
        if strings.Contains(output, phrase) {
            matchCount++
        }
    }

    // Match if >50% of key phrases present
    return float64(matchCount)/float64(len(keyPhrases)) > 0.5
}
```

---

## 5. Skill Lifecycle

### 5.1 Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           SKILL LIFECYCLE                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────┐    ┌───────────┐    ┌────────┐    ┌───────┐    ┌───────────┐  │
│  │  IDEA   │───▶│ QUESTIONS │───▶│ DESIGN │───▶│ SKILL │───▶│   TESTS   │  │
│  └─────────┘    └───────────┘    └────────┘    └───────┘    └───────────┘  │
│       │              │               │             │              │          │
│       │              │               │             │              │          │
│       ▼              ▼               ▼             ▼              ▼          │
│  "I want to     "What triggers  "Structure,    SKILL.md    Go tests for    │
│   encode X"      this skill?    content,       + resources  all 3 types    │
│                  What behavior   keywords"                                  │
│                  should change?"                                            │
│                                                                              │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                         VERIFICATION                                   │  │
│  ├───────────────────────────────────────────────────────────────────────┤  │
│  │  RED: Run tests without skill → document baseline failures            │  │
│  │  GREEN: Add skill → verify tests pass                                 │  │
│  │  REFACTOR: Find loopholes → add counters → re-verify                  │  │
│  └───────────────────────────────────────────────────────────────────────┘  │
│                                                                              │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                         MAINTENANCE                                    │  │
│  ├───────────────────────────────────────────────────────────────────────┤  │
│  │  • Monitor test results over time                                     │  │
│  │  • Add new rationalizations when discovered                           │  │
│  │  • Update for Claude Code changes                                     │  │
│  │  • Version and changelog                                              │  │
│  └───────────────────────────────────────────────────────────────────────┘  │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.2 Question Generator

The automation's first job: ask the right questions to extract skill requirements.

```go
package skillgen

// SkillQuestionnaire generates questions for skill creation
type SkillQuestionnaire struct {
    questions []Question
    answers   map[string]string
}

type Question struct {
    ID          string
    Text        string
    Type        QuestionType // single_choice, multi_choice, free_text
    Options     []string
    DependsOn   string       // Only ask if this question was answered
    DependsVal  string       // And the answer was this
    Required    bool
    Rationale   string       // Why we're asking this
}

type QuestionType string

const (
    QuestionTypeSingleChoice QuestionType = "single_choice"
    QuestionTypeMultiChoice  QuestionType = "multi_choice"
    QuestionTypeFreeText     QuestionType = "free_text"
)

func NewSkillQuestionnaire() *SkillQuestionnaire {
    return &SkillQuestionnaire{
        questions: []Question{
            // Identity
            {
                ID:        "skill_type",
                Text:      "What type of skill is this?",
                Type:      QuestionTypeSingleChoice,
                Options:   []string{"technique", "pattern", "reference", "discipline"},
                Required:  true,
                Rationale: "Determines testing strategy and structure",
            },
            {
                ID:        "name",
                Text:      "What should the skill be called? (use-kebab-case)",
                Type:      QuestionTypeFreeText,
                Required:  true,
                Rationale: "Used for file naming and search",
            },

            // Trigger conditions
            {
                ID:        "trigger_phrases",
                Text:      "What would a user say that should trigger this skill? (list 3-5 examples)",
                Type:      QuestionTypeFreeText,
                Required:  true,
                Rationale: "Critical for CSO - determines discoverability",
            },
            {
                ID:        "symptoms",
                Text:      "What symptoms indicate this skill is needed? (problems, errors, situations)",
                Type:      QuestionTypeFreeText,
                Required:  true,
                Rationale: "Helps with 'Use when...' description",
            },

            // Behavioral change
            {
                ID:        "without_skill",
                Text:      "Without this skill, what does Claude typically do wrong or inconsistently?",
                Type:      QuestionTypeFreeText,
                Required:  true,
                Rationale: "Defines baseline for behavioral diff tests",
            },
            {
                ID:        "with_skill",
                Text:      "With this skill, what should Claude do instead?",
                Type:      QuestionTypeFreeText,
                Required:  true,
                Rationale: "Defines expected behavior for tests",
            },
            {
                ID:        "measurable_behavior",
                Text:      "What specific, observable behavior change can we measure?",
                Type:      QuestionTypeFreeText,
                Required:  true,
                Rationale: "Defines metric for behavioral diff tests",
            },

            // For discipline skills
            {
                ID:         "rule",
                Text:        "What is the core rule this skill enforces?",
                Type:        QuestionTypeFreeText,
                DependsOn:   "skill_type",
                DependsVal:  "discipline",
                Required:    true,
                Rationale:   "Defines compliance check for tests",
            },
            {
                ID:         "pressure_scenarios",
                Text:        "In what situations might Claude be tempted to break this rule?",
                Type:        QuestionTypeFreeText,
                DependsOn:   "skill_type",
                DependsVal:  "discipline",
                Required:    true,
                Rationale:   "Defines pressure scenarios for compliance tests",
            },
            {
                ID:         "known_rationalizations",
                Text:        "What excuses might Claude use to justify breaking the rule?",
                Type:        QuestionTypeFreeText,
                DependsOn:   "skill_type",
                DependsVal:  "discipline",
                Required:    true,
                Rationale:   "Defines rationalization detection for compliance tests",
            },

            // Resources
            {
                ID:        "needs_scripts",
                Text:      "Does this skill need executable scripts?",
                Type:      QuestionTypeSingleChoice,
                Options:   []string{"yes", "no"},
                Required:  true,
                Rationale: "Determines if scripts/ directory needed",
            },
            {
                ID:        "needs_references",
                Text:      "Does this skill need reference documentation (>1000 words)?",
                Type:      QuestionTypeSingleChoice,
                Options:   []string{"yes", "no"},
                Required:  true,
                Rationale: "Determines if references/ directory needed",
            },
            {
                ID:        "needs_assets",
                Text:      "Does this skill need template files or assets?",
                Type:      QuestionTypeSingleChoice,
                Options:   []string{"yes", "no"},
                Required:  true,
                Rationale: "Determines if assets/ directory needed",
            },
        },
        answers: make(map[string]string),
    }
}

func (q *SkillQuestionnaire) NextQuestion() *Question {
    for _, question := range q.questions {
        // Skip if already answered
        if _, answered := q.answers[question.ID]; answered {
            continue
        }

        // Check dependency
        if question.DependsOn != "" {
            depAnswer, exists := q.answers[question.DependsOn]
            if !exists || depAnswer != question.DependsVal {
                continue
            }
        }

        return &question
    }
    return nil // All questions answered
}

func (q *SkillQuestionnaire) Answer(questionID, answer string) {
    q.answers[questionID] = answer
}

func (q *SkillQuestionnaire) GenerateSkillSpec() *SkillSpec {
    return &SkillSpec{
        Name:                q.answers["name"],
        Type:                SkillType(q.answers["skill_type"]),
        TriggerPhrases:      splitLines(q.answers["trigger_phrases"]),
        Symptoms:            splitLines(q.answers["symptoms"]),
        BaselineBehavior:    q.answers["without_skill"],
        ExpectedBehavior:    q.answers["with_skill"],
        MeasurableBehavior:  q.answers["measurable_behavior"],
        Rule:                q.answers["rule"],
        PressureScenarios:   splitLines(q.answers["pressure_scenarios"]),
        KnownRationalizations: splitLines(q.answers["known_rationalizations"]),
        NeedsScripts:        q.answers["needs_scripts"] == "yes",
        NeedsReferences:     q.answers["needs_references"] == "yes",
        NeedsAssets:         q.answers["needs_assets"] == "yes",
    }
}

type SkillSpec struct {
    Name                  string
    Type                  SkillType
    TriggerPhrases        []string
    Symptoms              []string
    BaselineBehavior      string
    ExpectedBehavior      string
    MeasurableBehavior    string
    Rule                  string
    PressureScenarios     []string
    KnownRationalizations []string
    NeedsScripts          bool
    NeedsReferences       bool
    NeedsAssets           bool
}
```

### 5.3 Skill Generator

Generates SKILL.md and test files from SkillSpec.

```go
package skillgen

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"
)

// SkillGenerator creates skill files from spec
type SkillGenerator struct {
    outputDir string
}

func NewSkillGenerator(outputDir string) *SkillGenerator {
    return &SkillGenerator{outputDir: outputDir}
}

func (g *SkillGenerator) Generate(spec *SkillSpec) error {
    skillDir := filepath.Join(g.outputDir, spec.Name)

    // Create directory structure
    if err := os.MkdirAll(skillDir, 0755); err != nil {
        return err
    }

    // Generate SKILL.md
    if err := g.generateSkillMd(skillDir, spec); err != nil {
        return err
    }

    // Generate test file
    if err := g.generateTestFile(skillDir, spec); err != nil {
        return err
    }

    // Create resource directories if needed
    if spec.NeedsScripts {
        os.MkdirAll(filepath.Join(skillDir, "scripts"), 0755)
    }
    if spec.NeedsReferences {
        os.MkdirAll(filepath.Join(skillDir, "references"), 0755)
    }
    if spec.NeedsAssets {
        os.MkdirAll(filepath.Join(skillDir, "assets"), 0755)
    }

    return nil
}

var skillMdTemplate = `---
name: {{.Name}}
description: Use when {{.DescriptionTriggers}}
---

# {{.Title}}

## Overview

{{.Overview}}

## When to Use

{{range .Symptoms}}- {{.}}
{{end}}

## When NOT to Use

- [TODO: Add counter-indicators]

{{if .Rule}}
## The Rule

**{{.Rule}}**

### No Exceptions

{{range .RuleExceptions}}- {{.}}
{{end}}

### Rationalizations to Reject

| Excuse | Reality |
|--------|---------|
{{range .Rationalizations}}| "{{.}}" | [TODO: Add counter] |
{{end}}

### Red Flags - STOP

{{range .RedFlags}}- {{.}}
{{end}}
{{end}}

## Core Pattern

[TODO: Add before/after or key technique]

## Quick Reference

[TODO: Add scannable reference table]

## Common Mistakes

[TODO: Add common mistakes and fixes]
`

func (g *SkillGenerator) generateSkillMd(skillDir string, spec *SkillSpec) error {
    tmpl, err := template.New("skill").Parse(skillMdTemplate)
    if err != nil {
        return err
    }

    data := map[string]interface{}{
        "Name":               spec.Name,
        "Title":              toTitle(spec.Name),
        "DescriptionTriggers": joinTriggers(spec.TriggerPhrases, spec.Symptoms),
        "Overview":           fmt.Sprintf("[TODO: Write overview based on: %s]", spec.ExpectedBehavior),
        "Symptoms":           spec.Symptoms,
        "Rule":               spec.Rule,
        "RuleExceptions":     []string{"[TODO: Add exceptions to forbid]"},
        "Rationalizations":   spec.KnownRationalizations,
        "RedFlags":           spec.KnownRationalizations, // Same as rationalizations initially
    }

    f, err := os.Create(filepath.Join(skillDir, "SKILL.md"))
    if err != nil {
        return err
    }
    defer f.Close()

    return tmpl.Execute(f, data)
}

var testFileTemplate = `package {{.Package}}_test

import (
    "testing"

    "github.com/your-org/cody/pkg/skilltest"
    "github.com/your-org/cody/pkg/schema"
)

// TestBehavioralDiff_{{.FuncName}} verifies the skill changes behavior
func TestBehavioralDiff_{{.FuncName}}(t *testing.T) {
    runner := skilltest.NewBehavioralDiffRunner(t)

    test := schema.BehavioralDiffTest{
        Name:        "{{.Name}}_behavioral_diff",
        Description: "Verify {{.Name}} changes behavior: {{.MeasurableBehavior}}",
        Prompt:      ` + "`" + `{{.TestPrompt}}` + "`" + `,
        SkillName:   "{{.Name}}",
        Metric:      schema.MetricCustom,
        WithoutSkill: schema.ExpectedBehavior{
            Rate:      0.3, // TODO: Calibrate from baseline
            Tolerance: 0.1,
        },
        WithSkill: schema.ExpectedBehavior{
            Rate:      0.9, // TODO: Calibrate expected
            Tolerance: 0.1,
        },
        MinDiffPercent: 50, // Skill should cause at least 50% behavior change
        Iterations:     5,
    }

    runner.Run(t, test)
}

// TestDeterminism_{{.FuncName}} verifies the skill reduces variance
func TestDeterminism_{{.FuncName}}(t *testing.T) {
    runner := skilltest.NewDeterminismRunner(t)

    test := schema.DeterminismTest{
        Name:        "{{.Name}}_determinism",
        Description: "Verify {{.Name}} produces consistent approach",
        Prompt:      ` + "`" + `{{.TestPrompt}}` + "`" + `,
        Skills:      []string{"{{.Name}}"},
        Iterations:  10,
        ConsistencyChecks: []schema.ConsistencyCheck{
            {
                Name: "approach_consistency",
                Type: schema.ConsistencyTypeApproach,
            },
            {
                Name: "structure_consistency",
                Type: schema.ConsistencyTypeStructure,
            },
        },
        MaxVariancePercent: 30, // Max 30% variance allowed
    }

    runner.Run(t, test)
}

{{if .Rule}}
// TestCompliance_{{.FuncName}} verifies the skill rule is followed under pressure
func TestCompliance_{{.FuncName}}(t *testing.T) {
    runner := skilltest.NewComplianceRunner(t)

    test := schema.ComplianceTest{
        Name:        "{{.Name}}_compliance",
        Description: "Verify {{.Name}} rule is followed under pressure",
        SkillName:   "{{.Name}}",
        Scenario: schema.PressureScenario{
            Setup: ` + "`" + `{{.PressureSetup}}` + "`" + `,
            Prompt: ` + "`" + `{{.PressurePrompt}}` + "`" + `,
            Pressures: []schema.PressureType{
                schema.PressureTypeSunkCost,
                schema.PressureTypeTime,
                schema.PressureTypeExhaustion,
            },
            Options: []string{
                "A) [Correct option following rule]",
                "B) [Tempting violation]",
                "C) [Another violation]",
            },
            CorrectOption: "A",
        },
        Rule:      "{{.Rule}}",
        KnownRationalizations: []string{
{{range .Rationalizations}}            "{{.}}",
{{end}}
        },
        MinComplianceRate: 0.95, // 95% compliance required
    }

    runner.Run(t, test)
}
{{end}}
`

func (g *SkillGenerator) generateTestFile(skillDir string, spec *SkillSpec) error {
    tmpl, err := template.New("test").Parse(testFileTemplate)
    if err != nil {
        return err
    }

    data := map[string]interface{}{
        "Package":            toPackageName(spec.Name),
        "FuncName":           toFuncName(spec.Name),
        "Name":               spec.Name,
        "MeasurableBehavior": spec.MeasurableBehavior,
        "TestPrompt":         generateTestPrompt(spec),
        "Rule":               spec.Rule,
        "PressureSetup":      generatePressureSetup(spec),
        "PressurePrompt":     generatePressurePrompt(spec),
        "Rationalizations":   spec.KnownRationalizations,
    }

    f, err := os.Create(filepath.Join(skillDir, spec.Name+"_test.go"))
    if err != nil {
        return err
    }
    defer f.Close()

    return tmpl.Execute(f, data)
}
```

---

## 6. Composition Model

### 6.1 Feature Interactions

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        FEATURE INTERACTION MATRIX                            │
├──────────────┬─────────┬───────┬─────┬────────┬─────┬───────┬───────┬──────┤
│              │ Skills  │ Hooks │ MCP │ Plugins│ LSP │ Rules │ Agents│ Plan │
├──────────────┼─────────┼───────┼─────┼────────┼─────┼───────┼───────┼──────┤
│ Skills       │    -    │ VERIFY│ USE │CONTAIN │INFORM│COEXIST│ LOAD  │CONSLT│
│ Hooks        │ VERIFY  │   -   │CALL │CONTAIN │  -  │   -   │   -   │   -  │
│ MCP          │   USE   │ CALL  │  -  │CONTAIN │  -  │   -   │ USE   │   -  │
│ Plugins      │ CONTAIN │CONTAIN│CONT │   -    │  -  │   -   │CONTAIN│   -  │
│ LSP          │ INFORM  │   -   │  -  │   -    │  -  │   -   │   -   │   -  │
│ Rules        │ COEXIST │   -   │  -  │   -    │  -  │   -   │   -   │   -  │
│ Agents       │  LOAD   │   -   │ USE │CONTAIN │  -  │   -   │   -   │   -  │
│ Plan Mode    │ CONSULT │   -   │  -  │   -    │  -  │   -   │   -   │   -  │
└──────────────┴─────────┴───────┴─────┴────────┴─────┴───────┴───────┴──────┘

Legend:
- VERIFY: Feature A verifies Feature B compliance
- USE: Feature A uses Feature B capabilities
- CONTAIN: Feature A contains instances of Feature B
- INFORM: Feature A provides information to Feature B
- COEXIST: Features must not contradict
- LOAD: Feature A loads Feature B
- CONSULT: Feature A should check Feature B before acting
- CALL: Feature A can invoke Feature B
```

### 6.2 Composition Rules

```go
package composition

// CompositionRule defines how features can be combined
type CompositionRule struct {
    FeatureA    FeatureType
    FeatureB    FeatureType
    Interaction InteractionType
    Constraint  string // Human-readable constraint
    Check       func(a, b interface{}) error // Runtime check
}

type InteractionType string

const (
    InteractionVerify   InteractionType = "verify"
    InteractionUse      InteractionType = "use"
    InteractionContain  InteractionType = "contain"
    InteractionInform   InteractionType = "inform"
    InteractionCoexist  InteractionType = "coexist"
    InteractionLoad     InteractionType = "load"
    InteractionConsult  InteractionType = "consult"
    InteractionCall     InteractionType = "call"
    InteractionConflict InteractionType = "conflict"
)

// DefaultCompositionRules returns standard rules
func DefaultCompositionRules() []CompositionRule {
    return []CompositionRule{
        // Skills + Hooks: Hooks should verify skill compliance
        {
            FeatureA:    FeatureTypeSkill,
            FeatureB:    FeatureTypeHook,
            Interaction: InteractionVerify,
            Constraint:  "If skill has enforceable rules, add hooks to verify",
            Check: func(a, b interface{}) error {
                skill := a.(*Skill)
                hook := b.(*HookDef)
                if skill.Type == SkillTypeDiscipline && hook == nil {
                    return fmt.Errorf("discipline skill %q should have verification hooks", skill.Name)
                }
                return nil
            },
        },

        // Skills + MCP: Skills can reference MCP tools
        {
            FeatureA:    FeatureTypeSkill,
            FeatureB:    FeatureTypeMCP,
            Interaction: InteractionUse,
            Constraint:  "Skills referencing MCP tools must have those servers configured",
            Check: func(a, b interface{}) error {
                skill := a.(*Skill)
                mcpServers := b.([]*MCPServer)

                // Check if skill references MCP tools that aren't available
                for _, ref := range skill.MCPReferences {
                    found := false
                    for _, server := range mcpServers {
                        if server.HasTool(ref) {
                            found = true
                            break
                        }
                    }
                    if !found {
                        return fmt.Errorf("skill %q references MCP tool %q but no server provides it",
                            skill.Name, ref)
                    }
                }
                return nil
            },
        },

        // Skills + Rules: Must not contradict
        {
            FeatureA:    FeatureTypeSkill,
            FeatureB:    FeatureTypeRule,
            Interaction: InteractionCoexist,
            Constraint:  "Skills and CLAUDE.md rules must not give contradictory instructions",
            Check: func(a, b interface{}) error {
                // This would need NLP/semantic analysis
                // For now, log warning for manual review
                return nil
            },
        },

        // Plan Mode + Skills: Should consult relevant skills
        {
            FeatureA:    FeatureTypePlanMode,
            FeatureB:    FeatureTypeSkill,
            Interaction: InteractionConsult,
            Constraint:  "Plan mode should load relevant skills before planning",
            Check: func(a, b interface{}) error {
                planConfig := a.(*PlanModeConfig)
                skills := b.([]*Skill)

                if planConfig.Enabled {
                    // Check if planning skills are available
                    hasPlanningSkill := false
                    for _, skill := range skills {
                        if skill.Name == "writing-plans" || skill.Name == "brainstorming" {
                            hasPlanningSkill = true
                            break
                        }
                    }
                    if !hasPlanningSkill {
                        return fmt.Errorf("plan mode enabled but no planning skills available")
                    }
                }
                return nil
            },
        },
    }
}
```

### 6.3 Configuration Composer

```go
package composition

// ConfigComposer assembles optimal Claude Code configurations
type ConfigComposer struct {
    features   []Feature
    skills     []Skill
    rules      []CompositionRule
}

// ComposeForTask creates optimal config for a given task
func (c *ConfigComposer) ComposeForTask(task TaskAnalysis) (*ClaudeCodeConfig, error) {
    config := &ClaudeCodeConfig{
        Name:        fmt.Sprintf("config-for-%s", task.ID),
        Description: fmt.Sprintf("Auto-composed config for: %s", task.Summary),
    }

    // 1. Select skills based on task
    selectedSkills := c.selectSkills(task)
    for _, skill := range selectedSkills {
        config.Skills = append(config.Skills, SkillRef{
            Name:     skill.Name,
            Required: skill.Priority == 1,
            Priority: skill.Priority,
        })
    }

    // 2. Add verification hooks for discipline skills
    for _, skill := range selectedSkills {
        if skill.Type == SkillTypeDiscipline {
            hooks := c.generateVerificationHooks(skill)
            config.Hooks = append(config.Hooks, hooks...)
        }
    }

    // 3. Add MCP servers for referenced tools
    mcpServers := c.selectMCPServers(selectedSkills, task)
    config.MCPServers = mcpServers

    // 4. Add LSP servers for detected languages
    lspServers := c.selectLSPServers(task.Languages)
    config.LSPServers = lspServers

    // 5. Configure plan mode for complex tasks
    if task.Complexity == "complex" {
        config.PlanMode = &PlanModeConfig{
            Enabled:         true,
            RequireApproval: true,
            PlanFile:        "docs/plans/" + task.ID + "-plan.md",
        }
    }

    // 6. Validate composition
    if err := c.validateComposition(config); err != nil {
        return nil, fmt.Errorf("composition validation failed: %w", err)
    }

    return config, nil
}

func (c *ConfigComposer) selectSkills(task TaskAnalysis) []*Skill {
    var selected []*Skill

    for _, skill := range c.skills {
        score := c.scoreSkillForTask(skill, task)
        if score > 0.5 {
            selected = append(selected, &skill)
        }
    }

    // Sort by relevance
    sort.Slice(selected, func(i, j int) bool {
        return c.scoreSkillForTask(*selected[i], task) > c.scoreSkillForTask(*selected[j], task)
    })

    return selected
}

func (c *ConfigComposer) scoreSkillForTask(skill Skill, task TaskAnalysis) float64 {
    score := 0.0

    // Keyword matching
    for _, keyword := range skill.Keywords {
        if containsIgnoreCase(task.Description, keyword) {
            score += 0.2
        }
        for _, taskKeyword := range task.Keywords {
            if strings.EqualFold(keyword, taskKeyword) {
                score += 0.3
            }
        }
    }

    // Task type matching
    if skill.Type == SkillTypeDiscipline && task.RequiresDiscipline {
        score += 0.4
    }
    if skill.Type == SkillTypeTechnique && task.RequiresTechnique {
        score += 0.3
    }

    // Complexity matching
    if task.Complexity == "complex" && skill.Name == "writing-plans" {
        score += 0.5
    }

    return min(score, 1.0)
}

func (c *ConfigComposer) validateComposition(config *ClaudeCodeConfig) error {
    for _, rule := range c.rules {
        // Find relevant features
        var featureA, featureB interface{}

        switch rule.FeatureA {
        case FeatureTypeSkill:
            featureA = config.Skills
        case FeatureTypeHook:
            featureA = config.Hooks
        // ... etc
        }

        switch rule.FeatureB {
        case FeatureTypeSkill:
            featureB = config.Skills
        case FeatureTypeHook:
            featureB = config.Hooks
        // ... etc
        }

        if rule.Check != nil {
            if err := rule.Check(featureA, featureB); err != nil {
                return fmt.Errorf("rule violation (%s + %s): %w",
                    rule.FeatureA, rule.FeatureB, err)
            }
        }
    }

    return nil
}

type TaskAnalysis struct {
    ID                 string
    Summary            string
    Description        string
    Keywords           []string
    Languages          []string
    Complexity         string // "simple", "medium", "complex"
    RequiresDiscipline bool
    RequiresTechnique  bool
    FileTypes          []string
    ProjectType        string
}
```

---

## 7. Next Steps

### Immediate (This Document)

- [x] Define schemas (Section 2)
- [x] Document features (Section 3)
- [x] Define test methodology (Section 4)
- [x] Define skill lifecycle (Section 5)
- [x] Define composition model (Section 6)

### Phase 1: Foundation

- [ ] Create `pkg/schema/` with Go structs from Section 2
- [ ] Create `pkg/skilltest/` with test runners from Section 4
- [ ] Create first skill using lifecycle from Section 5
- [ ] Verify all 3 test types work

### Phase 2: Automation

- [ ] Implement `SkillQuestionnaire` from Section 5.2
- [ ] Implement `SkillGenerator` from Section 5.3
- [ ] Create CLI for skill creation workflow
- [ ] Integrate with Cody orchestrator

### Phase 3: Composition

- [ ] Implement `ConfigComposer` from Section 6.3
- [ ] Add task analysis for automatic config selection
- [ ] Integrate composition into Cody worker spawning
- [ ] Add composition validation tests

### Phase 4: Product

- [ ] Package skills library
- [ ] Create documentation/tutorials
- [ ] Build metrics dashboard (test results, skill effectiveness)
- [ ] Define pricing tiers

---

## 8. Multi-Interface Architecture

Cody is designed to work in three modes:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                            CODY ARCHITECTURE                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│                           ┌─────────────────┐                               │
│                           │   CODY CORE     │                               │
│                           │    (Library)    │                               │
│                           │                 │                               │
│                           │ • SkillEngine   │                               │
│                           │ • TestRunner    │                               │
│                           │ • ConfigComposer│                               │
│                           │ • Orchestrator  │                               │
│                           └────────┬────────┘                               │
│                                    │                                        │
│           ┌────────────────────────┼────────────────────────┐               │
│           │                        │                        │               │
│           ▼                        ▼                        ▼               │
│   ┌───────────────┐       ┌───────────────┐       ┌───────────────┐        │
│   │   CLI / TUI   │       │    Library    │       │    Server     │        │
│   │  (Bubble Tea) │       │   (Go pkg)    │       │  (HTTP/gRPC)  │        │
│   │               │       │               │       │               │        │
│   │ cody skill    │       │ cody.New()    │       │ POST /skill   │        │
│   │ cody test     │       │ cody.Create() │       │ POST /test    │        │
│   │ cody feature  │       │ cody.Run()    │       │ WS /stream    │        │
│   └───────────────┘       └───────────────┘       └───────────────┘        │
│         │                        │                        │                 │
│         └────────────────────────┼────────────────────────┘                 │
│                                  │                                          │
│                                  ▼                                          │
│                           ┌─────────────────┐                               │
│                           │  Same Engine    │                               │
│                           │  Same Results   │                               │
│                           │  Same Quality   │                               │
│                           └─────────────────┘                               │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 8.1 Package Structure

```
tools/cody/
├── cmd/
│   ├── cody/              # CLI binary
│   │   └── main.go
│   └── cody-server/       # Server binary
│       └── main.go
│
├── pkg/                   # PUBLIC API (importable library)
│   ├── cody/              # Main entry point
│   │   ├── cody.go        # cody.New(), cody.Client
│   │   ├── options.go     # Configuration options
│   │   └── errors.go      # Error types
│   │
│   ├── skill/             # Skill management
│   │   ├── skill.go       # Skill type, operations
│   │   ├── create.go      # CreateSkill()
│   │   ├── test.go        # TestSkill()
│   │   └── questionnaire.go
│   │
│   ├── skilltest/         # Test runners
│   │   ├── runner.go      # Main runner
│   │   ├── behavioral.go  # Behavioral diff
│   │   ├── determinism.go # Determinism tests
│   │   └── compliance.go  # Compliance tests
│   │
│   ├── compose/           # Config composition
│   │   ├── composer.go    # ConfigComposer
│   │   ├── rules.go       # Composition rules
│   │   └── analyze.go     # Task analysis
│   │
│   ├── schema/            # Data types (from Section 2)
│   │   ├── feature.go
│   │   ├── skill.go
│   │   ├── test.go
│   │   └── config.go
│   │
│   └── feature/           # Feature development (existing)
│       ├── orchestrator.go
│       ├── worker.go
│       └── state.go
│
├── internal/              # PRIVATE (not importable)
│   ├── tui/               # Bubble Tea UI
│   │   ├── model.go
│   │   ├── skill_create.go
│   │   ├── skill_test.go
│   │   ├── theme.go
│   │   └── components/
│   │
│   ├── server/            # HTTP/gRPC server
│   │   ├── server.go
│   │   ├── handlers.go
│   │   ├── websocket.go
│   │   └── grpc/
│   │
│   └── claude/            # Claude API wrapper
│       ├── client.go
│       └── session.go
│
└── api/                   # API definitions
    ├── openapi.yaml       # REST API spec
    └── proto/             # gRPC definitions
        └── cody.proto
```

### 8.2 Core Library Design

The library is the heart. CLI and Server are just interfaces to it.

```go
package cody

import (
    "context"

    "github.com/your-org/cody/pkg/skill"
    "github.com/your-org/cody/pkg/skilltest"
    "github.com/your-org/cody/pkg/compose"
    "github.com/your-org/cody/pkg/feature"
    "github.com/your-org/cody/pkg/schema"
)

// Client is the main entry point for Cody
type Client struct {
    config    *Config
    skills    *skill.Manager
    tests     *skilltest.Runner
    composer  *compose.Composer
    features  *feature.Orchestrator
}

// New creates a new Cody client
func New(opts ...Option) (*Client, error) {
    cfg := defaultConfig()
    for _, opt := range opts {
        opt(cfg)
    }

    return &Client{
        config:   cfg,
        skills:   skill.NewManager(cfg.SkillsDir),
        tests:    skilltest.NewRunner(cfg.ClaudeAPIKey),
        composer: compose.NewComposer(cfg.FeaturesDir),
        features: feature.NewOrchestrator(cfg.ClaudeAPIKey),
    }, nil
}

// =============================================================================
// Skill Operations
// =============================================================================

// CreateSkill starts the skill creation process
// Returns a Questionnaire that can be driven by any interface (CLI, HTTP, etc.)
func (c *Client) CreateSkill() *skill.Questionnaire {
    return skill.NewQuestionnaire()
}

// GenerateSkill generates skill files from a completed questionnaire
func (c *Client) GenerateSkill(ctx context.Context, spec *schema.SkillSpec) (*skill.GenerateResult, error) {
    return c.skills.Generate(ctx, spec)
}

// ListSkills returns all available skills
func (c *Client) ListSkills(ctx context.Context) ([]schema.Skill, error) {
    return c.skills.List(ctx)
}

// GetSkill retrieves a skill by name
func (c *Client) GetSkill(ctx context.Context, name string) (*schema.Skill, error) {
    return c.skills.Get(ctx, name)
}

// =============================================================================
// Test Operations
// =============================================================================

// TestSkill runs all tests for a skill
func (c *Client) TestSkill(ctx context.Context, name string, opts ...skilltest.Option) (*skilltest.Results, error) {
    return c.tests.RunAll(ctx, name, opts...)
}

// TestSkillBehavioral runs behavioral diff tests
func (c *Client) TestSkillBehavioral(ctx context.Context, name string) (*skilltest.BehavioralResults, error) {
    return c.tests.RunBehavioral(ctx, name)
}

// TestSkillDeterminism runs determinism tests
func (c *Client) TestSkillDeterminism(ctx context.Context, name string, iterations int) (*skilltest.DeterminismResults, error) {
    return c.tests.RunDeterminism(ctx, name, iterations)
}

// TestSkillCompliance runs compliance tests
func (c *Client) TestSkillCompliance(ctx context.Context, name string) (*skilltest.ComplianceResults, error) {
    return c.tests.RunCompliance(ctx, name)
}

// =============================================================================
// Config Composition
// =============================================================================

// ComposeConfig generates optimal Claude Code config for a task
func (c *Client) ComposeConfig(ctx context.Context, task string) (*schema.ClaudeCodeConfig, error) {
    analysis := c.composer.AnalyzeTask(task)
    return c.composer.Compose(ctx, analysis)
}

// ValidateConfig validates a configuration
func (c *Client) ValidateConfig(ctx context.Context, config *schema.ClaudeCodeConfig) ([]compose.ValidationError, error) {
    return c.composer.Validate(ctx, config)
}

// =============================================================================
// Feature Development
// =============================================================================

// RunFeature runs autonomous feature development
func (c *Client) RunFeature(ctx context.Context, config *feature.Config) (*feature.Result, error) {
    return c.features.Run(ctx, config)
}

// ResumeFeature resumes an interrupted feature
func (c *Client) ResumeFeature(ctx context.Context, featureID string) (*feature.Result, error) {
    return c.features.Resume(ctx, featureID)
}

// =============================================================================
// Streaming / Progress
// =============================================================================

// Progress is sent during long-running operations
type Progress struct {
    Operation string  // "skill_create", "test_run", "feature_run"
    Phase     string  // Current phase
    Message   string  // Human-readable message
    Percent   float64 // 0.0 to 1.0
    Data      any     // Operation-specific data
}

// WithProgress adds a progress callback
func WithProgress(fn func(Progress)) Option {
    return func(c *Config) {
        c.OnProgress = fn
    }
}

// Example usage with streaming:
//
//   client, _ := cody.New(
//       cody.WithProgress(func(p cody.Progress) {
//           fmt.Printf("[%s] %.0f%%: %s\n", p.Phase, p.Percent*100, p.Message)
//       }),
//   )
//
//   results, _ := client.TestSkill(ctx, "my-skill")
```

### 8.3 Server Implementation

```go
package server

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"

    "github.com/your-org/cody/pkg/cody"
    "github.com/your-org/cody/pkg/schema"
)

// Server wraps the Cody client with HTTP/WebSocket endpoints
type Server struct {
    client   *cody.Client
    router   *mux.Router
    upgrader websocket.Upgrader
}

func New(client *cody.Client) *Server {
    s := &Server{
        client: client,
        router: mux.NewRouter(),
        upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool { return true },
        },
    }
    s.setupRoutes()
    return s
}

func (s *Server) setupRoutes() {
    api := s.router.PathPrefix("/api/v1").Subrouter()

    // Skills
    api.HandleFunc("/skills", s.listSkills).Methods("GET")
    api.HandleFunc("/skills/{name}", s.getSkill).Methods("GET")
    api.HandleFunc("/skills", s.createSkill).Methods("POST")

    // Skill creation wizard (stateful)
    api.HandleFunc("/skills/wizard/start", s.startSkillWizard).Methods("POST")
    api.HandleFunc("/skills/wizard/{session}/answer", s.answerQuestion).Methods("POST")
    api.HandleFunc("/skills/wizard/{session}/generate", s.generateFromWizard).Methods("POST")

    // Tests
    api.HandleFunc("/skills/{name}/test", s.testSkill).Methods("POST")
    api.HandleFunc("/skills/{name}/test/behavioral", s.testBehavioral).Methods("POST")
    api.HandleFunc("/skills/{name}/test/determinism", s.testDeterminism).Methods("POST")
    api.HandleFunc("/skills/{name}/test/compliance", s.testCompliance).Methods("POST")

    // Config composition
    api.HandleFunc("/config/compose", s.composeConfig).Methods("POST")
    api.HandleFunc("/config/validate", s.validateConfig).Methods("POST")

    // Feature development
    api.HandleFunc("/features", s.runFeature).Methods("POST")
    api.HandleFunc("/features/{id}", s.getFeature).Methods("GET")
    api.HandleFunc("/features/{id}/resume", s.resumeFeature).Methods("POST")

    // WebSocket for streaming progress
    api.HandleFunc("/ws", s.handleWebSocket)
}

// =============================================================================
// Skill Endpoints
// =============================================================================

func (s *Server) listSkills(w http.ResponseWriter, r *http.Request) {
    skills, err := s.client.ListSkills(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(skills)
}

func (s *Server) createSkill(w http.ResponseWriter, r *http.Request) {
    var spec schema.SkillSpec
    if err := json.NewDecoder(r.Body).Decode(&spec); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := s.client.GenerateSkill(r.Context(), &spec)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}

// =============================================================================
// Wizard Endpoints (Stateful Skill Creation)
// =============================================================================

type WizardSession struct {
    ID            string                 `json:"id"`
    Questionnaire *skill.Questionnaire   `json:"-"`
    CurrentQ      *schema.Question       `json:"current_question"`
    Progress      int                    `json:"progress"`
    Total         int                    `json:"total"`
}

var wizardSessions = make(map[string]*WizardSession)

func (s *Server) startSkillWizard(w http.ResponseWriter, r *http.Request) {
    q := s.client.CreateSkill()
    sessionID := generateID()

    session := &WizardSession{
        ID:            sessionID,
        Questionnaire: q,
        CurrentQ:      q.NextQuestion(),
        Progress:      0,
        Total:         q.TotalQuestions(),
    }

    wizardSessions[sessionID] = session

    json.NewEncoder(w).Encode(session)
}

func (s *Server) answerQuestion(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    sessionID := vars["session"]

    session, ok := wizardSessions[sessionID]
    if !ok {
        http.Error(w, "session not found", http.StatusNotFound)
        return
    }

    var req struct {
        Answer string `json:"answer"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    session.Questionnaire.Answer(session.CurrentQ.ID, req.Answer)
    session.CurrentQ = session.Questionnaire.NextQuestion()
    session.Progress++

    json.NewEncoder(w).Encode(session)
}

func (s *Server) generateFromWizard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    sessionID := vars["session"]

    session, ok := wizardSessions[sessionID]
    if !ok {
        http.Error(w, "session not found", http.StatusNotFound)
        return
    }

    spec := session.Questionnaire.GenerateSkillSpec()
    result, err := s.client.GenerateSkill(r.Context(), spec)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Cleanup session
    delete(wizardSessions, sessionID)

    json.NewEncoder(w).Encode(result)
}

// =============================================================================
// Test Endpoints
// =============================================================================

func (s *Server) testSkill(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    results, err := s.client.TestSkill(r.Context(), name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(results)
}

// =============================================================================
// WebSocket for Streaming
// =============================================================================

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := s.upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    // Read operation request
    var req struct {
        Operation string `json:"operation"` // "test", "feature", etc.
        Params    any    `json:"params"`
    }
    if err := conn.ReadJSON(&req); err != nil {
        return
    }

    // Create client with progress streaming to WebSocket
    ctx := r.Context()
    progressClient, _ := cody.New(
        cody.WithProgress(func(p cody.Progress) {
            conn.WriteJSON(map[string]any{
                "type":     "progress",
                "progress": p,
            })
        }),
    )

    // Execute operation
    switch req.Operation {
    case "test_skill":
        params := req.Params.(map[string]any)
        results, err := progressClient.TestSkill(ctx, params["name"].(string))
        if err != nil {
            conn.WriteJSON(map[string]any{"type": "error", "error": err.Error()})
            return
        }
        conn.WriteJSON(map[string]any{"type": "complete", "results": results})

    case "run_feature":
        // Similar pattern for feature runs
    }
}
```

### 8.4 OpenAPI Specification

```yaml
# api/openapi.yaml
openapi: 3.0.3
info:
  title: Cody API
  description: Autonomous development agent with Claude Code mastery
  version: 1.0.0

paths:
  /api/v1/skills:
    get:
      summary: List all skills
      responses:
        200:
          description: List of skills
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Skill'
    post:
      summary: Create a skill from spec
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/SkillSpec'
      responses:
        201:
          description: Skill created
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/GenerateResult'

  /api/v1/skills/wizard/start:
    post:
      summary: Start skill creation wizard
      responses:
        200:
          description: Wizard session started
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/WizardSession'

  /api/v1/skills/wizard/{session}/answer:
    post:
      summary: Answer current question
      parameters:
        - name: session
          in: path
          required: true
          schema:
            type: string
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                answer:
                  type: string
      responses:
        200:
          description: Answer accepted, next question returned
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/WizardSession'

  /api/v1/skills/{name}/test:
    post:
      summary: Run all tests for a skill
      parameters:
        - name: name
          in: path
          required: true
          schema:
            type: string
      responses:
        200:
          description: Test results
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TestResults'

  /api/v1/config/compose:
    post:
      summary: Generate optimal config for a task
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                task:
                  type: string
                  description: Task description
      responses:
        200:
          description: Composed configuration
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ClaudeCodeConfig'

  /api/v1/ws:
    get:
      summary: WebSocket for streaming operations
      description: |
        Connect via WebSocket, send operation request, receive progress updates.

        Request format:
        ```json
        {"operation": "test_skill", "params": {"name": "my-skill"}}
        ```

        Progress format:
        ```json
        {"type": "progress", "progress": {"phase": "behavioral", "percent": 0.5}}
        ```

        Complete format:
        ```json
        {"type": "complete", "results": {...}}
        ```

components:
  schemas:
    Skill:
      type: object
      properties:
        name:
          type: string
        description:
          type: string
        type:
          type: string
          enum: [technique, pattern, reference, discipline]
        tests:
          $ref: '#/components/schemas/SkillTests'

    SkillSpec:
      type: object
      required: [name, type, trigger_phrases, expected_behavior]
      properties:
        name:
          type: string
        type:
          type: string
        trigger_phrases:
          type: array
          items:
            type: string
        symptoms:
          type: array
          items:
            type: string
        baseline_behavior:
          type: string
        expected_behavior:
          type: string
        measurable_behavior:
          type: string
        rule:
          type: string
        pressure_scenarios:
          type: array
          items:
            type: string
        known_rationalizations:
          type: array
          items:
            type: string

    WizardSession:
      type: object
      properties:
        id:
          type: string
        current_question:
          $ref: '#/components/schemas/Question'
        progress:
          type: integer
        total:
          type: integer

    Question:
      type: object
      properties:
        id:
          type: string
        text:
          type: string
        type:
          type: string
          enum: [single_choice, multi_choice, free_text]
        options:
          type: array
          items:
            type: string
        rationale:
          type: string

    TestResults:
      type: object
      properties:
        skill_name:
          type: string
        behavioral:
          $ref: '#/components/schemas/BehavioralResults'
        determinism:
          $ref: '#/components/schemas/DeterminismResults'
        compliance:
          $ref: '#/components/schemas/ComplianceResults'
        overall_pass:
          type: boolean

    ClaudeCodeConfig:
      type: object
      properties:
        name:
          type: string
        skills:
          type: array
          items:
            type: object
        hooks:
          type: array
          items:
            type: object
        mcp_servers:
          type: array
          items:
            type: object
        plan_mode:
          type: object
```

### 8.5 gRPC Service Definition

For high-performance remote scenarios:

```protobuf
// api/proto/cody.proto
syntax = "proto3";

package cody.v1;

option go_package = "github.com/your-org/cody/api/proto/codyv1";

// =============================================================================
// Skill Service
// =============================================================================

service SkillService {
    rpc ListSkills(ListSkillsRequest) returns (ListSkillsResponse);
    rpc GetSkill(GetSkillRequest) returns (Skill);
    rpc CreateSkill(CreateSkillRequest) returns (CreateSkillResponse);

    // Wizard with streaming
    rpc SkillWizard(stream WizardMessage) returns (stream WizardMessage);
}

message Skill {
    string name = 1;
    string description = 2;
    SkillType type = 3;
    repeated string keywords = 4;
}

enum SkillType {
    SKILL_TYPE_UNSPECIFIED = 0;
    SKILL_TYPE_TECHNIQUE = 1;
    SKILL_TYPE_PATTERN = 2;
    SKILL_TYPE_REFERENCE = 3;
    SKILL_TYPE_DISCIPLINE = 4;
}

message WizardMessage {
    oneof message {
        WizardStart start = 1;
        WizardQuestion question = 2;
        WizardAnswer answer = 3;
        WizardComplete complete = 4;
    }
}

message WizardStart {}

message WizardQuestion {
    string id = 1;
    string text = 2;
    QuestionType type = 3;
    repeated string options = 4;
    string rationale = 5;
    int32 progress = 6;
    int32 total = 7;
}

enum QuestionType {
    QUESTION_TYPE_UNSPECIFIED = 0;
    QUESTION_TYPE_FREE_TEXT = 1;
    QUESTION_TYPE_SINGLE_CHOICE = 2;
    QUESTION_TYPE_MULTI_CHOICE = 3;
}

message WizardAnswer {
    string question_id = 1;
    string answer = 2;
}

message WizardComplete {
    string skill_path = 1;
    string test_path = 2;
}

// =============================================================================
// Test Service
// =============================================================================

service TestService {
    // Unary
    rpc TestSkill(TestSkillRequest) returns (TestResults);

    // Server streaming for progress
    rpc TestSkillStream(TestSkillRequest) returns (stream TestProgress);
}

message TestSkillRequest {
    string skill_name = 1;
    repeated TestType test_types = 2; // Empty = all
    int32 iterations = 3; // For determinism tests
}

enum TestType {
    TEST_TYPE_UNSPECIFIED = 0;
    TEST_TYPE_BEHAVIORAL = 1;
    TEST_TYPE_DETERMINISM = 2;
    TEST_TYPE_COMPLIANCE = 3;
}

message TestProgress {
    TestType current_type = 1;
    string phase = 2;
    string message = 3;
    float percent = 4;
}

message TestResults {
    string skill_name = 1;
    BehavioralResults behavioral = 2;
    DeterminismResults determinism = 3;
    ComplianceResults compliance = 4;
    bool overall_pass = 5;
}

message BehavioralResults {
    bool passed = 1;
    float baseline_rate = 2;
    float with_skill_rate = 3;
    float diff = 4;
    string summary = 5;
}

message DeterminismResults {
    bool passed = 1;
    float variance = 2;
    repeated ConsistencyResult checks = 3;
}

message ConsistencyResult {
    string name = 1;
    float variance = 2;
    bool passed = 3;
}

message ComplianceResults {
    bool passed = 1;
    float compliance_rate = 2;
    repeated string detected_rationalizations = 3;
}

// =============================================================================
// Config Service
// =============================================================================

service ConfigService {
    rpc ComposeConfig(ComposeConfigRequest) returns (ClaudeCodeConfig);
    rpc ValidateConfig(ClaudeCodeConfig) returns (ValidationResult);
}

message ComposeConfigRequest {
    string task_description = 1;
    repeated string languages = 2;
    string complexity = 3;
}

message ClaudeCodeConfig {
    string name = 1;
    string description = 2;
    repeated SkillRef skills = 3;
    repeated HookDef hooks = 4;
    repeated MCPServer mcp_servers = 5;
    PlanModeConfig plan_mode = 6;
}

message SkillRef {
    string name = 1;
    bool required = 2;
    int32 priority = 3;
}

message HookDef {
    string event = 1;
    string matcher = 2;
    string command = 3;
}

message MCPServer {
    string name = 1;
    string command = 2;
    repeated string args = 3;
}

message PlanModeConfig {
    bool enabled = 1;
    bool require_approval = 2;
}

message ValidationResult {
    bool valid = 1;
    repeated ValidationError errors = 2;
}

message ValidationError {
    string field = 1;
    string message = 2;
    string severity = 3;
}
```

---

## 9. CLI Architecture (Bubble Tea TUI)

### 8.1 Design Philosophy

Cody's CLI should feel like Claude Code:
- **Start it and go** - No complex setup, just `cody`
- **Beautiful terminal UI** - Not raw text, but styled interactive components
- **Guided workflows** - The system asks the right questions, you answer
- **Magic in background** - Complex architecture invisible to user

### 8.2 Command Structure

```
cody                           # Start interactive mode (like claude)
cody skill                     # Skill management
cody skill create              # Interactive skill creation wizard
cody skill test <name>         # Run all tests for a skill
cody skill list                # List available skills
cody skill edit <name>         # Edit existing skill
cody skill validate <name>     # Validate skill structure

cody feature                   # Feature development (existing cody)
cody feature run <config>      # Run autonomous feature dev
cody feature resume <id>       # Resume interrupted feature

cody config                    # Configuration composition
cody config compose <task>     # Generate optimal config for task
cody config validate           # Validate current config

cody test                      # Testing
cody test behavioral <skill>   # Run behavioral diff tests
cody test determinism <skill>  # Run determinism tests
cody test compliance <skill>   # Run compliance tests
cody test all <skill>          # Run all test types
```

### 8.3 Bubble Tea Architecture

```go
package tui

import (
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/spinner"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbles/viewport"
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

// AppState represents the current state of the TUI
type AppState int

const (
    StateHome AppState = iota
    StateSkillCreate
    StateSkillTest
    StateFeatureRun
    StateConfigCompose
)

// Model is the main Bubble Tea model
type Model struct {
    state       AppState
    width       int
    height      int

    // Sub-models for different screens
    home        HomeModel
    skillCreate SkillCreateModel
    skillTest   SkillTestModel
    featureRun  FeatureRunModel

    // Shared components
    spinner     spinner.Model
    err         error
}

func NewModel() Model {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    return Model{
        state:       StateHome,
        spinner:     s,
        home:        NewHomeModel(),
        skillCreate: NewSkillCreateModel(),
    }
}

func (m Model) Init() tea.Cmd {
    return tea.Batch(
        m.spinner.Tick,
        tea.EnterAltScreen,
    )
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }

    // Delegate to current state's model
    var cmd tea.Cmd
    switch m.state {
    case StateHome:
        m.home, cmd = m.home.Update(msg)
        if m.home.Selected != "" {
            m = m.transitionTo(m.home.Selected)
        }
    case StateSkillCreate:
        m.skillCreate, cmd = m.skillCreate.Update(msg)
    case StateSkillTest:
        m.skillTest, cmd = m.skillTest.Update(msg)
    }

    return m, cmd
}

func (m Model) View() string {
    switch m.state {
    case StateHome:
        return m.home.View(m.width, m.height)
    case StateSkillCreate:
        return m.skillCreate.View(m.width, m.height)
    case StateSkillTest:
        return m.skillTest.View(m.width, m.height)
    default:
        return "Unknown state"
    }
}
```

### 8.4 Skill Create Wizard

The interactive skill creation experience:

```go
package tui

import (
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"

    "github.com/your-org/cody/pkg/skillgen"
)

// SkillCreateModel handles the skill creation wizard
type SkillCreateModel struct {
    questionnaire *skillgen.SkillQuestionnaire
    currentQ      *skillgen.Question

    // UI state
    phase         CreatePhase
    textInput     textinput.Model
    choices       []string
    selectedIdx   int
    multiSelected map[int]bool

    // Progress
    answered      int
    total         int

    // Results
    spec          *skillgen.SkillSpec
    generating    bool
    generated     bool
    generatedPath string

    // Styling
    styles        SkillCreateStyles
}

type CreatePhase int

const (
    PhaseQuestions CreatePhase = iota
    PhaseReview
    PhaseGenerating
    PhaseComplete
    PhaseTesting
)

type SkillCreateStyles struct {
    Title       lipgloss.Style
    Question    lipgloss.Style
    Option      lipgloss.Style
    Selected    lipgloss.Style
    Help        lipgloss.Style
    Progress    lipgloss.Style
    Success     lipgloss.Style
    Error       lipgloss.Style
    Subtle      lipgloss.Style
    Box         lipgloss.Style
}

func NewSkillCreateStyles() SkillCreateStyles {
    return SkillCreateStyles{
        Title: lipgloss.NewStyle().
            Bold(true).
            Foreground(lipgloss.Color("205")).
            MarginBottom(1),
        Question: lipgloss.NewStyle().
            Bold(true).
            Foreground(lipgloss.Color("255")),
        Option: lipgloss.NewStyle().
            PaddingLeft(2),
        Selected: lipgloss.NewStyle().
            PaddingLeft(2).
            Foreground(lipgloss.Color("205")).
            Bold(true),
        Help: lipgloss.NewStyle().
            Foreground(lipgloss.Color("241")).
            MarginTop(1),
        Progress: lipgloss.NewStyle().
            Foreground(lipgloss.Color("241")),
        Success: lipgloss.NewStyle().
            Foreground(lipgloss.Color("82")).
            Bold(true),
        Error: lipgloss.NewStyle().
            Foreground(lipgloss.Color("196")).
            Bold(true),
        Subtle: lipgloss.NewStyle().
            Foreground(lipgloss.Color("241")).
            Italic(true),
        Box: lipgloss.NewStyle().
            Border(lipgloss.RoundedBorder()).
            BorderForeground(lipgloss.Color("62")).
            Padding(1, 2),
    }
}

func NewSkillCreateModel() SkillCreateModel {
    ti := textinput.New()
    ti.Placeholder = "Type your answer..."
    ti.CharLimit = 500
    ti.Width = 60

    q := skillgen.NewSkillQuestionnaire()

    return SkillCreateModel{
        questionnaire: q,
        currentQ:      q.NextQuestion(),
        textInput:     ti,
        multiSelected: make(map[int]bool),
        styles:        NewSkillCreateStyles(),
        total:         q.TotalQuestions(),
    }
}

func (m SkillCreateModel) Update(msg tea.Msg) (SkillCreateModel, tea.Cmd) {
    switch m.phase {
    case PhaseQuestions:
        return m.updateQuestions(msg)
    case PhaseReview:
        return m.updateReview(msg)
    case PhaseGenerating:
        return m.updateGenerating(msg)
    case PhaseComplete:
        return m.updateComplete(msg)
    case PhaseTesting:
        return m.updateTesting(msg)
    }
    return m, nil
}

func (m SkillCreateModel) updateQuestions(msg tea.Msg) (SkillCreateModel, tea.Cmd) {
    if m.currentQ == nil {
        // All questions answered, move to review
        m.phase = PhaseReview
        m.spec = m.questionnaire.GenerateSkillSpec()
        return m, nil
    }

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            // Submit answer
            answer := m.getAnswer()
            if answer != "" {
                m.questionnaire.Answer(m.currentQ.ID, answer)
                m.currentQ = m.questionnaire.NextQuestion()
                m.answered++
                m.resetInput()
            }
        case "up", "k":
            if m.currentQ.Type != skillgen.QuestionTypeFreeText {
                m.selectedIdx = max(0, m.selectedIdx-1)
            }
        case "down", "j":
            if m.currentQ.Type != skillgen.QuestionTypeFreeText {
                m.selectedIdx = min(len(m.currentQ.Options)-1, m.selectedIdx+1)
            }
        case " ":
            if m.currentQ.Type == skillgen.QuestionTypeMultiChoice {
                m.multiSelected[m.selectedIdx] = !m.multiSelected[m.selectedIdx]
            }
        default:
            if m.currentQ.Type == skillgen.QuestionTypeFreeText {
                var cmd tea.Cmd
                m.textInput, cmd = m.textInput.Update(msg)
                return m, cmd
            }
        }
    }

    return m, nil
}

func (m SkillCreateModel) View(width, height int) string {
    switch m.phase {
    case PhaseQuestions:
        return m.viewQuestions(width, height)
    case PhaseReview:
        return m.viewReview(width, height)
    case PhaseGenerating:
        return m.viewGenerating(width, height)
    case PhaseComplete:
        return m.viewComplete(width, height)
    case PhaseTesting:
        return m.viewTesting(width, height)
    }
    return ""
}

func (m SkillCreateModel) viewQuestions(width, height int) string {
    var b strings.Builder

    // Header
    b.WriteString(m.styles.Title.Render("✨ Create New Skill"))
    b.WriteString("\n\n")

    // Progress bar
    progress := float64(m.answered) / float64(m.total)
    progressBar := renderProgressBar(progress, 40)
    b.WriteString(m.styles.Progress.Render(
        fmt.Sprintf("Progress: %s %d/%d", progressBar, m.answered, m.total)))
    b.WriteString("\n\n")

    if m.currentQ == nil {
        return b.String()
    }

    // Question
    b.WriteString(m.styles.Question.Render(m.currentQ.Text))
    b.WriteString("\n")

    // Rationale (why we're asking)
    if m.currentQ.Rationale != "" {
        b.WriteString(m.styles.Subtle.Render("(" + m.currentQ.Rationale + ")"))
        b.WriteString("\n")
    }
    b.WriteString("\n")

    // Input based on question type
    switch m.currentQ.Type {
    case skillgen.QuestionTypeFreeText:
        b.WriteString(m.textInput.View())
        b.WriteString("\n")
        b.WriteString(m.styles.Help.Render("Press Enter to submit"))

    case skillgen.QuestionTypeSingleChoice:
        for i, opt := range m.currentQ.Options {
            cursor := "  "
            style := m.styles.Option
            if i == m.selectedIdx {
                cursor = "▸ "
                style = m.styles.Selected
            }
            b.WriteString(style.Render(cursor + opt))
            b.WriteString("\n")
        }
        b.WriteString(m.styles.Help.Render("↑/↓ to move, Enter to select"))

    case skillgen.QuestionTypeMultiChoice:
        for i, opt := range m.currentQ.Options {
            cursor := "  "
            checkbox := "○"
            style := m.styles.Option
            if m.multiSelected[i] {
                checkbox = "●"
            }
            if i == m.selectedIdx {
                cursor = "▸ "
                style = m.styles.Selected
            }
            b.WriteString(style.Render(fmt.Sprintf("%s%s %s", cursor, checkbox, opt)))
            b.WriteString("\n")
        }
        b.WriteString(m.styles.Help.Render("↑/↓ to move, Space to toggle, Enter to confirm"))
    }

    return m.styles.Box.Width(width - 4).Render(b.String())
}

func (m SkillCreateModel) viewReview(width, height int) string {
    var b strings.Builder

    b.WriteString(m.styles.Title.Render("📋 Review Your Skill"))
    b.WriteString("\n\n")

    // Show spec summary
    b.WriteString(m.styles.Question.Render("Name: "))
    b.WriteString(m.spec.Name)
    b.WriteString("\n")

    b.WriteString(m.styles.Question.Render("Type: "))
    b.WriteString(string(m.spec.Type))
    b.WriteString("\n\n")

    b.WriteString(m.styles.Question.Render("Triggers:"))
    b.WriteString("\n")
    for _, t := range m.spec.TriggerPhrases {
        b.WriteString("  • " + t + "\n")
    }
    b.WriteString("\n")

    b.WriteString(m.styles.Question.Render("Expected Behavior:"))
    b.WriteString("\n")
    b.WriteString("  " + m.spec.ExpectedBehavior)
    b.WriteString("\n\n")

    if m.spec.Rule != "" {
        b.WriteString(m.styles.Question.Render("Rule:"))
        b.WriteString("\n")
        b.WriteString("  " + m.spec.Rule)
        b.WriteString("\n\n")
    }

    b.WriteString(m.styles.Help.Render("Press Enter to generate, Esc to go back"))

    return m.styles.Box.Width(width - 4).Render(b.String())
}

func (m SkillCreateModel) viewGenerating(width, height int) string {
    return m.styles.Box.Width(width - 4).Render(
        m.styles.Title.Render("⚙️  Generating Skill...") + "\n\n" +
        m.spinner.View() + " Creating SKILL.md and test files...",
    )
}

func (m SkillCreateModel) viewComplete(width, height int) string {
    var b strings.Builder

    b.WriteString(m.styles.Success.Render("✓ Skill Created Successfully!"))
    b.WriteString("\n\n")

    b.WriteString(m.styles.Question.Render("Generated files:"))
    b.WriteString("\n")
    b.WriteString(fmt.Sprintf("  📄 %s/SKILL.md\n", m.generatedPath))
    b.WriteString(fmt.Sprintf("  🧪 %s/%s_test.go\n", m.generatedPath, m.spec.Name))
    if m.spec.NeedsScripts {
        b.WriteString(fmt.Sprintf("  📁 %s/scripts/\n", m.generatedPath))
    }
    if m.spec.NeedsReferences {
        b.WriteString(fmt.Sprintf("  📁 %s/references/\n", m.generatedPath))
    }
    b.WriteString("\n")

    b.WriteString(m.styles.Help.Render("Press T to run tests, Enter to finish, Q to quit"))

    return m.styles.Box.Width(width - 4).Render(b.String())
}

func renderProgressBar(progress float64, width int) string {
    filled := int(progress * float64(width))
    empty := width - filled

    bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
    return bar
}
```

### 8.5 Skill Test Runner TUI

Real-time test execution with live results:

```go
package tui

import (
    "github.com/charmbracelet/bubbles/progress"
    "github.com/charmbracelet/bubbletea"

    "github.com/your-org/cody/pkg/skilltest"
)

// SkillTestModel handles test execution UI
type SkillTestModel struct {
    skillName    string
    testRunner   *skilltest.Runner

    // State
    phase        TestPhase
    currentTest  string
    results      []TestResult

    // Progress
    progress     progress.Model
    totalTests   int
    completed    int

    // Live output
    liveOutput   []string
    viewport     viewport.Model

    styles       SkillTestStyles
}

type TestPhase int

const (
    TestPhaseIdle TestPhase = iota
    TestPhaseBehavioralBaseline  // Running without skill
    TestPhaseBehavioralWithSkill // Running with skill
    TestPhaseDeterminism         // Running N iterations
    TestPhaseCompliance          // Running pressure scenarios
    TestPhaseComplete
)

type TestResult struct {
    Name     string
    Type     string // "behavioral", "determinism", "compliance"
    Passed   bool
    Message  string
    Metrics  map[string]float64
}

func (m SkillTestModel) View(width, height int) string {
    var b strings.Builder

    // Header
    b.WriteString(m.styles.Title.Render(
        fmt.Sprintf("🧪 Testing: %s", m.skillName)))
    b.WriteString("\n\n")

    // Progress
    b.WriteString(m.progress.View())
    b.WriteString("\n")
    b.WriteString(m.styles.Subtle.Render(
        fmt.Sprintf("%d/%d tests complete", m.completed, m.totalTests)))
    b.WriteString("\n\n")

    // Current phase
    phaseNames := map[TestPhase]string{
        TestPhaseBehavioralBaseline:  "📊 Behavioral Diff: Running baseline (without skill)...",
        TestPhaseBehavioralWithSkill: "📊 Behavioral Diff: Running with skill...",
        TestPhaseDeterminism:         "🎯 Determinism: Running iterations...",
        TestPhaseCompliance:          "🛡️  Compliance: Running pressure scenarios...",
        TestPhaseComplete:            "✅ All tests complete!",
    }

    if phaseName, ok := phaseNames[m.phase]; ok {
        b.WriteString(m.styles.Question.Render(phaseName))
        b.WriteString("\n\n")
    }

    // Live output viewport
    b.WriteString(m.styles.Box.Render(m.viewport.View()))
    b.WriteString("\n")

    // Results summary (if any complete)
    if len(m.results) > 0 {
        b.WriteString("\n")
        b.WriteString(m.styles.Question.Render("Results:"))
        b.WriteString("\n")

        for _, r := range m.results {
            icon := "✓"
            style := m.styles.Success
            if !r.Passed {
                icon = "✗"
                style = m.styles.Error
            }
            b.WriteString(style.Render(fmt.Sprintf("  %s %s: %s", icon, r.Name, r.Message)))
            b.WriteString("\n")

            // Show metrics
            for k, v := range r.Metrics {
                b.WriteString(m.styles.Subtle.Render(
                    fmt.Sprintf("      %s: %.2f%%", k, v*100)))
                b.WriteString("\n")
            }
        }
    }

    return b.String()
}

// Messages for test progress
type TestProgressMsg struct {
    Phase    TestPhase
    Message  string
    Progress float64
}

type TestResultMsg struct {
    Result TestResult
}

type TestCompleteMsg struct{}

func (m SkillTestModel) runTests() tea.Cmd {
    return func() tea.Msg {
        // This runs in background, sends messages back to update UI

        // Behavioral diff tests
        m.testRunner.OnProgress(func(phase string, msg string, pct float64) {
            // Send progress message (would need channel/subscription)
        })

        results := m.testRunner.RunBehavioralDiff(m.skillName)

        return TestResultMsg{Result: TestResult{
            Name:    "behavioral_diff",
            Type:    "behavioral",
            Passed:  results.Passed,
            Message: results.Summary,
            Metrics: map[string]float64{
                "baseline_rate":   results.BaselineRate,
                "with_skill_rate": results.WithSkillRate,
                "diff":            results.Diff,
            },
        }}
    }
}
```

### 8.6 Main Entry Point

```go
package main

import (
    "fmt"
    "os"

    "github.com/charmbracelet/bubbletea"
    "github.com/spf13/cobra"

    "github.com/your-org/cody/internal/tui"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "cody",
        Short: "Autonomous development agent with Claude Code mastery",
        Run: func(cmd *cobra.Command, args []string) {
            // Default: launch interactive TUI
            runInteractive()
        },
    }

    // Skill commands
    skillCmd := &cobra.Command{
        Use:   "skill",
        Short: "Skill management",
    }

    skillCreateCmd := &cobra.Command{
        Use:   "create",
        Short: "Create a new skill interactively",
        Run: func(cmd *cobra.Command, args []string) {
            runSkillCreate()
        },
    }

    skillTestCmd := &cobra.Command{
        Use:   "test [skill-name]",
        Short: "Run tests for a skill",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            runSkillTest(args[0])
        },
    }

    skillListCmd := &cobra.Command{
        Use:   "list",
        Short: "List available skills",
        Run: func(cmd *cobra.Command, args []string) {
            runSkillList()
        },
    }

    skillCmd.AddCommand(skillCreateCmd, skillTestCmd, skillListCmd)
    rootCmd.AddCommand(skillCmd)

    // Feature commands (existing cody functionality)
    featureCmd := &cobra.Command{
        Use:   "feature",
        Short: "Feature development",
    }

    featureRunCmd := &cobra.Command{
        Use:   "run [config-file]",
        Short: "Run autonomous feature development",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            runFeature(args[0])
        },
    }

    featureCmd.AddCommand(featureRunCmd)
    rootCmd.AddCommand(featureCmd)

    // Config commands
    configCmd := &cobra.Command{
        Use:   "config",
        Short: "Configuration composition",
    }

    configComposeCmd := &cobra.Command{
        Use:   "compose [task-description]",
        Short: "Generate optimal Claude Code config for a task",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            runConfigCompose(args[0])
        },
    }

    configCmd.AddCommand(configComposeCmd)
    rootCmd.AddCommand(configCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func runInteractive() {
    p := tea.NewProgram(
        tui.NewModel(),
        tea.WithAltScreen(),
        tea.WithMouseCellMotion(),
    )

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

func runSkillCreate() {
    p := tea.NewProgram(
        tui.NewSkillCreateModel(),
        tea.WithAltScreen(),
    )

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

func runSkillTest(skillName string) {
    p := tea.NewProgram(
        tui.NewSkillTestModel(skillName),
        tea.WithAltScreen(),
    )

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

### 8.7 Visual Design System

```go
package tui

import "github.com/charmbracelet/lipgloss"

// Theme defines the visual language
var Theme = struct {
    // Colors
    Primary    lipgloss.Color
    Secondary  lipgloss.Color
    Success    lipgloss.Color
    Error      lipgloss.Color
    Warning    lipgloss.Color
    Muted      lipgloss.Color
    Background lipgloss.Color

    // Typography
    Title     lipgloss.Style
    Heading   lipgloss.Style
    Body      lipgloss.Style
    Code      lipgloss.Style
    Subtle    lipgloss.Style

    // Components
    Box         lipgloss.Style
    ActiveBox   lipgloss.Style
    Button      lipgloss.Style
    ActiveButton lipgloss.Style
    Input       lipgloss.Style
}{
    // Cody brand colors
    Primary:    lipgloss.Color("205"),  // Pink/magenta
    Secondary:  lipgloss.Color("62"),   // Purple
    Success:    lipgloss.Color("82"),   // Green
    Error:      lipgloss.Color("196"),  // Red
    Warning:    lipgloss.Color("214"),  // Orange
    Muted:      lipgloss.Color("241"),  // Gray
    Background: lipgloss.Color("236"),  // Dark gray

    Title: lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("205")).
        MarginBottom(1),

    Heading: lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("255")),

    Body: lipgloss.NewStyle().
        Foreground(lipgloss.Color("252")),

    Code: lipgloss.NewStyle().
        Foreground(lipgloss.Color("229")).
        Background(lipgloss.Color("236")).
        Padding(0, 1),

    Subtle: lipgloss.NewStyle().
        Foreground(lipgloss.Color("241")).
        Italic(true),

    Box: lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("62")).
        Padding(1, 2),

    ActiveBox: lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("205")).
        Padding(1, 2),

    Button: lipgloss.NewStyle().
        Foreground(lipgloss.Color("252")).
        Background(lipgloss.Color("62")).
        Padding(0, 2).
        MarginRight(1),

    ActiveButton: lipgloss.NewStyle().
        Foreground(lipgloss.Color("255")).
        Background(lipgloss.Color("205")).
        Bold(true).
        Padding(0, 2).
        MarginRight(1),

    Input: lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("241")).
        Padding(0, 1),
}

// Logo renders the Cody ASCII art logo
func Logo() string {
    logo := `
   ██████╗ ██████╗ ██████╗ ██╗   ██╗
  ██╔════╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
  ██║     ██║   ██║██║  ██║ ╚████╔╝
  ██║     ██║   ██║██║  ██║  ╚██╔╝
  ╚██████╗╚██████╔╝██████╔╝   ██║
   ╚═════╝ ╚═════╝ ╚═════╝    ╚═╝
    `
    return lipgloss.NewStyle().
        Foreground(Theme.Primary).
        Bold(true).
        Render(logo)
}

// StatusLine renders a status bar
func StatusLine(left, center, right string, width int) string {
    leftStyle := lipgloss.NewStyle().
        Foreground(Theme.Muted).
        Width(width / 3)
    centerStyle := lipgloss.NewStyle().
        Foreground(Theme.Primary).
        Bold(true).
        Width(width / 3).
        Align(lipgloss.Center)
    rightStyle := lipgloss.NewStyle().
        Foreground(Theme.Muted).
        Width(width / 3).
        Align(lipgloss.Right)

    return lipgloss.JoinHorizontal(
        lipgloss.Top,
        leftStyle.Render(left),
        centerStyle.Render(center),
        rightStyle.Render(right),
    )
}
```

### 8.8 User Experience Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│    $ cody skill create                                                       │
│                                                                              │
│    ┌────────────────────────────────────────────────────────────────────┐   │
│    │  ✨ Create New Skill                                                │   │
│    │                                                                     │   │
│    │  Progress: ████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 3/12            │   │
│    │                                                                     │   │
│    │  What type of skill is this?                                       │   │
│    │  (Determines testing strategy and structure)                       │   │
│    │                                                                     │   │
│    │  ▸ technique                                                       │   │
│    │    pattern                                                         │   │
│    │    reference                                                       │   │
│    │    discipline                                                      │   │
│    │                                                                     │   │
│    │  ↑/↓ to move, Enter to select                                     │   │
│    └────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

                                    ↓ After answering all questions...

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│    ┌────────────────────────────────────────────────────────────────────┐   │
│    │  📋 Review Your Skill                                               │   │
│    │                                                                     │   │
│    │  Name: tdd-for-api-handlers                                        │   │
│    │  Type: discipline                                                  │   │
│    │                                                                     │   │
│    │  Triggers:                                                         │   │
│    │    • "write an API handler"                                        │   │
│    │    • "implement endpoint"                                          │   │
│    │    • "create REST API"                                             │   │
│    │                                                                     │   │
│    │  Expected Behavior:                                                │   │
│    │    Write tests before implementation code                          │   │
│    │                                                                     │   │
│    │  Rule:                                                             │   │
│    │    Never write handler code before tests exist                     │   │
│    │                                                                     │   │
│    │  Press Enter to generate, Esc to go back                          │   │
│    └────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

                                    ↓ After generation...

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│    ┌────────────────────────────────────────────────────────────────────┐   │
│    │  ✓ Skill Created Successfully!                                     │   │
│    │                                                                     │   │
│    │  Generated files:                                                  │   │
│    │    📄 skills/tdd-for-api-handlers/SKILL.md                        │   │
│    │    🧪 skills/tdd-for-api-handlers/tdd_for_api_handlers_test.go    │   │
│    │                                                                     │   │
│    │  Press T to run tests, Enter to finish, Q to quit                 │   │
│    └────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

                                    ↓ After pressing T...

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│    🧪 Testing: tdd-for-api-handlers                                         │
│                                                                              │
│    ████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░                     │
│    2/6 tests complete                                                       │
│                                                                              │
│    📊 Behavioral Diff: Running with skill...                                │
│                                                                              │
│    ┌────────────────────────────────────────────────────────────────────┐   │
│    │  Run 3/5: Baseline - agent wrote code first (no TDD)               │   │
│    │  Run 3/5: With skill - agent wrote tests first ✓                   │   │
│    │  Calculating behavioral diff...                                    │   │
│    └────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│    Results:                                                                 │
│      ✓ behavioral_diff: Skill increases TDD compliance                     │
│          baseline_rate: 20.00%                                             │
│          with_skill_rate: 80.00%                                           │
│          diff: 60.00%                                                      │
│      ✓ determinism: Consistent approach across runs                        │
│          variance: 15.00%                                                  │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```
