package main

import "fmt"

// Approach: Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func maxProfitBruteForce(prices []int) int {
	maxProfit := 0
	for i := range prices {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// Approach: One Pass
// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func main() {
	// test case 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices1)) // Output: 5

	// test case 2
	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices2)) // Output: 0
}
