# Reproduction Protocol

Detailed steps for Phase 2: REPRODUCE.

## Build and Flash

1. Build the firmware as-is: `west build -b <board> <app-path> -- -DCONFIG_LOG=y -DCONFIG_LOG_BACKEND_UART=y`
2. Flash to target: call `mtib_flash` with the built `zephyr.hex` or `zephyr.bin`
3. Reset the target: call `mtib_reset`

## Capture Logs

4. Capture logs from both sources:
   - Call `mtib_uart_read` with a timeout appropriate for the bug (default 30s)
   - Call `mtib_zephyr_logs` to capture Zephyr log backend output (RTT or other)
   - Use both outputs for analysis — UART may have bootloader messages, RTT has structured Zephyr logs

## Capture Debug Data (if crash)

5. If the bug involves a crash/HardFault:
   - Call `mtib_debug_connect`
   - Call `mtib_debug_backtrace` to get the call stack
   - Call `mtib_debug_status` to check halt reason

## Baseline Power

6. Capture baseline power measurement: call `mtib_power_measure` with `duration_ms: 5000`

## Document Evidence

7. Document what was observed vs. what was expected

## Confirmation Criteria

Bug is confirmed reproducible if:
- **Crash**: backtrace or fault registers captured AND crash reason matches bug description
- **Wrong output**: logs show the incorrect behavior described in the bug
- **Hang**: device stops responding and UART output freezes

If NONE of these occur after 3 attempts with clean builds: report "Bug not reproducible on current hardware/build. Verify: (1) correct target board, (2) correct firmware app, (3) correct steps to reproduce." HALT workflow.
