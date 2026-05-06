# I/O in Go

Executive Summary:
- I/O can be either buffered or unbuffered.
- Package `fmt` implements formatted I/O with functions analogous to C's `printf` and `scanf`.
- Package `bufio` implements buffered I/O.

The `bufio` package provides buffered I/O. It wraps an `io.Reader` or `io.Writer` object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

## Output

### Unbuffered Printing

The `fmt` package provides four families of printing functions defined by their output destinations:

| Function | Destination |
| --- | --- |
| `Print`, `Println`, `Printf` | Standard output (`os.Stdout`) |
| `Fprint`, `Fprintln`, `Fprintf` | Specified `io.Writer` |
| `Sprint`, `Sprintln`, `Sprintf` | Returns formatted string |
| `Append`, `Appendln`, `Appendf` | Appends to a byte slice and returns the updated slice |

The `Print` adds a space between operands if neither is a string. `Println` adds a space between operands and a newline at the end. `Printf` formats according to a format specifier.

Printf format specifiers include:
- `%v` the value in a default format
- `%+v` the value in a default format with field names
- `%#v` a Go-syntax representation of the value
- `%T` a Go-syntax representation of the type of the value
- `%%` a literal percent sign; consumes no value

There are other format specifiers for specific types, such as `%d` for integers, `%s` for strings, etc.

### Buffered Printing

.....

## Input

### Unbuffered Reading/Scanning

Similarly, the `fmt` package provides functions for scanning input from various sources:

| Function | Source |
| --- | --- |
| `Scan`, `Scanln`, `Scanf` | Standard input (`os.Stdin`) |
| `Fscan`, `Fscanln`, `Fscanf` | Specified `io.Reader` |
| `Sscan`, `Sscanln`, `Sscanf` | Scans from a string |

`Scan`, `Fscan`, and `Sscan` read space-separated values and treat a newline in the input as a space.

`Scanln`, `Fscanln`, and `Sscanln` stop scanning at a newline and require that the last scanned item be followed by a newline or EOF.

**Note**: `Scanln` doesn't read an entire line; it only stops at a newline. To read an entire line, must use other functions.

`Scanf`, `Fscanf`, and `Sscanf` read according to a format specifier, similar to `Printf`.

### Buffered Reading

The `bufio` package provides two main types: `bufio.Reader` and `bufio.Scanner`.

#### Scanner

The `bufio.Scanner` type provides a convenient interface for reading data such as a file of newline-delimited lines of text. It is designed to be used in a loop to read successive **tokens** from a file or other input source.

The `Scanner` type provides a `Scan` method that advances the scanner to the next token. Successive calls to the `Scan` method will step through the tokens in the input. The next token will then be available through the `Text` method. The `Scan` method returns `false` when the scan stops, either by reaching the end of the input or an error.

The specification of a "token" is defined by a split function (of a `SplitFunc` type), which can be set using the `Split` method. The default split function is `ScanLines`, which splits the input into lines.

Standard split functions:
| Split Function | Description |
| --- | --- |
| `ScanLines` | Splits on newlines |
| `ScanWords` | Splits on spaces |
| `ScanRunes` | Splits into Unicode code points |
| `ScanBytes` | Splits into bytes |

Scanning stops unrecoverably at EOF, the first I/O error, or a token too large to fit in the `Scanner.Buffer`. When a scan stops, the underline `io.Reader` may have advanced arbitrarily far past the last token.

The `Scanner.Err` method returns any error that occurred during scanning, except that if it returned `false` because it reached the end of the input, `Err` will return `nil`.

#### Reader

The `bufio.Reader` type implements buffering for an `io.Reader` object. It also provides some help for textual I/O, such as the `ReadString` method, which reads until the first occurrence of a specified delimiter and returns a string containing the data up to and including the delimiter.

This type is more flexible than `Scanner` and can be used to read data in various ways, such as reading a specific number of bytes, reading until a certain condition is met, etc.

| Method | Description |
| --- | --- |
| `Read` | Reads data into a byte slice |
| `ReadByte` | Reads and returns a single byte |
| `ReadRune` | Reads and returns a single Unicode code point |
| `ReadString` | Reads until the first occurrence of a specified delimiter and returns a string; often used for reading lines as in `ReadString('\n')` |
| `ReadBytes` | Reads until the first occurrence of a specified delimiter and returns a byte slice; often used for reading lines as in `ReadBytes('\n')` |
| `ReadLine` | (Low-level) Reads a single line, not including the end-of-line bytes; use `ReadBytes('\n')` or `ReadString('\n')` instead |
| `ReadSlice` | Reads until the first occurrence of a specified delimiter and returns a slice of the data |
| `Peek` | Returns the next n bytes without advancing the reader |
| `UnreadByte` | Unreads the last byte returned by `ReadByte` |
| `UnreadRune` | Unreads the last rune returned by `ReadRune` |
| `Reset` | Resets the reader to read from a new `io.Reader` |
| `Buffered` | Returns the number of bytes that can be read from the current buffer |
| `Size` | Returns the size of the buffer |
| `WriteTo` | Writes data to an `io.Writer` until the buffer is empty |

