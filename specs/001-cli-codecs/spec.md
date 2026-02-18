# Feature Specification: CLI Codec Tools

**Feature Branch**: `001-cli-codecs`
**Created**: 2026-02-18
**Status**: Draft
**Input**: User description: "I want you to create a cli tool for encoding and decoding. ... follow linux phsophy."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Pipeline Processing (Priority: P1)

As a developer or sysadmin, I want to pipe data from other commands into the encoder/decoder tools so that I can process data streams in a Unix-like pipeline.

**Why this priority**: Core to the "Linux philosophy" requirement and typical usage of such tools.

**Independent Test**: Can be tested using standard shell pipes without any other dependencies.

**Acceptance Scenarios**:

1. **Given** a string "hello world", **When** piped to `ende enb64` (`echo -n "hello world" | ende enb64`), **Then** the output is "aGVsbG8gd29ybGQ=".
2. **Given** a base64 string "aGVsbG8gd29ybGQ=", **When** piped to `ende deb64` (`echo -n "aGVsbG8gd29ybGQ=" | ende deb64`), **Then** the output is "hello world".

---

### User Story 2 - Direct Argument Input (Priority: P2)

As a user, I want to provide the input string directly as a command-line argument so that I can quickly encode/decode short strings without setting up a pipe.

**Why this priority**: Improves usability for quick, ad-hoc tasks.

**Independent Test**: Execute the command with an argument and verify output.

**Acceptance Scenarios**:

1. **Given** the command `ende enhex "foo"`, **When** executed, **Then** the output is "666f6f".
2. **Given** the command `ende dehex "666f6f"`, **When** executed, **Then** the output is "foo".
3. **Given** input provided via BOTH pipe and argument, **When** executed (`echo "ignore" | ende enhex "priority"`), **Then** the argument input "priority" is processed and pipe is ignored/discarded (or handled as per standard convention, but spec requires argument preference).

---

### Edge Cases

- **Empty Input**:
  - Argument provided but empty (`""`): Output should be empty.
  - Stdin empty: Output should be empty.
- **Invalid Input (Decoder)**:
  - `ende deb64` with invalid base64 chars: Should output error message to stderr and exit with non-zero status.
  - `ende dehex` with odd-length string: Should output error message to stderr.
- **Missing Arguments & Empty Stdin**: Program should likely wait for stdin (standard Unix behavior) or exit if non-interactive detection is implemented (but standard simple filter behavior is to wait).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST provide a single executable CLI tool `ende` with the following subcommands:
  - `enb64`, `deb64` (Base64)
  - `enb32`, `deb32` (Base32)
  - `enb16`, `deb16` (Base16)
  - `enhex`, `dehex` (Hexadecimal)
  - `enhtml`, `dehtml` (HTML Entity)
  - `enurl`, `deurl` (URL Encoding)
  - `enrot13`, `derot13` (ROT13)
- **FR-002**: Each subcommand MUST accept input from command-line arguments.
- **FR-003**: Each subcommand MUST read from Standard Input (stdin) if no command-line arguments are provided.
- **FR-004**: If command-line arguments are provided, the tool MUST process the arguments and ignore stdin.
- **FR-005**: Output content MUST be written to Standard Output (stdout).
- **FR-006**: Error messages MUST be written to Standard Error (stderr).
- **FR-007**: `enb64` / `deb64` MUST conform to RFC 4648 standard for Base64.
- **FR-008**: `enb32` / `deb32` MUST conform to RFC 4648 standard for Base32.
- **FR-009**: `enhex` / `dehex` and `enb16` / `deb16` MUST perform standard hexadecimal encoding (case-insensitive decoding, lowercase optional/configurable or standard output).
- **FR-010**: `enhtml` / `dehtml` MUST encode/decode standard HTML5 named entities.
- **FR-011**: `enurl` / `deurl` MUST perform percent-encoding suitable for URL components.
- **FR-012**: `enrot13` / `derot13` MUST perform the standard ROT13 substitution on alphabetic characters only.

### Key Entities

- **Input Stream**: Byte sequence or text to be processed.
- **Encoded/Decoded Output**: Resulting byte sequence or text.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: **Accuracy**: 100% of standard test vectors for each encoding scheme must round-trip correctly (Encode -> Decode -> Original).
- **SC-002**: **Performance**: CLI startup and execution time for small inputs (<1KB) must be under 50ms to ensure "instant" feel.
- **SC-003**: **Usability**: All tools must strictly adhere to the input priority rule (Arguments > Stdin) verified by automated tests.
