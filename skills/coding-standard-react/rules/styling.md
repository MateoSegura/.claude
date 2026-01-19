# Styling & Naming Conventions (RCT-STY-*)

Consistent naming conventions make code predictable and maintainable. These rules establish conventions for naming components, hooks, props, and other React elements.

## General Principles

- Names should be descriptive and reveal intent
- Use domain terminology consistently
- Component names should describe what they render
- Hook names should describe what they provide

## Naming Conventions Summary

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

---

## RCT-STY-001: Use PascalCase for Component Names :red_circle:

**Tier**: Critical

**Rationale**: React treats components starting with lowercase letters as DOM tags. PascalCase clearly distinguishes custom components from HTML elements and is the universally accepted convention.

```tsx
// Correct
function UserProfile() {
  return <div>User Profile</div>;
}

function NavigationMenu() {
  return <nav>Menu</nav>;
}

const ButtonGroup = ({ children }: ButtonGroupProps) => (
  <div className="button-group">{children}</div>
);

// Incorrect
function userProfile() {
  return <div>User Profile</div>;
}

function navigation_menu() {
  return <nav>Menu</nav>;
}

const button_group = ({ children }) => (
  <div>{children}</div>
);
```

---

## RCT-STY-002: Use camelCase for Prop Names :yellow_circle:

**Tier**: Required

**Rationale**: Follows JavaScript conventions for object properties and maintains consistency with React's built-in props.

```tsx
// Correct
<UserCard
  userName="John"
  isActive={true}
  onProfileClick={handleClick}
  maxItemCount={10}
/>

interface UserCardProps {
  userName: string;
  isActive: boolean;
  onProfileClick: () => void;
  maxItemCount: number;
}

// Incorrect
<UserCard
  user_name="John"
  IsActive={true}
  OnProfileClick={handleClick}
  MaxItemCount={10}
/>
```

---

## RCT-STY-003: Prefix Custom Hooks with `use` :red_circle:

**Tier**: Critical

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

function useDebounce<T>(value: T, delay: number): T {
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

function AuthHook() { // Wrong: PascalCase implies component
  // ...
}
```

---

## File Naming Conventions

| Type | Convention | Example |
|------|------------|---------|
| Component files | PascalCase with `.tsx` | `UserProfile.tsx` |
| Hook files | camelCase with `use` prefix | `useAuth.ts` |
| Test files | Match source with `.test.tsx` | `UserProfile.test.tsx` |
| Story files | Match source with `.stories.tsx` | `UserProfile.stories.tsx` |
| Utility files | camelCase | `formatDate.ts` |
| Type files | camelCase or PascalCase | `user.types.ts` |
| Index files | `index.ts` | `index.ts` |

---

## Directory Structure

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

---

## Boolean Prop Naming

Use prefixes that indicate boolean nature:

```tsx
// Correct
interface ButtonProps {
  isLoading: boolean;
  isDisabled: boolean;
  hasError: boolean;
  shouldAnimate: boolean;
  canSubmit: boolean;
}

// Usage
<Button
  isLoading={loading}
  isDisabled={!isValid}
  hasError={errors.length > 0}
/>

// Incorrect
interface ButtonProps {
  loading: boolean;     // Ambiguous - could be a noun
  disabled: boolean;    // OK for native attributes, not custom
  error: boolean;       // Ambiguous
  animate: boolean;     // Ambiguous - could be a verb
}
```

---

## Context and Provider Naming

```tsx
// Correct
const AuthContext = createContext<AuthContextType | null>(null);

function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
}

function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within AuthProvider');
  }
  return context;
}

// Incorrect
const Auth = createContext(null);           // Missing Context suffix
function AuthWrapper({ children }) { ... }  // Ambiguous name
function getAuth() { ... }                  // Missing use prefix
```
