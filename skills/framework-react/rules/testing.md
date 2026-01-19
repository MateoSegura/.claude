# Testing (RCT-TST-*)

Testing React components ensures reliability and enables confident refactoring. These rules follow the React Testing Library philosophy of testing components the way users interact with them.

## Testing Philosophy

- Test components the way users interact with them
- Query by accessibility roles, not implementation details
- Use userEvent for realistic user interactions
- Avoid testing implementation details

## Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% | 85% |
| Integration tests | 50% | 70% |

---

## RCT-TST-001: Query by Role and Accessible Name :red_circle:

**Tier**: Critical

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

**Query Priority (most to least preferred)**:
1. `getByRole` - Accessible to everyone
2. `getByLabelText` - For form fields
3. `getByPlaceholderText` - If label unavailable
4. `getByText` - For non-interactive elements
5. `getByDisplayValue` - For filled-in form elements
6. `getByAltText` - For images
7. `getByTitle` - If no other option
8. `getByTestId` - Last resort

---

## RCT-TST-002: Use userEvent Over fireEvent :yellow_circle:

**Tier**: Required

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

test('user can select option from dropdown', async () => {
  const user = userEvent.setup();
  render(<SelectField options={['Red', 'Green', 'Blue']} />);

  const select = screen.getByRole('combobox');
  await user.selectOptions(select, 'Green');

  expect(select).toHaveValue('Green');
});

test('user can check and uncheck checkbox', async () => {
  const user = userEvent.setup();
  render(<Checkbox label="Accept terms" />);

  const checkbox = screen.getByRole('checkbox');
  expect(checkbox).not.toBeChecked();

  await user.click(checkbox);
  expect(checkbox).toBeChecked();

  await user.click(checkbox);
  expect(checkbox).not.toBeChecked();
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

---

## RCT-TST-003: Test Behavior, Not Implementation :red_circle:

**Tier**: Critical

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

// Correct - Test user-visible outcomes
test('shows error message for invalid email', async () => {
  const user = userEvent.setup();
  render(<LoginForm />);

  await user.type(screen.getByLabelText(/email/i), 'invalid-email');
  await user.click(screen.getByRole('button', { name: /submit/i }));

  expect(screen.getByRole('alert')).toHaveTextContent(/valid email/i);
});

// Correct - Test integration behavior
test('filtered list updates as user types', async () => {
  const user = userEvent.setup();
  render(<FilterableList items={items} />);

  expect(screen.getAllByRole('listitem')).toHaveLength(10);

  await user.type(screen.getByRole('searchbox'), 'react');

  expect(screen.getAllByRole('listitem')).toHaveLength(3);
});

// Incorrect - Test implementation details
test('counter state updates', () => {
  const { result } = renderHook(() => useCounter());

  // Testing internal state, not user-facing behavior
  expect(result.current.state.count).toBe(0);
  act(() => result.current.increment());
  expect(result.current.state.count).toBe(1);
});

// Incorrect - Testing internal function calls
test('calls internal validate function', () => {
  const validateSpy = jest.spyOn(validation, 'validateEmail');
  render(<LoginForm />);

  // Implementation detail - doesn't test user-visible behavior
  fireEvent.blur(screen.getByLabelText(/email/i));
  expect(validateSpy).toHaveBeenCalled();
});
```

---

## RCT-TST-004: Use waitFor for Async Assertions :yellow_circle:

**Tier**: Required

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

// Correct - findBy for single async element
test('displays search results', async () => {
  const user = userEvent.setup();
  render(<SearchPage />);

  await user.type(screen.getByRole('searchbox'), 'react');
  await user.click(screen.getByRole('button', { name: /search/i }));

  // findBy waits for element to appear
  const results = await screen.findAllByRole('listitem');
  expect(results).toHaveLength(5);
});

// Incorrect - No waiting for async updates
test('loads user data', () => {
  render(<UserProfile userId="123" />);

  // Fails: assertion runs before data loads
  expect(screen.getByText('John Doe')).toBeInTheDocument();
});
```

---

## Testing Custom Hooks

```tsx
import { renderHook, act } from '@testing-library/react';

test('useCounter increments and decrements', () => {
  const { result } = renderHook(() => useCounter(0));

  expect(result.current.count).toBe(0);

  act(() => {
    result.current.increment();
  });
  expect(result.current.count).toBe(1);

  act(() => {
    result.current.decrement();
  });
  expect(result.current.count).toBe(0);
});
```
