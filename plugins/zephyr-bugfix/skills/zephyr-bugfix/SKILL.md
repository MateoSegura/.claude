---
name: zephyr-bugfix
description: "Diagnose and fix Zephyr RTOS firmware bugs on Nordic nRF52/nRF53 hardware using MTIB test bench — from Asana ticket through hardware reproduction, fault analysis, fix, Twister tests, power verification, and Bitbucket push. REQUIRES: MTIB MCP server, Zephyr project, Nordic nRF dev kit."
user-invocable: true
argument-hint: "<bug-description-or-asana-url>"
---

# Zephyr Bug-Fix Workflow

End-to-end firmware bug resolution for Nordic nRF chips running Zephyr RTOS via MTIB hardware test bench.

## Requirements

- **MTIB MCP server** running and accessible (provides `mtib_*` tools). To use a custom host/port, set `MTIB_HOST` and `MTIB_PORT` environment variables or modify `.mcp.json`.
- **west** installed and configured for the project
- **Bitbucket** repo access via SSH or HTTPS
- **Asana** access if starting from a ticket URL

## Files

| File | Purpose | Load When |
|------|---------|-----------|
| [reference/reproduction-protocol.md](reference/reproduction-protocol.md) | Build, flash, capture logs, confirm bug | Phase 2 (REPRODUCE) |
| [reference/nrf-fault-patterns.md](reference/nrf-fault-patterns.md) | ARM Cortex-M fault register decode, common nRF crash patterns | Phase 3 (DIAGNOSE) |
| [reference/debug-playbook.md](reference/debug-playbook.md) | Breakpoint strategy, register inspection, backtrace analysis | Phase 3 (DIAGNOSE) |
| [reference/fix-protocol.md](reference/fix-protocol.md) | Edit, build, flash, verify symptom gone | Phase 4 (FIX) |
| [reference/power-budgets.md](reference/power-budgets.md) | Power measurement methodology, regression thresholds | Phase 5 (VERIFY) |

## Workflow

### Phase 1: RECEIVE

**Goal**: Parse the bug and set up the workspace.

**Input**: Bug description (free text) or Asana task URL.

Steps:
1. If Asana URL: check `echo $ASANA_ACCESS_TOKEN` via Bash. If exists, call Asana API via `curl -s -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" "https://app.asana.com/api/1.0/tasks/<task-id>"`. If no token, ask user to paste title/description/steps from the ticket.
2. If free text, use as-is. Extract: repo URL, board/chip, affected component.
3. Check repo: `git remote get-url origin 2>/dev/null`. If matches target, verify `git status` is clean. Otherwise `git clone <repo-url>`. If clone fails: HALT, inform user about SSH/HTTPS credentials.
4. Create branch: `git checkout -b fix/<bug-slug>` (slug from title, lowercase, hyphens, max 50 chars)
5. Identify Zephyr board target (e.g., `nrf52840dk/nrf52840`, `nrf5340dk/nrf5340/cpuapp`)
6. Call `mtib_health_check`. If fails: HALT — "MTIB MCP server not reachable."
7. Call `mtib_list_targets`. If target board not found: HALT — show available targets.

**Exit criteria**: Git branch exists, MTIB responds, target hardware detected.

---

### Phase 2: REPRODUCE

**Goal**: Confirm the bug exists on hardware.

**Load now**: [reference/reproduction-protocol.md](reference/reproduction-protocol.md).

Follow the protocol to build, flash, capture UART/RTT logs, debug data (if crash), and baseline power. Document observed vs expected.

**Exit criteria**: Bug confirmed per the protocol's confirmation criteria. If not reproducible after 3 attempts: HALT.

---

### Phase 3: DIAGNOSE

**Goal**: Identify root cause.

**Load now**: [reference/nrf-fault-patterns.md](reference/nrf-fault-patterns.md) and [reference/debug-playbook.md](reference/debug-playbook.md).

Read the relevant source files from the backtrace. Then launch the fault-analyzer agent:

```
Task tool: subagent_type="general-purpose", model="sonnet"
Prompt:
  You are a firmware fault analysis specialist for ARM Cortex-M / Zephyr RTOS / Nordic nRF.

  Bug: <paste from Phase 1>
  UART/RTT logs: <paste from Phase 2>
  Backtrace: <paste if crash>
  Debug status: <paste if crash>
  Source files to read: <list paths>

  Analyze: classify fault, trace path, identify root cause (prioritize: registers → source → call chain → recent changes), classify category, assess confidence (HIGH/MEDIUM/LOW per criteria in agents/fault-analyzer.md).

  Output: ROOT CAUSE HYPOTHESIS with Confidence, Fault Type, Affected Files, Explanation, Suggested Fix Approach. If LOW: list additional data needed.

  Constraints: No fix code. No guessing registers. No assuming recent changes without evidence.
```

If confidence is LOW: use MTIB to gather more data (breakpoints, memory reads, Zephyr shell), then re-launch with additional evidence.

**Exit criteria**: Root cause at HIGH or MEDIUM confidence.

---

### Phase 4: FIX

**Goal**: Write the fix, rebuild, reflash.

**Load now**: [reference/fix-protocol.md](reference/fix-protocol.md).

Follow the protocol to edit source files, monitor hook auto-build, handle failures, flash, and verify symptom is gone.

**Exit criteria**: Build succeeds, firmware boots, original bug symptom absent.

---

### Phase 5: VERIFY

**Goal**: Comprehensive verification — tests, power, no regressions.

**Load now**: [reference/power-budgets.md](reference/power-budgets.md).

**5.1 — Twister tests**: Check for `tests/` directory. If exists, call `mtib_twister_run` or `west twister -T tests/ -p <board> --device-testing`. If no tests: skip, note "Manual verification required." All tests must pass or return to Phase 4.

**5.2 — Power**: Call `mtib_power_measure` (5000ms). Launch power-regressor agent:

```
Task tool: subagent_type="general-purpose", model="haiku"
Prompt:
  Compare power measurements. Baseline: <Phase 2 data>. Current: <Phase 5 data>.
  Thresholds: Average >10%, Peak >25%, Min >5% increase = FAIL.
  Output: POWER REGRESSION: PASS|FAIL with metric comparison table.
  Constraints: exact numbers only, no cause diagnosis, SKIP missing fields.
```

If FAIL: investigate (left-enabled peripherals, new timers, prevented sleep), return to Phase 4.

**5.3 — Comprehensive check**: Launch fix-verifier agent:

```
Task tool: subagent_type="general-purpose", model="sonnet"
Prompt:
  Final quality gate. Find reasons NOT to ship.
  Git diff: <paste>. Twister results: <paste or "No tests">. Power: <paste regressor output>. Build warnings: <paste>.
  Checklist: tests pass, power PASS, no new warnings, fix addresses root cause, no memory leaks, no ISR blocking calls, minimal changes, Zephyr-specific checks (Kconfig, DTS, IRQ priorities, stack sizes).
  Output: VERDICT: SHIP or BLOCK with issues list.
  Constraints: BLOCK if new warning, scope creep, or insufficient evidence. No style suggestions.
```

If BLOCK: address issues, re-run 5.3.

**Exit criteria**: Twister passes (or skipped), power within budget, fix-verifier says SHIP.

---

### Phase 6: SHIP

**Goal**: Commit, push, update tracking.

1. `git add <specific-files>` — only the fix files
2. Commit: `fix(<component>): <short description>` with explanation, Asana link, test results
3. `git push -u origin fix/<bug-slug>`
4. Inform user: branch name, remind to create Bitbucket PR. If Asana token exists, offer to update ticket via API. If no token, tell user to update manually.
5. Summary: bug, root cause, fix, verification results.

**Exit criteria**: Branch pushed, user informed.
