package main

import (
	"fmt"
)

func min(a, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
    var ls int32 = int32(len(s))
	var lt int32 = int32(len(t))

	var i int32 = 0

	for i < min(ls, lt) {
		if s[i] != t[i] {
			break
		}
		i++
	}

	var totalOps int32 = (ls - i) + (lt - i)

	// Must be exactly k ops
	if totalOps > k {
		return "No"
	} else if (totalOps%2 == k%2) || (ls+lt <= k) {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	var s, t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)

	var k int32
	fmt.Scanln(&k)

	result := appendAndDelete(s, t, k)
	fmt.Println(result)
}