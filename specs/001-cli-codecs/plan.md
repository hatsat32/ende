# Implementation Plan: CLI Codec Tools

**Branch**: `001-cli-codecs` | **Date**: 2026-02-18 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `specs/001-cli-codecs/spec.md`

## Summary

Implement a single CLI tool `ende` with subcommands (`enb64`, `deb64`, `enhex`, etc.) for encoding and decoding data streams.
**Technical Approach**: Use Go with `cobra` for the command structure. Stick to the standard library for all encoding/decoding logic.

## Technical Context

**Language/Version**: Go 1.21+ (Standard modern Go)
**Primary Dependencies**:
- `github.com/spf13/cobra` (CLI framework)
- Standard Library (`encoding/base64`, `encoding/hex`, `encoding/base32`, `html`, `net/url`)
**Storage**: N/A (Stateless CLI tool)
**Testing**: Go standard `testing` package (Table-driven tests)
**Target Platform**: Linux/Unix (Cross-platform Go support)
**Project Type**: CLI Tool / Single Module
**Performance Goals**: Instant startup (<50ms), stream processing efficiency (io.Reader/io.Writer)
**Constraints**:
- Low binary size overhead
- Strict adherence to Unix philosophy (stdin/stdout piping)
- No unnecessary 3rd party libs

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **Simplicity First**: Go standard lib is sufficient for all codecs. Logic is straightforward mapping of input -> codec -> output.
- **Explicit Over Implicit**: Cobra commands will explicitly define flags and arguments. Error handling will be explicit.
- **Readability & Quality**: Standard Go formatting (`gofmt`) and idiomatic project structure.
- **Testing Standards**: Table-driven unit tests for all codecs. Integration tests for CLI command execution.
- **User Experience**: Fast startup, standard flags (`--help`), consistent behavior across subcommands.
- **Consistency**: All subcommands follow the same pattern and use shared helper functions for input processing.
- **Performance**: streaming I/O to handle large inputs without loading everything into memory.

## Project Structure

### Documentation (this feature)

```text
specs/001-cli-codecs/
├── plan.md              # This file
├── research.md          # Reference for codec details (if needed)
├── data-model.md        # N/A for this simple CLI
├── quickstart.md        # Usage guide
└── tasks.md             # Implementation tasks
```

### Source Code (repository root)

Standard Go project layout with Cobra:

```text
cmd/
└── ende/
    ├── main.go          # Entry point (initializes root command)
    ├── root.go          # Root command definition
    ├── b64.go           # enb64/deb64 commands
    ├── hex.go           # enhex/dehex commands
    └── ...              # Other codec commands

internal/
├── codecs/          # Shared logic if specific wrappers are needed
│   ├── ...
├── cli/             # Shared CLI helpers (input processing)
    └── input.go     # Helper to handle stdin vs args
```

**Structure Decision**: Single binary `ende` defined in `cmd/ende`. The `main.go` will execute the root command, and individual source files in `cmd/ende` (or `internal/commands` if it grows large, but simpler is better for now) will define the subcommands.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| (None) | Single binary approach simplifies distribution and maintenance. | |
