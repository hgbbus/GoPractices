package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Set[T comparable] map[T]struct{}

// Method 1: Sorting the array first
func countDistinct1(arr []int) int {
	slices.Sort(arr)
	//fmt.Println(arr)

	var count int = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			count++
		}
	}
	return count
}

// Method 2: Use set
func countDistinct(arr []int) int {
	set := make(Set[int])
	for _, v := range arr {
		set[v] = struct{}{}
	}
	return len(set)
}

func main() {
	// use bufio.Scanner to read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	// By default split func, scanner returns lines

	// First line contains number n
	scanner.Scan() // returns true or false
	n, _ := strconv.Atoi(scanner.Text())

	// Second line contains n numbers
	arr := make([]int, n)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i := range n {
		arr[i], _ = strconv.Atoi(fields[i])
	}

	fmt.Println(countDistinct(arr))
}
