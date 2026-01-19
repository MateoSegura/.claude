# Accessibility (RCT-A11-*)

Accessibility ensures applications work for all users, including those with disabilities. These rules align with WCAG 2.1 AA guidelines.

## Accessibility Requirements

All components must be accessible to users with disabilities, following WCAG 2.1 AA guidelines:
- Perceivable - Information must be presentable in ways users can perceive
- Operable - UI components must be operable by all users
- Understandable - Information and UI operation must be understandable
- Robust - Content must be robust enough for assistive technologies

---

## RCT-A11-001: Use Semantic HTML Elements :red_circle:

**Tier**: Critical

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

function SearchForm({ onSearch }: { onSearch: (q: string) => void }) {
  return (
    <form role="search" onSubmit={handleSubmit}>
      <label htmlFor="search">Search</label>
      <input type="search" id="search" name="search" />
      <button type="submit">Search</button>
    </form>
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

**Semantic elements to use**:
- `<nav>` for navigation
- `<main>` for main content
- `<article>` for self-contained content
- `<section>` for thematic grouping
- `<header>`, `<footer>` for headers/footers
- `<button>` for clickable actions
- `<a>` for navigation links
- `<form>`, `<label>`, `<input>` for forms

---

## RCT-A11-002: Provide Text Alternatives for Images :red_circle:

**Tier**: Critical

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

// Correct - Icon buttons with accessible labels
function IconButton({ icon, label, onClick }: IconButtonProps) {
  return (
    <button onClick={onClick} aria-label={label}>
      <img src={icon} alt="" />
    </button>
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

---

## RCT-A11-003: Ensure Keyboard Navigation :red_circle:

**Tier**: Critical

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

// Correct - Custom button with keyboard support
function CustomButton({ onClick, children }: CustomButtonProps) {
  const handleKeyDown = (event: React.KeyboardEvent) => {
    if (event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      onClick();
    }
  };

  return (
    <div
      role="button"
      tabIndex={0}
      onClick={onClick}
      onKeyDown={handleKeyDown}
    >
      {children}
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

---

## RCT-A11-004: Use ARIA Attributes Correctly :yellow_circle:

**Tier**: Required

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

// Correct - Progress indicator
function LoadingBar({ progress }: { progress: number }) {
  return (
    <div
      role="progressbar"
      aria-valuenow={progress}
      aria-valuemin={0}
      aria-valuemax={100}
      aria-label="Loading progress"
    >
      <div style={{ width: `${progress}%` }} />
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

---

## RCT-A11-005: Maintain Sufficient Color Contrast :yellow_circle:

**Tier**: Required

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

  return <div style={styles[type]} role="alert">{children}</div>;
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

**Contrast Requirements**:
- Normal text (<18px): 4.5:1 ratio minimum
- Large text (18px+ or 14px+ bold): 3:1 ratio minimum
- UI components and graphical objects: 3:1 ratio minimum

**Tools for checking contrast**:
- WebAIM Contrast Checker
- Chrome DevTools Accessibility panel
- axe DevTools browser extension
