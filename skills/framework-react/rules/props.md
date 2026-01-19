# Props Patterns (RCT-PRP-*)

Props are the primary mechanism for passing data between components. These rules ensure props are well-typed, organized, and used effectively.

## Props Best Practices

- Always type props with TypeScript interfaces or types
- Destructure props for cleaner code
- Use sensible default values
- Prefer specific types over `any`

---

## RCT-PRP-001: Destructure Props in Function Signature :green_circle:

**Tier**: Recommended

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

---

## RCT-PRP-002: Use TypeScript for Prop Types :red_circle:

**Tier**: Critical

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

// Correct - Extending HTML element props
interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'danger';
  isLoading?: boolean;
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

---

## RCT-PRP-003: Avoid Spreading Props Blindly :yellow_circle:

**Tier**: Required

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

// Correct - Separating custom props from DOM props
interface InputFieldProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
  error?: string;
}

function InputField({ label, error, ...inputProps }: InputFieldProps) {
  return (
    <div className="field">
      <label>{label}</label>
      <input {...inputProps} aria-invalid={!!error} />
      {error && <span className="error">{error}</span>}
    </div>
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

---

## RCT-PRP-004: Use Children for Composition :green_circle:

**Tier**: Recommended

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

// Correct - Render props for more control
interface DataFetcherProps<T> {
  url: string;
  children: (data: T, loading: boolean) => React.ReactNode;
}

function DataFetcher<T>({ url, children }: DataFetcherProps<T>) {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(url)
      .then(res => res.json())
      .then(setData)
      .finally(() => setLoading(false));
  }, [url]);

  return <>{children(data as T, loading)}</>;
}

// Usage
<DataFetcher url="/api/users">
  {(users, loading) => (
    loading ? <Spinner /> : <UserList users={users} />
  )}
</DataFetcher>

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

**Types for children**:
- `React.ReactNode` - Most flexible, accepts anything renderable
- `React.ReactElement` - Only accepts JSX elements
- `string` - Only accepts text
- `(args: T) => React.ReactNode` - Render prop pattern
