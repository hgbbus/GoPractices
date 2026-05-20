package main

import (
	"fmt"
)

func birthdayCakeCandles(candles []int32) int32 {
	var tallest int32 = candles[0]
	var count int32 = 1

	for i := 1; i < len(candles); i++ {
		if candles[i] == tallest {
			count++
		} else if candles[i] > tallest {
			tallest = candles[i]
			count = 1
		}
	}

	return count
}

func main() {
	var n int32
	fmt.Scan(&n)
	
	var candles []int32 = make([]int32, n)
	for i := range n {
		fmt.Scan(&candles[i])
	}
	//fmt.Println(candles)

	fmt.Println(birthdayCakeCandles(candles))
}