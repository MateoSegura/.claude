---
description: "Measure power consumption of a Nordic nRF target via MTIB. Returns average, peak, and minimum current."
---

# /power-check

Measure power consumption.

## Usage

```
/power-check [target] [duration-seconds]
```

## Examples

```
/power-check
/power-check nrf52840dk 10
/power-check nrf5340dk 30
```

## Instructions

1. Call `mtib_health_check` to verify MTIB is reachable
2. If no target specified, call `mtib_list_targets` and let the user pick
3. Default duration is 5 seconds if not specified. Convert to milliseconds for the API call.
4. Call `mtib_power_measure` with `duration_ms` set to the duration
5. Report the results in a clear format:

```
Power Measurement (5.0s)
  Average: 3.2 mA
  Peak:    12.4 mA
  Min:     1.8 uA
  Energy:  16.0 mJ
```

6. If average current seems abnormal for the expected device state, add a note:
   - If > 5 mA while supposedly idle: "Note: Current seems high for idle state. Check if all peripherals are properly disabled."
   - If < 1 uA while supposedly active: "Note: Current seems low. Verify the target is powered and running."
