# Quickstart: CLI Codecs

**Feature**: `001-cli-codecs`

## Build

To build the tool, run:

```bash
# Build the single binary
go build -o bin/ende ./cmd/ende
```

## Usage

### Piping (Unix Filter)

```bash
# Encode a string
echo -n "hello" | ./bin/ende enb64
# Output: aGVsbG8=

# Decode a string
echo -n "aGVsbG8=" | ./bin/ende deb64
# Output: hello

# Chain commands
echo "data" | ./bin/ende enb64 | ./bin/ende enhex
```

### Arguments

```bash
# Encode directly
./bin/ende enb64 "hello"
# Output: aGVsbG8=

# Decode directly
./bin/ende deb64 "aGVsbG8="
# Output: hello
```

## Supported Codecs

| Subcommand | Description |
|------------|-------------|
| `enb64`, `deb64` | Base64 Encoding (RFC 4648) |
| `enb32`, `deb32` | Base32 Encoding (RFC 4648) |
| `enhex`, `dehex` | Hexadecimal Encoding |
| `enb16`, `deb16` | Base16 Encoding (Alias for Hex) |
| `enhtml`, `dehtml` | HTML Entity Encoding |
| `enurl`, `deurl` | URL Percent Encoding |
| `enrot13`, `derot13` | ROT13 Substitution |
