---
description: Update an existing extension
---

# /update-extension

Update any extension in this knowledge base.

## Usage

```
/update-extension [name]
```

If no name provided, lists all extensions for selection.

## Process

1. **Discover** all extensions in the repo

2. **Select** which to update (if not specified)

3. **Analyze** current state:
   - Read all files
   - Identify patterns and rules
   - Check dependencies

4. **Research** updates:
   - Web search for latest patterns
   - Check library changelogs
   - Look for breaking changes

5. **Plan** changes:
   - What to add/modify/remove
   - Impact analysis
   - Testing strategy

6. **Implement** with approval

7. **Test** the updated extension

## Update Types

### Library Version Update

1. Check installed version vs skill version
2. Search for changelog
3. Update affected patterns
4. Test with new version

### Add Missing Patterns

1. Search codebase for uncovered patterns
2. Research best practices
3. Add new rules/scaffolds
4. Update quick reference

### Fix Incorrect Examples

1. Identify broken examples
2. Find correct patterns
3. Update and verify

### Expand Scope

1. Identify gaps
2. Research new areas
3. Add new sections
4. Maintain consistency

## Validation

After update:
- [ ] All existing tests pass
- [ ] New functionality tested
- [ ] No regressions
- [ ] Documentation updated
