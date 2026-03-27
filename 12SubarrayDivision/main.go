package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int, d int, m int) int {
	if m > len(s) { return 0 }
	
    count := 0
	for i := 0; i < len(s) - m + 1; i++ {
		sum := 0
		for j := range m { sum += s[i+j] }
		if sum == d { count++ }
	}
	return count
}

func birthday2(s []int, d int, m int) int {
	if m > len(s) { return 0 }

    count := 0

	// calculate initial sliding window sum
	sliding_window_sum := 0
	for i := range m { sliding_window_sum += s[i] }
	if sliding_window_sum == d { count++ }

	// check the rest
	for i := m; i < len(s); i++ {
		// update the sliding window sum by removing the leftmost 
		// element and adding the new rightmost element
		sliding_window_sum += s[i] - s[i - m]
		if sliding_window_sum == d { count++ }
	}
	
	return count
}

func main() {
	// this exercise uses bufio.Scanner + strings.Fields
	scanner := bufio.NewScanner(os.Stdin)

	// first line is a single integer n
	scanner.Scan()	// read the line
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// second line contains n integers
	scanner.Scan()	// read the line
	parts := strings.Fields(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(parts[i])
	}

	// third line contains d and m
	scanner.Scan()
	parts2 := strings.Fields(scanner.Text())
	d, _ := strconv.Atoi(parts2[0])
	m, _ := strconv.Atoi(parts2[1])

	fmt.Println(birthday(arr, d, m))
}