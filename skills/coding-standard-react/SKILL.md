---
name: react-standard
description: Complete React coding standard reference (project)
---

# React Coding Standard

> **Version**: 1.0.0 | **Status**: Active
> **Base Standards**: Airbnb React/JSX Style Guide, React Documentation Best Practices

This standard establishes coding conventions for React development, ensuring consistency, quality, maintainability, performance, and accessibility across all React projects.

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [Components](rules/components.md) | RCT-CMP-* | Functional components, composition |
| [Hooks](rules/hooks.md) | RCT-HKS-* | Hooks rules, custom hooks |
| [State](rules/state.md) | RCT-STT-* | State management patterns |
| [Props](rules/props.md) | RCT-PRP-* | Props typing and patterns |
| [Events](rules/events.md) | RCT-EVT-* | Event handling conventions |
| [Rendering](rules/rendering.md) | RCT-RND-* | Conditional rendering, keys |
| [Performance](rules/performance.md) | RCT-PRF-* | Memoization, optimization |
| [Testing](rules/testing.md) | RCT-TST-* | Testing Library best practices |
| [Accessibility](rules/accessibility.md) | RCT-A11-* | WCAG compliance, ARIA |
| [Styling](rules/styling.md) | RCT-STY-* | Naming conventions |

### Tooling

| Tool | File |
|------|------|
| [ESLint](tooling/eslint.md) | ESLint configuration for React |
| [Prettier](tooling/prettier.md) | Code formatting |
| [Testing Library](tooling/testing-library.md) | React Testing Library setup |
| [Pre-commit](tooling/pre-commit.md) | Git hooks setup |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table with tiers |
| [Code Review Checklist](reference/code-review.md) | Review checklist by category |

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Critical Rules (Always Apply)

These rules are non-negotiable and must be followed in all code.

### RCT-STY-001: Use PascalCase for Component Names :red_circle:

React treats components starting with lowercase letters as DOM tags.

```tsx
// Correct
function UserProfile() {
  return <div>User Profile</div>;
}

// Incorrect
function userProfile() {
  return <div>User Profile</div>;
}
```

### RCT-STY-003: Prefix Custom Hooks with `use` :red_circle:

Required by React to enable the Rules of Hooks linting.

```tsx
// Correct
function useAuth() {
  const [user, setUser] = useState(null);
  return { user, setUser };
}

// Incorrect
function getAuth() {
  const [user, setUser] = useState(null); // Linter cannot verify hooks rules
  return { user, setUser };
}
```

### RCT-CMP-001: Use Functional Components :red_circle:

Functional components with hooks are the modern React standard.

```tsx
// Correct
function UserProfile({ name, email }: UserProfileProps) {
  const [isEditing, setIsEditing] = useState(false);
  return <div><h1>{name}</h1></div>;
}

// Incorrect
class UserProfile extends React.Component<UserProfileProps> {
  render() {
    return <div><h1>{this.props.name}</h1></div>;
  }
}
```

### RCT-HKS-001: Call Hooks at Top Level Only :red_circle:

React relies on the order of hooks calls to preserve state between renders.

```tsx
// Correct
function UserProfile({ userId }: { userId: string }) {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    if (userId) {
      fetchUser(userId).then(setUser);
    }
  }, [userId]);

  return <div>{user?.name}</div>;
}

// Incorrect
function UserProfile({ userId }: { userId: string }) {
  if (userId) {
    const [user, setUser] = useState<User | null>(null); // WRONG!
  }
  return <div>{user?.name}</div>;
}
```

### RCT-HKS-002: Include All Dependencies in useEffect :red_circle:

Missing dependencies cause stale closures and hard-to-debug bugs.

```tsx
// Correct
useEffect(() => {
  searchAPI(query, filters).then(setResults);
}, [query, filters]); // All dependencies listed

// Incorrect
useEffect(() => {
  searchAPI(query, filters).then(setResults);
}, [query]); // Missing 'filters' dependency
```

### RCT-HKS-003: Clean Up Effects Properly :red_circle:

Effects that set up subscriptions must clean up to prevent memory leaks.

```tsx
// Correct
useEffect(() => {
  window.addEventListener('resize', handleResize);
  return () => {
    window.removeEventListener('resize', handleResize);
  };
}, []);

// Incorrect - Missing cleanup
useEffect(() => {
  window.addEventListener('resize', handleResize);
}, []);
```

### RCT-PRP-002: Use TypeScript for Prop Types :red_circle:

TypeScript provides compile-time checking and self-documenting code.

```tsx
// Correct
interface UserCardProps {
  user: { id: string; name: string; };
  isSelected?: boolean;
  onSelect?: (id: string) => void;
}

function UserCard({ user, isSelected = false, onSelect }: UserCardProps) {
  return <div onClick={() => onSelect?.(user.id)}>{user.name}</div>;
}

// Incorrect - No types
function UserCard({ user, isSelected, onSelect }) {
  return <div onClick={() => onSelect(user.id)}>{user.name}</div>;
}
```

### RCT-RND-002: Avoid Short-Circuit with Numbers :red_circle:

Short-circuit evaluation with falsy numbers (0) will render "0" instead of nothing.

```tsx
// Correct
function ItemCount({ count }: { count: number }) {
  return count > 0 && <span>{count} items</span>;
}

// Incorrect - Will render "0"
function ItemCount({ count }: { count: number }) {
  return count && <span>{count} items</span>;
}
```

### RCT-RND-003: Use Stable Keys for Lists :red_circle:

Keys help React identify which items changed. Array indices cause bugs.

```tsx
// Correct
{todos.map(todo => (
  <TodoItem key={todo.id} todo={todo} />
))}

// Incorrect
{todos.map((todo, index) => (
  <TodoItem key={index} todo={todo} />
))}
```

### RCT-TST-001: Query by Role and Accessible Name :red_circle:

Queries by role reflect how assistive technologies see your UI.

```tsx
// Correct
const emailInput = screen.getByRole('textbox', { name: /email/i });
const submitButton = screen.getByRole('button', { name: /sign in/i });

// Incorrect
const emailInput = screen.getByTestId('email-input');
```

### RCT-TST-003: Test Behavior, Not Implementation :red_circle:

Tests that check internal state break when refactoring.

```tsx
// Correct - Test observable behavior
await user.click(screen.getByRole('button', { name: /increment/i }));
expect(screen.getByText('Count: 1')).toBeInTheDocument();

// Incorrect - Test implementation
expect(result.current.state.count).toBe(1);
```

### RCT-A11-001: Use Semantic HTML Elements :red_circle:

Semantic HTML provides built-in accessibility features.

```tsx
// Correct
<nav aria-label="Main navigation">
  <ul><li><a href="/home">Home</a></li></ul>
</nav>

// Incorrect
<div className="nav">
  <div onClick={() => navigate('/home')}>Home</div>
</div>
```

### RCT-A11-002: Provide Text Alternatives for Images :red_circle:

Screen readers need text descriptions for images.

```tsx
// Correct
<img src={product.imageUrl} alt={`${product.name} - ${product.color}`} />

// Incorrect
<img src={product.imageUrl} />
```

### RCT-A11-003: Ensure Keyboard Navigation :red_circle:

All interactive elements must be focusable and operable via keyboard.

```tsx
// Correct - Keyboard accessible
<div
  role="button"
  tabIndex={0}
  onClick={handleClick}
  onKeyDown={(e) => e.key === 'Enter' && handleClick()}
>
  Click me
</div>

// Incorrect - Not keyboard accessible
<div onClick={handleClick}>Click me</div>
```

---

## Quick Rule Lookup

| ID | Rule | Tier |
|----|------|------|
| RCT-STY-001 | Use PascalCase for Component Names | Critical |
| RCT-STY-002 | Use camelCase for Prop Names | Required |
| RCT-STY-003 | Prefix Custom Hooks with `use` | Critical |
| RCT-CMP-001 | Use Functional Components | Critical |
| RCT-CMP-002 | One Component Per File | Required |
| RCT-CMP-003 | Prefer Composition Over Prop Drilling | Required |
| RCT-CMP-004 | Keep Components Focused and Small | Required |
| RCT-HKS-001 | Call Hooks at Top Level Only | Critical |
| RCT-HKS-002 | Include All Dependencies in useEffect | Critical |
| RCT-HKS-003 | Clean Up Effects Properly | Critical |
| RCT-HKS-004 | Extract Reusable Logic into Custom Hooks | Required |
| RCT-HKS-005 | Keep Custom Hooks Focused | Recommended |
| RCT-STT-001 | Keep State Close to Where It's Used | Required |
| RCT-STT-002 | Use useReducer for Complex State Logic | Required |
| RCT-STT-003 | Derive State When Possible | Required |
| RCT-STT-004 | Avoid Redundant State | Required |
| RCT-PRP-001 | Destructure Props in Function Signature | Recommended |
| RCT-PRP-002 | Use TypeScript for Prop Types | Critical |
| RCT-PRP-003 | Avoid Spreading Props Blindly | Required |
| RCT-PRP-004 | Use Children for Composition | Recommended |
| RCT-EVT-001 | Use Consistent Event Handler Naming | Required |
| RCT-EVT-002 | Avoid Inline Arrow Functions (Performance Critical) | Recommended |
| RCT-RND-001 | Use Appropriate Conditional Rendering Patterns | Required |
| RCT-RND-002 | Avoid Short-Circuit with Numbers | Critical |
| RCT-RND-003 | Use Stable Keys for Lists | Critical |
| RCT-PRF-001 | Memoize Expensive Computations | Required |
| RCT-PRF-002 | Use React.memo for Pure Components | Required |
| RCT-PRF-003 | Avoid Creating Objects in Render | Required |
| RCT-PRF-004 | Use useCallback for Function Props | Required |
| RCT-TST-001 | Query by Role and Accessible Name | Critical |
| RCT-TST-002 | Use userEvent Over fireEvent | Required |
| RCT-TST-003 | Test Behavior, Not Implementation | Critical |
| RCT-TST-004 | Use waitFor for Async Assertions | Required |
| RCT-A11-001 | Use Semantic HTML Elements | Critical |
| RCT-A11-002 | Provide Text Alternatives for Images | Critical |
| RCT-A11-003 | Ensure Keyboard Navigation | Critical |
| RCT-A11-004 | Use ARIA Attributes Correctly | Required |
| RCT-A11-005 | Maintain Sufficient Color Contrast | Required |

---

## References

- [Airbnb React/JSX Style Guide](https://airbnb.io/javascript/react/)
- [React Documentation](https://react.dev)
- [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react)
- [eslint-plugin-react-hooks](https://react.dev/reference/eslint-plugin-react-hooks)
- [eslint-plugin-jsx-a11y](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- [WCAG 2.1 Quick Reference](https://www.w3.org/WAI/WCAG21/quickref/)
