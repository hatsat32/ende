# Tasks: CLI Codecs

**Feature**: `001-cli-codecs`
**Status**: Planning

## Phase 1: Setup
<!-- Goal: Initialize project structure and dependencies -->

- [x] T001 Initialize Go module `ende` in root
- [x] T002 Initialize Cobra application structure
- [x] T003 Create `cmd/ende` directory structure

## Phase 2: Foundational
<!-- Goal: Core input processing logic for stdin vs arguments -->

- [x] T004 Implement input helper `GetInput` in `internal/cli/input.go`
- [x] T005 [P] Create `internal/cli/input_test.go` to verify argument/stdin priority (SC-003)
- [x] T006 Implement root command in `cmd/ende/root.go`

## Phase 3: Pipeline Processing (User Story 1 - P1)
<!-- Goal: Enable data processing via pipes for all codecs -->

- [x] T007 [US1] Implement `enb64` and `deb64` commands in `cmd/ende/b64.go`
- [x] T008 [P] [US1] Create unit tests for Base64 in `cmd/ende/b64_test.go`
- [x] T009 [US1] Implement `enb32` and `deb32` commands in `cmd/ende/b32.go`
- [x] T010 [P] [US1] Create unit tests for Base32 in `cmd/ende/b32_test.go`
- [x] T011 [US1] Implement `enhex`, `dehex`, `enb16`, `deb16` commands in `cmd/ende/hex.go`
- [x] T012 [P] [US1] Create unit tests for Hex/Base16 in `cmd/ende/hex_test.go`
- [x] T013 [US1] Implement `enhtml`, `dehtml`, `enurl`, `deurl` commands in `cmd/ende/web.go`
- [x] T014 [P] [US1] Create unit tests for HTML/URL in `cmd/ende/web_test.go`
- [x] T015 [US1] Implement `enrot13`, `derot13` commands in `cmd/ende/rot13.go`
- [x] T016 [P] [US1] Create unit tests for ROT13 in `cmd/ende/rot13_test.go`
- [x] T017 [US1] Verify pipeline processing functionality manually for all commands

## Phase 4: Direct Argument Input (User Story 2 - P2)
<!-- Goal: Ensure all commands handle direct argument input correctly -->

- [x] T018 [US2] Verify argument vs stdin priority for `enb64`/`deb64` (SC-003)
- [x] T019 [US2] Verify argument handling for all other codecs
- [x] T020 [US2] Document direct argument usage in `README.md` or help output

## Phase 5: Polish & Edge Cases
<!-- Goal: Error handling, final documentation, and binary size optimization -->

- [x] T021 Implement error handling for invalid inputs (FR-006)
- [x] T022 Verify behavior with empty inputs (Edge Case)
- [x] T023 Finalize help messages and examples for all subcommands
- [x] T024 Check binary size and optimization (Constraints)

## Dependencies

- T006 (Root Cmd) depends on T001, T002
- T007-T016 (Codecs) depend on T004 (Input Helper)
- T018-T019 (Arg Verification) depend on T007-T016 (Codecs Implementation)

## Parallel Execution

- T007 (Base64), T009 (Base32), T011 (Hex), T013 (Web), T015 (ROT13) can be implemented in parallel after T004 is done.
- Unit tests (T008, T010, T012, T014, T016) can be written in parallel with their respective implementations.

## Independent Test Criteria

- **US1 (Pipeline)**: `echo "data" | ende enb64` produces expected output.
- **US2 (Arguments)**: `ende enb64 "data"` produces expected output and ignores any piped input.
