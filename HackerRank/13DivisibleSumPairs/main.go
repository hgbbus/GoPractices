package main

import (
	"fmt"
)

func main() {
	// Read n and k from standard input
	var n, k int
	fmt.Scan(&n, &k)
	//fmt.Println(n, k)

	// Read n integers
	var arr []int = make([]int, n)
	for i := range n {
		fmt.Scan(&arr[i])	// ignore error for simplicity
	}
	//fmt.Println(arr)

	var count int
	for i := range n {
		for j := i + 1; j < n; j++ {
			if (arr[i] + arr[j])%k == 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
