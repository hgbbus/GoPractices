package main

import (
	"fmt"
)

func viralAdvertising(n int) int {
	shared := 5
	liked := 0
	cumulative := 0

	for range n {
		liked = shared / 2
		cumulative += liked
		shared = liked * 3
	}

	return cumulative
}

func main() {
	var n int
	fmt.Scan(&n)
	result := viralAdvertising(n)
	fmt.Println(result)
}