# Kconfig Rules (C-KCF-*)

Kconfig is Zephyr's configuration system for compile-time options.

---

## C-KCF-001: Access CONFIG_ Values Only Via Kconfig :yellow_circle:

**Tier**: Required

**Rationale**: Never hardcode CONFIG_ macro values. Always define them in Kconfig files.

```c
/* Correct - use CONFIG_ values from Kconfig */
#if CONFIG_MY_FEATURE_ENABLED
static void my_feature_init(void)
{
	uint8_t buffer[CONFIG_MY_FEATURE_BUFFER_SIZE];
}
#endif

/* Incorrect - hardcoded values */
#define MY_BUFFER_SIZE 256  /* Should be in Kconfig */
```

---

## C-KCF-002: Document All Kconfig Options :yellow_circle:

**Tier**: Required

**Rationale**: Every Kconfig option must have help text explaining its purpose.

```kconfig
# Correct - documented option
config MY_MODULE_BUFFER_SIZE
	int "Buffer size for my module"
	default 256
	range 64 4096
	help
	  Size in bytes of the internal buffer used for message
	  processing. Larger values allow handling bigger messages
	  but consume more RAM.

# Incorrect - undocumented option
config MY_BUFFER
	int
	default 256
```
