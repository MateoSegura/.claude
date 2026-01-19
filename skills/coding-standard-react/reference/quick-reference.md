# Quick Reference - All Rules

Complete table of all rules with IDs, descriptions, and tiers.

## Rule Tiers

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Rules by Category

### Styling & Naming (RCT-STY-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-STY-001 | Use PascalCase for Component Names | Critical |
| RCT-STY-002 | Use camelCase for Prop Names | Required |
| RCT-STY-003 | Prefix Custom Hooks with `use` | Critical |

### Components (RCT-CMP-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-CMP-001 | Use Functional Components | Critical |
| RCT-CMP-002 | One Component Per File | Required |
| RCT-CMP-003 | Prefer Composition Over Prop Drilling | Required |
| RCT-CMP-004 | Keep Components Focused and Small | Required |

### Hooks (RCT-HKS-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-HKS-001 | Call Hooks at Top Level Only | Critical |
| RCT-HKS-002 | Include All Dependencies in useEffect | Critical |
| RCT-HKS-003 | Clean Up Effects Properly | Critical |
| RCT-HKS-004 | Extract Reusable Logic into Custom Hooks | Required |
| RCT-HKS-005 | Keep Custom Hooks Focused | Recommended |

### State Management (RCT-STT-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-STT-001 | Keep State Close to Where It's Used | Required |
| RCT-STT-002 | Use useReducer for Complex State Logic | Required |
| RCT-STT-003 | Derive State When Possible | Required |
| RCT-STT-004 | Avoid Redundant State | Required |

### Props (RCT-PRP-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-PRP-001 | Destructure Props in Function Signature | Recommended |
| RCT-PRP-002 | Use TypeScript for Prop Types | Critical |
| RCT-PRP-003 | Avoid Spreading Props Blindly | Required |
| RCT-PRP-004 | Use Children for Composition | Recommended |

### Events (RCT-EVT-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-EVT-001 | Use Consistent Event Handler Naming | Required |
| RCT-EVT-002 | Avoid Inline Arrow Functions (Performance Critical) | Recommended |

### Rendering (RCT-RND-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-RND-001 | Use Appropriate Conditional Rendering Patterns | Required |
| RCT-RND-002 | Avoid Short-Circuit with Numbers | Critical |
| RCT-RND-003 | Use Stable Keys for Lists | Critical |

### Performance (RCT-PRF-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-PRF-001 | Memoize Expensive Computations | Required |
| RCT-PRF-002 | Use React.memo for Pure Components | Required |
| RCT-PRF-003 | Avoid Creating Objects in Render | Required |
| RCT-PRF-004 | Use useCallback for Function Props | Required |

### Testing (RCT-TST-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-TST-001 | Query by Role and Accessible Name | Critical |
| RCT-TST-002 | Use userEvent Over fireEvent | Required |
| RCT-TST-003 | Test Behavior, Not Implementation | Critical |
| RCT-TST-004 | Use waitFor for Async Assertions | Required |

### Accessibility (RCT-A11-*)

| ID | Rule | Tier |
|----|------|------|
| RCT-A11-001 | Use Semantic HTML Elements | Critical |
| RCT-A11-002 | Provide Text Alternatives for Images | Critical |
| RCT-A11-003 | Ensure Keyboard Navigation | Critical |
| RCT-A11-004 | Use ARIA Attributes Correctly | Required |
| RCT-A11-005 | Maintain Sufficient Color Contrast | Required |

---

## Summary by Tier

### Critical Rules (15 total)

- RCT-STY-001, RCT-STY-003
- RCT-CMP-001
- RCT-HKS-001, RCT-HKS-002, RCT-HKS-003
- RCT-PRP-002
- RCT-RND-002, RCT-RND-003
- RCT-TST-001, RCT-TST-003
- RCT-A11-001, RCT-A11-002, RCT-A11-003

### Required Rules (19 total)

- RCT-STY-002
- RCT-CMP-002, RCT-CMP-003, RCT-CMP-004
- RCT-HKS-004
- RCT-STT-001, RCT-STT-002, RCT-STT-003, RCT-STT-004
- RCT-PRP-003
- RCT-EVT-001
- RCT-RND-001
- RCT-PRF-001, RCT-PRF-002, RCT-PRF-003, RCT-PRF-004
- RCT-TST-002, RCT-TST-004
- RCT-A11-004, RCT-A11-005

### Recommended Rules (4 total)

- RCT-HKS-005
- RCT-PRP-001, RCT-PRP-004
- RCT-EVT-002

---

## Quick Decision Guide

### Should I use `useMemo`?

- Yes: Expensive computations (sorting, filtering large arrays)
- Yes: Objects/arrays passed to memoized children
- No: Simple string concatenation or arithmetic
- No: Primitive values

### Should I use `useCallback`?

- Yes: Callbacks passed to memoized children
- Yes: Callbacks used as useEffect dependencies
- No: Simple handlers for non-memoized elements

### Should I use `React.memo`?

- Yes: Component renders often with same props
- Yes: Component is expensive to render
- No: Component receives new props on every render
- No: Component is simple and renders quickly

### What key should I use for lists?

- Best: Unique ID from data (`item.id`)
- OK: Composite key (`${item.category}-${item.slug}`)
- Last resort: Index (only if list never changes)
