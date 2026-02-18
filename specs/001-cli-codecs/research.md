# Research: CLI Codec Tools

**Feature**: `001-cli-codecs`
**Status**: Complete

## Technology Decisions

### 1. CLI Framework
- **Decision**: `spf13/cobra`
- **Rationale**: User explicitly requested it. It is the standard for modern Go CLIs (used by Kubernetes, Hugo, GitHub CLI).
- **Alternatives**: `urfave/cli`, standard `flag` package (rejected due to user request).

### 2. Encoding Libraries
- **Decision**: Go Standard Library
- **Rationale**: User preference for standard library. The standard lib covers all requested formats except ROT13 (trivial to implement).
- **Mappings**:
  - `enb64` / `deb64` -> `encoding/base64` (StdEncoding)
  - `enb32` / `deb32` -> `encoding/base32` (StdEncoding)
  - `enb16` / `deb16` -> `encoding/hex`
  - `enhex` / `dehex` -> `encoding/hex` (Alias for base16)
  - `enhtml` / `dehtml` -> `html` (`EscapeString` / `UnescapeString`)
  - `enurl` / `deurl` -> `net/url` (`QueryEscape` / `QueryUnescape`)
  - `enrot13` / `derot13` -> Custom implementation (simple substitution)

### 3. I/O Handling
- **Decision**: Streaming `io.Reader` and `io.Writer`
- **Rationale**: Efficiently handles large inputs (e.g., streaming a file through base64) without buffering the entire content in memory.
- **Pattern**:
  - Check `len(os.Args)`.
  - If args present, wrap `strings.NewReader(arg)` as input.
  - Else, use `os.Stdin` as input.
  - Pass input reader and `os.Stdout` writer to the codec function.

## Open Questions Resolved

- **Q**: How to handle newlines in Base64?
- **A**: Standard `base64` command on Linux wraps at 76 chars by default. Go's `encoding/base64` does not wrap by default.
- **Decision**: We will follow Go's default (no wrapping) for simplicity and machine-readability, unless `base64` compatibility is strictly required (not specified in spec). For decoding, Go's decoder ignores whitespace, so it's compatible.

- **Q**: CLI Structure?
- **A**: Single binary `ende` with subcommands using `cobra`. This simplifies distribution and adheres to the user's latest request, while `cobra` manages the subcommands effectively.
