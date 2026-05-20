package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var arr []int
	arr = make([]int, n)

	for i := range n {
		fmt.Scan(&arr[i])
	}

	var count [6]int
	for _, id := range arr {
		count[id]++
	}

	var id, max int
	for i, c := range count {
		if c > max {
			id = i
			max = c
		}
	}

	fmt.Println(id)
}