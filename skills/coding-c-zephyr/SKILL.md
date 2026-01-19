---
name: coding-c-zephyr
description: C programming standards for Zephyr RTOS embedded systems
---

# C/Zephyr RTOS Coding Standard

> **Version**: 2.0.0 | **Status**: Active
> **Base Standards**: Zephyr Project Coding Guidelines, MISRA-C:2012, SEI CERT C

This standard establishes coding conventions for C development targeting Zephyr RTOS applications, ensuring consistency, safety, reliability, maintainability, and security.

---

## Navigation

### Rules by Category

| Category | File | Rules |
|----------|------|-------|
| [Naming](rules/naming.md) | C-NAM-* | Descriptive names, subsystem prefixes, static scope |
| [Types](rules/types.md) | C-TYP-* | stdint.h types, Zephyr time types, booleans |
| [Comments](rules/comments.md) | C-CMT-* | C89 block comments, document rationale |
| [Memory](rules/memory.md) | C-MEM-* | Null checks, use-after-free, allocation strategy |
| [Concurrency](rules/concurrency.md) | C-CON-* | Mutex protection, lock ordering, ISR safety |
| [Error Handling](rules/error-handling.md) | C-ERR-* | Check returns, error codes, cleanup |
| [Control Flow](rules/control-flow.md) | C-CTL-* | Braces, switch completeness, nesting |
| [Macros](rules/macros.md) | C-MAC-* | Parenthesization, inline functions |
| [Security](rules/security.md) | C-SEC/INT/ARR/STR-* | Input validation, bounds checking |
| [Testing](rules/testing.md) | C-TST-* | Ztest framework, coverage |
| [Device Tree](rules/device-tree.md) | C-DT-* | Node labels, device readiness |
| [Kconfig](rules/kconfig.md) | C-KCF-* | Configuration options |
| [Power Management](rules/power-management.md) | C-PM-* | PM hooks, sleep constraints |
| [Documentation](rules/documentation.md) | C-DOC-* | Doxygen, magic numbers |
| [Formatting](rules/formatting.md) | - | Indentation, braces, includes |
| [Structure](rules/structure.md) | - | Project layout, directories |
| [CI/CD](rules/cicd.md) | C-CICD-* | Build, test, lint, security scanning |

### Tooling

| Tool | File |
|------|------|
| [clang-format](tooling/clang-format.md) | Code formatting configuration |
| [clang-tidy](tooling/clang-tidy.md) | Static analysis configuration |
| [cppcheck](tooling/cppcheck.md) | Additional static analysis |
| [Pre-commit](tooling/pre-commit.md) | Git hooks setup |

### Reference

| Document | Purpose |
|----------|---------|
| [Quick Reference](reference/quick-reference.md) | Complete rule table with tiers |
| [Code Review Checklist](reference/code-review.md) | Review checklist by category |
| [Full Standard](STANDARD.md) | Complete standard document |

---

## Rule Classification

| Tier | Marker | Enforcement | Response |
|------|--------|-------------|----------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Critical Rules (Always Apply)

These rules are non-negotiable and must be followed in all code.

### C-MEM-001: Check Allocation Return Values :red_circle:

Memory allocation can fail. Always check for NULL.

```c
/* Correct */
struct sensor_data *data = k_malloc(sizeof(*data));
if (data == NULL) {
	LOG_ERR("Failed to allocate sensor data");
	return -ENOMEM;
}

/* Incorrect */
struct sensor_data *data = k_malloc(sizeof(*data));
data->timestamp = k_uptime_get();  /* Crash if allocation failed */
```

### C-CON-001: Protect Shared Data with Synchronization :red_circle:

Unsynchronized access causes race conditions.

```c
/* Correct */
k_mutex_lock(&data_mutex, K_FOREVER);
shared_state.value = new_value;
k_mutex_unlock(&data_mutex);

/* Incorrect */
shared_state.value = new_value;  /* Race condition */
```

### C-CON-006: ISR-Safe Operations Only in Interrupt Context :red_circle:

ISRs cannot block. Use K_NO_WAIT or defer to work queue.

```c
/* Correct */
void uart_isr(const struct device *dev, void *user_data)
{
	k_msgq_put(&rx_queue, &byte, K_NO_WAIT);  /* Non-blocking */
}

/* Incorrect */
void bad_isr(const struct device *dev, void *user_data)
{
	k_mutex_lock(&data_mutex, K_FOREVER);  /* Will hang! */
}
```

### C-ERR-001: Check All Function Return Values :red_circle:

Unchecked errors lead to undefined behavior.

```c
/* Correct */
int ret = gpio_pin_configure(gpio_dev, PIN, GPIO_OUTPUT);
if (ret != 0) {
	LOG_ERR("GPIO configure failed: %d", ret);
	return ret;
}

/* Incorrect */
gpio_pin_configure(gpio_dev, PIN, GPIO_OUTPUT);  /* May fail silently */
```

### C-SEC-001: Validate All External Input :red_circle:

External data can be malformed or malicious.

```c
/* Correct */
if (len == 0 || len > MAX_CMD_LEN) {
	return -EINVAL;
}

/* Incorrect */
process_command(data, len);  /* No validation */
```

### C-DT-002: Validate Device Existence Before Use :red_circle:

Devices may not exist or may fail initialization.

```c
/* Correct */
const struct device *dev = DEVICE_DT_GET(DT_NODELABEL(my_sensor));
if (!device_is_ready(dev)) {
	return -ENODEV;
}

/* Incorrect */
const struct device *dev = DEVICE_DT_GET(DT_NODELABEL(my_sensor));
sensor_sample_fetch(dev);  /* May crash */
```

---

## Quick Rule Lookup

| ID | Rule | Tier |
|----|------|------|
| C-NAM-003 | Use Static for File-Scoped Identifiers | Critical |
| C-TYP-001 | Use stdint.h Types | Critical |
| C-CMT-001 | Use C89 Block Comments Only | Critical |
| C-MEM-001 | Check Allocation Return Values | Critical |
| C-MEM-002 | Never Access Freed Memory | Critical |
| C-MEM-003 | Prevent Double Free | Critical |
| C-CON-001 | Protect Shared Data with Synchronization | Critical |
| C-CON-002 | Avoid Deadlocks with Lock Ordering | Critical |
| C-CON-006 | ISR-Safe Operations Only | Critical |
| C-ERR-001 | Check All Function Return Values | Critical |
| C-ERR-004 | Clean Up Resources on Error Paths | Critical |
| C-CTL-001 | Always Use Braces for Control Structures | Critical |
| C-MAC-001 | Parenthesize Macro Arguments | Critical |
| C-SEC-001 | Validate All External Input | Critical |
| C-SEC-002 | Use Bounded String Functions | Critical |
| C-DT-002 | Validate Device Existence Before Use | Critical |

---

## References

- [Zephyr Project Coding Guidelines](https://docs.zephyrproject.org/latest/contribute/coding_guidelines/)
- [SEI CERT C Coding Standard](https://wiki.sei.cmu.edu/confluence/display/c/SEI+CERT+C+Coding+Standard)
- [MISRA C:2012 Guidelines](https://www.misra.org.uk/misra-c/)
- [Linux Kernel Coding Style](https://www.kernel.org/doc/html/latest/process/coding-style.html)
