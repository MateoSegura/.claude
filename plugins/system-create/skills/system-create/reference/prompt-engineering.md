# Prompt Engineering for Claude Code Extensions

Loaded during Phase 4 (Implementation) and Phase 5 (Review). Techniques that measurably improve Claude's output quality in skills and agents.

---

## Core Principles

### 1. Specificity Beats Length

A 10-line prompt with exact instructions outperforms a 100-line prompt with vague guidance.

```
BAD:  "Review the code carefully and provide detailed feedback on any issues."
GOOD: "Check each function for: (1) unchecked error returns, (2) missing null guards
       on pointer dereference, (3) mutex locks without corresponding unlocks.
       For each issue found, output: file:line, severity (critical/warning), fix."
```

### 2. Structure Beats Prose

Numbered steps, tables, and explicit output formats produce consistent results. Prose paragraphs produce variable results.

```
BAD:  "Analyze the device tree and make sure everything looks correct and follows
       our conventions. Pay attention to node naming and property values."

GOOD: "## Device Tree Review Steps
       1. Check node names match pattern: <type>@<unit-address>
       2. Check `compatible` property exists on every non-root node
       3. Check `status` property is 'okay' or 'disabled' (not 'ok')
       4. Check `reg` property size matches #address-cells and #size-cells

       ## Output Format
       | Node | Issue | Expected | Actual |
       |------|-------|----------|--------|"
```

### 3. Constraints Reduce Hallucination

Telling Claude what NOT to do is often more effective than telling it what to do.

```
BAD:  "Generate a Zephyr device tree overlay."

GOOD: "Generate a Zephyr device tree overlay.
       Constraints:
       - Use ONLY nodes that exist in the base device tree for this board
       - Do NOT invent register addresses — look them up in the datasheet
       - Do NOT add properties that aren't in the Zephyr binding yaml
       - If unsure about a value, mark it with /* VERIFY: reason */ comment"
```

---

## Techniques

### Role Definition

Set the agent's identity and expertise at the top of the prompt. This anchors Claude's behavior.

```markdown
You are a firmware security auditor with expertise in embedded C,
Zephyr RTOS, and MISRA-C:2012 compliance. Your role is to identify
security vulnerabilities in embedded firmware code.
```

**When to use**: Every agent prompt. Always.

**Key rule**: The role should be specific enough that Claude knows what it is AND is not. "You are a code reviewer" is weak. "You are a code reviewer specializing in memory safety for embedded C on ARM Cortex-M" is strong.

### Decomposition

Break complex tasks into numbered steps with clear dependencies.

```markdown
## Process

1. Read the source file and identify all public functions
2. For each public function:
   a. Check if all parameters are validated before use
   b. Check if the return value communicates errors
   c. Check if resources acquired in the function are released on all paths
3. Compile findings into the output table
4. Sort by severity (critical first)
```

**When to use**: Any task with more than 2 steps. Unnumbered instructions get reordered or skipped.

### Output Specification

Define exactly what the output looks like. Ambiguous output requirements = inconsistent results.

```markdown
## Output Format

Return a JSON object:
{
  "status": "pass" | "fail",
  "issues": [
    {
      "file": "src/main.c",
      "line": 42,
      "severity": "critical" | "warning" | "info",
      "rule": "C-MEM-001",
      "message": "Allocation return value not checked",
      "fix": "Add null check after k_malloc"
    }
  ],
  "summary": "3 critical, 1 warning, 0 info"
}
```

**When to use**: Every agent that returns results to a skill. The skill needs to parse the output.

### Verification Steps

Add self-check instructions before Claude outputs its final result.

```markdown
## Before Submitting

Verify:
- [ ] Every file path you referenced actually exists (you read it, not guessed)
- [ ] Every line number is accurate (re-read the file to confirm)
- [ ] No findings are duplicates
- [ ] Severity ratings are justified (critical = crash/security, warning = bug, info = style)
```

**When to use**: Any task where accuracy matters more than speed. Especially code review agents and analysis tasks.

### Progressive Disclosure

Load information incrementally, not all at once.

```markdown
## Workflow

### Phase 1: Understand
Read the project's CMakeLists.txt and prj.conf to understand the build configuration.

### Phase 2: Analyze
NOW read [reference/build-flags.md](reference/build-flags.md) and compare against the project config.

### Phase 3: Report
NOW read [reference/common-issues.md](reference/common-issues.md) to check if any known issues apply.
```

**When to use**: Skills with reference files. Loading all references upfront wastes context on content that might not be needed.

### Few-Shot Examples

Show Claude what good output looks like with 1-2 concrete examples.

```markdown
## Example

Input: `k_mutex_lock(&mtx, K_FOREVER);` without error check
Output:
| File | Line | Rule | Issue | Fix |
|------|------|------|-------|-----|
| src/sensor.c | 87 | C-ERR-001 | `k_mutex_lock` return value unchecked | Wrap in `if (k_mutex_lock(...) != 0) { LOG_ERR(...); return -EAGAIN; }` |
```

**When to use**: When the output format is non-obvious, or when Claude needs to understand the level of detail expected. 1-2 examples is optimal — more adds context cost without improving quality.

### Negative Examples

Show what BAD output looks like and why it's bad.

```markdown
## Do NOT produce output like this:

❌ "The code has some potential issues with error handling."
   Why bad: vague, no file/line, no actionable fix.

❌ "Line 42: Consider using a different approach."
   Why bad: no specific recommendation, "consider" is not actionable.

✅ "src/main.c:42 — k_malloc return unchecked. Add: if (ptr == NULL) { return -ENOMEM; }"
   Why good: specific file, line, issue, and copy-pasteable fix.
```

**When to use**: When Claude tends to produce vague or wishy-washy output for a particular task type. Negative examples are more effective than positive examples for fixing specific failure modes.

### Scope Boundaries

Explicitly define what is in-scope and out-of-scope.

```markdown
## Scope

IN SCOPE:
- Memory safety issues (buffer overflow, use-after-free, double-free)
- Concurrency issues (race conditions, deadlocks, ISR violations)
- Error handling (unchecked returns, missing cleanup)

OUT OF SCOPE:
- Code style (formatting, naming) — handled by clang-format hook
- Performance optimization — separate skill
- Documentation quality — separate skill
```

**When to use**: Every agent prompt. Without scope boundaries, agents expand their review to cover everything, producing low-confidence findings across many categories instead of high-confidence findings in their specialty.

### Confidence Thresholds

When agents score findings, set explicit thresholds.

```markdown
## Scoring

Rate each finding 0-100 for confidence:
- 90-100: Certain bug/vulnerability. You can point to the exact line and explain the failure mode.
- 70-89: Likely issue. Pattern matches known problems but you can't rule out intentional design.
- 50-69: Possible issue. Suspicious pattern but context might justify it.
- Below 50: Do not report. Speculation wastes the user's time.

Only include findings with confidence >= 70.
```

**When to use**: Code review agents, security scanners, any agent that produces findings that a human will act on. False positives destroy trust faster than false negatives.

---

## Anti-Patterns

### The Wall of Text
Prompt is 3000 words of unstructured prose. Claude gets lost in the middle. Fix: numbered steps, tables, headers.

### The Infinite Scope
"Review everything and report all issues." Claude produces 50 low-quality findings. Fix: scope boundaries + confidence thresholds.

### The Missing Role
No role definition. Claude defaults to generic assistant behavior. Fix: specific role with domain expertise.

### The Implicit Format
"Provide your analysis." Claude outputs a different format every time. Fix: explicit output specification with example.

### The Redundant Instruction
Same instruction repeated 3 times in different words "for emphasis." Wastes context. Fix: say it once, clearly.

### The Hedge Invitation
"If you're not sure, just do your best." Invites low-confidence guessing. Fix: "If unsure about a value, mark it with a VERIFY comment and explain what needs checking."
