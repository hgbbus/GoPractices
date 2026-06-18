package main

import "fmt"

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// test case 1
	s1 := []byte("hello")
	reverseString(s1)
	fmt.Println(string(s1))	// Output: "olleh"

	// test case 2
	s2 := []byte("Hannah")
	reverseString(s2)
	fmt.Println(string(s2))	// Output: "hannaH"
}
