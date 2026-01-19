# Code Review Checklist

Use this checklist during code reviews to ensure compliance with the React coding standard.

## Component Structure

- [ ] Components use functional syntax, not classes
- [ ] One component per file (unless small, tightly coupled helpers)
- [ ] Component names are PascalCase
- [ ] File names match component names
- [ ] Components are focused and reasonably sized (<300 lines)

**Related Rules**: RCT-CMP-001, RCT-CMP-002, RCT-CMP-003, RCT-CMP-004, RCT-STY-001

---

## Hooks Usage

- [ ] Hooks called at top level only (not in conditions/loops)
- [ ] Custom hooks prefixed with `use`
- [ ] Effect dependencies are complete (no eslint-disable exhaustive-deps)
- [ ] Effects clean up subscriptions/timers/listeners
- [ ] No unnecessary custom hooks (simple logic can be inline)

**Related Rules**: RCT-HKS-001, RCT-HKS-002, RCT-HKS-003, RCT-HKS-004, RCT-HKS-005, RCT-STY-003

---

## State Management

- [ ] State is colocated with where it's used
- [ ] No redundant state (derive when possible)
- [ ] `useReducer` used for complex state logic
- [ ] Context values are memoized
- [ ] No state that can be calculated from props

**Related Rules**: RCT-STT-001, RCT-STT-002, RCT-STT-003, RCT-STT-004

---

## Props

- [ ] Props are typed with TypeScript interfaces
- [ ] Props are destructured
- [ ] Sensible defaults provided where appropriate
- [ ] `children` used for composition
- [ ] No blind spreading of unknown props

**Related Rules**: RCT-PRP-001, RCT-PRP-002, RCT-PRP-003, RCT-PRP-004

---

## Events

- [ ] Event handlers named with `handle` prefix
- [ ] Event props named with `on` prefix
- [ ] Callbacks memoized when passed to memoized children
- [ ] Event types properly typed

**Related Rules**: RCT-EVT-001, RCT-EVT-002

---

## Rendering

- [ ] Stable keys used for list items (not array indices)
- [ ] No short-circuit with potentially falsy numbers
- [ ] Conditional rendering uses appropriate pattern
- [ ] No nested ternaries (use object map or early returns)
- [ ] Fragments used to avoid unnecessary wrapper divs

**Related Rules**: RCT-RND-001, RCT-RND-002, RCT-RND-003

---

## Performance

- [ ] `useMemo` used only for expensive computations
- [ ] `React.memo` used appropriately with stable props
- [ ] `useCallback` used for callbacks passed to memoized children
- [ ] No object literals created inline in JSX (for memoized components)
- [ ] Context providers memoize their values

**Related Rules**: RCT-PRF-001, RCT-PRF-002, RCT-PRF-003, RCT-PRF-004

---

## Testing

- [ ] Tests query by role/accessible name
- [ ] Tests use userEvent, not fireEvent
- [ ] Tests check behavior, not implementation
- [ ] Async operations use waitFor/findBy
- [ ] No test IDs when accessible queries work

**Related Rules**: RCT-TST-001, RCT-TST-002, RCT-TST-003, RCT-TST-004

---

## Accessibility

- [ ] Semantic HTML used (nav, main, article, button, etc.)
- [ ] Images have meaningful alt text
- [ ] Custom components are keyboard navigable
- [ ] ARIA attributes used correctly
- [ ] Focus management handled for modals/dynamic content
- [ ] Color contrast meets WCAG AA (4.5:1 for text)

**Related Rules**: RCT-A11-001, RCT-A11-002, RCT-A11-003, RCT-A11-004, RCT-A11-005

---

## Naming Conventions

- [ ] Components use PascalCase
- [ ] Props use camelCase
- [ ] Custom hooks use `use` prefix
- [ ] Boolean props use `is`, `has`, `should` prefix
- [ ] Event handlers use `handle` prefix
- [ ] Event props use `on` prefix

**Related Rules**: RCT-STY-001, RCT-STY-002, RCT-STY-003

---

## Security

- [ ] No `dangerouslySetInnerHTML` with user input
- [ ] Links with `target="_blank"` have `rel="noopener noreferrer"`
- [ ] No sensitive data in client-side state
- [ ] User input is validated before use
- [ ] API responses are typed and validated

---

## TypeScript

- [ ] No `any` types (use `unknown` if type is truly unknown)
- [ ] Props interfaces defined for all components
- [ ] Event handlers properly typed
- [ ] Generic types used appropriately
- [ ] Strict null checks handled

---

## Code Organization

- [ ] Imports grouped: React, third-party, internal
- [ ] Related code is colocated (components with tests/stories)
- [ ] Shared components in `components/` directory
- [ ] Feature-specific code in `features/` directory
- [ ] Custom hooks in appropriate location

---

## Quick Checks

| Check | Pass/Fail |
|-------|-----------|
| No class components (except error boundaries) | |
| No hooks in conditionals or loops | |
| All effects have complete dependencies | |
| All lists have stable keys | |
| All images have alt text | |
| All interactive elements are keyboard accessible | |
| TypeScript types for all props | |
| No inline objects for memoized components | |
