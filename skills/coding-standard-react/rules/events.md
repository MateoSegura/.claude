# Event Handling (RCT-EVT-*)

Event handling is a core part of React development. These rules ensure event handlers are named consistently and perform well.

## Event Handler Conventions

- Name handler functions with `handle` prefix (e.g., `handleClick`)
- Name handler props with `on` prefix (e.g., `onClick`)
- Always type event parameters
- Use `useCallback` when passing handlers to memoized children

---

## RCT-EVT-001: Use Consistent Event Handler Naming :yellow_circle:

**Tier**: Required

**Rationale**: Consistent naming makes code predictable and easier to understand. The `handle`/`on` convention is a widely adopted React pattern.

```tsx
// Correct
function LoginForm({ onSubmit, onCancel }: LoginFormProps) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    onSubmit({ email, password });
  };

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="email" onChange={handleEmailChange} />
      <input type="password" onChange={handlePasswordChange} />
      <button type="submit">Sign In</button>
      <button type="button" onClick={onCancel}>Cancel</button>
    </form>
  );
}

// Incorrect - Inconsistent naming
function LoginForm({ submitCallback, cancelAction }: LoginFormProps) {
  const submit = (e) => { /* ... */ };
  const emailChanged = (e) => { /* ... */ };

  return (
    <form onSubmit={submit}>
      <input onChange={emailChanged} />
      <button onClick={cancelAction}>Cancel</button>
    </form>
  );
}
```

**Naming pattern**:
- Internal handlers: `handle` + `Element` + `Event` (e.g., `handleEmailChange`, `handleFormSubmit`)
- Props: `on` + `Action` (e.g., `onSubmit`, `onCancel`, `onSelect`)

---

## RCT-EVT-002: Avoid Inline Arrow Functions in JSX (When Performance Critical) :green_circle:

**Tier**: Recommended

**Rationale**: Inline functions create new references on each render, which can cause unnecessary re-renders in memoized children. Use `useCallback` for handlers passed to memoized components.

```tsx
// Correct - useCallback for memoized children
function ParentList({ items }: { items: Item[] }) {
  const handleItemClick = useCallback((id: string) => {
    console.log('Clicked:', id);
  }, []);

  return (
    <ul>
      {items.map(item => (
        <MemoizedItem
          key={item.id}
          item={item}
          onClick={handleItemClick}
        />
      ))}
    </ul>
  );
}

// Correct - useCallback with item-specific handler
function ParentList({ items, onSelect }: ListProps) {
  const handleItemClick = useCallback((id: string) => {
    onSelect(id);
  }, [onSelect]);

  return (
    <ul>
      {items.map(item => (
        <MemoizedItem
          key={item.id}
          item={item}
          onClick={handleItemClick}
        />
      ))}
    </ul>
  );
}

// Acceptable - Inline for non-memoized children
function SimpleList({ items }: { items: Item[] }) {
  return (
    <ul>
      {items.map(item => (
        <li key={item.id} onClick={() => console.log(item.id)}>
          {item.name}
        </li>
      ))}
    </ul>
  );
}

// Incorrect - Creating new function for memoized component
function ParentList({ items }: { items: Item[] }) {
  return (
    <ul>
      {items.map(item => (
        <MemoizedItem
          key={item.id}
          item={item}
          onClick={() => console.log(item.id)} // New function each render
        />
      ))}
    </ul>
  );
}
```

**When to worry about inline functions**:
- The handler is passed to a memoized component (`React.memo`)
- The list is long (100+ items)
- Profiling shows the handler creation is a bottleneck

**When inline functions are fine**:
- The child is not memoized
- The list is short
- It's a one-off component, not rendered frequently

---

## Event Typing Reference

Common event types for TypeScript:

```tsx
// Click events
const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
  event.preventDefault();
};

// Form events
const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
  event.preventDefault();
};

// Input change events
const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  setValue(event.target.value);
};

// Select change events
const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
  setOption(event.target.value);
};

// Keyboard events
const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
  if (event.key === 'Enter') {
    handleSubmit();
  }
};

// Focus events
const handleFocus = (event: React.FocusEvent<HTMLInputElement>) => {
  setFocused(true);
};

// Drag events
const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
  event.preventDefault();
  const files = event.dataTransfer.files;
};
```
