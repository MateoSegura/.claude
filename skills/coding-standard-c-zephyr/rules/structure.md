# Project Structure

Consistent project structure enables developers to navigate unfamiliar codebases and ensures build system compatibility.

---

## Directory Layout

```
project-root/
+-- CMakeLists.txt
+-- prj.conf
+-- app.overlay            /* Device tree overlay */
+-- Kconfig                /* Application Kconfig */
+-- src/
|   +-- main.c
|   +-- app_config.h       /* Application configuration */
|   +-- subsystem/
|       +-- subsystem.c
|       +-- subsystem.h
|       +-- subsystem_internal.h
+-- include/
|   +-- public_api.h       /* Public interfaces */
+-- boards/
|   +-- <board_name>.conf
|   +-- <board_name>.overlay
+-- tests/
|   +-- unit/
|   +-- integration/
+-- doc/
```

---

## Key Directories

| Directory | Purpose |
|-----------|---------|
| `src/` | Application source code |
| `include/` | Public headers for external use |
| `boards/` | Board-specific configurations |
| `tests/` | Unit and integration tests |
| `doc/` | Documentation |

---

## Configuration Files

| File | Purpose |
|------|---------|
| `CMakeLists.txt` | Build configuration |
| `prj.conf` | Default Kconfig settings |
| `app.overlay` | Device tree modifications |
| `Kconfig` | Application-specific options |
