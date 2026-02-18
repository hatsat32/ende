# Feature Specification: Rename Module

**Feature Branch**: `002-rename-module`
**Created**: 2026-02-18
**Status**: Draft
**Input**: User request: "please rename project name from 'ende' to 'github.com/hatsat32/ende'"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Maintain Usage Compatibility (Priority: P1)

As a developer, I want to use the tool exactly as before so that my workflows are not disrupted by the internal module rename.

**Why this priority**: Essential to ensure the refactor is transparent to end users.

**Independent Test**: Build and run the tool; behavior should be identical.

**Acceptance Scenarios**:

1. **Given** the source code, **When** I run `go build`, **Then** it produces a working `ende` binary.
2. **Given** the new binary, **When** I run `./ende enb64 "hello"`, **Then** it outputs `aGVsbG8=`.
3. **Given** the codebase, **When** I run `go test ./...`, **Then** all tests pass without import errors.

---

### User Story 2 - Module Compatibility (Priority: P2)

As a Go developer importing this project, I want to import it using `github.com/hatsat32/ende` so that I can manage it as a standard Go dependency.

**Why this priority**: Standard Go ecosystem requirement.

**Independent Test**: `go mod verify` succeeds.

**Acceptance Scenarios**:

1. **Given** the `go.mod` file, **When** inspected, **Then** the module declaration is `module github.com/hatsat32/ende`.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The module path in `go.mod` MUST be updated to `github.com/hatsat32/ende`.
- **FR-002**: All internal import paths within `.go` files MUST be updated to use the new module path `github.com/hatsat32/ende` (replacing `ende`).
- **FR-003**: The build process MUST continue to work without manual intervention.
- **FR-004**: All existing tests MUST pass with the new import paths.

### Key Entities

- **Go Module**: The unit of code distribution and versioning.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: **Accuracy**: 100% of internal imports are updated correctly.
- **SC-002**: **Build**: `go build` succeeds with exit code 0.
- **SC-003**: **Test**: `go test ./...` passes all tests.

## Key Decisions

- This is a refactor-only change; no functional code changes will be made to the CLI logic.
