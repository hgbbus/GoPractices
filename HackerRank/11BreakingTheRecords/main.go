package main

import (
	"fmt"
)

func breakingRecords(scores []int32) []int32 {
	var min, max int32
	var cmin, cmax int32

	// set up min and max from the first game score
	min = scores[0]
	max = scores[0]

	// go over the rest of the slice
	for i := 1; i < len(scores); i++ {
		if scores[i] < min {
			min = scores[i]
			cmin++
		}
		if scores[i] > max {
			max = scores[i]
			cmax++
		}
	}

	return []int32 {cmax, cmin}
}

func main() {
	var n int32
	fmt.Scan(&n)

	// preallocate the entire array (slice)
	scores := make([]int32, n)
	for i := range n {
		fmt.Scan(&scores[i])
	}

	// print the slice
	//for _, v := range scores {
	//	fmt.Println(v)
	//}

	ans := breakingRecords(scores)
	fmt.Println(ans[0], ans[1])
}