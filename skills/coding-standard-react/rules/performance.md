# Performance (RCT-PRF-*)

React provides tools for optimizing rendering performance. These rules ensure performance optimizations are applied correctly and only when needed.

## Performance Principles

- Measure before optimizing
- Avoid premature optimization
- Understand React's rendering behavior
- Use the right tool for the job

---

## RCT-PRF-001: Memoize Expensive Computations :yellow_circle:

**Tier**: Required

**Rationale**: `useMemo` prevents recalculation of expensive values on every render. Use it when computations are genuinely expensive.

```tsx
// Correct - useMemo for expensive computation
function DataTable({ data, sortConfig }: DataTableProps) {
  const sortedData = useMemo(() => {
    // Expensive sort operation
    return [...data].sort((a, b) => {
      const aValue = a[sortConfig.key];
      const bValue = b[sortConfig.key];
      return sortConfig.direction === 'asc'
        ? aValue.localeCompare(bValue)
        : bValue.localeCompare(aValue);
    });
  }, [data, sortConfig]);

  return (
    <table>
      {sortedData.map(row => (
        <TableRow key={row.id} data={row} />
      ))}
    </table>
  );
}

// Correct - useMemo for filtered data
function UserList({ users, filter }: UserListProps) {
  const filteredUsers = useMemo(() => {
    return users.filter(user => {
      return (
        user.name.toLowerCase().includes(filter.toLowerCase()) &&
        user.status === 'active'
      );
    });
  }, [users, filter]);

  return <List items={filteredUsers} />;
}

// Incorrect - useMemo for trivial computation
function Greeting({ firstName, lastName }: GreetingProps) {
  // Unnecessary - string concatenation is cheap
  const fullName = useMemo(
    () => `${firstName} ${lastName}`,
    [firstName, lastName]
  );

  return <h1>Hello, {fullName}</h1>;
}

// Correct - No memoization needed for cheap operations
function Greeting({ firstName, lastName }: GreetingProps) {
  const fullName = `${firstName} ${lastName}`;
  return <h1>Hello, {fullName}</h1>;
}
```

**When to use useMemo**:
- Sorting or filtering large arrays
- Complex calculations
- Creating objects passed to memoized children
- Expensive transformations

**When NOT to use useMemo**:
- Simple computations (addition, string concatenation)
- Creating primitive values
- When the computation is faster than the memoization overhead

---

## RCT-PRF-002: Use React.memo for Pure Components :yellow_circle:

**Tier**: Required

**Rationale**: `React.memo` prevents re-renders when props haven't changed. Use it for components that render often with the same props.

```tsx
// Correct - Memoized list item
interface UserCardProps {
  user: User;
  onSelect: (id: string) => void;
}

const UserCard = memo(function UserCard({ user, onSelect }: UserCardProps) {
  return (
    <div onClick={() => onSelect(user.id)}>
      <img src={user.avatarUrl} alt={user.name} />
      <h3>{user.name}</h3>
    </div>
  );
});

// Parent with stable callback
function UserList({ users }: { users: User[] }) {
  const handleSelect = useCallback((id: string) => {
    console.log('Selected:', id);
  }, []);

  return (
    <div>
      {users.map(user => (
        <UserCard key={user.id} user={user} onSelect={handleSelect} />
      ))}
    </div>
  );
}

// Correct - Custom comparison function
const ExpensiveChart = memo(
  function ExpensiveChart({ data, options }: ChartProps) {
    return <canvas>{/* Complex rendering */}</canvas>;
  },
  (prevProps, nextProps) => {
    // Only re-render if data length or options changed
    return (
      prevProps.data.length === nextProps.data.length &&
      prevProps.options.type === nextProps.options.type
    );
  }
);

// Incorrect - memo without stable props
const UserCard = memo(function UserCard({ user, onSelect }: UserCardProps) {
  return <div onClick={() => onSelect(user.id)}>{user.name}</div>;
});

function UserList({ users }: { users: User[] }) {
  return (
    <div>
      {users.map(user => (
        <UserCard
          key={user.id}
          user={user}
          onSelect={(id) => console.log(id)} // New function each render - memo useless
        />
      ))}
    </div>
  );
}
```

---

## RCT-PRF-003: Avoid Creating Objects in Render :yellow_circle:

**Tier**: Required

**Rationale**: Object literals in JSX create new references each render, defeating memoization and causing unnecessary re-renders.

```tsx
// Correct - Stable object references
const defaultStyle = { color: 'blue', fontSize: 16 };

function StyledText({ children }: { children: React.ReactNode }) {
  return <span style={defaultStyle}>{children}</span>;
}

// Correct - useMemo for dynamic objects
function StyledText({ color, size }: StyledTextProps) {
  const style = useMemo(
    () => ({ color, fontSize: size }),
    [color, size]
  );

  return <span style={style}>Text</span>;
}

// Correct - Memoized context value
function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [theme, setTheme] = useState('light');

  const value = useMemo(() => ({ theme, setTheme }), [theme]);

  return (
    <ThemeContext.Provider value={value}>
      {children}
    </ThemeContext.Provider>
  );
}

// Incorrect - New object every render
function StyledText({ children }: { children: React.ReactNode }) {
  return <span style={{ color: 'blue', fontSize: 16 }}>{children}</span>;
  // New style object created every render
}

// Incorrect with context - causes all consumers to re-render
function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [theme, setTheme] = useState('light');

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {/* New object every render - all consumers re-render */}
      {children}
    </ThemeContext.Provider>
  );
}
```

---

## RCT-PRF-004: Use useCallback for Function Props :yellow_circle:

**Tier**: Required

**Rationale**: `useCallback` maintains stable function references, enabling effective memoization of child components.

```tsx
// Correct - Stable callback reference
function SearchForm({ onSearch }: { onSearch: (query: string) => void }) {
  const [query, setQuery] = useState('');

  const handleSubmit = useCallback((e: React.FormEvent) => {
    e.preventDefault();
    onSearch(query);
  }, [onSearch, query]);

  return (
    <form onSubmit={handleSubmit}>
      <input value={query} onChange={e => setQuery(e.target.value)} />
      <MemoizedButton type="submit">Search</MemoizedButton>
    </form>
  );
}

// Correct - useCallback with functional update to avoid dependencies
function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const handleToggle = useCallback((id: string) => {
    setTodos(prev => prev.map(todo =>
      todo.id === id ? { ...todo, completed: !todo.completed } : todo
    ));
  }, []); // No dependencies - uses functional update

  const handleDelete = useCallback((id: string) => {
    setTodos(prev => prev.filter(todo => todo.id !== id));
  }, []);

  return (
    <ul>
      {todos.map(todo => (
        <MemoizedTodoItem
          key={todo.id}
          todo={todo}
          onToggle={handleToggle}
          onDelete={handleDelete}
        />
      ))}
    </ul>
  );
}
```

**When to use useCallback**:
- Passing callbacks to memoized children
- Callbacks used as dependencies in useEffect
- Event handlers in performance-critical sections

**When NOT to use useCallback**:
- Simple event handlers with no memoized children
- Functions that don't need stable references
- When the component rarely re-renders anyway
