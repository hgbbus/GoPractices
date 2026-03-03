package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diagonalDifference(arr [][]int32) int32 {
	var primaryDiagonalSum int32
	var secondaryDiagonalSum int32

	for i := range arr {// index only; value ignored
		primaryDiagonalSum += arr[i][i]
		secondaryDiagonalSum += arr[i][len(arr)-1-i]
	}

	if primaryDiagonalSum > secondaryDiagonalSum {
		return primaryDiagonalSum - secondaryDiagonalSum
	}
	return secondaryDiagonalSum - primaryDiagonalSum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line of input and convert it to an integer
	scanner.Scan()
	line := scanner.Text()
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// Read the next n lines of input and store them in a 2D slice
	arr := make([][]int32, n)
	for i := range n {
		scanner.Scan()
		line := scanner.Text()
		arr[i] = make([]int32, n)
		for j, w := range strings.Fields(line) {
			e, _ := strconv.Atoi(w)
			arr[i][j] = int32(e)
		}
	}

	// Call the diagonalDifference function and print the result
	result := diagonalDifference(arr)
	fmt.Println(result)
}