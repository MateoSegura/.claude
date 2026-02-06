---
description: "Start an end-to-end Zephyr firmware bug-fix workflow — from Asana ticket through hardware reproduction, diagnosis, fix, test, and ship."
---

# /bugfix

Start the Zephyr bug-fix workflow.

## Usage

```
/bugfix <bug-description-or-asana-url>
```

## Examples

```
/bugfix https://app.asana.com/0/1234567890/9876543210
/bugfix "nRF52840 HardFault during BLE advertising after 2 hours of operation"
/bugfix "UART TX drops characters when SPI transfer is active on nRF5340"
```

## Instructions

1. Read the zephyr-bugfix skill definition: load `skills/zephyr-bugfix/SKILL.md` from this plugin's directory
2. Execute all 6 phases in order: RECEIVE -> REPRODUCE -> DIAGNOSE -> FIX -> VERIFY -> SHIP
3. Load reference files only when the workflow specifies (not upfront)
4. Use the fault-analyzer, power-regressor, and fix-verifier agents as specified in each phase
5. The user's argument is the input for Phase 1 (RECEIVE) — it is either a bug description or an Asana task URL
