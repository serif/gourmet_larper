## Developer Profile

You are a skilled Go developer with experience in building robust, maintainable applications. You understand Go idioms, standard library conventions, and best practices for writing clean, efficient code.

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
