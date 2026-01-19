# Quick Reference - All Rules

Complete table of all rules with IDs, descriptions, and tiers.

## Rule Tiers

| Tier | Marker | Enforcement | Violation Response |
|------|--------|-------------|-------------------|
| **Critical** | :red_circle: | CI blocking | Build fails |
| **Required** | :yellow_circle: | CI warning | Must fix before merge |
| **Recommended** | :green_circle: | Linter hint | Fix encouraged |

---

## Rules by Category

### Naming (C-NAM-*)

| ID | Rule | Tier |
|----|------|------|
| C-NAM-001 | Use Descriptive Function Names | Required |
| C-NAM-002 | Prefix Subsystem Functions | Required |
| C-NAM-003 | Use Static for File-Scoped Identifiers | Critical |

### Types (C-TYP-*)

| ID | Rule | Tier |
|----|------|------|
| C-TYP-001 | Use stdint.h Types | Critical |
| C-TYP-002 | Use Zephyr Time Types | Required |
| C-TYP-003 | Use Boolean Type Correctly | Required |

### Comments (C-CMT-*)

| ID | Rule | Tier |
|----|------|------|
| C-CMT-001 | Use C89 Block Comments Only | Critical |
| C-CMT-002 | Document Non-Obvious Code | Required |

### Memory (C-MEM-*)

| ID | Rule | Tier |
|----|------|------|
| C-MEM-001 | Check Allocation Return Values | Critical |
| C-MEM-002 | Never Access Freed Memory | Critical |
| C-MEM-003 | Prevent Double Free | Critical |
| C-MEM-004 | Free Memory in Consistent Order | Required |
| C-MEM-005 | Use Appropriate Allocation Strategy | Required |
| C-MEM-006 | Specify Sufficient Allocation Size | Critical |

### Concurrency (C-CON-*)

| ID | Rule | Tier |
|----|------|------|
| C-CON-001 | Protect Shared Data with Synchronization | Critical |
| C-CON-002 | Avoid Deadlocks with Lock Ordering | Critical |
| C-CON-003 | Do Not Hold Locks Across Blocking Calls | Required |
| C-CON-004 | Use Appropriate Synchronization Primitives | Required |
| C-CON-005 | Clean Up Thread-Specific Resources | Required |
| C-CON-006 | ISR-Safe Operations Only in Interrupt Context | Critical |

### Error Handling (C-ERR-*)

| ID | Rule | Tier |
|----|------|------|
| C-ERR-001 | Check All Function Return Values | Critical |
| C-ERR-002 | Use Zephyr Error Codes Consistently | Required |
| C-ERR-003 | Handle errno Correctly | Required |
| C-ERR-004 | Clean Up Resources on Error Paths | Critical |
| C-ERR-005 | Use Assertions for Invariants, Not Errors | Required |

### Control Flow (C-CTL-*)

| ID | Rule | Tier |
|----|------|------|
| C-CTL-001 | Always Use Braces for Control Structures | Critical |
| C-CTL-002 | Terminate if-else-if Chains with else | Required |
| C-CTL-003 | Switch Statements Must Be Complete | Required |
| C-CTL-004 | Avoid Deep Nesting | Required |

### Macros (C-MAC-*)

| ID | Rule | Tier |
|----|------|------|
| C-MAC-001 | Parenthesize Macro Arguments and Expressions | Critical |
| C-MAC-002 | Prefer Inline Functions Over Function-Like Macros | Required |
| C-MAC-003 | Do Not Redefine Standard Macros | Critical |

### Security (C-SEC-*, C-INT-*, C-ARR-*, C-STR-*)

| ID | Rule | Tier |
|----|------|------|
| C-SEC-001 | Validate All External Input | Critical |
| C-SEC-002 | Use Bounded String Functions | Critical |
| C-SEC-003 | Prevent Format String Attacks | Critical |
| C-SEC-004 | Validate Array Indices Before Access | Critical |
| C-INT-001 | Check for Integer Overflow Before Arithmetic | Critical |
| C-INT-002 | Validate Integer Conversions | Critical |
| C-INT-003 | Check for Division by Zero | Critical |
| C-ARR-001 | Never Access Arrays Out of Bounds | Critical |
| C-STR-001 | Ensure String Buffers Have Space for Null Terminator | Critical |
| C-STR-002 | Null-Terminate Strings Before Library Functions | Critical |

### Testing (C-TST-*)

| ID | Rule | Tier |
|----|------|------|
| C-TST-001 | Test All Public Functions | Required |
| C-TST-002 | Test Error Paths | Required |
| C-TST-003 | Use Ztest Framework | Required |
| C-TST-004 | Use Test Fixtures for Setup/Teardown | Recommended |

### Device Tree (C-DT-*)

| ID | Rule | Tier |
|----|------|------|
| C-DT-001 | Use DT_NODELABEL for Static Node References | Required |
| C-DT-002 | Validate Device Existence Before Use | Critical |

### Kconfig (C-KCF-*)

| ID | Rule | Tier |
|----|------|------|
| C-KCF-001 | Access CONFIG_ Values Only Via Kconfig | Required |
| C-KCF-002 | Document All Kconfig Options | Required |

### Power Management (C-PM-*)

| ID | Rule | Tier |
|----|------|------|
| C-PM-001 | Register PM Hooks for Device State Management | Required |
| C-PM-002 | Use PM Constraints for Critical Sections | Required |

### Documentation (C-DOC-*)

| ID | Rule | Tier |
|----|------|------|
| C-DOC-001 | Document Public APIs with Doxygen | Required |
| C-DOC-002 | Document Magic Numbers | Required |

---

## Summary by Tier

### Critical Rules (21 total)

- C-NAM-003, C-TYP-001, C-CMT-001
- C-MEM-001, C-MEM-002, C-MEM-003, C-MEM-006
- C-CON-001, C-CON-002, C-CON-006
- C-ERR-001, C-ERR-004, C-CTL-001
- C-MAC-001, C-MAC-003
- C-SEC-001, C-SEC-002, C-SEC-003, C-SEC-004
- C-INT-001, C-INT-002, C-INT-003
- C-ARR-001, C-STR-001, C-STR-002, C-DT-002

### Required Rules (25 total)

- C-NAM-001, C-NAM-002, C-TYP-002, C-TYP-003, C-CMT-002
- C-MEM-004, C-MEM-005
- C-CON-003, C-CON-004, C-CON-005
- C-ERR-002, C-ERR-003, C-ERR-005
- C-CTL-002, C-CTL-003, C-CTL-004, C-MAC-002
- C-TST-001, C-TST-002, C-TST-003
- C-DT-001, C-KCF-001, C-KCF-002, C-PM-001, C-PM-002
- C-DOC-001, C-DOC-002

### Recommended Rules (1 total)

- C-TST-004
