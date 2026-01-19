# Memory Management Rules (C-MEM-*)

Application code MAY use dynamic allocation with appropriate safeguards. All rules derive from SEI CERT C MEM30-C through MEM36-C.

## Allocation Strategy

| Strategy | When to Use |
|----------|-------------|
| Static allocation | Fixed-size, long-lived data |
| Stack allocation | Small, short-lived data |
| Memory pools | Fixed-size repeated allocations |
| k_malloc | Variable-size, application-lifetime data |

---

## C-MEM-001: Check Allocation Return Values :red_circle:

**Tier**: Critical

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

---

## C-MEM-002: Never Access Freed Memory :red_circle:

**Tier**: Critical

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

---

## C-MEM-003: Prevent Double Free :red_circle:

**Tier**: Critical

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

---

## C-MEM-004: Free Memory in Consistent Order :yellow_circle:

**Tier**: Required

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

---

## C-MEM-005: Use Appropriate Allocation Strategy :yellow_circle:

**Tier**: Required

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

---

## C-MEM-006: Specify Sufficient Allocation Size :red_circle:

**Tier**: Critical

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
