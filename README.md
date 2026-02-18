# ende

A simple, Unix-compliant CLI tool for encoding and decoding data streams.

## Features

- **Standard Codecs**: Base64, Base32, Hex (Base16), HTML Entities, URL Encoding.
- **Fun Codecs**: ROT13.
- **Unix Philosophy**: Reads from standard input (pipes) or command-line arguments.
- **Zero Dependencies**: Built with Go standard library (mostly).

## Installation

```bash
go install github.com/hatsat32/ende@latest
```

## Usage

`ende` provides subcommands for each encoding format. Prefix with `en` to encode and `de` to decode.

### Piping (Unix Filter)

```bash
$ echo -n "hello" | ende enb64
aGVsbG8=

$ echo -n "aGVsbG8=" | ende deb64
hello
```

### Arguments

```bash
$ ende enhex "hello"
68656c6c6f
```

### Supported Commands

| Encode | Decode | Description |
|--------|--------|-------------|
| `enb64` | `deb64` | Base64 (RFC 4648) |
| `enb32` | `deb32` | Base32 (RFC 4648) |
| `enhex`, `enb16` | `dehex`, `deb16` | Hexadecimal / Base16 |
| `enhtml` | `dehtml` | HTML Entities |
| `enurl` | `deurl` | URL Percent Encoding |
| `enrot13` | `derot13` | ROT13 Substitution |

## License

MIT
