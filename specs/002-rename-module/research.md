# Research: Rename Module

**Feature**: `002-rename-module`
**Status**: Completed

## Summary

The goal is to rename the Go module from `ende` to `github.com/hatsat32/ende`. This is a standard operation in the Go ecosystem.

## Decisions

- **Decision**: Use `go mod edit -module` for the `go.mod` file and `find + sed` (or compatible tool) for source code imports.
- **Rationale**: `go mod edit` is the canonical tool for modifying `go.mod`. `sed` is a standard Unix tool available in the environment for text replacement.
- **Alternatives Considered**:
  - *Manual editing*: Error-prone and slow.
  - *IDE Refactoring*: Not available in this CLI environment.

## Execution Steps Verification

1. `go mod edit -module github.com/hatsat32/ende`
2. `find . -type f -name "*.go" -exec sed -i 's|"ende/|"github.com/hatsat32/ende/|g' {} +`
3. `find . -type f -name "*.go" -exec sed -i 's|"ende"|"github.com/hatsat32/ende"|g' {} +` (Careful with this one, might match too broadly if not quoted properly in imports)
   - Better pattern for imports: `s| "ende/| "github.com/hatsat32/ende/|g`
4. `go mod tidy` to clean up.
5. `go test ./...` to verify.

## Open Questions

- **Q**: Are there any other files (README, docs) that mention `ende` as a module?
- **A**: `quickstart.md` or similar might. We should check. `grep -r "ende" .` will help.

## Action Items

- [ ] Execute rename commands.
- [ ] Verify build and tests.
- [ ] Check documentation for stale references.
