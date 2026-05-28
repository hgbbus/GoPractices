package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// use bufio.Scanner to read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	// By default split func, scanner returns lines

	// First line contains number n
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "Failed to read number of customers")
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	// Read next n lines, each line contains two integers (arrival time and leaving time)
	arrivals := make([]int, n)
	leavings := make([]int, n)
	for i := range n {
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Failed to read arrival and leaving times")
			return
		}
		fields := strings.Fields(scanner.Text())
		arrivals[i], _ = strconv.Atoi(fields[0])
		leavings[i], _ = strconv.Atoi(fields[1])
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	// Sort arrival and leaving times using slices.Sort
	slices.Sort(arrivals)
	slices.Sort(leavings)

	// Use two pointers to count the number of customers in the restaurant at any time
	var maxCustomers int
	var currentCustomers int
	var i, j int
	for i < n && j < n {
		if arrivals[i] < leavings[j] {
			currentCustomers++
			maxCustomers = max(maxCustomers, currentCustomers)
			i++
		} else {
			currentCustomers--
			j++
		}
	}

	fmt.Println(maxCustomers)
}