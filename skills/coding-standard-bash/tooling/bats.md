# Bats Testing Framework

Bats (Bash Automated Testing System) is the standard testing framework for shell scripts.

## Installation

```bash
# Using npm
npm install -g bats

# Using Homebrew (macOS)
brew install bats-core

# From source
git clone https://github.com/bats-core/bats-core.git
cd bats-core
./install.sh /usr/local

# Install helper libraries
git clone https://github.com/bats-core/bats-support.git test/test_helper/bats-support
git clone https://github.com/bats-core/bats-assert.git test/test_helper/bats-assert
```

## Required Version

- **Bats**: 1.10.0+

---

## Test File Structure

```
project/
├── scripts/
│   ├── deploy.sh
│   └── utils.sh
└── tests/
    ├── test_helper/
    │   ├── bats-support/
    │   ├── bats-assert/
    │   └── common.bash
    ├── deploy_test.bats
    ├── utils_test.bats
    └── fixtures/
        └── sample_config.json
```

---

## Basic Test File

```bash
#!/usr/bin/env bats

# Load test helpers
load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

# Setup runs before each test
setup() {
  # Load script being tested
  source "${BATS_TEST_DIRNAME}/../scripts/utils.sh"

  # Create temp directory
  TEST_TEMP_DIR="$(mktemp -d)"
}

# Teardown runs after each test
teardown() {
  rm -rf "${TEST_TEMP_DIR}"
}

# Test functions start with @test
@test "function returns success for valid input" {
  run validate_input "valid"
  assert_success
}

@test "function returns failure for invalid input" {
  run validate_input ""
  assert_failure
}

@test "function output contains expected string" {
  run get_greeting "World"
  assert_success
  assert_output "Hello, World!"
}

@test "function handles files with spaces" {
  local test_file="${TEST_TEMP_DIR}/file with spaces.txt"
  echo "content" > "${test_file}"

  run process_file "${test_file}"
  assert_success
}
```

---

## Running Tests

```bash
# Run all tests in directory
bats tests/

# Run specific test file
bats tests/utils_test.bats

# Run with verbose output
bats --verbose-run tests/

# Run with timing
bats --timing tests/

# Run in parallel (faster)
bats --jobs 4 tests/

# Output formats
bats --formatter tap tests/              # TAP format
bats --formatter junit tests/            # JUnit XML
bats --formatter pretty tests/           # Pretty print (default)

# Filter tests by name
bats --filter "valid" tests/             # Run tests matching "valid"
```

---

## Assertions (bats-assert)

```bash
# Load bats-assert
load 'test_helper/bats-assert/load'

@test "assert_success - command exits 0" {
  run echo "hello"
  assert_success
}

@test "assert_failure - command exits non-zero" {
  run false
  assert_failure
}

@test "assert_failure with specific code" {
  run bash -c "exit 2"
  assert_failure 2
}

@test "assert_output - exact match" {
  run echo "hello"
  assert_output "hello"
}

@test "assert_output - partial match" {
  run echo "hello world"
  assert_output --partial "world"
}

@test "assert_output - regex match" {
  run echo "hello123"
  assert_output --regexp "^hello[0-9]+$"
}

@test "assert_line - specific line" {
  run echo -e "line1\nline2\nline3"
  assert_line --index 1 "line2"
}

@test "refute_output - should not contain" {
  run echo "success"
  refute_output --partial "error"
}
```

---

## Common Test Helper

```bash
# tests/test_helper/common.bash

# Common setup for all tests
setup_common() {
  TEST_TEMP_DIR="$(mktemp -d)"
  export TEST_TEMP_DIR

  # Add scripts to PATH
  export PATH="${BATS_TEST_DIRNAME}/../scripts:${PATH}"

  # Load common utilities
  source "${BATS_TEST_DIRNAME}/../scripts/utils.sh"
}

teardown_common() {
  rm -rf "${TEST_TEMP_DIR}"
}

# Helper to create test file
create_test_file() {
  local filename="$1"
  local content="${2:-test content}"
  local filepath="${TEST_TEMP_DIR}/${filename}"

  mkdir -p "$(dirname "${filepath}")"
  echo "${content}" > "${filepath}"
  echo "${filepath}"
}

# Helper to assert file exists
assert_file_exists() {
  local file="$1"
  if [[ ! -f "${file}" ]]; then
    echo "Expected file to exist: ${file}" >&2
    return 1
  fi
}

# Helper to assert file contains
assert_file_contains() {
  local file="$1"
  local pattern="$2"

  if ! grep -q "${pattern}" "${file}"; then
    echo "Expected file to contain: ${pattern}" >&2
    echo "File contents:" >&2
    cat "${file}" >&2
    return 1
  fi
}
```

Usage in test file:

```bash
#!/usr/bin/env bats

load 'test_helper/common'

setup() {
  setup_common
}

teardown() {
  teardown_common
}

@test "creates output file" {
  local input_file
  input_file=$(create_test_file "input.txt" "test data")

  run process_file "${input_file}" "${TEST_TEMP_DIR}/output.txt"

  assert_success
  assert_file_exists "${TEST_TEMP_DIR}/output.txt"
  assert_file_contains "${TEST_TEMP_DIR}/output.txt" "processed"
}
```

---

## Mocking Commands

```bash
# Create mock in setup
setup() {
  TEST_TEMP_DIR="$(mktemp -d)"
  MOCK_BIN="${TEST_TEMP_DIR}/bin"
  mkdir -p "${MOCK_BIN}"
  export PATH="${MOCK_BIN}:${PATH}"
}

@test "handles curl failure gracefully" {
  # Create mock curl that fails
  cat > "${MOCK_BIN}/curl" << 'EOF'
#!/usr/bin/env bash
echo "Connection refused" >&2
exit 1
EOF
  chmod +x "${MOCK_BIN}/curl"

  run download_file "http://example.com/file"

  assert_failure
  assert_output --partial "Download failed"
}

@test "processes API response correctly" {
  # Create mock curl with fixed output
  cat > "${MOCK_BIN}/curl" << 'EOF'
#!/usr/bin/env bash
echo '{"status": "ok", "data": "test"}'
EOF
  chmod +x "${MOCK_BIN}/curl"

  run fetch_data "http://api.example.com"

  assert_success
  assert_output --partial "test"
}
```

---

## CI Integration

### GitHub Actions

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Bats
        uses: bats-core/bats-action@1.5.0

      - name: Run tests
        run: bats --formatter junit tests/ > test-results.xml

      - name: Upload results
        uses: actions/upload-artifact@v3
        with:
          name: test-results
          path: test-results.xml
```

### GitLab CI

```yaml
test:
  image: bats/bats:latest
  script:
    - bats --tap tests/
  artifacts:
    reports:
      junit: test-results.xml
```

---

## Debugging Tests

```bash
# Run single test with verbose output
bats --verbose-run --filter "specific test name" tests/

# Print debug info in test
@test "debug example" {
  echo "# Debug: variable=${variable}" >&3
  run some_command
  echo "# Debug: output=${output}" >&3
  echo "# Debug: status=${status}" >&3
  assert_success
}

# Use set -x in tested script
@test "trace execution" {
  run bash -x "${BATS_TEST_DIRNAME}/../scripts/script.sh"
  # Output includes trace
}
```
