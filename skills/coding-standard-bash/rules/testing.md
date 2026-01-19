# Testing Rules (SH-TST-*)

Testing shell scripts catches regressions and documents expected behavior. Shell scripts are prone to subtle bugs that tests can catch.

## Testing Strategy

- Use Bats (Bash Automated Testing System)
- Test edge cases involving special characters
- Test error conditions
- Use setup/teardown for isolation

---

## Testing Framework

Use [Bats](https://github.com/bats-core/bats-core) (Bash Automated Testing System) for testing shell scripts.

### Installation

```bash
# Using npm
npm install -g bats

# Using Homebrew (macOS)
brew install bats-core

# From source
git clone https://github.com/bats-core/bats-core.git
cd bats-core
./install.sh /usr/local
```

---

## SH-TST-001: Write Tests for Shell Scripts :yellow_circle:

**Tier**: Required

**Rationale**: Tests catch regressions and document expected behavior.

```bash
# tests/utils_test.bats

#!/usr/bin/env bats

setup() {
  # Load the script being tested
  source "${BATS_TEST_DIRNAME}/../scripts/utils.sh"

  # Create temp directory for test files
  TEST_TEMP_DIR="$(mktemp -d)"
}

teardown() {
  # Clean up
  rm -rf "${TEST_TEMP_DIR}"
}

@test "validate_email accepts valid email" {
  run validate_email "user@example.com"
  [ "$status" -eq 0 ]
}

@test "validate_email rejects invalid email" {
  run validate_email "not-an-email"
  [ "$status" -eq 1 ]
  [[ "$output" =~ "Invalid email" ]]
}

@test "process_file handles spaces in filename" {
  local test_file="${TEST_TEMP_DIR}/file with spaces.txt"
  echo "test content" > "${test_file}"

  run process_file "${test_file}"
  [ "$status" -eq 0 ]
}

@test "process_file fails on missing file" {
  run process_file "/nonexistent/file.txt"
  [ "$status" -eq 1 ]
}
```

---

## SH-TST-002: Test Edge Cases :yellow_circle:

**Tier**: Required

**Rationale**: Shell scripts are particularly vulnerable to edge cases involving special characters, empty values, and unusual input.

```bash
@test "handles empty string input" {
  run process_input ""
  [ "$status" -eq 1 ]
  [[ "$output" =~ "Input cannot be empty" ]]
}

@test "handles input with spaces" {
  run process_input "hello world"
  [ "$status" -eq 0 ]
}

@test "handles input with special characters" {
  run process_input 'test$var`cmd`$(injection)'
  [ "$status" -eq 0 ]
  # Verify no command execution occurred
  [[ ! "$output" =~ "injection" ]]
}

@test "handles input with newlines" {
  run process_input $'line1\nline2'
  [ "$status" -eq 0 ]
}

@test "handles unicode input" {
  run process_input "Hello 世界"
  [ "$status" -eq 0 ]
}

@test "handles extremely long input" {
  local long_input
  long_input="$(printf 'a%.0s' {1..10000})"
  run process_input "${long_input}"
  [ "$status" -eq 0 ]
}

@test "handles filename with glob characters" {
  local test_file="${TEST_TEMP_DIR}/file*.txt"
  touch "${test_file}"

  run process_file "${test_file}"
  [ "$status" -eq 0 ]
}

@test "handles dash-prefixed arguments" {
  run process_input "-n"
  [ "$status" -eq 0 ]
  # Should treat as data, not flag
}
```

---

## Test File Structure

```
project/
  scripts/
    deploy.sh
    utils.sh
  tests/
    deploy_test.bats
    utils_test.bats
    test_helper.bash
    fixtures/
      sample_config.json
      expected_output.txt
```

### Test Helper

```bash
# tests/test_helper.bash

# Common setup for all tests
setup_common() {
  TEST_TEMP_DIR="$(mktemp -d)"
  export TEST_TEMP_DIR

  # Load scripts
  source "${BATS_TEST_DIRNAME}/../scripts/utils.sh"
}

# Common teardown
teardown_common() {
  rm -rf "${TEST_TEMP_DIR}"
}

# Assertion helpers
assert_file_exists() {
  local file="$1"
  if [[ ! -f "${file}" ]]; then
    echo "Expected file to exist: ${file}" >&2
    return 1
  fi
}

assert_output_contains() {
  local expected="$1"
  if [[ ! "${output}" =~ ${expected} ]]; then
    echo "Expected output to contain: ${expected}" >&2
    echo "Actual output: ${output}" >&2
    return 1
  fi
}
```

### Using Test Helper

```bash
# tests/deploy_test.bats

#!/usr/bin/env bats

load test_helper

setup() {
  setup_common
}

teardown() {
  teardown_common
}

@test "deploy creates output directory" {
  run deploy --output "${TEST_TEMP_DIR}/deploy"
  [ "$status" -eq 0 ]
  assert_file_exists "${TEST_TEMP_DIR}/deploy/manifest.json"
}
```

---

## Additional Testing Patterns

### Testing Exit Codes

```bash
@test "returns 0 on success" {
  run my_command "valid_input"
  [ "$status" -eq 0 ]
}

@test "returns 1 on general error" {
  run my_command "invalid_input"
  [ "$status" -eq 1 ]
}

@test "returns 2 on invalid arguments" {
  run my_command
  [ "$status" -eq 2 ]
}
```

### Testing stdout and stderr

```bash
@test "success message goes to stdout" {
  run my_command "input"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "Success" ]]
}

@test "error message goes to stderr" {
  run bash -c 'my_command "bad" 2>&1'
  [[ "$output" =~ "Error" ]]
}
```

### Mocking Commands

```bash
setup() {
  # Create mock commands directory
  MOCK_DIR="${BATS_TEST_DIRNAME}/mocks"
  mkdir -p "${MOCK_DIR}"

  # Add to PATH
  export PATH="${MOCK_DIR}:${PATH}"
}

@test "handles curl failure" {
  # Create mock curl that fails
  cat > "${MOCK_DIR}/curl" << 'EOF'
#!/usr/bin/env bash
exit 1
EOF
  chmod +x "${MOCK_DIR}/curl"

  run download_file "http://example.com/file"
  [ "$status" -eq 1 ]
  [[ "$output" =~ "Download failed" ]]
}
```

### Testing with Fixtures

```bash
@test "parses config file correctly" {
  run parse_config "${BATS_TEST_DIRNAME}/fixtures/sample_config.json"
  [ "$status" -eq 0 ]

  # Compare with expected output
  local expected
  expected="$(cat "${BATS_TEST_DIRNAME}/fixtures/expected_output.txt")"
  [ "$output" = "$expected" ]
}
```

### Running Tests

```bash
# Run all tests
bats tests/

# Run specific test file
bats tests/utils_test.bats

# Run with verbose output
bats --verbose-run tests/

# Run with timing
bats --timing tests/

# Output TAP format (for CI)
bats --tap tests/
```
