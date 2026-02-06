---
description: "Read UART or RTT logs from a Nordic nRF target via MTIB. Quick standalone command for checking device output."
---

# /logs

Read device logs.

## Usage

```
/logs [target] [duration-seconds]
```

## Examples

```
/logs
/logs nrf52840dk 30
/logs nrf5340dk 10
```

## Instructions

1. Call `mtib_health_check` to verify MTIB is reachable
2. If no target specified, call `mtib_list_targets` and let the user pick
3. Default duration is 10 seconds if not specified
4. Capture logs from both sources in parallel:
   - Call `mtib_uart_read` with the specified timeout
   - Call `mtib_zephyr_logs` to get Zephyr log backend output
5. Present the combined output with source labels:

```
=== UART Output (10s) ===
[00:00:00.000,000] <inf> main: Application started
[00:00:00.012,000] <inf> ble: Advertising started
...

=== Zephyr Logs ===
[00:00:00.000,000] <inf> main: Application started
...
```

6. If any errors (`<err>`) or warnings (`<wrn>`) appear in the logs, highlight them at the end:

```
Warnings/Errors detected:
  Line 42: <err> spi: TX transfer timeout
  Line 67: <wrn> ble: Disconnected (reason 0x08)
```
