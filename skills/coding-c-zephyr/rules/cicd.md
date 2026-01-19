---
description: Zephyr CI/CD rules - build, test, lint, security scanning for Zephyr RTOS projects
---

# Zephyr CI/CD Rules

Zephyr RTOS-specific CI/CD configurations and tooling. For universal CI/CD principles, see [devops-standard/rules/cicd-principles.md](../../devops-standard/rules/cicd-principles.md).

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## 1. Build Stage

### C-CICD-001: Use `west build` for Compilation :red_circle:

**Rationale**: West is Zephyr's meta-tool that ensures consistent builds across platforms.

```bash
# Basic build
west build -b <board> <application>

# Build with pristine (clean)
west build -b nrf52840dk_nrf52840 samples/hello_world -p

# Build with specific configuration
west build -b nrf52840dk_nrf52840 -- -DCONF_FILE=prj_debug.conf

# Build multiple boards
for board in nrf52840dk_nrf52840 stm32f4_disco esp32; do
    west build -b $board app -p
done
```

**Required flags for CI**:
- `-p` or `--pristine` - Clean build (ensures reproducibility)
- `--` followed by CMake options for configuration

### C-CICD-002: Pin Zephyr SDK Version :red_circle:

**Rationale**: Reproducible builds require consistent toolchain versions.

```bash
# Specify SDK version in CI
ZEPHYR_SDK_VERSION=0.16.4

# Download and install specific version
wget https://github.com/zephyrproject-rtos/sdk-ng/releases/download/v${ZEPHYR_SDK_VERSION}/zephyr-sdk-${ZEPHYR_SDK_VERSION}_linux-x86_64.tar.xz
tar xf zephyr-sdk-${ZEPHYR_SDK_VERSION}_linux-x86_64.tar.xz
./zephyr-sdk-${ZEPHYR_SDK_VERSION}/setup.sh
```

### C-CICD-003: Use `west manifest` for Dependencies :red_circle:

**Rationale**: West manifest ensures reproducible dependency versions.

```bash
# Initialize workspace
west init -m https://github.com/org/project --mr v1.0.0

# Update dependencies
west update

# Freeze manifest for reproducibility
west manifest --freeze > west.yml.lock
```

---

## 2. Test Stage

### C-CICD-004: Use Twister for Testing :red_circle:

**Rationale**: Twister is Zephyr's test framework that runs tests across multiple platforms.

```bash
# Run all tests
./scripts/twister -T tests/

# Run tests for specific platform
./scripts/twister -p native_posix -T tests/

# Run with coverage
./scripts/twister -p native_posix --coverage -T tests/

# Run specific test suite
./scripts/twister -T tests/kernel/threads -p native_posix

# Generate JUnit XML report
./scripts/twister -T tests/ --report-junit report.xml
```

**Twister options for CI**:
- `--report-junit` - JUnit XML output for CI parsing
- `--coverage` - Enable code coverage collection
- `-p native_posix` - Run on host for unit tests
- `-j <n>` - Parallel jobs
- `--retry-failed 2` - Retry failed tests

### C-CICD-005: Use Native POSIX for Unit Tests :yellow_circle:

**Rationale**: native_posix enables fast unit testing without hardware.

```bash
# Build for native_posix
west build -b native_posix tests/unit/my_test -p

# Run
./build/zephyr/zephyr.exe

# With Twister
./scripts/twister -p native_posix -T tests/unit/
```

### C-CICD-006: Test on Multiple Boards :yellow_circle:

**Rationale**: Testing on multiple boards catches platform-specific issues.

```bash
# Define test platforms in testcase.yaml
tests:
  kernel.threads:
    platform_allow:
      - native_posix
      - nrf52840dk_nrf52840
      - stm32f4_disco
    integration_platforms:
      - native_posix
```

---

## 3. Static Analysis Stage

### C-CICD-007: Use clang-tidy :red_circle:

**Rationale**: clang-tidy catches bugs, style violations, and modernization opportunities.

```bash
# Generate compile_commands.json
west build -b native_posix app -p -- -DCMAKE_EXPORT_COMPILE_COMMANDS=ON

# Run clang-tidy
clang-tidy -p build/compile_commands.json src/*.c

# With Zephyr-specific checks
clang-tidy -p build/compile_commands.json \
    --checks='-*,bugprone-*,cert-*,clang-analyzer-*,misc-*,performance-*' \
    src/*.c
```

**Recommended `.clang-tidy`**:

```yaml
Checks: >
  bugprone-*,
  cert-*,
  clang-analyzer-*,
  misc-*,
  performance-*,
  readability-*,
  -readability-magic-numbers,
  -bugprone-easily-swappable-parameters

WarningsAsErrors: '*'

HeaderFilterRegex: 'src/.*'

CheckOptions:
  - key: readability-identifier-naming.FunctionCase
    value: lower_case
  - key: readability-identifier-naming.VariableCase
    value: lower_case
  - key: readability-identifier-naming.MacroDefinitionCase
    value: UPPER_CASE
```

### C-CICD-008: Use cppcheck :yellow_circle:

**Rationale**: cppcheck is lightweight and catches errors clang-tidy may miss.

```bash
# Run cppcheck
cppcheck --enable=warning,style,performance,portability \
    --error-exitcode=1 \
    --suppress=missingInclude \
    --inline-suppr \
    src/

# With MISRA checks (requires premium)
cppcheck --enable=all --addon=misra.py src/
```

### C-CICD-009: Use clang-format for Formatting :red_circle:

**Rationale**: Consistent formatting eliminates style debates and reduces diff noise.

```bash
# Check formatting (CI)
find src include -name '*.c' -o -name '*.h' | xargs clang-format --dry-run --Werror

# Apply formatting
find src include -name '*.c' -o -name '*.h' | xargs clang-format -i
```

**Zephyr-compatible `.clang-format`**:

```yaml
BasedOnStyle: LLVM
IndentWidth: 8
UseTab: Always
BreakBeforeBraces: Linux
AllowShortIfStatementsOnASingleLine: false
IndentCaseLabels: false
ColumnLimit: 100
```

### C-CICD-010: Enforce C89 Comments :yellow_circle:

**Rationale**: Zephyr style requires C89 block comments only.

```bash
# Check for C++ style comments
grep -rn '//' src/ include/ && echo "ERROR: Use /* */ comments only" && exit 1
```

---

## 4. Security Stage

### C-CICD-011: Use scan-build for Static Analysis :red_circle:

**Rationale**: scan-build uses Clang Static Analyzer for deep bug detection.

```bash
# Run scan-build with west
scan-build west build -b native_posix app -p

# Generate HTML report
scan-build -o /tmp/scan-results west build -b native_posix app -p
```

### C-CICD-012: Use Coverity or CodeQL :yellow_circle:

**Rationale**: Enterprise-grade static analysis catches security vulnerabilities.

```yaml
# GitHub Actions - CodeQL
- name: Initialize CodeQL
  uses: github/codeql-action/init@v3
  with:
    languages: cpp

- name: Build
  run: west build -b native_posix app -p

- name: Perform CodeQL Analysis
  uses: github/codeql-action/analyze@v3
```

### C-CICD-013: Stack Usage Analysis :yellow_circle:

**Rationale**: Embedded systems have limited stack; analysis prevents overflows.

```bash
# Enable stack usage in CMake
west build -b nrf52840dk_nrf52840 app -- -DCONFIG_STACK_USAGE=y

# Analyze .su files
find build -name '*.su' -exec cat {} \; | sort -t: -k3 -n -r | head -20
```

---

## 5. Complete GitHub Actions Pipeline

```yaml
name: Zephyr CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  ZEPHYR_SDK_VERSION: 0.16.4

jobs:
  build:
    runs-on: ubuntu-latest
    container:
      image: ghcr.io/zephyrproject-rtos/ci:v0.26.4

    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: West init and update
        run: |
          west init -l .
          west update

      - name: Build for nRF52840
        run: |
          west build -b nrf52840dk_nrf52840 app -p

      - name: Build for STM32
        run: |
          west build -b stm32f4_disco app -p

  test:
    runs-on: ubuntu-latest
    container:
      image: ghcr.io/zephyrproject-rtos/ci:v0.26.4

    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: West init and update
        run: |
          west init -l .
          west update

      - name: Run Twister tests
        run: |
          ./scripts/twister -p native_posix -T tests/ \
            --report-junit report.xml \
            --coverage

      - name: Upload test results
        uses: actions/upload-artifact@v4
        if: always()
        with:
          name: test-results
          path: twister-out/

      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          directory: twister-out/coverage/

  lint:
    runs-on: ubuntu-latest
    container:
      image: ghcr.io/zephyrproject-rtos/ci:v0.26.4

    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: West init and update
        run: |
          west init -l .
          west update

      - name: Check formatting
        run: |
          find src include -name '*.c' -o -name '*.h' | \
            xargs clang-format --dry-run --Werror

      - name: Generate compile_commands.json
        run: |
          west build -b native_posix app -p -- \
            -DCMAKE_EXPORT_COMPILE_COMMANDS=ON

      - name: Run clang-tidy
        run: |
          clang-tidy -p build/compile_commands.json \
            --checks='-*,bugprone-*,cert-*,clang-analyzer-*' \
            src/*.c

      - name: Run cppcheck
        run: |
          cppcheck --enable=warning,style,performance \
            --error-exitcode=1 \
            --suppress=missingInclude \
            src/

  security:
    runs-on: ubuntu-latest
    container:
      image: ghcr.io/zephyrproject-rtos/ci:v0.26.4

    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: West init and update
        run: |
          west init -l .
          west update

      - name: Run scan-build
        run: |
          scan-build -o scan-results \
            west build -b native_posix app -p

      - name: Upload scan-build results
        uses: actions/upload-artifact@v4
        if: always()
        with:
          name: scan-build-results
          path: scan-results/
```

---

## 6. Makefile Targets

```makefile
.PHONY: build test lint security all clean

BOARD ?= nrf52840dk_nrf52840
APP ?= app

all: lint test build

build:
	west build -b $(BOARD) $(APP) -p

test:
	./scripts/twister -p native_posix -T tests/ \
		--report-junit report.xml

lint:
	clang-format --dry-run --Werror src/*.c include/*.h
	west build -b native_posix $(APP) -p -- -DCMAKE_EXPORT_COMPILE_COMMANDS=ON
	clang-tidy -p build/compile_commands.json src/*.c
	cppcheck --enable=warning,style --error-exitcode=1 src/

security:
	scan-build west build -b native_posix $(APP) -p

coverage:
	./scripts/twister -p native_posix -T tests/ --coverage
	genhtml twister-out/coverage/*.info -o coverage-report

clean:
	rm -rf build twister-out scan-results coverage-report
```

---

## Quick Reference

| ID | Rule | Tier |
|----|------|------|
| C-CICD-001 | Use `west build` for compilation | Critical |
| C-CICD-002 | Pin Zephyr SDK version | Critical |
| C-CICD-003 | Use `west manifest` for dependencies | Critical |
| C-CICD-004 | Use Twister for testing | Critical |
| C-CICD-005 | Use native POSIX for unit tests | Required |
| C-CICD-006 | Test on multiple boards | Required |
| C-CICD-007 | Use clang-tidy | Critical |
| C-CICD-008 | Use cppcheck | Required |
| C-CICD-009 | Use clang-format for formatting | Critical |
| C-CICD-010 | Enforce C89 comments | Required |
| C-CICD-011 | Use scan-build for static analysis | Critical |
| C-CICD-012 | Use Coverity or CodeQL | Required |
| C-CICD-013 | Stack usage analysis | Required |

---

## References

- [Zephyr Build System](https://docs.zephyrproject.org/latest/build/index.html)
- [Zephyr Twister](https://docs.zephyrproject.org/latest/develop/test/twister.html)
- [Zephyr CI/CD](https://docs.zephyrproject.org/latest/develop/ci/index.html)
- [clang-tidy](https://clang.llvm.org/extra/clang-tidy/)
