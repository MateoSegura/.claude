---
description: "Flash firmware to a Nordic nRF target via MTIB. Quick standalone command for flashing without the full bug-fix workflow."
---

# /flash

Flash firmware to hardware.

## Usage

```
/flash [target] [firmware-path]
```

## Examples

```
/flash
/flash nrf52840dk build/zephyr/zephyr.hex
/flash nrf5340dk/cpuapp
```

## Instructions

1. Call `mtib_health_check` to verify MTIB is reachable
2. If no target specified, call `mtib_list_targets` and let the user pick
3. If no firmware path specified, look for the most recent build output:
   - Check `build/zephyr/zephyr.hex` first
   - Then `build/zephyr/zephyr.bin`
   - If neither exists, ask the user
4. Call `mtib_flash` with the target and firmware path
5. Call `mtib_reset` to restart the target after flashing
6. Call `mtib_uart_read` with a 5-second timeout to capture boot logs
7. Report: firmware flashed, device reset, boot log summary (first 10 lines)
