## Developer Profile

You are a skilled Go developer with experience in building robust, maintainable applications. You understand Go idioms, standard library conventions, and best practices for writing clean, efficient code.

## Application Overview

This application scans Chromium-based browser installations for malicious extensions associated with the ShadyPanda malware campaign. It automatically discovers and scans all browser profiles on the system, including the Default profile and any additional profiles (Profile 1, Profile 2, etc.).

### About ShadyPanda Campaign

The ShadyPanda campaign is a sophisticated 7-year malware operation that infected 4.3 million Chrome and Edge browsers. These extensions operated normally for years before being weaponized in mid-2024. The malicious extensions include popular tools like:

- Clean Master: the best Chrome Cache Cleaner
- Speedtest Pro-Free Online Internet Speed Test
- BlockSite
- Various "New Tab" extensions (Infinity, WeTab, SafeSwift, Dream Afar)
- Download Manager Pro
- Wallpaper/HomePage extensions

Once activated, these extensions:
- Execute remote code hourly
- Monitor every website visit
- Exfiltrate browsing history (encrypted)
- Collect complete browser fingerprints

**Reference:** [The Hacker News - ShadyPanda Report](https://thehackernews.com/2025/12/shadypanda-turns-popular-browser.html)

### Supported Browsers

- **Google Chrome**: Full support for all profiles
- **Brave Browser**: Full support for all profiles

The scanner automatically detects which browsers are installed and scans all available profiles across all detected browsers.

### Features

- **Multi-browser support**: Automatically detects Chrome and Brave installations
- **Multi-profile scanning**: Scans all profiles within each detected browser
- **Known malware detection**: Checks against a curated list of malicious extension IDs
- **Comprehensive reporting**: Displays scan results grouped by browser and profile
- **Safe operation**: Read-only scanning with no modifications to browser data
- **Graceful failure handling**: Continues scanning even if individual profiles fail

## Code Style Guidelines

### Import Organization

Organize imports into three distinct groups separated by blank lines:
- Standard library packages
- Third-party dependencies
- Local repository packages

Example:
```go
import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/yourorg/project/internal/services"
)
```

When using formatting tools like `gofmt` or `goimports`, verify that import grouping remains intact after execution.

### Naming Conventions

Use clear, unabbreviated names for aliases, constants, and variables (both package-level and function-scoped).

In Go, capitalized identifiers are exported (publicly accessible), while lowercase identifiers are unexported (package-private). Only export identifiers when external access is required to maintain encapsulation and API clarity.

Acceptable abbreviations:
* Standard shorthand variable names:
  * ctx for context.Context
  * i, j for loop indices
  * t for *testing.T
  * other idiomatic short names used consistently throughout the codebase
* Variables that conflict with package imports:
  * cfg when config package is imported

### Function Design

Function names should be clear and descriptive while remaining concise. Avoid redundancy by excluding the package name from function names. Only include the return type in the function name when multiple variants exist.

### Line of Sight Principle

Structure code to keep the happy path left-aligned and immediately visible. Handle errors and edge cases early with early returns to avoid deep nesting. This approach makes the main logic flow easier to follow and understand at a glance.

Example:
```go
func ProcessData(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }

    parsed, err := parse(data)
    if err != nil {
        return fmt.Errorf("parse failed: %w", err)
    }

    // Happy path continues left-aligned
    return process(parsed)
}
```

### Documentation Standards

Wrap all comments and documentation at 80 characters to improve readability in editors and make diffs easier to review. When editing existing files, only rewrap documentation if you're already modifying that section.

#### Function Parameters

Function parameters follow the same naming rules as other identifiers, with one exception: receiver variables should typically be a single character. All parameters must include explicit types and use camelCase formatting.

### Error Handling

Always wrap errors with descriptive context when propagating them:

```go
if err != nil {
    return fmt.Errorf("failed to process request: %w", err)
}

if err := performOperation(); err != nil {
    return fmt.Errorf("operation failed during execution: %w", err)
}
```

Provide enough context in error messages to understand the failure without requiring additional debugging.

### File Organization

Organize file contents in this order:

* Package imports
* Package-level constants
* Package-level variables
* Type aliases
* Interface definitions
* Struct definitions
* Exported (public) functions
* Unexported (private) functions

This consistent structure makes files easier to navigate and understand.

## Architecture Notes

### Browser Discovery

The application uses the `discoverBrowsers` function to detect installed Chromium-based browsers by checking standard installation paths on macOS:

- **Chrome**: `~/Library/Application Support/Google/Chrome`
- **Brave**: `~/Library/Application Support/BraveSoftware/Brave-Browser`

Each detected browser is represented by a `browserInfo` struct containing:
- `name`: Human-readable browser name (e.g., "Chrome", "Brave")
- `directory`: Full filesystem path to the browser's data directory

### Profile Discovery

The `discoverBrowserProfiles` function scans each browser directory to find all user profiles:
- The "Default" profile directory
- Any directories matching the "Profile *" naming pattern (e.g., "Profile 1", "Profile 2")

This approach ensures comprehensive scanning across all browsers and profiles without requiring manual configuration or user input.

### Scan Result Structure

The `profileScanResult` struct encapsulates scan results for each profile:
- `browserName`: Name of the browser (e.g., "Chrome", "Brave")
- `profileName`: The profile directory name (e.g., "Default", "Profile 1")
- `foundMalicious`: Slice of malicious extension IDs detected in this profile
- `installedCount`: Total number of extensions installed in this profile
- `extensionsPath`: Full filesystem path to the profile's extensions directory

### Error Handling Strategy

The scanner follows a fail-graceful approach at multiple levels:

**Browser Level:**
- If a browser is not installed, it's silently skipped
- Other browsers continue to be scanned normally

**Profile Level:**
- If a profile's extensions directory doesn't exist, it's skipped silently
- If scanning a profile fails, other profiles continue
- Only critical errors (no browsers found, home directory inaccessible) terminate the scan

This multi-level graceful failure ensures that issues with one browser or profile don't prevent scanning others.
