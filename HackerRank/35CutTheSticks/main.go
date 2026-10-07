package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

/*
 * Complete the 'cutTheSticks' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func cutTheSticks(arr []int32) []int32 {
	slices.Sort(arr)

	// No need to do actual cutting, 
	// just count the number of sticks left after each cut

	var result []int32
	n := len(arr)
	for i := 0; i < n; {
		result = append(result, int32(n-i))
		cur := arr[i]
		for i < n && arr[i] == cur {
			i++
		}
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)

	// read first line
	var n int
	fmt.Fscanln(in, &n)

	// read second line
	sticks := make([]int32, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &sticks[i])
	}

	result := cutTheSticks(sticks)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
