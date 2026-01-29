# Naming Conventions (C-NAM-*)

Consistent naming improves code readability and prevents symbol collisions across large codebases.

## General Principles

- Names reveal intent and purpose
- Length proportional to scope (longer for wider visibility)
- No Hungarian notation
- Abbreviations only for well-known terms (e.g., `cfg`, `ctx`, `buf`)

## Naming Conventions Table

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

---

## C-NAM-001: Use Descriptive Function Names :yellow_circle:

**Tier**: Required

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

---

## C-NAM-002: Prefix Subsystem Functions :yellow_circle:

**Tier**: Required

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

---

## C-NAM-003: Use Static for File-Scoped Identifiers :red_circle:

**Tier**: Critical

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

## C-NAM-004: Use Named Constants Instead of Inline Magic Numbers :yellow_circle:

**Tier**: Required

**Rationale**: Magic numbers embedded in code are unmaintainable and unclear. Use named constants with descriptive names. This complements C-DOC-002 by requiring constants, not just documentation.

```c
/* Correct - named constants with clear purpose */
#define TELEMETRY_MAX_METRICS    16
#define TELEMETRY_MAX_SPANS       8
#define SENSOR_RETRY_COUNT        3

for (size_t i = 0; i < TELEMETRY_MAX_METRICS; i++) {
	process_metric(&metrics[i]);
}

/* Correct - constants defined where value derives from constraints */
#define UART_FIFO_SIZE           64    /* Hardware FIFO depth */
#define PROTOBUF_MAX_FIELDS      32    /* nanopb field limit */

/* Incorrect - inline magic numbers */
for (size_t i = 0; i < 16; i++) {       /* What is 16? */
	process_metric(&metrics[i]);
}

if (retry_count > 3) {                   /* Why 3? */
	return -ETIMEDOUT;
}

/* Incorrect - arbitrary limits without constants */
if (metric_count < 16 && span_count < 8) {
	serialize_telemetry();
}
```

**Exception**: Well-known values like 0, 1, -1 for common operations (initialization, incrementing, error returns) do not require named constants.
