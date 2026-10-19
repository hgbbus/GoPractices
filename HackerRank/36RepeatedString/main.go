package main

import "fmt"

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	l := int64(len(s))

	count := int64(0)
	for _, c := range s {
		if c == 'a' {
			count++
		}
	}
	
	fullRepeats := n / l
	remainder := n % l

	count = count * fullRepeats
	for i := range remainder {
		if s[i] == 'a' {
			count++
		}
	}

	return count
}

func main() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var n int64
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	count := repeatedString(s, n)
	fmt.Println(count)
}
