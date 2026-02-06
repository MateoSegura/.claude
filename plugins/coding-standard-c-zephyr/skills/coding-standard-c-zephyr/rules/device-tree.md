# Device Tree Rules (C-DT-*)

Device tree bindings are central to Zephyr's hardware abstraction.

---

## C-DT-001: Use DT_NODELABEL for Static Node References :yellow_circle:

**Tier**: Required

**Rationale**: Using node labels instead of hardcoded paths makes code portable across boards.

```c
/* Correct - use node labels */
#define SENSOR_NODE DT_NODELABEL(my_sensor)
static const struct device *sensor = DEVICE_DT_GET(SENSOR_NODE);

/* Incorrect - hardcoded path */
#define SENSOR_NODE DT_PATH(soc, i2c_40003000, sensor_48)
```

---

## C-DT-002: Validate Device Existence Before Use :red_circle:

**Tier**: Critical

**Rationale**: Devices may not exist on all boards or may fail initialization.

```c
/* Correct - check device is ready */
const struct device *dev = DEVICE_DT_GET(DT_NODELABEL(my_sensor));
if (!device_is_ready(dev)) {
	LOG_ERR("Sensor device not ready");
	return -ENODEV;
}

/* Incorrect - no readiness check */
const struct device *dev = DEVICE_DT_GET(DT_NODELABEL(my_sensor));
sensor_sample_fetch(dev);  /* May crash if device not ready */
```
