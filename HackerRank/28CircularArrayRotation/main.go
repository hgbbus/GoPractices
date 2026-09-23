package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func circularArrayRotation(a []int, k int, queries []int) []int {
	result := make([]int, len(queries))

	// calculate the effective number of rotations
	k = k % len(a)

	// for each query, find the corresponding element in the rotated array
	for i, q := range queries {
		// calculate the original index before rotation
		originalIndex := (q - k + len(a)) % len(a)
		result[i] = a[originalIndex]
	}

	return result
}

func main() {
	// since there can be lots of data to read, let's use buffered input
	scanner := bufio.NewScanner(os.Stdin)	// default by scanning line by line
	buf := make([]byte, 0, 1024*1024) 		// initial 1MB buffer
	scanner.Buffer(buf, 16*1024*1024)		// can grow up to 16MB

	// first line contains n, k, q
	var n, k, q int
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d %d", &n, &k, &q)
	}

	// second line contains the array elements
	a := make([]int, n)
	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for i := range n {
			a[i], _ = strconv.Atoi(fields[i])
		}
	}

	// next q lines contain the queries
	queries := make([]int, q)
	for i := range q {
		if scanner.Scan() {
			queries[i], _ = strconv.Atoi(scanner.Text())
		}
	}

	// check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		return
	}

	// get the result
	result := circularArrayRotation(a, k, queries)

	// print the result
	for _, v := range result {
		fmt.Println(v)
	}
}