# 🛡️ Chrome-Based Browser Extension Malware Scanner

> **⚠️ SECURITY ALERT**: Scan your Mac for ShadyPanda malware extensions

A simple tool to scan your Mac's Chrome-based browser extensions for malicious extensions from the ShadyPanda campaign that infected 4.3 million browsers.

**Supported Browsers:** Chrome, Brave | **Platform:** macOS | **Language:** Go | **Scans for:** 27 malicious extensions

## ⚡ Quick Check

**Have Chrome or Brave on your Mac?** Run this scanner to check if you're infected.

**Takes less than 5 seconds** and requires no technical knowledge.

---

## 🔍 About ShadyPanda

The ShadyPanda campaign is a 7-year malware operation that infected 4.3 million Chrome and Edge browsers. These extensions operated normally for years before being weaponized in mid-2024 to:
- Execute remote code hourly
- Monitor every website visit
- Exfiltrate browsing history (encrypted)
- Collect complete browser fingerprints

This scanner checks for all 27 Chrome extension IDs identified in the ShadyPanda campaign.

## Quick Start (For Non-Developers)

(This section is for MacOS. For Windows, check the bottom. For Linux, you know what to do.)

### Step 1: Open Terminal

1. Press `Command + Space` to open Spotlight Search
2. Type "Terminal" and press Enter
3. A black or white window will open - this is the Terminal

### Step 2: Install Homebrew (if not already installed)

Homebrew is a package manager that makes installing software on Mac easier.

In Terminal, paste this command and press Enter:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Follow the on-screen instructions. You may need to enter your Mac password.

### Step 3: Install Git and Go

In Terminal, run these commands one at a time:

```bash
brew install git
brew install go
```

Wait for the installations to complete. This may take a few minutes.

### Step 4: Download This Scanner

In Terminal, run:

```bash
cd ~/Downloads
git clone https://github.com/soniah/gourmet_larper.git
cd gourmet_larper
```

This downloads the scanner to your Downloads folder and moves you into that directory.

### Step 5: Run the Scanner

In Terminal (make sure you're in the gourmet_larper folder), run:

```bash
go run main.go
```

The scanner will automatically detect and check all Chrome and Brave browser profiles, then display the results.

### What You'll See

**If your system is clean:**

```
🔍 Scanning browser extensions for ShadyPanda malware...
=============================================================
Scan Summary:
  • Chrome: 1 profile(s)
  • Brave: 2 profile(s)

📊 Total profiles scanned: 3
📦 Total extensions found: 12
🛡️  Malicious extensions checked: 27

✅ GOOD NEWS: No malicious extensions detected!

All scanned browser profiles appear to be clean from the
ShadyPanda malware campaign extensions.
```

**If malware is detected:**

```
⚠️  ALERT: MALICIOUS EXTENSIONS DETECTED!
=============================================================

🚨 Browser: Chrome | Profile: Default
Found 1 malicious extension(s):

  1. eagiakjmjnblliacokhcalebgnhellfi (Clean Master)
     Path: /Users/yourname/Library/.../Extensions/eagiakjmjnblliacokhcalebgnhellfi

⚡ RECOMMENDED ACTIONS:
  1. Remove these extensions immediately from your browser
  2. Go to chrome://extensions (Chrome) or brave://extensions (Brave)
  3. Enable 'Developer mode' to see extension IDs
  4. Remove any extensions matching the IDs above
  5. Change your passwords across all accounts
  6. Run a full antivirus scan
```

## Alternative: Build and Run

If you want to create a standalone program you can run anytime:

```bash
cd ~/Downloads/gourmet_larper
go build -o chrome-scanner main.go
./chrome-scanner
```

This creates a program called `chrome-scanner` that you can run with `./chrome-scanner`

## What It Does

1. Automatically detects Chrome-based browsers (Chrome and Brave) installed on your Mac

2. Locates all browser profiles for each detected browser:
   - Chrome: `~/Library/Application Support/Google/Chrome/`
   - Brave: `~/Library/Application Support/BraveSoftware/Brave-Browser/`

3. Scans all installed extensions across all profiles

4. Compares extension IDs against the known malicious list

5. Reports any matches with full paths, organized by browser and profile

## If Malware Is Found

If the scanner detects malicious extensions:

1. **Remove immediately** from the affected browser
2. Visit extensions page in your browser:
   - Chrome: `chrome://extensions`
   - Brave: `brave://extensions`
3. Enable "Developer mode" to see extension IDs
4. Remove any extensions matching the detected IDs
5. **Change all your passwords** (the extensions may have stolen credentials)
6. Run a full antivirus scan
7. Consider reviewing your browser history for sensitive sites visited

## Known Malicious Extensions

This scanner checks for 27 known ShadyPanda malicious extension IDs. The campaign includes these popular extensions:

- **Clean Master: the best Chrome Cache Cleaner**
- **Speedtest Pro-Free Online Internet Speed Test**
- **BlockSite**
- **Address bar search engine switcher**
- **SafeSwift New Tab**
- **Infinity V+ New Tab**
- **OneTab Plus:Tab Manage & Productivity**
- **WeTab 新标签页** (WeTab New Tab Page)
- **Infinity New Tab for Mobile**
- **Infinity New Tab (Pro)**
- **Infinity New Tab**
- **Dream Afar New Tab**
- **Download Manager Pro**
- **Galaxy Theme Wallpaper HD 4k HomePage**
- **Halo 4K Wallpaper HD HomePage**

**Note:** This list represents some of the identified malicious extensions, but the campaign includes additional extensions not listed here. The scanner checks for all 27 known malicious extension IDs.

## Troubleshooting

### "go: command not found" or "git: command not found"

Go or Git is not installed. Go back to Step 3 in the Quick Start section and run:

```bash
brew install git
brew install go
```

### "No supported browsers found"

This means neither Chrome nor Brave is installed on your Mac, or they haven't been run yet.

Solution: Open Chrome or Brave at least once to initialize the browser directories, then run the scanner again.

### "Permission denied" when running ./chrome-scanner

The file doesn't have execute permissions. Run:

```bash
chmod +x chrome-scanner
```

Then try running `./chrome-scanner` again.

### Already have Git and Go installed?

Check by running in Terminal:

```bash
git --version
go version
```

If you see version numbers for both, you can skip Step 3 and go straight to Step 4.

## For Developers

**Quick install and run:**

```bash
git clone https://github.com/soniah/gourmet_larper.git
cd gourmet_larper
go run main.go
```

**Or build:**

```bash
go build -o chrome-scanner main.go
./chrome-scanner
```

## For Windows

Right-click the Start button and select Terminal, then run:

```cmd
winget install GoLang.Go
winget install git
cd %USERPROFILE%\Downloads
git clone https://github.com/soniah/gourmet_larper.git
cd gourmet_larper
go run main.go
```

## For Linux

Open a terminal and run:

```bash
# Install Git and Go (Ubuntu/Debian)
sudo apt update
sudo apt install git golang-go

# Or for Fedora/RHEL
sudo dnf install git golang

# Clone and run
cd ~/Downloads
git clone https://github.com/soniah/gourmet_larper.git
cd gourmet_larper
go run main.go
```

## Limitations

- Only supports Chrome and Brave (not Edge or other Chromium browsers)
- Requires browsers to have been run at least once to create profile directories
- Windows and Linux support is experimental; primarily tested on macOS
- Linux: Browser paths may vary by distribution; Snap installations use different paths

## How It Works

The scanner:
1. Detects Chrome and Brave installations on your system (macOS, Windows, or Linux)
2. Discovers all profiles within each browser (Default, Profile 1, Profile 2, etc.)
3. Reads the extensions folder for each profile
4. Gets the unique ID of each installed extension
5. Compares those IDs against a hardcoded list of 27 known malicious extension IDs from the ShadyPanda campaign
6. Reports any matches, organized by browser and profile

No data leaves your computer. This tool only reads local files.

## References

- [The Hacker News - ShadyPanda Turns Popular Browser Extensions Into Spyware](https://thehackernews.com/2025/12/shadypanda-turns-popular-browser.html)
- [Koi Security Report](https://koi.ai/blog/4-million-browsers-infected-inside-shadypanda-7-year-malware-campaign)
- [BleepingComputer Coverage](https://www.bleepingcomputer.com/news/security/shadypanda-browser-extensions-amass-43m-installs-in-malicious-campaign/)
- [Malwarebytes Analysis](https://www.malwarebytes.com/blog/news/2025/12/sleeper-browser-extensions-woke-up-as-spyware-on-4-million-devices)

## FAQ

**Q: Is this safe to run?**
A: Yes. The code is open source and only reads your local Chrome extensions folder. No data is sent anywhere.

**Q: Will this slow down my computer?**
A: No. The scan takes less than 5 seconds and doesn't run in the background.

**Q: What if I use multiple browser profiles?**
A: The scanner automatically detects and scans all profiles in both Chrome and Brave, including Default, Profile 1, Profile 2, etc.

**Q: What about other Chromium browsers like Edge or Opera?**
A: Currently only Chrome and Brave are supported. Support for additional browsers could be added in the future.

**Q: Does this work on Windows or Linux?**
A: Yes! The scanner now supports macOS, Windows, and Linux. See the platform-specific installation instructions above. Note that Windows and Linux support is experimental and primarily tested on macOS.

**Q: The extensions were already removed from Chrome Web Store. Am I safe?**
A: Not necessarily! If you installed them before they were removed, they may still be on your computer. Run this scanner to check.

## License

See the [LICENSE](LICENSE) file for details.
