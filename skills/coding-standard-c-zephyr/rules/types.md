# Type Usage Rules (C-TYP-*)

Proper type usage ensures portability across different embedded architectures and prevents subtle bugs.

---

## C-TYP-001: Use stdint.h Types :red_circle:

**Tier**: Critical

**Rationale**: Platform-independent sizes prevent bugs when code runs on different architectures (ARM Cortex-M0 vs M4 vs M7).

```c
/* Correct - explicit sizes */
uint8_t  byte_value;
int16_t  signed_offset;
uint32_t milliseconds;
int64_t  timestamp_ns;
size_t   buffer_length;    /* For sizes and indices */
ssize_t  bytes_read;       /* For sizes that can be negative */

/* Incorrect - platform-dependent sizes */
int      count;            /* 16-bit on some platforms, 32-bit on others */
long     timestamp;        /* 32-bit or 64-bit depending on platform */
unsigned size;             /* Ambiguous size */
```

---

## C-TYP-002: Use Zephyr Time Types :yellow_circle:

**Tier**: Required

**Rationale**: Zephyr provides type-safe duration and timing APIs that prevent unit confusion.

```c
/* Correct - Zephyr duration macros */
k_timeout_t timeout = K_MSEC(100);
k_sleep(K_SECONDS(1));
int64_t uptime = k_uptime_get();

/* Correct - explicit units in variable names when using raw values */
uint32_t interval_ms = 500;
uint64_t deadline_ticks = k_uptime_ticks() + k_ms_to_ticks_ceil32(100);

/* Incorrect - ambiguous time units */
int timeout = 100;         /* Milliseconds? Microseconds? Ticks? */
sleep(1);                  /* What unit? Not Zephyr API */
```

---

## C-TYP-003: Use Boolean Type Correctly :yellow_circle:

**Tier**: Required

**Rationale**: `bool` communicates intent; integer comparisons to 1/0 are error-prone.

```c
#include <stdbool.h>

/* Correct */
bool is_initialized = false;
bool connection_active = true;

if (is_initialized) {
	/* ... */
}

bool check_ready(void)
{
	return (state == STATE_READY);
}

/* Incorrect */
int is_initialized = 0;    /* Use bool, not int */
if (is_initialized == 1) { /* Never compare bool to 1 */
	/* ... */
}
if (!is_initialized == true) { /* Convoluted */
	/* ... */
}
```
