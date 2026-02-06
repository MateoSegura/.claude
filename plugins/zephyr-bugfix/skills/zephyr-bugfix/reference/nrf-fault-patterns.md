# nRF Fault Patterns

Common crash and fault patterns on Nordic nRF52/nRF53 series running Zephyr RTOS.

## ARM Cortex-M Fault Registers

When a HardFault occurs, read these registers to classify the fault:

| Register | Address | Purpose |
|----------|---------|---------|
| CFSR | 0xE000ED28 | Configurable Fault Status (combined UsageFault + BusFault + MemManage) |
| HFSR | 0xE000ED2C | HardFault Status |
| MMFAR | 0xE000ED34 | MemManage Fault Address |
| BFAR | 0xE000ED38 | BusFault Address |
| AFSR | 0xE000ED3C | Auxiliary Fault Status |

### CFSR Bit Decode

**MemManage (bits 7:0)**:
| Bit | Name | Meaning |
|-----|------|---------|
| 7 | MMARVALID | MMFAR holds valid address |
| 5 | MLSPERR | Lazy FP state preservation fault |
| 4 | MSTKERR | Stacking fault (exception entry) |
| 3 | MUNSTKERR | Unstacking fault (exception return) |
| 1 | DACCVIOL | Data access violation — most common, check MMFAR for the bad address |
| 0 | IACCVIOL | Instruction access violation — jumped to non-executable region |

**BusFault (bits 15:8)**:
| Bit | Name | Meaning |
|-----|------|---------|
| 15 | BFARVALID | BFAR holds valid address |
| 13 | LSPERR | Lazy FP state preservation bus fault |
| 12 | STKERR | Stacking bus fault |
| 11 | UNSTKERR | Unstacking bus fault |
| 10 | IMPRECISERR | Imprecise data bus error — address unknown, often DMA or write buffer |
| 9 | PRECISERR | Precise data bus error — BFAR has the address |
| 8 | IBUSERR | Instruction bus error — bad fetch address |

**UsageFault (bits 25:16)**:
| Bit | Name | Meaning |
|-----|------|---------|
| 25 | DIVBYZERO | Division by zero (only if DIV_0_TRP enabled) |
| 24 | UNALIGNED | Unaligned access (only if UNALIGN_TRP enabled) |
| 19 | NOCP | Coprocessor access (FPU not enabled?) |
| 18 | INVPC | Invalid PC on exception return |
| 17 | INVSTATE | Invalid EPSR.T bit (tried to execute ARM instruction in Thumb-only core) |
| 16 | UNDEFINSTR | Undefined instruction — corrupted code, bad function pointer |

## Common nRF Crash Patterns

### 1. Stack Overflow
**Symptoms**: HardFault with MSTKERR or DACCVIOL, fault address near a thread's stack base.
**Zephyr clue**: `CONFIG_STACK_SENTINEL` or `CONFIG_MPU_STACK_GUARD` triggers `z_fatal_error` with reason `K_ERR_STACK_CHK_FAIL`.
**Debug steps**:
1. Check `kernel stacks` via Zephyr shell — look for usage near 100%
2. Read the thread's stack base from `kernel threads` output
3. Compare MMFAR with stack boundaries
**Common causes**: Deep call chains, large local arrays, recursive functions, unbounded `printk`/`LOG_*` in ISRs.
**Fix patterns**: Increase `CONFIG_<THREAD>_STACK_SIZE`, move large buffers to heap or static, reduce call depth.

### 2. Null Pointer Dereference
**Symptoms**: DACCVIOL with MMFAR at `0x00000000` or low address (< 0x100).
**Debug steps**:
1. Backtrace to find the faulting function
2. Look for uninitialized pointers, failed allocations not checked, callback pointers not set
**Fix patterns**: Add null checks, initialize pointers, verify allocations.

### 3. Use-After-Free / Double-Free
**Symptoms**: DACCVIOL or PRECISERR with address in heap region but content is garbage. Often intermittent.
**Zephyr clue**: If using `k_heap` or `k_malloc`, the freed block may be reused.
**Debug steps**:
1. Enable `CONFIG_SYS_HEAP_VALIDATE` and `CONFIG_SYS_HEAP_RUNTIME_STATS`
2. Set watchpoint on the suspect address: `mtib_debug_breakpoint` with type `WATCHPOINT`
3. Trace what writes to it after free
**Fix patterns**: Set pointers to NULL after free, use ownership patterns, consider `k_object_alloc`.

### 4. ISR-Context Violation
**Symptoms**: `z_fatal_error` with reason `K_ERR_KERNEL_OOPS`, log message about calling blocking API from ISR.
**Zephyr clue**: Log shows "cannot call from ISR context" or similar.
**Common causes**: Calling `k_sem_take` with timeout, `k_mutex_lock`, `k_sleep`, or `printk` (if UART backend) from an interrupt handler.
**Fix patterns**: Use `k_sem_give` (non-blocking) from ISR, defer work to a thread via `k_work_submit`, use `LOG_*` macros with deferred logging.

### 5. MPU Fault (nRF53 specific)
**Symptoms**: MemManage fault when accessing peripheral or memory region.
**nRF53 specifics**: The application core and network core have separate MPU configs. SPU (System Protection Unit) adds another layer.
**Debug steps**:
1. Check if the address is in a peripheral region (0x40000000-0x50000000)
2. Verify the peripheral is assigned to the correct domain in the device tree
3. Check `CONFIG_NRF_SPU_FLASH_REGION_SIZE` and `CONFIG_NRF_SPU_RAM_REGION_SIZE`
**Fix patterns**: Assign peripheral in device tree overlay, configure SPU permissions.

### 6. BLE SoftDevice / Controller Conflicts (nRF52)
**Symptoms**: Random crashes during BLE activity, often in `mpsl_*` or `sdc_*` functions.
**Common causes**: Interrupt priority conflict, insufficient MPSL timeslot, radio ISR preempted.
**Debug steps**:
1. Verify `CONFIG_MPSL` and `CONFIG_BT_CTLR` settings
2. Check that no user ISR uses priorities 0-1 (reserved for radio)
3. Look for flash write operations during BLE activity (nRF52 radio and flash share bus)
**Fix patterns**: Use `CONFIG_SOC_FLASH_NRF_RADIO_SYNC_MPSL=y`, avoid ISR priority 0-1, increase `CONFIG_BT_CTLR_SDC_TX_PACKET_COUNT`.

### 7. Watchdog Timeout
**Symptoms**: Device resets, reset reason register shows WDT. No fault registers set.
**Debug steps**:
1. Check `NRF_POWER->RESETREAS` or `NRF_RESET->RESETREAS` (nRF53)
2. If WDT bit set, a thread failed to feed the watchdog in time
3. Enable `CONFIG_THREAD_ANALYZER` to find stuck or starved threads
**Common causes**: Thread blocked on semaphore/mutex indefinitely, tight loop without yielding, priority inversion.
**Fix patterns**: Add timeout to all blocking calls, use `k_mutex_lock` with `K_MSEC(timeout)`, configure `CONFIG_WDT_ALLOW_CALLBACK` for debug.

### 8. Clock/Power Domain Issues
**Symptoms**: Peripheral not responding, reads return 0xFFFFFFFF, or device hangs after sleep.
**nRF specifics**: Peripherals are gated by PSEL registers and clock domains.
**Debug steps**:
1. Verify peripheral is enabled in device tree
2. Check `NRF_CLOCK->HFCLKSTAT` and `NRF_CLOCK->LFCLKSTAT`
3. After system-off: verify RAM retention settings and GPIO latch configuration
**Fix patterns**: Ensure `status = "okay"` in DTS, configure `CONFIG_PM_DEVICE` for peripherals that need explicit power management.

## Register Inspection Quick Reference

For Cortex-M33 (nRF53) and Cortex-M4 (nRF52), the stacked exception frame at the stack pointer contains:

| Offset | Register |
|--------|----------|
| 0x00 | R0 |
| 0x04 | R1 |
| 0x08 | R2 |
| 0x0C | R3 |
| 0x10 | R12 |
| 0x14 | LR (return address) |
| 0x18 | PC (faulting instruction) |
| 0x1C | xPSR |

To find the faulting instruction:
1. Get SP from backtrace
2. Read memory at SP+0x18 — this is the PC value at fault time
3. Use `addr2line` or map file to convert to source:line
