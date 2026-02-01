# Formatting Rules

Consistent formatting improves readability and reduces merge conflicts. Zephyr follows a style derived from the Linux kernel.

---

## Indentation and Spacing

| Aspect | Rule |
|--------|------|
| Indentation | Tabs (8-space width display) |
| Line length | 100 characters maximum (80 preferred) |
| Trailing whitespace | Forbidden |
| Final newline | Required |
| Spaces after keywords | Required (`if (`, `for (`, `while (`) |
| Spaces around operators | Required (`a + b`, not `a+b`) |
| No space after function name | `func(arg)`, not `func (arg)` |

---

## Brace Style

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

---

## Include Ordering

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

## Header File Organization

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

## File Naming

| Type | Convention | Example |
|------|------------|---------|
| Source files | lowercase_with_underscores | `sensor_driver.c` |
| Header files | lowercase_with_underscores | `sensor_driver.h` |
| Internal headers | suffix `_internal` | `sensor_driver_internal.h` |
| Test files | suffix `_test` | `sensor_driver_test.c` |
| Kconfig fragments | descriptive `.conf` | `debug_logging.conf` |
