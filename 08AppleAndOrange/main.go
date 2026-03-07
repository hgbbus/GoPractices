package main

import (
	"fmt"
)

// CountApplesAndOranges counts the number of apples and oranges 
// that fall on Sam's house and prints the results.
//
// Parameters:
// s is the starting point of Sam's house location.
// t is the ending location of Sam's house location.
// a is the location of the Apple tree.
// b is the location of the Orange tree.
// apples is a slice of distances at which each apple falls from the tree.
// oranges is a slice of distances at which each orange falls from the tree.
func countApplesAndOranges(s int32, t int32, a int32, b int32, 
	apples []int32, oranges []int32) {
	var count_apples, count_oranges int

	for _, apple := range apples {
		fall := a + apple
		if fall >= s && fall <= t {
			count_apples++
		}
	}
	fmt.Println(count_apples)

	for _, orange := range oranges {
		fall := b + orange
		if fall >= s && fall <= t {
			count_oranges++
		}
	}
	fmt.Println(count_oranges)
}

func main() {
	var s, t, a, b int32
	var m, n int32

	fmt.Scanf("%d%d\n", &s, &t)
	fmt.Scanf("%d%d\n", &a, &b)
	fmt.Scanf("%d%d\n", &m, &n)

	var apples []int32 = make([]int32, m)
	var oranges []int32 = make([]int32, n)
	for i := range m {
		fmt.Scan(&apples[i])
	}
	for i := range n {
		fmt.Scan(&oranges[i])
	}

	countApplesAndOranges(s, t, a, b, apples, oranges)
}