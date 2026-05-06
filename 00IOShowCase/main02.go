package main

import (
	"fmt"
)

func main() {
	// For Scan, the input is space-separated by default.
	// Multiple values can be scanned in one call.
	// Types of the scanned variables are the types of the corresponding arguments.
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	// For Scanln, the input is also space-separated, but it stops scanning at the first newline.
	var city string
	fmt.Print("Enter your city: ")
	fmt.Scanln(&city)
	fmt.Printf("You live in %s.\n", city)

	// For Scanf, the input is formatted according to a format string.
	// This is very much like C's scanf.
	var country string
	fmt.Print("Enter your country: ")
	fmt.Scanf("%s", &country)
	fmt.Printf("You are from %s.\n", country)
}