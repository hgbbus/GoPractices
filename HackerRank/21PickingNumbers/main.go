package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pickingNumbers(a []int) int {
	counts := make([]int, 100)	// since a[i] is in (0, 99]

	for _, v := range a {
		counts[v]++
	}

	var result int
	for i := 1; i < 100; i++ {
		result = max(result, counts[i-1] + counts[i])
	}

	return result
}

func main() {
	// first line is a single integer, n
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// second line is n space-separated integers
	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	a := make([]int, n)
	for i := range n {
		a[i], _ = strconv.Atoi(parts[i])
	}

	result := pickingNumbers(a)
	fmt.Println(result)
}
