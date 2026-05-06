package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//
	// bufio package
	//

	// The Scan function advances the Scanner to the next token, which will then be available
	// through either the Bytes or Text method. It returns a boolean.

	// When testing, use Ctrl+D (Unix) or Ctrl+Z (Windows) to signal the end of input.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // Default split function is ScanLines, which splits on newlines.
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}