# Data Model: CLI Codecs

**Feature**: `001-cli-codecs`

## Overview

The CLI Codecs operate as stateless filters, transforming an input stream into an output stream. There are no persistent data entities or state management.

## Data Flow

### Input Handling

1.  **Argument Mode** (Priority): If command-line arguments are provided.
    -   **Source**: `os.Args[1:]` (concatenated or processed individually).
    -   **Type**: In-memory string/byte slice.
2.  **Pipe Mode**: If no arguments are provided.
    -   **Source**: `os.Stdin`.
    -   **Type**: `io.Reader` stream.

### Transformation

-   **Process**: Read from Source -> Apply Codec -> Write to Destination.
-   **Codecs**:
    -   Base64 (Standard/URL)
    -   Base32 (Std/Hex)
    -   Hex (Base16)
    -   HTML Entity
    -   URL Encoding
    -   ROT13

### Output Handling

-   **Success**: Write transformed data to `os.Stdout`.
-   **Error**: Write error message to `os.Stderr` and exit with non-zero status code.

## Schema

N/A - No structured data storage.
