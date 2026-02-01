# Control Flow Rules (C-CTL-*)

Control flow rules ensure code is readable, maintainable, and follows Zephyr/MISRA-C conventions.

---

## C-CTL-001: Always Use Braces for Control Structures :red_circle:

**Tier**: Critical

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

---

## C-CTL-002: Terminate if-else-if Chains with else :yellow_circle:

**Tier**: Required

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

---

## C-CTL-003: Switch Statements Must Be Complete :yellow_circle:

**Tier**: Required

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

---

## C-CTL-004: Avoid Deep Nesting :yellow_circle:

**Tier**: Required

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
