# Power Management Rules (C-PM-*)

Proper power management extends battery life and meets energy requirements.

---

## C-PM-001: Register PM Hooks for Device State Management :yellow_circle:

**Tier**: Required

**Rationale**: Devices that hold state must register power management callbacks to save/restore state during sleep transitions.

```c
/* Correct - register PM callbacks */
static int my_driver_pm_action(const struct device *dev,
			       enum pm_device_action action)
{
	switch (action) {
	case PM_DEVICE_ACTION_SUSPEND:
		/* Save device state before sleep */
		save_device_state(dev);
		return 0;
	case PM_DEVICE_ACTION_RESUME:
		/* Restore device state after wake */
		restore_device_state(dev);
		return 0;
	default:
		return -ENOTSUP;
	}
}

PM_DEVICE_DT_DEFINE(DT_NODELABEL(my_device), my_driver_pm_action);
```

---

## C-PM-002: Use PM Constraints for Critical Sections :yellow_circle:

**Tier**: Required

**Rationale**: Use `pm_policy_state_lock_get()` to prevent sleep during time-critical operations.

```c
/* Correct - prevent sleep during critical operation */
void time_critical_transfer(void)
{
	/* Prevent system from entering sleep */
	pm_policy_state_lock_get(PM_STATE_SUSPEND_TO_RAM, PM_ALL_SUBSTATES);

	/* Perform time-critical operation */
	start_dma_transfer();
	wait_for_completion();

	/* Allow sleep again */
	pm_policy_state_lock_put(PM_STATE_SUSPEND_TO_RAM, PM_ALL_SUBSTATES);
}

/* Incorrect - no sleep prevention */
void bad_transfer(void)
{
	start_dma_transfer();
	/* System may sleep here, corrupting transfer */
	wait_for_completion();
}
```
