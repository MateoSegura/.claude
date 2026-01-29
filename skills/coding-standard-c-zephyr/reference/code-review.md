# Code Review Checklist

Quick reference for code reviewers evaluating C/Zephyr changes.

---

## Structure and Organization

- [ ] Files are in correct locations per project structure
- [ ] Header guards present and correctly formatted
- [ ] Include ordering follows convention (Zephyr, stdlib, project)
- [ ] Static used for file-scoped identifiers

## Naming and Style

- [ ] Function/variable names are descriptive and lowercase_with_underscores
- [ ] Macros/constants are UPPERCASE_WITH_UNDERSCORES
- [ ] Subsystem functions have consistent prefix
- [ ] No Hungarian notation or abbreviations (except standard ones)

## Formatting

- [ ] Tabs used for indentation (not spaces)
- [ ] Braces on all control structures
- [ ] K&R brace style (functions on new line, control on same line)
- [ ] Lines under 100 characters
- [ ] Only C89-style comments (`/* */`)

## Memory Safety (Critical)

- [ ] All allocations checked for NULL
- [ ] No use-after-free (pointers NULLed after free)
- [ ] No double-free potential
- [ ] Resources freed on all error paths
- [ ] sizeof(*pointer) pattern used for allocations
- [ ] Buffer sizes verified before access

## Concurrency Safety (Critical)

- [ ] Shared data protected by mutex/atomic
- [ ] Lock ordering documented and consistent
- [ ] No blocking calls while holding locks
- [ ] ISRs use only K_NO_WAIT and non-blocking APIs
- [ ] Thread resources cleaned up on exit

## Error Handling

- [ ] All function return values checked
- [ ] Zephyr error codes used consistently
- [ ] errno handled correctly (cleared before, checked after)
- [ ] Resources cleaned up on error paths
- [ ] Assertions used only for invariants, not runtime errors

## Control Flow

- [ ] if-else-if chains end with else
- [ ] Switch statements have default case
- [ ] No implicit fallthrough (use __fallthrough)
- [ ] Nesting depth <= 4 levels

## Macros

- [ ] Arguments and expressions fully parenthesized
- [ ] Statement macros use do-while(0)
- [ ] No redefinition of standard macros
- [ ] Inline functions preferred over complex macros

## Security

- [ ] All external inputs validated
- [ ] Bounded string functions used (snprintf, not sprintf)
- [ ] No format string vulnerabilities
- [ ] Array indices bounds-checked
- [ ] Integer operations checked for overflow

## Device Tree & Kconfig

- [ ] DT_NODELABEL used for device references
- [ ] device_is_ready() called before device use
- [ ] Kconfig options have help text
- [ ] No hardcoded CONFIG_ values

## Testing

- [ ] New functions have corresponding tests
- [ ] Error cases tested
- [ ] Edge cases considered

## Documentation

- [ ] Public APIs have Doxygen comments
- [ ] Complex logic explained
- [ ] Magic numbers explained or named
- [ ] Workarounds reference issues/errata

---

## Glossary

| Term | Definition |
|------|------------|
| ISR | Interrupt Service Routine - code executed in interrupt context |
| K_NO_WAIT | Zephyr timeout value meaning "do not block" |
| K_FOREVER | Zephyr timeout value meaning "block indefinitely" |
| MISRA-C | Motor Industry Software Reliability Association C Guidelines |
| RTOS | Real-Time Operating System |
| SEI CERT C | Software Engineering Institute CERT C Coding Standard |
| UB | Undefined Behavior - behavior not specified by the C standard |
