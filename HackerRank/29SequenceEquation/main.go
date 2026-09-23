package main

import (
	"fmt"
)

func permutationEquation(p []int32) []int32 {
    // Write your code here
	n := len(p)
	r := make([]int32, n)

	// p is a permutation.
	//
	// View p as a function, for each x, we need to find y such that 
	// p(p(y)) = x. Mathematically, this is to start from y, apply 
	// function p twice to reach x.
	//
	//      p(y)      p(z)
	//   y ------> z ------> x
	//
	// Thus, to find y while given x, it is going backward twice.
	// For this reason, we need the reverse function of p.
	//

	q := make([]int32, n)
	for i, v := range p {	// i is zero-based; v is one-based
		q[v-1] = int32(i + 1)
	}

	for i := range r {
		r[i] = q[q[i]-1]
	}

	return r
}

func main() {
	// Only at most 50 numbers, we can use unbuffered input

	// first line contains an integer n
	var n int32
	fmt.Scanln(&n)

	// second line contains n space-separated integers p[i]
	p := make([]int32, n)
	for i := range n {
		fmt.Scan(&p[i])
	}

	result := permutationEquation(p)
	for _, r := range result {
		fmt.Println(r)
	}
}
