package main

import (
	"fmt"
	"slices"
)

// Approach 1: Dynamic Programming
// Note: dp[i] represents the length of the longest increasing subsequence 
//       that ends with nums[i]; that is why we must keep track of the maximum 
//       value in dp[i] and not just return dp[n-1]
func lengthOfLIS1(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    dp := make([]int, n)
    ans := 1
    for i := range n {
        dp[i] = 1
        for j := range i {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        if dp[i] > ans {
            ans = dp[i]
        }
    }

    return ans
}

// Approach 2: Dynamic Programming with Binary Search
// Note: tails[i] is the smallest tail of all increasing subsequences 
//       with length i+1 in nums; a common misconception is that tails[i] 
// 		 is the largest tail of all increasing subsequences
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	tails := []int{}
	for _, num := range nums {
		pos, _ := slices.BinarySearch(tails, num)
		if pos == len(tails) {
			tails = append(tails, num)
		} else {
			tails[pos] = num
		}
	}

	return len(tails)
}

// Approach 2a: Dynamic Programming with Binary Search (using DIY binary search)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	tails := make([]int, n)
	size := 0
	for _, num := range nums {
		lo, hi := 0, size
		for lo < hi {
			mid := lo + (hi-lo)/2
			if tails[mid] < num {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		tails[lo] = num
		if lo == size {
			size++
		}
	}
	return size
}

func main() {
	// test case 1
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))

	// test case 2
	nums = []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))

	// test case 3
	nums = []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(lengthOfLIS(nums))
}