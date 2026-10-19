package main

import "fmt"

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int) int {
	n := len(c)
	ans := 0
	i := 0
	for i < n-1 {
		i++
		if i < n-1 && c[i+1] == 0 {
			i++
		}
		ans++
	}
	return ans
}

func main() {
	// read first line for number of clouds
	var n int
	fmt.Scanln(&n)

	// read second line for cloud values
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}

	result := jumpingOnClouds(c)
	fmt.Println(result)
}
