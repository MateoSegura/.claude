# Comment Rules (C-CMT-*)

Comments explain intent and rationale. They should add value beyond what the code itself communicates.

---

## C-CMT-001: Use C89 Block Comments Only :red_circle:

**Tier**: Critical

**Rationale**: Zephyr mandates C89-style comments for consistency with the kernel and broader compatibility.

```c
/* Correct - C89 style block comments */

/* Single line comment */

/*
 * Multi-line comment explaining
 * complex logic or algorithm
 */

/**
 * @brief Doxygen documentation comment
 * @param value Input value to process
 * @return Processed result
 */

/* Incorrect - C99/C++ style comments */

// Single line comment - NOT ALLOWED

int x = 5; // Trailing comment - NOT ALLOWED
```

---

## C-CMT-002: Document Non-Obvious Code :yellow_circle:

**Tier**: Required

**Rationale**: Comments explain *why*, not *what*. Code should be self-documenting for the *what*.

```c
/* Correct - explains rationale */

/*
 * Use a timeout of 2x the expected response time to account
 * for bus arbitration delays in multi-master configurations.
 * See datasheet section 4.3.2 for timing requirements.
 */
#define I2C_TIMEOUT_MS (EXPECTED_RESPONSE_MS * 2)

/*
 * Clear pending interrupt before enabling to prevent
 * spurious interrupt on first enable (hardware errata #42).
 */
clear_pending_irq(IRQ_NUM);
enable_irq(IRQ_NUM);

/* Incorrect - states the obvious */

/* Increment counter */
counter++;

/* Check if buffer is full */
if (buffer_is_full()) {
```
