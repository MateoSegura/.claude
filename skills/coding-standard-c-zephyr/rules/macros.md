# Macro and Preprocessor Rules (C-MAC-*)

Macros are powerful but can introduce subtle bugs. These rules ensure safe macro usage.

---

## C-MAC-001: Parenthesize Macro Arguments and Expressions :red_circle:

**Tier**: Critical

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

---

## C-MAC-002: Prefer Inline Functions Over Function-Like Macros :yellow_circle:

**Tier**: Required

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

---

## C-MAC-003: Do Not Redefine Standard Macros :red_circle:

**Tier**: Critical

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
