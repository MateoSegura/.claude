# Testing Rules (C-TST-*)

Testing ensures code correctness and prevents regressions in embedded systems where debugging is often difficult.

## Coverage Requirements

| Type | Minimum Coverage | Target Coverage |
|------|------------------|-----------------|
| Unit tests | 70% line coverage | 85% line coverage |
| Branch coverage | 60% | 75% |
| Integration tests | Critical paths | All public APIs |

## Test Organization

```
tests/
+-- unit/
|   +-- test_sensor_driver.c
|   +-- test_protocol_parser.c
|   +-- CMakeLists.txt
|   +-- testcase.yaml
+-- integration/
|   +-- test_system_init.c
|   +-- CMakeLists.txt
|   +-- testcase.yaml
+-- boards/
    +-- native_posix.conf
    +-- qemu_cortex_m3.conf
```

## Test Naming

| Element | Convention | Example |
|---------|------------|---------|
| Test files | `test_<module>.c` | `test_sensor_driver.c` |
| Test functions | `test_<function>_<scenario>` | `test_sensor_read_valid_channel` |
| Test suites | `<module>_tests` | `sensor_driver_tests` |

---

## C-TST-001: Test All Public Functions :yellow_circle:

**Tier**: Required

**Rationale**: Public APIs are contracts; tests verify the contract is fulfilled.

```c
/* Correct - comprehensive test coverage */
ZTEST(sensor_tests, test_sensor_init_success)
{
	struct sensor_config cfg = {
		.channel = 0,
		.sample_rate = 100,
	};

	int ret = sensor_init(&cfg);

	zassert_equal(ret, 0, "Expected success, got %d", ret);
	zassert_true(sensor_is_initialized(), "Sensor should be initialized");
}

ZTEST(sensor_tests, test_sensor_init_null_config)
{
	int ret = sensor_init(NULL);

	zassert_equal(ret, -EINVAL, "Expected -EINVAL for NULL config");
}

ZTEST(sensor_tests, test_sensor_init_invalid_channel)
{
	struct sensor_config cfg = {
		.channel = 255,  /* Invalid */
		.sample_rate = 100,
	};

	int ret = sensor_init(&cfg);

	zassert_equal(ret, -EINVAL, "Expected -EINVAL for invalid channel");
}
```

---

## C-TST-002: Test Error Paths :yellow_circle:

**Tier**: Required

**Rationale**: Error handling code is often under-tested and contains bugs.

```c
/* Correct - test both success and error cases */
ZTEST(driver_tests, test_i2c_read_success)
{
	uint8_t data[4];
	int ret = i2c_read_bytes(dev, REG_ADDR, data, sizeof(data));
	zassert_equal(ret, 0, "Expected success");
}

ZTEST(driver_tests, test_i2c_read_null_buffer)
{
	int ret = i2c_read_bytes(dev, REG_ADDR, NULL, 4);
	zassert_equal(ret, -EINVAL, "Expected -EINVAL for NULL buffer");
}

ZTEST(driver_tests, test_i2c_read_zero_length)
{
	uint8_t data[4];
	int ret = i2c_read_bytes(dev, REG_ADDR, data, 0);
	zassert_equal(ret, -EINVAL, "Expected -EINVAL for zero length");
}
```

---

## C-TST-003: Use Ztest Framework :yellow_circle:

**Tier**: Required

**Rationale**: Ztest is Zephyr's native test framework with built-in support for test discovery and reporting.

```c
#include <zephyr/ztest.h>

/* Define test suite */
ZTEST_SUITE(my_module_tests, NULL, NULL, NULL, NULL, NULL);

/* Test case */
ZTEST(my_module_tests, test_function_behavior)
{
	/* Arrange */
	struct input in = { .value = 42 };

	/* Act */
	int result = function_under_test(&in);

	/* Assert */
	zassert_equal(result, EXPECTED_VALUE, "Unexpected result: %d", result);
}
```

---

## C-TST-004: Use Test Fixtures for Setup/Teardown :green_circle:

**Tier**: Recommended

**Rationale**: Fixtures ensure consistent test state and prevent test interdependencies.

```c
/* Test fixture */
struct test_fixture {
	struct device *dev;
	uint8_t buffer[64];
};

static void *test_setup(void)
{
	struct test_fixture *fixture = k_malloc(sizeof(*fixture));
	zassert_not_null(fixture, "Failed to allocate fixture");

	fixture->dev = device_get_binding("SENSOR_0");
	zassert_not_null(fixture->dev, "Device not found");

	return fixture;
}

static void test_teardown(void *f)
{
	struct test_fixture *fixture = f;
	k_free(fixture);
}

ZTEST_SUITE(sensor_tests, NULL, test_setup, NULL, NULL, test_teardown);

ZTEST_F(sensor_tests, test_sensor_read)
{
	int32_t value;
	int ret = sensor_read(fixture->dev, &value);
	zassert_equal(ret, 0, "Read failed");
}
```
