# Code Review Standard

> **Version**: 1.0.0
> **Status**: Active
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes code review practices that ensure:

- **Quality**: Catching defects before they reach production
- **Knowledge Sharing**: Spreading expertise across the team
- **Consistency**: Maintaining uniform code standards
- **Mentorship**: Growing engineers through constructive feedback
- **Collective Ownership**: Building shared responsibility for the codebase

### 1.2 Scope

This standard applies to:
- All pull requests to protected branches
- All production code changes
- Configuration and infrastructure-as-code changes
- Test code modifications

### 1.3 Audience

- All software engineers submitting code changes
- All engineers performing code reviews
- Tech leads establishing team practices

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | CRITICAL | Blocking | PR cannot be merged |
| **Required** | REQUIRED | Blocking | Must address before approval |
| **Recommended** | RECOMMENDED | Non-blocking | Address at author's discretion |

Each rule includes:
- **Rule ID**: Unique identifier (e.g., `REV-001`)
- **Tier**: Enforcement level
- **Rationale**: Why this rule exists

---

## 3. Review Responsibilities

### 3.1 Author Responsibilities

#### REV-001: Complete Self-Review Before Requesting Review REQUIRED

**Rationale**: Self-review catches obvious issues and demonstrates respect for reviewers' time.

Before requesting review, the author must:

1. **Re-read the entire diff** line by line
2. **Run all tests** and verify they pass
3. **Run linters/formatters** and address all warnings
4. **Test the change manually** for the happy path and key edge cases
5. **Check for debugging artifacts** (console.log, print statements, TODO comments)
6. **Verify the build succeeds** in CI

#### REV-002: Provide Complete PR Description REQUIRED

**Rationale**: Context enables effective review and serves as documentation.

Every PR description must include:

```markdown
## Summary
[What this change does and why]

## Changes
- [Bullet point list of specific changes]

## Testing
[How this was tested - manual steps, new tests, etc.]

## Related
- [Links to issues, tickets, or related PRs]
```

For complex changes, also include:
- **Design decisions**: Why this approach was chosen
- **Alternatives considered**: What was rejected and why
- **Risk assessment**: What could go wrong

#### REV-003: Respond to All Review Comments REQUIRED

**Rationale**: Unaddressed comments create ambiguity and technical debt.

Authors must:
- Respond to every comment (even if just acknowledging)
- Explain if declining a suggestion
- Mark conversations as resolved when addressed
- Re-request review after addressing feedback

### 3.2 Reviewer Responsibilities

#### REV-004: Respond Within 24 Business Hours REQUIRED

**Rationale**: Timely reviews maintain development velocity.

| PR Size | Target Response Time |
|---------|---------------------|
| Small (<100 lines) | 4 hours |
| Medium (100-400 lines) | 8 hours |
| Large (400+ lines) | 24 hours |

If unable to review in time:
- Comment that you've seen it and when you'll review
- Suggest another reviewer if you're blocked

#### REV-005: Provide Thorough Review REQUIRED

**Rationale**: Superficial reviews miss defects and waste the review process.

Reviewers must:
- **Read every line** of changed code
- **Understand the context** (read surrounding code if needed)
- **Run the code mentally** or actually
- **Consider edge cases** and failure modes
- **Check test coverage** for the changes

#### REV-006: Maintain Constructive Tone CRITICAL

**Rationale**: Respectful communication builds trust and psychological safety.

Reviewers must:
- Critique the code, not the person
- Assume good intent
- Ask questions rather than make accusations
- Acknowledge good work and improvements
- Be specific and actionable

---

## 4. Review Criteria

### 4.1 Correctness

**Primary question**: Does the code work as intended?

| Check | What to Look For |
|-------|------------------|
| Logic | Off-by-one errors, null handling, boundary conditions |
| State | Race conditions, stale data, inconsistent state |
| Integration | API contracts, data format assumptions |
| Error paths | Exception handling, error propagation |

#### REV-007: Verify Correctness CRITICAL

**Rationale**: Incorrect code is worse than no code.

Reviewers must verify:
- Logic handles all documented requirements
- Edge cases are considered
- Error conditions are handled appropriately
- The change doesn't break existing functionality

### 4.2 Design

**Primary question**: Is the code well-structured?

| Check | What to Look For |
|-------|------------------|
| Modularity | Single responsibility, clear boundaries |
| Abstraction | Appropriate levels, not over-engineered |
| Coupling | Minimal dependencies between components |
| Extensibility | Easy to modify for future requirements |

#### REV-008: Evaluate Design Appropriateness REQUIRED

**Rationale**: Poor design compounds technical debt over time.

Consider:
- Does this belong in this location?
- Is the abstraction level appropriate?
- Will this be maintainable in 6 months?
- Does this align with existing patterns?

### 4.3 Complexity

**Primary question**: Can another engineer understand this quickly?

| Complexity Indicator | Threshold |
|---------------------|-----------|
| Function length | >40 lines warrants scrutiny |
| Cyclomatic complexity | >10 requires justification |
| Nesting depth | >3 levels needs refactoring |
| Parameters | >5 suggests object needed |

#### REV-009: Challenge Unnecessary Complexity REQUIRED

**Rationale**: Complex code hides bugs and slows future development.

Ask:
- Is there a simpler way to achieve this?
- Can this be broken into smaller pieces?
- Would a different approach be clearer?

### 4.4 Tests

**Primary question**: Are the tests sufficient and well-written?

| Test Quality Aspect | Requirement |
|--------------------|-------------|
| Coverage | All new code paths tested |
| Independence | Tests don't depend on each other |
| Clarity | Test name describes scenario and expectation |
| Assertions | Specific assertions, not just "no error" |

#### REV-010: Require Adequate Test Coverage REQUIRED

**Rationale**: Untested code is unverified code.

Verify:
- New functionality has corresponding tests
- Bug fixes include regression tests
- Edge cases are covered
- Tests actually test the changed code

### 4.5 Naming

**Primary question**: Do names clearly communicate intent?

| Element | Good Name | Bad Name |
|---------|-----------|----------|
| Function | `calculateMonthlyPayment` | `calc` |
| Variable | `activeUserCount` | `n` |
| Boolean | `isEnabled`, `hasPermission` | `flag`, `status` |
| Constant | `MAX_RETRY_ATTEMPTS` | `MAGIC_NUMBER` |

#### REV-011: Ensure Clear Naming REQUIRED

**Rationale**: Good names reduce cognitive load and serve as documentation.

Names should:
- Reveal intent without needing comments
- Be pronounceable and searchable
- Avoid abbreviations (except ubiquitous ones)
- Match domain terminology

### 4.6 Comments

**Primary question**: Do comments add value without stating the obvious?

| Good Comments | Bad Comments |
|--------------|--------------|
| Explain "why" for non-obvious decisions | Restate what code does |
| Document complex algorithms | Commented-out code |
| Note external constraints | Outdated information |
| Warn about non-obvious gotchas | Obvious explanations |

#### REV-012: Review Comment Quality RECOMMENDED

**Rationale**: Comments should enhance understanding, not clutter code.

Comments should:
- Explain the "why", not the "what"
- Be accurate and up-to-date
- Not compensate for unclear code
- Include references for complex algorithms

### 4.7 Style

**Primary question**: Does the code follow established conventions?

#### REV-013: Enforce Style Consistency REQUIRED

**Rationale**: Consistent style reduces cognitive load and diffs.

Verify:
- Code passes all linter checks
- Formatting matches project standards
- Import ordering follows conventions
- No style inconsistencies with surrounding code

### 4.8 Documentation

**Primary question**: Is documentation accurate and complete?

| Documentation Type | When Required |
|-------------------|---------------|
| API documentation | All public interfaces |
| README updates | User-facing changes |
| Architecture docs | Design changes |
| Changelog entries | Release-worthy changes |

#### REV-014: Verify Documentation Updates REQUIRED

**Rationale**: Outdated documentation is worse than no documentation.

Check:
- Public APIs have clear documentation
- README reflects any setup changes
- Migration guides for breaking changes
- Inline docs updated for behavior changes

---

## 5. Review Process

### 5.1 Response Time Expectations

| Stage | Expected Time | Maximum Time |
|-------|--------------|--------------|
| Initial review | 8 hours | 24 hours |
| Follow-up review | 4 hours | 12 hours |
| Final approval | 2 hours | 8 hours |

**Exceptions**:
- Complex architectural changes: up to 48 hours
- Cross-team reviews: add 24 hours
- During on-call rotations: delegate to available reviewer

### 5.2 Review Iterations

#### REV-015: Limit Review Cycles RECOMMENDED

**Rationale**: Excessive iterations indicate process or communication issues.

| Scenario | Expected Iterations |
|----------|-------------------|
| Routine change | 1-2 |
| Complex feature | 2-3 |
| Architecture change | 3-4 |

If exceeding limits:
- Sync meeting to align on direction
- Pair programming session
- Break PR into smaller pieces

### 5.3 Handling Disagreements

When reviewer and author disagree:

1. **Clarify understanding**: Ensure both parties understand each other's position
2. **Reference standards**: Check if existing standards address the issue
3. **Discuss tradeoffs**: Document pros/cons of each approach
4. **Seek third opinion**: Involve another senior engineer
5. **Escalate if needed**: Tech lead makes final call
6. **Document decision**: Record rationale for future reference

#### REV-016: Resolve Disagreements Constructively REQUIRED

**Rationale**: Unresolved disagreements block progress and damage relationships.

Rules of engagement:
- Focus on technical merits, not personal preference
- Be willing to compromise on style, not correctness
- Default to team conventions when no clear winner
- Accept that "good enough" is often good enough

### 5.4 Approval Criteria

#### When to Approve (LGTM)

The change:
- Meets all CRITICAL and REQUIRED rules
- Passes all automated checks
- Has adequate test coverage
- Documentation is updated
- No blocking concerns remain

#### When to Request Changes

The change has:
- Correctness issues or bugs
- Security vulnerabilities
- Missing required tests
- Incomplete error handling
- Violations of critical standards

#### When to Comment (No Block)

The change has:
- Style suggestions
- Alternative approaches worth considering
- Minor improvements for future
- Educational observations

### 5.5 LGTM with Comments

Reviewers may approve while leaving non-blocking feedback:

```
LGTM with minor suggestions.

nit: Consider renaming `x` to `count` for clarity.
suggestion: This could be simplified with a map/filter pattern.

These are non-blocking - merge at your discretion.
```

---

## 6. Comment Guidelines

### 6.1 Comment Prefixes

Use prefixes to clarify intent and priority:

| Prefix | Meaning | Blocking? |
|--------|---------|-----------|
| `blocking:` | Must be addressed before merge | Yes |
| `question:` | Seeking clarification | Depends |
| `suggestion:` | Consider this alternative | No |
| `nit:` | Minor style/preference issue | No |
| `note:` | FYI, no action needed | No |
| `praise:` | Positive feedback | No |

### 6.2 Writing Effective Comments

#### REV-017: Provide Actionable Feedback REQUIRED

**Rationale**: Vague feedback wastes time and creates frustration.

**Structure of a good comment**:
1. **What**: Identify the specific issue
2. **Why**: Explain why it matters
3. **How**: Suggest a solution or alternative

### 6.3 Example Comments

#### Good Comments

```
blocking: This doesn't handle the case where `user` is null, which can happen
for anonymous requests. Consider adding a null check or using optional chaining.

Example:
  const name = user?.name ?? 'Anonymous';
```

```
suggestion: This loop could be replaced with a filter/map chain, which would
be more idiomatic and avoid the mutable accumulator:

  const activeUsers = users
    .filter(u => u.isActive)
    .map(u => u.name);
```

```
question: I'm not familiar with this API - does `fetchAll()` return an empty
array or null when there are no results? The error handling assumes non-null.
```

```
nit: Minor naming suggestion - `getData` could be more specific like
`fetchUserProfile` to indicate what data is being retrieved.
```

```
praise: Nice refactor here! The extraction of this logic into a separate
function makes the flow much clearer.
```

#### Bad Comments (Avoid These)

| Bad Comment | Problem | Better Version |
|-------------|---------|----------------|
| "This is wrong" | No explanation or guidance | "This returns the wrong type - expected `User[]` but returning `User`. See line 45 for the expected signature." |
| "I wouldn't do it this way" | Personal preference without rationale | "suggestion: A factory pattern might be cleaner here because it would allow easier testing and extension." |
| "???" | Unclear what's being asked | "question: I'm confused about why we need to call `refresh()` here - the data should already be current from line 30." |
| "Fix this" | Not actionable | "blocking: This comparison uses `==` which will fail for null values. Use `===` or add explicit null handling." |
| "Looks fine I guess" | Not constructive approval | "LGTM - logic is sound and tests cover the key scenarios." |

### 6.4 Receiving Feedback

For authors receiving feedback:

1. **Assume positive intent**: Reviewers want to help
2. **Ask clarifying questions**: If feedback is unclear
3. **Explain your reasoning**: If you disagree, share context
4. **Thank reviewers**: They spent time helping you improve
5. **Don't take it personally**: Feedback is about code, not you

---

## 7. Review Size

### 7.1 PR Size Guidelines

| Size | Lines Changed | Review Time | Defect Detection |
|------|--------------|-------------|------------------|
| Small | <100 | 15-30 min | ~90% |
| Medium | 100-400 | 30-60 min | ~70% |
| Large | 400-1000 | 60-120 min | ~50% |
| Too Large | >1000 | N/A | <30% |

#### REV-018: Keep PRs Small REQUIRED

**Rationale**: Large PRs have lower review quality and longer cycle times.

**Guidelines**:
- Target <400 lines per PR
- Absolute maximum of 1000 lines (except generated code)
- If larger, must justify or split

### 7.2 Breaking Up Large Changes

Strategies for splitting large changes:

| Strategy | When to Use | Example |
|----------|-------------|---------|
| **Vertical slicing** | Feature development | Auth: model -> API -> UI (3 PRs) |
| **Horizontal slicing** | Cross-cutting changes | Refactor module A, then B, then C |
| **Preparatory refactoring** | Before new features | Extract interface (PR 1) -> implement (PR 2) |
| **Feature flags** | Incomplete features | Ship disabled code, enable separately |

#### REV-019: Split Large Changes Appropriately REQUIRED

**Rationale**: Reviewable chunks enable thorough review.

Before submitting a large PR:
1. Can this be a refactoring + feature?
2. Can layers be submitted separately?
3. Can it be behind a feature flag?
4. Is any part independently useful?

### 7.3 When Large PRs Are Acceptable

| Scenario | Justification | Mitigation |
|----------|---------------|------------|
| Generated code | Machine-produced | Separate from manual changes |
| Mass rename/refactor | Single logical change | Automated tool + spot check |
| Vendor updates | External dependency | Focus review on integration |
| Configuration | Infrastructure files | Document key changes |

---

## 8. Review Checklist

### 8.1 Quick Review Checklist

Use this checklist for every review:

#### Before Reviewing
- [ ] PR description is complete and clear
- [ ] CI checks are passing
- [ ] PR is appropriately sized

#### Correctness
- [ ] Code does what the description claims
- [ ] Edge cases are handled
- [ ] Error handling is appropriate
- [ ] No obvious bugs or logic errors

#### Design
- [ ] Code is in the right location
- [ ] Abstractions are appropriate
- [ ] No unnecessary dependencies added
- [ ] Consistent with existing patterns

#### Quality
- [ ] Names are clear and descriptive
- [ ] No unnecessary complexity
- [ ] Comments explain "why" not "what"
- [ ] No code duplication

#### Testing
- [ ] New code has tests
- [ ] Tests cover edge cases
- [ ] Tests are readable and maintainable
- [ ] Tests actually test the change

#### Security
- [ ] No sensitive data exposure
- [ ] Input validation present
- [ ] No injection vulnerabilities
- [ ] Dependencies are trusted

#### Documentation
- [ ] Public APIs documented
- [ ] README updated if needed
- [ ] Breaking changes documented
- [ ] Inline docs accurate

### 8.2 Extended Checklist (Complex Changes)

Additional checks for architectural or security-sensitive changes:

#### Architecture
- [ ] Fits within system architecture
- [ ] Scaling implications considered
- [ ] Backwards compatibility maintained
- [ ] Migration path documented

#### Performance
- [ ] No N+1 queries
- [ ] Appropriate caching
- [ ] Resource limits considered
- [ ] No memory leaks

#### Observability
- [ ] Logging is appropriate
- [ ] Metrics added if needed
- [ ] Errors are traceable
- [ ] Health checks updated

---

## 9. Building Review Culture

### 9.1 Principles

1. **Reviews are conversations, not gates**: Collaborate to improve code
2. **Everyone's code gets reviewed**: No exceptions for seniority
3. **Learning goes both ways**: Reviewers learn from authors too
4. **Speed and quality balance**: Neither blocked nor rubber-stamped
5. **Psychological safety first**: Safe to make mistakes and learn

### 9.2 Recognizing Good Practices

Celebrate:
- Thorough, constructive reviews
- Authors who respond gracefully to feedback
- Reviews that catch significant issues
- Knowledge sharing through comments

### 9.3 Anti-Patterns to Avoid

| Anti-Pattern | Why It's Harmful | Alternative |
|--------------|------------------|-------------|
| Rubber stamping | Misses defects, signals reviews don't matter | Take time to actually review |
| Nitpick paralysis | Blocks progress on minor issues | Use `nit:` prefix, don't block |
| Gatekeeping | Creates bottlenecks, damages trust | Distribute review responsibility |
| Seagull reviewing | Drive-by criticism without context | Understand before commenting |
| Ghost reviewers | Assigned but never respond | Reassign or ping for status |

---

## 10. Tooling and Automation

### 10.1 Automated Checks (Pre-Review)

These should pass before human review:

| Check | Purpose |
|-------|---------|
| Linting | Style consistency |
| Formatting | Code formatting |
| Unit tests | Basic correctness |
| Type checking | Type safety |
| Security scanning | Known vulnerabilities |
| Build | Compilation success |

### 10.2 Review Assignment

Configure automatic reviewer assignment:
- CODEOWNERS for critical paths
- Round-robin for team distribution
- Expertise matching when possible

### 10.3 Review Metrics (Use Carefully)

| Metric | What It Measures | Caution |
|--------|------------------|---------|
| Time to first review | Responsiveness | Don't sacrifice quality for speed |
| Review iterations | Alignment efficiency | Some iteration is healthy |
| PR size | Reviewability | Context matters |
| Comments per PR | Engagement | More isn't always better |

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| REV-001 | Complete Self-Review Before Requesting Review | REQUIRED |
| REV-002 | Provide Complete PR Description | REQUIRED |
| REV-003 | Respond to All Review Comments | REQUIRED |
| REV-004 | Respond Within 24 Business Hours | REQUIRED |
| REV-005 | Provide Thorough Review | REQUIRED |
| REV-006 | Maintain Constructive Tone | CRITICAL |
| REV-007 | Verify Correctness | CRITICAL |
| REV-008 | Evaluate Design Appropriateness | REQUIRED |
| REV-009 | Challenge Unnecessary Complexity | REQUIRED |
| REV-010 | Require Adequate Test Coverage | REQUIRED |
| REV-011 | Ensure Clear Naming | REQUIRED |
| REV-012 | Review Comment Quality | RECOMMENDED |
| REV-013 | Enforce Style Consistency | REQUIRED |
| REV-014 | Verify Documentation Updates | REQUIRED |
| REV-015 | Limit Review Cycles | RECOMMENDED |
| REV-016 | Resolve Disagreements Constructively | REQUIRED |
| REV-017 | Provide Actionable Feedback | REQUIRED |
| REV-018 | Keep PRs Small | REQUIRED |
| REV-019 | Split Large Changes Appropriately | REQUIRED |

---

## Appendix B: Comment Templates

### Approval Templates

**Standard approval**:
```
LGTM! Code is correct, well-tested, and follows our standards.
```

**Approval with suggestions**:
```
LGTM with minor suggestions below. Merge at your discretion.
```

**Conditional approval**:
```
LGTM once CI passes. No other concerns.
```

### Request Changes Templates

**Blocking issue**:
```
blocking: [Issue description]

This needs to be addressed because [reason].

Suggested fix:
[code or approach]
```

**Multiple issues**:
```
A few things need addressing before this can merge:

1. blocking: [Issue 1]
2. blocking: [Issue 2]
3. suggestion: [Non-blocking suggestion]

Happy to discuss any of these!
```

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- Google Engineering Practices: Code Review Guidelines
- Microsoft Code Review Best Practices
- Thoughtbot Code Review Guide
- SmartBear Best Practices for Code Review
