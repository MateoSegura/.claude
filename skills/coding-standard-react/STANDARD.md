# React Coding Standard

> **Version**: 1.0.0
> **Status**: Active
> **Base Standard**: Airbnb React/JSX Style Guide, React Documentation Best Practices
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes coding conventions for React development. It ensures:

- **Consistency**: Uniform code style across all React projects
- **Quality**: Reduced defects through proven patterns and hooks best practices
- **Maintainability**: Code that is easy to read, modify, and extend
- **Performance**: Optimized rendering through proper memoization and state management
- **Accessibility**: Inclusive applications that work for all users

### 1.2 Scope

This standard applies to:
- [x] All React source files (`.jsx`, `.tsx`) in production repositories
- [x] React component tests using React Testing Library
- [x] Custom hooks and utility functions
- [x] Storybook stories and documentation

### 1.3 Audience

- Software engineers writing React applications
- Code reviewers evaluating React component changes
- Tech leads establishing project standards
- QA engineers writing component tests

### 1.4 Relationship to Industry Standards

| Standard | Relationship |
|----------|--------------|
| [Airbnb React/JSX Style Guide](https://airbnb.io/javascript/react/) | **Base** - All rules apply unless documented otherwise |
| [React Documentation](https://react.dev) | **Authoritative** - Official patterns and hooks rules |
| [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) | **Enforcement** - Automated rule checking |
| [eslint-plugin-react-hooks](https://react.dev/reference/eslint-plugin-react-hooks) | **Enforcement** - Hooks rules validation |
| [eslint-plugin-jsx-a11y](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y) | **Enforcement** - Accessibility linting |
| [WCAG 2.1](https://www.w3.org/WAI/WCAG21/quickref/) | **Reference** - Accessibility compliance target |

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

Each rule includes:
- **Rule ID**: Unique identifier (e.g., `RCT-CMP-001`)
- **Tier**: Enforcement level
- **Rationale**: Why this rule exists
- **Example**: Correct and incorrect code

### Rule ID Format

Rule IDs follow the pattern `RCT-XXX-NNN` where:
- `RCT` = React
- `XXX` = Category code (see below)
- `NNN` = Sequential number

| Category Code | Area |
|--------------|------|
| `CMP` | Component Patterns |
| `HKS` | Hooks Usage |
| `STT` | State Management |
| `PRP` | Props Patterns |
| `EVT` | Event Handling |
| `RND` | Rendering Patterns |
| `PRF` | Performance |
| `TST` | Testing |
| `A11` | Accessibility |
| `STY` | Styling & JSX |

---

## 3. Project Structure

### 3.1 Directory Layout

```
src/
├── components/           # Shared/reusable components
│   ├── Button/
│   │   ├── Button.tsx
│   │   ├── Button.test.tsx
│   │   ├── Button.stories.tsx
│   │   └── index.ts
│   └── ...
├── features/            # Feature-based modules
│   ├── auth/
│   │   ├── components/
│   │   ├── hooks/
│   │   ├── utils/
│   │   └── index.ts
│   └── ...
├── hooks/               # Shared custom hooks
├── utils/               # Utility functions
├── types/               # TypeScript type definitions
├── constants/           # Application constants
├── context/             # React Context providers
└── App.tsx
```

### 3.2 File Naming

| Type | Convention | Example |
|------|------------|---------|
| Component files | PascalCase with `.tsx` | `UserProfile.tsx` |
| Hook files | camelCase with `use` prefix | `useAuth.ts` |
| Test files | Match source with `.test.tsx` | `UserProfile.test.tsx` |
| Story files | Match source with `.stories.tsx` | `UserProfile.stories.tsx` |
| Utility files | camelCase | `formatDate.ts` |
| Type files | camelCase or PascalCase | `user.types.ts` |
| Index files | `index.ts` | `index.ts` |

### 3.3 Component Organization

Components should be organized by feature when they are feature-specific, or in a shared `components/` directory when reusable across features. Each component should have its own directory containing the component file, tests, stories, and an index file for clean exports.

---

## 4. Naming Conventions

### 4.1 General Principles

- Names should be descriptive and reveal intent
- Use domain terminology consistently
- Component names should describe what they render
- Hook names should describe what they provide

### 4.2 Specific Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Components | PascalCase | `UserProfile`, `NavigationMenu` |
| Component instances | camelCase | `const userProfile = <UserProfile />` |
| Hooks | camelCase with `use` prefix | `useAuth`, `useLocalStorage` |
| Event handlers | camelCase with `handle` prefix | `handleClick`, `handleSubmit` |
| Event handler props | camelCase with `on` prefix | `onClick`, `onSubmit` |
| Boolean props | `is`, `has`, `should` prefix | `isLoading`, `hasError` |
| Context | PascalCase with `Context` suffix | `AuthContext`, `ThemeContext` |
| Providers | PascalCase with `Provider` suffix | `AuthProvider`, `ThemeProvider` |
| HOCs | `with` prefix | `withAuth`, `withTheme` |

### 4.3 Naming Rules

#### RCT-STY-001: Use PascalCase for Component Names :red_circle:

**Rationale**: React treats components starting with lowercase letters as DOM tags. PascalCase clearly distinguishes custom components from HTML elements and is the universally accepted convention.

```tsx
// Correct
function UserProfile() {
  return <div>User Profile</div>;
}

function NavigationMenu() {
  return <nav>Menu</nav>;
}

// Incorrect
function userProfile() {
  return <div>User Profile</div>;
}

function navigation_menu() {
  return <nav>Menu</nav>;
}
```

#### RCT-STY-002: Use camelCase for Prop Names :yellow_circle:

**Rationale**: Follows JavaScript conventions for object properties and maintains consistency with React's built-in props.

```tsx
// Correct
<UserCard
  userName="John"
  isActive={true}
  onProfileClick={handleClick}
  maxItemCount={10}
/>

// Incorrect
<UserCard
  user_name="John"
  IsActive={true}
  OnProfileClick={handleClick}
  MaxItemCount={10}
/>
```

#### RCT-STY-003: Prefix Custom Hooks with `use` :red_circle:

**Rationale**: Required by React to enable the Rules of Hooks linting. The `use` prefix signals to React and developers that the function follows hooks rules.

```tsx
// Correct
function useAuth() {
  const [user, setUser] = useState(null);
  return { user, setUser };
}

function useLocalStorage(key: string) {
  // ...hook implementation
}

// Incorrect
function getAuth() {
  const [user, setUser] = useState(null); // Linter cannot verify hooks rules
  return { user, setUser };
}

function localStorageHook(key: string) {
  // ...
}
```

---

## 5. Component Patterns

### 5.1 Component Types

Modern React development uses functional components exclusively. Class components should only be used for error boundaries until React provides a hooks-based alternative.

### 5.2 Component Rules

#### RCT-CMP-001: Use Functional Components :red_circle:

**Rationale**: Functional components with hooks are the modern React standard. They are simpler, more testable, and provide better performance through easier optimization.

```tsx
// Correct
function UserProfile({ name, email }: UserProfileProps) {
  const [isEditing, setIsEditing] = useState(false);

  return (
    <div>
      <h1>{name}</h1>
      <p>{email}</p>
    </div>
  );
}

// Also correct: Arrow function for simple components
const Avatar = ({ src, alt }: AvatarProps) => (
  <img src={src} alt={alt} className="avatar" />
);

// Incorrect
class UserProfile extends React.Component<UserProfileProps> {
  state = { isEditing: false };

  render() {
    return (
      <div>
        <h1>{this.props.name}</h1>
        <p>{this.props.email}</p>
      </div>
    );
  }
}
```

#### RCT-CMP-002: One Component Per File :yellow_circle:

**Rationale**: Improves maintainability and makes components easier to locate. Exception: Small, tightly coupled helper components may be co-located.

```tsx
// Correct - UserProfile.tsx
function UserProfile({ user }: UserProfileProps) {
  return (
    <div>
      <UserAvatar user={user} />
      <UserDetails user={user} />
    </div>
  );
}

export default UserProfile;

// Incorrect - Multiple unrelated components in one file
function UserProfile({ user }: UserProfileProps) {
  // ...
}

function ProductCard({ product }: ProductCardProps) {
  // Unrelated component in same file
}

function NavigationMenu() {
  // Another unrelated component
}
```

#### RCT-CMP-003: Prefer Composition Over Prop Drilling :yellow_circle:

**Rationale**: Composition using `children` and render props reduces coupling, improves flexibility, and avoids passing props through many levels.

```tsx
// Correct - Using composition
function Card({ children }: { children: React.ReactNode }) {
  return <div className="card">{children}</div>;
}

function App() {
  return (
    <Card>
      <CardHeader title="Welcome" />
      <CardBody>
        <UserProfile user={user} />
      </CardBody>
    </Card>
  );
}

// Correct - Using children for flexible content
function Modal({ children, isOpen, onClose }: ModalProps) {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        {children}
      </div>
    </div>
  );
}

// Incorrect - Excessive prop drilling
function App() {
  return (
    <Card
      headerTitle="Welcome"
      headerSubtitle="Back"
      bodyContent={<UserProfile user={user} />}
      footerActions={actions}
      footerAlignment="right"
    />
  );
}
```

#### RCT-CMP-004: Keep Components Focused and Small :yellow_circle:

**Rationale**: Single-responsibility components are easier to understand, test, and reuse. If a component exceeds 200-300 lines, consider splitting it.

```tsx
// Correct - Focused components
function UserList({ users }: UserListProps) {
  return (
    <ul>
      {users.map(user => (
        <UserListItem key={user.id} user={user} />
      ))}
    </ul>
  );
}

function UserListItem({ user }: UserListItemProps) {
  return (
    <li>
      <UserAvatar user={user} />
      <UserInfo user={user} />
      <UserActions user={user} />
    </li>
  );
}

// Incorrect - Monolithic component doing too much
function UserList({ users }: UserListProps) {
  // 500+ lines handling list, items, avatars, forms, modals...
  return (
    <ul>
      {users.map(user => (
        <li key={user.id}>
          <img src={user.avatar} alt={user.name} />
          <div>
            <h3>{user.name}</h3>
            <p>{user.email}</p>
            {/* ... hundreds more lines */}
          </div>
        </li>
      ))}
    </ul>
  );
}
```

---

## 6. Hooks Usage

### 6.1 Rules of Hooks

The Rules of Hooks are enforced by `eslint-plugin-react-hooks` and are non-negotiable:

1. **Only call hooks at the top level** - Never inside loops, conditions, or nested functions
2. **Only call hooks from React functions** - Either functional components or custom hooks

### 6.2 Hooks Rules

#### RCT-HKS-001: Call Hooks at Top Level Only :red_circle:

**Rationale**: React relies on the order of hooks calls to correctly preserve state between renders. Conditional or loop-based hooks break this mechanism.

```tsx
// Correct
function UserProfile({ userId }: { userId: string }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (userId) {
      fetchUser(userId).then(setUser);
    }
  }, [userId]);

  if (isLoading) return <Spinner />;
  return <div>{user?.name}</div>;
}

// Incorrect
function UserProfile({ userId }: { userId: string }) {
  if (userId) {
    // Hooks inside condition - WILL CAUSE BUGS
    const [user, setUser] = useState<User | null>(null);
    useEffect(() => {
      fetchUser(userId).then(setUser);
    }, [userId]);
  }

  return <div>{user?.name}</div>;
}
```

#### RCT-HKS-002: Include All Dependencies in useEffect :red_circle:

**Rationale**: Missing dependencies cause stale closures and hard-to-debug bugs. The `exhaustive-deps` rule catches these issues.

```tsx
// Correct
function SearchResults({ query, filters }: SearchProps) {
  const [results, setResults] = useState<Result[]>([]);

  useEffect(() => {
    const controller = new AbortController();

    searchAPI(query, filters, controller.signal)
      .then(setResults)
      .catch(err => {
        if (err.name !== 'AbortError') console.error(err);
      });

    return () => controller.abort();
  }, [query, filters]); // All dependencies listed

  return <ResultsList results={results} />;
}

// Incorrect
function SearchResults({ query, filters }: SearchProps) {
  const [results, setResults] = useState<Result[]>([]);

  useEffect(() => {
    // Bug: filters changes won't trigger re-fetch
    searchAPI(query, filters).then(setResults);
  }, [query]); // Missing 'filters' dependency

  return <ResultsList results={results} />;
}
```

#### RCT-HKS-003: Clean Up Effects Properly :red_circle:

**Rationale**: Effects that set up subscriptions, timers, or event listeners must clean up to prevent memory leaks and stale callbacks.

```tsx
// Correct
function WindowSize() {
  const [size, setSize] = useState({ width: 0, height: 0 });

  useEffect(() => {
    function handleResize() {
      setSize({ width: window.innerWidth, height: window.innerHeight });
    }

    handleResize(); // Set initial size
    window.addEventListener('resize', handleResize);

    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return <span>{size.width} x {size.height}</span>;
}

// Incorrect
function WindowSize() {
  const [size, setSize] = useState({ width: 0, height: 0 });

  useEffect(() => {
    function handleResize() {
      setSize({ width: window.innerWidth, height: window.innerHeight });
    }

    window.addEventListener('resize', handleResize);
    // Missing cleanup - memory leak!
  }, []);

  return <span>{size.width} x {size.height}</span>;
}
```

#### RCT-HKS-004: Extract Reusable Logic into Custom Hooks :yellow_circle:

**Rationale**: Custom hooks enable code reuse and separation of concerns. If the same stateful logic appears in multiple components, extract it.

```tsx
// Correct - Reusable custom hook
function useLocalStorage<T>(key: string, initialValue: T) {
  const [storedValue, setStoredValue] = useState<T>(() => {
    try {
      const item = window.localStorage.getItem(key);
      return item ? JSON.parse(item) : initialValue;
    } catch {
      return initialValue;
    }
  });

  const setValue = useCallback((value: T | ((val: T) => T)) => {
    try {
      const valueToStore = value instanceof Function ? value(storedValue) : value;
      setStoredValue(valueToStore);
      window.localStorage.setItem(key, JSON.stringify(valueToStore));
    } catch (error) {
      console.error(error);
    }
  }, [key, storedValue]);

  return [storedValue, setValue] as const;
}

// Usage
function Settings() {
  const [theme, setTheme] = useLocalStorage('theme', 'light');
  const [language, setLanguage] = useLocalStorage('language', 'en');
  // ...
}

// Incorrect - Duplicated logic in multiple components
function Settings() {
  const [theme, setTheme] = useState(() => {
    const saved = localStorage.getItem('theme');
    return saved ? JSON.parse(saved) : 'light';
  });

  useEffect(() => {
    localStorage.setItem('theme', JSON.stringify(theme));
  }, [theme]);
  // ... same pattern repeated for other settings
}
```

#### RCT-HKS-005: Keep Custom Hooks Focused :green_circle:

**Rationale**: A hook should do one thing well. Avoid "god hooks" that manage multiple unrelated concerns.

```tsx
// Correct - Focused hooks
function useUser(userId: string) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchUser(userId)
      .then(setUser)
      .catch(setError)
      .finally(() => setLoading(false));
  }, [userId]);

  return { user, loading, error };
}

function useUserPermissions(user: User | null) {
  return useMemo(() => ({
    canEdit: user?.role === 'admin' || user?.role === 'editor',
    canDelete: user?.role === 'admin',
  }), [user?.role]);
}

// Incorrect - God hook doing too much
function useEverything(userId: string) {
  // Manages user, permissions, notifications, theme, analytics...
  const [user, setUser] = useState(null);
  const [permissions, setPermissions] = useState({});
  const [notifications, setNotifications] = useState([]);
  const [theme, setTheme] = useState('light');
  // 200+ lines of unrelated logic
}
```

---

## 7. State Management

### 7.1 State Management Strategy

Choose the right tool for the scope of state:

| Scope | Solution |
|-------|----------|
| Local component state | `useState` |
| Complex local state with actions | `useReducer` |
| Shared state (small scope) | Lift state up + props |
| Shared state (medium scope) | React Context |
| Global/complex state | External library (Zustand, Jotai, Redux) |
| Server state | React Query, SWR, Apollo |

### 7.2 State Rules

#### RCT-STT-001: Keep State Close to Where It's Used :yellow_circle:

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

#### RCT-STT-002: Use useReducer for Complex State Logic :yellow_circle:

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
    // ... other cases
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

#### RCT-STT-003: Derive State When Possible :yellow_circle:

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

#### RCT-STT-004: Avoid Redundant State :yellow_circle:

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

---

## 8. Props Patterns

### 8.1 Props Best Practices

- Always type props with TypeScript interfaces or types
- Destructure props for cleaner code
- Use sensible default values
- Prefer specific types over `any`

### 8.2 Props Rules

#### RCT-PRP-001: Destructure Props in Function Signature :green_circle:

**Rationale**: Destructuring makes it clear what props a component uses and enables default values inline.

```tsx
// Correct - Destructured in signature
function Button({
  children,
  variant = 'primary',
  size = 'medium',
  disabled = false,
  onClick,
}: ButtonProps) {
  return (
    <button
      className={`btn btn-${variant} btn-${size}`}
      disabled={disabled}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

// Also acceptable - Destructured in body for many props
function ComplexForm(props: ComplexFormProps) {
  const {
    initialValues,
    onSubmit,
    onCancel,
    validationSchema,
    // ... many more props
  } = props;

  // ...
}

// Incorrect - Using props.x throughout
function Button(props: ButtonProps) {
  return (
    <button
      className={`btn btn-${props.variant} btn-${props.size}`}
      disabled={props.disabled}
      onClick={props.onClick}
    >
      {props.children}
    </button>
  );
}
```

#### RCT-PRP-002: Use TypeScript for Prop Types :red_circle:

**Rationale**: TypeScript provides compile-time checking, better IDE support, and self-documenting code.

```tsx
// Correct - TypeScript interface
interface UserCardProps {
  user: {
    id: string;
    name: string;
    email: string;
    avatarUrl?: string;
  };
  isSelected?: boolean;
  onSelect?: (id: string) => void;
}

function UserCard({ user, isSelected = false, onSelect }: UserCardProps) {
  return (
    <div className={isSelected ? 'selected' : ''} onClick={() => onSelect?.(user.id)}>
      <img src={user.avatarUrl ?? '/default-avatar.png'} alt={user.name} />
      <h3>{user.name}</h3>
    </div>
  );
}

// Incorrect - PropTypes or no types
function UserCard({ user, isSelected, onSelect }) {
  // No type safety
  return (
    <div onClick={() => onSelect(user.id)}>
      {/* Runtime error if user is undefined */}
      <h3>{user.name}</h3>
    </div>
  );
}

UserCard.propTypes = {
  user: PropTypes.object,
  isSelected: PropTypes.bool,
};
```

#### RCT-PRP-003: Avoid Spreading Props Blindly :yellow_circle:

**Rationale**: Spreading unknown props can pass invalid attributes to DOM elements and makes component APIs unclear.

```tsx
// Correct - Explicit prop handling
interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary';
  isLoading?: boolean;
}

function Button({
  variant = 'primary',
  isLoading = false,
  disabled,
  children,
  ...domProps
}: ButtonProps) {
  return (
    <button
      {...domProps}
      disabled={disabled || isLoading}
      className={`btn btn-${variant}`}
    >
      {isLoading ? <Spinner /> : children}
    </button>
  );
}

// Incorrect - Blind spreading
function Button(props: any) {
  return (
    <button {...props}>
      {/* Unknown props might cause React warnings */}
      {props.children}
    </button>
  );
}
```

#### RCT-PRP-004: Use Children for Composition :green_circle:

**Rationale**: The `children` prop is the primary mechanism for component composition in React, enabling flexible and reusable components.

```tsx
// Correct - Using children
interface CardProps {
  children: React.ReactNode;
  className?: string;
}

function Card({ children, className }: CardProps) {
  return <div className={`card ${className ?? ''}`}>{children}</div>;
}

// Flexible usage
<Card>
  <CardHeader>Title</CardHeader>
  <CardBody>Any content here</CardBody>
</Card>

// Incorrect - Rigid content props
interface CardProps {
  title: string;
  subtitle?: string;
  content: string;
  footer?: React.ReactNode;
}

function Card({ title, subtitle, content, footer }: CardProps) {
  return (
    <div className="card">
      <h2>{title}</h2>
      {subtitle && <h3>{subtitle}</h3>}
      <p>{content}</p>
      {footer}
    </div>
  );
}
```

---

## 9. Event Handling

### 9.1 Event Handler Conventions

- Name handler functions with `handle` prefix (e.g., `handleClick`)
- Name handler props with `on` prefix (e.g., `onClick`)
- Always type event parameters

### 9.2 Event Handling Rules

#### RCT-EVT-001: Use Consistent Event Handler Naming :yellow_circle:

**Rationale**: Consistent naming makes code predictable and easier to understand. The `handle`/`on` convention is a widely adopted React pattern.

```tsx
// Correct
function LoginForm({ onSubmit, onCancel }: LoginFormProps) {
  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    // ... validation logic
    onSubmit(formData);
  };

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input onChange={handleEmailChange} />
      <button type="button" onClick={onCancel}>Cancel</button>
    </form>
  );
}

// Incorrect
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

#### RCT-EVT-002: Avoid Inline Arrow Functions in JSX (When Performance Critical) :green_circle:

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

---

## 10. Rendering Patterns

### 10.1 Conditional Rendering

Use appropriate patterns for different conditional rendering scenarios.

### 10.2 Rendering Rules

#### RCT-RND-001: Use Appropriate Conditional Rendering Patterns :yellow_circle:

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

// Correct - Object map for multiple conditions
function ComplexStatus({ status }: { status: string }) {
  const statusMap: Record<string, string> = {
    loading: 'Loading...',
    error: 'Error!',
    success: 'Done!',
  };

  return <span>{statusMap[status] ?? 'Unknown'}</span>;
}
```

#### RCT-RND-002: Avoid Short-Circuit with Numbers :red_circle:

**Rationale**: Short-circuit evaluation with falsy numbers (0) will render "0" instead of nothing.

```tsx
// Correct
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
```

#### RCT-RND-003: Use Stable Keys for Lists :red_circle:

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

// Incorrect - Array index as key
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

---

## 11. Performance

### 11.1 Performance Principles

- Measure before optimizing
- Avoid premature optimization
- Understand React's rendering behavior
- Use the right tool for the job

### 11.2 Performance Rules

#### RCT-PRF-001: Memoize Expensive Computations :yellow_circle:

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

#### RCT-PRF-002: Use React.memo for Pure Components :yellow_circle:

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

#### RCT-PRF-003: Avoid Creating Objects in Render :yellow_circle:

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
```

#### RCT-PRF-004: Use useCallback for Function Props :yellow_circle:

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

// Correct - useCallback with dependencies
function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const handleToggle = useCallback((id: string) => {
    setTodos(prev => prev.map(todo =>
      todo.id === id ? { ...todo, completed: !todo.completed } : todo
    ));
  }, []); // No dependencies - uses functional update

  return (
    <ul>
      {todos.map(todo => (
        <MemoizedTodoItem
          key={todo.id}
          todo={todo}
          onToggle={handleToggle}
        />
      ))}
    </ul>
  );
}
```

---

## 12. Testing Requirements

### 12.1 Testing Philosophy

Follow the React Testing Library philosophy: test components the way users interact with them, not implementation details.

### 12.2 Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% | 85% |
| Integration tests | 50% | 70% |

### 12.3 Testing Rules

#### RCT-TST-001: Query by Role and Accessible Name :red_circle:

**Rationale**: Queries by role reflect how assistive technologies see your UI and encourage accessible markup.

```tsx
// Correct - Query by role
import { render, screen } from '@testing-library/react';

test('submits form with user data', async () => {
  const handleSubmit = jest.fn();
  render(<LoginForm onSubmit={handleSubmit} />);

  // Query by role and accessible name
  const emailInput = screen.getByRole('textbox', { name: /email/i });
  const passwordInput = screen.getByLabelText(/password/i);
  const submitButton = screen.getByRole('button', { name: /sign in/i });

  await userEvent.type(emailInput, 'test@example.com');
  await userEvent.type(passwordInput, 'password123');
  await userEvent.click(submitButton);

  expect(handleSubmit).toHaveBeenCalledWith({
    email: 'test@example.com',
    password: 'password123',
  });
});

// Incorrect - Query by test ID or class
test('submits form', async () => {
  render(<LoginForm onSubmit={handleSubmit} />);

  // Avoid: couples test to implementation
  const emailInput = screen.getByTestId('email-input');
  const submitButton = document.querySelector('.submit-btn');
});
```

#### RCT-TST-002: Use userEvent Over fireEvent :yellow_circle:

**Rationale**: `userEvent` simulates real user interactions more accurately than `fireEvent`, including focus, blur, and keyboard events.

```tsx
// Correct - userEvent for realistic interactions
import userEvent from '@testing-library/user-event';

test('user can type in search box and submit', async () => {
  const user = userEvent.setup();
  render(<SearchForm onSearch={mockSearch} />);

  const searchInput = screen.getByRole('searchbox');
  await user.type(searchInput, 'react hooks');
  await user.keyboard('{Enter}');

  expect(mockSearch).toHaveBeenCalledWith('react hooks');
});

// Incorrect - fireEvent for user actions
import { fireEvent } from '@testing-library/react';

test('user can type in search box', () => {
  render(<SearchForm onSearch={mockSearch} />);

  const searchInput = screen.getByRole('searchbox');
  fireEvent.change(searchInput, { target: { value: 'react hooks' } });
  // Less realistic - doesn't trigger focus, keydown, etc.
});
```

#### RCT-TST-003: Test Behavior, Not Implementation :red_circle:

**Rationale**: Tests that check internal state or implementation details break when refactoring, even if behavior is correct.

```tsx
// Correct - Test observable behavior
test('counter increments when button is clicked', async () => {
  const user = userEvent.setup();
  render(<Counter />);

  expect(screen.getByText('Count: 0')).toBeInTheDocument();

  await user.click(screen.getByRole('button', { name: /increment/i }));

  expect(screen.getByText('Count: 1')).toBeInTheDocument();
});

// Incorrect - Test implementation details
test('counter state updates', () => {
  const { result } = renderHook(() => useCounter());

  // Testing internal state, not user-facing behavior
  expect(result.current.state.count).toBe(0);
  act(() => result.current.increment());
  expect(result.current.state.count).toBe(1);
});
```

#### RCT-TST-004: Use waitFor for Async Assertions :yellow_circle:

**Rationale**: Async operations require waiting for state updates. Using `waitFor` or `findBy` queries handles timing correctly.

```tsx
// Correct - waitFor for async updates
test('loads and displays user data', async () => {
  render(<UserProfile userId="123" />);

  // Initially shows loading
  expect(screen.getByText(/loading/i)).toBeInTheDocument();

  // Wait for data to load
  expect(await screen.findByText('John Doe')).toBeInTheDocument();
  expect(screen.queryByText(/loading/i)).not.toBeInTheDocument();
});

// Correct - waitFor for multiple assertions
test('form shows validation errors', async () => {
  const user = userEvent.setup();
  render(<ContactForm />);

  await user.click(screen.getByRole('button', { name: /submit/i }));

  await waitFor(() => {
    expect(screen.getByText(/email is required/i)).toBeInTheDocument();
    expect(screen.getByText(/name is required/i)).toBeInTheDocument();
  });
});

// Incorrect - No waiting for async updates
test('loads user data', () => {
  render(<UserProfile userId="123" />);

  // Fails: assertion runs before data loads
  expect(screen.getByText('John Doe')).toBeInTheDocument();
});
```

---

## 13. Accessibility

### 13.1 Accessibility Requirements

All components must be accessible to users with disabilities, following WCAG 2.1 AA guidelines.

### 13.2 Accessibility Rules

#### RCT-A11-001: Use Semantic HTML Elements :red_circle:

**Rationale**: Semantic HTML provides built-in accessibility features. Divs and spans with click handlers are not accessible.

```tsx
// Correct - Semantic HTML
function Navigation() {
  return (
    <nav aria-label="Main navigation">
      <ul>
        <li><a href="/home">Home</a></li>
        <li><a href="/about">About</a></li>
        <li><a href="/contact">Contact</a></li>
      </ul>
    </nav>
  );
}

function ArticleCard({ article }: { article: Article }) {
  return (
    <article>
      <header>
        <h2>{article.title}</h2>
        <time dateTime={article.date}>{formatDate(article.date)}</time>
      </header>
      <p>{article.summary}</p>
      <footer>
        <a href={`/articles/${article.slug}`}>Read more</a>
      </footer>
    </article>
  );
}

// Incorrect - Divs for everything
function Navigation() {
  return (
    <div className="nav">
      <div className="nav-item" onClick={() => navigate('/home')}>Home</div>
      <div className="nav-item" onClick={() => navigate('/about')}>About</div>
    </div>
  );
}
```

#### RCT-A11-002: Provide Text Alternatives for Images :red_circle:

**Rationale**: Screen readers need text descriptions for images. Empty alt is acceptable only for decorative images.

```tsx
// Correct - Descriptive alt text
function ProductCard({ product }: { product: Product }) {
  return (
    <div>
      <img
        src={product.imageUrl}
        alt={`${product.name} - ${product.color} ${product.category}`}
      />
      <h3>{product.name}</h3>
    </div>
  );
}

// Correct - Empty alt for decorative images
function PageHeader() {
  return (
    <header>
      <img src="/decorative-pattern.svg" alt="" role="presentation" />
      <h1>Welcome</h1>
    </header>
  );
}

// Incorrect - Missing or unhelpful alt
function ProductCard({ product }: { product: Product }) {
  return (
    <div>
      <img src={product.imageUrl} /> {/* Missing alt */}
      <img src={product.imageUrl} alt="image" /> {/* Unhelpful */}
      <img src={product.imageUrl} alt={product.imageUrl} /> {/* URL as alt */}
    </div>
  );
}
```

#### RCT-A11-003: Ensure Keyboard Navigation :red_circle:

**Rationale**: Many users navigate using only a keyboard. All interactive elements must be focusable and operable via keyboard.

```tsx
// Correct - Keyboard accessible custom component
function Dropdown({ options, value, onChange }: DropdownProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [focusedIndex, setFocusedIndex] = useState(0);

  const handleKeyDown = (event: React.KeyboardEvent) => {
    switch (event.key) {
      case 'Enter':
      case ' ':
        event.preventDefault();
        if (isOpen) {
          onChange(options[focusedIndex]);
          setIsOpen(false);
        } else {
          setIsOpen(true);
        }
        break;
      case 'ArrowDown':
        event.preventDefault();
        if (isOpen) {
          setFocusedIndex(i => Math.min(i + 1, options.length - 1));
        } else {
          setIsOpen(true);
        }
        break;
      case 'ArrowUp':
        event.preventDefault();
        setFocusedIndex(i => Math.max(i - 1, 0));
        break;
      case 'Escape':
        setIsOpen(false);
        break;
    }
  };

  return (
    <div
      role="combobox"
      aria-expanded={isOpen}
      aria-haspopup="listbox"
      tabIndex={0}
      onKeyDown={handleKeyDown}
    >
      {/* ... */}
    </div>
  );
}

// Incorrect - Not keyboard accessible
function Dropdown({ options, value, onChange }: DropdownProps) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div onClick={() => setIsOpen(!isOpen)}>
      {/* No keyboard support, no ARIA attributes */}
      <span>{value}</span>
      {isOpen && (
        <div>
          {options.map(opt => (
            <div key={opt} onClick={() => onChange(opt)}>{opt}</div>
          ))}
        </div>
      )}
    </div>
  );
}
```

#### RCT-A11-004: Use ARIA Attributes Correctly :yellow_circle:

**Rationale**: ARIA attributes enhance accessibility but must be used correctly. Incorrect ARIA is worse than no ARIA.

```tsx
// Correct - Proper ARIA usage
function Modal({ isOpen, onClose, title, children }: ModalProps) {
  const titleId = useId();
  const descriptionId = useId();

  if (!isOpen) return null;

  return (
    <div
      role="dialog"
      aria-modal="true"
      aria-labelledby={titleId}
      aria-describedby={descriptionId}
    >
      <h2 id={titleId}>{title}</h2>
      <div id={descriptionId}>{children}</div>
      <button onClick={onClose} aria-label="Close modal">
        <CloseIcon />
      </button>
    </div>
  );
}

// Correct - Live region for dynamic content
function Notification({ message }: { message: string }) {
  return (
    <div role="alert" aria-live="polite">
      {message}
    </div>
  );
}

// Incorrect - Misused ARIA
function Modal({ isOpen, children }: ModalProps) {
  return (
    <div
      role="dialog"
      aria-hidden={!isOpen} // Wrong: aria-hidden hides from AT
      aria-label={children} // Wrong: children is ReactNode, not string
    >
      {children}
    </div>
  );
}
```

#### RCT-A11-005: Maintain Sufficient Color Contrast :yellow_circle:

**Rationale**: Text must have sufficient contrast against its background for readability. WCAG requires 4.5:1 for normal text and 3:1 for large text.

```tsx
// Correct - Sufficient contrast
const theme = {
  colors: {
    text: {
      primary: '#1a1a1a',    // High contrast on white
      secondary: '#4a4a4a',  // Still readable
      onPrimary: '#ffffff',  // White on dark backgrounds
    },
    background: {
      primary: '#ffffff',
      accent: '#0066cc',     // 4.5:1 with white text
    },
  },
};

function Alert({ type, children }: AlertProps) {
  const styles = {
    error: { background: '#d32f2f', color: '#ffffff' },     // 4.6:1
    warning: { background: '#f9a825', color: '#000000' },   // 4.5:1
    success: { background: '#2e7d32', color: '#ffffff' },   // 4.5:1
  };

  return <div style={styles[type]}>{children}</div>;
}

// Incorrect - Insufficient contrast
function LowContrastText() {
  return (
    <p style={{ color: '#999999', background: '#ffffff' }}>
      {/* 2.85:1 contrast - fails WCAG AA */}
      This text is hard to read
    </p>
  );
}
```

---

## 14. Tooling Configuration

### 14.1 Required Tools

| Tool | Purpose | Version |
|------|---------|---------|
| ESLint | Static analysis | ^8.0.0 |
| eslint-plugin-react | React-specific rules | ^7.33.0 |
| eslint-plugin-react-hooks | Hooks rules | ^4.6.0 |
| eslint-plugin-jsx-a11y | Accessibility linting | ^6.7.0 |
| Prettier | Code formatting | ^3.0.0 |
| TypeScript | Type checking | ^5.0.0 |

### 14.2 ESLint Configuration

```javascript
// eslint.config.js (ESLint 9+ flat config)
import js from '@eslint/js';
import tseslint from 'typescript-eslint';
import react from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import jsxA11y from 'eslint-plugin-jsx-a11y';

export default tseslint.config(
  js.configs.recommended,
  ...tseslint.configs.recommended,
  {
    files: ['**/*.{ts,tsx}'],
    plugins: {
      react,
      'react-hooks': reactHooks,
      'jsx-a11y': jsxA11y,
    },
    languageOptions: {
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
      },
    },
    settings: {
      react: {
        version: 'detect',
      },
    },
    rules: {
      // React rules
      'react/jsx-uses-react': 'off', // Not needed with new JSX transform
      'react/react-in-jsx-scope': 'off', // Not needed with new JSX transform
      'react/prop-types': 'off', // Using TypeScript
      'react/jsx-no-target-blank': 'error',
      'react/jsx-key': 'error',
      'react/jsx-no-duplicate-props': 'error',
      'react/jsx-no-undef': 'error',
      'react/no-children-prop': 'error',
      'react/no-danger-with-children': 'error',
      'react/no-deprecated': 'error',
      'react/no-direct-mutation-state': 'error',
      'react/no-unescaped-entities': 'error',
      'react/no-unknown-property': 'error',
      'react/self-closing-comp': 'error',
      'react/jsx-boolean-value': ['error', 'never'],
      'react/jsx-curly-brace-presence': ['error', { props: 'never', children: 'never' }],
      'react/jsx-fragments': ['error', 'syntax'],
      'react/jsx-no-useless-fragment': 'error',
      'react/jsx-pascal-case': 'error',
      'react/no-array-index-key': 'warn',
      'react/no-unstable-nested-components': 'error',
      'react/function-component-definition': ['error', {
        namedComponents: 'function-declaration',
        unnamedComponents: 'arrow-function',
      }],

      // React Hooks rules
      'react-hooks/rules-of-hooks': 'error',
      'react-hooks/exhaustive-deps': 'warn',

      // Accessibility rules
      'jsx-a11y/alt-text': 'error',
      'jsx-a11y/anchor-has-content': 'error',
      'jsx-a11y/anchor-is-valid': 'error',
      'jsx-a11y/aria-props': 'error',
      'jsx-a11y/aria-proptypes': 'error',
      'jsx-a11y/aria-role': 'error',
      'jsx-a11y/aria-unsupported-elements': 'error',
      'jsx-a11y/click-events-have-key-events': 'error',
      'jsx-a11y/heading-has-content': 'error',
      'jsx-a11y/html-has-lang': 'error',
      'jsx-a11y/img-redundant-alt': 'error',
      'jsx-a11y/interactive-supports-focus': 'error',
      'jsx-a11y/label-has-associated-control': 'error',
      'jsx-a11y/media-has-caption': 'error',
      'jsx-a11y/mouse-events-have-key-events': 'error',
      'jsx-a11y/no-access-key': 'error',
      'jsx-a11y/no-autofocus': 'warn',
      'jsx-a11y/no-distracting-elements': 'error',
      'jsx-a11y/no-interactive-element-to-noninteractive-role': 'error',
      'jsx-a11y/no-noninteractive-element-interactions': 'error',
      'jsx-a11y/no-noninteractive-element-to-interactive-role': 'error',
      'jsx-a11y/no-noninteractive-tabindex': 'error',
      'jsx-a11y/no-redundant-roles': 'error',
      'jsx-a11y/no-static-element-interactions': 'error',
      'jsx-a11y/role-has-required-aria-props': 'error',
      'jsx-a11y/role-supports-aria-props': 'error',
      'jsx-a11y/scope': 'error',
      'jsx-a11y/tabindex-no-positive': 'error',
    },
  }
);
```

### 14.3 Legacy ESLint Configuration

```json
// .eslintrc.json (ESLint 8 and below)
{
  "env": {
    "browser": true,
    "es2021": true
  },
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:jsx-a11y/recommended"
  ],
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "ecmaFeatures": {
      "jsx": true
    },
    "ecmaVersion": "latest",
    "sourceType": "module"
  },
  "plugins": [
    "react",
    "react-hooks",
    "@typescript-eslint",
    "jsx-a11y"
  ],
  "settings": {
    "react": {
      "version": "detect"
    }
  },
  "rules": {
    "react/react-in-jsx-scope": "off",
    "react/prop-types": "off",
    "react-hooks/rules-of-hooks": "error",
    "react-hooks/exhaustive-deps": "warn"
  }
}
```

### 14.4 Prettier Configuration

```json
// .prettierrc
{
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "es5",
  "printWidth": 100,
  "jsxSingleQuote": false,
  "bracketSpacing": true,
  "bracketSameLine": false,
  "arrowParens": "avoid"
}
```

### 14.5 TypeScript Configuration

```json
// tsconfig.json
{
  "compilerOptions": {
    "target": "ES2020",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "moduleResolution": "bundler",
    "jsx": "react-jsx",
    "strict": true,
    "noEmit": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true
  },
  "include": ["src"]
}
```

### 14.6 Pre-commit Hooks

```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: eslint
        name: ESLint
        entry: npx eslint --fix
        language: system
        types: [javascript, jsx, typescript, tsx]
        pass_filenames: true

      - id: prettier
        name: Prettier
        entry: npx prettier --write
        language: system
        types: [javascript, jsx, typescript, tsx, json, css, scss]
        pass_filenames: true

      - id: typescript
        name: TypeScript
        entry: npx tsc --noEmit
        language: system
        types: [typescript, tsx]
        pass_filenames: false
```

---

## 15. Code Review Checklist

Quick reference for code reviewers:

### Component Structure
- [ ] Components use functional syntax, not classes
- [ ] One component per file (unless small, tightly coupled helpers)
- [ ] Component names are PascalCase
- [ ] File names match component names
- [ ] Components are focused and reasonably sized (<300 lines)

### Hooks Usage
- [ ] Hooks called at top level only (not in conditions/loops)
- [ ] Custom hooks prefixed with `use`
- [ ] Effect dependencies are complete
- [ ] Effects clean up subscriptions/timers
- [ ] No unnecessary custom hooks (simple logic can be inline)

### State Management
- [ ] State is colocated with where it's used
- [ ] No redundant state (derive when possible)
- [ ] `useReducer` used for complex state logic
- [ ] Context values are memoized

### Props
- [ ] Props are typed with TypeScript interfaces
- [ ] Props are destructured
- [ ] Sensible defaults provided where appropriate
- [ ] `children` used for composition

### Events
- [ ] Event handlers named with `handle` prefix
- [ ] Event props named with `on` prefix
- [ ] Callbacks memoized when passed to memoized children

### Rendering
- [ ] Stable keys used for list items (not array indices)
- [ ] No short-circuit with potentially falsy numbers
- [ ] Conditional rendering uses appropriate pattern
- [ ] No nested ternaries

### Performance
- [ ] `useMemo` used only for expensive computations
- [ ] `React.memo` used appropriately with stable props
- [ ] `useCallback` used for callbacks passed to memoized children
- [ ] No object literals created inline in JSX

### Testing
- [ ] Tests query by role/accessible name
- [ ] Tests use userEvent, not fireEvent
- [ ] Tests check behavior, not implementation
- [ ] Async operations use waitFor/findBy

### Accessibility
- [ ] Semantic HTML used (nav, main, article, button, etc.)
- [ ] Images have meaningful alt text
- [ ] Custom components are keyboard navigable
- [ ] ARIA attributes used correctly
- [ ] Focus management handled for modals/dynamic content

### Security
- [ ] No `dangerouslySetInnerHTML` with user input
- [ ] Links with `target="_blank"` have `rel="noopener noreferrer"`
- [ ] No sensitive data in client-side state

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| RCT-STY-001 | Use PascalCase for Component Names | :red_circle: |
| RCT-STY-002 | Use camelCase for Prop Names | :yellow_circle: |
| RCT-STY-003 | Prefix Custom Hooks with `use` | :red_circle: |
| RCT-CMP-001 | Use Functional Components | :red_circle: |
| RCT-CMP-002 | One Component Per File | :yellow_circle: |
| RCT-CMP-003 | Prefer Composition Over Prop Drilling | :yellow_circle: |
| RCT-CMP-004 | Keep Components Focused and Small | :yellow_circle: |
| RCT-HKS-001 | Call Hooks at Top Level Only | :red_circle: |
| RCT-HKS-002 | Include All Dependencies in useEffect | :red_circle: |
| RCT-HKS-003 | Clean Up Effects Properly | :red_circle: |
| RCT-HKS-004 | Extract Reusable Logic into Custom Hooks | :yellow_circle: |
| RCT-HKS-005 | Keep Custom Hooks Focused | :green_circle: |
| RCT-STT-001 | Keep State Close to Where It's Used | :yellow_circle: |
| RCT-STT-002 | Use useReducer for Complex State Logic | :yellow_circle: |
| RCT-STT-003 | Derive State When Possible | :yellow_circle: |
| RCT-STT-004 | Avoid Redundant State | :yellow_circle: |
| RCT-PRP-001 | Destructure Props in Function Signature | :green_circle: |
| RCT-PRP-002 | Use TypeScript for Prop Types | :red_circle: |
| RCT-PRP-003 | Avoid Spreading Props Blindly | :yellow_circle: |
| RCT-PRP-004 | Use Children for Composition | :green_circle: |
| RCT-EVT-001 | Use Consistent Event Handler Naming | :yellow_circle: |
| RCT-EVT-002 | Avoid Inline Arrow Functions (Performance Critical) | :green_circle: |
| RCT-RND-001 | Use Appropriate Conditional Rendering Patterns | :yellow_circle: |
| RCT-RND-002 | Avoid Short-Circuit with Numbers | :red_circle: |
| RCT-RND-003 | Use Stable Keys for Lists | :red_circle: |
| RCT-PRF-001 | Memoize Expensive Computations | :yellow_circle: |
| RCT-PRF-002 | Use React.memo for Pure Components | :yellow_circle: |
| RCT-PRF-003 | Avoid Creating Objects in Render | :yellow_circle: |
| RCT-PRF-004 | Use useCallback for Function Props | :yellow_circle: |
| RCT-TST-001 | Query by Role and Accessible Name | :red_circle: |
| RCT-TST-002 | Use userEvent Over fireEvent | :yellow_circle: |
| RCT-TST-003 | Test Behavior, Not Implementation | :red_circle: |
| RCT-TST-004 | Use waitFor for Async Assertions | :yellow_circle: |
| RCT-A11-001 | Use Semantic HTML Elements | :red_circle: |
| RCT-A11-002 | Provide Text Alternatives for Images | :red_circle: |
| RCT-A11-003 | Ensure Keyboard Navigation | :red_circle: |
| RCT-A11-004 | Use ARIA Attributes Correctly | :yellow_circle: |
| RCT-A11-005 | Maintain Sufficient Color Contrast | :yellow_circle: |

---

## Appendix B: Glossary

| Term | Definition |
|------|------------|
| Component | A reusable piece of UI that can accept inputs (props) and return React elements |
| Hook | A function that lets you use React features in functional components |
| Props | Short for "properties" - read-only inputs passed to components |
| State | Data that changes over time and triggers re-renders when updated |
| Effect | Side effects that run after render, managed by useEffect |
| Memoization | Caching technique to avoid recalculating values or re-rendering |
| JSX | JavaScript syntax extension that allows writing HTML-like code in JavaScript |
| ARIA | Accessible Rich Internet Applications - attributes for accessibility |
| WCAG | Web Content Accessibility Guidelines |
| a11y | Numeronym for "accessibility" (11 letters between 'a' and 'y') |

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- [Airbnb React/JSX Style Guide](https://airbnb.io/javascript/react/)
- [React Documentation](https://react.dev)
- [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react)
- [eslint-plugin-react-hooks](https://react.dev/reference/eslint-plugin-react-hooks)
- [eslint-plugin-jsx-a11y](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- [Kent C. Dodds - Common Mistakes with React Testing Library](https://kentcdodds.com/blog/common-mistakes-with-react-testing-library)
- [WCAG 2.1 Quick Reference](https://www.w3.org/WAI/WCAG21/quickref/)
- [React Accessibility Documentation](https://legacy.reactjs.org/docs/accessibility.html)
