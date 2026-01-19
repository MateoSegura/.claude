# Hooks Usage (RCT-HKS-*)

Hooks are the primary way to manage state and side effects in functional components. These rules ensure hooks are used correctly and effectively.

## Rules of Hooks

The Rules of Hooks are enforced by `eslint-plugin-react-hooks` and are non-negotiable:

1. **Only call hooks at the top level** - Never inside loops, conditions, or nested functions
2. **Only call hooks from React functions** - Either functional components or custom hooks

---

## RCT-HKS-001: Call Hooks at Top Level Only :red_circle:

**Tier**: Critical

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

// Incorrect - Hook in loop
function MultiCounter({ count }: { count: number }) {
  const counters = [];
  for (let i = 0; i < count; i++) {
    counters.push(useState(0)); // WRONG!
  }
  return <div>{/* ... */}</div>;
}
```

---

## RCT-HKS-002: Include All Dependencies in useEffect :red_circle:

**Tier**: Critical

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

**If a dependency causes too many re-runs**:
- Move the function inside useEffect
- Wrap the function with useCallback
- Use a ref for values that shouldn't trigger re-runs
- Consider if the effect is doing too much

---

## RCT-HKS-003: Clean Up Effects Properly :red_circle:

**Tier**: Critical

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

// Correct - AbortController for fetch
function UserData({ userId }: { userId: string }) {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const controller = new AbortController();

    fetch(`/api/users/${userId}`, { signal: controller.signal })
      .then(res => res.json())
      .then(setUser)
      .catch(err => {
        if (err.name !== 'AbortError') throw err;
      });

    return () => controller.abort();
  }, [userId]);

  return <div>{user?.name}</div>;
}

// Incorrect - Memory leak
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

---

## RCT-HKS-004: Extract Reusable Logic into Custom Hooks :yellow_circle:

**Tier**: Required

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

---

## RCT-HKS-005: Keep Custom Hooks Focused :green_circle:

**Tier**: Recommended

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

**Custom Hook Naming Guidelines**:
- Always start with `use`
- Name should describe what the hook provides, not how it works
- Examples: `useAuth`, `useLocalStorage`, `useDebounce`, `useMediaQuery`
