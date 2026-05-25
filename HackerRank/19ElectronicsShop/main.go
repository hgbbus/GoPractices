package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Brute force approach - Try all combos - O(n*m)
func getMoneySpent1(keyboards []int, drives []int, b int) int {
	max := -1
	for i := range keyboards {
		for j := range drives {
			sum := keyboards[i] + drives[j]
			if sum <= b && sum > max {
				max = sum
			}
		}
	}
	return max
}

// More efficient approach - Sort the drives and use binary search - O(n log m)
func getMoneySpent2(keyboards []int, drives []int, b int) int {
	max := -1
	// sort the drives in ascending order
	sort.Ints(drives)

	for i := range keyboards {
		remaining := b - keyboards[i]
		if remaining < 0 {
			continue
		}
		// binary search for the largest drive that is <= remaining
		j := sort.Search(len(drives), func(j int) bool {
			return drives[j] > remaining
		})
		if j > 0 {
			sum := keyboards[i] + drives[j-1]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// Another approach - Sort both arrays and use two pointers - O(n log n + m log m)
func getMoneySpent(keyboards []int, drives []int, b int) int {
	max := -1
	// sort both arrays in ascending order
	sort.Ints(keyboards)
	sort.Ints(drives)

	// use two pointers to find the best combo
	i, j := 0, len(drives)-1
	for i < len(keyboards) && j >= 0 {
		sum := keyboards[i] + drives[j]
		if sum > b {
			j-- // try a cheaper drive
		} else {
			if sum > max {
				max = sum
			}
			i++ // try a more expensive keyboard
		}
	}
	return max
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// first line contains budget, num of keyboards, num of USB drives
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	b, _ := strconv.Atoi(fields[0])
	n, _ := strconv.Atoi(fields[1])
	m, _ := strconv.Atoi(fields[2])

	// second line is for the keyboards
	keyboards := make([]int, n)
	line, _ = reader.ReadString('\n')
	fields = strings.Fields(line)
	for i := range n {
		keyboards[i], _ = strconv.Atoi(fields[i])
	}

	// third line is for the drives
	drives := make([]int, m)
	line, _ = reader.ReadString('\n')
	fields = strings.Fields(line)
	for i := range m {
		drives[i], _ = strconv.Atoi(fields[i])
	}

	// call the function and print the result
	result := getMoneySpent(keyboards, drives, b)
	fmt.Println(result)
}