package main

import (
	"fmt"
)

func jumpingOnClouds(c []int32, k int32) int32 {
	var n int32 = int32(len(c))

	// solve by simulation
	var e int32 = 100

	i := int32(0)
	for {
		// jump k clouds
		i = (i + k) % n
		e-- // jump cost

		if c[i] == 1 {
			e -= 2 // thundercloud cost
		}

		if i == 0 {
			break
		}
	}

	return e
}

func main() {
	// read n and k
	var n, k int32
	fmt.Scanln(&n, &k)

	// read n numbers
	c := make([]int32, n)
	for i := range n {
		fmt.Scan(&c[i])
	}

	result := jumpingOnClouds(c, k)
	fmt.Println(result)
}
