# Rendering Patterns (RCT-RND-*)

Rendering is at the core of React. These rules ensure rendering logic is clear, correct, and performs well.

## Conditional Rendering Patterns

Use the appropriate pattern for different scenarios:

| Pattern | Use Case |
|---------|----------|
| `&&` short-circuit | Simple presence check (boolean) |
| Ternary `? :` | Two-way conditions |
| Early return | Guard clauses |
| Object map | Multiple conditions |

---

## RCT-RND-001: Use Appropriate Conditional Rendering Patterns :yellow_circle:

**Tier**: Required

**Rationale**: Different conditions call for different patterns. Using the right pattern improves readability.

```tsx
// Correct - Short-circuit for simple conditions
function Notification({ message, isVisible }: NotificationProps) {
  return isVisible && <div className="notification">{message}</div>;
}

// Correct - Ternary for two-way conditions
function StatusBadge({ isOnline }: { isOnline: boolean }) {
  return (
    <span className={isOnline ? 'badge-online' : 'badge-offline'}>
      {isOnline ? 'Online' : 'Offline'}
    </span>
  );
}

// Correct - Early return for guard clauses
function UserProfile({ user }: { user: User | null }) {
  if (!user) {
    return <div>Please log in</div>;
  }

  return (
    <div>
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}

// Correct - Object map for multiple conditions
function ComplexStatus({ status }: { status: string }) {
  const statusMap: Record<string, string> = {
    loading: 'Loading...',
    error: 'Error!',
    success: 'Done!',
  };

  return <span>{statusMap[status] ?? 'Unknown'}</span>;
}

// Correct - Component map for complex rendering
function StatusDisplay({ status }: { status: string }) {
  const StatusComponent = {
    loading: LoadingSpinner,
    error: ErrorMessage,
    success: SuccessMessage,
  }[status] ?? UnknownStatus;

  return <StatusComponent />;
}

// Incorrect - Nested ternaries
function ComplexStatus({ status }: { status: string }) {
  return (
    <span>
      {status === 'loading'
        ? 'Loading...'
        : status === 'error'
          ? 'Error!'
          : status === 'success'
            ? 'Done!'
            : 'Unknown'}
    </span>
  );
}
```

---

## RCT-RND-002: Avoid Short-Circuit with Numbers :red_circle:

**Tier**: Critical

**Rationale**: Short-circuit evaluation with falsy numbers (0) will render "0" instead of nothing.

```tsx
// Correct - Explicit comparison
function ItemCount({ count }: { count: number }) {
  return count > 0 && <span>{count} items</span>;
}

// Also correct - explicit boolean conversion
function ItemCount({ count }: { count: number }) {
  return !!count && <span>{count} items</span>;
}

// Also correct - ternary with null
function ItemCount({ count }: { count: number }) {
  return count ? <span>{count} items</span> : null;
}

// Incorrect - Will render "0"
function ItemCount({ count }: { count: number }) {
  return count && <span>{count} items</span>;
  // When count is 0, renders "0" instead of nothing
}

// Other falsy values to watch for:
function Examples() {
  const emptyString = '';
  const zero = 0;
  const nan = NaN;

  return (
    <>
      {emptyString && <div>Never shown</div>} {/* Renders nothing, OK */}
      {zero && <div>Never shown</div>}        {/* Renders "0", BAD */}
      {nan && <div>Never shown</div>}         {/* Renders nothing, OK */}
    </>
  );
}
```

---

## RCT-RND-003: Use Stable Keys for Lists :red_circle:

**Tier**: Critical

**Rationale**: Keys help React identify which items changed. Using array indices as keys causes bugs when list order changes.

```tsx
// Correct - Stable unique ID
function TodoList({ todos }: { todos: Todo[] }) {
  return (
    <ul>
      {todos.map(todo => (
        <TodoItem key={todo.id} todo={todo} />
      ))}
    </ul>
  );
}

// Correct - Composite key when no ID available
function SearchResults({ results }: { results: Result[] }) {
  return (
    <ul>
      {results.map(result => (
        <ResultItem key={`${result.category}-${result.slug}`} result={result} />
      ))}
    </ul>
  );
}

// Correct - Using index ONLY for static lists that never reorder
function StaticNav() {
  const links = ['Home', 'About', 'Contact']; // Never changes
  return (
    <nav>
      {links.map((link, index) => (
        <a key={index} href={`/${link.toLowerCase()}`}>{link}</a>
      ))}
    </nav>
  );
}

// Incorrect - Array index as key for dynamic lists
function TodoList({ todos }: { todos: Todo[] }) {
  return (
    <ul>
      {todos.map((todo, index) => (
        <TodoItem key={index} todo={todo} />
        // Bug: Reordering or removing items causes state issues
      ))}
    </ul>
  );
}
```

**Problems with index keys**:
- Reordering items causes incorrect state preservation
- Adding/removing items from the middle breaks component state
- Animations and transitions behave incorrectly
- Form inputs may show wrong values

**When index keys are acceptable**:
- List is static and never changes
- Items will never be reordered
- Items will never be filtered or deleted
- Items have no state or uncontrolled inputs

---

## Fragment Usage

```tsx
// Correct - Fragment to avoid extra DOM node
function UserInfo({ user }: { user: User }) {
  return (
    <>
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </>
  );
}

// Correct - Fragment with key in lists
function Glossary({ items }: { items: GlossaryItem[] }) {
  return (
    <dl>
      {items.map(item => (
        <React.Fragment key={item.id}>
          <dt>{item.term}</dt>
          <dd>{item.definition}</dd>
        </React.Fragment>
      ))}
    </dl>
  );
}

// Incorrect - Unnecessary wrapping div
function UserInfo({ user }: { user: User }) {
  return (
    <div> {/* Unnecessary wrapper */}
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}
```
