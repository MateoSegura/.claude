# Quality Standards

Loaded during Phase 5 (Review). The bar that every extension must clear.

---

## The 5 Dimensions

Every extension is scored 1-5 on each dimension. All must be >= 4 to ship.

### 1. Routing (Will Claude find and invoke this correctly?)

**Score 1**: Generic description like "helps with coding" — Claude will either never invoke it or invoke it for everything.

**Score 3**: Decent description but overlaps with other skills, or is too broad. "Helps debug Zephyr applications" could trigger on any Zephyr question.

**Score 5**: Laser-specific trigger. "Analyze Zephyr RTOS UART serial logs to identify HardFault causes, stack overflows, and kernel panics in firmware debug sessions." Claude knows exactly when this applies.

**Checklist**:
- [ ] Description mentions the specific technology/domain
- [ ] Description mentions the specific action/workflow
- [ ] Description mentions the trigger context (when to use)
- [ ] Description does NOT overlap with existing extensions
- [ ] A developer reading just the description knows if this skill is for them

### 2. Architecture (Is context managed efficiently?)

**Score 1**: Everything dumped in SKILL.md. All rules, all examples, all references inline. Claude's context fills up before work begins.

**Score 3**: Files exist but SKILL.md loads too many of them, or loads them too early. Reference table exists but "Load When" is vague.

**Score 5**: SKILL.md is a thin router. Each file has a clear "Load When" trigger tied to a specific workflow phase. Hooks handle all repetitive automation. Agents run in isolated context. The main conversation stays lean.

**Checklist**:
- [ ] SKILL.md body is under 200 lines (the router, not the knowledge)
- [ ] Every file in the reference table has a specific "Load When" condition
- [ ] Repetitive actions (format, lint, validate) are hooks, not skill steps
- [ ] Heavy analysis tasks are delegated to agents with `context: fork` or Task tool
- [ ] Agent model selections are cost-appropriate (haiku for simple, sonnet for analysis)

### 3. Completeness (Does it actually work end-to-end?)

**Score 1**: Stubs, TODOs, placeholder text. "Add your rules here." Missing files referenced in SKILL.md.

**Score 3**: Core workflow works but edge cases are missing. No error handling. No guidance for when things go wrong.

**Score 5**: Every file exists. Every workflow step is actionable — Claude can follow it without guessing. Error cases are handled. The extension works on first use.

**Checklist**:
- [ ] Every file referenced in SKILL.md exists and has real content
- [ ] Zero instances of TODO, FIXME, TBD, "placeholder", "add here"
- [ ] Workflow steps are specific enough that Claude doesn't need to guess
- [ ] Error/failure paths are addressed ("if X fails, do Y")
- [ ] CLI commands include actual flags and arguments, not pseudocode
- [ ] For scaffolds: code compiles/runs as-is (no broken templates)

### 4. Prompt Quality (Are the instructions effective for Claude?)

**Score 1**: Vague instructions. "Analyze the code and provide feedback." No structure, no examples, no constraints.

**Score 3**: Clear instructions but missing structure. Claude will produce something but the output quality varies between invocations.

**Score 5**: Instructions use proven prompt engineering techniques. Output is consistent and high-quality across invocations.

**Checklist** (see prompt-engineering.md for full techniques):
- [ ] Agent prompts define a clear ROLE ("You are a security analyst specializing in...")
- [ ] Output format is explicitly specified (table, list, structured sections)
- [ ] Constraints state what NOT to do (reduces hallucination, scope creep)
- [ ] Complex tasks are decomposed into numbered steps
- [ ] Verification steps exist ("Before outputting, check that...")
- [ ] Where relevant, few-shot examples show the expected input → output

### 5. Engineering (Is this built to last?)

**Score 1**: Brittle assumptions, hardcoded paths, no error handling, runs on opus when haiku would do.

**Score 3**: Works but wasteful. Loads unnecessary files. Uses expensive models for simple tasks. No thought given to maintenance.

**Score 5**: Minimal context cost. Right model for each task. Graceful failure. Easy to extend. Follows marketplace conventions.

**Checklist**:
- [ ] File paths are relative, not absolute
- [ ] No hardcoded usernames, paths, or environment-specific values
- [ ] Model selection matches task complexity
- [ ] Hook commands are fast (< 2 seconds)
- [ ] Agent tool sets are minimal (no Write access for read-only agents)
- [ ] Plugin can be installed in any project without modification

---

## Common Failure Modes

### The Knowledge Dump
SKILL.md contains 500+ lines of rules, references, and examples inline. Fix: break into separate files with lazy loading.

### The Vague Router
SKILL.md says "read the appropriate reference file" without specifying which file for which phase. Fix: explicit file-to-phase mapping in the reference table.

### The Overpowered Agent
Agent has Read, Write, Edit, Bash, Grep, Glob when it only needs Read and Grep. Fix: minimum viable tool set.

### The Context Hog
Skill loads 10 files in Phase 1 "just in case." Fix: defer file loading to the phase where the content is actually needed.

### The Template Trap
Scaffold code has `// YOUR CODE HERE` markers that Claude interprets literally. Fix: use descriptive markers that Claude can understand contextually, or provide complete examples.

### The Missing Constraint
Agent prompt says "review the code" but doesn't say what to look for or what to ignore. Fix: explicit evaluation criteria and out-of-scope declaration.

### The Generic Description
Description is "helps with development" — overlaps with 50 other skills. Fix: include technology, action, and context in the description.

---

## The Gut Check

After scoring all 5 dimensions, ask:

> If I installed this extension from someone else's marketplace, would I keep it or uninstall it after the first use?

If the answer is "uninstall" — it's not ready to ship.
