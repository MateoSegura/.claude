# Concurrency Rules (C-CON-*)

Zephyr applications commonly use threads, ISRs, and various synchronization primitives. Rules derive from SEI CERT C CON30-C through CON43-C.

## Synchronization Primitives

| Primitive | Use Case |
|-----------|----------|
| Mutex | Mutual exclusion with ownership |
| Semaphore | Signaling and counting |
| Message queue | Data passing between threads |
| Atomic operations | Simple counters and flags |

---

## C-CON-001: Protect Shared Data with Synchronization :red_circle:

**Tier**: Critical

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

---

## C-CON-002: Avoid Deadlocks with Lock Ordering :red_circle:

**Tier**: Critical

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

---

## C-CON-003: Do Not Hold Locks Across Blocking Calls :yellow_circle:

**Tier**: Required

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

---

## C-CON-004: Use Appropriate Synchronization Primitives :yellow_circle:

**Tier**: Required

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

---

## C-CON-005: Clean Up Thread-Specific Resources :yellow_circle:

**Tier**: Required

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

---

## C-CON-006: ISR-Safe Operations Only in Interrupt Context :red_circle:

**Tier**: Critical

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
