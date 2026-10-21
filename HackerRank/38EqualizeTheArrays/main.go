package main

import (
	"fmt"
	"slices"
)

/*
 * Complete the 'equalizeArray' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func equalizeArray(arr []int) int {
	n := len(arr) // n >= 1

	slices.Sort(arr)

	// count the number of occurrences of each element
	count := 1
	maxCount := 1
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 1
		}
	}
	if count > maxCount {
		maxCount = count
	}

	return n - maxCount
}

/* func equalizeArray(arr []int) int {
	n := len(arr) // n >= 1
	count := make(map[int]int)
	for _, c := range arr {
		count[c]++
	}
	m := 0
	for _, c := range count {
		if c > m {
			m = c
		}
	}
	return n - m
} */

/*
	 func equalizeArray(arr []int) int {
	    // Write your code here
		n := len(arr)	// n >= 1
		count := make([]int, 101)	// used as counting array

		// count
		for _, a := range arr {
			count[a]++;
		}

		// find an element that occurs the most
		//e := 0
		m := 0	// current max
		for _, c := range count {
			if c > m {
				//e = _
				m = c
			}
		}

		return n - m
	}
*/
func main() {
	// first line: number of elements in the array
	var n int
	fmt.Scanln(&n)

	// second line: the elements of the array
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	result := equalizeArray(arr)
	fmt.Println(result)
}
