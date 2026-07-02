package main

import (
	"fmt"
)

// Function to generate all subsets of a given set
func generateSubsets(nums []int) [][]int {
	var result [][] int
	var scratch []int

	backtrack(nums, 0, scratch, &result)

	return result
}

func backtrack(nums []int, k int, scratch []int, result *[][]int) {
	if (k == len(nums)) {
		// Make a copy of the current subset and add it to the result
		subset := make([]int, len(scratch))
		copy(subset, scratch)
		*result = append(*result, subset)
		return
	}

	// Include the current element
	scratch = append(scratch, nums[k])
	backtrack(nums, k+1, scratch, result)

	// Exclude the current element
	scratch = scratch[:len(scratch)-1]
	backtrack(nums, k+1, scratch, result)
}

func main() {
	nums := []int{1, 2, 3}
	subsets := generateSubsets(nums)
	fmt.Println("All subsets:")
	for _, subset := range subsets {
		fmt.Println(subset)
	}
}
