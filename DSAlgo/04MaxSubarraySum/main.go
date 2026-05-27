package main

import "fmt"

func main() {
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println("Max subarray sum:", maxSubarraySum(arr))
}

// Solution 1: Brute-force approach O(n^3)
func maxSubarraySum1(arr []int) int {
	n := len(arr)
	maxSum := arr[0]

	for i := range n {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}

// Solution 2: Improved brute-force approach O(n^2)
// Avoid recalculating the sum from the scratch 
// for all subarrays starting at the same index i
func maxSubarraySum2(arr []int) int {
	n := len(arr)
	maxSum := arr[0]

	for i := range n {
		sum := 0

		for j := i; j < n; j++ {
			sum += arr[j]
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}

// Solution 3: Divide and conquer approach O(n \log n)
// Split the array into half. The maximum subarray must be one of the following:
//  - entirely in left half
//  - entirely in right half
//  - crossing the midpoint
func maxSubarraySum3(arr []int) int {
	return rec_helper(arr, 0, len(arr)-1)
}

func rec_helper(arr []int, lo int, hi int) int {
	if lo == hi {	// trivial base
		return arr[lo]
	}

	mid := (lo + hi) / 2

	// recursive cases
	left_max := rec_helper(arr, lo, mid)
	right_max := rec_helper(arr, mid+1, hi)

	// crossing mid max sum

	// left first
	sum := 0
	left_sum := -1 << 31
	for i := mid; i >= lo; i-- {
		sum += arr[i]
		left_sum = max(left_sum, sum)
	}

	// right next
	sum = 0
	right_sum := -1 << 31
	for i := mid + 1; i <= hi; i++ {
		sum += arr[i]
		right_sum = max(right_sum, sum)
	}

	crossing_max := left_sum + right_sum

	return max(max(left_max, right_max), crossing_max)
}

// Solution 4: Kadane's algorithm O(n)
// Key idea is that, if the current running sum becomes negative, discard it
// because keeping the negative hurts the running sum
func maxSubarraySum(arr []int) int {
	maxSum := arr[0]
	runningSum := 0

	for _, num := range arr {
		runningSum += num
		maxSum = max(maxSum, runningSum)

		if runningSum < 0 {
			runningSum = 0
		}
	}
	return maxSum
}