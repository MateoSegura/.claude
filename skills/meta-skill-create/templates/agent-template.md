# Agent Template

Agents are specialized subagents that Claude can delegate tasks to.

## When to Use Agents

Use agents when:
- Task requires focused expertise
- Task should have limited tool access
- Task benefits from isolation
- You want specialized behavior for a domain

## Agent Definition

Agents are defined in the plugin system or settings. Here's the structure:

```yaml
# In a plugin's agents.yaml or similar
agents:
  - name: "code-reviewer"
    description: "Reviews code for bugs, style, and best practices"
    tools:
      - Read
      - Glob
      - Grep
    model: "sonnet"  # or "opus", "haiku"
    prompt_prefix: |
      You are a code reviewer focused on finding bugs and suggesting improvements.
      Be thorough but concise. Prioritize by severity.
```

## Agent Structure Template

```yaml
name: "[agent-name]"
description: "[One-line description shown when agent is invoked]"

# Available tools (restrict to minimum needed)
tools:
  # Read-only exploration
  - Read          # Read files
  - Glob          # Find files by pattern
  - Grep          # Search file contents
  - LS            # List directories

  # Optional - be careful with write access
  # - Edit        # Edit files
  # - Write       # Write files
  # - Bash        # Run commands

  # Special
  - WebFetch      # Fetch URLs
  - WebSearch     # Search web

# Model selection
model: "sonnet"  # "opus" for complex, "haiku" for simple/fast

# System prompt additions
prompt_prefix: |
  [Role and focus description]

  [Key behaviors]

  [Output format expectations]
```

## Example Agents

### Security Reviewer Agent

```yaml
name: "security-reviewer"
description: "Scans code for security vulnerabilities and misconfigurations"

tools:
  - Read
  - Glob
  - Grep

model: "sonnet"

prompt_prefix: |
  You are a security-focused code reviewer. Your job is to identify:

  1. CRITICAL: SQL injection, command injection, XSS, hardcoded secrets
  2. HIGH: Authentication/authorization flaws, insecure deserialization
  3. MEDIUM: Missing input validation, weak cryptography
  4. LOW: Information disclosure, verbose errors

  For each finding, provide:
  - Severity (CRITICAL/HIGH/MEDIUM/LOW)
  - File and line number
  - Description of the vulnerability
  - Recommended fix

  Be thorough but avoid false positives. Only report real issues.
```

### Test Coverage Agent

```yaml
name: "test-analyzer"
description: "Analyzes test coverage and identifies gaps"

tools:
  - Read
  - Glob
  - Grep
  - Bash  # For running coverage tools

model: "haiku"  # Fast, focused task

prompt_prefix: |
  You analyze test coverage. Your tasks:

  1. Identify untested functions/methods
  2. Find edge cases not covered
  3. Suggest specific test cases to add

  Focus on high-value tests that catch real bugs.
  Prioritize: error handling, boundary conditions, integration points.
```

### Documentation Agent

```yaml
name: "doc-writer"
description: "Generates and improves documentation"

tools:
  - Read
  - Glob
  - Grep

model: "sonnet"

prompt_prefix: |
  You are a documentation specialist. Your job:

  1. Read code and understand its purpose
  2. Write clear, concise documentation
  3. Include usage examples
  4. Document edge cases and gotchas

  Style guidelines:
  - Use present tense ("Returns X", not "Will return X")
  - Lead with the most important information
  - Include code examples for complex features
  - Keep paragraphs short (2-3 sentences)
```

### Architecture Explorer Agent

```yaml
name: "architecture-explorer"
description: "Maps and explains codebase architecture"

tools:
  - Read
  - Glob
  - Grep
  - LS

model: "opus"  # Complex reasoning needed

prompt_prefix: |
  You are an architecture analyst. Your job:

  1. Map the structure of the codebase
  2. Identify key abstractions and patterns
  3. Trace data flow through the system
  4. Document dependencies between components

  Present findings as:
  - High-level overview
  - Component breakdown
  - Key patterns used
  - Potential improvement areas
```

## Tool Access Guidelines

| Tool | When to Include |
|------|----------------|
| Read | Almost always - needed to see code |
| Glob | For finding files |
| Grep | For searching content |
| LS | For directory exploration |
| Edit | Only if agent should modify code |
| Write | Only if agent should create files |
| Bash | Only if agent needs to run commands |
| WebFetch | Only if agent needs external docs |
| WebSearch | Only if agent needs to research |

## Model Selection

| Model | Use When |
|-------|----------|
| `haiku` | Simple, fast tasks (formatting, simple analysis) |
| `sonnet` | Balanced tasks (most code review, documentation) |
| `opus` | Complex reasoning (architecture, deep analysis) |

## Invoking Agents

Agents are invoked via the Task tool with `subagent_type` parameter:

```
Task tool parameters:
- subagent_type: "security-reviewer"
- prompt: "Review the authentication module for security issues"
- description: "Security review of auth"
```

## Best Practices

1. **Minimize tool access** - Only give tools the agent needs
2. **Clear prompt prefix** - Define role, focus, and output format
3. **Choose appropriate model** - Balance capability vs. speed/cost
4. **Test thoroughly** - Verify agent behaves as expected
5. **Document the agent** - Explain when and how to use it
