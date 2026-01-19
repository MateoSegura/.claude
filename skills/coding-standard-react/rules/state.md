# State Management (RCT-STT-*)

State management is fundamental to React applications. These rules ensure state is organized effectively, avoiding common pitfalls like redundant state and unnecessary re-renders.

## State Management Strategy

Choose the right tool for the scope of state:

| Scope | Solution |
|-------|----------|
| Local component state | `useState` |
| Complex local state with actions | `useReducer` |
| Shared state (small scope) | Lift state up + props |
| Shared state (medium scope) | React Context |
| Global/complex state | External library (Zustand, Jotai, Redux) |
| Server state | React Query, SWR, Apollo |

---

## RCT-STT-001: Keep State Close to Where It's Used :yellow_circle:

**Tier**: Required

**Rationale**: Colocating state with its consumers improves performance (fewer re-renders) and makes code easier to understand.

```tsx
// Correct - State close to usage
function SearchPage() {
  return (
    <div>
      <SearchForm /> {/* Manages its own search state */}
      <FilterPanel /> {/* Manages its own filter state */}
      <ResultsList /> {/* Receives data via props or context */}
    </div>
  );
}

function SearchForm() {
  const [query, setQuery] = useState('');
  // query state only needed here
  return (
    <form>
      <input value={query} onChange={e => setQuery(e.target.value)} />
    </form>
  );
}

// Incorrect - State too high in tree
function SearchPage() {
  const [query, setQuery] = useState('');
  const [filters, setFilters] = useState({});
  const [sortOrder, setSortOrder] = useState('asc');
  // All state at top, causing unnecessary re-renders

  return (
    <div>
      <SearchForm query={query} setQuery={setQuery} />
      <FilterPanel filters={filters} setFilters={setFilters} />
      <ResultsList sortOrder={sortOrder} />
    </div>
  );
}
```

---

## RCT-STT-002: Use useReducer for Complex State Logic :yellow_circle:

**Tier**: Required

**Rationale**: When state updates involve multiple sub-values or complex transitions, `useReducer` provides clearer, more maintainable code.

```tsx
// Correct - useReducer for complex state
type FormState = {
  values: Record<string, string>;
  errors: Record<string, string>;
  touched: Record<string, boolean>;
  isSubmitting: boolean;
};

type FormAction =
  | { type: 'SET_FIELD'; field: string; value: string }
  | { type: 'SET_ERROR'; field: string; error: string }
  | { type: 'TOUCH_FIELD'; field: string }
  | { type: 'SUBMIT_START' }
  | { type: 'SUBMIT_SUCCESS' }
  | { type: 'SUBMIT_ERROR'; errors: Record<string, string> };

function formReducer(state: FormState, action: FormAction): FormState {
  switch (action.type) {
    case 'SET_FIELD':
      return {
        ...state,
        values: { ...state.values, [action.field]: action.value },
        errors: { ...state.errors, [action.field]: '' },
      };
    case 'SUBMIT_START':
      return { ...state, isSubmitting: true };
    case 'SUBMIT_SUCCESS':
      return { ...state, isSubmitting: false };
    case 'SUBMIT_ERROR':
      return { ...state, isSubmitting: false, errors: action.errors };
    default:
      return state;
  }
}

function ContactForm() {
  const [state, dispatch] = useReducer(formReducer, initialState);
  // Clean, predictable state updates
}

// Incorrect - Multiple useState for related state
function ContactForm() {
  const [values, setValues] = useState({});
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);

  // Complex, error-prone updates scattered throughout
  const handleChange = (field: string, value: string) => {
    setValues(v => ({ ...v, [field]: value }));
    setErrors(e => ({ ...e, [field]: '' }));
    setTouched(t => ({ ...t, [field]: true }));
  };
}
```

**When to use useReducer**:
- State logic involves multiple sub-values
- Next state depends on the previous state
- State transitions are complex or need to be testable
- Multiple event handlers update the same state

---

## RCT-STT-003: Derive State When Possible :yellow_circle:

**Tier**: Required

**Rationale**: Derived state should be calculated during render, not stored. Storing derived values leads to synchronization bugs.

```tsx
// Correct - Derived during render
function ShoppingCart({ items }: { items: CartItem[] }) {
  // Derived values calculated during render
  const itemCount = items.length;
  const subtotal = items.reduce((sum, item) => sum + item.price * item.quantity, 0);
  const tax = subtotal * 0.1;
  const total = subtotal + tax;

  return (
    <div>
      <span>{itemCount} items</span>
      <span>Total: ${total.toFixed(2)}</span>
    </div>
  );
}

// Correct - useMemo for expensive derivations
function DataTable({ data, sortConfig }: DataTableProps) {
  const sortedData = useMemo(() => {
    return [...data].sort((a, b) => {
      // Expensive sort operation
      return sortConfig.direction === 'asc'
        ? a[sortConfig.key].localeCompare(b[sortConfig.key])
        : b[sortConfig.key].localeCompare(a[sortConfig.key]);
    });
  }, [data, sortConfig]);

  return <table>{/* ... */}</table>;
}

// Incorrect - Storing derived state
function ShoppingCart({ items }: { items: CartItem[] }) {
  const [itemCount, setItemCount] = useState(0);
  const [total, setTotal] = useState(0);

  // Bug-prone: must remember to update derived state
  useEffect(() => {
    setItemCount(items.length);
    setTotal(items.reduce((sum, item) => sum + item.price * item.quantity, 0) * 1.1);
  }, [items]);

  return (
    <div>
      <span>{itemCount} items</span>
      <span>Total: ${total.toFixed(2)}</span>
    </div>
  );
}
```

---

## RCT-STT-004: Avoid Redundant State :yellow_circle:

**Tier**: Required

**Rationale**: State that can be calculated from other state or props is redundant and creates synchronization issues.

```tsx
// Correct - Single source of truth
function UserSearch({ users }: { users: User[] }) {
  const [query, setQuery] = useState('');

  // Filtered list derived from query and users
  const filteredUsers = users.filter(user =>
    user.name.toLowerCase().includes(query.toLowerCase())
  );

  return (
    <div>
      <input value={query} onChange={e => setQuery(e.target.value)} />
      <UserList users={filteredUsers} />
    </div>
  );
}

// Incorrect - Redundant state
function UserSearch({ users }: { users: User[] }) {
  const [query, setQuery] = useState('');
  const [filteredUsers, setFilteredUsers] = useState(users);

  // Must manually sync - error prone
  useEffect(() => {
    setFilteredUsers(users.filter(user =>
      user.name.toLowerCase().includes(query.toLowerCase())
    ));
  }, [query, users]);

  return (
    <div>
      <input value={query} onChange={e => setQuery(e.target.value)} />
      <UserList users={filteredUsers} />
    </div>
  );
}
```

**Signs of redundant state**:
- An effect that only updates state based on other state/props
- State that is reset when props change
- State that can be computed from props
- Multiple pieces of state that must stay in sync

**Questions to ask before adding state**:
1. Can this be derived from existing state or props?
2. Does this duplicate information already available?
3. Can I compute this value during render instead?
