package main

import (
	"fmt"
)

func utopianTree(n int) int {
	height := 1

	for i := range n {
		if i % 2 == 0 {
			height *= 2
		} else {
			height++
		}
	}

	return height
}

func main() {
	var t int
	fmt.Scanln(&t)

	for range t {
		var n int
		fmt.Scanln(&n)
		fmt.Println(utopianTree(n))
	}
}

