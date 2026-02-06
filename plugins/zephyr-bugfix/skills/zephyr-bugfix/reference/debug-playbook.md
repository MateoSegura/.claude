# Debug Playbook

Systematic debugging patterns using MTIB hardware debug tools for Zephyr firmware on Nordic nRF.

## Triage: Which Debug Approach?

| Symptom | Start With | Tools |
|---------|-----------|-------|
| HardFault / crash | Fault Register Analysis | `mtib_debug_backtrace`, `mtib_debug_memory`, `mtib_debug_status` |
| Device hangs (no crash) | Thread Analysis | `mtib_zephyr_shell` (`kernel threads`), `mtib_debug_connect` + halt |
| Intermittent crash | Watchpoint Trapping | `mtib_debug_breakpoint` (watchpoint mode) |
| Wrong output / behavior | Breakpoint + Step | `mtib_debug_breakpoint`, `mtib_debug_connect`, register reads |
| Performance issue | Power + Timing | `mtib_power_measure`, `mtib_zephyr_shell` (`kernel stacks`) |
| Boot failure | UART Logs | `mtib_uart_read`, `mtib_zephyr_logs` |
| Peripheral not working | Register Dump | `mtib_debug_memory` at peripheral base address |

## Pattern 1: Fault Register Analysis

Use when the device hits a HardFault, MemManage, BusFault, or UsageFault.

```
Step 1: Connect debugger
  → mtib_debug_connect (target_id, probe_id)

Step 2: Check halt reason
  → mtib_debug_status
  If halted due to exception, proceed. If running, halt first.

Step 3: Read fault registers
  → mtib_debug_memory (address: 0xE000ED28, length: 20)
  This reads CFSR (4 bytes), HFSR (4 bytes), gap (4 bytes), MMFAR (4 bytes), BFAR (4 bytes)

Step 4: Get backtrace
  → mtib_debug_backtrace
  Gives the call stack from the faulting context.

Step 5: Decode
  Cross-reference CFSR bits with nrf-fault-patterns.md to classify the fault.
  Use the PC from backtrace to identify the exact source line.
```

## Pattern 2: Thread Deadlock / Hang Analysis

Use when the device is alive but not responding or stuck.

```
Step 1: Connect and halt
  → mtib_debug_connect
  → (device should be halted or use mtib_debug_status to check)

Step 2: Get Zephyr thread state
  → mtib_zephyr_shell (command: "kernel threads")
  Look for:
  - Threads in "pending" state (waiting on something)
  - Threads with 0% CPU that should be active
  - Priority inversion (high-priority thread pending, low-priority running)

Step 3: Check stack usage
  → mtib_zephyr_shell (command: "kernel stacks")
  Look for threads near 100% stack usage.

Step 4: Identify the blocked resource
  If a thread is pending on a mutex/semaphore:
  - Read the thread's wait_q to find which object it's waiting on
  - Use mtib_debug_memory to read the kernel object
  - Identify the owner/holder of the mutex
  - Check if that owner is also blocked → circular dependency = deadlock

Step 5: If no deadlock, check for starvation
  - A thread may be runnable but never scheduled due to higher-priority threads running continuously
  - Check CONFIG_TIMESLICE_SIZE and CONFIG_TIMESLICE_PRIORITY
```

## Pattern 3: Watchpoint Trapping (Intermittent Bugs)

Use when a memory location is being corrupted but you don't know by what.

```
Step 1: Identify the suspect address
  From the fault or from observed corruption.
  Use the linker map file to translate symbol → address if needed.

Step 2: Set a data watchpoint
  → mtib_debug_breakpoint (address: <suspect>, type: WATCHPOINT_WRITE)
  The CPU will halt when anything writes to this address.

Step 3: Run and wait
  → mtib_debug_connect (resume device)
  Wait for the watchpoint to trigger.

Step 4: When halted, get context
  → mtib_debug_backtrace
  → mtib_debug_status
  The backtrace shows exactly which code wrote to the address.

Step 5: Clear watchpoint when done
  → (remove breakpoint through the debug interface)
```

## Pattern 4: Breakpoint + Register Inspection

Use when you know which function is suspect and want to inspect state at a specific point.

```
Step 1: Find the address
  From the ELF/map file, get the address of the target function or line.
  Or use a symbol name if the debug tools support it.

Step 2: Set breakpoint
  → mtib_debug_breakpoint (address: <target>, type: BREAKPOINT_HW)

Step 3: Run until hit
  Resume the device. When the breakpoint hits, the CPU halts.

Step 4: Inspect registers
  → mtib_debug_backtrace (gives register context)
  → mtib_debug_memory (read local variables on stack or global data)

Step 5: Step through
  Use single-step to walk through the code:
  → Resume and re-halt (or use step commands if available)

Step 6: Clean up
  Remove breakpoints after diagnosis.
```

## Pattern 5: Peripheral Register Dump

Use when a peripheral (SPI, I2C, UART, GPIO, timer) is not behaving as expected.

### nRF52840 Peripheral Base Addresses

| Peripheral | Base Address | Key Registers |
|-----------|-------------|---------------|
| GPIO P0 | 0x50000000 | OUT (0x504), DIR (0x514), PIN_CNF[n] (0x700+4n) |
| GPIO P1 | 0x50000300 | Same offsets from base |
| UARTE0 | 0x40002000 | ENABLE (0x500), BAUDRATE (0x524), CONFIG (0x56C) |
| UARTE1 | 0x40028000 | Same offsets |
| SPIM0 | 0x40003000 | ENABLE (0x500), FREQUENCY (0x524), CONFIG (0x554) |
| SPIM3 | 0x4002F000 | Same offsets |
| TWIM0 | 0x40003000 | ENABLE (0x500), FREQUENCY (0x524), ADDRESS (0x588) |
| TIMER0 | 0x40008000 | MODE (0x504), BITMODE (0x508), PRESCALER (0x510) |
| RTC0 | 0x4000B000 | PRESCALER (0x508), COUNTER (0x504) |
| WDT | 0x40010000 | CONFIG (0x504), CRV (0x504), RREN (0x508) |
| POWER | 0x40000000 | RESETREAS (0x400), SYSTEMOFF (0x500) |
| CLOCK | 0x40000000 | HFCLKSTAT (0x40C), LFCLKSTAT (0x418) |

```
Step 1: Read the peripheral enable register
  → mtib_debug_memory (address: <base> + 0x500, length: 4)
  If 0 or unexpected value, the peripheral isn't enabled.

Step 2: Read configuration registers
  → mtib_debug_memory for PSEL, CONFIG, FREQUENCY as appropriate.

Step 3: Check events and status
  → mtib_debug_memory for EVENT registers (usually at offset 0x100-0x1FF).
  Events stuck at 0 = peripheral not triggering.

Step 4: Check pin assignments
  Read PSEL registers to verify correct pin routing.
  Read GPIO PIN_CNF for those pins to verify direction and pull config.
```

## Debugging Efficiency Tips

1. **Read the logs first**. 80% of bugs are diagnosable from UART/RTT output alone. Only attach the debugger when logs aren't sufficient.

2. **Check the reset reason register early**. It tells you if this was a WDT reset, soft reset, pin reset, or power-on — saves time guessing.

3. **Use Zephyr shell commands before raw memory reads**. `kernel threads`, `kernel stacks`, `device list` give structured data faster than manually reading kernel objects.

4. **Set at most 4 hardware breakpoints** on Cortex-M4 (nRF52). The CPU has limited breakpoint comparators. Software breakpoints require flash writes and are slower.

5. **nRF5340: Debug the right core**. The application core and network core are debugged separately. Make sure you're connected to the correct one. The network core's debug access may need to be enabled via `CONFIG_NRF53_CPUNET_ENABLE`.

6. **After System OFF, the debugger disconnects**. If the bug involves system-off → wake, you need to set a breakpoint in the wake handler and re-attach after the device wakes.
