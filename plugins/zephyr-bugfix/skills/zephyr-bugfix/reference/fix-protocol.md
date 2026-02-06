# Fix Protocol

Detailed steps for Phase 4: FIX.

## Edit and Build

1. Edit the affected source files to implement the fix
   - The `hooks.json` will auto-trigger `west build` after `.c`/`.h` file edits
2. Check the hook output after the edit completes:
   - Look for "FAILED" or "error:" in the output — this means build failure
   - If no hook output appears, manually run: `west build` via Bash

## Handle Build Failures

3. If build fails:
   - Read the error messages from the build output
   - Fix compilation errors in the source files
   - Save the file again (hook retriggers build automatically)
   - Repeat until build succeeds

## Flash and Verify

4. Once build succeeds, flash the new firmware: call `mtib_flash`
5. Reset and capture logs: call `mtib_reset`, then `mtib_uart_read`
6. Verify the specific bug symptom is gone in the logs
