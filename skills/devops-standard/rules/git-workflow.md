---
description: Git workflow rules - branching, commits, PRs, releases
---

# Git Workflow Standard

> **Version**: 1.0.0
> **Status**: Active
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes Git workflow conventions for all development activities. It ensures:

- **Traceability**: Every change links to a ticket or documented reason
- **Quality**: Code reaches production only through controlled gates
- **Collaboration**: Clear processes for team code review and integration
- **Reliability**: Protected branches maintain stable, deployable states

### 1.2 Scope

This standard applies to:

- All Git repositories in the organization
- All contributors (employees, contractors, and automated systems)
- All branches, commits, and merge operations

### 1.3 Audience

- Software engineers committing code
- Tech leads managing repositories
- DevOps engineers configuring CI/CD
- Code reviewers approving changes

### 1.4 Relationship to Industry Standards

| Standard | Relationship |
|----------|--------------|
| [Conventional Commits](https://www.conventionalcommits.org/) | **Base** - Commit message format |
| [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow) | **Base** - Branching model foundation |
| [Semantic Versioning](https://semver.org/) | **Supplementary** - Release versioning |

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Merge blocked |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Advisory | Fix encouraged |

---

## 3. Branching Strategy

### 3.1 Branch Hierarchy

```
main (protected)
 |
 +-- release/v1.2.0 (protected, when applicable)
 |
 +-- feature/PROJ-123-user-authentication
 |
 +-- fix/PROJ-456-null-pointer-crash
 |
 +-- hotfix/PROJ-789-security-patch
```

### 3.2 Branch Types

| Branch Type | Pattern | Base Branch | Merges To | Purpose |
|-------------|---------|-------------|-----------|---------|
| Main | `main` | - | - | Production-ready code |
| Feature | `feature/TICKET-description` | `main` | `main` | New functionality |
| Bugfix | `fix/TICKET-description` | `main` | `main` | Non-critical bug fixes |
| Hotfix | `hotfix/TICKET-description` | `main` | `main` + `release/*` | Critical production fixes |
| Release | `release/vX.Y.Z` | `main` | `main` | Release stabilization |
| Experimental | `experiment/description` | `main` | - | Proof of concept (never merged) |

### 3.3 Branch Naming Rules

#### GIT-001: Branch names must follow naming convention :red_circle:

**Rationale**: Consistent naming enables automation, improves discoverability, and links changes to tracking systems.

```bash
# Correct
feature/PROJ-123-add-user-authentication
fix/PROJ-456-resolve-null-pointer
hotfix/PROJ-789-patch-xss-vulnerability
release/v2.1.0
experiment/graphql-migration-poc

# Incorrect
feature/add-auth                    # Missing ticket number
PROJ-123-new-feature                # Missing type prefix
feature/PROJ-123                    # Missing description
Feature/PROJ-123-add-auth           # Uppercase type
feature/PROJ-123_add_auth           # Underscores instead of hyphens
feature/proj-123-add-auth           # Lowercase ticket prefix
```

**Naming Rules**:
- Use lowercase for type prefix
- Ticket ID in UPPERCASE (e.g., `PROJ-123`, `BUG-456`)
- Description in lowercase with hyphens
- Keep descriptions concise (2-4 words)
- No spaces, underscores, or special characters

#### GIT-002: Feature branches must be short-lived :yellow_circle:

**Rationale**: Long-lived branches accumulate merge conflicts and diverge from main, increasing integration risk.

| Timeframe | Status | Action Required |
|-----------|--------|-----------------|
| < 5 days | Healthy | Continue work |
| 5-10 days | Warning | Plan to merge or split |
| > 10 days | Stale | Must justify or close |

---

## 4. Commit Messages

### 4.1 Conventional Commits Format

All commits must follow the Conventional Commits specification:

```
<type>(<scope>): <subject>

[optional body]

[optional footer(s)]
```

### 4.2 Commit Types

| Type | Description | Example |
|------|-------------|---------|
| `feat` | New feature | `feat(auth): add OAuth2 login support` |
| `fix` | Bug fix | `fix(api): resolve null pointer in user lookup` |
| `docs` | Documentation only | `docs(readme): update installation instructions` |
| `style` | Formatting, no code change | `style(lint): fix whitespace violations` |
| `refactor` | Code change without feature/fix | `refactor(db): extract connection pooling logic` |
| `test` | Adding/updating tests | `test(auth): add unit tests for token validation` |
| `chore` | Maintenance tasks | `chore(deps): upgrade axios to v1.6.0` |
| `perf` | Performance improvement | `perf(query): optimize database index usage` |
| `ci` | CI/CD changes | `ci(github): add code coverage reporting` |
| `build` | Build system changes | `build(docker): optimize multi-stage build` |
| `revert` | Revert previous commit | `revert: feat(auth): add OAuth2 login support` |

### 4.3 Subject Line Rules

#### GIT-003: Subject line must follow format requirements :red_circle:

**Rationale**: Consistent subject lines enable automated changelog generation and improve git log readability.

| Rule | Requirement |
|------|-------------|
| Length | Maximum 50 characters |
| Mood | Imperative ("add" not "added" or "adds") |
| Capitalization | Lowercase after colon |
| Punctuation | No period at end |
| Content | What the commit does, not what you did |

```bash
# Correct
feat(auth): add password reset functionality
fix(api): handle empty response gracefully
docs(api): document rate limiting headers
refactor(db): extract query builder logic

# Incorrect
feat(auth): Added password reset functionality      # Past tense
feat(auth): Add password reset functionality.       # Trailing period
feat(auth): Add Password Reset Functionality        # Title case
feat: add password reset functionality              # Missing scope
feat(auth): add password reset functionality for users who forgot their passwords  # Too long
```

### 4.4 Commit Body Rules

#### GIT-004: Commit body should explain what and why :green_circle:

**Rationale**: The body provides context that helps future developers understand the reasoning behind changes.

| Rule | Requirement |
|------|-------------|
| Blank line | Required between subject and body |
| Line width | Wrap at 72 characters |
| Content | Explain what changed and why, not how |

```bash
# Good commit message with body
feat(auth): add rate limiting to login endpoint

Implement rate limiting to prevent brute force attacks on the login
endpoint. Limits are set to 5 attempts per minute per IP address.

This addresses security audit finding SEC-2024-003.

Refs: PROJ-789
```

### 4.5 Commit Footer Rules

#### GIT-005: Use footers for metadata and references :green_circle:

**Rationale**: Footers provide machine-readable metadata for automation and issue tracking.

```bash
# Breaking change footer
feat(api)!: change authentication response format

BREAKING CHANGE: The /auth/token endpoint now returns a JSON object
instead of a plain token string. All API clients must be updated.

# Issue references
fix(upload): resolve file corruption on large uploads

Fixes: PROJ-456
Refs: PROJ-123, PROJ-124

# Co-authorship
feat(dashboard): add real-time metrics display

Co-authored-by: Jane Doe <jane.doe@example.com>
```

### 4.6 Commit Message Examples

#### Good Examples

```bash
# Simple feature
feat(cart): add quantity validation for cart items

# Bug fix with context
fix(auth): prevent session fixation on login

Regenerate session ID after successful authentication to prevent
session fixation attacks. Previous session data is preserved.

Fixes: SEC-2024-007

# Breaking change
feat(api)!: remove deprecated v1 endpoints

BREAKING CHANGE: All /api/v1/* endpoints have been removed.
Clients must migrate to /api/v2/* endpoints.

See migration guide: docs/migration/v1-to-v2.md

# Chore with scope
chore(deps): upgrade React from 18.2 to 18.3

Security patch for CVE-2024-XXXXX.
```

#### Bad Examples

```bash
# Too vague
fix: bug fix

# No type
updated the login page

# Past tense, too long, period at end
Fixed the bug that was causing the application to crash when users tried to upload files larger than 10MB.

# Multiple unrelated changes
feat: add login page, fix database connection, update readme

# Meaningless message
wip

# Commit just for the sake of committing
misc changes
```

---

## 5. Pull Request Process

### 5.1 PR Requirements

#### GIT-006: All changes to protected branches must go through PRs :red_circle:

**Rationale**: PRs provide a checkpoint for code review, automated testing, and documentation.

| Requirement | Enforcement |
|-------------|-------------|
| PR required | No direct pushes to `main` or `release/*` |
| Template required | PR must use standard template |
| Description required | PR must explain changes |
| Linked issue | PR should reference ticket(s) |

### 5.2 PR Template

```markdown
## Summary

<!-- Briefly describe what this PR does -->

## Type of Change

- [ ] Feature (new functionality)
- [ ] Bug fix (non-breaking fix)
- [ ] Breaking change (fix or feature causing existing functionality to change)
- [ ] Documentation update
- [ ] Refactoring (no functional changes)
- [ ] Performance improvement
- [ ] Test addition or update
- [ ] Chore (maintenance, dependencies)

## Related Issues

<!-- Link to related issues: Fixes #123, Refs #456 -->

## Changes Made

<!-- List the key changes in this PR -->

-
-
-

## Testing

<!-- Describe how this was tested -->

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing performed

## Screenshots (if applicable)

<!-- Add screenshots for UI changes -->

## Checklist

- [ ] My code follows the project's coding standards
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published

## Additional Notes

<!-- Any additional context or notes for reviewers -->
```

### 5.3 PR Title Convention

#### GIT-007: PR titles must match commit message format :yellow_circle:

**Rationale**: PR titles become merge commit messages; consistency aids changelog generation.

```bash
# Correct
feat(auth): add two-factor authentication
fix(api): resolve race condition in connection pool
docs(readme): add deployment instructions

# Incorrect
Add two-factor authentication          # Missing type
feat: add 2FA                          # Missing scope, abbreviation
PROJ-123: Add auth feature             # Ticket in title (belongs in body)
```

### 5.4 Review Requirements

#### GIT-008: PRs require minimum number of approvals :red_circle:

**Rationale**: Multiple reviewers catch more issues and spread knowledge across the team.

| Target Branch | Required Approvals | Required Reviewers |
|--------------|-------------------|-------------------|
| `main` | 1 | At least one code owner |
| `release/*` | 2 | Code owner + tech lead |

### 5.5 CI Requirements

#### GIT-009: All CI checks must pass before merge :red_circle:

**Rationale**: Automated checks prevent introducing regressions and maintain code quality.

| Check | Blocking | Description |
|-------|----------|-------------|
| Build | Yes | Code compiles without errors |
| Unit tests | Yes | All tests pass |
| Lint | Yes | No linting errors |
| Security scan | Yes | No critical vulnerabilities |
| Coverage | No | Coverage does not decrease |

### 5.6 Merge Strategy

#### GIT-010: Use squash merge for feature branches :yellow_circle:

**Rationale**: Squash merging creates a clean, linear history with one commit per feature.

| Branch Type | Merge Strategy | Rationale |
|-------------|----------------|-----------|
| Feature | Squash and merge | Clean history, single commit per feature |
| Hotfix | Squash and merge | Single traceable fix |
| Release | Merge commit | Preserve release branch history |

**Squash Merge Process**:
1. All commits are combined into one
2. PR title becomes commit subject
3. PR description becomes commit body
4. Original commits are preserved in PR for reference

---

## 6. Branch Hygiene

### 6.1 Branch Lifecycle

#### GIT-011: Delete branches after merge :yellow_circle:

**Rationale**: Stale branches clutter the repository and create confusion about active work.

```bash
# Automatic deletion (configure in GitHub/GitLab)
# Enable "Automatically delete head branches" in repository settings

# Manual deletion after merge
git branch -d feature/PROJ-123-add-auth        # Local
git push origin --delete feature/PROJ-123-add-auth  # Remote
```

### 6.2 Stale Branch Policy

#### GIT-012: Address stale branches within 30 days :yellow_circle:

**Rationale**: Abandoned branches waste resources and create confusion.

| Age | Status | Action |
|-----|--------|--------|
| 0-14 days | Active | No action |
| 15-30 days | Aging | Owner notified |
| 31-60 days | Stale | Requires justification |
| 60+ days | Abandoned | Deleted with notice |

### 6.3 Rebase vs Merge Policy

#### GIT-013: Rebase feature branches before PR :green_circle:

**Rationale**: Rebasing creates a linear history and ensures the feature branch includes latest main changes.

```bash
# Update feature branch with latest main
git checkout feature/PROJ-123-add-auth
git fetch origin
git rebase origin/main

# Resolve any conflicts, then force push (only for feature branches)
git push --force-with-lease origin feature/PROJ-123-add-auth
```

**Rules**:
- Rebase feature branches onto main before requesting review
- Use `--force-with-lease` (not `--force`) when pushing rebased branches
- Never rebase shared branches (main, release/*)
- Never rebase after PR is approved (request re-review instead)

---

## 7. Protected Branches

### 7.1 Main Branch Protection

#### GIT-014: Main branch must be protected :red_circle:

**Rationale**: The main branch represents production-ready code and must be guarded against accidental or unauthorized changes.

| Protection Rule | Setting | Rationale |
|-----------------|---------|-----------|
| Require PR | Enabled | No direct pushes |
| Required reviews | 1+ | Peer validation |
| Dismiss stale reviews | Enabled | New changes require re-review |
| Require status checks | Enabled | CI must pass |
| Require up-to-date branch | Enabled | Must include latest main |
| Include administrators | Enabled | No exceptions |
| Restrict who can push | Code owners | Limit merge authority |
| Allow force pushes | Disabled | Preserve history |
| Allow deletions | Disabled | Prevent accidental deletion |

### 7.2 Release Branch Protection

#### GIT-015: Release branches must be protected :red_circle:

**Rationale**: Release branches stabilize code for production deployment.

| Protection Rule | Setting |
|-----------------|---------|
| Require PR | Enabled |
| Required reviews | 2 |
| Require status checks | Enabled |
| Allow force pushes | Disabled |
| Allow deletions | After release complete |

### 7.3 Branch Protection Configuration

```yaml
# Example GitHub branch protection (via API or settings)
protection_rules:
  main:
    required_pull_request_reviews:
      required_approving_review_count: 1
      dismiss_stale_reviews: true
      require_code_owner_reviews: true
    required_status_checks:
      strict: true
      contexts:
        - "build"
        - "test"
        - "lint"
        - "security-scan"
    enforce_admins: true
    restrictions:
      users: []
      teams: ["maintainers"]
    allow_force_pushes: false
    allow_deletions: false
```

---

## 8. Hotfix Process

### 8.1 Hotfix Workflow

#### GIT-016: Hotfixes follow expedited process :red_circle:

**Rationale**: Critical production issues require a fast but controlled fix process.

```
1. Create hotfix branch from main
   git checkout main
   git pull origin main
   git checkout -b hotfix/PROJ-789-security-patch

2. Implement minimal fix
   # Make focused changes only

3. Test thoroughly
   # Run full test suite
   # Perform manual verification

4. Create PR with "hotfix" label
   # Expedited review (15 minute SLA)
   # Requires 1 approval from on-call

5. Merge to main
   # Deploy immediately after merge

6. Cherry-pick to release branches (if applicable)
   git checkout release/v1.2.0
   git cherry-pick <hotfix-commit-sha>
   git push origin release/v1.2.0
```

### 8.2 Hotfix Criteria

| Severity | Examples | Response Time |
|----------|----------|---------------|
| Critical (P0) | Security breach, data loss, complete outage | Immediate (< 1 hour) |
| High (P1) | Major feature broken, significant user impact | Same day (< 4 hours) |
| Medium (P2) | Minor feature issues, workaround available | Normal process |

---

## 9. Release Process

### 9.1 Release Branch Workflow

#### GIT-017: Release branches stabilize code for deployment :yellow_circle:

**Rationale**: Release branches allow bug fixes while main continues receiving new features.

```
1. Create release branch from main
   git checkout main
   git pull origin main
   git checkout -b release/v1.2.0

2. Only bug fixes allowed on release branch
   # No new features
   # Cherry-pick or target fixes to release branch

3. Update version numbers
   # Update package.json, version.go, etc.

4. Create release tag
   git tag -a v1.2.0 -m "Release v1.2.0"
   git push origin v1.2.0

5. Merge release branch back to main
   # Capture any release-specific fixes
```

### 9.2 Version Tagging

#### GIT-018: Use semantic versioning for releases :yellow_circle:

**Rationale**: Semantic versioning communicates the impact of changes to consumers.

```bash
# Version format: vMAJOR.MINOR.PATCH

v1.0.0    # Initial release
v1.1.0    # New features, backward compatible
v1.1.1    # Bug fixes only
v2.0.0    # Breaking changes

# Pre-release versions
v1.2.0-alpha.1
v1.2.0-beta.1
v1.2.0-rc.1
```

---

## 10. Tooling Configuration

### 10.1 Commit Message Validation

```yaml
# .commitlintrc.yaml
extends:
  - '@commitlint/config-conventional'

rules:
  type-enum:
    - 2
    - always
    - [feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert]
  type-case:
    - 2
    - always
    - lower-case
  subject-case:
    - 2
    - always
    - lower-case
  subject-max-length:
    - 2
    - always
    - 50
  body-max-line-length:
    - 2
    - always
    - 72
  header-max-length:
    - 2
    - always
    - 72
```

### 10.2 Pre-commit Hooks

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.5.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
        args: ['--maxkb=1000']
      - id: check-merge-conflict
      - id: detect-private-key

  - repo: https://github.com/compilerla/conventional-pre-commit
    rev: v3.0.0
    hooks:
      - id: conventional-pre-commit
        stages: [commit-msg]
```

### 10.3 Git Hooks Setup

```bash
#!/bin/bash
# .git/hooks/commit-msg (or via husky/pre-commit)

# Validate commit message format
commit_msg_file=$1
commit_msg=$(cat "$commit_msg_file")

# Pattern: type(scope): subject
pattern='^(feat|fix|docs|style|refactor|test|chore|perf|ci|build|revert)(\([a-z0-9-]+\))?: .{1,50}$'

if ! echo "$commit_msg" | head -1 | grep -qE "$pattern"; then
    echo "ERROR: Invalid commit message format"
    echo "Expected: type(scope): subject (max 50 chars)"
    echo "Types: feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert"
    exit 1
fi
```

### 10.4 Branch Name Validation

```bash
#!/bin/bash
# .git/hooks/pre-push

branch=$(git rev-parse --abbrev-ref HEAD)

# Skip validation for main and release branches
if [[ "$branch" == "main" ]] || [[ "$branch" =~ ^release/ ]]; then
    exit 0
fi

# Pattern: type/TICKET-description
pattern='^(feature|fix|hotfix|experiment)/[A-Z]+-[0-9]+-[a-z0-9-]+$'

if ! echo "$branch" | grep -qE "$pattern"; then
    echo "ERROR: Invalid branch name: $branch"
    echo "Expected: type/TICKET-description"
    echo "Example: feature/PROJ-123-add-user-auth"
    exit 1
fi
```

---

## 11. Code Review Checklist

Quick reference for reviewers:

### Branch and Commit
- [ ] Branch name follows convention (`type/TICKET-description`)
- [ ] Commit messages follow Conventional Commits format
- [ ] No merge commits in feature branch (clean rebase)
- [ ] Single logical change per PR

### PR Quality
- [ ] PR title matches commit message format
- [ ] PR description explains what and why
- [ ] Related issue is linked
- [ ] Appropriate reviewers assigned

### CI Status
- [ ] All CI checks pass
- [ ] No decrease in test coverage
- [ ] No new security vulnerabilities

### Merge Readiness
- [ ] Branch is up-to-date with main
- [ ] Required approvals obtained
- [ ] No unresolved conversations

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| GIT-001 | Branch names must follow naming convention | :red_circle: |
| GIT-002 | Feature branches must be short-lived | :yellow_circle: |
| GIT-003 | Subject line must follow format requirements | :red_circle: |
| GIT-004 | Commit body should explain what and why | :green_circle: |
| GIT-005 | Use footers for metadata and references | :green_circle: |
| GIT-006 | All changes to protected branches must go through PRs | :red_circle: |
| GIT-007 | PR titles must match commit message format | :yellow_circle: |
| GIT-008 | PRs require minimum number of approvals | :red_circle: |
| GIT-009 | All CI checks must pass before merge | :red_circle: |
| GIT-010 | Use squash merge for feature branches | :yellow_circle: |
| GIT-011 | Delete branches after merge | :yellow_circle: |
| GIT-012 | Address stale branches within 30 days | :yellow_circle: |
| GIT-013 | Rebase feature branches before PR | :green_circle: |
| GIT-014 | Main branch must be protected | :red_circle: |
| GIT-015 | Release branches must be protected | :red_circle: |
| GIT-016 | Hotfixes follow expedited process | :red_circle: |
| GIT-017 | Release branches stabilize code for deployment | :yellow_circle: |
| GIT-018 | Use semantic versioning for releases | :yellow_circle: |

---

## Appendix B: Glossary

| Term | Definition |
|------|------------|
| **Conventional Commits** | A specification for adding human and machine readable meaning to commit messages |
| **Feature branch** | A branch created to develop a specific feature or user story |
| **Force push** | Overwriting remote branch history (use `--force-with-lease` for safety) |
| **Hotfix** | An urgent fix for a critical production issue |
| **Protected branch** | A branch with rules preventing direct pushes and requiring reviews |
| **Rebase** | Rewriting commit history to apply changes on top of another branch |
| **Squash merge** | Combining all commits from a branch into a single commit when merging |

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- [Conventional Commits Specification](https://www.conventionalcommits.org/)
- [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
- [Semantic Versioning](https://semver.org/)
- [Git Branching Strategies](https://www.atlassian.com/git/tutorials/comparing-workflows)
- [How to Write a Git Commit Message](https://cbea.ms/git-commit/)
