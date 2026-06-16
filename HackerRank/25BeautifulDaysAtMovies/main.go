package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func reverse(n int) int {
	var ans int

	for n > 0 {
		ans = 10 * ans + n % 10
		n = n / 10
	}

	return ans
}

func beautifulDays(i int, j int, k int) int {
	var ans int

	for x := i; x <= j; x++ {
		if abs(x - reverse(x)) % k == 0 {
			ans++
		}
	}

	return ans
}

func main() {
	var i, j, k int
	fmt.Scan(&i, &j, &k)
	result := beautifulDays(i, j, k)
	fmt.Println(result)
}
