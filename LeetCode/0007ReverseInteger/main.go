package main

import (
	"fmt"
	"math"
	"strconv"
)

// Approach 1: String Manipulation
func reverse(x int) int {
	// Handle negative numbers
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	// Convert to string, reverse, and convert back to int
	s := fmt.Sprintf("%d", x)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedStr := string(runes)

	// Convert back to int and apply sign
	reversedInt, err := strconv.ParseInt(reversedStr, 10, 32)
	if err != nil {
		return 0 // Handle overflow case
	}
	ans := sign * int(reversedInt)
	return ans
}

// Approach 2: Mathematical Reversal
func reverseMath(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	var reversed int
	for x > 0 {
		pop := x % 10
		x /= 10

		// Check for overflow before multiplying and adding
		if reversed > (math.MaxInt32-pop)/10 {
			return 0 // Handle overflow case
		}
		reversed = reversed*10 + pop
	}

	return sign * reversed
}

func main() {
	// test case 1
	x1 := 123
	result1 := reverse(x1)
	fmt.Println(result1) // Output: 321

	// test case 2
	x2 := -123
	result2 := reverse(x2)
	fmt.Println(result2) // Output: -321

	// test case 3
	x3 := 120
	result3 := reverse(x3)
	fmt.Println(result3) // Output: 21

	// test case 4
	x4 := 0
	result4 := reverse(x4)
	fmt.Println(result4) // Output: 0

	// test case 5 (overflow)
	x5 := math.MaxInt32
	result5 := reverse(x5)
	fmt.Println(result5) // Output: 0

	// test case 6 (underflow)
	x6 := math.MinInt32
	result6 := reverse(x6)
	fmt.Println(result6) // Output: 0

	// test case 7 (mathematical approach)
	x7 := 1534236469
	result7 := reverseMath(x7)
	fmt.Println(result7) // Output: 0 (overflow case)
}