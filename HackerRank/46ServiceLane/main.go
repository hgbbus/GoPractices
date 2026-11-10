package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'serviceLane' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY cases
 */

func serviceLane(width []int32, cases [][]int32) []int32 {
    // Write your code here
    result := make([]int32, len(cases))
    for i, c := range cases {
        start := c[0]
        end := c[1]
        min := width[start]
        for j := start + 1; j <= end; j++ {
            if width[j] < min {
                min = width[j]
            }
        }
        result[i] = min
    }
    return result
}

func main() {
	// Create a buffered reader to read input from standard input
	reader := bufio.NewReader(os.Stdin)

	// Read the first line containing n and t
	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	parts := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(parts[0])
	t, _ := strconv.Atoi(parts[1])

	// Read the second line containing the width of the service lane
	widthLine, _ := reader.ReadString('\n')
	widthLine = strings.TrimSpace(widthLine)
	widthParts := strings.Split(widthLine, " ")
	width := make([]int32, n)
	for i := 0; i < n; i++ {
		w, _ := strconv.Atoi(widthParts[i])
		width[i] = int32(w)
	}

	// Read the next t lines containing the cases
	cases := make([][]int32, t)
	for i := 0; i < t; i++ {
		caseLine, _ := reader.ReadString('\n')
		caseLine = strings.TrimSpace(caseLine)
		caseParts := strings.Split(caseLine, " ")
		start, _ := strconv.Atoi(caseParts[0])
		end, _ := strconv.Atoi(caseParts[1])
		cases[i] = []int32{int32(start), int32(end)}
	}

	// Call the serviceLane function and get the result
	result := serviceLane(width, cases)

	// Print the result
	for _, res := range result {
		fmt.Println(res)
	}
}