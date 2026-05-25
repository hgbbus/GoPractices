package main

import (
	"fmt"
	"slices"
)

func hurdleRace(k int, heights[]int) int {
	maxh := slices.Max(heights)
	return max(0, maxh-k)
}

func main() {
	// first line contains n and k
	var n, k int
	fmt.Scanln(&n, &k)

	// second line contains n numbers (heights)
	heights := make([]int, 10)
	for i := range n {
		fmt.Scan(&heights[i])
	}

	result := hurdleRace(k, heights)
	fmt.Println(result)
}