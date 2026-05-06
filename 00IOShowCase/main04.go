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

	// Set to split on words.
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter some lines of text (Ctrl+D to end):")

	// Count the words in the input.
	count := 0
	for scanner.Scan() {
		fmt.Printf("Word #%d: %s\n", count, scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Printf("Total words: %d\n", count)
}