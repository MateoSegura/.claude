---
name: app-data-extraction
description: Extract all content from a mobile app by controlling a USB-connected phone from WSL via ADB. Automates screenshotting, swiping, OCR, and database construction to archive an app's full contents before it disappears.
---

# App Data Extraction

Systematically extract every piece of content from a mobile app by driving a
real phone over ADB from a WSL terminal. The phone connects to a Windows PC via
USB; WSL talks to the Windows-side ADB server. The pipeline is:

**screenshot → OCR → store → navigate → repeat**

Until every reachable screen has been captured and its content saved to a local
SQLite database.

## When to Use

- An app you rely on is being retired and you need to archive its data
- The app has no export feature or public API
- The app is mobile-only (no web version to scrape conventionally)
- You have physical access to the phone and a Windows PC with WSL

## Supporting Files

- [reference.md](reference.md) — ADB-over-WSL setup, OCR tooling, troubleshooting

---

## Phase Overview

| #   | Phase       | Goal                                           | Gate                              |
| --- | ----------- | ---------------------------------------------- | --------------------------------- |
| 1   | **SETUP**   | ADB bridge working, OCR tools installed        | `adb devices` shows the phone     |
| 2   | **RECON**   | Map the app's screen graph and content types    | User approves the screen map      |
| 3   | **PLAN**    | Design extraction strategy, present to user     | User approves the plan            |
| 4   | **EXTRACT** | Automated capture loop across all screens       | All mapped screens visited        |
| 5   | **STRUCTURE** | Parse raw OCR into a clean, queryable database | Database passes integrity checks  |
| 6   | **EXPORT**  | Deliver final archive in user's preferred format | User confirms completeness       |

---

## Phase 1: SETUP

### 1A — Verify ADB Bridge (WSL → Windows)

ADB cannot talk to USB devices directly from WSL. The approach is to connect
through the Windows-side ADB server.

```bash
# Check if adb is available
which adb || echo "adb not found — see reference.md for install steps"

# Point WSL at the Windows ADB server
# The user's Windows IP on the WSL bridge is usually:
export ADB_SERVER_SOCKET=tcp:$(ip route show default | awk '{print $3}'):5037

# Verify
adb devices
```

If `adb devices` returns the phone's serial number, the bridge is working.
If not, walk the user through the setup in [reference.md](reference.md).

**Important**: The user must:
1. Enable Developer Options on their phone (tap Build Number 7 times)
2. Enable USB Debugging in Developer Options
3. Connect the phone via USB and tap "Allow" on the authorization prompt
4. Have `adb.exe` running on the Windows side (Android SDK Platform-Tools)

### 1B — Install OCR and Image Tools

```bash
# Core dependencies
sudo apt-get update && sudo apt-get install -y tesseract-ocr python3-pip imagemagick

# Python packages
pip3 install pytesseract Pillow sqlite-utils
```

### 1C — Create Project Directory

```bash
mkdir -p app-extraction/{screenshots,ocr-text,raw-data}
cd app-extraction
```

### 1D — Verify End-to-End Pipeline

Take one test screenshot and OCR it:

```bash
adb exec-out screencap -p > screenshots/test.png
python3 -c "
from PIL import Image
import pytesseract
text = pytesseract.image_to_string(Image.open('screenshots/test.png'))
print(text[:500])
"
```

If text appears, the pipeline works. Proceed to Phase 2.

---

## Phase 2: RECON

The goal is to understand the app's navigation structure before automating
anything. This is a **manual + assisted** phase.

### 2A — Identify Top-Level Sections

Ask the user to describe the app:
- What are the main tabs or sections?
- Roughly how many items/pages per section?
- What types of content exist? (text, images, lists, cards, forms, etc.)
- Does the app use infinite scroll, pagination, or fixed pages?
- Are there any sections behind paywalls or login walls?

### 2B — Build a Screen Map

Create a simple screen graph. Example:

```
App Root
├── Home (scrollable feed — ~200 items)
│   └── Item Detail (text + image)
├── Library (grid — ~50 items)
│   └── Item Detail (text + audio player)
├── Favorites (list — ~30 items)
│   └── Item Detail (same as Library detail)
└── Settings (skip — no user content)
```

### 2C — Determine Navigation Actions

For each transition in the screen map, record the ADB action needed:

| From → To                | Action                                          |
| ------------------------ | ----------------------------------------------- |
| Home → Item Detail       | `adb shell input tap X Y`                       |
| Item Detail → Back       | `adb shell input keyevent KEYCODE_BACK`         |
| Scroll down in feed      | `adb shell input swipe 540 1500 540 500 300`    |
| Switch to Library tab    | `adb shell input tap TAB_X TAB_Y`               |

To find coordinates, use:

```bash
# Take a screenshot and open it to visually identify tap targets
adb exec-out screencap -p > screenshots/nav-reference.png
# Use ImageMagick to annotate coordinates
identify -verbose screenshots/nav-reference.png | head -20
```

Or enable "Pointer location" in Developer Options to see coordinates on-screen
as you tap.

### 2D — Estimate Scale

Document the estimated total volume:

```
Sections:       4
Total items:    ~280
Content types:  text, images
Est. screenshots: ~600
Est. time:      depends on delay settings
```

**Present the screen map to the user for approval before continuing.**

---

## Phase 3: PLAN

### 3A — Design the Extraction Script

Based on the screen map, build a Python extraction script with this structure:

```python
#!/usr/bin/env python3
"""App data extraction script — generated for [APP NAME]."""

import subprocess, time, os, json, hashlib
from PIL import Image
import pytesseract

# --- Config ---
SCREENSHOT_DIR = "screenshots"
OCR_DIR = "ocr-text"
DELAY = 1.5          # seconds between actions (tune for app responsiveness)
SWIPE_PAUSE = 1.0    # pause after swipe to let content load
SCREEN_W = 1080      # phone screen width (from: adb shell wm size)
SCREEN_H = 2400      # phone screen height

def adb(cmd):
    """Run an ADB command and return output."""
    result = subprocess.run(
        ["adb", "shell"] + cmd.split(),
        capture_output=True, timeout=10
    )
    return result.stdout.decode(errors="replace")

def screenshot(name):
    """Capture screen and return file path."""
    path = f"{SCREENSHOT_DIR}/{name}.png"
    subprocess.run(["adb", "exec-out", "screencap", "-p"],
                   stdout=open(path, "wb"), timeout=10)
    return path

def ocr(image_path):
    """Extract text from screenshot."""
    img = Image.open(image_path)
    # Crop out status bar and nav bar for cleaner OCR
    cropped = img.crop((0, 80, img.width, img.height - 120))
    return pytesseract.image_to_string(cropped)

def tap(x, y):
    adb(f"input tap {x} {y}")
    time.sleep(DELAY)

def swipe_up():
    """Scroll content up (swipe from bottom to top)."""
    adb(f"input swipe {SCREEN_W//2} {int(SCREEN_H*0.7)} {SCREEN_W//2} {int(SCREEN_H*0.3)} 300")
    time.sleep(SWIPE_PAUSE)

def back():
    adb("input keyevent KEYCODE_BACK")
    time.sleep(DELAY)

def screen_hash(path):
    """Hash screenshot to detect duplicate/unchanged screens."""
    with open(path, "rb") as f:
        return hashlib.md5(f.read()).hexdigest()
```

### 3B — Design the Scroll-Until-End Detection

The key challenge is knowing when you've reached the bottom of a scrollable
list. Strategy: **hash consecutive screenshots**. If two consecutive screenshots
produce the same hash, scrolling has stopped.

```python
def scroll_and_capture(section_name, max_pages=500):
    """Scroll through a section, capturing each page."""
    pages = []
    prev_hash = None

    for i in range(max_pages):
        path = screenshot(f"{section_name}_{i:04d}")
        h = screen_hash(path)

        if h == prev_hash:
            print(f"  End of {section_name} at page {i}")
            os.remove(path)  # duplicate
            break

        text = ocr(path)
        pages.append({
            "section": section_name,
            "page": i,
            "screenshot": path,
            "text": text,
            "hash": h
        })
        prev_hash = h
        swipe_up()

    return pages
```

### 3C — Design the Detail Page Extraction

For apps where each list item opens a detail view:

```python
def extract_detail_items(section_name, item_coords, max_scroll=50):
    """Tap each item in a list, capture its detail page, then go back."""
    items = []

    for idx, (x, y) in enumerate(item_coords):
        tap(x, y)
        time.sleep(DELAY)  # wait for detail to load

        # Scroll through the detail page
        detail_pages = scroll_and_capture(f"{section_name}_item{idx:04d}_detail")
        items.append({
            "section": section_name,
            "item_index": idx,
            "detail_pages": detail_pages
        })

        back()
        time.sleep(DELAY)

    return items
```

### 3D — Plan for Dynamic Item Discovery

If the list of tappable items isn't fixed (e.g., an infinite-scroll feed),
use OCR on each scroll frame to detect item boundaries, then tap into each:

```python
def find_item_tap_targets(screenshot_path):
    """Use OCR with bounding boxes to find tappable item regions."""
    img = Image.open(screenshot_path)
    data = pytesseract.image_to_data(img, output_type=pytesseract.Output.DICT)

    # Group text blocks into probable items by Y-position clustering
    blocks = []
    for i, text in enumerate(data["text"]):
        if text.strip():
            blocks.append({
                "text": text,
                "x": data["left"][i] + data["width"][i] // 2,
                "y": data["top"][i] + data["height"][i] // 2,
            })

    # Cluster by Y position to find distinct rows/items
    # (implementation depends on app layout)
    return blocks
```

### 3E — Present the Plan

Show the user:
1. The sections to be extracted and their order
2. The navigation strategy per section (scroll-only vs. scroll-and-tap)
3. Estimated screenshot count
4. The delay/timing settings
5. Output format (SQLite database + screenshot archive)

**Wait for user approval before starting extraction.**

---

## Phase 4: EXTRACT

### 4A — Run the Extraction

Execute the extraction script section by section. Monitor progress:

```python
import sqlite3, json

def init_db(db_path="app_data.db"):
    conn = sqlite3.connect(db_path)
    conn.execute("""
        CREATE TABLE IF NOT EXISTS pages (
            id INTEGER PRIMARY KEY,
            section TEXT,
            page_num INTEGER,
            item_index INTEGER,
            screenshot_path TEXT,
            raw_text TEXT,
            screen_hash TEXT,
            extracted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    """)
    conn.commit()
    return conn

def save_page(conn, section, page_num, item_index, screenshot_path, raw_text, screen_hash):
    conn.execute(
        "INSERT INTO pages (section, page_num, item_index, screenshot_path, raw_text, screen_hash) VALUES (?,?,?,?,?,?)",
        (section, page_num, item_index, screenshot_path, raw_text, screen_hash)
    )
    conn.commit()
```

### 4B — Handle Interruptions

The extraction may take a long time. Build in resilience:

```python
def get_last_checkpoint(conn, section):
    """Resume from where we left off if interrupted."""
    row = conn.execute(
        "SELECT MAX(page_num) FROM pages WHERE section = ?", (section,)
    ).fetchone()
    return (row[0] or -1) + 1
```

### 4C — Progress Reporting

```python
def report_progress(conn):
    rows = conn.execute(
        "SELECT section, COUNT(*), MAX(extracted_at) FROM pages GROUP BY section"
    ).fetchall()
    print("\n--- Extraction Progress ---")
    for section, count, last_time in rows:
        print(f"  {section}: {count} pages (last: {last_time})")
    total = conn.execute("SELECT COUNT(*) FROM pages").fetchone()[0]
    print(f"  TOTAL: {total} pages captured\n")
```

### 4D — Error Recovery

If the app crashes, shows a popup, or the screen changes unexpectedly:

```python
def verify_still_in_app(expected_section):
    """Take a screenshot and do a quick OCR sanity check."""
    path = screenshot("_verify")
    text = ocr(path)
    os.remove(path)

    # Check for common interruptions
    interruptions = ["has stopped", "not responding", "sign in", "update available"]
    for phrase in interruptions:
        if phrase.lower() in text.lower():
            print(f"WARNING: Detected interruption: '{phrase}'")
            return False
    return True
```

---

## Phase 5: STRUCTURE

### 5A — Parse Raw OCR into Structured Records

Once all raw text is captured, parse it into meaningful fields. This is
app-specific — work with the user to define the schema.

```python
def create_structured_table(conn, table_name, fields):
    """Create a structured table based on the app's content model."""
    cols = ", ".join(f"{name} {dtype}" for name, dtype in fields)
    conn.execute(f"CREATE TABLE IF NOT EXISTS {table_name} ({cols})")
    conn.commit()

# Example for a recipe app:
# create_structured_table(conn, "recipes", [
#     ("id", "INTEGER PRIMARY KEY"),
#     ("title", "TEXT"),
#     ("ingredients", "TEXT"),
#     ("instructions", "TEXT"),
#     ("category", "TEXT"),
#     ("source_page_id", "INTEGER REFERENCES pages(id)")
# ])
```

### 5B — Text Parsing Strategies

Different apps need different parsing approaches:

| Content Pattern       | Strategy                                      |
| --------------------- | --------------------------------------------- |
| Key-value pairs       | Regex: `r"^(.+?):\s*(.+)$"` per line         |
| Titled paragraphs     | Split on lines that are ALL CAPS or bold-like |
| Numbered lists        | Regex: `r"^\d+[\.\)]\s*(.+)$"`               |
| Card-based UI         | Cluster by OCR bounding box Y-gaps            |
| Tabular data          | Use pytesseract `--psm 6` (uniform block)     |

### 5C — Deduplication

Remove duplicate entries from overlapping screenshots:

```python
def deduplicate(conn):
    """Remove rows with identical raw_text within the same section."""
    conn.execute("""
        DELETE FROM pages WHERE id NOT IN (
            SELECT MIN(id) FROM pages GROUP BY section, screen_hash
        )
    """)
    conn.commit()
```

### 5D — Image Extraction

If the app has images worth saving, extract them from screenshots:

```python
def extract_images_from_screenshots(conn, output_dir="raw-data/images"):
    """Crop and save image regions from screenshots."""
    os.makedirs(output_dir, exist_ok=True)
    rows = conn.execute("SELECT id, screenshot_path FROM pages").fetchall()

    for page_id, path in rows:
        if not os.path.exists(path):
            continue
        img = Image.open(path)
        # Heuristic: main content image is usually in the top 60% of screen
        # Adjust crop region based on the specific app layout
        content_region = img.crop((0, 80, img.width, int(img.height * 0.6)))
        content_region.save(f"{output_dir}/page_{page_id}.png")
```

---

## Phase 6: EXPORT

### 6A — SQLite Database (default)

The database is already built during extraction. Verify it:

```bash
sqlite3 app_data.db ".tables"
sqlite3 app_data.db "SELECT section, COUNT(*) FROM pages GROUP BY section"
sqlite3 app_data.db "SELECT COUNT(*) FROM pages"
```

### 6B — JSON Export

```python
import json, sqlite3

def export_json(db_path="app_data.db", output="app_data.json"):
    conn = sqlite3.connect(db_path)
    conn.row_factory = sqlite3.Row
    rows = conn.execute("SELECT * FROM pages ORDER BY section, page_num").fetchall()
    data = [dict(row) for row in rows]
    with open(output, "w") as f:
        json.dump(data, f, indent=2, default=str)
    print(f"Exported {len(data)} records to {output}")
```

### 6C — CSV Export

```python
import csv, sqlite3

def export_csv(db_path="app_data.db", output="app_data.csv"):
    conn = sqlite3.connect(db_path)
    conn.row_factory = sqlite3.Row
    rows = conn.execute("SELECT * FROM pages ORDER BY section, page_num").fetchall()
    if not rows:
        return
    with open(output, "w", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=rows[0].keys())
        writer.writeheader()
        writer.writerows(dict(row) for row in rows)
    print(f"Exported {len(rows)} records to {output}")
```

### 6D — HTML Archive

```python
def export_html(db_path="app_data.db", output_dir="html-archive"):
    """Generate a browsable HTML archive with embedded screenshots."""
    os.makedirs(output_dir, exist_ok=True)
    conn = sqlite3.connect(db_path)
    conn.row_factory = sqlite3.Row

    sections = conn.execute("SELECT DISTINCT section FROM pages").fetchall()

    index = ["<html><head><title>App Archive</title></head><body>"]
    index.append("<h1>App Data Archive</h1><ul>")

    for (section,) in sections:
        index.append(f'<li><a href="{section}.html">{section}</a></li>')
        pages = conn.execute(
            "SELECT * FROM pages WHERE section=? ORDER BY page_num", (section,)
        ).fetchall()

        with open(f"{output_dir}/{section}.html", "w") as f:
            f.write(f"<html><body><h1>{section}</h1>")
            for page in pages:
                f.write(f"<div class='page'><h3>Page {page['page_num']}</h3>")
                if page["screenshot_path"] and os.path.exists(page["screenshot_path"]):
                    f.write(f"<img src='../{page['screenshot_path']}' width='360'>")
                f.write(f"<pre>{page['raw_text']}</pre></div><hr>")
            f.write("</body></html>")

    index.append("</ul></body></html>")
    with open(f"{output_dir}/index.html", "w") as f:
        f.write("\n".join(index))
```

### 6E — Full Archive Bundle

Package everything into a single archive:

```bash
tar czf app-archive-$(date +%Y%m%d).tar.gz \
    app_data.db \
    app_data.json \
    screenshots/ \
    raw-data/ \
    html-archive/
```

---

## Tuning Guide

| Setting           | Conservative | Balanced | Aggressive |
| ----------------- | ------------ | -------- | ---------- |
| `DELAY`           | 2.5s         | 1.5s     | 0.8s       |
| `SWIPE_PAUSE`     | 2.0s         | 1.0s     | 0.5s       |
| Duplicate check   | Every screen | Every 3  | Every 5    |
| Detail extraction | All items    | Sampled  | Skip       |

- **Conservative**: Slow, reliable. Use for apps with heavy animations or network loads.
- **Balanced**: Good default. Works for most offline-capable apps.
- **Aggressive**: Fast, may miss content if app can't keep up. Use for simple/local apps.

---

## Troubleshooting Quick Reference

| Problem                        | Solution                                                |
| ------------------------------ | ------------------------------------------------------- |
| `adb devices` empty            | See [reference.md](reference.md) §ADB-WSL Bridge       |
| OCR returns garbage            | Increase `DELAY`; check screenshot quality; try `--psm` |
| Screenshots are black          | App may block capture (DRM); try `scrcpy` workaround    |
| App crashes during extraction  | Reduce speed; add `verify_still_in_app()` checks        |
| Duplicate content in database  | Run `deduplicate()` in Phase 5                          |
| Wrong coordinates for taps     | Re-calibrate with "Pointer location" in Dev Options     |
| WSL can't reach ADB server     | Check Windows firewall; verify `ADB_SERVER_SOCKET`      |
