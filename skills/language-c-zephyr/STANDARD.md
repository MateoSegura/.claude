# C/Zephyr RTOS Coding Standard

> **Version**: 2.0.0
> **Status**: Active
> **Base Standards**: Zephyr Project Coding Guidelines, MISRA-C:2012, SEI CERT C
> **Last Updated**: 2026-01-04

---

## 1. Purpose and Scope

### 1.1 Purpose

This standard establishes coding conventions for C development targeting Zephyr RTOS applications. It ensures:

- **Consistency**: Uniform code style aligned with the Zephyr ecosystem
- **Safety**: Prevention of common embedded systems vulnerabilities
- **Reliability**: Deterministic behavior in real-time environments
- **Maintainability**: Code that teams can understand and extend
- **Security**: Protection against memory corruption and concurrency bugs

### 1.2 Scope

This standard applies to:

- [x] Application-level C code running on Zephyr RTOS
- [x] Device drivers developed outside the Zephyr kernel tree
- [x] Test code (with documented exceptions)
- [x] Board support packages and custom configurations

**Explicitly Permitted** (unlike kernel code):
- Dynamic memory allocation via `k_malloc()`, `k_heap`, and memory pools
- Networking stacks and socket operations
- Multithreading with Zephyr primitives

**Out of Scope**:
- Zephyr kernel contributions (follow upstream guidelines)
- Assembly code
- Third-party libraries (document deviations)

### 1.3 Audience

- Embedded software engineers writing Zephyr applications
- Code reviewers evaluating changes
- Tech leads establishing project baselines
- QA engineers writing static analysis rules

### 1.4 Relationship to Industry Standards

| Standard | Relationship |
|----------|--------------|
| Zephyr Project Coding Guidelines | **Base** - Style and formatting rules |
| MISRA-C:2012 | **Base** - Safety-critical patterns |
| SEI CERT C | **Supplementary** - Security rules for memory, concurrency, errors |
| Linux Kernel Coding Style | **Reference** - Formatting conventions |

---

## 2. Rule Classification

Rules are classified by enforcement level:

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails; must fix immediately |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged; document exceptions |

Each rule includes:
- **Rule ID**: Unique identifier (`C-XXX-NNN` format)
- **Tier**: Enforcement level
- **Rationale**: Technical justification
- **Correct Example**: Code that follows the rule
- **Incorrect Example**: Code that violates the rule

---

## 3. Project Structure

### 3.1 Directory Layout

```
project-root/
+-- CMakeLists.txt
+-- prj.conf
+-- app.overlay            /* Device tree overlay */
+-- Kconfig                /* Application Kconfig */
+-- src/
|   +-- main.c
|   +-- app_config.h       /* Application configuration */
|   +-- subsystem/
|       +-- subsystem.c
|       +-- subsystem.h
|       +-- subsystem_internal.h
+-- include/
|   +-- public_api.h       /* Public interfaces */
+-- boards/
|   +-- <board_name>.conf
|   +-- <board_name>.overlay
+-- tests/
|   +-- unit/
|   +-- integration/
+-- doc/
```

### 3.2 File Naming

| Type | Convention | Example |
|------|------------|---------|
| Source files | lowercase_with_underscores | `sensor_driver.c` |
| Header files | lowercase_with_underscores | `sensor_driver.h` |
| Internal headers | suffix `_internal` | `sensor_driver_internal.h` |
| Test files | suffix `_test` | `sensor_driver_test.c` |
| Kconfig fragments | descriptive `.conf` | `debug_logging.conf` |

### 3.3 Header File Organization

```c
/* SPDX-License-Identifier: Apache-2.0 */

/**
 * @file sensor_driver.h
 * @brief Public API for sensor driver
 */

#ifndef SENSOR_DRIVER_H
#define SENSOR_DRIVER_H

#ifdef __cplusplus
extern "C" {
#endif

#include <zephyr/kernel.h>
#include <stdint.h>
#include <stdbool.h>

/* Public type definitions */
/* Public constants */
/* Public function declarations */

#ifdef __cplusplus
}
#endif

#endif /* SENSOR_DRIVER_H */
```

---

## 4. Formatting Rules

### 4.1 Indentation and Spacing

| Aspect | Rule |
|--------|------|
| Indentation | Tabs (8-space width display) |
| Line length | 100 characters maximum (80 preferred) |
| Trailing whitespace | Forbidden |
| Final newline | Required |
| Spaces after keywords | Required (`if (`, `for (`, `while (`) |
| Spaces around operators | Required (`a + b`, not `a+b`) |
| No space after function name | `func(arg)`, not `func (arg)` |

### 4.2 Brace Style

Zephyr follows K&R style with mandatory braces:

```c
/* Function definitions: brace on new line */
int sensor_read(uint8_t channel)
{
	/* body */
}

/* Control structures: brace on same line */
if (condition) {
	do_something();
} else {
	do_other();
}

/* Single statements still require braces */
if (error) {
	return -EINVAL;
}
```

### 4.3 Include Ordering

Order includes in these groups, separated by blank lines:

1. Corresponding header (for `.c` files)
2. Zephyr headers (`<zephyr/*.h>`)
3. Standard library headers (`<stdint.h>`, `<string.h>`)
4. Third-party library headers
5. Project headers (`"local_header.h"`)

```c
#include "sensor_driver.h"

#include <zephyr/kernel.h>
#include <zephyr/device.h>
#include <zephyr/drivers/gpio.h>

#include <stdint.h>
#include <string.h>

#include "app_config.h"
#include "utils.h"
```

---

## 5. Naming Conventions

### 5.1 General Principles

- Names reveal intent and purpose
- Length proportional to scope (longer for wider visibility)
- No Hungarian notation
- Abbreviations only for well-known terms (e.g., `cfg`, `ctx`, `buf`)

### 5.2 Specific Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Functions | `lowercase_with_underscores` | `sensor_read_value()` |
| Local variables | `lowercase_with_underscores` | `sample_count` |
| Global variables | `lowercase_with_underscores` | `system_state` |
| Constants | `UPPERCASE_WITH_UNDERSCORES` | `MAX_BUFFER_SIZE` |
| Macros | `UPPERCASE_WITH_UNDERSCORES` | `ARRAY_SIZE(x)` |
| Enum values | `UPPERCASE_WITH_UNDERSCORES` | `SENSOR_STATE_IDLE` |
| Struct types | `struct lowercase_name` | `struct sensor_config` |
| Typedef (sparingly) | `lowercase_t` suffix | `sensor_callback_t` |

### 5.3 Naming Rules

#### C-NAM-001: Use Descriptive Function Names :yellow_circle:

**Rationale**: Functions should clearly communicate their purpose. Verbs indicate action; nouns indicate the subject.

```c
/* Correct */
int sensor_read_temperature(const struct device *dev, int32_t *value);
void buffer_queue_flush(struct buffer_queue *queue);
bool connection_is_established(const struct connection *conn);

/* Incorrect */
int read(void *d, int *v);           /* Too generic, unclear parameters */
void do_stuff(struct buffer_queue *q); /* Meaningless name */
int temp(void);                       /* Ambiguous abbreviation */
```

#### C-NAM-002: Prefix Subsystem Functions :yellow_circle:

**Rationale**: Namespacing prevents symbol collisions and clarifies code ownership.

```c
/* Correct - clear subsystem prefix */
int mqtt_client_connect(struct mqtt_client *client);
int mqtt_client_publish(struct mqtt_client *client, const char *topic);
void mqtt_client_disconnect(struct mqtt_client *client);

/* Incorrect - no namespace, collision risk */
int connect(struct mqtt_client *client);  /* Conflicts with POSIX */
int publish(struct mqtt_client *client, const char *topic);
```

#### C-NAM-003: Use Static for File-Scoped Identifiers :red_circle:

**Rationale**: `static` limits visibility, prevents unintended linkage, and enables compiler optimizations.

```c
/* Correct - internal functions are static */
static int validate_input(const uint8_t *data, size_t len);
static void cleanup_resources(struct context *ctx);

/* File-scope variables must be static */
static uint32_t packet_counter;
static struct k_mutex data_lock;

/* Incorrect - global linkage for internal items */
int validate_input(const uint8_t *data, size_t len);  /* Visible externally */
uint32_t packet_counter;                              /* Namespace pollution */
```

---

## 6. Type Usage

### 6.1 Use Sized Integer Types

#### C-TYP-001: Use stdint.h Types :red_circle:

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

#### C-TYP-002: Use Zephyr Time Types :yellow_circle:

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

#### C-TYP-003: Use Boolean Type Correctly :yellow_circle:

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

---

## 7. Comments

### 7.1 Comment Style

#### C-CMT-001: Use C89 Block Comments Only :red_circle:

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

#### C-CMT-002: Document Non-Obvious Code :yellow_circle:

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

---

## 8. Memory Management

Application code MAY use dynamic allocation with appropriate safeguards. All rules derive from SEI CERT C MEM30-C through MEM36-C.

### 8.1 Memory Allocation Rules

#### C-MEM-001: Check Allocation Return Values :red_circle:

**Rationale**: Memory allocation can fail. Dereferencing NULL causes crashes or exploitable vulnerabilities. Based on SEI CERT MEM32-C.

```c
/* Correct - check and handle failure */
struct sensor_data *data = k_malloc(sizeof(*data));
if (data == NULL) {
	LOG_ERR("Failed to allocate sensor data");
	return -ENOMEM;
}

/* Using memory pools */
void *block;
int ret = k_mem_pool_alloc(&my_pool, &block, size, K_NO_WAIT);
if (ret != 0) {
	LOG_WRN("Pool allocation failed: %d", ret);
	return ret;
}

/* Incorrect - no null check */
struct sensor_data *data = k_malloc(sizeof(*data));
data->timestamp = k_uptime_get();  /* Crash if allocation failed */
```

#### C-MEM-002: Never Access Freed Memory :red_circle:

**Rationale**: Use-after-free is undefined behavior exploitable for code execution. Based on SEI CERT MEM30-C.

```c
/* Correct - nullify pointer after free */
void cleanup_context(struct context **ctx_ptr)
{
	if (ctx_ptr == NULL || *ctx_ptr == NULL) {
		return;
	}

	struct context *ctx = *ctx_ptr;

	/* Free child resources first */
	if (ctx->buffer != NULL) {
		k_free(ctx->buffer);
		ctx->buffer = NULL;
	}

	k_free(ctx);
	*ctx_ptr = NULL;  /* Prevent use-after-free */
}

/* Correct - save next pointer before freeing in list traversal */
void free_list(struct node *head)
{
	struct node *current = head;
	struct node *next;

	while (current != NULL) {
		next = current->next;  /* Save before free */
		k_free(current);
		current = next;
	}
}

/* Incorrect - use after free */
void bad_cleanup(struct context *ctx)
{
	k_free(ctx);
	LOG_DBG("Freed context %p", ctx);  /* Use after free */
	ctx->state = STATE_FREED;          /* Write after free - exploitable */
}

/* Incorrect - accessing freed memory in loop */
void bad_free_list(struct node *head)
{
	for (struct node *p = head; p != NULL; p = p->next) {
		k_free(p);  /* p->next accessed after free! */
	}
}
```

#### C-MEM-003: Prevent Double Free :red_circle:

**Rationale**: Double-free corrupts heap metadata and enables exploitation. Based on SEI CERT MEM31-C.

```c
/* Correct - nullify and check before free */
void safe_free(void **ptr)
{
	if (ptr != NULL && *ptr != NULL) {
		k_free(*ptr);
		*ptr = NULL;
	}
}

/* Correct - clear pointer immediately */
struct resource *res = k_malloc(sizeof(*res));
/* ... use res ... */
k_free(res);
res = NULL;  /* Prevents accidental double-free */

/* Incorrect - no protection against double-free */
void process_and_cleanup(struct data *d)
{
	if (should_cleanup()) {
		k_free(d);
	}
	/* ... more code ... */
	if (error_occurred()) {
		k_free(d);  /* Possible double-free */
	}
}
```

#### C-MEM-004: Free Memory in Consistent Order :yellow_circle:

**Rationale**: Free child resources before parent to prevent dangling pointers. Based on SEI CERT MEM31-C patterns.

```c
/* Correct - free children before parent */
void destroy_connection(struct connection *conn)
{
	if (conn == NULL) {
		return;
	}

	/* Free dynamically allocated members first */
	if (conn->rx_buffer != NULL) {
		k_free(conn->rx_buffer);
	}
	if (conn->tx_buffer != NULL) {
		k_free(conn->tx_buffer);
	}
	if (conn->hostname != NULL) {
		k_free(conn->hostname);
	}

	/* Finally free the struct itself */
	k_free(conn);
}

/* Incorrect - freeing parent with live child pointers */
void bad_destroy(struct connection *conn)
{
	k_free(conn);  /* Children now dangling */
	k_free(conn->rx_buffer);  /* Use after free */
}
```

#### C-MEM-005: Use Appropriate Allocation Strategy :yellow_circle:

**Rationale**: Choose allocation method based on lifetime, size predictability, and real-time requirements. Based on SEI CERT MEM35-C.

```c
/* Correct - static allocation for fixed-size, long-lived data */
static uint8_t uart_rx_buffer[CONFIG_UART_RX_BUF_SIZE];
static struct k_msgq sensor_queue;
K_MSGQ_DEFINE(cmd_queue, sizeof(struct command), 10, 4);

/* Correct - stack allocation for small, short-lived data */
void process_message(const uint8_t *msg, size_t len)
{
	uint8_t local_buf[64];  /* Known small size, function scope */
	/* ... */
}

/* Correct - memory pool for fixed-size allocations */
K_HEAP_DEFINE(packet_heap, 4096);

struct packet *alloc_packet(void)
{
	return k_heap_alloc(&packet_heap, sizeof(struct packet), K_NO_WAIT);
}

/* Correct - k_malloc for variable-size, application-lifetime data */
struct config *load_config(size_t data_size)
{
	struct config *cfg = k_malloc(sizeof(*cfg) + data_size);
	if (cfg == NULL) {
		return NULL;
	}
	/* ... */
	return cfg;
}

/* Incorrect - heap allocation for small fixed buffers */
void bad_process(void)
{
	uint8_t *buf = k_malloc(32);  /* Stack would be simpler */
	/* ... */
	k_free(buf);  /* Must remember to free */
}
```

#### C-MEM-006: Specify Sufficient Allocation Size :red_circle:

**Rationale**: Undersized allocations cause buffer overflows. Use `sizeof` on the dereferenced pointer, not the type. Based on SEI CERT MEM35-C.

```c
/* Correct - sizeof(*pointer) pattern */
struct sensor_reading *reading = k_malloc(sizeof(*reading));

/* Correct - calculating array allocation */
size_t count = 100;
int32_t *samples = k_malloc(count * sizeof(*samples));

/* Correct - flexible array member */
struct message {
	uint16_t length;
	uint8_t data[];
};

struct message *msg = k_malloc(sizeof(*msg) + payload_len);
if (msg != NULL) {
	msg->length = payload_len;
}

/* Incorrect - sizeof(type) can diverge from pointer type */
struct sensor_reading *reading = k_malloc(sizeof(struct sensor_reading_v2));
/* If pointer type changes, allocation size doesn't update */

/* Incorrect - forgetting to multiply by count */
int32_t *samples = k_malloc(sizeof(*samples));  /* Only 4 bytes, not array */
```

---

## 9. Concurrency

Zephyr applications commonly use threads, ISRs, and various synchronization primitives. Rules derive from SEI CERT C CON30-C through CON43-C.

### 9.1 Thread Safety Rules

#### C-CON-001: Protect Shared Data with Synchronization :red_circle:

**Rationale**: Unsynchronized access to shared data causes race conditions with unpredictable results. Based on SEI CERT CON32-C.

```c
/* Correct - mutex protection for shared state */
static struct k_mutex data_mutex;
static struct system_state shared_state;

int update_state(enum state_value new_state)
{
	int ret;

	ret = k_mutex_lock(&data_mutex, K_MSEC(100));
	if (ret != 0) {
		LOG_ERR("Failed to acquire mutex: %d", ret);
		return ret;
	}

	shared_state.current = new_state;
	shared_state.timestamp = k_uptime_get();

	k_mutex_unlock(&data_mutex);
	return 0;
}

/* Correct - atomic operations for simple counters */
static atomic_t packet_count = ATOMIC_INIT(0);

void on_packet_received(void)
{
	atomic_inc(&packet_count);
}

/* Incorrect - unprotected shared access */
static uint32_t counter;  /* Shared between threads */

void thread_a(void)
{
	counter++;  /* Race condition */
}

void thread_b(void)
{
	if (counter > 0) {  /* May read partial update */
		counter--;
	}
}
```

#### C-CON-002: Avoid Deadlocks with Lock Ordering :red_circle:

**Rationale**: Inconsistent lock ordering between threads causes deadlocks. Based on SEI CERT CON35-C.

```c
/* Correct - consistent lock ordering (alphabetical by convention) */

/*
 * Lock ordering: config_mutex -> data_mutex -> stats_mutex
 * All code acquiring multiple locks must follow this order.
 */

static K_MUTEX_DEFINE(config_mutex);
static K_MUTEX_DEFINE(data_mutex);
static K_MUTEX_DEFINE(stats_mutex);

void update_with_stats(void)
{
	k_mutex_lock(&data_mutex, K_FOREVER);
	k_mutex_lock(&stats_mutex, K_FOREVER);  /* Always after data_mutex */

	/* Update both data and stats */

	k_mutex_unlock(&stats_mutex);
	k_mutex_unlock(&data_mutex);
}

/* Incorrect - inconsistent lock ordering causes deadlock */

/* Thread A */
void thread_a_work(void)
{
	k_mutex_lock(&mutex_x, K_FOREVER);
	k_mutex_lock(&mutex_y, K_FOREVER);  /* Waits for Y */
	/* ... */
}

/* Thread B */
void thread_b_work(void)
{
	k_mutex_lock(&mutex_y, K_FOREVER);
	k_mutex_lock(&mutex_x, K_FOREVER);  /* Deadlock with Thread A */
	/* ... */
}
```

#### C-CON-003: Do Not Hold Locks Across Blocking Calls :yellow_circle:

**Rationale**: Holding locks during blocking operations (I/O, sleeps) starves other threads and can cause priority inversion. Based on SEI CERT CON36-C principles.

```c
/* Correct - release lock before blocking */
void send_and_log(const uint8_t *data, size_t len)
{
	k_mutex_lock(&send_mutex, K_FOREVER);
	memcpy(tx_buffer, data, len);
	size_t tx_len = len;
	k_mutex_unlock(&send_mutex);  /* Release before blocking I/O */

	/* Blocking send outside the lock */
	int ret = uart_tx(uart_dev, tx_buffer, tx_len, K_MSEC(1000));

	k_mutex_lock(&stats_mutex, K_FOREVER);
	if (ret == 0) {
		stats.bytes_sent += tx_len;
	} else {
		stats.send_errors++;
	}
	k_mutex_unlock(&stats_mutex);
}

/* Incorrect - blocking I/O while holding lock */
void bad_send(const uint8_t *data, size_t len)
{
	k_mutex_lock(&send_mutex, K_FOREVER);
	memcpy(tx_buffer, data, len);
	uart_tx(uart_dev, tx_buffer, len, K_MSEC(1000));  /* Blocks! */
	k_mutex_unlock(&send_mutex);  /* Other threads starve */
}
```

#### C-CON-004: Use Appropriate Synchronization Primitives :yellow_circle:

**Rationale**: Different synchronization needs require different primitives. Using wrong primitive causes bugs or performance issues. Based on SEI CERT CON37-C.

```c
/* Correct - mutex for mutual exclusion with ownership */
static K_MUTEX_DEFINE(resource_mutex);

void access_resource(void)
{
	k_mutex_lock(&resource_mutex, K_FOREVER);
	/* Only one thread at a time, same thread can unlock */
	k_mutex_unlock(&resource_mutex);
}

/* Correct - semaphore for signaling and counting */
static K_SEM_DEFINE(work_available, 0, 10);

void producer(void)
{
	/* Prepare work item */
	k_sem_give(&work_available);  /* Signal consumer */
}

void consumer(void)
{
	k_sem_take(&work_available, K_FOREVER);  /* Wait for work */
	/* Process work */
}

/* Correct - message queue for data passing between threads */
K_MSGQ_DEFINE(event_queue, sizeof(struct event), 16, 4);

void event_producer(struct event *evt)
{
	k_msgq_put(&event_queue, evt, K_NO_WAIT);
}

void event_consumer(void)
{
	struct event evt;
	while (k_msgq_get(&event_queue, &evt, K_FOREVER) == 0) {
		handle_event(&evt);
	}
}

/* Incorrect - semaphore where mutex needed */
static K_SEM_DEFINE(bad_lock, 1, 1);

void thread_a(void)
{
	k_sem_take(&bad_lock, K_FOREVER);
	/* Another thread could k_sem_give() without taking */
}
```

#### C-CON-005: Clean Up Thread-Specific Resources :yellow_circle:

**Rationale**: Resources allocated by threads must be freed when threads exit to prevent leaks. Based on SEI CERT CON30-C.

```c
/* Correct - cleanup before thread exit */
void worker_thread(void *p1, void *p2, void *p3)
{
	struct worker_context *ctx = k_malloc(sizeof(*ctx));
	if (ctx == NULL) {
		return;
	}

	ctx->initialized = true;

	while (!ctx->shutdown_requested) {
		do_work(ctx);
	}

	/* Cleanup before exit */
	if (ctx->buffer != NULL) {
		k_free(ctx->buffer);
	}
	k_free(ctx);
}

/* Correct - using thread resource table for cleanup */
struct thread_resources {
	void *buffer;
	int fd;
	bool active;
};

static struct thread_resources thread_res[MAX_THREADS];
static K_MUTEX_DEFINE(res_mutex);

void register_thread_resource(k_tid_t tid, void *buffer, int fd)
{
	k_mutex_lock(&res_mutex, K_FOREVER);
	/* Register resources for later cleanup */
	k_mutex_unlock(&res_mutex);
}

/* Incorrect - no cleanup */
void bad_worker(void *p1, void *p2, void *p3)
{
	void *data = k_malloc(1024);
	while (running) {
		process(data);
	}
	/* Thread exits without freeing data - memory leak */
}
```

### 9.2 Interrupt Context Rules

#### C-CON-006: ISR-Safe Operations Only in Interrupt Context :red_circle:

**Rationale**: ISRs cannot block; using blocking APIs in ISRs causes system hangs or crashes.

```c
/* Correct - non-blocking operations in ISR */
void uart_isr(const struct device *dev, void *user_data)
{
	uint8_t byte;

	while (uart_irq_rx_ready(dev)) {
		uart_fifo_read(dev, &byte, 1);
		/* Use _isr variants or K_NO_WAIT */
		if (k_msgq_put(&rx_queue, &byte, K_NO_WAIT) != 0) {
			dropped_bytes++;
		}
	}
}

/* Correct - defer work to thread context */
static K_WORK_DEFINE(process_work, process_handler);

void gpio_isr(const struct device *dev, struct gpio_callback *cb, uint32_t pins)
{
	/* Minimal work in ISR, defer processing */
	k_work_submit(&process_work);
}

void process_handler(struct k_work *work)
{
	/* Can use blocking APIs here */
	k_mutex_lock(&data_mutex, K_FOREVER);
	process_data();
	k_mutex_unlock(&data_mutex);
}

/* Incorrect - blocking in ISR */
void bad_isr(const struct device *dev, void *user_data)
{
	k_mutex_lock(&data_mutex, K_FOREVER);  /* Will hang! */
	k_sleep(K_MSEC(10));                   /* Never sleep in ISR! */
	k_msgq_put(&queue, &data, K_FOREVER);  /* May block forever! */
}
```

---

## 10. Error Handling

Robust error handling is critical for embedded systems reliability. Rules derive from SEI CERT C ERR30-C through ERR34-C.

### 10.1 Error Handling Rules

#### C-ERR-001: Check All Function Return Values :red_circle:

**Rationale**: Unchecked errors lead to undefined behavior and security vulnerabilities. Based on SEI CERT ERR33-C.

```c
/* Correct - check and handle errors */
int init_subsystem(void)
{
	int ret;

	ret = gpio_pin_configure(gpio_dev, PIN, GPIO_OUTPUT);
	if (ret != 0) {
		LOG_ERR("GPIO configure failed: %d", ret);
		return ret;
	}

	ret = i2c_configure(i2c_dev, I2C_SPEED_STANDARD);
	if (ret != 0) {
		LOG_ERR("I2C configure failed: %d", ret);
		return ret;
	}

	return 0;
}

/* Correct - explicit ignore with cast when truly unneeded */
(void)k_mutex_unlock(&mutex);  /* Cannot fail if lock was held */

/* Incorrect - ignoring return values */
void bad_init(void)
{
	gpio_pin_configure(gpio_dev, PIN, GPIO_OUTPUT);  /* May fail */
	i2c_write(i2c_dev, data, len, addr);             /* May fail */
	/* Continuing with uninitialized hardware */
}
```

#### C-ERR-002: Use Zephyr Error Codes Consistently :yellow_circle:

**Rationale**: Consistent error codes enable proper error propagation and debugging.

```c
/* Correct - use standard Zephyr/POSIX error codes */
#include <errno.h>

int validate_config(const struct config *cfg)
{
	if (cfg == NULL) {
		return -EINVAL;  /* Invalid argument */
	}

	if (cfg->size > MAX_CONFIG_SIZE) {
		return -ENOSPC;  /* No space */
	}

	if (!is_supported_version(cfg->version)) {
		return -ENOTSUP;  /* Not supported */
	}

	return 0;  /* Success */
}

/* Correct - propagate errors with context */
int process_request(const struct request *req)
{
	int ret;

	ret = validate_config(&req->config);
	if (ret != 0) {
		LOG_ERR("Config validation failed: %d", ret);
		return ret;  /* Propagate the specific error */
	}

	/* ... */
	return 0;
}

/* Incorrect - inconsistent error values */
int bad_validate(const struct config *cfg)
{
	if (cfg == NULL) {
		return -1;      /* What does -1 mean? */
	}
	if (cfg->size > MAX) {
		return 1;       /* Positive error? */
	}
	return 0;
}
```

#### C-ERR-003: Handle errno Correctly :yellow_circle:

**Rationale**: `errno` is only meaningful immediately after a function that sets it and only when that function indicates an error. Based on SEI CERT ERR30-C.

```c
/* Correct - errno usage with standard library functions */
#include <errno.h>
#include <stdlib.h>

int parse_integer(const char *str, long *result)
{
	char *endptr;

	errno = 0;  /* Clear errno before call */
	long value = strtol(str, &endptr, 10);

	/* Check error indicators first */
	if (endptr == str) {
		LOG_ERR("No digits found");
		return -EINVAL;
	}

	if (errno == ERANGE) {
		LOG_ERR("Value out of range");
		return -ERANGE;
	}

	if (*endptr != '\0') {
		LOG_ERR("Extra characters after number");
		return -EINVAL;
	}

	*result = value;
	return 0;
}

/* Incorrect - checking errno without error indication */
long bad_parse(const char *str)
{
	long value = strtol(str, NULL, 10);
	if (errno != 0) {  /* errno might be stale */
		return 0;
	}
	return value;
}
```

#### C-ERR-004: Clean Up Resources on Error Paths :red_circle:

**Rationale**: Error paths must release allocated resources to prevent leaks. Based on SEI CERT ERR33-C.

```c
/* Correct - cleanup on all error paths */
int create_session(struct session **out)
{
	struct session *sess = NULL;
	int ret;

	sess = k_malloc(sizeof(*sess));
	if (sess == NULL) {
		return -ENOMEM;
	}

	sess->rx_buf = k_malloc(RX_BUF_SIZE);
	if (sess->rx_buf == NULL) {
		ret = -ENOMEM;
		goto err_free_sess;
	}

	sess->tx_buf = k_malloc(TX_BUF_SIZE);
	if (sess->tx_buf == NULL) {
		ret = -ENOMEM;
		goto err_free_rx;
	}

	ret = session_init(sess);
	if (ret != 0) {
		goto err_free_tx;
	}

	*out = sess;
	return 0;

err_free_tx:
	k_free(sess->tx_buf);
err_free_rx:
	k_free(sess->rx_buf);
err_free_sess:
	k_free(sess);
	return ret;
}

/* Correct - cleanup helper function */
void destroy_session(struct session *sess)
{
	if (sess == NULL) {
		return;
	}
	k_free(sess->tx_buf);
	k_free(sess->rx_buf);
	k_free(sess);
}

/* Incorrect - resource leak on error */
int bad_create(struct session **out)
{
	struct session *sess = k_malloc(sizeof(*sess));
	if (sess == NULL) {
		return -ENOMEM;
	}

	sess->rx_buf = k_malloc(RX_BUF_SIZE);
	if (sess->rx_buf == NULL) {
		return -ENOMEM;  /* sess leaked! */
	}

	sess->tx_buf = k_malloc(TX_BUF_SIZE);
	if (sess->tx_buf == NULL) {
		return -ENOMEM;  /* sess and rx_buf leaked! */
	}

	*out = sess;
	return 0;
}
```

#### C-ERR-005: Use Assertions for Invariants, Not Errors :yellow_circle:

**Rationale**: Assertions document programmer assumptions and catch bugs during development. They are not for handling runtime errors. Based on SEI CERT ERR00-C.

```c
/* Correct - assertions for programming errors */
void process_buffer(uint8_t *buf, size_t len)
{
	/* These should never happen if API is used correctly */
	__ASSERT(buf != NULL, "Buffer cannot be NULL");
	__ASSERT(len > 0, "Length must be positive");
	__ASSERT(len <= MAX_BUF_SIZE, "Length exceeds maximum");

	/* Process buffer */
}

/* Correct - runtime checks for external input */
int handle_user_input(const uint8_t *data, size_t len)
{
	/* These CAN happen at runtime from external sources */
	if (data == NULL || len == 0) {
		return -EINVAL;
	}

	if (len > MAX_INPUT_SIZE) {
		LOG_WRN("Input too large: %zu", len);
		return -ENOSPC;
	}

	/* Process input */
	return 0;
}

/* Incorrect - using assert for runtime conditions */
int bad_handle(const uint8_t *data, size_t len)
{
	__ASSERT(data != NULL, "Data is null");  /* User could trigger this */
	__ASSERT(len <= MAX, "Too long");        /* Runtime condition */
	/* Asserts may be compiled out in release builds */
}
```

### 10.2 Logging Guidelines

| Level | When to Use | Example |
|-------|-------------|---------|
| `LOG_ERR` | Unrecoverable errors, operation failures | `LOG_ERR("Failed to init: %d", ret);` |
| `LOG_WRN` | Recoverable issues, degraded operation | `LOG_WRN("Retry %d of %d", attempt, max);` |
| `LOG_INF` | Significant state changes, milestones | `LOG_INF("System initialized");` |
| `LOG_DBG` | Detailed debugging information | `LOG_DBG("Processing packet len=%d", len);` |

```c
/* Correct logging setup */
#include <zephyr/logging/log.h>
LOG_MODULE_REGISTER(my_module, CONFIG_MY_MODULE_LOG_LEVEL);

int init_module(void)
{
	LOG_INF("Initializing module v%d.%d", MAJOR_VER, MINOR_VER);

	int ret = hardware_init();
	if (ret != 0) {
		LOG_ERR("Hardware init failed: %d", ret);
		return ret;
	}

	LOG_DBG("Hardware registers: 0x%08x", read_status_reg());
	return 0;
}
```

---

## 11. Control Flow

### 11.1 Control Flow Rules

#### C-CTL-001: Always Use Braces for Control Structures :red_circle:

**Rationale**: Braces prevent bugs when modifying code and improve readability. Mandated by Zephyr and MISRA-C.

```c
/* Correct - braces on all control structures */
if (condition) {
	do_something();
}

if (error) {
	return -EINVAL;
}

while (count > 0) {
	process();
	count--;
}

for (int i = 0; i < len; i++) {
	buffer[i] = 0;
}

/* Incorrect - missing braces */
if (condition)
	do_something();

if (error)
	return -EINVAL;

while (count > 0)
	count--;
```

#### C-CTL-002: Terminate if-else-if Chains with else :yellow_circle:

**Rationale**: Explicit else documents that all cases were considered and provides a catch-all. Mandated by MISRA-C Rule 15.7.

```c
/* Correct - else terminates chain */
if (state == STATE_IDLE) {
	start_operation();
} else if (state == STATE_RUNNING) {
	continue_operation();
} else if (state == STATE_ERROR) {
	handle_error();
} else {
	/* Default case - should not happen */
	LOG_WRN("Unexpected state: %d", state);
}

/* Correct - simple if without else is fine */
if (needs_update) {
	perform_update();
}

/* Incorrect - missing final else */
if (state == STATE_IDLE) {
	start_operation();
} else if (state == STATE_RUNNING) {
	continue_operation();
}
/* What if state is something else? */
```

#### C-CTL-003: Switch Statements Must Be Complete :yellow_circle:

**Rationale**: Switches on enums should handle all values explicitly. Default catches unexpected values.

```c
/* Correct - all cases handled */
enum sensor_state {
	SENSOR_STATE_OFF,
	SENSOR_STATE_INIT,
	SENSOR_STATE_READY,
	SENSOR_STATE_ERROR
};

const char *state_to_string(enum sensor_state state)
{
	switch (state) {
	case SENSOR_STATE_OFF:
		return "off";
	case SENSOR_STATE_INIT:
		return "initializing";
	case SENSOR_STATE_READY:
		return "ready";
	case SENSOR_STATE_ERROR:
		return "error";
	default:
		/* Catch invalid values */
		return "unknown";
	}
}

/* Correct - fallthrough must be explicit */
int categorize_char(char c)
{
	switch (c) {
	case 'a':
	case 'e':
	case 'i':
	case 'o':
	case 'u':
		return VOWEL;
	case ' ':
	case '\t':
	case '\n':
		/* Intentional fallthrough for whitespace */
		__fallthrough;
	case '\r':
		return WHITESPACE;
	default:
		return CONSONANT;
	}
}

/* Incorrect - missing cases, implicit fallthrough */
int bad_switch(enum sensor_state state)
{
	switch (state) {
	case SENSOR_STATE_OFF:
		/* Missing break - unintended fallthrough */
	case SENSOR_STATE_READY:
		return 0;
	/* Missing cases for INIT and ERROR */
	}
}
```

#### C-CTL-004: Avoid Deep Nesting :yellow_circle:

**Rationale**: Deep nesting makes code hard to follow. Maximum 4 levels recommended.

```c
/* Correct - early returns reduce nesting */
int process_packet(struct packet *pkt)
{
	if (pkt == NULL) {
		return -EINVAL;
	}

	if (!is_valid_header(pkt)) {
		return -EPROTO;
	}

	if (pkt->length > MAX_PACKET_SIZE) {
		return -EMSGSIZE;
	}

	/* Main logic at low nesting level */
	return handle_payload(pkt);
}

/* Correct - extract functions to reduce nesting */
static int validate_packet(const struct packet *pkt)
{
	if (pkt == NULL) {
		return -EINVAL;
	}
	if (!is_valid_header(pkt)) {
		return -EPROTO;
	}
	if (pkt->length > MAX_PACKET_SIZE) {
		return -EMSGSIZE;
	}
	return 0;
}

int process_packet(struct packet *pkt)
{
	int ret = validate_packet(pkt);
	if (ret != 0) {
		return ret;
	}
	return handle_payload(pkt);
}

/* Incorrect - deep nesting */
int bad_process(struct packet *pkt)
{
	if (pkt != NULL) {
		if (is_valid_header(pkt)) {
			if (pkt->length <= MAX_PACKET_SIZE) {
				if (check_crc(pkt)) {
					if (has_permission(pkt)) {
						/* Five levels deep! */
					}
				}
			}
		}
	}
	return -EINVAL;
}
```

---

## 12. Macros and Preprocessor

### 12.1 Macro Rules

#### C-MAC-001: Parenthesize Macro Arguments and Expressions :red_circle:

**Rationale**: Prevents operator precedence bugs when macros are used in expressions. Based on MISRA-C Rule 20.7.

```c
/* Correct - fully parenthesized */
#define MIN(a, b) (((a) < (b)) ? (a) : (b))
#define MAX(a, b) (((a) > (b)) ? (a) : (b))
#define SQUARE(x) ((x) * (x))
#define ARRAY_SIZE(arr) (sizeof(arr) / sizeof((arr)[0]))

/* Correct - statement-like macros use do-while(0) */
#define LOG_AND_RETURN(ret)  \
	do {                 \
		LOG_ERR("Error: %d", (ret)); \
		return (ret);    \
	} while (0)

/* Incorrect - missing parentheses */
#define BAD_SQUARE(x) x * x         /* BAD_SQUARE(1+2) = 1+2*1+2 = 5, not 9 */
#define BAD_ADD(a, b) a + b         /* 2 * BAD_ADD(1, 2) = 2*1+2 = 4, not 6 */

/* Incorrect - braces without do-while */
#define BAD_MACRO(x) { foo(x); bar(x); }  /* Breaks if-else */
```

#### C-MAC-002: Prefer Inline Functions Over Function-Like Macros :yellow_circle:

**Rationale**: Inline functions provide type checking and avoid multiple evaluation issues. Based on MISRA-C Rule 4.9.

```c
/* Correct - inline function */
static inline int32_t clamp_value(int32_t val, int32_t min, int32_t max)
{
	if (val < min) {
		return min;
	}
	if (val > max) {
		return max;
	}
	return val;
}

/* Correct - inline function with type safety */
static inline uint32_t safe_subtract(uint32_t a, uint32_t b)
{
	return (a > b) ? (a - b) : 0;
}

/* Acceptable - simple macros without side-effect risk */
#define IS_POWER_OF_TWO(x) (((x) != 0) && (((x) & ((x) - 1)) == 0))

/* Incorrect - complex macro with multiple evaluation */
#define BAD_CLAMP(val, min, max) \
	((val) < (min) ? (min) : ((val) > (max) ? (max) : (val)))
/* BAD_CLAMP(get_sensor_value(), 0, 100) calls get_sensor_value() up to 3 times */
```

#### C-MAC-003: Do Not Redefine Standard Macros :red_circle:

**Rationale**: Redefining standard names causes unpredictable behavior and maintenance nightmares. Based on Zephyr guidelines.

```c
/* Correct - use existing Zephyr/standard macros */
#include <zephyr/sys/util.h>

size_t count = ARRAY_SIZE(my_array);    /* Zephyr's ARRAY_SIZE */
int minimum = MIN(a, b);                 /* Zephyr's MIN */
int maximum = MAX(a, b);                 /* Zephyr's MAX */

/* Correct - namespaced project-specific macros */
#define MYPROJ_ALIGN_UP(x, align) (((x) + (align) - 1) & ~((align) - 1))

/* Incorrect - redefining common macros */
#define MIN(a, b) ...     /* Conflicts with Zephyr's MIN */
#define ARRAY_SIZE(x) ... /* Conflicts with Zephyr's ARRAY_SIZE */
#define NULL ((void*)0)   /* Never redefine NULL */
```

---

## 13. Testing Requirements

### 13.1 Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% line coverage | 85% line coverage |
| Branch coverage | 60% | 75% |
| Integration tests | Critical paths | All public APIs |

### 13.2 Test Organization

```
tests/
+-- unit/
|   +-- test_sensor_driver.c
|   +-- test_protocol_parser.c
|   +-- CMakeLists.txt
|   +-- testcase.yaml
+-- integration/
|   +-- test_system_init.c
|   +-- CMakeLists.txt
|   +-- testcase.yaml
+-- boards/
    +-- native_posix.conf
    +-- qemu_cortex_m3.conf
```

### 13.3 Test Naming

| Element | Convention | Example |
|---------|------------|---------|
| Test files | `test_<module>.c` | `test_sensor_driver.c` |
| Test functions | `test_<function>_<scenario>` | `test_sensor_read_valid_channel` |
| Test suites | `<module>_tests` | `sensor_driver_tests` |

### 13.4 Testing Rules

#### C-TST-001: Test All Public Functions :yellow_circle:

**Rationale**: Public APIs are contracts; tests verify the contract is fulfilled.

```c
/* Correct - comprehensive test coverage */
ZTEST(sensor_tests, test_sensor_init_success)
{
	struct sensor_config cfg = {
		.channel = 0,
		.sample_rate = 100,
	};

	int ret = sensor_init(&cfg);

	zassert_equal(ret, 0, "Expected success, got %d", ret);
	zassert_true(sensor_is_initialized(), "Sensor should be initialized");
}

ZTEST(sensor_tests, test_sensor_init_null_config)
{
	int ret = sensor_init(NULL);

	zassert_equal(ret, -EINVAL, "Expected -EINVAL for NULL config");
}

ZTEST(sensor_tests, test_sensor_init_invalid_channel)
{
	struct sensor_config cfg = {
		.channel = 255,  /* Invalid */
		.sample_rate = 100,
	};

	int ret = sensor_init(&cfg);

	zassert_equal(ret, -EINVAL, "Expected -EINVAL for invalid channel");
}
```

---

## 14. Security Considerations

Embedded systems are often targets for attacks. These rules protect against common vulnerabilities.

### 14.1 Input Validation

#### C-SEC-001: Validate All External Input 🔴

**Rule**: Validate all data from external sources (UART, network, sensors, user input) before processing.

**Rationale**: External data can be malformed or malicious. Unvalidated input leads to buffer overflows, injection attacks, and undefined behavior.

**Correct Example**:
```c
int process_uart_command(const uint8_t *data, size_t len)
{
    if (len == 0 || len > MAX_CMD_LEN) {
        return -EINVAL;
    }
    uint8_t cmd = data[0];
    if (cmd < CMD_MIN || cmd > CMD_MAX) {
        return -ENOTSUP;
    }
    return execute_command(cmd, &data[1], len - 1);
}
```

**Incorrect Example**:
```c
int process_uart_command(const uint8_t *data, size_t len)
{
    uint8_t cmd = data[0];
    return execute_command(cmd, &data[1], len - 1);  /* VULNERABLE */
}
```

#### C-SEC-002: Use Bounded String Functions 🔴

**Rule**: Use size-bounded functions (`snprintf`, `strncpy`) instead of unbounded versions.

**Rationale**: Unbounded functions can overflow buffers if input exceeds destination size.

**Correct Example**:
```c
int ret = snprintf(buf, buf_size, "%s_%d", prefix, id);
if (ret < 0 || (size_t)ret >= buf_size) {
    LOG_ERR("Device name truncated");
}
```

**Incorrect Example**:
```c
sprintf(buf, "%s_%d", prefix, id);  /* BUFFER OVERFLOW RISK */
```

#### C-SEC-003: Prevent Format String Attacks 🔴

**Rule**: Never pass user-controlled data as the format string to printf-family functions or LOG_* macros.

**Correct Example**:
```c
LOG_INF("User message: %s", user_msg);
```

**Incorrect Example**:
```c
LOG_INF(user_msg);  /* VULNERABLE: user controls format string */
```

#### C-SEC-004: Validate Array Indices Before Access 🔴

**Rule**: Validate that array indices are within bounds before accessing array elements.

**Correct Example**:
```c
if (sensor_id >= SENSOR_COUNT) {
    return -EINVAL;
}
return sensor_values[sensor_id];
```

**Incorrect Example**:
```c
return sensor_values[sensor_id];  /* NO BOUNDS CHECK */
```

---

## 15. Integer Safety

Integer errors are a leading cause of vulnerabilities in C code.

#### C-INT-001: Check for Integer Overflow Before Arithmetic 🔴

**Rule**: Before performing arithmetic that could overflow, verify the operation is safe (CERT INT32-C).

**Correct Example**:
```c
if (a > 0 && b > 0 && a > INT_MAX / b) {
    return -EOVERFLOW;
}
*result = a * b;
```

#### C-INT-002: Validate Integer Conversions 🔴

**Rule**: When converting between integer types, ensure the value fits in the destination type (CERT INT31-C).

**Correct Example**:
```c
if (requested_size > UINT16_MAX) {
    return -EINVAL;
}
config.buffer_size = (uint16_t)requested_size;
```

#### C-INT-003: Check for Division by Zero 🔴

**Rule**: Before division or modulo operations, verify the divisor is not zero (CERT INT33-C).

---

## 16. Array and String Safety

Buffer overflows remain the most exploited vulnerability class.

#### C-ARR-001: Never Access Arrays Out of Bounds 🔴

**Rule**: Ensure all array accesses use indices within the valid range [0, size-1] (CERT ARR30-C).

#### C-STR-001: Ensure String Buffers Have Space for Null Terminator 🔴

**Rule**: When allocating string buffers, include space for the null terminator (CERT STR31-C).

**Correct Example**:
```c
char device_name[MAX_NAME_LEN + 1];
strncpy(device_name, name, MAX_NAME_LEN);
device_name[MAX_NAME_LEN] = '\0';
```

#### C-STR-002: Null-Terminate Strings Before Library Functions 🔴

**Rule**: Ensure strings are properly null-terminated before passing to library functions (CERT STR32-C).

---

## 17. Device Tree

Device tree bindings are central to Zephyr's hardware abstraction.

#### C-DT-001: Use DT_NODELABEL for Static Node References 🟡

**Rule**: Use `DT_NODELABEL()` or `DT_ALIAS()` to reference device tree nodes, not hardcoded paths.

**Correct Example**:
```c
#define SENSOR_NODE DT_NODELABEL(my_sensor)
static const struct device *sensor = DEVICE_DT_GET(SENSOR_NODE);
```

#### C-DT-002: Validate Device Existence Before Use 🔴

**Rule**: Check that devices obtained from devicetree are ready before using them.

**Correct Example**:
```c
const struct device *dev = DEVICE_DT_GET(DT_NODELABEL(my_sensor));
if (!device_is_ready(dev)) {
    return -ENODEV;
}
```

---

## 18. Kconfig

Kconfig is Zephyr's configuration system.

#### C-KCF-001: Access CONFIG_ Values Only Via Kconfig 🟡

**Rule**: Never hardcode CONFIG_ macro values. Always define them in Kconfig files.

#### C-KCF-002: Document All Kconfig Options 🟡

**Rule**: Every Kconfig option must have a help text explaining its purpose.

---

## 19. Power Management

Proper power management extends battery life and meets energy requirements.

#### C-PM-001: Register PM Hooks for Device State Management 🟡

**Rule**: Devices that hold state must register power management callbacks to save/restore state during sleep transitions.

#### C-PM-002: Use PM Constraints for Critical Sections 🟡

**Rule**: Use `pm_policy_state_lock_get()` to prevent sleep during time-critical operations.

---

## 20. Documentation

### 20.1 Documentation Requirements

| Element | Documentation Required |
|---------|----------------------|
| Public functions | Yes - Doxygen format with @brief, @param, @return |
| Public types | Yes - Purpose and usage |
| Complex algorithms | Yes - Explain approach, reference materials |
| Magic numbers | Yes - Explain derivation or reference |
| Workarounds | Yes - Reference issue/errata |

### 14.2 Documentation Rules

#### C-DOC-001: Document Public APIs with Doxygen :yellow_circle:

**Rationale**: API documentation enables users to understand usage without reading implementation.

```c
/* Correct - complete Doxygen documentation */

/**
 * @brief Read temperature from sensor
 *
 * Performs a blocking read of the temperature sensor and converts
 * the raw value to millidegrees Celsius.
 *
 * @param dev Pointer to the sensor device structure
 * @param channel Sensor channel (0-3)
 * @param[out] value Pointer to store the temperature in millidegrees C
 *
 * @return 0 on success
 * @return -EINVAL if dev or value is NULL, or channel is invalid
 * @return -EIO if communication with sensor failed
 * @return -EBUSY if sensor is currently in use
 *
 * @note This function may sleep; do not call from ISR context.
 */
int sensor_read_temperature(const struct device *dev,
			    uint8_t channel,
			    int32_t *value);

/* Correct - document non-obvious parameters */

/**
 * @brief Configure sensor sampling parameters
 *
 * @param interval_ms Sample interval in milliseconds. Valid range: 10-10000.
 *                    Values below 100ms may increase power consumption significantly.
 * @param oversample Oversampling factor (1, 2, 4, 8, or 16).
 *                   Higher values improve accuracy but increase conversion time.
 */
int sensor_configure(uint32_t interval_ms, uint8_t oversample);

/* Incorrect - incomplete documentation */

/**
 * @brief Read temperature
 */
int sensor_read_temperature(const struct device *dev,
			    uint8_t channel,
			    int32_t *value);
/* Missing: parameter descriptions, return values, usage notes */
```

---

## 21. Tooling Configuration

### 15.1 Required Tools

| Tool | Purpose | Version |
|------|---------|---------|
| clang-format | Code formatting | 14.0+ |
| clang-tidy | Static analysis | 14.0+ |
| cppcheck | Static analysis | 2.7+ |
| west | Zephyr build/flash | Latest |

### 15.2 clang-format Configuration

Save as `.clang-format` in project root:

```yaml
---
Language: Cpp
BasedOnStyle: LLVM

# Indentation
IndentWidth: 8
TabWidth: 8
UseTab: Always
IndentCaseLabels: false
NamespaceIndentation: None

# Braces
BreakBeforeBraces: Linux
BraceWrapping:
  AfterFunction: true
  AfterControlStatement: false
  AfterEnum: false
  AfterStruct: false
  BeforeElse: false
  SplitEmptyFunction: false

# Line handling
ColumnLimit: 100
ReflowComments: true
MaxEmptyLinesToKeep: 1

# Alignment
AlignAfterOpenBracket: Align
AlignConsecutiveAssignments: false
AlignConsecutiveDeclarations: false
AlignEscapedNewlines: Left
AlignOperands: true
AlignTrailingComments: true

# Spacing
SpaceAfterCStyleCast: false
SpaceBeforeAssignmentOperators: true
SpaceBeforeParens: ControlStatements
SpaceInEmptyParentheses: false
SpacesInCStyleCastParentheses: false
SpacesInParentheses: false
SpacesInSquareBrackets: false

# Includes
SortIncludes: true
IncludeBlocks: Preserve
IncludeCategories:
  - Regex: '^".*\.h"'
    Priority: 1
  - Regex: '^<zephyr/.*>'
    Priority: 2
  - Regex: '^<(std|errno|string|limits).*>'
    Priority: 3
  - Regex: '^<.*>'
    Priority: 4

# Pointers and references
DerivePointerAlignment: false
PointerAlignment: Right

# Other
AllowShortBlocksOnASingleLine: false
AllowShortCaseLabelsOnASingleLine: false
AllowShortFunctionsOnASingleLine: None
AllowShortIfStatementsOnASingleLine: false
AllowShortLoopsOnASingleLine: false
BreakStringLiterals: true
Cpp11BracedListStyle: false
KeepEmptyLinesAtTheStartOfBlocks: false
...
```

### 15.3 clang-tidy Configuration

Save as `.clang-tidy` in project root:

```yaml
---
Checks: >
  -*,
  bugprone-*,
  -bugprone-easily-swappable-parameters,
  cert-*,
  clang-analyzer-*,
  -clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling,
  concurrency-*,
  misc-*,
  -misc-unused-parameters,
  modernize-use-bool-literals,
  performance-*,
  portability-*,
  readability-braces-around-statements,
  readability-else-after-return,
  readability-function-size,
  readability-identifier-naming,
  readability-implicit-bool-conversion,
  readability-isolate-declaration,
  readability-misleading-indentation,
  readability-misplaced-array-index,
  readability-redundant-*,
  readability-simplify-*

WarningsAsErrors: >
  bugprone-use-after-move,
  cert-err33-c,
  cert-mem30-c,
  cert-mem31-c,
  clang-analyzer-core.*,
  clang-analyzer-deadcode.*,
  concurrency-mt-unsafe,
  readability-braces-around-statements

CheckOptions:
  - key: readability-identifier-naming.FunctionCase
    value: lower_case
  - key: readability-identifier-naming.VariableCase
    value: lower_case
  - key: readability-identifier-naming.GlobalConstantCase
    value: UPPER_CASE
  - key: readability-identifier-naming.MacroDefinitionCase
    value: UPPER_CASE
  - key: readability-identifier-naming.EnumConstantCase
    value: UPPER_CASE
  - key: readability-identifier-naming.StructCase
    value: lower_case
  - key: readability-function-size.LineThreshold
    value: '100'
  - key: readability-function-size.StatementThreshold
    value: '50'
  - key: readability-function-size.BranchThreshold
    value: '10'
  - key: readability-function-size.ParameterThreshold
    value: '6'
  - key: readability-function-size.NestingThreshold
    value: '4'

HeaderFilterRegex: '.*'
...
```

### 15.4 cppcheck Configuration

Save as `cppcheck.cfg`:

```
--enable=warning,style,performance,portability
--std=c11
--platform=unix32
--suppress=missingIncludeSystem
--suppress=unusedFunction
--inline-suppr
--error-exitcode=1
--force
```

### 15.5 Pre-commit Hooks

Save as `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: local
    hooks:
      - id: clang-format
        name: clang-format
        entry: clang-format
        args: [-i, --style=file, --Werror]
        language: system
        types: [c]

      - id: clang-tidy
        name: clang-tidy
        entry: clang-tidy
        args: [--config-file=.clang-tidy]
        language: system
        types: [c]

      - id: cppcheck
        name: cppcheck
        entry: cppcheck
        args: [--error-exitcode=1, --enable=warning,style]
        language: system
        types: [c]

      - id: trailing-whitespace
        name: Trim Trailing Whitespace
        entry: trailing-whitespace-fixer
        language: system
        types: [c]

      - id: check-added-large-files
        name: Check for large files
        entry: check-added-large-files
        args: [--maxkb=100]
        language: system
```

---

## 22. Code Review Checklist

Quick reference for code reviewers evaluating C/Zephyr changes:

### Structure and Organization
- [ ] Files are in correct locations per project structure
- [ ] Header guards present and correctly formatted
- [ ] Include ordering follows convention (Zephyr, stdlib, project)
- [ ] Static used for file-scoped identifiers

### Naming and Style
- [ ] Function/variable names are descriptive and lowercase_with_underscores
- [ ] Macros/constants are UPPERCASE_WITH_UNDERSCORES
- [ ] Subsystem functions have consistent prefix
- [ ] No Hungarian notation or abbreviations (except standard ones)

### Formatting
- [ ] Tabs used for indentation (not spaces)
- [ ] Braces on all control structures
- [ ] K&R brace style (functions on new line, control on same line)
- [ ] Lines under 100 characters
- [ ] Only C89-style comments (`/* */`)

### Memory Safety (Critical)
- [ ] All allocations checked for NULL
- [ ] No use-after-free (pointers NULLed after free)
- [ ] No double-free potential
- [ ] Resources freed on all error paths
- [ ] sizeof(*pointer) pattern used for allocations
- [ ] Buffer sizes verified before access

### Concurrency Safety (Critical)
- [ ] Shared data protected by mutex/atomic
- [ ] Lock ordering documented and consistent
- [ ] No blocking calls while holding locks
- [ ] ISRs use only K_NO_WAIT and non-blocking APIs
- [ ] Thread resources cleaned up on exit

### Error Handling
- [ ] All function return values checked
- [ ] Zephyr error codes used consistently
- [ ] errno handled correctly (cleared before, checked after)
- [ ] Resources cleaned up on error paths
- [ ] Assertions used only for invariants, not runtime errors

### Control Flow
- [ ] if-else-if chains end with else
- [ ] Switch statements have default case
- [ ] No implicit fallthrough (use __fallthrough)
- [ ] Nesting depth <= 4 levels

### Macros
- [ ] Arguments and expressions fully parenthesized
- [ ] Statement macros use do-while(0)
- [ ] No redefinition of standard macros
- [ ] Inline functions preferred over complex macros

### Testing
- [ ] New functions have corresponding tests
- [ ] Error cases tested
- [ ] Edge cases considered

### Documentation
- [ ] Public APIs have Doxygen comments
- [ ] Complex logic explained
- [ ] Magic numbers explained or named
- [ ] Workarounds reference issues/errata

---

## Appendix A: Rule Quick Reference

| ID | Rule | Tier |
|----|------|------|
| C-NAM-001 | Use Descriptive Function Names | :yellow_circle: |
| C-NAM-002 | Prefix Subsystem Functions | :yellow_circle: |
| C-NAM-003 | Use Static for File-Scoped Identifiers | :red_circle: |
| C-TYP-001 | Use stdint.h Types | :red_circle: |
| C-TYP-002 | Use Zephyr Time Types | :yellow_circle: |
| C-TYP-003 | Use Boolean Type Correctly | :yellow_circle: |
| C-CMT-001 | Use C89 Block Comments Only | :red_circle: |
| C-CMT-002 | Document Non-Obvious Code | :yellow_circle: |
| C-MEM-001 | Check Allocation Return Values | :red_circle: |
| C-MEM-002 | Never Access Freed Memory | :red_circle: |
| C-MEM-003 | Prevent Double Free | :red_circle: |
| C-MEM-004 | Free Memory in Consistent Order | :yellow_circle: |
| C-MEM-005 | Use Appropriate Allocation Strategy | :yellow_circle: |
| C-MEM-006 | Specify Sufficient Allocation Size | :red_circle: |
| C-CON-001 | Protect Shared Data with Synchronization | :red_circle: |
| C-CON-002 | Avoid Deadlocks with Lock Ordering | :red_circle: |
| C-CON-003 | Do Not Hold Locks Across Blocking Calls | :yellow_circle: |
| C-CON-004 | Use Appropriate Synchronization Primitives | :yellow_circle: |
| C-CON-005 | Clean Up Thread-Specific Resources | :yellow_circle: |
| C-CON-006 | ISR-Safe Operations Only in Interrupt Context | :red_circle: |
| C-ERR-001 | Check All Function Return Values | :red_circle: |
| C-ERR-002 | Use Zephyr Error Codes Consistently | :yellow_circle: |
| C-ERR-003 | Handle errno Correctly | :yellow_circle: |
| C-ERR-004 | Clean Up Resources on Error Paths | :red_circle: |
| C-ERR-005 | Use Assertions for Invariants, Not Errors | :yellow_circle: |
| C-CTL-001 | Always Use Braces for Control Structures | :red_circle: |
| C-CTL-002 | Terminate if-else-if Chains with else | :yellow_circle: |
| C-CTL-003 | Switch Statements Must Be Complete | :yellow_circle: |
| C-CTL-004 | Avoid Deep Nesting | :yellow_circle: |
| C-MAC-001 | Parenthesize Macro Arguments and Expressions | :red_circle: |
| C-MAC-002 | Prefer Inline Functions Over Function-Like Macros | :yellow_circle: |
| C-MAC-003 | Do Not Redefine Standard Macros | :red_circle: |
| C-TST-001 | Test All Public Functions | :yellow_circle: |
| C-DOC-001 | Document Public APIs with Doxygen | :yellow_circle: |

---

## Appendix B: Glossary

| Term | Definition |
|------|------------|
| ISR | Interrupt Service Routine - code executed in interrupt context |
| K_NO_WAIT | Zephyr timeout value meaning "do not block" |
| K_FOREVER | Zephyr timeout value meaning "block indefinitely" |
| MISRA-C | Motor Industry Software Reliability Association C Guidelines |
| RTOS | Real-Time Operating System |
| SEI CERT C | Software Engineering Institute CERT C Coding Standard |
| UB | Undefined Behavior - behavior not specified by the C standard |

---

## Appendix C: Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-04 | Initial release |

---

## Appendix D: References

- [Zephyr Project Coding Guidelines](https://docs.zephyrproject.org/latest/contribute/coding_guidelines/)
- [SEI CERT C Coding Standard](https://wiki.sei.cmu.edu/confluence/display/c/SEI+CERT+C+Coding+Standard)
- [MISRA C:2012 Guidelines](https://www.misra.org.uk/misra-c/)
- [Linux Kernel Coding Style](https://www.kernel.org/doc/html/latest/process/coding-style.html)
- [Zephyr API Documentation](https://docs.zephyrproject.org/latest/doxygen/html/index.html)
