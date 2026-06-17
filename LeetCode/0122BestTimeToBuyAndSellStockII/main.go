package main

import "fmt"

// Approach 1: Greedy
// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxProfit += prices[i] - minPrice
			minPrice = prices[i]
		}
	}
	return maxProfit
}

// Approach 2: Optimal greedy
// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfitOptimal(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func main() {
	// test case 1
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // Output: 7
	fmt.Println(maxProfitOptimal(prices)) // Output: 7

	// test case 2
	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices)) // Output: 4
	fmt.Println(maxProfitOptimal(prices)) // Output: 4

	// test case 3
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices)) // Output: 0
	fmt.Println(maxProfitOptimal(prices)) // Output: 0
}
