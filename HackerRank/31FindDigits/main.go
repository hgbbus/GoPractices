package main

import (
	"fmt"
)

/*
 * Complete the 'findDigits' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func findDigits(n int32) int32 {
    // Write your code here
	var d, result int32
	var num int32 = n

	for num != 0 {
		d = num % 10

		if d != 0 && n % d == 0 {
			result++
		}
		
		num /= 10
	}

	return result
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	// number of queries
	var t int32
	Must(fmt.Scanln(&t))

	for range t {
		var n int32
		Must(fmt.Scanln(&n))
		fmt.Println(findDigits(n))
	}
}