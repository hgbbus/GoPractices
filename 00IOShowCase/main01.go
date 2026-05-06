package main

import (
	"fmt"
)

func main() {
	// For Print, space is only added between operands if neither is a string.
	fmt.Print("Hello, World!", 5, 3, "Go", "\n")	// Output: Hello, World!5 3Go
	
	// For Println, space is always added between operands.
	fmt.Println("Hello, World!", 5, 3, "Go")

	// For Printf, the printing is controlled by a sequence of "verbs"
	fmt.Printf("Hello, %s! You are %d years old.\n", "Alice", 30)
}