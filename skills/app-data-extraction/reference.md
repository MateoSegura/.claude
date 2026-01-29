# App Data Extraction — Reference

Supporting material for [SKILL.md](SKILL.md).

---

## ADB-WSL Bridge Setup

WSL cannot access USB devices directly. The solution is to run ADB on the
Windows side and connect from WSL through TCP.

### Option A: Connect to Windows ADB Server (Recommended)

**On Windows (PowerShell):**

1. Download [Android SDK Platform-Tools](https://developer.android.com/tools/releases/platform-tools)
2. Extract to `C:\platform-tools\`
3. Add to PATH or run directly
4. Connect phone via USB
5. Start the ADB server:

```powershell
C:\platform-tools\adb.exe start-server
C:\platform-tools\adb.exe devices
# Should show your device serial number
```

**On WSL:**

```bash
# Install ADB in WSL
sudo apt-get update && sudo apt-get install -y adb

# Find the Windows host IP (the default gateway from WSL's perspective)
WIN_IP=$(ip route show default | awk '{print $3}')
echo "Windows IP: $WIN_IP"

# Point ADB at the Windows server
export ADB_SERVER_SOCKET=tcp:${WIN_IP}:5037

# Verify
adb devices
```

If `adb devices` shows your phone, you're set. Add the export to `~/.bashrc`:

```bash
echo 'export ADB_SERVER_SOCKET=tcp:$(ip route show default | awk "{print \$3}"):5037' >> ~/.bashrc
```

### Option B: USB/IP Forwarding (Alternative)

Windows 11 supports `usbipd-win` to forward USB devices into WSL:

```powershell
# Windows (PowerShell as Admin)
winget install usbipd

# List USB devices
usbipd list

# Bind and attach the phone (replace BUSID with your device's bus ID)
usbipd bind --busid <BUSID>
usbipd attach --wsl --busid <BUSID>
```

```bash
# WSL
sudo apt-get install -y linux-tools-generic
adb devices
```

### Firewall Notes

If ADB can't connect from WSL, the Windows firewall may be blocking port 5037.

```powershell
# Windows PowerShell (Admin) — allow ADB through firewall
New-NetFirewallRule -DisplayName "ADB Server" -Direction Inbound -Protocol TCP -LocalPort 5037 -Action Allow
```

---

## Phone Preparation

### Enable Developer Options

1. Open **Settings → About Phone**
2. Tap **Build Number** 7 times
3. Enter PIN/password if prompted
4. "Developer mode has been enabled" toast appears

### Enable USB Debugging

1. Open **Settings → Developer Options**
2. Enable **USB Debugging**
3. Connect phone to PC via USB cable
4. Tap **Allow** on the "Allow USB debugging?" prompt on the phone
5. Check **Always allow from this computer** for convenience

### Useful Developer Options

| Setting               | Purpose                                        |
| --------------------- | ---------------------------------------------- |
| USB Debugging         | Required — allows ADB commands                 |
| Stay Awake            | Prevents screen timeout during extraction      |
| Pointer Location      | Shows X,Y coordinates on screen for calibration|
| Show Taps             | Visual feedback for automated taps             |
| Window Animation Scale| Set to 0.5x or Off to speed up navigation      |
| Transition Animation  | Set to 0.5x or Off to speed up navigation      |
| Animator Duration     | Set to 0.5x or Off to speed up navigation      |

Disabling animations is highly recommended — it makes transitions instant and
dramatically reduces the wait time between actions.

---

## Screen Information

Get the phone's screen dimensions (needed for swipe/tap coordinates):

```bash
# Screen resolution
adb shell wm size
# Example output: Physical size: 1080x2400

# Screen density
adb shell wm density
# Example output: Physical density: 420
```

---

## ADB Command Reference

### Screenshots

```bash
# Capture to file (fast, preferred)
adb exec-out screencap -p > screenshot.png

# Capture to device then pull (slower, more compatible)
adb shell screencap -p /sdcard/screen.png
adb pull /sdcard/screen.png screenshot.png
adb shell rm /sdcard/screen.png
```

### Touch and Gesture Input

```bash
# Tap at coordinates
adb shell input tap 540 1200

# Long press (hold for 1000ms)
adb shell input swipe 540 1200 540 1200 1000

# Swipe (scroll up — finger moves from bottom to top)
adb shell input swipe 540 1800 540 600 300

# Swipe (scroll down — finger moves from top to bottom)
adb shell input swipe 540 600 540 1800 300

# Swipe left (next page in horizontal pager)
adb shell input swipe 900 1200 180 1200 300

# Swipe right (previous page)
adb shell input swipe 180 1200 900 1200 300

# The last parameter (300) is duration in milliseconds
# Shorter = faster swipe, Longer = slower/more precise
```

### Navigation Keys

```bash
adb shell input keyevent KEYCODE_BACK          # Back button
adb shell input keyevent KEYCODE_HOME          # Home button
adb shell input keyevent KEYCODE_APP_SWITCH    # Recent apps
adb shell input keyevent KEYCODE_MENU          # Menu key
adb shell input keyevent KEYCODE_ENTER         # Enter/confirm
adb shell input keyevent KEYCODE_DPAD_DOWN     # D-pad down (for focus nav)
adb shell input keyevent KEYCODE_DPAD_UP       # D-pad up
adb shell input keyevent KEYCODE_TAB           # Tab (next field)
adb shell input keyevent KEYCODE_PAGE_DOWN     # Page down (in scrollable)
adb shell input keyevent KEYCODE_PAGE_UP       # Page up
```

### Text Input

```bash
# Type text (no spaces — use %s for spaces)
adb shell input text "hello%sworld"

# For complex text, use a broadcast intent instead
adb shell am broadcast -a ADB_INPUT_TEXT --es msg "text with spaces"
```

### App Management

```bash
# List running activities (find the current app)
adb shell dumpsys activity activities | head -30

# Get current focused window/app
adb shell dumpsys window | grep -i "mCurrentFocus"

# Launch an app by package name
adb shell monkey -p com.example.app -c android.intent.category.LAUNCHER 1

# Force stop an app
adb shell am force-stop com.example.app

# List installed packages
adb shell pm list packages | grep -i "keyword"
```

---

## OCR Configuration

### Tesseract Page Segmentation Modes (--psm)

Different modes work better for different layouts:

| PSM | Mode                         | Best For                        |
| --- | ---------------------------- | ------------------------------- |
| 3   | Fully automatic (default)    | General mixed content           |
| 4   | Single column of variable    | Main text content               |
| 6   | Uniform block of text        | Tables, forms, structured data  |
| 7   | Single text line             | Titles, headers, single fields  |
| 11  | Sparse text, no order        | Scattered UI elements           |
| 12  | Sparse text with OSD         | Scattered text with orientation |

```python
# Use a specific PSM
text = pytesseract.image_to_string(img, config="--psm 4")

# Get bounding boxes for each word (useful for finding tap targets)
data = pytesseract.image_to_data(img, output_type=pytesseract.Output.DICT)
```

### Improving OCR Quality

```python
from PIL import Image, ImageFilter, ImageEnhance

def preprocess_for_ocr(img_path):
    """Enhance screenshot for better OCR results."""
    img = Image.open(img_path)

    # Convert to grayscale
    img = img.convert("L")

    # Increase contrast
    enhancer = ImageEnhance.Contrast(img)
    img = enhancer.enhance(2.0)

    # Sharpen
    img = img.filter(ImageFilter.SHARPEN)

    # Optionally threshold to pure black/white
    img = img.point(lambda x: 0 if x < 128 else 255, "1")

    return img
```

### Language Packs

```bash
# Install additional language support
sudo apt-get install -y tesseract-ocr-spa  # Spanish
sudo apt-get install -y tesseract-ocr-fra  # French
sudo apt-get install -y tesseract-ocr-deu  # German
sudo apt-get install -y tesseract-ocr-jpn  # Japanese
sudo apt-get install -y tesseract-ocr-chi-sim  # Chinese (Simplified)

# Use in Python
text = pytesseract.image_to_string(img, lang="eng+spa")
```

---

## Alternative OCR Engines

If Tesseract quality is insufficient:

| Engine         | Pros                            | Install                           |
| -------------- | ------------------------------- | --------------------------------- |
| Tesseract      | Free, local, fast               | `apt install tesseract-ocr`       |
| EasyOCR        | Better accuracy, GPU support    | `pip install easyocr`             |
| PaddleOCR      | Strong on CJK languages         | `pip install paddleocr paddlepaddle` |
| Google Vision   | Highest accuracy, cloud-based  | `pip install google-cloud-vision` |

```python
# EasyOCR example
import easyocr
reader = easyocr.Reader(["en"])
results = reader.readtext("screenshot.png")
for (bbox, text, confidence) in results:
    print(f"{text} ({confidence:.2f})")
```

---

## Live View with scrcpy (Optional)

`scrcpy` mirrors the phone screen on your PC — useful during the RECON phase
to see what's happening in real time.

```bash
# Install
sudo apt-get install -y scrcpy

# Basic usage (will work if ADB is connected)
scrcpy

# Record while extracting (useful for debugging)
scrcpy --record extraction-session.mp4

# Display-only (disable control, just watch)
scrcpy --no-control
```

Note: scrcpy requires a display. In WSL, you'll need an X server (VcXsrv,
WSLg on Windows 11) or use the Windows-native scrcpy build.

---

## SQLite Quick Reference

```bash
# Open the database
sqlite3 app_data.db

# Useful queries
.tables                                          # List all tables
.schema pages                                    # Show table structure
SELECT COUNT(*) FROM pages;                      # Total pages captured
SELECT section, COUNT(*) FROM pages GROUP BY section;  # Pages per section
SELECT * FROM pages WHERE section='library' LIMIT 5;   # Sample data
SELECT raw_text FROM pages WHERE id = 42;        # View specific page text

# Export
.mode csv
.headers on
.output export.csv
SELECT * FROM pages;
.output stdout
```

---

## Troubleshooting

### ADB Issues

**`adb devices` shows "unauthorized":**
- Unlock the phone screen
- Look for the USB debugging authorization popup
- Tap "Allow" (and check "Always allow")
- If no popup: revoke authorizations in Developer Options, reconnect

**`adb devices` shows nothing:**
- Verify USB cable supports data (not charge-only)
- Try a different USB port
- On Windows: check Device Manager for "ADB Interface" driver
- Restart ADB: `adb kill-server && adb start-server`

**Connection drops during extraction:**
- Use `adb shell settings put global stay_on_while_plugged_in 3` to prevent sleep
- Check cable connection
- The script's checkpoint system allows resuming from last captured page

### OCR Issues

**OCR returns empty or garbage text:**
- Check the screenshot is not black (DRM protection)
- Try preprocessing: grayscale + contrast + threshold
- Try different `--psm` modes
- Check if the app uses custom/non-standard fonts
- For non-Latin scripts, install the appropriate language pack

**OCR misreads numbers:**
- Use `--psm 7` for isolated number fields
- Post-process with regex: `re.sub(r'[Oo]', '0', text)` for common confusions
- Consider a second pass with a different engine for validation

### Screenshot Issues

**Black screenshots (DRM/Secure flag):**
Some apps set `FLAG_SECURE` to block screenshots. Workarounds:
- Root the phone and use a Xposed/Magisk module to disable the flag
- Use `scrcpy` which sometimes bypasses this through video encoding
- Check if the app has a web version as an alternative
- Contact the app developer to request a data export

**Screenshots are slow:**
- Use `adb exec-out screencap -p` (pipes directly, faster than save+pull)
- Reduce screen resolution temporarily: `adb shell wm size 720x1560`
- Restore after: `adb shell wm size reset`
