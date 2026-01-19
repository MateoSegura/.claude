# React Testing Library Setup

React Testing Library is the recommended testing library for React applications. It encourages testing components the way users interact with them.

## Installation

```bash
# Core Testing Library packages
npm install --save-dev @testing-library/react @testing-library/dom

# User event library (for realistic interactions)
npm install --save-dev @testing-library/user-event

# Jest DOM matchers
npm install --save-dev @testing-library/jest-dom

# Jest (if not already installed)
npm install --save-dev jest @types/jest ts-jest
```

## Required Versions

- **@testing-library/react**: ^14.0.0
- **@testing-library/user-event**: ^14.0.0
- **@testing-library/jest-dom**: ^6.0.0
- **Jest**: ^29.0.0

---

## Jest Configuration

Create `jest.config.js`:

```javascript
/** @type {import('jest').Config} */
const config = {
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/src/setupTests.ts'],
  moduleNameMapper: {
    // Handle CSS imports
    '\\.(css|less|scss|sass)$': 'identity-obj-proxy',
    // Handle image imports
    '\\.(jpg|jpeg|png|gif|webp|svg)$': '<rootDir>/__mocks__/fileMock.js',
    // Handle module aliases
    '^@/(.*)$': '<rootDir>/src/$1',
  },
  transform: {
    '^.+\\.(ts|tsx)$': 'ts-jest',
  },
  testMatch: [
    '<rootDir>/src/**/*.test.{ts,tsx}',
    '<rootDir>/src/**/*.spec.{ts,tsx}',
  ],
  collectCoverageFrom: [
    'src/**/*.{ts,tsx}',
    '!src/**/*.d.ts',
    '!src/**/*.stories.{ts,tsx}',
    '!src/index.tsx',
  ],
  coverageThreshold: {
    global: {
      branches: 70,
      functions: 70,
      lines: 70,
      statements: 70,
    },
  },
};

module.exports = config;
```

---

## Setup File

Create `src/setupTests.ts`:

```typescript
import '@testing-library/jest-dom';

// Mock window.matchMedia
Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: jest.fn().mockImplementation(query => ({
    matches: false,
    media: query,
    onchange: null,
    addListener: jest.fn(),
    removeListener: jest.fn(),
    addEventListener: jest.fn(),
    removeEventListener: jest.fn(),
    dispatchEvent: jest.fn(),
  })),
});

// Mock IntersectionObserver
class MockIntersectionObserver {
  observe = jest.fn();
  disconnect = jest.fn();
  unobserve = jest.fn();
}

Object.defineProperty(window, 'IntersectionObserver', {
  writable: true,
  value: MockIntersectionObserver,
});

// Clean up after each test
afterEach(() => {
  jest.clearAllMocks();
});
```

---

## Test File Structure

```typescript
// src/components/Button/Button.test.tsx
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { Button } from './Button';

describe('Button', () => {
  it('renders button text', () => {
    render(<Button>Click me</Button>);

    expect(screen.getByRole('button', { name: /click me/i })).toBeInTheDocument();
  });

  it('calls onClick when clicked', async () => {
    const user = userEvent.setup();
    const handleClick = jest.fn();

    render(<Button onClick={handleClick}>Click me</Button>);

    await user.click(screen.getByRole('button'));

    expect(handleClick).toHaveBeenCalledTimes(1);
  });

  it('is disabled when disabled prop is true', () => {
    render(<Button disabled>Click me</Button>);

    expect(screen.getByRole('button')).toBeDisabled();
  });
});
```

---

## Query Priority

Use queries in this order of preference:

```typescript
// 1. getByRole - Most accessible
screen.getByRole('button', { name: /submit/i });
screen.getByRole('textbox', { name: /email/i });
screen.getByRole('heading', { level: 1 });

// 2. getByLabelText - For form fields
screen.getByLabelText(/email address/i);

// 3. getByPlaceholderText - If no label
screen.getByPlaceholderText(/enter email/i);

// 4. getByText - For non-interactive elements
screen.getByText(/welcome back/i);

// 5. getByDisplayValue - For filled inputs
screen.getByDisplayValue('current@email.com');

// 6. getByAltText - For images
screen.getByAltText(/user avatar/i);

// 7. getByTitle - If no other option
screen.getByTitle(/close modal/i);

// 8. getByTestId - Last resort
screen.getByTestId('custom-element');
```

---

## userEvent vs fireEvent

Always prefer userEvent:

```typescript
// Correct - userEvent
import userEvent from '@testing-library/user-event';

test('user interaction', async () => {
  const user = userEvent.setup();

  await user.click(button);
  await user.type(input, 'hello');
  await user.keyboard('{Enter}');
  await user.selectOptions(select, 'option1');
  await user.hover(element);
  await user.tab();
});

// Avoid - fireEvent
import { fireEvent } from '@testing-library/react';

test('user interaction', () => {
  fireEvent.click(button);
  fireEvent.change(input, { target: { value: 'hello' } });
});
```

---

## Async Testing

```typescript
// Using findBy for async elements
test('loads data', async () => {
  render(<UserProfile userId="123" />);

  // findBy waits for element to appear
  expect(await screen.findByText('John Doe')).toBeInTheDocument();
});

// Using waitFor for multiple assertions
test('form validation', async () => {
  const user = userEvent.setup();
  render(<ContactForm />);

  await user.click(screen.getByRole('button', { name: /submit/i }));

  await waitFor(() => {
    expect(screen.getByText(/email required/i)).toBeInTheDocument();
    expect(screen.getByText(/name required/i)).toBeInTheDocument();
  });
});

// Using waitForElementToBeRemoved
test('loading state', async () => {
  render(<DataLoader />);

  expect(screen.getByText(/loading/i)).toBeInTheDocument();

  await waitForElementToBeRemoved(() => screen.queryByText(/loading/i));

  expect(screen.getByText('Data loaded')).toBeInTheDocument();
});
```

---

## Custom Render

Create a custom render for providers:

```typescript
// src/test-utils.tsx
import { ReactElement } from 'react';
import { render, RenderOptions } from '@testing-library/react';
import { ThemeProvider } from './context/ThemeContext';
import { AuthProvider } from './context/AuthContext';

const AllProviders = ({ children }: { children: React.ReactNode }) => {
  return (
    <ThemeProvider>
      <AuthProvider>
        {children}
      </AuthProvider>
    </ThemeProvider>
  );
};

const customRender = (
  ui: ReactElement,
  options?: Omit<RenderOptions, 'wrapper'>
) => render(ui, { wrapper: AllProviders, ...options });

export * from '@testing-library/react';
export { customRender as render };
```

---

## Package.json Scripts

```json
{
  "scripts": {
    "test": "jest",
    "test:watch": "jest --watch",
    "test:coverage": "jest --coverage",
    "test:ci": "jest --ci --coverage --reporters=default --reporters=jest-junit"
  }
}
```
