# Security Rules (C-SEC-*, C-INT-*, C-ARR-*, C-STR-*)

Embedded systems are often targets for attacks. These rules protect against common vulnerabilities including buffer overflows, integer errors, and injection attacks.

---

## Input Validation

### C-SEC-001: Validate All External Input :red_circle:

**Tier**: Critical

**Rationale**: External data can be malformed or malicious. Unvalidated input leads to buffer overflows, injection attacks, and undefined behavior.

```c
/* Correct */
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

/* Incorrect */
int process_uart_command(const uint8_t *data, size_t len)
{
	uint8_t cmd = data[0];
	return execute_command(cmd, &data[1], len - 1);  /* VULNERABLE */
}
```

---

### C-SEC-002: Use Bounded String Functions :red_circle:

**Tier**: Critical

**Rationale**: Unbounded functions can overflow buffers if input exceeds destination size.

```c
/* Correct */
int ret = snprintf(buf, buf_size, "%s_%d", prefix, id);
if (ret < 0 || (size_t)ret >= buf_size) {
	LOG_ERR("Device name truncated");
}

/* Incorrect */
sprintf(buf, "%s_%d", prefix, id);  /* BUFFER OVERFLOW RISK */
```

---

### C-SEC-003: Prevent Format String Attacks :red_circle:

**Tier**: Critical

**Rationale**: User-controlled format strings can read/write arbitrary memory.

```c
/* Correct */
LOG_INF("User message: %s", user_msg);

/* Incorrect */
LOG_INF(user_msg);  /* VULNERABLE: user controls format string */
```

---

### C-SEC-004: Validate Array Indices Before Access :red_circle:

**Tier**: Critical

**Rationale**: Out-of-bounds access causes crashes or security vulnerabilities.

```c
/* Correct */
if (sensor_id >= SENSOR_COUNT) {
	return -EINVAL;
}
return sensor_values[sensor_id];

/* Incorrect */
return sensor_values[sensor_id];  /* NO BOUNDS CHECK */
```

---

## Integer Safety

### C-INT-001: Check for Integer Overflow Before Arithmetic :red_circle:

**Tier**: Critical

**Rationale**: Integer overflow causes undefined behavior and security vulnerabilities. Based on CERT INT32-C.

```c
/* Correct */
if (a > 0 && b > 0 && a > INT_MAX / b) {
	return -EOVERFLOW;
}
*result = a * b;

/* Incorrect */
*result = a * b;  /* May overflow silently */
```

---

### C-INT-002: Validate Integer Conversions :red_circle:

**Tier**: Critical

**Rationale**: When converting between integer types, ensure the value fits in the destination type. Based on CERT INT31-C.

```c
/* Correct */
if (requested_size > UINT16_MAX) {
	return -EINVAL;
}
config.buffer_size = (uint16_t)requested_size;

/* Incorrect */
config.buffer_size = (uint16_t)requested_size;  /* May truncate */
```

---

### C-INT-003: Check for Division by Zero :red_circle:

**Tier**: Critical

**Rationale**: Division by zero causes undefined behavior. Based on CERT INT33-C.

```c
/* Correct */
if (divisor == 0) {
	return -EINVAL;
}
result = dividend / divisor;

/* Incorrect */
result = dividend / divisor;  /* May divide by zero */
```

---

## Array and String Safety

### C-ARR-001: Never Access Arrays Out of Bounds :red_circle:

**Tier**: Critical

**Rationale**: Buffer overflows remain the most exploited vulnerability class. Based on CERT ARR30-C.

```c
/* Correct */
if (index >= ARRAY_SIZE(buffer)) {
	return -EINVAL;
}
buffer[index] = value;

/* Incorrect */
buffer[index] = value;  /* No bounds check */
```

---

### C-STR-001: Ensure String Buffers Have Space for Null Terminator :red_circle:

**Tier**: Critical

**Rationale**: Missing null terminator causes buffer over-reads. Based on CERT STR31-C.

```c
/* Correct */
char device_name[MAX_NAME_LEN + 1];
strncpy(device_name, name, MAX_NAME_LEN);
device_name[MAX_NAME_LEN] = '\0';

/* Incorrect */
char device_name[MAX_NAME_LEN];  /* No space for null terminator */
strncpy(device_name, name, MAX_NAME_LEN);  /* May not be terminated */
```

---

### C-STR-002: Null-Terminate Strings Before Library Functions :red_circle:

**Tier**: Critical

**Rationale**: Passing non-null-terminated strings to library functions causes buffer over-reads. Based on CERT STR32-C.

```c
/* Correct */
char buf[16];
memcpy(buf, data, MIN(len, sizeof(buf) - 1));
buf[MIN(len, sizeof(buf) - 1)] = '\0';
LOG_INF("Received: %s", buf);

/* Incorrect */
char buf[16];
memcpy(buf, data, len);  /* May not be null-terminated */
LOG_INF("Received: %s", buf);  /* Buffer over-read */
```
