# Power Budgets

Power measurement methodology and regression thresholds for Nordic nRF firmware verification via MTIB.

## Measurement Protocol

### Setup
1. Call `mtib_power_measure` with appropriate `duration_ms` (minimum 5000ms for stable readings)
2. Ensure the device is in a known, repeatable state before measuring:
   - Same firmware build
   - Same hardware configuration (no extra peripherals attached)
   - Same test scenario (idle, active radio, peripheral usage)
3. Take baseline measurement BEFORE making changes (Phase 2 of the bug-fix workflow)
4. Take verification measurement AFTER applying the fix (Phase 5)

### Measurement Modes

| Mode | Duration | Use Case |
|------|----------|----------|
| Snapshot | 1-2s | Quick check during development |
| Standard | 5s | Bug-fix verification baseline and comparison |
| Extended | 30s+ | Sleep/wake cycle analysis, BLE advertising intervals |
| Profile | 60s+ | Full power state machine characterization |

For bug-fix verification, **Standard (5s)** is the default. Use Extended only when the bug involves sleep modes or periodic behavior.

## Regression Thresholds

A power regression means the fix made things worse. These thresholds trigger a FAIL:

| Metric | Threshold | Action |
|--------|-----------|--------|
| Average current increase | >10% | FAIL — investigate |
| Peak current increase | >25% | FAIL — check for new peripheral activation |
| Sleep current increase | >5% | FAIL — likely left something enabled |
| Average current decrease | Any | INFO — note improvement, not a blocker |

### Interpreting Results

The MTIB `mtib_power_measure` returns:
- `average_current_ua`: Mean current in microamps
- `peak_current_ua`: Maximum current spike
- `min_current_ua`: Minimum current (ideally captures sleep floor)
- `energy_uj`: Total energy consumed during measurement window

**Comparison formula**:
```
delta_percent = ((after - before) / before) * 100
```

If `delta_percent` exceeds the threshold for that metric → FAIL.

## Typical nRF Power Budgets (Reference)

These are approximate ranges for sanity checking. Actual values depend on configuration.

### nRF52840
| State | Expected Range |
|-------|---------------|
| System OFF | 0.4 - 1.5 uA |
| System ON, idle (no peripherals) | 1.5 - 3 uA |
| CPU active (64 MHz, no radio) | 3 - 5 mA |
| BLE advertising (1s interval) | 10 - 30 uA average |
| BLE connected (100ms CI) | 30 - 80 uA average |
| BLE connected + throughput | 4 - 8 mA |
| Radio TX (0 dBm) | 5 - 6 mA peak |
| Radio TX (+8 dBm) | 14 - 16 mA peak |
| Flash write | 8 - 10 mA peak |
| UART active (115200) | 0.5 - 1.5 mA |

### nRF5340 (Application Core)
| State | Expected Range |
|-------|---------------|
| System OFF | 1 - 2 uA |
| System ON, idle | 2 - 5 uA |
| CPU active (128 MHz) | 3 - 6 mA |
| CPU active (64 MHz) | 2 - 4 mA |
| Network core active (BLE) | Add 3 - 5 mA |

### nRF52832
| State | Expected Range |
|-------|---------------|
| System OFF | 0.3 - 1 uA |
| System ON, idle | 1.2 - 2.5 uA |
| CPU active (64 MHz) | 3 - 4.5 mA |
| BLE advertising (1s interval) | 8 - 20 uA average |

## Common Power Regression Causes in Bug Fixes

1. **Left a peripheral enabled**: Fix enabled a peripheral (UART, SPI, timer) but didn't disable it after use. Check `pm_device_action_run()` calls.

2. **Changed timer behavior**: Fix added a periodic timer or shortened an interval. Check `k_timer_start` parameters.

3. **Prevented sleep entry**: Fix added a busy-wait, removed a `k_sleep`, or took a mutex that blocks the idle thread. Check that the idle thread can still reach `__WFE`/`__WFI`.

4. **Radio parameter change**: Fix modified BLE connection interval, advertising interval, or TX power. Check Bluetooth configuration.

5. **Logging overhead**: Fix added `LOG_DBG`/`LOG_INF` calls in a hot path. Log backends (UART, RTT) consume power. Check that debug logging is behind `CONFIG_LOG_DEFAULT_LEVEL`.

6. **Clock source change**: Fix required HFXO to be running when it was previously using HFINT. Check `CONFIG_CLOCK_CONTROL_NRF_K32SRC_*`.
