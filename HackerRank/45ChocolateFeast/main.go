package main

import (
	"fmt"
)

/*
 * Complete the 'chocolateFeast' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c
 *  3. INTEGER m
 */

func chocolateFeast(n int32, c int32, m int32) int32 {
    if m <= 1 {
		panic("m must be greater than 1")
	}
	
	var total, wrappers int32

	candies := n/c
	for candies > 0 {
		// candies to wrappers
		total += candies
		wrappers += candies
		candies = wrappers / m
		wrappers = wrappers % m
	}

	return total
}

func main() {
	// first line t
	var t int
	fmt.Scanln(&t)

	for range t {
		var n, c, m int
		fmt.Scanln(&n, &c, &m)
		fmt.Println(chocolateFeast(int32(n), int32(c), int32(m)))
	}
}