# Component Patterns (RCT-CMP-*)

Components are the building blocks of React applications. These rules ensure components are well-structured, maintainable, and follow modern best practices.

## Component Philosophy

- Use functional components exclusively (except error boundaries)
- Keep components focused on a single responsibility
- Prefer composition over prop drilling
- Components should be easy to test in isolation

---

## RCT-CMP-001: Use Functional Components :red_circle:

**Tier**: Critical

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

**Exception**: Error boundaries currently require class components until React provides a hooks-based alternative.

---

## RCT-CMP-002: One Component Per File :yellow_circle:

**Tier**: Required

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

**Acceptable Exception**: Small, tightly coupled sub-components that are only used within a parent component may be co-located:

```tsx
// Acceptable - ListItem only used within List
function ListItem({ item }: { item: Item }) {
  return <li>{item.name}</li>;
}

function List({ items }: ListProps) {
  return (
    <ul>
      {items.map(item => (
        <ListItem key={item.id} item={item} />
      ))}
    </ul>
  );
}

export default List;
```

---

## RCT-CMP-003: Prefer Composition Over Prop Drilling :yellow_circle:

**Tier**: Required

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

---

## RCT-CMP-004: Keep Components Focused and Small :yellow_circle:

**Tier**: Required

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

**Guidelines for splitting components**:
- If a component handles multiple unrelated concerns, split by concern
- If JSX becomes deeply nested (>3-4 levels), extract sub-components
- If the same logic appears in multiple places, extract to a custom hook
- If a section of JSX could be meaningfully named, it might be a component
