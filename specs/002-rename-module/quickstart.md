# Quickstart: Module Rename

**Feature**: `002-rename-module`

## Usage

### For Users (Binary)

No change in usage.

```bash
go install github.com/hatsat32/ende/cmd/ende@latest
```

### For Developers (Importing)

Update your `go.mod`:

```go
require github.com/hatsat32/ende v0.0.0
```

Update your imports:

```go
import "github.com/hatsat32/ende/pkg/..."
```

## Verification

Run the tests to ensure the rename didn't break anything:

```bash
go test ./...
```
