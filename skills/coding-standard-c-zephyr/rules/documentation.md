# Documentation Rules (C-DOC-*)

Documentation enables code understanding without reading implementation details.

## Documentation Requirements

| Element | Documentation Required |
|---------|----------------------|
| Public functions | Yes - Doxygen format with @brief, @param, @return |
| Public types | Yes - Purpose and usage |
| Complex algorithms | Yes - Explain approach, reference materials |
| Magic numbers | Yes - Explain derivation or reference |
| Workarounds | Yes - Reference issue/errata |

---

## C-DOC-001: Document Public APIs with Doxygen :yellow_circle:

**Tier**: Required

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

## C-DOC-002: Document Magic Numbers :yellow_circle:

**Tier**: Required

**Rationale**: Magic numbers without explanation are unmaintainable.

```c
/* Correct - explained constant */

/*
 * Timeout calculated as:
 * - Max conversion time: 15.5ms (from datasheet Table 7.5)
 * - I2C transaction overhead: ~2ms
 * - Safety margin: 2x
 * Total: (15.5 + 2) * 2 = 35ms, rounded to 50ms
 */
#define SENSOR_READ_TIMEOUT_MS 50

/* Incorrect - unexplained magic numbers */
#define TIMEOUT 50
```
