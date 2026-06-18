package main

import "fmt"

func reverseStr(s string, k int) string {
	// string is immutable and UTF-8 encoded, 
	// so we need to convert it to a slice of runes
	r := []rune(s)
	for i := 0; i < len(r); i += 2 * k {
		hi := min(i+k-1, len(r)-1)
		for lo := i; lo < hi; lo, hi = lo+1, hi-1 {
			r[lo], r[hi] = r[hi], r[lo]
		}
	}
	return string(r)
}

func main() {
	// test case 1
	s1 := "abcdefg"
	k1 := 2
	result1 := reverseStr(s1, k1)
	fmt.Println(result1) // Output: "bacdfeg"

	// test case 2
	s2 := "hannah"
	k2 := 3
	result2 := reverseStr(s2, k2)
	fmt.Println(result2) // Output: "nahnah"
}
