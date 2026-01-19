# Error Handling Rules (C-ERR-*)

Robust error handling is critical for embedded systems reliability. Rules derive from SEI CERT C ERR30-C through ERR34-C.

## Logging Levels

| Level | When to Use | Example |
|-------|-------------|---------|
| `LOG_ERR` | Unrecoverable errors, operation failures | `LOG_ERR("Failed to init: %d", ret);` |
| `LOG_WRN` | Recoverable issues, degraded operation | `LOG_WRN("Retry %d of %d", attempt, max);` |
| `LOG_INF` | Significant state changes, milestones | `LOG_INF("System initialized");` |
| `LOG_DBG` | Detailed debugging information | `LOG_DBG("Processing packet len=%d", len);` |

## Logging Setup

```c
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

## C-ERR-001: Check All Function Return Values :red_circle:

**Tier**: Critical

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

---

## C-ERR-002: Use Zephyr Error Codes Consistently :yellow_circle:

**Tier**: Required

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

---

## C-ERR-003: Handle errno Correctly :yellow_circle:

**Tier**: Required

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

---

## C-ERR-004: Clean Up Resources on Error Paths :red_circle:

**Tier**: Critical

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

---

## C-ERR-005: Use Assertions for Invariants, Not Errors :yellow_circle:

**Tier**: Required

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
